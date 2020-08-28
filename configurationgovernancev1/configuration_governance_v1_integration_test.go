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

package configurationgovernancev1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/configurationgovernancev1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

/**
 * This file contains an integration test for the configurationgovernancev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`ConfigurationGovernanceV1 Integration Tests`, func() {

	const externalConfigFile = "../configuration_governance_v1.env"

	var (
		err          error
		configurationGovernanceService *configurationgovernancev1.ConfigurationGovernanceV1
		serviceURL   string
		config       map[string]string
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
			config, err = core.GetServiceProperties(configurationgovernancev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			configurationGovernanceServiceOptions := &configurationgovernancev1.ConfigurationGovernanceV1Options{}

			configurationGovernanceService, err = configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(configurationGovernanceServiceOptions)

			Expect(err).To(BeNil())
			Expect(configurationGovernanceService).ToNot(BeNil())
			Expect(configurationGovernanceService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`CreateRules - Create rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRules(createRulesOptions *CreateRulesOptions)`, func() {

			uiSupportModel := &configurationgovernancev1.UiSupport{
				DisplayName: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			ruleImportModel := &configurationgovernancev1.RuleImport{
				Name: core.StringPtr("testString"),
				UiSupport: uiSupportModel,
			}

			ruleTargetAttributeModel := &configurationgovernancev1.RuleTargetAttribute{
				Name: core.StringPtr("resource_id"),
				Operator: core.StringPtr("string_equals"),
				Value: core.StringPtr("f0f8f7994e754ff38f9d370201966561"),
			}

			targetResourceModel := &configurationgovernancev1.TargetResource{
				ServiceName: core.StringPtr("iam-groups"),
				ResourceKind: core.StringPtr("zone"),
				AdditionalTargetAttributes: []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel},
			}

			ruleRequiredConfigModel := &configurationgovernancev1.RuleRequiredConfigSingleProperty{
				Description: core.StringPtr("testString"),
				Property: core.StringPtr("public_access_enabled"),
				Operator: core.StringPtr("is_true"),
				Value: core.StringPtr("testString"),
			}

			enforcementActionModel := &configurationgovernancev1.EnforcementAction{
				Action: core.StringPtr("disallow"),
			}

			ruleRequestModel := &configurationgovernancev1.RuleRequest{
				AccountID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Version: core.StringPtr("1.0.0"),
				RuleType: core.StringPtr("user_defined"),
				Imports: []configurationgovernancev1.RuleImport{*ruleImportModel},
				Target: targetResourceModel,
				RequiredConfig: ruleRequiredConfigModel,
				EnforcementActions: []configurationgovernancev1.EnforcementAction{*enforcementActionModel},
				Labels: []string{"testString"},
			}

			createRuleRequestModel := &configurationgovernancev1.CreateRuleRequest{
				RequestID: core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558"),
				Rule: ruleRequestModel,
			}

			createRulesOptions := &configurationgovernancev1.CreateRulesOptions{
				Rules: []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel},
				TransactionID: core.StringPtr("testString"),
			}

			createRulesResponse, response, err := configurationGovernanceService.CreateRules(createRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createRulesResponse).ToNot(BeNil())

		})
	})

	Describe(`ListRules - List rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions)`, func() {

			listRulesOptions := &configurationgovernancev1.ListRulesOptions{
				AccountID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				Attached: core.BoolPtr(true),
				Labels: core.StringPtr("SOC2,ITCS300"),
				Scopes: core.StringPtr("scope_id"),
				Limit: core.Int64Ptr(int64(1000)),
				Offset: core.Int64Ptr(int64(38)),
			}

			ruleList, response, err := configurationGovernanceService.ListRules(listRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleList).ToNot(BeNil())

		})
	})

	Describe(`GetRule - Get a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions)`, func() {

			getRuleOptions := &configurationgovernancev1.GetRuleOptions{
				RuleID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			rule, response, err := configurationGovernanceService.GetRule(getRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
	})

	Describe(`UpdateRule - Update a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateRule(updateRuleOptions *UpdateRuleOptions)`, func() {

			ruleTargetAttributeModel := &configurationgovernancev1.RuleTargetAttribute{
				Name: core.StringPtr("testString"),
				Operator: core.StringPtr("string_equals"),
				Value: core.StringPtr("testString"),
			}

			targetResourceModel := &configurationgovernancev1.TargetResource{
				ServiceName: core.StringPtr("iam-groups"),
				ResourceKind: core.StringPtr("service"),
				AdditionalTargetAttributes: []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel},
			}

			ruleRequiredConfigModel := &configurationgovernancev1.RuleRequiredConfigSingleProperty{
				Description: core.StringPtr("testString"),
				Property: core.StringPtr("public_access_enabled"),
				Operator: core.StringPtr("is_false"),
				Value: core.StringPtr("testString"),
			}

			enforcementActionModel := &configurationgovernancev1.EnforcementAction{
				Action: core.StringPtr("audit_log"),
			}

			uiSupportModel := &configurationgovernancev1.UiSupport{
				DisplayName: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			ruleImportModel := &configurationgovernancev1.RuleImport{
				Name: core.StringPtr("testString"),
				UiSupport: uiSupportModel,
			}

			updateRuleOptions := &configurationgovernancev1.UpdateRuleOptions{
				RuleID: core.StringPtr("testString"),
				IfMatch: core.StringPtr("testString"),
				Name: core.StringPtr("Disable public access"),
				Description: core.StringPtr("Ensure that public access to account resources is disabled."),
				Target: targetResourceModel,
				RequiredConfig: ruleRequiredConfigModel,
				EnforcementActions: []configurationgovernancev1.EnforcementAction{*enforcementActionModel},
				AccountID: core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
				Version: core.StringPtr("1.0.0"),
				RuleType: core.StringPtr("user_defined"),
				Imports: []configurationgovernancev1.RuleImport{*ruleImportModel},
				Labels: []string{"testString"},
				TransactionID: core.StringPtr("testString"),
			}

			rule, response, err := configurationGovernanceService.UpdateRule(updateRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
	})

	Describe(`CreateAttachments - Create attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAttachments(createAttachmentsOptions *CreateAttachmentsOptions)`, func() {

			ruleScopeModel := &configurationgovernancev1.RuleScope{
				Note: core.StringPtr("My enterprise"),
				ScopeID: core.StringPtr("282cf433ac91493ba860480d92519990"),
				ScopeType: core.StringPtr("enterprise"),
			}

			attachmentRequestModel := &configurationgovernancev1.AttachmentRequest{
				AccountID: core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
				IncludedScope: ruleScopeModel,
				ExcludedScopes: []configurationgovernancev1.RuleScope{*ruleScopeModel},
			}

			createAttachmentsOptions := &configurationgovernancev1.CreateAttachmentsOptions{
				RuleID: core.StringPtr("testString"),
				Attachments: []configurationgovernancev1.AttachmentRequest{*attachmentRequestModel},
				TransactionID: core.StringPtr("testString"),
			}

			createAttachmentsResponse, response, err := configurationGovernanceService.CreateAttachments(createAttachmentsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createAttachmentsResponse).ToNot(BeNil())

		})
	})

	Describe(`ListAttachments - List attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAttachments(listAttachmentsOptions *ListAttachmentsOptions)`, func() {

			listAttachmentsOptions := &configurationgovernancev1.ListAttachmentsOptions{
				RuleID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1000)),
				Offset: core.Int64Ptr(int64(38)),
			}

			attachmentList, response, err := configurationGovernanceService.ListAttachments(listAttachmentsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentList).ToNot(BeNil())

		})
	})

	Describe(`GetAttachment - Get an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAttachment(getAttachmentOptions *GetAttachmentOptions)`, func() {

			getAttachmentOptions := &configurationgovernancev1.GetAttachmentOptions{
				RuleID: core.StringPtr("testString"),
				AttachmentID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			attachment, response, err := configurationGovernanceService.GetAttachment(getAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())

		})
	})

	Describe(`UpdateAttachment - Update an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAttachment(updateAttachmentOptions *UpdateAttachmentOptions)`, func() {

			ruleScopeModel := &configurationgovernancev1.RuleScope{
				Note: core.StringPtr("testString"),
				ScopeID: core.StringPtr("testString"),
				ScopeType: core.StringPtr("enterprise"),
			}

			updateAttachmentOptions := &configurationgovernancev1.UpdateAttachmentOptions{
				RuleID: core.StringPtr("testString"),
				AttachmentID: core.StringPtr("testString"),
				IfMatch: core.StringPtr("testString"),
				AccountID: core.StringPtr("testString"),
				IncludedScope: ruleScopeModel,
				ExcludedScopes: []configurationgovernancev1.RuleScope{*ruleScopeModel},
				TransactionID: core.StringPtr("testString"),
			}

			attachment, response, err := configurationGovernanceService.UpdateAttachment(updateAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())

		})
	})

	Describe(`DeleteRule - Delete a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {

			deleteRuleOptions := &configurationgovernancev1.DeleteRuleOptions{
				RuleID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			response, err := configurationGovernanceService.DeleteRule(deleteRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteAttachment - Delete an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAttachment(deleteAttachmentOptions *DeleteAttachmentOptions)`, func() {

			deleteAttachmentOptions := &configurationgovernancev1.DeleteAttachmentOptions{
				RuleID: core.StringPtr("testString"),
				AttachmentID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			response, err := configurationGovernanceService.DeleteAttachment(deleteAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
