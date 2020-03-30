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

package usagemeteringv4_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/usagemeteringv4"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`UsageMeteringV4`, func() {
	Describe(`ReportResourceUsage(reportResourceUsageOptions *ReportResourceUsageOptions)`, func() {
		bearerToken := "0ui9876453"
		reportResourceUsagePath := "/v4/metering/resources/testString/usage"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(reportResourceUsagePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{"resources": [{"status": 6, "location": "Location", "code": "Code", "message": "Message"}]}`)
			}))
			It(`Invoke ReportResourceUsage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usagemeteringv4.NewUsageMeteringV4(&usagemeteringv4.UsageMeteringV4Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ReportResourceUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MeasureAndQuantity model
				measureAndQuantityModel := new(usagemeteringv4.MeasureAndQuantity)
				measureAndQuantityModel.Measure = core.StringPtr("STORAGE")
				measureAndQuantityModel.Quantity = CreateMockMap()

				// Construct an instance of the ResourceInstanceUsage model
				resourceInstanceUsageModel := new(usagemeteringv4.ResourceInstanceUsage)
				resourceInstanceUsageModel.ResourceInstanceID = core.StringPtr("crn:v1:bluemix:public:database-service:us-south:a/1c8ae972c35e470d994b6faff9494ce1:793ff3d3-9fe3-4329-9ea0-404703a3c371::")
				resourceInstanceUsageModel.PlanID = core.StringPtr("database-lite")
				resourceInstanceUsageModel.Region = core.StringPtr("us-south")
				resourceInstanceUsageModel.Start = core.Int64Ptr(int64(1485907200000))
				resourceInstanceUsageModel.End = core.Int64Ptr(int64(1485907200000))
				resourceInstanceUsageModel.MeasuredUsage = []usagemeteringv4.MeasureAndQuantity{*measureAndQuantityModel}
				resourceInstanceUsageModel.ConsumerID = core.StringPtr("cf-application:ed20abbe-8870-44e6-90f7-56d764c21127")

				// Construct an instance of the ReportResourceUsageOptions model
				reportResourceUsageOptionsModel := new(usagemeteringv4.ReportResourceUsageOptions)
				reportResourceUsageOptionsModel.ResourceID = core.StringPtr("testString")
				reportResourceUsageOptionsModel.ResourceInstanceUsage = []usagemeteringv4.ResourceInstanceUsage{*resourceInstanceUsageModel}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ReportResourceUsage(reportResourceUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ReportCfresourceUsage(reportCfresourceUsageOptions *ReportCfresourceUsageOptions)`, func() {
		bearerToken := "0ui9876453"
		reportCfresourceUsagePath := "/v1/metering/resources/testString/usage"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(reportCfresourceUsagePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{"resources": [{"status": 6, "location": "Location", "code": "Code", "message": "Message"}]}`)
			}))
			It(`Invoke ReportCfresourceUsage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usagemeteringv4.NewUsageMeteringV4(&usagemeteringv4.UsageMeteringV4Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ReportCfresourceUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MeasureAndQuantity model
				measureAndQuantityModel := new(usagemeteringv4.MeasureAndQuantity)
				measureAndQuantityModel.Measure = core.StringPtr("STORAGE")
				measureAndQuantityModel.Quantity = CreateMockMap()

				// Construct an instance of the CfResourceInstanceUsage model
				cfResourceInstanceUsageModel := new(usagemeteringv4.CfResourceInstanceUsage)
				cfResourceInstanceUsageModel.OrganizationID = core.StringPtr("Public: us-south:102d9527-315a-4a71-afc2-068b1db6d68e Syndicated: ibm:dys0:us-south:e2566380-89cf-4c38-9290-7eb48cfca8f9")
				cfResourceInstanceUsageModel.SpaceID = core.StringPtr("89f0839a-812c-4180-8494-a453514c55e6")
				cfResourceInstanceUsageModel.ResourceInstanceID = core.StringPtr("5ca426de-5091-4f2f-8d87-28d37e7ff711")
				cfResourceInstanceUsageModel.PlanID = core.StringPtr("database-lite")
				cfResourceInstanceUsageModel.Region = core.StringPtr("us-south")
				cfResourceInstanceUsageModel.Start = core.Int64Ptr(int64(1485907200000))
				cfResourceInstanceUsageModel.End = core.Int64Ptr(int64(1485907200000))
				cfResourceInstanceUsageModel.MeasuredUsage = []usagemeteringv4.MeasureAndQuantity{*measureAndQuantityModel}
				cfResourceInstanceUsageModel.ConsumerID = core.StringPtr("cf-application:ed20abbe-8870-44e6-90f7-56d764c21127")

				// Construct an instance of the ReportCfresourceUsageOptions model
				reportCfresourceUsageOptionsModel := new(usagemeteringv4.ReportCfresourceUsageOptions)
				reportCfresourceUsageOptionsModel.ResourceID = core.StringPtr("testString")
				reportCfresourceUsageOptionsModel.CfResourceInstanceUsage = []usagemeteringv4.CfResourceInstanceUsage{*cfResourceInstanceUsageModel}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ReportCfresourceUsage(reportCfresourceUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a sample service client instance`, func() {
			testService, _ := usagemeteringv4.NewUsageMeteringV4(&usagemeteringv4.UsageMeteringV4Options{
				URL:           "http://usagemeteringv4modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCfResourceInstanceUsage successfully`, func() {
				organizationID := "Public: us-south:102d9527-315a-4a71-afc2-068b1db6d68e Syndicated: ibm:dys0:us-south:e2566380-89cf-4c38-9290-7eb48cfca8f9"
				spaceID := "89f0839a-812c-4180-8494-a453514c55e6"
				resourceInstanceID := "5ca426de-5091-4f2f-8d87-28d37e7ff711"
				planID := "database-lite"
				region := "us-south"
				start := int64(1485907200000)
				end := int64(1485907200000)
				measuredUsage := []usagemeteringv4.MeasureAndQuantity{}
				model, err := testService.NewCfResourceInstanceUsage(organizationID, spaceID, resourceInstanceID, planID, region, start, end, measuredUsage)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMeasureAndQuantity successfully`, func() {
				measure := "STORAGE"
				var quantity interface{} = nil
				_, err := testService.NewMeasureAndQuantity(measure, quantity)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewResourceInstanceUsage successfully`, func() {
				resourceInstanceID := "crn:v1:bluemix:public:database-service:us-south:a/1c8ae972c35e470d994b6faff9494ce1:793ff3d3-9fe3-4329-9ea0-404703a3c371::"
				planID := "database-lite"
				start := int64(1485907200000)
				end := int64(1485907200000)
				measuredUsage := []usagemeteringv4.MeasureAndQuantity{}
				model, err := testService.NewResourceInstanceUsage(resourceInstanceID, planID, start, end, measuredUsage)
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
