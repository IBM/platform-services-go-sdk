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
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/usermanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

/**
 * This file contains an integration test for the usermanagementv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`UserManagementV1 Integration Tests`, func() {

	const externalConfigFile = "../user_management_v1.env"

	var (
		err          error
		userManagementService *usermanagementv1.UserManagementV1
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
			config, err = core.GetServiceProperties(usermanagementv1.DefaultServiceName)
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

			userManagementServiceOptions := &usermanagementv1.UserManagementV1Options{}

			userManagementService, err = usermanagementv1.NewUserManagementV1UsingExternalConfig(userManagementServiceOptions)

			Expect(err).To(BeNil())
			Expect(userManagementService).ToNot(BeNil())
			Expect(userManagementService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`GetUserSettings - Get user settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUserSettings(getUserSettingsOptions *GetUserSettingsOptions)`, func() {

			getUserSettingsOptions := &usermanagementv1.GetUserSettingsOptions{
				AccountID: core.StringPtr("testString"),
				IamID: core.StringPtr("testString"),
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
				AccountID: core.StringPtr("testString"),
				IamID: core.StringPtr("testString"),
				Language: core.StringPtr("testString"),
				NotificationLanguage: core.StringPtr("testString"),
				AllowedIpAddresses: core.StringPtr("32.96.110.50,172.16.254.1"),
				SelfManage: core.BoolPtr(true),
			}

			userSettings, response, err := userManagementService.UpdateUserSettings(updateUserSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userSettings).ToNot(BeNil())

		})
	})

	Describe(`ListUsers - List users`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListUsers(listUsersOptions *ListUsersOptions)`, func() {

			listUsersOptions := &usermanagementv1.ListUsersOptions{
				AccountID: core.StringPtr("testString"),
				State: core.StringPtr("testString"),
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
				Email: core.StringPtr("testString"),
				AccountRole: core.StringPtr("testString"),
			}

			roleModel := &usermanagementv1.Role{
				RoleID: core.StringPtr("testString"),
			}

			attributeModel := &usermanagementv1.Attribute{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			resourceModel := &usermanagementv1.Resource{
				Attributes: []usermanagementv1.Attribute{*attributeModel},
			}

			inviteUserIamPolicyModel := &usermanagementv1.InviteUserIamPolicy{
				Roles: []usermanagementv1.Role{*roleModel},
				Resources: []usermanagementv1.Resource{*resourceModel},
			}

			inviteUsersOptions := &usermanagementv1.InviteUsersOptions{
				AccountID: core.StringPtr("testString"),
				Users: []usermanagementv1.InviteUser{*inviteUserModel},
				IamPolicy: []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel},
				AccessGroups: []string{"testString"},
			}

			userList, response, err := userManagementService.InviteUsers(inviteUsersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(userList).ToNot(BeNil())

		})
	})

	Describe(`GetUserProfile - Get user profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUserProfile(getUserProfileOptions *GetUserProfileOptions)`, func() {

			getUserProfileOptions := &usermanagementv1.GetUserProfileOptions{
				AccountID: core.StringPtr("testString"),
				IamID: core.StringPtr("testString"),
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
				AccountID: core.StringPtr("testString"),
				IamID: core.StringPtr("testString"),
				Firstname: core.StringPtr("testString"),
				Lastname: core.StringPtr("testString"),
				State: core.StringPtr("testString"),
				Email: core.StringPtr("testString"),
				Phonenumber: core.StringPtr("testString"),
				Altphonenumber: core.StringPtr("testString"),
				Photo: core.StringPtr("testString"),
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
				AccountID: core.StringPtr("testString"),
				IamID: core.StringPtr("testString"),
			}

			response, err := userManagementService.RemoveUsers(removeUsersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
