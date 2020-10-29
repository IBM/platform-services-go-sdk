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

package resourcecontrollerv2_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const externalConfigFile = "../resource_controller.env"

var (
	resourceControllerService *resourcecontrollerv2.ResourceControllerV2
	config                    map[string]string
	configLoaded              bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`ResourceControllerV2 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(resourcecontrollerv2.DefaultServiceName)
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
		It(`ListResourceInstances request example`, func() {
			// begin-list_resource_instances

			listResourceInstancesOptions := resourceControllerService.NewListResourceInstancesOptions()
			listResourceInstancesOptions.SetUpdatedFrom("2019-01-08T00:00:00.000Z")
			listResourceInstancesOptions.SetUpdatedTo("2019-01-08T00:00:00.000Z")

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
		It(`CreateResourceInstance request example`, func() {
			// begin-create_resource_instance

			createResourceInstanceOptions := resourceControllerService.NewCreateResourceInstanceOptions(
				"my-instance",
				"bluemix-us-south",
				"5c49eabc-f5e8-5881-a37e-2d100a33b3df",
				"cloudant-standard",
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

		})
		It(`GetResourceInstance request example`, func() {
			// begin-get_resource_instance

			getResourceInstanceOptions := resourceControllerService.NewGetResourceInstanceOptions(
				"testString",
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
			// begin-update_resource_instance

			updateResourceInstanceOptions := resourceControllerService.NewUpdateResourceInstanceOptions(
				"testString",
			)

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
		It(`LockResourceInstance request example`, func() {
			// begin-lock_resource_instance

			lockResourceInstanceOptions := resourceControllerService.NewLockResourceInstanceOptions(
				"testString",
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
			// begin-list_resource_keys

			listResourceKeysOptions := resourceControllerService.NewListResourceKeysOptions()
			listResourceKeysOptions.SetUpdatedFrom("2019-01-08T00:00:00.000Z")
			listResourceKeysOptions.SetUpdatedTo("2019-01-08T00:00:00.000Z")

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
		It(`CreateResourceKey request example`, func() {
			// begin-create_resource_key

			createResourceKeyOptions := resourceControllerService.NewCreateResourceKeyOptions(
				"my-key",
				"25eba2a9-beef-450b-82cf-f5ad5e36c6dd",
			)

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

		})
		It(`GetResourceKey request example`, func() {
			// begin-get_resource_key

			getResourceKeyOptions := resourceControllerService.NewGetResourceKeyOptions(
				"testString",
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
			// begin-update_resource_key

			updateResourceKeyOptions := resourceControllerService.NewUpdateResourceKeyOptions(
				"testString",
				"my-new-key-name",
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
			// begin-list_resource_bindings

			listResourceBindingsOptions := resourceControllerService.NewListResourceBindingsOptions()
			listResourceBindingsOptions.SetUpdatedFrom("2019-01-08T00:00:00.000Z")
			listResourceBindingsOptions.SetUpdatedTo("2019-01-08T00:00:00.000Z")

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
		It(`CreateResourceBinding request example`, func() {
			// begin-create_resource_binding

			createResourceBindingOptions := resourceControllerService.NewCreateResourceBindingOptions(
				"25eba2a9-beef-450b-82cf-f5ad5e36c6dd",
				"crn:v1:cf:public:cf:us-south:s/0ba4dba0-a120-4a1e-a124-5a249a904b76::cf-application:a1caa40b-2c24-4da8-8267-ac2c1a42ad0c",
			)

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

		})
		It(`GetResourceBinding request example`, func() {
			// begin-get_resource_binding

			getResourceBindingOptions := resourceControllerService.NewGetResourceBindingOptions(
				"testString",
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
			// begin-update_resource_binding

			updateResourceBindingOptions := resourceControllerService.NewUpdateResourceBindingOptions(
				"testString",
				"my-new-binding-name",
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
			// begin-list_resource_aliases

			listResourceAliasesOptions := resourceControllerService.NewListResourceAliasesOptions()
			listResourceAliasesOptions.SetUpdatedFrom("2019-01-08T00:00:00.000Z")
			listResourceAliasesOptions.SetUpdatedTo("2019-01-08T00:00:00.000Z")

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
		It(`CreateResourceAlias request example`, func() {
			// begin-create_resource_alias

			createResourceAliasOptions := resourceControllerService.NewCreateResourceAliasOptions(
				"my-alias",
				"a8dff6d3-d287-4668-a81d-c87c55c2656d",
				"crn:v1:cf:public:cf:us-south:o/5e939cd5-6377-4383-b9e0-9db22cd11753::cf-space:66c8b915-101a-406c-a784-e6636676e4f5",
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

		})
		It(`GetResourceAlias request example`, func() {
			// begin-get_resource_alias

			getResourceAliasOptions := resourceControllerService.NewGetResourceAliasOptions(
				"testString",
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
			// begin-update_resource_alias

			updateResourceAliasOptions := resourceControllerService.NewUpdateResourceAliasOptions(
				"testString",
				"my-new-alias-name",
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
		It(`ListReclamations request example`, func() {
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
			// begin-unlock_resource_instance

			unlockResourceInstanceOptions := resourceControllerService.NewUnlockResourceInstanceOptions(
				"testString",
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
				"testString",
			)

			response, err := resourceControllerService.DeleteResourceKey(deleteResourceKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteResourceInstance request example`, func() {
			// begin-delete_resource_instance

			deleteResourceInstanceOptions := resourceControllerService.NewDeleteResourceInstanceOptions(
				"testString",
			)

			response, err := resourceControllerService.DeleteResourceInstance(deleteResourceInstanceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteResourceBinding request example`, func() {
			// begin-delete_resource_binding

			deleteResourceBindingOptions := resourceControllerService.NewDeleteResourceBindingOptions(
				"testString",
			)

			response, err := resourceControllerService.DeleteResourceBinding(deleteResourceBindingOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteResourceAlias request example`, func() {
			// begin-delete_resource_alias

			deleteResourceAliasOptions := resourceControllerService.NewDeleteResourceAliasOptions(
				"testString",
			)

			response, err := resourceControllerService.DeleteResourceAlias(deleteResourceAliasOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_alias

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
