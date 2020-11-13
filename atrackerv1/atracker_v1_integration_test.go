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
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
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
		err          error
		atrackerService *atrackerv1.AtrackerV1
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
				Endpoint: core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCrn: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket: core.StringPtr("my-atracker-bucket"),
				ApiKey: core.StringPtr("xxxxxxxxxxxxxx"),
			}

			createTargetOptions := &atrackerv1.CreateTargetOptions{
				Name: core.StringPtr("my-cos-target"),
				TargetType: core.StringPtr("cos"),
				CosEndpoint: cosEndpointModel,
			}

			target, response, err := atrackerService.CreateTarget(createTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

		})
	})

	Describe(`ListTargets - List Cloud Object Storage targets for the region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTargets(listTargetsOptions *ListTargetsOptions)`, func() {

			listTargetsOptions := &atrackerv1.ListTargetsOptions{
			}

			targetList, response, err := atrackerService.ListTargets(listTargetsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetList).ToNot(BeNil())

		})
	})

	Describe(`GetTarget - Retrieve a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {

			getTargetOptions := &atrackerv1.GetTargetOptions{
				ID: core.StringPtr("testString"),
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

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint: core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCrn: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket: core.StringPtr("my-atracker-bucket"),
				ApiKey: core.StringPtr("xxxxxxxxxxxxxx"),
			}

			replaceTargetOptions := &atrackerv1.ReplaceTargetOptions{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("my-cos-target"),
				TargetType: core.StringPtr("cos"),
				CosEndpoint: cosEndpointModel,
			}

			target, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})
	})

	Describe(`CreateRoute - Create a Route for the region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRoute(createRouteOptions *CreateRouteOptions)`, func() {

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"},
			}

			createRouteOptions := &atrackerv1.CreateRouteOptions{
				Name: core.StringPtr("my-route"),
				ReceiveGlobalEvents: core.BoolPtr(false),
				Rules: []atrackerv1.Rule{*ruleModel},
			}

			route, response, err := atrackerService.CreateRoute(createRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

		})
	})

	Describe(`ListRoutes - List routes for the region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRoutes(listRoutesOptions *ListRoutesOptions)`, func() {

			listRoutesOptions := &atrackerv1.ListRoutesOptions{
			}

			routeList, response, err := atrackerService.ListRoutes(listRoutesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeList).ToNot(BeNil())

		})
	})

	Describe(`GetRoute - Retrieve a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRoute(getRouteOptions *GetRouteOptions)`, func() {

			getRouteOptions := &atrackerv1.GetRouteOptions{
				ID: core.StringPtr("testString"),
			}

			route, response, err := atrackerService.GetRoute(getRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
	})

	Describe(`ReplaceRoute - Replace a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions)`, func() {

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"},
			}

			replaceRouteOptions := &atrackerv1.ReplaceRouteOptions{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("my-route"),
				ReceiveGlobalEvents: core.BoolPtr(false),
				Rules: []atrackerv1.Rule{*ruleModel},
			}

			route, response, err := atrackerService.ReplaceRoute(replaceRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
	})

	Describe(`DeleteTarget - Delete a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {

			deleteTargetOptions := &atrackerv1.DeleteTargetOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := atrackerService.DeleteTarget(deleteTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteRoute - Delete a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRoute(deleteRouteOptions *DeleteRouteOptions)`, func() {

			deleteRouteOptions := &atrackerv1.DeleteRouteOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := atrackerService.DeleteRoute(deleteRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
