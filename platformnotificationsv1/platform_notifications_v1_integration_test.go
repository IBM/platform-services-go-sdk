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

package platformnotificationsv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/platformnotificationsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the platformnotificationsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PlatformNotificationsV1 Integration Tests`, func() {
	const externalConfigFile = "../platform_notifications_v1.env"

	var (
		err                          error
		platformNotificationsService *platformnotificationsv1.PlatformNotificationsV1
		serviceURL                   string
		config                       map[string]string
		accountID                    string
		instanceID                   string
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
			config, err = core.GetServiceProperties(platformnotificationsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			// Load test account ID and instance ID from config
			accountID = config["TEST_ACCOUNT_ID"]
			if accountID == "" {
				Skip("PLATFORM_NOTIFICATIONS_TEST_ACCOUNT_ID not found in configuration, skipping tests")
			}

			instanceID = config["TEST_INSTANCE_ID"]
			if instanceID == "" {
				Skip("PLATFORM_NOTIFICATIONS_TEST_INSTANCE_ID not found in configuration, skipping tests")
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
			platformNotificationsServiceOptions := &platformnotificationsv1.PlatformNotificationsV1Options{}

			platformNotificationsService, err = platformnotificationsv1.NewPlatformNotificationsV1UsingExternalConfig(platformNotificationsServiceOptions)
			Expect(err).To(BeNil())
			Expect(platformNotificationsService).ToNot(BeNil())
			Expect(platformNotificationsService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			platformNotificationsService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListDistributionListDestinations - Get all destination entries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDistributionListDestinations(listDistributionListDestinationsOptions *ListDistributionListDestinationsOptions)`, func() {
			listDistributionListDestinationsOptions := &platformnotificationsv1.ListDistributionListDestinationsOptions{
				AccountID: core.StringPtr(accountID),
			}

			addDestinationCollection, response, err := platformNotificationsService.ListDistributionListDestinations(listDistributionListDestinationsOptions)
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
			addDestinationPrototypeModel := &platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype{
				DestinationID:   CreateMockUUID(instanceID),
				DestinationType: core.StringPtr("event_notifications"),
			}

			createDistributionListDestinationOptions := &platformnotificationsv1.CreateDistributionListDestinationOptions{
				AccountID:               core.StringPtr(accountID),
				AddDestinationPrototype: addDestinationPrototypeModel,
			}

			addDestination, response, err := platformNotificationsService.CreateDistributionListDestination(createDistributionListDestinationOptions)
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
			getDistributionListDestinationOptions := &platformnotificationsv1.GetDistributionListDestinationOptions{
				AccountID:     core.StringPtr(accountID),
				DestinationID: core.StringPtr(instanceID),
			}

			addDestination, response, err := platformNotificationsService.GetDistributionListDestination(getDistributionListDestinationOptions)
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
			testDestinationRequestBodyPrototypeModel := &platformnotificationsv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype{
				DestinationType:  core.StringPtr("event_notifications"),
				NotificationType: core.StringPtr("incident"),
			}

			testDistributionListDestinationOptions := &platformnotificationsv1.TestDistributionListDestinationOptions{
				AccountID:                           core.StringPtr(accountID),
				DestinationID:                       core.StringPtr(instanceID),
				TestDestinationRequestBodyPrototype: testDestinationRequestBodyPrototypeModel,
			}

			testDestinationResponseBody, response, err := platformNotificationsService.TestDistributionListDestination(testDistributionListDestinationOptions)
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
			deleteDistributionListDestinationOptions := &platformnotificationsv1.DeleteDistributionListDestinationOptions{
				AccountID:     core.StringPtr(accountID),
				DestinationID: core.StringPtr(instanceID),
			}

			response, err := platformNotificationsService.DeleteDistributionListDestination(deleteDistributionListDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
