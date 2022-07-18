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

package resourcecontrollerv2_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the resource_controller service.
//
// The following configuration properties are assumed to be defined:
// RESOURCE_CONTROLLER_URL=<service base url>
// RESOURCE_CONTROLLER_AUTH_TYPE=iam
// RESOURCE_CONTROLLER_APIKEY=<IAM apikey>
// RESOURCE_CONTROLLER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`ResourceControllerV2 Examples Tests`, func() {

	const externalConfigFile = "../resource_controller_v2.env"

	var (
		resourceControllerService *resourcecontrollerv2.ResourceControllerV2
		config       map[string]string

		// Variables to hold link values
		resourceAliasIDLink string
		resourceBindingIDLink string
		resourceInstanceIDLink string
		resourceKeyIDLink string
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
			config, err = core.GetServiceProperties(resourcecontrollerv2.DefaultServiceName)
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

			resourceControllerServiceOptions := &resourcecontrollerv2.ResourceControllerV2Options{}

			resourceControllerService, err = resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(resourceControllerServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(resourceControllerService).ToNot(BeNil())
		})
	})

	Describe(`ResourceControllerV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateResourceAlias request example`, func() {
			fmt.Println("\nCreateResourceAlias() result:")
			// begin-create_resource_alias

			createResourceAliasOptions := resourceControllerService.NewCreateResourceAliasOptions(
				"ExampleResourceAlias",
				"381fd51a-f251-4f95-aff4-2b03fa8caa63",
				"crn:v1:bluemix:public:bluemix:us-south:o/d35d4f0e-5076-4c89-9361-2522894b6548::cf-space:e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7",
			)

			resourceAlias, response, err := resourceControllerService.CreateResourceAlias(createResourceAliasOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceAlias, "", "  ")
			fmt.Println(string(b))

			// end-create_resource_alias

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resourceAlias).ToNot(BeNil())

			resourceAliasIDLink = *resourceAlias.GUID
			fmt.Fprintf(GinkgoWriter, "Saved resourceAliasIDLink value: %v\n", resourceAliasIDLink)
		})
		It(`CreateResourceBinding request example`, func() {
			fmt.Println("\nCreateResourceBinding() result:")
			// begin-create_resource_binding

			resourceBindingPostParametersModel := &resourcecontrollerv2.ResourceBindingPostParameters{
			}
			resourceBindingPostParametersModel.SetProperty("exampleParameter", core.StringPtr("exampleValue"))

			createResourceBindingOptions := resourceControllerService.NewCreateResourceBindingOptions(
				"faaec9d8-ec64-44d8-ab83-868632fac6a2",
				"crn:v1:staging:public:bluemix:us-south:s/e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7::cf-application:8d9457e0-1303-4f32-b4b3-5525575f6205",
			)
			createResourceBindingOptions.SetName("ExampleResourceBinding")
			createResourceBindingOptions.SetParameters(resourceBindingPostParametersModel)

			resourceBinding, response, err := resourceControllerService.CreateResourceBinding(createResourceBindingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceBinding, "", "  ")
			fmt.Println(string(b))

			// end-create_resource_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resourceBinding).ToNot(BeNil())

			resourceBindingIDLink = *resourceBinding.GUID
			fmt.Fprintf(GinkgoWriter, "Saved resourceBindingIDLink value: %v\n", resourceBindingIDLink)
		})
		It(`CreateResourceInstance request example`, func() {
			fmt.Println("\nCreateResourceInstance() result:")
			// begin-create_resource_instance

			createResourceInstanceOptions := resourceControllerService.NewCreateResourceInstanceOptions(
				"ExampleResourceInstance",
				"global",
				"13aa3ee48c3b44ddb64c05c79f7ab8ef",
				"a10e4960-3685-11e9-b210-d663bd873d93",
			)

			resourceInstance, response, err := resourceControllerService.CreateResourceInstance(createResourceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceInstance, "", "  ")
			fmt.Println(string(b))

			// end-create_resource_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resourceInstance).ToNot(BeNil())

			resourceInstanceIDLink = *resourceInstance.GUID
			fmt.Fprintf(GinkgoWriter, "Saved resourceInstanceIDLink value: %v\n", resourceInstanceIDLink)
		})
		It(`CreateResourceKey request example`, func() {
			fmt.Println("\nCreateResourceKey() result:")
			// begin-create_resource_key

			resourceKeyPostParametersModel := &resourcecontrollerv2.ResourceKeyPostParameters{
			}
			resourceKeyPostParametersModel.SetProperty("exampleParameter", core.StringPtr("exampleValue"))

			createResourceKeyOptions := resourceControllerService.NewCreateResourceKeyOptions(
				"ExampleResourceKey",
				resourceInstanceIDLink,
			)
			createResourceKeyOptions.SetParameters(resourceKeyPostParametersModel)

			resourceKey, response, err := resourceControllerService.CreateResourceKey(createResourceKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceKey, "", "  ")
			fmt.Println(string(b))

			// end-create_resource_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resourceKey).ToNot(BeNil())

			resourceKeyIDLink = *resourceKey.GUID
			fmt.Fprintf(GinkgoWriter, "Saved resourceKeyIDLink value: %v\n", resourceKeyIDLink)
		})
		It(`ListResourceInstances request example`, func() {
			fmt.Println("\nListResourceInstances() result:")
			// begin-list_resource_instances

			listResourceInstancesOptions := resourceControllerService.NewListResourceInstancesOptions()
			listResourceInstancesOptions.SetUpdatedFrom("2021-01-01")
			listResourceInstancesOptions.SetUpdatedTo("2021-01-01")

			resourceInstancesList, response, err := resourceControllerService.ListResourceInstances(listResourceInstancesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceInstancesList, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_instances

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstancesList).ToNot(BeNil())

		})
		It(`GetResourceInstance request example`, func() {
			fmt.Println("\nGetResourceInstance() result:")
			// begin-get_resource_instance

			getResourceInstanceOptions := resourceControllerService.NewGetResourceInstanceOptions(
				resourceInstanceIDLink,
			)

			resourceInstance, response, err := resourceControllerService.GetResourceInstance(getResourceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceInstance, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstance).ToNot(BeNil())

		})
		It(`UpdateResourceInstance request example`, func() {
			fmt.Println("\nUpdateResourceInstance() result:")
			// begin-update_resource_instance

			updateResourceInstanceOptions := resourceControllerService.NewUpdateResourceInstanceOptions(
				resourceInstanceIDLink,
			)
			updateResourceInstanceOptions.SetName("UpdatedExampleResourceInstance")

			resourceInstance, response, err := resourceControllerService.UpdateResourceInstance(updateResourceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceInstance, "", "  ")
			fmt.Println(string(b))

			// end-update_resource_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstance).ToNot(BeNil())

		})
		It(`ListResourceAliasesForInstance request example`, func() {
			fmt.Println("\nListResourceAliasesForInstance() result:")
			// begin-list_resource_aliases_for_instance

			listResourceAliasesForInstanceOptions := resourceControllerService.NewListResourceAliasesForInstanceOptions(
				"testString",
			)

			resourceAliasesList, response, err := resourceControllerService.ListResourceAliasesForInstance(listResourceAliasesForInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceAliasesList, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_aliases_for_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceAliasesList).ToNot(BeNil())

		})
		It(`ListResourceKeysForInstance request example`, func() {
			fmt.Println("\nListResourceKeysForInstance() result:")
			// begin-list_resource_keys_for_instance

			listResourceKeysForInstanceOptions := resourceControllerService.NewListResourceKeysForInstanceOptions(
				"testString",
			)

			resourceKeysList, response, err := resourceControllerService.ListResourceKeysForInstance(listResourceKeysForInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceKeysList, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_keys_for_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceKeysList).ToNot(BeNil())

		})
		It(`LockResourceInstance request example`, func() {
			fmt.Println("\nLockResourceInstance() result:")
			// begin-lock_resource_instance

			lockResourceInstanceOptions := resourceControllerService.NewLockResourceInstanceOptions(
				resourceInstanceIDLink,
			)

			resourceInstance, response, err := resourceControllerService.LockResourceInstance(lockResourceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceInstance, "", "  ")
			fmt.Println(string(b))

			// end-lock_resource_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstance).ToNot(BeNil())

		})
		It(`ListResourceKeys request example`, func() {
			fmt.Println("\nListResourceKeys() result:")
			// begin-list_resource_keys

			listResourceKeysOptions := resourceControllerService.NewListResourceKeysOptions()
			listResourceKeysOptions.SetUpdatedFrom("2021-01-01")
			listResourceKeysOptions.SetUpdatedTo("2021-01-01")

			resourceKeysList, response, err := resourceControllerService.ListResourceKeys(listResourceKeysOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceKeysList, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_keys

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceKeysList).ToNot(BeNil())

		})
		It(`GetResourceKey request example`, func() {
			fmt.Println("\nGetResourceKey() result:")
			// begin-get_resource_key

			getResourceKeyOptions := resourceControllerService.NewGetResourceKeyOptions(
				resourceKeyIDLink,
			)

			resourceKey, response, err := resourceControllerService.GetResourceKey(getResourceKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceKey, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceKey).ToNot(BeNil())

		})
		It(`UpdateResourceKey request example`, func() {
			fmt.Println("\nUpdateResourceKey() result:")
			// begin-update_resource_key

			updateResourceKeyOptions := resourceControllerService.NewUpdateResourceKeyOptions(
				resourceKeyIDLink,
				"UpdatedExampleResourceKey",
			)

			resourceKey, response, err := resourceControllerService.UpdateResourceKey(updateResourceKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceKey, "", "  ")
			fmt.Println(string(b))

			// end-update_resource_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceKey).ToNot(BeNil())

		})
		It(`ListResourceBindings request example`, func() {
			fmt.Println("\nListResourceBindings() result:")
			// begin-list_resource_bindings

			listResourceBindingsOptions := resourceControllerService.NewListResourceBindingsOptions()
			listResourceBindingsOptions.SetUpdatedFrom("2021-01-01")
			listResourceBindingsOptions.SetUpdatedTo("2021-01-01")

			resourceBindingsList, response, err := resourceControllerService.ListResourceBindings(listResourceBindingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceBindingsList, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_bindings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceBindingsList).ToNot(BeNil())

		})
		It(`GetResourceBinding request example`, func() {
			fmt.Println("\nGetResourceBinding() result:")
			// begin-get_resource_binding

			getResourceBindingOptions := resourceControllerService.NewGetResourceBindingOptions(
				resourceBindingIDLink,
			)

			resourceBinding, response, err := resourceControllerService.GetResourceBinding(getResourceBindingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceBinding, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceBinding).ToNot(BeNil())

		})
		It(`UpdateResourceBinding request example`, func() {
			fmt.Println("\nUpdateResourceBinding() result:")
			// begin-update_resource_binding

			updateResourceBindingOptions := resourceControllerService.NewUpdateResourceBindingOptions(
				resourceBindingIDLink,
				"UpdatedExampleResourceBinding",
			)

			resourceBinding, response, err := resourceControllerService.UpdateResourceBinding(updateResourceBindingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceBinding, "", "  ")
			fmt.Println(string(b))

			// end-update_resource_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceBinding).ToNot(BeNil())

		})
		It(`ListResourceAliases request example`, func() {
			fmt.Println("\nListResourceAliases() result:")
			// begin-list_resource_aliases

			listResourceAliasesOptions := resourceControllerService.NewListResourceAliasesOptions()
			listResourceAliasesOptions.SetUpdatedFrom("2021-01-01")
			listResourceAliasesOptions.SetUpdatedTo("2021-01-01")

			resourceAliasesList, response, err := resourceControllerService.ListResourceAliases(listResourceAliasesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceAliasesList, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_aliases

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceAliasesList).ToNot(BeNil())

		})
		It(`GetResourceAlias request example`, func() {
			fmt.Println("\nGetResourceAlias() result:")
			// begin-get_resource_alias

			getResourceAliasOptions := resourceControllerService.NewGetResourceAliasOptions(
				resourceAliasIDLink,
			)

			resourceAlias, response, err := resourceControllerService.GetResourceAlias(getResourceAliasOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceAlias, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_alias

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceAlias).ToNot(BeNil())

		})
		It(`UpdateResourceAlias request example`, func() {
			fmt.Println("\nUpdateResourceAlias() result:")
			// begin-update_resource_alias

			updateResourceAliasOptions := resourceControllerService.NewUpdateResourceAliasOptions(
				resourceAliasIDLink,
				"UpdatedExampleResourceAlias",
			)

			resourceAlias, response, err := resourceControllerService.UpdateResourceAlias(updateResourceAliasOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceAlias, "", "  ")
			fmt.Println(string(b))

			// end-update_resource_alias

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceAlias).ToNot(BeNil())

		})
		It(`ListResourceBindingsForAlias request example`, func() {
			fmt.Println("\nListResourceBindingsForAlias() result:")
			// begin-list_resource_bindings_for_alias

			listResourceBindingsForAliasOptions := resourceControllerService.NewListResourceBindingsForAliasOptions(
				"testString",
			)

			resourceBindingsList, response, err := resourceControllerService.ListResourceBindingsForAlias(listResourceBindingsForAliasOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceBindingsList, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_bindings_for_alias

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceBindingsList).ToNot(BeNil())

		})
		It(`ListReclamations request example`, func() {
			fmt.Println("\nListReclamations() result:")
			// begin-list_reclamations

			listReclamationsOptions := resourceControllerService.NewListReclamationsOptions()

			reclamationsList, response, err := resourceControllerService.ListReclamations(listReclamationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reclamationsList, "", "  ")
			fmt.Println(string(b))

			// end-list_reclamations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reclamationsList).ToNot(BeNil())

		})
		It(`RunReclamationAction request example`, func() {
			fmt.Println("\nRunReclamationAction() result:")
			// begin-run_reclamation_action

			runReclamationActionOptions := resourceControllerService.NewRunReclamationActionOptions(
				"testString",
				"testString",
			)

			reclamation, response, err := resourceControllerService.RunReclamationAction(runReclamationActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reclamation, "", "  ")
			fmt.Println(string(b))

			// end-run_reclamation_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reclamation).ToNot(BeNil())

		})
		It(`UnlockResourceInstance request example`, func() {
			fmt.Println("\nUnlockResourceInstance() result:")
			// begin-unlock_resource_instance

			unlockResourceInstanceOptions := resourceControllerService.NewUnlockResourceInstanceOptions(
				resourceInstanceIDLink,
			)

			resourceInstance, response, err := resourceControllerService.UnlockResourceInstance(unlockResourceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceInstance, "", "  ")
			fmt.Println(string(b))

			// end-unlock_resource_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstance).ToNot(BeNil())

		})
		It(`DeleteResourceKey request example`, func() {
			// begin-delete_resource_key

			deleteResourceKeyOptions := resourceControllerService.NewDeleteResourceKeyOptions(
				resourceKeyIDLink,
			)

			response, err := resourceControllerService.DeleteResourceKey(deleteResourceKeyOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteResourceKey(): %d\n", response.StatusCode)
			}

			// end-delete_resource_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteResourceInstance request example`, func() {
			// begin-delete_resource_instance

			deleteResourceInstanceOptions := resourceControllerService.NewDeleteResourceInstanceOptions(
				resourceInstanceIDLink,
			)

			response, err := resourceControllerService.DeleteResourceInstance(deleteResourceInstanceOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteResourceInstance(): %d\n", response.StatusCode)
			}

			// end-delete_resource_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteResourceBinding request example`, func() {
			// begin-delete_resource_binding

			deleteResourceBindingOptions := resourceControllerService.NewDeleteResourceBindingOptions(
				resourceBindingIDLink,
			)

			response, err := resourceControllerService.DeleteResourceBinding(deleteResourceBindingOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteResourceBinding(): %d\n", response.StatusCode)
			}

			// end-delete_resource_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteResourceAlias request example`, func() {
			// begin-delete_resource_alias

			deleteResourceAliasOptions := resourceControllerService.NewDeleteResourceAliasOptions(
				resourceAliasIDLink,
			)

			response, err := resourceControllerService.DeleteResourceAlias(deleteResourceAliasOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteResourceAlias(): %d\n", response.StatusCode)
			}

			// end-delete_resource_alias

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`CancelLastopResourceInstance request example`, func() {
			fmt.Println("\nCancelLastopResourceInstance() result:")
			// begin-cancel_lastop_resource_instance

			cancelLastopResourceInstanceOptions := resourceControllerService.NewCancelLastopResourceInstanceOptions(
				resourceInstanceIDLink,
			)

			resourceInstance, response, err := resourceControllerService.CancelLastopResourceInstance(cancelLastopResourceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceInstance, "", "  ")
			fmt.Println(string(b))

			// end-cancel_lastop_resource_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstance).ToNot(BeNil())

		})
	})
})
