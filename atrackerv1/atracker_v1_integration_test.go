// +build integration

/**
 * (C) Copyright IBM Corp. 2020.
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

	"github.com/IBM/go-sdk-core/v4/core"
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

	var (
		err             error
		atrackerService *atrackerv1.AtrackerV1
		serviceURL      string
		config          map[string]string
	)

	var (
		targetID string
		routeID  string
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
		})
	})

	Describe(`CreateTarget - Create a Cloud Object Storage target for a region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				ApiKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			createTargetOptions := &atrackerv1.CreateTargetOptions{
				Name:        core.StringPtr("my-cos-target"),
				TargetType:  core.StringPtr(atrackerv1.CreateTargetOptionsTargetTypeCloudObjectStorageConst),
				CosEndpoint: cosEndpointModel,
			}

			target, response, err := atrackerService.CreateTarget(createTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())
			targetID = *target.ID
		})
	})

	Describe(`ListTargets - List Cloud Object Storage targets for the region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
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
		It(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {

			getTargetOptions := &atrackerv1.GetTargetOptions{
				ID: core.StringPtr(targetID),
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
		It(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions)`, func() {

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				ApiKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			replaceTargetOptions := &atrackerv1.ReplaceTargetOptions{
				ID:          core.StringPtr(targetID),
				Name:        core.StringPtr("my-cos-target-modified"),
				TargetType:  core.StringPtr("cos"),
				CosEndpoint: cosEndpointModel,
			}

			target, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
			Expect(*target.Name).To(Equal("my-cos-target-modified"))
		})
	})

	Describe(`CreateRoute - Create a Route for the region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRoute(createRouteOptions *CreateRouteOptions)`, func() {

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{targetID},
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
			routeID = *route.ID
		})
	})

	Describe(`ListRoutes - List routes for the region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
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
		It(`GetRoute(getRouteOptions *GetRouteOptions)`, func() {

			getRouteOptions := &atrackerv1.GetRouteOptions{
				ID: core.StringPtr(routeID),
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
		It(`ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions)`, func() {

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{targetID},
			}

			replaceRouteOptions := &atrackerv1.ReplaceRouteOptions{
				ID:                  core.StringPtr(routeID),
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

	// delete route first
	Describe(`DeleteRoute - Delete a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRoute(deleteRouteOptions *DeleteRouteOptions)`, func() {

			deleteRouteOptions := &atrackerv1.DeleteRouteOptions{
				ID: core.StringPtr(routeID),
			}

			response, err := atrackerService.DeleteRoute(deleteRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	// then delete target
	Describe(`DeleteTarget - Delete a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {

			deleteTargetOptions := &atrackerv1.DeleteTargetOptions{
				ID: core.StringPtr(targetID),
			}

			response, err := atrackerService.DeleteTarget(deleteTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
