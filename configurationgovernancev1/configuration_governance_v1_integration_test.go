// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	. "github.com/IBM/platform-services-go-sdk/configurationgovernancev1"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the configurationgovernancev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

const verbose bool = true

var transactionID string

var _ = Describe(`ConfigurationGovernanceV1 Integration Tests`, func() {

	const externalConfigFile = "../configuration_governance.env"
	const TestLabel = "GoSDKIntegrationTest"

	var (
		err             error
		service         *ConfigurationGovernanceV1
		serviceNoAccess *ConfigurationGovernanceV1
		serviceURL      string
		config          map[string]string

		// Test-related config properties.
		accountID         string
		testServiceName   string
		enterpriseScopeID string
		subacctScopeID    string

		// Sample rules and attachments.
		sampleRule1   *RuleRequest
		sampleRule2   *RuleRequest
		badSampleRule *RuleRequest

		enterpriseScope *RuleScope
		accountScope    *RuleScope
		badScope        *RuleScope

		// Variables to hold various id's and object instances (these could perhaps be configured via links).
		ruleID1   string
		rule1     *Rule
		ruleEtag1 string
		ruleID2   string

		attachmentID1   string
		attachment1     *Attachment
		attachmentEtag1 string
		attachmentID2   string
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
			config, err = core.GetServiceProperties(DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			// Retrieve test-related config properties.
			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			testServiceName = config["TEST_SERVICE_NAME"]
			Expect(testServiceName).ToNot(BeEmpty())

			enterpriseScopeID = config["ENTERPRISE_SCOPE_ID"]
			Expect(enterpriseScopeID).ToNot(BeEmpty())

			subacctScopeID = config["SUBACCT_SCOPE_ID"]
			Expect(subacctScopeID).ToNot(BeEmpty())

			transactionID = uuid.New().String()

			fmt.Fprintf(GinkgoWriter, "\nService URL: %s", serviceURL)
			fmt.Fprintf(GinkgoWriter, "TransactionID: %s", transactionID)

			shouldSkipTest = func() {}

		})
	})

	Describe(`Initialize sample data`, func() {
		It("Successfully construct model instances", func() {
			// Initialize some structs to serve as sample rules and attachments.

			ruleTargetAttributeModel := &RuleTargetAttribute{
				Name:     core.StringPtr("resource_id"),
				Operator: core.StringPtr("is_not_empty"),
			}
			targetResourceModel := &TargetResource{
				ServiceName:                &testServiceName,
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []RuleTargetAttribute{*ruleTargetAttributeModel},
			}
			allowedGBCondition := &RuleConditionSingleProperty{
				Property: core.StringPtr("allowed_gb"),
				Operator: core.StringPtr("num_less_than_equals"),
				Value:    core.StringPtr("20"),
			}
			locationCondition := &RuleConditionSingleProperty{
				Property: core.StringPtr("location"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("us-east"),
			}
			ruleRequiredConfigModel1 := &RuleRequiredConfigMultiplePropertiesConditionAnd{
				Description: core.StringPtr("allowed_gb<=20 && location=='us-east'"),
				And:         []RuleConditionIntf{allowedGBCondition, locationCondition},
			}
			ruleRequiredConfigModel2 := &RuleRequiredConfigSingleProperty{
				Description: core.StringPtr("allowed_gb <= 30"),
				Property:    core.StringPtr("allowed_gb"),
				Operator:    core.StringPtr("num_less_than_equals"),
				Value:       core.StringPtr("30"),
			}
			enforcementActionModel := &EnforcementAction{
				Action: core.StringPtr("disallow"),
			}

			// Sample rules.
			sampleRule1 = &RuleRequest{
				AccountID:          &accountID,
				Name:               core.StringPtr("Go Test Rule #1"),
				Description:        core.StringPtr("This is the description for Go Test Rule #1."),
				RuleType:           core.StringPtr("user_defined"),
				Target:             targetResourceModel,
				RequiredConfig:     ruleRequiredConfigModel1,
				EnforcementActions: []EnforcementAction{*enforcementActionModel},
				Labels:             []string{TestLabel},
			}
			sampleRule2 = &RuleRequest{
				AccountID:          &accountID,
				Name:               core.StringPtr("Go Test Rule #2"),
				Description:        core.StringPtr("This is the description for Go Test Rule #2."),
				RuleType:           core.StringPtr("user_defined"),
				Target:             targetResourceModel,
				RequiredConfig:     ruleRequiredConfigModel2,
				EnforcementActions: []EnforcementAction{*enforcementActionModel},
				Labels:             []string{TestLabel},
			}
			badSampleRule = &RuleRequest{
				AccountID:          &accountID,
				Name:               core.StringPtr("Go Test Rule #3"),
				Description:        core.StringPtr("This is the description for Go Test Rule #3."),
				RuleType:           core.StringPtr("service_defined"),
				Target:             targetResourceModel,
				RequiredConfig:     ruleRequiredConfigModel2,
				EnforcementActions: []EnforcementAction{*enforcementActionModel},
				Labels:             []string{TestLabel},
			}

			// Sample rule scopes.
			enterpriseScope = &RuleScope{
				Note:      core.StringPtr("enterprise"),
				ScopeID:   core.StringPtr(enterpriseScopeID),
				ScopeType: core.StringPtr("enterprise"),
			}
			accountScope = &RuleScope{
				Note:      core.StringPtr("leaf account"),
				ScopeID:   core.StringPtr(subacctScopeID),
				ScopeType: core.StringPtr("enterprise.account"),
			}
			badScope = &RuleScope{
				Note:      core.StringPtr("leaf account"),
				ScopeID:   core.StringPtr(subacctScopeID),
				ScopeType: core.StringPtr("enterprise.BOGUS"),
			}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			configurationGovernanceServiceOptions := &ConfigurationGovernanceV1Options{}

			service, err = NewConfigurationGovernanceV1UsingExternalConfig(configurationGovernanceServiceOptions)

			Expect(err).To(BeNil())
			Expect(service).ToNot(BeNil())
			Expect(service.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			service.EnableRetries(4, 30*time.Second)
		})
		It("Successfully construct addition 'no-access' service client instance", func() {

			configurationGovernanceServiceOptions := &ConfigurationGovernanceV1Options{
				ServiceName: "NO_ACCESS",
			}

			serviceNoAccess, err = NewConfigurationGovernanceV1UsingExternalConfig(configurationGovernanceServiceOptions)

			Expect(err).To(BeNil())
			Expect(serviceNoAccess).ToNot(BeNil())
			Expect(serviceNoAccess.Service.Options.URL).To(Equal(serviceURL))
		})

		It("Successfully setup the environment for tests", func() {
			fmt.Fprintln(GinkgoWriter, "Setup...")
			cleanRules(service, accountID, TestLabel)
			fmt.Fprintln(GinkgoWriter, "Finished setup.")
		})
	})

	Describe(`CreateRules - Create rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully create rule #1`, func() {

			createRuleRequestModel := &CreateRuleRequest{
				RequestID: core.StringPtr("request-0"),
				Rule:      sampleRule1,
			}

			createRulesOptions := &CreateRulesOptions{
				Rules:         []CreateRuleRequest{*createRuleRequestModel},
				TransactionID: &transactionID,
			}

			createRulesResponse, response, err := service.CreateRules(createRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createRulesResponse).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "\nReceived response:\n%s", common.ToJSON(createRulesResponse))
			Expect(len(createRulesResponse.Rules)).To(Equal(1))
			ruleResponse1 := createRulesResponse.Rules[0]
			Expect(ruleResponse1).ToNot(BeNil())
			Expect(*ruleResponse1.RequestID).To(Equal("request-0"))
			Expect(*ruleResponse1.StatusCode).To(Equal(int64(201)))
			Expect(ruleResponse1.Rule).ToNot(BeNil())

			ruleID1 = *ruleResponse1.Rule.RuleID
			Expect(ruleID1).ToNot(BeEmpty())

		})
		It(`Successfully create rule #2`, func() {

			createRuleRequestModel := &CreateRuleRequest{
				Rule: sampleRule2,
			}

			createRulesOptions := &CreateRulesOptions{
				Rules:         []CreateRuleRequest{*createRuleRequestModel},
				TransactionID: &transactionID,
			}

			createRulesResponse, response, err := service.CreateRules(createRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createRulesResponse).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nReceived response:\n%s", common.ToJSON(createRulesResponse))

			Expect(len(createRulesResponse.Rules)).To(Equal(1))
			ruleResponse1 := createRulesResponse.Rules[0]
			Expect(ruleResponse1).ToNot(BeNil())
			Expect(*ruleResponse1.RequestID).ToNot(BeEmpty())
			Expect(*ruleResponse1.StatusCode).To(Equal(int64(201)))
			Expect(ruleResponse1.Rule).ToNot(BeNil())

			ruleID2 = *ruleResponse1.Rule.RuleID
			Expect(ruleID2).ToNot(BeEmpty())
		})
		It(`Fail to create invalid rule`, func() {

			createRuleRequestModel := &CreateRuleRequest{
				Rule: badSampleRule,
			}

			createRulesOptions := &CreateRulesOptions{
				Rules:         []CreateRuleRequest{*createRuleRequestModel},
				TransactionID: &transactionID,
			}

			// An error will be reported within the CreateRuleResponse entry, but the operation itself will
			// return a 207 status code.
			createRulesResponse, response, err := service.CreateRules(createRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(207))
			Expect(createRulesResponse).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "\nReceived response:\n%s", common.ToJSON(createRulesResponse))
			Expect(len(createRulesResponse.Rules)).To(Equal(1))
			ruleResponse1 := createRulesResponse.Rules[0]
			Expect(ruleResponse1).ToNot(BeNil())
			Expect(*ruleResponse1.RequestID).ToNot(BeEmpty())
			Expect(*ruleResponse1.StatusCode).To(Equal(int64(400)))
			Expect(ruleResponse1.Rule).To(BeNil())
			Expect(len(ruleResponse1.Errors)).To(Equal(1))
			error1 := ruleResponse1.Errors[0]
			Expect(*error1.Code).To(Equal("rule_error"))
		})
		It(`Fail to create rule with unauthorized user`, func() {

			createRuleRequestModel := &CreateRuleRequest{
				Rule: sampleRule1,
			}

			createRulesOptions := &CreateRulesOptions{
				Rules:         []CreateRuleRequest{*createRuleRequestModel},
				TransactionID: &transactionID,
			}

			_, response, err := serviceNoAccess.CreateRules(createRulesOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(403))
			fmt.Fprintf(GinkgoWriter, "\nExpected error: %s", err.Error())
		})
	})

	Describe(`ListRules - List rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully list rules`, func() {

			listRulesOptions := &ListRulesOptions{
				AccountID:     &accountID,
				Labels:        core.StringPtr(TestLabel),
				Offset:        core.Int64Ptr(int64(0)),
				Limit:         core.Int64Ptr(int64(1000)),
				TransactionID: &transactionID,
			}

			ruleList, response, err := service.ListRules(listRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleList).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nReceived response:\n%s", common.ToJSON(ruleList))

			Expect(*ruleList.TotalCount).To(Equal(int64(2)))
			Expect(*ruleList.Offset).To(Equal(int64(0)))
			Expect(*ruleList.Limit).To(Equal(int64(1000)))
			Expect(*ruleList.First).ToNot(BeNil())
			Expect(*ruleList.Last).ToNot(BeNil())
		})
		It(`Fail to list rules with unauthorized user`, func() {

			listRulesOptions := &ListRulesOptions{
				AccountID:     &accountID,
				Labels:        core.StringPtr(TestLabel),
				Offset:        core.Int64Ptr(int64(0)),
				Limit:         core.Int64Ptr(int64(1000)),
				TransactionID: &transactionID,
			}

			_, response, err := serviceNoAccess.ListRules(listRulesOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(403))
			fmt.Fprintf(GinkgoWriter, "\nExpected error: %s", err.Error())
		})
	})

	Describe(`GetRule - Get a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully get rule #1`, func() {
			Expect(ruleID1).ToNot(BeEmpty())

			getRuleOptions := &GetRuleOptions{
				RuleID:        &ruleID1,
				TransactionID: &transactionID,
			}

			rule, response, err := service.GetRule(getRuleOptions)
			fmt.Fprintf(GinkgoWriter, "\nReceived detailed response:\n%s", common.ToJSON(response))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
			rule1 = rule

			// Grab the Etag value from the response for use in the update operation.
			ruleEtag1 = response.GetHeaders().Get("Etag")
			Expect(ruleEtag1).ToNot(BeEmpty())
		})
		It(`Fail to get rule with invalid rule id`, func() {

			getRuleOptions := &GetRuleOptions{
				RuleID:        core.StringPtr("BOGUS_ID"),
				TransactionID: &transactionID,
			}

			rule, response, err := service.GetRule(getRuleOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(response.Result).ToNot(BeNil())
			Expect(rule).To(BeNil())

			fmt.Fprintf(GinkgoWriter, "\nReceived detailed response:\n%s", common.ToJSON(response))

			errorResponse, ok := response.Result.(map[string]interface{})
			Expect(ok).To(BeTrue())
			Expect(errorResponse["trace"]).To(Equal(transactionID))
			Expect(err.Error()).To(ContainSubstring("not found"))
		})
	})

	Describe(`UpdateRule - Update a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully update rule #1`, func() {
			Expect(ruleID1).ToNot(BeEmpty())
			Expect(rule1).ToNot(BeNil())
			Expect(ruleEtag1).ToNot(BeEmpty())

			newDescription := fmt.Sprintf("Updated: %s", *rule1.Description)
			updateRuleOptions := &UpdateRuleOptions{
				RuleID:             &ruleID1,
				IfMatch:            &ruleEtag1,
				Name:               rule1.Name,
				Description:        &newDescription,
				Target:             rule1.Target,
				RequiredConfig:     rule1.RequiredConfig,
				EnforcementActions: rule1.EnforcementActions,
				AccountID:          rule1.AccountID,
				RuleType:           rule1.RuleType,
				Labels:             rule1.Labels,
				TransactionID:      &transactionID,
			}

			rule, response, err := service.UpdateRule(updateRuleOptions)
			fmt.Fprintf(GinkgoWriter, "\nReceived detailed response:\n%s", common.ToJSON(response))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
			Expect(*rule.Description).To(ContainSubstring("Updated:"))
		})
		It(`Fail to update rule using invalid Etag`, func() {
			Expect(ruleID1).ToNot(BeEmpty())
			Expect(rule1).ToNot(BeNil())

			newDescription := fmt.Sprintf("Updated: %s", *rule1.Description)
			updateRuleOptions := &UpdateRuleOptions{
				RuleID:             &ruleID1,
				IfMatch:            core.StringPtr("BOGUS_ETAG"),
				Name:               rule1.Name,
				Description:        &newDescription,
				Target:             rule1.Target,
				RequiredConfig:     rule1.RequiredConfig,
				EnforcementActions: rule1.EnforcementActions,
				AccountID:          rule1.AccountID,
				RuleType:           rule1.RuleType,
				Labels:             rule1.Labels,
				TransactionID:      &transactionID,
			}

			rule, response, err := service.UpdateRule(updateRuleOptions)
			fmt.Fprintf(GinkgoWriter, "\nReceived detailed response:\n%s", common.ToJSON(response))

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(400))
			Expect(rule).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("If-Match"))
		})
	})

	Describe(`DeleteRule - Delete a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully delete rule #2`, func() {
			Expect(ruleID2).ToNot(BeEmpty())

			deleteRuleOptions := &DeleteRuleOptions{
				RuleID:        &ruleID2,
				TransactionID: &transactionID,
			}

			response, err := service.DeleteRule(deleteRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			// Next, check to make sure ListRules() returns only 1 rule.
			listRulesOptions := &ListRulesOptions{
				AccountID:     &accountID,
				Labels:        core.StringPtr(TestLabel),
				Offset:        core.Int64Ptr(int64(0)),
				Limit:         core.Int64Ptr(int64(1000)),
				TransactionID: &transactionID,
			}

			ruleList, response, err := service.ListRules(listRulesOptions)
			Expect(err).To(BeNil())
			Expect(ruleList).ToNot(BeNil())
			Expect(*ruleList.TotalCount).To(Equal(int64(1)))

			// Next, make sure we can't do a get on the deleted rule.
			rule := getRule(service, ruleID2)
			Expect(rule).To(BeNil())
		})
		It(`Fail to delete rule with invalid rule id`, func() {

			deleteRuleOptions := &DeleteRuleOptions{
				RuleID:        core.StringPtr("BOGUS_ID"),
				TransactionID: &transactionID,
			}

			response, err := service.DeleteRule(deleteRuleOptions)
			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(err.Error()).To(ContainSubstring("not found"))
		})
	})

	Describe(`CreateAttachments - Create attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully create attachment #1`, func() {
			Expect(ruleID1).ToNot(BeEmpty())

			attachmentRequestModel := &AttachmentRequest{
				AccountID:      &accountID,
				IncludedScope:  enterpriseScope,
				ExcludedScopes: []RuleScope{*accountScope},
			}

			createAttachmentsOptions := &CreateAttachmentsOptions{
				RuleID:        &ruleID1,
				Attachments:   []AttachmentRequest{*attachmentRequestModel},
				TransactionID: &transactionID,
			}

			createAttachmentsResponse, response, err := service.CreateAttachments(createAttachmentsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createAttachmentsResponse).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nReceived response:\n%s", common.ToJSON(createAttachmentsResponse))

			Expect(len(createAttachmentsResponse.Attachments)).To(Equal(1))
			attachment := createAttachmentsResponse.Attachments[0]
			Expect(attachment).ToNot(BeNil())
			Expect(attachment.AttachmentID).ToNot(BeNil())
			attachmentID1 = *attachment.AttachmentID
			Expect(attachmentID1).ToNot(BeNil())

			// Now retrieve the rule and make sure the number_of_attachments is 1.
			rule := getRule(service, ruleID1)
			Expect(rule).ToNot(BeNil())
			Expect(rule.NumberOfAttachments).ToNot(BeNil())
			Expect(*rule.NumberOfAttachments).To(Equal(int64(1)))
		})
		It(`Successfully create attachment #2`, func() {
			Expect(ruleID1).ToNot(BeEmpty())

			attachmentRequestModel := &AttachmentRequest{
				AccountID:     &accountID,
				IncludedScope: accountScope,
			}

			createAttachmentsOptions := &CreateAttachmentsOptions{
				RuleID:        &ruleID1,
				Attachments:   []AttachmentRequest{*attachmentRequestModel},
				TransactionID: &transactionID,
			}

			createAttachmentsResponse, response, err := service.CreateAttachments(createAttachmentsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createAttachmentsResponse).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nReceived response:\n%s", common.ToJSON(createAttachmentsResponse))

			Expect(len(createAttachmentsResponse.Attachments)).To(Equal(1))
			attachment := createAttachmentsResponse.Attachments[0]
			Expect(attachment).ToNot(BeNil())
			Expect(attachment.AttachmentID).ToNot(BeNil())
			attachmentID2 = *attachment.AttachmentID
			Expect(attachmentID2).ToNot(BeEmpty())

			// Now retrieve the rule and make sure the number_of_attachments is 1.
			rule := getRule(service, ruleID1)
			Expect(rule).ToNot(BeNil())
			Expect(rule.NumberOfAttachments).ToNot(BeNil())
			Expect(*rule.NumberOfAttachments).To(Equal(int64(2)))
		})
		It(`Fail to create attachment with an invalid scope type`, func() {
			Expect(ruleID1).ToNot(BeEmpty())
			Expect(badScope).ToNot(BeNil())

			attachmentRequestModel := &AttachmentRequest{
				AccountID:      &accountID,
				IncludedScope:  enterpriseScope,
				ExcludedScopes: []RuleScope{*badScope},
			}

			createAttachmentsOptions := &CreateAttachmentsOptions{
				RuleID:        &ruleID1,
				Attachments:   []AttachmentRequest{*attachmentRequestModel},
				TransactionID: &transactionID,
			}

			createAttachmentsResponse, response, err := service.CreateAttachments(createAttachmentsOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(400))
			Expect(createAttachmentsResponse).To(BeNil())
		})
	})

	Describe(`GetAttachment - Get an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully get attachment #1`, func() {
			Expect(ruleID1).ToNot(BeEmpty())
			Expect(attachmentID1).ToNot(BeEmpty())

			getAttachmentOptions := &GetAttachmentOptions{
				RuleID:        &ruleID1,
				AttachmentID:  &attachmentID1,
				TransactionID: &transactionID,
			}

			attachment, response, err := service.GetAttachment(getAttachmentOptions)
			fmt.Fprintf(GinkgoWriter, "\nReceived detailed response:\n%s", common.ToJSON(response))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())
			attachment1 = attachment

			Expect(*attachment.AccountID).To(Equal(accountID))
			Expect(*attachment.RuleID).To(Equal(ruleID1))
			Expect(*attachment.AttachmentID).To(Equal(attachmentID1))
			Expect(*attachment.IncludedScope.Note).To(Equal("enterprise"))
			Expect(len(attachment.ExcludedScopes)).To(Equal(1))

			// Grab the Etag value from the response for use in the update operation.
			attachmentEtag1 = response.GetHeaders().Get("Etag")
			Expect(attachmentEtag1).ToNot(BeEmpty())
		})
		It(`Fail to get attachment with invalid attachment id`, func() {
			Expect(ruleID1).ToNot(BeEmpty())

			getAttachmentOptions := &GetAttachmentOptions{
				RuleID:        &ruleID1,
				AttachmentID:  core.StringPtr("BOGUS_ID"),
				TransactionID: &transactionID,
			}

			attachment, response, err := service.GetAttachment(getAttachmentOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(attachment).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("not found"))
		})
	})

	Describe(`ListAttachments - List attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully list attachments for rule #1`, func() {
			Expect(ruleID1).ToNot(BeEmpty())

			listAttachmentsOptions := &ListAttachmentsOptions{
				RuleID:        &ruleID1,
				TransactionID: &transactionID,
				Offset:        core.Int64Ptr(int64(0)),
				Limit:         core.Int64Ptr(int64(1000)),
			}

			attachmentsList, response, err := service.ListAttachments(listAttachmentsOptions)
			fmt.Fprintf(GinkgoWriter, "\nReceived detailed response:\n%s", common.ToJSON(response))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentsList).ToNot(BeNil())
			Expect(*attachmentsList.TotalCount).To(Equal(int64(2)))
			Expect(*attachmentsList.Offset).To(Equal(int64(0)))
			Expect(*attachmentsList.Limit).To(Equal(int64(1000)))
			Expect(attachmentsList.First).ToNot(BeNil())
			Expect(attachmentsList.Last).ToNot(BeNil())
			for _, att := range attachmentsList.Attachments {
				if attachmentID1 == *att.AttachmentID {
					Expect(*att.IncludedScope.Note).To(Equal("enterprise"))
					Expect(len(att.ExcludedScopes)).To(Equal(1))
				} else if attachmentID2 == *att.AttachmentID {
					Expect(*att.IncludedScope.Note).To(Equal("leaf account"))
					Expect(att.ExcludedScopes).To(BeEmpty())
				} else {
					Fail(fmt.Sprintf("Unrecognized attachmentID: %s", *att.AttachmentID))
				}
			}
		})
	})

	Describe(`UpdateAttachment - Update an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully update attachment #1`, func() {
			Expect(ruleID1).ToNot(BeEmpty())
			Expect(attachment1).ToNot(BeNil())
			Expect(attachmentID1).ToNot(BeEmpty())
			Expect(attachmentEtag1).ToNot(BeEmpty())

			// Update the note field of attachment1's IncludedScope field.
			updatedNote := fmt.Sprintf("Updated: %s", *attachment1.IncludedScope.Note)
			updatedScope := &RuleScope{
				Note:      &updatedNote,
				ScopeID:   attachment1.IncludedScope.ScopeID,
				ScopeType: attachment1.IncludedScope.ScopeType,
			}

			updateAttachmentOptions := &UpdateAttachmentOptions{
				RuleID:         &ruleID1,
				AttachmentID:   &attachmentID1,
				IfMatch:        &attachmentEtag1,
				AccountID:      &accountID,
				IncludedScope:  updatedScope,
				ExcludedScopes: attachment1.ExcludedScopes,
				TransactionID:  &transactionID,
			}

			attachment, response, err := service.UpdateAttachment(updateAttachmentOptions)
			fmt.Fprintf(GinkgoWriter, "\nReceived detailed response:\n%s", common.ToJSON(response))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())
			Expect(*attachment.IncludedScope.Note).To(ContainSubstring("Updated:"))
		})
		It(`Fail to update attachment with invalid etag`, func() {
			Expect(ruleID1).ToNot(BeEmpty())
			Expect(attachment1).ToNot(BeNil())

			// Update the note field of attachment1's IncludedScope field.
			updatedNote := fmt.Sprintf("Updated: %s", *attachment1.IncludedScope.Note)
			updatedScope := &RuleScope{
				Note:      &updatedNote,
				ScopeID:   attachment1.IncludedScope.ScopeID,
				ScopeType: attachment1.IncludedScope.ScopeType,
			}

			updateAttachmentOptions := &UpdateAttachmentOptions{
				RuleID:         &ruleID1,
				AttachmentID:   &attachmentID1,
				IfMatch:        core.StringPtr("BOGUS_ETAG"),
				AccountID:      &accountID,
				IncludedScope:  updatedScope,
				ExcludedScopes: attachment1.ExcludedScopes,
				TransactionID:  &transactionID,
			}

			attachment, response, err := service.UpdateAttachment(updateAttachmentOptions)
			fmt.Fprintf(GinkgoWriter, "\nReceived detailed response:\n%s", common.ToJSON(response))

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(400))
			Expect(attachment).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("If-Match"))
		})
	})

	Describe(`DeleteAttachment - Delete an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Successfully delete attachment #2`, func() {
			Expect(ruleID1).ToNot(BeEmpty())
			Expect(attachmentID2).ToNot(BeEmpty())

			deleteAttachmentOptions := &DeleteAttachmentOptions{
				RuleID:        &ruleID1,
				AttachmentID:  &attachmentID2,
				TransactionID: &transactionID,
			}

			response, err := service.DeleteAttachment(deleteAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`Fail to delete attachment with invalid attachment id`, func() {
			Expect(ruleID1).ToNot(BeEmpty())

			deleteAttachmentOptions := &DeleteAttachmentOptions{
				RuleID:        &ruleID1,
				AttachmentID:  core.StringPtr("BOGUS_ID"),
				TransactionID: &transactionID,
			}

			response, err := service.DeleteAttachment(deleteAttachmentOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(err.Error()).To(ContainSubstring("not found"))
		})
	})

	Describe(`Teardown - clean up test data`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Clean rules`, func() {
			fmt.Fprintln(GinkgoWriter, "Teardown...")
			cleanRules(service, accountID, TestLabel)
			fmt.Fprintln(GinkgoWriter, "Finished teardown.")
		})
	})
})

func cleanRules(service *ConfigurationGovernanceV1, accountID string, label string) {
	fmt.Fprintln(GinkgoWriter, "Cleaning rules...")

	listRulesOptions := &ListRulesOptions{
		AccountID:     &accountID,
		Labels:        core.StringPtr(label),
		Offset:        core.Int64Ptr(int64(0)),
		Limit:         core.Int64Ptr(int64(1000)),
		TransactionID: &transactionID,
	}

	ruleList, response, err := service.ListRules(listRulesOptions)

	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))
	Expect(ruleList).ToNot(BeNil())
	Expect(ruleList.TotalCount).ToNot(BeNil())

	fmt.Fprintf(GinkgoWriter, "Found %d rule(s) to be cleaned", *ruleList.TotalCount)

	if *ruleList.TotalCount > 0 {
		for _, rule := range ruleList.Rules {

			fmt.Fprintf(GinkgoWriter, "Deleting rule: name='%s' id='%s'", *rule.Name, *rule.RuleID)

			response, err := service.DeleteRule(
				&DeleteRuleOptions{
					RuleID:        rule.RuleID,
					TransactionID: &transactionID,
				})
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		}
	}

	fmt.Fprintln(GinkgoWriter, "Finished cleaning rules...")
}

func getRule(service *ConfigurationGovernanceV1, ruleID string) (rule *Rule) {
	rule, _, _ = service.GetRule(&GetRuleOptions{
		RuleID:        core.StringPtr(ruleID),
		TransactionID: &transactionID,
	})
	return
}
