// +build integration

/**
 * (C) Copyright IBM Corp. 2021.
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

package atrackerv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the atrackerv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`AtrackerV1 Integration Tests`, func() {

	const externalConfigFile = "../atracker_v1.env"

	const notFoundTargetID = "deadbeef-2222-2222-2222-222222222222"

	const notFoundRouteID = "deadbeef-2222-2222-2222-222222222222"

	const badTargetType = "bad_target_type"

	var (
		err                          error
		atrackerService              *atrackerv1.AtrackerV1
		atrackerServiceNotAuthorized *atrackerv1.AtrackerV1
		serviceURL                   string
		config                       map[string]string

		refreshTokenNotAuthorized string
	)

	// Globlal variables to hold link values
	var (
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
			config, err = core.GetServiceProperties(atrackerv1.DefaultServiceName)
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

			atrackerServiceOptions := &atrackerv1.AtrackerV1Options{}

			atrackerService, err = atrackerv1.NewAtrackerV1UsingExternalConfig(atrackerServiceOptions)

			Expect(err).To(BeNil())
			Expect(atrackerService).ToNot(BeNil())
			Expect(atrackerService.Service.Options.URL).To(Equal(serviceURL))

			atrackerUnauthorizedServiceOptions := &atrackerv1.AtrackerV1Options{
				ServiceName: "NOT_AUTHORIZED",
			}
			atrackerServiceNotAuthorized, err = atrackerv1.NewAtrackerV1UsingExternalConfig(atrackerUnauthorizedServiceOptions)
			Expect(err).To(BeNil())
			Expect(atrackerServiceNotAuthorized).ToNot(BeNil())
			Expect(atrackerServiceNotAuthorized.Service.Options.URL).To(Equal(serviceURL))

			tokenNotAuthorized, err := atrackerServiceNotAuthorized.Service.Options.Authenticator.(*core.IamAuthenticator).RequestToken()
			Expect(err).To(BeNil())
			refreshTokenNotAuthorized = tokenNotAuthorized.RefreshToken
			Expect(refreshTokenNotAuthorized).ToNot(BeNil())
		})
	})

	Describe(`CreateTarget - Create a Cloud Object Storage target for a region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 400 when backend input validation fails`, func() {
			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			createTargetOptions := &atrackerv1.CreateTargetOptions{
				Name:        core.StringPtr("my-cos-target"),
				TargetType:  core.StringPtr(badTargetType),
				CosEndpoint: cosEndpointModel,
			}

			_, response, err := atrackerService.CreateTarget(createTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			createTargetOptions := &atrackerv1.CreateTargetOptions{
				Name:        core.StringPtr("my-cos-target"),
				TargetType:  core.StringPtr("cloud_object_storage"),
				CosEndpoint: cosEndpointModel,
			}

			_, response, err := atrackerServiceNotAuthorized.CreateTarget(createTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			createTargetOptions := &atrackerv1.CreateTargetOptions{
				Name:        core.StringPtr("my-cos-target"),
				TargetType:  core.StringPtr("cloud_object_storage"),
				CosEndpoint: cosEndpointModel,
			}

			target, response, err := atrackerService.CreateTarget(createTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID

		})
	})

	Describe(`ListTargets - List Cloud Object Storage targets for the region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			listTargetsOptions := &atrackerv1.ListTargetsOptions{}

			_, response, err := atrackerServiceNotAuthorized.ListTargets(listTargetsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`ListTargets(listTargetsOptions *ListTargetsOptions)`, func() {

			listTargetsOptions := &atrackerv1.ListTargetsOptions{}

			targetList, response, err := atrackerService.ListTargets(listTargetsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetList).ToNot(BeNil())
			Expect(len(targetList.Targets)).To(BeNumerically(">", 0))

		})
	})

	Describe(`GetTarget - Retrieve a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			getTargetOptions := &atrackerv1.GetTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.GetTarget(getTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when target id is not found`, func() {

			getTargetOptions := &atrackerv1.GetTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			_, response, err := atrackerService.GetTarget(getTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))

		})

		It(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {

			getTargetOptions := &atrackerv1.GetTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			target, response, err := atrackerService.GetTarget(getTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
			Expect(*target.Name).To(Equal("my-cos-target"))

		})
	})

	Describe(`ReplaceTarget - Update a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			replaceTargetOptions := &atrackerv1.ReplaceTargetOptions{
				ID:          core.StringPtr(targetIDLink),
				Name:        core.StringPtr("my-cos-target-modified"),
				TargetType:  core.StringPtr("cloud_object_storage"),
				CosEndpoint: cosEndpointModel,
			}

			_, response, err := atrackerServiceNotAuthorized.ReplaceTarget(replaceTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when target fields are not valid`, func() {

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			replaceTargetOptions := &atrackerv1.ReplaceTargetOptions{
				ID:          core.StringPtr(targetIDLink),
				Name:        core.StringPtr("my-cos-target-modified"),
				TargetType:  core.StringPtr(badTargetType),
				CosEndpoint: cosEndpointModel,
			}
			_, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))

		})

		It(`Returns 404 when target id is not found`, func() {

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			replaceTargetOptions := &atrackerv1.ReplaceTargetOptions{
				ID:          core.StringPtr(notFoundTargetID),
				Name:        core.StringPtr("my-cos-target-modified"),
				TargetType:  core.StringPtr("cloud_object_storage"),
				CosEndpoint: cosEndpointModel,
			}
			_, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))

		})

		It(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions)`, func() {

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			replaceTargetOptions := &atrackerv1.ReplaceTargetOptions{
				ID:          core.StringPtr(targetIDLink),
				Name:        core.StringPtr("my-cos-target-modified"),
				TargetType:  core.StringPtr("cloud_object_storage"),
				CosEndpoint: cosEndpointModel,
			}

			target, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
			Expect(*target.Name).To(Equal("my-cos-target-modified"))

		})
	})

	Describe(`ValidateTarget - Update a target with cos validation results`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			validateTargetOptions := &atrackerv1.ValidateTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.ValidateTarget(validateTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when target id is not found`, func() {

			validateTargetOptions := &atrackerv1.ValidateTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			_, response, err := atrackerService.ValidateTarget(validateTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))

		})

		It(`ValidateTarget(validateTargetOptions *ValidateTargetOptions)`, func() {

			validateTargetOptions := &atrackerv1.ValidateTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			target, response, err := atrackerService.ValidateTarget(validateTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})

	})

	Describe(`CreateRoute - Create a Route for the region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{targetIDLink},
			}

			createRouteOptions := &atrackerv1.CreateRouteOptions{
				Name:                core.StringPtr("my-route"),
				ReceiveGlobalEvents: core.BoolPtr(false),
				Rules:               []atrackerv1.Rule{*ruleModel},
			}

			_, response, err := atrackerServiceNotAuthorized.CreateRoute(createRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when input validation fails`, func() {

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{notFoundTargetID},
			}

			createRouteOptions := &atrackerv1.CreateRouteOptions{
				Name:                core.StringPtr("my-route"),
				ReceiveGlobalEvents: core.BoolPtr(false),
				Rules:               []atrackerv1.Rule{*ruleModel},
			}

			_, response, err := atrackerService.CreateRoute(createRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))

		})

		It(`CreateRoute(createRouteOptions *CreateRouteOptions)`, func() {

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{targetIDLink},
			}

			createRouteOptions := &atrackerv1.CreateRouteOptions{
				Name:                core.StringPtr("my-route"),
				ReceiveGlobalEvents: core.BoolPtr(false),
				Rules:               []atrackerv1.Rule{*ruleModel},
			}

			route, response, err := atrackerService.CreateRoute(createRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

			routeIDLink = *route.ID

		})
	})

	Describe(`ListRoutes - List routes for the region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			listRoutesOptions := &atrackerv1.ListRoutesOptions{}

			_, response, err := atrackerServiceNotAuthorized.ListRoutes(listRoutesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`ListRoutes(listRoutesOptions *ListRoutesOptions)`, func() {

			listRoutesOptions := &atrackerv1.ListRoutesOptions{}

			routeList, response, err := atrackerService.ListRoutes(listRoutesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeList).ToNot(BeNil())
			Expect(len(routeList.Routes)).To(BeNumerically(">", 0))

		})
	})

	Describe(`GetRoute - Retrieve a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			getRouteOptions := &atrackerv1.GetRouteOptions{
				ID: core.StringPtr(routeIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.GetRoute(getRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {

			getRouteOptions := &atrackerv1.GetRouteOptions{
				ID: core.StringPtr(notFoundRouteID),
			}

			_, response, err := atrackerService.GetRoute(getRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))

		})

		It(`GetRoute(getRouteOptions *GetRouteOptions)`, func() {

			getRouteOptions := &atrackerv1.GetRouteOptions{
				ID: core.StringPtr(routeIDLink),
			}

			route, response, err := atrackerService.GetRoute(getRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
			Expect(*route.Name).To(Equal("my-route"))

		})
	})

	Describe(`ReplaceRoute - Replace a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{targetIDLink},
			}

			replaceRouteOptions := &atrackerv1.ReplaceRouteOptions{
				ID:                  core.StringPtr(routeIDLink),
				Name:                core.StringPtr("my-route-modified"),
				ReceiveGlobalEvents: core.BoolPtr(false),
				Rules:               []atrackerv1.Rule{*ruleModel},
			}

			_, response, err := atrackerServiceNotAuthorized.ReplaceRoute(replaceRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{targetIDLink},
			}

			replaceRouteOptions := &atrackerv1.ReplaceRouteOptions{
				ID:                  core.StringPtr(notFoundRouteID),
				Name:                core.StringPtr("my-route-modified"),
				ReceiveGlobalEvents: core.BoolPtr(false),
				Rules:               []atrackerv1.Rule{*ruleModel},
			}

			_, response, err := atrackerService.ReplaceRoute(replaceRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))

		})

		It(`ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions)`, func() {

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{targetIDLink},
			}

			replaceRouteOptions := &atrackerv1.ReplaceRouteOptions{
				ID:                  core.StringPtr(routeIDLink),
				Name:                core.StringPtr("my-route-modified"),
				ReceiveGlobalEvents: core.BoolPtr(false),
				Rules:               []atrackerv1.Rule{*ruleModel},
			}

			route, response, err := atrackerService.ReplaceRoute(replaceRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
			Expect(*route.Name).To(Equal("my-route-modified"))

		})
	})

	Describe(`DeleteRoute - Delete a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			deleteRouteOptions := &atrackerv1.DeleteRouteOptions{
				ID: core.StringPtr(routeIDLink),
			}

			response, err := atrackerServiceNotAuthorized.DeleteRoute(deleteRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {

			deleteRouteOptions := &atrackerv1.DeleteRouteOptions{
				ID: core.StringPtr(notFoundRouteID),
			}

			response, err := atrackerService.DeleteRoute(deleteRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))

		})

		It(`DeleteRoute(deleteRouteOptions *DeleteRouteOptions)`, func() {

			deleteRouteOptions := &atrackerv1.DeleteRouteOptions{
				ID: core.StringPtr(routeIDLink),
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

		It(`Returns 403 when user is not authorized`, func() {

			deleteTargetOptions := &atrackerv1.DeleteTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.DeleteTarget(deleteTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when target id is not found`, func() {

			deleteTargetOptions := &atrackerv1.DeleteTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			_, response, err := atrackerService.DeleteTarget(deleteTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))

		})

		It(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {

			deleteTargetOptions := &atrackerv1.DeleteTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			warningReport, response, err := atrackerService.DeleteTarget(deleteTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(warningReport).ToNot(BeNil())

		})
	})

	Describe(`GetEndpoints - Get endpoints`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			getEndpointsOptions := &atrackerv1.GetEndpointsOptions{}

			_, response, err := atrackerServiceNotAuthorized.GetEndpoints(getEndpointsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`GetEndpoints(getEndpointsOptions *GetEndpointsOptions)`, func() {

			getEndpointsOptions := &atrackerv1.GetEndpointsOptions{}

			endpoints, response, err := atrackerService.GetEndpoints(getEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoints).ToNot(BeNil())
			// more validations
			Expect(endpoints.APIEndpoint.PublicURL).ToNot(BeNil())
			Expect(endpoints.APIEndpoint.PrivateURL).ToNot(BeNil())
			Expect(*endpoints.APIEndpoint.PrivateEnabled).To(Equal(true))
		})
	})

	Describe(`PatchEndpoints - Modify endpoints`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			endpointsRequestAPIEndpointModel := &atrackerv1.EndpointsRequestAPIEndpoint{
				PublicEnabled: core.BoolPtr(true),
			}

			patchEndpointsOptions := &atrackerv1.PatchEndpointsOptions{
				APIEndpoint: endpointsRequestAPIEndpointModel,
			}

			_, response, err := atrackerServiceNotAuthorized.PatchEndpoints(patchEndpointsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`PatchEndpoints(patchEndpointsOptions *PatchEndpointsOptions)`, func() {

			endpointsRequestAPIEndpointModel := &atrackerv1.EndpointsRequestAPIEndpoint{
				PublicEnabled: core.BoolPtr(true),
			}

			patchEndpointsOptions := &atrackerv1.PatchEndpointsOptions{
				APIEndpoint: endpointsRequestAPIEndpointModel,
			}

			endpoints, response, err := atrackerService.PatchEndpoints(patchEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoints).ToNot(BeNil())
			Expect(endpoints.APIEndpoint).ToNot(BeNil())
			// more validations
			Expect(endpoints.APIEndpoint.PublicURL).ToNot(BeNil())
			Expect(endpoints.APIEndpoint.PrivateURL).ToNot(BeNil())
			Expect(*endpoints.APIEndpoint.PublicEnabled).To(Equal(true))
			Expect(*endpoints.APIEndpoint.PrivateEnabled).To(Equal(true))
		})
	})

})

//
// Utility functions are declared in the unit test file
//
