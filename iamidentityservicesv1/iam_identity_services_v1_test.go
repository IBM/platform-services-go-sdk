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

package iamidentityservicesv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/iamidentityservicesv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`IamIdentityServicesV1`, func() {
	Describe(`ListApiKeys(listApiKeysOptions *ListApiKeysOptions)`, func() {
		listApiKeysPath := "/v1/apikeys"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listApiKeysPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["pagesize"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"context": {"requestId": "RequestID", "requestType": "RequestType", "userAgent": "UserAgent", "clientIp": "ClientIp", "url": "URL", "instanceId": "InstanceID", "threadId": "ThreadID", "host": "Host", "startTime": "StartTime", "endTime": "EndTime", "elapsedTime": "ElapsedTime", "locale": "Locale", "clusterName": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "apikeys": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}]}`)
			}))
			It(`Invoke ListApiKeys successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListApiKeys(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListApiKeysOptions model
				listApiKeysOptionsModel := new(iamidentityservicesv1.ListApiKeysOptions)
				listApiKeysOptionsModel.AccountID = core.StringPtr("testString")
				listApiKeysOptionsModel.IamID = core.StringPtr("testString")
				listApiKeysOptionsModel.Pagesize = core.StringPtr("testString")
				listApiKeysOptionsModel.Pagetoken = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListApiKeys(listApiKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateApiKey(createApiKeyOptions *CreateApiKeyOptions)`, func() {
		createApiKeyPath := "/v1/apikeys"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createApiKeyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
				Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}`)
			}))
			It(`Invoke CreateApiKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateApiKeyOptions model
				createApiKeyOptionsModel := new(iamidentityservicesv1.CreateApiKeyOptions)
				createApiKeyOptionsModel.Name = core.StringPtr("testString")
				createApiKeyOptionsModel.IamID = core.StringPtr("testString")
				createApiKeyOptionsModel.Description = core.StringPtr("testString")
				createApiKeyOptionsModel.AccountID = core.StringPtr("testString")
				createApiKeyOptionsModel.Apikey = core.StringPtr("testString")
				createApiKeyOptionsModel.EntityLock = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateApiKey(createApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetApiKeyDetails(getApiKeyDetailsOptions *GetApiKeyDetailsOptions)`, func() {
		getApiKeyDetailsPath := "/v1/apikeys/details"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getApiKeyDetailsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Iam-Apikey"]).ToNot(BeNil())
				Expect(req.Header["Iam-Apikey"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}`)
			}))
			It(`Invoke GetApiKeyDetails successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetApiKeyDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApiKeyDetailsOptions model
				getApiKeyDetailsOptionsModel := new(iamidentityservicesv1.GetApiKeyDetailsOptions)
				getApiKeyDetailsOptionsModel.IAMApiKey = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetApiKeyDetails(getApiKeyDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetApiKey(getApiKeyOptions *GetApiKeyOptions)`, func() {
		getApiKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getApiKeyPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}`)
			}))
			It(`Invoke GetApiKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApiKeyOptions model
				getApiKeyOptionsModel := new(iamidentityservicesv1.GetApiKeyOptions)
				getApiKeyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetApiKey(getApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateApiKey(updateApiKeyOptions *UpdateApiKeyOptions)`, func() {
		updateApiKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateApiKeyPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["If-Match"]).ToNot(BeNil())
				Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}`)
			}))
			It(`Invoke UpdateApiKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateApiKeyOptions model
				updateApiKeyOptionsModel := new(iamidentityservicesv1.UpdateApiKeyOptions)
				updateApiKeyOptionsModel.ID = core.StringPtr("testString")
				updateApiKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateApiKeyOptionsModel.Name = core.StringPtr("testString")
				updateApiKeyOptionsModel.Description = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateApiKey(updateApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteApiKey(deleteApiKeyOptions *DeleteApiKeyOptions)`, func() {
		deleteApiKeyPath := "/v1/apikeys/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteApiKeyPath))
				Expect(req.Method).To(Equal("DELETE"))
				res.WriteHeader(204)
			}))
			It(`Invoke DeleteApiKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteApiKeyOptions model
				deleteApiKeyOptionsModel := new(iamidentityservicesv1.DeleteApiKeyOptions)
				deleteApiKeyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteApiKey(deleteApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`LockApiKey(lockApiKeyOptions *LockApiKeyOptions)`, func() {
		lockApiKeyPath := "/v1/apikeys/testString/lock"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(lockApiKeyPath))
				Expect(req.Method).To(Equal("POST"))
				res.WriteHeader(204)
			}))
			It(`Invoke LockApiKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.LockApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the LockApiKeyOptions model
				lockApiKeyOptionsModel := new(iamidentityservicesv1.LockApiKeyOptions)
				lockApiKeyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.LockApiKey(lockApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`UnlockApiKey(unlockApiKeyOptions *UnlockApiKeyOptions)`, func() {
		unlockApiKeyPath := "/v1/apikeys/testString/lock"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(unlockApiKeyPath))
				Expect(req.Method).To(Equal("DELETE"))
				res.WriteHeader(204)
			}))
			It(`Invoke UnlockApiKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UnlockApiKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UnlockApiKeyOptions model
				unlockApiKeyOptionsModel := new(iamidentityservicesv1.UnlockApiKeyOptions)
				unlockApiKeyOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UnlockApiKey(unlockApiKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListServiceIds(listServiceIdsOptions *ListServiceIdsOptions)`, func() {
		listServiceIdsPath := "/v1/serviceids"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listServiceIdsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["pagesize"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["pagetoken"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["order"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"context": {"requestId": "RequestID", "requestType": "RequestType", "userAgent": "UserAgent", "clientIp": "ClientIp", "url": "URL", "instanceId": "InstanceID", "threadId": "ThreadID", "host": "Host", "startTime": "StartTime", "endTime": "EndTime", "elapsedTime": "ElapsedTime", "locale": "Locale", "clusterName": "ClusterName"}, "offset": 6, "limit": 5, "first": "First", "previous": "Previous", "next": "Next", "serviceids": [{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}}]}`)
			}))
			It(`Invoke ListServiceIds successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListServiceIds(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListServiceIdsOptions model
				listServiceIdsOptionsModel := new(iamidentityservicesv1.ListServiceIdsOptions)
				listServiceIdsOptionsModel.AccountID = core.StringPtr("testString")
				listServiceIdsOptionsModel.Name = core.StringPtr("testString")
				listServiceIdsOptionsModel.Pagesize = core.StringPtr("testString")
				listServiceIdsOptionsModel.Pagetoken = core.StringPtr("testString")
				listServiceIdsOptionsModel.Sort = core.StringPtr("testString")
				listServiceIdsOptionsModel.Order = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListServiceIds(listServiceIdsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateServiceID(createServiceIdOptions *CreateServiceIdOptions)`, func() {
		createServiceIDPath := "/v1/serviceids"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createServiceIDPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
				Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}}`)
			}))
			It(`Invoke CreateServiceID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateApiKeyRequest model
				createApiKeyRequestModel := new(iamidentityservicesv1.CreateApiKeyRequest)
				createApiKeyRequestModel.Name = core.StringPtr("testString")
				createApiKeyRequestModel.Description = core.StringPtr("testString")
				createApiKeyRequestModel.IamID = core.StringPtr("testString")
				createApiKeyRequestModel.AccountID = core.StringPtr("testString")
				createApiKeyRequestModel.Apikey = core.StringPtr("testString")

				// Construct an instance of the CreateServiceIdOptions model
				createServiceIdOptionsModel := new(iamidentityservicesv1.CreateServiceIdOptions)
				createServiceIdOptionsModel.AccountID = core.StringPtr("testString")
				createServiceIdOptionsModel.Name = core.StringPtr("testString")
				createServiceIdOptionsModel.Description = core.StringPtr("testString")
				createServiceIdOptionsModel.UniqueInstanceCrns = []string{"testString"}
				createServiceIdOptionsModel.Apikey = createApiKeyRequestModel
				createServiceIdOptionsModel.EntityLock = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateServiceID(createServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetServiceID(getServiceIdOptions *GetServiceIdOptions)`, func() {
		getServiceIDPath := "/v1/serviceids/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getServiceIDPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}}`)
			}))
			It(`Invoke GetServiceID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServiceIdOptions model
				getServiceIdOptionsModel := new(iamidentityservicesv1.GetServiceIdOptions)
				getServiceIdOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetServiceID(getServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateServiceID(updateServiceIdOptions *UpdateServiceIdOptions)`, func() {
		updateServiceIDPath := "/v1/serviceids/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateServiceIDPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["If-Match"]).ToNot(BeNil())
				Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}}`)
			}))
			It(`Invoke UpdateServiceID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateServiceIdOptions model
				updateServiceIdOptionsModel := new(iamidentityservicesv1.UpdateServiceIdOptions)
				updateServiceIdOptionsModel.ID = core.StringPtr("testString")
				updateServiceIdOptionsModel.IfMatch = core.StringPtr("testString")
				updateServiceIdOptionsModel.Name = core.StringPtr("testString")
				updateServiceIdOptionsModel.Description = core.StringPtr("testString")
				updateServiceIdOptionsModel.UniqueInstanceCrns = []string{"testString"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateServiceID(updateServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteServiceID(deleteServiceIdOptions *DeleteServiceIdOptions)`, func() {
		deleteServiceIDPath := "/v1/serviceids/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteServiceIDPath))
				Expect(req.Method).To(Equal("DELETE"))
				res.WriteHeader(204)
			}))
			It(`Invoke DeleteServiceID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteServiceIdOptions model
				deleteServiceIdOptionsModel := new(iamidentityservicesv1.DeleteServiceIdOptions)
				deleteServiceIdOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteServiceID(deleteServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`LockServiceID(lockServiceIdOptions *LockServiceIdOptions)`, func() {
		lockServiceIDPath := "/v1/serviceids/testString/lock"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(lockServiceIDPath))
				Expect(req.Method).To(Equal("POST"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}}`)
			}))
			It(`Invoke LockServiceID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.LockServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LockServiceIdOptions model
				lockServiceIdOptionsModel := new(iamidentityservicesv1.LockServiceIdOptions)
				lockServiceIdOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.LockServiceID(lockServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UnlockServiceID(unlockServiceIdOptions *UnlockServiceIdOptions)`, func() {
		unlockServiceIDPath := "/v1/serviceids/testString/lock"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(unlockServiceIDPath))
				Expect(req.Method).To(Equal("DELETE"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "iam_id": "IamID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "account_id": "AccountID", "name": "Name", "description": "Description", "unique_instance_crns": ["UniqueInstanceCrns"], "apikey": {"context": {"transaction_id": "TransactionID", "operation": "Operation", "user_agent": "UserAgent", "client_ip": "ClientIp", "url": "URL", "instance_id": "InstanceID", "thread_id": "ThreadID", "host": "Host", "start_time": "StartTime", "end_time": "EndTime", "elapsed_time": "ElapsedTime", "cluster_name": "ClusterName"}, "id": "ID", "entity_tag": "EntityTag", "crn": "Crn", "locked": true, "created_at": "2019-01-01T12:00:00", "modified_at": "2019-01-01T12:00:00", "name": "Name", "description": "Description", "iam_id": "IamID", "account_id": "AccountID", "apikey": "Apikey"}}`)
			}))
			It(`Invoke UnlockServiceID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UnlockServiceID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UnlockServiceIdOptions model
				unlockServiceIdOptionsModel := new(iamidentityservicesv1.UnlockServiceIdOptions)
				unlockServiceIdOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UnlockServiceID(unlockServiceIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a sample service client instance`, func() {
			testService, _ := iamidentityservicesv1.NewIamIdentityServicesV1(&iamidentityservicesv1.IamIdentityServicesV1Options{
				URL:           "http://iamidentityservicesv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateApiKeyRequest successfully`, func() {
				name := "testString"
				iamID := "testString"
				model, err := testService.NewCreateApiKeyRequest(name, iamID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockMap() successfully`, func() {
			mockMap := CreateMockMap()
			Expect(mockMap).ToNot(BeNil())
		})
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

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, len(mockData))
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
	d := strfmt.Date(time.Now())
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Now())
	return &d
}
