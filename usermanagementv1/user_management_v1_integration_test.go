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

package usermanagementv1_test

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/usermanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the usermanagementv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`UserManagementV1 Integration Tests`, func() {

	const externalConfigFile = "../user_management.env"

	var (
		err                   error
		userManagementService *usermanagementv1.UserManagementV1
		alternateService      *usermanagementv1.UserManagementV1
		serviceURL            string
		config                map[string]string
		deleteUserID          string
		accountID             = "1aa434630b594b8a88b961a44c9eb2a9"
		userID                = "IBMid-550009SDCP"
		memberEmail           = "aminttest+linked_account_owner_11@mail.test.ibm.com"
		viewerRoleID          = "crn:v1:bluemix:public:iam::::role:Viewer"
		accessGroupID         = "AccessGroupId-51675919-2bd7-4ce3-86e4-5faff8065574"
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
			config, err = core.GetServiceProperties(usermanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client main instance", func() {
			userManagementServiceOptions := &usermanagementv1.UserManagementV1Options{
				ServiceName: usermanagementv1.DefaultServiceName,
			}
			userManagementService, err = usermanagementv1.NewUserManagementV1UsingExternalConfig(userManagementServiceOptions)
			Expect(err).To(BeNil())
			Expect(userManagementService).ToNot(BeNil())
			Expect(userManagementService.Service.Options.URL).To(Equal(serviceURL))
			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			userManagementService.EnableRetries(4, 30*time.Second)
		})
		It("Successfully construct the service client alternate instance", func() {
			userManagementServiceOptions := &usermanagementv1.UserManagementV1Options{
				ServiceName: "USERMGMT2",
			}
			alternateService, err = usermanagementv1.NewUserManagementV1UsingExternalConfig(userManagementServiceOptions)
			Expect(err).To(BeNil())
			Expect(alternateService).ToNot(BeNil())
			Expect(alternateService.Service.Options.URL).To(Equal(serviceURL))
			alternateService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetUserSettings - Get user settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUserSettings(getUserSettingsOptions *GetUserSettingsOptions)`, func() {

			getUserSettingsOptions := &usermanagementv1.GetUserSettingsOptions{
				AccountID: &accountID,
				IamID:     &userID,
			}

			result, response, err := userManagementService.GetUserSettings(getUserSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetUserSettings() result:\n%s\n", common.ToJSON(result))
		})
	})

	Describe(`UpdateUserSettings - Partially update user settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateUserSettings(updateUserSettingsOptions *UpdateUserSettingsOptions)`, func() {

			updateUserSettingsOptions := &usermanagementv1.UpdateUserSettingsOptions{
				AccountID:          &accountID,
				IamID:              &userID,
				AllowedIPAddresses: core.StringPtr("32.96.110.50,172.16.254.1"),
				SelfManage:         core.BoolPtr(true),
			}

			response, err := userManagementService.UpdateUserSettings(updateUserSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`ListUsers - List users`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListUsers(listUsersOptions *ListUsersOptions)`, func() {

			results := []usermanagementv1.UserProfile{}
			var moreResults bool = true
			var pageStart *string = nil

			for moreResults {
				listUsersOptions := &usermanagementv1.ListUsersOptions{
					AccountID: &accountID,
					Limit:     core.Int64Ptr(10),
					Start:     pageStart,
				}
				result, response, err := userManagementService.ListUsers(listUsersOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())
				// fmt.Fprintf(GinkgoWriter, "ListUsers() result:\n%s\n", common.ToJSON(result))

				for _, userProfile := range result.Resources {
					results = append(results, userProfile)
				}

				if result.NextURL != nil {
					pageStart = getStartTokenFromURL(result.NextURL)
				} else {
					moreResults = false
				}
			}

			fmt.Fprintf(GinkgoWriter, "Received a total of %d user profiles.\n", len(results))
		})
	})

	Describe(`InviteUsers - Invite users`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`InviteUsers(inviteUsersOptions *InviteUsersOptions)`, func() {

			inviteUserModel := &usermanagementv1.InviteUser{
				Email:       &memberEmail,
				AccountRole: core.StringPtr("Member"),
			}

			roleModel := &usermanagementv1.Role{
				RoleID: &viewerRoleID,
			}

			attributeModel := &usermanagementv1.Attribute{
				Name:  core.StringPtr("accountId"),
				Value: &accountID,
			}

			attributeModel2 := &usermanagementv1.Attribute{
				Name:  core.StringPtr("resourceGroupId"),
				Value: core.StringPtr("*"),
			}

			resourceModel := &usermanagementv1.Resource{
				Attributes: []usermanagementv1.Attribute{*attributeModel, *attributeModel2},
			}

			inviteUserIamPolicyModel := &usermanagementv1.InviteUserIamPolicy{
				Type:      core.StringPtr("access"),
				Roles:     []usermanagementv1.Role{*roleModel},
				Resources: []usermanagementv1.Resource{*resourceModel},
			}

			inviteUsersOptions := &usermanagementv1.InviteUsersOptions{
				AccountID:    &accountID,
				Users:        []usermanagementv1.InviteUser{*inviteUserModel},
				IamPolicy:    []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel},
				AccessGroups: []string{accessGroupID},
			}

			result, response, err := alternateService.InviteUsers(inviteUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "InviteUsers() result:\n%s\n", common.ToJSON(result))

			for _, res := range result.Resources {
				deleteUserID = *res.ID
			}
		})
	})

	Describe(`GetUserProfile - Get user profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUserProfile(getUserProfileOptions *GetUserProfileOptions)`, func() {

			getUserProfileOptions := &usermanagementv1.GetUserProfileOptions{
				AccountID: &accountID,
				IamID:     &userID,
			}

			result, response, err := userManagementService.GetUserProfile(getUserProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetUserProfile() result:\n%s\n", common.ToJSON(result))
		})
	})

	Describe(`UpdateUserProfile - Partially update user profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateUserProfile(updateUserProfileOptions *UpdateUserProfileOptions)`, func() {

			updateUserProfilesOptions := &usermanagementv1.UpdateUserProfileOptions{
				AccountID: &accountID,
				IamID:     &userID,
				State:     core.StringPtr("ACTIVE"),
			}

			response, err := userManagementService.UpdateUserProfile(updateUserProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`RemoveUser - Remove user from account`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveUser(removeUserOptions *RemoveUserOptions)`, func() {
			Expect(deleteUserID).ToNot(BeEmpty())

			removeUserOptions := &usermanagementv1.RemoveUserOptions{
				AccountID: &accountID,
				IamID:     &deleteUserID,
			}

			response, err := userManagementService.RemoveUser(removeUserOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

func getStartTokenFromURL(sptr *string) *string {
	if sptr == nil {
		return nil
	}

	s := *sptr
	if s == "" {
		return nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return nil
	}

	if u.RawQuery == "" {
		return nil
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil
	}

	token := q.Get("_start")
	if token == "" {
		return nil
	}
	return &token
}
