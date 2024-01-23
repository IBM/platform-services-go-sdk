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

package partnerusagereportsv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnerusagereportsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Partner Usage Reports service.
//
// The following configuration properties are assumed to be defined:
// PARTNER_USAGE_REPORTS_URL=<service base url>
// PARTNER_USAGE_REPORTS_AUTH_TYPE=iam
// PARTNER_USAGE_REPORTS_APIKEY=<IAM apikey>
// PARTNER_USAGE_REPORTS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// PARTNER_USAGE_REPORTS_PARTNER_ID=<Enterprise ID of the distributor or reseller for which the report is requested>
// PARTNER_USAGE_REPORTS_CUSTOMER_ID=<Enterprise ID of the child customer for which the report is requested>
// PARTNER_USAGE_REPORTS_RESELLER_ID=<Enterprise ID of the reseller for which the report is requested>
// PARTNER_USAGE_REPORTS_BILLING_MONTH=<The billing month for which the usage report is requested. Format is `yyyy-mm`>
// PARTNER_USAGE_REPORTS_VIEWPOINT=<Enables partner to view the cost of provisioned services as applicable at each level of the hierarchy>

// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`PartnerUsageReportsV1 Examples Tests`, func() {

	const externalConfigFile = "../partner_usage_reports_v1.env"

	var (
		partnerUsageReportsService *partnerusagereportsv1.PartnerUsageReportsV1
		config                     map[string]string

		partnerId    string
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
			config, err = core.GetServiceProperties(partnerusagereportsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			partnerId = config["PARTNER_ID"]
			Expect(partnerId).ToNot(BeEmpty())

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

			partnerUsageReportsServiceOptions := &partnerusagereportsv1.PartnerUsageReportsV1Options{}

			partnerUsageReportsService, err = partnerusagereportsv1.NewPartnerUsageReportsV1UsingExternalConfig(partnerUsageReportsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(partnerUsageReportsService).ToNot(BeNil())
		})
	})

	Describe(`PartnerUsageReportsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageReport request example`, func() {
			fmt.Println("\nGetResourceUsageReport() result:")
			// begin-get_resource_usage_report
			getResourceUsageReportOptions := &partnerusagereportsv1.GetResourceUsageReportOptions{
				PartnerID: &partnerId,
				Month:     &billingMonth,
				Limit:     core.Int64Ptr(int64(30)),
			}

			pager, err := partnerUsageReportsService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			if err != nil {
				panic(err)
			}

			var allResults []partnerusagereportsv1.PartnerUsageReport
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-get_resource_usage_report
		})
	})
})
