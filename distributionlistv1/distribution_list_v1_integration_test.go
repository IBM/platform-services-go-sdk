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

package distributionlistv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/distributionlistv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the distributionlistv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`DistributionListV1 Integration Tests`, func() {
	const externalConfigFile = "../distribution_list_v1.env"

	var (
		err          error
		distributionListService *distributionlistv1.DistributionListV1
		serviceURL   string
		config       map[string]string
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
			config, err = core.GetServiceProperties(distributionlistv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			distributionListServiceOptions := &distributionlistv1.DistributionListV1Options{}

			distributionListService, err = distributionlistv1.NewDistributionListV1UsingExternalConfig(distributionListServiceOptions)
			Expect(err).To(BeNil())
			Expect(distributionListService).ToNot(BeNil())
			Expect(distributionListService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			distributionListService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListDistributionListDestinations - Get all destination entries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDistributionListDestinations(listDistributionListDestinationsOptions *ListDistributionListDestinationsOptions)`, func() {
			listDistributionListDestinationsOptions := &distributionlistv1.ListDistributionListDestinationsOptions{
				AccountID: core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"),
			}

			addDestinationCollection, response, err := distributionListService.ListDistributionListDestinations(listDistributionListDestinationsOptions)
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
			addDestinationPrototypeModel := &distributionlistv1.AddDestinationPrototypeEventNotificationDestinationPrototype{
				DestinationID: CreateMockUUID("12345678-1234-1234-1234-123456789012"),
				DestinationType: core.StringPtr("event_notifications"),
			}

			createDistributionListDestinationOptions := &distributionlistv1.CreateDistributionListDestinationOptions{
				AccountID: core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"),
				AddDestinationPrototype: addDestinationPrototypeModel,
			}

			addDestination, response, err := distributionListService.CreateDistributionListDestination(createDistributionListDestinationOptions)
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
			getDistributionListDestinationOptions := &distributionlistv1.GetDistributionListDestinationOptions{
				AccountID: core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"),
				DestinationID: core.StringPtr("12345678-1234-1234-1234-123456789012"),
			}

			addDestination, response, err := distributionListService.GetDistributionListDestination(getDistributionListDestinationOptions)
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
			testDestinationRequestBodyPrototypeModel := &distributionlistv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestination{
				DestinationType: core.StringPtr("event_notifications"),
				NotificationType: core.StringPtr("incident"),
			}

			testDistributionListDestinationOptions := &distributionlistv1.TestDistributionListDestinationOptions{
				AccountID: core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"),
				DestinationID: core.StringPtr("12345678-1234-1234-1234-123456789012"),
				TestDestinationRequestBodyPrototype: testDestinationRequestBodyPrototypeModel,
			}

			testDestinationResponseBody, response, err := distributionListService.TestDistributionListDestination(testDistributionListDestinationOptions)
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
			deleteDistributionListDestinationOptions := &distributionlistv1.DeleteDistributionListDestinationOptions{
				AccountID: core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"),
				DestinationID: core.StringPtr("12345678-1234-1234-1234-123456789012"),
			}

			response, err := distributionListService.DeleteDistributionListDestination(deleteDistributionListDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
