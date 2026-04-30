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
		IamID                        string
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

			IamID = config["TEST_IAM_ID"]
			if IamID == "" {
				Skip("PLATFORM_NOTIFICATIONS_TEST_IAM_ID not found in configuration, skipping tests")
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

	Describe(`CreatePreferences - Create communication preferences`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePreferences(createPreferencesOptions *CreatePreferencesOptions)`, func() {
			preferenceValueWithUpdatesModel := &platformnotificationsv1.PreferenceValueWithUpdates{
				Channels: []string{"email"},
				Updates:  core.BoolPtr(true),
			}

			preferenceValueWithoutUpdatesModel := &platformnotificationsv1.PreferenceValueWithoutUpdates{
				Channels: []string{"email"},
			}

			createPreferencesOptions := &platformnotificationsv1.CreatePreferencesOptions{
				IamID:                   core.StringPtr(IamID),
				IncidentSeverity1:       preferenceValueWithUpdatesModel,
				IncidentSeverity2:       preferenceValueWithUpdatesModel,
				IncidentSeverity3:       preferenceValueWithUpdatesModel,
				IncidentSeverity4:       preferenceValueWithUpdatesModel,
				MaintenanceHigh:         preferenceValueWithUpdatesModel,
				MaintenanceMedium:       preferenceValueWithUpdatesModel,
				MaintenanceLow:          preferenceValueWithUpdatesModel,
				AnnouncementsMajor:      preferenceValueWithoutUpdatesModel,
				AnnouncementsMinor:      preferenceValueWithoutUpdatesModel,
				SecurityNormal:          preferenceValueWithoutUpdatesModel,
				AccountNormal:           preferenceValueWithoutUpdatesModel,
				BillingAndUsageOrder:    preferenceValueWithoutUpdatesModel,
				BillingAndUsageInvoices: preferenceValueWithoutUpdatesModel,
				BillingAndUsagePayments: preferenceValueWithoutUpdatesModel,
				BillingAndUsageSubscriptionsAndPromoCodes: preferenceValueWithoutUpdatesModel,
				BillingAndUsageSpendingAlerts:             preferenceValueWithoutUpdatesModel,
				ResourceactivityNormal:                    preferenceValueWithoutUpdatesModel,
				OrderingReview:                            preferenceValueWithoutUpdatesModel,
				OrderingApproved:                          preferenceValueWithoutUpdatesModel,
				OrderingApprovedVsi:                       preferenceValueWithoutUpdatesModel,
				OrderingApprovedServer:                    preferenceValueWithoutUpdatesModel,
				ProvisioningReloadComplete:                preferenceValueWithoutUpdatesModel,
				ProvisioningCompleteVsi:                   preferenceValueWithoutUpdatesModel,
				ProvisioningCompleteServer:                preferenceValueWithoutUpdatesModel,
				AccountID:                                 core.StringPtr(accountID),
			}

			preferencesObject, response, err := platformNotificationsService.CreatePreferences(createPreferencesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(preferencesObject).ToNot(BeNil())
		})
	})

	Describe(`GetPreferences - Get all communication preferences for a user in an account`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPreferences(getPreferencesOptions *GetPreferencesOptions)`, func() {
			getPreferencesOptions := &platformnotificationsv1.GetPreferencesOptions{
				IamID:     core.StringPtr(IamID),
				AccountID: core.StringPtr(accountID),
			}

			preferencesObject, response, err := platformNotificationsService.GetPreferences(getPreferencesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(preferencesObject).ToNot(BeNil())
		})
	})

	Describe(`ReplaceNotificationPreferences - Update communication preferences`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceNotificationPreferences(replaceNotificationPreferencesOptions *ReplaceNotificationPreferencesOptions)`, func() {
			preferenceValueWithUpdatesModel := &platformnotificationsv1.PreferenceValueWithUpdates{
				Channels: []string{"email"},
				Updates:  core.BoolPtr(true),
			}

			preferenceValueWithoutUpdatesModel := &platformnotificationsv1.PreferenceValueWithoutUpdates{
				Channels: []string{"email"},
			}

			replaceNotificationPreferencesOptions := &platformnotificationsv1.ReplaceNotificationPreferencesOptions{
				IamID:                   core.StringPtr(IamID),
				IncidentSeverity1:       preferenceValueWithUpdatesModel,
				IncidentSeverity2:       preferenceValueWithUpdatesModel,
				IncidentSeverity3:       preferenceValueWithUpdatesModel,
				IncidentSeverity4:       preferenceValueWithUpdatesModel,
				MaintenanceHigh:         preferenceValueWithUpdatesModel,
				MaintenanceMedium:       preferenceValueWithUpdatesModel,
				MaintenanceLow:          preferenceValueWithUpdatesModel,
				AnnouncementsMajor:      preferenceValueWithoutUpdatesModel,
				AnnouncementsMinor:      preferenceValueWithoutUpdatesModel,
				SecurityNormal:          preferenceValueWithoutUpdatesModel,
				AccountNormal:           preferenceValueWithoutUpdatesModel,
				BillingAndUsageOrder:    preferenceValueWithoutUpdatesModel,
				BillingAndUsageInvoices: preferenceValueWithoutUpdatesModel,
				BillingAndUsagePayments: preferenceValueWithoutUpdatesModel,
				BillingAndUsageSubscriptionsAndPromoCodes: preferenceValueWithoutUpdatesModel,
				BillingAndUsageSpendingAlerts:             preferenceValueWithoutUpdatesModel,
				ResourceactivityNormal:                    preferenceValueWithoutUpdatesModel,
				OrderingReview:                            preferenceValueWithoutUpdatesModel,
				OrderingApproved:                          preferenceValueWithoutUpdatesModel,
				OrderingApprovedVsi:                       preferenceValueWithoutUpdatesModel,
				OrderingApprovedServer:                    preferenceValueWithoutUpdatesModel,
				ProvisioningReloadComplete:                preferenceValueWithoutUpdatesModel,
				ProvisioningCompleteVsi:                   preferenceValueWithoutUpdatesModel,
				ProvisioningCompleteServer:                preferenceValueWithoutUpdatesModel,
				AccountID:                                 core.StringPtr(accountID),
			}

			preferencesObject, response, err := platformNotificationsService.ReplaceNotificationPreferences(replaceNotificationPreferencesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(preferencesObject).ToNot(BeNil())
		})
	})

	Describe(`ListNotifications - Get user notifications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNotifications(listNotificationsOptions *ListNotificationsOptions) with pagination`, func() {
			listNotificationsOptions := &platformnotificationsv1.ListNotificationsOptions{
				AccountID: core.StringPtr(accountID),
				Start:     core.StringPtr("3fe78a36b9aa7f26"),
				Limit:     core.Int64Ptr(int64(50)),
			}

			listNotificationsOptions.Start = nil
			listNotificationsOptions.Limit = core.Int64Ptr(1)

			var allResults []platformnotificationsv1.Notification
			for {
				notificationCollection, response, err := platformNotificationsService.ListNotifications(listNotificationsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(notificationCollection).ToNot(BeNil())
				allResults = append(allResults, notificationCollection.Notifications...)

				listNotificationsOptions.Start, err = notificationCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listNotificationsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListNotifications(listNotificationsOptions *ListNotificationsOptions) using NotificationsPager`, func() {
			listNotificationsOptions := &platformnotificationsv1.ListNotificationsOptions{
				AccountID: core.StringPtr(accountID),
				Limit:     core.Int64Ptr(int64(50)),
			}

			// Test GetNext().
			pager, err := platformNotificationsService.NewNotificationsPager(listNotificationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []platformnotificationsv1.Notification
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = platformNotificationsService.NewNotificationsPager(listNotificationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListNotifications() returned a total of %d item(s) using NotificationsPager.\n", len(allResults))
		})
	})

	// Describe(`GetAcknowledgment - Get user's last acknowledged notification Id`, func() {
	// 	BeforeEach(func() {
	// 		shouldSkipTest()
	// 	})
	// 	It(`GetAcknowledgment(getAcknowledgmentOptions *GetAcknowledgmentOptions)`, func() {
	// 		getAcknowledgmentOptions := &platformnotificationsv1.GetAcknowledgmentOptions{
	// 			AccountID: core.StringPtr(accountID),
	// 		}

	// 		acknowledgment, response, err := platformNotificationsService.GetAcknowledgment(getAcknowledgmentOptions)
	// 		Expect(err).To(BeNil())
	// 		Expect(response.StatusCode).To(Equal(200))
	// 		Expect(acknowledgment).ToNot(BeNil())
	// 	})
	// })

	// Describe(`ReplaceNotificationAcknowledgment - Update user's last acknowledged notification`, func() {
	// 	BeforeEach(func() {
	// 		shouldSkipTest()
	// 	})
	// 	It(`ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptions *ReplaceNotificationAcknowledgmentOptions)`, func() {
	// 		replaceNotificationAcknowledgmentOptions := &platformnotificationsv1.ReplaceNotificationAcknowledgmentOptions{
	// 			LastAcknowledgedID: core.StringPtr("1772804159452"),
	// 			AccountID:          core.StringPtr(accountID),
	// 		}

	// 		acknowledgment, response, err := platformNotificationsService.ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptions)
	// 		Expect(err).To(BeNil())
	// 		Expect(response.StatusCode).To(Equal(200))
	// 		Expect(acknowledgment).ToNot(BeNil())
	// 	})
	// })

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

	Describe(`DeleteNotificationPreferences - Resets all preferences to their default values`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNotificationPreferences(deleteNotificationPreferencesOptions *DeleteNotificationPreferencesOptions)`, func() {
			deleteNotificationPreferencesOptions := &platformnotificationsv1.DeleteNotificationPreferencesOptions{
				IamID:     core.StringPtr(IamID),
				AccountID: core.StringPtr(accountID),
			}

			response, err := platformNotificationsService.DeleteNotificationPreferences(deleteNotificationPreferencesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
