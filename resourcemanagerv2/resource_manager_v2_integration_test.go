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
package resourcemanagerv2_test

import (
	"github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
)

const externalConfigFile = "../resource_manager.env"

var (
	service1           *resourcemanagerv2.ResourceManagerV2
	service2           *resourcemanagerv2.ResourceManagerV2
	err                error
	testQuotaID        string = "7ce89f4a-4381-4600-b814-3cd9a4f4bdf4"
	testUserAccountID  string = "60ce10d1d94749bf8dceff12065db1b0"
	newResourceGroupID string
	configLoaded       bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe("Resource Manager - Integration Tests", func() {
	It("Successfully load the configuration", func() {
		_, err = os.Stat(externalConfigFile)
		if err == nil {
			err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			if err == nil {
				configLoaded = true
			}
		}
		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}
	})
	It(`Successfully created ResourceManagerV2 service instances`, func() {
		shouldSkipTest()
		options1 := &resourcemanagerv2.ResourceManagerV2Options{
			ServiceName: "RMGR1",
		}
		service1, err = resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(options1)
		Expect(err).To(BeNil())
		Expect(service1).ToNot(BeNil())

		options2 := &resourcemanagerv2.ResourceManagerV2Options{
			ServiceName: "RMGR2",
		}
		service2, err = resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(options2)
		Expect(err).To(BeNil())
		Expect(service2).ToNot(BeNil())
	})

	It("Get list of all quota definition", func() {
		shouldSkipTest()
		listQuotaDefinitionOptionsModel := service1.NewListQuotaDefinitionsOptions()
		result, detailedResponse, err := service1.ListQuotaDefinitions(listQuotaDefinitionOptionsModel)
		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))
		Expect(result.Resources).NotTo(BeNil())
	})

	It("Get a quota definition by id", func() {
		shouldSkipTest()
		getQuotaDefinitionOptionsModel := service1.NewGetQuotaDefinitionOptions(testQuotaID)
		result, detailedResponse, err := service1.GetQuotaDefinition(getQuotaDefinitionOptionsModel)
		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))
		Expect(result).NotTo(BeNil())
	})

	Describe("Get a List of all resource groups in an account", func() {
		It("Successfully retrieved list of resource groups in an account", func() {
			shouldSkipTest()

			listResourceGroupsOptionsModel := service1.NewListResourceGroupsOptions()
			listResourceGroupsOptionsModel.SetAccountID(testUserAccountID)
			result, detailedResponse, err := service1.ListResourceGroups(listResourceGroupsOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(len(result.Resources)).To(BeNumerically(">=", 1))
			Expect(result.Resources[0]).NotTo(BeNil())
			Expect(result.Resources[0].ID).NotTo(BeNil())
			Expect(result.Resources[0].Name).NotTo(BeNil())
			Expect(result.Resources[0].Crn).NotTo(BeNil())
			Expect(result.Resources[0].AccountID).NotTo(BeNil())
			Expect(result.Resources[0].QuotaID).NotTo(BeNil())
			Expect(result.Resources[0].QuotaURL).NotTo(BeNil())
			Expect(result.Resources[0].CreatedAt).NotTo(BeNil())
			Expect(result.Resources[0].UpdatedAt).NotTo(BeNil())
		})
	})

	Describe("Create a new resource group in an account", func() {
		It("Successfully created new resource group in an account", func() {
			shouldSkipTest()

			createResourceGroupOptionsModel := service1.NewCreateResourceGroupOptions()
			createResourceGroupOptionsModel.SetAccountID(testUserAccountID)
			createResourceGroupOptionsModel.SetName("TestGroup")
			result, detailedResponse, err := service1.CreateResourceGroup(createResourceGroupOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(result).NotTo(BeNil())
			Expect(result.ID).NotTo(BeNil())
			newResourceGroupID = *result.ID
		})
	})

	Describe("Get a resource group by ID", func() {
		It("Successfully retrieved resource group by ID", func() {
			shouldSkipTest()

			getResourceGroupOptionsModel := service1.NewGetResourceGroupOptions(newResourceGroupID)
			result, detailedResponse, err := service1.GetResourceGroup(getResourceGroupOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).NotTo(BeNil())
		})
	})

	Describe("Update a resource group by ID", func() {
		It("Successfully updated resource group", func() {
			shouldSkipTest()

			updateResourceGroupOptionsModel := service1.NewUpdateResourceGroupOptions(newResourceGroupID)
			result, detailedResponse, err := service1.UpdateResourceGroup(updateResourceGroupOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).NotTo(BeNil())
		})
	})

	Describe("Delete a resource group by ID", func() {
		It("Successfully deleted resource group", func() {
			shouldSkipTest()

			deleteResourceGroupOptionsModel := service2.NewDeleteResourceGroupOptions(newResourceGroupID)
			detailedResponse, err := service2.DeleteResourceGroup(deleteResourceGroupOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(204))
		})
	})
})
