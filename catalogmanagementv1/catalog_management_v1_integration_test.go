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

package catalogmanagementv1_test

import (
	"fmt"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	"github.com/joho/godotenv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	externalConfigFile   = "../catalog_mgmt.env"
	expectedAccount      = "67d27f28d43948b2b3bda9138f251a13"
	expectedLabel        = "integration-test"
	expectedShortDesc    = "test"
	expectedURL          = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s"
	expectedOfferingsURL = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings"
	fakeName             = "bogus"
)

var (
	service      *catalogmanagementv1.CatalogManagementV1
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe("Catalog Management - Integration Tests", func() {
	It("Successfully load the configuration", func() {
		err := godotenv.Load(externalConfigFile)
		if err == nil {
			configLoaded = true
		} else {
			Skip("External configuration could not be loaded, skipping...")
		}
	})

	It(`Successfully created CatalogManagementV1 service instance`, func() {
		var err error

		shouldSkipTest()

		service, err = catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(
			&catalogmanagementv1.CatalogManagementV1Options{},
		)

		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())
	})

	Describe("Run integration tests", func() {
		JustBeforeEach(func() {
			shouldSkipTest()

			listResult, _, _ := service.ListCatalogs(service.NewListCatalogsOptions())
			if listResult != nil && listResult.Resources != nil {
				for _, resource := range listResult.Resources {
					service.DeleteCatalog(service.NewDeleteCatalogOptions(*resource.ID))
				}
			}
		})

		JustAfterEach(func() {
			shouldSkipTest()

			listResult, _, _ := service.ListCatalogs(service.NewListCatalogsOptions())
			if listResult != nil && listResult.Resources != nil {
				for _, resource := range listResult.Resources {
					service.DeleteCatalog(service.NewDeleteCatalogOptions(*resource.ID))
				}
			}
		})

		It("Get catalog account", func() {
			const expectedAccount = "67d27f28d43948b2b3bda9138f251a13"
			shouldSkipTest()

			options := service.NewGetCatalogAccountOptions()
			result, response, err := service.GetCatalogAccount(options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(expectedAccount))
			Expect(*result.AccountFilters.IncludeAll).To(BeTrue())
			Expect(len(result.AccountFilters.CategoryFilters)).To(BeZero())
			Expect(result.AccountFilters.IdFilters.Include).To(BeNil())
			Expect(result.AccountFilters.IdFilters.Exclude).To(BeNil())
		})

		It("Get catalog account filters", func() {
			shouldSkipTest()

			options := service.NewGetCatalogAccountFiltersOptions()
			result, response, err := service.GetCatalogAccountFilters(options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(len(result.AccountFilters)).To(Equal(1))
			Expect(*result.AccountFilters[0].IncludeAll).To(BeTrue())
			Expect(len(result.AccountFilters[0].CategoryFilters)).To(BeZero())
			Expect(result.AccountFilters[0].IdFilters.Include).To(BeNil())
			Expect(result.AccountFilters[0].IdFilters.Exclude).To(BeNil())
			Expect(len(result.CatalogFilters)).To(BeZero())
		})

		It("Get list of catalogs", func() {
			const (
				expectedTotalCount    int64 = 1
				expectedResourceCount       = 1
			)

			shouldSkipTest()

			createOptions := service.NewCreateCatalogOptions()
			createOptions.SetLabel(expectedLabel)
			createOptions.SetShortDescription(expectedShortDesc)
			createResult, _, _ := service.CreateCatalog(createOptions)

			listOptions := service.NewListCatalogsOptions()
			listResult, listResponse, err := service.ListCatalogs(listOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(*createResult.ID))

			Expect(err).To(BeNil())
			Expect(listResponse.StatusCode).To(Equal(200))
			Expect(*listResult.Offset).To(BeZero())
			Expect(*listResult.Limit).To(BeZero())
			Expect(*listResult.TotalCount).To(Equal(expectedTotalCount))
			Expect(listResult.Last).To(BeNil())
			Expect(listResult.Prev).To(BeNil())
			Expect(listResult.Next).To(BeNil())
			Expect(len(listResult.Resources)).To(Equal(expectedResourceCount))

			Expect(*listResult.Resources[0].Label).To(Equal(expectedLabel))
			Expect(*listResult.Resources[0].ShortDescription).To(Equal(expectedShortDesc))
			Expect(*listResult.Resources[0].URL).To(Equal(fmt.Sprintf(expectedURL, *createResult.ID)))
			Expect(*listResult.Resources[0].OfferingsURL).To(Equal(fmt.Sprintf(expectedOfferingsURL, *createResult.ID)))
			Expect(*listResult.Resources[0].OwningAccount).To(Equal(expectedAccount))
			Expect(*listResult.Resources[0].CatalogFilters.IncludeAll).To(BeFalse())
			Expect(len(listResult.Resources[0].CatalogFilters.CategoryFilters)).To(BeZero())
			Expect(listResult.Resources[0].CatalogFilters.IdFilters.Include).To(BeNil())
			Expect(listResult.Resources[0].CatalogFilters.IdFilters.Exclude).To(BeNil())
		})

		It("Create a catalog", func() {
			shouldSkipTest()

			options := service.NewCreateCatalogOptions()
			options.SetLabel(expectedLabel)
			options.SetShortDescription(expectedShortDesc)
			result, response, err := service.CreateCatalog(options)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(*result.ID))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(*result.Label).To(Equal(expectedLabel))
			Expect(*result.ShortDescription).To(Equal(expectedShortDesc))
			Expect(*result.URL).To(Equal(fmt.Sprintf(expectedURL, *result.ID)))
			Expect(*result.OfferingsURL).To(Equal(fmt.Sprintf(expectedOfferingsURL, *result.ID)))
			Expect(*result.OwningAccount).To(Equal(expectedAccount))
			Expect(*result.CatalogFilters.IncludeAll).To(BeFalse())
			Expect(len(result.CatalogFilters.CategoryFilters)).To(BeZero())
			Expect(result.CatalogFilters.IdFilters.Include).To(BeNil())
			Expect(result.CatalogFilters.IdFilters.Exclude).To(BeNil())
		})

		It("Get a catalog", func() {
			shouldSkipTest()

			createOptions := service.NewCreateCatalogOptions()
			createOptions.SetLabel(expectedLabel)
			createOptions.SetShortDescription(expectedShortDesc)
			createResult, _, _ := service.CreateCatalog(createOptions)

			id := *createResult.ID
			getOptions := service.NewGetCatalogOptions(id)
			getResult, getResponse, err := service.GetCatalog(getOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(id))

			Expect(err).To(BeNil())
			Expect(getResponse.StatusCode).To(Equal(200))
			Expect(*getResult.Label).To(Equal(expectedLabel))
			Expect(*getResult.ShortDescription).To(Equal(expectedShortDesc))
			Expect(*getResult.URL).To(Equal(fmt.Sprintf(expectedURL, id)))
			Expect(*getResult.OfferingsURL).To(Equal(fmt.Sprintf(expectedOfferingsURL, id)))
			Expect(*getResult.OwningAccount).To(Equal(expectedAccount))
			Expect(*getResult.CatalogFilters.IncludeAll).To(BeFalse())
			Expect(len(getResult.CatalogFilters.CategoryFilters)).To(BeZero())
			Expect(getResult.CatalogFilters.IdFilters.Include).To(BeNil())
			Expect(getResult.CatalogFilters.IdFilters.Exclude).To(BeNil())
		})

		It("Fail to get a catalog that does not exist", func() {
			shouldSkipTest()

			id := fakeName
			getOptions := service.NewGetCatalogOptions(id)
			_, getResponse, err := service.GetCatalog(getOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(id))

			Expect(err).ToNot(BeNil())
			Expect(getResponse.StatusCode).To(Equal(404))
		})

		It("Update a catalog", func() {
			const (
				expectedLabelUpdated     = "test2"
				expectedShortDescUpdated = "integration-test-update"
			)

			shouldSkipTest()

			createOptions := service.NewCreateCatalogOptions()
			createOptions.SetLabel(expectedLabel)
			createOptions.SetShortDescription(expectedShortDesc)
			createResult, _, _ := service.CreateCatalog(createOptions)

			id := *createResult.ID

			replaceOptions := service.NewReplaceCatalogOptions(id)
			replaceOptions.SetCatalogIdentifier(id)
			replaceOptions.SetID(id)
			replaceOptions.SetLabel(expectedLabelUpdated)
			replaceOptions.SetShortDescription(expectedShortDescUpdated)
			replaceResult, replaceResponse, err := service.ReplaceCatalog(replaceOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(id))

			Expect(err).To(BeNil())
			Expect(replaceResponse.StatusCode).To(Equal(200))
			Expect(*replaceResult.Label).To(Equal(expectedLabelUpdated))
			Expect(*replaceResult.ShortDescription).To(Equal(expectedShortDescUpdated))
			Expect(*replaceResult.URL).To(Equal(fmt.Sprintf(expectedURL, id)))
			Expect(*replaceResult.OfferingsURL).To(Equal(fmt.Sprintf(expectedOfferingsURL, id)))
			Expect(*replaceResult.OwningAccount).To(Equal(expectedAccount))
			Expect(*replaceResult.CatalogFilters.IncludeAll).To(BeTrue())
			Expect(len(replaceResult.CatalogFilters.CategoryFilters)).To(BeZero())
			Expect(replaceResult.CatalogFilters.IdFilters.Include).To(BeNil())
			Expect(replaceResult.CatalogFilters.IdFilters.Exclude).To(BeNil())
		})

		It("Fail to update a catalog that does not exist", func() {
			shouldSkipTest()

			id := fakeName
			replaceOptions := service.NewReplaceCatalogOptions(id)
			replaceOptions.SetCatalogIdentifier(id)
			replaceOptions.SetID(id)
			_, replaceResponse, err := service.ReplaceCatalog(replaceOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(id))

			Expect(err).ToNot(BeNil())
			Expect(replaceResponse.StatusCode).To(Equal(404))
		})

		It("Delete a catalog", func() {
			shouldSkipTest()

			createOptions := service.NewCreateCatalogOptions()
			createOptions.SetLabel(expectedLabel)
			createOptions.SetShortDescription(expectedShortDesc)
			createResult, _, _ := service.CreateCatalog(createOptions)

			id := *createResult.ID

			deleteResponse, deleteErr := service.DeleteCatalog(service.NewDeleteCatalogOptions(id))
			Expect(deleteErr).To(BeNil())
			Expect(deleteResponse.StatusCode).To(Equal(200))

			getOptions := service.NewGetCatalogOptions(id)
			_, getResponse, getErr := service.GetCatalog(getOptions)

			Expect(getErr).ToNot(BeNil())
			Expect(getResponse.StatusCode).To(Equal(403))
		})

		It("Fail to delete a catalog that does not exist", func() {
			shouldSkipTest()

			id := fakeName
			deleteResponse, deleteErr := service.DeleteCatalog(service.NewDeleteCatalogOptions(id))

			Expect(deleteErr).To(BeNil())
			Expect(deleteResponse.StatusCode).To(Equal(200))
		})

		It("Create an offering", func() {
			const (
				expectedName  = "test-offering"
				expectedLabel = "test"
				expectedURL   = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings/%s"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewCreateOfferingOptions(catalogID)
			offeringOptions.SetName(expectedName)
			offeringOptions.SetLabel(expectedLabel)
			offeringResult, offeringResponse, err := service.CreateOffering(offeringOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			offeringID := *offeringResult.ID

			Expect(err).To(BeNil())
			Expect(offeringResponse.StatusCode).To(Equal(201))
			Expect(*offeringResult.Name).To(Equal(expectedName))
			Expect(*offeringResult.URL).To(Equal(fmt.Sprintf(expectedURL, catalogID, offeringID)))
			Expect(*offeringResult.Label).To(Equal(expectedLabel))
		})

		It("Get an offering", func() {
			const (
				expectedName  = "test-offering"
				expectedLabel = "test"
				expectedURL   = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings/%s"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewCreateOfferingOptions(catalogID)
			offeringOptions.SetName(expectedName)
			offeringOptions.SetLabel(expectedLabel)
			offeringResult, _, _ := service.CreateOffering(offeringOptions)
			offeringID := *offeringResult.ID

			getOptions := service.NewGetOfferingOptions(catalogID, offeringID)
			getResult, getResponse, err := service.GetOffering(getOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).To(BeNil())
			Expect(getResponse.StatusCode).To(Equal(200))
			Expect(*getResult.Name).To(Equal(expectedName))
			Expect(*getResult.URL).To(Equal(fmt.Sprintf(expectedURL, catalogID, offeringID)))
			Expect(*getResult.Label).To(Equal(expectedLabel))
		})

		It("Fail to get an offering that does not exist", func() {
			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID
			offeringID := fakeName

			getOptions := service.NewGetOfferingOptions(catalogID, offeringID)
			_, getResponse, err := service.GetOffering(getOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).ToNot(BeNil())
			Expect(getResponse.StatusCode).To(Equal(404))

			_, getResponse, err = service.GetOffering(getOptions)

			Expect(err).ToNot(BeNil())
			Expect(getResponse.StatusCode).To(Equal(403))
		})

		It("List offerings", func() {
			const (
				expectedName  = "test-offering"
				expectedLabel = "test"
				expectedURL   = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings/%s"

				expectedLimit         int64 = 100
				expectedTotalCount    int64 = 1
				expectedResourceCount int64 = 1
				expectedResouceLen          = 1
				expectedFirst               = "/api/v1-beta/catalogs/%s/offerings?limit=100&sort=label"
				expectedLast                = "/api/v1-beta/catalogs/%s/offerings?limit=100&sort=label"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewCreateOfferingOptions(catalogID)
			offeringOptions.SetName(expectedName)
			offeringOptions.SetLabel(expectedLabel)
			offeringResult, _, _ := service.CreateOffering(offeringOptions)
			offeringID := *offeringResult.ID

			listOptions := service.NewListOfferingsOptions(catalogID)
			listResult, listResponse, err := service.ListOfferings(listOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).To(BeNil())
			Expect(listResponse.StatusCode).To(Equal(200))
			Expect(*listResult.Offset).To(BeZero())
			Expect(*listResult.Limit).To(Equal(expectedLimit))
			Expect(*listResult.TotalCount).To(Equal(expectedTotalCount))
			Expect(*listResult.ResourceCount).To(Equal(expectedResourceCount))
			Expect(*listResult.First).To(Equal(fmt.Sprintf(expectedFirst, catalogID)))
			Expect(*listResult.Last).To(Equal(fmt.Sprintf(expectedLast, catalogID)))
			Expect(len(listResult.Resources)).To(Equal(expectedResouceLen))

			Expect(*listResult.Resources[0].ID).To(Equal(offeringID))
			Expect(*listResult.Resources[0].URL).To(Equal(fmt.Sprintf(expectedURL, catalogID, offeringID)))
			Expect(*listResult.Resources[0].Label).To(Equal(expectedLabel))
			Expect(*listResult.Resources[0].Name).To(Equal(expectedName))
			Expect(*listResult.Resources[0].CatalogID).To(Equal(catalogID))
			Expect(*listResult.Resources[0].CatalogName).To(Equal(expectedLabel))

		})

		It("Delete an offering", func() {
			const (
				expectedName  = "test-offering"
				expectedLabel = "test"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewCreateOfferingOptions(catalogID)
			offeringOptions.SetName(expectedName)
			offeringOptions.SetLabel(expectedLabel)
			offeringResult, _, _ := service.CreateOffering(offeringOptions)
			offeringID := *offeringResult.ID

			deleteResponse, err := service.DeleteOffering(service.NewDeleteOfferingOptions(catalogID, offeringID))
			Expect(err).To(BeNil())
			Expect(deleteResponse.StatusCode).To(Equal(200))

			_, getResponse, err := service.GetOffering(service.NewGetOfferingOptions(catalogID, offeringID))

			Expect(err).ToNot(BeNil())
			Expect(getResponse.StatusCode).To(Equal(404))

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))
		})

		It("Fail to delete an offering that does not exist", func() {
			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID
			offeringID := fakeName

			deleteResponse, err := service.DeleteOffering(service.NewDeleteOfferingOptions(catalogID, offeringID))
			Expect(err).To(BeNil())
			Expect(deleteResponse.StatusCode).To(Equal(200))

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			deleteResponse, err = service.DeleteOffering(service.NewDeleteOfferingOptions(catalogID, offeringID))
			Expect(err).ToNot(BeNil())
			Expect(deleteResponse.StatusCode).To(Equal(403))
		})

		It("Update an offering", func() {
			const (
				expectedName            = "test-offering"
				expectedLabel           = "test"
				expectedLabelUpdate     = "test-update"
				expectedShortDesc       = "test-desc"
				expectedShortDescUpdate = "test-desc-update"
				expectedURL             = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings/%s"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewCreateOfferingOptions(catalogID)
			offeringOptions.SetName(expectedName)
			offeringOptions.SetLabel(expectedLabel)
			offeringOptions.SetShortDescription(expectedShortDesc)
			offeringResult, _, _ := service.CreateOffering(offeringOptions)
			offeringID := *offeringResult.ID
			rev := *offeringResult.Rev

			updateOptions := service.NewReplaceOfferingOptions(catalogID, offeringID)
			updateOptions.SetID(offeringID)
			updateOptions.SetLabel(expectedLabelUpdate)
			updateOptions.SetShortDescription(expectedShortDescUpdate)
			updateOptions.SetRev(rev)
			updateResult, updateResponse, err := service.ReplaceOffering(updateOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).To(BeNil())
			Expect(updateResponse.StatusCode).To(Equal(200))
			Expect(*updateResult.ShortDescription).To(Equal(expectedShortDescUpdate))
			Expect(*updateResult.URL).To(Equal(fmt.Sprintf(expectedURL, catalogID, offeringID)))
			Expect(*updateResult.Label).To(Equal(expectedLabelUpdate))
		})

		It("Fail to update an offering that does not exist", func() {
			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID
			offeringID := fakeName
			rev := fakeName

			updateOptions := service.NewReplaceOfferingOptions(catalogID, offeringID)
			updateOptions.SetID(offeringID)
			updateOptions.SetRev(rev)
			_, updateResponse, err := service.ReplaceOffering(updateOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).ToNot(BeNil())
			Expect(updateResponse.StatusCode).To(Equal(404))

			_, updateResponse, err = service.ReplaceOffering(updateOptions)
			Expect(err).ToNot(BeNil())
			Expect(updateResponse.StatusCode).To(Equal(403))
		})

		It("Get list of offerings for consumption", func() {
			shouldSkipTest()

			options := service.NewGetConsumptionOfferingsOptions()
			result, response, err := service.GetConsumptionOfferings(options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(*result.Offset).To(BeZero())
			Expect(*result.Limit).To(BeZero())
			Expect(*result.TotalCount).ToNot(BeZero())
			Expect(result.Last).To(BeNil())
			Expect(result.Prev).To(BeNil())
			Expect(result.Next).To(BeNil())
			Expect(len(result.Resources)).ToNot(BeZero())
		})

		It("Import an offering", func() {
			const (
				expectedOfferingName       = "jenkins-operator"
				expectedOfferingLabel      = "Jenkins Operator"
				expectedOfferingTargetKind = "roks"
				expectedOfferingVersion    = "0.4.0"
				expectedOfferingVersions   = 1
				expectedOfferingKinds      = 1
				expectedOfferingShortDesc  = "Kubernetes native operator which fully manages Jenkins on Openshift."
				expectedOfferingURL        = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings/%s"
				expectedOfferingZipURL     = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.4.0/jenkins-operator.v0.4.0.clusterserviceversion.yaml"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewImportOfferingOptions(catalogID, expectedOfferingZipURL)
			offeringResult, offeringResponse, err := service.ImportOffering(offeringOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			offeringID := *offeringResult.ID

			Expect(err).To(BeNil())
			Expect(offeringResponse.StatusCode).To(Equal(201))
			Expect(*offeringResult.Name).To(Equal(expectedOfferingName))
			Expect(*offeringResult.URL).To(Equal(fmt.Sprintf(expectedOfferingURL, catalogID, offeringID)))
			Expect(*offeringResult.Label).To(Equal(expectedOfferingLabel))
			Expect(*offeringResult.ShortDescription).To(Equal(expectedOfferingShortDesc))
			Expect(*offeringResult.CatalogName).To(Equal(expectedLabel))
			Expect(*offeringResult.CatalogID).To(Equal(catalogID))
			Expect(len(offeringResult.Kinds)).To(Equal(expectedOfferingKinds))
			Expect(*offeringResult.Kinds[0].TargetKind).To(Equal(expectedOfferingTargetKind))
			Expect(len(offeringResult.Kinds[0].Versions)).To(Equal(expectedOfferingVersions))
			Expect(*offeringResult.Kinds[0].Versions[0].Version).To(Equal(expectedOfferingVersion))
			Expect(*offeringResult.Kinds[0].Versions[0].TgzURL).To(Equal(expectedOfferingZipURL))
		})

		It("Import new version to offering", func() {
			const (
				expectedOfferingName         = "jenkins-operator"
				expectedOfferingLabel        = "Jenkins Operator"
				expectedOfferingTargetKind   = "roks"
				expectedOfferingKinds        = 1
				expectedOfferingVersions     = 2
				expectedOfferingVersion1     = "0.3.31"
				expectedOfferingVersion2     = "0.4.0"
				expectedOfferingShortDesc    = "Kubernetes native operator which fully manages Jenkins on Openshift."
				expectedOfferingURL          = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings/%s"
				expectedOfferingZipURL       = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.3.31/jenkins-operator.v0.3.31.clusterserviceversion.yaml"
				expectedOfferingZipURLUpdate = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.4.0/jenkins-operator.v0.4.0.clusterserviceversion.yaml"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewImportOfferingOptions(catalogID, expectedOfferingZipURL)
			offeringResult, _, _ := service.ImportOffering(offeringOptions)

			offeringID := *offeringResult.ID

			importOptions := service.NewImportOfferingVersionOptions(catalogID, offeringID, expectedOfferingZipURLUpdate)
			importResult, importResponse, err := service.ImportOfferingVersion(importOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).To(BeNil())
			Expect(importResponse.StatusCode).To(Equal(201))
			Expect(*importResult.Name).To(Equal(expectedOfferingName))
			Expect(*importResult.URL).To(Equal(fmt.Sprintf(expectedOfferingURL, catalogID, offeringID)))
			Expect(*importResult.Label).To(Equal(expectedOfferingLabel))
			Expect(*importResult.ShortDescription).To(Equal(expectedOfferingShortDesc))
			Expect(*importResult.CatalogName).To(Equal(expectedLabel))
			Expect(*importResult.CatalogID).To(Equal(catalogID))
			Expect(len(importResult.Kinds)).To(Equal(expectedOfferingKinds))
			Expect(*importResult.Kinds[0].TargetKind).To(Equal(expectedOfferingTargetKind))
			Expect(len(importResult.Kinds[0].Versions)).To(Equal(expectedOfferingVersions))
			Expect(*importResult.Kinds[0].Versions[0].Version).To(Equal(expectedOfferingVersion1))
			Expect(*importResult.Kinds[0].Versions[0].TgzURL).To(Equal(expectedOfferingZipURL))
			Expect(*importResult.Kinds[0].Versions[1].Version).To(Equal(expectedOfferingVersion2))
			Expect(*importResult.Kinds[0].Versions[1].TgzURL).To(Equal(expectedOfferingZipURLUpdate))
		})

		It("Fail to import new version to offering that does not exist", func() {
			const expectedOfferingZipURLUpdate = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.4.0/jenkins-operator.v0.4.0.clusterserviceversion.yaml"

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringID := fakeName

			importOptions := service.NewImportOfferingVersionOptions(catalogID, offeringID, expectedOfferingZipURLUpdate)
			_, importResponse, err := service.ImportOfferingVersion(importOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).ToNot(BeNil())
			Expect(importResponse.StatusCode).To(Equal(404))

			_, importResponse, err = service.ImportOfferingVersion(importOptions)
			Expect(err).ToNot(BeNil())
			Expect(importResponse.StatusCode).To(Equal(403))
		})

		It("Reload an offering", func() {
			const (
				expectedOfferingName       = "jenkins-operator"
				expectedOfferingLabel      = "Jenkins Operator"
				expectedOfferingTargetKind = "roks"
				expectedOfferingVersion    = "0.4.0"
				expectedOfferingVersions   = 1
				expectedOfferingKinds      = 1
				expectedOfferingShortDesc  = "Kubernetes native operator which fully manages Jenkins on Openshift."
				expectedOfferingURL        = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings/%s"
				expectedOfferingZipURL     = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.4.0/jenkins-operator.v0.4.0.clusterserviceversion.yaml"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewImportOfferingOptions(catalogID, expectedOfferingZipURL)
			offeringResult, _, _ := service.ImportOffering(offeringOptions)
			offeringID := *offeringResult.ID

			reloadOptions := service.NewReloadOfferingOptions(catalogID, offeringID, expectedOfferingZipURL, expectedOfferingVersion)
			reloadResult, reloadResponse, err := service.ReloadOffering(reloadOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).To(BeNil())
			Expect(reloadResponse.StatusCode).To(Equal(200))
			Expect(*reloadResult.Name).To(Equal(expectedOfferingName))
			Expect(*reloadResult.URL).To(Equal(fmt.Sprintf(expectedOfferingURL, catalogID, offeringID)))
			Expect(*reloadResult.Label).To(Equal(expectedOfferingLabel))
			Expect(*reloadResult.ShortDescription).To(Equal(expectedOfferingShortDesc))
			Expect(*reloadResult.CatalogName).To(Equal(expectedLabel))
			Expect(*reloadResult.CatalogID).To(Equal(catalogID))
			Expect(len(reloadResult.Kinds)).To(Equal(expectedOfferingKinds))
			Expect(*reloadResult.Kinds[0].TargetKind).To(Equal(expectedOfferingTargetKind))
			Expect(len(reloadResult.Kinds[0].Versions)).To(Equal(expectedOfferingVersions))
			Expect(*reloadResult.Kinds[0].Versions[0].Version).To(Equal(expectedOfferingVersion))
			Expect(*reloadResult.Kinds[0].Versions[0].TgzURL).To(Equal(expectedOfferingZipURL))
		})

		It("Fail to reload an offering that does not exist", func() {
			const (
				expectedOfferingVersion = "0.4.0"
				expectedOfferingZipURL  = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.4.0/jenkins-operator.v0.4.0.clusterserviceversion.yaml"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID
			offeringID := fakeName

			reloadOptions := service.NewReloadOfferingOptions(catalogID, offeringID, expectedOfferingZipURL, expectedOfferingVersion)
			_, reloadResponse, err := service.ReloadOffering(reloadOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).ToNot(BeNil())
			Expect(reloadResponse.StatusCode).To(Equal(404))

			_, reloadResponse, err = service.ReloadOffering(reloadOptions)
			Expect(err).ToNot(BeNil())
			Expect(reloadResponse.StatusCode).To(Equal(403))
		})

		It("Get a version", func() {
			const (
				expectedOfferingName       = "jenkins-operator"
				expectedOfferingLabel      = "Jenkins Operator"
				expectedOfferingTargetKind = "roks"
				expectedOfferingVersion    = "0.4.0"
				expectedOfferingVersions   = 1
				expectedOfferingKinds      = 1
				expectedOfferingShortDesc  = "Kubernetes native operator which fully manages Jenkins on Openshift."
				expectedOfferingURL        = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings/%s"
				expectedOfferingZipURL     = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.4.0/jenkins-operator.v0.4.0.clusterserviceversion.yaml"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewImportOfferingOptions(catalogID, expectedOfferingZipURL)
			offeringResult, _, _ := service.ImportOffering(offeringOptions)
			offeringID := *offeringResult.ID
			versionLocator := *offeringResult.Kinds[0].Versions[0].VersionLocator

			versionOptions := service.NewGetVersionOptions(versionLocator)
			versionResult, versionResponse, err := service.GetVersion(versionOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).To(BeNil())
			Expect(versionResponse.StatusCode).To(Equal(200))
			Expect(*versionResult.Name).To(Equal(expectedOfferingName))
			Expect(*versionResult.URL).To(Equal(fmt.Sprintf(expectedOfferingURL, catalogID, offeringID)))
			Expect(*versionResult.Label).To(Equal(expectedOfferingLabel))
			Expect(*versionResult.ShortDescription).To(Equal(expectedOfferingShortDesc))
			Expect(*versionResult.CatalogName).To(Equal(expectedLabel))
			Expect(*versionResult.CatalogID).To(Equal(catalogID))
			Expect(len(versionResult.Kinds)).To(Equal(expectedOfferingKinds))
			Expect(*versionResult.Kinds[0].TargetKind).To(Equal(expectedOfferingTargetKind))
			Expect(len(versionResult.Kinds[0].Versions)).To(Equal(expectedOfferingVersions))
			Expect(*versionResult.Kinds[0].Versions[0].Version).To(Equal(expectedOfferingVersion))
			Expect(*versionResult.Kinds[0].Versions[0].TgzURL).To(Equal(expectedOfferingZipURL))
		})

		It("Fail to get a version that does not exist", func() {
			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID
			versionLocator := fakeName

			versionOptions := service.NewGetVersionOptions(versionLocator)
			_, versionResponse, err := service.GetVersion(versionOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).ToNot(BeNil())
			Expect(versionResponse.StatusCode).To(Equal(400))

			_, versionResponse, err = service.GetVersion(versionOptions)
			Expect(err).ToNot(BeNil())
			Expect(versionResponse.StatusCode).To(Equal(400))

		})

		It("Delete a version", func() {
			const expectedOfferingZipURL = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.4.0/jenkins-operator.v0.4.0.clusterserviceversion.yaml"

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewImportOfferingOptions(catalogID, expectedOfferingZipURL)
			offeringResult, _, _ := service.ImportOffering(offeringOptions)
			versionLocator := *offeringResult.Kinds[0].Versions[0].VersionLocator

			deleteOptions := service.NewDeleteVersionOptions(versionLocator)
			deleteResponse, err := service.DeleteVersion(deleteOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).To(BeNil())
			Expect(deleteResponse.StatusCode).To(Equal(200))
		})

		It("Failed to delete a version that does not exist", func() {
			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID
			versionLocator := fakeName

			deleteOptions := service.NewDeleteVersionOptions(versionLocator)
			deleteResponse, err := service.DeleteVersion(deleteOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).ToNot(BeNil())
			Expect(deleteResponse.StatusCode).To(Equal(400))

			deleteResponse, err = service.DeleteVersion(deleteOptions)
			Expect(err).ToNot(BeNil())
			Expect(deleteResponse.StatusCode).To(Equal(400))
		})

		It("Get version about", func() {
			const expectedOfferingZipURL = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.4.0/jenkins-operator.v0.4.0.clusterserviceversion.yaml"

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewImportOfferingOptions(catalogID, expectedOfferingZipURL)
			offeringResult, _, _ := service.ImportOffering(offeringOptions)
			versionLocator := *offeringResult.Kinds[0].Versions[0].VersionLocator

			getOptions := service.NewGetVersionAboutOptions(versionLocator)
			getResult, getResponse, err := service.GetVersionAbout(getOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).To(BeNil())
			Expect(getResponse.StatusCode).To(Equal(200))
			Expect(len(*getResult)).ToNot(BeZero())
		})

		It("Fail to get version about for a version that does not exist", func() {
			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID
			versionLocator := fakeName

			getOptions := service.NewGetVersionAboutOptions(versionLocator)
			_, getResponse, err := service.GetVersionAbout(getOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).ToNot(BeNil())
			Expect(getResponse.StatusCode).To(Equal(400))

			_, getResponse, err = service.GetVersionAbout(getOptions)

			Expect(err).ToNot(BeNil())
			Expect(getResponse.StatusCode).To(Equal(400))
		})

		It("Get version updates", func() {
			const (
				expectedOfferingUpdates      = 1
				expectedOfferingVersion2     = "0.4.0"
				expectedOfferingZipURL       = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.3.31/jenkins-operator.v0.3.31.clusterserviceversion.yaml"
				expectedOfferingZipURLUpdate = "https://github.com/operator-framework/community-operators/blob/master/community-operators/jenkins-operator/0.4.0/jenkins-operator.v0.4.0.clusterserviceversion.yaml"
			)

			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID

			offeringOptions := service.NewImportOfferingOptions(catalogID, expectedOfferingZipURL)
			offeringResult, _, _ := service.ImportOffering(offeringOptions)

			offeringID := *offeringResult.ID
			versionLocator1 := *offeringResult.Kinds[0].Versions[0].VersionLocator

			importOptions := service.NewImportOfferingVersionOptions(catalogID, offeringID, expectedOfferingZipURLUpdate)
			importResult, _, _ := service.ImportOfferingVersion(importOptions)
			versionLocator2 := *importResult.Kinds[0].Versions[1].VersionLocator

			getOptions := service.NewGetVersionUpdatesOptions(versionLocator1)
			getResult, getResponse, err := service.GetVersionUpdates(getOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).To(BeNil())
			Expect(getResponse.StatusCode).To(Equal(200))
			Expect(len(getResult)).To(Equal(expectedOfferingUpdates))
			Expect(*getResult[0].VersionLocator).To(Equal(versionLocator2))
			Expect(*getResult[0].Version).To(Equal(expectedOfferingVersion2))
			Expect(*getResult[0].PackageVersion).To(Equal(expectedOfferingVersion2))
			Expect(*getResult[0].CanUpdate).To(BeTrue())

		})
		It("Fail to get version updates for version that does not exist", func() {
			shouldSkipTest()

			catalogOptions := service.NewCreateCatalogOptions()
			catalogOptions.SetLabel(expectedLabel)
			catalogResult, _, _ := service.CreateCatalog(catalogOptions)
			catalogID := *catalogResult.ID
			versionLocator := fakeName

			getOptions := service.NewGetVersionUpdatesOptions(versionLocator)
			_, getResponse, err := service.GetVersionUpdates(getOptions)

			service.DeleteCatalog(service.NewDeleteCatalogOptions(catalogID))

			Expect(err).ToNot(BeNil())
			Expect(getResponse.StatusCode).To(Equal(400))

			_, getResponse, err = service.GetVersionUpdates(getOptions)

			Expect(err).ToNot(BeNil())
			Expect(getResponse.StatusCode).To(Equal(400))
		})

		It("Get license providers", func() {
			const (
				expectedResourceCount       = 1
				expectedTotalResults  int64 = 1
				expectedTotalPages    int64 = 1
				expectedName                = "IBM Passport Advantage"
				expectedOfferingType        = "content"
				expectedCreateURL           = "https://www.ibm.com/software/passportadvantage/aboutpassport.html"
				expectedInfoURL             = "https://www.ibm.com/software/passportadvantage/"
				expectedURL                 = "/v1/licensing/license_providers/11cabc37-c4a7-410b-894d-8cb3586423f1"
				expectedState               = "active"
			)

			shouldSkipTest()

			getOptions := service.NewGetLicenseProvidersOptions()
			getResult, getResponse, err := service.GetLicenseProviders(getOptions)

			Expect(err).To(BeNil())
			Expect(getResponse.StatusCode).To(Equal(200))
			Expect(len(getResult.Resources)).To(Equal(expectedResourceCount))
			Expect(*getResult.TotalResults).To(Equal(expectedTotalResults))
			Expect(*getResult.TotalPages).To(Equal(expectedTotalPages))
			Expect(*getResult.Resources[0].Name).To(Equal(expectedName))
			Expect(*getResult.Resources[0].OfferingType).To(Equal(expectedOfferingType))
			Expect(*getResult.Resources[0].CreateURL).To(Equal(expectedCreateURL))
			Expect(*getResult.Resources[0].InfoURL).To(Equal(expectedInfoURL))
			Expect(*getResult.Resources[0].URL).To(Equal(expectedURL))
			Expect(*getResult.Resources[0].State).To(Equal(expectedState))
		})

		It("Get license entitlements", func() {
			const (
				expectedResourceCount       = 0
				expectedTotalResults  int64 = 0
				expectedTotalPages    int64 = 1
			)

			shouldSkipTest()

			listOptions := service.NewListLicenseEntitlementsOptions()
			listResult, listResponse, err := service.ListLicenseEntitlements(listOptions)

			Expect(err).To(BeNil())
			Expect(listResponse.StatusCode).To(Equal(200))
			Expect(len(listResult.Resources)).To(Equal(expectedResourceCount))
			Expect(*listResult.TotalResults).To(Equal(expectedTotalResults))
			Expect(*listResult.TotalPages).To(Equal(expectedTotalPages))
		})

		It("Fail to search license versions", func() {
			shouldSkipTest()

			searchOptions := service.NewSearchLicenseVersionsOptions("")
			searchResponse, err := service.SearchLicenseVersions(searchOptions)

			Expect(err).ToNot(BeNil())
			Expect(searchResponse.StatusCode).To(Equal(403))
		})

		It("Fail to search license offerings", func() {
			shouldSkipTest()

			searchOptions := service.NewSearchLicenseOfferingsOptions("")
			searchResponse, err := service.SearchLicenseOfferings(searchOptions)

			Expect(err).ToNot(BeNil())
			Expect(searchResponse.StatusCode).To(Equal(403))
		})
	})
})
