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

package atrackerv2_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the atrackerv2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`AtrackerV2 Integration Tests`, func() {
	const externalConfigFile = "../atracker_v2.env"

	var (
		err             error
		atrackerService *atrackerv2.AtrackerV2
		serviceURL      string
		config          map[string]string

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
			config, err = core.GetServiceProperties(atrackerv2.DefaultServiceName)
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
			atrackerServiceOptions := &atrackerv2.AtrackerV2Options{}

			atrackerService, err = atrackerv2.NewAtrackerV2UsingExternalConfig(atrackerServiceOptions)
			Expect(err).To(BeNil())
			Expect(atrackerService).ToNot(BeNil())
			Expect(atrackerService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			atrackerService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateTarget - Create a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {
			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:                core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:                  core.StringPtr("my-atracker-bucket"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(true),
			}

			eventstreamsEndpointPrototypeModel := &atrackerv2.EventstreamsEndpointPrototype{
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Brokers:                 []string{"kafka-x:9094"},
				Topic:                   core.StringPtr("my-topic"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(false),
			}

			cloudLogsEndpointPrototypeModel := &atrackerv2.CloudLogsEndpointPrototype{
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
			}

			appconfigEndpointPrototypeModel := &atrackerv2.AppconfigEndpointPrototype{
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:apprapp:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
			}

			createTargetOptions := &atrackerv2.CreateTargetOptions{
				Name:                 core.StringPtr("my-cos-target"),
				TargetType:           core.StringPtr("cloud_object_storage"),
				CosEndpoint:          cosEndpointPrototypeModel,
				EventstreamsEndpoint: eventstreamsEndpointPrototypeModel,
				CloudlogsEndpoint:    cloudLogsEndpointPrototypeModel,
				AppconfigEndpoint:    appconfigEndpointPrototypeModel,
				Region:               core.StringPtr("us-south"),
				ManagedBy:            core.StringPtr("enterprise"),
			}

			target, response, err := atrackerService.CreateTarget(createTargetOptions)
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
			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"},
				Locations: []string{"us-south"},
			}

			createRouteOptions := &atrackerv2.CreateRouteOptions{
				Name:      core.StringPtr("my-route"),
				Rules:     []atrackerv2.RulePrototype{*rulePrototypeModel},
				ManagedBy: core.StringPtr("enterprise"),
			}

			route, response, err := atrackerService.CreateRoute(createRouteOptions)
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
			listTargetsOptions := &atrackerv2.ListTargetsOptions{
				Region: core.StringPtr("testString"),
			}

			targetList, response, err := atrackerService.ListTargets(listTargetsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetList).ToNot(BeNil())
		})
	})

	Describe(`GetTarget - Get details of a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {
			getTargetOptions := &atrackerv2.GetTargetOptions{
				ID: &targetIDLink,
			}

			target, response, err := atrackerService.GetTarget(getTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
	})

	Describe(`ReplaceTarget - Update a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions)`, func() {
			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:                core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:                  core.StringPtr("my-atracker-bucket"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(true),
			}

			eventstreamsEndpointPrototypeModel := &atrackerv2.EventstreamsEndpointPrototype{
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Brokers:                 []string{"kafka-x:9094"},
				Topic:                   core.StringPtr("my-topic"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(false),
			}

			cloudLogsEndpointPrototypeModel := &atrackerv2.CloudLogsEndpointPrototype{
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
			}

			appconfigEndpointPrototypeModel := &atrackerv2.AppconfigEndpointPrototype{
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:apprapp:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
			}

			replaceTargetOptions := &atrackerv2.ReplaceTargetOptions{
				ID:                   &targetIDLink,
				Name:                 core.StringPtr("my-cos-target"),
				CosEndpoint:          cosEndpointPrototypeModel,
				EventstreamsEndpoint: eventstreamsEndpointPrototypeModel,
				CloudlogsEndpoint:    cloudLogsEndpointPrototypeModel,
				AppconfigEndpoint:    appconfigEndpointPrototypeModel,
			}

			target, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
	})

	Describe(`ValidateTarget - Validate a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ValidateTarget(validateTargetOptions *ValidateTargetOptions)`, func() {
			validateTargetOptions := &atrackerv2.ValidateTargetOptions{
				ID: &targetIDLink,
			}

			target, response, err := atrackerService.ValidateTarget(validateTargetOptions)
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
			listRoutesOptions := &atrackerv2.ListRoutesOptions{}

			routeList, response, err := atrackerService.ListRoutes(listRoutesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeList).ToNot(BeNil())
		})
	})

	Describe(`GetRoute - Get details of a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRoute(getRouteOptions *GetRouteOptions)`, func() {
			getRouteOptions := &atrackerv2.GetRouteOptions{
				ID: &routeIDLink,
			}

			route, response, err := atrackerService.GetRoute(getRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
		})
	})

	Describe(`ReplaceRoute - Update a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions)`, func() {
			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"},
				Locations: []string{"us-south"},
			}

			replaceRouteOptions := &atrackerv2.ReplaceRouteOptions{
				ID:        &routeIDLink,
				Name:      core.StringPtr("my-route"),
				Rules:     []atrackerv2.RulePrototype{*rulePrototypeModel},
				ManagedBy: core.StringPtr("enterprise"),
			}

			route, response, err := atrackerService.ReplaceRoute(replaceRouteOptions)
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
			getSettingsOptions := &atrackerv2.GetSettingsOptions{}

			settings, response, err := atrackerService.GetSettings(getSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())
		})
	})

	Describe(`PutSettings - Modify settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutSettings(putSettingsOptions *PutSettingsOptions)`, func() {
			putSettingsOptions := &atrackerv2.PutSettingsOptions{
				MetadataRegionPrimary:  core.StringPtr("us-south"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
				DefaultTargets:         []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"},
				PermittedTargetRegions: []string{"us-south"},
				MetadataRegionBackup:   core.StringPtr("eu-de"),
			}

			settings, response, err := atrackerService.PutSettings(putSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(settings).ToNot(BeNil())
		})
	})

	Describe(`DeleteRoute - Delete a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRoute(deleteRouteOptions *DeleteRouteOptions)`, func() {
			deleteRouteOptions := &atrackerv2.DeleteRouteOptions{
				ID: &routeIDLink,
			}

			response, err := atrackerService.DeleteRoute(deleteRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTarget - Delete a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {
			deleteTargetOptions := &atrackerv2.DeleteTargetOptions{
				ID: &targetIDLink,
			}

			warningReport, response, err := atrackerService.DeleteTarget(deleteTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(warningReport).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
