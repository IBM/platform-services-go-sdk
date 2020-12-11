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
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/platform-services-go-sdk/common"
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
	apikeyName    string = "Go-SDK-IT-APIKey"
	serviceIDName string = "Go-SDK-IT-ServiceId"
	accountID     string
	iamID         string
	iamAPIKey     string

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

			iamAPIKey = config["APIKEY"]
			Expect(iamAPIKey).ToNot(BeEmpty())

			fmt.Fprintf(GinkgoWriter, "Service URL: %s\n", serviceURL)
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

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags)))
			iamIdentityService.EnableRetries(4, 30*time.Second)
		})
		It("Successfully setup the environment for tests", func() {
			fmt.Fprintln(GinkgoWriter, "Setup...")
			cleanupResources(iamIdentityService)
			fmt.Fprintln(GinkgoWriter, "Finished setup.")
		})
	})

	Describe(`CreateAPIKey1 - Create API key #1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAPIKey(createAPIKeyOptions *CreateAPIKeyOptions)`, func() {

			createAPIKeyOptions := &iamidentityv1.CreateAPIKeyOptions{
				Name:        &apikeyName,
				IamID:       &iamID,
				Description: core.StringPtr("GoSDK test apikey #1"),
			}

			apiKey, response, err := iamIdentityService.CreateAPIKey(createAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateAPIKey #1 response:\n%s\n", common.ToJSON(apiKey))

			apikeyId1 = *apiKey.ID
			Expect(apikeyId1).ToNot(BeNil())
		})
	})

	Describe(`CreateAPIKey2 - Create API key #2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAPIKey(createAPIKeyOptions *CreateAPIKeyOptions)`, func() {

			createAPIKeyOptions := &iamidentityv1.CreateAPIKeyOptions{
				Name:        &apikeyName,
				IamID:       &iamID,
				Description: core.StringPtr("GoSDK test apikey #2"),
			}

			apiKey, response, err := iamIdentityService.CreateAPIKey(createAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateAPIKey #2 response:\n%s\n", common.ToJSON(apiKey))

			apikeyId2 = *apiKey.ID
			Expect(apikeyId2).ToNot(BeNil())
		})
	})

	Describe(`GetAPIKey - Get details of an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAPIKey(getAPIKeyOptions *GetAPIKeyOptions)`, func() {
			Expect(apikeyId1).ToNot(BeNil())

			getAPIKeyOptions := &iamidentityv1.GetAPIKeyOptions{
				ID:             &apikeyId1,
				IncludeHistory: core.BoolPtr(true),
			}

			apiKey, response, err := iamIdentityService.GetAPIKey(getAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetAPIKey response:\n%s\n", common.ToJSON(apiKey))

			Expect(*apiKey.ID).To(Equal(apikeyId1))
			Expect(*apiKey.Name).To(Equal(apikeyName))
			Expect(*apiKey.IamID).To(Equal(iamID))
			Expect(*apiKey.AccountID).To(Equal(accountID))
			Expect(*apiKey.CreatedBy).To(Equal(iamID))
			Expect(*apiKey.CreatedAt).ToNot(BeNil())
			Expect(*apiKey.Locked).To(BeFalse())
			Expect(*apiKey.CRN).ToNot(BeNil())
			Expect(apiKey.History).ToNot(BeEmpty())

			// Grab the Etag value from the response for use in the update operation.
			apikeyEtag1 = response.GetHeaders().Get("Etag")
			Expect(apikeyEtag1).ToNot(BeEmpty())
		})
	})

	Describe(`GetAPIKeysDetails - Get details of an API key by its value`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAPIKeysDetails(getAPIKeysDetailsOptions *GetAPIKeysDetailsOptions)`, func() {

			getAPIKeysDetailsOptions := &iamidentityv1.GetAPIKeysDetailsOptions{
				IamAPIKey:      &iamAPIKey,
				IncludeHistory: core.BoolPtr(true),
			}

			apiKey, response, err := iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetAPIKeyDetails response:\n%s\n", common.ToJSON(apiKey))

			Expect(*apiKey.AccountID).To(Equal(accountID))
			Expect(*apiKey.IamID).To(Equal(iamID))
			Expect(*apiKey.Locked).To(BeFalse())
			Expect(apiKey.History).ToNot(BeEmpty())
		})
	})

	Describe(`ListAPIKeys - Get API keys for a given service or user IAM ID and account ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAPIKeys(listAPIKeysOptions *ListAPIKeysOptions)`, func() {

			apikeys := []iamidentityv1.APIKey{}

			// var pageToken *string = nil
			var pageTokenPresent bool = true
			var pageToken *string = nil

			// for ok := true; ok; ok = (pageToken != nil) {
			for pageTokenPresent {
				listAPIKeysOptions := &iamidentityv1.ListAPIKeysOptions{
					AccountID: &accountID,
					IamID:     &iamID,
					Pagetoken: pageToken,
					Pagesize:  core.Int64Ptr(int64(1)),
				}

				apiKeyList, response, err := iamIdentityService.ListAPIKeys(listAPIKeysOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(apiKeyList).ToNot(BeNil())
				fmt.Fprintf(GinkgoWriter, "ListAPIKeys response:\n%s\n", common.ToJSON(apiKeyList))

				// Walk through the returned results and save off the apikeys that we created earlier.
				for _, apikey := range apiKeyList.Apikeys {
					if apikeyName == *apikey.Name {
						apikeys = append(apikeys, apikey)
					}
				}

				pageToken = getPageTokenFromURL(apiKeyList.Next)
				pageTokenPresent = (pageToken != nil)
			}

			// Make sure we got back two apikeys.
			Expect(len(apikeys)).To(Equal(2))
		})
	})

	Describe(`UpdateAPIKey - Updates an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAPIKey(updateAPIKeyOptions *UpdateAPIKeyOptions)`, func() {
			Expect(apikeyId1).ToNot(BeEmpty())
			Expect(apikeyEtag1).ToNot(BeEmpty())

			updateAPIKeyOptions := &iamidentityv1.UpdateAPIKeyOptions{
				ID:          &apikeyId1,
				IfMatch:     &apikeyEtag1,
				Description: &newDescription,
			}

			apiKey, response, err := iamIdentityService.UpdateAPIKey(updateAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "UpdateAPIKey response:\n%s\n", common.ToJSON(apiKey))

			Expect(*apiKey.ID).To(Equal(apikeyId1))
			Expect(*apiKey.Description).To(Equal(newDescription))
		})
	})

	Describe(`LockAPIKey - Lock the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockAPIKey(lockAPIKeyOptions *LockAPIKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			lockAPIKeyOptions := &iamidentityv1.LockAPIKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.LockAPIKey(lockAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getAPIkey(iamIdentityService, apikeyId2)
			Expect(apiKey).ToNot(BeNil())
			Expect(*apiKey.Locked).To(BeTrue())
		})
	})

	Describe(`UnlockAPIKey - Unlock the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockAPIKey(unlockAPIKeyOptions *UnlockAPIKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			unlockAPIKeyOptions := &iamidentityv1.UnlockAPIKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.UnlockAPIKey(unlockAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getAPIkey(iamIdentityService, apikeyId2)
			Expect(apiKey).ToNot(BeNil())
			Expect(*apiKey.Locked).To(BeFalse())
		})
	})

	Describe(`DeleteAPIKey1 - Deletes an API key1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAPIKey(deleteAPIKeyOptions *DeleteAPIKeyOptions)`, func() {
			Expect(apikeyId1).ToNot(BeEmpty())

			deleteAPIKeyOptions := &iamidentityv1.DeleteAPIKeyOptions{
				ID: &apikeyId1,
			}

			response, err := iamIdentityService.DeleteAPIKey(deleteAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getAPIkey(iamIdentityService, apikeyId1)
			Expect(apiKey).To(BeNil())
		})
	})

	Describe(`DeleteAPIKey2 - Deletes an API key2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAPIKey(deleteAPIKeyOptions *DeleteAPIKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			deleteAPIKeyOptions := &iamidentityv1.DeleteAPIKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.DeleteAPIKey(deleteAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getAPIkey(iamIdentityService, apikeyId2)
			Expect(apiKey).To(BeNil())
		})
	})

	Describe(`CreateServiceID - Create a service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateServiceID(createServiceIdOptions *CreateServiceIdOptions)`, func() {

			createServiceIDOptions := &iamidentityv1.CreateServiceIDOptions{
				AccountID:   &accountID,
				Name:        &serviceIDName,
				Description: core.StringPtr("GoSDK test serviceId"),
			}

			serviceID, response, err := iamIdentityService.CreateServiceID(createServiceIDOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(serviceID).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateServiceID response:\n%s\n", common.ToJSON(serviceID))

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
			getServiceIDOptions := &iamidentityv1.GetServiceIDOptions{
				ID:             &serviceId1,
				IncludeHistory: core.BoolPtr(true),
			}

			serviceID, response, err := iamIdentityService.GetServiceID(getServiceIDOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetServiceID response:\n%s\n", common.ToJSON(serviceID))

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
			fmt.Fprintf(GinkgoWriter, "ListServiceIds response:\n%s\n", common.ToJSON(serviceIdList))

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

			updateServiceIDOptions := &iamidentityv1.UpdateServiceIDOptions{
				ID:          &serviceId1,
				IfMatch:     &serviceIdEtag1,
				Description: &newDescription,
			}

			serviceID, response, err := iamIdentityService.UpdateServiceID(updateServiceIDOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "UpdateServiceID response:\n%s\n", common.ToJSON(serviceID))

			Expect(*serviceID.Description).To(Equal(newDescription))
		})
	})

	Describe(`LockServiceID - Lock the service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockServiceID(lockServiceIdOptions *LockServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())

			lockServiceIDOptions := &iamidentityv1.LockServiceIDOptions{
				ID: &serviceId1,
			}

			response, err := iamIdentityService.LockServiceID(lockServiceIDOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
			fmt.Fprintf(GinkgoWriter, "LockServiceID response:\n%v\n", response)

			serviceID := getServiceID(iamIdentityService, serviceId1)
			Expect(serviceID).ToNot(BeNil())
			Expect(*serviceID.Locked).To(BeTrue())
		})
	})

	Describe(`UnlockServiceID - Unlock the service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockServiceID(unlockServiceIdOptions *UnlockServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())

			unlockServiceIDOptions := &iamidentityv1.UnlockServiceIDOptions{
				ID: &serviceId1,
			}

			response, err := iamIdentityService.UnlockServiceID(unlockServiceIDOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
			fmt.Fprintf(GinkgoWriter, "UnlockServiceID response:\n%v\n", response)

			serviceID := getServiceID(iamIdentityService, serviceId1)
			Expect(serviceID).ToNot(BeNil())
			Expect(*serviceID.Locked).To(BeFalse())
		})
	})

	Describe(`DeleteServiceID - Deletes a service ID and associated API keys`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteServiceID(deleteServiceIdOptions *DeleteServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())

			deleteServiceIDOptions := &iamidentityv1.DeleteServiceIDOptions{
				ID: &serviceId1,
			}

			response, err := iamIdentityService.DeleteServiceID(deleteServiceIDOptions)
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

func getAPIkey(service *iamidentityv1.IamIdentityV1, apikeyID string) *iamidentityv1.APIKey {
	getAPIKeyOptions := &iamidentityv1.GetAPIKeyOptions{
		ID: &apikeyID,
	}
	apiKey, _, _ := service.GetAPIKey(getAPIKeyOptions)
	return apiKey
}

func getServiceID(iamIdentityService *iamidentityv1.IamIdentityV1, serviceID string) *iamidentityv1.ServiceID {
	getServiceIDOptions := &iamidentityv1.GetServiceIDOptions{
		ID: &serviceID,
	}
	result, _, _ := iamIdentityService.GetServiceID(getServiceIDOptions)
	return result
}

func getPageTokenFromURL(sptr *string) *string {
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

	token := q.Get("pagetoken")
	if token == "" {
		return nil
	}
	return &token
}

func cleanupResources(service *iamidentityv1.IamIdentityV1) {
	if service == nil {
		panic("'service' cannot be nil!")
	}

	listAPIKeysOptions := &iamidentityv1.ListAPIKeysOptions{
		AccountID: &accountID,
		IamID:     &iamID,
		Pagesize:  core.Int64Ptr(int64(100)),
	}

	apiKeyList, response, err := service.ListAPIKeys(listAPIKeysOptions)
	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))

	numAPIKeys := len(apiKeyList.Apikeys)
	fmt.Fprintf(GinkgoWriter, ">>> Cleanup found %d apikeys.\n", numAPIKeys)

	if numAPIKeys > 0 {
		for _, element := range apiKeyList.Apikeys {
			if *element.Name == apikeyName {
				fmt.Fprintf(GinkgoWriter, ">>> Deleting apikey: %s\n", *element.ID)
				deleteAPIKeyOptions := &iamidentityv1.DeleteAPIKeyOptions{
					ID: element.ID,
				}
				response, err := service.DeleteAPIKey(deleteAPIKeyOptions)
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

	serviceIDList, response, err := service.ListServiceIds(listServiceIdsOptions)

	numServiceIds := len(serviceIDList.Serviceids)
	fmt.Fprintf(GinkgoWriter, ">>> Cleanup found %d serviceIDs.\n", numServiceIds)

	if numServiceIds > 0 {
		for _, element := range serviceIDList.Serviceids {
			fmt.Fprintf(GinkgoWriter, ">>> Deleting serviceId: %s\n", *element.ID)
			deleteServiceIDOptions := &iamidentityv1.DeleteServiceIDOptions{
				ID: element.ID,
			}
			response, err := service.DeleteServiceID(deleteServiceIDOptions)
			Expect(response).ToNot(BeNil())
			Expect(err).To(BeNil())
		}
	}
}
