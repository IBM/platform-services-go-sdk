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

package iamidentityv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "profiles": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}]}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "profiles": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}]}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "name": "Name", "description": "Description", "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "iam_id": "IamID", "account_id": "AccountID", "ims_account_id": 12, "ims_user_id": 9, "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "activity": {"last_authn": "LastAuthn", "authn_count": 10}}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity"}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity"}`)
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

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamidentityv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.RestrictCreateServiceID = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Mfa = core.StringPtr("NONE")
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("86400")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				updateAccountSettingsOptionsModel.MaxSessionsPerIdentity = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity"}`)
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

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamidentityv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.RestrictCreateServiceID = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Mfa = core.StringPtr("NONE")
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("86400")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				updateAccountSettingsOptionsModel.MaxSessionsPerIdentity = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "NOT_SET", "restrict_create_platform_apikey": "NOT_SET", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "86400", "session_invalidation_in_seconds": "7200", "max_sessions_per_identity": "MaxSessionsPerIdentity"}`)
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

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamidentityv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.RestrictCreateServiceID = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Mfa = core.StringPtr("NONE")
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("86400")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				updateAccountSettingsOptionsModel.MaxSessionsPerIdentity = core.StringPtr("testString")
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

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamidentityv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.RestrictCreateServiceID = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Mfa = core.StringPtr("NONE")
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("86400")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				updateAccountSettingsOptionsModel.MaxSessionsPerIdentity = core.StringPtr("testString")
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

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamidentityv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.RestrictCreateServiceID = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey = core.StringPtr("RESTRICTED")
				updateAccountSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Mfa = core.StringPtr("NONE")
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("86400")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("7200")
				updateAccountSettingsOptionsModel.MaxSessionsPerIdentity = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"created_by": "CreatedBy", "reference": "Reference", "report_duration": "ReportDuration", "report_start_time": "ReportStartTime", "report_end_time": "ReportEndTime", "users": [{"iam_id": "IamID", "username": "Username", "last_authn": "LastAuthn"}], "apikeys": [{"id": "ID", "name": "Name", "last_authn": "LastAuthn"}], "serviceids": [{"id": "ID", "name": "Name", "last_authn": "LastAuthn"}], "profiles": [{"id": "ID", "name": "Name", "last_authn": "LastAuthn"}]}`)
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
					fmt.Fprintf(res, "%s", `{"created_by": "CreatedBy", "reference": "Reference", "report_duration": "ReportDuration", "report_start_time": "ReportStartTime", "report_end_time": "ReportEndTime", "users": [{"iam_id": "IamID", "username": "Username", "last_authn": "LastAuthn"}], "apikeys": [{"id": "ID", "name": "Name", "last_authn": "LastAuthn"}], "serviceids": [{"id": "ID", "name": "Name", "last_authn": "LastAuthn"}], "profiles": [{"id": "ID", "name": "Name", "last_authn": "LastAuthn"}]}`)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			iamIdentityService, _ := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
				URL:           "http://iamidentityv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAPIKeyInsideCreateServiceIDRequest successfully`, func() {
				name := "testString"
				_model, err := iamIdentityService.NewAPIKeyInsideCreateServiceIDRequest(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
			It(`Invoke NewProfileClaimRuleConditions successfully`, func() {
				claim := "testString"
				operator := "testString"
				value := "testString"
				_model, err := iamIdentityService.NewProfileClaimRuleConditions(claim, operator, value)
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
			It(`Invoke NewUpdateAccountSettingsOptions successfully`, func() {
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
				updateAccountSettingsOptionsModel.SetSessionExpirationInSeconds("86400")
				updateAccountSettingsOptionsModel.SetSessionInvalidationInSeconds("7200")
				updateAccountSettingsOptionsModel.SetMaxSessionsPerIdentity("testString")
				updateAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(updateAccountSettingsOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.RestrictCreateServiceID).To(Equal(core.StringPtr("RESTRICTED")))
				Expect(updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey).To(Equal(core.StringPtr("RESTRICTED")))
				Expect(updateAccountSettingsOptionsModel.AllowedIPAddresses).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.Mfa).To(Equal(core.StringPtr("NONE")))
				Expect(updateAccountSettingsOptionsModel.SessionExpirationInSeconds).To(Equal(core.StringPtr("86400")))
				Expect(updateAccountSettingsOptionsModel.SessionInvalidationInSeconds).To(Equal(core.StringPtr("7200")))
				Expect(updateAccountSettingsOptionsModel.MaxSessionsPerIdentity).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
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
