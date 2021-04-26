/**
 * (C) Copyright IBM Corp. 2021.
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
				"IAM_IDENTITY_URL":       "https://iamidentityv1/api",
				"IAM_IDENTITY_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{})
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
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{})
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
				"IAM_IDENTITY_URL":       "https://iamidentityv1/api",
				"IAM_IDENTITY_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(iamIdentityService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_IDENTITY_AUTH_TYPE": "NOAuth",
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
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				listAPIKeysOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "apikeys": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}]}`)
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
				listAPIKeysOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "apikeys": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}]}`)
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
				listAPIKeysOptionsModel.IncludeHistory = core.BoolPtr(true)
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
				listAPIKeysOptionsModel.IncludeHistory = core.BoolPtr(true)
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
	})
	Describe(`CreateAPIKey(createAPIKeyOptions *CreateAPIKeyOptions) - Operation response error`, func() {
		createAPIKeyPath := "/v1/apikeys"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAPIKeyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				createAPIKeyOptionsModel.EntityLock = core.StringPtr("testString")
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
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
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
				createAPIKeyOptionsModel.EntityLock = core.StringPtr("testString")
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
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
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
				createAPIKeyOptionsModel.EntityLock = core.StringPtr("testString")
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
				createAPIKeyOptionsModel.EntityLock = core.StringPtr("testString")
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
	})
	Describe(`GetAPIKeysDetails(getAPIKeysDetailsOptions *GetAPIKeysDetailsOptions) - Operation response error`, func() {
		getAPIKeysDetailsPath := "/v1/apikeys/details"
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
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getAPIKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
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
				getAPIKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
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
				getAPIKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
				getAPIKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
	})
	Describe(`GetAPIKey(getAPIKeyOptions *GetAPIKeyOptions) - Operation response error`, func() {
		getAPIKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAPIKeyPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getAPIKeyOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
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
				getAPIKeyOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
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
				getAPIKeyOptionsModel.IncludeHistory = core.BoolPtr(true)
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
				getAPIKeyOptionsModel.IncludeHistory = core.BoolPtr(true)
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
	})
	Describe(`UpdateAPIKey(updateAPIKeyOptions *UpdateAPIKeyOptions) - Operation response error`, func() {
		updateAPIKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint`, func() {
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
					fmt.Fprintf(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
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
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				listServiceIdsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "serviceids": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}]}`)
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
				listServiceIdsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "serviceids": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}]}`)
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
				listServiceIdsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
				listServiceIdsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
	})
	Describe(`CreateServiceID(createServiceIDOptions *CreateServiceIDOptions) - Operation response error`, func() {
		createServiceIDPath := "/v1/serviceids/"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createServiceIDPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				createServiceIDOptionsModel.EntityLock = core.StringPtr("testString")
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
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}`)
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
				createServiceIDOptionsModel.EntityLock = core.StringPtr("testString")
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
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}`)
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
				createServiceIDOptionsModel.EntityLock = core.StringPtr("testString")
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
				createServiceIDOptionsModel.EntityLock = core.StringPtr("testString")
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
	})
	Describe(`GetServiceID(getServiceIDOptions *GetServiceIDOptions) - Operation response error`, func() {
		getServiceIDPath := "/v1/serviceids/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceIDPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getServiceIDOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}`)
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
				getServiceIDOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}`)
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
				getServiceIDOptionsModel.IncludeHistory = core.BoolPtr(true)
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
				getServiceIDOptionsModel.IncludeHistory = core.BoolPtr(true)
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
	})
	Describe(`UpdateServiceID(updateServiceIDOptions *UpdateServiceIDOptions) - Operation response error`, func() {
		updateServiceIDPath := "/v1/serviceids/testString"
		Context(`Using mock server endpoint`, func() {
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
					fmt.Fprintf(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}`)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "modified_at": "2019-01-01T12:00:00.000Z", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "CRN", "locked": true, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00.000Z", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}`)
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
				"IAM_IDENTITY_URL":       "https://iamidentityv1/api",
				"IAM_IDENTITY_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{})
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
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{})
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
				"IAM_IDENTITY_URL":       "https://iamidentityv1/api",
				"IAM_IDENTITY_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1UsingExternalConfig(&iamidentityv1.IamIdentityV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(iamIdentityService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_IDENTITY_AUTH_TYPE": "NOAuth",
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
	Describe(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) - Operation response error`, func() {
		getAccountSettingsPath := "/v1/accounts/testString/settings/identity"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for include_history query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getAccountSettingsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "RESTRICTED", "restrict_create_platform_apikey": "RESTRICTED", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "SessionExpirationInSeconds", "session_invalidation_in_seconds": "SessionInvalidationInSeconds", "max_sessions_per_identity": "MaxSessionsPerIdentity"}`)
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
				getAccountSettingsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "RESTRICTED", "restrict_create_platform_apikey": "RESTRICTED", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "SessionExpirationInSeconds", "session_invalidation_in_seconds": "SessionInvalidationInSeconds", "max_sessions_per_identity": "MaxSessionsPerIdentity"}`)
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
				getAccountSettingsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
				getAccountSettingsOptionsModel.IncludeHistory = core.BoolPtr(true)
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
	})
	Describe(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions) - Operation response error`, func() {
		updateAccountSettingsPath := "/v1/accounts/testString/settings/identity"
		Context(`Using mock server endpoint`, func() {
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
					fmt.Fprintf(res, `} this is not valid json {`)
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
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "RESTRICTED", "restrict_create_platform_apikey": "RESTRICTED", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "SessionExpirationInSeconds", "session_invalidation_in_seconds": "SessionInvalidationInSeconds", "max_sessions_per_identity": "MaxSessionsPerIdentity"}`)
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
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "account_id": "AccountID", "restrict_create_service_id": "RESTRICTED", "restrict_create_platform_apikey": "RESTRICTED", "allowed_ip_addresses": "AllowedIPAddresses", "entity_tag": "EntityTag", "mfa": "NONE", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIDAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "session_expiration_in_seconds": "SessionExpirationInSeconds", "session_invalidation_in_seconds": "SessionInvalidationInSeconds", "max_sessions_per_identity": "MaxSessionsPerIdentity"}`)
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
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("testString")
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
				updateAccountSettingsOptionsModel.SessionExpirationInSeconds = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.SessionInvalidationInSeconds = core.StringPtr("testString")
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
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			iamIdentityService, _ := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
				URL:           "http://iamidentityv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAPIKeyInsideCreateServiceIDRequest successfully`, func() {
				name := "testString"
				model, err := iamIdentityService.NewAPIKeyInsideCreateServiceIDRequest(name)
				Expect(model).ToNot(BeNil())
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
				createAPIKeyOptionsModel.SetEntityLock("testString")
				createAPIKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAPIKeyOptionsModel).ToNot(BeNil())
				Expect(createAPIKeyOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.Apikey).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.StoreValue).To(Equal(core.BoolPtr(true)))
				Expect(createAPIKeyOptionsModel.EntityLock).To(Equal(core.StringPtr("testString")))
				Expect(createAPIKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				createServiceIDOptionsModel.SetEntityLock("testString")
				createServiceIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createServiceIDOptionsModel).ToNot(BeNil())
				Expect(createServiceIDOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createServiceIDOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createServiceIDOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createServiceIDOptionsModel.UniqueInstanceCrns).To(Equal([]string{"testString"}))
				Expect(createServiceIDOptionsModel.Apikey).To(Equal(apiKeyInsideCreateServiceIDRequestModel))
				Expect(createServiceIDOptionsModel.EntityLock).To(Equal(core.StringPtr("testString")))
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
				getAccountSettingsOptionsModel.SetIncludeHistory(true)
				getAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(getAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(true)))
				Expect(getAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAPIKeyOptions successfully`, func() {
				// Construct an instance of the GetAPIKeyOptions model
				id := "testString"
				getAPIKeyOptionsModel := iamIdentityService.NewGetAPIKeyOptions(id)
				getAPIKeyOptionsModel.SetID("testString")
				getAPIKeyOptionsModel.SetIncludeHistory(true)
				getAPIKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAPIKeyOptionsModel).ToNot(BeNil())
				Expect(getAPIKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getAPIKeyOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(true)))
				Expect(getAPIKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAPIKeysDetailsOptions successfully`, func() {
				// Construct an instance of the GetAPIKeysDetailsOptions model
				getAPIKeysDetailsOptionsModel := iamIdentityService.NewGetAPIKeysDetailsOptions()
				getAPIKeysDetailsOptionsModel.SetIamAPIKey("testString")
				getAPIKeysDetailsOptionsModel.SetIncludeHistory(true)
				getAPIKeysDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAPIKeysDetailsOptionsModel).ToNot(BeNil())
				Expect(getAPIKeysDetailsOptionsModel.IamAPIKey).To(Equal(core.StringPtr("testString")))
				Expect(getAPIKeysDetailsOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(true)))
				Expect(getAPIKeysDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServiceIDOptions successfully`, func() {
				// Construct an instance of the GetServiceIDOptions model
				id := "testString"
				getServiceIDOptionsModel := iamIdentityService.NewGetServiceIDOptions(id)
				getServiceIDOptionsModel.SetID("testString")
				getServiceIDOptionsModel.SetIncludeHistory(true)
				getServiceIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServiceIDOptionsModel).ToNot(BeNil())
				Expect(getServiceIDOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getServiceIDOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(true)))
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
				listAPIKeysOptionsModel.SetIncludeHistory(true)
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
				Expect(listAPIKeysOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(true)))
				Expect(listAPIKeysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				listServiceIdsOptionsModel.SetIncludeHistory(true)
				listServiceIdsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listServiceIdsOptionsModel).ToNot(BeNil())
				Expect(listServiceIdsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listServiceIdsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listServiceIdsOptionsModel.Pagesize).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listServiceIdsOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listServiceIdsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listServiceIdsOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listServiceIdsOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(true)))
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
				updateAccountSettingsOptionsModel.SetSessionExpirationInSeconds("testString")
				updateAccountSettingsOptionsModel.SetSessionInvalidationInSeconds("testString")
				updateAccountSettingsOptionsModel.SetMaxSessionsPerIdentity("testString")
				updateAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(updateAccountSettingsOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.RestrictCreateServiceID).To(Equal(core.StringPtr("RESTRICTED")))
				Expect(updateAccountSettingsOptionsModel.RestrictCreatePlatformApikey).To(Equal(core.StringPtr("RESTRICTED")))
				Expect(updateAccountSettingsOptionsModel.AllowedIPAddresses).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.Mfa).To(Equal(core.StringPtr("NONE")))
				Expect(updateAccountSettingsOptionsModel.SessionExpirationInSeconds).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.SessionInvalidationInSeconds).To(Equal(core.StringPtr("testString")))
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
