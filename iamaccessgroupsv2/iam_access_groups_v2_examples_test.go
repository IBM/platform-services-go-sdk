// +build examples

/**
 * (C) Copyright IBM Corp. 2023.
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

package iamaccessgroupsv2_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the iam-access-groups service.
//
// The following configuration properties are assumed to be defined:
// IAM_ACCESS_GROUPS_URL=<service base url>
// IAM_ACCESS_GROUPS_AUTH_TYPE=iam
// IAM_ACCESS_GROUPS_APIKEY=<IAM apikey>
// IAM_ACCESS_GROUPS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`IamAccessGroupsV2 Examples Tests`, func() {

	const externalConfigFile = "../iam_access_groups_v2.env"

	var (
		iamAccessGroupsService *iamaccessgroupsv2.IamAccessGroupsV2
		config       map[string]string

		// Variables to hold link values
		accessGroupETagLink string
		accessGroupIDLink string
		testAccountID       string
		testProfileID       string
		testClaimRuleID     string
		testClaimRuleEtag   string
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
			config, err = core.GetServiceProperties(iamaccessgroupsv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			testAccountID = config["TEST_ACCOUNT_ID"]
			testProfileID = config["TEST_PROFILE_ID"]
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

			iamAccessGroupsServiceOptions := &iamaccessgroupsv2.IamAccessGroupsV2Options{}

			iamAccessGroupsService, err = iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(iamAccessGroupsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(iamAccessGroupsService).ToNot(BeNil())
		})
	})

	Describe(`IamAccessGroupsV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAccessGroup request example`, func() {
			fmt.Println("\nCreateAccessGroup() result:")
			// begin-create_access_group

			createAccessGroupOptions := iamAccessGroupsService.NewCreateAccessGroupOptions(
				testAccountID,
				"Managers",
			)
			createAccessGroupOptions.SetDescription("Group for managers")

			group, response, err := iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(group, "", "  ")
			fmt.Println(string(b))

			// end-create_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(group).ToNot(BeNil())

			accessGroupIDLink = *group.ID
			fmt.Fprintf(GinkgoWriter, "Saved accessGroupIDLink value: %v\n", accessGroupIDLink)
		})
		It(`GetAccessGroup request example`, func() {
			fmt.Println("\nGetAccessGroup() result:")
			// begin-get_access_group

			getAccessGroupOptions := iamAccessGroupsService.NewGetAccessGroupOptions(
				accessGroupIDLink,
			)

			group, response, err := iamAccessGroupsService.GetAccessGroup(getAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(group, "", "  ")
			fmt.Println(string(b))

			// end-get_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(group).ToNot(BeNil())

			accessGroupETagLink = response.Headers.Get("ETag")
			fmt.Fprintf(GinkgoWriter, "Saved accessGroupETagLink value: %v\n", accessGroupETagLink)
		})
		It(`ListAccessGroups request example`, func() {
			fmt.Println("\nListAccessGroups() result:")
			// begin-list_access_groups
			listAccessGroupsOptions := &iamaccessgroupsv2.ListAccessGroupsOptions{
				AccountID: &testAccountID,
				TransactionID: core.StringPtr("testString"),
				IamID: core.StringPtr("testString"),
				Search: core.StringPtr("testString"),
				MembershipType: core.StringPtr("static"),
				Limit: core.Int64Ptr(int64(10)),
				Sort: core.StringPtr("name"),
				ShowFederated: core.BoolPtr(false),
				HidePublicAccess: core.BoolPtr(false),
			}

			pager, err := iamAccessGroupsService.NewAccessGroupsPager(listAccessGroupsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []iamaccessgroupsv2.Group
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_access_groups
		})
		It(`UpdateAccessGroup request example`, func() {
			fmt.Println("\nUpdateAccessGroup() result:")
			// begin-update_access_group

			updateAccessGroupOptions := iamAccessGroupsService.NewUpdateAccessGroupOptions(
				accessGroupIDLink,
				accessGroupETagLink,
			)
			updateAccessGroupOptions.SetName("Awesome Managers")
			updateAccessGroupOptions.SetDescription("Group for awesome managers.")

			group, response, err := iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(group, "", "  ")
			fmt.Println(string(b))

			// end-update_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(group).ToNot(BeNil())
		})
		It(`IsMemberOfAccessGroup request example`, func() {
			// begin-is_member_of_access_group

			isMemberOfAccessGroupOptions := iamAccessGroupsService.NewIsMemberOfAccessGroupOptions(
				accessGroupIDLink,
				"testString",
			)

			response, err := iamAccessGroupsService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from IsMemberOfAccessGroup(): %d\n", response.StatusCode)
			}

			// end-is_member_of_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`AddMembersToAccessGroup request example`, func() {
			fmt.Println("\nAddMembersToAccessGroup() result:")
			// begin-add_members_to_access_group

			addGroupMembersRequestMembersItemModel := &iamaccessgroupsv2.AddGroupMembersRequestMembersItem{
				IamID: core.StringPtr("IBMid-user1"),
				Type: core.StringPtr("user"),
			}

			addMembersToAccessGroupOptions := iamAccessGroupsService.NewAddMembersToAccessGroupOptions(
				accessGroupIDLink,
			)
			addMembersToAccessGroupOptions.SetMembers([]iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel})

			addGroupMembersResponse, response, err := iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addGroupMembersResponse, "", "  ")
			fmt.Println(string(b))

			// end-add_members_to_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(207))
			Expect(addGroupMembersResponse).ToNot(BeNil())
		})
		It(`ListAccessGroupMembers request example`, func() {
			fmt.Println("\nListAccessGroupMembers() result:")
			// begin-list_access_group_members
			listAccessGroupMembersOptions := &iamaccessgroupsv2.ListAccessGroupMembersOptions{
				AccessGroupID: &accessGroupIDLink,
				TransactionID: core.StringPtr("testString"),
				MembershipType: core.StringPtr("static"),
				Limit: core.Int64Ptr(int64(10)),
				Type: core.StringPtr("testString"),
				Verbose: core.BoolPtr(false),
				Sort: core.StringPtr("testString"),
			}

			pager, err := iamAccessGroupsService.NewAccessGroupMembersPager(listAccessGroupMembersOptions)
			if err != nil {
				panic(err)
			}

			var allResults []iamaccessgroupsv2.ListGroupMembersResponseMember
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_access_group_members
		})
		It(`RemoveMembersFromAccessGroup request example`, func() {
			fmt.Println("\nRemoveMembersFromAccessGroup() result:")
			// begin-remove_members_from_access_group

			removeMembersFromAccessGroupOptions := iamAccessGroupsService.NewRemoveMembersFromAccessGroupOptions(
				accessGroupIDLink,
			)
			removeMembersFromAccessGroupOptions.SetMembers([]string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"})

			deleteGroupBulkMembersResponse, response, err := iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteGroupBulkMembersResponse, "", "  ")
			fmt.Println(string(b))

			// end-remove_members_from_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(207))
			Expect(deleteGroupBulkMembersResponse).ToNot(BeNil())
		})
		It(`AddMemberToMultipleAccessGroups request example`, func() {
			fmt.Println("\nAddMemberToMultipleAccessGroups() result:")
			// begin-add_member_to_multiple_access_groups

			addMemberToMultipleAccessGroupsOptions := iamAccessGroupsService.NewAddMemberToMultipleAccessGroupsOptions(
				testAccountID,
				"IBMid-user1",
			)
			addMemberToMultipleAccessGroupsOptions.SetType("user")
			addMemberToMultipleAccessGroupsOptions.SetGroups([]string{"access-group-id-1"})

			addMembershipMultipleGroupsResponse, response, err := iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addMembershipMultipleGroupsResponse, "", "  ")
			fmt.Println(string(b))

			// end-add_member_to_multiple_access_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(207))
			Expect(addMembershipMultipleGroupsResponse).ToNot(BeNil())
		})
		It(`AddAccessGroupRule request example`, func() {
			fmt.Println("\nAddAccessGroupRule() result:")
			// begin-add_access_group_rule

			ruleConditionsModel := &iamaccessgroupsv2.RuleConditions{
				Claim: core.StringPtr("isManager"),
				Operator: core.StringPtr("EQUALS"),
				Value: core.StringPtr("true"),
			}

			addAccessGroupRuleOptions := iamAccessGroupsService.NewAddAccessGroupRuleOptions(
				accessGroupIDLink,
				int64(12),
				"https://idp.example.org/SAML2",
				[]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel},
			)
			addAccessGroupRuleOptions.SetName("Manager group rule")

			rule, response, err := iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-add_access_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())
			testClaimRuleID = *rule.ID
		})
		It(`ListAccessGroupRules request example`, func() {
			fmt.Println("\nListAccessGroupRules() result:")
			// begin-list_access_group_rules

			listAccessGroupRulesOptions := iamAccessGroupsService.NewListAccessGroupRulesOptions(
				accessGroupIDLink,
			)

			rulesList, response, err := iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rulesList, "", "  ")
			fmt.Println(string(b))

			// end-list_access_group_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesList).ToNot(BeNil())
		})
		It(`GetAccessGroupRule request example`, func() {
			fmt.Println("\nGetAccessGroupRule() result:")
			// begin-get_access_group_rule

			getAccessGroupRuleOptions := iamAccessGroupsService.NewGetAccessGroupRuleOptions(
				accessGroupIDLink,
				testClaimRuleID,
			)

			rule, response, err := iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-get_access_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
			testClaimRuleEtag = response.GetHeaders().Get("Etag")
		})
		It(`ReplaceAccessGroupRule request example`, func() {
			fmt.Println("\nReplaceAccessGroupRule() result:")
			// begin-replace_access_group_rule

			ruleConditionsModel := &iamaccessgroupsv2.RuleConditions{
				Claim: core.StringPtr("isManager"),
				Operator: core.StringPtr("EQUALS"),
				Value: core.StringPtr("true"),
			}

			replaceAccessGroupRuleOptions := iamAccessGroupsService.NewReplaceAccessGroupRuleOptions(
				accessGroupIDLink,
				testClaimRuleID,
				testClaimRuleEtag,
				int64(12),
				"https://idp.example.org/SAML2",
				[]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel},
			)
			replaceAccessGroupRuleOptions.SetName("Manager group rule")

			rule, response, err := iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-replace_access_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
		It(`GetAccountSettings request example`, func() {
			fmt.Println("\nGetAccountSettings() result:")
			// begin-get_account_settings

			getAccountSettingsOptions := iamAccessGroupsService.NewGetAccountSettingsOptions(
				testAccountID,
			)

			accountSettings, response, err := iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_account_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())
		})
		It(`UpdateAccountSettings request example`, func() {
			fmt.Println("\nUpdateAccountSettings() result:")
			// begin-update_account_settings

			updateAccountSettingsOptions := iamAccessGroupsService.NewUpdateAccountSettingsOptions(
				testAccountID,
			)
			updateAccountSettingsOptions.SetPublicAccessEnabled(true)

			accountSettings, response, err := iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-update_account_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())
		})
		It(`CreateTemplate request example`, func() {
			fmt.Println("\nCreateTemplate() result:")
			// begin-create_template

			membersActionControlsModel := &iamaccessgroupsv2.MembersActionControls{
				Add: core.BoolPtr(true),
				Remove: core.BoolPtr(false),
			}

			membersInputModel := &iamaccessgroupsv2.MembersInput{
				Users: []string{"IBMid-123", "IBMid-234"},
				ActionControls: membersActionControlsModel,
			}

			conditionInputModel := &iamaccessgroupsv2.ConditionInput{
				Claim: core.StringPtr("blueGroup"),
				Operator: core.StringPtr("CONTAINS"),
				Value: core.StringPtr("test-bluegroup-saml"),
			}

			rulesActionControlsModel := &iamaccessgroupsv2.RulesActionControls{
				Remove: core.BoolPtr(false),
				Update: core.BoolPtr(false),
			}

			ruleInputModel := &iamaccessgroupsv2.RuleInput{
				Name: core.StringPtr("Manager group rule"),
				Expiration: core.Int64Ptr(int64(12)),
				RealmName: core.StringPtr("https://idp.example.org/SAML2"),
				Conditions: []iamaccessgroupsv2.ConditionInput{*conditionInputModel},
				ActionControls: rulesActionControlsModel,
			}

			assertionsActionControlsModel := &iamaccessgroupsv2.AssertionsActionControls{
				Add: core.BoolPtr(false),
				Remove: core.BoolPtr(true),
				Update: core.BoolPtr(true),
			}

			assertionsInputModel := &iamaccessgroupsv2.AssertionsInput{
				Rules: []iamaccessgroupsv2.RuleInput{*ruleInputModel},
				ActionControls: assertionsActionControlsModel,
			}

			accessActionControlsModel := &iamaccessgroupsv2.AccessActionControls{
				Add: core.BoolPtr(false),
			}

			groupActionControlsModel := &iamaccessgroupsv2.GroupActionControls{
				Access: accessActionControlsModel,
			}

			accessGroupInputModel := &iamaccessgroupsv2.AccessGroupInput{
				Name: core.StringPtr("IAM Admin Group"),
				Description: core.StringPtr("This access group template allows admin access to all IAM platform services in the account."),
				Members: membersInputModel,
				Assertions: assertionsInputModel,
				ActionControls: groupActionControlsModel,
			}

			policyTemplatesInputModel := &iamaccessgroupsv2.PolicyTemplatesInput{
				ID: core.StringPtr("policyTemplateId-123"),
				Version: core.StringPtr("1"),
			}

			createTemplateOptions := iamAccessGroupsService.NewCreateTemplateOptions(
				"IAM Admin Group template",
				"This access group template allows admin access to all IAM platform services in the account.",
				"accountID-123",
			)
			createTemplateOptions.SetAccessGroup(accessGroupInputModel)
			createTemplateOptions.SetPolicyTemplates([]iamaccessgroupsv2.PolicyTemplatesInput{*policyTemplatesInputModel})

			createTemplateResponse, response, err := iamAccessGroupsService.CreateTemplate(createTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`ListTemplates request example`, func() {
			fmt.Println("\nListTemplates() result:")
			// begin-list_templates
			listTemplatesOptions := &iamaccessgroupsv2.ListTemplatesOptions{
				AccountID: core.StringPtr("accountID-123"),
				TransactionID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(50)),
				Verbose: core.BoolPtr(true),
			}

			pager, err := iamAccessGroupsService.NewTemplatesPager(listTemplatesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []iamaccessgroupsv2.TemplateItem
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_templates
		})
		It(`CreateTemplateVersion request example`, func() {
			fmt.Println("\nCreateTemplateVersion() result:")
			// begin-create_template_version

			membersInputModel := &iamaccessgroupsv2.MembersInput{
				Users: []string{"IBMid-123", "IBMid-234"},
			}

			conditionInputModel := &iamaccessgroupsv2.ConditionInput{
				Claim: core.StringPtr("blueGroup"),
				Operator: core.StringPtr("CONTAINS"),
				Value: core.StringPtr("test-bluegroup-saml"),
			}

			ruleInputModel := &iamaccessgroupsv2.RuleInput{
				Name: core.StringPtr("Manager group rule"),
				Expiration: core.Int64Ptr(int64(12)),
				RealmName: core.StringPtr("https://idp.example.org/SAML2"),
				Conditions: []iamaccessgroupsv2.ConditionInput{*conditionInputModel},
			}

			assertionsInputModel := &iamaccessgroupsv2.AssertionsInput{
				Rules: []iamaccessgroupsv2.RuleInput{*ruleInputModel},
			}

			accessGroupInputModel := &iamaccessgroupsv2.AccessGroupInput{
				Name: core.StringPtr("IAM Admin Group 8"),
				Description: core.StringPtr("This access group template allows admin access to all IAM platform services in the account."),
				Members: membersInputModel,
				Assertions: assertionsInputModel,
			}

			policyTemplatesInputModel := &iamaccessgroupsv2.PolicyTemplatesInput{
				ID: core.StringPtr("policyTemplateId-123"),
				Version: core.StringPtr("1"),
			}

			createTemplateVersionOptions := iamAccessGroupsService.NewCreateTemplateVersionOptions(
				"testString",
			)
			createTemplateVersionOptions.SetName("IAM Admin Group template 2")
			createTemplateVersionOptions.SetDescription("This access group template allows admin access to all IAM platform services in the account.")
			createTemplateVersionOptions.SetAccessGroup(accessGroupInputModel)
			createTemplateVersionOptions.SetPolicyTemplates([]iamaccessgroupsv2.PolicyTemplatesInput{*policyTemplatesInputModel})

			createTemplateResponse, response, err := iamAccessGroupsService.CreateTemplateVersion(createTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`ListTemplateVersions request example`, func() {
			fmt.Println("\nListTemplateVersions() result:")
			// begin-list_template_versions
			listTemplateVersionsOptions := &iamaccessgroupsv2.ListTemplateVersionsOptions{
				TemplateID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(100)),
			}

			pager, err := iamAccessGroupsService.NewTemplateVersionsPager(listTemplateVersionsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []iamaccessgroupsv2.ListTemplatesVersionsResponse
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_template_versions
		})
		It(`GetTemplateVersion request example`, func() {
			fmt.Println("\nGetTemplateVersion() result:")
			// begin-get_template_version

			getTemplateVersionOptions := iamAccessGroupsService.NewGetTemplateVersionOptions(
				"testString",
				"testString",
			)

			createTemplateResponse, response, err := iamAccessGroupsService.GetTemplateVersion(getTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`UpdateTemplateVersion request example`, func() {
			fmt.Println("\nUpdateTemplateVersion() result:")
			// begin-update_template_version

			membersInputModel := &iamaccessgroupsv2.MembersInput{
				Users: []string{"IBMid-5500085Q21"},
			}

			conditionInputModel := &iamaccessgroupsv2.ConditionInput{
				Claim: core.StringPtr("blueGroup"),
				Operator: core.StringPtr("CONTAINS"),
				Value: core.StringPtr("test-bluegroup-saml"),
			}

			ruleInputModel := &iamaccessgroupsv2.RuleInput{
				Name: core.StringPtr("Manager group rule"),
				Expiration: core.Int64Ptr(int64(12)),
				RealmName: core.StringPtr("https://idp.example.org/SAML2"),
				Conditions: []iamaccessgroupsv2.ConditionInput{*conditionInputModel},
			}

			assertionsInputModel := &iamaccessgroupsv2.AssertionsInput{
				Rules: []iamaccessgroupsv2.RuleInput{*ruleInputModel},
			}

			accessGroupInputModel := &iamaccessgroupsv2.AccessGroupInput{
				Name: core.StringPtr("IAM Admin Group 8"),
				Description: core.StringPtr("This access group template allows admin access to all IAM platform services in the account."),
				Members: membersInputModel,
				Assertions: assertionsInputModel,
			}

			policyTemplatesInputModel := &iamaccessgroupsv2.PolicyTemplatesInput{
				ID: core.StringPtr("policyTemplateId-123"),
				Version: core.StringPtr("1"),
			}

			updateTemplateVersionOptions := iamAccessGroupsService.NewUpdateTemplateVersionOptions(
				"testString",
				"testString",
				"testString",
			)
			updateTemplateVersionOptions.SetName("IAM Admin Group template 2")
			updateTemplateVersionOptions.SetDescription("This access group template allows admin access to all IAM platform services in the account.")
			updateTemplateVersionOptions.SetAccessGroup(accessGroupInputModel)
			updateTemplateVersionOptions.SetPolicyTemplates([]iamaccessgroupsv2.PolicyTemplatesInput{*policyTemplatesInputModel})
			updateTemplateVersionOptions.SetTransactionID("83adf5bd-de790caa3")

			createTemplateResponse, response, err := iamAccessGroupsService.UpdateTemplateVersion(updateTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`CommitTemplate request example`, func() {
			fmt.Println("\nCommitTemplate() result:")
			// begin-commit_template

			commitTemplateOptions := iamAccessGroupsService.NewCommitTemplateOptions(
				"testString",
				"testString",
				"testString",
			)

			createTemplateResponse, response, err := iamAccessGroupsService.CommitTemplate(commitTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-commit_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`GetLatestTemplateVersion request example`, func() {
			fmt.Println("\nGetLatestTemplateVersion() result:")
			// begin-get_latest_template_version

			getLatestTemplateVersionOptions := iamAccessGroupsService.NewGetLatestTemplateVersionOptions(
				"testString",
			)

			createTemplateResponse, response, err := iamAccessGroupsService.GetLatestTemplateVersion(getLatestTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_latest_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`CreateAssignment request example`, func() {
			fmt.Println("\nCreateAssignment() result:")
			// begin-create_assignment

			createAssignmentOptions := iamAccessGroupsService.NewCreateAssignmentOptions(
				"AccessGroupTemplateId-4be4",
				"1",
				"accountGroup",
				"0a45594d0f-123",
			)

			templateCreateAssignmentResponse, response, err := iamAccessGroupsService.CreateAssignment(createAssignmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateCreateAssignmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(templateCreateAssignmentResponse).ToNot(BeNil())
		})
		It(`ListAssignments request example`, func() {
			fmt.Println("\nListAssignments() result:")
			// begin-list_assignments

			listAssignmentsOptions := iamAccessGroupsService.NewListAssignmentsOptions()
			listAssignmentsOptions.SetAccountID("accountID-123")

			templatesListAssignmentResponse, response, err := iamAccessGroupsService.ListAssignments(listAssignmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templatesListAssignmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_assignments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templatesListAssignmentResponse).ToNot(BeNil())
		})
		It(`GetAssignment request example`, func() {
			fmt.Println("\nGetAssignment() result:")
			// begin-get_assignment

			getAssignmentOptions := iamAccessGroupsService.NewGetAssignmentOptions(
				"testString",
			)

			getTemplateAssignmentResponse, response, err := iamAccessGroupsService.GetAssignment(getAssignmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getTemplateAssignmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(getTemplateAssignmentResponse).ToNot(BeNil())
		})
		It(`DeleteAccessGroup request example`, func() {
			// begin-delete_access_group

			deleteAccessGroupOptions := iamAccessGroupsService.NewDeleteAccessGroupOptions(
				accessGroupIDLink,
			)

			response, err := iamAccessGroupsService.DeleteAccessGroup(deleteAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteAccessGroup(): %d\n", response.StatusCode)
			}

			// end-delete_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`RemoveMemberFromAccessGroup request example`, func() {
			// begin-remove_member_from_access_group

			removeMemberFromAccessGroupOptions := iamAccessGroupsService.NewRemoveMemberFromAccessGroupOptions(
				accessGroupIDLink,
				testProfileID,
			)

			response, err := iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from RemoveMemberFromAccessGroup(): %d\n", response.StatusCode)
			}

			// end-remove_member_from_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`RemoveAccessGroupRule request example`, func() {
			// begin-remove_access_group_rule

			removeAccessGroupRuleOptions := iamAccessGroupsService.NewRemoveAccessGroupRuleOptions(
				accessGroupIDLink,
				testClaimRuleID,
			)

			response, err := iamAccessGroupsService.RemoveAccessGroupRule(removeAccessGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from RemoveAccessGroupRule(): %d\n", response.StatusCode)
			}

			// end-remove_access_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`RemoveMemberFromAllAccessGroups request example`, func() {
			fmt.Println("\nRemoveMemberFromAllAccessGroups() result:")
			// begin-remove_member_from_all_access_groups

			removeMemberFromAllAccessGroupsOptions := iamAccessGroupsService.NewRemoveMemberFromAllAccessGroupsOptions(
				testAccountID,
				"IBMid-user1",
			)

			deleteFromAllGroupsResponse, response, err := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteFromAllGroupsResponse, "", "  ")
			fmt.Println(string(b))

			// end-remove_member_from_all_access_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(207))
			Expect(deleteFromAllGroupsResponse).ToNot(BeNil())
		})
		It(`DeleteTemplateVersion request example`, func() {
			// begin-delete_template_version

			deleteTemplateVersionOptions := iamAccessGroupsService.NewDeleteTemplateVersionOptions(
				"testString",
				"testString",
			)

			response, err := iamAccessGroupsService.DeleteTemplateVersion(deleteTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTemplateVersion(): %d\n", response.StatusCode)
			}

			// end-delete_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteTemplate request example`, func() {
			// begin-delete_template

			deleteTemplateOptions := iamAccessGroupsService.NewDeleteTemplateOptions(
				"testString",
			)

			response, err := iamAccessGroupsService.DeleteTemplate(deleteTemplateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTemplate(): %d\n", response.StatusCode)
			}

			// end-delete_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteAssignment request example`, func() {
			// begin-delete_assignment

			deleteAssignmentOptions := iamAccessGroupsService.NewDeleteAssignmentOptions(
				"testString",
			)

			response, err := iamAccessGroupsService.DeleteAssignment(deleteAssignmentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteAssignment(): %d\n", response.StatusCode)
			}

			// end-delete_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
