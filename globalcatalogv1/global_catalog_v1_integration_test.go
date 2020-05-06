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

package globalcatalogv1_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	externalConfigFile    = "../global_catalog.env"
	visibilityRestriction = "private"
	artifact              = `{"someKey": "someValue"}`
)

var (
	service          *globalcatalogv1.GlobalCatalogV1
	defaultCreate    *globalcatalogv1.CreateCatalogEntryOptions
	defaultDelete    *globalcatalogv1.DeleteCatalogEntryOptions
	defaultGet       *globalcatalogv1.GetCatalogEntryOptions
	defaultUpdate    *globalcatalogv1.UpdateCatalogEntryOptions
	defaultList      *globalcatalogv1.ListCatalogEntriesOptions
	defaultChild     *globalcatalogv1.CreateCatalogEntryOptions
	getChild         *globalcatalogv1.GetChildObjectsOptions
	deleteChild      *globalcatalogv1.DeleteCatalogEntryOptions
	defaultRestore   *globalcatalogv1.RestoreCatalogEntryOptions
	bogusRestore     *globalcatalogv1.RestoreCatalogEntryOptions
	getVisibility    *globalcatalogv1.GetVisibilityOptions
	updateVisibility *globalcatalogv1.UpdateVisibilityOptions
	getPricing       *globalcatalogv1.GetPricingOptions
	defaultArtifact  *globalcatalogv1.UploadArtifactOptions
	listArtifacts    *globalcatalogv1.ListArtifactsOptions
	getArtifact      *globalcatalogv1.GetArtifactOptions
	deleteArtifact   *globalcatalogv1.DeleteArtifactOptions
	configLoaded     bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe("Global Catalog - Integration Tests", func() {
	It("Successfully load the configuration", func() {

		err := godotenv.Load(externalConfigFile)
		if err == nil {
			configLoaded = true
		} else {
			Skip("External configuration could not be loaded, skipping...")
		}
	})

	It(`Successfully created GlobalCatalogV1 service instance`, func() {
		const (
			kind                   = "service"
			disabled               = false
			email                  = "bogus@us.ibm.com"
			displayName            = "display"
			displayLongDesc        = "long"
			displayDesc            = "desc"
			displayNameUpdated     = "displayUpdated"
			displayLongDescUpdated = "longUpdated"
			displayDescUpdated     = "descUpdated"
			en                     = "en"
			providerName           = "someName"
			providerNameUpdated    = "someNameUpdated"
			artifactId             = "someArtifactId.json"
		)

		var (
			id                  = fmt.Sprintf("someId%d", time.Now().Unix())
			idChild             = fmt.Sprintf("someChildId%d", time.Now().Unix())
			name                = fmt.Sprintf("someName%d", time.Now().Unix())
			nameChild           = fmt.Sprintf("someChildName%d", time.Now().Unix())
			nameUpdated         = fmt.Sprintf("someNameUpdated%d", time.Now().Unix())
			image               = "image"
			imageUpdated        = "image"
			smallImage          = "small"
			smallImageUpdated   = "small"
			mediumImage         = "medium"
			mediumImageUpdated  = "medium"
			featureImage        = "feature"
			featureImageUpdated = "feature"
			tags                = []string{"a", "b", "c"}
			tagsUpdated         = []string{"x", "y", "z"}
			overviewUi          = &globalcatalogv1.OverviewUI{}
			overviewUiUpdated   = &globalcatalogv1.OverviewUI{}
			overview, _         = service.NewOverview(displayName, displayLongDesc, displayDesc)
			overviewUpdated, _  = service.NewOverview(displayNameUpdated, displayLongDescUpdated, displayDescUpdated)
			images              = &globalcatalogv1.Image{Image: &image,
				SmallImage:   &smallImage,
				MediumImage:  &mediumImage,
				FeatureImage: &featureImage,
			}
			imagesUpdated = &globalcatalogv1.Image{Image: &imageUpdated,
				SmallImage:   &smallImageUpdated,
				MediumImage:  &mediumImageUpdated,
				FeatureImage: &featureImageUpdated,
			}
			provider, _        = service.NewProvider(email, providerName)
			providerUpdated, _ = service.NewProvider(email, providerNameUpdated)
			err                error
		)

		shouldSkipTest()

		service, err = globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(
			&globalcatalogv1.GlobalCatalogV1Options{},
		)

		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())

		overviewUi.SetProperty(en, overview)
		overviewUiUpdated.SetProperty(en, overviewUpdated)

		defaultCreate = service.NewCreateCatalogEntryOptions(name,
			kind,
			overviewUi,
			images,
			disabled,
			tags,
			provider,
			id)
		defaultDelete = service.NewDeleteCatalogEntryOptions(id)
		defaultGet = service.NewGetCatalogEntryOptions(id)
		defaultUpdate = service.NewUpdateCatalogEntryOptions(id,
			nameUpdated,
			kind,
			overviewUiUpdated,
			imagesUpdated,
			disabled,
			tagsUpdated,
			providerUpdated)
		defaultList = service.NewListCatalogEntriesOptions()
		defaultChild = service.NewCreateCatalogEntryOptions(nameChild,
			kind,
			overviewUi,
			images,
			disabled,
			tags,
			provider,
			idChild)
		getChild = service.NewGetChildObjectsOptions(idChild, kind)
		defaultRestore = service.NewRestoreCatalogEntryOptions(id)
		bogusRestore = service.NewRestoreCatalogEntryOptions("bogus")
		getVisibility = service.NewGetVisibilityOptions(id)
		updateVisibility = service.NewUpdateVisibilityOptions(id)
		getPricing = service.NewGetPricingOptions(id)
		defaultArtifact = service.NewUploadArtifactOptions(id, artifactId)
		listArtifacts = service.NewListArtifactsOptions(id)
		getArtifact = service.NewGetArtifactOptions(id, artifactId)
		deleteArtifact = service.NewDeleteArtifactOptions(id, artifactId)
		deleteChild = service.NewDeleteCatalogEntryOptions(idChild)

		defaultChild.SetParentID(id)
		defaultArtifact.SetArtifact(ioutil.NopCloser(strings.NewReader(artifact)))
	})

	Describe("Run integration tests", func() {
		JustBeforeEach(func() {
			shouldSkipTest()

			service.DeleteCatalogEntry(defaultDelete)
			service.DeleteCatalogEntry(deleteChild)
		})

		JustAfterEach(func() {
			shouldSkipTest()

			service.DeleteCatalogEntry(defaultDelete)
			service.DeleteCatalogEntry(deleteChild)
		})

		It("Create a catalog entry", func() {
			shouldSkipTest()

			result, detailedResponse, err := service.CreateCatalogEntry(defaultCreate)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(result.ID).To(Equal(defaultCreate.ID))
			Expect(result.Name).To(Equal(defaultCreate.Name))
			Expect(result.Kind).To(Equal(defaultCreate.Kind))
			Expect(result.Images).To(Equal(defaultCreate.Images))
			Expect(result.Tags).To(Equal(defaultCreate.Tags))
			Expect(result.Provider).To(Equal(defaultCreate.Provider))
		})

		It("Get a catalog entry", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)
			result, detailedResponse, err := service.GetCatalogEntry(defaultGet)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.ID).To(Equal(defaultCreate.ID))
			Expect(result.Name).To(Equal(defaultCreate.Name))
			Expect(result.Kind).To(Equal(defaultCreate.Kind))
			Expect(result.Images).To(Equal(defaultCreate.Images))
			Expect(result.Tags).To(Equal(defaultCreate.Tags))
			Expect(result.Provider).To(Equal(defaultCreate.Provider))
		})

		It("Update a catalog entry", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)
			result, detailedResponse, err := service.UpdateCatalogEntry(defaultUpdate)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.ID).To(Equal(defaultUpdate.ID))
			Expect(result.Name).To(Equal(defaultUpdate.Name))
			Expect(result.Kind).To(Equal(defaultUpdate.Kind))
			Expect(result.Images).To(Equal(defaultUpdate.Images))
			Expect(result.Tags).To(Equal(defaultUpdate.Tags))
			Expect(result.Provider).To(Equal(defaultUpdate.Provider))
		})

		It("Delete a catalog entry", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)
			detailedResponse, err := service.DeleteCatalogEntry(defaultDelete)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})

		It("Fail to get a catalog entry after deletion", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)
			service.DeleteCatalogEntry(defaultDelete)

			_, detailedResponse, err := service.GetCatalogEntry(defaultGet)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

		It("Fail to get a catalog entry that does not exist", func() {
			shouldSkipTest()

			_, detailedResponse, err := service.GetCatalogEntry(defaultGet)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

		It("Fail to delete a catalog entry that does not exist", func() {
			shouldSkipTest()

			detailedResponse, err := service.DeleteCatalogEntry(defaultDelete)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})

		It("Fail to update a catalog entry that does not exist", func() {
			shouldSkipTest()

			_, detailedResponse, err := service.UpdateCatalogEntry(defaultUpdate)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

		It("Fail to create a catalog entry that already exists", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)

			_, detailedResponse, err := service.CreateCatalogEntry(defaultCreate)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(409))
		})

		It("List catalog entries", func() {
			shouldSkipTest()

			result, detailedResponse, err := service.ListCatalogEntries(defaultList)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.Resources).NotTo(BeNil())
			Expect(len(result.Resources)).NotTo(BeZero())
		})

		It("Get child catalog entry", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)
			service.CreateCatalogEntry(defaultChild)
			result, detailedResponse, err := service.GetChildObjects(getChild)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Resources[0].Name).To(Equal(defaultUpdate.Name))
			Expect(result.Resources[0].Kind).To(Equal(defaultUpdate.Kind))
			Expect(result.Resources[0].Images).To(Equal(defaultUpdate.Images))
			Expect(result.Resources[0].Tags).To(Equal(defaultUpdate.Tags))
			Expect(result.Resources[0].Provider).To(Equal(defaultUpdate.Provider))
		})

		It("Fail to get a child catalog entry that does not exist", func() {
			shouldSkipTest()

			_, detailedResponse, err := service.GetChildObjects(getChild)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

		It("Restore a catalog entry", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)
			service.DeleteCatalogEntry(defaultDelete)

			detailedResponseRestore, errRestore := service.RestoreCatalogEntry(defaultRestore)
			Expect(errRestore).To(BeNil())
			Expect(detailedResponseRestore.StatusCode).To(Equal(200))

			result, detailedResponseGet, errGet := service.GetCatalogEntry(defaultGet)
			Expect(errGet).To(BeNil())
			Expect(detailedResponseGet.StatusCode).To(Equal(200))
			Expect(result.ID).To(Equal(defaultCreate.ID))
			Expect(result.Name).To(Equal(defaultCreate.Name))
			Expect(result.Kind).To(Equal(defaultCreate.Kind))
			Expect(result.Images).To(Equal(defaultCreate.Images))
			Expect(result.Tags).To(Equal(defaultCreate.Tags))
			Expect(result.Provider).To(Equal(defaultCreate.Provider))
		})

		It("Fail to restore catalog entry that never existed", func() {
			shouldSkipTest()

			detailedResponse, err := service.RestoreCatalogEntry(bogusRestore)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

		It("Get visibility for catalog entry", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)

			result, detailedResponse, err := service.GetVisibility(getVisibility)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.Restrictions).To(Equal(visibilityRestriction))
		})

		It("Fail to get visibility of catalog entry that does not exist", func() {
			shouldSkipTest()

			_, detailedResponse, err := service.GetVisibility(getVisibility)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

		It("Update visibility for catalog entry", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)

			detailedResponse, err := service.UpdateVisibility(updateVisibility)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(403))
		})

		It("Fail to update visibility for catalog entry that does not exist", func() {
			shouldSkipTest()

			detailedResponse, err := service.UpdateVisibility(updateVisibility)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

		It("Fail to get pricing", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)

			_, detailedResponseExists, errExists := service.GetPricing(getPricing)
			Expect(errExists).NotTo(BeNil())
			Expect(detailedResponseExists.StatusCode).To(Equal(404))

			service.DeleteCatalogEntry(defaultDelete)

			_, detailedResponseNotExists, errNotExists := service.GetPricing(getPricing)
			Expect(errNotExists).NotTo(BeNil())
			Expect(detailedResponseNotExists.StatusCode).To(Equal(404))
		})

		It("List artifacts for a catalog entry", func() {
			const expectedCount int64 = 1
			const expectedSize int64 = 24

			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)
			service.UploadArtifact(defaultArtifact)

			result, detailedResponse, err := service.ListArtifacts(listArtifacts)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.Count).To(Equal(expectedCount))
			Expect(len(result.Resources)).To(Equal(1))
			Expect(result.Resources[0].Name).To(Equal(defaultArtifact.ArtifactID))
			Expect(*result.Resources[0].URL).To(Equal(fmt.Sprintf("%s/%s/artifacts/%s", "https://globalcatalog.test.cloud.ibm.com/api/v1", *defaultCreate.ID, *defaultArtifact.ArtifactID)))
			Expect(*result.Resources[0].Size).To(Equal(expectedSize))
		})

		It("Fail to list artifacts for a catalog entry that does not exist", func() {
			const expectedCount int64 = 0

			shouldSkipTest()

			result, detailedResponse, err := service.ListArtifacts(listArtifacts)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.Count).To(Equal(expectedCount))
		})

		It("Get artifact for a catalog entry", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)
			service.UploadArtifact(defaultArtifact)

			result, detailedResponse, err := service.GetArtifact(getArtifact)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(detailedResponse.Result).To(Equal(result))
		})

		It("Fail to get artifacts that do not exists", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)

			_, detailedResponseExists, errExists := service.GetArtifact(getArtifact)
			Expect(errExists).NotTo(BeNil())
			Expect(detailedResponseExists.StatusCode).To(Equal(404))

			service.DeleteCatalogEntry(defaultDelete)

			_, detailedResponseNotExists, errNotExists := service.GetArtifact(getArtifact)
			Expect(errNotExists).NotTo(BeNil())
			Expect(detailedResponseNotExists.StatusCode).To(Equal(404))
		})

		It("Create artifact for a catalog entry", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)

			detailedResponse, err := service.UploadArtifact(defaultArtifact)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})

		It("Fail to artifact for a catalog entry that does not exist", func() {
			shouldSkipTest()

			detailedResponse, err := service.UploadArtifact(defaultArtifact)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

		It("Delete artifact for a catalog entry", func() {
			shouldSkipTest()

			service.CreateCatalogEntry(defaultCreate)
			service.UploadArtifact(defaultArtifact)

			detailedResponse, err := service.DeleteArtifact(deleteArtifact)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})

		It("Fail to delete artifact for a catalog entry that does not exist", func() {
			shouldSkipTest()

			detailedResponse, err := service.DeleteArtifact(deleteArtifact)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

	})

})
