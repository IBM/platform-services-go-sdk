// +build examples

/**
 * (C) Copyright IBM Corp. 2020, 2021.
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

package configurationgovernancev1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/configurationgovernancev1"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Configuration Governance service.
//
// The following configuration properties are assumed to be defined:
//
// CONFIGURATION_GOVERNANCE_URL=<service url>
// CONFIGURATION_GOVERNANCE_AUTHTYPE=iam
// CONFIGURATION_GOVERNANCE_APIKEY=<IAM api key of user with authority to create rules>
// CONFIGURATION_GOVERNANCE_AUTH_URL=<IAM token service URL - omit this if using the production environment>
// CONFIGURATION_GOVERNANCE_ACCOUNT_ID=<the id of the account under which rules/attachments should be created>
// CONFIGURATION_GOVERNANCE_EXAMPLE_SERVICE_NAME=<the name of the service to be associated with rule>
// CONFIGURATION_GOVERNANCE_ENTERPRISE_SCOPE_ID=<the id of the "enterprise" scope to be used in the examples>
// CONFIGURATION_GOVERNANCE_SUBACCT_SCOPE_ID=<the id of the "leaf account" scope to be used in the examples>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../configuration_governance.env"

var (
	configurationGovernanceService *configurationgovernancev1.ConfigurationGovernanceV1
	config                         map[string]string
	configLoaded                   bool = false

	// Test-related configuration properties.
	accountID         string
	serviceName       string
	enterpriseScopeID string
	subacctScopeID    string

	transactionID string = uuid.New().String()
)

// Global variables to hold various values shared between operations
var (
	ruleIDLink           string
	ruleToUpdateLink     *configurationgovernancev1.Rule
	ruleToUpdateEtagLink string

	attachmentIDLink           string
	attachmentToUpdateLink     *configurationgovernancev1.Attachment
	attachmentToUpdateEtagLink string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`ConfigurationGovernanceV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(configurationgovernancev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0

			if configLoaded {
				// Retrieve test-related properties from the external configuration.
				accountID = config["ACCOUNT_ID"]
				serviceName = config["EXAMPLE_SERVICE_NAME"]
				enterpriseScopeID = config["ENTERPRISE_SCOPE_ID"]
				subacctScopeID = config["SUBACCT_SCOPE_ID"]
			}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			configurationGovernanceServiceOptions := &configurationgovernancev1.ConfigurationGovernanceV1Options{}

			configurationGovernanceService, err =
				configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(
					configurationGovernanceServiceOptions,
				)
			if err != nil {
				panic(err)
			}

			// end-common

			Expect(configurationGovernanceService).ToNot(BeNil())
		})
	})

	Describe(`ConfigurationGovernanceV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRules request example`, func() {
			// begin-create_rules

			ruleRequestModel := &configurationgovernancev1.RuleRequest{
				AccountID:   &accountID,
				Name:        core.StringPtr("Disable public access"),
				Description: core.StringPtr("Ensure that public access to account resources is disabled."),
				Target: &configurationgovernancev1.TargetResource{
					ServiceName:  &serviceName,
					ResourceKind: core.StringPtr("service"),
				},
				RequiredConfig: &configurationgovernancev1.RuleRequiredConfigSingleProperty{
					Description: core.StringPtr("Ensure public access is disabled."),
					Property:    core.StringPtr("public_access_enabled"),
					Operator:    core.StringPtr("is_false"),
				},
				EnforcementActions: []configurationgovernancev1.EnforcementAction{
					configurationgovernancev1.EnforcementAction{
						Action: core.StringPtr("audit_log"),
					},
					configurationgovernancev1.EnforcementAction{
						Action: core.StringPtr("disallow"),
					},
				},
				Labels: []string{"test_label"},
			}

			createRulesOptions := configurationGovernanceService.NewCreateRulesOptions(
				[]configurationgovernancev1.CreateRuleRequest{
					configurationgovernancev1.CreateRuleRequest{
						Rule: ruleRequestModel,
					},
				},
			)
			createRulesOptions.SetTransactionID(uuid.New().String())

			createRulesResponse, response, err := configurationGovernanceService.CreateRules(createRulesOptions)
			// For a 207 status code, check the response entry for an error
			if response.StatusCode == 207 {
				for _, responseEntry := range createRulesResponse.Rules {
					if *responseEntry.StatusCode > int64(299) {
						err = fmt.Errorf("%s: %s", *responseEntry.Errors[0].Code, *responseEntry.Errors[0].Message)
					}
				}
			}
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createRulesResponse, "", "  ")
			fmt.Printf("\nCreateRules() result:\n%s\n", string(b))

			// end-create_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createRulesResponse).ToNot(BeNil())
			Expect(createRulesResponse.Rules).ToNot(BeEmpty())
			Expect(createRulesResponse.Rules[0]).ToNot(BeNil())
			Expect(createRulesResponse.Rules[0].Rule).ToNot(BeNil())
			Expect(createRulesResponse.Rules[0].Rule.RuleID).ToNot(BeNil())

			ruleIDLink = *createRulesResponse.Rules[0].Rule.RuleID
			Expect(ruleIDLink).ToNot(BeEmpty())
		})
		It(`ListRules request example`, func() {
			// begin-list_rules

			listRulesOptions := configurationGovernanceService.NewListRulesOptions(accountID)
			listRulesOptions.SetLabels("test_label")
			listRulesOptions.SetTransactionID(uuid.New().String())

			ruleList, response, err := configurationGovernanceService.ListRules(listRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ruleList, "", "  ")
			fmt.Printf("\nListRules() result:\n%s\n", string(b))

			// end-list_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleList).ToNot(BeNil())

		})
		It(`GetRule request example`, func() {
			Expect(ruleIDLink).ToNot(BeEmpty())

			// begin-get_rule

			getRuleOptions := configurationGovernanceService.NewGetRuleOptions(ruleIDLink)
			getRuleOptions.SetTransactionID(uuid.New().String())

			rule, response, err := configurationGovernanceService.GetRule(getRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Printf("\nGetRule() result:\n%s\n", string(b))

			// end-get_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

			ruleToUpdateLink = rule

			// Retrieve the Etag header from the response for use in the UpdateRule() operation.
			ruleToUpdateEtagLink = response.GetHeaders().Get("Etag")
			Expect(ruleToUpdateEtagLink).ToNot(BeEmpty())

		})
		It(`UpdateRule request example`, func() {
			Expect(ruleIDLink).ToNot(BeEmpty())
			Expect(ruleToUpdateLink).ToNot(BeNil())
			Expect(ruleToUpdateEtagLink).ToNot(BeEmpty())

			// begin-update_rule

			// Update the existing rule's description.
			updateRuleOptions := configurationGovernanceService.NewUpdateRuleOptions(
				ruleIDLink,
				ruleToUpdateEtagLink,
				*ruleToUpdateLink.Name,
				"This is an updated description.",
				ruleToUpdateLink.Target,
				ruleToUpdateLink.RequiredConfig,
				ruleToUpdateLink.EnforcementActions,
			)
			updateRuleOptions.SetAccountID(*ruleToUpdateLink.AccountID)
			updateRuleOptions.SetRuleType(*ruleToUpdateLink.RuleType)
			updateRuleOptions.SetLabels(ruleToUpdateLink.Labels)
			updateRuleOptions.SetTransactionID(uuid.New().String())

			rule, response, err := configurationGovernanceService.UpdateRule(updateRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Printf("\nUpdateRule() result:\n%s\n", string(b))

			// end-update_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
		It(`CreateAttachments request example`, func() {
			Expect(ruleIDLink).ToNot(BeEmpty())

			// begin-create_attachments

			createAttachmentRequest := configurationgovernancev1.AttachmentRequest{
				AccountID: &accountID,
				IncludedScope: &configurationgovernancev1.RuleScope{
					Note:      core.StringPtr("My enterprise"),
					ScopeID:   &enterpriseScopeID,
					ScopeType: core.StringPtr("enterprise"),
				},
				ExcludedScopes: []configurationgovernancev1.RuleScope{
					configurationgovernancev1.RuleScope{
						Note:      core.StringPtr("leaf account"),
						ScopeID:   &subacctScopeID,
						ScopeType: core.StringPtr("enterprise.account"),
					},
				},
			}

			createAttachmentsOptions := configurationGovernanceService.NewCreateAttachmentsOptions(
				ruleIDLink,
				[]configurationgovernancev1.AttachmentRequest{
					createAttachmentRequest,
				},
			)
			createAttachmentsOptions.SetTransactionID(uuid.New().String())

			createAttachmentsResponse, response, err := configurationGovernanceService.CreateAttachments(createAttachmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createAttachmentsResponse, "", "  ")
			fmt.Printf("\nCreateAttachments() result:\n%s\n", string(b))

			// end-create_attachments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createAttachmentsResponse).ToNot(BeNil())
			Expect(createAttachmentsResponse.Attachments).ToNot(BeEmpty())
			Expect(createAttachmentsResponse.Attachments[0]).ToNot(BeNil())
			Expect(createAttachmentsResponse.Attachments[0].AttachmentID).ToNot(BeNil())

			attachmentIDLink = *createAttachmentsResponse.Attachments[0].AttachmentID
			Expect(attachmentIDLink).ToNot(BeEmpty())

		})
		It(`ListAttachments request example`, func() {
			Expect(ruleIDLink).ToNot(BeEmpty())

			// begin-list_attachments

			listAttachmentsOptions := configurationGovernanceService.NewListAttachmentsOptions(ruleIDLink)
			listAttachmentsOptions.SetTransactionID(uuid.New().String())

			attachmentList, response, err := configurationGovernanceService.ListAttachments(listAttachmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentList, "", "  ")
			fmt.Printf("\nListAttachments() result:\n%s\n", string(b))

			// end-list_attachments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentList).ToNot(BeNil())
		})
		It(`GetAttachment request example`, func() {
			Expect(ruleIDLink).ToNot(BeEmpty())
			Expect(attachmentIDLink).ToNot(BeEmpty())

			// begin-get_attachment

			getAttachmentOptions := configurationGovernanceService.NewGetAttachmentOptions(
				ruleIDLink,
				attachmentIDLink,
			)
			getAttachmentOptions.SetTransactionID(uuid.New().String())

			attachment, response, err := configurationGovernanceService.GetAttachment(getAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachment, "", "  ")
			fmt.Printf("\nGetAttachment() result:\n%s\n", string(b))

			// end-get_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())

			attachmentToUpdateLink = attachment

			// Retrieve the Etag header from the response for use in the update operation.
			attachmentToUpdateEtagLink = response.GetHeaders().Get("Etag")
			Expect(attachmentToUpdateEtagLink).ToNot(BeEmpty())
		})
		It(`UpdateAttachment request example`, func() {
			Expect(ruleIDLink).ToNot(BeEmpty())
			Expect(attachmentIDLink).ToNot(BeEmpty())
			Expect(attachmentToUpdateLink).ToNot(BeNil())
			Expect(attachmentToUpdateEtagLink).ToNot(BeEmpty())

			// begin-update_attachment

			// Update the Note field within the existing attachment's IncludedScope.
			updatedIncludedScope := &configurationgovernancev1.RuleScope{
				Note:      core.StringPtr("This is a new note."),
				ScopeID:   attachmentToUpdateLink.IncludedScope.ScopeID,
				ScopeType: attachmentToUpdateLink.IncludedScope.ScopeType,
			}

			updateAttachmentOptions := configurationGovernanceService.NewUpdateAttachmentOptions(
				ruleIDLink,
				attachmentIDLink,
				attachmentToUpdateEtagLink,
				*attachmentToUpdateLink.AccountID,
				updatedIncludedScope,
			)
			updateAttachmentOptions.SetExcludedScopes(attachmentToUpdateLink.ExcludedScopes)
			updateAttachmentOptions.SetTransactionID(uuid.New().String())

			attachment, response, err := configurationGovernanceService.UpdateAttachment(updateAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachment, "", "  ")
			fmt.Printf("UpdateAttachment() result:\n%s\n", string(b))

			// end-update_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())
		})
		It(`DeleteAttachment request example`, func() {
			Expect(ruleIDLink).ToNot(BeEmpty())
			Expect(attachmentIDLink).ToNot(BeEmpty())

			// begin-delete_attachment

			deleteAttachmentOptions := configurationGovernanceService.NewDeleteAttachmentOptions(
				ruleIDLink,
				attachmentIDLink,
			)

			response, err := configurationGovernanceService.DeleteAttachment(deleteAttachmentOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("DeleteAttachment() response status code: %d\n", response.StatusCode)

			// end-delete_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteRule request example`, func() {
			Expect(ruleIDLink).ToNot(BeEmpty())

			// begin-delete_rule

			deleteRuleOptions := configurationGovernanceService.NewDeleteRuleOptions(ruleIDLink)

			response, err := configurationGovernanceService.DeleteRule(deleteRuleOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("DeleteRule() response status code: %d\n", response.StatusCode)

			// end-delete_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
