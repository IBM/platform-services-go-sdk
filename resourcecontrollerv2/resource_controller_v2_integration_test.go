// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the resourcecontrollerv2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`ResourceControllerV2 Integration Tests`, func() {
	const externalConfigFile = "../resource_controller_v2.env"

	var (
		err          error
		resourceControllerService *resourcecontrollerv2.ResourceControllerV2
		serviceURL   string
		config       map[string]string

		// Variables to hold link values
		resourceAliasIDLink string
		resourceBindingIDLink string
		resourceInstanceIDLink string
		resourceKeyIDLink string
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
			config, err = core.GetServiceProperties(resourcecontrollerv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			resourceControllerServiceOptions := &resourcecontrollerv2.ResourceControllerV2Options{}

			resourceControllerService, err = resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(resourceControllerServiceOptions)
			Expect(err).To(BeNil())
			Expect(resourceControllerService).ToNot(BeNil())
			Expect(resourceControllerService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			resourceControllerService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateResourceAlias - Create a new resource alias`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateResourceAlias(createResourceAliasOptions *CreateResourceAliasOptions)`, func() {
			createResourceAliasOptions := &resourcecontrollerv2.CreateResourceAliasOptions{
				Name: core.StringPtr("ExampleResourceAlias"),
				Source: core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63"),
				Target: core.StringPtr("crn:v1:bluemix:public:bluemix:us-south:o/d35d4f0e-5076-4c89-9361-2522894b6548::cf-space:e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7"),
			}

			resourceAlias, response, err := resourceControllerService.CreateResourceAlias(createResourceAliasOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resourceAlias).ToNot(BeNil())

			resourceAliasIDLink = *resourceAlias.GUID
			fmt.Fprintf(GinkgoWriter, "Saved resourceAliasIDLink value: %v\n", resourceAliasIDLink)
		})
	})

	Describe(`CreateResourceBinding - Create a new resource binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateResourceBinding(createResourceBindingOptions *CreateResourceBindingOptions)`, func() {
			resourceBindingPostParametersModel := &resourcecontrollerv2.ResourceBindingPostParameters{
				ServiceidCRN: core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393"),
			}
			resourceBindingPostParametersModel.SetProperty("exampleParameter", core.StringPtr("exampleValue"))

			createResourceBindingOptions := &resourcecontrollerv2.CreateResourceBindingOptions{
				Source: core.StringPtr("faaec9d8-ec64-44d8-ab83-868632fac6a2"),
				Target: core.StringPtr("crn:v1:staging:public:bluemix:us-south:s/e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7::cf-application:8d9457e0-1303-4f32-b4b3-5525575f6205"),
				Name: core.StringPtr("ExampleResourceBinding"),
				Parameters: resourceBindingPostParametersModel,
				Role: core.StringPtr("Writer"),
			}

			resourceBinding, response, err := resourceControllerService.CreateResourceBinding(createResourceBindingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resourceBinding).ToNot(BeNil())

			resourceBindingIDLink = *resourceBinding.GUID
			fmt.Fprintf(GinkgoWriter, "Saved resourceBindingIDLink value: %v\n", resourceBindingIDLink)
		})
	})

	Describe(`CreateResourceInstance - Create (provision) a new resource instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateResourceInstance(createResourceInstanceOptions *CreateResourceInstanceOptions)`, func() {
			createResourceInstanceOptions := &resourcecontrollerv2.CreateResourceInstanceOptions{
				Name: core.StringPtr("ExampleResourceInstance"),
				Target: core.StringPtr("global"),
				ResourceGroup: core.StringPtr("13aa3ee48c3b44ddb64c05c79f7ab8ef"),
				ResourcePlanID: core.StringPtr("a10e4960-3685-11e9-b210-d663bd873d93"),
				Tags: []string{"testString"},
				AllowCleanup: core.BoolPtr(false),
				Parameters: make(map[string]interface{}),
				EntityLock: core.BoolPtr(false),
			}

			resourceInstance, response, err := resourceControllerService.CreateResourceInstance(createResourceInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resourceInstance).ToNot(BeNil())

			resourceInstanceIDLink = *resourceInstance.GUID
			fmt.Fprintf(GinkgoWriter, "Saved resourceInstanceIDLink value: %v\n", resourceInstanceIDLink)
		})
	})

	Describe(`CreateResourceKey - Create a new resource key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateResourceKey(createResourceKeyOptions *CreateResourceKeyOptions)`, func() {
			resourceKeyPostParametersModel := &resourcecontrollerv2.ResourceKeyPostParameters{
				ServiceidCRN: core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393"),
			}
			resourceKeyPostParametersModel.SetProperty("exampleParameter", core.StringPtr("exampleValue"))

			createResourceKeyOptions := &resourcecontrollerv2.CreateResourceKeyOptions{
				Name: core.StringPtr("ExampleResourceKey"),
				Source: &resourceInstanceIDLink,
				Parameters: resourceKeyPostParametersModel,
				Role: core.StringPtr("Writer"),
			}

			resourceKey, response, err := resourceControllerService.CreateResourceKey(createResourceKeyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resourceKey).ToNot(BeNil())

			resourceKeyIDLink = *resourceKey.GUID
			fmt.Fprintf(GinkgoWriter, "Saved resourceKeyIDLink value: %v\n", resourceKeyIDLink)
		})
	})

	Describe(`ListResourceInstances - Get a list of all resource instances`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceInstances(listResourceInstancesOptions *ListResourceInstancesOptions)`, func() {
			listResourceInstancesOptions := &resourcecontrollerv2.ListResourceInstancesOptions{
				GUID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				ResourcePlanID: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
				SubType: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
				State: core.StringPtr("active"),
				UpdatedFrom: core.StringPtr("2021-01-01"),
				UpdatedTo: core.StringPtr("2021-01-01"),
			}

			resourceInstancesList, response, err := resourceControllerService.ListResourceInstances(listResourceInstancesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstancesList).ToNot(BeNil())
		})
	})

	Describe(`GetResourceInstance - Get a resource instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceInstance(getResourceInstanceOptions *GetResourceInstanceOptions)`, func() {
			getResourceInstanceOptions := &resourcecontrollerv2.GetResourceInstanceOptions{
				ID: &resourceInstanceIDLink,
			}

			resourceInstance, response, err := resourceControllerService.GetResourceInstance(getResourceInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstance).ToNot(BeNil())
		})
	})

	Describe(`UpdateResourceInstance - Update a resource instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateResourceInstance(updateResourceInstanceOptions *UpdateResourceInstanceOptions)`, func() {
			updateResourceInstanceOptions := &resourcecontrollerv2.UpdateResourceInstanceOptions{
				ID: &resourceInstanceIDLink,
				Name: core.StringPtr("UpdatedExampleResourceInstance"),
				Parameters: make(map[string]interface{}),
				ResourcePlanID: core.StringPtr("testString"),
				AllowCleanup: core.BoolPtr(true),
			}

			resourceInstance, response, err := resourceControllerService.UpdateResourceInstance(updateResourceInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstance).ToNot(BeNil())
		})
	})

	Describe(`ListResourceAliasesForInstance - Get a list of all resource aliases for the instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceAliasesForInstance(listResourceAliasesForInstanceOptions *ListResourceAliasesForInstanceOptions)`, func() {
			listResourceAliasesForInstanceOptions := &resourcecontrollerv2.ListResourceAliasesForInstanceOptions{
				ID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			resourceAliasesList, response, err := resourceControllerService.ListResourceAliasesForInstance(listResourceAliasesForInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceAliasesList).ToNot(BeNil())
		})
	})

	Describe(`ListResourceKeysForInstance - Get a list of all the resource keys for the instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceKeysForInstance(listResourceKeysForInstanceOptions *ListResourceKeysForInstanceOptions)`, func() {
			listResourceKeysForInstanceOptions := &resourcecontrollerv2.ListResourceKeysForInstanceOptions{
				ID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			resourceKeysList, response, err := resourceControllerService.ListResourceKeysForInstance(listResourceKeysForInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceKeysList).ToNot(BeNil())
		})
	})

	Describe(`LockResourceInstance - Lock a resource instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockResourceInstance(lockResourceInstanceOptions *LockResourceInstanceOptions)`, func() {
			lockResourceInstanceOptions := &resourcecontrollerv2.LockResourceInstanceOptions{
				ID: &resourceInstanceIDLink,
			}

			resourceInstance, response, err := resourceControllerService.LockResourceInstance(lockResourceInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstance).ToNot(BeNil())
		})
	})

	Describe(`ListResourceKeys - Get a list of all of the resource keys`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceKeys(listResourceKeysOptions *ListResourceKeysOptions)`, func() {
			listResourceKeysOptions := &resourcecontrollerv2.ListResourceKeysOptions{
				GUID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
				UpdatedFrom: core.StringPtr("2021-01-01"),
				UpdatedTo: core.StringPtr("2021-01-01"),
			}

			resourceKeysList, response, err := resourceControllerService.ListResourceKeys(listResourceKeysOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceKeysList).ToNot(BeNil())
		})
	})

	Describe(`GetResourceKey - Get resource key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceKey(getResourceKeyOptions *GetResourceKeyOptions)`, func() {
			getResourceKeyOptions := &resourcecontrollerv2.GetResourceKeyOptions{
				ID: &resourceKeyIDLink,
			}

			resourceKey, response, err := resourceControllerService.GetResourceKey(getResourceKeyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceKey).ToNot(BeNil())
		})
	})

	Describe(`UpdateResourceKey - Update a resource key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateResourceKey(updateResourceKeyOptions *UpdateResourceKeyOptions)`, func() {
			updateResourceKeyOptions := &resourcecontrollerv2.UpdateResourceKeyOptions{
				ID: &resourceKeyIDLink,
				Name: core.StringPtr("UpdatedExampleResourceKey"),
			}

			resourceKey, response, err := resourceControllerService.UpdateResourceKey(updateResourceKeyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceKey).ToNot(BeNil())
		})
	})

	Describe(`ListResourceBindings - Get a list of all resource bindings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceBindings(listResourceBindingsOptions *ListResourceBindingsOptions)`, func() {
			listResourceBindingsOptions := &resourcecontrollerv2.ListResourceBindingsOptions{
				GUID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				RegionBindingID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
				UpdatedFrom: core.StringPtr("2021-01-01"),
				UpdatedTo: core.StringPtr("2021-01-01"),
			}

			resourceBindingsList, response, err := resourceControllerService.ListResourceBindings(listResourceBindingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceBindingsList).ToNot(BeNil())
		})
	})

	Describe(`GetResourceBinding - Get a resource binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceBinding(getResourceBindingOptions *GetResourceBindingOptions)`, func() {
			getResourceBindingOptions := &resourcecontrollerv2.GetResourceBindingOptions{
				ID: &resourceBindingIDLink,
			}

			resourceBinding, response, err := resourceControllerService.GetResourceBinding(getResourceBindingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceBinding).ToNot(BeNil())
		})
	})

	Describe(`UpdateResourceBinding - Update a resource binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateResourceBinding(updateResourceBindingOptions *UpdateResourceBindingOptions)`, func() {
			updateResourceBindingOptions := &resourcecontrollerv2.UpdateResourceBindingOptions{
				ID: &resourceBindingIDLink,
				Name: core.StringPtr("UpdatedExampleResourceBinding"),
			}

			resourceBinding, response, err := resourceControllerService.UpdateResourceBinding(updateResourceBindingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceBinding).ToNot(BeNil())
		})
	})

	Describe(`ListResourceAliases - Get a list of all resource aliases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceAliases(listResourceAliasesOptions *ListResourceAliasesOptions)`, func() {
			listResourceAliasesOptions := &resourcecontrollerv2.ListResourceAliasesOptions{
				GUID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ResourceInstanceID: core.StringPtr("testString"),
				RegionInstanceID: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
				UpdatedFrom: core.StringPtr("2021-01-01"),
				UpdatedTo: core.StringPtr("2021-01-01"),
			}

			resourceAliasesList, response, err := resourceControllerService.ListResourceAliases(listResourceAliasesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceAliasesList).ToNot(BeNil())
		})
	})

	Describe(`GetResourceAlias - Get a resource alias`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceAlias(getResourceAliasOptions *GetResourceAliasOptions)`, func() {
			getResourceAliasOptions := &resourcecontrollerv2.GetResourceAliasOptions{
				ID: &resourceAliasIDLink,
			}

			resourceAlias, response, err := resourceControllerService.GetResourceAlias(getResourceAliasOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceAlias).ToNot(BeNil())
		})
	})

	Describe(`UpdateResourceAlias - Update a resource alias`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateResourceAlias(updateResourceAliasOptions *UpdateResourceAliasOptions)`, func() {
			updateResourceAliasOptions := &resourcecontrollerv2.UpdateResourceAliasOptions{
				ID: &resourceAliasIDLink,
				Name: core.StringPtr("UpdatedExampleResourceAlias"),
			}

			resourceAlias, response, err := resourceControllerService.UpdateResourceAlias(updateResourceAliasOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceAlias).ToNot(BeNil())
		})
	})

	Describe(`ListResourceBindingsForAlias - Get a list of all resource bindings for the alias`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceBindingsForAlias(listResourceBindingsForAliasOptions *ListResourceBindingsForAliasOptions)`, func() {
			listResourceBindingsForAliasOptions := &resourcecontrollerv2.ListResourceBindingsForAliasOptions{
				ID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			resourceBindingsList, response, err := resourceControllerService.ListResourceBindingsForAlias(listResourceBindingsForAliasOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceBindingsList).ToNot(BeNil())
		})
	})

	Describe(`ListReclamations - Get a list of all reclamations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReclamations(listReclamationsOptions *ListReclamationsOptions)`, func() {
			listReclamationsOptions := &resourcecontrollerv2.ListReclamationsOptions{
				AccountID: core.StringPtr("testString"),
				ResourceInstanceID: core.StringPtr("testString"),
			}

			reclamationsList, response, err := resourceControllerService.ListReclamations(listReclamationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reclamationsList).ToNot(BeNil())
		})
	})

	Describe(`RunReclamationAction - Perform a reclamation action`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunReclamationAction(runReclamationActionOptions *RunReclamationActionOptions)`, func() {
			runReclamationActionOptions := &resourcecontrollerv2.RunReclamationActionOptions{
				ID: core.StringPtr("testString"),
				ActionName: core.StringPtr("testString"),
				RequestBy: core.StringPtr("testString"),
				Comment: core.StringPtr("testString"),
			}

			reclamation, response, err := resourceControllerService.RunReclamationAction(runReclamationActionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reclamation).ToNot(BeNil())
		})
	})

	Describe(`UnlockResourceInstance - Unlock a resource instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockResourceInstance(unlockResourceInstanceOptions *UnlockResourceInstanceOptions)`, func() {
			unlockResourceInstanceOptions := &resourcecontrollerv2.UnlockResourceInstanceOptions{
				ID: &resourceInstanceIDLink,
			}

			resourceInstance, response, err := resourceControllerService.UnlockResourceInstance(unlockResourceInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstance).ToNot(BeNil())
		})
	})

	Describe(`DeleteResourceKey - Delete a resource key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteResourceKey(deleteResourceKeyOptions *DeleteResourceKeyOptions)`, func() {
			deleteResourceKeyOptions := &resourcecontrollerv2.DeleteResourceKeyOptions{
				ID: &resourceKeyIDLink,
			}

			response, err := resourceControllerService.DeleteResourceKey(deleteResourceKeyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteResourceInstance - Delete a resource instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteResourceInstance(deleteResourceInstanceOptions *DeleteResourceInstanceOptions)`, func() {
			deleteResourceInstanceOptions := &resourcecontrollerv2.DeleteResourceInstanceOptions{
				ID: &resourceInstanceIDLink,
				Recursive: core.BoolPtr(false),
			}

			response, err := resourceControllerService.DeleteResourceInstance(deleteResourceInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`DeleteResourceBinding - Delete a resource binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteResourceBinding(deleteResourceBindingOptions *DeleteResourceBindingOptions)`, func() {
			deleteResourceBindingOptions := &resourcecontrollerv2.DeleteResourceBindingOptions{
				ID: &resourceBindingIDLink,
			}

			response, err := resourceControllerService.DeleteResourceBinding(deleteResourceBindingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteResourceAlias - Delete a resource alias`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteResourceAlias(deleteResourceAliasOptions *DeleteResourceAliasOptions)`, func() {
			deleteResourceAliasOptions := &resourcecontrollerv2.DeleteResourceAliasOptions{
				ID: &resourceAliasIDLink,
				Recursive: core.BoolPtr(false),
			}

			response, err := resourceControllerService.DeleteResourceAlias(deleteResourceAliasOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`CancelLastopResourceInstance - Cancel the in progress last operation of the resource instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CancelLastopResourceInstance(cancelLastopResourceInstanceOptions *CancelLastopResourceInstanceOptions)`, func() {
			cancelLastopResourceInstanceOptions := &resourcecontrollerv2.CancelLastopResourceInstanceOptions{
				ID: &resourceInstanceIDLink,
			}

			resourceInstance, response, err := resourceControllerService.CancelLastopResourceInstance(cancelLastopResourceInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceInstance).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
