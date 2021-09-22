//go:build integration
// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	NonExistentID = "1234567890abcdef1234567890abcdef"
	InvalidID     = "this_is_an_invalid_id"
)

/**
 * This file contains an integration test for the contextbasedrestrictionsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`ContextBasedRestrictionsV1 Integration Tests`, func() {

	const externalConfigFile = "../context_based_restrictions_v1.env"

	var (
		err                             error
		contextBasedRestrictionsService *contextbasedrestrictionsv1.ContextBasedRestrictionsV1
		serviceURL                      string
		config                          map[string]string
		testAccountID                   string
		testServiceName                 string
		zoneID                          string
		zoneRev                         string
		ruleID                          string
		ruleRev                         string
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

			err := os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			if err != nil {
				Skip("Error setting IBM_CREDENTIALS_FILE environment variable, skipping tests: " + err.Error())
			}

			config, err = core.GetServiceProperties(contextbasedrestrictionsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			testAccountID = config["TEST_ACCOUNT_ID"]
			if testAccountID == "" {
				Skip("Unable to load TEST_ACCOUNT_ID configuration property, skipping tests")
			}

			testServiceName = config["TEST_SERVICE_NAME"]
			if testServiceName == "" {
				Skip("Unable to load TEST_SERVICE_NAME configuration property, skipping tests")
			}

			fmt.Printf("\nService URL: %s\n", serviceURL)
			fmt.Printf("Test Account ID: %s\n", testAccountID)
			fmt.Printf("Test Service Name: %s\n", testServiceName)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			contextBasedRestrictionsServiceOptions := &contextbasedrestrictionsv1.Options{}
			contextBasedRestrictionsService, err = contextbasedrestrictionsv1.NewContextBasedRestrictionsV1UsingExternalConfig(contextBasedRestrictionsServiceOptions)

			Expect(err).To(BeNil())
			Expect(contextBasedRestrictionsService).ToNot(BeNil())
			Expect(contextBasedRestrictionsService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			contextBasedRestrictionsService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateZone - Create a zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateZone(createZoneOptions *CreateZoneOptions)`, func() {
			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			createZoneOptions := &contextbasedrestrictionsv1.CreateZoneOptions{
				Name:          core.StringPtr("SDK TEST - an example of zone"),
				AccountID:     core.StringPtr(testAccountID),
				Description:   core.StringPtr("SDK TEST - this is an example of zone"),
				Addresses:     []contextbasedrestrictionsv1.AddressIntf{addressModel},
				TransactionID: core.StringPtr("sdk-create-zone-" + uuid.New().String()),
			}

			outZone, response, err := contextBasedRestrictionsService.CreateZone(createZoneOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(outZone).ToNot(BeNil())
			zoneID = *outZone.ID
		})
	})

	Describe(`ListZones - List zones`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListZones(listZonesOptions *ListZonesOptions)`, func() {
			listZonesOptions := &contextbasedrestrictionsv1.ListZonesOptions{
				AccountID:     core.StringPtr(testAccountID),
				TransactionID: core.StringPtr("sdk-list-zones-" + uuid.New().String()),
			}

			outZonePage, response, err := contextBasedRestrictionsService.ListZones(listZonesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outZonePage).ToNot(BeNil())
		})
	})

	Describe(`GetZone - Get the specified zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetZone(getZoneOptions *GetZoneOptions)`, func() {
			getZoneOptions := &contextbasedrestrictionsv1.GetZoneOptions{
				ZoneID:        core.StringPtr(zoneID),
				TransactionID: core.StringPtr("sdk-get-zone-" + uuid.New().String()),
			}

			outZone, response, err := contextBasedRestrictionsService.GetZone(getZoneOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outZone).ToNot(BeNil())
			zoneRev = response.Headers.Get("Etag")
		})
	})

	Describe(`ReplaceZone - Update the specified zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceZone(replaceZoneOptions *ReplaceZoneOptions)`, func() {
			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			replaceZoneOptions := &contextbasedrestrictionsv1.ReplaceZoneOptions{
				ZoneID:        core.StringPtr(zoneID),
				IfMatch:       core.StringPtr(zoneRev),
				Name:          core.StringPtr("SDK TEST - an example of updated zone"),
				AccountID:     core.StringPtr(testAccountID),
				Description:   core.StringPtr("SDK TEST - this is an example of updated zone"),
				Addresses:     []contextbasedrestrictionsv1.AddressIntf{addressModel},
				TransactionID: core.StringPtr("sdk-replace-zone-" + uuid.New().String()),
			}

			outZone, response, err := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outZone).ToNot(BeNil())
		})
	})

	Describe(`ListAvailableServiceRefTargets - List available service reference targets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAvailableServiceRefTargets(listAvailableServiceRefTargetsOptions *ListAvailableServiceRefTargetsOptions)`, func() {
			listAvailableServiceRefTargetsOptions := &contextbasedrestrictionsv1.ListAvailableServiceRefTargetsOptions{
				Type: core.StringPtr(contextbasedrestrictionsv1.ListAvailableServiceRefTargetsOptionsTypeAll),
			}

			serviceRefTargetPage, response, err := contextBasedRestrictionsService.ListAvailableServiceRefTargets(listAvailableServiceRefTargetsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceRefTargetPage).ToNot(BeNil())
		})
	})

	Describe(`CreateRule - Create a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRule(createRuleOptions *CreateRuleOptions)`, func() {
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

			createRuleOptions := &contextbasedrestrictionsv1.CreateRuleOptions{
				Description:   core.StringPtr("SDK TEST - this is an example of rule"),
				Contexts:      []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:     []contextbasedrestrictionsv1.Resource{*resourceModel},
				TransactionID: core.StringPtr("sdk-create-rule-" + uuid.New().String()),
			}

			outRule, response, err := contextBasedRestrictionsService.CreateRule(createRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(outRule).ToNot(BeNil())
			ruleID = *outRule.ID
		})
	})

	Describe(`ListRules - List rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions)`, func() {
			listRulesOptions := &contextbasedrestrictionsv1.ListRulesOptions{
				AccountID:     core.StringPtr(testAccountID),
				TransactionID: core.StringPtr("sdk-list-rules-" + uuid.New().String()),
			}

			outRulePage, response, err := contextBasedRestrictionsService.ListRules(listRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outRulePage).ToNot(BeNil())
		})
	})

	Describe(`GetRule - Get the specified rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
			getRuleOptions := &contextbasedrestrictionsv1.GetRuleOptions{
				RuleID:        core.StringPtr(ruleID),
				TransactionID: core.StringPtr("sdk-get-rule-" + uuid.New().String()),
			}

			outRule, response, err := contextBasedRestrictionsService.GetRule(getRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outRule).ToNot(BeNil())
			ruleRev = response.Headers.Get("Etag")
		})
	})

	Describe(`ReplaceRule - Update the specified rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions)`, func() {
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

			replaceRuleOptions := &contextbasedrestrictionsv1.ReplaceRuleOptions{
				RuleID:        core.StringPtr(ruleID),
				IfMatch:       core.StringPtr(ruleRev),
				Description:   core.StringPtr("SDK TEST - this is an example of updated rule"),
				Contexts:      []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:     []contextbasedrestrictionsv1.Resource{*resourceModel},
				TransactionID: core.StringPtr("sdk-replace-rule-" + uuid.New().String()),
			}

			outRule, response, err := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outRule).ToNot(BeNil())
		})
	})

	Describe(`GetAccountSettings - Get the specified account settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions)`, func() {
			getAccountSettingsOptions := &contextbasedrestrictionsv1.GetAccountSettingsOptions{
				AccountID:     core.StringPtr(testAccountID),
				TransactionID: core.StringPtr("sdk-get-account-settings-" + uuid.New().String()),
			}

			outAccountSettings, response, err := contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outAccountSettings).ToNot(BeNil())
		})
	})

	//
	// Requests with errors
	//

	Describe(`CreateZone - Create a zone with 'invalid ip address format' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateZone(createZoneOptions *CreateZoneOptions) with 'invalid ip address format' error (400)`, func() {
			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234."),
			}

			createZoneOptions := &contextbasedrestrictionsv1.CreateZoneOptions{
				Name:          core.StringPtr("SDK TEST - an example of zone"),
				AccountID:     core.StringPtr(testAccountID),
				Description:   core.StringPtr("SDK TEST - this is an example of zone"),
				Addresses:     []contextbasedrestrictionsv1.AddressIntf{addressModel},
				TransactionID: core.StringPtr("sdk-create-zone-" + uuid.New().String()),
			}

			outZone, response, err := contextBasedRestrictionsService.CreateZone(createZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(outZone).To(BeNil())
		})
	})

	Describe(`ListZones - List zones with 'invalid account id parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListZones(listZonesOptions *ListZonesOptions) with 'invalid account id parameter' error (400)`, func() {
			listZonesOptions := &contextbasedrestrictionsv1.ListZonesOptions{
				AccountID:     core.StringPtr(InvalidID),
				TransactionID: core.StringPtr("sdk-list-zones-" + uuid.New().String()),
			}

			outZonePage, response, err := contextBasedRestrictionsService.ListZones(listZonesOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(outZonePage).To(BeNil())
		})
	})

	Describe(`GetZone - Get the specified zone with 'zone not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetZone(getZoneOptions *GetZoneOptions) with 'zone not found' error (404)`, func() {
			getZoneOptions := &contextbasedrestrictionsv1.GetZoneOptions{
				ZoneID:        core.StringPtr(NonExistentID),
				TransactionID: core.StringPtr("sdk-get-zone-" + uuid.New().String()),
			}

			outZone, response, err := contextBasedRestrictionsService.GetZone(getZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
			Expect(outZone).To(BeNil())

		})
	})

	Describe(`ReplaceZone - Update the specified zone with 'zone not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceZone(replaceZoneOptions *ReplaceZoneOptions) with 'zone not found' error (404)`, func() {
			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			replaceZoneOptions := &contextbasedrestrictionsv1.ReplaceZoneOptions{
				ZoneID:        core.StringPtr(NonExistentID),
				IfMatch:       core.StringPtr("abc"),
				Name:          core.StringPtr("SDK TEST - an example of zone"),
				AccountID:     core.StringPtr(testAccountID),
				Description:   core.StringPtr("SDK TEST - this is an example of zone"),
				Addresses:     []contextbasedrestrictionsv1.AddressIntf{addressModel},
				TransactionID: core.StringPtr("sdk-replace-zone-" + uuid.New().String()),
			}

			outZone, response, err := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
			Expect(outZone).To(BeNil())
		})
	})

	Describe(`ListAvailableServiceRefTargets - List available service reference targets with 'invalid type parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAvailableServiceRefTargets(listAvailableServiceRefTargetsOptions *ListAvailableServiceRefTargetsOptions) with 'invalid type parameter' error (400)`, func() {
			listAvailableServiceRefTargetsOptions := &contextbasedrestrictionsv1.ListAvailableServiceRefTargetsOptions{
				Type: core.StringPtr("invalid-type"),
			}

			serviceRefTargetPage, response, err := contextBasedRestrictionsService.ListAvailableServiceRefTargets(listAvailableServiceRefTargetsOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(serviceRefTargetPage).To(BeNil())
		})
	})

	Describe(`CreateRule - Create a rule with 'service not cbr enabled' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRule(createRuleOptions *CreateRuleOptions) with 'service not cbr enabled' error (400)`, func() {
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
						Value: core.StringPtr("cbr-not-enabled"),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("tagValue"),
					},
				},
			}

			createRuleOptions := &contextbasedrestrictionsv1.CreateRuleOptions{
				Description:   core.StringPtr("SDK TEST - this is an example of rule"),
				Contexts:      []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:     []contextbasedrestrictionsv1.Resource{*resourceModel},
				TransactionID: core.StringPtr("sdk-create-rule-" + uuid.New().String()),
			}

			outRule, response, err := contextBasedRestrictionsService.CreateRule(createRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(outRule).To(BeNil())
		})
	})

	Describe(`ListRules - List rules with 'invalid account id parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions) with 'invalid account id parameter' error (400)`, func() {
			listRulesOptions := &contextbasedrestrictionsv1.ListRulesOptions{
				AccountID:     core.StringPtr(InvalidID),
				TransactionID: core.StringPtr("sdk-list-rules-" + uuid.New().String()),
			}

			outRulePage, response, err := contextBasedRestrictionsService.ListRules(listRulesOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(outRulePage).To(BeNil())
		})
	})

	Describe(`GetRule - Get the specified rule with error with 'rule not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions) with 'rule not found' error (404)`, func() {
			getRuleOptions := &contextbasedrestrictionsv1.GetRuleOptions{
				RuleID:        core.StringPtr(NonExistentID),
				TransactionID: core.StringPtr("sdk-get-rule-" + uuid.New().String()),
			}

			outRule, response, err := contextBasedRestrictionsService.GetRule(getRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
			Expect(outRule).To(BeNil())
		})
	})

	Describe(`ReplaceRule - Update the specified rule with error with 'rule not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions) with 'rule not found' error (404)`, func() {
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
						Value: core.StringPtr("cbr-not-enabled"),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("updatedTagValue"),
					},
				},
			}

			replaceRuleOptions := &contextbasedrestrictionsv1.ReplaceRuleOptions{
				RuleID:        core.StringPtr(NonExistentID),
				IfMatch:       core.StringPtr("abc"),
				Description:   core.StringPtr("SDK TEST - this is an example of rule"),
				Contexts:      []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:     []contextbasedrestrictionsv1.Resource{*resourceModel},
				TransactionID: core.StringPtr("sdk-replace-rule-" + uuid.New().String()),
			}

			outRule, response, err := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
			Expect(outRule).To(BeNil())
		})
	})

	Describe(`GetAccountSettings - Get the specified account settings with 'invalid account id parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) with 'invalid account id parameter' error (400)`, func() {
			getAccountSettingsOptions := &contextbasedrestrictionsv1.GetAccountSettingsOptions{
				AccountID:     core.StringPtr(InvalidID),
				TransactionID: core.StringPtr("sdk-get-account-settings-" + uuid.New().String()),
			}

			outAccountSettings, response, err := contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(outAccountSettings).To(BeNil())
		})
	})

	Describe(`DeleteRule - Delete the specified rule with 'rule not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions) with 'rule not found' error (404)`, func() {
			deleteRuleOptions := &contextbasedrestrictionsv1.DeleteRuleOptions{
				RuleID:        core.StringPtr(NonExistentID),
				TransactionID: core.StringPtr("sdk-delete-rule-" + uuid.New().String()),
			}

			response, err := contextBasedRestrictionsService.DeleteRule(deleteRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`DeleteZone - Delete the specified zone with 'zone not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteZone(deleteZoneOptions *DeleteZoneOptions) with 'zone not found' error (404)`, func() {
			deleteZoneOptions := &contextbasedrestrictionsv1.DeleteZoneOptions{
				ZoneID:        core.StringPtr(NonExistentID),
				TransactionID: core.StringPtr("sdk-delete-zone-" + uuid.New().String()),
			}

			response, err := contextBasedRestrictionsService.DeleteZone(deleteZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	//
	// Cleanup the created zones and rules
	//

	Describe(`DeleteRule - Delete the specified rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
			deleteRuleOptions := &contextbasedrestrictionsv1.DeleteRuleOptions{
				RuleID:        core.StringPtr(ruleID),
				TransactionID: core.StringPtr("sdk-delete-rule-" + uuid.New().String()),
			}

			response, err := contextBasedRestrictionsService.DeleteRule(deleteRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteZone - Delete the specified zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteZone(deleteZoneOptions *DeleteZoneOptions)`, func() {
			deleteZoneOptions := &contextbasedrestrictionsv1.DeleteZoneOptions{
				ZoneID:        core.StringPtr(zoneID),
				TransactionID: core.StringPtr("sdk-delete-zone-" + uuid.New().String()),
			}

			response, err := contextBasedRestrictionsService.DeleteZone(deleteZoneOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
