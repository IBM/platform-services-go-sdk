// +build examples

/**
 * (C) Copyright IBM Corp. 2021.
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
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"os"
)

const externalConfigFile = "../global_catalog_v1.env"

var (
	globalCatalogService *globalcatalogv1.GlobalCatalogV1
	config       map[string]string
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`GlobalCatalogV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
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
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			globalCatalogServiceOptions := &globalcatalogv1.GlobalCatalogV1Options{}

			globalCatalogService, err = globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(globalCatalogServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(globalCatalogService).ToNot(BeNil())
		})
	})

	Describe(`GlobalCatalogV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCatalogEntries request example`, func() {
			// begin-list_catalog_entries

			listCatalogEntriesOptions := globalCatalogService.NewListCatalogEntriesOptions()

			entrySearchResult, response, err := globalCatalogService.ListCatalogEntries(listCatalogEntriesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(entrySearchResult, "", "  ")
			fmt.Println(string(b))

			// end-list_catalog_entries

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(entrySearchResult).ToNot(BeNil())

		})
		It(`CreateCatalogEntry request example`, func() {
			// begin-create_catalog_entry

			overviewModel := &globalcatalogv1.Overview{
				DisplayName: core.StringPtr("testString"),
				LongDescription: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			imageModel := &globalcatalogv1.Image{
				Image: core.StringPtr("testString"),
			}

			providerModel := &globalcatalogv1.Provider{
				Email: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			createCatalogEntryOptions := globalCatalogService.NewCreateCatalogEntryOptions(
				"testString",
				"service",
				make(map[string]globalcatalogv1.Overview),
				imageModel,
				true,
				[]string{"testString"},
				providerModel,
				"testString",
			)
			createCatalogEntryOptions.OverviewUI["foo"] = *overviewModel

			catalogEntry, response, err := globalCatalogService.CreateCatalogEntry(createCatalogEntryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogEntry, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog_entry

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogEntry).ToNot(BeNil())

		})
		It(`GetCatalogEntry request example`, func() {
			// begin-get_catalog_entry

			getCatalogEntryOptions := globalCatalogService.NewGetCatalogEntryOptions(
				"testString",
			)

			catalogEntry, response, err := globalCatalogService.GetCatalogEntry(getCatalogEntryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogEntry, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_entry

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogEntry).ToNot(BeNil())

		})
		It(`UpdateCatalogEntry request example`, func() {
			// begin-update_catalog_entry

			overviewModel := &globalcatalogv1.Overview{
				DisplayName: core.StringPtr("testString"),
				LongDescription: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			imageModel := &globalcatalogv1.Image{
				Image: core.StringPtr("testString"),
			}

			providerModel := &globalcatalogv1.Provider{
				Email: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			updateCatalogEntryOptions := globalCatalogService.NewUpdateCatalogEntryOptions(
				"testString",
				"testString",
				"service",
				make(map[string]globalcatalogv1.Overview),
				imageModel,
				true,
				[]string{"testString"},
				providerModel,
			)
			updateCatalogEntryOptions.OverviewUI["foo"] = *overviewModel

			catalogEntry, response, err := globalCatalogService.UpdateCatalogEntry(updateCatalogEntryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogEntry, "", "  ")
			fmt.Println(string(b))

			// end-update_catalog_entry

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogEntry).ToNot(BeNil())

		})
		It(`GetChildObjects request example`, func() {
			// begin-get_child_objects

			getChildObjectsOptions := globalCatalogService.NewGetChildObjectsOptions(
				"testString",
				"testString",
			)

			entrySearchResult, response, err := globalCatalogService.GetChildObjects(getChildObjectsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(entrySearchResult, "", "  ")
			fmt.Println(string(b))

			// end-get_child_objects

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(entrySearchResult).ToNot(BeNil())

		})
		It(`RestoreCatalogEntry request example`, func() {
			// begin-restore_catalog_entry

			restoreCatalogEntryOptions := globalCatalogService.NewRestoreCatalogEntryOptions(
				"testString",
			)

			response, err := globalCatalogService.RestoreCatalogEntry(restoreCatalogEntryOptions)
			if err != nil {
				panic(err)
			}

			// end-restore_catalog_entry

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`GetVisibility request example`, func() {
			// begin-get_visibility

			getVisibilityOptions := globalCatalogService.NewGetVisibilityOptions(
				"testString",
			)

			visibility, response, err := globalCatalogService.GetVisibility(getVisibilityOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(visibility, "", "  ")
			fmt.Println(string(b))

			// end-get_visibility

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(visibility).ToNot(BeNil())

		})
		It(`UpdateVisibility request example`, func() {
			// begin-update_visibility

			updateVisibilityOptions := globalCatalogService.NewUpdateVisibilityOptions(
				"testString",
			)

			response, err := globalCatalogService.UpdateVisibility(updateVisibilityOptions)
			if err != nil {
				panic(err)
			}

			// end-update_visibility

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`GetPricing request example`, func() {
			// begin-get_pricing

			getPricingOptions := globalCatalogService.NewGetPricingOptions(
				"testString",
			)

			pricingGet, response, err := globalCatalogService.GetPricing(getPricingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pricingGet, "", "  ")
			fmt.Println(string(b))

			// end-get_pricing

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pricingGet).ToNot(BeNil())

		})
		It(`GetAuditLogs request example`, func() {
			// begin-get_audit_logs

			getAuditLogsOptions := globalCatalogService.NewGetAuditLogsOptions(
				"testString",
			)

			auditSearchResult, response, err := globalCatalogService.GetAuditLogs(getAuditLogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(auditSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-get_audit_logs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditSearchResult).ToNot(BeNil())

		})
		It(`ListArtifacts request example`, func() {
			// begin-list_artifacts

			listArtifactsOptions := globalCatalogService.NewListArtifactsOptions(
				"testString",
			)

			artifacts, response, err := globalCatalogService.ListArtifacts(listArtifactsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(artifacts, "", "  ")
			fmt.Println(string(b))

			// end-list_artifacts

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(artifacts).ToNot(BeNil())

		})
		It(`GetArtifact request example`, func() {
			// begin-get_artifact

			getArtifactOptions := globalCatalogService.NewGetArtifactOptions(
				"testString",
				"testString",
			)

			result, response, err := globalCatalogService.GetArtifact(getArtifactOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil { panic(err) }
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil { panic(err) }
			}

			// end-get_artifact

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`UploadArtifact request example`, func() {
			// begin-upload_artifact

			uploadArtifactOptions := globalCatalogService.NewUploadArtifactOptions(
				"testString",
				"testString",
			)

			response, err := globalCatalogService.UploadArtifact(uploadArtifactOptions)
			if err != nil {
				panic(err)
			}

			// end-upload_artifact

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteCatalogEntry request example`, func() {
			// begin-delete_catalog_entry

			deleteCatalogEntryOptions := globalCatalogService.NewDeleteCatalogEntryOptions(
				"testString",
			)

			response, err := globalCatalogService.DeleteCatalogEntry(deleteCatalogEntryOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_catalog_entry

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteArtifact request example`, func() {
			// begin-delete_artifact

			deleteArtifactOptions := globalCatalogService.NewDeleteArtifactOptions(
				"testString",
				"testString",
			)

			response, err := globalCatalogService.DeleteArtifact(deleteArtifactOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_artifact

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})
