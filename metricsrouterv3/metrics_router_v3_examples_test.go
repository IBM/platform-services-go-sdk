//go:build examples
// +build examples

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

package metricsrouterv3_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the metrics-router service.
//
// The following configuration properties are assumed to be defined:
// METRICS_ROUTER_URL=<service base url>
// METRICS_ROUTER_AUTH_TYPE=iam
// METRICS_ROUTER_APIKEY=<IAM apikey>
// METRICS_ROUTER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`MetricsRouterV3 Examples Tests`, func() {

	const externalConfigFile = "../metrics_router_v3.env"

	var (
		metricsRouterService *metricsrouterv3.MetricsRouterV3
		config               map[string]string

		// Variables to hold link values
		routeIDLink  string
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
			config, err = core.GetServiceProperties(metricsrouterv3.DefaultServiceName)
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

			metricsRouterServiceOptions := &metricsrouterv3.MetricsRouterV3Options{}

			metricsRouterService, err = metricsrouterv3.NewMetricsRouterV3UsingExternalConfig(metricsRouterServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(metricsRouterService).ToNot(BeNil())
		})
	})

	Describe(`MetricsRouterV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget request example`, func() {
			fmt.Println("\nCreateTarget() result:")
			// begin-create_target

			createTargetOptions := metricsRouterService.NewCreateTargetOptions(
				"my-mr-target",
				"crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::",
			)

			target, response, err := metricsRouterService.CreateTarget(createTargetOptions)
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

			inclusionFilterModel := &metricsrouterv3.InclusionFilter{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Value:    []string{"teststring"},
			}

			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				TargetIds:        []string{targetIDLink},
				InclusionFilters: []metricsrouterv3.InclusionFilter{*inclusionFilterModel},
			}

			createRouteOptions := metricsRouterService.NewCreateRouteOptions(
				"my-route",
				[]metricsrouterv3.RulePrototype{*rulePrototypeModel},
			)

			route, response, err := metricsRouterService.CreateRoute(createRouteOptions)
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

			listTargetsOptions := metricsRouterService.NewListTargetsOptions()

			targetList, response, err := metricsRouterService.ListTargets(listTargetsOptions)
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

			getTargetOptions := metricsRouterService.NewGetTargetOptions(
				targetIDLink,
			)

			target, response, err := metricsRouterService.GetTarget(getTargetOptions)
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

			replaceTargetOptions := metricsRouterService.NewReplaceTargetOptions(
				targetIDLink,
				"my-mr-target",
				"crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::",
			)

			target, response, err := metricsRouterService.ReplaceTarget(replaceTargetOptions)
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

			validateTargetOptions := metricsRouterService.NewValidateTargetOptions(
				targetIDLink,
			)

			target, response, err := metricsRouterService.ValidateTarget(validateTargetOptions)
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
		It(`ListRoutes request example`, func() {
			fmt.Println("\nListRoutes() result:")
			// begin-list_routes

			listRoutesOptions := metricsRouterService.NewListRoutesOptions()

			routeList, response, err := metricsRouterService.ListRoutes(listRoutesOptions)
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

			getRouteOptions := metricsRouterService.NewGetRouteOptions(
				routeIDLink,
			)

			route, response, err := metricsRouterService.GetRoute(getRouteOptions)
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

			inclusionFilterModel := &metricsrouterv3.InclusionFilter{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Value:    core.StringPtr("teststring"),
			}

			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				TargetIds:        []string{targetIDLink},
				InclusionFilters: []metricsrouterv3.InclusionFilter{*inclusionFilterModel},
			}

			replaceRouteOptions := metricsRouterService.NewReplaceRouteOptions(
				routeIDLink,
				"my-route",
				[]metricsrouterv3.RulePrototype{*rulePrototypeModel},
			)

			route, response, err := metricsRouterService.ReplaceRoute(replaceRouteOptions)
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
		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-get_settings

			getSettingsOptions := metricsRouterService.NewGetSettingsOptions()

			settings, response, err := metricsRouterService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-get_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())
		})
		It(`ReplaceSettings request example`, func() {
			fmt.Println("\nReplaceSettings() result:")
			// begin-replace_settings

			replaceSettingsOptions := metricsRouterService.NewReplaceSettingsOptions(
				"us-south",
				false,
			)

			settings, response, err := metricsRouterService.ReplaceSettings(replaceSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-replace_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(settings).ToNot(BeNil())
		})
		It(`DeleteRoute request example`, func() {
			// begin-delete_route

			deleteRouteOptions := metricsRouterService.NewDeleteRouteOptions(
				routeIDLink,
			)

			response, err := metricsRouterService.DeleteRoute(deleteRouteOptions)
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
			fmt.Println("\nDeleteTarget() result:")
			// begin-delete_target

			deleteTargetOptions := metricsRouterService.NewDeleteTargetOptions(
				targetIDLink,
			)

			warningReport, response, err := metricsRouterService.DeleteTarget(deleteTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(warningReport, "", "  ")
			fmt.Println(string(b))

			// end-delete_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
