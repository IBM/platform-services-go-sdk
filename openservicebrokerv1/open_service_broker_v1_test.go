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

package openservicebrokerv1_test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/openservicebrokerv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`OpenServiceBrokerV1`, func() {
	Describe(`GetServiceInstanceState(getServiceInstanceStateOptions *GetServiceInstanceStateOptions)`, func() {
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		getServiceInstanceStatePath := "/bluemix_v1/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getServiceInstanceStatePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"active": true, "enabled": false, "last_active": 10}`)
			}))
			It(`Invoke GetServiceInstanceState successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetServiceInstanceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServiceInstanceStateOptions model
				getServiceInstanceStateOptionsModel := new(openservicebrokerv1.GetServiceInstanceStateOptions)
				getServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetServiceInstanceState(getServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ReplaceState(replaceStateOptions *ReplaceStateOptions)`, func() {
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		replaceStatePath := "/bluemix_v1/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(replaceStatePath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"active": true, "enabled": false, "last_active": 10}`)
			}))
			It(`Invoke ReplaceState successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ReplaceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceStateOptions model
				replaceStateOptionsModel := new(openservicebrokerv1.ReplaceStateOptions)
				replaceStateOptionsModel.InstanceID = core.StringPtr("testString")
				replaceStateOptionsModel.Enabled = core.BoolPtr(false)
				replaceStateOptionsModel.InitiatorID = core.StringPtr("null")
				replaceStateOptionsModel.ReasonCode = core.StringPtr("null")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ReplaceState(replaceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions)`, func() {
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		replaceServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(replaceServiceInstancePath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))

				// TODO: Add check for accepts_incomplete query parameter

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"dashboard_url": "DashboardURL", "operation": "Operation"}`)
			}))
			It(`Invoke ReplaceServiceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ReplaceServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the Parameters model
				parametersModel := new(openservicebrokerv1.Parameters)
				parametersModel.Parameter1 = core.Int64Ptr(int64(38))
				parametersModel.Parameter2 = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(openservicebrokerv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = []openservicebrokerv1.Context{*contextModel}
				replaceServiceInstanceOptionsModel.OrganizationGuid = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.Parameters = []openservicebrokerv1.Parameters{*parametersModel}
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.SpaceGuid = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions)`, func() {
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		updateServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateServiceInstancePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				Expect(req.URL.Query()["accepts_incomplete"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `"This is a string response..."`)
			}))
			It(`Invoke UpdateServiceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the Parameters model
				parametersModel := new(openservicebrokerv1.Parameters)
				parametersModel.Parameter1 = core.Int64Ptr(int64(38))
				parametersModel.Parameter2 = core.StringPtr("null")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.Context = []openservicebrokerv1.Context{*contextModel}
				updateServiceInstanceOptionsModel.Parameters = parametersModel
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions)`, func() {
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteServiceInstancePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))


				// TODO: Add check for accepts_incomplete query parameter

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `"This is a string response..."`)
			}))
			It(`Invoke DeleteServiceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(openservicebrokerv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListCatalog(listCatalogOptions *ListCatalogOptions)`, func() {
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listCatalogPath := "/v2/catalog"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCatalogPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `[{}]`)
			}))
			It(`Invoke ListCatalog successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCatalogOptions model
				listCatalogOptionsModel := new(openservicebrokerv1.ListCatalogOptions)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListCatalog(listCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListLastOperation(listLastOperationOptions *ListLastOperationOptions)`, func() {
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listLastOperationPath := "/v2/service_instances/testString/last_operation"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listLastOperationPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				Expect(req.URL.Query()["operation"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"description": "Description", "state": "State"}`)
			}))
			It(`Invoke ListLastOperation successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListLastOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLastOperationOptions model
				listLastOperationOptionsModel := new(openservicebrokerv1.ListLastOperationOptions)
				listLastOperationOptionsModel.InstanceID = core.StringPtr("testString")
				listLastOperationOptionsModel.Operation = core.StringPtr("testString")
				listLastOperationOptionsModel.PlanID = core.StringPtr("testString")
				listLastOperationOptionsModel.ServiceID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListLastOperation(listLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ReplaceServiceBinding(replaceServiceBindingOptions *ReplaceServiceBindingOptions)`, func() {
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		replaceServiceBindingPath := "/v2/service_instances/testString/service_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(replaceServiceBindingPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `"This is a string response..."`)
			}))
			It(`Invoke ReplaceServiceBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ReplaceServiceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BindResource model
				bindResourceModel := new(openservicebrokerv1.BindResource)
				bindResourceModel.AccountID = core.StringPtr("null")
				bindResourceModel.ServiceidCrn = core.StringPtr("null")
				bindResourceModel.TargetCrn = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceBindingOptions model
				replaceServiceBindingOptionsModel := new(openservicebrokerv1.ReplaceServiceBindingOptions)
				replaceServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.BindResource = []openservicebrokerv1.BindResource{*bindResourceModel}
				replaceServiceBindingOptionsModel.Parameters = CreateMockMap()
				replaceServiceBindingOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceBindingOptionsModel.ServiceID = core.StringPtr("null")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ReplaceServiceBinding(replaceServiceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteServiceBinding(deleteServiceBindingOptions *DeleteServiceBindingOptions)`, func() {
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteServiceBindingPath := "/v2/service_instances/testString/service_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteServiceBindingPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `"This is a string response..."`)
			}))
			It(`Invoke DeleteServiceBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteServiceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteServiceBindingOptions model
				deleteServiceBindingOptionsModel := new(openservicebrokerv1.DeleteServiceBindingOptions)
				deleteServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.ServiceID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteServiceBinding(deleteServiceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
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
