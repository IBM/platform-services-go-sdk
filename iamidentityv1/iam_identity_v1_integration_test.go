// +build integration

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
	"container/list"
	"fmt"
	"net/url"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the iamidentityv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */
const verbose bool = true

var _ = Describe(`IamIdentityV1 Integration Tests`, func() {

	const externalConfigFile = "../iam_identity.env"

	var (
		err                error
		iamIdentityService *iamidentityv1.IamIdentityV1
		serviceURL         string
		config             map[string]string

		apikeyName    string = "Go-SDK-IT-ApiKey"
		serviceIDName string = "Go-SDK-IT-ServiceId"
		accountID     string
		iamID         string
		iamApiKey     string

		//iamIDInvalid     string = "IAM-InvalidId"
		//accountIDInvalid string = "Account-InvalidId"

		apikeyId1   string
		apikeyId2   string
		apikeyEtag1 string

		serviceId1     string
		serviceIdEtag1 string
		//pageToken      string
		newDescription string = "This is an updated description"
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
			config, err = core.GetServiceProperties(iamidentityv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
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
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			iamIdentityServiceOptions := &iamidentityv1.IamIdentityV1Options{}

			iamIdentityService, err = iamidentityv1.NewIamIdentityV1UsingExternalConfig(iamIdentityServiceOptions)

			Expect(err).To(BeNil())
			Expect(iamIdentityService).ToNot(BeNil())
			Expect(iamIdentityService.Service.Options.URL).To(Equal(serviceURL))
		})

	})

	Describe(`CreateApiKey1 - Create an API key1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApiKey(createApiKeyOptions *CreateApiKeyOptions)`, func() {

			createApiKeyOptions := &iamidentityv1.CreateApiKeyOptions{
				Name:        core.StringPtr(apikeyName),
				IamID:       &iamID,
				Description: core.StringPtr("GoSDK test apikey #1"),
				AccountID:   &accountID,
			}

			apiKey, response, err := iamIdentityService.CreateApiKey(createApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())

			apikeyId1 = *apiKey.ID
			Expect(apikeyId1).ToNot(BeNil())

		})
	})

	Describe(`CreateApiKey2 - Create an API key2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApiKey(createApiKeyOptions *CreateApiKeyOptions)`, func() {

			createApiKeyOptions := &iamidentityv1.CreateApiKeyOptions{
				Name:        core.StringPtr(apikeyName),
				IamID:       &iamID,
				Description: core.StringPtr("GoSDK test apikey #2"),
				AccountID:   &accountID,
			}

			apiKey, response, err := iamIdentityService.CreateApiKey(createApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())

			apikeyId2 = *apiKey.ID
			Expect(apikeyId2).ToNot(BeNil())

		})
	})

	Describe(`GetApiKey - Get details of an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApiKey(getApiKeyOptions *GetApiKeyOptions)`, func() {

			getApiKeyOptions := &iamidentityv1.GetApiKeyOptions{
				ID:             core.StringPtr(apikeyId1),
				IncludeHistory: core.BoolPtr(true),
			}

			apiKey, response, err := iamIdentityService.GetApiKey(getApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())

			Expect(*apiKey.ID).To(Equal(apikeyId1))
			Expect(*apiKey.Name).To(Equal(apikeyName))
			Expect(*apiKey.IamID).To(Equal(iamID))
			Expect(*apiKey.AccountID).To(Equal(accountID))
			Expect(*apiKey.CreatedBy).To(Equal(iamID))
			Expect(*apiKey.CreatedAt).ToNot(BeNil())
			Expect(*apiKey.Crn).ToNot(BeNil())

			apikeyEtag1 = *apiKey.EntityTag
		})
	})

	Describe(`GetApiKeysDetails - Get details of an API key by its value`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApiKeysDetails(getApiKeysDetailsOptions *GetApiKeysDetailsOptions)`, func() {

			getApiKeysDetailsOptions := &iamidentityv1.GetApiKeysDetailsOptions{
				IAMApiKey:      core.StringPtr(iamApiKey),
				IncludeHistory: core.BoolPtr(true),
			}

			apiKey, response, err := iamIdentityService.GetApiKeysDetails(getApiKeysDetailsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())

			Expect(*apiKey.AccountID).To(Equal(accountID))
			Expect(*apiKey.IamID).To(Equal(iamID))
		})
	})

	Describe(`ListApiKeys - Get API keys for a given service or user IAM ID and account ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListApiKeys(listApiKeysOptions *ListApiKeysOptions)`, func() {
			//pageToken = " "
			//for ok := true; ok; ok = (pageToken != nil) {
			listApiKeysOptions := &iamidentityv1.ListApiKeysOptions{
				AccountID: &accountID,
				IamID:     &iamID,
				//Pagetoken:		core.StringPtr(pageToken),
				Pagesize:       core.Int64Ptr(int64(1)),
				IncludeHistory: core.BoolPtr(false),
			}

			apiKeyList, response, err := iamIdentityService.ListApiKeys(listApiKeysOptions)

			apikeysListNew := list.New()

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKeyList).ToNot(BeNil())

			for i := 0; i < len(apiKeyList.Apikeys); i++ {
				if Expect(*apiKeyList.Apikeys[i].Name).To(Equal(apikeyName)) {
					apikeysListNew.PushBack(apikeyName)
				}
			}

			Expect(apikeysListNew.Len).To(Equal(2))
			//pageToken = getPageToken(apiKeyList.Next)
			//}
		})
	})

	Describe(`UpdateApiKey - Updates an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateApiKey(updateApiKeyOptions *UpdateApiKeyOptions)`, func() {
			Expect(apikeyId1).To(BeNil())
			Expect(apikeyEtag1).To(BeNil())
			updateApiKeyOptions := &iamidentityv1.UpdateApiKeyOptions{
				ID:          core.StringPtr(apikeyId1),
				IfMatch:     core.StringPtr(apikeyEtag1),
				Description: core.StringPtr(newDescription),
			}

			apiKey, response, err := iamIdentityService.UpdateApiKey(updateApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())

			Expect(*apiKey.ID).To(Equal(apikeyId1))
			Expect(*apiKey.Description).To(Equal(newDescription))
		})
	})

	Describe(`LockApiKey - Lock the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockApiKey(lockApiKeyOptions *LockApiKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeNil())
			lockApiKeyOptions := &iamidentityv1.LockApiKeyOptions{
				ID: core.StringPtr(apikeyId2),
			}

			response, err := iamIdentityService.LockApiKey(lockApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			//apiKey = getApikey(iamIdentityService, apikeyId2)
			//Expect(apiKey).ToNot(BeNil())
			//Expect(*apiKey.Lock).To(Equal(true))

		})
	})

	Describe(`UnlockApiKey - Unlock the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockApiKey(unlockApiKeyOptions *UnlockApiKeyOptions)`, func() {

			unlockApiKeyOptions := &iamidentityv1.UnlockApiKeyOptions{
				ID: core.StringPtr(apikeyId2),
			}

			response, err := iamIdentityService.UnlockApiKey(unlockApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			//apiKey = getApikey(iamIdentityService, apikeyId2)
			//Expect(apiKey).ToNot(BeNil())
			//Expect(*apiKey.Lock).To(Equal(false))
		})
	})

	Describe(`DeleteApiKey1 - Deletes an API key1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteApiKey(deleteApiKeyOptions *DeleteApiKeyOptions)`, func() {

			deleteApiKeyOptions := &iamidentityv1.DeleteApiKeyOptions{
				ID: core.StringPtr(apikeyId1),
			}

			response, err := iamIdentityService.DeleteApiKey(deleteApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//apiKey = getApikey(iamIdentityService, apikeyId1)
			//Expect(apiKey).To(BeNil())

		})
	})

	Describe(`DeleteApiKey2 - Deletes an API key2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteApiKey(deleteApiKeyOptions *DeleteApiKeyOptions)`, func() {

			deleteApiKeyOptions := &iamidentityv1.DeleteApiKeyOptions{
				ID: core.StringPtr(apikeyId2),
			}

			response, err := iamIdentityService.DeleteApiKey(deleteApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//apiKey = getApikey(iamIdentityService, apikeyId2)
			//Expect(apiKey).To(BeNil())

		})
	})

	Describe(`CreateServiceID - Create a service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateServiceID(createServiceIdOptions *CreateServiceIdOptions)`, func() {

			createServiceIdOptions := &iamidentityv1.CreateServiceIdOptions{
				AccountID:   &accountID,
				Name:        &serviceIDName,
				Description: core.StringPtr("GoSDK test serviceId"),
			}

			serviceID, response, err := iamIdentityService.CreateServiceID(createServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(serviceID).ToNot(BeNil())

			serviceId1 = *serviceID.ID
			Expect(serviceId1).ToNot(BeNil())

		})
	})

	Describe(`GetServiceID - Get details of a service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServiceID(getServiceIdOptions *GetServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeNil())
			getServiceIdOptions := &iamidentityv1.GetServiceIdOptions{
				ID:             core.StringPtr(serviceId1),
				IncludeHistory: core.BoolPtr(true),
			}

			serviceID, response, err := iamIdentityService.GetServiceID(getServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())

			Expect(serviceID.Name).To(Equal(serviceIDName))
			Expect(serviceID.Description).To(Equal("GoSDK test serviceId"))
			Expect(serviceID.History).ToNot(BeNil())
			Expect(len(serviceID.History)).ToNot(Equal(0))

			serviceIdEtag1 = *serviceID.EntityTag

		})
	})

	Describe(`ListServiceIds - List service IDs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListServiceIds(listServiceIdsOptions *ListServiceIdsOptions)`, func() {

			listServiceIdsOptions := &iamidentityv1.ListServiceIdsOptions{
				AccountID: core.StringPtr(accountID),
				Name:      core.StringPtr(serviceIDName),
				Pagesize:  core.Int64Ptr(int64(100)),
			}

			serviceIdList, response, err := iamIdentityService.ListServiceIds(listServiceIdsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIdList).ToNot(BeNil())

			Expect(serviceIdList.Serviceids).ToNot(BeNil())
			Expect(len(serviceIdList.Serviceids)).To(Equal(1))
			Expect(serviceIdList.Offset).ToNot(BeNil())
			Expect(serviceIdList.Next).ToNot(BeNil())
			Expect(serviceIdList.Serviceids[0].Name).To(Equal(serviceIDName))
		})
	})

	Describe(`UpdateServiceID - Update service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateServiceID(updateServiceIdOptions *UpdateServiceIdOptions)`, func() {

			updateServiceIdOptions := &iamidentityv1.UpdateServiceIdOptions{
				ID:          core.StringPtr(serviceId1),
				IfMatch:     core.StringPtr(serviceIdEtag1),
				Description: core.StringPtr(newDescription),
			}

			serviceID, response, err := iamIdentityService.UpdateServiceID(updateServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())

			Expect(serviceID.Description).To(Equal(newDescription))

		})
	})

	Describe(`LockServiceID - Lock the service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockServiceID(lockServiceIdOptions *LockServiceIdOptions)`, func() {

			lockServiceIdOptions := &iamidentityv1.LockServiceIdOptions{
				ID: core.StringPtr(serviceId1),
			}

			serviceID, response, err := iamIdentityService.LockServiceID(lockServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())

			Expect(serviceID.Locked).To(Equal(true))

		})
	})

	Describe(`UnlockServiceID - Unlock the service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockServiceID(unlockServiceIdOptions *UnlockServiceIdOptions)`, func() {

			unlockServiceIdOptions := &iamidentityv1.UnlockServiceIdOptions{
				ID: core.StringPtr(serviceId1),
			}

			serviceID, response, err := iamIdentityService.UnlockServiceID(unlockServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())

			Expect(serviceID.Locked).To(Equal(false))

		})
	})

	Describe(`DeleteServiceID - Deletes a service ID and associated API keys`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteServiceID(deleteServiceIdOptions *DeleteServiceIdOptions)`, func() {

			deleteServiceIdOptions := &iamidentityv1.DeleteServiceIdOptions{
				ID: core.StringPtr(serviceId1),
			}

			response, err := iamIdentityService.DeleteServiceID(deleteServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//serviceID = getServiceID(iamIdentityService, serviceId1)
			//Expect(serviceID).To(BeNil())

		})
	})

	Describe(`Teardown - clean up test data`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		tearDown(iamIdentityService, accountID, iamID, serviceIDName)
		log("Finished teardown.")

	})
})

func getApikey(iamIdentityService *iamidentityv1.IamIdentityV1, apikeyId string) {

	getApiKeyOptions := &iamidentityv1.GetApiKeyOptions{
		ID:             core.StringPtr(apikeyId),
		IncludeHistory: core.BoolPtr(true),
	}

	apiKey, response, err := iamIdentityService.GetApiKey(getApiKeyOptions)

	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))
	if Expect(apiKey).ToNot(BeNil()) {
		return
	}
}

func getServiceID(iamIdentityService *iamidentityv1.IamIdentityV1, serviceID string) {

	getServiceIdOptions := &iamidentityv1.GetServiceIdOptions{
		ID:             core.StringPtr(serviceID),
		IncludeHistory: core.BoolPtr(true),
	}

	serviceIDres, response, err := iamIdentityService.GetServiceID(getServiceIdOptions)

	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))

	if Expect(serviceIDres).ToNot(BeNil()) {
		return
	}

}

func getPageToken(s string) {

	Expect(s).ToNot(BeNil())

	u, e := url.Parse(s)
	Expect(e).To(BeNil())
	q := u.Query()
	Expect(q).ToNot(BeNil())
	//return string(q.Get("pagetoken"))
	if Expect(q.Get("pagetoken")).ToNot(BeNil()) {
		return
	}
}

func log(msg string) {
	if verbose {
		fmt.Printf("%s\n", msg)
	}
}

func tearDown(iamIdentityService *iamidentityv1.IamIdentityV1, accountID string, iamID string, serviceIDName string) {

	//Expect(getServiceID(serviceId1)).To(BeNil()

	listApiKeysOptions := &iamidentityv1.ListApiKeysOptions{
		AccountID: core.StringPtr(accountID),
		IamID:     core.StringPtr(iamID),
	}

	apiKeyList, response, err := iamIdentityService.ListApiKeys(listApiKeysOptions)

	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))

	if len(apiKeyList.Apikeys) > 0 {
		for _, element := range apiKeyList.Apikeys {
			//fetchedApikeyID string = apiKeyList.Apikeys[i].ID
			deleteApiKeyOptions := &iamidentityv1.DeleteApiKeyOptions{
				ID: core.StringPtr(*element.ID),
			}
			response, err := iamIdentityService.DeleteApiKey(deleteApiKeyOptions)
			Expect(response).ToNot(BeNil())
			Expect(err).To(BeNil())
		}
	}

	listServiceIdsOptions := &iamidentityv1.ListServiceIdsOptions{
		AccountID: core.StringPtr(accountID),
		Name:      core.StringPtr(serviceIDName),
		Pagesize:  core.Int64Ptr(int64(100)),
	}

	serviceIdList, response, err := iamIdentityService.ListServiceIds(listServiceIdsOptions)

	if len(serviceIdList.Serviceids) > 0 {
		for _, element := range serviceIdList.Serviceids {
			deleteServiceIdOptions := &iamidentityv1.DeleteServiceIdOptions{
				ID: core.StringPtr(*element.ID),
			}
			response, err := iamIdentityService.DeleteServiceID(deleteServiceIdOptions)
			Expect(response).ToNot(BeNil())
			Expect(err).To(BeNil())
		}
	}

}
