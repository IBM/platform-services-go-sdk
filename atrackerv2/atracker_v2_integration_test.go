//go:build integration
// +build integration

/**
 * (C) Copyright IBM Corp. 2022.
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

	const notFoundTargetID = "deadbeef-2222-2222-2222-222222222222"

	const notFoundRouteID = "deadbeef-2222-2222-2222-222222222222"

	const badTargetType = "bad_target_type"

	var (
		err                          error
		atrackerService              *atrackerv2.AtrackerV2
		atrackerServiceNotAuthorized *atrackerv2.AtrackerV2
		serviceURL                   string
		config                       map[string]string
		refreshTokenNotAuthorized    string
	)

	// Globlal variables to hold link values
	var (
		routeIDLink   string
		targetIDLink  string
		targetIDLink2 string
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

			fmt.Printf("Service URL: %s\n", serviceURL)
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

			atrackerUnauthorizedServiceOptions := &atrackerv2.AtrackerV2Options{
				ServiceName: "NOT_AUTHORIZED",
			}
			atrackerServiceNotAuthorized, err = atrackerv2.NewAtrackerV2UsingExternalConfig(atrackerUnauthorizedServiceOptions)
			Expect(err).To(BeNil())
			Expect(atrackerServiceNotAuthorized).ToNot(BeNil())
			Expect(atrackerServiceNotAuthorized.Service.Options.URL).To(Equal(serviceURL))

			tokenNotAuthorized, err := atrackerServiceNotAuthorized.Service.Options.Authenticator.(*core.IamAuthenticator).RequestToken()
			Expect(err).To(BeNil())
			refreshTokenNotAuthorized = tokenNotAuthorized.RefreshToken
			Expect(refreshTokenNotAuthorized).ToNot(BeNil())

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

			createTargetOptions := &atrackerv2.CreateTargetOptions{
				Name:        core.StringPtr("my-cos-target"),
				TargetType:  core.StringPtr("cloud_object_storage"),
				CosEndpoint: cosEndpointPrototypeModel,
			}

			target, response, err := atrackerService.CreateTarget(createTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID
		})

		It(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {

			logdnaEndpointPrototypeModel := &atrackerv2.LogdnaEndpointPrototype{
				TargetCRN:    core.StringPtr("crn:v1:bluemix:public:logdna:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				IngestionKey: core.StringPtr("xxxxxxxxxxxxxx"),
			}

			createTargetOptions := &atrackerv2.CreateTargetOptions{
				Name:           core.StringPtr("my-logdna-target"),
				TargetType:     core.StringPtr("logdna"),
				LogdnaEndpoint: logdnaEndpointPrototypeModel,
			}

			target, response, err := atrackerService.CreateTarget(createTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink2 = *target.ID
		})

		It(`Returns 400 when backend input validation fails`, func() {
			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:                core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:                  core.StringPtr("my-atracker-bucket"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(false),
			}

			createTargetOptions := &atrackerv2.CreateTargetOptions{
				Name:        core.StringPtr("my-cos-target"),
				TargetType:  core.StringPtr(badTargetType),
				CosEndpoint: cosEndpointPrototypeModel,
			}

			_, response, err := atrackerService.CreateTarget(createTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {

			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			createTargetOptions := &atrackerv2.CreateTargetOptions{
				Name:        core.StringPtr("my-cos-target"),
				TargetType:  core.StringPtr("cloud_object_storage"),
				CosEndpoint: cosEndpointPrototypeModel,
			}

			_, response, err := atrackerServiceNotAuthorized.CreateTarget(createTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`CreateRoute - Create a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRoute(createRouteOptions *CreateRouteOptions)`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{targetIDLink},
				Locations: []string{"us-south"},
			}

			createRouteOptions := &atrackerv2.CreateRouteOptions{
				Name:  core.StringPtr("my-route"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := atrackerService.CreateRoute(createRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

			routeIDLink = *route.ID
		})

		It(`Returns 403 when user is not authorized`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{targetIDLink},
				Locations: []string{"us-south"},
			}

			createRouteOptions := &atrackerv2.CreateRouteOptions{
				Name:  core.StringPtr("my-route"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := atrackerServiceNotAuthorized.CreateRoute(createRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when input validation fails`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{notFoundTargetID},
				Locations: []string{"us-south"},
			}

			createRouteOptions := &atrackerv2.CreateRouteOptions{
				Name:  core.StringPtr("my-route"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := atrackerService.CreateRoute(createRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})
	})

	Describe(`ListTargets - List targets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTargets(listTargetsOptions *ListTargetsOptions)`, func() {

			listTargetsOptions := &atrackerv2.ListTargetsOptions{}

			targetList, response, err := atrackerService.ListTargets(listTargetsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetList).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized)`, func() {

			listTargetsOptions := &atrackerv2.ListTargetsOptions{}

			_, response, err := atrackerServiceNotAuthorized.ListTargets(listTargetsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
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

		It(`Returns 403 when user is not authorized`, func() {

			getTargetOptions := &atrackerv2.GetTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.GetTarget(getTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when target id is not found`, func() {

			getTargetOptions := &atrackerv2.GetTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			_, response, err := atrackerService.GetTarget(getTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
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

			logdnaEndpointPrototypeModel := &atrackerv2.LogdnaEndpointPrototype{
				TargetCRN:    core.StringPtr("crn:v1:bluemix:public:logdna:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				IngestionKey: core.StringPtr("xxxxxxxxxxxxxx"),
			}

			replaceTargetOptions := &atrackerv2.ReplaceTargetOptions{
				ID:             &targetIDLink,
				Name:           core.StringPtr("my-cos-target"),
				TargetType:     core.StringPtr("cloud_object_storage"),
				CosEndpoint:    cosEndpointPrototypeModel,
				LogdnaEndpoint: logdnaEndpointPrototypeModel,
				Region:         core.StringPtr("us-south"),
			}

			target, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})

		It(`Returns 400 when target fields are not valid`, func() {

			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:                core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:                  core.StringPtr("my-atracker-bucket"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(false),
			}

			replaceTargetOptions := &atrackerv2.ReplaceTargetOptions{
				ID:          core.StringPtr(targetIDLink),
				Name:        core.StringPtr("my-cos-target-modified"),
				TargetType:  core.StringPtr(badTargetType),
				CosEndpoint: cosEndpointPrototypeModel,
			}
			_, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))

		})

		It(`Returns 404 when target id is not found`, func() {

			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:                core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:                  core.StringPtr("my-atracker-bucket"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(false),
			}

			replaceTargetOptions := &atrackerv2.ReplaceTargetOptions{
				ID:          core.StringPtr(notFoundTargetID),
				Name:        core.StringPtr("my-cos-target-modified"),
				TargetType:  core.StringPtr("cloud_object_storage"),
				CosEndpoint: cosEndpointPrototypeModel,
			}
			_, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
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

		It(`Returns 403 when user is not authorized`, func() {

			validateTargetOptions := &atrackerv2.ValidateTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.ValidateTarget(validateTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when target id is not found`, func() {

			validateTargetOptions := &atrackerv2.ValidateTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			_, response, err := atrackerService.ValidateTarget(validateTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
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

		It(`Returns 403 when user is not authorized`, func() {

			listRoutesOptions := &atrackerv2.ListRoutesOptions{}

			_, response, err := atrackerServiceNotAuthorized.ListRoutes(listRoutesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
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

		It(`Returns 403 when user is not authorized`, func() {

			getRouteOptions := &atrackerv2.GetRouteOptions{
				ID: core.StringPtr(routeIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.GetRoute(getRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {

			getRouteOptions := &atrackerv2.GetRouteOptions{
				ID: core.StringPtr(notFoundRouteID),
			}

			_, response, err := atrackerService.GetRoute(getRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`ReplaceRoute - Update a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions)`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{targetIDLink},
				Locations: []string{"us-south"},
			}

			replaceRouteOptions := &atrackerv2.ReplaceRouteOptions{
				ID:    &routeIDLink,
				Name:  core.StringPtr("my-route"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := atrackerService.ReplaceRoute(replaceRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{targetIDLink},
				Locations: []string{"us-south"},
			}

			replaceRouteOptions := &atrackerv2.ReplaceRouteOptions{
				ID:    core.StringPtr(routeIDLink),
				Name:  core.StringPtr("my-route-modified"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := atrackerServiceNotAuthorized.ReplaceRoute(replaceRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{targetIDLink},
				Locations: []string{"us-south"},
			}

			replaceRouteOptions := &atrackerv2.ReplaceRouteOptions{
				ID:    core.StringPtr(notFoundRouteID),
				Name:  core.StringPtr("my-route-modified"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := atrackerService.ReplaceRoute(replaceRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
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

		It(`Returns 403 when user is not authorized`, func() {

			getSettingsOptions := &atrackerv2.GetSettingsOptions{}

			_, response, err := atrackerServiceNotAuthorized.GetSettings(getSettingsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`PutSettings - Modify settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutSettings(putSettingsOptions *PutSettingsOptions)`, func() {

			putSettingsOptions := &atrackerv2.PutSettingsOptions{
				DefaultTargets:         []string{targetIDLink},
				PermittedTargetRegions: []string{"us-south"},
				MetadataRegionPrimary:  core.StringPtr("us-south"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
			}

			settings, response, err := atrackerService.PutSettings(putSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(settings).ToNot(BeNil())
		})

		It(`Removing default targets`, func() {

			putSettingsOptions := &atrackerv2.PutSettingsOptions{
				DefaultTargets:         []string{},
				PermittedTargetRegions: []string{"us-south"},
				MetadataRegionPrimary:  core.StringPtr("us-south"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
			}

			settings, response, err := atrackerService.PutSettings(putSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(settings).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			putSettingsOptions := &atrackerv2.PutSettingsOptions{
				DefaultTargets:         []string{targetIDLink2},
				PermittedTargetRegions: []string{"us-south"},
				MetadataRegionPrimary:  core.StringPtr("us-south"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
			}

			_, response, err := atrackerServiceNotAuthorized.PutSettings(putSettingsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`PostMigration - Migrate v1 atracker resources to v2 atracker resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostMigration(postMigrationOptions *PostMigrationOptions)`, func() {

			postMigrationOptions := &atrackerv2.PostMigrationOptions{}

			migration, response, err := atrackerService.PostMigration(postMigrationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(migration).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			postMigrationOptions := &atrackerv2.PostMigrationOptions{}

			_, response, err := atrackerServiceNotAuthorized.PostMigration(postMigrationOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`GetMigration - get migration status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMigration(getMigrationOptions *GetMigrationOptions)`, func() {

			getMigrationOptions := &atrackerv2.GetMigrationOptions{}

			migration, response, err := atrackerService.GetMigration(getMigrationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(migration).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			getMigrationOptions := &atrackerv2.GetMigrationOptions{}

			_, response, err := atrackerServiceNotAuthorized.GetMigration(getMigrationOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
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

		It(`Returns 403 when user is not authorized`, func() {

			deleteRouteOptions := &atrackerv2.DeleteRouteOptions{
				ID: core.StringPtr(routeIDLink),
			}

			response, err := atrackerServiceNotAuthorized.DeleteRoute(deleteRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {

			deleteRouteOptions := &atrackerv2.DeleteRouteOptions{
				ID: core.StringPtr(notFoundRouteID),
			}

			response, err := atrackerService.DeleteRoute(deleteRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`DeleteTarget - Delete a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			deleteTargetOptions := &atrackerv2.DeleteTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.DeleteTarget(deleteTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
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

		It(`Returns 404 when target id is not found`, func() {

			deleteTargetOptions := &atrackerv2.DeleteTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			_, response, err := atrackerService.DeleteTarget(deleteTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
