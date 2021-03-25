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

package iamidentityv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IAM Identity service.
//
// The following configuration properties are assumed to be defined:
//
// IAM_IDENTITY_URL=<service url>
// IAM_IDENTITY_AUTHTYPE=iam
// IAM_IDENTITY_AUTH_URL=<IAM Token Service url>
// IAM_IDENTITY_APIKEY=<IAM APIKEY for the User>
// IAM_IDENTITY_ACCOUNT_ID=<AccountID which is unique to the User>
// IAM_IDENTITY_IAM_ID=<IAM ID which is unique to the User account>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../iam_identity.env"

var (
	iamIdentityService *iamidentityv1.IamIdentityV1
	config             map[string]string
	configLoaded       bool = false

	err        error
	serviceURL string

	apikeyName    string = "Example-ApiKey"
	serviceIDName string = "Example-ServiceId"
	accountID     string
	iamID         string
	iamAPIKey     string

	apikeyID   string
	apikeyEtag string

	svcID     string
	svcIDEtag string

	accountSettingEtag string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`IamIdentityV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(iamidentityv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0

			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			iamID = config["IAM_ID"]
			Expect(iamID).ToNot(BeEmpty())

			iamAPIKey = config["APIKEY"]
			Expect(iamAPIKey).ToNot(BeEmpty())

			fmt.Printf("Service URL: %s\n", serviceURL)
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			iamIdentityServiceOptions := &iamidentityv1.IamIdentityV1Options{}

			iamIdentityService, err = iamidentityv1.NewIamIdentityV1UsingExternalConfig(iamIdentityServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(err).To(BeNil())
			Expect(iamIdentityService).ToNot(BeNil())
			Expect(iamIdentityService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`IamIdentityV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAPIKey request example`, func() {
			// begin-create_api_key

			createAPIKeyOptions := iamIdentityService.NewCreateAPIKeyOptions(apikeyName, iamID)
			createAPIKeyOptions.SetDescription("Example ApiKey")

			apiKey, response, err := iamIdentityService.CreateAPIKey(createAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Printf("\nCreateAPIKey() result:\n%s\n", string(b))
			apikeyID = *apiKey.ID

			// end-create_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())
			Expect(apikeyID).ToNot(BeNil())
		})
		It(`ListAPIKeys request example`, func() {
			// begin-list_api_keys

			listAPIKeysOptions := iamIdentityService.NewListAPIKeysOptions()
			listAPIKeysOptions.SetAccountID(accountID)
			listAPIKeysOptions.SetIamID(iamID)
			listAPIKeysOptions.SetIncludeHistory(true)

			apiKeyList, response, err := iamIdentityService.ListAPIKeys(listAPIKeysOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKeyList, "", "  ")
			fmt.Printf("\nListAPIKeys() result:\n%s\n", string(b))

			// end-list_api_keys

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKeyList).ToNot(BeNil())
		})
		It(`GetAPIKeysDetails request example`, func() {
			// begin-get_api_keys_details

			getAPIKeysDetailsOptions := iamIdentityService.NewGetAPIKeysDetailsOptions()
			getAPIKeysDetailsOptions.SetIamAPIKey(iamAPIKey)
			getAPIKeysDetailsOptions.SetIncludeHistory(false)

			apiKey, response, err := iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Printf("\nGetAPIKeysDetails() result:\n%s\n", string(b))

			// end-get_api_keys_details

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
		})
		It(`GetAPIKey request example`, func() {
			// begin-get_api_key

			getAPIKeyOptions := iamIdentityService.NewGetAPIKeyOptions(apikeyID)

			apiKey, response, err := iamIdentityService.GetAPIKey(getAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			apikeyEtag = response.GetHeaders().Get("Etag")
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Printf("\nGetAPIKey() result:\n%s\n", string(b))

			// end-get_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			Expect(apikeyEtag).ToNot(BeEmpty())
		})
		It(`UpdateAPIKey request example`, func() {
			// begin-update_api_key

			updateAPIKeyOptions := iamIdentityService.NewUpdateAPIKeyOptions(apikeyID, apikeyEtag)
			updateAPIKeyOptions.SetDescription("This is an updated description")

			apiKey, response, err := iamIdentityService.UpdateAPIKey(updateAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Printf("\nUpdateAPIKey() result:\n%s\n", string(b))

			// end-update_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
		})
		It(`LockAPIKey request example`, func() {
			// begin-lock_api_key

			lockAPIKeyOptions := iamIdentityService.NewLockAPIKeyOptions(apikeyID)

			response, err := iamIdentityService.LockAPIKey(lockAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nLockAPIKey() response status code: %d\n", response.StatusCode)

			// end-lock_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`UnlockAPIKey request example`, func() {
			// begin-unlock_api_key

			unlockAPIKeyOptions := iamIdentityService.NewUnlockAPIKeyOptions(apikeyID)

			response, err := iamIdentityService.UnlockAPIKey(unlockAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nUnlockAPIKey() response status code: %d\n", response.StatusCode)

			// end-unlock_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteAPIKey request example`, func() {
			// begin-delete_api_key

			deleteAPIKeyOptions := iamIdentityService.NewDeleteAPIKeyOptions(apikeyID)

			response, err := iamIdentityService.DeleteAPIKey(deleteAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteAPIKey() response status code: %d\n", response.StatusCode)

			// end-delete_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateServiceID request example`, func() {
			// begin-create_service_id

			createServiceIDOptions := iamIdentityService.NewCreateServiceIDOptions(accountID, serviceIDName)
			createServiceIDOptions.SetDescription("Example ServiceId")

			serviceID, response, err := iamIdentityService.CreateServiceID(createServiceIDOptions)
			if err != nil {
				panic(err)
			}
			svcID = *serviceID.ID
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Printf("\nCreateServiceID() result:\n%s\n", string(b))

			// end-create_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(serviceID).ToNot(BeNil())
			Expect(svcID).ToNot(BeEmpty())
		})
		It(`GetServiceID request example`, func() {
			// begin-get_service_id

			getServiceIDOptions := iamIdentityService.NewGetServiceIDOptions(svcID)

			serviceID, response, err := iamIdentityService.GetServiceID(getServiceIDOptions)
			if err != nil {
				panic(err)
			}
			svcIDEtag = response.GetHeaders().Get("Etag")
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Printf("\nGetServiceID() result:\n%s\n", string(b))

			// end-get_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			Expect(svcIDEtag).ToNot(BeEmpty())
		})
		It(`ListServiceIds request example`, func() {
			// begin-list_service_ids

			listServiceIdsOptions := iamIdentityService.NewListServiceIdsOptions()
			listServiceIdsOptions.SetAccountID(accountID)
			listServiceIdsOptions.SetName(serviceIDName)

			serviceIDList, response, err := iamIdentityService.ListServiceIds(listServiceIdsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceIDList, "", "  ")
			fmt.Printf("\nListServiceIds() result:\n%s\n", string(b))

			// end-list_service_ids

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIDList).ToNot(BeNil())
		})
		It(`UpdateServiceID request example`, func() {
			// begin-update_service_id

			updateServiceIDOptions := iamIdentityService.NewUpdateServiceIDOptions(svcID, svcIDEtag)
			updateServiceIDOptions.SetDescription("This is an updated description")

			serviceID, response, err := iamIdentityService.UpdateServiceID(updateServiceIDOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Printf("\nUpdateServiceID() result:\n%s\n", string(b))

			// end-update_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
		})
		It(`LockServiceID request example`, func() {
			// begin-lock_service_id

			lockServiceIDOptions := iamIdentityService.NewLockServiceIDOptions(svcID)

			response, err := iamIdentityService.LockServiceID(lockServiceIDOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nLockServiceID() response status code: %d\n", response.StatusCode)

			// end-lock_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`UnlockServiceID request example`, func() {
			// begin-unlock_service_id

			unlockServiceIDOptions := iamIdentityService.NewUnlockServiceIDOptions(svcID)

			response, err := iamIdentityService.UnlockServiceID(unlockServiceIDOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nUnlockServiceID() response status code: %d\n", response.StatusCode)

			// end-unlock_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteServiceID request example`, func() {
			// begin-delete_service_id

			deleteServiceIDOptions := iamIdentityService.NewDeleteServiceIDOptions(svcID)

			response, err := iamIdentityService.DeleteServiceID(deleteServiceIDOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteServiceID() response status code: %d\n", response.StatusCode)

			// end-delete_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`GetAccountSettings request example`, func() {
			// begin-getAccountSettings

			getAccountSettingsOptions := iamIdentityService.NewGetAccountSettingsOptions(accountID)

			accountSettingsResponse, response, err := iamIdentityService.GetAccountSettings(getAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettingsResponse, "", "  ")
			fmt.Printf("\nGetAccountSettings() result:\n%s\n", string(b))

			// end-getAccountSettings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsResponse).ToNot(BeNil())

			accountSettingEtag = response.GetHeaders().Get("Etag")
			Expect(accountSettingEtag).ToNot(BeEmpty())
		})
		It(`UpdateAccountSettings request example`, func() {
			// begin-updateAccountSettings

			updateAccountSettingsOptions := iamIdentityService.NewUpdateAccountSettingsOptions(
				accountSettingEtag,
				accountID,
			)
			updateAccountSettingsOptions.SetSessionExpirationInSeconds("86400")
			updateAccountSettingsOptions.SetSessionInvalidationInSeconds("7200")
			updateAccountSettingsOptions.SetMfa("NONE")
			updateAccountSettingsOptions.SetRestrictCreatePlatformApikey("NOT_RESTRICTED")
			updateAccountSettingsOptions.SetRestrictCreatePlatformApikey("NOT_RESTRICTED")

			accountSettingsResponse, response, err := iamIdentityService.UpdateAccountSettings(updateAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettingsResponse, "", "  ")
			fmt.Printf("\nUpdateAccountSettings() result:\n%s\n", string(b))

			// end-updateAccountSettings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsResponse).ToNot(BeNil())
		})
	})
})
