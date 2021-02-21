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
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Resource Controller service.
//
// The following configuration properties are assumed to be defined:
//
// RESOURCE_CONTROLLER_URL=<service url>
// RESOURCE_CONTROLLER_AUTH_TYPE=iam
// RESOURCE_CONTROLLER_AUTH_URL=<IAM Token Service url>
// RESOURCE_CONTROLLER_APIKEY=<User's IAM API Key>
// RESOURCE_CONTROLLER_RESOURCE_GROUP=<Short ID of the user's resource group>
// RESOURCE_CONTROLLER_PLAN_ID=<Unique ID of the plan associated with the offering>
// RESOURCE_CONTROLLER_ACCOUNT_ID=<User's account ID>
// RESOURCE_CONTROLLER_ALIAS_TARGET_CRN=<The CRN of target name(space) in a specific environment>
// RESOURCE_CONTROLLER_BINDING_TARGET_CRN=<The CRN of application to bind to in a specific environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../resource_controller.env"

var (
	resourceControllerService *resourcecontrollerv2.ResourceControllerV2
	config                    map[string]string
	configLoaded              bool = false

	instanceGUID               string
	aliasGUID                  string
	bindingGUID                string
	instanceKeyGUID            string
	resourceGroup              string
	resourcePlanID             string
	accountID                  string
	aliasTargetCRN             string
	bindingTargetCRN           string
	reclamationID              string
	resourceInstanceName       string = "RcSdkInstance1Go"
	resourceInstanceUpdateName string = "RcSdkInstanceUpdate1Go"
	aliasName                  string = "RcSdkAlias1Go"
	aliasUpdateName            string = "RcSdkAliasUpdate1Go"
	bindingName                string = "RcSdkBinding1Go"
	bindingUpdateName          string = "RcSdkBindingUpdate1Go"
	keyName                    string = "RcSdkKey1Go"
	keyUpdateName              string = "RcSdkKeyUpdate1Go"
	targetRegion               string = "global"
	reclaimAction              string = "reclaim"
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

			resourceGroup = config["RESOURCE_GROUP"]
			Expect(resourceGroup).ToNot(BeEmpty())

			resourcePlanID = config["RECLAMATION_PLAN_ID"]
			Expect(resourcePlanID).ToNot(BeEmpty())

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			aliasTargetCRN = config["ALIAS_TARGET_CRN"]
			Expect(aliasTargetCRN).ToNot(BeEmpty())

			bindingTargetCRN = config["BINDING_TARGET_CRN"]
			Expect(bindingTargetCRN).ToNot(BeEmpty())
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			options := &resourcecontrollerv2.ResourceControllerV2Options{}

			resourceControllerService, err = resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(options)

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
		It(`CreateResourceInstance request example`, func() {
			// begin-create_resource_instance

			createResourceInstanceOptions := resourceControllerService.NewCreateResourceInstanceOptions(
				resourceInstanceName,
				targetRegion,
				resourceGroup,
				resourcePlanID,
			)

			resourceInstance, response, err := resourceControllerService.CreateResourceInstance(createResourceInstanceOptions)
			if err != nil {
				panic(err)
			}
			instanceGUID = *resourceInstance.GUID

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
				instanceGUID,
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

			params := make(map[string]interface{}, 0)
			params["example"] = "property"
			updateResourceInstanceOptions := resourceControllerService.NewUpdateResourceInstanceOptions(
				instanceGUID,
			)
			updateResourceInstanceOptions = updateResourceInstanceOptions.SetName(resourceInstanceUpdateName)
			updateResourceInstanceOptions = updateResourceInstanceOptions.SetParameters(params)

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
		It(`ListResourceInstances request example`, func() {
			// begin-list_resource_instances

			listResourceInstancesOptions := resourceControllerService.NewListResourceInstancesOptions()
			listResourceInstancesOptions = listResourceInstancesOptions.SetName(resourceInstanceName)

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
		It(`CreateResourceAlias request example`, func() {
			// begin-create_resource_alias

			createResourceAliasOptions := resourceControllerService.NewCreateResourceAliasOptions(
				aliasName,
				instanceGUID,
				aliasTargetCRN,
			)

			resourceAlias, response, err := resourceControllerService.CreateResourceAlias(createResourceAliasOptions)
			if err != nil {
				panic(err)
			}
			aliasGUID = *resourceAlias.GUID

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
				aliasGUID,
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
				aliasGUID,
				aliasUpdateName,
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
		It(`ListResourceAliases request example`, func() {
			// begin-list_resource_aliases

			listResourceAliasesOptions := resourceControllerService.NewListResourceAliasesOptions()
			listResourceAliasesOptions = listResourceAliasesOptions.SetName(aliasName)

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
		It(`CreateResourceBinding request example`, func() {
			// begin-create_resource_binding

			createResourceBindingOptions := resourceControllerService.NewCreateResourceBindingOptions(
				aliasGUID,
				bindingTargetCRN,
			)
			createResourceBindingOptions = createResourceBindingOptions.SetName(bindingName)

			resourceBinding, response, err := resourceControllerService.CreateResourceBinding(createResourceBindingOptions)
			if err != nil {
				panic(err)
			}
			bindingGUID = *resourceBinding.GUID

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
				bindingGUID,
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
				bindingGUID,
				bindingUpdateName,
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
		It(`ListResourceBindings request example`, func() {
			// begin-list_resource_bindings

			listResourceBindingsOptions := resourceControllerService.NewListResourceBindingsOptions()
			listResourceBindingsOptions = listResourceBindingsOptions.SetName(bindingName)

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
		It(`CreateResourceKey request example`, func() {
			// begin-create_resource_key

			createResourceKeyOptions := resourceControllerService.NewCreateResourceKeyOptions(
				keyName,
				instanceGUID,
			)

			resourceKey, response, err := resourceControllerService.CreateResourceKey(createResourceKeyOptions)
			if err != nil {
				panic(err)
			}
			instanceKeyGUID = *resourceKey.GUID

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
				instanceKeyGUID,
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
				instanceKeyGUID,
				keyUpdateName,
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
		It(`ListResourceKeys request example`, func() {
			// begin-list_resource_keys

			listResourceKeysOptions := resourceControllerService.NewListResourceKeysOptions()
			listResourceKeysOptions = listResourceKeysOptions.SetName(keyName)

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
		It(`DeleteResourceBinding request example`, func() {
			// begin-delete_resource_binding

			deleteResourceBindingOptions := resourceControllerService.NewDeleteResourceBindingOptions(
				bindingGUID,
			)

			response, err := resourceControllerService.DeleteResourceBinding(deleteResourceBindingOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteResourceKey request example`, func() {
			// begin-delete_resource_key

			deleteResourceKeyOptions := resourceControllerService.NewDeleteResourceKeyOptions(
				instanceKeyGUID,
			)

			response, err := resourceControllerService.DeleteResourceKey(deleteResourceKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteResourceAlias request example`, func() {
			// begin-delete_resource_alias

			deleteResourceAliasOptions := resourceControllerService.NewDeleteResourceAliasOptions(
				aliasGUID,
			)

			response, err := resourceControllerService.DeleteResourceAlias(deleteResourceAliasOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_alias

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`LockResourceInstance request example`, func() {
			// begin-lock_resource_instance

			lockResourceInstanceOptions := resourceControllerService.NewLockResourceInstanceOptions(
				instanceGUID,
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
		It(`UnlockResourceInstance request example`, func() {
			// begin-unlock_resource_instance

			unlockResourceInstanceOptions := resourceControllerService.NewUnlockResourceInstanceOptions(
				instanceGUID,
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
		It(`DeleteResourceInstance request example`, func() {
			// begin-delete_resource_instance

			deleteResourceInstanceOptions := resourceControllerService.NewDeleteResourceInstanceOptions(
				instanceGUID,
			)

			response, err := resourceControllerService.DeleteResourceInstance(deleteResourceInstanceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_instance

			time.Sleep(20 * time.Second)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListReclamations request example`, func() {
			// begin-list_reclamations

			listReclamationsOptions := resourceControllerService.NewListReclamationsOptions()
			listReclamationsOptions = listReclamationsOptions.SetAccountID(accountID)
			reclamationsList, response, err := resourceControllerService.ListReclamations(listReclamationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reclamationsList, "", "  ")
			fmt.Println(string(b))

			for _, res := range reclamationsList.Resources {
				if *res.ResourceInstanceID == instanceGUID {
					reclamationID = *res.ID
				}
			}

			// end-list_reclamations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reclamationsList).ToNot(BeNil())

		})
		It(`RunReclamationAction request example`, func() {
			// begin-run_reclamation_action

			runReclamationActionOptions := resourceControllerService.NewRunReclamationActionOptions(
				reclamationID,
				reclaimAction,
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
			//wait for reclamation object to be created
			time.Sleep(20 * time.Second)
		})
	})
})
