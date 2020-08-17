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

package usermanagementv1_test

import (
	"fmt"
	"os"
	"github.com/IBM/go-sdk-core/v4/core"
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
		deleteUserId          string
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
			config, err = core.GetServiceProperties("USERMGMT1")
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["AM_HOST"]
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
        It("Successfully construct the service client main instance", func() {
            userManagementServiceOptions := &usermanagementv1.UserManagementV1Options{
                ServiceName: "USERMGMT1",
            }
            userManagementService, err = usermanagementv1.NewUserManagementV1UsingExternalConfig(userManagementServiceOptions)
            Expect(err).To(BeNil())
            Expect(userManagementService).ToNot(BeNil())
            Expect(userManagementService.Service.Options.URL).To(Equal(serviceURL))
        })
        It("Successfully construct the service client alternate instance", func() {
            userManagementServiceOptions := &usermanagementv1.UserManagementV1Options{
                ServiceName: "USERMGMT2",
            }
            alternateService, err = usermanagementv1.NewUserManagementV1UsingExternalConfig(userManagementServiceOptions)
            Expect(err).To(BeNil())
            Expect(alternateService).ToNot(BeNil())
            Expect(alternateService.Service.Options.URL).To(Equal(serviceURL))
        })
    })

	Describe(`GetUserSettings - Get user settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUserSettings(getUserSettingsOptions *GetUserSettingsOptions)`, func() {

			getUserSettingsOptions := &usermanagementv1.GetUserSettingsOptions{
				AccountID: core.StringPtr("1aa434630b594b8a88b961a44c9eb2a9"),
				IamID:     core.StringPtr("IBMid-5500089E4W"),
			}

			userSettings, response, err := userManagementService.GetUserSettings(getUserSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userSettings).ToNot(BeNil())

		})
	})

	Describe(`UpdateUserSettings - Partially update user settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateUserSettings(updateUserSettingsOptions *UpdateUserSettingsOptions)`, func() {

			updateUserSettingsOptions := &usermanagementv1.UpdateUserSettingsOptions{
				AccountID:            core.StringPtr("1aa434630b594b8a88b961a44c9eb2a9"),
				IamID:                core.StringPtr("IBMid-5500089E4W"),
				Language:             core.StringPtr("testString"),
				NotificationLanguage: core.StringPtr("testString"),
				AllowedIpAddresses:   core.StringPtr("32.96.110.50,172.16.254.1"),
				SelfManage:           core.BoolPtr(true),
			}

			userSettings, response, err := userManagementService.UpdateUserSettings(updateUserSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
			Expect(userSettings).To(BeNil())
		})
	})

	Describe(`ListUsers - List users`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListUsers(listUsersOptions *ListUsersOptions)`, func() {

			listUsersOptions := &usermanagementv1.ListUsersOptions{
				AccountID: core.StringPtr("1aa434630b594b8a88b961a44c9eb2a9"),
				State:     core.StringPtr("testString"),
			}
			userList, response, err := userManagementService.ListUsers(listUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userList).ToNot(BeNil())
		})
	})

	Describe(`InviteUsers - Invite users`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`InviteUsers(inviteUsersOptions *InviteUsersOptions)`, func() {

			inviteUserModel := &usermanagementv1.InviteUser{
				Email:       core.StringPtr("aminttest+linked_account_owner_11@mail.test.ibm.com"),
				AccountRole: core.StringPtr("Member"),
			}

			roleModel := &usermanagementv1.Role{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Viewer"),
			}

			attributeModel := &usermanagementv1.Attribute{
				Name:  core.StringPtr("accountId"),
				Value: core.StringPtr("1aa434630b594b8a88b961a44c9eb2a9"),
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
				AccountID:    core.StringPtr("1aa434630b594b8a88b961a44c9eb2a9"),
				Users:        []usermanagementv1.InviteUser{*inviteUserModel},
				IamPolicy:    []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel},
				AccessGroups: []string{"AccessGroupId-51675919-2bd7-4ce3-86e4-5faff8065574"},
			}

			userList, response, err := alternateService.InviteUsers(inviteUsersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(userList).ToNot(BeNil())

			for _, res := range userList.Resources {
				deleteUserId = *res.ID
			}
		})
	})

	Describe(`GetUserProfile - Get user profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUserProfile(getUserProfileOptions *GetUserProfileOptions)`, func() {

			getUserProfileOptions := &usermanagementv1.GetUserProfileOptions{
				AccountID: core.StringPtr("1aa434630b594b8a88b961a44c9eb2a9"),
				IamID:     core.StringPtr("IBMid-5500089E4W"),
			}

			userProfile, response, err := userManagementService.GetUserProfile(getUserProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userProfile).ToNot(BeNil())
		})
	})

	Describe(`UpdateUserProfiles - Partially update user profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateUserProfiles(updateUserProfilesOptions *UpdateUserProfilesOptions)`, func() {

			updateUserProfilesOptions := &usermanagementv1.UpdateUserProfilesOptions{
				AccountID:      core.StringPtr("1aa434630b594b8a88b961a44c9eb2a9"),
				IamID:          core.StringPtr("IBMid-5500089E4W"),
				Firstname:      core.StringPtr("testString"),
				Lastname:       core.StringPtr("testString"),
				State:          core.StringPtr("ACTIVE"),
				Email:          core.StringPtr("testString"),
				Phonenumber:    core.StringPtr("testString"),
				Altphonenumber: core.StringPtr("testString"),
				Photo:          core.StringPtr("testString"),
			}

			response, err := userManagementService.UpdateUserProfiles(updateUserProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`RemoveUsers - Remove users`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveUsers(removeUsersOptions *RemoveUsersOptions)`, func() {

			removeUsersOptions := &usermanagementv1.RemoveUsersOptions{
				AccountID: core.StringPtr("1aa434630b594b8a88b961a44c9eb2a9"),
				IamID:     core.StringPtr(deleteUserId),
			}

			fmt.Println(*removeUsersOptions)

			response, err := userManagementService.RemoveUsers(removeUsersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
