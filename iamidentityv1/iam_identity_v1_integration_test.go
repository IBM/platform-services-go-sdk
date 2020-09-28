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
	"encoding/json"
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

var (
	apikeyName    string = "Go-SDK-IT-ApiKey"
	serviceIDName string = "Go-SDK-IT-ServiceId"
	accountID     string
	iamID         string
	iamApiKey     string

	iamIdentityService *iamidentityv1.IamIdentityV1
)

var _ = Describe(`IamIdentityV1 Integration Tests`, func() {

	const externalConfigFile = "../iam_identity.env"

	var (
		err        error
		serviceURL string
		config     map[string]string

		apikeyId1   string
		apikeyId2   string
		apikeyEtag1 string

		serviceId1     string
		serviceIdEtag1 string
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

			fmt.Printf("\nService URL: %s\n", serviceURL)
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
		It("Successfully setup the environment for tests", func() {
			fmt.Println("\nSetup...")
			cleanupResources(iamIdentityService)
			fmt.Println("Finished setup.")
		})
	})

	Describe(`CreateApiKey1 - Create API key #1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApiKey(createApiKeyOptions *CreateApiKeyOptions)`, func() {

			createApiKeyOptions := &iamidentityv1.CreateApiKeyOptions{
				Name:        &apikeyName,
				IamID:       &iamID,
				Description: core.StringPtr("GoSDK test apikey #1"),
			}

			apiKey, response, err := iamIdentityService.CreateApiKey(createApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())
			// fmt.Printf("\nCreateApiKey #1 response:\n%s", toJson(apiKey))

			apikeyId1 = *apiKey.ID
			Expect(apikeyId1).ToNot(BeNil())
		})
	})

	Describe(`CreateApiKey2 - Create API key #2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApiKey(createApiKeyOptions *CreateApiKeyOptions)`, func() {

			createApiKeyOptions := &iamidentityv1.CreateApiKeyOptions{
				Name:        &apikeyName,
				IamID:       &iamID,
				Description: core.StringPtr("GoSDK test apikey #2"),
			}

			apiKey, response, err := iamIdentityService.CreateApiKey(createApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())
			// fmt.Printf("\nCreateApiKey #2 response:\n%s", toJson(apiKey))

			apikeyId2 = *apiKey.ID
			Expect(apikeyId2).ToNot(BeNil())
		})
	})

	Describe(`GetApiKey - Get details of an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApiKey(getApiKeyOptions *GetApiKeyOptions)`, func() {
			Expect(apikeyId1).ToNot(BeNil())

			getApiKeyOptions := &iamidentityv1.GetApiKeyOptions{
				ID:             &apikeyId1,
				IncludeHistory: core.BoolPtr(true),
			}

			apiKey, response, err := iamIdentityService.GetApiKey(getApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			// fmt.Printf("\nGetApiKey response:\n%s", toJson(apiKey))

			Expect(*apiKey.ID).To(Equal(apikeyId1))
			Expect(*apiKey.Name).To(Equal(apikeyName))
			Expect(*apiKey.IamID).To(Equal(iamID))
			Expect(*apiKey.AccountID).To(Equal(accountID))
			Expect(*apiKey.CreatedBy).To(Equal(iamID))
			Expect(*apiKey.CreatedAt).ToNot(BeNil())
			Expect(*apiKey.Locked).To(BeFalse())
			Expect(*apiKey.Crn).ToNot(BeNil())
			Expect(apiKey.History).ToNot(BeEmpty())

			// Grab the Etag value from the response for use in the update operation.
			apikeyEtag1 = response.GetHeaders().Get("Etag")
			Expect(apikeyEtag1).ToNot(BeEmpty())
		})
	})

	Describe(`GetApiKeysDetails - Get details of an API key by its value`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApiKeysDetails(getApiKeysDetailsOptions *GetApiKeysDetailsOptions)`, func() {

			getApiKeysDetailsOptions := &iamidentityv1.GetApiKeysDetailsOptions{
				IAMApiKey:      &iamApiKey,
				IncludeHistory: core.BoolPtr(true),
			}

			apiKey, response, err := iamIdentityService.GetApiKeysDetails(getApiKeysDetailsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			// fmt.Printf("\nGetApiKeyDetails response:\n%s", toJson(apiKey))

			Expect(*apiKey.AccountID).To(Equal(accountID))
			Expect(*apiKey.IamID).To(Equal(iamID))
			Expect(*apiKey.Locked).To(BeFalse())
			Expect(apiKey.History).ToNot(BeEmpty())
		})
	})

	Describe(`ListApiKeys - Get API keys for a given service or user IAM ID and account ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListApiKeys(listApiKeysOptions *ListApiKeysOptions)`, func() {

			apikeys := []iamidentityv1.ApiKey{}

			// var pageToken *string = nil

			// for ok := true; ok; ok = (pageToken != nil) {

			listApiKeysOptions := &iamidentityv1.ListApiKeysOptions{
				AccountID: &accountID,
				IamID:     &iamID,
				//Pagetoken:		core.StringPtr(pageToken),
				Pagesize:       core.Int64Ptr(int64(100)),
				IncludeHistory: core.BoolPtr(false),
			}

			apiKeyList, response, err := iamIdentityService.ListApiKeys(listApiKeysOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKeyList).ToNot(BeNil())
			// fmt.Printf("\nListApiKeys response:\n%s", toJson(apiKeyList))

			// Walk through the returned results and save off the apikeys that we created earlier.
			for _, apikey := range apiKeyList.Apikeys {
				if apikeyName == *apikey.Name {
					apikeys = append(apikeys, apikey)
				}
			}

			// pageToken = getPageToken(apiKeyList.Next)
			// }

			// Make sure we got back two apikeys.
			Expect(len(apikeys)).To(Equal(2))
		})
	})

	Describe(`UpdateApiKey - Updates an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateApiKey(updateApiKeyOptions *UpdateApiKeyOptions)`, func() {
			Expect(apikeyId1).ToNot(BeEmpty())
			Expect(apikeyEtag1).ToNot(BeEmpty())

			updateApiKeyOptions := &iamidentityv1.UpdateApiKeyOptions{
				ID:          &apikeyId1,
				IfMatch:     &apikeyEtag1,
				Description: &newDescription,
			}

			apiKey, response, err := iamIdentityService.UpdateApiKey(updateApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			// fmt.Printf("\nUpdateApiKey response:\n%s", toJson(apiKey))

			Expect(*apiKey.ID).To(Equal(apikeyId1))
			Expect(*apiKey.Description).To(Equal(newDescription))
		})
	})

	Describe(`LockApiKey - Lock the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockApiKey(lockApiKeyOptions *LockApiKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			lockApiKeyOptions := &iamidentityv1.LockApiKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.LockApiKey(lockApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			apiKey := getApikey(iamIdentityService, apikeyId2)
			Expect(apiKey).ToNot(BeNil())
			Expect(*apiKey.Locked).To(BeTrue())
		})
	})

	Describe(`UnlockApiKey - Unlock the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockApiKey(unlockApiKeyOptions *UnlockApiKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			unlockApiKeyOptions := &iamidentityv1.UnlockApiKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.UnlockApiKey(unlockApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			apiKey := getApikey(iamIdentityService, apikeyId2)
			Expect(apiKey).ToNot(BeNil())
			Expect(*apiKey.Locked).To(BeFalse())
		})
	})

	Describe(`DeleteApiKey1 - Deletes an API key1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteApiKey(deleteApiKeyOptions *DeleteApiKeyOptions)`, func() {
			Expect(apikeyId1).ToNot(BeEmpty())

			deleteApiKeyOptions := &iamidentityv1.DeleteApiKeyOptions{
				ID: &apikeyId1,
			}

			response, err := iamIdentityService.DeleteApiKey(deleteApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getApikey(iamIdentityService, apikeyId1)
			Expect(apiKey).To(BeNil())
		})
	})

	Describe(`DeleteApiKey2 - Deletes an API key2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteApiKey(deleteApiKeyOptions *DeleteApiKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			deleteApiKeyOptions := &iamidentityv1.DeleteApiKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.DeleteApiKey(deleteApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getApikey(iamIdentityService, apikeyId2)
			Expect(apiKey).To(BeNil())
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
			// fmt.Printf("\nCreateServiceID response:\n%s", toJson(serviceID))

			serviceId1 = *serviceID.ID
			Expect(serviceId1).ToNot(BeNil())
		})
	})

	Describe(`GetServiceID - Get details of a service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServiceID(getServiceIdOptions *GetServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())
			getServiceIdOptions := &iamidentityv1.GetServiceIdOptions{
				ID:             &serviceId1,
				IncludeHistory: core.BoolPtr(true),
			}

			serviceID, response, err := iamIdentityService.GetServiceID(getServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			// fmt.Printf("\nGetServiceID response:\n%s", toJson(serviceID))

			Expect(*serviceID.Name).To(Equal(serviceIDName))
			Expect(*serviceID.Description).To(Equal("GoSDK test serviceId"))
			Expect(serviceID.History).ToNot(BeEmpty())

			// Grab the Etag value from the response for use in the update operation.
			serviceIdEtag1 = response.GetHeaders().Get("Etag")
			Expect(serviceIdEtag1).ToNot(BeEmpty())
		})
	})

	Describe(`ListServiceIds - List service IDs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListServiceIds(listServiceIdsOptions *ListServiceIdsOptions)`, func() {

			listServiceIdsOptions := &iamidentityv1.ListServiceIdsOptions{
				AccountID: &accountID,
				Name:      &serviceIDName,
				Pagesize:  core.Int64Ptr(int64(100)),
			}

			serviceIdList, response, err := iamIdentityService.ListServiceIds(listServiceIdsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIdList).ToNot(BeNil())
			// fmt.Printf("\nListServiceIds response:\n%s", toJson(serviceIdList))

			Expect(len(serviceIdList.Serviceids)).To(Equal(1))
			Expect(serviceIdList.Offset).ToNot(BeNil())
			Expect(serviceIdList.Next).To(BeNil())
			Expect(*serviceIdList.Serviceids[0].Name).To(Equal(serviceIDName))
		})
	})

	Describe(`UpdateServiceID - Update service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateServiceID(updateServiceIdOptions *UpdateServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())
			Expect(serviceIdEtag1).ToNot(BeEmpty())

			updateServiceIdOptions := &iamidentityv1.UpdateServiceIdOptions{
				ID:          &serviceId1,
				IfMatch:     &serviceIdEtag1,
				Description: &newDescription,
			}

			serviceID, response, err := iamIdentityService.UpdateServiceID(updateServiceIdOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			// fmt.Printf("\nUpdateServiceID response:\n%s", toJson(serviceID))

			Expect(*serviceID.Description).To(Equal(newDescription))
		})
	})

	Describe(`LockServiceID - Lock the service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockServiceID(lockServiceIdOptions *LockServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())

			lockServiceIdOptions := &iamidentityv1.LockServiceIdOptions{
				ID: &serviceId1,
			}

			serviceID, response, err := iamIdentityService.LockServiceID(lockServiceIdOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			// fmt.Printf("\nLockServiceID response:\n%s", toJson(serviceID))

			Expect(*serviceID.Locked).To(BeTrue())
		})
	})

	Describe(`UnlockServiceID - Unlock the service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockServiceID(unlockServiceIdOptions *UnlockServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())

			unlockServiceIdOptions := &iamidentityv1.UnlockServiceIdOptions{
				ID: &serviceId1,
			}

			serviceID, response, err := iamIdentityService.UnlockServiceID(unlockServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			// fmt.Printf("\nUnlockServiceID response:\n%s", toJson(serviceID))

			Expect(*serviceID.Locked).To(BeFalse())
		})
	})

	Describe(`DeleteServiceID - Deletes a service ID and associated API keys`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteServiceID(deleteServiceIdOptions *DeleteServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())

			deleteServiceIdOptions := &iamidentityv1.DeleteServiceIdOptions{
				ID: &serviceId1,
			}

			response, err := iamIdentityService.DeleteServiceID(deleteServiceIdOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			serviceID := getServiceID(iamIdentityService, serviceId1)
			Expect(serviceID).To(BeNil())
		})
	})
})

var _ = AfterSuite(func() {
	fmt.Println("\nBeginning teardown.")
	cleanupResources(iamIdentityService)
	fmt.Println("Finished teardown.")
})

func getApikey(service *iamidentityv1.IamIdentityV1, apikeyId string) *iamidentityv1.ApiKey {
	getApiKeyOptions := &iamidentityv1.GetApiKeyOptions{
		ID: &apikeyId,
	}
	apiKey, _, _ := service.GetApiKey(getApiKeyOptions)
	return apiKey
}

func getServiceID(iamIdentityService *iamidentityv1.IamIdentityV1, serviceID string) *iamidentityv1.ServiceID {
	getServiceIdOptions := &iamidentityv1.GetServiceIdOptions{
		ID: &serviceID,
	}
	result, _, _ := iamIdentityService.GetServiceID(getServiceIdOptions)
	return result
}

func getPageTokenFromURL(s string) string {
	if s == "" {
		return ""
	}

	u, err := url.Parse(s)
	if err != nil {
		return ""
	}

	if u.RawQuery == "" {
		return ""
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return ""
	}

	return q.Get("pagetoken")
}

func toJson(obj interface{}) string {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func cleanupResources(service *iamidentityv1.IamIdentityV1) {
	if service == nil {
		panic("'service' cannot be nil!")
	}

	listApiKeysOptions := &iamidentityv1.ListApiKeysOptions{
		AccountID: &accountID,
		IamID:     &iamID,
		Pagesize:  core.Int64Ptr(int64(100)),
	}

	apiKeyList, response, err := service.ListApiKeys(listApiKeysOptions)
	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))

	numApiKeys := len(apiKeyList.Apikeys)
	fmt.Printf(">>> Cleanup found %d apikeys.\n", numApiKeys)

	if numApiKeys > 0 {
		for _, element := range apiKeyList.Apikeys {
			if *element.Name == apikeyName {
				fmt.Printf(">>> Deleting apikey: %s\n", *element.ID)
				deleteApiKeyOptions := &iamidentityv1.DeleteApiKeyOptions{
					ID: element.ID,
				}
				response, err := service.DeleteApiKey(deleteApiKeyOptions)
				Expect(response).ToNot(BeNil())
				Expect(err).To(BeNil())
			}
		}
	}

	listServiceIdsOptions := &iamidentityv1.ListServiceIdsOptions{
		AccountID: &accountID,
		Name:      &serviceIDName,
		Pagesize:  core.Int64Ptr(int64(100)),
	}

	serviceIdList, response, err := service.ListServiceIds(listServiceIdsOptions)

	numServiceIds := len(serviceIdList.Serviceids)
	fmt.Printf(">>> Cleanup found %d serviceIDs.\n", numServiceIds)

	if numServiceIds > 0 {
		for _, element := range serviceIdList.Serviceids {
			fmt.Printf(">>> Deleting serviceId: %s\n", *element.ID)
			deleteServiceIdOptions := &iamidentityv1.DeleteServiceIdOptions{
				ID: element.ID,
			}
			response, err := service.DeleteServiceID(deleteServiceIdOptions)
			Expect(response).ToNot(BeNil())
			Expect(err).To(BeNil())
		}
	}
}
