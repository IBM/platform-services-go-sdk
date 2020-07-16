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

package casemanagementv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/casemanagementv1"
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

var _ = Describe(`CaseManagementV1`, func() {
	var testServer *httptest.Server
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
				URL: "https://casemanagementv1/api",
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
				"CASE_MANAGEMENT_URL": "https://casemanagementv1/api",
				"CASE_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1UsingExternalConfig(&casemanagementv1.CaseManagementV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1UsingExternalConfig(&casemanagementv1.CaseManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1UsingExternalConfig(&casemanagementv1.CaseManagementV1Options{
				})
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
				"CASE_MANAGEMENT_URL": "https://casemanagementv1/api",
				"CASE_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := casemanagementv1.NewCaseManagementV1UsingExternalConfig(&casemanagementv1.CaseManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CASE_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := casemanagementv1.NewCaseManagementV1UsingExternalConfig(&casemanagementv1.CaseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetCases(getCasesOptions *GetCasesOptions) - Operation response error`, func() {
		getCasesPath := "/cases"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCasesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"number"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCases with error: Operation response processing error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCasesOptions model
				getCasesOptionsModel := new(casemanagementv1.GetCasesOptions)
				getCasesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Search = core.StringPtr("testString")
				getCasesOptionsModel.Sort = core.StringPtr("number")
				getCasesOptionsModel.Status = []string{"new"}
				getCasesOptionsModel.Fields = []string{"number"}
				getCasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetCases(getCasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCases(getCasesOptions *GetCasesOptions)`, func() {
		getCasesPath := "/cases"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCasesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"number"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"total_count": 10, "first": {"href": "Href"}, "next": {"href": "Href"}, "previous": {"href": "Href"}, "last": {"href": "Href"}, "cases": [{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "Crn", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}]}`)
				}))
			})
			It(`Invoke GetCases successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCases(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCasesOptions model
				getCasesOptionsModel := new(casemanagementv1.GetCasesOptions)
				getCasesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Search = core.StringPtr("testString")
				getCasesOptionsModel.Sort = core.StringPtr("number")
				getCasesOptionsModel.Status = []string{"new"}
				getCasesOptionsModel.Fields = []string{"number"}
 				getCasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCases(getCasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetCases with error: Operation request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCasesOptions model
				getCasesOptionsModel := new(casemanagementv1.GetCasesOptions)
				getCasesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Search = core.StringPtr("testString")
				getCasesOptionsModel.Sort = core.StringPtr("number")
				getCasesOptionsModel.Status = []string{"new"}
				getCasesOptionsModel.Fields = []string{"number"}
				getCasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetCases(getCasesOptionsModel)
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
	Describe(`CreateCase(createCaseOptions *CreateCaseOptions) - Operation response error`, func() {
		createCasePath := "/cases"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createCasePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCase with error: Operation response processing error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the OfferingType model
				offeringTypeModel := new(casemanagementv1.OfferingType)
				offeringTypeModel.Group = core.StringPtr("crn_service_name")
				offeringTypeModel.Key = core.StringPtr("testString")
				offeringTypeModel.Kind = core.StringPtr("testString")
				offeringTypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the CasePayloadEu model
				casePayloadEuModel := new(casemanagementv1.CasePayloadEu)
				casePayloadEuModel.Supported = core.BoolPtr(true)
				casePayloadEuModel.DataCenter = core.Int64Ptr(int64(38))

				// Construct an instance of the Offering model
				offeringModel := new(casemanagementv1.Offering)
				offeringModel.Name = core.StringPtr("testString")
				offeringModel.Type = offeringTypeModel

				// Construct an instance of the ResourcePayload model
				resourcePayloadModel := new(casemanagementv1.ResourcePayload)
				resourcePayloadModel.Crn = core.StringPtr("testString")
				resourcePayloadModel.Type = core.StringPtr("testString")
				resourcePayloadModel.ID = core.Float64Ptr(float64(72.5))
				resourcePayloadModel.Note = core.StringPtr("testString")

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsModel := new(casemanagementv1.CreateCaseOptions)
				createCaseOptionsModel.Type = core.StringPtr("technical")
				createCaseOptionsModel.Subject = core.StringPtr("testString")
				createCaseOptionsModel.Description = core.StringPtr("testString")
				createCaseOptionsModel.Severity = core.Int64Ptr(int64(1))
				createCaseOptionsModel.Eu = casePayloadEuModel
				createCaseOptionsModel.Offering = offeringModel
				createCaseOptionsModel.Resources = []casemanagementv1.ResourcePayload{*resourcePayloadModel}
				createCaseOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				createCaseOptionsModel.InvoiceNumber = core.StringPtr("testString")
				createCaseOptionsModel.SlaCreditRequest = core.BoolPtr(true)
				createCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateCase(createCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateCase(createCaseOptions *CreateCaseOptions)`, func() {
		createCasePath := "/cases"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createCasePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "Crn", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
				}))
			})
			It(`Invoke CreateCase successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateCase(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the OfferingType model
				offeringTypeModel := new(casemanagementv1.OfferingType)
				offeringTypeModel.Group = core.StringPtr("crn_service_name")
				offeringTypeModel.Key = core.StringPtr("testString")
				offeringTypeModel.Kind = core.StringPtr("testString")
				offeringTypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the CasePayloadEu model
				casePayloadEuModel := new(casemanagementv1.CasePayloadEu)
				casePayloadEuModel.Supported = core.BoolPtr(true)
				casePayloadEuModel.DataCenter = core.Int64Ptr(int64(38))

				// Construct an instance of the Offering model
				offeringModel := new(casemanagementv1.Offering)
				offeringModel.Name = core.StringPtr("testString")
				offeringModel.Type = offeringTypeModel

				// Construct an instance of the ResourcePayload model
				resourcePayloadModel := new(casemanagementv1.ResourcePayload)
				resourcePayloadModel.Crn = core.StringPtr("testString")
				resourcePayloadModel.Type = core.StringPtr("testString")
				resourcePayloadModel.ID = core.Float64Ptr(float64(72.5))
				resourcePayloadModel.Note = core.StringPtr("testString")

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsModel := new(casemanagementv1.CreateCaseOptions)
				createCaseOptionsModel.Type = core.StringPtr("technical")
				createCaseOptionsModel.Subject = core.StringPtr("testString")
				createCaseOptionsModel.Description = core.StringPtr("testString")
				createCaseOptionsModel.Severity = core.Int64Ptr(int64(1))
				createCaseOptionsModel.Eu = casePayloadEuModel
				createCaseOptionsModel.Offering = offeringModel
				createCaseOptionsModel.Resources = []casemanagementv1.ResourcePayload{*resourcePayloadModel}
				createCaseOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				createCaseOptionsModel.InvoiceNumber = core.StringPtr("testString")
				createCaseOptionsModel.SlaCreditRequest = core.BoolPtr(true)
 				createCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateCase(createCaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateCase with error: Operation validation and request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the OfferingType model
				offeringTypeModel := new(casemanagementv1.OfferingType)
				offeringTypeModel.Group = core.StringPtr("crn_service_name")
				offeringTypeModel.Key = core.StringPtr("testString")
				offeringTypeModel.Kind = core.StringPtr("testString")
				offeringTypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the CasePayloadEu model
				casePayloadEuModel := new(casemanagementv1.CasePayloadEu)
				casePayloadEuModel.Supported = core.BoolPtr(true)
				casePayloadEuModel.DataCenter = core.Int64Ptr(int64(38))

				// Construct an instance of the Offering model
				offeringModel := new(casemanagementv1.Offering)
				offeringModel.Name = core.StringPtr("testString")
				offeringModel.Type = offeringTypeModel

				// Construct an instance of the ResourcePayload model
				resourcePayloadModel := new(casemanagementv1.ResourcePayload)
				resourcePayloadModel.Crn = core.StringPtr("testString")
				resourcePayloadModel.Type = core.StringPtr("testString")
				resourcePayloadModel.ID = core.Float64Ptr(float64(72.5))
				resourcePayloadModel.Note = core.StringPtr("testString")

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsModel := new(casemanagementv1.CreateCaseOptions)
				createCaseOptionsModel.Type = core.StringPtr("technical")
				createCaseOptionsModel.Subject = core.StringPtr("testString")
				createCaseOptionsModel.Description = core.StringPtr("testString")
				createCaseOptionsModel.Severity = core.Int64Ptr(int64(1))
				createCaseOptionsModel.Eu = casePayloadEuModel
				createCaseOptionsModel.Offering = offeringModel
				createCaseOptionsModel.Resources = []casemanagementv1.ResourcePayload{*resourcePayloadModel}
				createCaseOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				createCaseOptionsModel.InvoiceNumber = core.StringPtr("testString")
				createCaseOptionsModel.SlaCreditRequest = core.BoolPtr(true)
				createCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateCase(createCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCaseOptions model with no property values
				createCaseOptionsModelNew := new(casemanagementv1.CreateCaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateCase(createCaseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCase(getCaseOptions *GetCaseOptions) - Operation response error`, func() {
		getCasePath := "/cases/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCasePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCase with error: Operation response processing error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCaseOptions model
				getCaseOptionsModel := new(casemanagementv1.GetCaseOptions)
				getCaseOptionsModel.CaseNumber = core.StringPtr("testString")
				getCaseOptionsModel.Fields = []string{"number"}
				getCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetCase(getCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCase(getCaseOptions *GetCaseOptions)`, func() {
		getCasePath := "/cases/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCasePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "Crn", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
				}))
			})
			It(`Invoke GetCase successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCase(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCaseOptions model
				getCaseOptionsModel := new(casemanagementv1.GetCaseOptions)
				getCaseOptionsModel.CaseNumber = core.StringPtr("testString")
				getCaseOptionsModel.Fields = []string{"number"}
 				getCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCase(getCaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetCase with error: Operation validation and request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCaseOptions model
				getCaseOptionsModel := new(casemanagementv1.GetCaseOptions)
				getCaseOptionsModel.CaseNumber = core.StringPtr("testString")
				getCaseOptionsModel.Fields = []string{"number"}
				getCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetCase(getCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCaseOptions model with no property values
				getCaseOptionsModelNew := new(casemanagementv1.GetCaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetCase(getCaseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCaseStatus(updateCaseStatusOptions *UpdateCaseStatusOptions) - Operation response error`, func() {
		updateCaseStatusPath := "/cases/testString/status"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateCaseStatusPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCaseStatus with error: Operation response processing error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResolvePayload model
				statusPayloadModel := new(casemanagementv1.ResolvePayload)
				statusPayloadModel.Action = core.StringPtr("resolve")
				statusPayloadModel.Comment = core.StringPtr("testString")
				statusPayloadModel.ResolutionCode = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdateCaseStatusOptions model
				updateCaseStatusOptionsModel := new(casemanagementv1.UpdateCaseStatusOptions)
				updateCaseStatusOptionsModel.CaseNumber = core.StringPtr("testString")
				updateCaseStatusOptionsModel.StatusPayload = statusPayloadModel
				updateCaseStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateCaseStatus(updateCaseStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCaseStatus(updateCaseStatusOptions *UpdateCaseStatusOptions)`, func() {
		updateCaseStatusPath := "/cases/testString/status"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateCaseStatusPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "Crn", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
				}))
			})
			It(`Invoke UpdateCaseStatus successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateCaseStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResolvePayload model
				statusPayloadModel := new(casemanagementv1.ResolvePayload)
				statusPayloadModel.Action = core.StringPtr("resolve")
				statusPayloadModel.Comment = core.StringPtr("testString")
				statusPayloadModel.ResolutionCode = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdateCaseStatusOptions model
				updateCaseStatusOptionsModel := new(casemanagementv1.UpdateCaseStatusOptions)
				updateCaseStatusOptionsModel.CaseNumber = core.StringPtr("testString")
				updateCaseStatusOptionsModel.StatusPayload = statusPayloadModel
 				updateCaseStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateCaseStatus(updateCaseStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateCaseStatus with error: Operation validation and request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResolvePayload model
				statusPayloadModel := new(casemanagementv1.ResolvePayload)
				statusPayloadModel.Action = core.StringPtr("resolve")
				statusPayloadModel.Comment = core.StringPtr("testString")
				statusPayloadModel.ResolutionCode = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdateCaseStatusOptions model
				updateCaseStatusOptionsModel := new(casemanagementv1.UpdateCaseStatusOptions)
				updateCaseStatusOptionsModel.CaseNumber = core.StringPtr("testString")
				updateCaseStatusOptionsModel.StatusPayload = statusPayloadModel
				updateCaseStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateCaseStatus(updateCaseStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCaseStatusOptions model with no property values
				updateCaseStatusOptionsModelNew := new(casemanagementv1.UpdateCaseStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateCaseStatus(updateCaseStatusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddComment(addCommentOptions *AddCommentOptions) - Operation response error`, func() {
		addCommentPath := "/cases/testString/comments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(addCommentPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddComment with error: Operation response processing error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the AddCommentOptions model
				addCommentOptionsModel := new(casemanagementv1.AddCommentOptions)
				addCommentOptionsModel.CaseNumber = core.StringPtr("testString")
				addCommentOptionsModel.Comment = core.StringPtr("This is a test comment")
				addCommentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.AddComment(addCommentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddComment(addCommentOptions *AddCommentOptions)`, func() {
		addCommentPath := "/cases/testString/comments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(addCommentPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}`)
				}))
			})
			It(`Invoke AddComment successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.AddComment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddCommentOptions model
				addCommentOptionsModel := new(casemanagementv1.AddCommentOptions)
				addCommentOptionsModel.CaseNumber = core.StringPtr("testString")
				addCommentOptionsModel.Comment = core.StringPtr("This is a test comment")
 				addCommentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AddComment(addCommentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke AddComment with error: Operation validation and request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the AddCommentOptions model
				addCommentOptionsModel := new(casemanagementv1.AddCommentOptions)
				addCommentOptionsModel.CaseNumber = core.StringPtr("testString")
				addCommentOptionsModel.Comment = core.StringPtr("This is a test comment")
				addCommentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.AddComment(addCommentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddCommentOptions model with no property values
				addCommentOptionsModelNew := new(casemanagementv1.AddCommentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.AddComment(addCommentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddWatchlist(addWatchlistOptions *AddWatchlistOptions) - Operation response error`, func() {
		addWatchlistPath := "/cases/testString/watchlist"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(addWatchlistPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddWatchlist with error: Operation response processing error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the AddWatchlistOptions model
				addWatchlistOptionsModel := new(casemanagementv1.AddWatchlistOptions)
				addWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				addWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				addWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.AddWatchlist(addWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddWatchlist(addWatchlistOptions *AddWatchlistOptions)`, func() {
		addWatchlistPath := "/cases/testString/watchlist"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(addWatchlistPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"added": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "failed": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}]}`)
				}))
			})
			It(`Invoke AddWatchlist successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.AddWatchlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the AddWatchlistOptions model
				addWatchlistOptionsModel := new(casemanagementv1.AddWatchlistOptions)
				addWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				addWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
 				addWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AddWatchlist(addWatchlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke AddWatchlist with error: Operation validation and request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the AddWatchlistOptions model
				addWatchlistOptionsModel := new(casemanagementv1.AddWatchlistOptions)
				addWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				addWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				addWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.AddWatchlist(addWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddWatchlistOptions model with no property values
				addWatchlistOptionsModelNew := new(casemanagementv1.AddWatchlistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.AddWatchlist(addWatchlistOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveWatchlist(removeWatchlistOptions *RemoveWatchlistOptions) - Operation response error`, func() {
		removeWatchlistPath := "/cases/testString/watchlist"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(removeWatchlistPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RemoveWatchlist with error: Operation response processing error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the RemoveWatchlistOptions model
				removeWatchlistOptionsModel := new(casemanagementv1.RemoveWatchlistOptions)
				removeWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				removeWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				removeWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.RemoveWatchlist(removeWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RemoveWatchlist(removeWatchlistOptions *RemoveWatchlistOptions)`, func() {
		removeWatchlistPath := "/cases/testString/watchlist"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(removeWatchlistPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}]}`)
				}))
			})
			It(`Invoke RemoveWatchlist successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.RemoveWatchlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the RemoveWatchlistOptions model
				removeWatchlistOptionsModel := new(casemanagementv1.RemoveWatchlistOptions)
				removeWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				removeWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
 				removeWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.RemoveWatchlist(removeWatchlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke RemoveWatchlist with error: Operation validation and request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the RemoveWatchlistOptions model
				removeWatchlistOptionsModel := new(casemanagementv1.RemoveWatchlistOptions)
				removeWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				removeWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				removeWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.RemoveWatchlist(removeWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveWatchlistOptions model with no property values
				removeWatchlistOptionsModelNew := new(casemanagementv1.RemoveWatchlistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.RemoveWatchlist(removeWatchlistOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddResource(addResourceOptions *AddResourceOptions) - Operation response error`, func() {
		addResourcePath := "/cases/testString/resources"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(addResourcePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddResource with error: Operation response processing error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the AddResourceOptions model
				addResourceOptionsModel := new(casemanagementv1.AddResourceOptions)
				addResourceOptionsModel.CaseNumber = core.StringPtr("testString")
				addResourceOptionsModel.Crn = core.StringPtr("testString")
				addResourceOptionsModel.Type = core.StringPtr("testString")
				addResourceOptionsModel.ID = core.Float64Ptr(float64(72.5))
				addResourceOptionsModel.Note = core.StringPtr("testString")
				addResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.AddResource(addResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddResource(addResourceOptions *AddResourceOptions)`, func() {
		addResourcePath := "/cases/testString/resources"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(addResourcePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"crn": "Crn", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}`)
				}))
			})
			It(`Invoke AddResource successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.AddResource(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddResourceOptions model
				addResourceOptionsModel := new(casemanagementv1.AddResourceOptions)
				addResourceOptionsModel.CaseNumber = core.StringPtr("testString")
				addResourceOptionsModel.Crn = core.StringPtr("testString")
				addResourceOptionsModel.Type = core.StringPtr("testString")
				addResourceOptionsModel.ID = core.Float64Ptr(float64(72.5))
				addResourceOptionsModel.Note = core.StringPtr("testString")
 				addResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AddResource(addResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke AddResource with error: Operation validation and request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the AddResourceOptions model
				addResourceOptionsModel := new(casemanagementv1.AddResourceOptions)
				addResourceOptionsModel.CaseNumber = core.StringPtr("testString")
				addResourceOptionsModel.Crn = core.StringPtr("testString")
				addResourceOptionsModel.Type = core.StringPtr("testString")
				addResourceOptionsModel.ID = core.Float64Ptr(float64(72.5))
				addResourceOptionsModel.Note = core.StringPtr("testString")
				addResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.AddResource(addResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddResourceOptions model with no property values
				addResourceOptionsModelNew := new(casemanagementv1.AddResourceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.AddResource(addResourceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UploadFile(uploadFileOptions *UploadFileOptions) - Operation response error`, func() {
		uploadFilePath := "/cases/testString/attachments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(uploadFilePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UploadFile with error: Operation response processing error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the FileWithMetadata model
				fileWithMetadataModel := new(casemanagementv1.FileWithMetadata)
				fileWithMetadataModel.Data = CreateMockReader("This is a mock file.")
				fileWithMetadataModel.Filename = core.StringPtr("testString")
				fileWithMetadataModel.ContentType = core.StringPtr("testString")

				// Construct an instance of the UploadFileOptions model
				uploadFileOptionsModel := new(casemanagementv1.UploadFileOptions)
				uploadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				uploadFileOptionsModel.File = []casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}
				uploadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UploadFile(uploadFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UploadFile(uploadFileOptions *UploadFileOptions)`, func() {
		uploadFilePath := "/cases/testString/attachments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(uploadFilePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}`)
				}))
			})
			It(`Invoke UploadFile successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UploadFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileWithMetadata model
				fileWithMetadataModel := new(casemanagementv1.FileWithMetadata)
				fileWithMetadataModel.Data = CreateMockReader("This is a mock file.")
				fileWithMetadataModel.Filename = core.StringPtr("testString")
				fileWithMetadataModel.ContentType = core.StringPtr("testString")

				// Construct an instance of the UploadFileOptions model
				uploadFileOptionsModel := new(casemanagementv1.UploadFileOptions)
				uploadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				uploadFileOptionsModel.File = []casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}
 				uploadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UploadFile(uploadFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UploadFile with error: Operation validation and request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the FileWithMetadata model
				fileWithMetadataModel := new(casemanagementv1.FileWithMetadata)
				fileWithMetadataModel.Data = CreateMockReader("This is a mock file.")
				fileWithMetadataModel.Filename = core.StringPtr("testString")
				fileWithMetadataModel.ContentType = core.StringPtr("testString")

				// Construct an instance of the UploadFileOptions model
				uploadFileOptionsModel := new(casemanagementv1.UploadFileOptions)
				uploadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				uploadFileOptionsModel.File = []casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}
				uploadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UploadFile(uploadFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UploadFileOptions model with no property values
				uploadFileOptionsModelNew := new(casemanagementv1.UploadFileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UploadFile(uploadFileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DownloadFile(downloadFileOptions *DownloadFileOptions)`, func() {
		downloadFilePath := "/cases/testString/attachments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(downloadFilePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/octet-stream")
					res.WriteHeader(200)
					fmt.Fprintf(res, `This is a mock binary response.`)
				}))
			})
			It(`Invoke DownloadFile successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DownloadFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DownloadFileOptions model
				downloadFileOptionsModel := new(casemanagementv1.DownloadFileOptions)
				downloadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				downloadFileOptionsModel.FileID = core.StringPtr("testString")
 				downloadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DownloadFile(downloadFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DownloadFile with error: Operation validation and request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DownloadFileOptions model
				downloadFileOptionsModel := new(casemanagementv1.DownloadFileOptions)
				downloadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				downloadFileOptionsModel.FileID = core.StringPtr("testString")
				downloadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DownloadFile(downloadFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DownloadFileOptions model with no property values
				downloadFileOptionsModelNew := new(casemanagementv1.DownloadFileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DownloadFile(downloadFileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteFile(deleteFileOptions *DeleteFileOptions) - Operation response error`, func() {
		deleteFilePath := "/cases/testString/attachments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteFilePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteFile with error: Operation response processing error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteFileOptions model
				deleteFileOptionsModel := new(casemanagementv1.DeleteFileOptions)
				deleteFileOptionsModel.CaseNumber = core.StringPtr("testString")
				deleteFileOptionsModel.FileID = core.StringPtr("testString")
				deleteFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteFile(deleteFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteFile(deleteFileOptions *DeleteFileOptions)`, func() {
		deleteFilePath := "/cases/testString/attachments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteFilePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}]}`)
				}))
			})
			It(`Invoke DeleteFile successfully`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteFileOptions model
				deleteFileOptionsModel := new(casemanagementv1.DeleteFileOptions)
				deleteFileOptionsModel.CaseNumber = core.StringPtr("testString")
				deleteFileOptionsModel.FileID = core.StringPtr("testString")
 				deleteFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteFile(deleteFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteFile with error: Operation validation and request error`, func() {
				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteFileOptions model
				deleteFileOptionsModel := new(casemanagementv1.DeleteFileOptions)
				deleteFileOptionsModel.CaseNumber = core.StringPtr("testString")
				deleteFileOptionsModel.FileID = core.StringPtr("testString")
				deleteFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteFile(deleteFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteFileOptions model with no property values
				deleteFileOptionsModelNew := new(casemanagementv1.DeleteFileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteFile(deleteFileOptionsModelNew)
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
			testService, _ := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
				URL:           "http://casemanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddCommentOptions successfully`, func() {
				// Construct an instance of the AddCommentOptions model
				caseNumber := "testString"
				addCommentOptionsComment := "This is a test comment"
				addCommentOptionsModel := testService.NewAddCommentOptions(caseNumber, addCommentOptionsComment)
				addCommentOptionsModel.SetCaseNumber("testString")
				addCommentOptionsModel.SetComment("This is a test comment")
				addCommentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addCommentOptionsModel).ToNot(BeNil())
				Expect(addCommentOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(addCommentOptionsModel.Comment).To(Equal(core.StringPtr("This is a test comment")))
				Expect(addCommentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddResourceOptions successfully`, func() {
				// Construct an instance of the AddResourceOptions model
				caseNumber := "testString"
				addResourceOptionsModel := testService.NewAddResourceOptions(caseNumber)
				addResourceOptionsModel.SetCaseNumber("testString")
				addResourceOptionsModel.SetCrn("testString")
				addResourceOptionsModel.SetType("testString")
				addResourceOptionsModel.SetID(float64(72.5))
				addResourceOptionsModel.SetNote("testString")
				addResourceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addResourceOptionsModel).ToNot(BeNil())
				Expect(addResourceOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(addResourceOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(addResourceOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(addResourceOptionsModel.ID).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(addResourceOptionsModel.Note).To(Equal(core.StringPtr("testString")))
				Expect(addResourceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddWatchlistOptions successfully`, func() {
				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				Expect(userModel).ToNot(BeNil())
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")
				Expect(userModel.Realm).To(Equal(core.StringPtr("IBMid")))
				Expect(userModel.UserID).To(Equal(core.StringPtr("abc@ibm.com")))

				// Construct an instance of the AddWatchlistOptions model
				caseNumber := "testString"
				addWatchlistOptionsModel := testService.NewAddWatchlistOptions(caseNumber)
				addWatchlistOptionsModel.SetCaseNumber("testString")
				addWatchlistOptionsModel.SetWatchlist([]casemanagementv1.User{*userModel})
				addWatchlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addWatchlistOptionsModel).ToNot(BeNil())
				Expect(addWatchlistOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(addWatchlistOptionsModel.Watchlist).To(Equal([]casemanagementv1.User{*userModel}))
				Expect(addWatchlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCaseOptions successfully`, func() {
				// Construct an instance of the OfferingType model
				offeringTypeModel := new(casemanagementv1.OfferingType)
				Expect(offeringTypeModel).ToNot(BeNil())
				offeringTypeModel.Group = core.StringPtr("crn_service_name")
				offeringTypeModel.Key = core.StringPtr("testString")
				offeringTypeModel.Kind = core.StringPtr("testString")
				offeringTypeModel.ID = core.StringPtr("testString")
				Expect(offeringTypeModel.Group).To(Equal(core.StringPtr("crn_service_name")))
				Expect(offeringTypeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(offeringTypeModel.Kind).To(Equal(core.StringPtr("testString")))
				Expect(offeringTypeModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CasePayloadEu model
				casePayloadEuModel := new(casemanagementv1.CasePayloadEu)
				Expect(casePayloadEuModel).ToNot(BeNil())
				casePayloadEuModel.Supported = core.BoolPtr(true)
				casePayloadEuModel.DataCenter = core.Int64Ptr(int64(38))
				Expect(casePayloadEuModel.Supported).To(Equal(core.BoolPtr(true)))
				Expect(casePayloadEuModel.DataCenter).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the Offering model
				offeringModel := new(casemanagementv1.Offering)
				Expect(offeringModel).ToNot(BeNil())
				offeringModel.Name = core.StringPtr("testString")
				offeringModel.Type = offeringTypeModel
				Expect(offeringModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(offeringModel.Type).To(Equal(offeringTypeModel))

				// Construct an instance of the ResourcePayload model
				resourcePayloadModel := new(casemanagementv1.ResourcePayload)
				Expect(resourcePayloadModel).ToNot(BeNil())
				resourcePayloadModel.Crn = core.StringPtr("testString")
				resourcePayloadModel.Type = core.StringPtr("testString")
				resourcePayloadModel.ID = core.Float64Ptr(float64(72.5))
				resourcePayloadModel.Note = core.StringPtr("testString")
				Expect(resourcePayloadModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(resourcePayloadModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(resourcePayloadModel.ID).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(resourcePayloadModel.Note).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				Expect(userModel).ToNot(BeNil())
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")
				Expect(userModel.Realm).To(Equal(core.StringPtr("IBMid")))
				Expect(userModel.UserID).To(Equal(core.StringPtr("abc@ibm.com")))

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsType := "technical"
				createCaseOptionsSubject := "testString"
				createCaseOptionsDescription := "testString"
				createCaseOptionsModel := testService.NewCreateCaseOptions(createCaseOptionsType, createCaseOptionsSubject, createCaseOptionsDescription)
				createCaseOptionsModel.SetType("technical")
				createCaseOptionsModel.SetSubject("testString")
				createCaseOptionsModel.SetDescription("testString")
				createCaseOptionsModel.SetSeverity(int64(1))
				createCaseOptionsModel.SetEu(casePayloadEuModel)
				createCaseOptionsModel.SetOffering(offeringModel)
				createCaseOptionsModel.SetResources([]casemanagementv1.ResourcePayload{*resourcePayloadModel})
				createCaseOptionsModel.SetWatchlist([]casemanagementv1.User{*userModel})
				createCaseOptionsModel.SetInvoiceNumber("testString")
				createCaseOptionsModel.SetSlaCreditRequest(true)
				createCaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCaseOptionsModel).ToNot(BeNil())
				Expect(createCaseOptionsModel.Type).To(Equal(core.StringPtr("technical")))
				Expect(createCaseOptionsModel.Subject).To(Equal(core.StringPtr("testString")))
				Expect(createCaseOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createCaseOptionsModel.Severity).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createCaseOptionsModel.Eu).To(Equal(casePayloadEuModel))
				Expect(createCaseOptionsModel.Offering).To(Equal(offeringModel))
				Expect(createCaseOptionsModel.Resources).To(Equal([]casemanagementv1.ResourcePayload{*resourcePayloadModel}))
				Expect(createCaseOptionsModel.Watchlist).To(Equal([]casemanagementv1.User{*userModel}))
				Expect(createCaseOptionsModel.InvoiceNumber).To(Equal(core.StringPtr("testString")))
				Expect(createCaseOptionsModel.SlaCreditRequest).To(Equal(core.BoolPtr(true)))
				Expect(createCaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFileOptions successfully`, func() {
				// Construct an instance of the DeleteFileOptions model
				caseNumber := "testString"
				fileID := "testString"
				deleteFileOptionsModel := testService.NewDeleteFileOptions(caseNumber, fileID)
				deleteFileOptionsModel.SetCaseNumber("testString")
				deleteFileOptionsModel.SetFileID("testString")
				deleteFileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFileOptionsModel).ToNot(BeNil())
				Expect(deleteFileOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(deleteFileOptionsModel.FileID).To(Equal(core.StringPtr("testString")))
				Expect(deleteFileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDownloadFileOptions successfully`, func() {
				// Construct an instance of the DownloadFileOptions model
				caseNumber := "testString"
				fileID := "testString"
				downloadFileOptionsModel := testService.NewDownloadFileOptions(caseNumber, fileID)
				downloadFileOptionsModel.SetCaseNumber("testString")
				downloadFileOptionsModel.SetFileID("testString")
				downloadFileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(downloadFileOptionsModel).ToNot(BeNil())
				Expect(downloadFileOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(downloadFileOptionsModel.FileID).To(Equal(core.StringPtr("testString")))
				Expect(downloadFileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFileWithMetadata successfully`, func() {
				data := CreateMockReader("This is a mock file.")
				model, err := testService.NewFileWithMetadata(data)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetCaseOptions successfully`, func() {
				// Construct an instance of the GetCaseOptions model
				caseNumber := "testString"
				getCaseOptionsModel := testService.NewGetCaseOptions(caseNumber)
				getCaseOptionsModel.SetCaseNumber("testString")
				getCaseOptionsModel.SetFields([]string{"number"})
				getCaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCaseOptionsModel).ToNot(BeNil())
				Expect(getCaseOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(getCaseOptionsModel.Fields).To(Equal([]string{"number"}))
				Expect(getCaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCasesOptions successfully`, func() {
				// Construct an instance of the GetCasesOptions model
				getCasesOptionsModel := testService.NewGetCasesOptions()
				getCasesOptionsModel.SetOffset(int64(38))
				getCasesOptionsModel.SetLimit(int64(38))
				getCasesOptionsModel.SetSearch("testString")
				getCasesOptionsModel.SetSort("number")
				getCasesOptionsModel.SetStatus([]string{"new"})
				getCasesOptionsModel.SetFields([]string{"number"})
				getCasesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCasesOptionsModel).ToNot(BeNil())
				Expect(getCasesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getCasesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getCasesOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(getCasesOptionsModel.Sort).To(Equal(core.StringPtr("number")))
				Expect(getCasesOptionsModel.Status).To(Equal([]string{"new"}))
				Expect(getCasesOptionsModel.Fields).To(Equal([]string{"number"}))
				Expect(getCasesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewOffering successfully`, func() {
				name := "testString"
				var typeVar *casemanagementv1.OfferingType = nil
				_, err := testService.NewOffering(name, typeVar)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewOfferingType successfully`, func() {
				group := "crn_service_name"
				key := "testString"
				model, err := testService.NewOfferingType(group, key)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRemoveWatchlistOptions successfully`, func() {
				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				Expect(userModel).ToNot(BeNil())
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")
				Expect(userModel.Realm).To(Equal(core.StringPtr("IBMid")))
				Expect(userModel.UserID).To(Equal(core.StringPtr("abc@ibm.com")))

				// Construct an instance of the RemoveWatchlistOptions model
				caseNumber := "testString"
				removeWatchlistOptionsModel := testService.NewRemoveWatchlistOptions(caseNumber)
				removeWatchlistOptionsModel.SetCaseNumber("testString")
				removeWatchlistOptionsModel.SetWatchlist([]casemanagementv1.User{*userModel})
				removeWatchlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeWatchlistOptionsModel).ToNot(BeNil())
				Expect(removeWatchlistOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(removeWatchlistOptionsModel.Watchlist).To(Equal([]casemanagementv1.User{*userModel}))
				Expect(removeWatchlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCaseStatusOptions successfully`, func() {
				// Construct an instance of the ResolvePayload model
				statusPayloadModel := new(casemanagementv1.ResolvePayload)
				Expect(statusPayloadModel).ToNot(BeNil())
				statusPayloadModel.Action = core.StringPtr("resolve")
				statusPayloadModel.Comment = core.StringPtr("testString")
				statusPayloadModel.ResolutionCode = core.Int64Ptr(int64(1))
				Expect(statusPayloadModel.Action).To(Equal(core.StringPtr("resolve")))
				Expect(statusPayloadModel.Comment).To(Equal(core.StringPtr("testString")))
				Expect(statusPayloadModel.ResolutionCode).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the UpdateCaseStatusOptions model
				caseNumber := "testString"
				var statusPayload casemanagementv1.StatusPayloadIntf = nil
				updateCaseStatusOptionsModel := testService.NewUpdateCaseStatusOptions(caseNumber, statusPayload)
				updateCaseStatusOptionsModel.SetCaseNumber("testString")
				updateCaseStatusOptionsModel.SetStatusPayload(statusPayloadModel)
				updateCaseStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCaseStatusOptionsModel).ToNot(BeNil())
				Expect(updateCaseStatusOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(updateCaseStatusOptionsModel.StatusPayload).To(Equal(statusPayloadModel))
				Expect(updateCaseStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadFileOptions successfully`, func() {
				// Construct an instance of the FileWithMetadata model
				fileWithMetadataModel := new(casemanagementv1.FileWithMetadata)
				Expect(fileWithMetadataModel).ToNot(BeNil())
				fileWithMetadataModel.Data = CreateMockReader("This is a mock file.")
				fileWithMetadataModel.Filename = core.StringPtr("testString")
				fileWithMetadataModel.ContentType = core.StringPtr("testString")
				Expect(fileWithMetadataModel.Data).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(fileWithMetadataModel.Filename).To(Equal(core.StringPtr("testString")))
				Expect(fileWithMetadataModel.ContentType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UploadFileOptions model
				caseNumber := "testString"
				file := []casemanagementv1.FileWithMetadata{}
				uploadFileOptionsModel := testService.NewUploadFileOptions(caseNumber, file)
				uploadFileOptionsModel.SetCaseNumber("testString")
				uploadFileOptionsModel.SetFile([]casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }})
				uploadFileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadFileOptionsModel).ToNot(BeNil())
				Expect(uploadFileOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(uploadFileOptionsModel.File).To(Equal([]casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}))
				Expect(uploadFileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUser successfully`, func() {
				realm := "IBMid"
				userID := "abc@ibm.com"
				model, err := testService.NewUser(realm, userID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAcceptPayload successfully`, func() {
				action := "accept"
				model, err := testService.NewAcceptPayload(action)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResolvePayload successfully`, func() {
				action := "resolve"
				resolutionCode := int64(1)
				model, err := testService.NewResolvePayload(action, resolutionCode)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUnresolvePayload successfully`, func() {
				action := "unresolve"
				comment := "testString"
				model, err := testService.NewUnresolvePayload(action, comment)
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
