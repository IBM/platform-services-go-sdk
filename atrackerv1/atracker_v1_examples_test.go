// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the atracker service.
//
// The following configuration properties are assumed to be defined:
// ATRACKER_URL=<service base url>
// ATRACKER_AUTH_TYPE=iam
// ATRACKER_APIKEY=<IAM apikey>
// ATRACKER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
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
			fmt.Println("\nCreateTarget() result:")
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
			fmt.Println(string(b))

			// end-create_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID

		})
		It(`ListTargets request example`, func() {
			fmt.Println("\nListTargets() result:")
			// begin-list_targets

			listTargetsOptions := atrackerService.NewListTargetsOptions()

			targetList, response, err := atrackerService.ListTargets(listTargetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(targetList, "", "  ")
			fmt.Println(string(b))

			// end-list_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetList).ToNot(BeNil())

		})
		It(`GetTarget request example`, func() {
			fmt.Println("\nGetTarget() result:")
			// begin-get_target

			getTargetOptions := atrackerService.NewGetTargetOptions(
				targetIDLink,
			)

			target, response, err := atrackerService.GetTarget(getTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-get_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})
		It(`ReplaceTarget request example`, func() {
			fmt.Println("\nReplaceTarget() result:")
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
			fmt.Println(string(b))

			// end-replace_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})
		It(`ValidateTarget request example`, func() {
			fmt.Println("\nValidateTarget() result:")
			// begin-validate_target

			validateTargetOptions := atrackerService.NewValidateTargetOptions(
				targetIDLink,
			)

			target, response, err := atrackerService.ValidateTarget(validateTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-validate_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})
		It(`CreateRoute request example`, func() {
			fmt.Println("\nCreateRoute() result:")
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
			fmt.Println(string(b))

			// end-create_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

			routeIDLink = *route.ID

		})
		It(`ListRoutes request example`, func() {
			fmt.Println("\nListRoutes() result:")
			// begin-list_routes

			listRoutesOptions := atrackerService.NewListRoutesOptions()

			routeList, response, err := atrackerService.ListRoutes(listRoutesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routeList, "", "  ")
			fmt.Println(string(b))

			// end-list_routes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeList).ToNot(BeNil())

		})
		It(`GetRoute request example`, func() {
			fmt.Println("\nGetRoute() result:")
			// begin-get_route

			getRouteOptions := atrackerService.NewGetRouteOptions(
				routeIDLink,
			)

			route, response, err := atrackerService.GetRoute(getRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-get_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`ReplaceRoute request example`, func() {
			fmt.Println("\nReplaceRoute() result:")
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
			fmt.Println(string(b))

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
			// end-delete_route
			fmt.Printf("\nDeleteRoute() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTarget request example`, func() {
			fmt.Println("\nDeleteTarget() result:")
			// begin-delete_target

			deleteTargetOptions := atrackerService.NewDeleteTargetOptions(
				targetIDLink,
			)

			warningReport, response, err := atrackerService.DeleteTarget(deleteTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(warningReport, "", "  ")
			fmt.Println(string(b))

			// end-delete_target
			fmt.Printf("\nDeleteTarget() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(warningReport).ToNot(BeNil())

		})
		It(`GetEndpoints request example`, func() {
			fmt.Println("\nGetEndpoints() result:")
			// begin-get_endpoints

			getEndpointsOptions := atrackerService.NewGetEndpointsOptions()

			endpoints, response, err := atrackerService.GetEndpoints(getEndpointsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(endpoints, "", "  ")
			fmt.Println(string(b))

			// end-get_endpoints

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoints).ToNot(BeNil())

		})
		It(`PatchEndpoints request example`, func() {
			fmt.Println("\nPatchEndpoints() result:")
			// begin-patch_endpoints

			patchEndpointsOptions := atrackerService.NewPatchEndpointsOptions()

			// set public enabled flag
			publicEnabled := true
			patchEndpointsOptions.SetAPIEndpoint(&atrackerv1.EndpointsRequestAPIEndpoint{PublicEnabled: &publicEnabled})

			endpoints, response, err := atrackerService.PatchEndpoints(patchEndpointsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(endpoints, "", "  ")
			fmt.Println(string(b))

			// end-patch_endpoints

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoints).ToNot(BeNil())

		})
	})
})
