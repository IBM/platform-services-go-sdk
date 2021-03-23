// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/usagereportsv4"
)

//
// This file provides an example of how to use the Usage Reports service.
//
// The following configuration properties are assumed to be defined:
// USAGE_REPORTS_URL=<service url>
// USAGE_REPORTS_AUTHTYPE=iam
// USAGE_REPORTS_APIKEY=<IAM api key of user with authority to create rules>
// USAGE_REPORTS_AUTH_URL=<IAM token service URL - omit this if using the production environment>
// USAGE_REPORTS_ACCOUNT_ID=<the id of the account whose usage info will be retrieved>
// USAGE_REPORTS_RESOURCE_GROUP_ID=<the id of the resource group whose usage info will be retrieved>
// USAGE_REPORTS_ORG_ID=<the id of the organization whose usage info will be retrieved>
// USAGE_REPORTS_BILLING_MONTH=<the billing month (yyyy-mm) for which usage info will be retrieved>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../usage_reports.env"

var (
	usageReportsService *usagereportsv4.UsageReportsV4
	config              map[string]string
	configLoaded        bool = false

	accountID       string
	resourceGroupID string
	orgID           string
	billingMonth    string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`UsageReportsV4 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(usagereportsv4.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			resourceGroupID = config["RESOURCE_GROUP_ID"]
			Expect(resourceGroupID).ToNot(BeEmpty())

			orgID = config["ORG_ID"]
			Expect(orgID).ToNot(BeEmpty())

			billingMonth = config["BILLING_MONTH"]
			Expect(billingMonth).ToNot(BeEmpty())

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			usageReportsServiceOptions := &usagereportsv4.UsageReportsV4Options{}

			usageReportsService, err = usagereportsv4.NewUsageReportsV4UsingExternalConfig(usageReportsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(usageReportsService).ToNot(BeNil())
		})
	})

	Describe(`UsageReportsV4 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSummary request example`, func() {
			// begin-get_account_summary

			getAccountSummaryOptions := usageReportsService.NewGetAccountSummaryOptions(
				accountID,
				billingMonth,
			)

			accountSummary, response, err := usageReportsService.GetAccountSummary(getAccountSummaryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSummary, "", "  ")
			fmt.Printf("\nGetAccountSummary() result:\n %s \n", string(b))

			// end-get_account_summary

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSummary).ToNot(BeNil())
		})
		It(`GetAccountUsage request example`, func() {
			// begin-get_account_usage

			getAccountUsageOptions := usageReportsService.NewGetAccountUsageOptions(
				accountID,
				billingMonth,
			)

			accountUsage, response, err := usageReportsService.GetAccountUsage(getAccountUsageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountUsage, "", "  ")
			fmt.Printf("\nGetAccountUsage() result:\n %s \n", string(b))

			// end-get_account_usage

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountUsage).ToNot(BeNil())
		})
		It(`GetResourceGroupUsage request example`, func() {
			// begin-get_resource_group_usage

			getResourceGroupUsageOptions := usageReportsService.NewGetResourceGroupUsageOptions(
				accountID,
				resourceGroupID,
				billingMonth,
			)

			resourceGroupUsage, response, err := usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceGroupUsage, "", "  ")
			fmt.Printf("\nGetResourceGroupUsage() result:\n %s \n", string(b))

			// end-get_resource_group_usage

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceGroupUsage).ToNot(BeNil())
		})
		It(`GetOrgUsage request example`, func() {
			// begin-get_org_usage

			getOrgUsageOptions := usageReportsService.NewGetOrgUsageOptions(
				accountID,
				orgID,
				billingMonth,
			)

			orgUsage, response, err := usageReportsService.GetOrgUsage(getOrgUsageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(orgUsage, "", "  ")
			fmt.Printf("\nGetOrgUsage() result:\n %s \n", string(b))

			// end-get_org_usage

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(orgUsage).ToNot(BeNil())
		})
		It(`GetResourceUsageAccount request example`, func() {
			// begin-get_resource_usage_account

			getResourceUsageAccountOptions := usageReportsService.NewGetResourceUsageAccountOptions(
				accountID,
				billingMonth,
			)

			instancesUsage, response, err := usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instancesUsage, "", "  ")
			fmt.Printf("\nGetResourceUsageAccount() result:\n %s \n", string(b))

			// end-get_resource_usage_account

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instancesUsage).ToNot(BeNil())
		})
		It(`GetResourceUsageResourceGroup request example`, func() {
			// begin-get_resource_usage_resource_group

			getResourceUsageResourceGroupOptions := usageReportsService.NewGetResourceUsageResourceGroupOptions(
				accountID,
				resourceGroupID,
				billingMonth,
			)

			instancesUsage, response, err := usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instancesUsage, "", "  ")
			fmt.Printf("\nGetResourceUsageResourceGroup() result:\n %s \n", string(b))

			// end-get_resource_usage_resource_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instancesUsage).ToNot(BeNil())
		})
		It(`GetResourceUsageOrg request example`, func() {
			// begin-get_resource_usage_org

			getResourceUsageOrgOptions := usageReportsService.NewGetResourceUsageOrgOptions(
				accountID,
				orgID,
				billingMonth,
			)

			instancesUsage, response, err := usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instancesUsage, "", "  ")
			fmt.Printf("\nGetResourceUsageOrg() result:\n %s \n", string(b))

			// end-get_resource_usage_org

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instancesUsage).ToNot(BeNil())
		})
	})
})
