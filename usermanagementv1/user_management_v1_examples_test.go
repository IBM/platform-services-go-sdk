// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/usermanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const externalConfigFile = "../user_management.env"

var (
	userManagementService *usermanagementv1.UserManagementV1
	config                map[string]string
	configLoaded          bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`UserManagementV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(usermanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

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

			userManagementServiceOptions := &usermanagementv1.UserManagementV1Options{}

			userManagementService, err = usermanagementv1.NewUserManagementV1UsingExternalConfig(userManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(userManagementService).ToNot(BeNil())
		})
	})

	Describe(`UserManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InviteUsers request example`, func() {
			// begin-invite_users

			inviteUsersOptions := userManagementService.NewInviteUsersOptions(
				"testString",
			)

			invitedUserList, response, err := userManagementService.InviteUsers(inviteUsersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(invitedUserList, "", "  ")
			fmt.Println(string(b))

			// end-invite_users

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(invitedUserList).ToNot(BeNil())

		})
		It(`ListUsers request example`, func() {
			// begin-list_users

			listUsersOptions := userManagementService.NewListUsersOptions(
				"testString",
			)

			userList, response, err := userManagementService.ListUsers(listUsersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(userList, "", "  ")
			fmt.Println(string(b))

			// end-list_users

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userList).ToNot(BeNil())

		})
		It(`RemoveUser request example`, func() {
			// begin-remove_user

			removeUserOptions := userManagementService.NewRemoveUserOptions(
				"testString",
				"testString",
			)

			response, err := userManagementService.RemoveUser(removeUserOptions)
			if err != nil {
				panic(err)
			}

			// end-remove_user

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`GetUserProfile request example`, func() {
			// begin-get_user_profile

			getUserProfileOptions := userManagementService.NewGetUserProfileOptions(
				"testString",
				"testString",
			)

			userProfile, response, err := userManagementService.GetUserProfile(getUserProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(userProfile, "", "  ")
			fmt.Println(string(b))

			// end-get_user_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userProfile).ToNot(BeNil())

		})
		It(`UpdateUserProfile request example`, func() {
			// begin-update_user_profile

			updateUserProfileOptions := userManagementService.NewUpdateUserProfileOptions(
				"testString",
				"testString",
			)

			response, err := userManagementService.UpdateUserProfile(updateUserProfileOptions)
			if err != nil {
				panic(err)
			}

			// end-update_user_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`GetUserSettings request example`, func() {
			// begin-get_user_settings

			getUserSettingsOptions := userManagementService.NewGetUserSettingsOptions(
				"testString",
				"testString",
			)

			userSettings, response, err := userManagementService.GetUserSettings(getUserSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(userSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_user_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userSettings).ToNot(BeNil())

		})
		It(`UpdateUserSettings request example`, func() {
			// begin-update_user_settings

			updateUserSettingsOptions := userManagementService.NewUpdateUserSettingsOptions(
				"testString",
				"testString",
			)

			response, err := userManagementService.UpdateUserSettings(updateUserSettingsOptions)
			if err != nil {
				panic(err)
			}

			// end-update_user_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
