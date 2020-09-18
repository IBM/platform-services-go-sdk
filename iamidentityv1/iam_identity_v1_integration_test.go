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

var _ = Describe(`IamIdentityV1 Integration Tests`, func() {

	const externalConfigFile = "../iam_identity.env"

	var (
		err                error
		iamIdentityService *iamidentityv1.IamIdentityV1
		serviceURL         string
		config             map[string]string
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

	Describe(`ListApiKeys - Get API keys for a given service or user IAM ID and account ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListApiKeys(listApiKeysOptions *ListApiKeysOptions)`, func() {

			listApiKeysOptions := &iamidentityv1.ListApiKeysOptions{
				AccountID:      core.StringPtr("testString"),
				IamID:          core.StringPtr("testString"),
				Pagesize:       core.Int64Ptr(int64(38)),
				Pagetoken:      core.StringPtr("testString"),
				Scope:          core.StringPtr("entity"),
				Type:           core.StringPtr("user"),
				Sort:           core.StringPtr("testString"),
				Order:          core.StringPtr("asc"),
				IncludeHistory: core.BoolPtr(true),
			}

			apiKeyList, response, err := iamIdentityService.ListApiKeys(listApiKeysOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKeyList).ToNot(BeNil())

		})
	})

	Describe(`CreateApiKey - Create an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApiKey(createApiKeyOptions *CreateApiKeyOptions)`, func() {

			createApiKeyOptions := &iamidentityv1.CreateApiKeyOptions{
				Name:        core.StringPtr("testString"),
				IamID:       core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				AccountID:   core.StringPtr("testString"),
				Apikey:      core.StringPtr("testString"),
				StoreValue:  core.BoolPtr(true),
				EntityLock:  core.StringPtr("testString"),
			}

			apiKey, response, err := iamIdentityService.CreateApiKey(createApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())

		})
	})

	Describe(`GetApiKeysDetails - Get details of an API key by its value`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApiKeysDetails(getApiKeysDetailsOptions *GetApiKeysDetailsOptions)`, func() {

			getApiKeysDetailsOptions := &iamidentityv1.GetApiKeysDetailsOptions{
				IAMApiKey:      core.StringPtr("testString"),
				IncludeHistory: core.BoolPtr(true),
			}

			apiKey, response, err := iamIdentityService.GetApiKeysDetails(getApiKeysDetailsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())

		})
	})

	Describe(`GetApiKey - Get details of an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApiKey(getApiKeyOptions *GetApiKeyOptions)`, func() {

			getApiKeyOptions := &iamidentityv1.GetApiKeyOptions{
				ID:             core.StringPtr("testString"),
				IncludeHistory: core.BoolPtr(true),
			}

			apiKey, response, err := iamIdentityService.GetApiKey(getApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())

		})
	})

	Describe(`UpdateApiKey - Updates an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateApiKey(updateApiKeyOptions *UpdateApiKeyOptions)`, func() {

			updateApiKeyOptions := &iamidentityv1.UpdateApiKeyOptions{
				ID:          core.StringPtr("testString"),
				IfMatch:     core.StringPtr("testString"),
				Name:        core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			apiKey, response, err := iamIdentityService.UpdateApiKey(updateApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())

		})
	})

	Describe(`LockApiKey - Lock the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockApiKey(lockApiKeyOptions *LockApiKeyOptions)`, func() {

			lockApiKeyOptions := &iamidentityv1.LockApiKeyOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := iamIdentityService.LockApiKey(lockApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`ListServiceIds - List service IDs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListServiceIds(listServiceIdsOptions *ListServiceIdsOptions)`, func() {

			listServiceIdsOptions := &iamidentityv1.ListServiceIdsOptions{
				AccountID:      core.StringPtr("testString"),
				Name:           core.StringPtr("testString"),
				Pagesize:       core.Int64Ptr(int64(38)),
				Pagetoken:      core.StringPtr("testString"),
				Sort:           core.StringPtr("testString"),
				Order:          core.StringPtr("asc"),
				IncludeHistory: core.BoolPtr(true),
			}

			serviceIdList, response, err := iamIdentityService.ListServiceIds(listServiceIdsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIdList).ToNot(BeNil())

		})
	})

	Describe(`CreateServiceID - Create a service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateServiceID(createServiceIdOptions *CreateServiceIdOptions)`, func() {

			createApiKeyRequestModel := &iamidentityv1.CreateApiKeyRequest{
				Name:        core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				IamID:       core.StringPtr("testString"),
				AccountID:   core.StringPtr("testString"),
				Apikey:      core.StringPtr("testString"),
				StoreValue:  core.BoolPtr(true),
			}

			createServiceIdOptions := &iamidentityv1.CreateServiceIdOptions{
				AccountID:          core.StringPtr("testString"),
				Name:               core.StringPtr("testString"),
				Description:        core.StringPtr("testString"),
				UniqueInstanceCrns: []string{"testString"},
				Apikey:             createApiKeyRequestModel,
				EntityLock:         core.StringPtr("testString"),
			}

			serviceID, response, err := iamIdentityService.CreateServiceID(createServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(serviceID).ToNot(BeNil())

		})
	})

	Describe(`GetServiceID - Get details of a service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServiceID(getServiceIdOptions *GetServiceIdOptions)`, func() {

			getServiceIdOptions := &iamidentityv1.GetServiceIdOptions{
				ID:             core.StringPtr("testString"),
				IncludeHistory: core.BoolPtr(true),
			}

			serviceID, response, err := iamIdentityService.GetServiceID(getServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())

		})
	})

	Describe(`UpdateServiceID - Update service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateServiceID(updateServiceIdOptions *UpdateServiceIdOptions)`, func() {

			updateServiceIdOptions := &iamidentityv1.UpdateServiceIdOptions{
				ID:                 core.StringPtr("testString"),
				IfMatch:            core.StringPtr("testString"),
				Name:               core.StringPtr("testString"),
				Description:        core.StringPtr("testString"),
				UniqueInstanceCrns: []string{"testString"},
			}

			serviceID, response, err := iamIdentityService.UpdateServiceID(updateServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())

		})
	})

	Describe(`LockServiceID - Lock the service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockServiceID(lockServiceIdOptions *LockServiceIdOptions)`, func() {

			lockServiceIdOptions := &iamidentityv1.LockServiceIdOptions{
				ID: core.StringPtr("testString"),
			}

			serviceID, response, err := iamIdentityService.LockServiceID(lockServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())

		})
	})

	Describe(`UnlockServiceID - Unlock the service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockServiceID(unlockServiceIdOptions *UnlockServiceIdOptions)`, func() {

			unlockServiceIdOptions := &iamidentityv1.UnlockServiceIdOptions{
				ID: core.StringPtr("testString"),
			}

			serviceID, response, err := iamIdentityService.UnlockServiceID(unlockServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())

		})
	})

	Describe(`UnlockApiKey - Unlock the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockApiKey(unlockApiKeyOptions *UnlockApiKeyOptions)`, func() {

			unlockApiKeyOptions := &iamidentityv1.UnlockApiKeyOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := iamIdentityService.UnlockApiKey(unlockApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`DeleteServiceID - Deletes a service ID and associated API keys`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteServiceID(deleteServiceIdOptions *DeleteServiceIdOptions)`, func() {

			deleteServiceIdOptions := &iamidentityv1.DeleteServiceIdOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := iamIdentityService.DeleteServiceID(deleteServiceIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteApiKey - Deletes an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteApiKey(deleteApiKeyOptions *DeleteApiKeyOptions)`, func() {

			deleteApiKeyOptions := &iamidentityv1.DeleteApiKeyOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := iamIdentityService.DeleteApiKey(deleteApiKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
