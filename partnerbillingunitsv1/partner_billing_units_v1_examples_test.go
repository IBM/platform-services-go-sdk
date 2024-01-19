//go:build examples
// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnerbillingunitsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Partner Billing Units service.
//
// The following configuration properties are assumed to be defined:
// PARTNER_BILLING_UNITS_URL=<service base url>
// PARTNER_BILLING_UNITS_AUTH_TYPE=iam
// PARTNER_BILLING_UNITS_APIKEY=<IAM apikey>
// PARTNER_BILLING_UNITS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// PARTNER_BILLING_UNITS_PARTNER_ID=<Enterprise ID of the distributor or reseller for which the report is requested>
// PARTNER_BILLING_UNITS_CUSTOMER_ID=<Enterprise ID of the customer for which the report is requested>
// PARTNER_BILLING_UNITS_RESELLER_ID=<Enterprise ID of the reseller for which the report is requested>
// PARTNER_BILLING_UNITS_BILLING_MONTH=<The billing month (yyyy-mm) for which usage report will be retrieved>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`PartnerBillingUnitsV1 Examples Tests`, func() {

	const externalConfigFile = "../partner_billing_units_v1.env"

	var (
		partnerBillingUnitsService *partnerbillingunitsv1.PartnerBillingUnitsV1
		config                     map[string]string

		partnerID    string
		customerID   string
		resellerID   string
		billingMonth string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(partnerbillingunitsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

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
			var err error

			// begin-common

			partnerBillingUnitsServiceOptions := &partnerbillingunitsv1.PartnerBillingUnitsV1Options{}

			partnerBillingUnitsService, err = partnerbillingunitsv1.NewPartnerBillingUnitsV1UsingExternalConfig(partnerBillingUnitsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(partnerBillingUnitsService).ToNot(BeNil())
		})
	})

	Describe(`PartnerBillingUnitsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBillingOptions request example of a partner`, func() {
			fmt.Println("\nGetBillingOptions() result:")
			// begin-get_billing_options

			getBillingOptionsOptions := partnerBillingUnitsService.NewGetBillingOptionsOptions(
				partnerID,
				billingMonth,
			)

			billingOptionsSummary, response, err := partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(billingOptionsSummary, "", "  ")
			fmt.Println(string(b))

			// end-get_billing_options

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingOptionsSummary).ToNot(BeNil())
		})
		It(`GetCreditPoolsReport request example of a partner`, func() {
			fmt.Println("\nGetCreditPoolsReport() result:")
			// begin-get_credit_pools_report

			getCreditPoolsReportOptions := partnerBillingUnitsService.NewGetCreditPoolsReportOptions(
				partnerID,
				billingMonth,
			)

			creditPoolsReportSummary, response, err := partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(creditPoolsReportSummary, "", "  ")
			fmt.Println(string(b))

			// end-get_credit_pools_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(creditPoolsReportSummary).ToNot(BeNil())
		})
	})
})
