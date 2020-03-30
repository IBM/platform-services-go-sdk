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

package enterprisebillingunitsv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/enterprisebillingunitsv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`EnterpriseBillingUnitsV1`, func() {
	Describe(`GetBillingOptionByQuery(getBillingOptionByQueryOptions *GetBillingOptionByQueryOptions)`, func() {
		bearerToken := "0ui9876453"
		getBillingOptionByQueryPath := "/v1/billing-options"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getBillingOptionByQueryPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "billing_unit_id": "BillingUnitID", "start_date": "2019-05-01T00:00:00.000Z", "end_date": "2020-05-01T00:00:00.000Z", "state": "ACTIVE", "type": "SUBSCRIPTION", "category": "PLATFORM", "payment_instrument": {"anyKey": "anyValue"}, "duration_in_months": 11, "line_item_id": 10, "billing_system": {"anyKey": "anyValue"}, "renewal_mode_code": "RenewalModeCode", "updated_at": "2019-06-01T00:00:00.000Z"}`)
			}))
			It(`Invoke GetBillingOptionByQuery successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetBillingOptionByQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBillingOptionByQueryOptions model
				getBillingOptionByQueryOptionsModel := new(enterprisebillingunitsv1.GetBillingOptionByQueryOptions)
				getBillingOptionByQueryOptionsModel.BillingUnitID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetBillingOptionByQuery(getBillingOptionByQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetBillingUnitByID(getBillingUnitByIdOptions *GetBillingUnitByIdOptions)`, func() {
		bearerToken := "0ui9876453"
		getBillingUnitByIDPath := "/v1/billing-units/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getBillingUnitByIDPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "crn": "crn:v1:bluemix:public:billing::a/<<enterprise_account_id>>::billing-unit:<<billing_unit_id>>", "name": "Name", "enterprise_id": "EnterpriseID", "currency_code": "USD", "country_code": "USA", "master": true, "created_at": "2019-05-01T00:00:00.000Z"}`)
			}))
			It(`Invoke GetBillingUnitByID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetBillingUnitByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBillingUnitByIdOptions model
				getBillingUnitByIdOptionsModel := new(enterprisebillingunitsv1.GetBillingUnitByIdOptions)
				getBillingUnitByIdOptionsModel.BillingUnitID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetBillingUnitByID(getBillingUnitByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetBillingUnitByQuery(getBillingUnitByQueryOptions *GetBillingUnitByQueryOptions)`, func() {
		bearerToken := "0ui9876453"
		getBillingUnitByQueryPath := "/v1/billing-units"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getBillingUnitByQueryPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "crn": "crn:v1:bluemix:public:billing::a/<<enterprise_account_id>>::billing-unit:<<billing_unit_id>>", "name": "Name", "enterprise_id": "EnterpriseID", "currency_code": "USD", "country_code": "USA", "master": true, "created_at": "2019-05-01T00:00:00.000Z"}]}`)
			}))
			It(`Invoke GetBillingUnitByQuery successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetBillingUnitByQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBillingUnitByQueryOptions model
				getBillingUnitByQueryOptionsModel := new(enterprisebillingunitsv1.GetBillingUnitByQueryOptions)
				getBillingUnitByQueryOptionsModel.AccountID = core.StringPtr("testString")
				getBillingUnitByQueryOptionsModel.EnterpriseID = core.StringPtr("testString")
				getBillingUnitByQueryOptionsModel.AccountGroupID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetBillingUnitByQuery(getBillingUnitByQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCreditPools(getCreditPoolsOptions *GetCreditPoolsOptions)`, func() {
		bearerToken := "0ui9876453"
		getCreditPoolsPath := "/v1/credit-pools"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCreditPoolsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["date"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"rows_count": 2, "next_url": "NextURL", "resources": [{"type": "Type", "currency_code": "USD", "billing_unit_id": "BillingUnitID", "term_credits": [{"billing_option_id": "JWX986YRGFSHACQUEFOI", "category": "PLATFORM", "start_date": "2019-05-01T00:00:00.000Z", "end_date": "2020-04-30T23:59:29.999Z", "total_credits": 10000, "starting_balance": 9000, "used_credits": 9500, "current_balance": 0, "resources": [{"anyKey": "anyValue"}]}], "overage": {"cost": 500, "resources": [{"anyKey": "anyValue"}]}}]}`)
			}))
			It(`Invoke GetCreditPools successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCreditPools(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCreditPoolsOptions model
				getCreditPoolsOptionsModel := new(enterprisebillingunitsv1.GetCreditPoolsOptions)
				getCreditPoolsOptionsModel.BillingUnitID = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Date = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Type = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCreditPools(getCreditPoolsOptionsModel)
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
