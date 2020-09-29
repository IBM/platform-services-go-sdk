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

package iamidentityv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IAM-IDENTITY service.
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
// in a "credentials" file and then:
// export IBM_CREDENTIALS_FILE=<name of credentials file>
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
	iamApiKey     string

	apikeyID   string
	apikeyEtag string

	svcID     string
	svcIDEtag string
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

			iamApiKey = config["APIKEY"]
			Expect(iamApiKey).ToNot(BeEmpty())

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
		It(`CreateApiKey request example`, func() {
			// begin-create_api_key

			createApiKeyOptions := iamIdentityService.NewCreateApiKeyOptions(apikeyName, iamID)
			createApiKeyOptions.SetDescription("Example ApiKey")

			apiKey, response, err := iamIdentityService.CreateApiKey(createApiKeyOptions)
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
		It(`ListApiKeys request example`, func() {
			// begin-list_api_keys

			listApiKeysOptions := iamIdentityService.NewListApiKeysOptions()
			listApiKeysOptions.SetAccountID(accountID)
			listApiKeysOptions.SetIamID(iamID)
			listApiKeysOptions.SetIncludeHistory(true)

			apiKeyList, response, err := iamIdentityService.ListApiKeys(listApiKeysOptions)
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
		It(`GetApiKeysDetails request example`, func() {
			// begin-get_api_keys_details

			getApiKeysDetailsOptions := iamIdentityService.NewGetApiKeysDetailsOptions()
			getApiKeysDetailsOptions.SetIAMApiKey(iamApiKey)
			getApiKeysDetailsOptions.SetIncludeHistory(false)

			apiKey, response, err := iamIdentityService.GetApiKeysDetails(getApiKeysDetailsOptions)
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
		It(`GetApiKey request example`, func() {
			// begin-get_api_key

			getApiKeyOptions := iamIdentityService.NewGetApiKeyOptions(apikeyID)

			apiKey, response, err := iamIdentityService.GetApiKey(getApiKeyOptions)
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
		It(`UpdateApiKey request example`, func() {
			// begin-update_api_key

			updateApiKeyOptions := iamIdentityService.NewUpdateApiKeyOptions(apikeyID, apikeyEtag)
			updateApiKeyOptions.SetDescription("This is an updated description")

			apiKey, response, err := iamIdentityService.UpdateApiKey(updateApiKeyOptions)
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
		It(`LockApiKey request example`, func() {
			// begin-lock_api_key

			lockApiKeyOptions := iamIdentityService.NewLockApiKeyOptions(apikeyID)

			response, err := iamIdentityService.LockApiKey(lockApiKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-lock_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
		It(`UnlockApiKey request example`, func() {
			// begin-unlock_api_key

			unlockApiKeyOptions := iamIdentityService.NewUnlockApiKeyOptions(apikeyID)

			response, err := iamIdentityService.UnlockApiKey(unlockApiKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-unlock_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
		It(`DeleteApiKey request example`, func() {
			// begin-delete_api_key

			deleteApiKeyOptions := iamIdentityService.NewDeleteApiKeyOptions(apikeyID)

			response, err := iamIdentityService.DeleteApiKey(deleteApiKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateServiceID request example`, func() {
			// begin-create_service_id

			createServiceIdOptions := iamIdentityService.NewCreateServiceIdOptions(accountID, serviceIDName)
			createServiceIdOptions.SetDescription("Example ServiceId")

			serviceID, response, err := iamIdentityService.CreateServiceID(createServiceIdOptions)
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
			// begin-get_service_id

			getServiceIdOptions := iamIdentityService.NewGetServiceIdOptions(svcID)

			serviceID, response, err := iamIdentityService.GetServiceID(getServiceIdOptions)
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
			// begin-list_service_ids

			listServiceIdsOptions := iamIdentityService.NewListServiceIdsOptions()
			listServiceIdsOptions.SetAccountID(accountID)
			listServiceIdsOptions.SetName(serviceIDName)

			serviceIdList, response, err := iamIdentityService.ListServiceIds(listServiceIdsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceIdList, "", "  ")
			fmt.Println(string(b))

			// end-list_service_ids

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIdList).ToNot(BeNil())
		})
		It(`UpdateServiceID request example`, func() {
			// begin-update_service_id

			updateServiceIdOptions := iamIdentityService.NewUpdateServiceIdOptions(svcID, svcIDEtag)
			updateServiceIdOptions.SetDescription("This is an updated description")

			serviceID, response, err := iamIdentityService.UpdateServiceID(updateServiceIdOptions)
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

			lockServiceIdOptions := iamIdentityService.NewLockServiceIdOptions(svcID)

			serviceID, response, err := iamIdentityService.LockServiceID(lockServiceIdOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Println(string(b))

			// end-lock_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
		})
		It(`UnlockServiceID request example`, func() {
			// begin-unlock_service_id

			unlockServiceIdOptions := iamIdentityService.NewUnlockServiceIdOptions(svcID)

			serviceID, response, err := iamIdentityService.UnlockServiceID(unlockServiceIdOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Println(string(b))

			// end-unlock_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
		})
		It(`DeleteServiceID request example`, func() {
			// begin-delete_service_id

			deleteServiceIdOptions := iamIdentityService.NewDeleteServiceIdOptions(svcID)

			response, err := iamIdentityService.DeleteServiceID(deleteServiceIdOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
