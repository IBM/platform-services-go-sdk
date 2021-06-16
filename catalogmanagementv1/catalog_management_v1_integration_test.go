// +build integration

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
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"os"
)

/**
 * This file contains an integration test for the catalogmanagementv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`CatalogManagementV1 Integration Tests`, func() {

	const externalConfigFile = "../catalog_mgmt.env"

	var (
		err                                   error
		catalogManagementServiceAuthorized    *catalogmanagementv1.CatalogManagementV1
		catalogManagementServiceNotAuthorized *catalogmanagementv1.CatalogManagementV1
		serviceURL                            string
		config                                map[string]string
		importOfferingZipUrl                  = "https://github.com/rhm-samples/node-red-operator/blob/master/node-red-operator/bundle/0.0.2/node-red-operator.v0.0.2.clusterserviceversion.yaml"
		labelGoSdk                            = "go-sdk"
		kindVPE                               = "vpe"
		kindRoks                              = "roks"
		namespaceGoSDK                        = "java"
		bogusRevision                         = "bogus-revision"
		bogusVersionLocatorID                 = "bogus-version-locator-id"
		repoTypeGitPublic                     = "git_public"
		objectName                            = "object_created_by_go_sdk6"
		regionUSSouth                         = "us-south"
		objectCRN                             = "crn:v1:bluemix:public:iam-global-endpoint:global:::endpoint:private.iam.cloud.ibm.com"
		accountID                             string
		gitAuthToken                          string
		catalogID                             string
		offeringID                            string
		versionLocatorID                      string
		objectID                              string
		clusterID                             string
		refreshTokenAuthorized                string
		refreshTokenNotAuthorized             string
		offeringInstanceID                    string
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
			config, err = core.GetServiceProperties(catalogmanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).NotTo(BeNil())

			gitAuthToken = config["GIT_TOKEN"]
			Expect(gitAuthToken).NotTo(BeNil())

			clusterID = config["CLUSTER_ID"]
			Expect(clusterID).NotTo(BeNil())

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			core.SetLogger(
				core.NewLogger(
					core.LevelDebug,
					log.New(GinkgoWriter, "", log.LstdFlags),
					log.New(GinkgoWriter, "", log.LstdFlags)))

			catalogManagementServiceOptions := &catalogmanagementv1.CatalogManagementV1Options{}
			catalogManagementServiceAuthorized, err = catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(
				catalogManagementServiceOptions)

			Expect(err).To(BeNil())
			Expect(catalogManagementServiceAuthorized).ToNot(BeNil())
			Expect(catalogManagementServiceAuthorized.Service.Options.URL).To(Equal(serviceURL))

			tokenAuthorized, err := catalogManagementServiceAuthorized.Service.Options.Authenticator.(*core.IamAuthenticator).RequestToken()
			Expect(err).To(BeNil())
			refreshTokenAuthorized = tokenAuthorized.RefreshToken
			Expect(refreshTokenAuthorized).ToNot(BeNil())

			catalogManagementUnauthorizedServiceOptions := &catalogmanagementv1.CatalogManagementV1Options{
				ServiceName: "NOT_AUTHORIZED",
			}
			catalogManagementServiceNotAuthorized, err = catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(
				catalogManagementUnauthorizedServiceOptions)
			Expect(err).To(BeNil())
			Expect(catalogManagementServiceNotAuthorized).ToNot(BeNil())
			Expect(catalogManagementServiceNotAuthorized.Service.Options.URL).To(Equal(serviceURL))

			tokenNotAuthorized, err := catalogManagementServiceAuthorized.Service.Options.Authenticator.(*core.IamAuthenticator).RequestToken()
			Expect(err).To(BeNil())
			refreshTokenNotAuthorized = tokenNotAuthorized.RefreshToken
			Expect(refreshTokenNotAuthorized).ToNot(BeNil())

		})
	})

	Describe(`CreateCatalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when user is not authorized`, func() {

			createCatalogOptions := &catalogmanagementv1.CreateCatalogOptions{
				Label:         &labelGoSdk,
				Tags:          []string{"go", "sdk"},
				OwningAccount: &accountID,
				Kind:          &kindVPE,
			}

			_, response, err := catalogManagementServiceNotAuthorized.CreateCatalog(createCatalogOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			createCatalogOptions := &catalogmanagementv1.CreateCatalogOptions{
				Label:         &labelGoSdk,
				Rev:           &bogusRevision,
				Tags:          []string{"go", "sdk"},
				OwningAccount: &accountID,
				Kind:          &kindVPE,
			}

			_, response, err := catalogManagementServiceAuthorized.CreateCatalog(createCatalogOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Creates a catalog`, func() {

			createCatalogOptions := &catalogmanagementv1.CreateCatalogOptions{
				Label:         &labelGoSdk,
				Tags:          []string{"go", "sdk"},
				OwningAccount: &accountID,
				Kind:          &kindVPE,
			}

			result, response, err := catalogManagementServiceAuthorized.CreateCatalog(createCatalogOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(result).NotTo(BeNil())
			Expect(result.ID).NotTo(BeNil())
			catalogID = *result.ID
		})

	})

	Describe(`Get Catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 404 when no such catalog`, func() {
			Expect(catalogID).NotTo(BeNil())

			invalidCatalogId := "invalid-" + catalogID
			getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{
				CatalogIdentifier: &invalidCatalogId,
			}

			_, response, err := catalogManagementServiceAuthorized.GetCatalog(getCatalogOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())

			getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{
				CatalogIdentifier: &catalogID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetCatalog(getCatalogOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns the catalog`, func() {
			Expect(catalogID).NotTo(BeNil())

			getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{
				CatalogIdentifier: &catalogID,
			}

			catalog, response, err := catalogManagementServiceAuthorized.GetCatalog(getCatalogOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).NotTo(BeNil())
			Expect(*catalog.ID).To(Equal(catalogID))
		})
	})

	Describe(`Replace Catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())

			replaceCatalogOptions := &catalogmanagementv1.ReplaceCatalogOptions{
				CatalogIdentifier: &catalogID,
				ID:                &catalogID,
				OwningAccount:     &accountID,
				Kind:              &kindVPE,
			}

			_, response, err := catalogManagementServiceNotAuthorized.ReplaceCatalog(replaceCatalogOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when backend input validation fails`, func() {
			Expect(catalogID).NotTo(BeNil())

			invalidCatalogID := "invalid-" + catalogID
			replaceCatalogOptions := &catalogmanagementv1.ReplaceCatalogOptions{
				CatalogIdentifier: &catalogID,
				ID:                &invalidCatalogID,
				OwningAccount:     &accountID,
				Kind:              &kindVPE,
			}

			_, response, err := catalogManagementServiceAuthorized.ReplaceCatalog(replaceCatalogOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such catalog`, func() {
			Expect(catalogID).NotTo(BeNil())

			invalidCatalogID := "invalid-" + catalogID
			replaceCatalogOptions := &catalogmanagementv1.ReplaceCatalogOptions{
				CatalogIdentifier: &invalidCatalogID,
				ID:                &invalidCatalogID,
				OwningAccount:     &accountID,
				Kind:              &kindVPE,
			}

			_, response, err := catalogManagementServiceAuthorized.ReplaceCatalog(replaceCatalogOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Replaces the catalog`, func() {
			Expect(catalogID).NotTo(BeNil())

			tags := []string{"go", "sdk", "update"}
			replaceCatalogOptions := &catalogmanagementv1.ReplaceCatalogOptions{
				CatalogIdentifier: &catalogID,
				ID:                &catalogID,
				Tags:              tags,
				OwningAccount:     &accountID,
				Kind:              &kindVPE,
			}

			catalog, response, err := catalogManagementServiceAuthorized.ReplaceCatalog(replaceCatalogOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).NotTo(BeNil())
			Expect(catalog.Tags).To(Equal(tags))
		})
	})

	Describe(`List Catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Lists Catalogs`, func() {

			listCatalogsOptions := &catalogmanagementv1.ListCatalogsOptions{}

			catalogSearchResult, response, err := catalogManagementServiceAuthorized.ListCatalogs(listCatalogsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogSearchResult.Resources).ToNot(BeNil())

			contains := false
			for _, ctl := range catalogSearchResult.Resources {
				if *ctl.ID == catalogID {
					contains = true
					break
				}
			}
			Expect(contains).To(BeTrue())
		})
	})

	Describe(`Create Offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 404 when no such catalog`, func() {
			Expect(catalogID).NotTo(BeNil())

			invalidCatalogID := "invalid-" + catalogID
			createOfferingOptions := &catalogmanagementv1.CreateOfferingOptions{
				CatalogIdentifier: &invalidCatalogID,
			}

			_, response, err := catalogManagementServiceAuthorized.CreateOffering(createOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {
			Expect(catalogID).NotTo(BeNil())

			offeringName := "offering created by go sdk"
			createOfferingOptions := &catalogmanagementv1.CreateOfferingOptions{
				CatalogIdentifier: &catalogID,
				CatalogID:         &catalogID,
				Name:              &offeringName,
			}

			_, response, err := catalogManagementServiceAuthorized.CreateOffering(createOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())

			offeringName := "offering-created-by-go-sdk"
			createOfferingOptions := &catalogmanagementv1.CreateOfferingOptions{
				CatalogIdentifier: &catalogID,
				ID:                &catalogID,
				Name:              &offeringName,
			}

			_, response, err := catalogManagementServiceNotAuthorized.CreateOffering(createOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Creates an offering`, func() {
			Expect(catalogID).NotTo(BeNil())

			offeringName := "offering-created-by-go-sdk"
			createOfferingOptions := &catalogmanagementv1.CreateOfferingOptions{
				CatalogIdentifier: &catalogID,
				ID:                &catalogID,
				Name:              &offeringName,
			}

			offering, response, err := catalogManagementServiceAuthorized.CreateOffering(createOfferingOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).NotTo(BeNil())
			Expect(offering.ID).NotTo(BeNil())

			offeringID = *offering.ID
		})
	})

	Describe(`Get Offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 404 when no such offering`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			invalidOfferingID := "invalid-" + offeringID
			getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &invalidOfferingID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOffering(getOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetOffering(getOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns the offering`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
			}

			offering, response, err := catalogManagementServiceAuthorized.GetOffering(getOfferingOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).NotTo(BeNil())
			Expect(*offering.ID).To(Equal(offeringID))
			Expect(*offering.CatalogID).To(Equal(catalogID))
		})
	})

	Describe(`Replace Offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 404 when no such offering`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			invalidOfferingID := "invalid-" + offeringID
			replaceOfferingOptions := &catalogmanagementv1.ReplaceOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &invalidOfferingID,
				ID:                &invalidOfferingID,
				Name:              core.StringPtr("updated-offering-name-by-go-sdk"),
			}

			_, response, err := catalogManagementServiceAuthorized.ReplaceOffering(replaceOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			updatedOfferingName := "updated offering name by go sdk"
			replaceOfferingOptions := &catalogmanagementv1.ReplaceOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				ID:                &offeringID,
				Name:              &updatedOfferingName,
			}

			_, response, err := catalogManagementServiceAuthorized.ReplaceOffering(replaceOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			replaceOfferingOptions := &catalogmanagementv1.ReplaceOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				ID:                &offeringID,
				Name:              core.StringPtr("updated-offering-name-by-go-sdk"),
			}

			_, response, err := catalogManagementServiceNotAuthorized.ReplaceOffering(replaceOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		// once the version related conflict is resolved this test requires a conflict case
		It(`Returns 409 when conflict occurs`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			replaceOfferingOptions := &catalogmanagementv1.ReplaceOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				ID:                &offeringID,
				Name:              core.StringPtr("updated-offering-name-by-go-sdk"),
			}

			_, response, err := catalogManagementServiceAuthorized.ReplaceOffering(replaceOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(409))
		})

		It(`Updates the offering`, func() {
			Skip("Revision conflict.")

			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			updatedOfferingName := "updated-offering-name-by-go-sdk"
			replaceOfferingOptions := &catalogmanagementv1.ReplaceOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				ID:                &offeringID,
				Name:              &updatedOfferingName,
			}

			offering, response, err := catalogManagementServiceAuthorized.ReplaceOffering(replaceOfferingOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).NotTo(BeNil())

			Expect(*offering.ID).To(Equal(offeringID))
			Expect(*offering.CatalogID).To(Equal(catalogID))
			Expect(*offering.Name).To(Equal(updatedOfferingName))
		})
	})

	Describe(`List Offerings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())

			listOfferingsOptions := &catalogmanagementv1.ListOfferingsOptions{
				CatalogIdentifier: &catalogID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.ListOfferings(listOfferingsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when backend input validation fails`, func() {
			Expect(catalogID).NotTo(BeNil())

			listOfferingsOptions := &catalogmanagementv1.ListOfferingsOptions{
				CatalogIdentifier: &catalogID,
				Digest:            core.BoolPtr(true),
				Sort:              core.StringPtr("bogus-sort-value"),
			}

			_, response, err := catalogManagementServiceAuthorized.ListOfferings(listOfferingsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such catalog`, func() {
			Expect(catalogID).NotTo(BeNil())

			invalidCatalogId := "invalid-" + catalogID
			listOfferingsOptions := &catalogmanagementv1.ListOfferingsOptions{
				CatalogIdentifier: &invalidCatalogId,
			}

			_, response, err := catalogManagementServiceAuthorized.ListOfferings(listOfferingsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns list of offerings`, func() {
			Expect(catalogID).NotTo(BeNil())

			var offset int64 = 0
			var limit int64 = 50
			fetch := true
			amountOfOfferings := 0
			contains := false

			for fetch == true {
				listOfferingsOptions := &catalogmanagementv1.ListOfferingsOptions{
					CatalogIdentifier: &catalogID,
					Offset:            &offset,
					Limit:             &limit,
				}

				offeringList, response, err := catalogManagementServiceAuthorized.ListOfferings(listOfferingsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offeringList).NotTo(BeNil())

				if len(offeringList.Resources) > 0 {
					amountOfOfferings += len(offeringList.Resources)
					offset += 50

					if contains == false {
						for _, offering := range offeringList.Resources {
							if *offering.ID == offeringID {
								contains = true
								break
							}
						}
					}

				} else {
					fetch = false
				}
			}
			Expect(contains).To(BeTrue())

			fmt.Printf("Amount of Offerings: %d", amountOfOfferings)
		})
	})

	Describe(`ImportOffering - Import offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			importOfferingOptions := &catalogmanagementv1.ImportOfferingOptions{
				CatalogIdentifier: &catalogID,
				Tags:              []string{"go", "sdk"},
				TargetKinds:       []string{kindVPE},
				Zipurl:            &importOfferingZipUrl,
				OfferingID:        &offeringID,
				TargetVersion:     core.StringPtr("0.0.3"),
				RepoType:          &repoTypeGitPublic,
				XAuthToken:        &gitAuthToken,
			}

			_, response, err := catalogManagementServiceNotAuthorized.ImportOffering(importOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when backend input validation fails`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			importOfferingOptions := &catalogmanagementv1.ImportOfferingOptions{
				CatalogIdentifier: &catalogID,
				Tags:              []string{"go", "sdk"},
				TargetKinds:       []string{"rocks"},
				Zipurl:            &importOfferingZipUrl,
				OfferingID:        &offeringID,
				TargetVersion:     core.StringPtr("0.0.2-patch"),
				RepoType:          &repoTypeGitPublic,
				XAuthToken:        &gitAuthToken,
			}

			_, response, err := catalogManagementServiceAuthorized.ImportOffering(importOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such catalog`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			invalidCatalogId := "invalid-" + catalogID
			importOfferingOptions := &catalogmanagementv1.ImportOfferingOptions{
				CatalogIdentifier: &invalidCatalogId,
				Tags:              []string{"go", "sdk"},
				TargetKinds:       []string{kindRoks},
				Zipurl:            &importOfferingZipUrl,
				OfferingID:        &offeringID,
				TargetVersion:     core.StringPtr("0.0.2"),
				RepoType:          &repoTypeGitPublic,
				XAuthToken:        &gitAuthToken,
			}

			_, response, err := catalogManagementServiceAuthorized.ImportOffering(importOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Imports the offering`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			importOfferingOptions := &catalogmanagementv1.ImportOfferingOptions{
				CatalogIdentifier: &catalogID,
				Tags:              []string{"go", "sdk"},
				TargetKinds:       []string{kindRoks},
				Zipurl:            &importOfferingZipUrl,
				OfferingID:        &offeringID,
				TargetVersion:     core.StringPtr("0.0.2"),
				RepoType:          &repoTypeGitPublic,
				XAuthToken:        &gitAuthToken,
			}

			offering, response, err := catalogManagementServiceAuthorized.ImportOffering(importOfferingOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).NotTo(BeNil())
			Expect(offering.Kinds[0].Versions[0].VersionLocator).NotTo(BeNil())
			versionLocatorID = *offering.Kinds[0].Versions[0].VersionLocator
		})

		It(`Returns 409 when conflict occurs`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			importOfferingOptions := &catalogmanagementv1.ImportOfferingOptions{
				CatalogIdentifier: &catalogID,
				Tags:              []string{"go", "sdk"},
				TargetKinds:       []string{kindRoks},
				Zipurl:            &importOfferingZipUrl,
				OfferingID:        &offeringID,
				TargetVersion:     core.StringPtr("0.0.2"),
				RepoType:          &repoTypeGitPublic,
				XAuthToken:        &gitAuthToken,
			}

			_, response, err := catalogManagementServiceAuthorized.ImportOffering(importOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(409))
		})
	})

	Describe(`Reload Offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 404 when no such offering`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			invalidOfferingId := "invalid-" + offeringID
			reloadOfferingOptions := &catalogmanagementv1.ReloadOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &invalidOfferingId,
				TargetVersion:     core.StringPtr("0.0.2"),
				TargetKinds:       []string{kindRoks},
				Zipurl:            &importOfferingZipUrl,
				RepoType:          &repoTypeGitPublic,
			}

			_, response, err := catalogManagementServiceAuthorized.ReloadOffering(reloadOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			reloadOfferingOptions := &catalogmanagementv1.ReloadOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				TargetVersion:     core.StringPtr("0.0.2"),
				Zipurl:            &importOfferingZipUrl,
				TargetKinds:       []string{kindVPE},
				RepoType:          &repoTypeGitPublic,
			}

			_, response, err := catalogManagementServiceNotAuthorized.ReloadOffering(reloadOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		// Error: Could not find a kind with a target/format value of roks:operator for the current offering, Code: 400
		It(`Reloads the offering`, func() {
			Skip("Could not find kind.")

			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			reloadOfferingOptions := &catalogmanagementv1.ReloadOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				TargetVersion:     core.StringPtr("0.0.2"),
				TargetKinds:       []string{kindRoks},
				Zipurl:            &importOfferingZipUrl,
				RepoType:          &repoTypeGitPublic,
			}

			offering, response, err := catalogManagementServiceNotAuthorized.ReloadOffering(reloadOfferingOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).NotTo(BeNil())
		})
	})

	Describe(`Create Object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 400 when backend input validation fails`, func() {
			Expect(catalogID).NotTo(BeNil())

			publishObjectModel := &catalogmanagementv1.PublishObject{
				PermitIBMPublicPublish: core.BoolPtr(true),
				IBMApproved:            core.BoolPtr(true),
				PublicApproved:         core.BoolPtr(true),
			}

			stateModel := &catalogmanagementv1.State{
				Current: core.StringPtr("new"),
			}

			bogusRegionName := "bogus region name"
			createObjectOptions := &catalogmanagementv1.CreateObjectOptions{
				CatalogIdentifier: &catalogID,
				CatalogID:         &catalogID,
				Name:              &objectName,
				CRN:               &objectCRN,
				ParentID:          &bogusRegionName,
				Kind:              core.StringPtr(kindVPE),
				Publish:           publishObjectModel,
				State:             stateModel,
			}

			_, response, err := catalogManagementServiceAuthorized.CreateObject(createObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())

			publishObjectModel := &catalogmanagementv1.PublishObject{
				PermitIBMPublicPublish: core.BoolPtr(true),
				IBMApproved:            core.BoolPtr(true),
				PublicApproved:         core.BoolPtr(true),
			}

			stateModel := &catalogmanagementv1.State{
				Current: core.StringPtr("new"),
			}

			createObjectOptions := &catalogmanagementv1.CreateObjectOptions{
				CatalogIdentifier: &catalogID,
				CatalogID:         &catalogID,
				Name:              &objectName,
				CRN:               &objectCRN,
				ParentID:          &regionUSSouth,
				Kind:              core.StringPtr(kindVPE),
				Publish:           publishObjectModel,
				State:             stateModel,
			}

			_, response, err := catalogManagementServiceNotAuthorized.CreateObject(createObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such catalog`, func() {
			Expect(catalogID).NotTo(BeNil())

			publishObjectModel := &catalogmanagementv1.PublishObject{
				PermitIBMPublicPublish: core.BoolPtr(true),
				IBMApproved:            core.BoolPtr(true),
				PublicApproved:         core.BoolPtr(true),
			}

			stateModel := &catalogmanagementv1.State{
				Current: core.StringPtr("new"),
			}

			invalidCatalogID := "invalid-" + catalogID
			createObjectOptions := &catalogmanagementv1.CreateObjectOptions{
				CatalogIdentifier: &invalidCatalogID,
				CatalogID:         &invalidCatalogID,
				Name:              &objectName,
				CRN:               &objectCRN,
				ParentID:          &regionUSSouth,
				Kind:              core.StringPtr(kindVPE),
				Publish:           publishObjectModel,
				State:             stateModel,
			}

			_, response, err := catalogManagementServiceAuthorized.CreateObject(createObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Creates an object`, func() {
			Expect(catalogID).NotTo(BeNil())

			publishObjectModel := &catalogmanagementv1.PublishObject{
				PermitIBMPublicPublish: core.BoolPtr(true),
				IBMApproved:            core.BoolPtr(true),
				PublicApproved:         core.BoolPtr(true),
			}

			stateModel := &catalogmanagementv1.State{
				Current: core.StringPtr("new"),
			}

			createObjectOptions := &catalogmanagementv1.CreateObjectOptions{
				CatalogIdentifier: &catalogID,
				CatalogID:         &catalogID,
				Name:              &objectName,
				CRN:               &objectCRN,
				ParentID:          &regionUSSouth,
				Kind:              core.StringPtr(kindVPE),
				Publish:           publishObjectModel,
				State:             stateModel,
			}

			catalogObject, response, err := catalogManagementServiceAuthorized.CreateObject(createObjectOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogObject).NotTo(BeNil())
			Expect(catalogObject.ID).NotTo(BeNil())

			objectID = *catalogObject.ID
		})
	})

	Describe(`Get Offering Audit`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 200 when no such offering`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			invalidOfferingId := "invalid-" + offeringID
			getOfferingAuditOptions := &catalogmanagementv1.GetOfferingAuditOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &invalidOfferingId,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingAudit(getOfferingAuditOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			getOfferingAuditOptions := &catalogmanagementv1.GetOfferingAuditOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetOfferingAudit(getOfferingAuditOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns offering audit logs`, func() {
			Expect(catalogID).NotTo(BeNil())
			Expect(offeringID).NotTo(BeNil())

			getOfferingAuditOptions := &catalogmanagementv1.GetOfferingAuditOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
			}

			auditLog, response, err := catalogManagementServiceAuthorized.GetOfferingAudit(getOfferingAuditOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditLog).NotTo(BeNil())
		})
	})

	Describe(`Get Catalog Account`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns the catalog account`, func() {

			getCatalogAccountOptions := &catalogmanagementv1.GetCatalogAccountOptions{}

			account, response, err := catalogManagementServiceAuthorized.GetCatalogAccount(getCatalogAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(account).ToNot(BeNil())
			Expect(*account.ID).To(Equal(accountID))
		})
	})

	Describe(`Update Catalog Account`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when no such account`, func() {
			Expect(accountID).NotTo(BeNil())

			invalidAccountID := "invalid-" + accountID
			updateCatalogAccountOptions := &catalogmanagementv1.UpdateCatalogAccountOptions{
				ID: &invalidAccountID,
			}

			response, err := catalogManagementServiceAuthorized.UpdateCatalogAccount(updateCatalogAccountOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(accountID).NotTo(BeNil())

			updateCatalogAccountOptions := &catalogmanagementv1.UpdateCatalogAccountOptions{
				ID: &accountID,
			}

			response, err := catalogManagementServiceNotAuthorized.UpdateCatalogAccount(updateCatalogAccountOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when backend input validation fails`, func() {
			Skip("User is not granted.")
			Expect(accountID).NotTo(BeNil())

			// a body with failing data comes here

			updateCatalogAccountOptions := &catalogmanagementv1.UpdateCatalogAccountOptions{
				ID: &accountID,
			}

			response, err := catalogManagementServiceNotAuthorized.UpdateCatalogAccount(updateCatalogAccountOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Updates catalog account`, func() {
			Skip("User is not granted.")
			Expect(accountID).NotTo(BeNil())

			// data

			updateCatalogAccountOptions := &catalogmanagementv1.UpdateCatalogAccountOptions{
				ID: &accountID,
			}

			response, err := catalogManagementServiceNotAuthorized.UpdateCatalogAccount(updateCatalogAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(response.Result).NotTo(BeNil())
		})
	})

	Describe(`Get Catalog Account Audit`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			getCatalogAccountAuditOptions := &catalogmanagementv1.GetCatalogAccountAuditOptions{}

			_, response, err := catalogManagementServiceNotAuthorized.GetCatalogAccountAudit(getCatalogAccountAuditOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns catalog account audit logs`, func() {

			getCatalogAccountAuditOptions := &catalogmanagementv1.GetCatalogAccountAuditOptions{}

			catalogAccountAuditLogs, response, err := catalogManagementServiceAuthorized.GetCatalogAccountAudit(getCatalogAccountAuditOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogAccountAuditLogs).NotTo(BeNil())
		})
	})

	Describe(`Get Catalog Account Filters`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).ToNot(BeNil())

			getCatalogAccountFiltersOptions := &catalogmanagementv1.GetCatalogAccountFiltersOptions{
				Catalog: &catalogID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetCatalogAccountFilters(getCatalogAccountFiltersOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such catalog`, func() {
			Expect(catalogID).ToNot(BeNil())

			invalidCatalogID := "invalid-" + catalogID
			getCatalogAccountFiltersOptions := &catalogmanagementv1.GetCatalogAccountFiltersOptions{
				Catalog: &invalidCatalogID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetCatalogAccountFilters(getCatalogAccountFiltersOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns the account filters`, func() {
			Expect(catalogID).ToNot(BeNil())

			getCatalogAccountFiltersOptions := &catalogmanagementv1.GetCatalogAccountFiltersOptions{
				Catalog: &catalogID,
			}

			accumulatedFilters, response, err := catalogManagementServiceAuthorized.GetCatalogAccountFilters(getCatalogAccountFiltersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accumulatedFilters).NotTo(BeNil())
		})
	})

	Describe(`Get Catalog Audit - Get catalog audit log`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 404 when no such catalog`, func() {
			Expect(catalogID).ToNot(BeNil())

			invalidCatalogID := "invalid-" + catalogID
			getCatalogAuditOptions := &catalogmanagementv1.GetCatalogAuditOptions{
				CatalogIdentifier: &invalidCatalogID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetCatalogAudit(getCatalogAuditOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).ToNot(BeNil())

			getCatalogAuditOptions := &catalogmanagementv1.GetCatalogAuditOptions{
				CatalogIdentifier: &catalogID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetCatalogAudit(getCatalogAuditOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns the catalog audit logs`, func() {
			Expect(catalogID).ToNot(BeNil())

			getCatalogAuditOptions := &catalogmanagementv1.GetCatalogAuditOptions{
				CatalogIdentifier: &catalogID,
			}

			catalogAuditLogs, response, err := catalogManagementServiceAuthorized.GetCatalogAudit(getCatalogAuditOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogAuditLogs).NotTo(BeNil())
		})
	})

	Describe(`Get Consumption Offerings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).ToNot(BeNil())

			getConsumptionOfferingsOptions := &catalogmanagementv1.GetConsumptionOfferingsOptions{
				Catalog: &catalogID,
				Select:  core.StringPtr("all"),
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetConsumptionOfferings(getConsumptionOfferingsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such catalog`, func() {
			Expect(catalogID).ToNot(BeNil())

			invalidCatalogID := "invalid-" + catalogID
			getConsumptionOfferingsOptions := &catalogmanagementv1.GetConsumptionOfferingsOptions{
				Catalog: &invalidCatalogID,
				Select:  core.StringPtr("all"),
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetConsumptionOfferings(getConsumptionOfferingsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns the catalog consumption offerings`, func() {
			Expect(catalogID).ToNot(BeNil())

			getConsumptionOfferingsOptions := &catalogmanagementv1.GetConsumptionOfferingsOptions{
				Catalog: &catalogID,
				Select:  core.StringPtr("all"),
			}

			_, response, err := catalogManagementServiceAuthorized.GetConsumptionOfferings(getConsumptionOfferingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(response.Result).ToNot(BeNil())
		})
	})

	Describe(`Import Offering Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {
			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			importOfferingVersionOptions := &catalogmanagementv1.ImportOfferingVersionOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				TargetKinds:       []string{"rocks"},
				Zipurl:            &importOfferingZipUrl,
				TargetVersion:     core.StringPtr("0.0.3"),
				RepoType:          &repoTypeGitPublic,
			}

			_, response, err := catalogManagementServiceAuthorized.ImportOfferingVersion(importOfferingVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such offering`, func() {
			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			invalidOfferingID := "invalid-" + offeringID
			importOfferingVersionOptions := &catalogmanagementv1.ImportOfferingVersionOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &invalidOfferingID,
				TargetKinds:       []string{kindRoks},
				Zipurl:            &importOfferingZipUrl,
				TargetVersion:     core.StringPtr("0.0.3"),
				RepoType:          &repoTypeGitPublic,
			}

			_, response, err := catalogManagementServiceAuthorized.ImportOfferingVersion(importOfferingVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			importOfferingVersionOptions := &catalogmanagementv1.ImportOfferingVersionOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				TargetKinds:       []string{kindRoks},
				Zipurl:            &importOfferingZipUrl,
				TargetVersion:     core.StringPtr("0.0.3"),
				RepoType:          &repoTypeGitPublic,
			}

			_, response, err := catalogManagementServiceNotAuthorized.ImportOfferingVersion(importOfferingVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns imported offering version`, func() {
			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			importOfferingVersionOptions := &catalogmanagementv1.ImportOfferingVersionOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				TargetKinds:       []string{kindRoks},
				Zipurl:            &importOfferingZipUrl,
				TargetVersion:     core.StringPtr("0.0.3"),
				RepoType:          &repoTypeGitPublic,
			}

			offeringResult, response, err := catalogManagementServiceAuthorized.ImportOfferingVersion(importOfferingVersionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offeringResult).NotTo(BeNil())
		})
	})

	Describe(`Replace Offering Icon`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 404 when no such offerings`, func() {
			Skip("This functionality is disabled.")

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			invalidOfferingID := "invalid-" + offeringID
			replaceOfferingIconOptions := &catalogmanagementv1.ReplaceOfferingIconOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &invalidOfferingID,
				FileName:          core.StringPtr("filename.jpg"),
			}

			_, response, err := catalogManagementServiceAuthorized.ReplaceOfferingIcon(replaceOfferingIconOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Skip("This functionality is disabled.")

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			replaceOfferingIconOptions := &catalogmanagementv1.ReplaceOfferingIconOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				FileName:          core.StringPtr("filename.jpg"),
			}

			_, response, err := catalogManagementServiceNotAuthorized.ReplaceOfferingIcon(replaceOfferingIconOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Replaces the Offerings Icon`, func() {
			Skip("This functionality is disabled.")

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			replaceOfferingIconOptions := &catalogmanagementv1.ReplaceOfferingIconOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				FileName:          core.StringPtr("filename.jpg"),
			}

			offering, response, err := catalogManagementServiceAuthorized.ReplaceOfferingIcon(replaceOfferingIconOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).NotTo(BeNil())
		})

	})

	Describe(`Update Offering IBM`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {
			Skip("Once the user is granted for this operation this test can be executed.")

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			updateOfferingIBMOptions := &catalogmanagementv1.UpdateOfferingIBMOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				ApprovalType:      core.StringPtr("fake approval type"),
				Approved:          core.StringPtr("true"),
			}

			_, response, err := catalogManagementServiceAuthorized.UpdateOfferingIBM(updateOfferingIBMOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such offering`, func() {
			Skip("Once the user is granted for this operation this test can be executed.")

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			invalidOfferingID := "invalid-" + offeringID
			updateOfferingIBMOptions := &catalogmanagementv1.UpdateOfferingIBMOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &invalidOfferingID,
				ApprovalType:      core.StringPtr("allow_request"),
				Approved:          core.StringPtr("true"),
			}

			_, response, err := catalogManagementServiceAuthorized.UpdateOfferingIBM(updateOfferingIBMOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			updateOfferingIBMOptions := &catalogmanagementv1.UpdateOfferingIBMOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				ApprovalType:      core.StringPtr("allow_request"),
				Approved:          core.StringPtr("true"),
			}

			_, response, err := catalogManagementServiceNotAuthorized.UpdateOfferingIBM(updateOfferingIBMOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Updates offering IBM`, func() {
			Skip("This user is not permitted to execute this operation.")

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			updateOfferingIBMOptions := &catalogmanagementv1.UpdateOfferingIBMOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				ApprovalType:      core.StringPtr("allow_request"),
				Approved:          core.StringPtr("true"),
			}

			approvalResult, response, err := catalogManagementServiceAuthorized.UpdateOfferingIBM(updateOfferingIBMOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(approvalResult).NotTo(BeNil())
		})
	})

	Describe(`Get Offering Updates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			getOfferingUpdatesOptions := &catalogmanagementv1.GetOfferingUpdatesOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				Kind:              core.StringPtr("rocks"),
				Version:           core.StringPtr("0.0.2"),
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingUpdates(getOfferingUpdatesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such offerings`, func() {

			Skip("It requires some special offering type.")

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			// it always complaining about offering types which is somehow related to create/import offerings
			// once this is resolved there is a chance we can squeeze a 404 out from the service

			invalidOfferingID := "invalid-" + offeringID
			getOfferingUpdatesOptions := &catalogmanagementv1.GetOfferingUpdatesOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &invalidOfferingID,
				Kind:              core.StringPtr(kindVPE),
				Version:           core.StringPtr("0.0.2"),
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				Namespace:         &namespaceGoSDK,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingUpdates(getOfferingUpdatesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			getOfferingUpdatesOptions := &catalogmanagementv1.GetOfferingUpdatesOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				Kind:              core.StringPtr(kindRoks),
				Version:           core.StringPtr("0.0.2"),
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				Namespace:         &namespaceGoSDK,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetOfferingUpdates(getOfferingUpdatesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns offering updates`, func() {
			Skip("Requires an offering type different than helm, roks or vpe")

			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			getOfferingUpdatesOptions := &catalogmanagementv1.GetOfferingUpdatesOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
				Kind:              core.StringPtr(kindVPE),
				Version:           core.StringPtr("0.0.2"),
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				Namespace:         &namespaceGoSDK,
			}

			versionUpdateDescriptor, response, err := catalogManagementServiceAuthorized.GetOfferingUpdates(getOfferingUpdatesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(versionUpdateDescriptor).NotTo(BeNil())
		})
	})

	Describe(`Get Offering About`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 400 when backend input validation fails`, func() {

			getOfferingAboutOptions := &catalogmanagementv1.GetOfferingAboutOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingAbout(getOfferingAboutOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such version`, func() {
			Expect(versionLocatorID).ToNot(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			getOfferingAboutOptions := &catalogmanagementv1.GetOfferingAboutOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingAbout(getOfferingAboutOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).ToNot(BeNil())

			getOfferingAboutOptions := &catalogmanagementv1.GetOfferingAboutOptions{
				VersionLocID: &versionLocatorID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetOfferingAbout(getOfferingAboutOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns Offering About`, func() {

			Expect(versionLocatorID).ToNot(BeNil())

			getOfferingAboutOptions := &catalogmanagementv1.GetOfferingAboutOptions{
				VersionLocID: &versionLocatorID,
			}

			offeringAbout, response, err := catalogManagementServiceAuthorized.GetOfferingAbout(getOfferingAboutOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringAbout).NotTo(BeNil())
		})
	})

	Describe(`Get Offering License`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			getOfferingLicenseOptions := &catalogmanagementv1.GetOfferingLicenseOptions{
				VersionLocID: &bogusVersionLocatorID,
				LicenseID:    core.StringPtr("license-id-is-needed"),
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingLicense(getOfferingLicenseOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such version`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			getOfferingLicenseOptions := &catalogmanagementv1.GetOfferingLicenseOptions{
				VersionLocID: &invalidVersionLocatorID,
				LicenseID:    core.StringPtr("license-id-is-needed"),
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingLicense(getOfferingLicenseOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Skip("No license.")

			getOfferingLicenseOptions := &catalogmanagementv1.GetOfferingLicenseOptions{
				VersionLocID: &versionLocatorID,
				LicenseID:    core.StringPtr("license-id-is-needed"),
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetOfferingLicense(getOfferingLicenseOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns the offering license`, func() {
			Skip("No license.")

			Expect(versionLocatorID).NotTo(BeNil())

			getOfferingLicenseOptions := &catalogmanagementv1.GetOfferingLicenseOptions{
				VersionLocID: &versionLocatorID,
				LicenseID:    core.StringPtr("license-id-is-needed"),
			}

			offeringLicense, response, err := catalogManagementServiceAuthorized.GetOfferingLicense(getOfferingLicenseOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringLicense).NotTo(BeNil())
		})
	})

	Describe(`Get Offering Container Images`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 400 when backend input validation fails`, func() {

			getOfferingContainerImagesOptions := &catalogmanagementv1.GetOfferingContainerImagesOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingContainerImages(getOfferingContainerImagesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such version`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			getOfferingContainerImagesOptions := &catalogmanagementv1.GetOfferingContainerImagesOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingContainerImages(getOfferingContainerImagesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			getOfferingContainerImagesOptions := &catalogmanagementv1.GetOfferingContainerImagesOptions{
				VersionLocID: &versionLocatorID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetOfferingContainerImages(getOfferingContainerImagesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns offering container images`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			getOfferingContainerImagesOptions := &catalogmanagementv1.GetOfferingContainerImagesOptions{
				VersionLocID: &versionLocatorID,
			}

			containerImageManifest, response, err := catalogManagementServiceAuthorized.GetOfferingContainerImages(getOfferingContainerImagesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(containerImageManifest).NotTo(BeNil())
		})
	})

	Describe(`Deprecate Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 404 when no such version`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			deprecateVersionOptions := &catalogmanagementv1.DeprecateVersionOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.DeprecateVersion(deprecateVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			deprecateVersionOptions := &catalogmanagementv1.DeprecateVersionOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.DeprecateVersion(deprecateVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			deprecateVersionOptions := &catalogmanagementv1.DeprecateVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceNotAuthorized.DeprecateVersion(deprecateVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Deprecates the version`, func() {
			Skip("Order of states is needed")

			Expect(versionLocatorID).NotTo(BeNil())

			deprecateVersionOptions := &catalogmanagementv1.DeprecateVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.DeprecateVersion(deprecateVersionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`Account Publish Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			accountPublishVersionOptions := &catalogmanagementv1.AccountPublishVersionOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.AccountPublishVersion(accountPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such version`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			accountPublishVersionOptions := &catalogmanagementv1.AccountPublishVersionOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.AccountPublishVersion(accountPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			accountPublishVersionOptions := &catalogmanagementv1.AccountPublishVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceNotAuthorized.AccountPublishVersion(accountPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Publishes version`, func() {
			Skip("Order of states is needed")
			Expect(versionLocatorID).NotTo(BeNil())

			accountPublishVersionOptions := &catalogmanagementv1.AccountPublishVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.AccountPublishVersion(accountPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`IBM Publish Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			ibmPublishVersionOptions := &catalogmanagementv1.IBMPublishVersionOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.IBMPublishVersion(ibmPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such version`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			ibmPublishVersionOptions := &catalogmanagementv1.IBMPublishVersionOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.IBMPublishVersion(ibmPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			ibmPublishVersionOptions := &catalogmanagementv1.IBMPublishVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceNotAuthorized.IBMPublishVersion(ibmPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Publishes the version`, func() {
			Skip("User is not granted.")
			Expect(versionLocatorID).NotTo(BeNil())

			ibmPublishVersionOptions := &catalogmanagementv1.IBMPublishVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.IBMPublishVersion(ibmPublishVersionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`Public Publish Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			publicPublishVersionOptions := &catalogmanagementv1.PublicPublishVersionOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.PublicPublishVersion(publicPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such version`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			publicPublishVersionOptions := &catalogmanagementv1.PublicPublishVersionOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.PublicPublishVersion(publicPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			publicPublishVersionOptions := &catalogmanagementv1.PublicPublishVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceNotAuthorized.PublicPublishVersion(publicPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Publishes the version`, func() {
			Skip("User is not granted.")
			Expect(versionLocatorID).NotTo(BeNil())

			publicPublishVersionOptions := &catalogmanagementv1.PublicPublishVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.PublicPublishVersion(publicPublishVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`Commit Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			commitVersionOptions := &catalogmanagementv1.CommitVersionOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.CommitVersion(commitVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such version`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			commitVersionOptions := &catalogmanagementv1.CommitVersionOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.CommitVersion(commitVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			commitVersionOptions := &catalogmanagementv1.CommitVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceNotAuthorized.CommitVersion(commitVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Commits version`, func() {
			Skip("Workflow of versions")

			Expect(versionLocatorID).NotTo(BeNil())

			commitVersionOptions := &catalogmanagementv1.CommitVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.CommitVersion(commitVersionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`Copy Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			copyVersionOptions := &catalogmanagementv1.CopyVersionOptions{
				VersionLocID: &versionLocatorID,
				TargetKinds:  []string{kindRoks},
			}

			response, err := catalogManagementServiceNotAuthorized.CopyVersion(copyVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such version`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorId := "invalid-" + versionLocatorID
			copyVersionOptions := &catalogmanagementv1.CopyVersionOptions{
				VersionLocID: &invalidVersionLocatorId,
				TargetKinds:  []string{kindRoks},
			}

			response, err := catalogManagementServiceAuthorized.CopyVersion(copyVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend validation fails`, func() {

			copyVersionOptions := &catalogmanagementv1.CopyVersionOptions{
				VersionLocID: &bogusVersionLocatorID,
				TargetKinds:  []string{kindRoks},
			}

			response, err := catalogManagementServiceAuthorized.CopyVersion(copyVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Copies a version`, func() {
			Skip("Only for helm, but helm is not supported.")
			Expect(versionLocatorID).NotTo(BeNil())

			copyVersionOptions := &catalogmanagementv1.CopyVersionOptions{
				VersionLocID: &versionLocatorID,
				TargetKinds:  []string{kindRoks},
			}

			response, err := catalogManagementServiceAuthorized.CopyVersion(copyVersionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`Get Offering Working Copy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			getOfferingWorkingCopyOptions := &catalogmanagementv1.GetOfferingWorkingCopyOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingWorkingCopy(getOfferingWorkingCopyOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			getOfferingWorkingCopyOptions := &catalogmanagementv1.GetOfferingWorkingCopyOptions{
				VersionLocID: &versionLocatorID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetOfferingWorkingCopy(getOfferingWorkingCopyOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such version`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			getOfferingWorkingCopyOptions := &catalogmanagementv1.GetOfferingWorkingCopyOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingWorkingCopy(getOfferingWorkingCopyOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns the offering working copy`, func() {
			Skip("requires published state which this user cannot create")
			Expect(versionLocatorID).ToNot(BeNil())

			getOfferingWorkingCopyOptions := &catalogmanagementv1.GetOfferingWorkingCopyOptions{
				VersionLocID: &versionLocatorID,
			}

			version, response, err := catalogManagementServiceAuthorized.GetOfferingWorkingCopy(getOfferingWorkingCopyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(version).ToNot(BeNil())
		})
	})

	Describe(`Get Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			getVersionOptions := &catalogmanagementv1.GetVersionOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetVersion(getVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such version`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			getVersionOptions := &catalogmanagementv1.GetVersionOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetVersion(getVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			getVersionOptions := &catalogmanagementv1.GetVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetVersion(getVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns the offering version`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			getVersionOptions := &catalogmanagementv1.GetVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			offeringVersion, response, err := catalogManagementServiceAuthorized.GetVersion(getVersionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringVersion).NotTo(BeNil())
		})
	})

	Describe(`Get Cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Skip("possibly this user doesn't have right to execute this operation")

			getClusterOptions := &catalogmanagementv1.GetClusterOptions{
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				XAuthRefreshToken: &refreshTokenNotAuthorized,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetCluster(getClusterOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such cluster`, func() {

			invalidClusterID := "invalid-" + clusterID
			getClusterOptions := &catalogmanagementv1.GetClusterOptions{
				ClusterID:         &invalidClusterID,
				Region:            &regionUSSouth,
				XAuthRefreshToken: &refreshTokenAuthorized,
			}

			_, response, err := catalogManagementServiceAuthorized.GetCluster(getClusterOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns the cluster details`, func() {

			Skip("possibly this user doesn't have right to execute this operation")

			getClusterOptions := &catalogmanagementv1.GetClusterOptions{
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				XAuthRefreshToken: &refreshTokenAuthorized,
			}

			clusterInfo, response, err := catalogManagementServiceAuthorized.GetCluster(getClusterOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clusterInfo).NotTo(BeNil())
		})
	})

	Describe(`Get Namespaces`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 404 when no such cluster`, func() {

			invalidClusterID := "invalid-" + clusterID
			getNamespacesOptions := &catalogmanagementv1.GetNamespacesOptions{
				ClusterID:         &invalidClusterID,
				Region:            &regionUSSouth,
				XAuthRefreshToken: &refreshTokenAuthorized,
			}

			_, response, err := catalogManagementServiceAuthorized.GetNamespaces(getNamespacesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 401 when user is not authorized`, func() {

			Skip("It returns randomly either 401 and 404, so it is skipped.")

			getNamespacesOptions := &catalogmanagementv1.GetNamespacesOptions{
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				XAuthRefreshToken: &refreshTokenNotAuthorized,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetNamespaces(getNamespacesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns namespaces`, func() {
			Skip("Possibly the user is not granted.")

			getNamespacesOptions := &catalogmanagementv1.GetNamespacesOptions{
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				XAuthRefreshToken: &refreshTokenAuthorized,
			}

			namespacesResult, response, err := catalogManagementServiceAuthorized.GetNamespaces(getNamespacesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(namespacesResult).NotTo(BeNil())
		})
	})

	Describe(`Deploy Operators`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			deployOperatorsOptions := &catalogmanagementv1.DeployOperatorsOptions{
				XAuthRefreshToken: &refreshTokenNotAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				AllNamespaces:     core.BoolPtr(true),
				VersionLocatorID:  &versionLocatorID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.DeployOperators(deployOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such cluster`, func() {

			invalidClusterID := "invalid-" + clusterID
			deployOperatorsOptions := &catalogmanagementv1.DeployOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &invalidClusterID,
				Region:            &regionUSSouth,
				AllNamespaces:     core.BoolPtr(true),
				VersionLocatorID:  &versionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.DeployOperators(deployOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			deployOperatorsOptions := &catalogmanagementv1.DeployOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				AllNamespaces:     core.BoolPtr(true),
				VersionLocatorID:  &bogusVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.DeployOperators(deployOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Deploys operator`, func() {
			Skip("Possibly the user is not granted.")
			Expect(versionLocatorID).NotTo(BeNil())

			deployOperatorsOptions := &catalogmanagementv1.DeployOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				AllNamespaces:     core.BoolPtr(true),
				VersionLocatorID:  &versionLocatorID,
			}

			operatorDeployResult, response, err := catalogManagementServiceAuthorized.DeployOperators(deployOperatorsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatorDeployResult).NotTo(BeNil())
		})
	})

	Describe(`List Operators`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			listOperatorsOptions := &catalogmanagementv1.ListOperatorsOptions{
				XAuthRefreshToken: &refreshTokenNotAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.ListOperators(listOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			listOperatorsOptions := &catalogmanagementv1.ListOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &bogusVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.ListOperators(listOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such cluster`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			invalidClusterID := "invalid-" + clusterID
			listOperatorsOptions := &catalogmanagementv1.ListOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &invalidClusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.ListOperators(listOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns list of operators`, func() {
			Skip("Possibly this user is not granted.")

			Expect(versionLocatorID).NotTo(BeNil())

			listOperatorsOptions := &catalogmanagementv1.ListOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			listOperatorsResult, response, err := catalogManagementServiceAuthorized.ListOperators(listOperatorsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listOperatorsResult).NotTo(BeNil())
		})
	})

	Describe(`Replace Operators`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			replaceOperatorsOptions := &catalogmanagementv1.ReplaceOperatorsOptions{
				XAuthRefreshToken: &refreshTokenNotAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				AllNamespaces:     core.BoolPtr(true),
				VersionLocatorID:  &versionLocatorID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.ReplaceOperators(replaceOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such cluster`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidClusterID := "invalid-" + clusterID
			replaceOperatorsOptions := &catalogmanagementv1.ReplaceOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &invalidClusterID,
				Region:            &regionUSSouth,
				AllNamespaces:     core.BoolPtr(true),
				VersionLocatorID:  &versionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.ReplaceOperators(replaceOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			replaceOperatorsOptions := &catalogmanagementv1.ReplaceOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				AllNamespaces:     core.BoolPtr(true),
				VersionLocatorID:  &bogusVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.ReplaceOperators(replaceOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Replaces operator`, func() {
			Skip("Possibly this user is not granted.")

			Expect(versionLocatorID).NotTo(BeNil())

			replaceOperatorsOptions := &catalogmanagementv1.ReplaceOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				AllNamespaces:     core.BoolPtr(true),
				VersionLocatorID:  &versionLocatorID,
			}

			operatorDeployResult, response, err := catalogManagementServiceAuthorized.ReplaceOperators(replaceOperatorsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatorDeployResult).NotTo(BeNil())
		})
	})

	Describe(`Install Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			installVersionOptions := &catalogmanagementv1.InstallVersionOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenNotAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			response, err := catalogManagementServiceNotAuthorized.InstallVersion(installVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such cluster`, func() {
			Expect(versionLocatorID).NotTo(BeNil())

			invalidClusterId := "invalid-" + clusterID
			installVersionOptions := &catalogmanagementv1.InstallVersionOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &invalidClusterId,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.InstallVersion(installVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			installVersionOptions := &catalogmanagementv1.InstallVersionOptions{
				VersionLocID:      &bogusVersionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &bogusVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.InstallVersion(installVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Installs the version`, func() {
			Skip("Possibly this user is not granted.")

			Expect(versionLocatorID).NotTo(BeNil())

			installVersionOptions := &catalogmanagementv1.InstallVersionOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.InstallVersion(installVersionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`Preinstall Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			preinstallVersionOptions := &catalogmanagementv1.PreinstallVersionOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenNotAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			response, err := catalogManagementServiceNotAuthorized.PreinstallVersion(preinstallVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such cluster`, func() {
			Skip("Requires preinstall script.")

			Expect(versionLocatorID).NotTo(BeNil())

			// it requires a version where preinstall script is installed
			// but I don't know how to do it
			// one it is done possible to squeeze a 404 from the cluster
			// until then it checks 400

			invalidClusterID := "invalid-" + clusterID
			preinstallVersionOptions := &catalogmanagementv1.PreinstallVersionOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &invalidClusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.PreinstallVersion(preinstallVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			preinstallVersionOptions := &catalogmanagementv1.PreinstallVersionOptions{
				VersionLocID:      &bogusVersionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &bogusVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.PreinstallVersion(preinstallVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Preinstalls the version`, func() {

			Skip("Requires pre-install script.")

			Expect(versionLocatorID).NotTo(BeNil())

			preinstallVersionOptions := &catalogmanagementv1.PreinstallVersionOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.PreinstallVersion(preinstallVersionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`Get Preinstall`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			getPreinstallOptions := &catalogmanagementv1.GetPreinstallOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenNotAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetPreinstall(getPreinstallOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such version`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			getPreinstallOptions := &catalogmanagementv1.GetPreinstallOptions{
				VersionLocID:      &invalidVersionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
			}

			_, response, err := catalogManagementServiceAuthorized.GetPreinstall(getPreinstallOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			getPreinstallOptions := &catalogmanagementv1.GetPreinstallOptions{
				VersionLocID:      &bogusVersionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
			}

			_, response, err := catalogManagementServiceAuthorized.GetPreinstall(getPreinstallOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns preinstall`, func() {
			Skip("Pre-install script is required.")

			Expect(versionLocatorID).NotTo(BeNil())

			getPreinstallOptions := &catalogmanagementv1.GetPreinstallOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
			}

			installStatusResult, response, err := catalogManagementServiceAuthorized.GetPreinstall(getPreinstallOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(installStatusResult).NotTo(BeNil())
		})
	})

	Describe(`Validate Install`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			validateInstallOptions := &catalogmanagementv1.ValidateInstallOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenNotAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			response, err := catalogManagementServiceNotAuthorized.ValidateInstall(validateInstallOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such version`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			validateInstallOptions := &catalogmanagementv1.ValidateInstallOptions{
				VersionLocID:      &invalidVersionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &invalidVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.ValidateInstall(validateInstallOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			validateInstallOptions := &catalogmanagementv1.ValidateInstallOptions{
				VersionLocID:      &bogusVersionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &bogusVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.ValidateInstall(validateInstallOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Validates install`, func() {
			Skip("Possibly this user is not granted.")

			Expect(versionLocatorID).NotTo(BeNil())

			validateInstallOptions := &catalogmanagementv1.ValidateInstallOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.ValidateInstall(validateInstallOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`Get Validation Status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			getValidationStatusOptions := &catalogmanagementv1.GetValidationStatusOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenNotAuthorized,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetValidationStatus(getValidationStatusOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such version`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			getValidationStatusOptions := &catalogmanagementv1.GetValidationStatusOptions{
				VersionLocID:      &invalidVersionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
			}

			_, response, err := catalogManagementServiceAuthorized.GetValidationStatus(getValidationStatusOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			getValidationStatusOptions := &catalogmanagementv1.GetValidationStatusOptions{
				VersionLocID:      &bogusVersionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
			}

			_, response, err := catalogManagementServiceAuthorized.GetValidationStatus(getValidationStatusOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns validation status`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			getValidationStatusOptions := &catalogmanagementv1.GetValidationStatusOptions{
				VersionLocID:      &versionLocatorID,
				XAuthRefreshToken: &refreshTokenAuthorized,
			}

			validationStatus, response, err := catalogManagementServiceAuthorized.GetValidationStatus(getValidationStatusOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(validationStatus).NotTo(BeNil())
		})
	})

	Describe(`Get Override Values`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			getOverrideValuesOptions := &catalogmanagementv1.GetOverrideValuesOptions{
				VersionLocID: &versionLocatorID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetOverrideValues(getOverrideValuesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such version`, func() {

			Expect(versionLocatorID).NotTo(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			getOverrideValuesOptions := &catalogmanagementv1.GetOverrideValuesOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOverrideValues(getOverrideValuesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			getOverrideValuesOptions := &catalogmanagementv1.GetOverrideValuesOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOverrideValues(getOverrideValuesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns override values`, func() {
			Skip("requires validation run before this operation")

			Expect(versionLocatorID).NotTo(BeNil())

			getOverrideValuesOptions := &catalogmanagementv1.GetOverrideValuesOptions{
				VersionLocID: &versionLocatorID,
			}

			result, response, err := catalogManagementServiceAuthorized.GetOverrideValues(getOverrideValuesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).NotTo(BeNil())
		})
	})

	Describe(`Search Objects`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			searchObjectsOptions := &catalogmanagementv1.SearchObjectsOptions{
				Query:    core.StringPtr(""),
				Collapse: core.BoolPtr(true),
				Digest:   core.BoolPtr(true),
			}

			_, response, err := catalogManagementServiceAuthorized.SearchObjects(searchObjectsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 200 when user is not authorized`, func() {

			searchObjectsOptions := &catalogmanagementv1.SearchObjectsOptions{
				Query:    core.StringPtr("name: " + objectName),
				Collapse: core.BoolPtr(true),
				Digest:   core.BoolPtr(true),
			}

			searchResult, response, err := catalogManagementServiceNotAuthorized.SearchObjects(searchObjectsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(searchResult).NotTo(BeNil())
		})

		It(`Returns objects`, func() {

			var offset int64 = 0
			var limit int64 = 0
			fetch := true
			amountOfObjects := 0

			for fetch == true {
				searchObjectsOptions := &catalogmanagementv1.SearchObjectsOptions{
					Query:    core.StringPtr("name: object*"),
					Collapse: core.BoolPtr(true),
					Digest:   core.BoolPtr(true),
					Limit:    &limit,
					Offset:   &offset,
				}

				searchResult, response, err := catalogManagementServiceAuthorized.SearchObjects(searchObjectsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(searchResult).NotTo(BeNil())

				if len(searchResult.Resources) > 0 {
					amountOfObjects += len(searchResult.Resources)
					offset += 50
				} else {
					fetch = false
				}
			}

			fmt.Printf("Amount of objects: %d", amountOfObjects)
		})
	})

	Describe(`List Objects`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			Expect(catalogID).NotTo(BeNil())

			listObjectsOptions := &catalogmanagementv1.ListObjectsOptions{
				CatalogIdentifier: &catalogID,
				Name:              core.StringPtr(""),
				Sort:              core.StringPtr(""),
			}

			_, response, err := catalogManagementServiceAuthorized.ListObjects(listObjectsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {

			Skip("It returns the result for some reason, however python doesn't. Skipped.")
			Expect(catalogID).NotTo(BeNil())

			listObjectsOptions := &catalogmanagementv1.ListObjectsOptions{
				CatalogIdentifier: &catalogID,
			}

			_, response, err := catalogManagementServiceAuthorized.ListObjects(listObjectsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns list of objects`, func() {

			Expect(catalogID).NotTo(BeNil())

			var limit int64 = 50
			var offset int64 = 0
			amountOfObjects := 0
			contains := false
			fetch := true

			for fetch == true {

				listObjectsOptions := &catalogmanagementv1.ListObjectsOptions{
					CatalogIdentifier: &catalogID,
					Offset:            &offset,
					Limit:             &limit,
				}

				searchResult, response, err := catalogManagementServiceAuthorized.ListObjects(listObjectsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(searchResult).NotTo(BeNil())

				if len(searchResult.Resources) > 0 {
					amountOfObjects += len(searchResult.Resources)
					offset += 50

					if contains == false {
						for _, obj := range searchResult.Resources {
							if *obj.ID == objectID {
								contains = true
								break
							}
						}
					}
				} else {
					fetch = false
				}
			}
			Expect(contains).To(BeTrue())
			fmt.Printf("Amount of objects: %d", amountOfObjects)
		})
	})

	Describe(`Replace Object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			replaceObjectOptions := &catalogmanagementv1.ReplaceObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				ID:                &objectID,
				Name:              core.StringPtr("updated-object-name-by-go-sdk"),
				ParentID:          &regionUSSouth,
				Kind:              core.StringPtr(kindVPE),
				CatalogID:         &catalogID,
				Data:              make(map[string]interface{}),
			}

			_, response, err := catalogManagementServiceNotAuthorized.ReplaceObject(replaceObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			replaceObjectOptions := &catalogmanagementv1.ReplaceObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
				ID:                &invalidObjectID,
				Name:              core.StringPtr("updated-object-name-by-go-sdk"),
				ParentID:          &regionUSSouth,
				Kind:              core.StringPtr(kindVPE),
				CatalogID:         &catalogID,
				Data:              make(map[string]interface{}),
			}

			_, response, err := catalogManagementServiceAuthorized.ReplaceObject(replaceObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			replaceObjectOptions := &catalogmanagementv1.ReplaceObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				ID:                &objectID,
				Name:              core.StringPtr("updated object name by go sdk"),
				ParentID:          &regionUSSouth,
				Kind:              core.StringPtr(kindVPE),
				CatalogID:         &catalogID,
				Data:              make(map[string]interface{}),
			}

			_, response, err := catalogManagementServiceAuthorized.ReplaceObject(replaceObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Replaces object`, func() {

			Skip("Cannot change the name of the object.")

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			replaceObjectOptions := &catalogmanagementv1.ReplaceObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				ID:                &objectID,
				Name:              core.StringPtr("updated-object-name-by-go-sdk"),
				ParentID:          &regionUSSouth,
				Kind:              core.StringPtr(kindVPE),
				CatalogID:         &catalogID,
				Data:              make(map[string]interface{}),
			}

			replaceObjectResult, response, err := catalogManagementServiceAuthorized.ReplaceObject(replaceObjectOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(replaceObjectResult).NotTo(BeNil())
		})
	})

	Describe(`Get Object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			getObjectOptions := &catalogmanagementv1.GetObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetObject(getObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			getObjectOptions := &catalogmanagementv1.GetObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetObject(getObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns the object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			getObjectOptions := &catalogmanagementv1.GetObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			requestedObject, response, err := catalogManagementServiceAuthorized.GetObject(getObjectOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(requestedObject).NotTo(BeNil())
		})
	})

	Describe(`Get Object Audit`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			getObjectAuditOptions := &catalogmanagementv1.GetObjectAuditOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetObjectAudit(getObjectAuditOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 200 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			getObjectAuditOptions := &catalogmanagementv1.GetObjectAuditOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetObjectAudit(getObjectAuditOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`Returns the object's audit log`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			getObjectAuditOptions := &catalogmanagementv1.GetObjectAuditOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			auditLog, response, err := catalogManagementServiceAuthorized.GetObjectAudit(getObjectAuditOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditLog).NotTo(BeNil())
		})
	})

	Describe(`Account Publish Object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			accountPublishObjectOptions := &catalogmanagementv1.AccountPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			response, err := catalogManagementServiceNotAuthorized.AccountPublishObject(accountPublishObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			accountPublishObjectOptions := &catalogmanagementv1.AccountPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
			}

			response, err := catalogManagementServiceAuthorized.AccountPublishObject(accountPublishObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Publishes object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			accountPublishObjectOptions := &catalogmanagementv1.AccountPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			response, err := catalogManagementServiceAuthorized.AccountPublishObject(accountPublishObjectOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`Shared Publish Object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			sharedPublishObjectOptions := &catalogmanagementv1.SharedPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			response, err := catalogManagementServiceNotAuthorized.SharedPublishObject(sharedPublishObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			sharedPublishObjectOptions := &catalogmanagementv1.SharedPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
			}

			response, err := catalogManagementServiceAuthorized.SharedPublishObject(sharedPublishObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Publishes object`, func() {
			Skip("Invalid catalog object.")

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			sharedPublishObjectOptions := &catalogmanagementv1.SharedPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			response, err := catalogManagementServiceAuthorized.SharedPublishObject(sharedPublishObjectOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`IBM Publish Object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			ibmPublishObjectOptions := &catalogmanagementv1.IBMPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			response, err := catalogManagementServiceNotAuthorized.IBMPublishObject(ibmPublishObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			ibmPublishObjectOptions := &catalogmanagementv1.IBMPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
			}

			response, err := catalogManagementServiceAuthorized.IBMPublishObject(ibmPublishObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Publishes object`, func() {
			Skip("This user is not granted.")

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			ibmPublishObjectOptions := &catalogmanagementv1.IBMPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			response, err := catalogManagementServiceAuthorized.IBMPublishObject(ibmPublishObjectOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`Public Publish Object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			publicPublishObjectOptions := &catalogmanagementv1.PublicPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			response, err := catalogManagementServiceNotAuthorized.PublicPublishObject(publicPublishObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			publicPublishObjectOptions := &catalogmanagementv1.PublicPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
			}

			response, err := catalogManagementServiceAuthorized.PublicPublishObject(publicPublishObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Publishes object`, func() {
			Skip("This user is not granted.")

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			publicPublishObjectOptions := &catalogmanagementv1.PublicPublishObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			response, err := catalogManagementServiceAuthorized.PublicPublishObject(publicPublishObjectOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`Create Object Access`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			createObjectAccessOptions := &catalogmanagementv1.CreateObjectAccessOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				AccountIdentifier: &accountID,
			}

			response, err := catalogManagementServiceNotAuthorized.CreateObjectAccess(createObjectAccessOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			createObjectAccessOptions := &catalogmanagementv1.CreateObjectAccessOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
				AccountIdentifier: &accountID,
			}

			response, err := catalogManagementServiceAuthorized.CreateObjectAccess(createObjectAccessOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Creates object access`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			createObjectAccessOptions := &catalogmanagementv1.CreateObjectAccessOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				AccountIdentifier: &accountID,
			}

			response, err := catalogManagementServiceAuthorized.CreateObjectAccess(createObjectAccessOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
		})
	})

	Describe(`Get Object Access List`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			getObjectAccessListOptions := &catalogmanagementv1.GetObjectAccessListOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetObjectAccessList(getObjectAccessListOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 200 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			getObjectAccessListOptions := &catalogmanagementv1.GetObjectAccessListOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetObjectAccessList(getObjectAccessListOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`Returns object's access list`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			getObjectAccessListOptions := &catalogmanagementv1.GetObjectAccessListOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			objectAccessList, response, err := catalogManagementServiceAuthorized.GetObjectAccessList(getObjectAccessListOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectAccessList).NotTo(BeNil())
		})
	})

	Describe(`Get Object Access`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			getObjectAccessOptions := &catalogmanagementv1.GetObjectAccessOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				AccountIdentifier: &accountID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetObjectAccess(getObjectAccessOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 200 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			getObjectAccessOptions := &catalogmanagementv1.GetObjectAccessOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
				AccountIdentifier: &accountID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetObjectAccess(getObjectAccessOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns object's access`, func() {

			Skip("Strange not found error see comments.")
			// Error: Error loading version with id: 6e263640-4805-471d-a30c-d7667325581c.
			// e59ad442-d113-49e4-bcd4-5431990135fd: Error[404 Not Found]

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			getObjectAccessOptions := &catalogmanagementv1.GetObjectAccessOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				AccountIdentifier: &accountID,
			}

			objectAccessList, response, err := catalogManagementServiceAuthorized.GetObjectAccess(getObjectAccessOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectAccessList).NotTo(BeNil())
		})
	})

	Describe(`Add Object Access List`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			addObjectAccessListOptions := &catalogmanagementv1.AddObjectAccessListOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				Accounts:          []string{accountID},
			}

			_, response, err := catalogManagementServiceNotAuthorized.AddObjectAccessList(addObjectAccessListOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			addObjectAccessListOptions := &catalogmanagementv1.AddObjectAccessListOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
				Accounts:          []string{accountID},
			}

			_, response, err := catalogManagementServiceAuthorized.AddObjectAccessList(addObjectAccessListOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Adds object access list`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			addObjectAccessListOptions := &catalogmanagementv1.AddObjectAccessListOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				Accounts:          []string{accountID},
			}

			accessListResponse, response, err := catalogManagementServiceAuthorized.AddObjectAccessList(addObjectAccessListOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(accessListResponse).NotTo(BeNil())
		})
	})

	Describe(`Create Offering Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 404 when no such catalog`, func() {

			Skip("None of the known kinds work")

			Expect(offeringID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidCatalogID := "invalid-" + catalogID
			createOfferingInstanceOptions := &catalogmanagementv1.CreateOfferingInstanceOptions{
				XAuthRefreshToken:    &refreshTokenAuthorized,
				ID:                   &offeringID,
				CatalogID:            &invalidCatalogID,
				OfferingID:           &offeringID,
				KindFormat:           core.StringPtr(kindVPE),
				Version:              core.StringPtr("0.0.2"),
				ClusterID:            &clusterID,
				ClusterRegion:        &regionUSSouth,
				ClusterAllNamespaces: core.BoolPtr(true),
			}

			_, response, err := catalogManagementServiceAuthorized.CreateOfferingInstance(createOfferingInstanceOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {

			Skip("None of the known kinds work")

			Expect(offeringID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			createOfferingInstanceOptions := &catalogmanagementv1.CreateOfferingInstanceOptions{
				XAuthRefreshToken:    &refreshTokenNotAuthorized,
				ID:                   &offeringID,
				CatalogID:            &catalogID,
				OfferingID:           &offeringID,
				KindFormat:           core.StringPtr(kindVPE),
				Version:              core.StringPtr("0.0.2"),
				ClusterID:            &clusterID,
				ClusterRegion:        &regionUSSouth,
				ClusterAllNamespaces: core.BoolPtr(true),
			}

			_, response, err := catalogManagementServiceNotAuthorized.CreateOfferingInstance(createOfferingInstanceOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			Expect(offeringID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			createOfferingInstanceOptions := &catalogmanagementv1.CreateOfferingInstanceOptions{
				XAuthRefreshToken:    &refreshTokenAuthorized,
				ID:                   &offeringID,
				CatalogID:            &catalogID,
				OfferingID:           &offeringID,
				KindFormat:           core.StringPtr("bogus kind"),
				Version:              core.StringPtr("0.0.2"),
				ClusterID:            &clusterID,
				ClusterRegion:        &regionUSSouth,
				ClusterAllNamespaces: core.BoolPtr(true),
			}

			_, response, err := catalogManagementServiceAuthorized.CreateOfferingInstance(createOfferingInstanceOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Creates Offering Instance`, func() {

			Skip("None of the known kinds work")

			Expect(offeringID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			createOfferingInstanceOptions := &catalogmanagementv1.CreateOfferingInstanceOptions{
				XAuthRefreshToken:    &refreshTokenAuthorized,
				ID:                   &offeringID,
				CatalogID:            &catalogID,
				OfferingID:           &offeringID,
				KindFormat:           core.StringPtr(kindVPE),
				Version:              core.StringPtr("0.0.2"),
				ClusterID:            &clusterID,
				ClusterRegion:        &regionUSSouth,
				ClusterAllNamespaces: core.BoolPtr(true),
			}

			createdOfferingInstance, response, err := catalogManagementServiceAuthorized.CreateOfferingInstance(createOfferingInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createdOfferingInstance).NotTo(BeNil())
			Expect(createdOfferingInstance.ID).NotTo(BeNil())
			offeringInstanceID = *createdOfferingInstance.ID
		})
	})

	Describe(`Get Offering Instance - Get Offering Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {
			Skip("No offering instance id.")

			Expect(offeringInstanceID).ToNot(BeNil())

			getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{
				InstanceIdentifier: &offeringInstanceID,
			}

			_, response, err := catalogManagementServiceNotAuthorized.GetOfferingInstance(getOfferingInstanceOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such offering instance`, func() {

			Skip("No offering instance id.")

			Expect(offeringInstanceID).ToNot(BeNil())

			invalidOfferingInstanceID := "invalid-" + offeringInstanceID
			getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{
				InstanceIdentifier: &invalidOfferingInstanceID,
			}

			_, response, err := catalogManagementServiceAuthorized.GetOfferingInstance(getOfferingInstanceOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns the offering instance`, func() {
			Skip("No offering instance id.")

			Expect(offeringInstanceID).ToNot(BeNil())

			getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{
				InstanceIdentifier: &offeringInstanceID,
			}

			offeringInstance, response, err := catalogManagementServiceAuthorized.GetOfferingInstance(getOfferingInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringInstance).NotTo(BeNil())
		})
	})

	Describe(`Put Offering Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Skip("No offering instance id.")

			Expect(offeringInstanceID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			putOfferingInstanceOptions := &catalogmanagementv1.PutOfferingInstanceOptions{
				InstanceIdentifier:   &offeringInstanceID,
				XAuthRefreshToken:    &refreshTokenNotAuthorized,
				ID:                   &offeringInstanceID,
				CatalogID:            &catalogID,
				OfferingID:           &offeringID,
				KindFormat:           core.StringPtr(kindVPE),
				Version:              core.StringPtr("0.0.3"),
				ClusterID:            &clusterID,
				ClusterRegion:        &regionUSSouth,
				ClusterAllNamespaces: core.BoolPtr(true),
			}

			_, response, err := catalogManagementServiceNotAuthorized.PutOfferingInstance(putOfferingInstanceOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such catalog`, func() {

			Skip("No offering instance id.")

			Expect(offeringInstanceID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			invalidCatalogID := "invalid-" + catalogID
			putOfferingInstanceOptions := &catalogmanagementv1.PutOfferingInstanceOptions{
				InstanceIdentifier:   &offeringInstanceID,
				XAuthRefreshToken:    &refreshTokenAuthorized,
				ID:                   &offeringInstanceID,
				CatalogID:            &invalidCatalogID,
				OfferingID:           &offeringID,
				KindFormat:           core.StringPtr(kindVPE),
				Version:              core.StringPtr("0.0.3"),
				ClusterID:            &clusterID,
				ClusterRegion:        &regionUSSouth,
				ClusterAllNamespaces: core.BoolPtr(true),
			}

			_, response, err := catalogManagementServiceAuthorized.PutOfferingInstance(putOfferingInstanceOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			Skip("No offering instance id.")

			Expect(offeringInstanceID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			putOfferingInstanceOptions := &catalogmanagementv1.PutOfferingInstanceOptions{
				InstanceIdentifier:   &offeringInstanceID,
				XAuthRefreshToken:    &refreshTokenAuthorized,
				ID:                   &offeringInstanceID,
				CatalogID:            &catalogID,
				OfferingID:           &offeringID,
				KindFormat:           core.StringPtr("bogus kind"),
				Version:              core.StringPtr("0.0.3"),
				ClusterID:            &clusterID,
				ClusterRegion:        &regionUSSouth,
				ClusterAllNamespaces: core.BoolPtr(true),
			}

			_, response, err := catalogManagementServiceAuthorized.PutOfferingInstance(putOfferingInstanceOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Updates the offering instance`, func() {

			Skip("No offering instance id.")

			Expect(offeringInstanceID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())
			Expect(offeringID).ToNot(BeNil())

			putOfferingInstanceOptions := &catalogmanagementv1.PutOfferingInstanceOptions{
				InstanceIdentifier:   &offeringInstanceID,
				XAuthRefreshToken:    &refreshTokenAuthorized,
				ID:                   &offeringInstanceID,
				CatalogID:            &catalogID,
				OfferingID:           &offeringID,
				KindFormat:           core.StringPtr(kindVPE),
				Version:              core.StringPtr("0.0.3"),
				ClusterID:            &clusterID,
				ClusterRegion:        &regionUSSouth,
				ClusterAllNamespaces: core.BoolPtr(true),
			}

			offeringInstance, response, err := catalogManagementServiceAuthorized.PutOfferingInstance(putOfferingInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringInstance).NotTo(BeNil())
		})
	})

	Describe(`Delete Version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 400 when backend input validation fails`, func() {

			deleteVersionOptions := &catalogmanagementv1.DeleteVersionOptions{
				VersionLocID: &bogusVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteVersion(deleteVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 404 when no such version`, func() {

			Expect(versionLocatorID).ToNot(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			deleteVersionOptions := &catalogmanagementv1.DeleteVersionOptions{
				VersionLocID: &invalidVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteVersion(deleteVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 403 when user is not authorized`, func() {

			Expect(versionLocatorID).ToNot(BeNil())

			deleteVersionOptions := &catalogmanagementv1.DeleteVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceNotAuthorized.DeleteVersion(deleteVersionOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Deletes the version`, func() {

			Expect(versionLocatorID).ToNot(BeNil())

			deleteVersionOptions := &catalogmanagementv1.DeleteVersionOptions{
				VersionLocID: &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteVersion(deleteVersionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`Delete Operators`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(versionLocatorID).ToNot(BeNil())

			deleteOperatorsOptions := &catalogmanagementv1.DeleteOperatorsOptions{
				XAuthRefreshToken: &refreshTokenNotAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			response, err := catalogManagementServiceNotAuthorized.DeleteOperators(deleteOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such version`, func() {

			Expect(versionLocatorID).ToNot(BeNil())

			invalidVersionLocatorID := "invalid-" + versionLocatorID
			deleteOperatorsOptions := &catalogmanagementv1.DeleteOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &invalidVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteOperators(deleteOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Returns 400 when backend input validation fails`, func() {

			deleteOperatorsOptions := &catalogmanagementv1.DeleteOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &bogusVersionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteOperators(deleteOperatorsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Deletes the operator`, func() {
			Skip("Strange not found error, see comments.")
			// Error: Error loading version with id: fdeefb18-57aa-4390-a9e0-b66b551db803.
			// 2c187aa6-5009-4a2f-8f57-86533d2d3a18: Error[404 Not Found] -
			// Version not found: Catalog[fdeefb18-57aa-4390-a9e0-b66b551db803]:Version[2c187aa6-5009-4a2f-8f57-86533d2d3a18]

			Expect(versionLocatorID).ToNot(BeNil())

			deleteOperatorsOptions := &catalogmanagementv1.DeleteOperatorsOptions{
				XAuthRefreshToken: &refreshTokenAuthorized,
				ClusterID:         &clusterID,
				Region:            &regionUSSouth,
				VersionLocatorID:  &versionLocatorID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteOperators(deleteOperatorsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`Delete Offering Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {
			Skip("No Offering instance id.")

			Expect(offeringInstanceID).ToNot(BeNil())

			deleteOfferingInstanceOptions := &catalogmanagementv1.DeleteOfferingInstanceOptions{
				InstanceIdentifier: &offeringInstanceID,
				XAuthRefreshToken:  &refreshTokenNotAuthorized,
			}

			response, err := catalogManagementServiceNotAuthorized.DeleteOfferingInstance(deleteOfferingInstanceOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such offering instance`, func() {
			Skip("No Offering instance id.")

			Expect(offeringInstanceID).ToNot(BeNil())

			invalidOfferingInstanceID := "invalid-" + offeringInstanceID
			deleteOfferingInstanceOptions := &catalogmanagementv1.DeleteOfferingInstanceOptions{
				InstanceIdentifier: &invalidOfferingInstanceID,
				XAuthRefreshToken:  &refreshTokenAuthorized,
			}

			response, err := catalogManagementServiceAuthorized.DeleteOfferingInstance(deleteOfferingInstanceOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Deletes the offering instance`, func() {
			Skip("No offering instance.")

			Expect(offeringInstanceID).ToNot(BeNil())

			deleteOfferingInstanceOptions := &catalogmanagementv1.DeleteOfferingInstanceOptions{
				InstanceIdentifier: &offeringInstanceID,
				XAuthRefreshToken:  &refreshTokenAuthorized,
			}

			response, err := catalogManagementServiceAuthorized.DeleteOfferingInstance(deleteOfferingInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`Delete Object Access List`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			deleteObjectAccessListOptions := &catalogmanagementv1.DeleteObjectAccessListOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				Accounts:          []string{accountID},
			}

			_, response, err := catalogManagementServiceNotAuthorized.DeleteObjectAccessList(deleteObjectAccessListOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such catalog`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidCatalogID := "invalid-" + catalogID
			deleteObjectAccessListOptions := &catalogmanagementv1.DeleteObjectAccessListOptions{
				CatalogIdentifier: &invalidCatalogID,
				ObjectIdentifier:  &objectID,
				Accounts:          []string{accountID},
			}

			_, response, err := catalogManagementServiceAuthorized.DeleteObjectAccessList(deleteObjectAccessListOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Deletes object access list`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			deleteObjectAccessListOptions := &catalogmanagementv1.DeleteObjectAccessListOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				Accounts:          []string{accountID},
			}

			result, response, err := catalogManagementServiceAuthorized.DeleteObjectAccessList(deleteObjectAccessListOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).NotTo(BeNil())
		})
	})

	Describe(`Delete Object Access`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			deleteObjectAccessOptions := &catalogmanagementv1.DeleteObjectAccessOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				AccountIdentifier: &accountID,
			}

			response, err := catalogManagementServiceNotAuthorized.DeleteObjectAccess(deleteObjectAccessOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when no such catalog`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidCatalogID := "invalid-" + catalogID
			deleteObjectAccessOptions := &catalogmanagementv1.DeleteObjectAccessOptions{
				CatalogIdentifier: &invalidCatalogID,
				ObjectIdentifier:  &objectID,
				AccountIdentifier: &accountID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteObjectAccess(deleteObjectAccessOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`Deletes object access`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			deleteObjectAccessOptions := &catalogmanagementv1.DeleteObjectAccessOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
				AccountIdentifier: &accountID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteObjectAccess(deleteObjectAccessOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`Delete Object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Returns 403 when user is not authorized`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			deleteObjectOptions := &catalogmanagementv1.DeleteObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			response, err := catalogManagementServiceNotAuthorized.DeleteObject(deleteObjectOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 200 when no such object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidObjectID := "invalid-" + objectID
			deleteObjectOptions := &catalogmanagementv1.DeleteObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &invalidObjectID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteObject(deleteObjectOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`Deletes object`, func() {

			Expect(objectID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			deleteObjectOptions := &catalogmanagementv1.DeleteObjectOptions{
				CatalogIdentifier: &catalogID,
				ObjectIdentifier:  &objectID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteObject(deleteObjectOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`Delete Offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 200 when no such offering`, func() {

			Expect(offeringID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			invalidOfferingID := "invalid-" + offeringID
			deleteOfferingOptions := &catalogmanagementv1.DeleteOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &invalidOfferingID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteOffering(deleteOfferingOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`Returns 403 when user is not authorized`, func() {

			Expect(offeringID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			deleteOfferingOptions := &catalogmanagementv1.DeleteOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
			}

			response, err := catalogManagementServiceNotAuthorized.DeleteOffering(deleteOfferingOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Deletes offering`, func() {

			Expect(offeringID).ToNot(BeNil())
			Expect(catalogID).ToNot(BeNil())

			deleteOfferingOptions := &catalogmanagementv1.DeleteOfferingOptions{
				CatalogIdentifier: &catalogID,
				OfferingID:        &offeringID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteOffering(deleteOfferingOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`Delete Catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 404 when no such catalog`, func() {
			Expect(catalogID).NotTo(BeNil())

			invalidCatalogId := "invalid-" + catalogID
			deleteCatalogOptions := &catalogmanagementv1.DeleteCatalogOptions{
				CatalogIdentifier: &invalidCatalogId,
			}

			response, err := catalogManagementServiceAuthorized.DeleteCatalog(deleteCatalogOptions)

			Expect(err).To(BeNil())
			Expect(response).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`Returns 403 when user is not authorized`, func() {
			Expect(catalogID).NotTo(BeNil())

			deleteCatalogOptions := &catalogmanagementv1.DeleteCatalogOptions{
				CatalogIdentifier: &catalogID,
			}

			response, err := catalogManagementServiceNotAuthorized.DeleteCatalog(deleteCatalogOptions)

			Expect(err).NotTo(BeNil())
			Expect(response).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Deletes catalog`, func() {
			Expect(catalogID).NotTo(BeNil())

			deleteCatalogOptions := &catalogmanagementv1.DeleteCatalogOptions{
				CatalogIdentifier: &catalogID,
			}

			response, err := catalogManagementServiceAuthorized.DeleteCatalog(deleteCatalogOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	AfterSuite(func() {

		deleteObjectOptions := &catalogmanagementv1.DeleteObjectOptions{
			CatalogIdentifier: &catalogID,
			ObjectIdentifier:  &objectID,
		}

		_, err = catalogManagementServiceAuthorized.DeleteObject(deleteObjectOptions)
		if err != nil {
			fmt.Println("Object is already deleted.")
		}

		deleteOfferingOptions := &catalogmanagementv1.DeleteOfferingOptions{
			CatalogIdentifier: &catalogID,
			OfferingID:        &offeringID,
		}

		_, err = catalogManagementServiceAuthorized.DeleteOffering(deleteOfferingOptions)
		if err != nil {
			fmt.Println("Offering is already deleted.")
		}

		deleteCatalogOptions := &catalogmanagementv1.DeleteCatalogOptions{
			CatalogIdentifier: &catalogID,
		}

		_, err = catalogManagementServiceAuthorized.DeleteCatalog(deleteCatalogOptions)
		if err != nil {
			fmt.Println("Catalog is already deleted.")
		}
	})
})
