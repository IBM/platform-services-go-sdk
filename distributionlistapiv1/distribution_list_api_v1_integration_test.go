//go:build integration

/**
 * (C) Copyright IBM Corp. 2026.
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

package distributionlistapiv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/distributionlistapiv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the distributionlistapiv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`DistributionListApiV1 Integration Tests`, func() {
	const externalConfigFile = "../distribution_list_api_v1.env"

	var (
		err                        error
		distributionListApiService *distributionlistapiv1.DistributionListApiV1
		serviceURL                 string
		config                     map[string]string
		accountID                  string
		instanceID                 string
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
			config, err = core.GetServiceProperties(distributionlistapiv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			// Load test-specific configuration from environment variables
			accountID = os.Getenv("DISTRIBUTION_LIST_API_TEST_ACCOUNT_ID")
			if accountID == "" {
				Skip("DISTRIBUTION_LIST_API_TEST_ACCOUNT_ID not set, skipping tests")
			}

			instanceID = os.Getenv("DISTRIBUTION_LIST_API_TEST_INSTANCE_ID")
			if instanceID == "" {
				Skip("DISTRIBUTION_LIST_API_TEST_INSTANCE_ID not set, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			fmt.Fprintf(GinkgoWriter, "Account ID: %v\n", accountID)
			fmt.Fprintf(GinkgoWriter, "Instance ID: %v\n", instanceID)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			distributionListApiServiceOptions := &distributionlistapiv1.DistributionListApiV1Options{}

			distributionListApiService, err = distributionlistapiv1.NewDistributionListApiV1UsingExternalConfig(distributionListApiServiceOptions)
			Expect(err).To(BeNil())
			Expect(distributionListApiService).ToNot(BeNil())
			Expect(distributionListApiService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			distributionListApiService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListDistributionListDestinations - Get all destination entries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDistributionListDestinations(listDistributionListDestinationsOptions *ListDistributionListDestinationsOptions)`, func() {
			listDistributionListDestinationsOptions := &distributionlistapiv1.ListDistributionListDestinationsOptions{
				AccountID: core.StringPtr(accountID),
			}

			addDestinationCollection, response, err := distributionListApiService.ListDistributionListDestinations(listDistributionListDestinationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addDestinationCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateDistributionListDestination - Add a destination entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDistributionListDestination(createDistributionListDestinationOptions *CreateDistributionListDestinationOptions)`, func() {
			addDestinationPrototypeModel := &distributionlistapiv1.AddDestinationPrototypeEventNotificationDestination{
				ID:              CreateMockUUID(instanceID),
				DestinationType: core.StringPtr("event_notifications"),
			}

			createDistributionListDestinationOptions := &distributionlistapiv1.CreateDistributionListDestinationOptions{
				AccountID:               core.StringPtr(accountID),
				AddDestinationPrototype: addDestinationPrototypeModel,
			}

			addDestination, response, err := distributionListApiService.CreateDistributionListDestination(createDistributionListDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(addDestination).ToNot(BeNil())
		})
	})

	Describe(`GetDistributionListDestination - Get a destination entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDistributionListDestination(getDistributionListDestinationOptions *GetDistributionListDestinationOptions)`, func() {
			getDistributionListDestinationOptions := &distributionlistapiv1.GetDistributionListDestinationOptions{
				AccountID: core.StringPtr(accountID),
				ID:        core.StringPtr(instanceID),
			}

			addDestination, response, err := distributionListApiService.GetDistributionListDestination(getDistributionListDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addDestination).ToNot(BeNil())
		})
	})

	Describe(`TestDistributionListDestination - Test destination entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`TestDistributionListDestination(testDistributionListDestinationOptions *TestDistributionListDestinationOptions)`, func() {
			testDestinationRequestBodyPrototypeModel := &distributionlistapiv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestination{
				DestinationType:  core.StringPtr("event_notifications"),
				NotificationType: core.StringPtr("incident"),
			}

			testDistributionListDestinationOptions := &distributionlistapiv1.TestDistributionListDestinationOptions{
				AccountID:                           core.StringPtr(accountID),
				ID:                                  core.StringPtr(instanceID),
				TestDestinationRequestBodyPrototype: testDestinationRequestBodyPrototypeModel,
			}

			testDestinationResponseBody, response, err := distributionListApiService.TestDistributionListDestination(testDistributionListDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testDestinationResponseBody).ToNot(BeNil())
		})
	})

	Describe(`DeleteDistributionListDestination - Delete destination entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDistributionListDestination(deleteDistributionListDestinationOptions *DeleteDistributionListDestinationOptions)`, func() {
			deleteDistributionListDestinationOptions := &distributionlistapiv1.DeleteDistributionListDestinationOptions{
				AccountID: core.StringPtr(accountID),
				ID:        core.StringPtr(instanceID),
			}

			response, err := distributionListApiService.DeleteDistributionListDestination(deleteDistributionListDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
