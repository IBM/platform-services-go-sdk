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

package platformnotificationsv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/platformnotificationsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Platform Notifications service.
//
// The following configuration properties are assumed to be defined:
// PLATFORM_NOTIFICATIONS_URL=<service base url>
// PLATFORM_NOTIFICATIONS_AUTH_TYPE=iam
// PLATFORM_NOTIFICATIONS_APIKEY=<IAM apikey>
// PLATFORM_NOTIFICATIONS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`PlatformNotificationsV1 Examples Tests`, func() {

	const externalConfigFile = "../platform_notifications_v1.env"

	var (
		platformNotificationsService *platformnotificationsv1.PlatformNotificationsV1
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
			config, err = core.GetServiceProperties(platformnotificationsv1.DefaultServiceName)
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

			platformNotificationsServiceOptions := &platformnotificationsv1.PlatformNotificationsV1Options{}

			platformNotificationsService, err = platformnotificationsv1.NewPlatformNotificationsV1UsingExternalConfig(platformNotificationsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(platformNotificationsService).ToNot(BeNil())
		})
	})

	Describe(`PlatformNotificationsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDistributionListDestinations request example`, func() {
			fmt.Println("\nListDistributionListDestinations() result:")
			// begin-list_distribution_list_destinations

			listDistributionListDestinationsOptions := platformNotificationsService.NewListDistributionListDestinationsOptions(
				"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
			)

			addDestinationCollection, response, err := platformNotificationsService.ListDistributionListDestinations(listDistributionListDestinationsOptions)
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

			addDestinationPrototypeModel := &platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype{
				DestinationID: CreateMockUUID("12345678-1234-1234-1234-123456789012"),
				DestinationType: core.StringPtr("event_notifications"),
			}

			createDistributionListDestinationOptions := platformNotificationsService.NewCreateDistributionListDestinationOptions(
				"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
				addDestinationPrototypeModel,
			)

			addDestination, response, err := platformNotificationsService.CreateDistributionListDestination(createDistributionListDestinationOptions)
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

			getDistributionListDestinationOptions := platformNotificationsService.NewGetDistributionListDestinationOptions(
				"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
				"12345678-1234-1234-1234-123456789012",
			)

			addDestination, response, err := platformNotificationsService.GetDistributionListDestination(getDistributionListDestinationOptions)
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

			testDestinationRequestBodyPrototypeModel := &platformnotificationsv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype{
				DestinationType: core.StringPtr("event_notifications"),
				NotificationType: core.StringPtr("incident"),
			}

			testDistributionListDestinationOptions := platformNotificationsService.NewTestDistributionListDestinationOptions(
				"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
				"12345678-1234-1234-1234-123456789012",
				testDestinationRequestBodyPrototypeModel,
			)

			testDestinationResponseBody, response, err := platformNotificationsService.TestDistributionListDestination(testDistributionListDestinationOptions)
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
		It(`CreatePreferences request example`, func() {
			fmt.Println("\nCreatePreferences() result:")
			// begin-create_preferences

			preferenceValueWithUpdatesModel := &platformnotificationsv1.PreferenceValueWithUpdates{
				Channels: []string{"email"},
				Updates: core.BoolPtr(true),
			}

			preferenceValueWithoutUpdatesModel := &platformnotificationsv1.PreferenceValueWithoutUpdates{
				Channels: []string{"email"},
			}

			createPreferencesOptions := platformNotificationsService.NewCreatePreferencesOptions(
				"IBMid-1234567890",
			)
			createPreferencesOptions.SetIncidentSeverity1(preferenceValueWithUpdatesModel)
			createPreferencesOptions.SetOrderingReview(preferenceValueWithoutUpdatesModel)
			createPreferencesOptions.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")

			preferencesObject, response, err := platformNotificationsService.CreatePreferences(createPreferencesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(preferencesObject, "", "  ")
			fmt.Println(string(b))

			// end-create_preferences

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(preferencesObject).ToNot(BeNil())
		})
		It(`GetPreferences request example`, func() {
			fmt.Println("\nGetPreferences() result:")
			// begin-get_preferences

			getPreferencesOptions := platformNotificationsService.NewGetPreferencesOptions(
				"IBMid-1234567890",
			)
			getPreferencesOptions.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")

			preferencesObject, response, err := platformNotificationsService.GetPreferences(getPreferencesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(preferencesObject, "", "  ")
			fmt.Println(string(b))

			// end-get_preferences

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(preferencesObject).ToNot(BeNil())
		})
		It(`ReplaceNotificationPreferences request example`, func() {
			fmt.Println("\nReplaceNotificationPreferences() result:")
			// begin-replace_notification_preferences

			preferenceValueWithUpdatesModel := &platformnotificationsv1.PreferenceValueWithUpdates{
				Channels: []string{"email"},
				Updates: core.BoolPtr(true),
			}

			preferenceValueWithoutUpdatesModel := &platformnotificationsv1.PreferenceValueWithoutUpdates{
				Channels: []string{"email"},
			}

			replaceNotificationPreferencesOptions := platformNotificationsService.NewReplaceNotificationPreferencesOptions(
				"IBMid-1234567890",
			)
			replaceNotificationPreferencesOptions.SetIncidentSeverity1(preferenceValueWithUpdatesModel)
			replaceNotificationPreferencesOptions.SetOrderingReview(preferenceValueWithoutUpdatesModel)
			replaceNotificationPreferencesOptions.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")

			preferencesObject, response, err := platformNotificationsService.ReplaceNotificationPreferences(replaceNotificationPreferencesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(preferencesObject, "", "  ")
			fmt.Println(string(b))

			// end-replace_notification_preferences

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(preferencesObject).ToNot(BeNil())
		})
		It(`ListNotifications request example`, func() {
			fmt.Println("\nListNotifications() result:")
			// begin-list_notifications
			listNotificationsOptions := &platformnotificationsv1.ListNotificationsOptions{
				AccountID: core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"),
				Limit: core.Int64Ptr(int64(50)),
			}

			pager, err := platformNotificationsService.NewNotificationsPager(listNotificationsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []platformnotificationsv1.Notification
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_notifications
		})
		It(`GetAcknowledgment request example`, func() {
			fmt.Println("\nGetAcknowledgment() result:")
			// begin-get_acknowledgment

			getAcknowledgmentOptions := platformNotificationsService.NewGetAcknowledgmentOptions()
			getAcknowledgmentOptions.SetAccountID("1369339417d906e5620b8d861d40cfd7")

			acknowledgment, response, err := platformNotificationsService.GetAcknowledgment(getAcknowledgmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(acknowledgment, "", "  ")
			fmt.Println(string(b))

			// end-get_acknowledgment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(acknowledgment).ToNot(BeNil())
		})
		It(`ReplaceNotificationAcknowledgment request example`, func() {
			fmt.Println("\nReplaceNotificationAcknowledgment() result:")
			// begin-replace_notification_acknowledgment

			replaceNotificationAcknowledgmentOptions := platformNotificationsService.NewReplaceNotificationAcknowledgmentOptions(
				"1772804159452",
			)
			replaceNotificationAcknowledgmentOptions.SetAccountID("1369339417d906e5620b8d861d40cfd7")

			acknowledgment, response, err := platformNotificationsService.ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(acknowledgment, "", "  ")
			fmt.Println(string(b))

			// end-replace_notification_acknowledgment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(acknowledgment).ToNot(BeNil())
		})
		It(`DeleteDistributionListDestination request example`, func() {
			// begin-delete_distribution_list_destination

			deleteDistributionListDestinationOptions := platformNotificationsService.NewDeleteDistributionListDestinationOptions(
				"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
				"12345678-1234-1234-1234-123456789012",
			)

			response, err := platformNotificationsService.DeleteDistributionListDestination(deleteDistributionListDestinationOptions)
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
		It(`DeleteNotificationPreferences request example`, func() {
			// begin-delete_notification_preferences

			deleteNotificationPreferencesOptions := platformNotificationsService.NewDeleteNotificationPreferencesOptions(
				"IBMid-1234567890",
			)
			deleteNotificationPreferencesOptions.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")

			response, err := platformNotificationsService.DeleteNotificationPreferences(deleteNotificationPreferencesOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteNotificationPreferences(): %d\n", response.StatusCode)
			}

			// end-delete_notification_preferences

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
