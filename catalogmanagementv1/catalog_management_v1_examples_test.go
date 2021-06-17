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

package catalogmanagementv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Catalog Management service.
//
// The following configuration properties are assumed to be defined:
// CATALOG_MANAGEMENT_URL=<service base url>
// CATALOG_MANAGEMENT_AUTH_TYPE=iam
// CATALOG_MANAGEMENT_APIKEY=<IAM apikey>
// CATALOG_MANAGEMENT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../catalog_management_v1.env"

var (
	catalogManagementService *catalogmanagementv1.CatalogManagementV1
	config                   map[string]string
	configLoaded             bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`CatalogManagementV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(catalogmanagementv1.DefaultServiceName)
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

			catalogManagementServiceOptions := &catalogmanagementv1.CatalogManagementV1Options{}

			catalogManagementService, err = catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(catalogManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(catalogManagementService).ToNot(BeNil())
		})
	})

	Describe(`CatalogManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`CreateCatalog request example`, func() {
			fmt.Println("\nCreateCatalog() result:")
			// begin-create_catalog

			createCatalogOptions := catalogManagementService.NewCreateCatalogOptions()

			catalog, response, err := catalogManagementService.CreateCatalog(createCatalogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalog, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalog).ToNot(BeNil())

		})

		It(`GetCatalog request example`, func() {
			fmt.Println("\nGetCatalog() result:")
			// begin-get_catalog

			getCatalogOptions := catalogManagementService.NewGetCatalogOptions(
				"testString",
			)

			catalog, response, err := catalogManagementService.GetCatalog(getCatalogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalog, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())

		})

		It(`ReplaceCatalog request example`, func() {
			fmt.Println("\nReplaceCatalog() result:")
			// begin-replace_catalog

			replaceCatalogOptions := catalogManagementService.NewReplaceCatalogOptions(
				"testString",
			)

			catalog, response, err := catalogManagementService.ReplaceCatalog(replaceCatalogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalog, "", "  ")
			fmt.Println(string(b))

			// end-replace_catalog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())

		})

		It(`ListCatalogs request example`, func() {
			fmt.Println("\nListCatalogs() result:")
			// begin-list_catalogs

			listCatalogsOptions := catalogManagementService.NewListCatalogsOptions()

			catalogSearchResult, response, err := catalogManagementService.ListCatalogs(listCatalogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-list_catalogs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogSearchResult).ToNot(BeNil())

		})

		It(`CreateOffering request example`, func() {
			fmt.Println("\nCreateOffering() result:")
			// begin-create_offering

			createOfferingOptions := catalogManagementService.NewCreateOfferingOptions(
				"testString",
			)

			offering, response, err := catalogManagementService.CreateOffering(createOfferingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-create_offering

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())

		})

		It(`GetOffering request example`, func() {
			fmt.Println("\nGetOffering() result:")
			// begin-get_offering

			getOfferingOptions := catalogManagementService.NewGetOfferingOptions(
				"testString",
				"testString",
			)

			offering, response, err := catalogManagementService.GetOffering(getOfferingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-get_offering

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

		})

		It(`ReplaceOffering request example`, func() {
			fmt.Println("\nReplaceOffering() result:")
			// begin-replace_offering

			replaceOfferingOptions := catalogManagementService.NewReplaceOfferingOptions(
				"testString",
				"testString",
			)

			offering, response, err := catalogManagementService.ReplaceOffering(replaceOfferingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-replace_offering

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

		})

		It(`ListOfferings request example`, func() {
			fmt.Println("\nListOfferings() result:")
			// begin-list_offerings

			listOfferingsOptions := catalogManagementService.NewListOfferingsOptions(
				"testString",
			)

			offeringSearchResult, response, err := catalogManagementService.ListOfferings(listOfferingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offeringSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-list_offerings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringSearchResult).ToNot(BeNil())

		})

		It(`ImportOffering request example`, func() {
			fmt.Println("\nImportOffering() result:")
			// begin-import_offering

			importOfferingOptions := catalogManagementService.NewImportOfferingOptions(
				"testString",
			)

			offering, response, err := catalogManagementService.ImportOffering(importOfferingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-import_offering

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())

		})

		It(`ReloadOffering request example`, func() {
			fmt.Println("\nReloadOffering() result:")
			// begin-reload_offering

			reloadOfferingOptions := catalogManagementService.NewReloadOfferingOptions(
				"testString",
				"testString",
				"testString",
			)

			offering, response, err := catalogManagementService.ReloadOffering(reloadOfferingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-reload_offering

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())

		})

		It(`CreateObject request example`, func() {
			fmt.Println("\nCreateObject() result:")
			// begin-create_object

			createObjectOptions := catalogManagementService.NewCreateObjectOptions(
				"testString",
			)

			catalogObject, response, err := catalogManagementService.CreateObject(createObjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogObject, "", "  ")
			fmt.Println(string(b))

			// end-create_object

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogObject).ToNot(BeNil())

		})

		It(`GetOfferingAudit request example`, func() {
			fmt.Println("\nGetOfferingAudit() result:")
			// begin-get_offering_audit

			getOfferingAuditOptions := catalogManagementService.NewGetOfferingAuditOptions(
				"testString",
				"testString",
			)

			auditLog, response, err := catalogManagementService.GetOfferingAudit(getOfferingAuditOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(auditLog, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_audit

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditLog).ToNot(BeNil())

		})

		It(`GetCatalogAccount request example`, func() {
			fmt.Println("\nGetCatalogAccount() result:")
			// begin-get_catalog_account

			getCatalogAccountOptions := catalogManagementService.NewGetCatalogAccountOptions()

			account, response, err := catalogManagementService.GetCatalogAccount(getCatalogAccountOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(account, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_account

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(account).ToNot(BeNil())

		})

		It(`UpdateCatalogAccount request example`, func() {
			// begin-update_catalog_account

			updateCatalogAccountOptions := catalogManagementService.NewUpdateCatalogAccountOptions()

			response, err := catalogManagementService.UpdateCatalogAccount(updateCatalogAccountOptions)
			if err != nil {
				panic(err)
			}

			// end-update_catalog_account
			fmt.Printf("\nUpdateCatalogAccount() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`GetCatalogAccountAudit request example`, func() {
			fmt.Println("\nGetCatalogAccountAudit() result:")
			// begin-get_catalog_account_audit

			getCatalogAccountAuditOptions := catalogManagementService.NewGetCatalogAccountAuditOptions()

			auditLog, response, err := catalogManagementService.GetCatalogAccountAudit(getCatalogAccountAuditOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(auditLog, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_account_audit

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditLog).ToNot(BeNil())

		})

		It(`GetCatalogAccountFilters request example`, func() {
			fmt.Println("\nGetCatalogAccountFilters() result:")
			// begin-get_catalog_account_filters

			getCatalogAccountFiltersOptions := catalogManagementService.NewGetCatalogAccountFiltersOptions()

			accumulatedFilters, response, err := catalogManagementService.GetCatalogAccountFilters(getCatalogAccountFiltersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accumulatedFilters, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_account_filters

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accumulatedFilters).ToNot(BeNil())

		})

		It(`GetCatalogAudit request example`, func() {
			fmt.Println("\nGetCatalogAudit() result:")
			// begin-get_catalog_audit

			getCatalogAuditOptions := catalogManagementService.NewGetCatalogAuditOptions(
				"testString",
			)

			auditLog, response, err := catalogManagementService.GetCatalogAudit(getCatalogAuditOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(auditLog, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_audit

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditLog).ToNot(BeNil())

		})

		It(`GetConsumptionOfferings request example`, func() {
			fmt.Println("\nGetConsumptionOfferings() result:")
			// begin-get_consumption_offerings

			getConsumptionOfferingsOptions := catalogManagementService.NewGetConsumptionOfferingsOptions()

			offeringSearchResult, response, err := catalogManagementService.GetConsumptionOfferings(getConsumptionOfferingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offeringSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-get_consumption_offerings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringSearchResult).ToNot(BeNil())

		})

		It(`ImportOfferingVersion request example`, func() {
			fmt.Println("\nImportOfferingVersion() result:")
			// begin-import_offering_version

			importOfferingVersionOptions := catalogManagementService.NewImportOfferingVersionOptions(
				"testString",
				"testString",
			)

			offering, response, err := catalogManagementService.ImportOfferingVersion(importOfferingVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-import_offering_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())

		})

		It(`ReplaceOfferingIcon request example`, func() {
			fmt.Println("\nReplaceOfferingIcon() result:")
			// begin-replace_offering_icon

			replaceOfferingIconOptions := catalogManagementService.NewReplaceOfferingIconOptions(
				"testString",
				"testString",
				"testString",
			)

			offering, response, err := catalogManagementService.ReplaceOfferingIcon(replaceOfferingIconOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-replace_offering_icon

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

		})

		It(`UpdateOfferingIBM request example`, func() {
			fmt.Println("\nUpdateOfferingIBM() result:")
			// begin-update_offering_ibm

			updateOfferingIBMOptions := catalogManagementService.NewUpdateOfferingIBMOptions(
				"testString",
				"testString",
				"allow_request",
				"true",
			)

			approvalResult, response, err := catalogManagementService.UpdateOfferingIBM(updateOfferingIBMOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(approvalResult, "", "  ")
			fmt.Println(string(b))

			// end-update_offering_ibm

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(approvalResult).ToNot(BeNil())

		})

		It(`GetOfferingUpdates request example`, func() {
			fmt.Println("\nGetOfferingUpdates() result:")
			// begin-get_offering_updates

			getOfferingUpdatesOptions := catalogManagementService.NewGetOfferingUpdatesOptions(
				"testString",
				"testString",
				"testString",
			)

			versionUpdateDescriptor, response, err := catalogManagementService.GetOfferingUpdates(getOfferingUpdatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(versionUpdateDescriptor, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_updates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(versionUpdateDescriptor).ToNot(BeNil())

		})

		It(`GetOfferingAbout request example`, func() {
			fmt.Println("\nGetOfferingAbout() result:")
			// begin-get_offering_about

			getOfferingAboutOptions := catalogManagementService.NewGetOfferingAboutOptions(
				"testString",
			)

			result, response, err := catalogManagementService.GetOfferingAbout(getOfferingAboutOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_about

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})

		It(`GetOfferingLicense request example`, func() {
			fmt.Println("\nGetOfferingLicense() result:")
			// begin-get_offering_license

			getOfferingLicenseOptions := catalogManagementService.NewGetOfferingLicenseOptions(
				"testString",
				"testString",
			)

			result, response, err := catalogManagementService.GetOfferingLicense(getOfferingLicenseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_license

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})

		It(`GetOfferingContainerImages request example`, func() {
			fmt.Println("\nGetOfferingContainerImages() result:")
			// begin-get_offering_container_images

			getOfferingContainerImagesOptions := catalogManagementService.NewGetOfferingContainerImagesOptions(
				"testString",
			)

			imageManifest, response, err := catalogManagementService.GetOfferingContainerImages(getOfferingContainerImagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(imageManifest, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_container_images

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageManifest).ToNot(BeNil())

		})

		It(`DeprecateVersion request example`, func() {
			// begin-deprecate_version

			deprecateVersionOptions := catalogManagementService.NewDeprecateVersionOptions(
				"testString",
			)

			response, err := catalogManagementService.DeprecateVersion(deprecateVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-deprecate_version
			fmt.Printf("\nDeprecateVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`AccountPublishVersion request example`, func() {
			// begin-account_publish_version

			accountPublishVersionOptions := catalogManagementService.NewAccountPublishVersionOptions(
				"testString",
			)

			response, err := catalogManagementService.AccountPublishVersion(accountPublishVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-account_publish_version
			fmt.Printf("\nAccountPublishVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`IBMPublishVersion request example`, func() {
			// begin-ibm_publish_version

			ibmPublishVersionOptions := catalogManagementService.NewIBMPublishVersionOptions(
				"testString",
			)

			response, err := catalogManagementService.IBMPublishVersion(ibmPublishVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-ibm_publish_version
			fmt.Printf("\nIBMPublishVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`PublicPublishVersion request example`, func() {
			// begin-public_publish_version

			publicPublishVersionOptions := catalogManagementService.NewPublicPublishVersionOptions(
				"testString",
			)

			response, err := catalogManagementService.PublicPublishVersion(publicPublishVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-public_publish_version
			fmt.Printf("\nPublicPublishVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`CommitVersion request example`, func() {
			// begin-commit_version

			commitVersionOptions := catalogManagementService.NewCommitVersionOptions(
				"testString",
			)

			response, err := catalogManagementService.CommitVersion(commitVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-commit_version
			fmt.Printf("\nCommitVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`CopyVersion request example`, func() {
			// begin-copy_version

			copyVersionOptions := catalogManagementService.NewCopyVersionOptions(
				"testString",
			)

			response, err := catalogManagementService.CopyVersion(copyVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-copy_version
			fmt.Printf("\nCopyVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`GetOfferingWorkingCopy request example`, func() {
			fmt.Println("\nGetOfferingWorkingCopy() result:")
			// begin-get_offering_working_copy

			getOfferingWorkingCopyOptions := catalogManagementService.NewGetOfferingWorkingCopyOptions(
				"testString",
			)

			version, response, err := catalogManagementService.GetOfferingWorkingCopy(getOfferingWorkingCopyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(version, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_working_copy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(version).ToNot(BeNil())

		})

		It(`GetVersion request example`, func() {
			fmt.Println("\nGetVersion() result:")
			// begin-get_version

			getVersionOptions := catalogManagementService.NewGetVersionOptions(
				"testString",
			)

			offering, response, err := catalogManagementService.GetVersion(getVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offering, "", "  ")
			fmt.Println(string(b))

			// end-get_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

		})

		It(`GetCluster request example`, func() {
			fmt.Println("\nGetCluster() result:")
			// begin-get_cluster

			getClusterOptions := catalogManagementService.NewGetClusterOptions(
				"testString",
				"testString",
				"testString",
			)

			clusterInfo, response, err := catalogManagementService.GetCluster(getClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clusterInfo, "", "  ")
			fmt.Println(string(b))

			// end-get_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clusterInfo).ToNot(BeNil())

		})

		It(`GetNamespaces request example`, func() {
			fmt.Println("\nGetNamespaces() result:")
			// begin-get_namespaces

			getNamespacesOptions := catalogManagementService.NewGetNamespacesOptions(
				"testString",
				"testString",
				"testString",
			)

			namespaceSearchResult, response, err := catalogManagementService.GetNamespaces(getNamespacesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(namespaceSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-get_namespaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(namespaceSearchResult).ToNot(BeNil())

		})

		It(`DeployOperators request example`, func() {
			fmt.Println("\nDeployOperators() result:")
			// begin-deploy_operators

			deployOperatorsOptions := catalogManagementService.NewDeployOperatorsOptions(
				"testString",
			)

			operatorDeployResult, response, err := catalogManagementService.DeployOperators(deployOperatorsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(operatorDeployResult, "", "  ")
			fmt.Println(string(b))

			// end-deploy_operators

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatorDeployResult).ToNot(BeNil())

		})

		It(`ListOperators request example`, func() {
			fmt.Println("\nListOperators() result:")
			// begin-list_operators

			listOperatorsOptions := catalogManagementService.NewListOperatorsOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			operatorDeployResult, response, err := catalogManagementService.ListOperators(listOperatorsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(operatorDeployResult, "", "  ")
			fmt.Println(string(b))

			// end-list_operators

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatorDeployResult).ToNot(BeNil())

		})

		It(`InstallVersion request example`, func() {
			// begin-install_version

			installVersionOptions := catalogManagementService.NewInstallVersionOptions(
				"testString",
				"testString",
			)

			response, err := catalogManagementService.InstallVersion(installVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-install_version
			fmt.Printf("\nInstallVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`ReplaceOperators request example`, func() {
			fmt.Println("\nReplaceOperators() result:")
			// begin-replace_operators

			replaceOperatorsOptions := catalogManagementService.NewReplaceOperatorsOptions(
				"testString",
			)

			operatorDeployResult, response, err := catalogManagementService.ReplaceOperators(replaceOperatorsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(operatorDeployResult, "", "  ")
			fmt.Println(string(b))

			// end-replace_operators

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatorDeployResult).ToNot(BeNil())

		})

		It(`PreinstallVersion request example`, func() {
			// begin-preinstall_version

			preinstallVersionOptions := catalogManagementService.NewPreinstallVersionOptions(
				"testString",
				"testString",
			)

			response, err := catalogManagementService.PreinstallVersion(preinstallVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-preinstall_version
			fmt.Printf("\nPreinstallVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`GetPreinstall request example`, func() {
			fmt.Println("\nGetPreinstall() result:")
			// begin-get_preinstall

			getPreinstallOptions := catalogManagementService.NewGetPreinstallOptions(
				"testString",
				"testString",
			)

			installStatus, response, err := catalogManagementService.GetPreinstall(getPreinstallOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(installStatus, "", "  ")
			fmt.Println(string(b))

			// end-get_preinstall

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(installStatus).ToNot(BeNil())

		})

		It(`ValidateInstall request example`, func() {
			// begin-validate_install

			validateInstallOptions := catalogManagementService.NewValidateInstallOptions(
				"testString",
				"testString",
			)

			response, err := catalogManagementService.ValidateInstall(validateInstallOptions)
			if err != nil {
				panic(err)
			}

			// end-validate_install
			fmt.Printf("\nValidateInstall() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`GetValidationStatus request example`, func() {
			fmt.Println("\nGetValidationStatus() result:")
			// begin-get_validation_status

			getValidationStatusOptions := catalogManagementService.NewGetValidationStatusOptions(
				"testString",
				"testString",
			)

			validation, response, err := catalogManagementService.GetValidationStatus(getValidationStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(validation, "", "  ")
			fmt.Println(string(b))

			// end-get_validation_status

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(validation).ToNot(BeNil())

		})

		It(`GetOverrideValues request example`, func() {
			fmt.Println("\nGetOverrideValues() result:")
			// begin-get_override_values

			getOverrideValuesOptions := catalogManagementService.NewGetOverrideValuesOptions(
				"testString",
			)

			getOverrideValuesResponse, response, err := catalogManagementService.GetOverrideValues(getOverrideValuesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getOverrideValuesResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_override_values

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getOverrideValuesResponse).ToNot(BeNil())

		})

		It(`SearchObjects request example`, func() {
			fmt.Println("\nSearchObjects() result:")
			// begin-search_objects

			searchObjectsOptions := catalogManagementService.NewSearchObjectsOptions(
				"testString",
			)

			objectSearchResult, response, err := catalogManagementService.SearchObjects(searchObjectsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(objectSearchResult, "", "  ")
			fmt.Println(string(b))

			// end-search_objects

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectSearchResult).ToNot(BeNil())

		})

		It(`ListObjects request example`, func() {
			fmt.Println("\nListObjects() result:")
			// begin-list_objects

			listObjectsOptions := catalogManagementService.NewListObjectsOptions(
				"testString",
			)

			objectListResult, response, err := catalogManagementService.ListObjects(listObjectsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(objectListResult, "", "  ")
			fmt.Println(string(b))

			// end-list_objects

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectListResult).ToNot(BeNil())

		})

		It(`GetObject request example`, func() {
			fmt.Println("\nGetObject() result:")
			// begin-get_object

			getObjectOptions := catalogManagementService.NewGetObjectOptions(
				"testString",
				"testString",
			)

			catalogObject, response, err := catalogManagementService.GetObject(getObjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogObject, "", "  ")
			fmt.Println(string(b))

			// end-get_object

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogObject).ToNot(BeNil())

		})

		It(`ReplaceObject request example`, func() {
			fmt.Println("\nReplaceObject() result:")
			// begin-replace_object

			replaceObjectOptions := catalogManagementService.NewReplaceObjectOptions(
				"testString",
				"testString",
			)

			catalogObject, response, err := catalogManagementService.ReplaceObject(replaceObjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogObject, "", "  ")
			fmt.Println(string(b))

			// end-replace_object

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogObject).ToNot(BeNil())

		})

		It(`GetObjectAudit request example`, func() {
			fmt.Println("\nGetObjectAudit() result:")
			// begin-get_object_audit

			getObjectAuditOptions := catalogManagementService.NewGetObjectAuditOptions(
				"testString",
				"testString",
			)

			auditLog, response, err := catalogManagementService.GetObjectAudit(getObjectAuditOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(auditLog, "", "  ")
			fmt.Println(string(b))

			// end-get_object_audit

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditLog).ToNot(BeNil())

		})

		It(`AccountPublishObject request example`, func() {
			// begin-account_publish_object

			accountPublishObjectOptions := catalogManagementService.NewAccountPublishObjectOptions(
				"testString",
				"testString",
			)

			response, err := catalogManagementService.AccountPublishObject(accountPublishObjectOptions)
			if err != nil {
				panic(err)
			}

			// end-account_publish_object
			fmt.Printf("\nAccountPublishObject() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`SharedPublishObject request example`, func() {
			// begin-shared_publish_object

			sharedPublishObjectOptions := catalogManagementService.NewSharedPublishObjectOptions(
				"testString",
				"testString",
			)

			response, err := catalogManagementService.SharedPublishObject(sharedPublishObjectOptions)
			if err != nil {
				panic(err)
			}

			// end-shared_publish_object
			fmt.Printf("\nSharedPublishObject() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`IBMPublishObject request example`, func() {
			// begin-ibm_publish_object

			ibmPublishObjectOptions := catalogManagementService.NewIBMPublishObjectOptions(
				"testString",
				"testString",
			)

			response, err := catalogManagementService.IBMPublishObject(ibmPublishObjectOptions)
			if err != nil {
				panic(err)
			}

			// end-ibm_publish_object
			fmt.Printf("\nIBMPublishObject() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`PublicPublishObject request example`, func() {
			// begin-public_publish_object

			publicPublishObjectOptions := catalogManagementService.NewPublicPublishObjectOptions(
				"testString",
				"testString",
			)

			response, err := catalogManagementService.PublicPublishObject(publicPublishObjectOptions)
			if err != nil {
				panic(err)
			}

			// end-public_publish_object
			fmt.Printf("\nPublicPublishObject() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`CreateObjectAccess request example`, func() {
			// begin-create_object_access

			createObjectAccessOptions := catalogManagementService.NewCreateObjectAccessOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := catalogManagementService.CreateObjectAccess(createObjectAccessOptions)
			if err != nil {
				panic(err)
			}

			// end-create_object_access
			fmt.Printf("\nCreateObjectAccess() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})

		It(`GetObjectAccess request example`, func() {
			fmt.Println("\nGetObjectAccess() result:")
			// begin-get_object_access

			getObjectAccessOptions := catalogManagementService.NewGetObjectAccessOptions(
				"testString",
				"testString",
				"testString",
			)

			objectAccess, response, err := catalogManagementService.GetObjectAccess(getObjectAccessOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(objectAccess, "", "  ")
			fmt.Println(string(b))

			// end-get_object_access

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectAccess).ToNot(BeNil())

		})

		It(`AddObjectAccessList request example`, func() {
			fmt.Println("\nAddObjectAccessList() result:")
			// begin-add_object_access_list

			addObjectAccessListOptions := catalogManagementService.NewAddObjectAccessListOptions(
				"testString",
				"testString",
				[]string{"testString"},
			)

			accessListBulkResponse, response, err := catalogManagementService.AddObjectAccessList(addObjectAccessListOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accessListBulkResponse, "", "  ")
			fmt.Println(string(b))

			// end-add_object_access_list

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessListBulkResponse).ToNot(BeNil())

		})

		It(`GetObjectAccessList request example`, func() {
			fmt.Println("\nGetObjectAccessList() result:")
			// begin-get_object_access_list

			getObjectAccessListOptions := catalogManagementService.NewGetObjectAccessListOptions(
				"testString",
				"testString",
			)

			objectAccessListResult, response, err := catalogManagementService.GetObjectAccessList(getObjectAccessListOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(objectAccessListResult, "", "  ")
			fmt.Println(string(b))

			// end-get_object_access_list

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectAccessListResult).ToNot(BeNil())

		})

		It(`CreateOfferingInstance request example`, func() {
			fmt.Println("\nCreateOfferingInstance() result:")
			// begin-create_offering_instance

			createOfferingInstanceOptions := catalogManagementService.NewCreateOfferingInstanceOptions(
				"testString",
			)

			offeringInstance, response, err := catalogManagementService.CreateOfferingInstance(createOfferingInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offeringInstance, "", "  ")
			fmt.Println(string(b))

			// end-create_offering_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offeringInstance).ToNot(BeNil())

		})

		It(`GetOfferingInstance request example`, func() {
			fmt.Println("\nGetOfferingInstance() result:")
			// begin-get_offering_instance

			getOfferingInstanceOptions := catalogManagementService.NewGetOfferingInstanceOptions(
				"testString",
			)

			offeringInstance, response, err := catalogManagementService.GetOfferingInstance(getOfferingInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offeringInstance, "", "  ")
			fmt.Println(string(b))

			// end-get_offering_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringInstance).ToNot(BeNil())

		})

		It(`PutOfferingInstance request example`, func() {
			fmt.Println("\nPutOfferingInstance() result:")
			// begin-put_offering_instance

			putOfferingInstanceOptions := catalogManagementService.NewPutOfferingInstanceOptions(
				"testString",
				"testString",
			)

			offeringInstance, response, err := catalogManagementService.PutOfferingInstance(putOfferingInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(offeringInstance, "", "  ")
			fmt.Println(string(b))

			// end-put_offering_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringInstance).ToNot(BeNil())

		})

		It(`DeleteVersion request example`, func() {
			// begin-delete_version

			deleteVersionOptions := catalogManagementService.NewDeleteVersionOptions(
				"testString",
			)

			response, err := catalogManagementService.DeleteVersion(deleteVersionOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_version
			fmt.Printf("\nDeleteVersion() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`DeleteOperators request example`, func() {
			// begin-delete_operators

			deleteOperatorsOptions := catalogManagementService.NewDeleteOperatorsOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := catalogManagementService.DeleteOperators(deleteOperatorsOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_operators
			fmt.Printf("\nDeleteOperators() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`DeleteOfferingInstance request example`, func() {
			// begin-delete_offering_instance

			deleteOfferingInstanceOptions := catalogManagementService.NewDeleteOfferingInstanceOptions(
				"testString",
				"testString",
			)

			response, err := catalogManagementService.DeleteOfferingInstance(deleteOfferingInstanceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_offering_instance
			fmt.Printf("\nDeleteOfferingInstance() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`DeleteObjectAccessList request example`, func() {
			fmt.Println("\nDeleteObjectAccessList() result:")
			// begin-delete_object_access_list

			deleteObjectAccessListOptions := catalogManagementService.NewDeleteObjectAccessListOptions(
				"testString",
				"testString",
				[]string{"testString"},
			)

			accessListBulkResponse, response, err := catalogManagementService.DeleteObjectAccessList(deleteObjectAccessListOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accessListBulkResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_object_access_list

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessListBulkResponse).ToNot(BeNil())

		})

		It(`DeleteObjectAccess request example`, func() {
			// begin-delete_object_access

			deleteObjectAccessOptions := catalogManagementService.NewDeleteObjectAccessOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := catalogManagementService.DeleteObjectAccess(deleteObjectAccessOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_object_access
			fmt.Printf("\nDeleteObjectAccess() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`DeleteObject request example`, func() {
			// begin-delete_object

			deleteObjectOptions := catalogManagementService.NewDeleteObjectOptions(
				"testString",
				"testString",
			)

			response, err := catalogManagementService.DeleteObject(deleteObjectOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_object
			fmt.Printf("\nDeleteObject() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`DeleteOffering request example`, func() {
			// begin-delete_offering

			deleteOfferingOptions := catalogManagementService.NewDeleteOfferingOptions(
				"testString",
				"testString",
			)

			response, err := catalogManagementService.DeleteOffering(deleteOfferingOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_offering
			fmt.Printf("\nDeleteOffering() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`DeleteCatalog request example`, func() {
			// begin-delete_catalog

			deleteCatalogOptions := catalogManagementService.NewDeleteCatalogOptions(
				"testString",
			)

			response, err := catalogManagementService.DeleteCatalog(deleteCatalogOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_catalog
			fmt.Printf("\nDeleteCatalog() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})
