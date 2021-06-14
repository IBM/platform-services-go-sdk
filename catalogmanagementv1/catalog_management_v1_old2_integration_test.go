// +build integration2

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
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the catalogmanagementv1 package.
 *
 * s:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = FDescribe(`CatalogManagementV1 Integration Tests`, func() {

	const (
		externalConfigFile   = "../catalog_management.env"
		expectedShortDesc    = "test"
		expectedURL          = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s"
		expectedOfferingsURL = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings"
		fakeName             = "bogus"
		fakeVersionLocator   = "bogus.bogus"
		expectedOfferingName = "test-offering"
		expectedOfferingURL  = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings/%s"
		fakeRevision         = "rev_2.0"
	)

	var (
		err                      error
		catalogManagementService *catalogmanagementv1.CatalogManagementV1
		serviceURL               string
		config                   map[string]string

		expectedLabel string = fmt.Sprintf("integration-test-%d", time.Now().Unix())
		accountID     string
		gitToken      string
		clusterID     string

		catalogID  string
		offeringID string
		objectID   string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	FDescribe(`External configuration`, func() {
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
			if accountID == "" {
				Skip("Unable to load account ID configuration property, skipping tests")
			}

			gitToken = config["GIT_TOKEN"]
			if gitToken == "" {
				Skip("Unable to load Git token configuration property, skipping tests")
			}

			clusterID = config["CLUSTER_ID"]
			if clusterID == "" {
				Skip("Unable to load cluster ID configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	FDescribe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			catalogManagementServiceOptions := &catalogmanagementv1.CatalogManagementV1Options{}

			catalogManagementService, err = catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(catalogManagementServiceOptions)

			Expect(err).To(BeNil())
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(catalogManagementService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			catalogManagementService.EnableRetries(4, 30*time.Second)
		})
	})

	FDescribe(`Run integration tests`, func() {
		JustBeforeEach(func() {
			shouldSkipTest()

			return

			listResult, _, _ := catalogManagementService.ListCatalogs(catalogManagementService.NewListCatalogsOptions())
			if listResult != nil && listResult.Resources != nil {
				for _, resource := range listResult.Resources {
					if *resource.Label == expectedLabel {
						catalogManagementService.DeleteCatalog(catalogManagementService.NewDeleteCatalogOptions(*resource.ID))
					}
				}
			}
		})

		JustAfterEach(func() {
			shouldSkipTest()

			return

			listResult, _, _ := catalogManagementService.ListCatalogs(catalogManagementService.NewListCatalogsOptions())
			if listResult != nil && listResult.Resources != nil {
				for _, resource := range listResult.Resources {
					if *resource.Label == expectedLabel {
						catalogManagementService.DeleteCatalog(catalogManagementService.NewDeleteCatalogOptions(*resource.ID))
					}
				}
			}
		})
		// *************** ACCOUNTS
		FDescribe(`GetCatalogAccount - Get catalog account settings`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetCatalogAccount(getCatalogAccountOptions *GetCatalogAccountOptions)`, func() {

				getCatalogAccountOptions := &catalogmanagementv1.GetCatalogAccountOptions{}

				account, response, err := catalogManagementService.GetCatalogAccount(getCatalogAccountOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(account).ToNot(BeNil())
				Expect(*account.ID).To(Equal(accountID))
				Expect(len(account.AccountFilters.CategoryFilters)).To(BeZero())
				// NOTE1: Are these necessary?
				// Expect(account.AccountFilters.IDFilters.Include).ToNot(BeNil())
				// Expect(account.AccountFilters.IDFilters.Exclude).ToNot(BeNil())
			})
			It("GetCatalogAccount - Fail to get a catalog that does not exist", func() {
				getOptions := catalogManagementService.NewGetCatalogOptions(fakeName)
				_, getResponse, err := catalogManagementService.GetCatalog(getOptions)

				Expect(err).ToNot(BeNil())
				Expect(getResponse.StatusCode).To(Equal(404))
			})
		})

		// NOTE2: Can we run this? Doesn't this require another account to make the update? I guess it cannot update itself.
		Describe(`UpdateCatalogAccount - Update account settings`, func() {
			It(`UpdateCatalogAccount(updateCatalogAccountOptions *UpdateCatalogAccountOptions)`, func() {

				filterTermsModel := &catalogmanagementv1.FilterTerms{
					FilterTerms: []string{"filterTerm"},
				}

				categoryFilterModel := &catalogmanagementv1.CategoryFilter{
					Include: core.BoolPtr(true),
					Filter:  filterTermsModel,
				}

				idFilterModel := &catalogmanagementv1.IDFilter{
					Include: filterTermsModel,
					Exclude: filterTermsModel,
				}

				filtersModel := &catalogmanagementv1.Filters{
					IncludeAll:      core.BoolPtr(true),
					CategoryFilters: make(map[string]catalogmanagementv1.CategoryFilter),
					IDFilters:       idFilterModel,
				}
				filtersModel.CategoryFilters["category"] = *categoryFilterModel

				updateCatalogAccountOptions := &catalogmanagementv1.UpdateCatalogAccountOptions{
					ID:                  &accountID,
					HideIBMCloudCatalog: core.BoolPtr(true),
					AccountFilters:      filtersModel,
				}

				response, err := catalogManagementService.UpdateCatalogAccount(updateCatalogAccountOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		FDescribe(`GetCatalogAccountAudit - Get catalog account audit log`, func() {
			It(`GetCatalogAccountAudit(getCatalogAccountAuditOptions *GetCatalogAccountAuditOptions)`, func() {

				getCatalogAccountAuditOptions := &catalogmanagementv1.GetCatalogAccountAuditOptions{}

				auditLog, response, err := catalogManagementService.GetCatalogAccountAudit(getCatalogAccountAuditOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLog).ToNot(BeNil())

			})
		})

		FDescribe(`GetCatalogAccountFilters - Get catalog account filters`, func() {
			It(`GetCatalogAccountFilters(getCatalogAccountFiltersOptions *GetCatalogAccountFiltersOptions)`, func() {

				getCatalogAccountFiltersOptions := &catalogmanagementv1.GetCatalogAccountFiltersOptions{
					// NOTE3: Catalog: core.StringPtr("testString"),
				}

				accumulatedFilters, response, err := catalogManagementService.GetCatalogAccountFilters(getCatalogAccountFiltersOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(accumulatedFilters).ToNot(BeNil())
				Expect(*accumulatedFilters.AccountFilters[0].IncludeAll).To(BeTrue())
				Expect(len(accumulatedFilters.AccountFilters[0].CategoryFilters)).To(BeZero())
				// NOTE4: GOTO NOTE1
				// Expect(accumulatedFilters.AccountFilters[0].IDFilters.Include).ToNot(BeNil())
				// Expect(accumulatedFilters.AccountFilters[0].IDFilters.Exclude).ToNot(BeNil())
			})
		})

		// *************** CATALOGS
		FDescribe(`CreateCatalog - Create a catalog`, func() {

			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`CreateCatalog(createCatalogOptions *CreateCatalogOptions)`, func() {

				// NOTE5: What about these?
				// featureModel := &catalogmanagementv1.Feature{
				// 	Title:       core.StringPtr("title"),
				// 	Description: core.StringPtr("description"),
				// }

				// filterTermsModel := &catalogmanagementv1.FilterTerms{
				// 	FilterTerms: []string{"filterTerm"},
				// }

				// categoryFilterModel := &catalogmanagementv1.CategoryFilter{
				// 	Include: core.BoolPtr(true),
				// 	Filter:  filterTermsModel,
				// }

				// idFilterModel := &catalogmanagementv1.IDFilter{
				// 	Include: filterTermsModel,
				// 	Exclude: filterTermsModel,
				// }

				// filtersModel := &catalogmanagementv1.Filters{
				// 	IncludeAll:      core.BoolPtr(true),
				// 	CategoryFilters: make(map[string]catalogmanagementv1.CategoryFilter),
				// 	IDFilters:       idFilterModel,
				// }
				// filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// syndicationClusterModel := &catalogmanagementv1.SyndicationCluster{
				// 	Region:            core.StringPtr("US"),
				// 	ID:                core.StringPtr("testString"),
				// 	Name:              core.StringPtr("testString"),
				// 	ResourceGroupName: core.StringPtr("testString"),
				// 	Type:              core.StringPtr("testString"),
				// 	Namespaces:        []string{"testString"},
				// 	AllNamespaces:     core.BoolPtr(true),
				// }

				// syndicationHistoryModel := &catalogmanagementv1.SyndicationHistory{
				// 	Namespaces: []string{"testString"},
				// 	Clusters:   []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel},
				// 	LastRun:    CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// }

				// syndicationAuthorizationModel := &catalogmanagementv1.SyndicationAuthorization{
				// 	Token:   core.StringPtr("testString"),
				// 	LastRun: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// }

				// syndicationResourceModel := &catalogmanagementv1.SyndicationResource{
				// 	RemoveRelatedComponents: core.BoolPtr(true),
				// 	Clusters:                []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel},
				// 	History:                 syndicationHistoryModel,
				// 	Authorization:           syndicationAuthorizationModel,
				// }

				createCatalogOptions := &catalogmanagementv1.CreateCatalogOptions{
					Label:            core.StringPtr(expectedLabel),
					ShortDescription: core.StringPtr(expectedShortDesc),
					// CatalogIconURL:      core.StringPtr("testString"),
					// Tags:                []string{"testString"},
					// Features:            []catalogmanagementv1.Feature{*featureModel},
					// Disabled:            core.BoolPtr(true),
					// ResourceGroupID:     core.StringPtr("testString"),
					// OwningAccount:       core.StringPtr("testString"),
					// CatalogFilters:      filtersModel,
					// SyndicationSettings: syndicationResourceModel,
					// Kind:                core.StringPtr("testString"),
				}

				catalog, response, err := catalogManagementService.CreateCatalog(createCatalogOptions)

				// catalogManagementService.DeleteCatalog(catalogManagementService.NewDeleteCatalogOptions(*catalog.ID))

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(catalog).ToNot(BeNil())

				catalogID = *catalog.ID

				Expect(*catalog.Label).To(Equal(expectedLabel))
				Expect(*catalog.ShortDescription).To(Equal(expectedShortDesc))
				Expect(*catalog.URL).To(Equal(fmt.Sprintf(expectedURL, *catalog.ID)))
				Expect(*catalog.OfferingsURL).To(Equal(fmt.Sprintf(expectedOfferingsURL, *catalog.ID)))
				Expect(*catalog.OwningAccount).To(Equal(accountID))
				Expect(*catalog.CatalogFilters.IncludeAll).To(BeFalse())
				Expect(len(catalog.CatalogFilters.CategoryFilters)).To(BeZero())
				Expect(catalog.CatalogFilters.IDFilters.Include).To(BeNil())
				Expect(catalog.CatalogFilters.IDFilters.Exclude).To(BeNil())

			})
		})

		FDescribe(`ListCatalogs - Get list of catalogs`, func() {
			It(`ListCatalogs(listCatalogsOptions *ListCatalogsOptions)`, func() {
				const (
					expectedTotalCount    = 1
					expectedResourceCount = 1
				)

				catalogCount := 0
				catalogIndex := -1

				// createOptions := catalogManagementService.NewCreateCatalogOptions()
				// createOptions.SetLabel(expectedLabel)
				// createOptions.SetShortDescription(expectedShortDesc)
				// createResult, _, _ := catalogManagementService.CreateCatalog(createOptions)

				listCatalogsOptions := &catalogmanagementv1.ListCatalogsOptions{}

				catalogSearchResult, response, err := catalogManagementService.ListCatalogs(listCatalogsOptions)

				for i, resource := range catalogSearchResult.Resources {
					if *resource.Label == expectedLabel {
						catalogCount++
						catalogIndex = i
					}
				}

				// catalogManagementService.DeleteCatalog(catalogManagementService.NewDeleteCatalogOptions(*createResult.ID))

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(catalogSearchResult).ToNot(BeNil())

				Expect(catalogCount).To(Equal(expectedTotalCount))
				Expect(*catalogSearchResult.Resources[catalogIndex].ID).To(Equal(catalogID))
				Expect(*catalogSearchResult.Resources[catalogIndex].Label).To(Equal(expectedLabel))
				Expect(*catalogSearchResult.Resources[catalogIndex].ShortDescription).To(Equal(expectedShortDesc))
				Expect(*catalogSearchResult.Resources[catalogIndex].URL).To(Equal(fmt.Sprintf(expectedURL, catalogID)))
				Expect(*catalogSearchResult.Resources[catalogIndex].OfferingsURL).To(Equal(fmt.Sprintf(expectedOfferingsURL, catalogID)))
				Expect(*catalogSearchResult.Resources[catalogIndex].OwningAccount).To(Equal(accountID))
				Expect(*catalogSearchResult.Resources[catalogIndex].CatalogFilters.IncludeAll).To(BeFalse())
				Expect(len(catalogSearchResult.Resources[catalogIndex].CatalogFilters.CategoryFilters)).To(BeZero())
				Expect(catalogSearchResult.Resources[catalogIndex].CatalogFilters.IDFilters.Include).To(BeNil())
				Expect(catalogSearchResult.Resources[catalogIndex].CatalogFilters.IDFilters.Exclude).To(BeNil())
			})
		})

		FDescribe(`GetCatalog - Get catalog`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {
				Expect(catalogID).ToNot(BeEmpty())

				getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{
					CatalogIdentifier: &catalogID,
				}

				catalog, response, err := catalogManagementService.GetCatalog(getCatalogOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(catalog).ToNot(BeNil())
				Expect(*catalog.ID).To(Equal(catalogID))
				Expect(*catalog.Label).To(Equal(expectedLabel))
				Expect(*catalog.ShortDescription).To(Equal(expectedShortDesc))
				Expect(*catalog.URL).To(Equal(fmt.Sprintf(expectedURL, catalogID)))
				Expect(*catalog.OfferingsURL).To(Equal(fmt.Sprintf(expectedOfferingsURL, catalogID)))
				Expect(*catalog.OwningAccount).To(Equal(accountID))
				Expect(*catalog.CatalogFilters.IncludeAll).To(BeFalse())
				Expect(len(catalog.CatalogFilters.CategoryFilters)).To(BeZero())
				Expect(catalog.CatalogFilters.IDFilters.Include).To(BeNil())
				Expect(catalog.CatalogFilters.IDFilters.Exclude).To(BeNil())
			})
			It("GetCatalog(getCatalogOptions *GetCatalogOptions) - Fail to get a catalog that does not exist", func() {
				id := fakeName
				getOptions := catalogManagementService.NewGetCatalogOptions(id)
				_, getResponse, err := catalogManagementService.GetCatalog(getOptions)

				Expect(err).ToNot(BeNil())
				Expect(getResponse.StatusCode).To(Equal(404))
			})
		})

		FDescribe(`ReplaceCatalog - Update catalog`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ReplaceCatalog(replaceCatalogOptions *ReplaceCatalogOptions)`, func() {

				// featureModel := &catalogmanagementv1.Feature{
				// 	Title:       core.StringPtr("testString"),
				// 	Description: core.StringPtr("testString"),
				// }

				// filterTermsModel := &catalogmanagementv1.FilterTerms{
				// 	FilterTerms: []string{"testString"},
				// }

				// categoryFilterModel := &catalogmanagementv1.CategoryFilter{
				// 	Include: core.BoolPtr(true),
				// 	Filter:  filterTermsModel,
				// }

				// idFilterModel := &catalogmanagementv1.IDFilter{
				// 	Include: filterTermsModel,
				// 	Exclude: filterTermsModel,
				// }

				// filtersModel := &catalogmanagementv1.Filters{
				// 	IncludeAll:      core.BoolPtr(true),
				// 	CategoryFilters: make(map[string]catalogmanagementv1.CategoryFilter),
				// 	IDFilters:       idFilterModel,
				// }
				// filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// syndicationClusterModel := &catalogmanagementv1.SyndicationCluster{
				// 	Region:            core.StringPtr("testString"),
				// 	ID:                core.StringPtr("testString"),
				// 	Name:              core.StringPtr("testString"),
				// 	ResourceGroupName: core.StringPtr("testString"),
				// 	Type:              core.StringPtr("testString"),
				// 	Namespaces:        []string{"testString"},
				// 	AllNamespaces:     core.BoolPtr(true),
				// }

				// syndicationHistoryModel := &catalogmanagementv1.SyndicationHistory{
				// 	Namespaces: []string{"testString"},
				// 	Clusters:   []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel},
				// 	LastRun:    CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// }

				// syndicationAuthorizationModel := &catalogmanagementv1.SyndicationAuthorization{
				// 	Token:   core.StringPtr("testString"),
				// 	LastRun: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// }

				// syndicationResourceModel := &catalogmanagementv1.SyndicationResource{
				// 	RemoveRelatedComponents: core.BoolPtr(true),
				// 	Clusters:                []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel},
				// 	History:                 syndicationHistoryModel,
				// 	Authorization:           syndicationAuthorizationModel,
				// }

				replaceCatalogOptions := &catalogmanagementv1.ReplaceCatalogOptions{
					CatalogIdentifier: core.StringPtr(catalogID),
					ID:                core.StringPtr(catalogID),
					// Rev:                 core.StringPtr("testString"),
					// Label:               core.StringPtr("testString"),
					// ShortDescription:    core.StringPtr("testString"),
					// CatalogIconURL:      core.StringPtr("testString"),
					Tags: []string{"tag"},
					// Features:            []catalogmanagementv1.Feature{*featureModel},
					// Disabled:            core.BoolPtr(true),
					// ResourceGroupID:     core.StringPtr("testString"),
					// OwningAccount:       core.StringPtr("testString"),
					// CatalogFilters:      filtersModel,
					// SyndicationSettings: syndicationResourceModel,
					// Kind:                core.StringPtr("testString"),
				}

				catalog, response, err := catalogManagementService.ReplaceCatalog(replaceCatalogOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(catalog).ToNot(BeNil())
			})
		})

		FDescribe(`GetCatalogAudit - Get catalog audit log`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetCatalogAudit(getCatalogAuditOptions *GetCatalogAuditOptions)`, func() {

				getCatalogAuditOptions := &catalogmanagementv1.GetCatalogAuditOptions{
					CatalogIdentifier: &catalogID,
				}

				auditLog, response, err := catalogManagementService.GetCatalogAudit(getCatalogAuditOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLog).ToNot(BeNil())

			})
		})

		// *************** OFFERINGS
		FDescribe(`CreateOffering - Create offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`CreateOffering(createOfferingOptions *CreateOfferingOptions)`, func() {

				// ratingModel := &catalogmanagementv1.Rating{
				// 	OneStarCount:   core.Int64Ptr(int64(38)),
				// 	TwoStarCount:   core.Int64Ptr(int64(38)),
				// 	ThreeStarCount: core.Int64Ptr(int64(38)),
				// 	FourStarCount:  core.Int64Ptr(int64(38)),
				// }

				// featureModel := &catalogmanagementv1.Feature{
				// 	Title:       core.StringPtr("testString"),
				// 	Description: core.StringPtr("testString"),
				// }

				// configurationModel := &catalogmanagementv1.Configuration{
				// 	Key:             core.StringPtr("testString"),
				// 	Type:            core.StringPtr("testString"),
				// 	DefaultValue:    core.StringPtr("testString"),
				// 	ValueConstraint: core.StringPtr("testString"),
				// 	Description:     core.StringPtr("testString"),
				// 	Required:        core.BoolPtr(true),
				// 	Options:         []interface{}{"testString"},
				// 	Hidden:          core.BoolPtr(true),
				// }

				// validationModel := &catalogmanagementv1.Validation{
				// 	Validated:     CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Requested:     CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	State:         core.StringPtr("testString"),
				// 	LastOperation: core.StringPtr("testString"),
				// 	Target:        make(map[string]interface{}),
				// }

				// resourceModel := &catalogmanagementv1.Resource{
				// 	Type:  core.StringPtr("mem"),
				// 	Value: core.StringPtr("testString"),
				// }

				// scriptModel := &catalogmanagementv1.Script{
				// 	Instructions:     core.StringPtr("testString"),
				// 	Script:           core.StringPtr("testString"),
				// 	ScriptPermission: core.StringPtr("testString"),
				// 	DeleteScript:     core.StringPtr("testString"),
				// 	Scope:            core.StringPtr("testString"),
				// }

				// versionEntitlementModel := &catalogmanagementv1.VersionEntitlement{
				// 	ProviderName:  core.StringPtr("testString"),
				// 	ProviderID:    core.StringPtr("testString"),
				// 	ProductID:     core.StringPtr("testString"),
				// 	PartNumbers:   []string{"testString"},
				// 	ImageRepoName: core.StringPtr("testString"),
				// }

				// licenseModel := &catalogmanagementv1.License{
				// 	ID:          core.StringPtr("testString"),
				// 	Name:        core.StringPtr("testString"),
				// 	Type:        core.StringPtr("testString"),
				// 	URL:         core.StringPtr("testString"),
				// 	Description: core.StringPtr("testString"),
				// }

				// stateModel := &catalogmanagementv1.State{
				// 	Current:          core.StringPtr("testString"),
				// 	CurrentEntered:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Pending:          core.StringPtr("testString"),
				// 	PendingRequested: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Previous:         core.StringPtr("testString"),
				// }

				// versionModel := &catalogmanagementv1.Version{
				// 	ID:                  core.StringPtr("testString"),
				// 	Rev:                 core.StringPtr("testString"),
				// 	CRN:                 core.StringPtr("testString"),
				// 	Version:             core.StringPtr("testString"),
				// 	Sha:                 core.StringPtr("testString"),
				// 	Created:             CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Updated:             CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	OfferingID:          core.StringPtr("testString"),
				// 	CatalogID:           core.StringPtr("testString"),
				// 	KindID:              core.StringPtr("testString"),
				// 	Tags:                []string{"testString"},
				// 	RepoURL:             core.StringPtr("testString"),
				// 	SourceURL:           core.StringPtr("testString"),
				// 	TgzURL:              core.StringPtr("testString"),
				// 	Configuration:       []catalogmanagementv1.Configuration{*configurationModel},
				// 	Metadata:            make(map[string]interface{}),
				// 	Validation:          validationModel,
				// 	RequiredResources:   []catalogmanagementv1.Resource{*resourceModel},
				// 	SingleInstance:      core.BoolPtr(true),
				// 	Install:             scriptModel,
				// 	PreInstall:          []catalogmanagementv1.Script{*scriptModel},
				// 	Entitlement:         versionEntitlementModel,
				// 	Licenses:            []catalogmanagementv1.License{*licenseModel},
				// 	ImageManifestURL:    core.StringPtr("testString"),
				// 	Deprecated:          core.BoolPtr(true),
				// 	PackageVersion:      core.StringPtr("testString"),
				// 	State:               stateModel,
				// 	VersionLocator:      core.StringPtr("testString"),
				// 	ConsoleURL:          core.StringPtr("testString"),
				// 	LongDescription:     core.StringPtr("testString"),
				// 	WhitelistedAccounts: []string{"testString"},
				// }

				// deploymentModel := &catalogmanagementv1.Deployment{
				// 	ID:               core.StringPtr("testString"),
				// 	Label:            core.StringPtr("testString"),
				// 	Name:             core.StringPtr("testString"),
				// 	ShortDescription: core.StringPtr("testString"),
				// 	LongDescription:  core.StringPtr("testString"),
				// 	Metadata:         make(map[string]interface{}),
				// 	Tags:             []string{"testString"},
				// 	Created:          CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Updated:          CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// }

				// planModel := &catalogmanagementv1.Plan{
				// 	ID:                 core.StringPtr("testString"),
				// 	Label:              core.StringPtr("testString"),
				// 	Name:               core.StringPtr("testString"),
				// 	ShortDescription:   core.StringPtr("testString"),
				// 	LongDescription:    core.StringPtr("testString"),
				// 	Metadata:           make(map[string]interface{}),
				// 	Tags:               []string{"testString"},
				// 	AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
				// 	Created:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Updated:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Deployments:        []catalogmanagementv1.Deployment{*deploymentModel},
				// }

				// kindModel := &catalogmanagementv1.Kind{
				// 	ID:                 core.StringPtr("testString"),
				// 	FormatKind:         core.StringPtr("testString"),
				// 	TargetKind:         core.StringPtr("testString"),
				// 	Metadata:           make(map[string]interface{}),
				// 	InstallDescription: core.StringPtr("testString"),
				// 	Tags:               []string{"testString"},
				// 	AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
				// 	Created:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Updated:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Versions:           []catalogmanagementv1.Version{*versionModel},
				// 	Plans:              []catalogmanagementv1.Plan{*planModel},
				// }

				// repoInfoModel := &catalogmanagementv1.RepoInfo{
				// 	Token: core.StringPtr("testString"),
				// 	Type:  core.StringPtr("testString"),
				// }

				createOfferingOptions := &catalogmanagementv1.CreateOfferingOptions{
					CatalogIdentifier: &catalogID,
					// ID:                            core.StringPtr("testString"),
					// Rev:                           core.StringPtr("testString"),
					// URL:                           core.StringPtr("testString"),
					// CRN:                           core.StringPtr("testString"),
					Label: core.StringPtr(expectedLabel),
					Name:  core.StringPtr(expectedOfferingName),
					// OfferingIconURL:               core.StringPtr("testString"),
					// OfferingDocsURL:               core.StringPtr("testString"),
					// OfferingSupportURL:            core.StringPtr("testString"),
					// Tags:                          []string{"testString"},
					// Keywords:                      []string{"testString"},
					// Rating:                        ratingModel,
					// Created:                       CreateMockDateTime("2019-01-01T12:00:00.000Z"),
					// Updated:                       CreateMockDateTime("2019-01-01T12:00:00.000Z"),
					// ShortDescription:              core.StringPtr("testString"),
					// LongDescription:               core.StringPtr("testString"),
					// Features:                      []catalogmanagementv1.Feature{*featureModel},
					// Kinds:                         []catalogmanagementv1.Kind{*kindModel},
					// PermitRequestIBMPublicPublish: core.BoolPtr(true),
					// IBMPublishApproved:            core.BoolPtr(true),
					// PublicPublishApproved:         core.BoolPtr(true),
					// PublicOriginalCRN:             core.StringPtr("testString"),
					// PublishPublicCRN:              core.StringPtr("testString"),
					// PortalApprovalRecord:          core.StringPtr("testString"),
					// PortalUIURL:                   core.StringPtr("testString"),
					// CatalogID:                     core.StringPtr("testString"),
					// CatalogName:                   core.StringPtr("testString"),
					// Metadata:                      make(map[string]interface{}),
					// Disclaimer:                    core.StringPtr("testString"),
					// Hidden:                        core.BoolPtr(true),
					// Provider:                      core.StringPtr("testString"),
					// RepoInfo:                      repoInfoModel,
				}

				offering, response, err := catalogManagementService.CreateOffering(createOfferingOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(offering).ToNot(BeNil())

				offeringID = *offering.ID

				Expect(*offering.Name).To(Equal(expectedOfferingName))
				Expect(*offering.URL).To(Equal(fmt.Sprintf(expectedOfferingURL, catalogID, offeringID)))
				Expect(*offering.Label).To(Equal(expectedLabel))
			})
		})

		FDescribe(`GetOffering - Get offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOffering(getOfferingOptions *GetOfferingOptions)`, func() {
				getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{
					CatalogIdentifier: core.StringPtr(catalogID),
					OfferingID:        core.StringPtr(offeringID),
				}

				offering, response, err := catalogManagementService.GetOffering(getOfferingOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offering).ToNot(BeNil())
				Expect(*offering.URL).To(Equal(fmt.Sprintf(expectedOfferingURL, catalogID, offeringID)))
			})
			It(`GetOffering(getOfferingOptions *GetOfferingOptions) - Fail to get an offering that does not exist`, func() {
				getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{
					CatalogIdentifier: core.StringPtr(fakeName),
					OfferingID:        core.StringPtr(fakeName),
				}

				offering, response, err := catalogManagementService.GetOffering(getOfferingOptions)

				Expect(err).ToNot(BeNil())
				Expect(response.StatusCode).To(Equal(404))
				Expect(offering).To(BeNil())
			})
		})

		FDescribe(`ListOfferings - Get list of offerings`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ListOfferings(listOfferingsOptions *ListOfferingsOptions)`, func() {
				const (
					expectedLimit         int64 = 100
					expectedTotalCount    int64 = 1
					expectedResourceCount int64 = 1
					expectedResouceLen          = 1
					expectedFirst               = "/api/v1-beta/catalogs/%s/offerings?limit=100&sort=label"
					expectedLast                = "/api/v1-beta/catalogs/%s/offerings?limit=100&sort=label"
				)

				listOfferingsOptions := &catalogmanagementv1.ListOfferingsOptions{
					CatalogIdentifier: &catalogID,
				}

				offeringSearchResult, response, err := catalogManagementService.ListOfferings(listOfferingsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offeringSearchResult).ToNot(BeNil())
				Expect(*offeringSearchResult.Offset).To(BeZero())
				Expect(*offeringSearchResult.Limit).To(Equal(expectedLimit))
				Expect(*offeringSearchResult.TotalCount).To(Equal(expectedTotalCount))
				Expect(*offeringSearchResult.ResourceCount).To(Equal(expectedResourceCount))
				Expect(*offeringSearchResult.First).To(Equal(fmt.Sprintf(expectedFirst, catalogID)))
				Expect(*offeringSearchResult.Last).To(Equal(fmt.Sprintf(expectedLast, catalogID)))
				Expect(len(offeringSearchResult.Resources)).To(Equal(expectedResouceLen))

				Expect(*offeringSearchResult.Resources[0].ID).To(Equal(offeringID))
				Expect(*offeringSearchResult.Resources[0].URL).To(Equal(fmt.Sprintf(expectedOfferingURL, catalogID, offeringID)))
				Expect(*offeringSearchResult.Resources[0].Label).To(Equal(expectedLabel))
				Expect(*offeringSearchResult.Resources[0].Name).To(Equal(expectedOfferingName))
				Expect(*offeringSearchResult.Resources[0].CatalogID).To(Equal(catalogID))
				Expect(*offeringSearchResult.Resources[0].CatalogName).To(Equal(expectedLabel))

			})
		})

		FDescribe(`GetConsumptionOfferings - Get consumption offerings`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetConsumptionOfferings(getConsumptionOfferingsOptions *GetConsumptionOfferingsOptions)`, func() {

				getConsumptionOfferingsOptions := &catalogmanagementv1.GetConsumptionOfferingsOptions{
					// Digest:        core.BoolPtr(true),
					// Catalog:       &catalogID,
					// Select:        core.StringPtr("all"),
					// IncludeHidden: core.BoolPtr(true),
					// Limit:         core.Int64Ptr(int64(1000)),
					// Offset:        core.Int64Ptr(int64(38)),
				}

				offeringSearchResult, response, err := catalogManagementService.GetConsumptionOfferings(getConsumptionOfferingsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offeringSearchResult).ToNot(BeNil())
				Expect(*offeringSearchResult.Offset).To(BeZero())
				Expect(*offeringSearchResult.Limit).ToNot(BeZero())
				Expect(*offeringSearchResult.TotalCount).ToNot(BeZero())
				Expect(offeringSearchResult.Last).ToNot(BeNil())
				Expect(offeringSearchResult.Prev).To(BeNil())
				Expect(offeringSearchResult.Next).ToNot(BeNil())
				Expect(len(offeringSearchResult.Resources)).ToNot(BeZero())
			})
		})

		Describe(`ImportOfferingVersion - Import offering version`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ImportOfferingVersion(importOfferingVersionOptions *ImportOfferingVersionOptions)`, func() {

				importOfferingVersionOptions := &catalogmanagementv1.ImportOfferingVersionOptions{
					CatalogIdentifier: core.StringPtr(catalogID),
					OfferingID:        core.StringPtr(offeringID),
				}

				offering, response, err := catalogManagementService.ImportOfferingVersion(importOfferingVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(offering).ToNot(BeNil())
			})
		})

		FDescribe(`ImportOffering - Import offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ImportOffering(importOfferingOptions *ImportOfferingOptions)`, func() {
				Expect(catalogID).ToNot(BeEmpty())

				const (
					expectedOfferingName       = "node-red-operator-certified"
					expectedOfferingLabel      = "Node-RED Operator"
					expectedOfferingTargetKind = "roks"
					expectedOfferingVersion    = "0.0.2"
					expectedOfferingVersions   = 1
					expectedOfferingKinds      = 1
					expectedOfferingShortDesc  = "Node-RED is a programming tool for wiring together hardware devices, APIs and online services in new and interesting ways."
					expectedOfferingURL        = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings/%s"
					expectedOfferingZipURL     = "https://github.com/rhm-samples/node-red-operator/blob/nodered-1.2.8/node-red-operator/bundle/0.0.2/manifests/node-red-operator.v0.0.2.clusterserviceversion.yaml"
				)

				importOfferingOptions := &catalogmanagementv1.ImportOfferingOptions{
					CatalogIdentifier: core.StringPtr(catalogID),
					// Tags:              []string{"testString"},
					TargetKinds: []string{"roks"},
					// Content:           CreateMockByteArray("This is a mock byte array value."),
					Zipurl:        core.StringPtr(expectedOfferingZipURL),
					OfferingID:    core.StringPtr(offeringID),
					TargetVersion: core.StringPtr("0.0.2"),
					// IncludeConfig:     core.BoolPtr(true),
					// IsVsi:             core.BoolPtr(true),
					RepoType:   core.StringPtr("public_git"),
					XAuthToken: core.StringPtr(gitToken),
				}

				offering, response, err := catalogManagementService.ImportOffering(importOfferingOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(offering).ToNot(BeNil())

				// Expect(response.StatusCode).To(Equal(201))
				// Expect(*offering.Name).To(Equal(expectedOfferingName))
				// Expect(*offering.URL).To(Equal(fmt.Sprintf(expectedOfferingURL, catalogID, offeringID)))
				// Expect(*offering.Label).To(Equal(expectedOfferingLabel))
				// Expect(*offering.ShortDescription).To(Equal(expectedOfferingShortDesc))
				// Expect(*offering.CatalogName).To(Equal(expectedLabel))
				// Expect(*offering.CatalogID).To(Equal(catalogID))
				// Expect(len(offering.Kinds)).To(Equal(expectedOfferingKinds))
				// Expect(*offering.Kinds[0].TargetKind).To(Equal(expectedOfferingTargetKind))
				// Expect(len(offering.Kinds[0].Versions)).To(Equal(expectedOfferingVersions))
				// Expect(*offering.Kinds[0].Versions[0].Version).To(Equal(expectedOfferingVersion))

			})
		})

		Describe(`ReloadOffering - Reload offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ReloadOffering(reloadOfferingOptions *ReloadOfferingOptions)`, func() {

				reloadOfferingOptions := &catalogmanagementv1.ReloadOfferingOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					OfferingID:        core.StringPtr("testString"),
					TargetVersion:     core.StringPtr("testString"),
					Tags:              []string{"testString"},
					TargetKinds:       []string{"testString"},
					Content:           CreateMockByteArray("This is a mock byte array value."),
					Zipurl:            core.StringPtr("testString"),
					RepoType:          core.StringPtr("testString"),
				}

				offering, response, err := catalogManagementService.ReloadOffering(reloadOfferingOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(offering).ToNot(BeNil())

			})
		})

		Describe(`ReplaceOffering - Update offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ReplaceOffering(replaceOfferingOptions *ReplaceOfferingOptions)`, func() {

				// ratingModel := &catalogmanagementv1.Rating{
				// 	OneStarCount:   core.Int64Ptr(int64(38)),
				// 	TwoStarCount:   core.Int64Ptr(int64(38)),
				// 	ThreeStarCount: core.Int64Ptr(int64(38)),
				// 	FourStarCount:  core.Int64Ptr(int64(38)),
				// }

				// featureModel := &catalogmanagementv1.Feature{
				// 	Title:       core.StringPtr("testString"),
				// 	Description: core.StringPtr("testString"),
				// }

				// configurationModel := &catalogmanagementv1.Configuration{
				// 	Key:             core.StringPtr("testString"),
				// 	Type:            core.StringPtr("testString"),
				// 	DefaultValue:    core.StringPtr("testString"),
				// 	ValueConstraint: core.StringPtr("testString"),
				// 	Description:     core.StringPtr("testString"),
				// 	Required:        core.BoolPtr(true),
				// 	Options:         []interface{}{"testString"},
				// 	Hidden:          core.BoolPtr(true),
				// }

				// validationModel := &catalogmanagementv1.Validation{
				// 	Validated:     CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Requested:     CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	State:         core.StringPtr("testString"),
				// 	LastOperation: core.StringPtr("testString"),
				// 	Target:        make(map[string]interface{}),
				// }

				// resourceModel := &catalogmanagementv1.Resource{
				// 	Type:  core.StringPtr("mem"),
				// 	Value: core.StringPtr("testString"),
				// }

				// scriptModel := &catalogmanagementv1.Script{
				// 	Instructions:     core.StringPtr("testString"),
				// 	Script:           core.StringPtr("testString"),
				// 	ScriptPermission: core.StringPtr("testString"),
				// 	DeleteScript:     core.StringPtr("testString"),
				// 	Scope:            core.StringPtr("testString"),
				// }

				// versionEntitlementModel := &catalogmanagementv1.VersionEntitlement{
				// 	ProviderName:  core.StringPtr("testString"),
				// 	ProviderID:    core.StringPtr("testString"),
				// 	ProductID:     core.StringPtr("testString"),
				// 	PartNumbers:   []string{"testString"},
				// 	ImageRepoName: core.StringPtr("testString"),
				// }

				// licenseModel := &catalogmanagementv1.License{
				// 	ID:          core.StringPtr("testString"),
				// 	Name:        core.StringPtr("testString"),
				// 	Type:        core.StringPtr("testString"),
				// 	URL:         core.StringPtr("testString"),
				// 	Description: core.StringPtr("testString"),
				// }

				// stateModel := &catalogmanagementv1.State{
				// 	Current:          core.StringPtr("testString"),
				// 	CurrentEntered:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Pending:          core.StringPtr("testString"),
				// 	PendingRequested: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Previous:         core.StringPtr("testString"),
				// }

				// versionModel := &catalogmanagementv1.Version{
				// 	ID:                  core.StringPtr("testString"),
				// 	Rev:                 core.StringPtr("testString"),
				// 	CRN:                 core.StringPtr("testString"),
				// 	Version:             core.StringPtr("testString"),
				// 	Sha:                 core.StringPtr("testString"),
				// 	Created:             CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Updated:             CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	OfferingID:          core.StringPtr("testString"),
				// 	CatalogID:           core.StringPtr("testString"),
				// 	KindID:              core.StringPtr("testString"),
				// 	Tags:                []string{"testString"},
				// 	RepoURL:             core.StringPtr("testString"),
				// 	SourceURL:           core.StringPtr("testString"),
				// 	TgzURL:              core.StringPtr("testString"),
				// 	Configuration:       []catalogmanagementv1.Configuration{*configurationModel},
				// 	Metadata:            make(map[string]interface{}),
				// 	Validation:          validationModel,
				// 	RequiredResources:   []catalogmanagementv1.Resource{*resourceModel},
				// 	SingleInstance:      core.BoolPtr(true),
				// 	Install:             scriptModel,
				// 	PreInstall:          []catalogmanagementv1.Script{*scriptModel},
				// 	Entitlement:         versionEntitlementModel,
				// 	Licenses:            []catalogmanagementv1.License{*licenseModel},
				// 	ImageManifestURL:    core.StringPtr("testString"),
				// 	Deprecated:          core.BoolPtr(true),
				// 	PackageVersion:      core.StringPtr("testString"),
				// 	State:               stateModel,
				// 	VersionLocator:      core.StringPtr("testString"),
				// 	ConsoleURL:          core.StringPtr("testString"),
				// 	LongDescription:     core.StringPtr("testString"),
				// 	WhitelistedAccounts: []string{"testString"},
				// }

				// deploymentModel := &catalogmanagementv1.Deployment{
				// 	ID:               core.StringPtr("testString"),
				// 	Label:            core.StringPtr("testString"),
				// 	Name:             core.StringPtr("testString"),
				// 	ShortDescription: core.StringPtr("testString"),
				// 	LongDescription:  core.StringPtr("testString"),
				// 	Metadata:         make(map[string]interface{}),
				// 	Tags:             []string{"testString"},
				// 	Created:          CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Updated:          CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// }

				// planModel := &catalogmanagementv1.Plan{
				// 	ID:                 core.StringPtr("testString"),
				// 	Label:              core.StringPtr("testString"),
				// 	Name:               core.StringPtr("testString"),
				// 	ShortDescription:   core.StringPtr("testString"),
				// 	LongDescription:    core.StringPtr("testString"),
				// 	Metadata:           make(map[string]interface{}),
				// 	Tags:               []string{"testString"},
				// 	AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
				// 	Created:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Updated:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Deployments:        []catalogmanagementv1.Deployment{*deploymentModel},
				// }

				// kindModel := &catalogmanagementv1.Kind{
				// 	ID:                 core.StringPtr("testString"),
				// 	FormatKind:         core.StringPtr("testString"),
				// 	TargetKind:         core.StringPtr("testString"),
				// 	Metadata:           make(map[string]interface{}),
				// 	InstallDescription: core.StringPtr("testString"),
				// 	Tags:               []string{"testString"},
				// 	AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
				// 	Created:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Updated:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Versions:           []catalogmanagementv1.Version{*versionModel},
				// 	Plans:              []catalogmanagementv1.Plan{*planModel},
				// }

				// repoInfoModel := &catalogmanagementv1.RepoInfo{
				// 	Token: core.StringPtr("testString"),
				// 	Type:  core.StringPtr("testString"),
				// }

				replaceOfferingOptions := &catalogmanagementv1.ReplaceOfferingOptions{
					CatalogIdentifier: core.StringPtr(catalogID),
					OfferingID:        core.StringPtr(offeringID),
					ID:                core.StringPtr(offeringID),
					Rev:               core.StringPtr("2-8a366bcd418e559674c83cb73bbd9bb6"),
					// URL:                           core.StringPtr("testString"),
					// CRN:                           core.StringPtr("testString"),
					// Label:                         core.StringPtr("testString"),
					// Name:                          core.StringPtr("testString"),
					// OfferingIconURL:               core.StringPtr("testString"),
					// OfferingDocsURL:               core.StringPtr("testString"),
					// OfferingSupportURL:            core.StringPtr("testString"),
					Tags: []string{"tag"},
					// Keywords:                      []string{"testString"},
					// Rating:                        ratingModel,
					// Created:                       CreateMockDateTime("2019-01-01T12:00:00.000Z"),
					// Updated:                       CreateMockDateTime("2019-01-01T12:00:00.000Z"),
					// ShortDescription:              core.StringPtr("testString"),
					// LongDescription:               core.StringPtr("testString"),
					// Features:                      []catalogmanagementv1.Feature{*featureModel},
					// Kinds:                         []catalogmanagementv1.Kind{*kindModel},
					// PermitRequestIBMPublicPublish: core.BoolPtr(true),
					// IBMPublishApproved:            core.BoolPtr(true),
					// PublicPublishApproved:         core.BoolPtr(true),
					// PublicOriginalCRN:             core.StringPtr("testString"),
					// PublishPublicCRN:              core.StringPtr("testString"),
					// PortalApprovalRecord:          core.StringPtr("testString"),
					// PortalUIURL:                   core.StringPtr("testString"),
					// CatalogID:                     core.StringPtr("testString"),
					// CatalogName:                   core.StringPtr("testString"),
					// Metadata:                      make(map[string]interface{}),
					// Disclaimer:                    core.StringPtr("testString"),
					// Hidden:                        core.BoolPtr(true),
					// Provider:                      core.StringPtr("testString"),
					// RepoInfo:                      repoInfoModel,
				}

				offering, response, err := catalogManagementService.ReplaceOffering(replaceOfferingOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offering).ToNot(BeNil())

			})
		})

		FDescribe(`GetOfferingAudit - Get offering audit log`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOfferingAudit(getOfferingAuditOptions *GetOfferingAuditOptions)`, func() {

				getOfferingAuditOptions := &catalogmanagementv1.GetOfferingAuditOptions{
					CatalogIdentifier: core.StringPtr(catalogID),
					OfferingID:        core.StringPtr(offeringID),
				}

				auditLog, response, err := catalogManagementService.GetOfferingAudit(getOfferingAuditOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLog).ToNot(BeNil())

			})
		})

		Describe(`ReplaceOfferingIcon - Upload icon for offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ReplaceOfferingIcon(replaceOfferingIconOptions *ReplaceOfferingIconOptions)`, func() {

				replaceOfferingIconOptions := &catalogmanagementv1.ReplaceOfferingIconOptions{
					CatalogIdentifier: core.StringPtr(catalogID),
					OfferingID:        core.StringPtr(offeringID),
					FileName:          core.StringPtr("test_icon.png"),
				}

				offering, response, err := catalogManagementService.ReplaceOfferingIcon(replaceOfferingIconOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offering).ToNot(BeNil())

			})
		})

		// NOTE: not permitted to approve promotion
		Describe(`UpdateOfferingIBM - Allow offering to be published`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`UpdateOfferingIBM(updateOfferingIBMOptions *UpdateOfferingIBMOptions)`, func() {

				updateOfferingIBMOptions := &catalogmanagementv1.UpdateOfferingIBMOptions{
					CatalogIdentifier: core.StringPtr(catalogID),
					OfferingID:        core.StringPtr(offeringID),
					ApprovalType:      core.StringPtr("allow_request"),
					Approved:          core.StringPtr("true"),
				}

				approvalResult, response, err := catalogManagementService.UpdateOfferingIBM(updateOfferingIBMOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(approvalResult).ToNot(BeNil())

			})
		})

		Describe(`GetOfferingUpdates - Get version updates`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOfferingUpdates(getOfferingUpdatesOptions *GetOfferingUpdatesOptions)`, func() {

				getOfferingUpdatesOptions := &catalogmanagementv1.GetOfferingUpdatesOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					OfferingID:        core.StringPtr("testString"),
					Kind:              core.StringPtr("testString"),
					Version:           core.StringPtr("testString"),
					ClusterID:         core.StringPtr("testString"),
					Region:            core.StringPtr("testString"),
					ResourceGroupID:   core.StringPtr("testString"),
					Namespace:         core.StringPtr("testString"),
				}

				versionUpdateDescriptor, response, err := catalogManagementService.GetOfferingUpdates(getOfferingUpdatesOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(versionUpdateDescriptor).ToNot(BeNil())

			})
		})

		Describe(`GetOfferingAbout - Get version about information`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOfferingAbout(getOfferingAboutOptions *GetOfferingAboutOptions)`, func() {

				getOfferingAboutOptions := &catalogmanagementv1.GetOfferingAboutOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				result, response, err := catalogManagementService.GetOfferingAbout(getOfferingAboutOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

			})
		})

		Describe(`GetOfferingLicense - Get version license content`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOfferingLicense(getOfferingLicenseOptions *GetOfferingLicenseOptions)`, func() {

				getOfferingLicenseOptions := &catalogmanagementv1.GetOfferingLicenseOptions{
					VersionLocID: core.StringPtr("testString"),
					LicenseID:    core.StringPtr("testString"),
				}

				result, response, err := catalogManagementService.GetOfferingLicense(getOfferingLicenseOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

			})
		})

		Describe(`GetOfferingContainerImages - Get version's container images`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOfferingContainerImages(getOfferingContainerImagesOptions *GetOfferingContainerImagesOptions)`, func() {

				getOfferingContainerImagesOptions := &catalogmanagementv1.GetOfferingContainerImagesOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				imageManifest, response, err := catalogManagementService.GetOfferingContainerImages(getOfferingContainerImagesOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(imageManifest).ToNot(BeNil())

			})
		})

		Describe(`DeprecateVersion - Deprecate version`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`DeprecateVersion(deprecateVersionOptions *DeprecateVersionOptions)`, func() {

				deprecateVersionOptions := &catalogmanagementv1.DeprecateVersionOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				response, err := catalogManagementService.DeprecateVersion(deprecateVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`AccountPublishVersion - Publish version to account members`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`AccountPublishVersion(accountPublishVersionOptions *AccountPublishVersionOptions)`, func() {

				accountPublishVersionOptions := &catalogmanagementv1.AccountPublishVersionOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				response, err := catalogManagementService.AccountPublishVersion(accountPublishVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`IBMPublishVersion - Publish version to IBMers in public catalog`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`IBMPublishVersion(ibmPublishVersionOptions *IBMPublishVersionOptions)`, func() {

				ibmPublishVersionOptions := &catalogmanagementv1.IBMPublishVersionOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				response, err := catalogManagementService.IBMPublishVersion(ibmPublishVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`PublicPublishVersion - Publish version to all users in public catalog`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`PublicPublishVersion(publicPublishVersionOptions *PublicPublishVersionOptions)`, func() {

				publicPublishVersionOptions := &catalogmanagementv1.PublicPublishVersionOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				response, err := catalogManagementService.PublicPublishVersion(publicPublishVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`CommitVersion - Commit version`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`CommitVersion(commitVersionOptions *CommitVersionOptions)`, func() {

				commitVersionOptions := &catalogmanagementv1.CommitVersionOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				response, err := catalogManagementService.CommitVersion(commitVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		Describe(`CopyVersion - Copy version to new target kind`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`CopyVersion(copyVersionOptions *CopyVersionOptions)`, func() {

				copyVersionOptions := &catalogmanagementv1.CopyVersionOptions{
					VersionLocID: core.StringPtr("testString"),
					Tags:         []string{"testString"},
					TargetKinds:  []string{"testString"},
					Content:      CreateMockByteArray("This is a mock byte array value."),
				}

				response, err := catalogManagementService.CopyVersion(copyVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		Describe(`GetOfferingWorkingCopy - Create working copy of version`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOfferingWorkingCopy(getOfferingWorkingCopyOptions *GetOfferingWorkingCopyOptions)`, func() {

				getOfferingWorkingCopyOptions := &catalogmanagementv1.GetOfferingWorkingCopyOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				version, response, err := catalogManagementService.GetOfferingWorkingCopy(getOfferingWorkingCopyOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(version).ToNot(BeNil())

			})
		})

		Describe(`GetVersion - Get offering/kind/version 'branch'`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetVersion(getVersionOptions *GetVersionOptions)`, func() {

				getVersionOptions := &catalogmanagementv1.GetVersionOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				offering, response, err := catalogManagementService.GetVersion(getVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offering).ToNot(BeNil())

			})
		})

		// *************** CLUSTERS
		FDescribe(`GetCluster - Get kubernetes cluster`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetCluster(getClusterOptions *GetClusterOptions)`, func() {

				getClusterOptions := &catalogmanagementv1.GetClusterOptions{
					ClusterID:         core.StringPtr(clusterID),
					Region:            core.StringPtr("US"),
					XAuthRefreshToken: core.StringPtr(gitToken),
				}

				clusterInfo, response, err := catalogManagementService.GetCluster(getClusterOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(clusterInfo).ToNot(BeNil())
				Expect(*clusterInfo.ID).To(Equal(clusterID))
			})
		})

		Describe(`GetNamespaces - Get cluster namespaces`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetNamespaces(getNamespacesOptions *GetNamespacesOptions)`, func() {

				getNamespacesOptions := &catalogmanagementv1.GetNamespacesOptions{
					ClusterID:         core.StringPtr(clusterID),
					Region:            core.StringPtr("US"),
					XAuthRefreshToken: core.StringPtr(gitToken),
					// Limit:             core.Int64Ptr(int64(1000)),
					// Offset:            core.Int64Ptr(int64(38)),
				}

				namespaceSearchResult, response, err := catalogManagementService.GetNamespaces(getNamespacesOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(namespaceSearchResult).ToNot(BeNil())

			})
		})

		Describe(`DeployOperators - Deploy operators`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`DeployOperators(deployOperatorsOptions *DeployOperatorsOptions)`, func() {

				deployOperatorsOptions := &catalogmanagementv1.DeployOperatorsOptions{
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID:         core.StringPtr("testString"),
					Region:            core.StringPtr("testString"),
					Namespaces:        []string{"testString"},
					AllNamespaces:     core.BoolPtr(true),
					VersionLocatorID:  core.StringPtr("testString"),
				}

				operatorDeployResult, response, err := catalogManagementService.DeployOperators(deployOperatorsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(operatorDeployResult).ToNot(BeNil())

			})
		})

		Describe(`ListOperators - List operators`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ListOperators(listOperatorsOptions *ListOperatorsOptions)`, func() {

				listOperatorsOptions := &catalogmanagementv1.ListOperatorsOptions{
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID:         core.StringPtr("testString"),
					Region:            core.StringPtr("testString"),
					VersionLocatorID:  core.StringPtr("testString"),
				}

				operatorDeployResult, response, err := catalogManagementService.ListOperators(listOperatorsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(operatorDeployResult).ToNot(BeNil())

			})
		})

		Describe(`ReplaceOperators - Update operators`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ReplaceOperators(replaceOperatorsOptions *ReplaceOperatorsOptions)`, func() {

				replaceOperatorsOptions := &catalogmanagementv1.ReplaceOperatorsOptions{
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID:         core.StringPtr("testString"),
					Region:            core.StringPtr("testString"),
					Namespaces:        []string{"testString"},
					AllNamespaces:     core.BoolPtr(true),
					VersionLocatorID:  core.StringPtr("testString"),
				}

				operatorDeployResult, response, err := catalogManagementService.ReplaceOperators(replaceOperatorsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(operatorDeployResult).ToNot(BeNil())

			})
		})

		Describe(`InstallVersion - Install version`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`InstallVersion(installVersionOptions *InstallVersionOptions)`, func() {

				deployRequestBodySchematicsModel := &catalogmanagementv1.DeployRequestBodySchematics{
					Name:            core.StringPtr("testString"),
					Description:     core.StringPtr("testString"),
					Tags:            []string{"testString"},
					ResourceGroupID: core.StringPtr("testString"),
				}

				installVersionOptions := &catalogmanagementv1.InstallVersionOptions{
					VersionLocID:      core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID:         core.StringPtr("testString"),
					Region:            core.StringPtr("testString"),
					Namespace:         core.StringPtr("testString"),
					OverrideValues:    make(map[string]interface{}),
					EntitlementApikey: core.StringPtr("testString"),
					Schematics:        deployRequestBodySchematicsModel,
					Script:            core.StringPtr("testString"),
					ScriptID:          core.StringPtr("testString"),
					VersionLocatorID:  core.StringPtr("testString"),
					VcenterID:         core.StringPtr("testString"),
					VcenterUser:       core.StringPtr("testString"),
					VcenterPassword:   core.StringPtr("testString"),
					VcenterLocation:   core.StringPtr("testString"),
					VcenterDatastore:  core.StringPtr("testString"),
				}

				response, err := catalogManagementService.InstallVersion(installVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`PreinstallVersion - Pre-install version`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`PreinstallVersion(preinstallVersionOptions *PreinstallVersionOptions)`, func() {

				deployRequestBodySchematicsModel := &catalogmanagementv1.DeployRequestBodySchematics{
					Name:            core.StringPtr("testString"),
					Description:     core.StringPtr("testString"),
					Tags:            []string{"testString"},
					ResourceGroupID: core.StringPtr("testString"),
				}

				preinstallVersionOptions := &catalogmanagementv1.PreinstallVersionOptions{
					VersionLocID:      core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID:         core.StringPtr("testString"),
					Region:            core.StringPtr("testString"),
					Namespace:         core.StringPtr("testString"),
					OverrideValues:    make(map[string]interface{}),
					EntitlementApikey: core.StringPtr("testString"),
					Schematics:        deployRequestBodySchematicsModel,
					Script:            core.StringPtr("testString"),
					ScriptID:          core.StringPtr("testString"),
					VersionLocatorID:  core.StringPtr("testString"),
					VcenterID:         core.StringPtr("testString"),
					VcenterUser:       core.StringPtr("testString"),
					VcenterPassword:   core.StringPtr("testString"),
					VcenterLocation:   core.StringPtr("testString"),
					VcenterDatastore:  core.StringPtr("testString"),
				}

				response, err := catalogManagementService.PreinstallVersion(preinstallVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`GetPreinstall - Get version pre-install status`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetPreinstall(getPreinstallOptions *GetPreinstallOptions)`, func() {

				getPreinstallOptions := &catalogmanagementv1.GetPreinstallOptions{
					VersionLocID:      core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID:         core.StringPtr("testString"),
					Region:            core.StringPtr("testString"),
					Namespace:         core.StringPtr("testString"),
				}

				installStatus, response, err := catalogManagementService.GetPreinstall(getPreinstallOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(installStatus).ToNot(BeNil())

			})
		})

		Describe(`ValidateInstall - Validate offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ValidateInstall(validateInstallOptions *ValidateInstallOptions)`, func() {

				deployRequestBodySchematicsModel := &catalogmanagementv1.DeployRequestBodySchematics{
					Name:            core.StringPtr("testString"),
					Description:     core.StringPtr("testString"),
					Tags:            []string{"testString"},
					ResourceGroupID: core.StringPtr("testString"),
				}

				validateInstallOptions := &catalogmanagementv1.ValidateInstallOptions{
					VersionLocID:      core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID:         core.StringPtr("testString"),
					Region:            core.StringPtr("testString"),
					Namespace:         core.StringPtr("testString"),
					OverrideValues:    make(map[string]interface{}),
					EntitlementApikey: core.StringPtr("testString"),
					Schematics:        deployRequestBodySchematicsModel,
					Script:            core.StringPtr("testString"),
					ScriptID:          core.StringPtr("testString"),
					VersionLocatorID:  core.StringPtr("testString"),
					VcenterID:         core.StringPtr("testString"),
					VcenterUser:       core.StringPtr("testString"),
					VcenterPassword:   core.StringPtr("testString"),
					VcenterLocation:   core.StringPtr("testString"),
					VcenterDatastore:  core.StringPtr("testString"),
				}

				response, err := catalogManagementService.ValidateInstall(validateInstallOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`GetValidationStatus - Get offering install status`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetValidationStatus(getValidationStatusOptions *GetValidationStatusOptions)`, func() {

				getValidationStatusOptions := &catalogmanagementv1.GetValidationStatusOptions{
					VersionLocID:      core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
				}

				validation, response, err := catalogManagementService.GetValidationStatus(getValidationStatusOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(validation).ToNot(BeNil())

			})
		})

		Describe(`GetOverrideValues - Get override values`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOverrideValues(getOverrideValuesOptions *GetOverrideValuesOptions)`, func() {

				getOverrideValuesOptions := &catalogmanagementv1.GetOverrideValuesOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				result, response, err := catalogManagementService.GetOverrideValues(getOverrideValuesOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

			})
		})

		Describe(`SearchObjects - List objects across catalogs`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`SearchObjects(searchObjectsOptions *SearchObjectsOptions)`, func() {

				searchObjectsOptions := &catalogmanagementv1.SearchObjectsOptions{
					Query:    core.StringPtr("testString"),
					Limit:    core.Int64Ptr(int64(1000)),
					Offset:   core.Int64Ptr(int64(38)),
					Collapse: core.BoolPtr(true),
					Digest:   core.BoolPtr(true),
				}

				objectSearchResult, response, err := catalogManagementService.SearchObjects(searchObjectsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(objectSearchResult).ToNot(BeNil())

			})
		})

		Describe(`ListObjects - List objects within a catalog`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ListObjects(listObjectsOptions *ListObjectsOptions)`, func() {

				listObjectsOptions := &catalogmanagementv1.ListObjectsOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					Limit:             core.Int64Ptr(int64(1000)),
					Offset:            core.Int64Ptr(int64(38)),
					Name:              core.StringPtr("testString"),
					Sort:              core.StringPtr("testString"),
				}

				objectListResult, response, err := catalogManagementService.ListObjects(listObjectsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(objectListResult).ToNot(BeNil())

			})
		})

		Describe(`CreateObject - Create catalog object`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`CreateObject(createObjectOptions *CreateObjectOptions)`, func() {

				// publishObjectModel := &catalogmanagementv1.PublishObject{
				// 	PermitIBMPublicPublish: core.BoolPtr(true),
				// 	IBMApproved:            core.BoolPtr(true),
				// 	PublicApproved:         core.BoolPtr(true),
				// 	PortalApprovalRecord:   core.StringPtr("testString"),
				// 	PortalURL:              core.StringPtr("testString"),
				// }

				// stateModel := &catalogmanagementv1.State{
				// 	Current:          core.StringPtr("testString"),
				// 	CurrentEntered:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Pending:          core.StringPtr("testString"),
				// 	PendingRequested: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				// 	Previous:         core.StringPtr("testString"),
				// }

				createObjectOptions := &catalogmanagementv1.CreateObjectOptions{
					CatalogIdentifier: core.StringPtr(catalogID),
					// ID:                   core.StringPtr("testString"),
					Name: core.StringPtr(fakeName),
					// Rev:                  core.StringPtr("testString"),
					// CRN:                  core.StringPtr("testString"),
					// URL:                  core.StringPtr("testString"),
					// ParentID:             core.StringPtr("testString"),
					// LabelI18n:            core.StringPtr("testString"),
					Label: core.StringPtr(expectedLabel),
					// Tags:                 []string{"testString"},
					// Created:              CreateMockDateTime("2019-01-01T12:00:00.000Z"),
					// Updated:              CreateMockDateTime("2019-01-01T12:00:00.000Z"),
					// ShortDescription:     core.StringPtr("testString"),
					// ShortDescriptionI18n: core.StringPtr("testString"),
					Kind: core.StringPtr("terraform"),
					// Publish:              publishObjectModel,
					// State:                stateModel,
					// CatalogID:            core.StringPtr("testString"),
					// CatalogName:          core.StringPtr("testString"),
					// Data:                 make(map[string]interface{}),
				}

				catalogObject, response, err := catalogManagementService.CreateObject(createObjectOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(catalogObject).ToNot(BeNil())

			})
		})

		Describe(`GetObject - Get catalog object`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetObject(getObjectOptions *GetObjectOptions)`, func() {

				getObjectOptions := &catalogmanagementv1.GetObjectOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
				}

				catalogObject, response, err := catalogManagementService.GetObject(getObjectOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(catalogObject).ToNot(BeNil())

			})
		})

		Describe(`ReplaceObject - Update catalog object`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ReplaceObject(replaceObjectOptions *ReplaceObjectOptions)`, func() {

				publishObjectModel := &catalogmanagementv1.PublishObject{
					PermitIBMPublicPublish: core.BoolPtr(true),
					IBMApproved:            core.BoolPtr(true),
					PublicApproved:         core.BoolPtr(true),
					PortalApprovalRecord:   core.StringPtr("testString"),
					PortalURL:              core.StringPtr("testString"),
				}

				stateModel := &catalogmanagementv1.State{
					Current:          core.StringPtr("testString"),
					CurrentEntered:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
					Pending:          core.StringPtr("testString"),
					PendingRequested: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
					Previous:         core.StringPtr("testString"),
				}

				replaceObjectOptions := &catalogmanagementv1.ReplaceObjectOptions{
					CatalogIdentifier:    core.StringPtr("testString"),
					ObjectIdentifier:     core.StringPtr("testString"),
					ID:                   core.StringPtr("testString"),
					Name:                 core.StringPtr("testString"),
					Rev:                  core.StringPtr("testString"),
					CRN:                  core.StringPtr("testString"),
					URL:                  core.StringPtr("testString"),
					ParentID:             core.StringPtr("testString"),
					LabelI18n:            core.StringPtr("testString"),
					Label:                core.StringPtr("testString"),
					Tags:                 []string{"testString"},
					Created:              CreateMockDateTime("2019-01-01T12:00:00.000Z"),
					Updated:              CreateMockDateTime("2019-01-01T12:00:00.000Z"),
					ShortDescription:     core.StringPtr("testString"),
					ShortDescriptionI18n: core.StringPtr("testString"),
					Kind:                 core.StringPtr("testString"),
					Publish:              publishObjectModel,
					State:                stateModel,
					CatalogID:            core.StringPtr("testString"),
					CatalogName:          core.StringPtr("testString"),
					Data:                 make(map[string]interface{}),
				}

				catalogObject, response, err := catalogManagementService.ReplaceObject(replaceObjectOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(catalogObject).ToNot(BeNil())

			})
		})

		Describe(`GetObjectAudit - Get catalog object audit log`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetObjectAudit(getObjectAuditOptions *GetObjectAuditOptions)`, func() {

				getObjectAuditOptions := &catalogmanagementv1.GetObjectAuditOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
				}

				auditLog, response, err := catalogManagementService.GetObjectAudit(getObjectAuditOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLog).ToNot(BeNil())

			})
		})

		Describe(`AccountPublishObject - Publish object to account`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`AccountPublishObject(accountPublishObjectOptions *AccountPublishObjectOptions)`, func() {

				accountPublishObjectOptions := &catalogmanagementv1.AccountPublishObjectOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
				}

				response, err := catalogManagementService.AccountPublishObject(accountPublishObjectOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`SharedPublishObject - Publish object to share with allow list`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`SharedPublishObject(sharedPublishObjectOptions *SharedPublishObjectOptions)`, func() {

				sharedPublishObjectOptions := &catalogmanagementv1.SharedPublishObjectOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
				}

				response, err := catalogManagementService.SharedPublishObject(sharedPublishObjectOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`IBMPublishObject - Publish object to share with IBMers`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`IBMPublishObject(ibmPublishObjectOptions *IBMPublishObjectOptions)`, func() {

				ibmPublishObjectOptions := &catalogmanagementv1.IBMPublishObjectOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
				}

				response, err := catalogManagementService.IBMPublishObject(ibmPublishObjectOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`PublicPublishObject - Publish object to share with all users`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`PublicPublishObject(publicPublishObjectOptions *PublicPublishObjectOptions)`, func() {

				publicPublishObjectOptions := &catalogmanagementv1.PublicPublishObjectOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
				}

				response, err := catalogManagementService.PublicPublishObject(publicPublishObjectOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))

			})
		})

		Describe(`CreateObjectAccess - Add account ID to object access list`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`CreateObjectAccess(createObjectAccessOptions *CreateObjectAccessOptions)`, func() {

				createObjectAccessOptions := &catalogmanagementv1.CreateObjectAccessOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
					AccountIdentifier: core.StringPtr("testString"),
				}

				response, err := catalogManagementService.CreateObjectAccess(createObjectAccessOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))

			})
		})

		Describe(`GetObjectAccess - Check for account ID in object access list`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetObjectAccess(getObjectAccessOptions *GetObjectAccessOptions)`, func() {

				getObjectAccessOptions := &catalogmanagementv1.GetObjectAccessOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
					AccountIdentifier: core.StringPtr("testString"),
				}

				objectAccess, response, err := catalogManagementService.GetObjectAccess(getObjectAccessOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(objectAccess).ToNot(BeNil())

			})
		})

		Describe(`GetObjectAccessList - Get object access list`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetObjectAccessList(getObjectAccessListOptions *GetObjectAccessListOptions)`, func() {

				getObjectAccessListOptions := &catalogmanagementv1.GetObjectAccessListOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
					Limit:             core.Int64Ptr(int64(1000)),
					Offset:            core.Int64Ptr(int64(38)),
				}

				objectAccessListResult, response, err := catalogManagementService.GetObjectAccessList(getObjectAccessListOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(objectAccessListResult).ToNot(BeNil())

			})
		})

		Describe(`AddObjectAccessList - Add accounts to object access list`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`AddObjectAccessList(addObjectAccessListOptions *AddObjectAccessListOptions)`, func() {

				addObjectAccessListOptions := &catalogmanagementv1.AddObjectAccessListOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
					Accounts:          []string{"testString"},
				}

				accessListBulkResponse, response, err := catalogManagementService.AddObjectAccessList(addObjectAccessListOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(accessListBulkResponse).ToNot(BeNil())

			})
		})

		Describe(`CreateOfferingInstance - Create an offering resource instance`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`CreateOfferingInstance(createOfferingInstanceOptions *CreateOfferingInstanceOptions)`, func() {

				createOfferingInstanceOptions := &catalogmanagementv1.CreateOfferingInstanceOptions{
					XAuthRefreshToken:    core.StringPtr("testString"),
					ID:                   core.StringPtr("testString"),
					URL:                  core.StringPtr("testString"),
					CRN:                  core.StringPtr("testString"),
					Label:                core.StringPtr("testString"),
					CatalogID:            core.StringPtr("testString"),
					OfferingID:           core.StringPtr("testString"),
					KindFormat:           core.StringPtr("testString"),
					Version:              core.StringPtr("testString"),
					ClusterID:            core.StringPtr("testString"),
					ClusterRegion:        core.StringPtr("testString"),
					ClusterNamespaces:    []string{"testString"},
					ClusterAllNamespaces: core.BoolPtr(true),
				}

				offeringInstance, response, err := catalogManagementService.CreateOfferingInstance(createOfferingInstanceOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(offeringInstance).ToNot(BeNil())

			})
		})

		Describe(`GetOfferingInstance - Get Offering Instance`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOfferingInstance(getOfferingInstanceOptions *GetOfferingInstanceOptions)`, func() {

				getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{
					InstanceIdentifier: core.StringPtr("testString"),
				}

				offeringInstance, response, err := catalogManagementService.GetOfferingInstance(getOfferingInstanceOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offeringInstance).ToNot(BeNil())

			})
		})

		Describe(`PutOfferingInstance - Update Offering Instance`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`PutOfferingInstance(putOfferingInstanceOptions *PutOfferingInstanceOptions)`, func() {

				putOfferingInstanceOptions := &catalogmanagementv1.PutOfferingInstanceOptions{
					InstanceIdentifier:   core.StringPtr("testString"),
					XAuthRefreshToken:    core.StringPtr("testString"),
					ID:                   core.StringPtr("testString"),
					URL:                  core.StringPtr("testString"),
					CRN:                  core.StringPtr("testString"),
					Label:                core.StringPtr("testString"),
					CatalogID:            core.StringPtr("testString"),
					OfferingID:           core.StringPtr("testString"),
					KindFormat:           core.StringPtr("testString"),
					Version:              core.StringPtr("testString"),
					ClusterID:            core.StringPtr("testString"),
					ClusterRegion:        core.StringPtr("testString"),
					ClusterNamespaces:    []string{"testString"},
					ClusterAllNamespaces: core.BoolPtr(true),
				}

				offeringInstance, response, err := catalogManagementService.PutOfferingInstance(putOfferingInstanceOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offeringInstance).ToNot(BeNil())

			})
		})

		// *************** DELETE
		Describe(`DeleteVersion - Delete version`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`DeleteVersion(deleteVersionOptions *DeleteVersionOptions)`, func() {

				deleteVersionOptions := &catalogmanagementv1.DeleteVersionOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				response, err := catalogManagementService.DeleteVersion(deleteVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		Describe(`DeleteOperators - Delete operators`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`DeleteOperators(deleteOperatorsOptions *DeleteOperatorsOptions)`, func() {

				deleteOperatorsOptions := &catalogmanagementv1.DeleteOperatorsOptions{
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID:         core.StringPtr("testString"),
					Region:            core.StringPtr("testString"),
					VersionLocatorID:  core.StringPtr("testString"),
				}

				response, err := catalogManagementService.DeleteOperators(deleteOperatorsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		Describe(`DeleteOfferingInstance - Delete a version instance`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`DeleteOfferingInstance(deleteOfferingInstanceOptions *DeleteOfferingInstanceOptions)`, func() {

				deleteOfferingInstanceOptions := &catalogmanagementv1.DeleteOfferingInstanceOptions{
					InstanceIdentifier: core.StringPtr("testString"),
					XAuthRefreshToken:  core.StringPtr("testString"),
				}

				response, err := catalogManagementService.DeleteOfferingInstance(deleteOfferingInstanceOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		FDescribe(`DeleteOffering - Delete offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`DeleteOffering(deleteOfferingOptions *DeleteOfferingOptions)`, func() {

				deleteOfferingOptions := &catalogmanagementv1.DeleteOfferingOptions{
					CatalogIdentifier: &catalogID,
					OfferingID:        &offeringID,
				}

				response, err := catalogManagementService.DeleteOffering(deleteOfferingOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
			It(`DeleteOffering(deleteOfferingOptions *DeleteOfferingOptions) - Fail to delete an offering that does not exist`, func() {

				deleteOfferingOptions := &catalogmanagementv1.DeleteOfferingOptions{
					CatalogIdentifier: core.StringPtr(fakeName),
					OfferingID:        core.StringPtr(fakeName),
				}

				response, err := catalogManagementService.DeleteOffering(deleteOfferingOptions)

				Expect(err).ToNot(BeNil())
				Expect(response.StatusCode).To(Equal(404))
			})
		})

		Describe(`DeleteObjectAccessList - Delete accounts from object access list`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`DeleteObjectAccessList(deleteObjectAccessListOptions *DeleteObjectAccessListOptions)`, func() {

				deleteObjectAccessListOptions := &catalogmanagementv1.DeleteObjectAccessListOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
					Accounts:          []string{"testString"},
				}

				accessListBulkResponse, response, err := catalogManagementService.DeleteObjectAccessList(deleteObjectAccessListOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(accessListBulkResponse).ToNot(BeNil())

			})
		})

		Describe(`DeleteObjectAccess - Remove account ID from object access list`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`DeleteObjectAccess(deleteObjectAccessOptions *DeleteObjectAccessOptions)`, func() {

				deleteObjectAccessOptions := &catalogmanagementv1.DeleteObjectAccessOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier:  core.StringPtr("testString"),
					AccountIdentifier: core.StringPtr("testString"),
				}

				response, err := catalogManagementService.DeleteObjectAccess(deleteObjectAccessOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		Describe(`DeleteObject - Delete catalog object`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`DeleteObject(deleteObjectOptions *DeleteObjectOptions)`, func() {

				deleteObjectOptions := &catalogmanagementv1.DeleteObjectOptions{
					CatalogIdentifier: core.StringPtr(catalogID),
					ObjectIdentifier:  core.StringPtr(objectID),
				}

				response, err := catalogManagementService.DeleteObject(deleteObjectOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		FDescribe(`DeleteCatalog - Delete catalog`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			FIt(`DeleteCatalog(deleteCatalogOptions *DeleteCatalogOptions)`, func() {

				deleteCatalogOptions := &catalogmanagementv1.DeleteCatalogOptions{
					CatalogIdentifier: &catalogID,
				}

				response, err := catalogManagementService.DeleteCatalog(deleteCatalogOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
			})
			It(`DeleteCatalog(deleteCatalogOptions *DeleteCatalogOptions) - Fail to delete a catalog that does not exist`, func() {

				deleteCatalogOptions := &catalogmanagementv1.DeleteCatalogOptions{
					CatalogIdentifier: core.StringPtr(fakeName),
				}

				response, err := catalogManagementService.DeleteCatalog(deleteCatalogOptions)

				Expect(err).ToNot(BeNil())
				Expect(response.StatusCode).To(Equal(404))
			})
		})
	})

})

//
// Utility functions are declared in the unit test file
//
