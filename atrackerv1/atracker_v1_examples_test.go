// +build examples

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
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

const externalConfigFile = "../atracker_v1.env"

var (
	atrackerService *atrackerv1.AtrackerV1
	config          map[string]string
	configLoaded    bool = false
)

// Globlal variables to hold link values
var (
	routeIDLink  string
	targetIDLink string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`AtrackerV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(atrackerv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			atrackerServiceOptions := &atrackerv1.AtrackerV1Options{}

			atrackerService, err = atrackerv1.NewAtrackerV1UsingExternalConfig(atrackerServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(atrackerService).ToNot(BeNil())
		})
	})

	Describe(`AtrackerV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget request example`, func() {
			// begin-create_target

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			createTargetOptions := atrackerService.NewCreateTargetOptions(
				"my-cos-target",
				"cloud_object_storage",
				cosEndpointModel,
			)

			target, response, err := atrackerService.CreateTarget(createTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Printf("\nCreateTarget() result:\n%s\n", string(b))

			// end-create_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID

		})
		It(`ListTargets request example`, func() {
			// begin-list_targets

			listTargetsOptions := atrackerService.NewListTargetsOptions()

			targetList, response, err := atrackerService.ListTargets(listTargetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(targetList, "", "  ")
			fmt.Printf("\nListTargets() result:\n%s\n", string(b))

			// end-list_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetList).ToNot(BeNil())

		})
		It(`GetTarget request example`, func() {
			// begin-get_target

			getTargetOptions := atrackerService.NewGetTargetOptions(
				targetIDLink,
			)

			target, response, err := atrackerService.GetTarget(getTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Printf("\nGetTarget() result:\n%s\n", string(b))

			// end-get_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})
		It(`ReplaceTarget request example`, func() {
			// begin-replace_target

			cosEndpointModel := &atrackerv1.CosEndpoint{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			replaceTargetOptions := atrackerService.NewReplaceTargetOptions(
				targetIDLink,
				"my-cos-target-modified",
				"cloud_object_storage",
				cosEndpointModel,
			)

			target, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Printf("\nReplaceTarget() result:\n%s\n", string(b))

			// end-replace_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})
		It(`CreateRoute request example`, func() {
			// begin-create_route

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{targetIDLink},
			}

			createRouteOptions := atrackerService.NewCreateRouteOptions(
				"my-route",
				false,
				[]atrackerv1.Rule{*ruleModel},
			)

			route, response, err := atrackerService.CreateRoute(createRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Printf("\nCreateRoute() result:\n%s\n", string(b))

			// end-create_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

			routeIDLink = *route.ID

		})
		It(`ListRoutes request example`, func() {
			// begin-list_routes

			listRoutesOptions := atrackerService.NewListRoutesOptions()

			routeList, response, err := atrackerService.ListRoutes(listRoutesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routeList, "", "  ")
			fmt.Printf("\nListRoutes() result:\n%s\n", string(b))

			// end-list_routes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeList).ToNot(BeNil())

		})
		It(`GetRoute request example`, func() {
			// begin-get_route

			getRouteOptions := atrackerService.NewGetRouteOptions(
				routeIDLink,
			)

			route, response, err := atrackerService.GetRoute(getRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Printf("\nGetRoute() result:\n%s\n", string(b))

			// end-get_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`ReplaceRoute request example`, func() {
			// begin-replace_route

			ruleModel := &atrackerv1.Rule{
				TargetIds: []string{targetIDLink},
			}

			replaceRouteOptions := atrackerService.NewReplaceRouteOptions(
				routeIDLink,
				"my-route-modified",
				false,
				[]atrackerv1.Rule{*ruleModel},
			)

			route, response, err := atrackerService.ReplaceRoute(replaceRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Printf("\nReplaceRoute() result:\n%s\n", string(b))

			// end-replace_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`DeleteRoute request example`, func() {
			// begin-delete_route

			deleteRouteOptions := atrackerService.NewDeleteRouteOptions(
				routeIDLink,
			)

			response, err := atrackerService.DeleteRoute(deleteRouteOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteRoute() response status code: %d\n", response.StatusCode)
			// end-delete_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTarget request example`, func() {
			// begin-delete_target

			deleteTargetOptions := atrackerService.NewDeleteTargetOptions(
				targetIDLink,
			)

			response, err := atrackerService.DeleteTarget(deleteTargetOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteTarget() result status code: %d\n", response.StatusCode)
			// end-delete_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
