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

package usagereportsv4_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/usagereportsv4"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`UsageReportsV4`, func() {
	Describe(`GetAccountSummary(getAccountSummaryOptions *GetAccountSummaryOptions)`, func() {
		bearerToken := "0ui9876453"
		getAccountSummaryPath := "/v4/accounts/testString/summary/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccountSummaryPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"account_id": "AccountID", "billing_month": "BillingMonth", "billing_country_code": "BillingCountryCode", "billing_currency_code": "BillingCurrencyCode", "resources": {"billable_cost": 12, "non_billable_cost": 15}, "offers": [{"offer_id": "OfferID", "credits_total": 12, "offer_template": "OfferTemplate", "valid_from": "2019-01-01T12:00:00", "expires_on": "2019-01-01T12:00:00", "credits": {"total": 5, "starting_balance": 15, "used": 4, "balance": 7}}], "subscription": {"overage": 7, "subscriptions": [{"subscription_id": "SubscriptionID", "charge_agreement_number": "ChargeAgreementNumber", "type": "Type", "subscription_amount": 18, "start": "2019-01-01T12:00:00", "end": "2019-01-01T12:00:00", "credits_total": 12, "terms": [{"start": "2019-01-01T12:00:00", "end": "2019-01-01T12:00:00", "credits": {"total": 5, "starting_balance": 15, "used": 4, "balance": 7}}]}]}, "support": [{"cost": 4, "type": "Type", "overage": 7}]}`)
			}))
			It(`Invoke GetAccountSummary successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccountSummary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSummaryOptions model
				getAccountSummaryOptionsModel := new(usagereportsv4.GetAccountSummaryOptions)
				getAccountSummaryOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Billingmonth = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccountSummary(getAccountSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAccountUsage(getAccountUsageOptions *GetAccountUsageOptions)`, func() {
		bearerToken := "0ui9876453"
		getAccountUsagePath := "/v4/accounts/testString/usage/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccountUsagePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"account_id": "AccountID", "pricing_country": "USA", "currency_code": "USD", "month": "2017-08", "resources": [{"resource_id": "ResourceID", "billable_cost": 12, "non_billable_cost": 15, "plans": [{"plan_id": "PlanID", "pricing_region": "PricingRegion", "billable": true, "cost": 4, "usage": [{"metric": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "price": [{"mapKey": {"anyKey": "anyValue"}}], "unit": "HOURS", "non_chargeable": true}]}]}]}`)
			}))
			It(`Invoke GetAccountUsage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccountUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountUsageOptions model
				getAccountUsageOptionsModel := new(usagereportsv4.GetAccountUsageOptions)
				getAccountUsageOptionsModel.AccountID = core.StringPtr("testString")
				getAccountUsageOptionsModel.Billingmonth = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccountUsage(getAccountUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResourceGroupUsage(getResourceGroupUsageOptions *GetResourceGroupUsageOptions)`, func() {
		bearerToken := "0ui9876453"
		getResourceGroupUsagePath := "/v4/accounts/testString/resource_groups/testString/usage/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourceGroupUsagePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"account_id": "AccountID", "resource_group_id": "ResourceGroupID", "pricing_country": "USA", "currency_code": "USD", "month": "2017-08", "resources": [{"resource_id": "ResourceID", "billable_cost": 12, "non_billable_cost": 15, "plans": [{"plan_id": "PlanID", "pricing_region": "PricingRegion", "billable": true, "cost": 4, "usage": [{"metric": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "price": [{"mapKey": {"anyKey": "anyValue"}}], "unit": "HOURS", "non_chargeable": true}]}]}]}`)
			}))
			It(`Invoke GetResourceGroupUsage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceGroupUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceGroupUsageOptions model
				getResourceGroupUsageOptionsModel := new(usagereportsv4.GetResourceGroupUsageOptions)
				getResourceGroupUsageOptionsModel.AccountID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Billingmonth = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceGroupUsage(getResourceGroupUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetOrganizationUsage(getOrganizationUsageOptions *GetOrganizationUsageOptions)`, func() {
		bearerToken := "0ui9876453"
		getOrganizationUsagePath := "/v4/accounts/testString/organizations/testString/usage/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getOrganizationUsagePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"account_id": "AccountID", "organization_id": "OrganizationID", "pricing_country": "USA", "currency_code": "USD", "month": "2017-08", "resources": [{"resource_id": "ResourceID", "billable_cost": 12, "non_billable_cost": 15, "plans": [{"plan_id": "PlanID", "pricing_region": "PricingRegion", "billable": true, "cost": 4, "usage": [{"metric": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "price": [{"mapKey": {"anyKey": "anyValue"}}], "unit": "HOURS", "non_chargeable": true}]}]}]}`)
			}))
			It(`Invoke GetOrganizationUsage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetOrganizationUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOrganizationUsageOptions model
				getOrganizationUsageOptionsModel := new(usagereportsv4.GetOrganizationUsageOptions)
				getOrganizationUsageOptionsModel.AccountID = core.StringPtr("testString")
				getOrganizationUsageOptionsModel.OrganizationID = core.StringPtr("testString")
				getOrganizationUsageOptionsModel.Billingmonth = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetOrganizationUsage(getOrganizationUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAccountInstancesUsage(getAccountInstancesUsageOptions *GetAccountInstancesUsageOptions)`, func() {
		bearerToken := "0ui9876453"
		getAccountInstancesUsagePath := "/v4/accounts/testString/resource_instances/usage/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccountInstancesUsagePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["organization_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"limit": 5, "count": 5, "first": {"href": "Href", "offset": "Offset"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"account_id": "AccountID", "resource_instance_id": "ResourceInstanceID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "organization_id": "OrganizationID", "space": "Space", "consumer_id": "ConsumerID", "region": "Region", "pricing_region": "PricingRegion", "pricing_country": "USA", "currency_code": "USD", "billable": true, "plan_id": "PlanID", "month": "2017-08", "usage": [{"metric": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "price": [{"mapKey": {"anyKey": "anyValue"}}], "unit": "HOURS", "non_chargeable": true}]}]}`)
			}))
			It(`Invoke GetAccountInstancesUsage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccountInstancesUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountInstancesUsageOptions model
				getAccountInstancesUsageOptionsModel := new(usagereportsv4.GetAccountInstancesUsageOptions)
				getAccountInstancesUsageOptionsModel.AccountID = core.StringPtr("testString")
				getAccountInstancesUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountInstancesUsageOptionsModel.Limit = core.Int64Ptr(int64(38))
				getAccountInstancesUsageOptionsModel.Start = core.StringPtr("testString")
				getAccountInstancesUsageOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getAccountInstancesUsageOptionsModel.OrganizationID = core.StringPtr("testString")
				getAccountInstancesUsageOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getAccountInstancesUsageOptionsModel.ResourceID = core.StringPtr("testString")
				getAccountInstancesUsageOptionsModel.PlanID = core.StringPtr("testString")
				getAccountInstancesUsageOptionsModel.Region = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccountInstancesUsage(getAccountInstancesUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResourceGroupInstancesUsage(getResourceGroupInstancesUsageOptions *GetResourceGroupInstancesUsageOptions)`, func() {
		bearerToken := "0ui9876453"
		getResourceGroupInstancesUsagePath := "/v4/accounts/testString/resource_groups/testString/resource_instances/usage/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourceGroupInstancesUsagePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"limit": 5, "count": 5, "first": {"href": "Href", "offset": "Offset"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"account_id": "AccountID", "resource_instance_id": "ResourceInstanceID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "organization_id": "OrganizationID", "space": "Space", "consumer_id": "ConsumerID", "region": "Region", "pricing_region": "PricingRegion", "pricing_country": "USA", "currency_code": "USD", "billable": true, "plan_id": "PlanID", "month": "2017-08", "usage": [{"metric": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "price": [{"mapKey": {"anyKey": "anyValue"}}], "unit": "HOURS", "non_chargeable": true}]}]}`)
			}))
			It(`Invoke GetResourceGroupInstancesUsage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceGroupInstancesUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceGroupInstancesUsageOptions model
				getResourceGroupInstancesUsageOptionsModel := new(usagereportsv4.GetResourceGroupInstancesUsageOptions)
				getResourceGroupInstancesUsageOptionsModel.AccountID = core.StringPtr("testString")
				getResourceGroupInstancesUsageOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceGroupInstancesUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceGroupInstancesUsageOptionsModel.Limit = core.Int64Ptr(int64(38))
				getResourceGroupInstancesUsageOptionsModel.Start = core.StringPtr("testString")
				getResourceGroupInstancesUsageOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceGroupInstancesUsageOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceGroupInstancesUsageOptionsModel.PlanID = core.StringPtr("testString")
				getResourceGroupInstancesUsageOptionsModel.Region = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceGroupInstancesUsage(getResourceGroupInstancesUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetOrganizationInstancesUsage(getOrganizationInstancesUsageOptions *GetOrganizationInstancesUsageOptions)`, func() {
		bearerToken := "0ui9876453"
		getOrganizationInstancesUsagePath := "/v4/accounts/testString/organizations/testString/resource_instances/usage/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getOrganizationInstancesUsagePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"limit": 5, "count": 5, "first": {"href": "Href", "offset": "Offset"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"account_id": "AccountID", "resource_instance_id": "ResourceInstanceID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "organization_id": "OrganizationID", "space": "Space", "consumer_id": "ConsumerID", "region": "Region", "pricing_region": "PricingRegion", "pricing_country": "USA", "currency_code": "USD", "billable": true, "plan_id": "PlanID", "month": "2017-08", "usage": [{"metric": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "price": [{"mapKey": {"anyKey": "anyValue"}}], "unit": "HOURS", "non_chargeable": true}]}]}`)
			}))
			It(`Invoke GetOrganizationInstancesUsage successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetOrganizationInstancesUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOrganizationInstancesUsageOptions model
				getOrganizationInstancesUsageOptionsModel := new(usagereportsv4.GetOrganizationInstancesUsageOptions)
				getOrganizationInstancesUsageOptionsModel.AccountID = core.StringPtr("testString")
				getOrganizationInstancesUsageOptionsModel.OrganizationID = core.StringPtr("testString")
				getOrganizationInstancesUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getOrganizationInstancesUsageOptionsModel.Limit = core.Int64Ptr(int64(38))
				getOrganizationInstancesUsageOptionsModel.Start = core.StringPtr("testString")
				getOrganizationInstancesUsageOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getOrganizationInstancesUsageOptionsModel.ResourceID = core.StringPtr("testString")
				getOrganizationInstancesUsageOptionsModel.PlanID = core.StringPtr("testString")
				getOrganizationInstancesUsageOptionsModel.Region = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetOrganizationInstancesUsage(getOrganizationInstancesUsageOptionsModel)
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
