//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2021.
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

package contextbasedrestrictionsv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Context Based Restrictions service.
//
// The following configuration properties are assumed to be defined:
// CONTEXT_BASED_RESTRICTIONS_URL=<service base url>
// CONTEXT_BASED_RESTRICTIONS_AUTH_TYPE=iam
// CONTEXT_BASED_RESTRICTIONS_APIKEY=<IAM apikey>
// CONTEXT_BASED_RESTRICTIONS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../context_based_restrictions_v1.env"

var (
	contextBasedRestrictionsService *contextbasedrestrictionsv1.ContextBasedRestrictionsV1
	config                          map[string]string
	configLoaded                    bool = false
	testAccountID                   string
	testServiceName                 string
	zoneID                          string
	zoneRev                         string
	ruleID                          string
	ruleRev                         string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`contextbasedrestrictionsv1.ContextBasedRestrictionsV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(contextbasedrestrictionsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			testAccountID = config["TEST_ACCOUNT_ID"]
			if testAccountID == "" {
				Skip("Unable to load TEST_ACCOUNT_ID configuration property, skipping tests")
			}

			testServiceName = config["TEST_SERVICE_NAME"]
			if testServiceName == "" {
				Skip("Unable to load TEST_SERVICE_NAME configuration property, skipping tests")
			}

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

			contextBasedRestrictionsServiceOptions := &contextbasedrestrictionsv1.Options{}

			contextBasedRestrictionsService, err = contextbasedrestrictionsv1.NewContextBasedRestrictionsV1UsingExternalConfig(contextBasedRestrictionsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(contextBasedRestrictionsService).ToNot(BeNil())
		})
	})

	Describe(`contextbasedrestrictionsv1.ContextBasedRestrictionsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateZone request example`, func() {
			fmt.Println("\nCreateZone() result:")
			// begin-create_zone

			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			createZoneOptions := contextBasedRestrictionsService.NewCreateZoneOptions()
			createZoneOptions.SetName("SDK TEST - an example of zone")
			createZoneOptions.SetAccountID(testAccountID)
			createZoneOptions.SetDescription("this is an example of zone")
			createZoneOptions.SetAddresses([]contextbasedrestrictionsv1.AddressIntf{addressModel})

			outZone, response, err := contextBasedRestrictionsService.CreateZone(createZoneOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outZone, "", "  ")
			fmt.Println(string(b))

			// end-create_zone

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(outZone).ToNot(BeNil())
			zoneID = *outZone.ID

		})
		It(`ListZones request example`, func() {
			fmt.Println("\nListZones() result:")
			// begin-list_zones

			listZonesOptions := contextBasedRestrictionsService.NewListZonesOptions(
				testAccountID,
			)

			outZonePage, response, err := contextBasedRestrictionsService.ListZones(listZonesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outZonePage, "", "  ")
			fmt.Println(string(b))

			// end-list_zones

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outZonePage).ToNot(BeNil())

		})
		It(`GetZone request example`, func() {
			fmt.Println("\nGetZone() result:")
			// begin-get_zone

			getZoneOptions := contextBasedRestrictionsService.NewGetZoneOptions(
				zoneID,
			)

			outZone, response, err := contextBasedRestrictionsService.GetZone(getZoneOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outZone, "", "  ")
			fmt.Println(string(b))

			// end-get_zone

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outZone).ToNot(BeNil())
			zoneRev = response.Headers.Get("Etag")

		})
		It(`ReplaceZone request example`, func() {
			fmt.Println("\nReplaceZone() result:")
			// begin-replace_zone

			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			replaceZoneOptions := contextBasedRestrictionsService.NewReplaceZoneOptions(
				zoneID,
				zoneRev,
			)
			replaceZoneOptions.SetName("SDK TEST - an example of updated zone")
			replaceZoneOptions.SetAccountID(testAccountID)
			replaceZoneOptions.SetDescription("this is an example of updated zone")
			replaceZoneOptions.SetAddresses([]contextbasedrestrictionsv1.AddressIntf{addressModel})

			outZone, response, err := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outZone, "", "  ")
			fmt.Println(string(b))

			// end-replace_zone

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outZone).ToNot(BeNil())

		})
		It(`ListAvailableServiceRefTargets request example`, func() {
			fmt.Println("\nListAvailableServiceRefTargets() result:")
			// begin-list_available_serviceref_targets

			listAvailableServiceRefTargetsOptions := contextBasedRestrictionsService.NewListAvailableServiceRefTargetsOptions()

			serviceRefTargetPage, response, err := contextBasedRestrictionsService.ListAvailableServiceRefTargets(listAvailableServiceRefTargetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceRefTargetPage, "", "  ")
			fmt.Println(string(b))

			// end-list_available_serviceref_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceRefTargetPage).ToNot(BeNil())

		})
		It(`CreateRule request example`, func() {
			fmt.Println("\nCreateRule() result:")
			// begin-create_rule

			ruleContextAttributeModel := &contextbasedrestrictionsv1.RuleContextAttribute{
				Name:  core.StringPtr("networkZoneId"),
				Value: core.StringPtr(zoneID),
			}

			ruleContextModel := &contextbasedrestrictionsv1.RuleContext{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel},
			}

			resourceModel := &contextbasedrestrictionsv1.Resource{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{
						Name:  core.StringPtr("accountId"),
						Value: core.StringPtr(testAccountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr(testServiceName),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("tagValue"),
					},
				},
			}

			createRuleOptions := contextBasedRestrictionsService.NewCreateRuleOptions()
			createRuleOptions.SetDescription("SDK TEST - this is an example of rule")
			createRuleOptions.SetContexts([]contextbasedrestrictionsv1.RuleContext{*ruleContextModel})
			createRuleOptions.SetResources([]contextbasedrestrictionsv1.Resource{*resourceModel})

			outRule, response, err := contextBasedRestrictionsService.CreateRule(createRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outRule, "", "  ")
			fmt.Println(string(b))

			// end-create_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(outRule).ToNot(BeNil())
			ruleID = *outRule.ID

		})
		It(`ListRules request example`, func() {
			fmt.Println("\nListRules() result:")
			// begin-list_rules

			listRulesOptions := contextBasedRestrictionsService.NewListRulesOptions(
				testAccountID,
			)

			outRulePage, response, err := contextBasedRestrictionsService.ListRules(listRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outRulePage, "", "  ")
			fmt.Println(string(b))

			// end-list_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outRulePage).ToNot(BeNil())
		})
		It(`GetRule request example`, func() {
			fmt.Println("\nGetRule() result:")
			// begin-get_rule

			getRuleOptions := contextBasedRestrictionsService.NewGetRuleOptions(
				ruleID,
			)

			outRule, response, err := contextBasedRestrictionsService.GetRule(getRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outRule, "", "  ")
			fmt.Println(string(b))

			// end-get_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outRule).ToNot(BeNil())
			ruleRev = response.Headers.Get("Etag")
		})
		It(`ReplaceRule request example`, func() {
			fmt.Println("\nReplaceRule() result:")
			// begin-replace_rule

			ruleContextAttributeModel := &contextbasedrestrictionsv1.RuleContextAttribute{
				Name:  core.StringPtr("networkZoneId"),
				Value: core.StringPtr(zoneID),
			}

			ruleContextModel := &contextbasedrestrictionsv1.RuleContext{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel},
			}

			resourceModel := &contextbasedrestrictionsv1.Resource{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{
						Name:  core.StringPtr("accountId"),
						Value: core.StringPtr(testAccountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr(testServiceName),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("updatedTagValue"),
					},
				},
			}

			replaceRuleOptions := contextBasedRestrictionsService.NewReplaceRuleOptions(
				ruleID,
				ruleRev,
			)
			replaceRuleOptions.SetDescription("this is an example of updated rule")
			replaceRuleOptions.SetContexts([]contextbasedrestrictionsv1.RuleContext{*ruleContextModel})
			replaceRuleOptions.SetResources([]contextbasedrestrictionsv1.Resource{*resourceModel})

			outRule, response, err := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outRule, "", "  ")
			fmt.Println(string(b))

			// end-replace_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outRule).ToNot(BeNil())

		})
		It(`GetAccountSettings request example`, func() {
			fmt.Println("\nGetAccountSettings() result:")
			// begin-get_account_settings

			getAccountSettingsOptions := contextBasedRestrictionsService.NewGetAccountSettingsOptions(
				testAccountID,
			)

			outAccountSettings, response, err := contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outAccountSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_account_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outAccountSettings).ToNot(BeNil())

		})
		It(`DeleteRule request example`, func() {
			// begin-delete_rule

			deleteRuleOptions := contextBasedRestrictionsService.NewDeleteRuleOptions(
				ruleID,
			)

			response, err := contextBasedRestrictionsService.DeleteRule(deleteRuleOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_rule
			fmt.Printf("\nDeleteRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteZone request example`, func() {
			// begin-delete_zone

			deleteZoneOptions := contextBasedRestrictionsService.NewDeleteZoneOptions(
				zoneID,
			)

			response, err := contextBasedRestrictionsService.DeleteZone(deleteZoneOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_zone
			fmt.Printf("\nDeleteZone() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
