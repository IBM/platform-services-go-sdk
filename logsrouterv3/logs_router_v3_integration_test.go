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

package logsrouterv3_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the logsrouterv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`LogsRouterV3 Integration Tests`, func() {
	const externalConfigFile = "../logs_router_v3.env"

	var (
		err               error
		logsRouterService *logsrouterv3.LogsRouterV3
		serviceURL        string
		config            map[string]string

		// Variables to hold link values
		routeIDLink  string
		targetIDLink string
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
			config, err = core.GetServiceProperties(logsrouterv3.DefaultServiceName)
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
			logsRouterServiceOptions := &logsrouterv3.LogsRouterV3Options{}

			logsRouterService, err = logsrouterv3.NewLogsRouterV3UsingExternalConfig(logsRouterServiceOptions)
			Expect(err).To(BeNil())
			Expect(logsRouterService).ToNot(BeNil())
			Expect(logsRouterService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			logsRouterService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateTarget - Create a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {
			// Manual code to configure setting first
			updateSettingsOptions := &logsrouterv3.UpdateSettingsOptions{
				PermittedTargetRegions: []string{"us-south", "us-east"},
				PrimaryMetadataRegion:  core.StringPtr("us-south"),
				BackupMetadataRegion:   core.StringPtr("us-east"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
			}
			setting, response, err := logsRouterService.UpdateSettings(updateSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setting).ToNot(BeNil())

			createTargetOptions := &logsrouterv3.CreateTargetOptions{
				Name:           core.StringPtr("my-lr-target"),
				DestinationCRN: core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Region:         core.StringPtr("us-south"),
				ManagedBy:      core.StringPtr("enterprise"),
			}

			target, response, err := logsRouterService.CreateTarget(createTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID
			fmt.Fprintf(GinkgoWriter, "Saved targetIDLink value: %v\n", targetIDLink)
		})
	})

	Describe(`CreateRoute - Create a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRoute(createRouteOptions *CreateRouteOptions)`, func() {
			targetIdentityModel := &logsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}

			inclusionFilterPrototypeModel := &logsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Values:   []string{"us-south"},
			}

			rulePrototypeModel := &logsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []logsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}

			createRouteOptions := &logsrouterv3.CreateRouteOptions{
				Name:      core.StringPtr("my-route"),
				Rules:     []logsrouterv3.RulePrototype{*rulePrototypeModel},
				ManagedBy: core.StringPtr("enterprise"),
			}

			route, response, err := logsRouterService.CreateRoute(createRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

			routeIDLink = *route.ID
			fmt.Fprintf(GinkgoWriter, "Saved routeIDLink value: %v\n", routeIDLink)
		})
	})

	Describe(`ListTargets - List targets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTargets(listTargetsOptions *ListTargetsOptions)`, func() {
			listTargetsOptions := &logsrouterv3.ListTargetsOptions{}

			targetCollection, response, err := logsRouterService.ListTargets(listTargetsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetCollection).ToNot(BeNil())
		})
	})

	Describe(`GetTarget - Get details of a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {
			getTargetOptions := &logsrouterv3.GetTargetOptions{
				ID: &targetIDLink,
			}

			target, response, err := logsRouterService.GetTarget(getTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
	})

	Describe(`UpdateTarget - Update a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTarget(updateTargetOptions *UpdateTargetOptions)`, func() {
			updateTargetOptions := &logsrouterv3.UpdateTargetOptions{
				ID:             &targetIDLink,
				Name:           core.StringPtr("my-lr-target"),
				DestinationCRN: core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
			}

			target, response, err := logsRouterService.UpdateTarget(updateTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
	})

	Describe(`ListRoutes - List routes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRoutes(listRoutesOptions *ListRoutesOptions)`, func() {
			listRoutesOptions := &logsrouterv3.ListRoutesOptions{}

			routeCollection, response, err := logsRouterService.ListRoutes(listRoutesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeCollection).ToNot(BeNil())
		})
	})

	Describe(`GetRoute - Get details of a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRoute(getRouteOptions *GetRouteOptions)`, func() {
			getRouteOptions := &logsrouterv3.GetRouteOptions{
				ID: &routeIDLink,
			}

			route, response, err := logsRouterService.GetRoute(getRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
		})
	})

	Describe(`UpdateRoute - Update a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateRoute(updateRouteOptions *UpdateRouteOptions)`, func() {
			targetIdentityModel := &logsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}

			inclusionFilterPrototypeModel := &logsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Values:   []string{"us-south"},
			}

			rulePrototypeModel := &logsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []logsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}

			updateRouteOptions := &logsrouterv3.UpdateRouteOptions{
				ID:    &routeIDLink,
				Name:  core.StringPtr("my-route"),
				Rules: []logsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := logsRouterService.UpdateRoute(updateRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
		})
	})

	Describe(`GetSettings - Get settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
			getSettingsOptions := &logsrouterv3.GetSettingsOptions{}

			setting, response, err := logsRouterService.GetSettings(getSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setting).ToNot(BeNil())
		})
	})

	Describe(`UpdateSettings - Modify settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
			// Manual code: create a target without enterprise so that it can be used as default target
			createTargetOptions := &logsrouterv3.CreateTargetOptions{
				Name:           core.StringPtr("lr-target2"),
				DestinationCRN: core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Region:         core.StringPtr("us-south"),
			}

			target, response, err := logsRouterService.CreateTarget(createTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			defaultTargetIDLink := *target.ID
			// End of manual code

			targetIdentityModel := &logsrouterv3.TargetIdentity{
				ID: &defaultTargetIDLink,
			}

			updateSettingsOptions := &logsrouterv3.UpdateSettingsOptions{
				DefaultTargets:         []logsrouterv3.TargetIdentity{*targetIdentityModel},
				PermittedTargetRegions: []string{"us-south"},
				PrimaryMetadataRegion:  core.StringPtr("us-south"),
				BackupMetadataRegion:   core.StringPtr("us-east"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
			}

			setting, response, err := logsRouterService.UpdateSettings(updateSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setting).ToNot(BeNil())
		})
	})

	Describe(`QueryDestinations - Query Destinations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`QueryDestinations(queryDestinationsOptions *QueryDestinationsOptions)`, func() {
			crnPrototypeModel := &logsrouterv3.CRNPrototype{
				CRN: core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/6a1d10334a2e4dd197d4e301e8f87df9:22222222-2222-2222-2222-222222222222::"),
			}

			queryDestinationsOptions := &logsrouterv3.QueryDestinationsOptions{
				Crns: []logsrouterv3.CRNPrototype{*crnPrototypeModel},
			}

			destinationsQuery, response, err := logsRouterService.QueryDestinations(queryDestinationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationsQuery).ToNot(BeNil())
		})
	})

	Describe(`DeleteRoute - Delete a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRoute(deleteRouteOptions *DeleteRouteOptions)`, func() {
			deleteRouteOptions := &logsrouterv3.DeleteRouteOptions{
				ID: &routeIDLink,
			}

			response, err := logsRouterService.DeleteRoute(deleteRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTarget - Delete a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {
			deleteTargetOptions := &logsrouterv3.DeleteTargetOptions{
				ID: &targetIDLink,
			}

			response, err := logsRouterService.DeleteTarget(deleteTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
