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

package iamidentityv1_test

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
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`IamIdentityV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(iamIdentityService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(iamIdentityService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
				URL: "https://iamidentityv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(iamIdentityService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_IDENTITY_URL": "https://iamidentityv1/api",
				"IAM_IDENTITY_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{
				})
				Expect(iamIdentityService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := iamIdentityService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamIdentityService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamIdentityService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamIdentityService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{
					URL: "https://testService/api",
				})
				Expect(iamIdentityService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := iamIdentityService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamIdentityService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamIdentityService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamIdentityService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{
				})
				err := iamIdentityService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := iamIdentityService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamIdentityService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamIdentityService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamIdentityService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_IDENTITY_URL": "https://iamidentityv1/api",
				"IAM_IDENTITY_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamIdentityService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_IDENTITY_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamIdentityService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = iamidentityv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListAPIKeys(listAPIKeysOptions *ListAPIKeysOptions) - Operation response error`, func() {
		listAPIKeysPath := "/v1/apikeys"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAPIKeysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["pagesize"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["scope"]).To(Equal([]string{"entity"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"user"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAPIKeys with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListAPIKeysOptions model
				listAPIKeysOptionsModel := new(iamidentityv1.ListAPIKeysOptions)
				listAPIKeysOptionsModel.AccountID = core.StringPtr("testString")
				listAPIKeysOptionsModel.IamID = core.StringPtr("testString")
				listAPIKeysOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listAPIKeysOptionsModel.Pagetoken = core.StringPtr("testString")
				listAPIKeysOptionsModel.Scope = core.StringPtr("entity")
				listAPIKeysOptionsModel.Type = core.StringPtr("user")
				listAPIKeysOptionsModel.Sort = core.StringPtr("testString")
				listAPIKeysOptionsModel.Order = core.StringPtr("asc")
				listAPIKeysOptionsModel.IncludeHistory = core.BoolPtr(false)
				listAPIKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListAPIKeys(listAPIKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListAPIKeys(listAPIKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAPIKeys(listAPIKeysOptions *ListAPIKeysOptions)`, func() {
		listAPIKeysPath := "/v1/apikeys"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAPIKeysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["pagesize"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["scope"]).To(Equal([]string{"entity"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"user"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "apikeys": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}]}`)
				}))
			})
			It(`Invoke ListAPIKeys successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListAPIKeysOptions model
				listAPIKeysOptionsModel := new(iamidentityv1.ListAPIKeysOptions)
				listAPIKeysOptionsModel.AccountID = core.StringPtr("testString")
				listAPIKeysOptionsModel.IamID = core.StringPtr("testString")
				listAPIKeysOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listAPIKeysOptionsModel.Pagetoken = core.StringPtr("testString")
				listAPIKeysOptionsModel.Scope = core.StringPtr("entity")
				listAPIKeysOptionsModel.Type = core.StringPtr("user")
				listAPIKeysOptionsModel.Sort = core.StringPtr("testString")
				listAPIKeysOptionsModel.Order = core.StringPtr("asc")
				listAPIKeysOptionsModel.IncludeHistory = core.BoolPtr(false)
				listAPIKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListAPIKeysWithContext(ctx, listAPIKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListAPIKeys(listAPIKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListAPIKeysWithContext(ctx, listAPIKeysOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAPIKeysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["pagesize"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["scope"]).To(Equal([]string{"entity"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"user"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "apikeys": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}]}`)
				}))
			})
			It(`Invoke ListAPIKeys successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListAPIKeys(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAPIKeysOptions model
				listAPIKeysOptionsModel := new(iamidentityv1.ListAPIKeysOptions)
				listAPIKeysOptionsModel.AccountID = core.StringPtr("testString")
				listAPIKeysOptionsModel.IamID = core.StringPtr("testString")
				listAPIKeysOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listAPIKeysOptionsModel.Pagetoken = core.StringPtr("testString")
				listAPIKeysOptionsModel.Scope = core.StringPtr("entity")
				listAPIKeysOptionsModel.Type = core.StringPtr("user")
				listAPIKeysOptionsModel.Sort = core.StringPtr("testString")
				listAPIKeysOptionsModel.Order = core.StringPtr("asc")
				listAPIKeysOptionsModel.IncludeHistory = core.BoolPtr(false)
				listAPIKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListAPIKeys(listAPIKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAPIKeys with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListAPIKeysOptions model
				listAPIKeysOptionsModel := new(iamidentityv1.ListAPIKeysOptions)
				listAPIKeysOptionsModel.AccountID = core.StringPtr("testString")
				listAPIKeysOptionsModel.IamID = core.StringPtr("testString")
				listAPIKeysOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listAPIKeysOptionsModel.Pagetoken = core.StringPtr("testString")
				listAPIKeysOptionsModel.Scope = core.StringPtr("entity")
				listAPIKeysOptionsModel.Type = core.StringPtr("user")
				listAPIKeysOptionsModel.Sort = core.StringPtr("testString")
				listAPIKeysOptionsModel.Order = core.StringPtr("asc")
				listAPIKeysOptionsModel.IncludeHistory = core.BoolPtr(false)
				listAPIKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListAPIKeys(listAPIKeysOptionsModel)
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
			It(`Invoke ListAPIKeys successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListAPIKeysOptions model
				listAPIKeysOptionsModel := new(iamidentityv1.ListAPIKeysOptions)
				listAPIKeysOptionsModel.AccountID = core.StringPtr("testString")
				listAPIKeysOptionsModel.IamID = core.StringPtr("testString")
				listAPIKeysOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listAPIKeysOptionsModel.Pagetoken = core.StringPtr("testString")
				listAPIKeysOptionsModel.Scope = core.StringPtr("entity")
				listAPIKeysOptionsModel.Type = core.StringPtr("user")
				listAPIKeysOptionsModel.Sort = core.StringPtr("testString")
				listAPIKeysOptionsModel.Order = core.StringPtr("asc")
				listAPIKeysOptionsModel.IncludeHistory = core.BoolPtr(false)
				listAPIKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListAPIKeys(listAPIKeysOptionsModel)
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
	Describe(`CreateAPIKey(createAPIKeyOptions *CreateAPIKeyOptions) - Operation response error`, func() {
		createAPIKeyPath := "/v1/apikeys"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAPIKeyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "false")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAPIKey with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsModel := new(iamidentityv1.CreateAPIKeyOptions)
				createAPIKeyOptionsModel.Name = core.StringPtr("testString")
				createAPIKeyOptionsModel.IamID = core.StringPtr("testString")
				createAPIKeyOptionsModel.Description = core.StringPtr("testString")
				createAPIKeyOptionsModel.AccountID = core.StringPtr("testString")
				createAPIKeyOptionsModel.Apikey = core.StringPtr("testString")
				createAPIKeyOptionsModel.StoreValue = core.BoolPtr(true)
				createAPIKeyOptionsModel.EntityLock = core.StringPtr("false")
				createAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateAPIKey(createAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateAPIKey(createAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAPIKey(createAPIKeyOptions *CreateAPIKeyOptions)`, func() {
		createAPIKeyPath := "/v1/apikeys"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAPIKeyPath))
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

					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "false")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke CreateAPIKey successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsModel := new(iamidentityv1.CreateAPIKeyOptions)
				createAPIKeyOptionsModel.Name = core.StringPtr("testString")
				createAPIKeyOptionsModel.IamID = core.StringPtr("testString")
				createAPIKeyOptionsModel.Description = core.StringPtr("testString")
				createAPIKeyOptionsModel.AccountID = core.StringPtr("testString")
				createAPIKeyOptionsModel.Apikey = core.StringPtr("testString")
				createAPIKeyOptionsModel.StoreValue = core.BoolPtr(true)
				createAPIKeyOptionsModel.EntityLock = core.StringPtr("false")
				createAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateAPIKeyWithContext(ctx, createAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateAPIKey(createAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateAPIKeyWithContext(ctx, createAPIKeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAPIKeyPath))
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

					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "false")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke CreateAPIKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateAPIKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsModel := new(iamidentityv1.CreateAPIKeyOptions)
				createAPIKeyOptionsModel.Name = core.StringPtr("testString")
				createAPIKeyOptionsModel.IamID = core.StringPtr("testString")
				createAPIKeyOptionsModel.Description = core.StringPtr("testString")
				createAPIKeyOptionsModel.AccountID = core.StringPtr("testString")
				createAPIKeyOptionsModel.Apikey = core.StringPtr("testString")
				createAPIKeyOptionsModel.StoreValue = core.BoolPtr(true)
				createAPIKeyOptionsModel.EntityLock = core.StringPtr("false")
				createAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateAPIKey(createAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAPIKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsModel := new(iamidentityv1.CreateAPIKeyOptions)
				createAPIKeyOptionsModel.Name = core.StringPtr("testString")
				createAPIKeyOptionsModel.IamID = core.StringPtr("testString")
				createAPIKeyOptionsModel.Description = core.StringPtr("testString")
				createAPIKeyOptionsModel.AccountID = core.StringPtr("testString")
				createAPIKeyOptionsModel.Apikey = core.StringPtr("testString")
				createAPIKeyOptionsModel.StoreValue = core.BoolPtr(true)
				createAPIKeyOptionsModel.EntityLock = core.StringPtr("false")
				createAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateAPIKey(createAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAPIKeyOptions model with no property values
				createAPIKeyOptionsModelNew := new(iamidentityv1.CreateAPIKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateAPIKey(createAPIKeyOptionsModelNew)
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
			It(`Invoke CreateAPIKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsModel := new(iamidentityv1.CreateAPIKeyOptions)
				createAPIKeyOptionsModel.Name = core.StringPtr("testString")
				createAPIKeyOptionsModel.IamID = core.StringPtr("testString")
				createAPIKeyOptionsModel.Description = core.StringPtr("testString")
				createAPIKeyOptionsModel.AccountID = core.StringPtr("testString")
				createAPIKeyOptionsModel.Apikey = core.StringPtr("testString")
				createAPIKeyOptionsModel.StoreValue = core.BoolPtr(true)
				createAPIKeyOptionsModel.EntityLock = core.StringPtr("false")
				createAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateAPIKey(createAPIKeyOptionsModel)
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
	Describe(`GetAPIKeysDetails(getAPIKeysDetailsOptions *GetAPIKeysDetailsOptions) - Operation response error`, func() {
		getAPIKeysDetailsPath := "/v1/apikeys/details"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAPIKeysDetailsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Iam-Apikey"]).ToNot(BeNil())
					Expect(req.Header["Iam-Apikey"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAPIKeysDetails with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAPIKeysDetailsOptions model
				getAPIKeysDetailsOptionsModel := new(iamidentityv1.GetAPIKeysDetailsOptions)
				getAPIKeysDetailsOptionsModel.IamAPIKey = core.StringPtr("testString")
				getAPIKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAPIKeysDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAPIKeysDetails(getAPIKeysDetailsOptions *GetAPIKeysDetailsOptions)`, func() {
		getAPIKeysDetailsPath := "/v1/apikeys/details"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAPIKeysDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Iam-Apikey"]).ToNot(BeNil())
					Expect(req.Header["Iam-Apikey"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke GetAPIKeysDetails successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetAPIKeysDetailsOptions model
				getAPIKeysDetailsOptionsModel := new(iamidentityv1.GetAPIKeysDetailsOptions)
				getAPIKeysDetailsOptionsModel.IamAPIKey = core.StringPtr("testString")
				getAPIKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAPIKeysDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetAPIKeysDetailsWithContext(ctx, getAPIKeysDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetAPIKeysDetailsWithContext(ctx, getAPIKeysDetailsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAPIKeysDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Iam-Apikey"]).ToNot(BeNil())
					Expect(req.Header["Iam-Apikey"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke GetAPIKeysDetails successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetAPIKeysDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAPIKeysDetailsOptions model
				getAPIKeysDetailsOptionsModel := new(iamidentityv1.GetAPIKeysDetailsOptions)
				getAPIKeysDetailsOptionsModel.IamAPIKey = core.StringPtr("testString")
				getAPIKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAPIKeysDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAPIKeysDetails with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAPIKeysDetailsOptions model
				getAPIKeysDetailsOptionsModel := new(iamidentityv1.GetAPIKeysDetailsOptions)
				getAPIKeysDetailsOptionsModel.IamAPIKey = core.StringPtr("testString")
				getAPIKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAPIKeysDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptionsModel)
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
			It(`Invoke GetAPIKeysDetails successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAPIKeysDetailsOptions model
				getAPIKeysDetailsOptionsModel := new(iamidentityv1.GetAPIKeysDetailsOptions)
				getAPIKeysDetailsOptionsModel.IamAPIKey = core.StringPtr("testString")
				getAPIKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAPIKeysDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptionsModel)
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
	Describe(`GetAPIKey(getAPIKeyOptions *GetAPIKeyOptions) - Operation response error`, func() {
		getAPIKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAPIKeyPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					// TODO: Add check for include_activity query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAPIKey with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAPIKeyOptions model
				getAPIKeyOptionsModel := new(iamidentityv1.GetAPIKeyOptions)
				getAPIKeyOptionsModel.ID = core.StringPtr("testString")
				getAPIKeyOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAPIKeyOptionsModel.IncludeActivity = core.BoolPtr(false)
				getAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetAPIKey(getAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetAPIKey(getAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAPIKey(getAPIKeyOptions *GetAPIKeyOptions)`, func() {
		getAPIKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAPIKeyPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// TODO: Add check for include_activity query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke GetAPIKey successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetAPIKeyOptions model
				getAPIKeyOptionsModel := new(iamidentityv1.GetAPIKeyOptions)
				getAPIKeyOptionsModel.ID = core.StringPtr("testString")
				getAPIKeyOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAPIKeyOptionsModel.IncludeActivity = core.BoolPtr(false)
				getAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetAPIKeyWithContext(ctx, getAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetAPIKey(getAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetAPIKeyWithContext(ctx, getAPIKeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAPIKeyPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// TODO: Add check for include_activity query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke GetAPIKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetAPIKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAPIKeyOptions model
				getAPIKeyOptionsModel := new(iamidentityv1.GetAPIKeyOptions)
				getAPIKeyOptionsModel.ID = core.StringPtr("testString")
				getAPIKeyOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAPIKeyOptionsModel.IncludeActivity = core.BoolPtr(false)
				getAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetAPIKey(getAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAPIKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAPIKeyOptions model
				getAPIKeyOptionsModel := new(iamidentityv1.GetAPIKeyOptions)
				getAPIKeyOptionsModel.ID = core.StringPtr("testString")
				getAPIKeyOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAPIKeyOptionsModel.IncludeActivity = core.BoolPtr(false)
				getAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetAPIKey(getAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAPIKeyOptions model with no property values
				getAPIKeyOptionsModelNew := new(iamidentityv1.GetAPIKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetAPIKey(getAPIKeyOptionsModelNew)
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
			It(`Invoke GetAPIKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAPIKeyOptions model
				getAPIKeyOptionsModel := new(iamidentityv1.GetAPIKeyOptions)
				getAPIKeyOptionsModel.ID = core.StringPtr("testString")
				getAPIKeyOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAPIKeyOptionsModel.IncludeActivity = core.BoolPtr(false)
				getAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetAPIKey(getAPIKeyOptionsModel)
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
	Describe(`UpdateAPIKey(updateAPIKeyOptions *UpdateAPIKeyOptions) - Operation response error`, func() {
		updateAPIKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAPIKeyPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAPIKey with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateAPIKeyOptions model
				updateAPIKeyOptionsModel := new(iamidentityv1.UpdateAPIKeyOptions)
				updateAPIKeyOptionsModel.ID = core.StringPtr("testString")
				updateAPIKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Name = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Description = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateAPIKey(updateAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateAPIKey(updateAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAPIKey(updateAPIKeyOptions *UpdateAPIKeyOptions)`, func() {
		updateAPIKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAPIKeyPath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke UpdateAPIKey successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAPIKeyOptions model
				updateAPIKeyOptionsModel := new(iamidentityv1.UpdateAPIKeyOptions)
				updateAPIKeyOptionsModel.ID = core.StringPtr("testString")
				updateAPIKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Name = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Description = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.UpdateAPIKeyWithContext(ctx, updateAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.UpdateAPIKey(updateAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.UpdateAPIKeyWithContext(ctx, updateAPIKeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAPIKeyPath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke UpdateAPIKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateAPIKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAPIKeyOptions model
				updateAPIKeyOptionsModel := new(iamidentityv1.UpdateAPIKeyOptions)
				updateAPIKeyOptionsModel.ID = core.StringPtr("testString")
				updateAPIKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Name = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Description = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateAPIKey(updateAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAPIKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateAPIKeyOptions model
				updateAPIKeyOptionsModel := new(iamidentityv1.UpdateAPIKeyOptions)
				updateAPIKeyOptionsModel.ID = core.StringPtr("testString")
				updateAPIKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Name = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Description = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateAPIKey(updateAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAPIKeyOptions model with no property values
				updateAPIKeyOptionsModelNew := new(iamidentityv1.UpdateAPIKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateAPIKey(updateAPIKeyOptionsModelNew)
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
			It(`Invoke UpdateAPIKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateAPIKeyOptions model
				updateAPIKeyOptionsModel := new(iamidentityv1.UpdateAPIKeyOptions)
				updateAPIKeyOptionsModel.ID = core.StringPtr("testString")
				updateAPIKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Name = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Description = core.StringPtr("testString")
				updateAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.UpdateAPIKey(updateAPIKeyOptionsModel)
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
	Describe(`DeleteAPIKey(deleteAPIKeyOptions *DeleteAPIKeyOptions)`, func() {
		deleteAPIKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAPIKeyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAPIKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteAPIKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAPIKeyOptions model
				deleteAPIKeyOptionsModel := new(iamidentityv1.DeleteAPIKeyOptions)
				deleteAPIKeyOptionsModel.ID = core.StringPtr("testString")
				deleteAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteAPIKey(deleteAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAPIKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteAPIKeyOptions model
				deleteAPIKeyOptionsModel := new(iamidentityv1.DeleteAPIKeyOptions)
				deleteAPIKeyOptionsModel.ID = core.StringPtr("testString")
				deleteAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteAPIKey(deleteAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAPIKeyOptions model with no property values
				deleteAPIKeyOptionsModelNew := new(iamidentityv1.DeleteAPIKeyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteAPIKey(deleteAPIKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`LockAPIKey(lockAPIKeyOptions *LockAPIKeyOptions)`, func() {
		lockAPIKeyPath := "/v1/apikeys/testString/lock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(lockAPIKeyPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke LockAPIKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.LockAPIKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the LockAPIKeyOptions model
				lockAPIKeyOptionsModel := new(iamidentityv1.LockAPIKeyOptions)
				lockAPIKeyOptionsModel.ID = core.StringPtr("testString")
				lockAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.LockAPIKey(lockAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke LockAPIKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the LockAPIKeyOptions model
				lockAPIKeyOptionsModel := new(iamidentityv1.LockAPIKeyOptions)
				lockAPIKeyOptionsModel.ID = core.StringPtr("testString")
				lockAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.LockAPIKey(lockAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the LockAPIKeyOptions model with no property values
				lockAPIKeyOptionsModelNew := new(iamidentityv1.LockAPIKeyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.LockAPIKey(lockAPIKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UnlockAPIKey(unlockAPIKeyOptions *UnlockAPIKeyOptions)`, func() {
		unlockAPIKeyPath := "/v1/apikeys/testString/lock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unlockAPIKeyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UnlockAPIKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.UnlockAPIKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UnlockAPIKeyOptions model
				unlockAPIKeyOptionsModel := new(iamidentityv1.UnlockAPIKeyOptions)
				unlockAPIKeyOptionsModel.ID = core.StringPtr("testString")
				unlockAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.UnlockAPIKey(unlockAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UnlockAPIKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UnlockAPIKeyOptions model
				unlockAPIKeyOptionsModel := new(iamidentityv1.UnlockAPIKeyOptions)
				unlockAPIKeyOptionsModel.ID = core.StringPtr("testString")
				unlockAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.UnlockAPIKey(unlockAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UnlockAPIKeyOptions model with no property values
				unlockAPIKeyOptionsModelNew := new(iamidentityv1.UnlockAPIKeyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.UnlockAPIKey(unlockAPIKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListServiceIds(listServiceIdsOptions *ListServiceIdsOptions) - Operation response error`, func() {
		listServiceIdsPath := "/v1/serviceids/"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listServiceIdsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["pagesize"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListServiceIds with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListServiceIdsOptions model
				listServiceIdsOptionsModel := new(iamidentityv1.ListServiceIdsOptions)
				listServiceIdsOptionsModel.AccountID = core.StringPtr("testString")
				listServiceIdsOptionsModel.Name = core.StringPtr("testString")
				listServiceIdsOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listServiceIdsOptionsModel.Pagetoken = core.StringPtr("testString")
				listServiceIdsOptionsModel.Sort = core.StringPtr("testString")
				listServiceIdsOptionsModel.Order = core.StringPtr("asc")
				listServiceIdsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listServiceIdsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListServiceIds(listServiceIdsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListServiceIds(listServiceIdsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListServiceIds(listServiceIdsOptions *ListServiceIdsOptions)`, func() {
		listServiceIdsPath := "/v1/serviceids/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listServiceIdsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["pagesize"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "serviceids": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}, "activity": {"last_authn": "LastAuthn", "authn_count": 10}}]}`)
				}))
			})
			It(`Invoke ListServiceIds successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListServiceIdsOptions model
				listServiceIdsOptionsModel := new(iamidentityv1.ListServiceIdsOptions)
				listServiceIdsOptionsModel.AccountID = core.StringPtr("testString")
				listServiceIdsOptionsModel.Name = core.StringPtr("testString")
				listServiceIdsOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listServiceIdsOptionsModel.Pagetoken = core.StringPtr("testString")
				listServiceIdsOptionsModel.Sort = core.StringPtr("testString")
				listServiceIdsOptionsModel.Order = core.StringPtr("asc")
				listServiceIdsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listServiceIdsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListServiceIdsWithContext(ctx, listServiceIdsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListServiceIds(listServiceIdsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListServiceIdsWithContext(ctx, listServiceIdsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listServiceIdsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["pagesize"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "serviceids": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}, "activity": {"last_authn": "LastAuthn", "authn_count": 10}}]}`)
				}))
			})
			It(`Invoke ListServiceIds successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListServiceIds(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListServiceIdsOptions model
				listServiceIdsOptionsModel := new(iamidentityv1.ListServiceIdsOptions)
				listServiceIdsOptionsModel.AccountID = core.StringPtr("testString")
				listServiceIdsOptionsModel.Name = core.StringPtr("testString")
				listServiceIdsOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listServiceIdsOptionsModel.Pagetoken = core.StringPtr("testString")
				listServiceIdsOptionsModel.Sort = core.StringPtr("testString")
				listServiceIdsOptionsModel.Order = core.StringPtr("asc")
				listServiceIdsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listServiceIdsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListServiceIds(listServiceIdsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListServiceIds with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListServiceIdsOptions model
				listServiceIdsOptionsModel := new(iamidentityv1.ListServiceIdsOptions)
				listServiceIdsOptionsModel.AccountID = core.StringPtr("testString")
				listServiceIdsOptionsModel.Name = core.StringPtr("testString")
				listServiceIdsOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listServiceIdsOptionsModel.Pagetoken = core.StringPtr("testString")
				listServiceIdsOptionsModel.Sort = core.StringPtr("testString")
				listServiceIdsOptionsModel.Order = core.StringPtr("asc")
				listServiceIdsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listServiceIdsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListServiceIds(listServiceIdsOptionsModel)
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
			It(`Invoke ListServiceIds successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListServiceIdsOptions model
				listServiceIdsOptionsModel := new(iamidentityv1.ListServiceIdsOptions)
				listServiceIdsOptionsModel.AccountID = core.StringPtr("testString")
				listServiceIdsOptionsModel.Name = core.StringPtr("testString")
				listServiceIdsOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listServiceIdsOptionsModel.Pagetoken = core.StringPtr("testString")
				listServiceIdsOptionsModel.Sort = core.StringPtr("testString")
				listServiceIdsOptionsModel.Order = core.StringPtr("asc")
				listServiceIdsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listServiceIdsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListServiceIds(listServiceIdsOptionsModel)
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
	Describe(`CreateServiceID(createServiceIDOptions *CreateServiceIDOptions) - Operation response error`, func() {
		createServiceIDPath := "/v1/serviceids/"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createServiceIDPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "false")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateServiceID with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the APIKeyInsideCreateServiceIDRequest model
				apiKeyInsideCreateServiceIDRequestModel := new(iamidentityv1.APIKeyInsideCreateServiceIDRequest)
				apiKeyInsideCreateServiceIDRequestModel.Name = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Description = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Apikey = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.StoreValue = core.BoolPtr(true)

				// Construct an instance of the CreateServiceIDOptions model
				createServiceIDOptionsModel := new(iamidentityv1.CreateServiceIDOptions)
				createServiceIDOptionsModel.AccountID = core.StringPtr("testString")
				createServiceIDOptionsModel.Name = core.StringPtr("testString")
				createServiceIDOptionsModel.Description = core.StringPtr("testString")
				createServiceIDOptionsModel.UniqueInstanceCrns = []string{"testString"}
				createServiceIDOptionsModel.Apikey = apiKeyInsideCreateServiceIDRequestModel
				createServiceIDOptionsModel.EntityLock = core.StringPtr("false")
				createServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateServiceID(createServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateServiceID(createServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateServiceID(createServiceIDOptions *CreateServiceIDOptions)`, func() {
		createServiceIDPath := "/v1/serviceids/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createServiceIDPath))
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

					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "false")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}, "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke CreateServiceID successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the APIKeyInsideCreateServiceIDRequest model
				apiKeyInsideCreateServiceIDRequestModel := new(iamidentityv1.APIKeyInsideCreateServiceIDRequest)
				apiKeyInsideCreateServiceIDRequestModel.Name = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Description = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Apikey = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.StoreValue = core.BoolPtr(true)

				// Construct an instance of the CreateServiceIDOptions model
				createServiceIDOptionsModel := new(iamidentityv1.CreateServiceIDOptions)
				createServiceIDOptionsModel.AccountID = core.StringPtr("testString")
				createServiceIDOptionsModel.Name = core.StringPtr("testString")
				createServiceIDOptionsModel.Description = core.StringPtr("testString")
				createServiceIDOptionsModel.UniqueInstanceCrns = []string{"testString"}
				createServiceIDOptionsModel.Apikey = apiKeyInsideCreateServiceIDRequestModel
				createServiceIDOptionsModel.EntityLock = core.StringPtr("false")
				createServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateServiceIDWithContext(ctx, createServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateServiceID(createServiceIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateServiceIDWithContext(ctx, createServiceIDOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createServiceIDPath))
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

					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "false")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}, "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke CreateServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the APIKeyInsideCreateServiceIDRequest model
				apiKeyInsideCreateServiceIDRequestModel := new(iamidentityv1.APIKeyInsideCreateServiceIDRequest)
				apiKeyInsideCreateServiceIDRequestModel.Name = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Description = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Apikey = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.StoreValue = core.BoolPtr(true)

				// Construct an instance of the CreateServiceIDOptions model
				createServiceIDOptionsModel := new(iamidentityv1.CreateServiceIDOptions)
				createServiceIDOptionsModel.AccountID = core.StringPtr("testString")
				createServiceIDOptionsModel.Name = core.StringPtr("testString")
				createServiceIDOptionsModel.Description = core.StringPtr("testString")
				createServiceIDOptionsModel.UniqueInstanceCrns = []string{"testString"}
				createServiceIDOptionsModel.Apikey = apiKeyInsideCreateServiceIDRequestModel
				createServiceIDOptionsModel.EntityLock = core.StringPtr("false")
				createServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateServiceID(createServiceIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateServiceID with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the APIKeyInsideCreateServiceIDRequest model
				apiKeyInsideCreateServiceIDRequestModel := new(iamidentityv1.APIKeyInsideCreateServiceIDRequest)
				apiKeyInsideCreateServiceIDRequestModel.Name = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Description = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Apikey = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.StoreValue = core.BoolPtr(true)

				// Construct an instance of the CreateServiceIDOptions model
				createServiceIDOptionsModel := new(iamidentityv1.CreateServiceIDOptions)
				createServiceIDOptionsModel.AccountID = core.StringPtr("testString")
				createServiceIDOptionsModel.Name = core.StringPtr("testString")
				createServiceIDOptionsModel.Description = core.StringPtr("testString")
				createServiceIDOptionsModel.UniqueInstanceCrns = []string{"testString"}
				createServiceIDOptionsModel.Apikey = apiKeyInsideCreateServiceIDRequestModel
				createServiceIDOptionsModel.EntityLock = core.StringPtr("false")
				createServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateServiceID(createServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateServiceIDOptions model with no property values
				createServiceIDOptionsModelNew := new(iamidentityv1.CreateServiceIDOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateServiceID(createServiceIDOptionsModelNew)
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
			It(`Invoke CreateServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the APIKeyInsideCreateServiceIDRequest model
				apiKeyInsideCreateServiceIDRequestModel := new(iamidentityv1.APIKeyInsideCreateServiceIDRequest)
				apiKeyInsideCreateServiceIDRequestModel.Name = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Description = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Apikey = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.StoreValue = core.BoolPtr(true)

				// Construct an instance of the CreateServiceIDOptions model
				createServiceIDOptionsModel := new(iamidentityv1.CreateServiceIDOptions)
				createServiceIDOptionsModel.AccountID = core.StringPtr("testString")
				createServiceIDOptionsModel.Name = core.StringPtr("testString")
				createServiceIDOptionsModel.Description = core.StringPtr("testString")
				createServiceIDOptionsModel.UniqueInstanceCrns = []string{"testString"}
				createServiceIDOptionsModel.Apikey = apiKeyInsideCreateServiceIDRequestModel
				createServiceIDOptionsModel.EntityLock = core.StringPtr("false")
				createServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateServiceID(createServiceIDOptionsModel)
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
	Describe(`GetServiceID(getServiceIDOptions *GetServiceIDOptions) - Operation response error`, func() {
		getServiceIDPath := "/v1/serviceids/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceIDPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					// TODO: Add check for include_activity query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetServiceID with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetServiceIDOptions model
				getServiceIDOptionsModel := new(iamidentityv1.GetServiceIDOptions)
				getServiceIDOptionsModel.ID = core.StringPtr("testString")
				getServiceIDOptionsModel.IncludeHistory = core.BoolPtr(false)
				getServiceIDOptionsModel.IncludeActivity = core.BoolPtr(false)
				getServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetServiceID(getServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetServiceID(getServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetServiceID(getServiceIDOptions *GetServiceIDOptions)`, func() {
		getServiceIDPath := "/v1/serviceids/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceIDPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// TODO: Add check for include_activity query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}, "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke GetServiceID successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetServiceIDOptions model
				getServiceIDOptionsModel := new(iamidentityv1.GetServiceIDOptions)
				getServiceIDOptionsModel.ID = core.StringPtr("testString")
				getServiceIDOptionsModel.IncludeHistory = core.BoolPtr(false)
				getServiceIDOptionsModel.IncludeActivity = core.BoolPtr(false)
				getServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetServiceIDWithContext(ctx, getServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetServiceID(getServiceIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetServiceIDWithContext(ctx, getServiceIDOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getServiceIDPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// TODO: Add check for include_activity query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}, "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke GetServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServiceIDOptions model
				getServiceIDOptionsModel := new(iamidentityv1.GetServiceIDOptions)
				getServiceIDOptionsModel.ID = core.StringPtr("testString")
				getServiceIDOptionsModel.IncludeHistory = core.BoolPtr(false)
				getServiceIDOptionsModel.IncludeActivity = core.BoolPtr(false)
				getServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetServiceID(getServiceIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetServiceID with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetServiceIDOptions model
				getServiceIDOptionsModel := new(iamidentityv1.GetServiceIDOptions)
				getServiceIDOptionsModel.ID = core.StringPtr("testString")
				getServiceIDOptionsModel.IncludeHistory = core.BoolPtr(false)
				getServiceIDOptionsModel.IncludeActivity = core.BoolPtr(false)
				getServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetServiceID(getServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetServiceIDOptions model with no property values
				getServiceIDOptionsModelNew := new(iamidentityv1.GetServiceIDOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetServiceID(getServiceIDOptionsModelNew)
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
			It(`Invoke GetServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetServiceIDOptions model
				getServiceIDOptionsModel := new(iamidentityv1.GetServiceIDOptions)
				getServiceIDOptionsModel.ID = core.StringPtr("testString")
				getServiceIDOptionsModel.IncludeHistory = core.BoolPtr(false)
				getServiceIDOptionsModel.IncludeActivity = core.BoolPtr(false)
				getServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetServiceID(getServiceIDOptionsModel)
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
	Describe(`UpdateServiceID(updateServiceIDOptions *UpdateServiceIDOptions) - Operation response error`, func() {
		updateServiceIDPath := "/v1/serviceids/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateServiceIDPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateServiceID with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateServiceIDOptions model
				updateServiceIDOptionsModel := new(iamidentityv1.UpdateServiceIDOptions)
				updateServiceIDOptionsModel.ID = core.StringPtr("testString")
				updateServiceIDOptionsModel.IfMatch = core.StringPtr("testString")
				updateServiceIDOptionsModel.Name = core.StringPtr("testString")
				updateServiceIDOptionsModel.Description = core.StringPtr("testString")
				updateServiceIDOptionsModel.UniqueInstanceCrns = []string{"testString"}
				updateServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateServiceID(updateServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateServiceID(updateServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateServiceID(updateServiceIDOptions *UpdateServiceIDOptions)`, func() {
		updateServiceIDPath := "/v1/serviceids/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateServiceIDPath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}, "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke UpdateServiceID successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the UpdateServiceIDOptions model
				updateServiceIDOptionsModel := new(iamidentityv1.UpdateServiceIDOptions)
				updateServiceIDOptionsModel.ID = core.StringPtr("testString")
				updateServiceIDOptionsModel.IfMatch = core.StringPtr("testString")
				updateServiceIDOptionsModel.Name = core.StringPtr("testString")
				updateServiceIDOptionsModel.Description = core.StringPtr("testString")
				updateServiceIDOptionsModel.UniqueInstanceCrns = []string{"testString"}
				updateServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.UpdateServiceIDWithContext(ctx, updateServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.UpdateServiceID(updateServiceIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.UpdateServiceIDWithContext(ctx, updateServiceIDOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateServiceIDPath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}, "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke UpdateServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateServiceIDOptions model
				updateServiceIDOptionsModel := new(iamidentityv1.UpdateServiceIDOptions)
				updateServiceIDOptionsModel.ID = core.StringPtr("testString")
				updateServiceIDOptionsModel.IfMatch = core.StringPtr("testString")
				updateServiceIDOptionsModel.Name = core.StringPtr("testString")
				updateServiceIDOptionsModel.Description = core.StringPtr("testString")
				updateServiceIDOptionsModel.UniqueInstanceCrns = []string{"testString"}
				updateServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateServiceID(updateServiceIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateServiceID with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateServiceIDOptions model
				updateServiceIDOptionsModel := new(iamidentityv1.UpdateServiceIDOptions)
				updateServiceIDOptionsModel.ID = core.StringPtr("testString")
				updateServiceIDOptionsModel.IfMatch = core.StringPtr("testString")
				updateServiceIDOptionsModel.Name = core.StringPtr("testString")
				updateServiceIDOptionsModel.Description = core.StringPtr("testString")
				updateServiceIDOptionsModel.UniqueInstanceCrns = []string{"testString"}
				updateServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateServiceID(updateServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateServiceIDOptions model with no property values
				updateServiceIDOptionsModelNew := new(iamidentityv1.UpdateServiceIDOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateServiceID(updateServiceIDOptionsModelNew)
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
			It(`Invoke UpdateServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateServiceIDOptions model
				updateServiceIDOptionsModel := new(iamidentityv1.UpdateServiceIDOptions)
				updateServiceIDOptionsModel.ID = core.StringPtr("testString")
				updateServiceIDOptionsModel.IfMatch = core.StringPtr("testString")
				updateServiceIDOptionsModel.Name = core.StringPtr("testString")
				updateServiceIDOptionsModel.Description = core.StringPtr("testString")
				updateServiceIDOptionsModel.UniqueInstanceCrns = []string{"testString"}
				updateServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.UpdateServiceID(updateServiceIDOptionsModel)
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
	Describe(`DeleteServiceID(deleteServiceIDOptions *DeleteServiceIDOptions)`, func() {
		deleteServiceIDPath := "/v1/serviceids/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteServiceIDPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteServiceIDOptions model
				deleteServiceIDOptionsModel := new(iamidentityv1.DeleteServiceIDOptions)
				deleteServiceIDOptionsModel.ID = core.StringPtr("testString")
				deleteServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteServiceID(deleteServiceIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteServiceID with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceIDOptions model
				deleteServiceIDOptionsModel := new(iamidentityv1.DeleteServiceIDOptions)
				deleteServiceIDOptionsModel.ID = core.StringPtr("testString")
				deleteServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteServiceID(deleteServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteServiceIDOptions model with no property values
				deleteServiceIDOptionsModelNew := new(iamidentityv1.DeleteServiceIDOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteServiceID(deleteServiceIDOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`LockServiceID(lockServiceIDOptions *LockServiceIDOptions)`, func() {
		lockServiceIDPath := "/v1/serviceids/testString/lock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(lockServiceIDPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke LockServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.LockServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the LockServiceIDOptions model
				lockServiceIDOptionsModel := new(iamidentityv1.LockServiceIDOptions)
				lockServiceIDOptionsModel.ID = core.StringPtr("testString")
				lockServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.LockServiceID(lockServiceIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke LockServiceID with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the LockServiceIDOptions model
				lockServiceIDOptionsModel := new(iamidentityv1.LockServiceIDOptions)
				lockServiceIDOptionsModel.ID = core.StringPtr("testString")
				lockServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.LockServiceID(lockServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the LockServiceIDOptions model with no property values
				lockServiceIDOptionsModelNew := new(iamidentityv1.LockServiceIDOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.LockServiceID(lockServiceIDOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UnlockServiceID(unlockServiceIDOptions *UnlockServiceIDOptions)`, func() {
		unlockServiceIDPath := "/v1/serviceids/testString/lock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unlockServiceIDPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UnlockServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.UnlockServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UnlockServiceIDOptions model
				unlockServiceIDOptionsModel := new(iamidentityv1.UnlockServiceIDOptions)
				unlockServiceIDOptionsModel.ID = core.StringPtr("testString")
				unlockServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.UnlockServiceID(unlockServiceIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UnlockServiceID with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UnlockServiceIDOptions model
				unlockServiceIDOptionsModel := new(iamidentityv1.UnlockServiceIDOptions)
				unlockServiceIDOptionsModel.ID = core.StringPtr("testString")
				unlockServiceIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.UnlockServiceID(unlockServiceIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UnlockServiceIDOptions model with no property values
				unlockServiceIDOptionsModelNew := new(iamidentityv1.UnlockServiceIDOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.UnlockServiceID(unlockServiceIDOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProfile(createProfileOptions *CreateProfileOptions) - Operation response error`, func() {
		createProfilePath := "/v1/profiles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProfilePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProfile with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(iamidentityv1.CreateProfileOptions)
				createProfileOptionsModel.Name = core.StringPtr("testString")
				createProfileOptionsModel.AccountID = core.StringPtr("testString")
				createProfileOptionsModel.Description = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProfile(createProfileOptions *CreateProfileOptions)`, func() {
		createProfilePath := "/v1/profiles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProfilePath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "template_id": "TemplateID", "assignment_id": "AssignmentID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke CreateProfile successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(iamidentityv1.CreateProfileOptions)
				createProfileOptionsModel.Name = core.StringPtr("testString")
				createProfileOptionsModel.AccountID = core.StringPtr("testString")
				createProfileOptionsModel.Description = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateProfileWithContext(ctx, createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateProfileWithContext(ctx, createProfileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createProfilePath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "template_id": "TemplateID", "assignment_id": "AssignmentID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke CreateProfile successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(iamidentityv1.CreateProfileOptions)
				createProfileOptionsModel.Name = core.StringPtr("testString")
				createProfileOptionsModel.AccountID = core.StringPtr("testString")
				createProfileOptionsModel.Description = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProfile with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(iamidentityv1.CreateProfileOptions)
				createProfileOptionsModel.Name = core.StringPtr("testString")
				createProfileOptionsModel.AccountID = core.StringPtr("testString")
				createProfileOptionsModel.Description = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProfileOptions model with no property values
				createProfileOptionsModelNew := new(iamidentityv1.CreateProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateProfile(createProfileOptionsModelNew)
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
			It(`Invoke CreateProfile successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(iamidentityv1.CreateProfileOptions)
				createProfileOptionsModel.Name = core.StringPtr("testString")
				createProfileOptionsModel.AccountID = core.StringPtr("testString")
				createProfileOptionsModel.Description = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateProfile(createProfileOptionsModel)
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
	Describe(`ListProfiles(listProfilesOptions *ListProfilesOptions) - Operation response error`, func() {
		listProfilesPath := "/v1/profiles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["pagesize"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProfiles with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(iamidentityv1.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.Name = core.StringPtr("testString")
				listProfilesOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listProfilesOptionsModel.Sort = core.StringPtr("testString")
				listProfilesOptionsModel.Order = core.StringPtr("asc")
				listProfilesOptionsModel.IncludeHistory = core.BoolPtr(false)
				listProfilesOptionsModel.Pagetoken = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProfiles(listProfilesOptions *ListProfilesOptions)`, func() {
		listProfilesPath := "/v1/profiles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["pagesize"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "profiles": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "template_id": "TemplateID", "assignment_id": "AssignmentID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(iamidentityv1.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.Name = core.StringPtr("testString")
				listProfilesOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listProfilesOptionsModel.Sort = core.StringPtr("testString")
				listProfilesOptionsModel.Order = core.StringPtr("asc")
				listProfilesOptionsModel.IncludeHistory = core.BoolPtr(false)
				listProfilesOptionsModel.Pagetoken = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListProfilesWithContext(ctx, listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListProfilesWithContext(ctx, listProfilesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["pagesize"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "profiles": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "template_id": "TemplateID", "assignment_id": "AssignmentID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(iamidentityv1.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.Name = core.StringPtr("testString")
				listProfilesOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listProfilesOptionsModel.Sort = core.StringPtr("testString")
				listProfilesOptionsModel.Order = core.StringPtr("asc")
				listProfilesOptionsModel.IncludeHistory = core.BoolPtr(false)
				listProfilesOptionsModel.Pagetoken = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProfiles with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(iamidentityv1.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.Name = core.StringPtr("testString")
				listProfilesOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listProfilesOptionsModel.Sort = core.StringPtr("testString")
				listProfilesOptionsModel.Order = core.StringPtr("asc")
				listProfilesOptionsModel.IncludeHistory = core.BoolPtr(false)
				listProfilesOptionsModel.Pagetoken = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListProfilesOptions model with no property values
				listProfilesOptionsModelNew := new(iamidentityv1.ListProfilesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.ListProfiles(listProfilesOptionsModelNew)
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
			It(`Invoke ListProfiles successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(iamidentityv1.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.Name = core.StringPtr("testString")
				listProfilesOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listProfilesOptionsModel.Sort = core.StringPtr("testString")
				listProfilesOptionsModel.Order = core.StringPtr("asc")
				listProfilesOptionsModel.IncludeHistory = core.BoolPtr(false)
				listProfilesOptionsModel.Pagetoken = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListProfiles(listProfilesOptionsModel)
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
	Describe(`GetProfile(getProfileOptions *GetProfileOptions) - Operation response error`, func() {
		getProfilePath := "/v1/profiles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_activity query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfile with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(iamidentityv1.GetProfileOptions)
				getProfileOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileOptionsModel.IncludeActivity = core.BoolPtr(false)
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfile(getProfileOptions *GetProfileOptions)`, func() {
		getProfilePath := "/v1/profiles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_activity query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "template_id": "TemplateID", "assignment_id": "AssignmentID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke GetProfile successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(iamidentityv1.GetProfileOptions)
				getProfileOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileOptionsModel.IncludeActivity = core.BoolPtr(false)
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetProfileWithContext(ctx, getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetProfileWithContext(ctx, getProfileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_activity query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "template_id": "TemplateID", "assignment_id": "AssignmentID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke GetProfile successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(iamidentityv1.GetProfileOptions)
				getProfileOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileOptionsModel.IncludeActivity = core.BoolPtr(false)
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfile with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(iamidentityv1.GetProfileOptions)
				getProfileOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileOptionsModel.IncludeActivity = core.BoolPtr(false)
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileOptions model with no property values
				getProfileOptionsModelNew := new(iamidentityv1.GetProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetProfile(getProfileOptionsModelNew)
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
			It(`Invoke GetProfile successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(iamidentityv1.GetProfileOptions)
				getProfileOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileOptionsModel.IncludeActivity = core.BoolPtr(false)
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetProfile(getProfileOptionsModel)
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
	Describe(`UpdateProfile(updateProfileOptions *UpdateProfileOptions) - Operation response error`, func() {
		updateProfilePath := "/v1/profiles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProfilePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProfile with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateProfileOptions model
				updateProfileOptionsModel := new(iamidentityv1.UpdateProfileOptions)
				updateProfileOptionsModel.ProfileID = core.StringPtr("testString")
				updateProfileOptionsModel.IfMatch = core.StringPtr("testString")
				updateProfileOptionsModel.Name = core.StringPtr("testString")
				updateProfileOptionsModel.Description = core.StringPtr("testString")
				updateProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateProfile(updateProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateProfile(updateProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProfile(updateProfileOptions *UpdateProfileOptions)`, func() {
		updateProfilePath := "/v1/profiles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProfilePath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "template_id": "TemplateID", "assignment_id": "AssignmentID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke UpdateProfile successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the UpdateProfileOptions model
				updateProfileOptionsModel := new(iamidentityv1.UpdateProfileOptions)
				updateProfileOptionsModel.ProfileID = core.StringPtr("testString")
				updateProfileOptionsModel.IfMatch = core.StringPtr("testString")
				updateProfileOptionsModel.Name = core.StringPtr("testString")
				updateProfileOptionsModel.Description = core.StringPtr("testString")
				updateProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.UpdateProfileWithContext(ctx, updateProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.UpdateProfile(updateProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.UpdateProfileWithContext(ctx, updateProfileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProfilePath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "template_id": "TemplateID", "assignment_id": "AssignmentID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
				}))
			})
			It(`Invoke UpdateProfile successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateProfileOptions model
				updateProfileOptionsModel := new(iamidentityv1.UpdateProfileOptions)
				updateProfileOptionsModel.ProfileID = core.StringPtr("testString")
				updateProfileOptionsModel.IfMatch = core.StringPtr("testString")
				updateProfileOptionsModel.Name = core.StringPtr("testString")
				updateProfileOptionsModel.Description = core.StringPtr("testString")
				updateProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateProfile(updateProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProfile with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateProfileOptions model
				updateProfileOptionsModel := new(iamidentityv1.UpdateProfileOptions)
				updateProfileOptionsModel.ProfileID = core.StringPtr("testString")
				updateProfileOptionsModel.IfMatch = core.StringPtr("testString")
				updateProfileOptionsModel.Name = core.StringPtr("testString")
				updateProfileOptionsModel.Description = core.StringPtr("testString")
				updateProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateProfile(updateProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProfileOptions model with no property values
				updateProfileOptionsModelNew := new(iamidentityv1.UpdateProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateProfile(updateProfileOptionsModelNew)
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
			It(`Invoke UpdateProfile successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateProfileOptions model
				updateProfileOptionsModel := new(iamidentityv1.UpdateProfileOptions)
				updateProfileOptionsModel.ProfileID = core.StringPtr("testString")
				updateProfileOptionsModel.IfMatch = core.StringPtr("testString")
				updateProfileOptionsModel.Name = core.StringPtr("testString")
				updateProfileOptionsModel.Description = core.StringPtr("testString")
				updateProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.UpdateProfile(updateProfileOptionsModel)
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
	Describe(`DeleteProfile(deleteProfileOptions *DeleteProfileOptions)`, func() {
		deleteProfilePath := "/v1/profiles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProfilePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteProfile successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProfileOptions model
				deleteProfileOptionsModel := new(iamidentityv1.DeleteProfileOptions)
				deleteProfileOptionsModel.ProfileID = core.StringPtr("testString")
				deleteProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteProfile(deleteProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProfile with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileOptions model
				deleteProfileOptionsModel := new(iamidentityv1.DeleteProfileOptions)
				deleteProfileOptionsModel.ProfileID = core.StringPtr("testString")
				deleteProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteProfile(deleteProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProfileOptions model with no property values
				deleteProfileOptionsModelNew := new(iamidentityv1.DeleteProfileOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteProfile(deleteProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateClaimRule(createClaimRuleOptions *CreateClaimRuleOptions) - Operation response error`, func() {
		createClaimRulePath := "/v1/profiles/testString/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClaimRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateClaimRule with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")

				// Construct an instance of the CreateClaimRuleOptions model
				createClaimRuleOptionsModel := new(iamidentityv1.CreateClaimRuleOptions)
				createClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				createClaimRuleOptionsModel.Type = core.StringPtr("testString")
				createClaimRuleOptionsModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				createClaimRuleOptionsModel.Context = responseContextModel
				createClaimRuleOptionsModel.Name = core.StringPtr("testString")
				createClaimRuleOptionsModel.RealmName = core.StringPtr("testString")
				createClaimRuleOptionsModel.CrType = core.StringPtr("testString")
				createClaimRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				createClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateClaimRule(createClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateClaimRule(createClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateClaimRule(createClaimRuleOptions *CreateClaimRuleOptions)`, func() {
		createClaimRulePath := "/v1/profiles/testString/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClaimRulePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "type": "Type", "realm_name": "RealmName", "expiration": 10, "cr_type": "CrType", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}`)
				}))
			})
			It(`Invoke CreateClaimRule successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")

				// Construct an instance of the CreateClaimRuleOptions model
				createClaimRuleOptionsModel := new(iamidentityv1.CreateClaimRuleOptions)
				createClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				createClaimRuleOptionsModel.Type = core.StringPtr("testString")
				createClaimRuleOptionsModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				createClaimRuleOptionsModel.Context = responseContextModel
				createClaimRuleOptionsModel.Name = core.StringPtr("testString")
				createClaimRuleOptionsModel.RealmName = core.StringPtr("testString")
				createClaimRuleOptionsModel.CrType = core.StringPtr("testString")
				createClaimRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				createClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateClaimRuleWithContext(ctx, createClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateClaimRule(createClaimRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateClaimRuleWithContext(ctx, createClaimRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createClaimRulePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "type": "Type", "realm_name": "RealmName", "expiration": 10, "cr_type": "CrType", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}`)
				}))
			})
			It(`Invoke CreateClaimRule successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateClaimRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")

				// Construct an instance of the CreateClaimRuleOptions model
				createClaimRuleOptionsModel := new(iamidentityv1.CreateClaimRuleOptions)
				createClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				createClaimRuleOptionsModel.Type = core.StringPtr("testString")
				createClaimRuleOptionsModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				createClaimRuleOptionsModel.Context = responseContextModel
				createClaimRuleOptionsModel.Name = core.StringPtr("testString")
				createClaimRuleOptionsModel.RealmName = core.StringPtr("testString")
				createClaimRuleOptionsModel.CrType = core.StringPtr("testString")
				createClaimRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				createClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateClaimRule(createClaimRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateClaimRule with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")

				// Construct an instance of the CreateClaimRuleOptions model
				createClaimRuleOptionsModel := new(iamidentityv1.CreateClaimRuleOptions)
				createClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				createClaimRuleOptionsModel.Type = core.StringPtr("testString")
				createClaimRuleOptionsModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				createClaimRuleOptionsModel.Context = responseContextModel
				createClaimRuleOptionsModel.Name = core.StringPtr("testString")
				createClaimRuleOptionsModel.RealmName = core.StringPtr("testString")
				createClaimRuleOptionsModel.CrType = core.StringPtr("testString")
				createClaimRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				createClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateClaimRule(createClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateClaimRuleOptions model with no property values
				createClaimRuleOptionsModelNew := new(iamidentityv1.CreateClaimRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateClaimRule(createClaimRuleOptionsModelNew)
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
			It(`Invoke CreateClaimRule successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")

				// Construct an instance of the CreateClaimRuleOptions model
				createClaimRuleOptionsModel := new(iamidentityv1.CreateClaimRuleOptions)
				createClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				createClaimRuleOptionsModel.Type = core.StringPtr("testString")
				createClaimRuleOptionsModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				createClaimRuleOptionsModel.Context = responseContextModel
				createClaimRuleOptionsModel.Name = core.StringPtr("testString")
				createClaimRuleOptionsModel.RealmName = core.StringPtr("testString")
				createClaimRuleOptionsModel.CrType = core.StringPtr("testString")
				createClaimRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				createClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateClaimRule(createClaimRuleOptionsModel)
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
	Describe(`ListClaimRules(listClaimRulesOptions *ListClaimRulesOptions) - Operation response error`, func() {
		listClaimRulesPath := "/v1/profiles/testString/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClaimRulesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListClaimRules with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListClaimRulesOptions model
				listClaimRulesOptionsModel := new(iamidentityv1.ListClaimRulesOptions)
				listClaimRulesOptionsModel.ProfileID = core.StringPtr("testString")
				listClaimRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListClaimRules(listClaimRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListClaimRules(listClaimRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListClaimRules(listClaimRulesOptions *ListClaimRulesOptions)`, func() {
		listClaimRulesPath := "/v1/profiles/testString/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClaimRulesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "rules": [{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "type": "Type", "realm_name": "RealmName", "expiration": 10, "cr_type": "CrType", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}]}`)
				}))
			})
			It(`Invoke ListClaimRules successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListClaimRulesOptions model
				listClaimRulesOptionsModel := new(iamidentityv1.ListClaimRulesOptions)
				listClaimRulesOptionsModel.ProfileID = core.StringPtr("testString")
				listClaimRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListClaimRulesWithContext(ctx, listClaimRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListClaimRules(listClaimRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListClaimRulesWithContext(ctx, listClaimRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listClaimRulesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "rules": [{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "type": "Type", "realm_name": "RealmName", "expiration": 10, "cr_type": "CrType", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}]}`)
				}))
			})
			It(`Invoke ListClaimRules successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListClaimRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListClaimRulesOptions model
				listClaimRulesOptionsModel := new(iamidentityv1.ListClaimRulesOptions)
				listClaimRulesOptionsModel.ProfileID = core.StringPtr("testString")
				listClaimRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListClaimRules(listClaimRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListClaimRules with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListClaimRulesOptions model
				listClaimRulesOptionsModel := new(iamidentityv1.ListClaimRulesOptions)
				listClaimRulesOptionsModel.ProfileID = core.StringPtr("testString")
				listClaimRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListClaimRules(listClaimRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListClaimRulesOptions model with no property values
				listClaimRulesOptionsModelNew := new(iamidentityv1.ListClaimRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.ListClaimRules(listClaimRulesOptionsModelNew)
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
			It(`Invoke ListClaimRules successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListClaimRulesOptions model
				listClaimRulesOptionsModel := new(iamidentityv1.ListClaimRulesOptions)
				listClaimRulesOptionsModel.ProfileID = core.StringPtr("testString")
				listClaimRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListClaimRules(listClaimRulesOptionsModel)
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
	Describe(`GetClaimRule(getClaimRuleOptions *GetClaimRuleOptions) - Operation response error`, func() {
		getClaimRulePath := "/v1/profiles/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClaimRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetClaimRule with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetClaimRuleOptions model
				getClaimRuleOptionsModel := new(iamidentityv1.GetClaimRuleOptions)
				getClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				getClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				getClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetClaimRule(getClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetClaimRule(getClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetClaimRule(getClaimRuleOptions *GetClaimRuleOptions)`, func() {
		getClaimRulePath := "/v1/profiles/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClaimRulePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "type": "Type", "realm_name": "RealmName", "expiration": 10, "cr_type": "CrType", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}`)
				}))
			})
			It(`Invoke GetClaimRule successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetClaimRuleOptions model
				getClaimRuleOptionsModel := new(iamidentityv1.GetClaimRuleOptions)
				getClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				getClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				getClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetClaimRuleWithContext(ctx, getClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetClaimRule(getClaimRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetClaimRuleWithContext(ctx, getClaimRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getClaimRulePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "type": "Type", "realm_name": "RealmName", "expiration": 10, "cr_type": "CrType", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}`)
				}))
			})
			It(`Invoke GetClaimRule successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetClaimRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetClaimRuleOptions model
				getClaimRuleOptionsModel := new(iamidentityv1.GetClaimRuleOptions)
				getClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				getClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				getClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetClaimRule(getClaimRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetClaimRule with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetClaimRuleOptions model
				getClaimRuleOptionsModel := new(iamidentityv1.GetClaimRuleOptions)
				getClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				getClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				getClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetClaimRule(getClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetClaimRuleOptions model with no property values
				getClaimRuleOptionsModelNew := new(iamidentityv1.GetClaimRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetClaimRule(getClaimRuleOptionsModelNew)
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
			It(`Invoke GetClaimRule successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetClaimRuleOptions model
				getClaimRuleOptionsModel := new(iamidentityv1.GetClaimRuleOptions)
				getClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				getClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				getClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetClaimRule(getClaimRuleOptionsModel)
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
	Describe(`UpdateClaimRule(updateClaimRuleOptions *UpdateClaimRuleOptions) - Operation response error`, func() {
		updateClaimRulePath := "/v1/profiles/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateClaimRulePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateClaimRule with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")

				// Construct an instance of the UpdateClaimRuleOptions model
				updateClaimRuleOptionsModel := new(iamidentityv1.UpdateClaimRuleOptions)
				updateClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				updateClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateClaimRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Type = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				updateClaimRuleOptionsModel.Context = responseContextModel
				updateClaimRuleOptionsModel.Name = core.StringPtr("testString")
				updateClaimRuleOptionsModel.RealmName = core.StringPtr("testString")
				updateClaimRuleOptionsModel.CrType = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				updateClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateClaimRule(updateClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateClaimRule(updateClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateClaimRule(updateClaimRuleOptions *UpdateClaimRuleOptions)`, func() {
		updateClaimRulePath := "/v1/profiles/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateClaimRulePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "type": "Type", "realm_name": "RealmName", "expiration": 10, "cr_type": "CrType", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}`)
				}))
			})
			It(`Invoke UpdateClaimRule successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")

				// Construct an instance of the UpdateClaimRuleOptions model
				updateClaimRuleOptionsModel := new(iamidentityv1.UpdateClaimRuleOptions)
				updateClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				updateClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateClaimRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Type = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				updateClaimRuleOptionsModel.Context = responseContextModel
				updateClaimRuleOptionsModel.Name = core.StringPtr("testString")
				updateClaimRuleOptionsModel.RealmName = core.StringPtr("testString")
				updateClaimRuleOptionsModel.CrType = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				updateClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.UpdateClaimRuleWithContext(ctx, updateClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.UpdateClaimRule(updateClaimRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.UpdateClaimRuleWithContext(ctx, updateClaimRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateClaimRulePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "type": "Type", "realm_name": "RealmName", "expiration": 10, "cr_type": "CrType", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}`)
				}))
			})
			It(`Invoke UpdateClaimRule successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateClaimRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")

				// Construct an instance of the UpdateClaimRuleOptions model
				updateClaimRuleOptionsModel := new(iamidentityv1.UpdateClaimRuleOptions)
				updateClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				updateClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateClaimRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Type = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				updateClaimRuleOptionsModel.Context = responseContextModel
				updateClaimRuleOptionsModel.Name = core.StringPtr("testString")
				updateClaimRuleOptionsModel.RealmName = core.StringPtr("testString")
				updateClaimRuleOptionsModel.CrType = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				updateClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateClaimRule(updateClaimRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateClaimRule with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")

				// Construct an instance of the UpdateClaimRuleOptions model
				updateClaimRuleOptionsModel := new(iamidentityv1.UpdateClaimRuleOptions)
				updateClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				updateClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateClaimRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Type = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				updateClaimRuleOptionsModel.Context = responseContextModel
				updateClaimRuleOptionsModel.Name = core.StringPtr("testString")
				updateClaimRuleOptionsModel.RealmName = core.StringPtr("testString")
				updateClaimRuleOptionsModel.CrType = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				updateClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateClaimRule(updateClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateClaimRuleOptions model with no property values
				updateClaimRuleOptionsModelNew := new(iamidentityv1.UpdateClaimRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateClaimRule(updateClaimRuleOptionsModelNew)
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
			It(`Invoke UpdateClaimRule successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")

				// Construct an instance of the UpdateClaimRuleOptions model
				updateClaimRuleOptionsModel := new(iamidentityv1.UpdateClaimRuleOptions)
				updateClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				updateClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateClaimRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Type = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				updateClaimRuleOptionsModel.Context = responseContextModel
				updateClaimRuleOptionsModel.Name = core.StringPtr("testString")
				updateClaimRuleOptionsModel.RealmName = core.StringPtr("testString")
				updateClaimRuleOptionsModel.CrType = core.StringPtr("testString")
				updateClaimRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				updateClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.UpdateClaimRule(updateClaimRuleOptionsModel)
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
	Describe(`DeleteClaimRule(deleteClaimRuleOptions *DeleteClaimRuleOptions)`, func() {
		deleteClaimRulePath := "/v1/profiles/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteClaimRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteClaimRule successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteClaimRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteClaimRuleOptions model
				deleteClaimRuleOptionsModel := new(iamidentityv1.DeleteClaimRuleOptions)
				deleteClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				deleteClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteClaimRule(deleteClaimRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteClaimRule with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteClaimRuleOptions model
				deleteClaimRuleOptionsModel := new(iamidentityv1.DeleteClaimRuleOptions)
				deleteClaimRuleOptionsModel.ProfileID = core.StringPtr("testString")
				deleteClaimRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteClaimRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteClaimRule(deleteClaimRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteClaimRuleOptions model with no property values
				deleteClaimRuleOptionsModelNew := new(iamidentityv1.DeleteClaimRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteClaimRule(deleteClaimRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLink(createLinkOptions *CreateLinkOptions) - Operation response error`, func() {
		createLinkPath := "/v1/profiles/testString/links"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLinkPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLink with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateProfileLinkRequestLink model
				createProfileLinkRequestLinkModel := new(iamidentityv1.CreateProfileLinkRequestLink)
				createProfileLinkRequestLinkModel.CRN = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Namespace = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Name = core.StringPtr("testString")

				// Construct an instance of the CreateLinkOptions model
				createLinkOptionsModel := new(iamidentityv1.CreateLinkOptions)
				createLinkOptionsModel.ProfileID = core.StringPtr("testString")
				createLinkOptionsModel.CrType = core.StringPtr("testString")
				createLinkOptionsModel.Link = createProfileLinkRequestLinkModel
				createLinkOptionsModel.Name = core.StringPtr("testString")
				createLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateLink(createLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateLink(createLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLink(createLinkOptions *CreateLinkOptions)`, func() {
		createLinkPath := "/v1/profiles/testString/links"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLinkPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "cr_type": "CrType", "link": {"crn": "CRN", "namespace": "Namespace", "name": "Name"}}`)
				}))
			})
			It(`Invoke CreateLink successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the CreateProfileLinkRequestLink model
				createProfileLinkRequestLinkModel := new(iamidentityv1.CreateProfileLinkRequestLink)
				createProfileLinkRequestLinkModel.CRN = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Namespace = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Name = core.StringPtr("testString")

				// Construct an instance of the CreateLinkOptions model
				createLinkOptionsModel := new(iamidentityv1.CreateLinkOptions)
				createLinkOptionsModel.ProfileID = core.StringPtr("testString")
				createLinkOptionsModel.CrType = core.StringPtr("testString")
				createLinkOptionsModel.Link = createProfileLinkRequestLinkModel
				createLinkOptionsModel.Name = core.StringPtr("testString")
				createLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateLinkWithContext(ctx, createLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateLink(createLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateLinkWithContext(ctx, createLinkOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createLinkPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "cr_type": "CrType", "link": {"crn": "CRN", "namespace": "Namespace", "name": "Name"}}`)
				}))
			})
			It(`Invoke CreateLink successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateLink(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateProfileLinkRequestLink model
				createProfileLinkRequestLinkModel := new(iamidentityv1.CreateProfileLinkRequestLink)
				createProfileLinkRequestLinkModel.CRN = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Namespace = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Name = core.StringPtr("testString")

				// Construct an instance of the CreateLinkOptions model
				createLinkOptionsModel := new(iamidentityv1.CreateLinkOptions)
				createLinkOptionsModel.ProfileID = core.StringPtr("testString")
				createLinkOptionsModel.CrType = core.StringPtr("testString")
				createLinkOptionsModel.Link = createProfileLinkRequestLinkModel
				createLinkOptionsModel.Name = core.StringPtr("testString")
				createLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateLink(createLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateLink with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateProfileLinkRequestLink model
				createProfileLinkRequestLinkModel := new(iamidentityv1.CreateProfileLinkRequestLink)
				createProfileLinkRequestLinkModel.CRN = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Namespace = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Name = core.StringPtr("testString")

				// Construct an instance of the CreateLinkOptions model
				createLinkOptionsModel := new(iamidentityv1.CreateLinkOptions)
				createLinkOptionsModel.ProfileID = core.StringPtr("testString")
				createLinkOptionsModel.CrType = core.StringPtr("testString")
				createLinkOptionsModel.Link = createProfileLinkRequestLinkModel
				createLinkOptionsModel.Name = core.StringPtr("testString")
				createLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateLink(createLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateLinkOptions model with no property values
				createLinkOptionsModelNew := new(iamidentityv1.CreateLinkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateLink(createLinkOptionsModelNew)
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
			It(`Invoke CreateLink successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateProfileLinkRequestLink model
				createProfileLinkRequestLinkModel := new(iamidentityv1.CreateProfileLinkRequestLink)
				createProfileLinkRequestLinkModel.CRN = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Namespace = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Name = core.StringPtr("testString")

				// Construct an instance of the CreateLinkOptions model
				createLinkOptionsModel := new(iamidentityv1.CreateLinkOptions)
				createLinkOptionsModel.ProfileID = core.StringPtr("testString")
				createLinkOptionsModel.CrType = core.StringPtr("testString")
				createLinkOptionsModel.Link = createProfileLinkRequestLinkModel
				createLinkOptionsModel.Name = core.StringPtr("testString")
				createLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateLink(createLinkOptionsModel)
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
	Describe(`ListLinks(listLinksOptions *ListLinksOptions) - Operation response error`, func() {
		listLinksPath := "/v1/profiles/testString/links"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLinksPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLinks with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListLinksOptions model
				listLinksOptionsModel := new(iamidentityv1.ListLinksOptions)
				listLinksOptionsModel.ProfileID = core.StringPtr("testString")
				listLinksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListLinks(listLinksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListLinks(listLinksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListLinks(listLinksOptions *ListLinksOptions)`, func() {
		listLinksPath := "/v1/profiles/testString/links"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLinksPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"links": [{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "cr_type": "CrType", "link": {"crn": "CRN", "namespace": "Namespace", "name": "Name"}}]}`)
				}))
			})
			It(`Invoke ListLinks successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListLinksOptions model
				listLinksOptionsModel := new(iamidentityv1.ListLinksOptions)
				listLinksOptionsModel.ProfileID = core.StringPtr("testString")
				listLinksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListLinksWithContext(ctx, listLinksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListLinks(listLinksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListLinksWithContext(ctx, listLinksOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listLinksPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"links": [{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "cr_type": "CrType", "link": {"crn": "CRN", "namespace": "Namespace", "name": "Name"}}]}`)
				}))
			})
			It(`Invoke ListLinks successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListLinks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLinksOptions model
				listLinksOptionsModel := new(iamidentityv1.ListLinksOptions)
				listLinksOptionsModel.ProfileID = core.StringPtr("testString")
				listLinksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListLinks(listLinksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListLinks with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListLinksOptions model
				listLinksOptionsModel := new(iamidentityv1.ListLinksOptions)
				listLinksOptionsModel.ProfileID = core.StringPtr("testString")
				listLinksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListLinks(listLinksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListLinksOptions model with no property values
				listLinksOptionsModelNew := new(iamidentityv1.ListLinksOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.ListLinks(listLinksOptionsModelNew)
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
			It(`Invoke ListLinks successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListLinksOptions model
				listLinksOptionsModel := new(iamidentityv1.ListLinksOptions)
				listLinksOptionsModel.ProfileID = core.StringPtr("testString")
				listLinksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListLinks(listLinksOptionsModel)
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
	Describe(`GetLink(getLinkOptions *GetLinkOptions) - Operation response error`, func() {
		getLinkPath := "/v1/profiles/testString/links/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLinkPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLink with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetLinkOptions model
				getLinkOptionsModel := new(iamidentityv1.GetLinkOptions)
				getLinkOptionsModel.ProfileID = core.StringPtr("testString")
				getLinkOptionsModel.LinkID = core.StringPtr("testString")
				getLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetLink(getLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetLink(getLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLink(getLinkOptions *GetLinkOptions)`, func() {
		getLinkPath := "/v1/profiles/testString/links/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLinkPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "cr_type": "CrType", "link": {"crn": "CRN", "namespace": "Namespace", "name": "Name"}}`)
				}))
			})
			It(`Invoke GetLink successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetLinkOptions model
				getLinkOptionsModel := new(iamidentityv1.GetLinkOptions)
				getLinkOptionsModel.ProfileID = core.StringPtr("testString")
				getLinkOptionsModel.LinkID = core.StringPtr("testString")
				getLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetLinkWithContext(ctx, getLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetLink(getLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetLinkWithContext(ctx, getLinkOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLinkPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_tag": "EntityTag", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "cr_type": "CrType", "link": {"crn": "CRN", "namespace": "Namespace", "name": "Name"}}`)
				}))
			})
			It(`Invoke GetLink successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetLink(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLinkOptions model
				getLinkOptionsModel := new(iamidentityv1.GetLinkOptions)
				getLinkOptionsModel.ProfileID = core.StringPtr("testString")
				getLinkOptionsModel.LinkID = core.StringPtr("testString")
				getLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetLink(getLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLink with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetLinkOptions model
				getLinkOptionsModel := new(iamidentityv1.GetLinkOptions)
				getLinkOptionsModel.ProfileID = core.StringPtr("testString")
				getLinkOptionsModel.LinkID = core.StringPtr("testString")
				getLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetLink(getLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLinkOptions model with no property values
				getLinkOptionsModelNew := new(iamidentityv1.GetLinkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetLink(getLinkOptionsModelNew)
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
			It(`Invoke GetLink successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetLinkOptions model
				getLinkOptionsModel := new(iamidentityv1.GetLinkOptions)
				getLinkOptionsModel.ProfileID = core.StringPtr("testString")
				getLinkOptionsModel.LinkID = core.StringPtr("testString")
				getLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetLink(getLinkOptionsModel)
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
	Describe(`DeleteLink(deleteLinkOptions *DeleteLinkOptions)`, func() {
		deleteLinkPath := "/v1/profiles/testString/links/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLinkPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteLink successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteLink(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLinkOptions model
				deleteLinkOptionsModel := new(iamidentityv1.DeleteLinkOptions)
				deleteLinkOptionsModel.ProfileID = core.StringPtr("testString")
				deleteLinkOptionsModel.LinkID = core.StringPtr("testString")
				deleteLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteLink(deleteLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteLink with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteLinkOptions model
				deleteLinkOptionsModel := new(iamidentityv1.DeleteLinkOptions)
				deleteLinkOptionsModel.ProfileID = core.StringPtr("testString")
				deleteLinkOptionsModel.LinkID = core.StringPtr("testString")
				deleteLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteLink(deleteLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteLinkOptions model with no property values
				deleteLinkOptionsModelNew := new(iamidentityv1.DeleteLinkOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteLink(deleteLinkOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfileIdentities(getProfileIdentitiesOptions *GetProfileIdentitiesOptions) - Operation response error`, func() {
		getProfileIdentitiesPath := "/v1/profiles/testString/identities"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileIdentitiesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfileIdentities with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileIdentitiesOptions model
				getProfileIdentitiesOptionsModel := new(iamidentityv1.GetProfileIdentitiesOptions)
				getProfileIdentitiesOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileIdentitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetProfileIdentities(getProfileIdentitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetProfileIdentities(getProfileIdentitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfileIdentities(getProfileIdentitiesOptions *GetProfileIdentitiesOptions)`, func() {
		getProfileIdentitiesPath := "/v1/profiles/testString/identities"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileIdentitiesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"entity_tag": "EntityTag", "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}`)
				}))
			})
			It(`Invoke GetProfileIdentities successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileIdentitiesOptions model
				getProfileIdentitiesOptionsModel := new(iamidentityv1.GetProfileIdentitiesOptions)
				getProfileIdentitiesOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileIdentitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetProfileIdentitiesWithContext(ctx, getProfileIdentitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetProfileIdentities(getProfileIdentitiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetProfileIdentitiesWithContext(ctx, getProfileIdentitiesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProfileIdentitiesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"entity_tag": "EntityTag", "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}`)
				}))
			})
			It(`Invoke GetProfileIdentities successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetProfileIdentities(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileIdentitiesOptions model
				getProfileIdentitiesOptionsModel := new(iamidentityv1.GetProfileIdentitiesOptions)
				getProfileIdentitiesOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileIdentitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetProfileIdentities(getProfileIdentitiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfileIdentities with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileIdentitiesOptions model
				getProfileIdentitiesOptionsModel := new(iamidentityv1.GetProfileIdentitiesOptions)
				getProfileIdentitiesOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileIdentitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetProfileIdentities(getProfileIdentitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileIdentitiesOptions model with no property values
				getProfileIdentitiesOptionsModelNew := new(iamidentityv1.GetProfileIdentitiesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetProfileIdentities(getProfileIdentitiesOptionsModelNew)
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
			It(`Invoke GetProfileIdentities successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileIdentitiesOptions model
				getProfileIdentitiesOptionsModel := new(iamidentityv1.GetProfileIdentitiesOptions)
				getProfileIdentitiesOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileIdentitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetProfileIdentities(getProfileIdentitiesOptionsModel)
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
	Describe(`SetProfileIdentities(setProfileIdentitiesOptions *SetProfileIdentitiesOptions) - Operation response error`, func() {
		setProfileIdentitiesPath := "/v1/profiles/testString/identities"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setProfileIdentitiesPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetProfileIdentities with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the SetProfileIdentitiesOptions model
				setProfileIdentitiesOptionsModel := new(iamidentityv1.SetProfileIdentitiesOptions)
				setProfileIdentitiesOptionsModel.ProfileID = core.StringPtr("testString")
				setProfileIdentitiesOptionsModel.IfMatch = core.StringPtr("testString")
				setProfileIdentitiesOptionsModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}
				setProfileIdentitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.SetProfileIdentities(setProfileIdentitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.SetProfileIdentities(setProfileIdentitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetProfileIdentities(setProfileIdentitiesOptions *SetProfileIdentitiesOptions)`, func() {
		setProfileIdentitiesPath := "/v1/profiles/testString/identities"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setProfileIdentitiesPath))
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
					fmt.Fprintf(res, "%s", `{"entity_tag": "EntityTag", "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}`)
				}))
			})
			It(`Invoke SetProfileIdentities successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the SetProfileIdentitiesOptions model
				setProfileIdentitiesOptionsModel := new(iamidentityv1.SetProfileIdentitiesOptions)
				setProfileIdentitiesOptionsModel.ProfileID = core.StringPtr("testString")
				setProfileIdentitiesOptionsModel.IfMatch = core.StringPtr("testString")
				setProfileIdentitiesOptionsModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}
				setProfileIdentitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.SetProfileIdentitiesWithContext(ctx, setProfileIdentitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.SetProfileIdentities(setProfileIdentitiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.SetProfileIdentitiesWithContext(ctx, setProfileIdentitiesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setProfileIdentitiesPath))
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
					fmt.Fprintf(res, "%s", `{"entity_tag": "EntityTag", "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}`)
				}))
			})
			It(`Invoke SetProfileIdentities successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.SetProfileIdentities(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the SetProfileIdentitiesOptions model
				setProfileIdentitiesOptionsModel := new(iamidentityv1.SetProfileIdentitiesOptions)
				setProfileIdentitiesOptionsModel.ProfileID = core.StringPtr("testString")
				setProfileIdentitiesOptionsModel.IfMatch = core.StringPtr("testString")
				setProfileIdentitiesOptionsModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}
				setProfileIdentitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.SetProfileIdentities(setProfileIdentitiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetProfileIdentities with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the SetProfileIdentitiesOptions model
				setProfileIdentitiesOptionsModel := new(iamidentityv1.SetProfileIdentitiesOptions)
				setProfileIdentitiesOptionsModel.ProfileID = core.StringPtr("testString")
				setProfileIdentitiesOptionsModel.IfMatch = core.StringPtr("testString")
				setProfileIdentitiesOptionsModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}
				setProfileIdentitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.SetProfileIdentities(setProfileIdentitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetProfileIdentitiesOptions model with no property values
				setProfileIdentitiesOptionsModelNew := new(iamidentityv1.SetProfileIdentitiesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.SetProfileIdentities(setProfileIdentitiesOptionsModelNew)
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
			It(`Invoke SetProfileIdentities successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the SetProfileIdentitiesOptions model
				setProfileIdentitiesOptionsModel := new(iamidentityv1.SetProfileIdentitiesOptions)
				setProfileIdentitiesOptionsModel.ProfileID = core.StringPtr("testString")
				setProfileIdentitiesOptionsModel.IfMatch = core.StringPtr("testString")
				setProfileIdentitiesOptionsModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}
				setProfileIdentitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.SetProfileIdentities(setProfileIdentitiesOptionsModel)
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
	Describe(`SetProfileIdentity(setProfileIdentityOptions *SetProfileIdentityOptions) - Operation response error`, func() {
		setProfileIdentityPath := "/v1/profiles/testString/identities/user"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setProfileIdentityPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetProfileIdentity with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the SetProfileIdentityOptions model
				setProfileIdentityOptionsModel := new(iamidentityv1.SetProfileIdentityOptions)
				setProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				setProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				setProfileIdentityOptionsModel.Identifier = core.StringPtr("testString")
				setProfileIdentityOptionsModel.Type = core.StringPtr("user")
				setProfileIdentityOptionsModel.Accounts = []string{"testString"}
				setProfileIdentityOptionsModel.Description = core.StringPtr("testString")
				setProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.SetProfileIdentity(setProfileIdentityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.SetProfileIdentity(setProfileIdentityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetProfileIdentity(setProfileIdentityOptions *SetProfileIdentityOptions)`, func() {
		setProfileIdentityPath := "/v1/profiles/testString/identities/user"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setProfileIdentityPath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}`)
				}))
			})
			It(`Invoke SetProfileIdentity successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the SetProfileIdentityOptions model
				setProfileIdentityOptionsModel := new(iamidentityv1.SetProfileIdentityOptions)
				setProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				setProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				setProfileIdentityOptionsModel.Identifier = core.StringPtr("testString")
				setProfileIdentityOptionsModel.Type = core.StringPtr("user")
				setProfileIdentityOptionsModel.Accounts = []string{"testString"}
				setProfileIdentityOptionsModel.Description = core.StringPtr("testString")
				setProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.SetProfileIdentityWithContext(ctx, setProfileIdentityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.SetProfileIdentity(setProfileIdentityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.SetProfileIdentityWithContext(ctx, setProfileIdentityOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setProfileIdentityPath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}`)
				}))
			})
			It(`Invoke SetProfileIdentity successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.SetProfileIdentity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetProfileIdentityOptions model
				setProfileIdentityOptionsModel := new(iamidentityv1.SetProfileIdentityOptions)
				setProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				setProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				setProfileIdentityOptionsModel.Identifier = core.StringPtr("testString")
				setProfileIdentityOptionsModel.Type = core.StringPtr("user")
				setProfileIdentityOptionsModel.Accounts = []string{"testString"}
				setProfileIdentityOptionsModel.Description = core.StringPtr("testString")
				setProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.SetProfileIdentity(setProfileIdentityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetProfileIdentity with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the SetProfileIdentityOptions model
				setProfileIdentityOptionsModel := new(iamidentityv1.SetProfileIdentityOptions)
				setProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				setProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				setProfileIdentityOptionsModel.Identifier = core.StringPtr("testString")
				setProfileIdentityOptionsModel.Type = core.StringPtr("user")
				setProfileIdentityOptionsModel.Accounts = []string{"testString"}
				setProfileIdentityOptionsModel.Description = core.StringPtr("testString")
				setProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.SetProfileIdentity(setProfileIdentityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetProfileIdentityOptions model with no property values
				setProfileIdentityOptionsModelNew := new(iamidentityv1.SetProfileIdentityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.SetProfileIdentity(setProfileIdentityOptionsModelNew)
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
			It(`Invoke SetProfileIdentity successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the SetProfileIdentityOptions model
				setProfileIdentityOptionsModel := new(iamidentityv1.SetProfileIdentityOptions)
				setProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				setProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				setProfileIdentityOptionsModel.Identifier = core.StringPtr("testString")
				setProfileIdentityOptionsModel.Type = core.StringPtr("user")
				setProfileIdentityOptionsModel.Accounts = []string{"testString"}
				setProfileIdentityOptionsModel.Description = core.StringPtr("testString")
				setProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.SetProfileIdentity(setProfileIdentityOptionsModel)
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
	Describe(`GetProfileIdentity(getProfileIdentityOptions *GetProfileIdentityOptions) - Operation response error`, func() {
		getProfileIdentityPath := "/v1/profiles/testString/identities/user/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileIdentityPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfileIdentity with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileIdentityOptions model
				getProfileIdentityOptionsModel := new(iamidentityv1.GetProfileIdentityOptions)
				getProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				getProfileIdentityOptionsModel.IdentifierID = core.StringPtr("testString")
				getProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetProfileIdentity(getProfileIdentityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetProfileIdentity(getProfileIdentityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfileIdentity(getProfileIdentityOptions *GetProfileIdentityOptions)`, func() {
		getProfileIdentityPath := "/v1/profiles/testString/identities/user/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileIdentityPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}`)
				}))
			})
			It(`Invoke GetProfileIdentity successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileIdentityOptions model
				getProfileIdentityOptionsModel := new(iamidentityv1.GetProfileIdentityOptions)
				getProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				getProfileIdentityOptionsModel.IdentifierID = core.StringPtr("testString")
				getProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetProfileIdentityWithContext(ctx, getProfileIdentityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetProfileIdentity(getProfileIdentityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetProfileIdentityWithContext(ctx, getProfileIdentityOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProfileIdentityPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}`)
				}))
			})
			It(`Invoke GetProfileIdentity successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetProfileIdentity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileIdentityOptions model
				getProfileIdentityOptionsModel := new(iamidentityv1.GetProfileIdentityOptions)
				getProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				getProfileIdentityOptionsModel.IdentifierID = core.StringPtr("testString")
				getProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetProfileIdentity(getProfileIdentityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfileIdentity with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileIdentityOptions model
				getProfileIdentityOptionsModel := new(iamidentityv1.GetProfileIdentityOptions)
				getProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				getProfileIdentityOptionsModel.IdentifierID = core.StringPtr("testString")
				getProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetProfileIdentity(getProfileIdentityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileIdentityOptions model with no property values
				getProfileIdentityOptionsModelNew := new(iamidentityv1.GetProfileIdentityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetProfileIdentity(getProfileIdentityOptionsModelNew)
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
			It(`Invoke GetProfileIdentity successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileIdentityOptions model
				getProfileIdentityOptionsModel := new(iamidentityv1.GetProfileIdentityOptions)
				getProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				getProfileIdentityOptionsModel.IdentifierID = core.StringPtr("testString")
				getProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetProfileIdentity(getProfileIdentityOptionsModel)
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
	Describe(`DeleteProfileIdentity(deleteProfileIdentityOptions *DeleteProfileIdentityOptions)`, func() {
		deleteProfileIdentityPath := "/v1/profiles/testString/identities/user/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProfileIdentityPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteProfileIdentity successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteProfileIdentity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProfileIdentityOptions model
				deleteProfileIdentityOptionsModel := new(iamidentityv1.DeleteProfileIdentityOptions)
				deleteProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				deleteProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				deleteProfileIdentityOptionsModel.IdentifierID = core.StringPtr("testString")
				deleteProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteProfileIdentity(deleteProfileIdentityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProfileIdentity with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileIdentityOptions model
				deleteProfileIdentityOptionsModel := new(iamidentityv1.DeleteProfileIdentityOptions)
				deleteProfileIdentityOptionsModel.ProfileID = core.StringPtr("testString")
				deleteProfileIdentityOptionsModel.IdentityType = core.StringPtr("user")
				deleteProfileIdentityOptionsModel.IdentifierID = core.StringPtr("testString")
				deleteProfileIdentityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteProfileIdentity(deleteProfileIdentityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProfileIdentityOptions model with no property values
				deleteProfileIdentityOptionsModelNew := new(iamidentityv1.DeleteProfileIdentityOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteProfileIdentity(deleteProfileIdentityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) - Operation response error`, func() {
		getAccountSettingsPath := "/v1/accounts/testString/settings/identity"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountSettings with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamidentityv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetAccountSettings(getAccountSettingsOptionsModel)
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
		getAccountSettingsPath := "/v1/accounts/testString/settings/identity"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamidentityv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
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

					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamidentityv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountSettings with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamidentityv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountSettingsOptions model with no property values
				getAccountSettingsOptionsModelNew := new(iamidentityv1.GetAccountSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetAccountSettings(getAccountSettingsOptionsModelNew)
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
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamidentityv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetAccountSettings(getAccountSettingsOptionsModel)
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
		updateAccountSettingsPath := "/v1/accounts/testString/settings/identity"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountSettings with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamidentityv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.RestrictCreateServiceID = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Mfa = core.StringPtr("NONE")
				updateAccountSettingsOptionsModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("86400")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				updateAccountSettingsOptionsModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				updateAccountSettingsOptionsModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
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
		updateAccountSettingsPath := "/v1/accounts/testString/settings/identity"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsPath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}`)
				}))
			})
			It(`Invoke UpdateAccountSettings successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamidentityv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.RestrictCreateServiceID = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Mfa = core.StringPtr("NONE")
				updateAccountSettingsOptionsModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("86400")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				updateAccountSettingsOptionsModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				updateAccountSettingsOptionsModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.UpdateAccountSettingsWithContext(ctx, updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.UpdateAccountSettingsWithContext(ctx, updateAccountSettingsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}`)
				}))
			})
			It(`Invoke UpdateAccountSettings successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamidentityv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.RestrictCreateServiceID = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Mfa = core.StringPtr("NONE")
				updateAccountSettingsOptionsModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("86400")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				updateAccountSettingsOptionsModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				updateAccountSettingsOptionsModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccountSettings with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamidentityv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.RestrictCreateServiceID = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Mfa = core.StringPtr("NONE")
				updateAccountSettingsOptionsModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("86400")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				updateAccountSettingsOptionsModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				updateAccountSettingsOptionsModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountSettingsOptions model with no property values
				updateAccountSettingsOptionsModelNew := new(iamidentityv1.UpdateAccountSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateAccountSettings(updateAccountSettingsOptionsModelNew)
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
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamidentityv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.RestrictCreateServiceID = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Mfa = core.StringPtr("NONE")
				updateAccountSettingsOptionsModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("86400")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				updateAccountSettingsOptionsModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				updateAccountSettingsOptionsModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
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
	Describe(`GetMfaStatus(getMfaStatusOptions *GetMfaStatusOptions) - Operation response error`, func() {
		getMfaStatusPath := "/v1/mfa/accounts/testString/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMfaStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMfaStatus with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetMfaStatusOptions model
				getMfaStatusOptionsModel := new(iamidentityv1.GetMfaStatusOptions)
				getMfaStatusOptionsModel.AccountID = core.StringPtr("testString")
				getMfaStatusOptionsModel.IamID = core.StringPtr("testString")
				getMfaStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetMfaStatus(getMfaStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetMfaStatus(getMfaStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMfaStatus(getMfaStatusOptions *GetMfaStatusOptions)`, func() {
		getMfaStatusPath := "/v1/mfa/accounts/testString/status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMfaStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "effective_mfa_type": "EffectiveMfaType", "id_based_mfa": {"trait_account_default": "NONE", "trait_user_specific": "NONE", "trait_effective": "NONE", "complies": true, "comply_state": "false"}, "account_based_mfa": {"security_questions": {"required": true, "enrolled": true}, "totp": {"required": true, "enrolled": true}, "verisign": {"required": true, "enrolled": true}, "complies": true}}`)
				}))
			})
			It(`Invoke GetMfaStatus successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetMfaStatusOptions model
				getMfaStatusOptionsModel := new(iamidentityv1.GetMfaStatusOptions)
				getMfaStatusOptionsModel.AccountID = core.StringPtr("testString")
				getMfaStatusOptionsModel.IamID = core.StringPtr("testString")
				getMfaStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetMfaStatusWithContext(ctx, getMfaStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetMfaStatus(getMfaStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetMfaStatusWithContext(ctx, getMfaStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getMfaStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "effective_mfa_type": "EffectiveMfaType", "id_based_mfa": {"trait_account_default": "NONE", "trait_user_specific": "NONE", "trait_effective": "NONE", "complies": true, "comply_state": "false"}, "account_based_mfa": {"security_questions": {"required": true, "enrolled": true}, "totp": {"required": true, "enrolled": true}, "verisign": {"required": true, "enrolled": true}, "complies": true}}`)
				}))
			})
			It(`Invoke GetMfaStatus successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetMfaStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMfaStatusOptions model
				getMfaStatusOptionsModel := new(iamidentityv1.GetMfaStatusOptions)
				getMfaStatusOptionsModel.AccountID = core.StringPtr("testString")
				getMfaStatusOptionsModel.IamID = core.StringPtr("testString")
				getMfaStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetMfaStatus(getMfaStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMfaStatus with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetMfaStatusOptions model
				getMfaStatusOptionsModel := new(iamidentityv1.GetMfaStatusOptions)
				getMfaStatusOptionsModel.AccountID = core.StringPtr("testString")
				getMfaStatusOptionsModel.IamID = core.StringPtr("testString")
				getMfaStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetMfaStatus(getMfaStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetMfaStatusOptions model with no property values
				getMfaStatusOptionsModelNew := new(iamidentityv1.GetMfaStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetMfaStatus(getMfaStatusOptionsModelNew)
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
			It(`Invoke GetMfaStatus successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetMfaStatusOptions model
				getMfaStatusOptionsModel := new(iamidentityv1.GetMfaStatusOptions)
				getMfaStatusOptionsModel.AccountID = core.StringPtr("testString")
				getMfaStatusOptionsModel.IamID = core.StringPtr("testString")
				getMfaStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetMfaStatus(getMfaStatusOptionsModel)
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
	Describe(`CreateMfaReport(createMfaReportOptions *CreateMfaReportOptions) - Operation response error`, func() {
		createMfaReportPath := "/v1/mfa/accounts/testString/report"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createMfaReportPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateMfaReport with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateMfaReportOptions model
				createMfaReportOptionsModel := new(iamidentityv1.CreateMfaReportOptions)
				createMfaReportOptionsModel.AccountID = core.StringPtr("testString")
				createMfaReportOptionsModel.Type = core.StringPtr("testString")
				createMfaReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateMfaReport(createMfaReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateMfaReport(createMfaReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateMfaReport(createMfaReportOptions *CreateMfaReportOptions)`, func() {
		createMfaReportPath := "/v1/mfa/accounts/testString/report"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createMfaReportPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"reference": "Reference"}`)
				}))
			})
			It(`Invoke CreateMfaReport successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the CreateMfaReportOptions model
				createMfaReportOptionsModel := new(iamidentityv1.CreateMfaReportOptions)
				createMfaReportOptionsModel.AccountID = core.StringPtr("testString")
				createMfaReportOptionsModel.Type = core.StringPtr("testString")
				createMfaReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateMfaReportWithContext(ctx, createMfaReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateMfaReport(createMfaReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateMfaReportWithContext(ctx, createMfaReportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createMfaReportPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"reference": "Reference"}`)
				}))
			})
			It(`Invoke CreateMfaReport successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateMfaReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateMfaReportOptions model
				createMfaReportOptionsModel := new(iamidentityv1.CreateMfaReportOptions)
				createMfaReportOptionsModel.AccountID = core.StringPtr("testString")
				createMfaReportOptionsModel.Type = core.StringPtr("testString")
				createMfaReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateMfaReport(createMfaReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateMfaReport with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateMfaReportOptions model
				createMfaReportOptionsModel := new(iamidentityv1.CreateMfaReportOptions)
				createMfaReportOptionsModel.AccountID = core.StringPtr("testString")
				createMfaReportOptionsModel.Type = core.StringPtr("testString")
				createMfaReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateMfaReport(createMfaReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateMfaReportOptions model with no property values
				createMfaReportOptionsModelNew := new(iamidentityv1.CreateMfaReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateMfaReport(createMfaReportOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateMfaReport successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateMfaReportOptions model
				createMfaReportOptionsModel := new(iamidentityv1.CreateMfaReportOptions)
				createMfaReportOptionsModel.AccountID = core.StringPtr("testString")
				createMfaReportOptionsModel.Type = core.StringPtr("testString")
				createMfaReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateMfaReport(createMfaReportOptionsModel)
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
	Describe(`GetMfaReport(getMfaReportOptions *GetMfaReportOptions) - Operation response error`, func() {
		getMfaReportPath := "/v1/mfa/accounts/testString/report/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMfaReportPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMfaReport with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetMfaReportOptions model
				getMfaReportOptionsModel := new(iamidentityv1.GetMfaReportOptions)
				getMfaReportOptionsModel.AccountID = core.StringPtr("testString")
				getMfaReportOptionsModel.Reference = core.StringPtr("testString")
				getMfaReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetMfaReport(getMfaReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetMfaReport(getMfaReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMfaReport(getMfaReportOptions *GetMfaReportOptions)`, func() {
		getMfaReportPath := "/v1/mfa/accounts/testString/report/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMfaReportPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_by": "CreatedBy", "reference": "Reference", "report_time": "ReportTime", "account_id": "AccountID", "ims_account_id": "ImsAccountID", "users": [{"iam_id": "IamID", "name": "Name", "username": "Username", "email": "Email", "enrollments": {"effective_mfa_type": "EffectiveMfaType", "id_based_mfa": {"trait_account_default": "NONE", "trait_user_specific": "NONE", "trait_effective": "NONE", "complies": true, "comply_state": "false"}, "account_based_mfa": {"security_questions": {"required": true, "enrolled": true}, "totp": {"required": true, "enrolled": true}, "verisign": {"required": true, "enrolled": true}, "complies": true}}}]}`)
				}))
			})
			It(`Invoke GetMfaReport successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetMfaReportOptions model
				getMfaReportOptionsModel := new(iamidentityv1.GetMfaReportOptions)
				getMfaReportOptionsModel.AccountID = core.StringPtr("testString")
				getMfaReportOptionsModel.Reference = core.StringPtr("testString")
				getMfaReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetMfaReportWithContext(ctx, getMfaReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetMfaReport(getMfaReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetMfaReportWithContext(ctx, getMfaReportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getMfaReportPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_by": "CreatedBy", "reference": "Reference", "report_time": "ReportTime", "account_id": "AccountID", "ims_account_id": "ImsAccountID", "users": [{"iam_id": "IamID", "name": "Name", "username": "Username", "email": "Email", "enrollments": {"effective_mfa_type": "EffectiveMfaType", "id_based_mfa": {"trait_account_default": "NONE", "trait_user_specific": "NONE", "trait_effective": "NONE", "complies": true, "comply_state": "false"}, "account_based_mfa": {"security_questions": {"required": true, "enrolled": true}, "totp": {"required": true, "enrolled": true}, "verisign": {"required": true, "enrolled": true}, "complies": true}}}]}`)
				}))
			})
			It(`Invoke GetMfaReport successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetMfaReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMfaReportOptions model
				getMfaReportOptionsModel := new(iamidentityv1.GetMfaReportOptions)
				getMfaReportOptionsModel.AccountID = core.StringPtr("testString")
				getMfaReportOptionsModel.Reference = core.StringPtr("testString")
				getMfaReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetMfaReport(getMfaReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMfaReport with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetMfaReportOptions model
				getMfaReportOptionsModel := new(iamidentityv1.GetMfaReportOptions)
				getMfaReportOptionsModel.AccountID = core.StringPtr("testString")
				getMfaReportOptionsModel.Reference = core.StringPtr("testString")
				getMfaReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetMfaReport(getMfaReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetMfaReportOptions model with no property values
				getMfaReportOptionsModelNew := new(iamidentityv1.GetMfaReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetMfaReport(getMfaReportOptionsModelNew)
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
			It(`Invoke GetMfaReport successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetMfaReportOptions model
				getMfaReportOptionsModel := new(iamidentityv1.GetMfaReportOptions)
				getMfaReportOptionsModel.AccountID = core.StringPtr("testString")
				getMfaReportOptionsModel.Reference = core.StringPtr("testString")
				getMfaReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetMfaReport(getMfaReportOptionsModel)
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
	Describe(`ListAccountSettingsAssignments(listAccountSettingsAssignmentsOptions *ListAccountSettingsAssignmentsOptions) - Operation response error`, func() {
		listAccountSettingsAssignmentsPath := "/v1/account_settings_assignments/"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccountSettingsAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_type"]).To(Equal([]string{"Account"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccountSettingsAssignments with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListAccountSettingsAssignmentsOptions model
				listAccountSettingsAssignmentsOptionsModel := new(iamidentityv1.ListAccountSettingsAssignmentsOptions)
				listAccountSettingsAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TargetType = core.StringPtr("Account")
				listAccountSettingsAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listAccountSettingsAssignmentsOptionsModel.Pagetoken = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.Sort = core.StringPtr("created_at")
				listAccountSettingsAssignmentsOptionsModel.Order = core.StringPtr("asc")
				listAccountSettingsAssignmentsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listAccountSettingsAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListAccountSettingsAssignments(listAccountSettingsAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListAccountSettingsAssignments(listAccountSettingsAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccountSettingsAssignments(listAccountSettingsAssignmentsOptions *ListAccountSettingsAssignmentsOptions)`, func() {
		listAccountSettingsAssignmentsPath := "/v1/account_settings_assignments/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccountSettingsAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_type"]).To(Equal([]string{"Account"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "assignments": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}]}`)
				}))
			})
			It(`Invoke ListAccountSettingsAssignments successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListAccountSettingsAssignmentsOptions model
				listAccountSettingsAssignmentsOptionsModel := new(iamidentityv1.ListAccountSettingsAssignmentsOptions)
				listAccountSettingsAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TargetType = core.StringPtr("Account")
				listAccountSettingsAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listAccountSettingsAssignmentsOptionsModel.Pagetoken = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.Sort = core.StringPtr("created_at")
				listAccountSettingsAssignmentsOptionsModel.Order = core.StringPtr("asc")
				listAccountSettingsAssignmentsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listAccountSettingsAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListAccountSettingsAssignmentsWithContext(ctx, listAccountSettingsAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListAccountSettingsAssignments(listAccountSettingsAssignmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListAccountSettingsAssignmentsWithContext(ctx, listAccountSettingsAssignmentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAccountSettingsAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_type"]).To(Equal([]string{"Account"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "assignments": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}]}`)
				}))
			})
			It(`Invoke ListAccountSettingsAssignments successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListAccountSettingsAssignments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccountSettingsAssignmentsOptions model
				listAccountSettingsAssignmentsOptionsModel := new(iamidentityv1.ListAccountSettingsAssignmentsOptions)
				listAccountSettingsAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TargetType = core.StringPtr("Account")
				listAccountSettingsAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listAccountSettingsAssignmentsOptionsModel.Pagetoken = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.Sort = core.StringPtr("created_at")
				listAccountSettingsAssignmentsOptionsModel.Order = core.StringPtr("asc")
				listAccountSettingsAssignmentsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listAccountSettingsAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListAccountSettingsAssignments(listAccountSettingsAssignmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAccountSettingsAssignments with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListAccountSettingsAssignmentsOptions model
				listAccountSettingsAssignmentsOptionsModel := new(iamidentityv1.ListAccountSettingsAssignmentsOptions)
				listAccountSettingsAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TargetType = core.StringPtr("Account")
				listAccountSettingsAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listAccountSettingsAssignmentsOptionsModel.Pagetoken = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.Sort = core.StringPtr("created_at")
				listAccountSettingsAssignmentsOptionsModel.Order = core.StringPtr("asc")
				listAccountSettingsAssignmentsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listAccountSettingsAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListAccountSettingsAssignments(listAccountSettingsAssignmentsOptionsModel)
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
			It(`Invoke ListAccountSettingsAssignments successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListAccountSettingsAssignmentsOptions model
				listAccountSettingsAssignmentsOptionsModel := new(iamidentityv1.ListAccountSettingsAssignmentsOptions)
				listAccountSettingsAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.TargetType = core.StringPtr("Account")
				listAccountSettingsAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listAccountSettingsAssignmentsOptionsModel.Pagetoken = core.StringPtr("testString")
				listAccountSettingsAssignmentsOptionsModel.Sort = core.StringPtr("created_at")
				listAccountSettingsAssignmentsOptionsModel.Order = core.StringPtr("asc")
				listAccountSettingsAssignmentsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listAccountSettingsAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListAccountSettingsAssignments(listAccountSettingsAssignmentsOptionsModel)
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
	Describe(`CreateAccountSettingsAssignment(createAccountSettingsAssignmentOptions *CreateAccountSettingsAssignmentOptions) - Operation response error`, func() {
		createAccountSettingsAssignmentPath := "/v1/account_settings_assignments/"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountSettingsAssignmentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccountSettingsAssignment with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateAccountSettingsAssignmentOptions model
				createAccountSettingsAssignmentOptionsModel := new(iamidentityv1.CreateAccountSettingsAssignmentOptions)
				createAccountSettingsAssignmentOptionsModel.TemplateID = core.StringPtr("testString")
				createAccountSettingsAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				createAccountSettingsAssignmentOptionsModel.TargetType = core.StringPtr("Account")
				createAccountSettingsAssignmentOptionsModel.Target = core.StringPtr("testString")
				createAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateAccountSettingsAssignment(createAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateAccountSettingsAssignment(createAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccountSettingsAssignment(createAccountSettingsAssignmentOptions *CreateAccountSettingsAssignmentOptions)`, func() {
		createAccountSettingsAssignmentPath := "/v1/account_settings_assignments/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountSettingsAssignmentPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke CreateAccountSettingsAssignment successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the CreateAccountSettingsAssignmentOptions model
				createAccountSettingsAssignmentOptionsModel := new(iamidentityv1.CreateAccountSettingsAssignmentOptions)
				createAccountSettingsAssignmentOptionsModel.TemplateID = core.StringPtr("testString")
				createAccountSettingsAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				createAccountSettingsAssignmentOptionsModel.TargetType = core.StringPtr("Account")
				createAccountSettingsAssignmentOptionsModel.Target = core.StringPtr("testString")
				createAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateAccountSettingsAssignmentWithContext(ctx, createAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateAccountSettingsAssignment(createAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateAccountSettingsAssignmentWithContext(ctx, createAccountSettingsAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccountSettingsAssignmentPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke CreateAccountSettingsAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateAccountSettingsAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccountSettingsAssignmentOptions model
				createAccountSettingsAssignmentOptionsModel := new(iamidentityv1.CreateAccountSettingsAssignmentOptions)
				createAccountSettingsAssignmentOptionsModel.TemplateID = core.StringPtr("testString")
				createAccountSettingsAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				createAccountSettingsAssignmentOptionsModel.TargetType = core.StringPtr("Account")
				createAccountSettingsAssignmentOptionsModel.Target = core.StringPtr("testString")
				createAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateAccountSettingsAssignment(createAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccountSettingsAssignment with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateAccountSettingsAssignmentOptions model
				createAccountSettingsAssignmentOptionsModel := new(iamidentityv1.CreateAccountSettingsAssignmentOptions)
				createAccountSettingsAssignmentOptionsModel.TemplateID = core.StringPtr("testString")
				createAccountSettingsAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				createAccountSettingsAssignmentOptionsModel.TargetType = core.StringPtr("Account")
				createAccountSettingsAssignmentOptionsModel.Target = core.StringPtr("testString")
				createAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateAccountSettingsAssignment(createAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccountSettingsAssignmentOptions model with no property values
				createAccountSettingsAssignmentOptionsModelNew := new(iamidentityv1.CreateAccountSettingsAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateAccountSettingsAssignment(createAccountSettingsAssignmentOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateAccountSettingsAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateAccountSettingsAssignmentOptions model
				createAccountSettingsAssignmentOptionsModel := new(iamidentityv1.CreateAccountSettingsAssignmentOptions)
				createAccountSettingsAssignmentOptionsModel.TemplateID = core.StringPtr("testString")
				createAccountSettingsAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				createAccountSettingsAssignmentOptionsModel.TargetType = core.StringPtr("Account")
				createAccountSettingsAssignmentOptionsModel.Target = core.StringPtr("testString")
				createAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateAccountSettingsAssignment(createAccountSettingsAssignmentOptionsModel)
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
	Describe(`GetAccountSettingsAssignment(getAccountSettingsAssignmentOptions *GetAccountSettingsAssignmentOptions) - Operation response error`, func() {
		getAccountSettingsAssignmentPath := "/v1/account_settings_assignments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsAssignmentPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountSettingsAssignment with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsAssignmentOptions model
				getAccountSettingsAssignmentOptionsModel := new(iamidentityv1.GetAccountSettingsAssignmentOptions)
				getAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getAccountSettingsAssignmentOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetAccountSettingsAssignment(getAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetAccountSettingsAssignment(getAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountSettingsAssignment(getAccountSettingsAssignmentOptions *GetAccountSettingsAssignmentOptions)`, func() {
		getAccountSettingsAssignmentPath := "/v1/account_settings_assignments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsAssignmentPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke GetAccountSettingsAssignment successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountSettingsAssignmentOptions model
				getAccountSettingsAssignmentOptionsModel := new(iamidentityv1.GetAccountSettingsAssignmentOptions)
				getAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getAccountSettingsAssignmentOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetAccountSettingsAssignmentWithContext(ctx, getAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetAccountSettingsAssignment(getAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetAccountSettingsAssignmentWithContext(ctx, getAccountSettingsAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsAssignmentPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke GetAccountSettingsAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetAccountSettingsAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSettingsAssignmentOptions model
				getAccountSettingsAssignmentOptionsModel := new(iamidentityv1.GetAccountSettingsAssignmentOptions)
				getAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getAccountSettingsAssignmentOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetAccountSettingsAssignment(getAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountSettingsAssignment with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsAssignmentOptions model
				getAccountSettingsAssignmentOptionsModel := new(iamidentityv1.GetAccountSettingsAssignmentOptions)
				getAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getAccountSettingsAssignmentOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetAccountSettingsAssignment(getAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountSettingsAssignmentOptions model with no property values
				getAccountSettingsAssignmentOptionsModelNew := new(iamidentityv1.GetAccountSettingsAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetAccountSettingsAssignment(getAccountSettingsAssignmentOptionsModelNew)
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
			It(`Invoke GetAccountSettingsAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsAssignmentOptions model
				getAccountSettingsAssignmentOptionsModel := new(iamidentityv1.GetAccountSettingsAssignmentOptions)
				getAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getAccountSettingsAssignmentOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetAccountSettingsAssignment(getAccountSettingsAssignmentOptionsModel)
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
	Describe(`DeleteAccountSettingsAssignment(deleteAccountSettingsAssignmentOptions *DeleteAccountSettingsAssignmentOptions) - Operation response error`, func() {
		deleteAccountSettingsAssignmentPath := "/v1/account_settings_assignments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountSettingsAssignmentPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAccountSettingsAssignment with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountSettingsAssignmentOptions model
				deleteAccountSettingsAssignmentOptionsModel := new(iamidentityv1.DeleteAccountSettingsAssignmentOptions)
				deleteAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.DeleteAccountSettingsAssignment(deleteAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.DeleteAccountSettingsAssignment(deleteAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAccountSettingsAssignment(deleteAccountSettingsAssignmentOptions *DeleteAccountSettingsAssignmentOptions)`, func() {
		deleteAccountSettingsAssignmentPath := "/v1/account_settings_assignments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountSettingsAssignmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "status_code": "StatusCode", "errors": [{"code": "Code", "message_code": "MessageCode", "message": "Message", "details": "Details"}], "trace": "Trace"}`)
				}))
			})
			It(`Invoke DeleteAccountSettingsAssignment successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the DeleteAccountSettingsAssignmentOptions model
				deleteAccountSettingsAssignmentOptionsModel := new(iamidentityv1.DeleteAccountSettingsAssignmentOptions)
				deleteAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.DeleteAccountSettingsAssignmentWithContext(ctx, deleteAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.DeleteAccountSettingsAssignment(deleteAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.DeleteAccountSettingsAssignmentWithContext(ctx, deleteAccountSettingsAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountSettingsAssignmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "status_code": "StatusCode", "errors": [{"code": "Code", "message_code": "MessageCode", "message": "Message", "details": "Details"}], "trace": "Trace"}`)
				}))
			})
			It(`Invoke DeleteAccountSettingsAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.DeleteAccountSettingsAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAccountSettingsAssignmentOptions model
				deleteAccountSettingsAssignmentOptionsModel := new(iamidentityv1.DeleteAccountSettingsAssignmentOptions)
				deleteAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.DeleteAccountSettingsAssignment(deleteAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteAccountSettingsAssignment with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountSettingsAssignmentOptions model
				deleteAccountSettingsAssignmentOptionsModel := new(iamidentityv1.DeleteAccountSettingsAssignmentOptions)
				deleteAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.DeleteAccountSettingsAssignment(deleteAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteAccountSettingsAssignmentOptions model with no property values
				deleteAccountSettingsAssignmentOptionsModelNew := new(iamidentityv1.DeleteAccountSettingsAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.DeleteAccountSettingsAssignment(deleteAccountSettingsAssignmentOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteAccountSettingsAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountSettingsAssignmentOptions model
				deleteAccountSettingsAssignmentOptionsModel := new(iamidentityv1.DeleteAccountSettingsAssignmentOptions)
				deleteAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.DeleteAccountSettingsAssignment(deleteAccountSettingsAssignmentOptionsModel)
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
	Describe(`UpdateAccountSettingsAssignment(updateAccountSettingsAssignmentOptions *UpdateAccountSettingsAssignmentOptions) - Operation response error`, func() {
		updateAccountSettingsAssignmentPath := "/v1/account_settings_assignments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsAssignmentPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountSettingsAssignment with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountSettingsAssignmentOptions model
				updateAccountSettingsAssignmentOptionsModel := new(iamidentityv1.UpdateAccountSettingsAssignmentOptions)
				updateAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateAccountSettingsAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				updateAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateAccountSettingsAssignment(updateAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateAccountSettingsAssignment(updateAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountSettingsAssignment(updateAccountSettingsAssignmentOptions *UpdateAccountSettingsAssignmentOptions)`, func() {
		updateAccountSettingsAssignmentPath := "/v1/account_settings_assignments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsAssignmentPath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke UpdateAccountSettingsAssignment successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAccountSettingsAssignmentOptions model
				updateAccountSettingsAssignmentOptionsModel := new(iamidentityv1.UpdateAccountSettingsAssignmentOptions)
				updateAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateAccountSettingsAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				updateAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.UpdateAccountSettingsAssignmentWithContext(ctx, updateAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.UpdateAccountSettingsAssignment(updateAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.UpdateAccountSettingsAssignmentWithContext(ctx, updateAccountSettingsAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsAssignmentPath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke UpdateAccountSettingsAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateAccountSettingsAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAccountSettingsAssignmentOptions model
				updateAccountSettingsAssignmentOptionsModel := new(iamidentityv1.UpdateAccountSettingsAssignmentOptions)
				updateAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateAccountSettingsAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				updateAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateAccountSettingsAssignment(updateAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccountSettingsAssignment with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountSettingsAssignmentOptions model
				updateAccountSettingsAssignmentOptionsModel := new(iamidentityv1.UpdateAccountSettingsAssignmentOptions)
				updateAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateAccountSettingsAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				updateAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateAccountSettingsAssignment(updateAccountSettingsAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountSettingsAssignmentOptions model with no property values
				updateAccountSettingsAssignmentOptionsModelNew := new(iamidentityv1.UpdateAccountSettingsAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateAccountSettingsAssignment(updateAccountSettingsAssignmentOptionsModelNew)
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
			It(`Invoke UpdateAccountSettingsAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountSettingsAssignmentOptions model
				updateAccountSettingsAssignmentOptionsModel := new(iamidentityv1.UpdateAccountSettingsAssignmentOptions)
				updateAccountSettingsAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateAccountSettingsAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				updateAccountSettingsAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.UpdateAccountSettingsAssignment(updateAccountSettingsAssignmentOptionsModel)
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
	Describe(`ListAccountSettingsTemplates(listAccountSettingsTemplatesOptions *ListAccountSettingsTemplatesOptions) - Operation response error`, func() {
		listAccountSettingsTemplatesPath := "/v1/account_settings_templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccountSettingsTemplatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccountSettingsTemplates with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListAccountSettingsTemplatesOptions model
				listAccountSettingsTemplatesOptionsModel := new(iamidentityv1.ListAccountSettingsTemplatesOptions)
				listAccountSettingsTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listAccountSettingsTemplatesOptionsModel.Limit = core.StringPtr("20")
				listAccountSettingsTemplatesOptionsModel.Pagetoken = core.StringPtr("testString")
				listAccountSettingsTemplatesOptionsModel.Sort = core.StringPtr("created_at")
				listAccountSettingsTemplatesOptionsModel.Order = core.StringPtr("asc")
				listAccountSettingsTemplatesOptionsModel.IncludeHistory = core.StringPtr("false")
				listAccountSettingsTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListAccountSettingsTemplates(listAccountSettingsTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListAccountSettingsTemplates(listAccountSettingsTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccountSettingsTemplates(listAccountSettingsTemplatesOptions *ListAccountSettingsTemplatesOptions)`, func() {
		listAccountSettingsTemplatesPath := "/v1/account_settings_templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccountSettingsTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 20, "first": "First", "previous": "Previous", "next": "Next", "account_settings_templates": [{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListAccountSettingsTemplates successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListAccountSettingsTemplatesOptions model
				listAccountSettingsTemplatesOptionsModel := new(iamidentityv1.ListAccountSettingsTemplatesOptions)
				listAccountSettingsTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listAccountSettingsTemplatesOptionsModel.Limit = core.StringPtr("20")
				listAccountSettingsTemplatesOptionsModel.Pagetoken = core.StringPtr("testString")
				listAccountSettingsTemplatesOptionsModel.Sort = core.StringPtr("created_at")
				listAccountSettingsTemplatesOptionsModel.Order = core.StringPtr("asc")
				listAccountSettingsTemplatesOptionsModel.IncludeHistory = core.StringPtr("false")
				listAccountSettingsTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListAccountSettingsTemplatesWithContext(ctx, listAccountSettingsTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListAccountSettingsTemplates(listAccountSettingsTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListAccountSettingsTemplatesWithContext(ctx, listAccountSettingsTemplatesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAccountSettingsTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 20, "first": "First", "previous": "Previous", "next": "Next", "account_settings_templates": [{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListAccountSettingsTemplates successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListAccountSettingsTemplates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccountSettingsTemplatesOptions model
				listAccountSettingsTemplatesOptionsModel := new(iamidentityv1.ListAccountSettingsTemplatesOptions)
				listAccountSettingsTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listAccountSettingsTemplatesOptionsModel.Limit = core.StringPtr("20")
				listAccountSettingsTemplatesOptionsModel.Pagetoken = core.StringPtr("testString")
				listAccountSettingsTemplatesOptionsModel.Sort = core.StringPtr("created_at")
				listAccountSettingsTemplatesOptionsModel.Order = core.StringPtr("asc")
				listAccountSettingsTemplatesOptionsModel.IncludeHistory = core.StringPtr("false")
				listAccountSettingsTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListAccountSettingsTemplates(listAccountSettingsTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAccountSettingsTemplates with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListAccountSettingsTemplatesOptions model
				listAccountSettingsTemplatesOptionsModel := new(iamidentityv1.ListAccountSettingsTemplatesOptions)
				listAccountSettingsTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listAccountSettingsTemplatesOptionsModel.Limit = core.StringPtr("20")
				listAccountSettingsTemplatesOptionsModel.Pagetoken = core.StringPtr("testString")
				listAccountSettingsTemplatesOptionsModel.Sort = core.StringPtr("created_at")
				listAccountSettingsTemplatesOptionsModel.Order = core.StringPtr("asc")
				listAccountSettingsTemplatesOptionsModel.IncludeHistory = core.StringPtr("false")
				listAccountSettingsTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListAccountSettingsTemplates(listAccountSettingsTemplatesOptionsModel)
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
			It(`Invoke ListAccountSettingsTemplates successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListAccountSettingsTemplatesOptions model
				listAccountSettingsTemplatesOptionsModel := new(iamidentityv1.ListAccountSettingsTemplatesOptions)
				listAccountSettingsTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listAccountSettingsTemplatesOptionsModel.Limit = core.StringPtr("20")
				listAccountSettingsTemplatesOptionsModel.Pagetoken = core.StringPtr("testString")
				listAccountSettingsTemplatesOptionsModel.Sort = core.StringPtr("created_at")
				listAccountSettingsTemplatesOptionsModel.Order = core.StringPtr("asc")
				listAccountSettingsTemplatesOptionsModel.IncludeHistory = core.StringPtr("false")
				listAccountSettingsTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListAccountSettingsTemplates(listAccountSettingsTemplatesOptionsModel)
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
	Describe(`CreateAccountSettingsTemplate(createAccountSettingsTemplateOptions *CreateAccountSettingsTemplateOptions) - Operation response error`, func() {
		createAccountSettingsTemplatePath := "/v1/account_settings_templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountSettingsTemplatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccountSettingsTemplate with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the CreateAccountSettingsTemplateOptions model
				createAccountSettingsTemplateOptionsModel := new(iamidentityv1.CreateAccountSettingsTemplateOptions)
				createAccountSettingsTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.Name = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.Description = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.AccountSettings = accountSettingsComponentModel
				createAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateAccountSettingsTemplate(createAccountSettingsTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateAccountSettingsTemplate(createAccountSettingsTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccountSettingsTemplate(createAccountSettingsTemplateOptions *CreateAccountSettingsTemplateOptions)`, func() {
		createAccountSettingsTemplatePath := "/v1/account_settings_templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountSettingsTemplatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateAccountSettingsTemplate successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the CreateAccountSettingsTemplateOptions model
				createAccountSettingsTemplateOptionsModel := new(iamidentityv1.CreateAccountSettingsTemplateOptions)
				createAccountSettingsTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.Name = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.Description = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.AccountSettings = accountSettingsComponentModel
				createAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateAccountSettingsTemplateWithContext(ctx, createAccountSettingsTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateAccountSettingsTemplate(createAccountSettingsTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateAccountSettingsTemplateWithContext(ctx, createAccountSettingsTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccountSettingsTemplatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateAccountSettingsTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateAccountSettingsTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the CreateAccountSettingsTemplateOptions model
				createAccountSettingsTemplateOptionsModel := new(iamidentityv1.CreateAccountSettingsTemplateOptions)
				createAccountSettingsTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.Name = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.Description = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.AccountSettings = accountSettingsComponentModel
				createAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateAccountSettingsTemplate(createAccountSettingsTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccountSettingsTemplate with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the CreateAccountSettingsTemplateOptions model
				createAccountSettingsTemplateOptionsModel := new(iamidentityv1.CreateAccountSettingsTemplateOptions)
				createAccountSettingsTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.Name = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.Description = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.AccountSettings = accountSettingsComponentModel
				createAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateAccountSettingsTemplate(createAccountSettingsTemplateOptionsModel)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateAccountSettingsTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the CreateAccountSettingsTemplateOptions model
				createAccountSettingsTemplateOptionsModel := new(iamidentityv1.CreateAccountSettingsTemplateOptions)
				createAccountSettingsTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.Name = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.Description = core.StringPtr("testString")
				createAccountSettingsTemplateOptionsModel.AccountSettings = accountSettingsComponentModel
				createAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateAccountSettingsTemplate(createAccountSettingsTemplateOptionsModel)
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
	Describe(`GetLatestAccountSettingsTemplateVersion(getLatestAccountSettingsTemplateVersionOptions *GetLatestAccountSettingsTemplateVersionOptions) - Operation response error`, func() {
		getLatestAccountSettingsTemplateVersionPath := "/v1/account_settings_templates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestAccountSettingsTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLatestAccountSettingsTemplateVersion with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetLatestAccountSettingsTemplateVersionOptions model
				getLatestAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.GetLatestAccountSettingsTemplateVersionOptions)
				getLatestAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestAccountSettingsTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getLatestAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetLatestAccountSettingsTemplateVersion(getLatestAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetLatestAccountSettingsTemplateVersion(getLatestAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLatestAccountSettingsTemplateVersion(getLatestAccountSettingsTemplateVersionOptions *GetLatestAccountSettingsTemplateVersionOptions)`, func() {
		getLatestAccountSettingsTemplateVersionPath := "/v1/account_settings_templates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestAccountSettingsTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetLatestAccountSettingsTemplateVersion successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetLatestAccountSettingsTemplateVersionOptions model
				getLatestAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.GetLatestAccountSettingsTemplateVersionOptions)
				getLatestAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestAccountSettingsTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getLatestAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetLatestAccountSettingsTemplateVersionWithContext(ctx, getLatestAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetLatestAccountSettingsTemplateVersion(getLatestAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetLatestAccountSettingsTemplateVersionWithContext(ctx, getLatestAccountSettingsTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLatestAccountSettingsTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetLatestAccountSettingsTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetLatestAccountSettingsTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLatestAccountSettingsTemplateVersionOptions model
				getLatestAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.GetLatestAccountSettingsTemplateVersionOptions)
				getLatestAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestAccountSettingsTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getLatestAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetLatestAccountSettingsTemplateVersion(getLatestAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLatestAccountSettingsTemplateVersion with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetLatestAccountSettingsTemplateVersionOptions model
				getLatestAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.GetLatestAccountSettingsTemplateVersionOptions)
				getLatestAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestAccountSettingsTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getLatestAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetLatestAccountSettingsTemplateVersion(getLatestAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLatestAccountSettingsTemplateVersionOptions model with no property values
				getLatestAccountSettingsTemplateVersionOptionsModelNew := new(iamidentityv1.GetLatestAccountSettingsTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetLatestAccountSettingsTemplateVersion(getLatestAccountSettingsTemplateVersionOptionsModelNew)
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
			It(`Invoke GetLatestAccountSettingsTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetLatestAccountSettingsTemplateVersionOptions model
				getLatestAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.GetLatestAccountSettingsTemplateVersionOptions)
				getLatestAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestAccountSettingsTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getLatestAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetLatestAccountSettingsTemplateVersion(getLatestAccountSettingsTemplateVersionOptionsModel)
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
	Describe(`DeleteAllVersionsOfAccountSettingsTemplate(deleteAllVersionsOfAccountSettingsTemplateOptions *DeleteAllVersionsOfAccountSettingsTemplateOptions)`, func() {
		deleteAllVersionsOfAccountSettingsTemplatePath := "/v1/account_settings_templates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllVersionsOfAccountSettingsTemplatePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAllVersionsOfAccountSettingsTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteAllVersionsOfAccountSettingsTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAllVersionsOfAccountSettingsTemplateOptions model
				deleteAllVersionsOfAccountSettingsTemplateOptionsModel := new(iamidentityv1.DeleteAllVersionsOfAccountSettingsTemplateOptions)
				deleteAllVersionsOfAccountSettingsTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				deleteAllVersionsOfAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteAllVersionsOfAccountSettingsTemplate(deleteAllVersionsOfAccountSettingsTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAllVersionsOfAccountSettingsTemplate with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteAllVersionsOfAccountSettingsTemplateOptions model
				deleteAllVersionsOfAccountSettingsTemplateOptionsModel := new(iamidentityv1.DeleteAllVersionsOfAccountSettingsTemplateOptions)
				deleteAllVersionsOfAccountSettingsTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				deleteAllVersionsOfAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteAllVersionsOfAccountSettingsTemplate(deleteAllVersionsOfAccountSettingsTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAllVersionsOfAccountSettingsTemplateOptions model with no property values
				deleteAllVersionsOfAccountSettingsTemplateOptionsModelNew := new(iamidentityv1.DeleteAllVersionsOfAccountSettingsTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteAllVersionsOfAccountSettingsTemplate(deleteAllVersionsOfAccountSettingsTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVersionsOfAccountSettingsTemplate(listVersionsOfAccountSettingsTemplateOptions *ListVersionsOfAccountSettingsTemplateOptions) - Operation response error`, func() {
		listVersionsOfAccountSettingsTemplatePath := "/v1/account_settings_templates/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVersionsOfAccountSettingsTemplatePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListVersionsOfAccountSettingsTemplate with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListVersionsOfAccountSettingsTemplateOptions model
				listVersionsOfAccountSettingsTemplateOptionsModel := new(iamidentityv1.ListVersionsOfAccountSettingsTemplateOptions)
				listVersionsOfAccountSettingsTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.Limit = core.StringPtr("20")
				listVersionsOfAccountSettingsTemplateOptionsModel.Pagetoken = core.StringPtr("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.Sort = core.StringPtr("created_at")
				listVersionsOfAccountSettingsTemplateOptionsModel.Order = core.StringPtr("asc")
				listVersionsOfAccountSettingsTemplateOptionsModel.IncludeHistory = core.StringPtr("false")
				listVersionsOfAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListVersionsOfAccountSettingsTemplate(listVersionsOfAccountSettingsTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListVersionsOfAccountSettingsTemplate(listVersionsOfAccountSettingsTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVersionsOfAccountSettingsTemplate(listVersionsOfAccountSettingsTemplateOptions *ListVersionsOfAccountSettingsTemplateOptions)`, func() {
		listVersionsOfAccountSettingsTemplatePath := "/v1/account_settings_templates/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVersionsOfAccountSettingsTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 20, "first": "First", "previous": "Previous", "next": "Next", "account_settings_templates": [{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListVersionsOfAccountSettingsTemplate successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListVersionsOfAccountSettingsTemplateOptions model
				listVersionsOfAccountSettingsTemplateOptionsModel := new(iamidentityv1.ListVersionsOfAccountSettingsTemplateOptions)
				listVersionsOfAccountSettingsTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.Limit = core.StringPtr("20")
				listVersionsOfAccountSettingsTemplateOptionsModel.Pagetoken = core.StringPtr("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.Sort = core.StringPtr("created_at")
				listVersionsOfAccountSettingsTemplateOptionsModel.Order = core.StringPtr("asc")
				listVersionsOfAccountSettingsTemplateOptionsModel.IncludeHistory = core.StringPtr("false")
				listVersionsOfAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListVersionsOfAccountSettingsTemplateWithContext(ctx, listVersionsOfAccountSettingsTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListVersionsOfAccountSettingsTemplate(listVersionsOfAccountSettingsTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListVersionsOfAccountSettingsTemplateWithContext(ctx, listVersionsOfAccountSettingsTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listVersionsOfAccountSettingsTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 20, "first": "First", "previous": "Previous", "next": "Next", "account_settings_templates": [{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListVersionsOfAccountSettingsTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListVersionsOfAccountSettingsTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVersionsOfAccountSettingsTemplateOptions model
				listVersionsOfAccountSettingsTemplateOptionsModel := new(iamidentityv1.ListVersionsOfAccountSettingsTemplateOptions)
				listVersionsOfAccountSettingsTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.Limit = core.StringPtr("20")
				listVersionsOfAccountSettingsTemplateOptionsModel.Pagetoken = core.StringPtr("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.Sort = core.StringPtr("created_at")
				listVersionsOfAccountSettingsTemplateOptionsModel.Order = core.StringPtr("asc")
				listVersionsOfAccountSettingsTemplateOptionsModel.IncludeHistory = core.StringPtr("false")
				listVersionsOfAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListVersionsOfAccountSettingsTemplate(listVersionsOfAccountSettingsTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListVersionsOfAccountSettingsTemplate with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListVersionsOfAccountSettingsTemplateOptions model
				listVersionsOfAccountSettingsTemplateOptionsModel := new(iamidentityv1.ListVersionsOfAccountSettingsTemplateOptions)
				listVersionsOfAccountSettingsTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.Limit = core.StringPtr("20")
				listVersionsOfAccountSettingsTemplateOptionsModel.Pagetoken = core.StringPtr("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.Sort = core.StringPtr("created_at")
				listVersionsOfAccountSettingsTemplateOptionsModel.Order = core.StringPtr("asc")
				listVersionsOfAccountSettingsTemplateOptionsModel.IncludeHistory = core.StringPtr("false")
				listVersionsOfAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListVersionsOfAccountSettingsTemplate(listVersionsOfAccountSettingsTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListVersionsOfAccountSettingsTemplateOptions model with no property values
				listVersionsOfAccountSettingsTemplateOptionsModelNew := new(iamidentityv1.ListVersionsOfAccountSettingsTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.ListVersionsOfAccountSettingsTemplate(listVersionsOfAccountSettingsTemplateOptionsModelNew)
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
			It(`Invoke ListVersionsOfAccountSettingsTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListVersionsOfAccountSettingsTemplateOptions model
				listVersionsOfAccountSettingsTemplateOptionsModel := new(iamidentityv1.ListVersionsOfAccountSettingsTemplateOptions)
				listVersionsOfAccountSettingsTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.Limit = core.StringPtr("20")
				listVersionsOfAccountSettingsTemplateOptionsModel.Pagetoken = core.StringPtr("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.Sort = core.StringPtr("created_at")
				listVersionsOfAccountSettingsTemplateOptionsModel.Order = core.StringPtr("asc")
				listVersionsOfAccountSettingsTemplateOptionsModel.IncludeHistory = core.StringPtr("false")
				listVersionsOfAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListVersionsOfAccountSettingsTemplate(listVersionsOfAccountSettingsTemplateOptionsModel)
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
	Describe(`CreateAccountSettingsTemplateVersion(createAccountSettingsTemplateVersionOptions *CreateAccountSettingsTemplateVersionOptions) - Operation response error`, func() {
		createAccountSettingsTemplateVersionPath := "/v1/account_settings_templates/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountSettingsTemplateVersionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccountSettingsTemplateVersion with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the CreateAccountSettingsTemplateVersionOptions model
				createAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.CreateAccountSettingsTemplateVersionOptions)
				createAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.AccountSettings = accountSettingsComponentModel
				createAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateAccountSettingsTemplateVersion(createAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateAccountSettingsTemplateVersion(createAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccountSettingsTemplateVersion(createAccountSettingsTemplateVersionOptions *CreateAccountSettingsTemplateVersionOptions)`, func() {
		createAccountSettingsTemplateVersionPath := "/v1/account_settings_templates/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountSettingsTemplateVersionPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateAccountSettingsTemplateVersion successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the CreateAccountSettingsTemplateVersionOptions model
				createAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.CreateAccountSettingsTemplateVersionOptions)
				createAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.AccountSettings = accountSettingsComponentModel
				createAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateAccountSettingsTemplateVersionWithContext(ctx, createAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateAccountSettingsTemplateVersion(createAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateAccountSettingsTemplateVersionWithContext(ctx, createAccountSettingsTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccountSettingsTemplateVersionPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateAccountSettingsTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateAccountSettingsTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the CreateAccountSettingsTemplateVersionOptions model
				createAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.CreateAccountSettingsTemplateVersionOptions)
				createAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.AccountSettings = accountSettingsComponentModel
				createAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateAccountSettingsTemplateVersion(createAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccountSettingsTemplateVersion with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the CreateAccountSettingsTemplateVersionOptions model
				createAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.CreateAccountSettingsTemplateVersionOptions)
				createAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.AccountSettings = accountSettingsComponentModel
				createAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateAccountSettingsTemplateVersion(createAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccountSettingsTemplateVersionOptions model with no property values
				createAccountSettingsTemplateVersionOptionsModelNew := new(iamidentityv1.CreateAccountSettingsTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateAccountSettingsTemplateVersion(createAccountSettingsTemplateVersionOptionsModelNew)
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
			It(`Invoke CreateAccountSettingsTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the CreateAccountSettingsTemplateVersionOptions model
				createAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.CreateAccountSettingsTemplateVersionOptions)
				createAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createAccountSettingsTemplateVersionOptionsModel.AccountSettings = accountSettingsComponentModel
				createAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateAccountSettingsTemplateVersion(createAccountSettingsTemplateVersionOptionsModel)
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
	Describe(`GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptions *GetAccountSettingsTemplateVersionOptions) - Operation response error`, func() {
		getAccountSettingsTemplateVersionPath := "/v1/account_settings_templates/testString/versions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountSettingsTemplateVersion with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsTemplateVersionOptions model
				getAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.GetAccountSettingsTemplateVersionOptions)
				getAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getAccountSettingsTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptions *GetAccountSettingsTemplateVersionOptions)`, func() {
		getAccountSettingsTemplateVersionPath := "/v1/account_settings_templates/testString/versions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAccountSettingsTemplateVersion successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountSettingsTemplateVersionOptions model
				getAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.GetAccountSettingsTemplateVersionOptions)
				getAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getAccountSettingsTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetAccountSettingsTemplateVersionWithContext(ctx, getAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetAccountSettingsTemplateVersionWithContext(ctx, getAccountSettingsTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAccountSettingsTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetAccountSettingsTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSettingsTemplateVersionOptions model
				getAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.GetAccountSettingsTemplateVersionOptions)
				getAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getAccountSettingsTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountSettingsTemplateVersion with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsTemplateVersionOptions model
				getAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.GetAccountSettingsTemplateVersionOptions)
				getAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getAccountSettingsTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountSettingsTemplateVersionOptions model with no property values
				getAccountSettingsTemplateVersionOptionsModelNew := new(iamidentityv1.GetAccountSettingsTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptionsModelNew)
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
			It(`Invoke GetAccountSettingsTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsTemplateVersionOptions model
				getAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.GetAccountSettingsTemplateVersionOptions)
				getAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getAccountSettingsTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptionsModel)
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
	Describe(`UpdateAccountSettingsTemplateVersion(updateAccountSettingsTemplateVersionOptions *UpdateAccountSettingsTemplateVersionOptions) - Operation response error`, func() {
		updateAccountSettingsTemplateVersionPath := "/v1/account_settings_templates/testString/versions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsTemplateVersionPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountSettingsTemplateVersion with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the UpdateAccountSettingsTemplateVersionOptions model
				updateAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.UpdateAccountSettingsTemplateVersionOptions)
				updateAccountSettingsTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.AccountSettings = accountSettingsComponentModel
				updateAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateAccountSettingsTemplateVersion(updateAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateAccountSettingsTemplateVersion(updateAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountSettingsTemplateVersion(updateAccountSettingsTemplateVersionOptions *UpdateAccountSettingsTemplateVersionOptions)`, func() {
		updateAccountSettingsTemplateVersionPath := "/v1/account_settings_templates/testString/versions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsTemplateVersionPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke UpdateAccountSettingsTemplateVersion successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the UpdateAccountSettingsTemplateVersionOptions model
				updateAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.UpdateAccountSettingsTemplateVersionOptions)
				updateAccountSettingsTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.AccountSettings = accountSettingsComponentModel
				updateAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.UpdateAccountSettingsTemplateVersionWithContext(ctx, updateAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.UpdateAccountSettingsTemplateVersion(updateAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.UpdateAccountSettingsTemplateVersionWithContext(ctx, updateAccountSettingsTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsTemplateVersionPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "account_settings": {"restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "mfa": "NONE", "user_mfa": [{"iam_id": "IamID", "mfa": "NONE"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity", "system_access_token_expiration_in_seconds": "3600", "system_refresh_token_expiration_in_seconds": "259200"}, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke UpdateAccountSettingsTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateAccountSettingsTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the UpdateAccountSettingsTemplateVersionOptions model
				updateAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.UpdateAccountSettingsTemplateVersionOptions)
				updateAccountSettingsTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.AccountSettings = accountSettingsComponentModel
				updateAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateAccountSettingsTemplateVersion(updateAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccountSettingsTemplateVersion with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the UpdateAccountSettingsTemplateVersionOptions model
				updateAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.UpdateAccountSettingsTemplateVersionOptions)
				updateAccountSettingsTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.AccountSettings = accountSettingsComponentModel
				updateAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateAccountSettingsTemplateVersion(updateAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountSettingsTemplateVersionOptions model with no property values
				updateAccountSettingsTemplateVersionOptionsModelNew := new(iamidentityv1.UpdateAccountSettingsTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateAccountSettingsTemplateVersion(updateAccountSettingsTemplateVersionOptionsModelNew)
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
			It(`Invoke UpdateAccountSettingsTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

				// Construct an instance of the UpdateAccountSettingsTemplateVersionOptions model
				updateAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.UpdateAccountSettingsTemplateVersionOptions)
				updateAccountSettingsTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				updateAccountSettingsTemplateVersionOptionsModel.AccountSettings = accountSettingsComponentModel
				updateAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.UpdateAccountSettingsTemplateVersion(updateAccountSettingsTemplateVersionOptionsModel)
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
	Describe(`DeleteAccountSettingsTemplateVersion(deleteAccountSettingsTemplateVersionOptions *DeleteAccountSettingsTemplateVersionOptions)`, func() {
		deleteAccountSettingsTemplateVersionPath := "/v1/account_settings_templates/testString/versions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountSettingsTemplateVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAccountSettingsTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteAccountSettingsTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAccountSettingsTemplateVersionOptions model
				deleteAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.DeleteAccountSettingsTemplateVersionOptions)
				deleteAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				deleteAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				deleteAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteAccountSettingsTemplateVersion(deleteAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAccountSettingsTemplateVersion with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountSettingsTemplateVersionOptions model
				deleteAccountSettingsTemplateVersionOptionsModel := new(iamidentityv1.DeleteAccountSettingsTemplateVersionOptions)
				deleteAccountSettingsTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				deleteAccountSettingsTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				deleteAccountSettingsTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteAccountSettingsTemplateVersion(deleteAccountSettingsTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAccountSettingsTemplateVersionOptions model with no property values
				deleteAccountSettingsTemplateVersionOptionsModelNew := new(iamidentityv1.DeleteAccountSettingsTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteAccountSettingsTemplateVersion(deleteAccountSettingsTemplateVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CommitAccountSettingsTemplate(commitAccountSettingsTemplateOptions *CommitAccountSettingsTemplateOptions)`, func() {
		commitAccountSettingsTemplatePath := "/v1/account_settings_templates/testString/versions/testString/commit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(commitAccountSettingsTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke CommitAccountSettingsTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.CommitAccountSettingsTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CommitAccountSettingsTemplateOptions model
				commitAccountSettingsTemplateOptionsModel := new(iamidentityv1.CommitAccountSettingsTemplateOptions)
				commitAccountSettingsTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				commitAccountSettingsTemplateOptionsModel.Version = core.StringPtr("testString")
				commitAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.CommitAccountSettingsTemplate(commitAccountSettingsTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CommitAccountSettingsTemplate with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CommitAccountSettingsTemplateOptions model
				commitAccountSettingsTemplateOptionsModel := new(iamidentityv1.CommitAccountSettingsTemplateOptions)
				commitAccountSettingsTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				commitAccountSettingsTemplateOptionsModel.Version = core.StringPtr("testString")
				commitAccountSettingsTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.CommitAccountSettingsTemplate(commitAccountSettingsTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CommitAccountSettingsTemplateOptions model with no property values
				commitAccountSettingsTemplateOptionsModelNew := new(iamidentityv1.CommitAccountSettingsTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.CommitAccountSettingsTemplate(commitAccountSettingsTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateReport(createReportOptions *CreateReportOptions) - Operation response error`, func() {
		createReportPath := "/v1/activity/accounts/testString/report"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createReportPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"inactive"}))
					Expect(req.URL.Query()["duration"]).To(Equal([]string{"720"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateReport with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateReportOptions model
				createReportOptionsModel := new(iamidentityv1.CreateReportOptions)
				createReportOptionsModel.AccountID = core.StringPtr("testString")
				createReportOptionsModel.Type = core.StringPtr("inactive")
				createReportOptionsModel.Duration = core.StringPtr("720")
				createReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateReport(createReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateReport(createReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateReport(createReportOptions *CreateReportOptions)`, func() {
		createReportPath := "/v1/activity/accounts/testString/report"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createReportPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"inactive"}))
					Expect(req.URL.Query()["duration"]).To(Equal([]string{"720"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"reference": "Reference"}`)
				}))
			})
			It(`Invoke CreateReport successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the CreateReportOptions model
				createReportOptionsModel := new(iamidentityv1.CreateReportOptions)
				createReportOptionsModel.AccountID = core.StringPtr("testString")
				createReportOptionsModel.Type = core.StringPtr("inactive")
				createReportOptionsModel.Duration = core.StringPtr("720")
				createReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateReportWithContext(ctx, createReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateReport(createReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateReportWithContext(ctx, createReportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createReportPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"inactive"}))
					Expect(req.URL.Query()["duration"]).To(Equal([]string{"720"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"reference": "Reference"}`)
				}))
			})
			It(`Invoke CreateReport successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateReportOptions model
				createReportOptionsModel := new(iamidentityv1.CreateReportOptions)
				createReportOptionsModel.AccountID = core.StringPtr("testString")
				createReportOptionsModel.Type = core.StringPtr("inactive")
				createReportOptionsModel.Duration = core.StringPtr("720")
				createReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateReport(createReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateReport with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateReportOptions model
				createReportOptionsModel := new(iamidentityv1.CreateReportOptions)
				createReportOptionsModel.AccountID = core.StringPtr("testString")
				createReportOptionsModel.Type = core.StringPtr("inactive")
				createReportOptionsModel.Duration = core.StringPtr("720")
				createReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateReport(createReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateReportOptions model with no property values
				createReportOptionsModelNew := new(iamidentityv1.CreateReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateReport(createReportOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateReport successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateReportOptions model
				createReportOptionsModel := new(iamidentityv1.CreateReportOptions)
				createReportOptionsModel.AccountID = core.StringPtr("testString")
				createReportOptionsModel.Type = core.StringPtr("inactive")
				createReportOptionsModel.Duration = core.StringPtr("720")
				createReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateReport(createReportOptionsModel)
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
	Describe(`GetReport(getReportOptions *GetReportOptions) - Operation response error`, func() {
		getReportPath := "/v1/activity/accounts/testString/report/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReport with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(iamidentityv1.GetReportOptions)
				getReportOptionsModel.AccountID = core.StringPtr("testString")
				getReportOptionsModel.Reference = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetReport(getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetReport(getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReport(getReportOptions *GetReportOptions)`, func() {
		getReportPath := "/v1/activity/accounts/testString/report/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_by": "CreatedBy", "reference": "Reference", "report_duration": "ReportDuration", "report_start_time": "ReportStartTime", "report_end_time": "ReportEndTime", "users": [{"iam_id": "IamID", "name": "Name", "username": "Username", "email": "Email", "last_authn": "LastAuthn"}], "apikeys": [{"id": "ID", "name": "Name", "type": "Type", "serviceid": {"id": "ID", "name": "Name"}, "user": {"iam_id": "IamID", "name": "Name", "username": "Username", "email": "Email"}, "last_authn": "LastAuthn"}], "serviceids": [{"id": "ID", "name": "Name", "last_authn": "LastAuthn"}], "profiles": [{"id": "ID", "name": "Name", "last_authn": "LastAuthn"}]}`)
				}))
			})
			It(`Invoke GetReport successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(iamidentityv1.GetReportOptions)
				getReportOptionsModel.AccountID = core.StringPtr("testString")
				getReportOptionsModel.Reference = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetReportWithContext(ctx, getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetReport(getReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetReportWithContext(ctx, getReportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReportPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_by": "CreatedBy", "reference": "Reference", "report_duration": "ReportDuration", "report_start_time": "ReportStartTime", "report_end_time": "ReportEndTime", "users": [{"iam_id": "IamID", "name": "Name", "username": "Username", "email": "Email", "last_authn": "LastAuthn"}], "apikeys": [{"id": "ID", "name": "Name", "type": "Type", "serviceid": {"id": "ID", "name": "Name"}, "user": {"iam_id": "IamID", "name": "Name", "username": "Username", "email": "Email"}, "last_authn": "LastAuthn"}], "serviceids": [{"id": "ID", "name": "Name", "last_authn": "LastAuthn"}], "profiles": [{"id": "ID", "name": "Name", "last_authn": "LastAuthn"}]}`)
				}))
			})
			It(`Invoke GetReport successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(iamidentityv1.GetReportOptions)
				getReportOptionsModel.AccountID = core.StringPtr("testString")
				getReportOptionsModel.Reference = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetReport(getReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReport with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(iamidentityv1.GetReportOptions)
				getReportOptionsModel.AccountID = core.StringPtr("testString")
				getReportOptionsModel.Reference = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetReport(getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportOptions model with no property values
				getReportOptionsModelNew := new(iamidentityv1.GetReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetReport(getReportOptionsModelNew)
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
			It(`Invoke GetReport successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(iamidentityv1.GetReportOptions)
				getReportOptionsModel.AccountID = core.StringPtr("testString")
				getReportOptionsModel.Reference = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetReport(getReportOptionsModel)
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
	Describe(`ListTrustedProfileAssignments(listTrustedProfileAssignmentsOptions *ListTrustedProfileAssignmentsOptions) - Operation response error`, func() {
		listTrustedProfileAssignmentsPath := "/v1/profile_assignments/"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTrustedProfileAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_type"]).To(Equal([]string{"Account"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTrustedProfileAssignments with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListTrustedProfileAssignmentsOptions model
				listTrustedProfileAssignmentsOptionsModel := new(iamidentityv1.ListTrustedProfileAssignmentsOptions)
				listTrustedProfileAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TargetType = core.StringPtr("Account")
				listTrustedProfileAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listTrustedProfileAssignmentsOptionsModel.Pagetoken = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.Sort = core.StringPtr("created_at")
				listTrustedProfileAssignmentsOptionsModel.Order = core.StringPtr("asc")
				listTrustedProfileAssignmentsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listTrustedProfileAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListTrustedProfileAssignments(listTrustedProfileAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListTrustedProfileAssignments(listTrustedProfileAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTrustedProfileAssignments(listTrustedProfileAssignmentsOptions *ListTrustedProfileAssignmentsOptions)`, func() {
		listTrustedProfileAssignmentsPath := "/v1/profile_assignments/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTrustedProfileAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_type"]).To(Equal([]string{"Account"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "assignments": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}]}`)
				}))
			})
			It(`Invoke ListTrustedProfileAssignments successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListTrustedProfileAssignmentsOptions model
				listTrustedProfileAssignmentsOptionsModel := new(iamidentityv1.ListTrustedProfileAssignmentsOptions)
				listTrustedProfileAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TargetType = core.StringPtr("Account")
				listTrustedProfileAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listTrustedProfileAssignmentsOptionsModel.Pagetoken = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.Sort = core.StringPtr("created_at")
				listTrustedProfileAssignmentsOptionsModel.Order = core.StringPtr("asc")
				listTrustedProfileAssignmentsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listTrustedProfileAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListTrustedProfileAssignmentsWithContext(ctx, listTrustedProfileAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListTrustedProfileAssignments(listTrustedProfileAssignmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListTrustedProfileAssignmentsWithContext(ctx, listTrustedProfileAssignmentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTrustedProfileAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_type"]).To(Equal([]string{"Account"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "assignments": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}]}`)
				}))
			})
			It(`Invoke ListTrustedProfileAssignments successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListTrustedProfileAssignments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTrustedProfileAssignmentsOptions model
				listTrustedProfileAssignmentsOptionsModel := new(iamidentityv1.ListTrustedProfileAssignmentsOptions)
				listTrustedProfileAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TargetType = core.StringPtr("Account")
				listTrustedProfileAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listTrustedProfileAssignmentsOptionsModel.Pagetoken = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.Sort = core.StringPtr("created_at")
				listTrustedProfileAssignmentsOptionsModel.Order = core.StringPtr("asc")
				listTrustedProfileAssignmentsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listTrustedProfileAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListTrustedProfileAssignments(listTrustedProfileAssignmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTrustedProfileAssignments with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListTrustedProfileAssignmentsOptions model
				listTrustedProfileAssignmentsOptionsModel := new(iamidentityv1.ListTrustedProfileAssignmentsOptions)
				listTrustedProfileAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TargetType = core.StringPtr("Account")
				listTrustedProfileAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listTrustedProfileAssignmentsOptionsModel.Pagetoken = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.Sort = core.StringPtr("created_at")
				listTrustedProfileAssignmentsOptionsModel.Order = core.StringPtr("asc")
				listTrustedProfileAssignmentsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listTrustedProfileAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListTrustedProfileAssignments(listTrustedProfileAssignmentsOptionsModel)
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
			It(`Invoke ListTrustedProfileAssignments successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListTrustedProfileAssignmentsOptions model
				listTrustedProfileAssignmentsOptionsModel := new(iamidentityv1.ListTrustedProfileAssignmentsOptions)
				listTrustedProfileAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.TargetType = core.StringPtr("Account")
				listTrustedProfileAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(20))
				listTrustedProfileAssignmentsOptionsModel.Pagetoken = core.StringPtr("testString")
				listTrustedProfileAssignmentsOptionsModel.Sort = core.StringPtr("created_at")
				listTrustedProfileAssignmentsOptionsModel.Order = core.StringPtr("asc")
				listTrustedProfileAssignmentsOptionsModel.IncludeHistory = core.BoolPtr(false)
				listTrustedProfileAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListTrustedProfileAssignments(listTrustedProfileAssignmentsOptionsModel)
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
	Describe(`CreateTrustedProfileAssignment(createTrustedProfileAssignmentOptions *CreateTrustedProfileAssignmentOptions) - Operation response error`, func() {
		createTrustedProfileAssignmentPath := "/v1/profile_assignments/"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTrustedProfileAssignmentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTrustedProfileAssignment with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateTrustedProfileAssignmentOptions model
				createTrustedProfileAssignmentOptionsModel := new(iamidentityv1.CreateTrustedProfileAssignmentOptions)
				createTrustedProfileAssignmentOptionsModel.TemplateID = core.StringPtr("testString")
				createTrustedProfileAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				createTrustedProfileAssignmentOptionsModel.TargetType = core.StringPtr("Account")
				createTrustedProfileAssignmentOptionsModel.Target = core.StringPtr("testString")
				createTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateTrustedProfileAssignment(createTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateTrustedProfileAssignment(createTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTrustedProfileAssignment(createTrustedProfileAssignmentOptions *CreateTrustedProfileAssignmentOptions)`, func() {
		createTrustedProfileAssignmentPath := "/v1/profile_assignments/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTrustedProfileAssignmentPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke CreateTrustedProfileAssignment successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the CreateTrustedProfileAssignmentOptions model
				createTrustedProfileAssignmentOptionsModel := new(iamidentityv1.CreateTrustedProfileAssignmentOptions)
				createTrustedProfileAssignmentOptionsModel.TemplateID = core.StringPtr("testString")
				createTrustedProfileAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				createTrustedProfileAssignmentOptionsModel.TargetType = core.StringPtr("Account")
				createTrustedProfileAssignmentOptionsModel.Target = core.StringPtr("testString")
				createTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateTrustedProfileAssignmentWithContext(ctx, createTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateTrustedProfileAssignment(createTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateTrustedProfileAssignmentWithContext(ctx, createTrustedProfileAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTrustedProfileAssignmentPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke CreateTrustedProfileAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateTrustedProfileAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTrustedProfileAssignmentOptions model
				createTrustedProfileAssignmentOptionsModel := new(iamidentityv1.CreateTrustedProfileAssignmentOptions)
				createTrustedProfileAssignmentOptionsModel.TemplateID = core.StringPtr("testString")
				createTrustedProfileAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				createTrustedProfileAssignmentOptionsModel.TargetType = core.StringPtr("Account")
				createTrustedProfileAssignmentOptionsModel.Target = core.StringPtr("testString")
				createTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateTrustedProfileAssignment(createTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTrustedProfileAssignment with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateTrustedProfileAssignmentOptions model
				createTrustedProfileAssignmentOptionsModel := new(iamidentityv1.CreateTrustedProfileAssignmentOptions)
				createTrustedProfileAssignmentOptionsModel.TemplateID = core.StringPtr("testString")
				createTrustedProfileAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				createTrustedProfileAssignmentOptionsModel.TargetType = core.StringPtr("Account")
				createTrustedProfileAssignmentOptionsModel.Target = core.StringPtr("testString")
				createTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateTrustedProfileAssignment(createTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTrustedProfileAssignmentOptions model with no property values
				createTrustedProfileAssignmentOptionsModelNew := new(iamidentityv1.CreateTrustedProfileAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateTrustedProfileAssignment(createTrustedProfileAssignmentOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateTrustedProfileAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateTrustedProfileAssignmentOptions model
				createTrustedProfileAssignmentOptionsModel := new(iamidentityv1.CreateTrustedProfileAssignmentOptions)
				createTrustedProfileAssignmentOptionsModel.TemplateID = core.StringPtr("testString")
				createTrustedProfileAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				createTrustedProfileAssignmentOptionsModel.TargetType = core.StringPtr("Account")
				createTrustedProfileAssignmentOptionsModel.Target = core.StringPtr("testString")
				createTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateTrustedProfileAssignment(createTrustedProfileAssignmentOptionsModel)
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
	Describe(`GetTrustedProfileAssignment(getTrustedProfileAssignmentOptions *GetTrustedProfileAssignmentOptions) - Operation response error`, func() {
		getTrustedProfileAssignmentPath := "/v1/profile_assignments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrustedProfileAssignmentPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTrustedProfileAssignment with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetTrustedProfileAssignmentOptions model
				getTrustedProfileAssignmentOptionsModel := new(iamidentityv1.GetTrustedProfileAssignmentOptions)
				getTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getTrustedProfileAssignmentOptionsModel.IncludeHistory = core.BoolPtr(false)
				getTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetTrustedProfileAssignment(getTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetTrustedProfileAssignment(getTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTrustedProfileAssignment(getTrustedProfileAssignmentOptions *GetTrustedProfileAssignmentOptions)`, func() {
		getTrustedProfileAssignmentPath := "/v1/profile_assignments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrustedProfileAssignmentPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke GetTrustedProfileAssignment successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetTrustedProfileAssignmentOptions model
				getTrustedProfileAssignmentOptionsModel := new(iamidentityv1.GetTrustedProfileAssignmentOptions)
				getTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getTrustedProfileAssignmentOptionsModel.IncludeHistory = core.BoolPtr(false)
				getTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetTrustedProfileAssignmentWithContext(ctx, getTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetTrustedProfileAssignment(getTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetTrustedProfileAssignmentWithContext(ctx, getTrustedProfileAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTrustedProfileAssignmentPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke GetTrustedProfileAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetTrustedProfileAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTrustedProfileAssignmentOptions model
				getTrustedProfileAssignmentOptionsModel := new(iamidentityv1.GetTrustedProfileAssignmentOptions)
				getTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getTrustedProfileAssignmentOptionsModel.IncludeHistory = core.BoolPtr(false)
				getTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetTrustedProfileAssignment(getTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTrustedProfileAssignment with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetTrustedProfileAssignmentOptions model
				getTrustedProfileAssignmentOptionsModel := new(iamidentityv1.GetTrustedProfileAssignmentOptions)
				getTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getTrustedProfileAssignmentOptionsModel.IncludeHistory = core.BoolPtr(false)
				getTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetTrustedProfileAssignment(getTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTrustedProfileAssignmentOptions model with no property values
				getTrustedProfileAssignmentOptionsModelNew := new(iamidentityv1.GetTrustedProfileAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetTrustedProfileAssignment(getTrustedProfileAssignmentOptionsModelNew)
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
			It(`Invoke GetTrustedProfileAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetTrustedProfileAssignmentOptions model
				getTrustedProfileAssignmentOptionsModel := new(iamidentityv1.GetTrustedProfileAssignmentOptions)
				getTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getTrustedProfileAssignmentOptionsModel.IncludeHistory = core.BoolPtr(false)
				getTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetTrustedProfileAssignment(getTrustedProfileAssignmentOptionsModel)
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
	Describe(`DeleteTrustedProfileAssignment(deleteTrustedProfileAssignmentOptions *DeleteTrustedProfileAssignmentOptions) - Operation response error`, func() {
		deleteTrustedProfileAssignmentPath := "/v1/profile_assignments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTrustedProfileAssignmentPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteTrustedProfileAssignment with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteTrustedProfileAssignmentOptions model
				deleteTrustedProfileAssignmentOptionsModel := new(iamidentityv1.DeleteTrustedProfileAssignmentOptions)
				deleteTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.DeleteTrustedProfileAssignment(deleteTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.DeleteTrustedProfileAssignment(deleteTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTrustedProfileAssignment(deleteTrustedProfileAssignmentOptions *DeleteTrustedProfileAssignmentOptions)`, func() {
		deleteTrustedProfileAssignmentPath := "/v1/profile_assignments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTrustedProfileAssignmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "status_code": "StatusCode", "errors": [{"code": "Code", "message_code": "MessageCode", "message": "Message", "details": "Details"}], "trace": "Trace"}`)
				}))
			})
			It(`Invoke DeleteTrustedProfileAssignment successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the DeleteTrustedProfileAssignmentOptions model
				deleteTrustedProfileAssignmentOptionsModel := new(iamidentityv1.DeleteTrustedProfileAssignmentOptions)
				deleteTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.DeleteTrustedProfileAssignmentWithContext(ctx, deleteTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.DeleteTrustedProfileAssignment(deleteTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.DeleteTrustedProfileAssignmentWithContext(ctx, deleteTrustedProfileAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteTrustedProfileAssignmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "status_code": "StatusCode", "errors": [{"code": "Code", "message_code": "MessageCode", "message": "Message", "details": "Details"}], "trace": "Trace"}`)
				}))
			})
			It(`Invoke DeleteTrustedProfileAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.DeleteTrustedProfileAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteTrustedProfileAssignmentOptions model
				deleteTrustedProfileAssignmentOptionsModel := new(iamidentityv1.DeleteTrustedProfileAssignmentOptions)
				deleteTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.DeleteTrustedProfileAssignment(deleteTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteTrustedProfileAssignment with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteTrustedProfileAssignmentOptions model
				deleteTrustedProfileAssignmentOptionsModel := new(iamidentityv1.DeleteTrustedProfileAssignmentOptions)
				deleteTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.DeleteTrustedProfileAssignment(deleteTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteTrustedProfileAssignmentOptions model with no property values
				deleteTrustedProfileAssignmentOptionsModelNew := new(iamidentityv1.DeleteTrustedProfileAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.DeleteTrustedProfileAssignment(deleteTrustedProfileAssignmentOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteTrustedProfileAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteTrustedProfileAssignmentOptions model
				deleteTrustedProfileAssignmentOptionsModel := new(iamidentityv1.DeleteTrustedProfileAssignmentOptions)
				deleteTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.DeleteTrustedProfileAssignment(deleteTrustedProfileAssignmentOptionsModel)
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
	Describe(`UpdateTrustedProfileAssignment(updateTrustedProfileAssignmentOptions *UpdateTrustedProfileAssignmentOptions) - Operation response error`, func() {
		updateTrustedProfileAssignmentPath := "/v1/profile_assignments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTrustedProfileAssignmentPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTrustedProfileAssignment with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateTrustedProfileAssignmentOptions model
				updateTrustedProfileAssignmentOptionsModel := new(iamidentityv1.UpdateTrustedProfileAssignmentOptions)
				updateTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateTrustedProfileAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateTrustedProfileAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				updateTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateTrustedProfileAssignment(updateTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateTrustedProfileAssignment(updateTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTrustedProfileAssignment(updateTrustedProfileAssignmentOptions *UpdateTrustedProfileAssignmentOptions)`, func() {
		updateTrustedProfileAssignmentPath := "/v1/profile_assignments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTrustedProfileAssignmentPath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke UpdateTrustedProfileAssignment successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the UpdateTrustedProfileAssignmentOptions model
				updateTrustedProfileAssignmentOptionsModel := new(iamidentityv1.UpdateTrustedProfileAssignmentOptions)
				updateTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateTrustedProfileAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateTrustedProfileAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				updateTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.UpdateTrustedProfileAssignmentWithContext(ctx, updateTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.UpdateTrustedProfileAssignment(updateTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.UpdateTrustedProfileAssignmentWithContext(ctx, updateTrustedProfileAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTrustedProfileAssignmentPath))
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": 15, "target_type": "TargetType", "target": "Target", "status": "Status", "resources": [{"target": "Target", "profile": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "account_settings": {"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}, "policy_template_refs": [{"id": "ID", "version": "Version", "resource_created": {"id": "ID"}, "error_message": {"name": "Name", "errorCode": "ErrorCode", "message": "Message", "statusCode": "StatusCode"}, "status": "Status"}]}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "entity_tag": "EntityTag"}`)
				}))
			})
			It(`Invoke UpdateTrustedProfileAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateTrustedProfileAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTrustedProfileAssignmentOptions model
				updateTrustedProfileAssignmentOptionsModel := new(iamidentityv1.UpdateTrustedProfileAssignmentOptions)
				updateTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateTrustedProfileAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateTrustedProfileAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				updateTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateTrustedProfileAssignment(updateTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTrustedProfileAssignment with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateTrustedProfileAssignmentOptions model
				updateTrustedProfileAssignmentOptionsModel := new(iamidentityv1.UpdateTrustedProfileAssignmentOptions)
				updateTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateTrustedProfileAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateTrustedProfileAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				updateTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateTrustedProfileAssignment(updateTrustedProfileAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTrustedProfileAssignmentOptions model with no property values
				updateTrustedProfileAssignmentOptionsModelNew := new(iamidentityv1.UpdateTrustedProfileAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateTrustedProfileAssignment(updateTrustedProfileAssignmentOptionsModelNew)
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
			It(`Invoke UpdateTrustedProfileAssignment successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateTrustedProfileAssignmentOptions model
				updateTrustedProfileAssignmentOptionsModel := new(iamidentityv1.UpdateTrustedProfileAssignmentOptions)
				updateTrustedProfileAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateTrustedProfileAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateTrustedProfileAssignmentOptionsModel.TemplateVersion = core.Int64Ptr(int64(1))
				updateTrustedProfileAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.UpdateTrustedProfileAssignment(updateTrustedProfileAssignmentOptionsModel)
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
	Describe(`ListProfileTemplates(listProfileTemplatesOptions *ListProfileTemplatesOptions) - Operation response error`, func() {
		listProfileTemplatesPath := "/v1/profile_templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfileTemplatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProfileTemplates with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListProfileTemplatesOptions model
				listProfileTemplatesOptionsModel := new(iamidentityv1.ListProfileTemplatesOptions)
				listProfileTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listProfileTemplatesOptionsModel.Limit = core.StringPtr("20")
				listProfileTemplatesOptionsModel.Pagetoken = core.StringPtr("testString")
				listProfileTemplatesOptionsModel.Sort = core.StringPtr("created_at")
				listProfileTemplatesOptionsModel.Order = core.StringPtr("asc")
				listProfileTemplatesOptionsModel.IncludeHistory = core.StringPtr("false")
				listProfileTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListProfileTemplates(listProfileTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListProfileTemplates(listProfileTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProfileTemplates(listProfileTemplatesOptions *ListProfileTemplatesOptions)`, func() {
		listProfileTemplatesPath := "/v1/profile_templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfileTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 20, "first": "First", "previous": "Previous", "next": "Next", "profile_templates": [{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListProfileTemplates successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListProfileTemplatesOptions model
				listProfileTemplatesOptionsModel := new(iamidentityv1.ListProfileTemplatesOptions)
				listProfileTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listProfileTemplatesOptionsModel.Limit = core.StringPtr("20")
				listProfileTemplatesOptionsModel.Pagetoken = core.StringPtr("testString")
				listProfileTemplatesOptionsModel.Sort = core.StringPtr("created_at")
				listProfileTemplatesOptionsModel.Order = core.StringPtr("asc")
				listProfileTemplatesOptionsModel.IncludeHistory = core.StringPtr("false")
				listProfileTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListProfileTemplatesWithContext(ctx, listProfileTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListProfileTemplates(listProfileTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListProfileTemplatesWithContext(ctx, listProfileTemplatesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProfileTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 20, "first": "First", "previous": "Previous", "next": "Next", "profile_templates": [{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListProfileTemplates successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListProfileTemplates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProfileTemplatesOptions model
				listProfileTemplatesOptionsModel := new(iamidentityv1.ListProfileTemplatesOptions)
				listProfileTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listProfileTemplatesOptionsModel.Limit = core.StringPtr("20")
				listProfileTemplatesOptionsModel.Pagetoken = core.StringPtr("testString")
				listProfileTemplatesOptionsModel.Sort = core.StringPtr("created_at")
				listProfileTemplatesOptionsModel.Order = core.StringPtr("asc")
				listProfileTemplatesOptionsModel.IncludeHistory = core.StringPtr("false")
				listProfileTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListProfileTemplates(listProfileTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProfileTemplates with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListProfileTemplatesOptions model
				listProfileTemplatesOptionsModel := new(iamidentityv1.ListProfileTemplatesOptions)
				listProfileTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listProfileTemplatesOptionsModel.Limit = core.StringPtr("20")
				listProfileTemplatesOptionsModel.Pagetoken = core.StringPtr("testString")
				listProfileTemplatesOptionsModel.Sort = core.StringPtr("created_at")
				listProfileTemplatesOptionsModel.Order = core.StringPtr("asc")
				listProfileTemplatesOptionsModel.IncludeHistory = core.StringPtr("false")
				listProfileTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListProfileTemplates(listProfileTemplatesOptionsModel)
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
			It(`Invoke ListProfileTemplates successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListProfileTemplatesOptions model
				listProfileTemplatesOptionsModel := new(iamidentityv1.ListProfileTemplatesOptions)
				listProfileTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listProfileTemplatesOptionsModel.Limit = core.StringPtr("20")
				listProfileTemplatesOptionsModel.Pagetoken = core.StringPtr("testString")
				listProfileTemplatesOptionsModel.Sort = core.StringPtr("created_at")
				listProfileTemplatesOptionsModel.Order = core.StringPtr("asc")
				listProfileTemplatesOptionsModel.IncludeHistory = core.StringPtr("false")
				listProfileTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListProfileTemplates(listProfileTemplatesOptionsModel)
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
	Describe(`CreateProfileTemplate(createProfileTemplateOptions *CreateProfileTemplateOptions) - Operation response error`, func() {
		createProfileTemplatePath := "/v1/profile_templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProfileTemplatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProfileTemplate with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the CreateProfileTemplateOptions model
				createProfileTemplateOptionsModel := new(iamidentityv1.CreateProfileTemplateOptions)
				createProfileTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Name = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Description = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Profile = templateProfileComponentRequestModel
				createProfileTemplateOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				createProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateProfileTemplate(createProfileTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateProfileTemplate(createProfileTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProfileTemplate(createProfileTemplateOptions *CreateProfileTemplateOptions)`, func() {
		createProfileTemplatePath := "/v1/profile_templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProfileTemplatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateProfileTemplate successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the CreateProfileTemplateOptions model
				createProfileTemplateOptionsModel := new(iamidentityv1.CreateProfileTemplateOptions)
				createProfileTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Name = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Description = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Profile = templateProfileComponentRequestModel
				createProfileTemplateOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				createProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateProfileTemplateWithContext(ctx, createProfileTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateProfileTemplate(createProfileTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateProfileTemplateWithContext(ctx, createProfileTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createProfileTemplatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateProfileTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateProfileTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the CreateProfileTemplateOptions model
				createProfileTemplateOptionsModel := new(iamidentityv1.CreateProfileTemplateOptions)
				createProfileTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Name = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Description = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Profile = templateProfileComponentRequestModel
				createProfileTemplateOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				createProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateProfileTemplate(createProfileTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProfileTemplate with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the CreateProfileTemplateOptions model
				createProfileTemplateOptionsModel := new(iamidentityv1.CreateProfileTemplateOptions)
				createProfileTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Name = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Description = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Profile = templateProfileComponentRequestModel
				createProfileTemplateOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				createProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateProfileTemplate(createProfileTemplateOptionsModel)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateProfileTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the CreateProfileTemplateOptions model
				createProfileTemplateOptionsModel := new(iamidentityv1.CreateProfileTemplateOptions)
				createProfileTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Name = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Description = core.StringPtr("testString")
				createProfileTemplateOptionsModel.Profile = templateProfileComponentRequestModel
				createProfileTemplateOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				createProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateProfileTemplate(createProfileTemplateOptionsModel)
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
	Describe(`GetLatestProfileTemplateVersion(getLatestProfileTemplateVersionOptions *GetLatestProfileTemplateVersionOptions) - Operation response error`, func() {
		getLatestProfileTemplateVersionPath := "/v1/profile_templates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestProfileTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLatestProfileTemplateVersion with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetLatestProfileTemplateVersionOptions model
				getLatestProfileTemplateVersionOptionsModel := new(iamidentityv1.GetLatestProfileTemplateVersionOptions)
				getLatestProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestProfileTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getLatestProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetLatestProfileTemplateVersion(getLatestProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetLatestProfileTemplateVersion(getLatestProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLatestProfileTemplateVersion(getLatestProfileTemplateVersionOptions *GetLatestProfileTemplateVersionOptions)`, func() {
		getLatestProfileTemplateVersionPath := "/v1/profile_templates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestProfileTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetLatestProfileTemplateVersion successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetLatestProfileTemplateVersionOptions model
				getLatestProfileTemplateVersionOptionsModel := new(iamidentityv1.GetLatestProfileTemplateVersionOptions)
				getLatestProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestProfileTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getLatestProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetLatestProfileTemplateVersionWithContext(ctx, getLatestProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetLatestProfileTemplateVersion(getLatestProfileTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetLatestProfileTemplateVersionWithContext(ctx, getLatestProfileTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLatestProfileTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetLatestProfileTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetLatestProfileTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLatestProfileTemplateVersionOptions model
				getLatestProfileTemplateVersionOptionsModel := new(iamidentityv1.GetLatestProfileTemplateVersionOptions)
				getLatestProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestProfileTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getLatestProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetLatestProfileTemplateVersion(getLatestProfileTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLatestProfileTemplateVersion with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetLatestProfileTemplateVersionOptions model
				getLatestProfileTemplateVersionOptionsModel := new(iamidentityv1.GetLatestProfileTemplateVersionOptions)
				getLatestProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestProfileTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getLatestProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetLatestProfileTemplateVersion(getLatestProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLatestProfileTemplateVersionOptions model with no property values
				getLatestProfileTemplateVersionOptionsModelNew := new(iamidentityv1.GetLatestProfileTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetLatestProfileTemplateVersion(getLatestProfileTemplateVersionOptionsModelNew)
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
			It(`Invoke GetLatestProfileTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetLatestProfileTemplateVersionOptions model
				getLatestProfileTemplateVersionOptionsModel := new(iamidentityv1.GetLatestProfileTemplateVersionOptions)
				getLatestProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestProfileTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getLatestProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetLatestProfileTemplateVersion(getLatestProfileTemplateVersionOptionsModel)
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
	Describe(`DeleteAllVersionsOfProfileTemplate(deleteAllVersionsOfProfileTemplateOptions *DeleteAllVersionsOfProfileTemplateOptions)`, func() {
		deleteAllVersionsOfProfileTemplatePath := "/v1/profile_templates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllVersionsOfProfileTemplatePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAllVersionsOfProfileTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteAllVersionsOfProfileTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAllVersionsOfProfileTemplateOptions model
				deleteAllVersionsOfProfileTemplateOptionsModel := new(iamidentityv1.DeleteAllVersionsOfProfileTemplateOptions)
				deleteAllVersionsOfProfileTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				deleteAllVersionsOfProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteAllVersionsOfProfileTemplate(deleteAllVersionsOfProfileTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAllVersionsOfProfileTemplate with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteAllVersionsOfProfileTemplateOptions model
				deleteAllVersionsOfProfileTemplateOptionsModel := new(iamidentityv1.DeleteAllVersionsOfProfileTemplateOptions)
				deleteAllVersionsOfProfileTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				deleteAllVersionsOfProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteAllVersionsOfProfileTemplate(deleteAllVersionsOfProfileTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAllVersionsOfProfileTemplateOptions model with no property values
				deleteAllVersionsOfProfileTemplateOptionsModelNew := new(iamidentityv1.DeleteAllVersionsOfProfileTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteAllVersionsOfProfileTemplate(deleteAllVersionsOfProfileTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVersionsOfProfileTemplate(listVersionsOfProfileTemplateOptions *ListVersionsOfProfileTemplateOptions) - Operation response error`, func() {
		listVersionsOfProfileTemplatePath := "/v1/profile_templates/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVersionsOfProfileTemplatePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListVersionsOfProfileTemplate with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListVersionsOfProfileTemplateOptions model
				listVersionsOfProfileTemplateOptionsModel := new(iamidentityv1.ListVersionsOfProfileTemplateOptions)
				listVersionsOfProfileTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				listVersionsOfProfileTemplateOptionsModel.Limit = core.StringPtr("20")
				listVersionsOfProfileTemplateOptionsModel.Pagetoken = core.StringPtr("testString")
				listVersionsOfProfileTemplateOptionsModel.Sort = core.StringPtr("created_at")
				listVersionsOfProfileTemplateOptionsModel.Order = core.StringPtr("asc")
				listVersionsOfProfileTemplateOptionsModel.IncludeHistory = core.StringPtr("false")
				listVersionsOfProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListVersionsOfProfileTemplate(listVersionsOfProfileTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListVersionsOfProfileTemplate(listVersionsOfProfileTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVersionsOfProfileTemplate(listVersionsOfProfileTemplateOptions *ListVersionsOfProfileTemplateOptions)`, func() {
		listVersionsOfProfileTemplatePath := "/v1/profile_templates/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVersionsOfProfileTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 20, "first": "First", "previous": "Previous", "next": "Next", "profile_templates": [{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListVersionsOfProfileTemplate successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ListVersionsOfProfileTemplateOptions model
				listVersionsOfProfileTemplateOptionsModel := new(iamidentityv1.ListVersionsOfProfileTemplateOptions)
				listVersionsOfProfileTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				listVersionsOfProfileTemplateOptionsModel.Limit = core.StringPtr("20")
				listVersionsOfProfileTemplateOptionsModel.Pagetoken = core.StringPtr("testString")
				listVersionsOfProfileTemplateOptionsModel.Sort = core.StringPtr("created_at")
				listVersionsOfProfileTemplateOptionsModel.Order = core.StringPtr("asc")
				listVersionsOfProfileTemplateOptionsModel.IncludeHistory = core.StringPtr("false")
				listVersionsOfProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.ListVersionsOfProfileTemplateWithContext(ctx, listVersionsOfProfileTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.ListVersionsOfProfileTemplate(listVersionsOfProfileTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.ListVersionsOfProfileTemplateWithContext(ctx, listVersionsOfProfileTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listVersionsOfProfileTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"20"}))
					Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_at"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["include_history"]).To(Equal([]string{"false"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 20, "first": "First", "previous": "Previous", "next": "Next", "profile_templates": [{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListVersionsOfProfileTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListVersionsOfProfileTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVersionsOfProfileTemplateOptions model
				listVersionsOfProfileTemplateOptionsModel := new(iamidentityv1.ListVersionsOfProfileTemplateOptions)
				listVersionsOfProfileTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				listVersionsOfProfileTemplateOptionsModel.Limit = core.StringPtr("20")
				listVersionsOfProfileTemplateOptionsModel.Pagetoken = core.StringPtr("testString")
				listVersionsOfProfileTemplateOptionsModel.Sort = core.StringPtr("created_at")
				listVersionsOfProfileTemplateOptionsModel.Order = core.StringPtr("asc")
				listVersionsOfProfileTemplateOptionsModel.IncludeHistory = core.StringPtr("false")
				listVersionsOfProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListVersionsOfProfileTemplate(listVersionsOfProfileTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListVersionsOfProfileTemplate with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListVersionsOfProfileTemplateOptions model
				listVersionsOfProfileTemplateOptionsModel := new(iamidentityv1.ListVersionsOfProfileTemplateOptions)
				listVersionsOfProfileTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				listVersionsOfProfileTemplateOptionsModel.Limit = core.StringPtr("20")
				listVersionsOfProfileTemplateOptionsModel.Pagetoken = core.StringPtr("testString")
				listVersionsOfProfileTemplateOptionsModel.Sort = core.StringPtr("created_at")
				listVersionsOfProfileTemplateOptionsModel.Order = core.StringPtr("asc")
				listVersionsOfProfileTemplateOptionsModel.IncludeHistory = core.StringPtr("false")
				listVersionsOfProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListVersionsOfProfileTemplate(listVersionsOfProfileTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListVersionsOfProfileTemplateOptions model with no property values
				listVersionsOfProfileTemplateOptionsModelNew := new(iamidentityv1.ListVersionsOfProfileTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.ListVersionsOfProfileTemplate(listVersionsOfProfileTemplateOptionsModelNew)
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
			It(`Invoke ListVersionsOfProfileTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListVersionsOfProfileTemplateOptions model
				listVersionsOfProfileTemplateOptionsModel := new(iamidentityv1.ListVersionsOfProfileTemplateOptions)
				listVersionsOfProfileTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				listVersionsOfProfileTemplateOptionsModel.Limit = core.StringPtr("20")
				listVersionsOfProfileTemplateOptionsModel.Pagetoken = core.StringPtr("testString")
				listVersionsOfProfileTemplateOptionsModel.Sort = core.StringPtr("created_at")
				listVersionsOfProfileTemplateOptionsModel.Order = core.StringPtr("asc")
				listVersionsOfProfileTemplateOptionsModel.IncludeHistory = core.StringPtr("false")
				listVersionsOfProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.ListVersionsOfProfileTemplate(listVersionsOfProfileTemplateOptionsModel)
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
	Describe(`CreateProfileTemplateVersion(createProfileTemplateVersionOptions *CreateProfileTemplateVersionOptions) - Operation response error`, func() {
		createProfileTemplateVersionPath := "/v1/profile_templates/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProfileTemplateVersionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProfileTemplateVersion with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the CreateProfileTemplateVersionOptions model
				createProfileTemplateVersionOptionsModel := new(iamidentityv1.CreateProfileTemplateVersionOptions)
				createProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Profile = templateProfileComponentRequestModel
				createProfileTemplateVersionOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				createProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateProfileTemplateVersion(createProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateProfileTemplateVersion(createProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProfileTemplateVersion(createProfileTemplateVersionOptions *CreateProfileTemplateVersionOptions)`, func() {
		createProfileTemplateVersionPath := "/v1/profile_templates/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProfileTemplateVersionPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateProfileTemplateVersion successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the CreateProfileTemplateVersionOptions model
				createProfileTemplateVersionOptionsModel := new(iamidentityv1.CreateProfileTemplateVersionOptions)
				createProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Profile = templateProfileComponentRequestModel
				createProfileTemplateVersionOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				createProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.CreateProfileTemplateVersionWithContext(ctx, createProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.CreateProfileTemplateVersion(createProfileTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.CreateProfileTemplateVersionWithContext(ctx, createProfileTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createProfileTemplateVersionPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateProfileTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateProfileTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the CreateProfileTemplateVersionOptions model
				createProfileTemplateVersionOptionsModel := new(iamidentityv1.CreateProfileTemplateVersionOptions)
				createProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Profile = templateProfileComponentRequestModel
				createProfileTemplateVersionOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				createProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateProfileTemplateVersion(createProfileTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProfileTemplateVersion with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the CreateProfileTemplateVersionOptions model
				createProfileTemplateVersionOptionsModel := new(iamidentityv1.CreateProfileTemplateVersionOptions)
				createProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Profile = templateProfileComponentRequestModel
				createProfileTemplateVersionOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				createProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateProfileTemplateVersion(createProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProfileTemplateVersionOptions model with no property values
				createProfileTemplateVersionOptionsModelNew := new(iamidentityv1.CreateProfileTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateProfileTemplateVersion(createProfileTemplateVersionOptionsModelNew)
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
			It(`Invoke CreateProfileTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the CreateProfileTemplateVersionOptions model
				createProfileTemplateVersionOptionsModel := new(iamidentityv1.CreateProfileTemplateVersionOptions)
				createProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createProfileTemplateVersionOptionsModel.Profile = templateProfileComponentRequestModel
				createProfileTemplateVersionOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				createProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.CreateProfileTemplateVersion(createProfileTemplateVersionOptionsModel)
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
	Describe(`GetProfileTemplateVersion(getProfileTemplateVersionOptions *GetProfileTemplateVersionOptions) - Operation response error`, func() {
		getProfileTemplateVersionPath := "/v1/profile_templates/testString/versions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfileTemplateVersion with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileTemplateVersionOptions model
				getProfileTemplateVersionOptionsModel := new(iamidentityv1.GetProfileTemplateVersionOptions)
				getProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getProfileTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetProfileTemplateVersion(getProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetProfileTemplateVersion(getProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfileTemplateVersion(getProfileTemplateVersionOptions *GetProfileTemplateVersionOptions)`, func() {
		getProfileTemplateVersionPath := "/v1/profile_templates/testString/versions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetProfileTemplateVersion successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileTemplateVersionOptions model
				getProfileTemplateVersionOptionsModel := new(iamidentityv1.GetProfileTemplateVersionOptions)
				getProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getProfileTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.GetProfileTemplateVersionWithContext(ctx, getProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.GetProfileTemplateVersion(getProfileTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.GetProfileTemplateVersionWithContext(ctx, getProfileTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProfileTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetProfileTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetProfileTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileTemplateVersionOptions model
				getProfileTemplateVersionOptionsModel := new(iamidentityv1.GetProfileTemplateVersionOptions)
				getProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getProfileTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetProfileTemplateVersion(getProfileTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfileTemplateVersion with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileTemplateVersionOptions model
				getProfileTemplateVersionOptionsModel := new(iamidentityv1.GetProfileTemplateVersionOptions)
				getProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getProfileTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetProfileTemplateVersion(getProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileTemplateVersionOptions model with no property values
				getProfileTemplateVersionOptionsModelNew := new(iamidentityv1.GetProfileTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetProfileTemplateVersion(getProfileTemplateVersionOptionsModelNew)
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
			It(`Invoke GetProfileTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetProfileTemplateVersionOptions model
				getProfileTemplateVersionOptionsModel := new(iamidentityv1.GetProfileTemplateVersionOptions)
				getProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getProfileTemplateVersionOptionsModel.IncludeHistory = core.BoolPtr(false)
				getProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.GetProfileTemplateVersion(getProfileTemplateVersionOptionsModel)
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
	Describe(`UpdateProfileTemplateVersion(updateProfileTemplateVersionOptions *UpdateProfileTemplateVersionOptions) - Operation response error`, func() {
		updateProfileTemplateVersionPath := "/v1/profile_templates/testString/versions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProfileTemplateVersionPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProfileTemplateVersion with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the UpdateProfileTemplateVersionOptions model
				updateProfileTemplateVersionOptionsModel := new(iamidentityv1.UpdateProfileTemplateVersionOptions)
				updateProfileTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Profile = templateProfileComponentRequestModel
				updateProfileTemplateVersionOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				updateProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateProfileTemplateVersion(updateProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateProfileTemplateVersion(updateProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProfileTemplateVersion(updateProfileTemplateVersionOptions *UpdateProfileTemplateVersionOptions)`, func() {
		updateProfileTemplateVersionPath := "/v1/profile_templates/testString/versions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProfileTemplateVersionPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke UpdateProfileTemplateVersion successfully with retries`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the UpdateProfileTemplateVersionOptions model
				updateProfileTemplateVersionOptionsModel := new(iamidentityv1.UpdateProfileTemplateVersionOptions)
				updateProfileTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Profile = templateProfileComponentRequestModel
				updateProfileTemplateVersionOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				updateProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamIdentityService.UpdateProfileTemplateVersionWithContext(ctx, updateProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr := iamIdentityService.UpdateProfileTemplateVersion(updateProfileTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamIdentityService.UpdateProfileTemplateVersionWithContext(ctx, updateProfileTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProfileTemplateVersionPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "version": 7, "account_id": "AccountID", "name": "Name", "description": "Description", "committed": false, "profile": {"name": "Name", "description": "Description", "rules": [{"name": "Name", "type": "Profile-SAML", "realm_name": "RealmName", "expiration": 10, "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}]}], "identities": [{"iam_id": "IamID", "identifier": "Identifier", "type": "user", "accounts": ["Accounts"], "description": "Description"}]}, "policy_template_references": [{"id": "ID", "version": "Version"}], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "entity_tag": "EntityTag", "crn": "CRN", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke UpdateProfileTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateProfileTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the UpdateProfileTemplateVersionOptions model
				updateProfileTemplateVersionOptionsModel := new(iamidentityv1.UpdateProfileTemplateVersionOptions)
				updateProfileTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Profile = templateProfileComponentRequestModel
				updateProfileTemplateVersionOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				updateProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateProfileTemplateVersion(updateProfileTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProfileTemplateVersion with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the UpdateProfileTemplateVersionOptions model
				updateProfileTemplateVersionOptionsModel := new(iamidentityv1.UpdateProfileTemplateVersionOptions)
				updateProfileTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Profile = templateProfileComponentRequestModel
				updateProfileTemplateVersionOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				updateProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateProfileTemplateVersion(updateProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProfileTemplateVersionOptions model with no property values
				updateProfileTemplateVersionOptionsModelNew := new(iamidentityv1.UpdateProfileTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateProfileTemplateVersion(updateProfileTemplateVersionOptionsModelNew)
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
			It(`Invoke UpdateProfileTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")

				// Construct an instance of the UpdateProfileTemplateVersionOptions model
				updateProfileTemplateVersionOptionsModel := new(iamidentityv1.UpdateProfileTemplateVersionOptions)
				updateProfileTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.AccountID = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				updateProfileTemplateVersionOptionsModel.Profile = templateProfileComponentRequestModel
				updateProfileTemplateVersionOptionsModel.PolicyTemplateReferences = []iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}
				updateProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamIdentityService.UpdateProfileTemplateVersion(updateProfileTemplateVersionOptionsModel)
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
	Describe(`DeleteProfileTemplateVersion(deleteProfileTemplateVersionOptions *DeleteProfileTemplateVersionOptions)`, func() {
		deleteProfileTemplateVersionPath := "/v1/profile_templates/testString/versions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProfileTemplateVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteProfileTemplateVersion successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteProfileTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProfileTemplateVersionOptions model
				deleteProfileTemplateVersionOptionsModel := new(iamidentityv1.DeleteProfileTemplateVersionOptions)
				deleteProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				deleteProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				deleteProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteProfileTemplateVersion(deleteProfileTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProfileTemplateVersion with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileTemplateVersionOptions model
				deleteProfileTemplateVersionOptionsModel := new(iamidentityv1.DeleteProfileTemplateVersionOptions)
				deleteProfileTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				deleteProfileTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				deleteProfileTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteProfileTemplateVersion(deleteProfileTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProfileTemplateVersionOptions model with no property values
				deleteProfileTemplateVersionOptionsModelNew := new(iamidentityv1.DeleteProfileTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteProfileTemplateVersion(deleteProfileTemplateVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CommitProfileTemplate(commitProfileTemplateOptions *CommitProfileTemplateOptions)`, func() {
		commitProfileTemplatePath := "/v1/profile_templates/testString/versions/testString/commit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(commitProfileTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke CommitProfileTemplate successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.CommitProfileTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CommitProfileTemplateOptions model
				commitProfileTemplateOptionsModel := new(iamidentityv1.CommitProfileTemplateOptions)
				commitProfileTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				commitProfileTemplateOptionsModel.Version = core.StringPtr("testString")
				commitProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.CommitProfileTemplate(commitProfileTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CommitProfileTemplate with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CommitProfileTemplateOptions model
				commitProfileTemplateOptionsModel := new(iamidentityv1.CommitProfileTemplateOptions)
				commitProfileTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				commitProfileTemplateOptionsModel.Version = core.StringPtr("testString")
				commitProfileTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.CommitProfileTemplate(commitProfileTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CommitProfileTemplateOptions model with no property values
				commitProfileTemplateOptionsModelNew := new(iamidentityv1.CommitProfileTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.CommitProfileTemplate(commitProfileTemplateOptionsModelNew)
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
			iamIdentityService, _ := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
				URL:           "http://iamidentityv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAccountSettingsUserMfa successfully`, func() {
				iamID := "testString"
				mfa := "NONE"
				_model, err := iamIdentityService.NewAccountSettingsUserMfa(iamID, mfa)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAPIKeyInsideCreateServiceIDRequest successfully`, func() {
				name := "testString"
				_model, err := iamIdentityService.NewAPIKeyInsideCreateServiceIDRequest(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCommitAccountSettingsTemplateOptions successfully`, func() {
				// Construct an instance of the CommitAccountSettingsTemplateOptions model
				templateID := "testString"
				version := "testString"
				commitAccountSettingsTemplateOptionsModel := iamIdentityService.NewCommitAccountSettingsTemplateOptions(templateID, version)
				commitAccountSettingsTemplateOptionsModel.SetTemplateID("testString")
				commitAccountSettingsTemplateOptionsModel.SetVersion("testString")
				commitAccountSettingsTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(commitAccountSettingsTemplateOptionsModel).ToNot(BeNil())
				Expect(commitAccountSettingsTemplateOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(commitAccountSettingsTemplateOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(commitAccountSettingsTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCommitProfileTemplateOptions successfully`, func() {
				// Construct an instance of the CommitProfileTemplateOptions model
				templateID := "testString"
				version := "testString"
				commitProfileTemplateOptionsModel := iamIdentityService.NewCommitProfileTemplateOptions(templateID, version)
				commitProfileTemplateOptionsModel.SetTemplateID("testString")
				commitProfileTemplateOptionsModel.SetVersion("testString")
				commitProfileTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(commitProfileTemplateOptionsModel).ToNot(BeNil())
				Expect(commitProfileTemplateOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(commitProfileTemplateOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(commitProfileTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAccountSettingsAssignmentOptions successfully`, func() {
				// Construct an instance of the CreateAccountSettingsAssignmentOptions model
				createAccountSettingsAssignmentOptionsTemplateID := "testString"
				createAccountSettingsAssignmentOptionsTemplateVersion := int64(1)
				createAccountSettingsAssignmentOptionsTargetType := "Account"
				createAccountSettingsAssignmentOptionsTarget := "testString"
				createAccountSettingsAssignmentOptionsModel := iamIdentityService.NewCreateAccountSettingsAssignmentOptions(createAccountSettingsAssignmentOptionsTemplateID, createAccountSettingsAssignmentOptionsTemplateVersion, createAccountSettingsAssignmentOptionsTargetType, createAccountSettingsAssignmentOptionsTarget)
				createAccountSettingsAssignmentOptionsModel.SetTemplateID("testString")
				createAccountSettingsAssignmentOptionsModel.SetTemplateVersion(int64(1))
				createAccountSettingsAssignmentOptionsModel.SetTargetType("Account")
				createAccountSettingsAssignmentOptionsModel.SetTarget("testString")
				createAccountSettingsAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccountSettingsAssignmentOptionsModel).ToNot(BeNil())
				Expect(createAccountSettingsAssignmentOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(createAccountSettingsAssignmentOptionsModel.TemplateVersion).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createAccountSettingsAssignmentOptionsModel.TargetType).To(Equal(core.StringPtr("Account")))
				Expect(createAccountSettingsAssignmentOptionsModel.Target).To(Equal(core.StringPtr("testString")))
				Expect(createAccountSettingsAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAccountSettingsTemplateOptions successfully`, func() {
				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				Expect(accountSettingsUserMfaModel).ToNot(BeNil())
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")
				Expect(accountSettingsUserMfaModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(accountSettingsUserMfaModel.Mfa).To(Equal(core.StringPtr("NONE")))

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				Expect(accountSettingsComponentModel).ToNot(BeNil())
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")
				Expect(accountSettingsComponentModel.RestrictCreateServiceID).To(Equal(core.StringPtr("NOT_SET")))
				Expect(accountSettingsComponentModel.RestrictCreatePlatformApikey).To(Equal(core.StringPtr("NOT_SET")))
				Expect(accountSettingsComponentModel.AllowedIPAddresses).To(Equal(core.StringPtr("testString")))
				Expect(accountSettingsComponentModel.Mfa).To(Equal(core.StringPtr("NONE")))
				Expect(accountSettingsComponentModel.UserMfa).To(Equal([]iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}))
				Expect(accountSettingsComponentModel.SessionExpirationInSeconds).To(Equal(core.StringPtr("86400")))
				Expect(accountSettingsComponentModel.SessionInvalidationInSeconds).To(Equal(core.StringPtr("7200")))
				Expect(accountSettingsComponentModel.MaxSessionsPerIdentity).To(Equal(core.StringPtr("testString")))
				Expect(accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds).To(Equal(core.StringPtr("3600")))
				Expect(accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds).To(Equal(core.StringPtr("259200")))

				// Construct an instance of the CreateAccountSettingsTemplateOptions model
				createAccountSettingsTemplateOptionsModel := iamIdentityService.NewCreateAccountSettingsTemplateOptions()
				createAccountSettingsTemplateOptionsModel.SetAccountID("testString")
				createAccountSettingsTemplateOptionsModel.SetName("testString")
				createAccountSettingsTemplateOptionsModel.SetDescription("testString")
				createAccountSettingsTemplateOptionsModel.SetAccountSettings(accountSettingsComponentModel)
				createAccountSettingsTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccountSettingsTemplateOptionsModel).ToNot(BeNil())
				Expect(createAccountSettingsTemplateOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createAccountSettingsTemplateOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createAccountSettingsTemplateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createAccountSettingsTemplateOptionsModel.AccountSettings).To(Equal(accountSettingsComponentModel))
				Expect(createAccountSettingsTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAccountSettingsTemplateVersionOptions successfully`, func() {
				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				Expect(accountSettingsUserMfaModel).ToNot(BeNil())
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")
				Expect(accountSettingsUserMfaModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(accountSettingsUserMfaModel.Mfa).To(Equal(core.StringPtr("NONE")))

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				Expect(accountSettingsComponentModel).ToNot(BeNil())
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")
				Expect(accountSettingsComponentModel.RestrictCreateServiceID).To(Equal(core.StringPtr("NOT_SET")))
				Expect(accountSettingsComponentModel.RestrictCreatePlatformApikey).To(Equal(core.StringPtr("NOT_SET")))
				Expect(accountSettingsComponentModel.AllowedIPAddresses).To(Equal(core.StringPtr("testString")))
				Expect(accountSettingsComponentModel.Mfa).To(Equal(core.StringPtr("NONE")))
				Expect(accountSettingsComponentModel.UserMfa).To(Equal([]iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}))
				Expect(accountSettingsComponentModel.SessionExpirationInSeconds).To(Equal(core.StringPtr("86400")))
				Expect(accountSettingsComponentModel.SessionInvalidationInSeconds).To(Equal(core.StringPtr("7200")))
				Expect(accountSettingsComponentModel.MaxSessionsPerIdentity).To(Equal(core.StringPtr("testString")))
				Expect(accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds).To(Equal(core.StringPtr("3600")))
				Expect(accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds).To(Equal(core.StringPtr("259200")))

				// Construct an instance of the CreateAccountSettingsTemplateVersionOptions model
				templateID := "testString"
				createAccountSettingsTemplateVersionOptionsModel := iamIdentityService.NewCreateAccountSettingsTemplateVersionOptions(templateID)
				createAccountSettingsTemplateVersionOptionsModel.SetTemplateID("testString")
				createAccountSettingsTemplateVersionOptionsModel.SetAccountID("testString")
				createAccountSettingsTemplateVersionOptionsModel.SetName("testString")
				createAccountSettingsTemplateVersionOptionsModel.SetDescription("testString")
				createAccountSettingsTemplateVersionOptionsModel.SetAccountSettings(accountSettingsComponentModel)
				createAccountSettingsTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccountSettingsTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(createAccountSettingsTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(createAccountSettingsTemplateVersionOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createAccountSettingsTemplateVersionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createAccountSettingsTemplateVersionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createAccountSettingsTemplateVersionOptionsModel.AccountSettings).To(Equal(accountSettingsComponentModel))
				Expect(createAccountSettingsTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAPIKeyOptions successfully`, func() {
				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsName := "testString"
				createAPIKeyOptionsIamID := "testString"
				createAPIKeyOptionsModel := iamIdentityService.NewCreateAPIKeyOptions(createAPIKeyOptionsName, createAPIKeyOptionsIamID)
				createAPIKeyOptionsModel.SetName("testString")
				createAPIKeyOptionsModel.SetIamID("testString")
				createAPIKeyOptionsModel.SetDescription("testString")
				createAPIKeyOptionsModel.SetAccountID("testString")
				createAPIKeyOptionsModel.SetApikey("testString")
				createAPIKeyOptionsModel.SetStoreValue(true)
				createAPIKeyOptionsModel.SetEntityLock("false")
				createAPIKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAPIKeyOptionsModel).ToNot(BeNil())
				Expect(createAPIKeyOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.Apikey).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.StoreValue).To(Equal(core.BoolPtr(true)))
				Expect(createAPIKeyOptionsModel.EntityLock).To(Equal(core.StringPtr("false")))
				Expect(createAPIKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateClaimRuleOptions successfully`, func() {
				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				Expect(profileClaimRuleConditionsModel).ToNot(BeNil())
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")
				Expect(profileClaimRuleConditionsModel.Claim).To(Equal(core.StringPtr("testString")))
				Expect(profileClaimRuleConditionsModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(profileClaimRuleConditionsModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				Expect(responseContextModel).ToNot(BeNil())
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")
				Expect(responseContextModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.Operation).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.UserAgent).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.ThreadID).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.Host).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.StartTime).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.EndTime).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.ElapsedTime).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.ClusterName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateClaimRuleOptions model
				profileID := "testString"
				createClaimRuleOptionsType := "testString"
				createClaimRuleOptionsConditions := []iamidentityv1.ProfileClaimRuleConditions{}
				createClaimRuleOptionsModel := iamIdentityService.NewCreateClaimRuleOptions(profileID, createClaimRuleOptionsType, createClaimRuleOptionsConditions)
				createClaimRuleOptionsModel.SetProfileID("testString")
				createClaimRuleOptionsModel.SetType("testString")
				createClaimRuleOptionsModel.SetConditions([]iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel})
				createClaimRuleOptionsModel.SetContext(responseContextModel)
				createClaimRuleOptionsModel.SetName("testString")
				createClaimRuleOptionsModel.SetRealmName("testString")
				createClaimRuleOptionsModel.SetCrType("testString")
				createClaimRuleOptionsModel.SetExpiration(int64(38))
				createClaimRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createClaimRuleOptionsModel).ToNot(BeNil())
				Expect(createClaimRuleOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(createClaimRuleOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(createClaimRuleOptionsModel.Conditions).To(Equal([]iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}))
				Expect(createClaimRuleOptionsModel.Context).To(Equal(responseContextModel))
				Expect(createClaimRuleOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createClaimRuleOptionsModel.RealmName).To(Equal(core.StringPtr("testString")))
				Expect(createClaimRuleOptionsModel.CrType).To(Equal(core.StringPtr("testString")))
				Expect(createClaimRuleOptionsModel.Expiration).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createClaimRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateLinkOptions successfully`, func() {
				// Construct an instance of the CreateProfileLinkRequestLink model
				createProfileLinkRequestLinkModel := new(iamidentityv1.CreateProfileLinkRequestLink)
				Expect(createProfileLinkRequestLinkModel).ToNot(BeNil())
				createProfileLinkRequestLinkModel.CRN = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Namespace = core.StringPtr("testString")
				createProfileLinkRequestLinkModel.Name = core.StringPtr("testString")
				Expect(createProfileLinkRequestLinkModel.CRN).To(Equal(core.StringPtr("testString")))
				Expect(createProfileLinkRequestLinkModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(createProfileLinkRequestLinkModel.Name).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateLinkOptions model
				profileID := "testString"
				createLinkOptionsCrType := "testString"
				var createLinkOptionsLink *iamidentityv1.CreateProfileLinkRequestLink = nil
				createLinkOptionsModel := iamIdentityService.NewCreateLinkOptions(profileID, createLinkOptionsCrType, createLinkOptionsLink)
				createLinkOptionsModel.SetProfileID("testString")
				createLinkOptionsModel.SetCrType("testString")
				createLinkOptionsModel.SetLink(createProfileLinkRequestLinkModel)
				createLinkOptionsModel.SetName("testString")
				createLinkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLinkOptionsModel).ToNot(BeNil())
				Expect(createLinkOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(createLinkOptionsModel.CrType).To(Equal(core.StringPtr("testString")))
				Expect(createLinkOptionsModel.Link).To(Equal(createProfileLinkRequestLinkModel))
				Expect(createLinkOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createLinkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateMfaReportOptions successfully`, func() {
				// Construct an instance of the CreateMfaReportOptions model
				accountID := "testString"
				createMfaReportOptionsModel := iamIdentityService.NewCreateMfaReportOptions(accountID)
				createMfaReportOptionsModel.SetAccountID("testString")
				createMfaReportOptionsModel.SetType("testString")
				createMfaReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createMfaReportOptionsModel).ToNot(BeNil())
				Expect(createMfaReportOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createMfaReportOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(createMfaReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProfileLinkRequestLink successfully`, func() {
				crn := "testString"
				namespace := "testString"
				_model, err := iamIdentityService.NewCreateProfileLinkRequestLink(crn, namespace)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateProfileOptions successfully`, func() {
				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsName := "testString"
				createProfileOptionsAccountID := "testString"
				createProfileOptionsModel := iamIdentityService.NewCreateProfileOptions(createProfileOptionsName, createProfileOptionsAccountID)
				createProfileOptionsModel.SetName("testString")
				createProfileOptionsModel.SetAccountID("testString")
				createProfileOptionsModel.SetDescription("testString")
				createProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProfileOptionsModel).ToNot(BeNil())
				Expect(createProfileOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProfileTemplateOptions successfully`, func() {
				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				Expect(profileClaimRuleConditionsModel).ToNot(BeNil())
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")
				Expect(profileClaimRuleConditionsModel.Claim).To(Equal(core.StringPtr("testString")))
				Expect(profileClaimRuleConditionsModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(profileClaimRuleConditionsModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				Expect(trustedProfileTemplateClaimRuleModel).ToNot(BeNil())
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				Expect(trustedProfileTemplateClaimRuleModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(trustedProfileTemplateClaimRuleModel.Type).To(Equal(core.StringPtr("Profile-SAML")))
				Expect(trustedProfileTemplateClaimRuleModel.RealmName).To(Equal(core.StringPtr("testString")))
				Expect(trustedProfileTemplateClaimRuleModel.Expiration).To(Equal(core.Int64Ptr(int64(38))))
				Expect(trustedProfileTemplateClaimRuleModel.Conditions).To(Equal([]iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}))

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				Expect(profileIdentityRequestModel).ToNot(BeNil())
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")
				Expect(profileIdentityRequestModel.Identifier).To(Equal(core.StringPtr("testString")))
				Expect(profileIdentityRequestModel.Type).To(Equal(core.StringPtr("user")))
				Expect(profileIdentityRequestModel.Accounts).To(Equal([]string{"testString"}))
				Expect(profileIdentityRequestModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				Expect(templateProfileComponentRequestModel).ToNot(BeNil())
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}
				Expect(templateProfileComponentRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(templateProfileComponentRequestModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(templateProfileComponentRequestModel.Rules).To(Equal([]iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}))
				Expect(templateProfileComponentRequestModel.Identities).To(Equal([]iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}))

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				Expect(policyTemplateReferenceModel).ToNot(BeNil())
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")
				Expect(policyTemplateReferenceModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(policyTemplateReferenceModel.Version).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateProfileTemplateOptions model
				createProfileTemplateOptionsModel := iamIdentityService.NewCreateProfileTemplateOptions()
				createProfileTemplateOptionsModel.SetAccountID("testString")
				createProfileTemplateOptionsModel.SetName("testString")
				createProfileTemplateOptionsModel.SetDescription("testString")
				createProfileTemplateOptionsModel.SetProfile(templateProfileComponentRequestModel)
				createProfileTemplateOptionsModel.SetPolicyTemplateReferences([]iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel})
				createProfileTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProfileTemplateOptionsModel).ToNot(BeNil())
				Expect(createProfileTemplateOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileTemplateOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createProfileTemplateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createProfileTemplateOptionsModel.Profile).To(Equal(templateProfileComponentRequestModel))
				Expect(createProfileTemplateOptionsModel.PolicyTemplateReferences).To(Equal([]iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}))
				Expect(createProfileTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProfileTemplateVersionOptions successfully`, func() {
				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				Expect(profileClaimRuleConditionsModel).ToNot(BeNil())
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")
				Expect(profileClaimRuleConditionsModel.Claim).To(Equal(core.StringPtr("testString")))
				Expect(profileClaimRuleConditionsModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(profileClaimRuleConditionsModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				Expect(trustedProfileTemplateClaimRuleModel).ToNot(BeNil())
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				Expect(trustedProfileTemplateClaimRuleModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(trustedProfileTemplateClaimRuleModel.Type).To(Equal(core.StringPtr("Profile-SAML")))
				Expect(trustedProfileTemplateClaimRuleModel.RealmName).To(Equal(core.StringPtr("testString")))
				Expect(trustedProfileTemplateClaimRuleModel.Expiration).To(Equal(core.Int64Ptr(int64(38))))
				Expect(trustedProfileTemplateClaimRuleModel.Conditions).To(Equal([]iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}))

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				Expect(profileIdentityRequestModel).ToNot(BeNil())
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")
				Expect(profileIdentityRequestModel.Identifier).To(Equal(core.StringPtr("testString")))
				Expect(profileIdentityRequestModel.Type).To(Equal(core.StringPtr("user")))
				Expect(profileIdentityRequestModel.Accounts).To(Equal([]string{"testString"}))
				Expect(profileIdentityRequestModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				Expect(templateProfileComponentRequestModel).ToNot(BeNil())
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}
				Expect(templateProfileComponentRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(templateProfileComponentRequestModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(templateProfileComponentRequestModel.Rules).To(Equal([]iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}))
				Expect(templateProfileComponentRequestModel.Identities).To(Equal([]iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}))

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				Expect(policyTemplateReferenceModel).ToNot(BeNil())
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")
				Expect(policyTemplateReferenceModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(policyTemplateReferenceModel.Version).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateProfileTemplateVersionOptions model
				templateID := "testString"
				createProfileTemplateVersionOptionsModel := iamIdentityService.NewCreateProfileTemplateVersionOptions(templateID)
				createProfileTemplateVersionOptionsModel.SetTemplateID("testString")
				createProfileTemplateVersionOptionsModel.SetAccountID("testString")
				createProfileTemplateVersionOptionsModel.SetName("testString")
				createProfileTemplateVersionOptionsModel.SetDescription("testString")
				createProfileTemplateVersionOptionsModel.SetProfile(templateProfileComponentRequestModel)
				createProfileTemplateVersionOptionsModel.SetPolicyTemplateReferences([]iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel})
				createProfileTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProfileTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(createProfileTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileTemplateVersionOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileTemplateVersionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createProfileTemplateVersionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createProfileTemplateVersionOptionsModel.Profile).To(Equal(templateProfileComponentRequestModel))
				Expect(createProfileTemplateVersionOptionsModel.PolicyTemplateReferences).To(Equal([]iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}))
				Expect(createProfileTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateReportOptions successfully`, func() {
				// Construct an instance of the CreateReportOptions model
				accountID := "testString"
				createReportOptionsModel := iamIdentityService.NewCreateReportOptions(accountID)
				createReportOptionsModel.SetAccountID("testString")
				createReportOptionsModel.SetType("inactive")
				createReportOptionsModel.SetDuration("720")
				createReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createReportOptionsModel).ToNot(BeNil())
				Expect(createReportOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createReportOptionsModel.Type).To(Equal(core.StringPtr("inactive")))
				Expect(createReportOptionsModel.Duration).To(Equal(core.StringPtr("720")))
				Expect(createReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateServiceIDOptions successfully`, func() {
				// Construct an instance of the APIKeyInsideCreateServiceIDRequest model
				apiKeyInsideCreateServiceIDRequestModel := new(iamidentityv1.APIKeyInsideCreateServiceIDRequest)
				Expect(apiKeyInsideCreateServiceIDRequestModel).ToNot(BeNil())
				apiKeyInsideCreateServiceIDRequestModel.Name = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Description = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.Apikey = core.StringPtr("testString")
				apiKeyInsideCreateServiceIDRequestModel.StoreValue = core.BoolPtr(true)
				Expect(apiKeyInsideCreateServiceIDRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(apiKeyInsideCreateServiceIDRequestModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(apiKeyInsideCreateServiceIDRequestModel.Apikey).To(Equal(core.StringPtr("testString")))
				Expect(apiKeyInsideCreateServiceIDRequestModel.StoreValue).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreateServiceIDOptions model
				createServiceIDOptionsAccountID := "testString"
				createServiceIDOptionsName := "testString"
				createServiceIDOptionsModel := iamIdentityService.NewCreateServiceIDOptions(createServiceIDOptionsAccountID, createServiceIDOptionsName)
				createServiceIDOptionsModel.SetAccountID("testString")
				createServiceIDOptionsModel.SetName("testString")
				createServiceIDOptionsModel.SetDescription("testString")
				createServiceIDOptionsModel.SetUniqueInstanceCrns([]string{"testString"})
				createServiceIDOptionsModel.SetApikey(apiKeyInsideCreateServiceIDRequestModel)
				createServiceIDOptionsModel.SetEntityLock("false")
				createServiceIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createServiceIDOptionsModel).ToNot(BeNil())
				Expect(createServiceIDOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createServiceIDOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createServiceIDOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createServiceIDOptionsModel.UniqueInstanceCrns).To(Equal([]string{"testString"}))
				Expect(createServiceIDOptionsModel.Apikey).To(Equal(apiKeyInsideCreateServiceIDRequestModel))
				Expect(createServiceIDOptionsModel.EntityLock).To(Equal(core.StringPtr("false")))
				Expect(createServiceIDOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTrustedProfileAssignmentOptions successfully`, func() {
				// Construct an instance of the CreateTrustedProfileAssignmentOptions model
				createTrustedProfileAssignmentOptionsTemplateID := "testString"
				createTrustedProfileAssignmentOptionsTemplateVersion := int64(1)
				createTrustedProfileAssignmentOptionsTargetType := "Account"
				createTrustedProfileAssignmentOptionsTarget := "testString"
				createTrustedProfileAssignmentOptionsModel := iamIdentityService.NewCreateTrustedProfileAssignmentOptions(createTrustedProfileAssignmentOptionsTemplateID, createTrustedProfileAssignmentOptionsTemplateVersion, createTrustedProfileAssignmentOptionsTargetType, createTrustedProfileAssignmentOptionsTarget)
				createTrustedProfileAssignmentOptionsModel.SetTemplateID("testString")
				createTrustedProfileAssignmentOptionsModel.SetTemplateVersion(int64(1))
				createTrustedProfileAssignmentOptionsModel.SetTargetType("Account")
				createTrustedProfileAssignmentOptionsModel.SetTarget("testString")
				createTrustedProfileAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTrustedProfileAssignmentOptionsModel).ToNot(BeNil())
				Expect(createTrustedProfileAssignmentOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(createTrustedProfileAssignmentOptionsModel.TemplateVersion).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createTrustedProfileAssignmentOptionsModel.TargetType).To(Equal(core.StringPtr("Account")))
				Expect(createTrustedProfileAssignmentOptionsModel.Target).To(Equal(core.StringPtr("testString")))
				Expect(createTrustedProfileAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccountSettingsAssignmentOptions successfully`, func() {
				// Construct an instance of the DeleteAccountSettingsAssignmentOptions model
				assignmentID := "testString"
				deleteAccountSettingsAssignmentOptionsModel := iamIdentityService.NewDeleteAccountSettingsAssignmentOptions(assignmentID)
				deleteAccountSettingsAssignmentOptionsModel.SetAssignmentID("testString")
				deleteAccountSettingsAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccountSettingsAssignmentOptionsModel).ToNot(BeNil())
				Expect(deleteAccountSettingsAssignmentOptionsModel.AssignmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccountSettingsAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccountSettingsTemplateVersionOptions successfully`, func() {
				// Construct an instance of the DeleteAccountSettingsTemplateVersionOptions model
				templateID := "testString"
				version := "testString"
				deleteAccountSettingsTemplateVersionOptionsModel := iamIdentityService.NewDeleteAccountSettingsTemplateVersionOptions(templateID, version)
				deleteAccountSettingsTemplateVersionOptionsModel.SetTemplateID("testString")
				deleteAccountSettingsTemplateVersionOptionsModel.SetVersion("testString")
				deleteAccountSettingsTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccountSettingsTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(deleteAccountSettingsTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccountSettingsTemplateVersionOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccountSettingsTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAllVersionsOfAccountSettingsTemplateOptions successfully`, func() {
				// Construct an instance of the DeleteAllVersionsOfAccountSettingsTemplateOptions model
				templateID := "testString"
				deleteAllVersionsOfAccountSettingsTemplateOptionsModel := iamIdentityService.NewDeleteAllVersionsOfAccountSettingsTemplateOptions(templateID)
				deleteAllVersionsOfAccountSettingsTemplateOptionsModel.SetTemplateID("testString")
				deleteAllVersionsOfAccountSettingsTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAllVersionsOfAccountSettingsTemplateOptionsModel).ToNot(BeNil())
				Expect(deleteAllVersionsOfAccountSettingsTemplateOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAllVersionsOfAccountSettingsTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAllVersionsOfProfileTemplateOptions successfully`, func() {
				// Construct an instance of the DeleteAllVersionsOfProfileTemplateOptions model
				templateID := "testString"
				deleteAllVersionsOfProfileTemplateOptionsModel := iamIdentityService.NewDeleteAllVersionsOfProfileTemplateOptions(templateID)
				deleteAllVersionsOfProfileTemplateOptionsModel.SetTemplateID("testString")
				deleteAllVersionsOfProfileTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAllVersionsOfProfileTemplateOptionsModel).ToNot(BeNil())
				Expect(deleteAllVersionsOfProfileTemplateOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAllVersionsOfProfileTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAPIKeyOptions successfully`, func() {
				// Construct an instance of the DeleteAPIKeyOptions model
				id := "testString"
				deleteAPIKeyOptionsModel := iamIdentityService.NewDeleteAPIKeyOptions(id)
				deleteAPIKeyOptionsModel.SetID("testString")
				deleteAPIKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAPIKeyOptionsModel).ToNot(BeNil())
				Expect(deleteAPIKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAPIKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteClaimRuleOptions successfully`, func() {
				// Construct an instance of the DeleteClaimRuleOptions model
				profileID := "testString"
				ruleID := "testString"
				deleteClaimRuleOptionsModel := iamIdentityService.NewDeleteClaimRuleOptions(profileID, ruleID)
				deleteClaimRuleOptionsModel.SetProfileID("testString")
				deleteClaimRuleOptionsModel.SetRuleID("testString")
				deleteClaimRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteClaimRuleOptionsModel).ToNot(BeNil())
				Expect(deleteClaimRuleOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(deleteClaimRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteClaimRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLinkOptions successfully`, func() {
				// Construct an instance of the DeleteLinkOptions model
				profileID := "testString"
				linkID := "testString"
				deleteLinkOptionsModel := iamIdentityService.NewDeleteLinkOptions(profileID, linkID)
				deleteLinkOptionsModel.SetProfileID("testString")
				deleteLinkOptionsModel.SetLinkID("testString")
				deleteLinkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLinkOptionsModel).ToNot(BeNil())
				Expect(deleteLinkOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLinkOptionsModel.LinkID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLinkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProfileIdentityOptions successfully`, func() {
				// Construct an instance of the DeleteProfileIdentityOptions model
				profileID := "testString"
				identityType := "user"
				identifierID := "testString"
				deleteProfileIdentityOptionsModel := iamIdentityService.NewDeleteProfileIdentityOptions(profileID, identityType, identifierID)
				deleteProfileIdentityOptionsModel.SetProfileID("testString")
				deleteProfileIdentityOptionsModel.SetIdentityType("user")
				deleteProfileIdentityOptionsModel.SetIdentifierID("testString")
				deleteProfileIdentityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProfileIdentityOptionsModel).ToNot(BeNil())
				Expect(deleteProfileIdentityOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileIdentityOptionsModel.IdentityType).To(Equal(core.StringPtr("user")))
				Expect(deleteProfileIdentityOptionsModel.IdentifierID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileIdentityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProfileOptions successfully`, func() {
				// Construct an instance of the DeleteProfileOptions model
				profileID := "testString"
				deleteProfileOptionsModel := iamIdentityService.NewDeleteProfileOptions(profileID)
				deleteProfileOptionsModel.SetProfileID("testString")
				deleteProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProfileOptionsModel).ToNot(BeNil())
				Expect(deleteProfileOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProfileTemplateVersionOptions successfully`, func() {
				// Construct an instance of the DeleteProfileTemplateVersionOptions model
				templateID := "testString"
				version := "testString"
				deleteProfileTemplateVersionOptionsModel := iamIdentityService.NewDeleteProfileTemplateVersionOptions(templateID, version)
				deleteProfileTemplateVersionOptionsModel.SetTemplateID("testString")
				deleteProfileTemplateVersionOptionsModel.SetVersion("testString")
				deleteProfileTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProfileTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(deleteProfileTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileTemplateVersionOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteServiceIDOptions successfully`, func() {
				// Construct an instance of the DeleteServiceIDOptions model
				id := "testString"
				deleteServiceIDOptionsModel := iamIdentityService.NewDeleteServiceIDOptions(id)
				deleteServiceIDOptionsModel.SetID("testString")
				deleteServiceIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteServiceIDOptionsModel).ToNot(BeNil())
				Expect(deleteServiceIDOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceIDOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTrustedProfileAssignmentOptions successfully`, func() {
				// Construct an instance of the DeleteTrustedProfileAssignmentOptions model
				assignmentID := "testString"
				deleteTrustedProfileAssignmentOptionsModel := iamIdentityService.NewDeleteTrustedProfileAssignmentOptions(assignmentID)
				deleteTrustedProfileAssignmentOptionsModel.SetAssignmentID("testString")
				deleteTrustedProfileAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTrustedProfileAssignmentOptionsModel).ToNot(BeNil())
				Expect(deleteTrustedProfileAssignmentOptionsModel.AssignmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrustedProfileAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountSettingsAssignmentOptions successfully`, func() {
				// Construct an instance of the GetAccountSettingsAssignmentOptions model
				assignmentID := "testString"
				getAccountSettingsAssignmentOptionsModel := iamIdentityService.NewGetAccountSettingsAssignmentOptions(assignmentID)
				getAccountSettingsAssignmentOptionsModel.SetAssignmentID("testString")
				getAccountSettingsAssignmentOptionsModel.SetIncludeHistory(false)
				getAccountSettingsAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountSettingsAssignmentOptionsModel).ToNot(BeNil())
				Expect(getAccountSettingsAssignmentOptionsModel.AssignmentID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsAssignmentOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(getAccountSettingsAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountSettingsOptions successfully`, func() {
				// Construct an instance of the GetAccountSettingsOptions model
				accountID := "testString"
				getAccountSettingsOptionsModel := iamIdentityService.NewGetAccountSettingsOptions(accountID)
				getAccountSettingsOptionsModel.SetAccountID("testString")
				getAccountSettingsOptionsModel.SetIncludeHistory(false)
				getAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(getAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(getAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountSettingsTemplateVersionOptions successfully`, func() {
				// Construct an instance of the GetAccountSettingsTemplateVersionOptions model
				templateID := "testString"
				version := "testString"
				getAccountSettingsTemplateVersionOptionsModel := iamIdentityService.NewGetAccountSettingsTemplateVersionOptions(templateID, version)
				getAccountSettingsTemplateVersionOptionsModel.SetTemplateID("testString")
				getAccountSettingsTemplateVersionOptionsModel.SetVersion("testString")
				getAccountSettingsTemplateVersionOptionsModel.SetIncludeHistory(false)
				getAccountSettingsTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountSettingsTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(getAccountSettingsTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsTemplateVersionOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsTemplateVersionOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(getAccountSettingsTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAPIKeyOptions successfully`, func() {
				// Construct an instance of the GetAPIKeyOptions model
				id := "testString"
				getAPIKeyOptionsModel := iamIdentityService.NewGetAPIKeyOptions(id)
				getAPIKeyOptionsModel.SetID("testString")
				getAPIKeyOptionsModel.SetIncludeHistory(false)
				getAPIKeyOptionsModel.SetIncludeActivity(false)
				getAPIKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAPIKeyOptionsModel).ToNot(BeNil())
				Expect(getAPIKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getAPIKeyOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(getAPIKeyOptionsModel.IncludeActivity).To(Equal(core.BoolPtr(false)))
				Expect(getAPIKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAPIKeysDetailsOptions successfully`, func() {
				// Construct an instance of the GetAPIKeysDetailsOptions model
				getAPIKeysDetailsOptionsModel := iamIdentityService.NewGetAPIKeysDetailsOptions()
				getAPIKeysDetailsOptionsModel.SetIamAPIKey("testString")
				getAPIKeysDetailsOptionsModel.SetIncludeHistory(false)
				getAPIKeysDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAPIKeysDetailsOptionsModel).ToNot(BeNil())
				Expect(getAPIKeysDetailsOptionsModel.IamAPIKey).To(Equal(core.StringPtr("testString")))
				Expect(getAPIKeysDetailsOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(getAPIKeysDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetClaimRuleOptions successfully`, func() {
				// Construct an instance of the GetClaimRuleOptions model
				profileID := "testString"
				ruleID := "testString"
				getClaimRuleOptionsModel := iamIdentityService.NewGetClaimRuleOptions(profileID, ruleID)
				getClaimRuleOptionsModel.SetProfileID("testString")
				getClaimRuleOptionsModel.SetRuleID("testString")
				getClaimRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getClaimRuleOptionsModel).ToNot(BeNil())
				Expect(getClaimRuleOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(getClaimRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getClaimRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLatestAccountSettingsTemplateVersionOptions successfully`, func() {
				// Construct an instance of the GetLatestAccountSettingsTemplateVersionOptions model
				templateID := "testString"
				getLatestAccountSettingsTemplateVersionOptionsModel := iamIdentityService.NewGetLatestAccountSettingsTemplateVersionOptions(templateID)
				getLatestAccountSettingsTemplateVersionOptionsModel.SetTemplateID("testString")
				getLatestAccountSettingsTemplateVersionOptionsModel.SetIncludeHistory(false)
				getLatestAccountSettingsTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLatestAccountSettingsTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(getLatestAccountSettingsTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(getLatestAccountSettingsTemplateVersionOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(getLatestAccountSettingsTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLatestProfileTemplateVersionOptions successfully`, func() {
				// Construct an instance of the GetLatestProfileTemplateVersionOptions model
				templateID := "testString"
				getLatestProfileTemplateVersionOptionsModel := iamIdentityService.NewGetLatestProfileTemplateVersionOptions(templateID)
				getLatestProfileTemplateVersionOptionsModel.SetTemplateID("testString")
				getLatestProfileTemplateVersionOptionsModel.SetIncludeHistory(false)
				getLatestProfileTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLatestProfileTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(getLatestProfileTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(getLatestProfileTemplateVersionOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(getLatestProfileTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLinkOptions successfully`, func() {
				// Construct an instance of the GetLinkOptions model
				profileID := "testString"
				linkID := "testString"
				getLinkOptionsModel := iamIdentityService.NewGetLinkOptions(profileID, linkID)
				getLinkOptionsModel.SetProfileID("testString")
				getLinkOptionsModel.SetLinkID("testString")
				getLinkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLinkOptionsModel).ToNot(BeNil())
				Expect(getLinkOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(getLinkOptionsModel.LinkID).To(Equal(core.StringPtr("testString")))
				Expect(getLinkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMfaReportOptions successfully`, func() {
				// Construct an instance of the GetMfaReportOptions model
				accountID := "testString"
				reference := "testString"
				getMfaReportOptionsModel := iamIdentityService.NewGetMfaReportOptions(accountID, reference)
				getMfaReportOptionsModel.SetAccountID("testString")
				getMfaReportOptionsModel.SetReference("testString")
				getMfaReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMfaReportOptionsModel).ToNot(BeNil())
				Expect(getMfaReportOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getMfaReportOptionsModel.Reference).To(Equal(core.StringPtr("testString")))
				Expect(getMfaReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMfaStatusOptions successfully`, func() {
				// Construct an instance of the GetMfaStatusOptions model
				accountID := "testString"
				iamID := "testString"
				getMfaStatusOptionsModel := iamIdentityService.NewGetMfaStatusOptions(accountID, iamID)
				getMfaStatusOptionsModel.SetAccountID("testString")
				getMfaStatusOptionsModel.SetIamID("testString")
				getMfaStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMfaStatusOptionsModel).ToNot(BeNil())
				Expect(getMfaStatusOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getMfaStatusOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(getMfaStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileIdentitiesOptions successfully`, func() {
				// Construct an instance of the GetProfileIdentitiesOptions model
				profileID := "testString"
				getProfileIdentitiesOptionsModel := iamIdentityService.NewGetProfileIdentitiesOptions(profileID)
				getProfileIdentitiesOptionsModel.SetProfileID("testString")
				getProfileIdentitiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileIdentitiesOptionsModel).ToNot(BeNil())
				Expect(getProfileIdentitiesOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileIdentitiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileIdentityOptions successfully`, func() {
				// Construct an instance of the GetProfileIdentityOptions model
				profileID := "testString"
				identityType := "user"
				identifierID := "testString"
				getProfileIdentityOptionsModel := iamIdentityService.NewGetProfileIdentityOptions(profileID, identityType, identifierID)
				getProfileIdentityOptionsModel.SetProfileID("testString")
				getProfileIdentityOptionsModel.SetIdentityType("user")
				getProfileIdentityOptionsModel.SetIdentifierID("testString")
				getProfileIdentityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileIdentityOptionsModel).ToNot(BeNil())
				Expect(getProfileIdentityOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileIdentityOptionsModel.IdentityType).To(Equal(core.StringPtr("user")))
				Expect(getProfileIdentityOptionsModel.IdentifierID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileIdentityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileOptions successfully`, func() {
				// Construct an instance of the GetProfileOptions model
				profileID := "testString"
				getProfileOptionsModel := iamIdentityService.NewGetProfileOptions(profileID)
				getProfileOptionsModel.SetProfileID("testString")
				getProfileOptionsModel.SetIncludeActivity(false)
				getProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileOptionsModel).ToNot(BeNil())
				Expect(getProfileOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.IncludeActivity).To(Equal(core.BoolPtr(false)))
				Expect(getProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileTemplateVersionOptions successfully`, func() {
				// Construct an instance of the GetProfileTemplateVersionOptions model
				templateID := "testString"
				version := "testString"
				getProfileTemplateVersionOptionsModel := iamIdentityService.NewGetProfileTemplateVersionOptions(templateID, version)
				getProfileTemplateVersionOptionsModel.SetTemplateID("testString")
				getProfileTemplateVersionOptionsModel.SetVersion("testString")
				getProfileTemplateVersionOptionsModel.SetIncludeHistory(false)
				getProfileTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(getProfileTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileTemplateVersionOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(getProfileTemplateVersionOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(getProfileTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportOptions successfully`, func() {
				// Construct an instance of the GetReportOptions model
				accountID := "testString"
				reference := "testString"
				getReportOptionsModel := iamIdentityService.NewGetReportOptions(accountID, reference)
				getReportOptionsModel.SetAccountID("testString")
				getReportOptionsModel.SetReference("testString")
				getReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportOptionsModel).ToNot(BeNil())
				Expect(getReportOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getReportOptionsModel.Reference).To(Equal(core.StringPtr("testString")))
				Expect(getReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServiceIDOptions successfully`, func() {
				// Construct an instance of the GetServiceIDOptions model
				id := "testString"
				getServiceIDOptionsModel := iamIdentityService.NewGetServiceIDOptions(id)
				getServiceIDOptionsModel.SetID("testString")
				getServiceIDOptionsModel.SetIncludeHistory(false)
				getServiceIDOptionsModel.SetIncludeActivity(false)
				getServiceIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServiceIDOptionsModel).ToNot(BeNil())
				Expect(getServiceIDOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getServiceIDOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(getServiceIDOptionsModel.IncludeActivity).To(Equal(core.BoolPtr(false)))
				Expect(getServiceIDOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTrustedProfileAssignmentOptions successfully`, func() {
				// Construct an instance of the GetTrustedProfileAssignmentOptions model
				assignmentID := "testString"
				getTrustedProfileAssignmentOptionsModel := iamIdentityService.NewGetTrustedProfileAssignmentOptions(assignmentID)
				getTrustedProfileAssignmentOptionsModel.SetAssignmentID("testString")
				getTrustedProfileAssignmentOptionsModel.SetIncludeHistory(false)
				getTrustedProfileAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTrustedProfileAssignmentOptionsModel).ToNot(BeNil())
				Expect(getTrustedProfileAssignmentOptionsModel.AssignmentID).To(Equal(core.StringPtr("testString")))
				Expect(getTrustedProfileAssignmentOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(getTrustedProfileAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccountSettingsAssignmentsOptions successfully`, func() {
				// Construct an instance of the ListAccountSettingsAssignmentsOptions model
				listAccountSettingsAssignmentsOptionsModel := iamIdentityService.NewListAccountSettingsAssignmentsOptions()
				listAccountSettingsAssignmentsOptionsModel.SetAccountID("testString")
				listAccountSettingsAssignmentsOptionsModel.SetTemplateID("testString")
				listAccountSettingsAssignmentsOptionsModel.SetTemplateVersion("testString")
				listAccountSettingsAssignmentsOptionsModel.SetTarget("testString")
				listAccountSettingsAssignmentsOptionsModel.SetTargetType("Account")
				listAccountSettingsAssignmentsOptionsModel.SetLimit(int64(20))
				listAccountSettingsAssignmentsOptionsModel.SetPagetoken("testString")
				listAccountSettingsAssignmentsOptionsModel.SetSort("created_at")
				listAccountSettingsAssignmentsOptionsModel.SetOrder("asc")
				listAccountSettingsAssignmentsOptionsModel.SetIncludeHistory(false)
				listAccountSettingsAssignmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccountSettingsAssignmentsOptionsModel).ToNot(BeNil())
				Expect(listAccountSettingsAssignmentsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountSettingsAssignmentsOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountSettingsAssignmentsOptionsModel.TemplateVersion).To(Equal(core.StringPtr("testString")))
				Expect(listAccountSettingsAssignmentsOptionsModel.Target).To(Equal(core.StringPtr("testString")))
				Expect(listAccountSettingsAssignmentsOptionsModel.TargetType).To(Equal(core.StringPtr("Account")))
				Expect(listAccountSettingsAssignmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(20))))
				Expect(listAccountSettingsAssignmentsOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listAccountSettingsAssignmentsOptionsModel.Sort).To(Equal(core.StringPtr("created_at")))
				Expect(listAccountSettingsAssignmentsOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listAccountSettingsAssignmentsOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(listAccountSettingsAssignmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccountSettingsTemplatesOptions successfully`, func() {
				// Construct an instance of the ListAccountSettingsTemplatesOptions model
				listAccountSettingsTemplatesOptionsModel := iamIdentityService.NewListAccountSettingsTemplatesOptions()
				listAccountSettingsTemplatesOptionsModel.SetAccountID("testString")
				listAccountSettingsTemplatesOptionsModel.SetLimit("20")
				listAccountSettingsTemplatesOptionsModel.SetPagetoken("testString")
				listAccountSettingsTemplatesOptionsModel.SetSort("created_at")
				listAccountSettingsTemplatesOptionsModel.SetOrder("asc")
				listAccountSettingsTemplatesOptionsModel.SetIncludeHistory("false")
				listAccountSettingsTemplatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccountSettingsTemplatesOptionsModel).ToNot(BeNil())
				Expect(listAccountSettingsTemplatesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountSettingsTemplatesOptionsModel.Limit).To(Equal(core.StringPtr("20")))
				Expect(listAccountSettingsTemplatesOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listAccountSettingsTemplatesOptionsModel.Sort).To(Equal(core.StringPtr("created_at")))
				Expect(listAccountSettingsTemplatesOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listAccountSettingsTemplatesOptionsModel.IncludeHistory).To(Equal(core.StringPtr("false")))
				Expect(listAccountSettingsTemplatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAPIKeysOptions successfully`, func() {
				// Construct an instance of the ListAPIKeysOptions model
				listAPIKeysOptionsModel := iamIdentityService.NewListAPIKeysOptions()
				listAPIKeysOptionsModel.SetAccountID("testString")
				listAPIKeysOptionsModel.SetIamID("testString")
				listAPIKeysOptionsModel.SetPagesize(int64(38))
				listAPIKeysOptionsModel.SetPagetoken("testString")
				listAPIKeysOptionsModel.SetScope("entity")
				listAPIKeysOptionsModel.SetType("user")
				listAPIKeysOptionsModel.SetSort("testString")
				listAPIKeysOptionsModel.SetOrder("asc")
				listAPIKeysOptionsModel.SetIncludeHistory(false)
				listAPIKeysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAPIKeysOptionsModel).ToNot(BeNil())
				Expect(listAPIKeysOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listAPIKeysOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(listAPIKeysOptionsModel.Pagesize).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAPIKeysOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listAPIKeysOptionsModel.Scope).To(Equal(core.StringPtr("entity")))
				Expect(listAPIKeysOptionsModel.Type).To(Equal(core.StringPtr("user")))
				Expect(listAPIKeysOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listAPIKeysOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listAPIKeysOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(listAPIKeysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListClaimRulesOptions successfully`, func() {
				// Construct an instance of the ListClaimRulesOptions model
				profileID := "testString"
				listClaimRulesOptionsModel := iamIdentityService.NewListClaimRulesOptions(profileID)
				listClaimRulesOptionsModel.SetProfileID("testString")
				listClaimRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listClaimRulesOptionsModel).ToNot(BeNil())
				Expect(listClaimRulesOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(listClaimRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLinksOptions successfully`, func() {
				// Construct an instance of the ListLinksOptions model
				profileID := "testString"
				listLinksOptionsModel := iamIdentityService.NewListLinksOptions(profileID)
				listLinksOptionsModel.SetProfileID("testString")
				listLinksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLinksOptionsModel).ToNot(BeNil())
				Expect(listLinksOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(listLinksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProfileTemplatesOptions successfully`, func() {
				// Construct an instance of the ListProfileTemplatesOptions model
				listProfileTemplatesOptionsModel := iamIdentityService.NewListProfileTemplatesOptions()
				listProfileTemplatesOptionsModel.SetAccountID("testString")
				listProfileTemplatesOptionsModel.SetLimit("20")
				listProfileTemplatesOptionsModel.SetPagetoken("testString")
				listProfileTemplatesOptionsModel.SetSort("created_at")
				listProfileTemplatesOptionsModel.SetOrder("asc")
				listProfileTemplatesOptionsModel.SetIncludeHistory("false")
				listProfileTemplatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProfileTemplatesOptionsModel).ToNot(BeNil())
				Expect(listProfileTemplatesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listProfileTemplatesOptionsModel.Limit).To(Equal(core.StringPtr("20")))
				Expect(listProfileTemplatesOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listProfileTemplatesOptionsModel.Sort).To(Equal(core.StringPtr("created_at")))
				Expect(listProfileTemplatesOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listProfileTemplatesOptionsModel.IncludeHistory).To(Equal(core.StringPtr("false")))
				Expect(listProfileTemplatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProfilesOptions successfully`, func() {
				// Construct an instance of the ListProfilesOptions model
				accountID := "testString"
				listProfilesOptionsModel := iamIdentityService.NewListProfilesOptions(accountID)
				listProfilesOptionsModel.SetAccountID("testString")
				listProfilesOptionsModel.SetName("testString")
				listProfilesOptionsModel.SetPagesize(int64(38))
				listProfilesOptionsModel.SetSort("testString")
				listProfilesOptionsModel.SetOrder("asc")
				listProfilesOptionsModel.SetIncludeHistory(false)
				listProfilesOptionsModel.SetPagetoken("testString")
				listProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProfilesOptionsModel).ToNot(BeNil())
				Expect(listProfilesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Pagesize).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listProfilesOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listProfilesOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(listProfilesOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListServiceIdsOptions successfully`, func() {
				// Construct an instance of the ListServiceIdsOptions model
				listServiceIdsOptionsModel := iamIdentityService.NewListServiceIdsOptions()
				listServiceIdsOptionsModel.SetAccountID("testString")
				listServiceIdsOptionsModel.SetName("testString")
				listServiceIdsOptionsModel.SetPagesize(int64(38))
				listServiceIdsOptionsModel.SetPagetoken("testString")
				listServiceIdsOptionsModel.SetSort("testString")
				listServiceIdsOptionsModel.SetOrder("asc")
				listServiceIdsOptionsModel.SetIncludeHistory(false)
				listServiceIdsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listServiceIdsOptionsModel).ToNot(BeNil())
				Expect(listServiceIdsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listServiceIdsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listServiceIdsOptionsModel.Pagesize).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listServiceIdsOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listServiceIdsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listServiceIdsOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listServiceIdsOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(listServiceIdsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTrustedProfileAssignmentsOptions successfully`, func() {
				// Construct an instance of the ListTrustedProfileAssignmentsOptions model
				listTrustedProfileAssignmentsOptionsModel := iamIdentityService.NewListTrustedProfileAssignmentsOptions()
				listTrustedProfileAssignmentsOptionsModel.SetAccountID("testString")
				listTrustedProfileAssignmentsOptionsModel.SetTemplateID("testString")
				listTrustedProfileAssignmentsOptionsModel.SetTemplateVersion("testString")
				listTrustedProfileAssignmentsOptionsModel.SetTarget("testString")
				listTrustedProfileAssignmentsOptionsModel.SetTargetType("Account")
				listTrustedProfileAssignmentsOptionsModel.SetLimit(int64(20))
				listTrustedProfileAssignmentsOptionsModel.SetPagetoken("testString")
				listTrustedProfileAssignmentsOptionsModel.SetSort("created_at")
				listTrustedProfileAssignmentsOptionsModel.SetOrder("asc")
				listTrustedProfileAssignmentsOptionsModel.SetIncludeHistory(false)
				listTrustedProfileAssignmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTrustedProfileAssignmentsOptionsModel).ToNot(BeNil())
				Expect(listTrustedProfileAssignmentsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listTrustedProfileAssignmentsOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(listTrustedProfileAssignmentsOptionsModel.TemplateVersion).To(Equal(core.StringPtr("testString")))
				Expect(listTrustedProfileAssignmentsOptionsModel.Target).To(Equal(core.StringPtr("testString")))
				Expect(listTrustedProfileAssignmentsOptionsModel.TargetType).To(Equal(core.StringPtr("Account")))
				Expect(listTrustedProfileAssignmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(20))))
				Expect(listTrustedProfileAssignmentsOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listTrustedProfileAssignmentsOptionsModel.Sort).To(Equal(core.StringPtr("created_at")))
				Expect(listTrustedProfileAssignmentsOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listTrustedProfileAssignmentsOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(false)))
				Expect(listTrustedProfileAssignmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVersionsOfAccountSettingsTemplateOptions successfully`, func() {
				// Construct an instance of the ListVersionsOfAccountSettingsTemplateOptions model
				templateID := "testString"
				listVersionsOfAccountSettingsTemplateOptionsModel := iamIdentityService.NewListVersionsOfAccountSettingsTemplateOptions(templateID)
				listVersionsOfAccountSettingsTemplateOptionsModel.SetTemplateID("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.SetLimit("20")
				listVersionsOfAccountSettingsTemplateOptionsModel.SetPagetoken("testString")
				listVersionsOfAccountSettingsTemplateOptionsModel.SetSort("created_at")
				listVersionsOfAccountSettingsTemplateOptionsModel.SetOrder("asc")
				listVersionsOfAccountSettingsTemplateOptionsModel.SetIncludeHistory("false")
				listVersionsOfAccountSettingsTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVersionsOfAccountSettingsTemplateOptionsModel).ToNot(BeNil())
				Expect(listVersionsOfAccountSettingsTemplateOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(listVersionsOfAccountSettingsTemplateOptionsModel.Limit).To(Equal(core.StringPtr("20")))
				Expect(listVersionsOfAccountSettingsTemplateOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listVersionsOfAccountSettingsTemplateOptionsModel.Sort).To(Equal(core.StringPtr("created_at")))
				Expect(listVersionsOfAccountSettingsTemplateOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listVersionsOfAccountSettingsTemplateOptionsModel.IncludeHistory).To(Equal(core.StringPtr("false")))
				Expect(listVersionsOfAccountSettingsTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVersionsOfProfileTemplateOptions successfully`, func() {
				// Construct an instance of the ListVersionsOfProfileTemplateOptions model
				templateID := "testString"
				listVersionsOfProfileTemplateOptionsModel := iamIdentityService.NewListVersionsOfProfileTemplateOptions(templateID)
				listVersionsOfProfileTemplateOptionsModel.SetTemplateID("testString")
				listVersionsOfProfileTemplateOptionsModel.SetLimit("20")
				listVersionsOfProfileTemplateOptionsModel.SetPagetoken("testString")
				listVersionsOfProfileTemplateOptionsModel.SetSort("created_at")
				listVersionsOfProfileTemplateOptionsModel.SetOrder("asc")
				listVersionsOfProfileTemplateOptionsModel.SetIncludeHistory("false")
				listVersionsOfProfileTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVersionsOfProfileTemplateOptionsModel).ToNot(BeNil())
				Expect(listVersionsOfProfileTemplateOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(listVersionsOfProfileTemplateOptionsModel.Limit).To(Equal(core.StringPtr("20")))
				Expect(listVersionsOfProfileTemplateOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listVersionsOfProfileTemplateOptionsModel.Sort).To(Equal(core.StringPtr("created_at")))
				Expect(listVersionsOfProfileTemplateOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listVersionsOfProfileTemplateOptionsModel.IncludeHistory).To(Equal(core.StringPtr("false")))
				Expect(listVersionsOfProfileTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLockAPIKeyOptions successfully`, func() {
				// Construct an instance of the LockAPIKeyOptions model
				id := "testString"
				lockAPIKeyOptionsModel := iamIdentityService.NewLockAPIKeyOptions(id)
				lockAPIKeyOptionsModel.SetID("testString")
				lockAPIKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(lockAPIKeyOptionsModel).ToNot(BeNil())
				Expect(lockAPIKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(lockAPIKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLockServiceIDOptions successfully`, func() {
				// Construct an instance of the LockServiceIDOptions model
				id := "testString"
				lockServiceIDOptionsModel := iamIdentityService.NewLockServiceIDOptions(id)
				lockServiceIDOptionsModel.SetID("testString")
				lockServiceIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(lockServiceIDOptionsModel).ToNot(BeNil())
				Expect(lockServiceIDOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(lockServiceIDOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPolicyTemplateReference successfully`, func() {
				id := "testString"
				version := "testString"
				_model, err := iamIdentityService.NewPolicyTemplateReference(id, version)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProfileClaimRuleConditions successfully`, func() {
				claim := "testString"
				operator := "testString"
				value := "testString"
				_model, err := iamIdentityService.NewProfileClaimRuleConditions(claim, operator, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProfileIdentityRequest successfully`, func() {
				identifier := "testString"
				typeVar := "user"
				_model, err := iamIdentityService.NewProfileIdentityRequest(identifier, typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSetProfileIdentitiesOptions successfully`, func() {
				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				Expect(profileIdentityRequestModel).ToNot(BeNil())
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")
				Expect(profileIdentityRequestModel.Identifier).To(Equal(core.StringPtr("testString")))
				Expect(profileIdentityRequestModel.Type).To(Equal(core.StringPtr("user")))
				Expect(profileIdentityRequestModel.Accounts).To(Equal([]string{"testString"}))
				Expect(profileIdentityRequestModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SetProfileIdentitiesOptions model
				profileID := "testString"
				ifMatch := "testString"
				setProfileIdentitiesOptionsModel := iamIdentityService.NewSetProfileIdentitiesOptions(profileID, ifMatch)
				setProfileIdentitiesOptionsModel.SetProfileID("testString")
				setProfileIdentitiesOptionsModel.SetIfMatch("testString")
				setProfileIdentitiesOptionsModel.SetIdentities([]iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel})
				setProfileIdentitiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setProfileIdentitiesOptionsModel).ToNot(BeNil())
				Expect(setProfileIdentitiesOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(setProfileIdentitiesOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(setProfileIdentitiesOptionsModel.Identities).To(Equal([]iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}))
				Expect(setProfileIdentitiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetProfileIdentityOptions successfully`, func() {
				// Construct an instance of the SetProfileIdentityOptions model
				profileID := "testString"
				identityType := "user"
				setProfileIdentityOptionsIdentifier := "testString"
				setProfileIdentityOptionsType := "user"
				setProfileIdentityOptionsModel := iamIdentityService.NewSetProfileIdentityOptions(profileID, identityType, setProfileIdentityOptionsIdentifier, setProfileIdentityOptionsType)
				setProfileIdentityOptionsModel.SetProfileID("testString")
				setProfileIdentityOptionsModel.SetIdentityType("user")
				setProfileIdentityOptionsModel.SetIdentifier("testString")
				setProfileIdentityOptionsModel.SetType("user")
				setProfileIdentityOptionsModel.SetAccounts([]string{"testString"})
				setProfileIdentityOptionsModel.SetDescription("testString")
				setProfileIdentityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setProfileIdentityOptionsModel).ToNot(BeNil())
				Expect(setProfileIdentityOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(setProfileIdentityOptionsModel.IdentityType).To(Equal(core.StringPtr("user")))
				Expect(setProfileIdentityOptionsModel.Identifier).To(Equal(core.StringPtr("testString")))
				Expect(setProfileIdentityOptionsModel.Type).To(Equal(core.StringPtr("user")))
				Expect(setProfileIdentityOptionsModel.Accounts).To(Equal([]string{"testString"}))
				Expect(setProfileIdentityOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(setProfileIdentityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTemplateProfileComponentRequest successfully`, func() {
				name := "testString"
				_model, err := iamIdentityService.NewTemplateProfileComponentRequest(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTrustedProfileTemplateClaimRule successfully`, func() {
				typeVar := "Profile-SAML"
				conditions := []iamidentityv1.ProfileClaimRuleConditions{}
				_model, err := iamIdentityService.NewTrustedProfileTemplateClaimRule(typeVar, conditions)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUnlockAPIKeyOptions successfully`, func() {
				// Construct an instance of the UnlockAPIKeyOptions model
				id := "testString"
				unlockAPIKeyOptionsModel := iamIdentityService.NewUnlockAPIKeyOptions(id)
				unlockAPIKeyOptionsModel.SetID("testString")
				unlockAPIKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unlockAPIKeyOptionsModel).ToNot(BeNil())
				Expect(unlockAPIKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(unlockAPIKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUnlockServiceIDOptions successfully`, func() {
				// Construct an instance of the UnlockServiceIDOptions model
				id := "testString"
				unlockServiceIDOptionsModel := iamIdentityService.NewUnlockServiceIDOptions(id)
				unlockServiceIDOptionsModel.SetID("testString")
				unlockServiceIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unlockServiceIDOptionsModel).ToNot(BeNil())
				Expect(unlockServiceIDOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(unlockServiceIDOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountSettingsAssignmentOptions successfully`, func() {
				// Construct an instance of the UpdateAccountSettingsAssignmentOptions model
				assignmentID := "testString"
				ifMatch := "testString"
				updateAccountSettingsAssignmentOptionsTemplateVersion := int64(1)
				updateAccountSettingsAssignmentOptionsModel := iamIdentityService.NewUpdateAccountSettingsAssignmentOptions(assignmentID, ifMatch, updateAccountSettingsAssignmentOptionsTemplateVersion)
				updateAccountSettingsAssignmentOptionsModel.SetAssignmentID("testString")
				updateAccountSettingsAssignmentOptionsModel.SetIfMatch("testString")
				updateAccountSettingsAssignmentOptionsModel.SetTemplateVersion(int64(1))
				updateAccountSettingsAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountSettingsAssignmentOptionsModel).ToNot(BeNil())
				Expect(updateAccountSettingsAssignmentOptionsModel.AssignmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsAssignmentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsAssignmentOptionsModel.TemplateVersion).To(Equal(core.Int64Ptr(int64(1))))
				Expect(updateAccountSettingsAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountSettingsOptions successfully`, func() {
				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				Expect(accountSettingsUserMfaModel).ToNot(BeNil())
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")
				Expect(accountSettingsUserMfaModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(accountSettingsUserMfaModel.Mfa).To(Equal(core.StringPtr("NONE")))

				// Construct an instance of the UpdateAccountSettingsOptions model
				ifMatch := "testString"
				accountID := "testString"
				updateAccountSettingsOptionsModel := iamIdentityService.NewUpdateAccountSettingsOptions(ifMatch, accountID)
				updateAccountSettingsOptionsModel.SetIfMatch("testString")
				updateAccountSettingsOptionsModel.SetAccountID("testString")
				updateAccountSettingsOptionsModel.SetRestrictCreateServiceID("RESTRICTED")
				updateAccountSettingsOptionsModel.SetRestrictCreatePlatformApikey("RESTRICTED")
				updateAccountSettingsOptionsModel.SetAllowedIPAddresses("testString")
				updateAccountSettingsOptionsModel.SetMfa("NONE")
				updateAccountSettingsOptionsModel.SetUserMfa([]iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel})
				updateAccountSettingsOptionsModel.SetSessionExpirationInSeconds("86400")
				updateAccountSettingsOptionsModel.SetSessionInvalidationInSeconds("7200")
				updateAccountSettingsOptionsModel.SetMaxSessionsPerIdentity("testString")
				updateAccountSettingsOptionsModel.SetSystemAccessTokenExpirationInSeconds("3600")
				updateAccountSettingsOptionsModel.SetSystemRefreshTokenExpirationInSeconds("259200")
				updateAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(updateAccountSettingsOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.RestrictCreateServiceID).To(Equal(core.StringPtr("RESTRICTED")))
				Expect(updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey).To(Equal(core.StringPtr("RESTRICTED")))
				Expect(updateAccountSettingsOptionsModel.AllowedIPAddresses).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.Mfa).To(Equal(core.StringPtr("NONE")))
				Expect(updateAccountSettingsOptionsModel.UserMfa).To(Equal([]iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}))
				Expect(updateAccountSettingsOptionsModel.SessionExpirationInSeconds).To(Equal(core.StringPtr("86400")))
				Expect(updateAccountSettingsOptionsModel.SessionInvalidationInSeconds).To(Equal(core.StringPtr("7200")))
				Expect(updateAccountSettingsOptionsModel.MaxSessionsPerIdentity).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.SystemAccessTokenExpirationInSeconds).To(Equal(core.StringPtr("3600")))
				Expect(updateAccountSettingsOptionsModel.SystemRefreshTokenExpirationInSeconds).To(Equal(core.StringPtr("259200")))
				Expect(updateAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountSettingsTemplateVersionOptions successfully`, func() {
				// Construct an instance of the AccountSettingsUserMfa model
				accountSettingsUserMfaModel := new(iamidentityv1.AccountSettingsUserMfa)
				Expect(accountSettingsUserMfaModel).ToNot(BeNil())
				accountSettingsUserMfaModel.IamID = core.StringPtr("testString")
				accountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")
				Expect(accountSettingsUserMfaModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(accountSettingsUserMfaModel.Mfa).To(Equal(core.StringPtr("NONE")))

				// Construct an instance of the AccountSettingsComponent model
				accountSettingsComponentModel := new(iamidentityv1.AccountSettingsComponent)
				Expect(accountSettingsComponentModel).ToNot(BeNil())
				accountSettingsComponentModel.RestrictCreateServiceID = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
				accountSettingsComponentModel.AllowedIPAddresses = core.StringPtr("testString")
				accountSettingsComponentModel.Mfa = core.StringPtr("NONE")
				accountSettingsComponentModel.UserMfa = []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}
				accountSettingsComponentModel.SessionExpirationInSeconds = core.StringPtr("86400")
				accountSettingsComponentModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				accountSettingsComponentModel.MaxSessionsPerIdentity = core.StringPtr("testString")
				accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
				accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")
				Expect(accountSettingsComponentModel.RestrictCreateServiceID).To(Equal(core.StringPtr("NOT_SET")))
				Expect(accountSettingsComponentModel.RestrictCreatePlatformApikey).To(Equal(core.StringPtr("NOT_SET")))
				Expect(accountSettingsComponentModel.AllowedIPAddresses).To(Equal(core.StringPtr("testString")))
				Expect(accountSettingsComponentModel.Mfa).To(Equal(core.StringPtr("NONE")))
				Expect(accountSettingsComponentModel.UserMfa).To(Equal([]iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMfaModel}))
				Expect(accountSettingsComponentModel.SessionExpirationInSeconds).To(Equal(core.StringPtr("86400")))
				Expect(accountSettingsComponentModel.SessionInvalidationInSeconds).To(Equal(core.StringPtr("7200")))
				Expect(accountSettingsComponentModel.MaxSessionsPerIdentity).To(Equal(core.StringPtr("testString")))
				Expect(accountSettingsComponentModel.SystemAccessTokenExpirationInSeconds).To(Equal(core.StringPtr("3600")))
				Expect(accountSettingsComponentModel.SystemRefreshTokenExpirationInSeconds).To(Equal(core.StringPtr("259200")))

				// Construct an instance of the UpdateAccountSettingsTemplateVersionOptions model
				ifMatch := "testString"
				templateID := "testString"
				version := "testString"
				updateAccountSettingsTemplateVersionOptionsModel := iamIdentityService.NewUpdateAccountSettingsTemplateVersionOptions(ifMatch, templateID, version)
				updateAccountSettingsTemplateVersionOptionsModel.SetIfMatch("testString")
				updateAccountSettingsTemplateVersionOptionsModel.SetTemplateID("testString")
				updateAccountSettingsTemplateVersionOptionsModel.SetVersion("testString")
				updateAccountSettingsTemplateVersionOptionsModel.SetAccountID("testString")
				updateAccountSettingsTemplateVersionOptionsModel.SetName("testString")
				updateAccountSettingsTemplateVersionOptionsModel.SetDescription("testString")
				updateAccountSettingsTemplateVersionOptionsModel.SetAccountSettings(accountSettingsComponentModel)
				updateAccountSettingsTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountSettingsTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(updateAccountSettingsTemplateVersionOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsTemplateVersionOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsTemplateVersionOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsTemplateVersionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsTemplateVersionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsTemplateVersionOptionsModel.AccountSettings).To(Equal(accountSettingsComponentModel))
				Expect(updateAccountSettingsTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAPIKeyOptions successfully`, func() {
				// Construct an instance of the UpdateAPIKeyOptions model
				id := "testString"
				ifMatch := "testString"
				updateAPIKeyOptionsModel := iamIdentityService.NewUpdateAPIKeyOptions(id, ifMatch)
				updateAPIKeyOptionsModel.SetID("testString")
				updateAPIKeyOptionsModel.SetIfMatch("testString")
				updateAPIKeyOptionsModel.SetName("testString")
				updateAPIKeyOptionsModel.SetDescription("testString")
				updateAPIKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAPIKeyOptionsModel).ToNot(BeNil())
				Expect(updateAPIKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateAPIKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAPIKeyOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateAPIKeyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateAPIKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateClaimRuleOptions successfully`, func() {
				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				Expect(profileClaimRuleConditionsModel).ToNot(BeNil())
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")
				Expect(profileClaimRuleConditionsModel.Claim).To(Equal(core.StringPtr("testString")))
				Expect(profileClaimRuleConditionsModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(profileClaimRuleConditionsModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ResponseContext model
				responseContextModel := new(iamidentityv1.ResponseContext)
				Expect(responseContextModel).ToNot(BeNil())
				responseContextModel.TransactionID = core.StringPtr("testString")
				responseContextModel.Operation = core.StringPtr("testString")
				responseContextModel.UserAgent = core.StringPtr("testString")
				responseContextModel.URL = core.StringPtr("testString")
				responseContextModel.InstanceID = core.StringPtr("testString")
				responseContextModel.ThreadID = core.StringPtr("testString")
				responseContextModel.Host = core.StringPtr("testString")
				responseContextModel.StartTime = core.StringPtr("testString")
				responseContextModel.EndTime = core.StringPtr("testString")
				responseContextModel.ElapsedTime = core.StringPtr("testString")
				responseContextModel.ClusterName = core.StringPtr("testString")
				Expect(responseContextModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.Operation).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.UserAgent).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.ThreadID).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.Host).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.StartTime).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.EndTime).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.ElapsedTime).To(Equal(core.StringPtr("testString")))
				Expect(responseContextModel.ClusterName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateClaimRuleOptions model
				profileID := "testString"
				ruleID := "testString"
				ifMatch := "testString"
				updateClaimRuleOptionsType := "testString"
				updateClaimRuleOptionsConditions := []iamidentityv1.ProfileClaimRuleConditions{}
				updateClaimRuleOptionsModel := iamIdentityService.NewUpdateClaimRuleOptions(profileID, ruleID, ifMatch, updateClaimRuleOptionsType, updateClaimRuleOptionsConditions)
				updateClaimRuleOptionsModel.SetProfileID("testString")
				updateClaimRuleOptionsModel.SetRuleID("testString")
				updateClaimRuleOptionsModel.SetIfMatch("testString")
				updateClaimRuleOptionsModel.SetType("testString")
				updateClaimRuleOptionsModel.SetConditions([]iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel})
				updateClaimRuleOptionsModel.SetContext(responseContextModel)
				updateClaimRuleOptionsModel.SetName("testString")
				updateClaimRuleOptionsModel.SetRealmName("testString")
				updateClaimRuleOptionsModel.SetCrType("testString")
				updateClaimRuleOptionsModel.SetExpiration(int64(38))
				updateClaimRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateClaimRuleOptionsModel).ToNot(BeNil())
				Expect(updateClaimRuleOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(updateClaimRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(updateClaimRuleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateClaimRuleOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(updateClaimRuleOptionsModel.Conditions).To(Equal([]iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}))
				Expect(updateClaimRuleOptionsModel.Context).To(Equal(responseContextModel))
				Expect(updateClaimRuleOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateClaimRuleOptionsModel.RealmName).To(Equal(core.StringPtr("testString")))
				Expect(updateClaimRuleOptionsModel.CrType).To(Equal(core.StringPtr("testString")))
				Expect(updateClaimRuleOptionsModel.Expiration).To(Equal(core.Int64Ptr(int64(38))))
				Expect(updateClaimRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProfileOptions successfully`, func() {
				// Construct an instance of the UpdateProfileOptions model
				profileID := "testString"
				ifMatch := "testString"
				updateProfileOptionsModel := iamIdentityService.NewUpdateProfileOptions(profileID, ifMatch)
				updateProfileOptionsModel.SetProfileID("testString")
				updateProfileOptionsModel.SetIfMatch("testString")
				updateProfileOptionsModel.SetName("testString")
				updateProfileOptionsModel.SetDescription("testString")
				updateProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProfileOptionsModel).ToNot(BeNil())
				Expect(updateProfileOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProfileTemplateVersionOptions successfully`, func() {
				// Construct an instance of the ProfileClaimRuleConditions model
				profileClaimRuleConditionsModel := new(iamidentityv1.ProfileClaimRuleConditions)
				Expect(profileClaimRuleConditionsModel).ToNot(BeNil())
				profileClaimRuleConditionsModel.Claim = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Operator = core.StringPtr("testString")
				profileClaimRuleConditionsModel.Value = core.StringPtr("testString")
				Expect(profileClaimRuleConditionsModel.Claim).To(Equal(core.StringPtr("testString")))
				Expect(profileClaimRuleConditionsModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(profileClaimRuleConditionsModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TrustedProfileTemplateClaimRule model
				trustedProfileTemplateClaimRuleModel := new(iamidentityv1.TrustedProfileTemplateClaimRule)
				Expect(trustedProfileTemplateClaimRuleModel).ToNot(BeNil())
				trustedProfileTemplateClaimRuleModel.Name = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Type = core.StringPtr("Profile-SAML")
				trustedProfileTemplateClaimRuleModel.RealmName = core.StringPtr("testString")
				trustedProfileTemplateClaimRuleModel.Expiration = core.Int64Ptr(int64(38))
				trustedProfileTemplateClaimRuleModel.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}
				Expect(trustedProfileTemplateClaimRuleModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(trustedProfileTemplateClaimRuleModel.Type).To(Equal(core.StringPtr("Profile-SAML")))
				Expect(trustedProfileTemplateClaimRuleModel.RealmName).To(Equal(core.StringPtr("testString")))
				Expect(trustedProfileTemplateClaimRuleModel.Expiration).To(Equal(core.Int64Ptr(int64(38))))
				Expect(trustedProfileTemplateClaimRuleModel.Conditions).To(Equal([]iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditionsModel}))

				// Construct an instance of the ProfileIdentityRequest model
				profileIdentityRequestModel := new(iamidentityv1.ProfileIdentityRequest)
				Expect(profileIdentityRequestModel).ToNot(BeNil())
				profileIdentityRequestModel.Identifier = core.StringPtr("testString")
				profileIdentityRequestModel.Type = core.StringPtr("user")
				profileIdentityRequestModel.Accounts = []string{"testString"}
				profileIdentityRequestModel.Description = core.StringPtr("testString")
				Expect(profileIdentityRequestModel.Identifier).To(Equal(core.StringPtr("testString")))
				Expect(profileIdentityRequestModel.Type).To(Equal(core.StringPtr("user")))
				Expect(profileIdentityRequestModel.Accounts).To(Equal([]string{"testString"}))
				Expect(profileIdentityRequestModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TemplateProfileComponentRequest model
				templateProfileComponentRequestModel := new(iamidentityv1.TemplateProfileComponentRequest)
				Expect(templateProfileComponentRequestModel).ToNot(BeNil())
				templateProfileComponentRequestModel.Name = core.StringPtr("testString")
				templateProfileComponentRequestModel.Description = core.StringPtr("testString")
				templateProfileComponentRequestModel.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}
				templateProfileComponentRequestModel.Identities = []iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}
				Expect(templateProfileComponentRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(templateProfileComponentRequestModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(templateProfileComponentRequestModel.Rules).To(Equal([]iamidentityv1.TrustedProfileTemplateClaimRule{*trustedProfileTemplateClaimRuleModel}))
				Expect(templateProfileComponentRequestModel.Identities).To(Equal([]iamidentityv1.ProfileIdentityRequest{*profileIdentityRequestModel}))

				// Construct an instance of the PolicyTemplateReference model
				policyTemplateReferenceModel := new(iamidentityv1.PolicyTemplateReference)
				Expect(policyTemplateReferenceModel).ToNot(BeNil())
				policyTemplateReferenceModel.ID = core.StringPtr("testString")
				policyTemplateReferenceModel.Version = core.StringPtr("testString")
				Expect(policyTemplateReferenceModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(policyTemplateReferenceModel.Version).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateProfileTemplateVersionOptions model
				ifMatch := "testString"
				templateID := "testString"
				version := "testString"
				updateProfileTemplateVersionOptionsModel := iamIdentityService.NewUpdateProfileTemplateVersionOptions(ifMatch, templateID, version)
				updateProfileTemplateVersionOptionsModel.SetIfMatch("testString")
				updateProfileTemplateVersionOptionsModel.SetTemplateID("testString")
				updateProfileTemplateVersionOptionsModel.SetVersion("testString")
				updateProfileTemplateVersionOptionsModel.SetAccountID("testString")
				updateProfileTemplateVersionOptionsModel.SetName("testString")
				updateProfileTemplateVersionOptionsModel.SetDescription("testString")
				updateProfileTemplateVersionOptionsModel.SetProfile(templateProfileComponentRequestModel)
				updateProfileTemplateVersionOptionsModel.SetPolicyTemplateReferences([]iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel})
				updateProfileTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProfileTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(updateProfileTemplateVersionOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileTemplateVersionOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileTemplateVersionOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileTemplateVersionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileTemplateVersionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileTemplateVersionOptionsModel.Profile).To(Equal(templateProfileComponentRequestModel))
				Expect(updateProfileTemplateVersionOptionsModel.PolicyTemplateReferences).To(Equal([]iamidentityv1.PolicyTemplateReference{*policyTemplateReferenceModel}))
				Expect(updateProfileTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateServiceIDOptions successfully`, func() {
				// Construct an instance of the UpdateServiceIDOptions model
				id := "testString"
				ifMatch := "testString"
				updateServiceIDOptionsModel := iamIdentityService.NewUpdateServiceIDOptions(id, ifMatch)
				updateServiceIDOptionsModel.SetID("testString")
				updateServiceIDOptionsModel.SetIfMatch("testString")
				updateServiceIDOptionsModel.SetName("testString")
				updateServiceIDOptionsModel.SetDescription("testString")
				updateServiceIDOptionsModel.SetUniqueInstanceCrns([]string{"testString"})
				updateServiceIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateServiceIDOptionsModel).ToNot(BeNil())
				Expect(updateServiceIDOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceIDOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceIDOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceIDOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceIDOptionsModel.UniqueInstanceCrns).To(Equal([]string{"testString"}))
				Expect(updateServiceIDOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTrustedProfileAssignmentOptions successfully`, func() {
				// Construct an instance of the UpdateTrustedProfileAssignmentOptions model
				assignmentID := "testString"
				ifMatch := "testString"
				updateTrustedProfileAssignmentOptionsTemplateVersion := int64(1)
				updateTrustedProfileAssignmentOptionsModel := iamIdentityService.NewUpdateTrustedProfileAssignmentOptions(assignmentID, ifMatch, updateTrustedProfileAssignmentOptionsTemplateVersion)
				updateTrustedProfileAssignmentOptionsModel.SetAssignmentID("testString")
				updateTrustedProfileAssignmentOptionsModel.SetIfMatch("testString")
				updateTrustedProfileAssignmentOptionsModel.SetTemplateVersion(int64(1))
				updateTrustedProfileAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTrustedProfileAssignmentOptionsModel).ToNot(BeNil())
				Expect(updateTrustedProfileAssignmentOptionsModel.AssignmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateTrustedProfileAssignmentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateTrustedProfileAssignmentOptionsModel.TemplateVersion).To(Equal(core.Int64Ptr(int64(1))))
				Expect(updateTrustedProfileAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
