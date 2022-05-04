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
	profileName   string = "Example-Profile"
	accountID     string
	iamID         string
	iamAPIKey     string

	apikeyID   string
	apikeyEtag string

	svcID     string
	svcIDEtag string

	profileId     string
	profileIamId  string
	profileEtag   string
	claimRuleId   string
	claimRuleEtag string
	claimRuleType string = "Profile-SAML"
	realmName     string = "https://w3id.sso.ibm.com/auth/sps/samlidp2/saml20"
	linkId        string

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
			fmt.Println("\nCreateAPIKey() result:")
			// begin-create_api_key

			createAPIKeyOptions := iamIdentityService.NewCreateAPIKeyOptions(apikeyName, iamID)
			createAPIKeyOptions.SetDescription("Example ApiKey")

			apiKey, response, err := iamIdentityService.CreateAPIKey(createAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Println(string(b))
			apikeyID = *apiKey.ID

			// end-create_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())
			Expect(apikeyID).ToNot(BeNil())
		})
		It(`ListAPIKeys request example`, func() {
			fmt.Println("\nListAPIKeys() result:")
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
			fmt.Println(string(b))

			// end-list_api_keys

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKeyList).ToNot(BeNil())
		})
		It(`GetAPIKeysDetails request example`, func() {
			fmt.Println("\nGetAPIKeysDetails() result:")
			// begin-get_api_keys_details

			getAPIKeysDetailsOptions := iamIdentityService.NewGetAPIKeysDetailsOptions()
			getAPIKeysDetailsOptions.SetIamAPIKey(iamAPIKey)
			getAPIKeysDetailsOptions.SetIncludeHistory(false)

			apiKey, response, err := iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Println(string(b))

			// end-get_api_keys_details

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
		})
		It(`GetAPIKey request example`, func() {
			fmt.Println("\nGetAPIKey() result:")
			// begin-get_api_key

			getAPIKeyOptions := iamIdentityService.NewGetAPIKeyOptions(apikeyID)

			getAPIKeyOptions.SetIncludeHistory(false)
			getAPIKeyOptions.SetIncludeActivity(false)

			apiKey, response, err := iamIdentityService.GetAPIKey(getAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			apikeyEtag = response.GetHeaders().Get("Etag")
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Println(string(b))

			// end-get_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			Expect(apikeyEtag).ToNot(BeEmpty())
		})
		It(`UpdateAPIKey request example`, func() {
			fmt.Println("\nUpdateAPIKey() result:")
			// begin-update_api_key

			updateAPIKeyOptions := iamIdentityService.NewUpdateAPIKeyOptions(apikeyID, apikeyEtag)
			updateAPIKeyOptions.SetDescription("This is an updated description")

			apiKey, response, err := iamIdentityService.UpdateAPIKey(updateAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Println(string(b))

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

			// end-lock_api_key
			fmt.Printf("\nLockAPIKey() response status code: %d\n", response.StatusCode)

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

			// end-unlock_api_key
			fmt.Printf("\nUnlockAPIKey() response status code: %d\n", response.StatusCode)

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

			// end-delete_api_key
			fmt.Printf("\nDeleteAPIKey() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateServiceID request example`, func() {
			fmt.Println("\nCreateServiceID() result:")
			// begin-create_service_id

			createServiceIDOptions := iamIdentityService.NewCreateServiceIDOptions(accountID, serviceIDName)
			createServiceIDOptions.SetDescription("Example ServiceId")

			serviceID, response, err := iamIdentityService.CreateServiceID(createServiceIDOptions)
			if err != nil {
				panic(err)
			}
			svcID = *serviceID.ID
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Println(string(b))

			// end-create_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(serviceID).ToNot(BeNil())
			Expect(svcID).ToNot(BeEmpty())
		})
		It(`GetServiceID request example`, func() {
			fmt.Println("\nGetServiceID() result:")
			// begin-get_service_id

			getServiceIDOptions := iamIdentityService.NewGetServiceIDOptions(svcID)

			getServiceIDOptions.SetIncludeActivity(false)

			serviceID, response, err := iamIdentityService.GetServiceID(getServiceIDOptions)
			if err != nil {
				panic(err)
			}
			svcIDEtag = response.GetHeaders().Get("Etag")
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Println(string(b))

			// end-get_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			Expect(svcIDEtag).ToNot(BeEmpty())
		})
		It(`ListServiceIds request example`, func() {
			fmt.Println("\nListServiceIds() result:")
			// begin-list_service_ids

			listServiceIdsOptions := iamIdentityService.NewListServiceIdsOptions()
			listServiceIdsOptions.SetAccountID(accountID)
			listServiceIdsOptions.SetName(serviceIDName)

			serviceIDList, response, err := iamIdentityService.ListServiceIds(listServiceIdsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceIDList, "", "  ")
			fmt.Println(string(b))

			// end-list_service_ids

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIDList).ToNot(BeNil())
		})
		It(`UpdateServiceID request example`, func() {
			fmt.Println("\nUpdateServiceID() result:")
			// begin-update_service_id

			updateServiceIDOptions := iamIdentityService.NewUpdateServiceIDOptions(svcID, svcIDEtag)
			updateServiceIDOptions.SetDescription("This is an updated description")

			serviceID, response, err := iamIdentityService.UpdateServiceID(updateServiceIDOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Println(string(b))

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

			// end-lock_service_id
			fmt.Printf("\nLockServiceID() response status code: %d\n", response.StatusCode)

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

			// end-unlock_service_id
			fmt.Printf("\nUnlockServiceID() response status code: %d\n", response.StatusCode)

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

			// end-delete_service_id
			fmt.Printf("\nDeleteServiceID() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateProfile request example`, func() {
			fmt.Println("\nCreateProfile() result:")
			// begin-create_profile

			createProfileOptions := iamIdentityService.NewCreateProfileOptions(profileName, accountID)
			createProfileOptions.SetDescription("Example Profile")

			profile, response, err := iamIdentityService.CreateProfile(createProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))
			profileId = *profile.ID

			// end-create_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profile).ToNot(BeNil())
			Expect(profileId).ToNot(BeNil())
		})
		It(`GetProfile request example`, func() {
			fmt.Println("\nGetProfile() result:")
			// begin-get_profile

			getProfileOptions := iamIdentityService.NewGetProfileOptions(profileId)

			getProfileOptions.SetIncludeActivity(false)

			profile, response, err := iamIdentityService.GetProfile(getProfileOptions)
			if err != nil {
				panic(err)
			}
			profileEtag = response.GetHeaders().Get("Etag")
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-get_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
			Expect(profileEtag).ToNot(BeEmpty())
		})
		It(`ListProfiles request example`, func() {
			fmt.Println("\nListProfiles() result:")
			// begin-list_profiles

			listProfilesOptions := iamIdentityService.NewListProfilesOptions(accountID)
			listProfilesOptions.SetIncludeHistory(false)

			trustedProfiles, response, err := iamIdentityService.ListProfiles(listProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trustedProfiles, "", "  ")
			fmt.Println(string(b))

			// end-list_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trustedProfiles).ToNot(BeNil())
		})
		It(`UpdateProfile request example`, func() {
			fmt.Println("\nUpdateProfile() result:")
			// begin-update_profile

			updateProfileOptions := iamIdentityService.NewUpdateProfileOptions(profileId, profileEtag)
			updateProfileOptions.SetDescription("This is an updated description")

			profile, response, err := iamIdentityService.UpdateProfile(updateProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-update_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
		It(`CreateClaimRule request example`, func() {
			fmt.Println("\nCreateClaimRule() result:")
			// begin-create_claim_rule

			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("blueGroups")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"cloud-docs-dev\"")

			createClaimRuleOptions := iamIdentityService.NewCreateClaimRuleOptions(profileId, claimRuleType, []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions})
			createClaimRuleOptions.SetName("claimRule")
			createClaimRuleOptions.SetRealmName(realmName)
			createClaimRuleOptions.SetExpiration(int64(43200))

			claimRule, response, err := iamIdentityService.CreateClaimRule(createClaimRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(claimRule, "", "  ")
			fmt.Println(string(b))
			claimRuleId = *claimRule.ID

			// end-create_claim_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(claimRule).ToNot(BeNil())
			Expect(claimRuleId).ToNot(BeNil())
		})
		It(`GetClaimRule request example`, func() {
			fmt.Println("\nGetClaimRule() result:")
			// begin-get_claim_rule

			getClaimRuleOptions := iamIdentityService.NewGetClaimRuleOptions(profileId, claimRuleId)

			claimRule, response, err := iamIdentityService.GetClaimRule(getClaimRuleOptions)
			if err != nil {
				panic(err)
			}
			claimRuleEtag = response.GetHeaders().Get("Etag")
			b, _ := json.MarshalIndent(claimRule, "", "  ")
			fmt.Println(string(b))

			// end-get_claim_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(claimRule).ToNot(BeNil())
			Expect(claimRuleEtag).ToNot(BeEmpty())
		})
		It(`ListClaimRules request example`, func() {
			fmt.Println("\nListClaimRules() result:")
			// begin-list_claim_rules

			listClaimRulesOptions := iamIdentityService.NewListClaimRulesOptions(profileId)

			claimRulesList, response, err := iamIdentityService.ListClaimRules(listClaimRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(claimRulesList, "", "  ")
			fmt.Println(string(b))

			// end-list_claim_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(claimRulesList).ToNot(BeNil())
		})
		It(`UpdateClaimRule request example`, func() {
			fmt.Println("\nUpdateClaimRule() result:")
			// begin-update_claim_rule

			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("blueGroups")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"Europe_Group\"")

			updateClaimRuleOptions := iamIdentityService.NewUpdateClaimRuleOptions(profileId, claimRuleId, claimRuleEtag, claimRuleType, []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions})
			updateClaimRuleOptions.SetRealmName(realmName)
			updateClaimRuleOptions.SetExpiration(int64(33200))

			claimRule, response, err := iamIdentityService.UpdateClaimRule(updateClaimRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(claimRule, "", "  ")
			fmt.Println(string(b))

			// end-update_claim_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(claimRule).ToNot(BeNil())
		})
		It(`DeleteClaimRule request example`, func() {
			// begin-delete_claim_rule

			deleteClaimRuleOptions := iamIdentityService.NewDeleteClaimRuleOptions(profileId, claimRuleId)

			response, err := iamIdentityService.DeleteClaimRule(deleteClaimRuleOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_claim_rule
			fmt.Printf("\nDeleteClaimRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateLink request example`, func() {
			fmt.Println("\nCreateLink() result:")
			// begin-create_link

			createProfileLinkRequestLink := new(iamidentityv1.CreateProfileLinkRequestLink)
			createProfileLinkRequestLink.CRN = core.StringPtr("crn:v1:staging:public:iam-identity::a/" + accountID + "::computeresource:Fake-Compute-Resource")
			createProfileLinkRequestLink.Namespace = core.StringPtr("default")
			createProfileLinkRequestLink.Name = core.StringPtr("niceName")

			createLinkOptions := iamIdentityService.NewCreateLinkOptions(profileId, "ROKS_SA", createProfileLinkRequestLink)
			createLinkOptions.SetName("niceLink")

			link, response, err := iamIdentityService.CreateLink(createLinkOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(link, "", "  ")
			fmt.Println(string(b))
			linkId = *link.ID

			// end-create_link

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(link).ToNot(BeNil())
			Expect(linkId).ToNot(BeNil())
		})
		It(`GetLink request example`, func() {
			fmt.Println("\nGetLink() result:")
			// begin-get_link

			getLinkOptions := iamIdentityService.NewGetLinkOptions(profileId, linkId)

			link, response, err := iamIdentityService.GetLink(getLinkOptions)
			if err != nil {
				panic(err)
			}

			// end-get_link

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(link).ToNot(BeNil())
		})
		It(`ListLinks request example`, func() {
			fmt.Println("\nListLinks() result:")
			// begin-list_links

			listLinksOptions := iamIdentityService.NewListLinksOptions(profileId)

			linkList, response, err := iamIdentityService.ListLinks(listLinksOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(linkList, "", "  ")
			fmt.Println(string(b))

			// end-list_links

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(linkList).ToNot(BeNil())
		})
		It(`DeleteLink request example`, func() {
			// begin-delete_link

			deleteLinkOptions := iamIdentityService.NewDeleteLinkOptions(profileId, linkId)

			response, err := iamIdentityService.DeleteLink(deleteLinkOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_link
			fmt.Printf("\nDeleteLink() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteProfile request example`, func() {
			// begin-delete_profile

			deleteProfileOptions := iamIdentityService.NewDeleteProfileOptions(profileId)

			response, err := iamIdentityService.DeleteProfile(deleteProfileOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_profile
			fmt.Printf("\nDeleteProfile() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`GetAccountSettings request example`, func() {
			fmt.Println("\nGetAccountSettings() result:")
			// begin-getAccountSettings

			getAccountSettingsOptions := iamIdentityService.NewGetAccountSettingsOptions(accountID)

			accountSettingsResponse, response, err := iamIdentityService.GetAccountSettings(getAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettingsResponse, "", "  ")
			fmt.Println(string(b))

			// end-getAccountSettings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsResponse).ToNot(BeNil())

			accountSettingEtag = response.GetHeaders().Get("Etag")
			Expect(accountSettingEtag).ToNot(BeEmpty())
		})
		It(`UpdateAccountSettings request example`, func() {
			fmt.Println("\nUpdateAccountSettings() result:")
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
			fmt.Println(string(b))

			// end-updateAccountSettings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsResponse).ToNot(BeNil())
		})
		It(`CreateReport request example`, func() {
			fmt.Println("\nCreateReport() result:")
			// begin-create_report

			createReportOptions := iamIdentityService.NewCreateReportOptions(accountID)
			createReportOptions.SetType("inactive")
			createReportOptions.SetDuration("120")

			report, response, err := iamIdentityService.CreateReport(createReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(report, "", "  ")
			fmt.Println(string(b))

			// end-create_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(report).ToNot(BeNil())
		})
		It(`GetReport request example`, func() {
			fmt.Println("\nGetReport() result:")
			// begin-get_report

			getReportOptions := iamIdentityService.NewGetReportOptions(accountID, "latest")

			report, response, err := iamIdentityService.GetReport(getReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(report, "", "  ")
			fmt.Println(string(b))

			// end-get_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(report).ToNot(BeNil())
		})
	})
})
