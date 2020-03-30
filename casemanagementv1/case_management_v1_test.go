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
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/casemanagementv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`CaseManagementV1`, func() {
	Describe(`GetCases(getCasesOptions *GetCasesOptions)`, func() {
		bearerToken := "0ui9876453"
		getCasesPath := "/case-management/v1/cases"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCasesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["sort"]).To(Equal([]string{"number"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"total_count": 10, "first": {"href": "Href"}, "next": {"href": "Href"}, "previous": {"href": "Href"}, "last": {"href": "Href"}, "cases": [{"number": "Number", "short_escription": "ShortEscription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"id": "ID", "value": "Value"}, "resources": [{"crn": "Crn", "name": "Name", "type": "Type", "id": 2, "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}]}`)
			}))
			It(`Invoke GetCases successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCases(getCasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateCase(createCaseOptions *CreateCaseOptions)`, func() {
		bearerToken := "0ui9876453"
		createCasePath := "/case-management/v1/cases"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createCasePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"number": "Number", "short_escription": "ShortEscription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"id": "ID", "value": "Value"}, "resources": [{"crn": "Crn", "name": "Name", "type": "Type", "id": 2, "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
			}))
			It(`Invoke CreateCase successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateCase(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the OfferingPayloadType model
				offeringPayloadTypeModel := new(casemanagementv1.OfferingPayloadType)
				offeringPayloadTypeModel.Group = core.StringPtr("crn_service_name")
				offeringPayloadTypeModel.Key = core.StringPtr("testString")
				offeringPayloadTypeModel.ID = core.StringPtr("testString")
				offeringPayloadTypeModel.Kind = core.StringPtr("testString")

				// Construct an instance of the EuPayload model
				euPayloadModel := new(casemanagementv1.EuPayload)
				euPayloadModel.Supported = core.BoolPtr(true)
				euPayloadModel.DataCenter = core.Int64Ptr(int64(38))

				// Construct an instance of the OfferingPayload model
				offeringPayloadModel := new(casemanagementv1.OfferingPayload)
				offeringPayloadModel.Name = core.StringPtr("testString")
				offeringPayloadModel.Type = offeringPayloadTypeModel

				// Construct an instance of the Resource model
				resourceModel := new(casemanagementv1.Resource)
				resourceModel.Crn = core.StringPtr("testString")
				resourceModel.Name = core.StringPtr("testString")
				resourceModel.Type = core.StringPtr("testString")
				resourceModel.ID = core.Int64Ptr(int64(38))
				resourceModel.Note = core.StringPtr("testString")

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Name = core.StringPtr("testString")
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsModel := new(casemanagementv1.CreateCaseOptions)
				createCaseOptionsModel.Type = core.StringPtr("technical")
				createCaseOptionsModel.Subject = core.StringPtr("testString")
				createCaseOptionsModel.Description = core.StringPtr("testString")
				createCaseOptionsModel.Severity = core.Int64Ptr(int64(1))
				createCaseOptionsModel.Eu = euPayloadModel
				createCaseOptionsModel.Offering = offeringPayloadModel
				createCaseOptionsModel.Resources = []casemanagementv1.Resource{*resourceModel}
				createCaseOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				createCaseOptionsModel.InvoiceNumber = core.StringPtr("testString")
				createCaseOptionsModel.SlaCreditRequest = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateCase(createCaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCase(getCaseOptions *GetCaseOptions)`, func() {
		bearerToken := "0ui9876453"
		getCasePath := "/case-management/v1/cases/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCasePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"number": "Number", "short_escription": "ShortEscription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"id": "ID", "value": "Value"}, "resources": [{"crn": "Crn", "name": "Name", "type": "Type", "id": 2, "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
			}))
			It(`Invoke GetCase successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				getCaseOptionsModel.Fields = []string{"testString"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCase(getCaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateCaseStatus(updateCaseStatusOptions *UpdateCaseStatusOptions)`, func() {
		bearerToken := "0ui9876453"
		updateCaseStatusPath := "/case-management/v1/cases/testString/status"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateCaseStatusPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"number": "Number", "short_escription": "ShortEscription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"id": "ID", "value": "Value"}, "resources": [{"crn": "Crn", "name": "Name", "type": "Type", "id": 2, "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
			}))
			It(`Invoke UpdateCaseStatus successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateCaseStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCaseStatusOptions model
				updateCaseStatusOptionsModel := new(casemanagementv1.UpdateCaseStatusOptions)
				updateCaseStatusOptionsModel.CaseNumber = core.StringPtr("testString")
				updateCaseStatusOptionsModel.Action = core.StringPtr("resolve")
				updateCaseStatusOptionsModel.Comment = core.StringPtr("testString")
				updateCaseStatusOptionsModel.ResolutionCode = core.Int64Ptr(int64(1))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateCaseStatus(updateCaseStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddComment(addCommentOptions *AddCommentOptions)`, func() {
		bearerToken := "0ui9876453"
		addCommentPath := "/case-management/v1/cases/testString/comments"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addCommentPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}`)
			}))
			It(`Invoke AddComment successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				addCommentOptionsModel.Comment = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AddComment(addCommentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddWatchlist(addWatchlistOptions *AddWatchlistOptions)`, func() {
		bearerToken := "0ui9876453"
		addWatchlistPath := "/case-management/v1/cases/testString/watchlist"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addWatchlistPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"added": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "failed": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}]}`)
			}))
			It(`Invoke AddWatchlist successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.AddWatchlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UserIdAndRealm model
				userIdAndRealmModel := new(casemanagementv1.UserIdAndRealm)
				userIdAndRealmModel.Realm = core.StringPtr("IBMid")
				userIdAndRealmModel.UserID = core.StringPtr("testString")

				// Construct an instance of the AddWatchlistOptions model
				addWatchlistOptionsModel := new(casemanagementv1.AddWatchlistOptions)
				addWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				addWatchlistOptionsModel.Watchlist = []casemanagementv1.UserIdAndRealm{*userIdAndRealmModel}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AddWatchlist(addWatchlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`RemoveWatchlist(removeWatchlistOptions *RemoveWatchlistOptions)`, func() {
		bearerToken := "0ui9876453"
		removeWatchlistPath := "/case-management/v1/cases/testString/watchlist"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(removeWatchlistPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `[{}]`)
			}))
			It(`Invoke RemoveWatchlist successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.RemoveWatchlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UserIdAndRealm model
				userIdAndRealmModel := new(casemanagementv1.UserIdAndRealm)
				userIdAndRealmModel.Realm = core.StringPtr("IBMid")
				userIdAndRealmModel.UserID = core.StringPtr("testString")

				// Construct an instance of the RemoveWatchlistOptions model
				removeWatchlistOptionsModel := new(casemanagementv1.RemoveWatchlistOptions)
				removeWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				removeWatchlistOptionsModel.Watchlist = []casemanagementv1.UserIdAndRealm{*userIdAndRealmModel}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.RemoveWatchlist(removeWatchlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddResource(addResourceOptions *AddResourceOptions)`, func() {
		bearerToken := "0ui9876453"
		addResourcePath := "/case-management/v1/cases/testString/resources"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addResourcePath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"crn": "Crn", "name": "Name", "type": "Type", "id": 2, "note": "Note"}`)
			}))
			It(`Invoke AddResource successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				addResourceOptionsModel.Name = core.StringPtr("testString")
				addResourceOptionsModel.Type = core.StringPtr("testString")
				addResourceOptionsModel.ID = core.Int64Ptr(int64(38))
				addResourceOptionsModel.Note = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AddResource(addResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UploadFile(uploadFileOptions *UploadFileOptions)`, func() {
		bearerToken := "0ui9876453"
		uploadFilePath := "/case-management/v1/cases/testString/attachments"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(uploadFilePath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}`)
			}))
			It(`Invoke UploadFile successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UploadFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UploadFileOptions model
				uploadFileOptionsModel := new(casemanagementv1.UploadFileOptions)
				uploadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				uploadFileOptionsModel.File = CreateMockReader("This is a mock file.")
				uploadFileOptionsModel.FileContentType = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UploadFile(uploadFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DownloadFile(downloadFileOptions *DownloadFileOptions)`, func() {
		bearerToken := "0ui9876453"
		downloadFilePath := "/case-management/v1/cases/testString/attachments/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(downloadFilePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/octet-stream")
				res.WriteHeader(200)
				fmt.Fprintf(res, `Contents of response byte-stream...`)
			}))
			It(`Invoke DownloadFile successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DownloadFile(downloadFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteFile(deleteFileOptions *DeleteFileOptions)`, func() {
		bearerToken := "0ui9876453"
		deleteFilePath := "/case-management/v1/cases/testString/attachments/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteFilePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}]}`)
			}))
			It(`Invoke DeleteFile successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteFile(deleteFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetEUSupport(getEUSupportOptions *GetEUSupportOptions)`, func() {
		bearerToken := "0ui9876453"
		getEuSupportPath := "/case-management/utilities/v1/eu-support"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getEuSupportPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"property": "supported", "values": [{"anyKey": "anyValue"}]}`)
			}))
			It(`Invoke GetEUSupport successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetEUSupport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEUSupportOptions model
				getEuSupportOptionsModel := new(casemanagementv1.GetEUSupportOptions)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetEUSupport(getEuSupportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetTechnicalOfferings(getTechnicalOfferingsOptions *GetTechnicalOfferingsOptions)`, func() {
		bearerToken := "0ui9876453"
		getTechnicalOfferingsPath := "/case-management/utilities/v1/offerings/technical"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTechnicalOfferingsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"offerings": [{"name": "Name", "type": {"group": "crn_service_name", "key": "Key"}}]}`)
			}))
			It(`Invoke GetTechnicalOfferings successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetTechnicalOfferings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTechnicalOfferingsOptions model
				getTechnicalOfferingsOptionsModel := new(casemanagementv1.GetTechnicalOfferingsOptions)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetTechnicalOfferings(getTechnicalOfferingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResolutionCodes(getResolutionCodesOptions *GetResolutionCodesOptions)`, func() {
		bearerToken := "0ui9876453"
		getResolutionCodesPath := "/case-management/utilities/v1/constants/resolution-codes"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResolutionCodesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"resolution_codes": [{"id": 2, "value": "Value"}]}`)
			}))
			It(`Invoke GetResolutionCodes successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResolutionCodes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResolutionCodesOptions model
				getResolutionCodesOptionsModel := new(casemanagementv1.GetResolutionCodesOptions)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResolutionCodes(getResolutionCodesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetStatuses(getStatusesOptions *GetStatusesOptions)`, func() {
		bearerToken := "0ui9876453"
		getStatusesPath := "/case-management/utilities/v1/constants/statuses"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getStatusesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"statuses": [{"id": "ID", "description": "Description"}]}`)
			}))
			It(`Invoke GetStatuses successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetStatuses(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetStatusesOptions model
				getStatusesOptionsModel := new(casemanagementv1.GetStatusesOptions)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetStatuses(getStatusesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a sample service client instance`, func() {
			testService, _ := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
				URL:           "http://casemanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewOfferingPayload successfully`, func() {
				name := "testString"
				var typeVar *casemanagementv1.OfferingPayloadType = nil
				_, err := testService.NewOfferingPayload(name, typeVar)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewUserIdAndRealm successfully`, func() {
				realm := "IBMid"
				userID := "testString"
				model, err := testService.NewUserIdAndRealm(realm, userID)
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
