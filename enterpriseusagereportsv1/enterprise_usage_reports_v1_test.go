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

package enterpriseusagereportsv1_test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/enterpriseusagereportsv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`EnterpriseUsageReportsV1`, func() {
	Describe(`ListResourceUsageReport(listResourceUsageReportOptions *ListResourceUsageReportOptions)`, func() {
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listResourceUsageReportPath := "/v1/resource-usage-reports"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listResourceUsageReportPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))


				// TODO: Add check for children query parameter

				Expect(req.URL.Query()["month"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href"}, "reports": [{"entity_id": "de129b787b86403db7d3a14be2ae5f76", "entity_type": "enterprise", "entity_crn": "crn:v1:bluemix:public:enterprise::a/e9a57260546c4b4aa9ebfa316a82e56e::enterprise:de129b787b86403db7d3a14be2ae5f76", "entity_name": "Platform-Services", "billing_unit_id": "65719a07280a4022a9efa2f6ff4c3369", "billing_unit_crn": "crn:v1:bluemix:public:billing::a/3f99f8accbc848ea96f3c61a0ae22c44::billing-unit:65719a07280a4022a9efa2f6ff4c3369", "billing_unit_name": "Operations", "country_code": "USA", "currency_code": "USD", "month": "2017-08", "billable_cost": 12, "non_billable_cost": 15, "billable_rated_cost": 17, "non_billable_rated_cost": 20, "resources": [{"resource_id": "ResourceID", "billable_cost": 12, "billable_rated_cost": 17, "non_billable_cost": 15, "non_billable_rated_cost": 20, "plans": [{"plan_id": "PlanID", "pricing_region": "PricingRegion", "pricing_plan_id": "PricingPlanID", "billable": true, "cost": 4, "rated_cost": 9, "usage": [{"metric": "UP-TIME", "unit": "HOURS", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130, "price": [{"anyKey": "anyValue"}]}]}]}]}]}`)
			}))
			It(`Invoke ListResourceUsageReport successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListResourceUsageReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceUsageReportOptions model
				listResourceUsageReportOptionsModel := new(enterpriseusagereportsv1.ListResourceUsageReportOptions)
				listResourceUsageReportOptionsModel.EnterpriseID = core.StringPtr("testString")
				listResourceUsageReportOptionsModel.AccountGroupID = core.StringPtr("testString")
				listResourceUsageReportOptionsModel.AccountID = core.StringPtr("testString")
				listResourceUsageReportOptionsModel.Children = core.BoolPtr(true)
				listResourceUsageReportOptionsModel.Month = core.StringPtr("testString")
				listResourceUsageReportOptionsModel.BillingUnitID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceUsageReport(listResourceUsageReportOptionsModel)
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
