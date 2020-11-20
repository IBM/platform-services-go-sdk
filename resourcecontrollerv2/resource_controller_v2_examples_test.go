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

const externalConfigFile = "../resource_controller.env"

///////////////////////////////////////////////////////
///////////// Example config file fields //////////////
// RESOURCE_CONTROLLER_URL=https://resource-controller.cloud.ibm.com
// RESOURCE_CONTROLLER_AUTH_TYPE=iam
// RESOURCE_CONTROLLER_AUTH_URL=https://iam.cloud.ibm.com/identity/token
// RESOURCE_CONTROLLER_APIKEY=<User's IAM API Key>
// RESOURCE_CONTROLLER_RESOURCE_GROUP=5g9f447903254bb58972a2f3f5a4c711
// RESOURCE_CONTROLLER_RECLAMATION_PLAN_ID=0be5ad401ae913d8ff665d92680664ed
// RESOURCE_CONTROLLER_ACCOUNT_ID=b80a8b513ae24e178438b7a18bd8d609
// RESOURCE_CONTROLLER_ALIAS_TARGET_CRN=crn:v1:cf:public:cf:eu-gb:o/e242c7f0-9eb7-4541-ad3e-b5f5a45a1498::cf-space:f5038ca8-9d28-42a1-9e57-9b9fdd66bf8e
// RESOURCE_CONTROLLER_BINDING_TARGET_CRN=crn:v1:cf:public:cf:eu-gb:s/f5038ca8-9d28-42a1-9e57-9b9fdd66bf8e::cf-application:b04ddee1-2838-449a-96d3-02a03179e991
///////////////////////////////////////////////////////

var (
	resourceControllerService *resourcecontrollerv2.ResourceControllerV2
	config                    map[string]string
	configLoaded              bool = false

	instanceGuid               string
	aliasGuid                  string
	bindingGuid                string
	instanceKeyGuid            string
	resourceGroup              string
	resourcePlanId             string
	accountId                  string
	aliasTargetCRN             string
	bindingTargetCRN           string
	reclamationId              string
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

			resourcePlanId = config["RECLAMATION_PLAN_ID"]
			Expect(resourcePlanId).ToNot(BeEmpty())

			accountId = config["ACCOUNT_ID"]
			Expect(accountId).ToNot(BeEmpty())

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
				resourcePlanId,
			)

			resourceInstance, response, err := resourceControllerService.CreateResourceInstance(createResourceInstanceOptions)
			if err != nil {
				panic(err)
			}
			instanceGuid = *resourceInstance.Guid

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
				instanceGuid,
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
				instanceGuid,
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
				instanceGuid,
				aliasTargetCRN,
			)

			resourceAlias, response, err := resourceControllerService.CreateResourceAlias(createResourceAliasOptions)
			if err != nil {
				panic(err)
			}
			aliasGuid = *resourceAlias.Guid

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
				aliasGuid,
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
				aliasGuid,
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
				aliasGuid,
				bindingTargetCRN,
			)
			createResourceBindingOptions = createResourceBindingOptions.SetName(bindingName)

			resourceBinding, response, err := resourceControllerService.CreateResourceBinding(createResourceBindingOptions)
			if err != nil {
				panic(err)
			}
			bindingGuid = *resourceBinding.Guid

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
				bindingGuid,
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
				bindingGuid,
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
				instanceGuid,
			)

			resourceKey, response, err := resourceControllerService.CreateResourceKey(createResourceKeyOptions)
			if err != nil {
				panic(err)
			}
			instanceKeyGuid = *resourceKey.Guid

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
				instanceKeyGuid,
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
				instanceKeyGuid,
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
				bindingGuid,
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
				instanceKeyGuid,
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
				aliasGuid,
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
				instanceGuid,
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
				instanceGuid,
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
				instanceGuid,
			)

			response, err := resourceControllerService.DeleteResourceInstance(deleteResourceInstanceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListReclamations request example`, func() {
			// begin-list_reclamations

			listReclamationsOptions := resourceControllerService.NewListReclamationsOptions()
			listReclamationsOptions = listReclamationsOptions.SetAccountID(accountId)
			reclamationsList, response, err := resourceControllerService.ListReclamations(listReclamationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reclamationsList, "", "  ")
			fmt.Println(string(b))

			for _, res := range reclamationsList.Resources {
				if *res.ResourceInstanceID == instanceGuid {
					reclamationId = *res.ID
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
				reclamationId,
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
