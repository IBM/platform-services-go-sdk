//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/distributionlistapiv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Distribution List API service.
//
// The following configuration properties are assumed to be defined:
// DISTRIBUTION_LIST_API_URL=<service base url>
// DISTRIBUTION_LIST_API_AUTH_TYPE=iam
// DISTRIBUTION_LIST_API_APIKEY=<IAM apikey>
// DISTRIBUTION_LIST_API_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`DistributionListApiV1 Examples Tests`, func() {

	const externalConfigFile = "../distribution_list_api_v1.env"

	var (
		distributionListApiService *distributionlistapiv1.DistributionListApiV1
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(distributionlistapiv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			distributionListApiServiceOptions := &distributionlistapiv1.DistributionListApiV1Options{}

			distributionListApiService, err = distributionlistapiv1.NewDistributionListApiV1UsingExternalConfig(distributionListApiServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(distributionListApiService).ToNot(BeNil())
		})
	})

	Describe(`DistributionListApiV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDistributionListDestinations request example`, func() {
			fmt.Println("\nListDistributionListDestinations() result:")
			// begin-list_distribution_list_destinations

			listDistributionListDestinationsOptions := distributionListApiService.NewListDistributionListDestinationsOptions(
				"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
			)

			addDestinationCollection, response, err := distributionListApiService.ListDistributionListDestinations(listDistributionListDestinationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addDestinationCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_distribution_list_destinations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addDestinationCollection).ToNot(BeNil())
		})
		It(`CreateDistributionListDestination request example`, func() {
			fmt.Println("\nCreateDistributionListDestination() result:")
			// begin-create_distribution_list_destination

			addDestinationPrototypeModel := &distributionlistapiv1.AddDestinationPrototypeEventNotificationDestination{
				ID: CreateMockUUID("12345678-1234-1234-1234-123456789012"),
				DestinationType: core.StringPtr("event_notifications"),
			}

			createDistributionListDestinationOptions := distributionListApiService.NewCreateDistributionListDestinationOptions(
				"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
				addDestinationPrototypeModel,
			)

			addDestination, response, err := distributionListApiService.CreateDistributionListDestination(createDistributionListDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addDestination, "", "  ")
			fmt.Println(string(b))

			// end-create_distribution_list_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(addDestination).ToNot(BeNil())
		})
		It(`GetDistributionListDestination request example`, func() {
			fmt.Println("\nGetDistributionListDestination() result:")
			// begin-get_distribution_list_destination

			getDistributionListDestinationOptions := distributionListApiService.NewGetDistributionListDestinationOptions(
				"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
				"12345678-1234-1234-1234-123456789012",
			)

			addDestination, response, err := distributionListApiService.GetDistributionListDestination(getDistributionListDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addDestination, "", "  ")
			fmt.Println(string(b))

			// end-get_distribution_list_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addDestination).ToNot(BeNil())
		})
		It(`TestDistributionListDestination request example`, func() {
			fmt.Println("\nTestDistributionListDestination() result:")
			// begin-test_distribution_list_destination

			testDestinationRequestBodyPrototypeModel := &distributionlistapiv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestination{
				DestinationType: core.StringPtr("event_notifications"),
				NotificationType: core.StringPtr("incident"),
			}

			testDistributionListDestinationOptions := distributionListApiService.NewTestDistributionListDestinationOptions(
				"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
				"12345678-1234-1234-1234-123456789012",
				testDestinationRequestBodyPrototypeModel,
			)

			testDestinationResponseBody, response, err := distributionListApiService.TestDistributionListDestination(testDistributionListDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(testDestinationResponseBody, "", "  ")
			fmt.Println(string(b))

			// end-test_distribution_list_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testDestinationResponseBody).ToNot(BeNil())
		})
		It(`DeleteDistributionListDestination request example`, func() {
			// begin-delete_distribution_list_destination

			deleteDistributionListDestinationOptions := distributionListApiService.NewDeleteDistributionListDestinationOptions(
				"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
				"12345678-1234-1234-1234-123456789012",
			)

			response, err := distributionListApiService.DeleteDistributionListDestination(deleteDistributionListDestinationOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteDistributionListDestination(): %d\n", response.StatusCode)
			}

			// end-delete_distribution_list_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
