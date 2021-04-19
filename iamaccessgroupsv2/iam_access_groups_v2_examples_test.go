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
// This file provides an example of how to use the IAM Access Groups service.
//
// The following configuration properties are assumed to be defined:
//
// IAM_ACCESS_GROUPS_URL=<service url>
// IAM_ACCESS_GROUPS_AUTHTYPE=iam
// IAM_ACCESS_GROUPS_APIKEY=<your iam apikey>
// IAM_ACCESS_GROUPS_AUTH_URL=<IAM token service URL - omit this if using the production environment>
// IAM_ACCESS_GROUPS_TEST_ACCOUNT_ID=<id of an account used for testing>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../iam_access_groups.env"

var (
	iamAccessGroupsService *iamaccessgroupsv2.IamAccessGroupsV2
	config                 map[string]string
	configLoaded           bool = false

	testAccountID       string
	testGroupEtag       string
	testGroupID         string
	testClaimRuleID     string
	testClaimRuleEtag   string
	testAccountSettings *iamaccessgroupsv2.AccountSettings
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`IamAccessGroupsV2 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(iamaccessgroupsv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			testAccountID = config["TEST_ACCOUNT_ID"]

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

			createAccessGroupOptions := iamAccessGroupsService.NewCreateAccessGroupOptions(testAccountID, "Managers")
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

			testGroupID = *group.ID
		})
		It(`GetAccessGroup request example`, func() {
			// begin-get_access_group
			fmt.Println("\nGetAccessGroup() result:")

			getAccessGroupOptions := iamAccessGroupsService.NewGetAccessGroupOptions(testGroupID)
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

			testGroupEtag = response.GetHeaders().Get("ETag")
		})
		It(`UpdateAccessGroup request example`, func() {
			fmt.Println("\nUpdateAccessGroup() result:")
			// begin-update_access_group

			updateAccessGroupOptions := iamAccessGroupsService.NewUpdateAccessGroupOptions(testGroupID, testGroupEtag)
			updateAccessGroupOptions.SetName("Awesome Managers")
			updateAccessGroupOptions.SetDescription("Group for awesome managers")
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
		It(`ListAccessGroups request example`, func() {
			fmt.Println("\nListAccessGroups() result:")
			// begin-list_access_groups

			listAccessGroupsOptions := iamAccessGroupsService.NewListAccessGroupsOptions(testAccountID)
			groupsList, response, err := iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(groupsList, "", "  ")
			fmt.Println(string(b))

			// end-list_access_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(groupsList).ToNot(BeNil())
		})
		It(`AddMembersToAccessGroup request example`, func() {
			fmt.Println("\nAddMembersToAccessGroup() result:")
			// begin-add_members_to_access_group

			groupMembers := []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{
				iamaccessgroupsv2.AddGroupMembersRequestMembersItem{
					IamID: core.StringPtr("IBMid-user1"),
					Type:  core.StringPtr("user"),
				},
				iamaccessgroupsv2.AddGroupMembersRequestMembersItem{
					IamID: core.StringPtr("iam-ServiceId-123"),
					Type:  core.StringPtr("service"),
				},
			}
			addMembersToAccessGroupOptions := iamAccessGroupsService.NewAddMembersToAccessGroupOptions(testGroupID)
			addMembersToAccessGroupOptions.SetMembers(groupMembers)
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
		It(`IsMemberOfAccessGroup request example`, func() {
			// begin-is_member_of_access_group

			isMemberOfAccessGroupOptions := iamAccessGroupsService.NewIsMemberOfAccessGroupOptions(testGroupID, "IBMid-user1")
			response, err := iamAccessGroupsService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptions)
			if err != nil {
				panic(err)
			}

			// end-is_member_of_access_group
			fmt.Printf("\nIsMemberOfAccessGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`ListAccessGroupMembers request example`, func() {
			fmt.Println("\nListAccessGroupMembers() result:")
			// begin-list_access_group_members

			listAccessGroupMembersOptions := iamAccessGroupsService.NewListAccessGroupMembersOptions(testGroupID)
			groupMembersList, response, err := iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(groupMembersList, "", "  ")
			fmt.Println(string(b))

			// end-list_access_group_members

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(groupMembersList).ToNot(BeNil())
		})
		It(`RemoveMemberFromAccessGroup request example`, func() {
			// begin-remove_member_from_access_group

			removeMemberFromAccessGroupOptions := iamAccessGroupsService.NewRemoveMemberFromAccessGroupOptions(testGroupID, "IBMid-user1")
			response, err := iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions)
			if err != nil {
				panic(err)
			}

			// end-remove_member_from_access_group
			fmt.Printf("\nRemoveMemberFromAccessGroup() response status code: %d\n:", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`RemoveMembersFromAccessGroup request example`, func() {
			fmt.Println("\nRemoveMembersFromAccessGroup() result:")
			// begin-remove_members_from_access_group

			removeMembersFromAccessGroupOptions := iamAccessGroupsService.NewRemoveMembersFromAccessGroupOptions(testGroupID)
			removeMembersFromAccessGroupOptions.SetMembers([]string{"iam-ServiceId-123"})
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

			addMemberToMultipleAccessGroupsOptions := iamAccessGroupsService.NewAddMemberToMultipleAccessGroupsOptions(testAccountID, "IBMid-user1")
			addMemberToMultipleAccessGroupsOptions.SetType("user")
			addMemberToMultipleAccessGroupsOptions.SetGroups([]string{testGroupID})
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
		It(`RemoveMemberFromAllAccessGroups request example`, func() {
			fmt.Println("\nRemoveMemberFromAllAccessGroups() result:")
			// begin-remove_member_from_all_access_groups

			removeMemberFromAllAccessGroupsOptions := iamAccessGroupsService.NewRemoveMemberFromAllAccessGroupsOptions(testAccountID, "IBMid-user1")
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
		It(`AddAccessGroupRule request example`, func() {
			fmt.Println("\nAddAccessGroupRule() result:")
			// begin-add_access_group_rule

			ruleConditionsModel := iamaccessgroupsv2.RuleConditions{
				Claim:    core.StringPtr("isManager"),
				Operator: core.StringPtr("EQUALS"),
				Value:    core.StringPtr("true"),
			}
			addAccessGroupRuleOptions := iamAccessGroupsService.NewAddAccessGroupRuleOptions(
				testGroupID,
				int64(12),
				"https://idp.example.org/SAML2",
				[]iamaccessgroupsv2.RuleConditions{ruleConditionsModel},
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
		It(`GetAccessGroupRule request example`, func() {
			fmt.Println("\nGetAccessGroupRule() result:")
			// begin-get_access_group_rule

			getAccessGroupRuleOptions := iamAccessGroupsService.NewGetAccessGroupRuleOptions(testGroupID, testClaimRuleID)
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

			ruleConditionsModel := iamaccessgroupsv2.RuleConditions{
				Claim:    core.StringPtr("isManager"),
				Operator: core.StringPtr("EQUALS"),
				Value:    core.StringPtr("true"),
			}
			replaceAccessGroupRuleOptions := iamAccessGroupsService.NewReplaceAccessGroupRuleOptions(
				testGroupID,
				testClaimRuleID,
				testClaimRuleEtag,
				int64(24),
				"https://idp.example.org/SAML2",
				[]iamaccessgroupsv2.RuleConditions{ruleConditionsModel},
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
		It(`ListAccessGroupRules request example`, func() {
			fmt.Println("\nListAccessGroupRules() result:")
			// begin-list_access_group_rules

			listAccessGroupRulesOptions := iamAccessGroupsService.NewListAccessGroupRulesOptions(testGroupID)
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
		It(`RemoveAccessGroupRule request example`, func() {
			// begin-remove_access_group_rule

			removeAccessGroupRuleOptions := iamAccessGroupsService.NewRemoveAccessGroupRuleOptions(testGroupID, testClaimRuleID)
			response, err := iamAccessGroupsService.RemoveAccessGroupRule(removeAccessGroupRuleOptions)
			if err != nil {
				panic(err)
			}

			// end-remove_access_group_rule
			fmt.Printf("\nRemoveAccessGroupRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`GetAccountSettings request example`, func() {
			fmt.Println("\nGetAccountSettings() result:")
			// begin-get_account_settings

			getAccountSettingsOptions := iamAccessGroupsService.NewGetAccountSettingsOptions(testAccountID)
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

			testAccountSettings = accountSettings
		})
		It(`UpdateAccountSettings request example`, func() {
			fmt.Println("\nUpdateAccountSettings() result:")
			// begin-update_account_settings

			updateAccountSettingsOptions := iamAccessGroupsService.NewUpdateAccountSettingsOptions(testAccountID)
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
		It(`DeleteAccessGroup request example`, func() {
			// begin-delete_access_group

			deleteAccessGroupOptions := iamAccessGroupsService.NewDeleteAccessGroupOptions(testGroupID)
			response, err := iamAccessGroupsService.DeleteAccessGroup(deleteAccessGroupOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_access_group
			fmt.Printf("\nDeleteAccessGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
