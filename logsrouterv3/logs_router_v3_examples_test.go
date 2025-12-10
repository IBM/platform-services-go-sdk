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

package logsrouterv3_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the logs-router service.
//
// The following configuration properties are assumed to be defined:
// LOGS_ROUTER_URL=<service base url>
// LOGS_ROUTER_AUTH_TYPE=iam
// LOGS_ROUTER_APIKEY=<IAM apikey>
// LOGS_ROUTER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`LogsRouterV3 Examples Tests`, func() {

	const externalConfigFile = "../logs_router_v3.env"

	var (
		logsRouterService *logsrouterv3.LogsRouterV3
		config       map[string]string

		// Variables to hold link values
		routeIDLink string
		targetIDLink string
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
			config, err = core.GetServiceProperties(logsrouterv3.DefaultServiceName)
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

			logsRouterServiceOptions := &logsrouterv3.LogsRouterV3Options{}

			logsRouterService, err = logsrouterv3.NewLogsRouterV3UsingExternalConfig(logsRouterServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(logsRouterService).ToNot(BeNil())
		})
	})

	Describe(`LogsRouterV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget request example`, func() {
			fmt.Println("\nCreateTarget() result:")
			// begin-create_target

			createTargetOptions := logsRouterService.NewCreateTargetOptions(
				"my-lr-target",
				"crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::",
			)

			target, response, err := logsRouterService.CreateTarget(createTargetOptions)
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
			fmt.Fprintf(GinkgoWriter, "Saved targetIDLink value: %v\n", targetIDLink)
		})
		It(`CreateRoute request example`, func() {
			fmt.Println("\nCreateRoute() result:")
			// begin-create_route

			targetIdentityModel := &logsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}

			inclusionFilterPrototypeModel := &logsrouterv3.InclusionFilterPrototype{
				Operand: core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Values: []string{"us-south"},
			}

			rulePrototypeModel := &logsrouterv3.RulePrototype{
				Targets: []logsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}

			createRouteOptions := logsRouterService.NewCreateRouteOptions(
				"my-route",
				[]logsrouterv3.RulePrototype{*rulePrototypeModel},
			)

			route, response, err := logsRouterService.CreateRoute(createRouteOptions)
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
			fmt.Fprintf(GinkgoWriter, "Saved routeIDLink value: %v\n", routeIDLink)
		})
		It(`ListTargets request example`, func() {
			fmt.Println("\nListTargets() result:")
			// begin-list_targets

			listTargetsOptions := logsRouterService.NewListTargetsOptions()

			targetCollection, response, err := logsRouterService.ListTargets(listTargetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(targetCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetCollection).ToNot(BeNil())
		})
		It(`GetTarget request example`, func() {
			fmt.Println("\nGetTarget() result:")
			// begin-get_target

			getTargetOptions := logsRouterService.NewGetTargetOptions(
				targetIDLink,
			)

			target, response, err := logsRouterService.GetTarget(getTargetOptions)
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
		It(`UpdateTarget request example`, func() {
			fmt.Println("\nUpdateTarget() result:")
			// begin-update_target

			updateTargetOptions := logsRouterService.NewUpdateTargetOptions(
				targetIDLink,
			)

			target, response, err := logsRouterService.UpdateTarget(updateTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-update_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
		It(`ListRoutes request example`, func() {
			fmt.Println("\nListRoutes() result:")
			// begin-list_routes

			listRoutesOptions := logsRouterService.NewListRoutesOptions()

			routeCollection, response, err := logsRouterService.ListRoutes(listRoutesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routeCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_routes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeCollection).ToNot(BeNil())
		})
		It(`GetRoute request example`, func() {
			fmt.Println("\nGetRoute() result:")
			// begin-get_route

			getRouteOptions := logsRouterService.NewGetRouteOptions(
				routeIDLink,
			)

			route, response, err := logsRouterService.GetRoute(getRouteOptions)
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
		It(`UpdateRoute request example`, func() {
			fmt.Println("\nUpdateRoute() result:")
			// begin-update_route

			updateRouteOptions := logsRouterService.NewUpdateRouteOptions(
				routeIDLink,
			)

			route, response, err := logsRouterService.UpdateRoute(updateRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-update_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
		})
		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-get_settings

			getSettingsOptions := logsRouterService.NewGetSettingsOptions()

			setting, response, err := logsRouterService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(setting, "", "  ")
			fmt.Println(string(b))

			// end-get_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setting).ToNot(BeNil())
		})
		It(`UpdateSettings request example`, func() {
			fmt.Println("\nUpdateSettings() result:")
			// begin-update_settings

			updateSettingsOptions := logsRouterService.NewUpdateSettingsOptions()

			setting, response, err := logsRouterService.UpdateSettings(updateSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(setting, "", "  ")
			fmt.Println(string(b))

			// end-update_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setting).ToNot(BeNil())
		})
		It(`QueryDestinations request example`, func() {
			fmt.Println("\nQueryDestinations() result:")
			// begin-query_destinations

			crnPrototypeModel := &logsrouterv3.CRNPrototype{
				CRN: core.StringPtr("crn:v1:bluemix:public:codeengine:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::"),
			}

			queryDestinationsOptions := logsRouterService.NewQueryDestinationsOptions(
				[]logsrouterv3.CRNPrototype{*crnPrototypeModel},
			)

			destinationsQuery, response, err := logsRouterService.QueryDestinations(queryDestinationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(destinationsQuery, "", "  ")
			fmt.Println(string(b))

			// end-query_destinations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationsQuery).ToNot(BeNil())
		})
		It(`MigrateActions request example`, func() {
			fmt.Println("\nMigrateActions() result:")
			// begin-migrate_actions

			migrateActionsOptions := logsRouterService.NewMigrateActionsOptions(
				"complete",
			)

			migrationComplete, response, err := logsRouterService.MigrateActions(migrateActionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(migrationComplete, "", "  ")
			fmt.Println(string(b))

			// end-migrate_actions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(migrationComplete).ToNot(BeNil())
		})
		It(`DeleteRoute request example`, func() {
			// begin-delete_route

			deleteRouteOptions := logsRouterService.NewDeleteRouteOptions(
				routeIDLink,
			)

			response, err := logsRouterService.DeleteRoute(deleteRouteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteRoute(): %d\n", response.StatusCode)
			}

			// end-delete_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteTarget request example`, func() {
			// begin-delete_target

			deleteTargetOptions := logsRouterService.NewDeleteTargetOptions(
				targetIDLink,
			)

			response, err := logsRouterService.DeleteTarget(deleteTargetOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTarget(): %d\n", response.StatusCode)
			}

			// end-delete_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
