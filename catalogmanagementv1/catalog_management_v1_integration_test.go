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
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	common "github.com/IBM/platform-services-go-sdk/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the catalogmanagementv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`CatalogManagementV1 Integration Tests`, func() {

	const (
		externalConfigFile   = "../catalog_mgmt.env"
		expectedShortDesc    = "test"
		expectedURL          = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s"
		expectedOfferingsURL = "https://cm.globalcatalog.test.cloud.ibm.com/api/v1-beta/catalogs/%s/offerings"
		expectedAccount      = "67d27f28d43948b2b3bda9138f251a13"
	)

	var (
		err                      error
		catalogManagementService *catalogmanagementv1.CatalogManagementV1
		configLoaded             bool = false
		serviceURL               string
		config                   map[string]string
		testCatalogID            string
		testOfferingID           string
		expectedLabel            = fmt.Sprintf("integration-test-%d", time.Now().Unix())
		gitToken                 string
		refreshToken             string
		testVersionInstanceId    string
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping...")
		}
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
			gitToken = config["GIT_TOKEN"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Get Refresh Token`, func() {
		It("successfully creates a refresh token", func() {
			authenticator, err := core.GetAuthenticatorFromEnvironment("catalog_management")
			iamAuthenticator := authenticator.(*core.IamAuthenticator)

			Expect(err).To(BeNil())

			tokenServerResponse, err := iamAuthenticator.RequestToken()
			refreshToken = tokenServerResponse.RefreshToken

			Expect(err).To(BeNil())
			Expect(refreshToken).ToNot(BeNil())
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			catalogManagementServiceOptions := &catalogmanagementv1.CatalogManagementV1Options{}

			catalogManagementService, err = catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(catalogManagementServiceOptions)

			Expect(err).To(BeNil())
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(catalogManagementService.Service.Options.URL).To(Equal(serviceURL))

		})
	})
	/*
		Describe(`GetCatalogAccount - Get catalog account settings`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetCatalogAccount(getCatalogAccountOptions *GetCatalogAccountOptions)`, func() {

				getCatalogAccountOptions := &catalogmanagementv1.GetCatalogAccountOptions{
				}

				account, response, err := catalogManagementService.GetCatalogAccount(getCatalogAccountOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(account).ToNot(BeNil())

			})
		})

		Describe(`UpdateCatalogAccount - Update account settings`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`UpdateCatalogAccount(updateCatalogAccountOptions *UpdateCatalogAccountOptions)`, func() {

				filterTermsModel := &catalogmanagementv1.FilterTerms{
					FilterTerms: []string{"testString"},
				}

				categoryFilterModel := &catalogmanagementv1.CategoryFilter{
					Include: core.BoolPtr(true),
					Filter: filterTermsModel,
				}

				idFilterModel := &catalogmanagementv1.IDFilter{
					Include: filterTermsModel,
					Exclude: filterTermsModel,
				}

				filtersModel := &catalogmanagementv1.Filters{
					IncludeAll: core.BoolPtr(true),
					CategoryFilters: make(map[string]catalogmanagementv1.CategoryFilter),
					IDFilters: idFilterModel,
				}
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				updateCatalogAccountOptions := &catalogmanagementv1.UpdateCatalogAccountOptions{
					ID: core.StringPtr("testString"),
					HideIBMCloudCatalog: core.BoolPtr(true),
					AccountFilters: filtersModel,
				}

				response, err := catalogManagementService.UpdateCatalogAccount(updateCatalogAccountOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		Describe(`GetCatalogAccountAudit - Get catalog account audit log`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetCatalogAccountAudit(getCatalogAccountAuditOptions *GetCatalogAccountAuditOptions)`, func() {

				getCatalogAccountAuditOptions := &catalogmanagementv1.GetCatalogAccountAuditOptions{
				}

				auditLog, response, err := catalogManagementService.GetCatalogAccountAudit(getCatalogAccountAuditOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLog).ToNot(BeNil())

			})
		})

		Describe(`GetCatalogAccountFilters - Get catalog account filters`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetCatalogAccountFilters(getCatalogAccountFiltersOptions *GetCatalogAccountFiltersOptions)`, func() {

				getCatalogAccountFiltersOptions := &catalogmanagementv1.GetCatalogAccountFiltersOptions{
					Catalog: core.StringPtr("testString"),
				}

				accumulatedFilters, response, err := catalogManagementService.GetCatalogAccountFilters(getCatalogAccountFiltersOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(accumulatedFilters).ToNot(BeNil())

			})
		})

		Describe(`ListCatalogs - Get list of catalogs`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ListCatalogs(listCatalogsOptions *ListCatalogsOptions)`, func() {

				listCatalogsOptions := &catalogmanagementv1.ListCatalogsOptions{
				}

				catalogSearchResult, response, err := catalogManagementService.ListCatalogs(listCatalogsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(catalogSearchResult).ToNot(BeNil())

			})
		})
	*/
	Describe(`CreateCatalog - Create a catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCatalog(createCatalogOptions *CreateCatalogOptions)`, func() {

			options := catalogManagementService.NewCreateCatalogOptions()
			options.SetLabel(expectedLabel)
			options.SetShortDescription(expectedShortDesc)
			result, response, err := catalogManagementService.CreateCatalog(options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateCatalog() result:\n%s\n", common.ToJSON(result))

			Expect(*result.Label).To(Equal(expectedLabel))
			Expect(*result.ShortDescription).To(Equal(expectedShortDesc))
			Expect(*result.URL).To(Equal(fmt.Sprintf(expectedURL, *result.ID)))
			Expect(*result.OfferingsURL).To(Equal(fmt.Sprintf(expectedOfferingsURL, *result.ID)))
			Expect(*result.OwningAccount).To(Equal(expectedAccount))
			Expect(*result.CatalogFilters.IncludeAll).To(BeFalse())
			Expect(len(result.CatalogFilters.CategoryFilters)).To(BeZero())
			Expect(result.CatalogFilters.IDFilters.Include).To(BeNil())
			Expect(result.CatalogFilters.IDFilters.Exclude).To(BeNil())

			testCatalogID = *result.ID

		})
	})
	/*
		Describe(`GetCatalog - Get catalog`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {

				getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{
					CatalogIdentifier: core.StringPtr("testString"),
				}

				catalog, response, err := catalogManagementService.GetCatalog(getCatalogOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(catalog).ToNot(BeNil())

			})
		})

		Describe(`ReplaceCatalog - Update catalog`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ReplaceCatalog(replaceCatalogOptions *ReplaceCatalogOptions)`, func() {

				featureModel := &catalogmanagementv1.Feature{
					Title: core.StringPtr("testString"),
					Description: core.StringPtr("testString"),
				}

				filterTermsModel := &catalogmanagementv1.FilterTerms{
					FilterTerms: []string{"testString"},
				}

				categoryFilterModel := &catalogmanagementv1.CategoryFilter{
					Include: core.BoolPtr(true),
					Filter: filterTermsModel,
				}

				idFilterModel := &catalogmanagementv1.IDFilter{
					Include: filterTermsModel,
					Exclude: filterTermsModel,
				}

				filtersModel := &catalogmanagementv1.Filters{
					IncludeAll: core.BoolPtr(true),
					CategoryFilters: make(map[string]catalogmanagementv1.CategoryFilter),
					IDFilters: idFilterModel,
				}
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				syndicationClusterModel := &catalogmanagementv1.SyndicationCluster{
					Region: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					ResourceGroupName: core.StringPtr("testString"),
					Type: core.StringPtr("testString"),
					Namespaces: []string{"testString"},
					AllNamespaces: core.BoolPtr(true),
				}

				syndicationHistoryModel := &catalogmanagementv1.SyndicationHistory{
					Namespaces: []string{"testString"},
					Clusters: []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel},
					LastRun: CreateMockDateTime(),
				}

				syndicationAuthorizationModel := &catalogmanagementv1.SyndicationAuthorization{
					Token: core.StringPtr("testString"),
					LastRun: CreateMockDateTime(),
				}

				syndicationResourceModel := &catalogmanagementv1.SyndicationResource{
					RemoveRelatedComponents: core.BoolPtr(true),
					Clusters: []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel},
					History: syndicationHistoryModel,
					Authorization: syndicationAuthorizationModel,
				}

				replaceCatalogOptions := &catalogmanagementv1.ReplaceCatalogOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					Rev: core.StringPtr("testString"),
					Label: core.StringPtr("testString"),
					ShortDescription: core.StringPtr("testString"),
					CatalogIconURL: core.StringPtr("testString"),
					Tags: []string{"testString"},
					Features: []catalogmanagementv1.Feature{*featureModel},
					Disabled: core.BoolPtr(true),
					ResourceGroupID: core.StringPtr("testString"),
					OwningAccount: core.StringPtr("testString"),
					CatalogFilters: filtersModel,
					SyndicationSettings: syndicationResourceModel,
				}

				catalog, response, err := catalogManagementService.ReplaceCatalog(replaceCatalogOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(catalog).ToNot(BeNil())

			})
		})

		Describe(`GetCatalogAudit - Get catalog audit log`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetCatalogAudit(getCatalogAuditOptions *GetCatalogAuditOptions)`, func() {

				getCatalogAuditOptions := &catalogmanagementv1.GetCatalogAuditOptions{
					CatalogIdentifier: core.StringPtr("testString"),
				}

				auditLog, response, err := catalogManagementService.GetCatalogAudit(getCatalogAuditOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLog).ToNot(BeNil())

			})
		})

		Describe(`GetEnterprise - Get enterprise settings`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions)`, func() {

				getEnterpriseOptions := &catalogmanagementv1.GetEnterpriseOptions{
					EnterpriseID: core.StringPtr("testString"),
				}

				enterprise, response, err := catalogManagementService.GetEnterprise(getEnterpriseOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(enterprise).ToNot(BeNil())

			})
		})

		Describe(`UpdateEnterprise - Update enterprise settings`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`UpdateEnterprise(updateEnterpriseOptions *UpdateEnterpriseOptions)`, func() {

				filterTermsModel := &catalogmanagementv1.FilterTerms{
					FilterTerms: []string{"testString"},
				}

				categoryFilterModel := &catalogmanagementv1.CategoryFilter{
					Include: core.BoolPtr(true),
					Filter: filterTermsModel,
				}

				idFilterModel := &catalogmanagementv1.IDFilter{
					Include: filterTermsModel,
					Exclude: filterTermsModel,
				}

				filtersModel := &catalogmanagementv1.Filters{
					IncludeAll: core.BoolPtr(true),
					CategoryFilters: make(map[string]catalogmanagementv1.CategoryFilter),
					IDFilters: idFilterModel,
				}
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				accountGroupModel := &catalogmanagementv1.AccountGroup{
					ID: core.StringPtr("testString"),
					AccountFilters: filtersModel,
				}

				enterpriseAccountGroupsModel := &catalogmanagementv1.EnterpriseAccountGroups{
					Keys: accountGroupModel,
				}

				updateEnterpriseOptions := &catalogmanagementv1.UpdateEnterpriseOptions{
					EnterpriseID: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					Rev: core.StringPtr("testString"),
					AccountFilters: filtersModel,
					AccountGroups: enterpriseAccountGroupsModel,
				}

				response, err := catalogManagementService.UpdateEnterprise(updateEnterpriseOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		Describe(`GetEnterpriseAudit - Get enterprise audit log`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetEnterpriseAudit(getEnterpriseAuditOptions *GetEnterpriseAuditOptions)`, func() {

				getEnterpriseAuditOptions := &catalogmanagementv1.GetEnterpriseAuditOptions{
					EnterpriseID: core.StringPtr("testString"),
				}

				auditLog, response, err := catalogManagementService.GetEnterpriseAudit(getEnterpriseAuditOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLog).ToNot(BeNil())

			})
		})

		Describe(`GetConsumptionOfferings - Get consumption offerings`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetConsumptionOfferings(getConsumptionOfferingsOptions *GetConsumptionOfferingsOptions)`, func() {

				getConsumptionOfferingsOptions := &catalogmanagementv1.GetConsumptionOfferingsOptions{
					Digest: core.BoolPtr(true),
					Catalog: core.StringPtr("testString"),
					Select: core.StringPtr("all"),
					IncludeHidden: core.BoolPtr(true),
					Limit: core.Int64Ptr(int64(1000)),
					Offset: core.Int64Ptr(int64(38)),
				}

				offeringSearchResult, response, err := catalogManagementService.GetConsumptionOfferings(getConsumptionOfferingsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offeringSearchResult).ToNot(BeNil())

			})
		})

		Describe(`ListOfferings - Get list of offerings`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ListOfferings(listOfferingsOptions *ListOfferingsOptions)`, func() {

				listOfferingsOptions := &catalogmanagementv1.ListOfferingsOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					Digest: core.BoolPtr(true),
					Limit: core.Int64Ptr(int64(1000)),
					Offset: core.Int64Ptr(int64(38)),
					Name: core.StringPtr("testString"),
					Sort: core.StringPtr("testString"),
				}

				offeringSearchResult, response, err := catalogManagementService.ListOfferings(listOfferingsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offeringSearchResult).ToNot(BeNil())

			})
		})

		Describe(`CreateOffering - Create offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`CreateOffering(createOfferingOptions *CreateOfferingOptions)`, func() {

				ratingModel := &catalogmanagementv1.Rating{
					OneStarCount: core.Int64Ptr(int64(38)),
					TwoStarCount: core.Int64Ptr(int64(38)),
					ThreeStarCount: core.Int64Ptr(int64(38)),
					FourStarCount: core.Int64Ptr(int64(38)),
				}

				featureModel := &catalogmanagementv1.Feature{
					Title: core.StringPtr("testString"),
					Description: core.StringPtr("testString"),
				}

				configurationModel := &catalogmanagementv1.Configuration{
					Key: core.StringPtr("testString"),
					Type: core.StringPtr("testString"),
					DefaultValue: core.StringPtr("testString"),
					ValueConstraint: core.StringPtr("testString"),
					Description: core.StringPtr("testString"),
					Required: core.BoolPtr(true),
					Options: []interface{}{"testString"},
					Hidden: core.BoolPtr(true),
				}

				validationModel := &catalogmanagementv1.Validation{
					Validated: CreateMockDateTime(),
					Requested: CreateMockDateTime(),
					State: core.StringPtr("testString"),
					LastOperation: core.StringPtr("testString"),
					Target: make(map[string]interface{}),
				}

				resourceModel := &catalogmanagementv1.Resource{
					Type: core.StringPtr("mem"),
					Value: core.StringPtr("testString"),
				}

				scriptModel := &catalogmanagementv1.Script{
					Instructions: core.StringPtr("testString"),
					Script: core.StringPtr("testString"),
					ScriptPermission: core.StringPtr("testString"),
					DeleteScript: core.StringPtr("testString"),
					Scope: core.StringPtr("testString"),
				}

				versionEntitlementModel := &catalogmanagementv1.VersionEntitlement{
					ProviderName: core.StringPtr("testString"),
					ProviderID: core.StringPtr("testString"),
					ProductID: core.StringPtr("testString"),
					PartNumbers: []string{"testString"},
					ImageRepoName: core.StringPtr("testString"),
				}

				licenseModel := &catalogmanagementv1.License{
					ID: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					Type: core.StringPtr("testString"),
					URL: core.StringPtr("testString"),
					Description: core.StringPtr("testString"),
				}

				stateModel := &catalogmanagementv1.State{
					Current: core.StringPtr("testString"),
					CurrentEntered: CreateMockDateTime(),
					Pending: core.StringPtr("testString"),
					PendingRequested: CreateMockDateTime(),
					Previous: core.StringPtr("testString"),
				}

				versionModel := &catalogmanagementv1.Version{
					ID: core.StringPtr("testString"),
					Rev: core.StringPtr("testString"),
					CRN: core.StringPtr("testString"),
					Version: core.StringPtr("testString"),
					Sha: core.StringPtr("testString"),
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
					OfferingID: core.StringPtr("testString"),
					CatalogID: core.StringPtr("testString"),
					KindID: core.StringPtr("testString"),
					Tags: []string{"testString"},
					RepoURL: core.StringPtr("testString"),
					SourceURL: core.StringPtr("testString"),
					TgzURL: core.StringPtr("testString"),
					Configuration: []catalogmanagementv1.Configuration{*configurationModel},
					Metadata: make(map[string]interface{}),
					Validation: validationModel,
					RequiredResources: []catalogmanagementv1.Resource{*resourceModel},
					SingleInstance: core.BoolPtr(true),
					Install: scriptModel,
					PreInstall: []catalogmanagementv1.Script{*scriptModel},
					Entitlement: versionEntitlementModel,
					Licenses: []catalogmanagementv1.License{*licenseModel},
					ImageManifestURL: core.StringPtr("testString"),
					Deprecated: core.BoolPtr(true),
					PackageVersion: core.StringPtr("testString"),
					State: stateModel,
					VersionLocator: core.StringPtr("testString"),
					ConsoleURL: core.StringPtr("testString"),
					LongDescription: core.StringPtr("testString"),
					WhitelistedAccounts: []string{"testString"},
				}

				deploymentModel := &catalogmanagementv1.Deployment{
					ID: core.StringPtr("testString"),
					Label: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					ShortDescription: core.StringPtr("testString"),
					LongDescription: core.StringPtr("testString"),
					Metadata: make(map[string]interface{}),
					Tags: []string{"testString"},
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
				}

				planModel := &catalogmanagementv1.Plan{
					ID: core.StringPtr("testString"),
					Label: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					ShortDescription: core.StringPtr("testString"),
					LongDescription: core.StringPtr("testString"),
					Metadata: make(map[string]interface{}),
					Tags: []string{"testString"},
					AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
					Deployments: []catalogmanagementv1.Deployment{*deploymentModel},
				}

				kindModel := &catalogmanagementv1.Kind{
					ID: core.StringPtr("testString"),
					FormatKind: core.StringPtr("testString"),
					TargetKind: core.StringPtr("testString"),
					Metadata: make(map[string]interface{}),
					InstallDescription: core.StringPtr("testString"),
					Tags: []string{"testString"},
					AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
					Versions: []catalogmanagementv1.Version{*versionModel},
					Plans: []catalogmanagementv1.Plan{*planModel},
				}

				repoInfoModel := &catalogmanagementv1.RepoInfo{
					Token: core.StringPtr("testString"),
					Type: core.StringPtr("testString"),
				}

				createOfferingOptions := &catalogmanagementv1.CreateOfferingOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					Rev: core.StringPtr("testString"),
					URL: core.StringPtr("testString"),
					CRN: core.StringPtr("testString"),
					Label: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					OfferingIconURL: core.StringPtr("testString"),
					OfferingDocsURL: core.StringPtr("testString"),
					OfferingSupportURL: core.StringPtr("testString"),
					Tags: []string{"testString"},
					Rating: ratingModel,
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
					ShortDescription: core.StringPtr("testString"),
					LongDescription: core.StringPtr("testString"),
					Features: []catalogmanagementv1.Feature{*featureModel},
					Kinds: []catalogmanagementv1.Kind{*kindModel},
					PermitRequestIBMPublicPublish: core.BoolPtr(true),
					IBMPublishApproved: core.BoolPtr(true),
					PublicPublishApproved: core.BoolPtr(true),
					PublicOriginalCRN: core.StringPtr("testString"),
					PublishPublicCRN: core.StringPtr("testString"),
					PortalApprovalRecord: core.StringPtr("testString"),
					PortalUIURL: core.StringPtr("testString"),
					CatalogID: core.StringPtr("testString"),
					CatalogName: core.StringPtr("testString"),
					Metadata: make(map[string]interface{}),
					Disclaimer: core.StringPtr("testString"),
					Hidden: core.BoolPtr(true),
					Provider: core.StringPtr("testString"),
					RepoInfo: repoInfoModel,
				}

				offering, response, err := catalogManagementService.CreateOffering(createOfferingOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(offering).ToNot(BeNil())

			})
		})

		Describe(`ImportOfferingVersion - Import offering version`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ImportOfferingVersion(importOfferingVersionOptions *ImportOfferingVersionOptions)`, func() {

				importOfferingVersionOptions := &catalogmanagementv1.ImportOfferingVersionOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					OfferingID: core.StringPtr("testString"),
					Tags: []string{"testString"},
					TargetKinds: []string{"testString"},
					Content: CreateMockByteArray("This is a mock byte array value."),
					Zipurl: core.StringPtr("testString"),
					TargetVersion: core.StringPtr("testString"),
					IncludeConfig: core.BoolPtr(true),
					RepoType: core.StringPtr("testString"),
				}

				offering, response, err := catalogManagementService.ImportOfferingVersion(importOfferingVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(offering).ToNot(BeNil())

			})
		})
	*/
	Describe(`ImportOffering - Import offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ImportOffering(importOfferingOptions *ImportOfferingOptions)`, func() {
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
			offeringOptions := catalogManagementService.NewImportOfferingOptions(testCatalogID)
			offeringOptions.SetZipurl(expectedOfferingZipURL)
			offeringOptions.SetXAuthToken(gitToken)
			offeringOptions.SetTargetKinds([]string{"roks"})
			offeringOptions.SetTargetVersion("0.0.2")
			offeringOptions.SetRepoType("public_git")
			offering, response, err := catalogManagementService.ImportOffering(offeringOptions)

			Expect(err).To(BeNil())

			testOfferingID = *offering.ID

			Expect(response.StatusCode).To(Equal(201))
			Expect(*offering.Name).To(Equal(expectedOfferingName))
			Expect(*offering.URL).To(Equal(fmt.Sprintf(expectedOfferingURL, testCatalogID, testOfferingID)))
			Expect(*offering.Label).To(Equal(expectedOfferingLabel))
			Expect(*offering.ShortDescription).To(Equal(expectedOfferingShortDesc))
			Expect(*offering.CatalogName).To(Equal(expectedLabel))
			Expect(*offering.CatalogID).To(Equal(testCatalogID))
			Expect(len(offering.Kinds)).To(Equal(expectedOfferingKinds))
			Expect(*offering.Kinds[0].TargetKind).To(Equal(expectedOfferingTargetKind))
			Expect(len(offering.Kinds[0].Versions)).To(Equal(expectedOfferingVersions))
			Expect(*offering.Kinds[0].Versions[0].Version).To(Equal(expectedOfferingVersion))
			Expect(*offering.Kinds[0].Versions[0].TgzURL).To(Equal(expectedOfferingZipURL))
		})
	})
	/*
		Describe(`ReloadOffering - Reload offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ReloadOffering(reloadOfferingOptions *ReloadOfferingOptions)`, func() {

				reloadOfferingOptions := &catalogmanagementv1.ReloadOfferingOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					OfferingID: core.StringPtr("testString"),
					TargetVersion: core.StringPtr("testString"),
					Tags: []string{"testString"},
					TargetKinds: []string{"testString"},
					Content: CreateMockByteArray("This is a mock byte array value."),
					Zipurl: core.StringPtr("testString"),
					RepoType: core.StringPtr("testString"),
				}

				offering, response, err := catalogManagementService.ReloadOffering(reloadOfferingOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(offering).ToNot(BeNil())

			})
		})

		Describe(`GetOffering - Get offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOffering(getOfferingOptions *GetOfferingOptions)`, func() {

				getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					OfferingID: core.StringPtr("testString"),
				}

				offering, response, err := catalogManagementService.GetOffering(getOfferingOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offering).ToNot(BeNil())

			})
		})

		Describe(`ReplaceOffering - Update offering`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ReplaceOffering(replaceOfferingOptions *ReplaceOfferingOptions)`, func() {

				ratingModel := &catalogmanagementv1.Rating{
					OneStarCount: core.Int64Ptr(int64(38)),
					TwoStarCount: core.Int64Ptr(int64(38)),
					ThreeStarCount: core.Int64Ptr(int64(38)),
					FourStarCount: core.Int64Ptr(int64(38)),
				}

				featureModel := &catalogmanagementv1.Feature{
					Title: core.StringPtr("testString"),
					Description: core.StringPtr("testString"),
				}

				configurationModel := &catalogmanagementv1.Configuration{
					Key: core.StringPtr("testString"),
					Type: core.StringPtr("testString"),
					DefaultValue: core.StringPtr("testString"),
					ValueConstraint: core.StringPtr("testString"),
					Description: core.StringPtr("testString"),
					Required: core.BoolPtr(true),
					Options: []interface{}{"testString"},
					Hidden: core.BoolPtr(true),
				}

				validationModel := &catalogmanagementv1.Validation{
					Validated: CreateMockDateTime(),
					Requested: CreateMockDateTime(),
					State: core.StringPtr("testString"),
					LastOperation: core.StringPtr("testString"),
					Target: make(map[string]interface{}),
				}

				resourceModel := &catalogmanagementv1.Resource{
					Type: core.StringPtr("mem"),
					Value: core.StringPtr("testString"),
				}

				scriptModel := &catalogmanagementv1.Script{
					Instructions: core.StringPtr("testString"),
					Script: core.StringPtr("testString"),
					ScriptPermission: core.StringPtr("testString"),
					DeleteScript: core.StringPtr("testString"),
					Scope: core.StringPtr("testString"),
				}

				versionEntitlementModel := &catalogmanagementv1.VersionEntitlement{
					ProviderName: core.StringPtr("testString"),
					ProviderID: core.StringPtr("testString"),
					ProductID: core.StringPtr("testString"),
					PartNumbers: []string{"testString"},
					ImageRepoName: core.StringPtr("testString"),
				}

				licenseModel := &catalogmanagementv1.License{
					ID: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					Type: core.StringPtr("testString"),
					URL: core.StringPtr("testString"),
					Description: core.StringPtr("testString"),
				}

				stateModel := &catalogmanagementv1.State{
					Current: core.StringPtr("testString"),
					CurrentEntered: CreateMockDateTime(),
					Pending: core.StringPtr("testString"),
					PendingRequested: CreateMockDateTime(),
					Previous: core.StringPtr("testString"),
				}

				versionModel := &catalogmanagementv1.Version{
					ID: core.StringPtr("testString"),
					Rev: core.StringPtr("testString"),
					CRN: core.StringPtr("testString"),
					Version: core.StringPtr("testString"),
					Sha: core.StringPtr("testString"),
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
					OfferingID: core.StringPtr("testString"),
					CatalogID: core.StringPtr("testString"),
					KindID: core.StringPtr("testString"),
					Tags: []string{"testString"},
					RepoURL: core.StringPtr("testString"),
					SourceURL: core.StringPtr("testString"),
					TgzURL: core.StringPtr("testString"),
					Configuration: []catalogmanagementv1.Configuration{*configurationModel},
					Metadata: make(map[string]interface{}),
					Validation: validationModel,
					RequiredResources: []catalogmanagementv1.Resource{*resourceModel},
					SingleInstance: core.BoolPtr(true),
					Install: scriptModel,
					PreInstall: []catalogmanagementv1.Script{*scriptModel},
					Entitlement: versionEntitlementModel,
					Licenses: []catalogmanagementv1.License{*licenseModel},
					ImageManifestURL: core.StringPtr("testString"),
					Deprecated: core.BoolPtr(true),
					PackageVersion: core.StringPtr("testString"),
					State: stateModel,
					VersionLocator: core.StringPtr("testString"),
					ConsoleURL: core.StringPtr("testString"),
					LongDescription: core.StringPtr("testString"),
					WhitelistedAccounts: []string{"testString"},
				}

				deploymentModel := &catalogmanagementv1.Deployment{
					ID: core.StringPtr("testString"),
					Label: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					ShortDescription: core.StringPtr("testString"),
					LongDescription: core.StringPtr("testString"),
					Metadata: make(map[string]interface{}),
					Tags: []string{"testString"},
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
				}

				planModel := &catalogmanagementv1.Plan{
					ID: core.StringPtr("testString"),
					Label: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					ShortDescription: core.StringPtr("testString"),
					LongDescription: core.StringPtr("testString"),
					Metadata: make(map[string]interface{}),
					Tags: []string{"testString"},
					AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
					Deployments: []catalogmanagementv1.Deployment{*deploymentModel},
				}

				kindModel := &catalogmanagementv1.Kind{
					ID: core.StringPtr("testString"),
					FormatKind: core.StringPtr("testString"),
					TargetKind: core.StringPtr("testString"),
					Metadata: make(map[string]interface{}),
					InstallDescription: core.StringPtr("testString"),
					Tags: []string{"testString"},
					AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
					Versions: []catalogmanagementv1.Version{*versionModel},
					Plans: []catalogmanagementv1.Plan{*planModel},
				}

				repoInfoModel := &catalogmanagementv1.RepoInfo{
					Token: core.StringPtr("testString"),
					Type: core.StringPtr("testString"),
				}

				replaceOfferingOptions := &catalogmanagementv1.ReplaceOfferingOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					OfferingID: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					Rev: core.StringPtr("testString"),
					URL: core.StringPtr("testString"),
					CRN: core.StringPtr("testString"),
					Label: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					OfferingIconURL: core.StringPtr("testString"),
					OfferingDocsURL: core.StringPtr("testString"),
					OfferingSupportURL: core.StringPtr("testString"),
					Tags: []string{"testString"},
					Rating: ratingModel,
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
					ShortDescription: core.StringPtr("testString"),
					LongDescription: core.StringPtr("testString"),
					Features: []catalogmanagementv1.Feature{*featureModel},
					Kinds: []catalogmanagementv1.Kind{*kindModel},
					PermitRequestIBMPublicPublish: core.BoolPtr(true),
					IBMPublishApproved: core.BoolPtr(true),
					PublicPublishApproved: core.BoolPtr(true),
					PublicOriginalCRN: core.StringPtr("testString"),
					PublishPublicCRN: core.StringPtr("testString"),
					PortalApprovalRecord: core.StringPtr("testString"),
					PortalUIURL: core.StringPtr("testString"),
					CatalogID: core.StringPtr("testString"),
					CatalogName: core.StringPtr("testString"),
					Metadata: make(map[string]interface{}),
					Disclaimer: core.StringPtr("testString"),
					Hidden: core.BoolPtr(true),
					Provider: core.StringPtr("testString"),
					RepoInfo: repoInfoModel,
				}

				offering, response, err := catalogManagementService.ReplaceOffering(replaceOfferingOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offering).ToNot(BeNil())

			})
		})

		Describe(`GetOfferingAudit - Get offering audit log`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetOfferingAudit(getOfferingAuditOptions *GetOfferingAuditOptions)`, func() {

				getOfferingAuditOptions := &catalogmanagementv1.GetOfferingAuditOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					OfferingID: core.StringPtr("testString"),
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
					CatalogIdentifier: core.StringPtr("testString"),
					OfferingID: core.StringPtr("testString"),
					FileName: core.StringPtr("testString"),
				}

				offering, response, err := catalogManagementService.ReplaceOfferingIcon(replaceOfferingIconOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offering).ToNot(BeNil())

			})
		})

		Describe(`UpdateOfferingIBM - Allow offering to be published`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`UpdateOfferingIBM(updateOfferingIBMOptions *UpdateOfferingIBMOptions)`, func() {

				updateOfferingIBMOptions := &catalogmanagementv1.UpdateOfferingIBMOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					OfferingID: core.StringPtr("testString"),
					ApprovalType: core.StringPtr("allow_request"),
					Approved: core.StringPtr("true"),
				}

				approvalResult, response, err := catalogManagementService.UpdateOfferingIBM(updateOfferingIBMOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(approvalResult).ToNot(BeNil())

			})
		})

		Describe(`GetVersionUpdates - Get version updates`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetVersionUpdates(getVersionUpdatesOptions *GetVersionUpdatesOptions)`, func() {

				getVersionUpdatesOptions := &catalogmanagementv1.GetVersionUpdatesOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					OfferingID: core.StringPtr("testString"),
					Kind: core.StringPtr("testString"),
					Version: core.StringPtr("testString"),
					ClusterID: core.StringPtr("testString"),
					Region: core.StringPtr("testString"),
					ResourceGroupID: core.StringPtr("testString"),
					Namespace: core.StringPtr("testString"),
				}

				versionUpdateDescriptor, response, err := catalogManagementService.GetVersionUpdates(getVersionUpdatesOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(versionUpdateDescriptor).ToNot(BeNil())

			})
		})

		Describe(`GetVersionAbout - Get version about information`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetVersionAbout(getVersionAboutOptions *GetVersionAboutOptions)`, func() {

				getVersionAboutOptions := &catalogmanagementv1.GetVersionAboutOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				result, response, err := catalogManagementService.GetVersionAbout(getVersionAboutOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

			})
		})

		Describe(`GetVersionLicense - Get version license content`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetVersionLicense(getVersionLicenseOptions *GetVersionLicenseOptions)`, func() {

				getVersionLicenseOptions := &catalogmanagementv1.GetVersionLicenseOptions{
					VersionLocID: core.StringPtr("testString"),
					LicenseID: core.StringPtr("testString"),
				}

				result, response, err := catalogManagementService.GetVersionLicense(getVersionLicenseOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

			})
		})

		Describe(`GetVersionContainerImages - Get version's container images`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetVersionContainerImages(getVersionContainerImagesOptions *GetVersionContainerImagesOptions)`, func() {

				getVersionContainerImagesOptions := &catalogmanagementv1.GetVersionContainerImagesOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				imageManifest, response, err := catalogManagementService.GetVersionContainerImages(getVersionContainerImagesOptions)

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
					Tags: []string{"testString"},
					TargetKinds: []string{"testString"},
					Content: CreateMockByteArray("This is a mock byte array value."),
				}

				response, err := catalogManagementService.CopyVersion(copyVersionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))

			})
		})

		Describe(`GetVersionWorkingCopy - Create working copy of version`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetVersionWorkingCopy(getVersionWorkingCopyOptions *GetVersionWorkingCopyOptions)`, func() {

				getVersionWorkingCopyOptions := &catalogmanagementv1.GetVersionWorkingCopyOptions{
					VersionLocID: core.StringPtr("testString"),
				}

				version, response, err := catalogManagementService.GetVersionWorkingCopy(getVersionWorkingCopyOptions)

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

		Describe(`GetRepos - List a repository's entries`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetRepos(getReposOptions *GetReposOptions)`, func() {

				getReposOptions := &catalogmanagementv1.GetReposOptions{
					Type: core.StringPtr("testString"),
					Repourl: core.StringPtr("testString"),
				}

				helmRepoList, response, err := catalogManagementService.GetRepos(getReposOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(helmRepoList).ToNot(BeNil())

			})
		})

		Describe(`GetRepo - Get repository contents`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetRepo(getRepoOptions *GetRepoOptions)`, func() {

				getRepoOptions := &catalogmanagementv1.GetRepoOptions{
					Type: core.StringPtr("testString"),
					Charturl: core.StringPtr("testString"),
				}

				helmPackage, response, err := catalogManagementService.GetRepo(getRepoOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(helmPackage).ToNot(BeNil())

			})
		})

		Describe(`GetCluster - Get kubernetes cluster`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetCluster(getClusterOptions *GetClusterOptions)`, func() {

				getClusterOptions := &catalogmanagementv1.GetClusterOptions{
					ClusterID: core.StringPtr("testString"),
					Region: core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
				}

				clusterInfo, response, err := catalogManagementService.GetCluster(getClusterOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(clusterInfo).ToNot(BeNil())

			})
		})

		Describe(`GetNamespaces - Get cluster namespaces`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`GetNamespaces(getNamespacesOptions *GetNamespacesOptions)`, func() {

				getNamespacesOptions := &catalogmanagementv1.GetNamespacesOptions{
					ClusterID: core.StringPtr("testString"),
					Region: core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(1000)),
					Offset: core.Int64Ptr(int64(38)),
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
					ClusterID: core.StringPtr("testString"),
					Region: core.StringPtr("testString"),
					Namespaces: []string{"testString"},
					AllNamespaces: core.BoolPtr(true),
					VersionLocatorID: core.StringPtr("testString"),
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
					ClusterID: core.StringPtr("testString"),
					Region: core.StringPtr("testString"),
					VersionLocatorID: core.StringPtr("testString"),
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
					ClusterID: core.StringPtr("testString"),
					Region: core.StringPtr("testString"),
					Namespaces: []string{"testString"},
					AllNamespaces: core.BoolPtr(true),
					VersionLocatorID: core.StringPtr("testString"),
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
					Name: core.StringPtr("testString"),
					Description: core.StringPtr("testString"),
					Tags: []string{"testString"},
					ResourceGroupID: core.StringPtr("testString"),
				}

				installVersionOptions := &catalogmanagementv1.InstallVersionOptions{
					VersionLocID: core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID: core.StringPtr("testString"),
					Region: core.StringPtr("testString"),
					Namespace: core.StringPtr("testString"),
					OverrideValues: make(map[string]interface{}),
					EntitlementApikey: core.StringPtr("testString"),
					Schematics: deployRequestBodySchematicsModel,
					Script: core.StringPtr("testString"),
					ScriptID: core.StringPtr("testString"),
					VersionLocatorID: core.StringPtr("testString"),
					VcenterID: core.StringPtr("testString"),
					VcenterUser: core.StringPtr("testString"),
					VcenterPassword: core.StringPtr("testString"),
					VcenterLocation: core.StringPtr("testString"),
					VcenterDatastore: core.StringPtr("testString"),
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
					Name: core.StringPtr("testString"),
					Description: core.StringPtr("testString"),
					Tags: []string{"testString"},
					ResourceGroupID: core.StringPtr("testString"),
				}

				preinstallVersionOptions := &catalogmanagementv1.PreinstallVersionOptions{
					VersionLocID: core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID: core.StringPtr("testString"),
					Region: core.StringPtr("testString"),
					Namespace: core.StringPtr("testString"),
					OverrideValues: make(map[string]interface{}),
					EntitlementApikey: core.StringPtr("testString"),
					Schematics: deployRequestBodySchematicsModel,
					Script: core.StringPtr("testString"),
					ScriptID: core.StringPtr("testString"),
					VersionLocatorID: core.StringPtr("testString"),
					VcenterID: core.StringPtr("testString"),
					VcenterUser: core.StringPtr("testString"),
					VcenterPassword: core.StringPtr("testString"),
					VcenterLocation: core.StringPtr("testString"),
					VcenterDatastore: core.StringPtr("testString"),
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
					VersionLocID: core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID: core.StringPtr("testString"),
					Region: core.StringPtr("testString"),
					Namespace: core.StringPtr("testString"),
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
					Name: core.StringPtr("testString"),
					Description: core.StringPtr("testString"),
					Tags: []string{"testString"},
					ResourceGroupID: core.StringPtr("testString"),
				}

				validateInstallOptions := &catalogmanagementv1.ValidateInstallOptions{
					VersionLocID: core.StringPtr("testString"),
					XAuthRefreshToken: core.StringPtr("testString"),
					ClusterID: core.StringPtr("testString"),
					Region: core.StringPtr("testString"),
					Namespace: core.StringPtr("testString"),
					OverrideValues: make(map[string]interface{}),
					EntitlementApikey: core.StringPtr("testString"),
					Schematics: deployRequestBodySchematicsModel,
					Script: core.StringPtr("testString"),
					ScriptID: core.StringPtr("testString"),
					VersionLocatorID: core.StringPtr("testString"),
					VcenterID: core.StringPtr("testString"),
					VcenterUser: core.StringPtr("testString"),
					VcenterPassword: core.StringPtr("testString"),
					VcenterLocation: core.StringPtr("testString"),
					VcenterDatastore: core.StringPtr("testString"),
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
					VersionLocID: core.StringPtr("testString"),
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

		Describe(`CreateLicenseEntitlement - Create license entitlement`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`CreateLicenseEntitlement(createLicenseEntitlementOptions *CreateLicenseEntitlementOptions)`, func() {

				createLicenseEntitlementOptions := &catalogmanagementv1.CreateLicenseEntitlementOptions{
					Name: core.StringPtr("testString"),
					EffectiveFrom: core.StringPtr("testString"),
					EffectiveUntil: core.StringPtr("testString"),
					VersionID: core.StringPtr("testString"),
					LicenseID: core.StringPtr("testString"),
					LicenseOwnerID: core.StringPtr("testString"),
					LicenseProviderID: core.StringPtr("testString"),
					LicenseProductID: core.StringPtr("testString"),
					AccountID: core.StringPtr("testString"),
				}

				licenseEntitlement, response, err := catalogManagementService.CreateLicenseEntitlement(createLicenseEntitlementOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(licenseEntitlement).ToNot(BeNil())

			})
		})

		Describe(`SearchObjects - List objects across catalogs`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`SearchObjects(searchObjectsOptions *SearchObjectsOptions)`, func() {

				searchObjectsOptions := &catalogmanagementv1.SearchObjectsOptions{
					Query: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(1000)),
					Offset: core.Int64Ptr(int64(38)),
					Collapse: core.BoolPtr(true),
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
					Limit: core.Int64Ptr(int64(1000)),
					Offset: core.Int64Ptr(int64(38)),
					Name: core.StringPtr("testString"),
					Sort: core.StringPtr("testString"),
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

				publishObjectModel := &catalogmanagementv1.PublishObject{
					PermitIBMPublicPublish: core.BoolPtr(true),
					IBMApproved: core.BoolPtr(true),
					PublicApproved: core.BoolPtr(true),
					PortalApprovalRecord: core.StringPtr("testString"),
					PortalURL: core.StringPtr("testString"),
				}

				stateModel := &catalogmanagementv1.State{
					Current: core.StringPtr("testString"),
					CurrentEntered: CreateMockDateTime(),
					Pending: core.StringPtr("testString"),
					PendingRequested: CreateMockDateTime(),
					Previous: core.StringPtr("testString"),
				}

				createObjectOptions := &catalogmanagementv1.CreateObjectOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					Rev: core.StringPtr("testString"),
					CRN: core.StringPtr("testString"),
					URL: core.StringPtr("testString"),
					ParentID: core.StringPtr("testString"),
					LabelI18n: core.StringPtr("testString"),
					Label: core.StringPtr("testString"),
					Tags: []string{"testString"},
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
					ShortDescription: core.StringPtr("testString"),
					ShortDescriptionI18n: core.StringPtr("testString"),
					Kind: core.StringPtr("testString"),
					Publish: publishObjectModel,
					State: stateModel,
					CatalogID: core.StringPtr("testString"),
					CatalogName: core.StringPtr("testString"),
					Data: make(map[string]interface{}),
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
					ObjectIdentifier: core.StringPtr("testString"),
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
					IBMApproved: core.BoolPtr(true),
					PublicApproved: core.BoolPtr(true),
					PortalApprovalRecord: core.StringPtr("testString"),
					PortalURL: core.StringPtr("testString"),
				}

				stateModel := &catalogmanagementv1.State{
					Current: core.StringPtr("testString"),
					CurrentEntered: CreateMockDateTime(),
					Pending: core.StringPtr("testString"),
					PendingRequested: CreateMockDateTime(),
					Previous: core.StringPtr("testString"),
				}

				replaceObjectOptions := &catalogmanagementv1.ReplaceObjectOptions{
					CatalogIdentifier: core.StringPtr("testString"),
					ObjectIdentifier: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					Name: core.StringPtr("testString"),
					Rev: core.StringPtr("testString"),
					CRN: core.StringPtr("testString"),
					URL: core.StringPtr("testString"),
					ParentID: core.StringPtr("testString"),
					LabelI18n: core.StringPtr("testString"),
					Label: core.StringPtr("testString"),
					Tags: []string{"testString"},
					Created: CreateMockDateTime(),
					Updated: CreateMockDateTime(),
					ShortDescription: core.StringPtr("testString"),
					ShortDescriptionI18n: core.StringPtr("testString"),
					Kind: core.StringPtr("testString"),
					Publish: publishObjectModel,
					State: stateModel,
					CatalogID: core.StringPtr("testString"),
					CatalogName: core.StringPtr("testString"),
					Data: make(map[string]interface{}),
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
					ObjectIdentifier: core.StringPtr("testString"),
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
					ObjectIdentifier: core.StringPtr("testString"),
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
					ObjectIdentifier: core.StringPtr("testString"),
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
					ObjectIdentifier: core.StringPtr("testString"),
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
					ObjectIdentifier: core.StringPtr("testString"),
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
					ObjectIdentifier: core.StringPtr("testString"),
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
					ObjectIdentifier: core.StringPtr("testString"),
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
					ObjectIdentifier: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(1000)),
					Offset: core.Int64Ptr(int64(38)),
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
					ObjectIdentifier: core.StringPtr("testString"),
					Accounts: []string{"testString"},
				}

				accessListBulkResponse, response, err := catalogManagementService.AddObjectAccessList(addObjectAccessListOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(accessListBulkResponse).ToNot(BeNil())

			})
		})
	*/
	Describe(`CreateVersionInstance - Create an offering version resource instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVersionInstance(createVersionInstanceOptions *CreateVersionInstanceOptions)`, func() {

			versionInstanceOptions := catalogManagementService.NewCreateVersionInstanceOptions(refreshToken)
			versionInstanceOptions.SetCatalogID(testCatalogID)
			versionInstanceOptions.SetOfferingID(testOfferingID)
			versionInstanceOptions.SetKindFormat("operator")
			versionInstanceOptions.SetVersion("0.0.2")
			versionInstanceOptions.SetClusterID("c07cn9h20vsge6l0e8o")
			versionInstanceOptions.SetClusterRegion("us-south")
			versionInstanceOptions.SetClusterNamespaces([]string{"sdk-test"})

			versionInstance, response, err := catalogManagementService.CreateVersionInstance(versionInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(versionInstance).ToNot(BeNil())

			testVersionInstanceId = *versionInstance.ID

			Expect(testVersionInstanceId).ToNot(BeNil())

		})
	})
	Describe(`GetVersionInstance - Get Version Instrance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVersionInstance(getVersionInstanceOptions *GetVersionInstanceOptions)`, func() {

			getVersionInstanceOptions := &catalogmanagementv1.GetVersionInstanceOptions{
				InstanceIdentifier: &testVersionInstanceId,
			}

			versionInstance, response, err := catalogManagementService.GetVersionInstance(getVersionInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(versionInstance).ToNot(BeNil())

		})
	})

	Describe(`PutVersionInstance - Update Version Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutVersionInstance(putVersionInstanceOptions *PutVersionInstanceOptions)`, func() {

			putVersionInstanceOptions := catalogManagementService.NewPutVersionInstanceOptions(testVersionInstanceId, refreshToken)

			versionInstance, response, err := catalogManagementService.PutVersionInstance(putVersionInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(versionInstance).ToNot(BeNil())

		})
	})

	Describe(`DeleteVersionInstance - Delete a version instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVersionInstance(deleteVersionInstanceOptions *DeleteVersionInstanceOptions)`, func() {

			deleteVersionInstanceOptions := catalogManagementService.NewDeleteVersionInstanceOptions(testVersionInstanceId)

			response, err := catalogManagementService.DeleteVersionInstance(deleteVersionInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})

	/* 	Describe(`DeleteVersion - Delete version`, func() {
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
	   				ClusterID: core.StringPtr("testString"),
	   				Region: core.StringPtr("testString"),
	   				VersionLocatorID: core.StringPtr("testString"),
	   			}

	   			response, err := catalogManagementService.DeleteOperators(deleteOperatorsOptions)

	   			Expect(err).To(BeNil())
	   			Expect(response.StatusCode).To(Equal(200))

	   		})
	   	})

	   	Describe(`DeleteOffering - Delete offering`, func() {
	   		BeforeEach(func() {
	   			shouldSkipTest()
	   		})
	   		It(`DeleteOffering(deleteOfferingOptions *DeleteOfferingOptions)`, func() {

	   			deleteOfferingOptions := &catalogmanagementv1.DeleteOfferingOptions{
	   				CatalogIdentifier: core.StringPtr("testString"),
	   				OfferingID: core.StringPtr("testString"),
	   			}

	   			response, err := catalogManagementService.DeleteOffering(deleteOfferingOptions)

	   			Expect(err).To(BeNil())
	   			Expect(response.StatusCode).To(Equal(200))

	   		})
	   	})

	   	Describe(`DeleteObjectAccessList - Delete accounts from object access list`, func() {
	   		BeforeEach(func() {
	   			shouldSkipTest()
	   		})
	   		It(`DeleteObjectAccessList(deleteObjectAccessListOptions *DeleteObjectAccessListOptions)`, func() {

	   			deleteObjectAccessListOptions := &catalogmanagementv1.DeleteObjectAccessListOptions{
	   				CatalogIdentifier: core.StringPtr("testString"),
	   				ObjectIdentifier: core.StringPtr("testString"),
	   				Accounts: []string{"testString"},
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
	   				ObjectIdentifier: core.StringPtr("testString"),
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
	   				CatalogIdentifier: core.StringPtr("testString"),
	   				ObjectIdentifier: core.StringPtr("testString"),
	   			}

	   			response, err := catalogManagementService.DeleteObject(deleteObjectOptions)

	   			Expect(err).To(BeNil())
	   			Expect(response.StatusCode).To(Equal(200))

	   		})
	   	})

	   	Describe(`DeleteLicenseEntitlement - Delete license entitlement`, func() {
	   		BeforeEach(func() {
	   			shouldSkipTest()
	   		})
	   		It(`DeleteLicenseEntitlement(deleteLicenseEntitlementOptions *DeleteLicenseEntitlementOptions)`, func() {

	   			deleteLicenseEntitlementOptions := &catalogmanagementv1.DeleteLicenseEntitlementOptions{
	   				EntitlementID: core.StringPtr("testString"),
	   				AccountID: core.StringPtr("testString"),
	   			}

	   			response, err := catalogManagementService.DeleteLicenseEntitlement(deleteLicenseEntitlementOptions)

	   			Expect(err).To(BeNil())
	   			Expect(response.StatusCode).To(Equal(200))

	   		})
	   	})
	*/

	Describe(`DeleteCatalog - Delete catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCatalog(deleteCatalogOptions *DeleteCatalogOptions)`, func() {

			deleteCatalogOptions := &catalogmanagementv1.DeleteCatalogOptions{
				CatalogIdentifier: &testCatalogID,
			}

			response, err := catalogManagementService.DeleteCatalog(deleteCatalogOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
