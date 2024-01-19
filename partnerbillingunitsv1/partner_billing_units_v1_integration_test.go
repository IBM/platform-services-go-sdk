//go:build integration
// +build integration

/**
 * (C) Copyright IBM Corp. 2024.
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

package partnerbillingunitsv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnerbillingunitsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the partnerbillingunitsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PartnerBillingUnitsV1 Integration Tests`, func() {
	const externalConfigFile = "../partner_billing_units_v1.env"

	var (
		err                        error
		partnerBillingUnitsService *partnerbillingunitsv1.PartnerBillingUnitsV1
		serviceURL                 string
		config                     map[string]string

		partnerID    string
		customerID   string
		resellerID   string
		billingMonth string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(partnerbillingunitsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)

			partnerID = config["PARTNER_ID"]
			Expect(partnerID).ToNot(BeEmpty())

			customerID = config["CUSTOMER_ID"]
			Expect(customerID).ToNot(BeEmpty())

			resellerID = config["RESELLER_ID"]
			Expect(resellerID).ToNot(BeEmpty())

			billingMonth = config["BILLING_MONTH"]
			Expect(billingMonth).ToNot(BeEmpty())

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			partnerBillingUnitsServiceOptions := &partnerbillingunitsv1.PartnerBillingUnitsV1Options{}

			partnerBillingUnitsService, err = partnerbillingunitsv1.NewPartnerBillingUnitsV1UsingExternalConfig(partnerBillingUnitsServiceOptions)
			Expect(err).To(BeNil())
			Expect(partnerBillingUnitsService).ToNot(BeNil())
			Expect(partnerBillingUnitsService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			partnerBillingUnitsService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetBillingOptions - Get customers billing options of a partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBillingOptions(getBillingOptionsOptions *GetBillingOptionsOptions)`, func() {
			getBillingOptionsOptions := &partnerbillingunitsv1.GetBillingOptionsOptions{
				PartnerID: &partnerID,
				Date:      &billingMonth,
				Limit:     core.Int64Ptr(int64(30)),
			}

			billingOptionsSummary, response, err := partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingOptionsSummary).ToNot(BeNil())
		})
	})
	Describe(`GetBillingOptions - Get customers billing options of a Reseller for a specific partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBillingOptions(getBillingOptionsOptions *GetBillingOptionsOptions)`, func() {
			getBillingOptionsOptions := &partnerbillingunitsv1.GetBillingOptionsOptions{
				PartnerID:  &partnerID,
				ResellerID: &resellerID,
				Date:       &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			billingOptionsSummary, response, err := partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingOptionsSummary).ToNot(BeNil())
		})
	})
	Describe(`GetBillingOptions - Get customers billing options of an end customer for a specific partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBillingOptions(getBillingOptionsOptions *GetBillingOptionsOptions)`, func() {
			getBillingOptionsOptions := &partnerbillingunitsv1.GetBillingOptionsOptions{
				PartnerID:  &partnerID,
				CustomerID: &customerID,
				Date:       &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			billingOptionsSummary, response, err := partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingOptionsSummary).ToNot(BeNil())
		})
	})

	Describe(`GetCreditPoolsReport - Get subscription burn-down report of a partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCreditPoolsReport(getCreditPoolsReportOptions *GetCreditPoolsReportOptions)`, func() {
			getCreditPoolsReportOptions := &partnerbillingunitsv1.GetCreditPoolsReportOptions{
				PartnerID: &partnerID,
				Date:      &billingMonth,
				Limit:     core.Int64Ptr(int64(30)),
			}

			creditPoolsReportSummary, response, err := partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(creditPoolsReportSummary).ToNot(BeNil())
		})
	})
	Describe(`GetCreditPoolsReport - Get subscription burn-down report of a Reseller for a specific partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCreditPoolsReport(getCreditPoolsReportOptions *GetCreditPoolsReportOptions)`, func() {
			getCreditPoolsReportOptions := &partnerbillingunitsv1.GetCreditPoolsReportOptions{
				PartnerID:  &partnerID,
				ResellerID: &resellerID,
				Date:       &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			creditPoolsReportSummary, response, err := partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(creditPoolsReportSummary).ToNot(BeNil())
		})
	})
	Describe(`GetCreditPoolsReport - Get subscription burn-down report of an end customer for a specific partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCreditPoolsReport(getCreditPoolsReportOptions *GetCreditPoolsReportOptions)`, func() {
			getCreditPoolsReportOptions := &partnerbillingunitsv1.GetCreditPoolsReportOptions{
				PartnerID:  &partnerID,
				CustomerID: &customerID,
				Date:       &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			creditPoolsReportSummary, response, err := partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(creditPoolsReportSummary).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
