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

package iamidentityv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
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
	Describe(`ListApiKeys(listApiKeysOptions *ListApiKeysOptions) - Operation response error`, func() {
		listApiKeysPath := "/v1/apikeys"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listApiKeysPath))
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
			It(`Invoke ListApiKeys with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListApiKeysOptions model
				listApiKeysOptionsModel := new(iamidentityv1.ListApiKeysOptions)
				listApiKeysOptionsModel.AccountID = core.StringPtr("testString")
				listApiKeysOptionsModel.IamID = core.StringPtr("testString")
				listApiKeysOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listApiKeysOptionsModel.Pagetoken = core.StringPtr("testString")
				listApiKeysOptionsModel.Scope = core.StringPtr("entity")
				listApiKeysOptionsModel.Type = core.StringPtr("user")
				listApiKeysOptionsModel.Sort = core.StringPtr("testString")
				listApiKeysOptionsModel.Order = core.StringPtr("asc")
				listApiKeysOptionsModel.IncludeHistory = core.BoolPtr(true)
				listApiKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.ListApiKeys(listApiKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.ListApiKeys(listApiKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListApiKeys(listApiKeysOptions *ListApiKeysOptions)`, func() {
		listApiKeysPath := "/v1/apikeys"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listApiKeysPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "apikeys": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke ListApiKeys successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.ListApiKeys(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListApiKeysOptions model
				listApiKeysOptionsModel := new(iamidentityv1.ListApiKeysOptions)
				listApiKeysOptionsModel.AccountID = core.StringPtr("testString")
				listApiKeysOptionsModel.IamID = core.StringPtr("testString")
				listApiKeysOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listApiKeysOptionsModel.Pagetoken = core.StringPtr("testString")
				listApiKeysOptionsModel.Scope = core.StringPtr("entity")
				listApiKeysOptionsModel.Type = core.StringPtr("user")
				listApiKeysOptionsModel.Sort = core.StringPtr("testString")
				listApiKeysOptionsModel.Order = core.StringPtr("asc")
				listApiKeysOptionsModel.IncludeHistory = core.BoolPtr(true)
				listApiKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.ListApiKeys(listApiKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.ListApiKeysWithContext(ctx, listApiKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr = iamIdentityService.ListApiKeys(listApiKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.ListApiKeysWithContext(ctx, listApiKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListApiKeys with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ListApiKeysOptions model
				listApiKeysOptionsModel := new(iamidentityv1.ListApiKeysOptions)
				listApiKeysOptionsModel.AccountID = core.StringPtr("testString")
				listApiKeysOptionsModel.IamID = core.StringPtr("testString")
				listApiKeysOptionsModel.Pagesize = core.Int64Ptr(int64(38))
				listApiKeysOptionsModel.Pagetoken = core.StringPtr("testString")
				listApiKeysOptionsModel.Scope = core.StringPtr("entity")
				listApiKeysOptionsModel.Type = core.StringPtr("user")
				listApiKeysOptionsModel.Sort = core.StringPtr("testString")
				listApiKeysOptionsModel.Order = core.StringPtr("asc")
				listApiKeysOptionsModel.IncludeHistory = core.BoolPtr(true)
				listApiKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.ListApiKeys(listApiKeysOptionsModel)
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
	Describe(`CreateApiKey(createApiKeyOptions *CreateApiKeyOptions) - Operation response error`, func() {
		createApiKeyPath := "/v1/apikeys"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApiKeyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateApiKey with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateApiKeyOptions model
				createApiKeyOptionsModel := new(iamidentityv1.CreateApiKeyOptions)
				createApiKeyOptionsModel.Name = core.StringPtr("testString")
				createApiKeyOptionsModel.IamID = core.StringPtr("testString")
				createApiKeyOptionsModel.Description = core.StringPtr("testString")
				createApiKeyOptionsModel.AccountID = core.StringPtr("testString")
				createApiKeyOptionsModel.Apikey = core.StringPtr("testString")
				createApiKeyOptionsModel.StoreValue = core.BoolPtr(true)
				createApiKeyOptionsModel.EntityLock = core.StringPtr("testString")
				createApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateApiKey(createApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateApiKey(createApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateApiKey(createApiKeyOptions *CreateApiKeyOptions)`, func() {
		createApiKeyPath := "/v1/apikeys"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApiKeyPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
				}))
			})
			It(`Invoke CreateApiKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateApiKeyOptions model
				createApiKeyOptionsModel := new(iamidentityv1.CreateApiKeyOptions)
				createApiKeyOptionsModel.Name = core.StringPtr("testString")
				createApiKeyOptionsModel.IamID = core.StringPtr("testString")
				createApiKeyOptionsModel.Description = core.StringPtr("testString")
				createApiKeyOptionsModel.AccountID = core.StringPtr("testString")
				createApiKeyOptionsModel.Apikey = core.StringPtr("testString")
				createApiKeyOptionsModel.StoreValue = core.BoolPtr(true)
				createApiKeyOptionsModel.EntityLock = core.StringPtr("testString")
				createApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateApiKey(createApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.CreateApiKeyWithContext(ctx, createApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr = iamIdentityService.CreateApiKey(createApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.CreateApiKeyWithContext(ctx, createApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateApiKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the CreateApiKeyOptions model
				createApiKeyOptionsModel := new(iamidentityv1.CreateApiKeyOptions)
				createApiKeyOptionsModel.Name = core.StringPtr("testString")
				createApiKeyOptionsModel.IamID = core.StringPtr("testString")
				createApiKeyOptionsModel.Description = core.StringPtr("testString")
				createApiKeyOptionsModel.AccountID = core.StringPtr("testString")
				createApiKeyOptionsModel.Apikey = core.StringPtr("testString")
				createApiKeyOptionsModel.StoreValue = core.BoolPtr(true)
				createApiKeyOptionsModel.EntityLock = core.StringPtr("testString")
				createApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateApiKey(createApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateApiKeyOptions model with no property values
				createApiKeyOptionsModelNew := new(iamidentityv1.CreateApiKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateApiKey(createApiKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApiKeysDetails(getApiKeysDetailsOptions *GetApiKeysDetailsOptions) - Operation response error`, func() {
		getApiKeysDetailsPath := "/v1/apikeys/details"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApiKeysDetailsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Iam-Apikey"]).ToNot(BeNil())
					Expect(req.Header["Iam-Apikey"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for include_history query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApiKeysDetails with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetApiKeysDetailsOptions model
				getApiKeysDetailsOptionsModel := new(iamidentityv1.GetApiKeysDetailsOptions)
				getApiKeysDetailsOptionsModel.IAMApiKey = core.StringPtr("testString")
				getApiKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(true)
				getApiKeysDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetApiKeysDetails(getApiKeysDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetApiKeysDetails(getApiKeysDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetApiKeysDetails(getApiKeysDetailsOptions *GetApiKeysDetailsOptions)`, func() {
		getApiKeysDetailsPath := "/v1/apikeys/details"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApiKeysDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Iam-Apikey"]).ToNot(BeNil())
					Expect(req.Header["Iam-Apikey"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for include_history query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
				}))
			})
			It(`Invoke GetApiKeysDetails successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetApiKeysDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApiKeysDetailsOptions model
				getApiKeysDetailsOptionsModel := new(iamidentityv1.GetApiKeysDetailsOptions)
				getApiKeysDetailsOptionsModel.IAMApiKey = core.StringPtr("testString")
				getApiKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(true)
				getApiKeysDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetApiKeysDetails(getApiKeysDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.GetApiKeysDetailsWithContext(ctx, getApiKeysDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr = iamIdentityService.GetApiKeysDetails(getApiKeysDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.GetApiKeysDetailsWithContext(ctx, getApiKeysDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetApiKeysDetails with error: Operation request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetApiKeysDetailsOptions model
				getApiKeysDetailsOptionsModel := new(iamidentityv1.GetApiKeysDetailsOptions)
				getApiKeysDetailsOptionsModel.IAMApiKey = core.StringPtr("testString")
				getApiKeysDetailsOptionsModel.IncludeHistory = core.BoolPtr(true)
				getApiKeysDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetApiKeysDetails(getApiKeysDetailsOptionsModel)
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
	Describe(`GetApiKey(getApiKeyOptions *GetApiKeyOptions) - Operation response error`, func() {
		getApiKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApiKeyPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for include_history query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApiKey with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetApiKeyOptions model
				getApiKeyOptionsModel := new(iamidentityv1.GetApiKeyOptions)
				getApiKeyOptionsModel.ID = core.StringPtr("testString")
				getApiKeyOptionsModel.IncludeHistory = core.BoolPtr(true)
				getApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetApiKey(getApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetApiKey(getApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetApiKey(getApiKeyOptions *GetApiKeyOptions)`, func() {
		getApiKeyPath := "/v1/apikeys/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApiKeyPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for include_history query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
				}))
			})
			It(`Invoke GetApiKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApiKeyOptions model
				getApiKeyOptionsModel := new(iamidentityv1.GetApiKeyOptions)
				getApiKeyOptionsModel.ID = core.StringPtr("testString")
				getApiKeyOptionsModel.IncludeHistory = core.BoolPtr(true)
				getApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetApiKey(getApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.GetApiKeyWithContext(ctx, getApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr = iamIdentityService.GetApiKey(getApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.GetApiKeyWithContext(ctx, getApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetApiKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetApiKeyOptions model
				getApiKeyOptionsModel := new(iamidentityv1.GetApiKeyOptions)
				getApiKeyOptionsModel.ID = core.StringPtr("testString")
				getApiKeyOptionsModel.IncludeHistory = core.BoolPtr(true)
				getApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetApiKey(getApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetApiKeyOptions model with no property values
				getApiKeyOptionsModelNew := new(iamidentityv1.GetApiKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetApiKey(getApiKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateApiKey(updateApiKeyOptions *UpdateApiKeyOptions) - Operation response error`, func() {
		updateApiKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateApiKeyPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateApiKey with error: Operation response processing error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateApiKeyOptions model
				updateApiKeyOptionsModel := new(iamidentityv1.UpdateApiKeyOptions)
				updateApiKeyOptionsModel.ID = core.StringPtr("testString")
				updateApiKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateApiKeyOptionsModel.Name = core.StringPtr("testString")
				updateApiKeyOptionsModel.Description = core.StringPtr("testString")
				updateApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateApiKey(updateApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateApiKey(updateApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateApiKey(updateApiKeyOptions *UpdateApiKeyOptions)`, func() {
		updateApiKeyPath := "/v1/apikeys/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateApiKeyPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}`)
				}))
			})
			It(`Invoke UpdateApiKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateApiKeyOptions model
				updateApiKeyOptionsModel := new(iamidentityv1.UpdateApiKeyOptions)
				updateApiKeyOptionsModel.ID = core.StringPtr("testString")
				updateApiKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateApiKeyOptionsModel.Name = core.StringPtr("testString")
				updateApiKeyOptionsModel.Description = core.StringPtr("testString")
				updateApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateApiKey(updateApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.UpdateApiKeyWithContext(ctx, updateApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr = iamIdentityService.UpdateApiKey(updateApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.UpdateApiKeyWithContext(ctx, updateApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateApiKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateApiKeyOptions model
				updateApiKeyOptionsModel := new(iamidentityv1.UpdateApiKeyOptions)
				updateApiKeyOptionsModel.ID = core.StringPtr("testString")
				updateApiKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateApiKeyOptionsModel.Name = core.StringPtr("testString")
				updateApiKeyOptionsModel.Description = core.StringPtr("testString")
				updateApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateApiKey(updateApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateApiKeyOptions model with no property values
				updateApiKeyOptionsModelNew := new(iamidentityv1.UpdateApiKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateApiKey(updateApiKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteApiKey(deleteApiKeyOptions *DeleteApiKeyOptions)`, func() {
		deleteApiKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteApiKeyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteApiKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteApiKeyOptions model
				deleteApiKeyOptionsModel := new(iamidentityv1.DeleteApiKeyOptions)
				deleteApiKeyOptionsModel.ID = core.StringPtr("testString")
				deleteApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteApiKey(deleteApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				response, operationErr = iamIdentityService.DeleteApiKey(deleteApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteApiKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the DeleteApiKeyOptions model
				deleteApiKeyOptionsModel := new(iamidentityv1.DeleteApiKeyOptions)
				deleteApiKeyOptionsModel.ID = core.StringPtr("testString")
				deleteApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteApiKey(deleteApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteApiKeyOptions model with no property values
				deleteApiKeyOptionsModelNew := new(iamidentityv1.DeleteApiKeyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteApiKey(deleteApiKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`LockApiKey(lockApiKeyOptions *LockApiKeyOptions)`, func() {
		lockApiKeyPath := "/v1/apikeys/testString/lock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(lockApiKeyPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke LockApiKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.LockApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the LockApiKeyOptions model
				lockApiKeyOptionsModel := new(iamidentityv1.LockApiKeyOptions)
				lockApiKeyOptionsModel.ID = core.StringPtr("testString")
				lockApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.LockApiKey(lockApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				response, operationErr = iamIdentityService.LockApiKey(lockApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke LockApiKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the LockApiKeyOptions model
				lockApiKeyOptionsModel := new(iamidentityv1.LockApiKeyOptions)
				lockApiKeyOptionsModel.ID = core.StringPtr("testString")
				lockApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.LockApiKey(lockApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the LockApiKeyOptions model with no property values
				lockApiKeyOptionsModelNew := new(iamidentityv1.LockApiKeyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.LockApiKey(lockApiKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UnlockApiKey(unlockApiKeyOptions *UnlockApiKeyOptions)`, func() {
		unlockApiKeyPath := "/v1/apikeys/testString/lock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unlockApiKeyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UnlockApiKey successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.UnlockApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UnlockApiKeyOptions model
				unlockApiKeyOptionsModel := new(iamidentityv1.UnlockApiKeyOptions)
				unlockApiKeyOptionsModel.ID = core.StringPtr("testString")
				unlockApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.UnlockApiKey(unlockApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				response, operationErr = iamIdentityService.UnlockApiKey(unlockApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UnlockApiKey with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UnlockApiKeyOptions model
				unlockApiKeyOptionsModel := new(iamidentityv1.UnlockApiKeyOptions)
				unlockApiKeyOptionsModel.ID = core.StringPtr("testString")
				unlockApiKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.UnlockApiKey(unlockApiKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UnlockApiKeyOptions model with no property values
				unlockApiKeyOptionsModelNew := new(iamidentityv1.UnlockApiKeyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.UnlockApiKey(unlockApiKeyOptionsModelNew)
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "serviceids": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}]}`)
				}))
			})
			It(`Invoke ListServiceIds successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

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

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.ListServiceIdsWithContext(ctx, listServiceIdsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr = iamIdentityService.ListServiceIds(listServiceIdsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.ListServiceIdsWithContext(ctx, listServiceIdsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
	Describe(`CreateServiceID(createServiceIdOptions *CreateServiceIdOptions) - Operation response error`, func() {
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

				// Construct an instance of the ApiKeyInsideCreateServiceIdRequest model
				apiKeyInsideCreateServiceIdRequestModel := new(iamidentityv1.ApiKeyInsideCreateServiceIdRequest)
				apiKeyInsideCreateServiceIdRequestModel.Name = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.Description = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.Apikey = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.StoreValue = core.BoolPtr(true)

				// Construct an instance of the CreateServiceIdOptions model
				createServiceIdOptionsModel := new(iamidentityv1.CreateServiceIdOptions)
				createServiceIdOptionsModel.AccountID = core.StringPtr("testString")
				createServiceIdOptionsModel.Name = core.StringPtr("testString")
				createServiceIdOptionsModel.Description = core.StringPtr("testString")
				createServiceIdOptionsModel.UniqueInstanceCrns = []string{"testString"}
				createServiceIdOptionsModel.Apikey = apiKeyInsideCreateServiceIdRequestModel
				createServiceIdOptionsModel.EntityLock = core.StringPtr("testString")
				createServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.CreateServiceID(createServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.CreateServiceID(createServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateServiceID(createServiceIdOptions *CreateServiceIdOptions)`, func() {
		createServiceIDPath := "/v1/serviceids/"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}`)
				}))
			})
			It(`Invoke CreateServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.CreateServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ApiKeyInsideCreateServiceIdRequest model
				apiKeyInsideCreateServiceIdRequestModel := new(iamidentityv1.ApiKeyInsideCreateServiceIdRequest)
				apiKeyInsideCreateServiceIdRequestModel.Name = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.Description = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.Apikey = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.StoreValue = core.BoolPtr(true)

				// Construct an instance of the CreateServiceIdOptions model
				createServiceIdOptionsModel := new(iamidentityv1.CreateServiceIdOptions)
				createServiceIdOptionsModel.AccountID = core.StringPtr("testString")
				createServiceIdOptionsModel.Name = core.StringPtr("testString")
				createServiceIdOptionsModel.Description = core.StringPtr("testString")
				createServiceIdOptionsModel.UniqueInstanceCrns = []string{"testString"}
				createServiceIdOptionsModel.Apikey = apiKeyInsideCreateServiceIdRequestModel
				createServiceIdOptionsModel.EntityLock = core.StringPtr("testString")
				createServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.CreateServiceID(createServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.CreateServiceIDWithContext(ctx, createServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr = iamIdentityService.CreateServiceID(createServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.CreateServiceIDWithContext(ctx, createServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateServiceID with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the ApiKeyInsideCreateServiceIdRequest model
				apiKeyInsideCreateServiceIdRequestModel := new(iamidentityv1.ApiKeyInsideCreateServiceIdRequest)
				apiKeyInsideCreateServiceIdRequestModel.Name = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.Description = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.Apikey = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.StoreValue = core.BoolPtr(true)

				// Construct an instance of the CreateServiceIdOptions model
				createServiceIdOptionsModel := new(iamidentityv1.CreateServiceIdOptions)
				createServiceIdOptionsModel.AccountID = core.StringPtr("testString")
				createServiceIdOptionsModel.Name = core.StringPtr("testString")
				createServiceIdOptionsModel.Description = core.StringPtr("testString")
				createServiceIdOptionsModel.UniqueInstanceCrns = []string{"testString"}
				createServiceIdOptionsModel.Apikey = apiKeyInsideCreateServiceIdRequestModel
				createServiceIdOptionsModel.EntityLock = core.StringPtr("testString")
				createServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.CreateServiceID(createServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateServiceIdOptions model with no property values
				createServiceIdOptionsModelNew := new(iamidentityv1.CreateServiceIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.CreateServiceID(createServiceIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetServiceID(getServiceIdOptions *GetServiceIdOptions) - Operation response error`, func() {
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

				// Construct an instance of the GetServiceIdOptions model
				getServiceIdOptionsModel := new(iamidentityv1.GetServiceIdOptions)
				getServiceIdOptionsModel.ID = core.StringPtr("testString")
				getServiceIdOptionsModel.IncludeHistory = core.BoolPtr(true)
				getServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.GetServiceID(getServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.GetServiceID(getServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetServiceID(getServiceIdOptions *GetServiceIdOptions)`, func() {
		getServiceIDPath := "/v1/serviceids/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceIDPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for include_history query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}`)
				}))
			})
			It(`Invoke GetServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.GetServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServiceIdOptions model
				getServiceIdOptionsModel := new(iamidentityv1.GetServiceIdOptions)
				getServiceIdOptionsModel.ID = core.StringPtr("testString")
				getServiceIdOptionsModel.IncludeHistory = core.BoolPtr(true)
				getServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.GetServiceID(getServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.GetServiceIDWithContext(ctx, getServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr = iamIdentityService.GetServiceID(getServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.GetServiceIDWithContext(ctx, getServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetServiceID with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the GetServiceIdOptions model
				getServiceIdOptionsModel := new(iamidentityv1.GetServiceIdOptions)
				getServiceIdOptionsModel.ID = core.StringPtr("testString")
				getServiceIdOptionsModel.IncludeHistory = core.BoolPtr(true)
				getServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.GetServiceID(getServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetServiceIdOptions model with no property values
				getServiceIdOptionsModelNew := new(iamidentityv1.GetServiceIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.GetServiceID(getServiceIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateServiceID(updateServiceIdOptions *UpdateServiceIdOptions) - Operation response error`, func() {
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

				// Construct an instance of the UpdateServiceIdOptions model
				updateServiceIdOptionsModel := new(iamidentityv1.UpdateServiceIdOptions)
				updateServiceIdOptionsModel.ID = core.StringPtr("testString")
				updateServiceIdOptionsModel.IfMatch = core.StringPtr("testString")
				updateServiceIdOptionsModel.Name = core.StringPtr("testString")
				updateServiceIdOptionsModel.Description = core.StringPtr("testString")
				updateServiceIdOptionsModel.UniqueInstanceCrns = []string{"testString"}
				updateServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamIdentityService.UpdateServiceID(updateServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamIdentityService.EnableRetries(0, 0)
				result, response, operationErr = iamIdentityService.UpdateServiceID(updateServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateServiceID(updateServiceIdOptions *UpdateServiceIdOptions)`, func() {
		updateServiceIDPath := "/v1/serviceids/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey", "history": [{"timestamp": "Timestamp", "iam_id": "IamID", "iam_id_account": "IamIdAccount", "action": "Action", "params": ["Params"], "message": "Message"}]}}`)
				}))
			})
			It(`Invoke UpdateServiceID successfully`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamIdentityService.UpdateServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateServiceIdOptions model
				updateServiceIdOptionsModel := new(iamidentityv1.UpdateServiceIdOptions)
				updateServiceIdOptionsModel.ID = core.StringPtr("testString")
				updateServiceIdOptionsModel.IfMatch = core.StringPtr("testString")
				updateServiceIdOptionsModel.Name = core.StringPtr("testString")
				updateServiceIdOptionsModel.Description = core.StringPtr("testString")
				updateServiceIdOptionsModel.UniqueInstanceCrns = []string{"testString"}
				updateServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamIdentityService.UpdateServiceID(updateServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.UpdateServiceIDWithContext(ctx, updateServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				result, response, operationErr = iamIdentityService.UpdateServiceID(updateServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamIdentityService.UpdateServiceIDWithContext(ctx, updateServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateServiceID with error: Operation validation and request error`, func() {
				iamIdentityService, serviceErr := iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamIdentityService).ToNot(BeNil())

				// Construct an instance of the UpdateServiceIdOptions model
				updateServiceIdOptionsModel := new(iamidentityv1.UpdateServiceIdOptions)
				updateServiceIdOptionsModel.ID = core.StringPtr("testString")
				updateServiceIdOptionsModel.IfMatch = core.StringPtr("testString")
				updateServiceIdOptionsModel.Name = core.StringPtr("testString")
				updateServiceIdOptionsModel.Description = core.StringPtr("testString")
				updateServiceIdOptionsModel.UniqueInstanceCrns = []string{"testString"}
				updateServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamIdentityService.UpdateServiceID(updateServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateServiceIdOptions model with no property values
				updateServiceIdOptionsModelNew := new(iamidentityv1.UpdateServiceIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamIdentityService.UpdateServiceID(updateServiceIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteServiceID(deleteServiceIdOptions *DeleteServiceIdOptions)`, func() {
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
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.DeleteServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteServiceIdOptions model
				deleteServiceIdOptionsModel := new(iamidentityv1.DeleteServiceIdOptions)
				deleteServiceIdOptionsModel.ID = core.StringPtr("testString")
				deleteServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.DeleteServiceID(deleteServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				response, operationErr = iamIdentityService.DeleteServiceID(deleteServiceIdOptionsModel)
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

				// Construct an instance of the DeleteServiceIdOptions model
				deleteServiceIdOptionsModel := new(iamidentityv1.DeleteServiceIdOptions)
				deleteServiceIdOptionsModel.ID = core.StringPtr("testString")
				deleteServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.DeleteServiceID(deleteServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteServiceIdOptions model with no property values
				deleteServiceIdOptionsModelNew := new(iamidentityv1.DeleteServiceIdOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.DeleteServiceID(deleteServiceIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`LockServiceID(lockServiceIdOptions *LockServiceIdOptions)`, func() {
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
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.LockServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the LockServiceIdOptions model
				lockServiceIdOptionsModel := new(iamidentityv1.LockServiceIdOptions)
				lockServiceIdOptionsModel.ID = core.StringPtr("testString")
				lockServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.LockServiceID(lockServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				response, operationErr = iamIdentityService.LockServiceID(lockServiceIdOptionsModel)
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

				// Construct an instance of the LockServiceIdOptions model
				lockServiceIdOptionsModel := new(iamidentityv1.LockServiceIdOptions)
				lockServiceIdOptionsModel.ID = core.StringPtr("testString")
				lockServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.LockServiceID(lockServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the LockServiceIdOptions model with no property values
				lockServiceIdOptionsModelNew := new(iamidentityv1.LockServiceIdOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.LockServiceID(lockServiceIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UnlockServiceID(unlockServiceIdOptions *UnlockServiceIdOptions)`, func() {
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
				iamIdentityService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamIdentityService.UnlockServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UnlockServiceIdOptions model
				unlockServiceIdOptionsModel := new(iamidentityv1.UnlockServiceIdOptions)
				unlockServiceIdOptionsModel.ID = core.StringPtr("testString")
				unlockServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamIdentityService.UnlockServiceID(unlockServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				iamIdentityService.DisableRetries()
				response, operationErr = iamIdentityService.UnlockServiceID(unlockServiceIdOptionsModel)
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

				// Construct an instance of the UnlockServiceIdOptions model
				unlockServiceIdOptionsModel := new(iamidentityv1.UnlockServiceIdOptions)
				unlockServiceIdOptionsModel.ID = core.StringPtr("testString")
				unlockServiceIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamIdentityService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamIdentityService.UnlockServiceID(unlockServiceIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UnlockServiceIdOptions model with no property values
				unlockServiceIdOptionsModelNew := new(iamidentityv1.UnlockServiceIdOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamIdentityService.UnlockServiceID(unlockServiceIdOptionsModelNew)
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
			It(`Invoke NewApiKeyInsideCreateServiceIdRequest successfully`, func() {
				name := "testString"
				model, err := iamIdentityService.NewApiKeyInsideCreateServiceIdRequest(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateApiKeyOptions successfully`, func() {
				// Construct an instance of the CreateApiKeyOptions model
				createApiKeyOptionsName := "testString"
				createApiKeyOptionsIamID := "testString"
				createApiKeyOptionsModel := iamIdentityService.NewCreateApiKeyOptions(createApiKeyOptionsName, createApiKeyOptionsIamID)
				createApiKeyOptionsModel.SetName("testString")
				createApiKeyOptionsModel.SetIamID("testString")
				createApiKeyOptionsModel.SetDescription("testString")
				createApiKeyOptionsModel.SetAccountID("testString")
				createApiKeyOptionsModel.SetApikey("testString")
				createApiKeyOptionsModel.SetStoreValue(true)
				createApiKeyOptionsModel.SetEntityLock("testString")
				createApiKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createApiKeyOptionsModel).ToNot(BeNil())
				Expect(createApiKeyOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createApiKeyOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(createApiKeyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createApiKeyOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createApiKeyOptionsModel.Apikey).To(Equal(core.StringPtr("testString")))
				Expect(createApiKeyOptionsModel.StoreValue).To(Equal(core.BoolPtr(true)))
				Expect(createApiKeyOptionsModel.EntityLock).To(Equal(core.StringPtr("testString")))
				Expect(createApiKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateServiceIdOptions successfully`, func() {
				// Construct an instance of the ApiKeyInsideCreateServiceIdRequest model
				apiKeyInsideCreateServiceIdRequestModel := new(iamidentityv1.ApiKeyInsideCreateServiceIdRequest)
				Expect(apiKeyInsideCreateServiceIdRequestModel).ToNot(BeNil())
				apiKeyInsideCreateServiceIdRequestModel.Name = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.Description = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.Apikey = core.StringPtr("testString")
				apiKeyInsideCreateServiceIdRequestModel.StoreValue = core.BoolPtr(true)
				Expect(apiKeyInsideCreateServiceIdRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(apiKeyInsideCreateServiceIdRequestModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(apiKeyInsideCreateServiceIdRequestModel.Apikey).To(Equal(core.StringPtr("testString")))
				Expect(apiKeyInsideCreateServiceIdRequestModel.StoreValue).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreateServiceIdOptions model
				createServiceIdOptionsAccountID := "testString"
				createServiceIdOptionsName := "testString"
				createServiceIdOptionsModel := iamIdentityService.NewCreateServiceIdOptions(createServiceIdOptionsAccountID, createServiceIdOptionsName)
				createServiceIdOptionsModel.SetAccountID("testString")
				createServiceIdOptionsModel.SetName("testString")
				createServiceIdOptionsModel.SetDescription("testString")
				createServiceIdOptionsModel.SetUniqueInstanceCrns([]string{"testString"})
				createServiceIdOptionsModel.SetApikey(apiKeyInsideCreateServiceIdRequestModel)
				createServiceIdOptionsModel.SetEntityLock("testString")
				createServiceIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createServiceIdOptionsModel).ToNot(BeNil())
				Expect(createServiceIdOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createServiceIdOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createServiceIdOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createServiceIdOptionsModel.UniqueInstanceCrns).To(Equal([]string{"testString"}))
				Expect(createServiceIdOptionsModel.Apikey).To(Equal(apiKeyInsideCreateServiceIdRequestModel))
				Expect(createServiceIdOptionsModel.EntityLock).To(Equal(core.StringPtr("testString")))
				Expect(createServiceIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteApiKeyOptions successfully`, func() {
				// Construct an instance of the DeleteApiKeyOptions model
				id := "testString"
				deleteApiKeyOptionsModel := iamIdentityService.NewDeleteApiKeyOptions(id)
				deleteApiKeyOptionsModel.SetID("testString")
				deleteApiKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteApiKeyOptionsModel).ToNot(BeNil())
				Expect(deleteApiKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteApiKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteServiceIdOptions successfully`, func() {
				// Construct an instance of the DeleteServiceIdOptions model
				id := "testString"
				deleteServiceIdOptionsModel := iamIdentityService.NewDeleteServiceIdOptions(id)
				deleteServiceIdOptionsModel.SetID("testString")
				deleteServiceIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteServiceIdOptionsModel).ToNot(BeNil())
				Expect(deleteServiceIdOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApiKeyOptions successfully`, func() {
				// Construct an instance of the GetApiKeyOptions model
				id := "testString"
				getApiKeyOptionsModel := iamIdentityService.NewGetApiKeyOptions(id)
				getApiKeyOptionsModel.SetID("testString")
				getApiKeyOptionsModel.SetIncludeHistory(true)
				getApiKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApiKeyOptionsModel).ToNot(BeNil())
				Expect(getApiKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getApiKeyOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(true)))
				Expect(getApiKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApiKeysDetailsOptions successfully`, func() {
				// Construct an instance of the GetApiKeysDetailsOptions model
				getApiKeysDetailsOptionsModel := iamIdentityService.NewGetApiKeysDetailsOptions()
				getApiKeysDetailsOptionsModel.SetIAMApiKey("testString")
				getApiKeysDetailsOptionsModel.SetIncludeHistory(true)
				getApiKeysDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApiKeysDetailsOptionsModel).ToNot(BeNil())
				Expect(getApiKeysDetailsOptionsModel.IAMApiKey).To(Equal(core.StringPtr("testString")))
				Expect(getApiKeysDetailsOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(true)))
				Expect(getApiKeysDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServiceIdOptions successfully`, func() {
				// Construct an instance of the GetServiceIdOptions model
				id := "testString"
				getServiceIdOptionsModel := iamIdentityService.NewGetServiceIdOptions(id)
				getServiceIdOptionsModel.SetID("testString")
				getServiceIdOptionsModel.SetIncludeHistory(true)
				getServiceIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServiceIdOptionsModel).ToNot(BeNil())
				Expect(getServiceIdOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getServiceIdOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(true)))
				Expect(getServiceIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListApiKeysOptions successfully`, func() {
				// Construct an instance of the ListApiKeysOptions model
				listApiKeysOptionsModel := iamIdentityService.NewListApiKeysOptions()
				listApiKeysOptionsModel.SetAccountID("testString")
				listApiKeysOptionsModel.SetIamID("testString")
				listApiKeysOptionsModel.SetPagesize(int64(38))
				listApiKeysOptionsModel.SetPagetoken("testString")
				listApiKeysOptionsModel.SetScope("entity")
				listApiKeysOptionsModel.SetType("user")
				listApiKeysOptionsModel.SetSort("testString")
				listApiKeysOptionsModel.SetOrder("asc")
				listApiKeysOptionsModel.SetIncludeHistory(true)
				listApiKeysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listApiKeysOptionsModel).ToNot(BeNil())
				Expect(listApiKeysOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listApiKeysOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(listApiKeysOptionsModel.Pagesize).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listApiKeysOptionsModel.Pagetoken).To(Equal(core.StringPtr("testString")))
				Expect(listApiKeysOptionsModel.Scope).To(Equal(core.StringPtr("entity")))
				Expect(listApiKeysOptionsModel.Type).To(Equal(core.StringPtr("user")))
				Expect(listApiKeysOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listApiKeysOptionsModel.Order).To(Equal(core.StringPtr("asc")))
				Expect(listApiKeysOptionsModel.IncludeHistory).To(Equal(core.BoolPtr(true)))
				Expect(listApiKeysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewLockApiKeyOptions successfully`, func() {
				// Construct an instance of the LockApiKeyOptions model
				id := "testString"
				lockApiKeyOptionsModel := iamIdentityService.NewLockApiKeyOptions(id)
				lockApiKeyOptionsModel.SetID("testString")
				lockApiKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(lockApiKeyOptionsModel).ToNot(BeNil())
				Expect(lockApiKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(lockApiKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLockServiceIdOptions successfully`, func() {
				// Construct an instance of the LockServiceIdOptions model
				id := "testString"
				lockServiceIdOptionsModel := iamIdentityService.NewLockServiceIdOptions(id)
				lockServiceIdOptionsModel.SetID("testString")
				lockServiceIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(lockServiceIdOptionsModel).ToNot(BeNil())
				Expect(lockServiceIdOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(lockServiceIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUnlockApiKeyOptions successfully`, func() {
				// Construct an instance of the UnlockApiKeyOptions model
				id := "testString"
				unlockApiKeyOptionsModel := iamIdentityService.NewUnlockApiKeyOptions(id)
				unlockApiKeyOptionsModel.SetID("testString")
				unlockApiKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unlockApiKeyOptionsModel).ToNot(BeNil())
				Expect(unlockApiKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(unlockApiKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUnlockServiceIdOptions successfully`, func() {
				// Construct an instance of the UnlockServiceIdOptions model
				id := "testString"
				unlockServiceIdOptionsModel := iamIdentityService.NewUnlockServiceIdOptions(id)
				unlockServiceIdOptionsModel.SetID("testString")
				unlockServiceIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unlockServiceIdOptionsModel).ToNot(BeNil())
				Expect(unlockServiceIdOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(unlockServiceIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateApiKeyOptions successfully`, func() {
				// Construct an instance of the UpdateApiKeyOptions model
				id := "testString"
				ifMatch := "testString"
				updateApiKeyOptionsModel := iamIdentityService.NewUpdateApiKeyOptions(id, ifMatch)
				updateApiKeyOptionsModel.SetID("testString")
				updateApiKeyOptionsModel.SetIfMatch("testString")
				updateApiKeyOptionsModel.SetName("testString")
				updateApiKeyOptionsModel.SetDescription("testString")
				updateApiKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateApiKeyOptionsModel).ToNot(BeNil())
				Expect(updateApiKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateApiKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateApiKeyOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateApiKeyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateApiKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateServiceIdOptions successfully`, func() {
				// Construct an instance of the UpdateServiceIdOptions model
				id := "testString"
				ifMatch := "testString"
				updateServiceIdOptionsModel := iamIdentityService.NewUpdateServiceIdOptions(id, ifMatch)
				updateServiceIdOptionsModel.SetID("testString")
				updateServiceIdOptionsModel.SetIfMatch("testString")
				updateServiceIdOptionsModel.SetName("testString")
				updateServiceIdOptionsModel.SetDescription("testString")
				updateServiceIdOptionsModel.SetUniqueInstanceCrns([]string{"testString"})
				updateServiceIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateServiceIdOptionsModel).ToNot(BeNil())
				Expect(updateServiceIdOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceIdOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceIdOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceIdOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceIdOptionsModel.UniqueInstanceCrns).To(Equal([]string{"testString"}))
				Expect(updateServiceIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
