//go:build integration
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
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Global Catalog - Integration Tests", func() {

	const (
		externalConfigFile    = "../global_catalog.env"
		visibilityRestriction = "private"
		artifact              = `{"someKey": "someValue"}`
	)

	var (
		service                     *globalcatalogv1.GlobalCatalogV1
		defaultCreate               *globalcatalogv1.CreateCatalogEntryOptions
		defaultDelete               *globalcatalogv1.DeleteCatalogEntryOptions
		forceDelete                 *globalcatalogv1.DeleteCatalogEntryOptions
		defaultGet                  *globalcatalogv1.GetCatalogEntryOptions
		defaultUpdate               *globalcatalogv1.UpdateCatalogEntryOptions
		defaultList                 *globalcatalogv1.ListCatalogEntriesOptions
		defaultChild                *globalcatalogv1.CreateCatalogEntryOptions
		getChild                    *globalcatalogv1.GetChildObjectsOptions
		defaultRestore              *globalcatalogv1.RestoreCatalogEntryOptions
		bogusRestore                *globalcatalogv1.RestoreCatalogEntryOptions
		getVisibility               *globalcatalogv1.GetVisibilityOptions
		updateVisibility            *globalcatalogv1.UpdateVisibilityOptions
		getPricing                  *globalcatalogv1.GetPricingOptions
		defaultArtifact             *globalcatalogv1.UploadArtifactOptions
		uploadArtifactList          *globalcatalogv1.UploadArtifactOptions
		uploadArtifactCreate        *globalcatalogv1.UploadArtifactOptions
		uploadArtifactCreateFailure *globalcatalogv1.UploadArtifactOptions
		uploadArtifactDelete        *globalcatalogv1.UploadArtifactOptions
		listArtifacts               *globalcatalogv1.ListArtifactsOptions
		getArtifact                 *globalcatalogv1.GetArtifactOptions
		deleteArtifact              *globalcatalogv1.DeleteArtifactOptions
		config                      map[string]string
		configLoaded                bool = false
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping...")
		}
	}

	It("Successfully load the configuration", func() {
		var err error
		_, err = os.Stat(externalConfigFile)
		if err != nil {
			Skip("External configuration file not found, skipping tests: " + err.Error())
		}

		os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
		config, err = core.GetServiceProperties(globalcatalogv1.DefaultServiceName)
		if err != nil {
			Skip("Error loading service properties, skipping tests: " + err.Error())
		}

		configLoaded = len(config) > 0
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
			overviewUi          = make(map[string]globalcatalogv1.Overview)
			overviewUiUpdated   = make(map[string]globalcatalogv1.Overview)
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

		core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
		service.EnableRetries(4, 30*time.Second)

		overviewUi[en] = *overview
		overviewUiUpdated[en] = *overviewUpdated

		defaultCreate = service.NewCreateCatalogEntryOptions(name,
			kind,
			overviewUi,
			images,
			disabled,
			tags,
			provider,
			id)
		defaultDelete = service.NewDeleteCatalogEntryOptions(id)
		forceDelete = service.NewDeleteCatalogEntryOptions(id)
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
		getChild = service.NewGetChildObjectsOptions(id, kind)
		defaultRestore = service.NewRestoreCatalogEntryOptions(id)
		bogusRestore = service.NewRestoreCatalogEntryOptions("bogus")
		getVisibility = service.NewGetVisibilityOptions(id)
		updateVisibility = service.NewUpdateVisibilityOptions(id)
		getPricing = service.NewGetPricingOptions(id)
		defaultArtifact = service.NewUploadArtifactOptions(id, artifactId)
		uploadArtifactList = service.NewUploadArtifactOptions(id, artifactId)
		uploadArtifactCreate = service.NewUploadArtifactOptions(id, artifactId)
		uploadArtifactCreateFailure = service.NewUploadArtifactOptions(id, artifactId)
		uploadArtifactDelete = service.NewUploadArtifactOptions(id, artifactId)
		listArtifacts = service.NewListArtifactsOptions(id)
		getArtifact = service.NewGetArtifactOptions(id, artifactId)
		deleteArtifact = service.NewDeleteArtifactOptions(id, artifactId)

		defaultChild.SetParentID(id)
		defaultArtifact.SetArtifact(ioutil.NopCloser(strings.NewReader(artifact)))
		uploadArtifactList.SetArtifact(ioutil.NopCloser(strings.NewReader(artifact)))
		uploadArtifactCreate.SetArtifact(ioutil.NopCloser(strings.NewReader(artifact)))
		uploadArtifactCreateFailure.SetArtifact(ioutil.NopCloser(strings.NewReader(artifact)))
		uploadArtifactDelete.SetArtifact(ioutil.NopCloser(strings.NewReader(artifact)))
		forceDelete.SetForce(true)
	})

	Describe("Run integration tests", func() {
		JustBeforeEach(func() {
			shouldSkipTest()

			_, _ = service.DeleteCatalogEntry(forceDelete)
		})

		JustAfterEach(func() {
			shouldSkipTest()

			_, _ = service.DeleteCatalogEntry(forceDelete)
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

			fmt.Fprintf(GinkgoWriter, "CreateCatalogEntry() result:\n%s", common.ToJSON(result))
		})

		It("Get a catalog entry", func() {
			shouldSkipTest()

			_, _, _ = service.CreateCatalogEntry(defaultCreate)
			result, detailedResponse, err := service.GetCatalogEntry(defaultGet)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.ID).To(Equal(defaultCreate.ID))
			Expect(result.Name).To(Equal(defaultCreate.Name))
			Expect(result.Kind).To(Equal(defaultCreate.Kind))
			Expect(result.Images).To(Equal(defaultCreate.Images))
			Expect(result.Tags).To(Equal(defaultCreate.Tags))
			Expect(result.Provider).To(Equal(defaultCreate.Provider))

			fmt.Fprintf(GinkgoWriter, "GetCatalogEntry() result:\n%s", common.ToJSON(result))
		})

		It("Update a catalog entry", func() {
			shouldSkipTest()

			_, _, _ = service.CreateCatalogEntry(defaultCreate)
			result, detailedResponse, err := service.UpdateCatalogEntry(defaultUpdate)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.ID).To(Equal(defaultUpdate.ID))
			Expect(result.Name).To(Equal(defaultUpdate.Name))
			Expect(result.Kind).To(Equal(defaultUpdate.Kind))
			Expect(result.Images).To(Equal(defaultUpdate.Images))
			Expect(result.Tags).To(Equal(defaultUpdate.Tags))
			Expect(result.Provider).To(Equal(defaultUpdate.Provider))

			fmt.Fprintf(GinkgoWriter, "UpdateCatalogEntry() result:\n%s", common.ToJSON(result))
		})

		It("Delete a catalog entry", func() {
			shouldSkipTest()

			_, _, _ = service.CreateCatalogEntry(defaultCreate)
			detailedResponse, err := service.DeleteCatalogEntry(forceDelete)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})

		It("Fail to get a catalog entry after deletion", func() {
			shouldSkipTest()

			_, _, _ = service.CreateCatalogEntry(defaultCreate)
			_, _ = service.DeleteCatalogEntry(forceDelete)

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

			detailedResponse, err := service.DeleteCatalogEntry(forceDelete)
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

			_, _, _ = service.CreateCatalogEntry(defaultCreate)

			_, detailedResponse, err := service.CreateCatalogEntry(defaultCreate)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(409))
		})

		It("List catalog entries", func() {
			shouldSkipTest()

			result, detailedResponse, err := service.ListCatalogEntries(defaultList)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListCatalogEntries() result:\n%s", common.ToJSON(result))
			Expect(result.Resources).NotTo(BeNil())
			Expect(len(result.Resources)).NotTo(BeZero())
		})

		It("Get child catalog entry", func() {
			const expectedOffset int64 = 0
			const expectedCount int64 = 1
			const expectedResourceCount int64 = 1

			shouldSkipTest()

			_, _, _ = service.CreateCatalogEntry(defaultCreate)
			_, _, _ = service.CreateCatalogEntry(defaultChild)

			result, detailedResponse, err := service.GetChildObjects(getChild)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetChildObjects() result:\n%s", common.ToJSON(result))
			Expect(*result.Offset).To(Equal(expectedOffset))
			Expect(*result.Count).To(Equal(expectedCount))
			Expect(*result.ResourceCount).To(Equal(expectedResourceCount))
			Expect(result.Resources[0].Name).To(Equal(defaultChild.Name))
			Expect(result.Resources[0].Kind).To(Equal(defaultChild.Kind))
			Expect(result.Resources[0].Images).To(Equal(defaultChild.Images))
			Expect(result.Resources[0].Tags).To(Equal(defaultChild.Tags))
			Expect(result.Resources[0].Provider).To(Equal(defaultChild.Provider))
		})

		It("Fail to get a child catalog entry that does not exist", func() {
			shouldSkipTest()

			_, detailedResponse, err := service.GetChildObjects(getChild)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

		It("Restore a catalog entry", func() {
			shouldSkipTest()

			_, _, _ = service.CreateCatalogEntry(defaultCreate)
			_, _ = service.DeleteCatalogEntry(defaultDelete)

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

			_, _, _ = service.CreateCatalogEntry(defaultCreate)

			result, detailedResponse, err := service.GetVisibility(getVisibility)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetVisibility() result:\n%s", common.ToJSON(result))
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

			_, _, _ = service.CreateCatalogEntry(defaultCreate)

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

			_, _, _ = service.CreateCatalogEntry(defaultCreate)

			_, detailedResponseExists, errExists := service.GetPricing(getPricing)
			Expect(errExists).NotTo(BeNil())
			Expect(detailedResponseExists.StatusCode).To(Equal(404))

			_, _ = service.DeleteCatalogEntry(forceDelete)

			_, detailedResponseNotExists, errNotExists := service.GetPricing(getPricing)
			Expect(errNotExists).NotTo(BeNil())
			Expect(detailedResponseNotExists.StatusCode).To(Equal(404))
		})

		It("List artifacts for a catalog entry", func() {
			const expectedCount int64 = 1
			const expectedSize int64 = 24

			shouldSkipTest()

			_, _, _ = service.CreateCatalogEntry(defaultCreate)
			_, _ = service.UploadArtifact(uploadArtifactList)

			result, detailedResponse, err := service.ListArtifacts(listArtifacts)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListArtifacts() result:\n%s", common.ToJSON(result))
			Expect(*result.Count).To(Equal(expectedCount))
			Expect(len(result.Resources)).To(Equal(1))
			Expect(result.Resources[0].Name).To(Equal(uploadArtifactList.ArtifactID))
			Expect(*result.Resources[0].URL).To(Equal(fmt.Sprintf("%s/%s/artifacts/%s", "https://globalcatalog.test.cloud.ibm.com/api/v1", *defaultCreate.ID, *uploadArtifactList.ArtifactID)))
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

			_, _, _ = service.CreateCatalogEntry(defaultCreate)
			_, _ = service.UploadArtifact(defaultArtifact)

			result, detailedResponse, err := service.GetArtifact(getArtifact)
			Expect(result).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(result)
			Expect(buf.String()).To(Equal(artifact))
		})

		It("Fail to get artifacts that do not exists", func() {
			shouldSkipTest()

			_, _, _ = service.CreateCatalogEntry(defaultCreate)

			_, detailedResponseExists, errExists := service.GetArtifact(getArtifact)
			Expect(errExists).NotTo(BeNil())
			Expect(detailedResponseExists.StatusCode).To(Equal(404))

			_, _ = service.DeleteCatalogEntry(forceDelete)

			_, detailedResponseNotExists, errNotExists := service.GetArtifact(getArtifact)
			Expect(errNotExists).NotTo(BeNil())
			Expect(detailedResponseNotExists.StatusCode).To(Equal(404))
		})

		It("Create artifact for a catalog entry", func() {
			shouldSkipTest()

			_, _, _ = service.CreateCatalogEntry(defaultCreate)

			detailedResponse, err := service.UploadArtifact(uploadArtifactCreate)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})

		It("Fail to artifact for a catalog entry that does not exist", func() {
			shouldSkipTest()

			detailedResponse, err := service.UploadArtifact(uploadArtifactCreateFailure)
			Expect(err).NotTo(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})

		It("Delete artifact for a catalog entry", func() {
			shouldSkipTest()

			_, _, _ = service.CreateCatalogEntry(defaultCreate)
			_, _ = service.UploadArtifact(uploadArtifactDelete)

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
