//go:build integration
// +build integration

/**
 * (C) Copyright IBM Corp. 2020, 2022.
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
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
)

var _ = Describe("IAM Access Groups - Integration Tests", func() {
	const externalConfigFile = "../iam_access_groups_v2.env"

	var (
		iamAccessGroupsService *iamaccessgroupsv2.IamAccessGroupsV2
		err                    error
		config                 map[string]string
		configLoaded           bool = false

		testAccountID        string
		testGroupName        string = "SDK Test Group - Golang"
		testGroupDescription string = "This group is used for integration test purposes. It can be deleted at any time."
		testGroupEtag        string
		testGroupID          string
		testUserID           string = "IBMid-" + strconv.Itoa(rand.Intn(100000))
		testClaimRuleID      string
		testClaimRuleEtag    string
		testAccountSettings  *iamaccessgroupsv2.AccountSettings

		userType   string = "user"
		etagHeader string = "Etag"
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping...")
		}
	}

	It("Successfully load the configuration", func() {
		err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
		if err != nil {
			Skip("Could not set IBM_CREDENTIALS_FILE environment variable: " + err.Error())
		}

		config, err = core.GetServiceProperties(iamaccessgroupsv2.DefaultServiceName)
		if err == nil {
			testAccountID = config["TEST_ACCOUNT_ID"]
			if testAccountID != "" {
				configLoaded = true
			}
		}
		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}
	})

	It(`Successfully created IamAccessGroupsV2 service instance`, func() {
		shouldSkipTest()

		iamAccessGroupsService, err = iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(
			&iamaccessgroupsv2.IamAccessGroupsV2Options{},
		)

		Expect(err).To(BeNil())
		Expect(iamAccessGroupsService).ToNot(BeNil())

		core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
		iamAccessGroupsService.EnableRetries(4, 30*time.Second)
	})

	Describe("Create an access group", func() {

		It("Successfully created an access group", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewCreateAccessGroupOptions(testAccountID, testGroupName)
			result, detailedResponse, err := iamAccessGroupsService.CreateAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.Name).To(Equal(testGroupName))

			testGroupID = *result.ID
		})
	})

	Describe("Get an access group", func() {

		It("Successfully retrieved an access group", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewGetAccessGroupOptions(testGroupID)
			result, detailedResponse, err := iamAccessGroupsService.GetAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ID).To(Equal(testGroupID))
			Expect(*result.Name).To(Equal(testGroupName))
			Expect(*result.Description).To(Equal(""))

			testGroupEtag = detailedResponse.GetHeaders().Get(etagHeader)
		})
	})

	Describe("Update an access group description", func() {

		It("Successfully updated an access group description", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewUpdateAccessGroupOptions(testGroupID, testGroupEtag)
			options.SetDescription(testGroupDescription)
			result, detailedResponse, err := iamAccessGroupsService.UpdateAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.Name).To(Equal(testGroupName))
			Expect(*result.ID).To(Equal(testGroupID))
			Expect(*result.Description).To(Equal(testGroupDescription))
		})
	})

	Describe("List access groups", func() {

		It("Successfully listed the account's access groups", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewListAccessGroupsOptions(testAccountID)
			options.SetHidePublicAccess(true)
			result, detailedResponse, err := iamAccessGroupsService.ListAccessGroups(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			// confirm the test group is present
			testGroupPresent := false
			for _, group := range result.Groups {
				if *group.ID == testGroupID {
					testGroupPresent = true
				}
			}
			Expect(testGroupPresent).To(BeTrue())
		})

		It(`ListAccessGroups(listAccessGroupsOptions *ListAccessGroupsOptions) using AccessGroupsPager`, func() {
			listAccessGroupsOptions := &iamaccessgroupsv2.ListAccessGroupsOptions{
				AccountID:        &testAccountID,
				HidePublicAccess: core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := iamAccessGroupsService.NewAccessGroupsPager(listAccessGroupsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []iamaccessgroupsv2.Group
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = iamAccessGroupsService.NewAccessGroupsPager(listAccessGroupsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListAccessGroups() returned a total of %d item(s) using AccessGroupsPager.\n", len(allResults))

			// confirm the test group is present
			testGroupPresent := false
			for _, group := range allResults {
				if *group.ID == testGroupID {
					testGroupPresent = true
				}
			}
			Expect(testGroupPresent).To(BeTrue())
		})
	})

	Describe("Add members to an access group", func() {

		It("Successfully added members to an access group", func() {
			shouldSkipTest()

			addMemberItem, err := iamAccessGroupsService.NewAddGroupMembersRequestMembersItem(testUserID, userType)
			Expect(err).To(BeNil())

			options := iamAccessGroupsService.NewAddMembersToAccessGroupOptions(testGroupID)
			options.Members = append(options.Members, *addMemberItem)
			result, detailedResponse, err := iamAccessGroupsService.AddMembersToAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(207))

			// confirm the test user is present
			testUserPresent := false
			for _, member := range result.Members {
				if *member.IamID == testUserID {
					testUserPresent = true
					Expect(*member.Type).To(Equal(userType))
					Expect(*member.StatusCode).To(Equal(int64(200)))
				}
			}
			Expect(testUserPresent).To(BeTrue())

		})

		It("Successfully added member to multiple access groups", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewAddMemberToMultipleAccessGroupsOptions(testAccountID, testUserID)
			options.SetType(userType)
			options.Groups = append(options.Groups, testGroupID)
			result, detailedResponse, err := iamAccessGroupsService.AddMemberToMultipleAccessGroups(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(207))

			// confirm the test user is present
			testGroupPresent := false
			for _, group := range result.Groups {
				if *group.AccessGroupID == testGroupID {
					testGroupPresent = true
					Expect(*group.StatusCode).To(Equal(int64(200)))
				}
			}
			Expect(testGroupPresent).To(BeTrue())

		})
	})

	Describe("Check access group membership", func() {

		It("Successfully checked the membership", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewIsMemberOfAccessGroupOptions(testGroupID, testUserID)
			detailedResponse, err := iamAccessGroupsService.IsMemberOfAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(204))
		})
	})

	Describe("List access group memberships", func() {

		It("Successfully listed the memberships", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewListAccessGroupMembersOptions(testGroupID)
			result, detailedResponse, err := iamAccessGroupsService.ListAccessGroupMembers(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			// confirm the test user is present
			testUserPresent := false
			for _, member := range result.Members {
				if *member.IamID == testUserID {
					testUserPresent = true
				}
			}
			Expect(testUserPresent).To(BeTrue())
		})
	})

	Describe("Delete access group membership", func() {

		It("Successfully deleted the membership", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewRemoveMemberFromAccessGroupOptions(testGroupID, testUserID)
			detailedResponse, err := iamAccessGroupsService.RemoveMemberFromAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(204))
		})
	})

	Describe("Delete member from all groups", func() {

		It("Returned that the membership was not found", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewRemoveMemberFromAllAccessGroupsOptions(testAccountID, testUserID)
			result, detailedResponse, err := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(options)
			Expect(err).To(Not(BeNil()))
			Expect(result).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})
	})

	Describe("Delete multiple members from access group", func() {

		It("Returned that the membership was not found", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewRemoveMembersFromAccessGroupOptions(testGroupID)
			options.Members = append(options.Members, testUserID)
			result, detailedResponse, err := iamAccessGroupsService.RemoveMembersFromAccessGroup(options)
			Expect(err).To(Not(BeNil()))
			Expect(result).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})
	})

	Describe("Create an access group rule", func() {

		It("Successfully created an access group rule", func() {
			shouldSkipTest()

			testExpiration := int64(24)
			condition, err := iamAccessGroupsService.NewRuleConditions("test claim", "EQUALS", "1")
			Expect(err).To(BeNil())

			options := iamAccessGroupsService.NewAddAccessGroupRuleOptions(testGroupID, testExpiration, "test realm name", []iamaccessgroupsv2.RuleConditions{*condition})

			result, detailedResponse, err := iamAccessGroupsService.AddAccessGroupRule(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(*result.AccessGroupID).To(Equal(testGroupID))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.Expiration).To(Equal(testExpiration))

			testClaimRuleID = *result.ID
		})
	})

	Describe("Get an access group rule", func() {

		It("Successfully retrieved the rule", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewGetAccessGroupRuleOptions(testGroupID, testClaimRuleID)
			result, detailedResponse, err := iamAccessGroupsService.GetAccessGroupRule(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testClaimRuleID))
			Expect(*result.AccessGroupID).To(Equal(testGroupID))
			Expect(*result.AccountID).To(Equal(testAccountID))

			testClaimRuleEtag = detailedResponse.GetHeaders().Get(etagHeader)
		})
	})

	Describe("List access group rules", func() {

		It("Successfully listed the rules", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewListAccessGroupRulesOptions(testGroupID)
			result, detailedResponse, err := iamAccessGroupsService.ListAccessGroupRules(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			// confirm the test user is present
			testClaimRulePresent := false
			for _, claimRule := range result.Rules {
				if *claimRule.ID == testClaimRuleID {
					testClaimRulePresent = true
				}
			}
			Expect(testClaimRulePresent).To(BeTrue())
		})
	})

	Describe("Update an access group rule", func() {

		It("Successfully updated the rule", func() {
			shouldSkipTest()

			testExpiration := int64(24)
			condition, err := iamAccessGroupsService.NewRuleConditions("test claim", "EQUALS", "1")
			Expect(err).To(BeNil())

			options := iamAccessGroupsService.NewReplaceAccessGroupRuleOptions(testGroupID, testClaimRuleID, testClaimRuleEtag, testExpiration, "updated test realm name", []iamaccessgroupsv2.RuleConditions{*condition})

			result, detailedResponse, err := iamAccessGroupsService.ReplaceAccessGroupRule(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testClaimRuleID))
			Expect(*result.AccessGroupID).To(Equal(testGroupID))
			Expect(*result.AccountID).To(Equal(testAccountID))
		})
	})

	Describe("Delete access group rule", func() {

		It("Successfully deleted the rule", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewRemoveAccessGroupRuleOptions(testGroupID, testClaimRuleID)
			detailedResponse, err := iamAccessGroupsService.RemoveAccessGroupRule(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(204))
		})
	})

	Describe("Get account settings", func() {

		It("Successfully retrieved the settings", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewGetAccountSettingsOptions(testAccountID)
			result, detailedResponse, err := iamAccessGroupsService.GetAccountSettings(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.AccountID).To(Equal(testAccountID))

			testAccountSettings = result
		})
	})

	Describe("Update account settings", func() {

		It("Successfully updated the settings", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewUpdateAccountSettingsOptions(testAccountID)
			options.SetPublicAccessEnabled(*testAccountSettings.PublicAccessEnabled)
			result, detailedResponse, err := iamAccessGroupsService.UpdateAccountSettings(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.PublicAccessEnabled).To(Equal(*options.PublicAccessEnabled))
		})
	})

	// clean up all test groups
	AfterSuite(func() {
		if !configLoaded {
			return
		}

		// list all groups in the account (minus the public access group)
		options := iamAccessGroupsService.NewListAccessGroupsOptions(testAccountID)
		options.SetHidePublicAccess(true)
		result, detailedResponse, err := iamAccessGroupsService.ListAccessGroups(options)
		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))

		// iterate across the groups
		for _, group := range result.Groups {

			// force delete the test group (or any test groups older than 5 minutes)
			if *group.Name == testGroupName {

				createdAt := time.Time(*group.CreatedAt)
				fiveMinutesAgo := time.Now().Add(-(time.Duration(5) * time.Minute))

				if *group.ID == testGroupID || createdAt.Before(fiveMinutesAgo) {
					options := iamAccessGroupsService.NewDeleteAccessGroupOptions(*group.ID)
					options.SetForce(true)
					detailedResponse, err := iamAccessGroupsService.DeleteAccessGroup(options)
					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(204))
				}
			}
		}
	})
})
