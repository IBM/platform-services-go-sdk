// +build integration

/**
 * (C) Copyright IBM Corp. 2023.
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
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`CatalogManagementV1 Integration Tests`, func() {
	const externalConfigFile = "../catalog_management_v1.env"

	var (
		err          error
		catalogManagementService *catalogmanagementv1.CatalogManagementV1
		serviceURL   string
		config       map[string]string

		// Variables to hold link values
		accountRevLink string
		catalogIDLink string
		catalogRevLink string
		objectIDLink string
		objectRevLink string
		offeringIDLink string
		offeringRevLink string
		versionIDLink string
		versionLocatorLink string
		versionRevLink string
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

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
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

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			catalogManagementService.EnableRetries(4, 30*time.Second)
		})
	})

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

			accountRevLink = *account.Rev
			fmt.Fprintf(GinkgoWriter, "Saved accountRevLink value: %v\n", accountRevLink)
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
				HideIBMCloudCatalog: core.BoolPtr(true),
				AccountFilters: filtersModel,
			}

			account, response, err := catalogManagementService.UpdateCatalogAccount(updateCatalogAccountOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(account).ToNot(BeNil())

			accountRevLink = *account.Rev
			fmt.Fprintf(GinkgoWriter, "Saved accountRevLink value: %v\n", accountRevLink)
		})
	})

	Describe(`CreateCatalog - Create a catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCatalog(createCatalogOptions *CreateCatalogOptions)`, func() {
			featureModel := &catalogmanagementv1.Feature{
				Title: core.StringPtr("testString"),
				TitleI18n: make(map[string]string),
				Description: core.StringPtr("testString"),
				DescriptionI18n: make(map[string]string),
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
				LastRun: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			syndicationAuthorizationModel := &catalogmanagementv1.SyndicationAuthorization{
				Token: core.StringPtr("testString"),
				LastRun: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			syndicationResourceModel := &catalogmanagementv1.SyndicationResource{
				RemoveRelatedComponents: core.BoolPtr(true),
				Clusters: []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel},
				History: syndicationHistoryModel,
				Authorization: syndicationAuthorizationModel,
			}

			createCatalogOptions := &catalogmanagementv1.CreateCatalogOptions{
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				ShortDescription: core.StringPtr("testString"),
				ShortDescriptionI18n: make(map[string]string),
				CatalogIconURL: core.StringPtr("testString"),
				CatalogBannerURL: core.StringPtr("testString"),
				Tags: []string{"testString"},
				Features: []catalogmanagementv1.Feature{*featureModel},
				Disabled: core.BoolPtr(true),
				ResourceGroupID: core.StringPtr("testString"),
				OwningAccount: core.StringPtr("testString"),
				CatalogFilters: filtersModel,
				SyndicationSettings: syndicationResourceModel,
				Kind: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
			}

			catalog, response, err := catalogManagementService.CreateCatalog(createCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalog).ToNot(BeNil())

			catalogIDLink = *catalog.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogIDLink value: %v\n", catalogIDLink)
			catalogRevLink = *catalog.Rev
			fmt.Fprintf(GinkgoWriter, "Saved catalogRevLink value: %v\n", catalogRevLink)
		})
	})

	Describe(`GetCatalog - Get catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {
			getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{
				CatalogIdentifier: &catalogIDLink,
			}

			catalog, response, err := catalogManagementService.GetCatalog(getCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())

			catalogRevLink = *catalog.Rev
			fmt.Fprintf(GinkgoWriter, "Saved catalogRevLink value: %v\n", catalogRevLink)
		})
	})

	Describe(`ReplaceCatalog - Update catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceCatalog(replaceCatalogOptions *ReplaceCatalogOptions)`, func() {
			featureModel := &catalogmanagementv1.Feature{
				Title: core.StringPtr("testString"),
				TitleI18n: make(map[string]string),
				Description: core.StringPtr("testString"),
				DescriptionI18n: make(map[string]string),
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
				LastRun: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			syndicationAuthorizationModel := &catalogmanagementv1.SyndicationAuthorization{
				Token: core.StringPtr("testString"),
				LastRun: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			syndicationResourceModel := &catalogmanagementv1.SyndicationResource{
				RemoveRelatedComponents: core.BoolPtr(true),
				Clusters: []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel},
				History: syndicationHistoryModel,
				Authorization: syndicationAuthorizationModel,
			}

			replaceCatalogOptions := &catalogmanagementv1.ReplaceCatalogOptions{
				CatalogIdentifier: &catalogIDLink,
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				ShortDescription: core.StringPtr("testString"),
				ShortDescriptionI18n: make(map[string]string),
				CatalogIconURL: core.StringPtr("testString"),
				CatalogBannerURL: core.StringPtr("testString"),
				Tags: []string{"testString"},
				Features: []catalogmanagementv1.Feature{*featureModel},
				Disabled: core.BoolPtr(true),
				ResourceGroupID: core.StringPtr("testString"),
				OwningAccount: core.StringPtr("testString"),
				CatalogFilters: filtersModel,
				SyndicationSettings: syndicationResourceModel,
				Kind: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
			}

			catalog, response, err := catalogManagementService.ReplaceCatalog(replaceCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())

			catalogRevLink = *catalog.Rev
			fmt.Fprintf(GinkgoWriter, "Saved catalogRevLink value: %v\n", catalogRevLink)
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
				TitleI18n: make(map[string]string),
				Description: core.StringPtr("testString"),
				DescriptionI18n: make(map[string]string),
			}

			flavorModel := &catalogmanagementv1.Flavor{
				Name: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Index: core.Int64Ptr(int64(38)),
			}

			renderTypeAssociationsParametersItemModel := &catalogmanagementv1.RenderTypeAssociationsParametersItem{
				Name: core.StringPtr("testString"),
				OptionsRefresh: core.BoolPtr(true),
			}

			renderTypeAssociationsModel := &catalogmanagementv1.RenderTypeAssociations{
				Parameters: []catalogmanagementv1.RenderTypeAssociationsParametersItem{*renderTypeAssociationsParametersItemModel},
			}

			renderTypeModel := &catalogmanagementv1.RenderType{
				Type: core.StringPtr("testString"),
				Grouping: core.StringPtr("testString"),
				OriginalGrouping: core.StringPtr("testString"),
				GroupingIndex: core.Int64Ptr(int64(38)),
				ConfigConstraints: map[string]interface{}{"anyKey": "anyValue"},
				Associations: renderTypeAssociationsModel,
			}

			configurationModel := &catalogmanagementv1.Configuration{
				Key: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				DisplayName: core.StringPtr("testString"),
				ValueConstraint: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Required: core.BoolPtr(true),
				Options: []interface{}{"testString"},
				Hidden: core.BoolPtr(true),
				CustomConfig: renderTypeModel,
				TypeMetadata: core.StringPtr("testString"),
			}

			outputModel := &catalogmanagementv1.Output{
				Key: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			iamResourceModel := &catalogmanagementv1.IamResource{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				RoleCrns: []string{"testString"},
			}

			iamPermissionModel := &catalogmanagementv1.IamPermission{
				ServiceName: core.StringPtr("testString"),
				RoleCrns: []string{"testString"},
				Resources: []catalogmanagementv1.IamResource{*iamResourceModel},
			}

			validationModel := &catalogmanagementv1.Validation{
				Validated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Requested: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				State: core.StringPtr("testString"),
				LastOperation: core.StringPtr("testString"),
				Target: map[string]interface{}{"anyKey": "anyValue"},
				Message: core.StringPtr("testString"),
			}

			resourceModel := &catalogmanagementv1.Resource{
				Type: core.StringPtr("mem"),
				Value: core.StringPtr("testString"),
			}

			scriptModel := &catalogmanagementv1.Script{
				Instructions: core.StringPtr("testString"),
				InstructionsI18n: make(map[string]string),
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
				CurrentEntered: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Pending: core.StringPtr("testString"),
				PendingRequested: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Previous: core.StringPtr("testString"),
			}

			deprecatePendingModel := &catalogmanagementv1.DeprecatePending{
				DeprecateDate: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				DeprecateState: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			urlProxyModel := &catalogmanagementv1.URLProxy{
				URL: core.StringPtr("testString"),
				Sha: core.StringPtr("testString"),
			}

			mediaItemModel := &catalogmanagementv1.MediaItem{
				URL: core.StringPtr("testString"),
				APIURL: core.StringPtr("testString"),
				URLProxy: urlProxyModel,
				Caption: core.StringPtr("testString"),
				CaptionI18n: make(map[string]string),
				Type: core.StringPtr("testString"),
				ThumbnailURL: core.StringPtr("testString"),
			}

			architectureDiagramModel := &catalogmanagementv1.ArchitectureDiagram{
				Diagram: mediaItemModel,
				Description: core.StringPtr("testString"),
				DescriptionI18n: make(map[string]string),
			}

			costComponentModel := &catalogmanagementv1.CostComponent{
				Name: core.StringPtr("testString"),
				Unit: core.StringPtr("testString"),
				HourlyQuantity: core.StringPtr("testString"),
				MonthlyQuantity: core.StringPtr("testString"),
				Price: core.StringPtr("testString"),
				HourlyCost: core.StringPtr("testString"),
				MonthlyCost: core.StringPtr("testString"),
			}

			costResourceModel := &catalogmanagementv1.CostResource{
				Name: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				HourlyCost: core.StringPtr("testString"),
				MonthlyCost: core.StringPtr("testString"),
				CostComponents: []catalogmanagementv1.CostComponent{*costComponentModel},
			}

			costBreakdownModel := &catalogmanagementv1.CostBreakdown{
				TotalHourlyCost: core.StringPtr("testString"),
				TotalMonthlyCost: core.StringPtr("testString"),
				Resources: []catalogmanagementv1.CostResource{*costResourceModel},
			}

			costSummaryModel := &catalogmanagementv1.CostSummary{
				TotalDetectedResources: core.Int64Ptr(int64(38)),
				TotalSupportedResources: core.Int64Ptr(int64(38)),
				TotalUnsupportedResources: core.Int64Ptr(int64(38)),
				TotalUsageBasedResources: core.Int64Ptr(int64(38)),
				TotalNoPriceResources: core.Int64Ptr(int64(38)),
				UnsupportedResourceCounts: make(map[string]int64),
				NoPriceResourceCounts: make(map[string]int64),
			}

			projectModel := &catalogmanagementv1.Project{
				Name: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				PastBreakdown: costBreakdownModel,
				Breakdown: costBreakdownModel,
				Diff: costBreakdownModel,
				Summary: costSummaryModel,
			}

			costEstimateModel := &catalogmanagementv1.CostEstimate{
				Version: core.StringPtr("testString"),
				Currency: core.StringPtr("testString"),
				Projects: []catalogmanagementv1.Project{*projectModel},
				Summary: costSummaryModel,
				TotalHourlyCost: core.StringPtr("testString"),
				TotalMonthlyCost: core.StringPtr("testString"),
				PastTotalHourlyCost: core.StringPtr("testString"),
				PastTotalMonthlyCost: core.StringPtr("testString"),
				DiffTotalHourlyCost: core.StringPtr("testString"),
				DiffTotalMonthlyCost: core.StringPtr("testString"),
				TimeGenerated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			offeringReferenceModel := &catalogmanagementv1.OfferingReference{
				CatalogID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Kind: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				Flavors: []string{"testString"},
			}

			solutionInfoModel := &catalogmanagementv1.SolutionInfo{
				ArchitectureDiagrams: []catalogmanagementv1.ArchitectureDiagram{*architectureDiagramModel},
				Features: []catalogmanagementv1.Feature{*featureModel},
				CostEstimate: costEstimateModel,
				Dependencies: []catalogmanagementv1.OfferingReference{*offeringReferenceModel},
			}

			complianceControlSccProfileModel := &catalogmanagementv1.ComplianceControlSccProfile{
				Type: core.StringPtr("testString"),
			}

			complianceControlFamilyModel := &catalogmanagementv1.ComplianceControlFamily{
				ID: core.StringPtr("testString"),
				ExternalID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				UIHref: core.StringPtr("testString"),
			}

			goalModel := &catalogmanagementv1.Goal{
				ID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				UIHref: core.StringPtr("testString"),
			}

			complianceControlValidationModel := &catalogmanagementv1.ComplianceControlValidation{
				Certified: core.BoolPtr(true),
				Results: map[string]interface{}{"anyKey": "anyValue"},
			}

			complianceControlModel := &catalogmanagementv1.ComplianceControl{
				SccProfile: complianceControlSccProfileModel,
				Family: complianceControlFamilyModel,
				Goals: []catalogmanagementv1.Goal{*goalModel},
				Validation: complianceControlValidationModel,
			}

			versionModel := &catalogmanagementv1.Version{
				CRN: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				Flavor: flavorModel,
				Sha: core.StringPtr("testString"),
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				OfferingID: &offeringIDLink,
				CatalogID: &catalogIDLink,
				KindID: core.StringPtr("testString"),
				Tags: []string{"testString"},
				RepoURL: core.StringPtr("testString"),
				SourceURL: core.StringPtr("testString"),
				TgzURL: core.StringPtr("testString"),
				Configuration: []catalogmanagementv1.Configuration{*configurationModel},
				Outputs: []catalogmanagementv1.Output{*outputModel},
				IamPermissions: []catalogmanagementv1.IamPermission{*iamPermissionModel},
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
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
				VersionLocator: &versionIDLink,
				LongDescription: core.StringPtr("testString"),
				LongDescriptionI18n: make(map[string]string),
				WhitelistedAccounts: []string{"testString"},
				ImagePullKeyName: core.StringPtr("testString"),
				DeprecatePending: deprecatePendingModel,
				SolutionInfo: solutionInfoModel,
				IsConsumable: core.BoolPtr(true),
				Compliance: []catalogmanagementv1.ComplianceControl{*complianceControlModel},
			}

			deploymentModel := &catalogmanagementv1.Deployment{
				ID: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ShortDescription: core.StringPtr("testString"),
				LongDescription: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				Tags: []string{"testString"},
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			planModel := &catalogmanagementv1.Plan{
				ID: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ShortDescription: core.StringPtr("testString"),
				LongDescription: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				Tags: []string{"testString"},
				AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Deployments: []catalogmanagementv1.Deployment{*deploymentModel},
			}

			kindModel := &catalogmanagementv1.Kind{
				ID: core.StringPtr("testString"),
				FormatKind: core.StringPtr("testString"),
				InstallKind: core.StringPtr("testString"),
				TargetKind: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				Tags: []string{"testString"},
				AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Versions: []catalogmanagementv1.Version{*versionModel},
				Plans: []catalogmanagementv1.Plan{*planModel},
			}

			providerInfoModel := &catalogmanagementv1.ProviderInfo{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			repoInfoModel := &catalogmanagementv1.RepoInfo{
				Token: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
			}

			imagePullKeyModel := &catalogmanagementv1.ImagePullKey{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			supportWaitTimeModel := &catalogmanagementv1.SupportWaitTime{
				Value: core.Int64Ptr(int64(38)),
				Type: core.StringPtr("testString"),
			}

			supportTimeModel := &catalogmanagementv1.SupportTime{
				Day: core.Int64Ptr(int64(38)),
				StartTime: core.StringPtr("testString"),
				EndTime: core.StringPtr("testString"),
			}

			supportAvailabilityModel := &catalogmanagementv1.SupportAvailability{
				Times: []catalogmanagementv1.SupportTime{*supportTimeModel},
				Timezone: core.StringPtr("testString"),
				AlwaysAvailable: core.BoolPtr(true),
			}

			supportDetailModel := &catalogmanagementv1.SupportDetail{
				Type: core.StringPtr("testString"),
				Contact: core.StringPtr("testString"),
				ResponseWaitTime: supportWaitTimeModel,
				Availability: supportAvailabilityModel,
			}

			supportEscalationModel := &catalogmanagementv1.SupportEscalation{
				EscalationWaitTime: supportWaitTimeModel,
				ResponseWaitTime: supportWaitTimeModel,
				Contact: core.StringPtr("testString"),
			}

			supportModel := &catalogmanagementv1.Support{
				URL: core.StringPtr("testString"),
				Process: core.StringPtr("testString"),
				ProcessI18n: make(map[string]string),
				Locations: []string{"testString"},
				SupportDetails: []catalogmanagementv1.SupportDetail{*supportDetailModel},
				SupportEscalation: supportEscalationModel,
				SupportType: core.StringPtr("testString"),
			}

			learnMoreLinksModel := &catalogmanagementv1.LearnMoreLinks{
				FirstParty: core.StringPtr("testString"),
				ThirdParty: core.StringPtr("testString"),
			}

			constraintModel := &catalogmanagementv1.Constraint{
				Type: core.StringPtr("testString"),
				Rule: core.StringPtr("testString"),
			}

			badgeModel := &catalogmanagementv1.Badge{
				ID: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Description: core.StringPtr("testString"),
				DescriptionI18n: make(map[string]string),
				Icon: core.StringPtr("testString"),
				Authority: core.StringPtr("testString"),
				Tag: core.StringPtr("testString"),
				LearnMoreLinks: learnMoreLinksModel,
				Constraints: []catalogmanagementv1.Constraint{*constraintModel},
			}

			replaceOfferingOptions := &catalogmanagementv1.ReplaceOfferingOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				URL: core.StringPtr("testString"),
				CRN: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Name: core.StringPtr("testString"),
				OfferingIconURL: core.StringPtr("testString"),
				OfferingDocsURL: core.StringPtr("testString"),
				OfferingSupportURL: core.StringPtr("testString"),
				Tags: []string{"testString"},
				Keywords: []string{"testString"},
				Rating: ratingModel,
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				ShortDescription: core.StringPtr("testString"),
				ShortDescriptionI18n: make(map[string]string),
				LongDescription: core.StringPtr("testString"),
				LongDescriptionI18n: make(map[string]string),
				Features: []catalogmanagementv1.Feature{*featureModel},
				Kinds: []catalogmanagementv1.Kind{*kindModel},
				PcManaged: core.BoolPtr(true),
				PublishApproved: core.BoolPtr(true),
				ShareWithAll: core.BoolPtr(true),
				ShareWithIBM: core.BoolPtr(true),
				ShareEnabled: core.BoolPtr(true),
				PermitRequestIBMPublicPublish: core.BoolPtr(true),
				IBMPublishApproved: core.BoolPtr(true),
				PublicPublishApproved: core.BoolPtr(true),
				PublicOriginalCRN: core.StringPtr("testString"),
				PublishPublicCRN: core.StringPtr("testString"),
				PortalApprovalRecord: core.StringPtr("testString"),
				PortalUIURL: core.StringPtr("testString"),
				CatalogID: &catalogIDLink,
				CatalogName: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				Disclaimer: core.StringPtr("testString"),
				Hidden: core.BoolPtr(true),
				Provider: core.StringPtr("testString"),
				ProviderInfo: providerInfoModel,
				RepoInfo: repoInfoModel,
				ImagePullKeys: []catalogmanagementv1.ImagePullKey{*imagePullKeyModel},
				Support: supportModel,
				Media: []catalogmanagementv1.MediaItem{*mediaItemModel},
				DeprecatePending: deprecatePendingModel,
				ProductKind: core.StringPtr("testString"),
				Badges: []catalogmanagementv1.Badge{*badgeModel},
			}

			offering, response, err := catalogManagementService.ReplaceOffering(replaceOfferingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

			offeringIDLink = *offering.ID
			fmt.Fprintf(GinkgoWriter, "Saved offeringIDLink value: %v\n", offeringIDLink)
			offeringRevLink = *offering.Rev
			fmt.Fprintf(GinkgoWriter, "Saved offeringRevLink value: %v\n", offeringRevLink)
			versionLocatorLink = *offering.Kinds[0].Versions[0].VersionLocator
			fmt.Fprintf(GinkgoWriter, "Saved versionLocatorLink value: %v\n", versionLocatorLink)
			versionIDLink = *offering.Kinds[0].Versions[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved versionIDLink value: %v\n", versionIDLink)
			versionRevLink = *offering.Kinds[0].Versions[0].Rev
			fmt.Fprintf(GinkgoWriter, "Saved versionRevLink value: %v\n", versionRevLink)
		})
	})

	Describe(`ImportOfferingVersion - Import offering version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ImportOfferingVersion(importOfferingVersionOptions *ImportOfferingVersionOptions)`, func() {
			flavorModel := &catalogmanagementv1.Flavor{
				Name: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Index: core.Int64Ptr(int64(38)),
			}

			importOfferingBodyMetadataOperatingSystemModel := &catalogmanagementv1.ImportOfferingBodyMetadataOperatingSystem{
				DedicatedHostOnly: core.BoolPtr(true),
				Vendor: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Href: core.StringPtr("testString"),
				DisplayName: core.StringPtr("testString"),
				Family: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				Architecture: core.StringPtr("testString"),
			}

			importOfferingBodyMetadataFileModel := &catalogmanagementv1.ImportOfferingBodyMetadataFile{
				Size: core.Int64Ptr(int64(38)),
			}

			importOfferingBodyMetadataImagesItemModel := &catalogmanagementv1.ImportOfferingBodyMetadataImagesItem{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			importOfferingBodyMetadataModel := &catalogmanagementv1.ImportOfferingBodyMetadata{
				OperatingSystem: importOfferingBodyMetadataOperatingSystemModel,
				File: importOfferingBodyMetadataFileModel,
				MinimumProvisionedSize: core.Int64Ptr(int64(38)),
				Images: []catalogmanagementv1.ImportOfferingBodyMetadataImagesItem{*importOfferingBodyMetadataImagesItemModel},
			}

			importOfferingVersionOptions := &catalogmanagementv1.ImportOfferingVersionOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				Tags: []string{"testString"},
				Content: CreateMockByteArray("This is a mock byte array value."),
				Name: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				InstallKind: core.StringPtr("testString"),
				TargetKinds: []string{"testString"},
				FormatKind: core.StringPtr("testString"),
				ProductKind: core.StringPtr("testString"),
				Sha: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				Flavor: flavorModel,
				Metadata: importOfferingBodyMetadataModel,
				WorkingDirectory: core.StringPtr("testString"),
				Zipurl: core.StringPtr("testString"),
				TargetVersion: core.StringPtr("testString"),
				IncludeConfig: core.BoolPtr(true),
				IsVsi: core.BoolPtr(true),
				Repotype: core.StringPtr("testString"),
				XAuthToken: core.StringPtr("testString"),
			}

			offering, response, err := catalogManagementService.ImportOfferingVersion(importOfferingVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())

			offeringIDLink = *offering.ID
			fmt.Fprintf(GinkgoWriter, "Saved offeringIDLink value: %v\n", offeringIDLink)
			offeringRevLink = *offering.Rev
			fmt.Fprintf(GinkgoWriter, "Saved offeringRevLink value: %v\n", offeringRevLink)
			versionLocatorLink = *offering.Kinds[0].Versions[0].VersionLocator
			fmt.Fprintf(GinkgoWriter, "Saved versionLocatorLink value: %v\n", versionLocatorLink)
			versionIDLink = *offering.Kinds[0].Versions[0].VersionLocator
			fmt.Fprintf(GinkgoWriter, "Saved versionIDLink value: %v\n", versionIDLink)
			versionRevLink = *offering.Kinds[0].Versions[0].Rev
			fmt.Fprintf(GinkgoWriter, "Saved versionRevLink value: %v\n", versionRevLink)
		})
	})

	Describe(`ImportOffering - Import offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ImportOffering(importOfferingOptions *ImportOfferingOptions)`, func() {
			flavorModel := &catalogmanagementv1.Flavor{
				Name: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Index: core.Int64Ptr(int64(38)),
			}

			importOfferingBodyMetadataOperatingSystemModel := &catalogmanagementv1.ImportOfferingBodyMetadataOperatingSystem{
				DedicatedHostOnly: core.BoolPtr(true),
				Vendor: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Href: core.StringPtr("testString"),
				DisplayName: core.StringPtr("testString"),
				Family: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				Architecture: core.StringPtr("testString"),
			}

			importOfferingBodyMetadataFileModel := &catalogmanagementv1.ImportOfferingBodyMetadataFile{
				Size: core.Int64Ptr(int64(38)),
			}

			importOfferingBodyMetadataImagesItemModel := &catalogmanagementv1.ImportOfferingBodyMetadataImagesItem{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			importOfferingBodyMetadataModel := &catalogmanagementv1.ImportOfferingBodyMetadata{
				OperatingSystem: importOfferingBodyMetadataOperatingSystemModel,
				File: importOfferingBodyMetadataFileModel,
				MinimumProvisionedSize: core.Int64Ptr(int64(38)),
				Images: []catalogmanagementv1.ImportOfferingBodyMetadataImagesItem{*importOfferingBodyMetadataImagesItemModel},
			}

			importOfferingOptions := &catalogmanagementv1.ImportOfferingOptions{
				CatalogIdentifier: &catalogIDLink,
				Tags: []string{"testString"},
				Content: CreateMockByteArray("This is a mock byte array value."),
				Name: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				InstallKind: core.StringPtr("testString"),
				TargetKinds: []string{"testString"},
				FormatKind: core.StringPtr("testString"),
				ProductKind: core.StringPtr("testString"),
				Sha: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				Flavor: flavorModel,
				Metadata: importOfferingBodyMetadataModel,
				WorkingDirectory: core.StringPtr("testString"),
				Zipurl: core.StringPtr("testString"),
				OfferingID: &offeringIDLink,
				TargetVersion: core.StringPtr("testString"),
				IncludeConfig: core.BoolPtr(true),
				IsVsi: core.BoolPtr(true),
				Repotype: core.StringPtr("testString"),
				XAuthToken: core.StringPtr("testString"),
			}

			offering, response, err := catalogManagementService.ImportOffering(importOfferingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())

			offeringRevLink = *offering.Rev
			fmt.Fprintf(GinkgoWriter, "Saved offeringRevLink value: %v\n", offeringRevLink)
			versionLocatorLink = *offering.Kinds[0].Versions[0].VersionLocator
			fmt.Fprintf(GinkgoWriter, "Saved versionLocatorLink value: %v\n", versionLocatorLink)
		})
	})

	Describe(`GetOffering - Get offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOffering(getOfferingOptions *GetOfferingOptions)`, func() {
			getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				Type: core.StringPtr("testString"),
				Digest: core.BoolPtr(true),
			}

			offering, response, err := catalogManagementService.GetOffering(getOfferingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

			offeringRevLink = *offering.Rev
			fmt.Fprintf(GinkgoWriter, "Saved offeringRevLink value: %v\n", offeringRevLink)
			versionLocatorLink = *offering.Kinds[0].Versions[0].VersionLocator
			fmt.Fprintf(GinkgoWriter, "Saved versionLocatorLink value: %v\n", versionLocatorLink)
			versionIDLink = *offering.Kinds[0].Versions[0].VersionLocator
			fmt.Fprintf(GinkgoWriter, "Saved versionIDLink value: %v\n", versionIDLink)
			versionRevLink = *offering.Kinds[0].Versions[0].Rev
			fmt.Fprintf(GinkgoWriter, "Saved versionRevLink value: %v\n", versionRevLink)
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
				TitleI18n: make(map[string]string),
				Description: core.StringPtr("testString"),
				DescriptionI18n: make(map[string]string),
			}

			flavorModel := &catalogmanagementv1.Flavor{
				Name: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Index: core.Int64Ptr(int64(38)),
			}

			renderTypeAssociationsParametersItemModel := &catalogmanagementv1.RenderTypeAssociationsParametersItem{
				Name: core.StringPtr("testString"),
				OptionsRefresh: core.BoolPtr(true),
			}

			renderTypeAssociationsModel := &catalogmanagementv1.RenderTypeAssociations{
				Parameters: []catalogmanagementv1.RenderTypeAssociationsParametersItem{*renderTypeAssociationsParametersItemModel},
			}

			renderTypeModel := &catalogmanagementv1.RenderType{
				Type: core.StringPtr("testString"),
				Grouping: core.StringPtr("testString"),
				OriginalGrouping: core.StringPtr("testString"),
				GroupingIndex: core.Int64Ptr(int64(38)),
				ConfigConstraints: map[string]interface{}{"anyKey": "anyValue"},
				Associations: renderTypeAssociationsModel,
			}

			configurationModel := &catalogmanagementv1.Configuration{
				Key: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				DisplayName: core.StringPtr("testString"),
				ValueConstraint: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Required: core.BoolPtr(true),
				Options: []interface{}{"testString"},
				Hidden: core.BoolPtr(true),
				CustomConfig: renderTypeModel,
				TypeMetadata: core.StringPtr("testString"),
			}

			outputModel := &catalogmanagementv1.Output{
				Key: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			iamResourceModel := &catalogmanagementv1.IamResource{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				RoleCrns: []string{"testString"},
			}

			iamPermissionModel := &catalogmanagementv1.IamPermission{
				ServiceName: core.StringPtr("testString"),
				RoleCrns: []string{"testString"},
				Resources: []catalogmanagementv1.IamResource{*iamResourceModel},
			}

			validationModel := &catalogmanagementv1.Validation{
				Validated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Requested: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				State: core.StringPtr("testString"),
				LastOperation: core.StringPtr("testString"),
				Target: map[string]interface{}{"anyKey": "anyValue"},
				Message: core.StringPtr("testString"),
			}

			resourceModel := &catalogmanagementv1.Resource{
				Type: core.StringPtr("mem"),
				Value: core.StringPtr("testString"),
			}

			scriptModel := &catalogmanagementv1.Script{
				Instructions: core.StringPtr("testString"),
				InstructionsI18n: make(map[string]string),
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
				CurrentEntered: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Pending: core.StringPtr("testString"),
				PendingRequested: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Previous: core.StringPtr("testString"),
			}

			deprecatePendingModel := &catalogmanagementv1.DeprecatePending{
				DeprecateDate: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				DeprecateState: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			urlProxyModel := &catalogmanagementv1.URLProxy{
				URL: core.StringPtr("testString"),
				Sha: core.StringPtr("testString"),
			}

			mediaItemModel := &catalogmanagementv1.MediaItem{
				URL: core.StringPtr("testString"),
				APIURL: core.StringPtr("testString"),
				URLProxy: urlProxyModel,
				Caption: core.StringPtr("testString"),
				CaptionI18n: make(map[string]string),
				Type: core.StringPtr("testString"),
				ThumbnailURL: core.StringPtr("testString"),
			}

			architectureDiagramModel := &catalogmanagementv1.ArchitectureDiagram{
				Diagram: mediaItemModel,
				Description: core.StringPtr("testString"),
				DescriptionI18n: make(map[string]string),
			}

			costComponentModel := &catalogmanagementv1.CostComponent{
				Name: core.StringPtr("testString"),
				Unit: core.StringPtr("testString"),
				HourlyQuantity: core.StringPtr("testString"),
				MonthlyQuantity: core.StringPtr("testString"),
				Price: core.StringPtr("testString"),
				HourlyCost: core.StringPtr("testString"),
				MonthlyCost: core.StringPtr("testString"),
			}

			costResourceModel := &catalogmanagementv1.CostResource{
				Name: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				HourlyCost: core.StringPtr("testString"),
				MonthlyCost: core.StringPtr("testString"),
				CostComponents: []catalogmanagementv1.CostComponent{*costComponentModel},
			}

			costBreakdownModel := &catalogmanagementv1.CostBreakdown{
				TotalHourlyCost: core.StringPtr("testString"),
				TotalMonthlyCost: core.StringPtr("testString"),
				Resources: []catalogmanagementv1.CostResource{*costResourceModel},
			}

			costSummaryModel := &catalogmanagementv1.CostSummary{
				TotalDetectedResources: core.Int64Ptr(int64(38)),
				TotalSupportedResources: core.Int64Ptr(int64(38)),
				TotalUnsupportedResources: core.Int64Ptr(int64(38)),
				TotalUsageBasedResources: core.Int64Ptr(int64(38)),
				TotalNoPriceResources: core.Int64Ptr(int64(38)),
				UnsupportedResourceCounts: make(map[string]int64),
				NoPriceResourceCounts: make(map[string]int64),
			}

			projectModel := &catalogmanagementv1.Project{
				Name: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				PastBreakdown: costBreakdownModel,
				Breakdown: costBreakdownModel,
				Diff: costBreakdownModel,
				Summary: costSummaryModel,
			}

			costEstimateModel := &catalogmanagementv1.CostEstimate{
				Version: core.StringPtr("testString"),
				Currency: core.StringPtr("testString"),
				Projects: []catalogmanagementv1.Project{*projectModel},
				Summary: costSummaryModel,
				TotalHourlyCost: core.StringPtr("testString"),
				TotalMonthlyCost: core.StringPtr("testString"),
				PastTotalHourlyCost: core.StringPtr("testString"),
				PastTotalMonthlyCost: core.StringPtr("testString"),
				DiffTotalHourlyCost: core.StringPtr("testString"),
				DiffTotalMonthlyCost: core.StringPtr("testString"),
				TimeGenerated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			offeringReferenceModel := &catalogmanagementv1.OfferingReference{
				CatalogID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Kind: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				Flavors: []string{"testString"},
			}

			solutionInfoModel := &catalogmanagementv1.SolutionInfo{
				ArchitectureDiagrams: []catalogmanagementv1.ArchitectureDiagram{*architectureDiagramModel},
				Features: []catalogmanagementv1.Feature{*featureModel},
				CostEstimate: costEstimateModel,
				Dependencies: []catalogmanagementv1.OfferingReference{*offeringReferenceModel},
			}

			complianceControlSccProfileModel := &catalogmanagementv1.ComplianceControlSccProfile{
				Type: core.StringPtr("testString"),
			}

			complianceControlFamilyModel := &catalogmanagementv1.ComplianceControlFamily{
				ID: core.StringPtr("testString"),
				ExternalID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				UIHref: core.StringPtr("testString"),
			}

			goalModel := &catalogmanagementv1.Goal{
				ID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				UIHref: core.StringPtr("testString"),
			}

			complianceControlValidationModel := &catalogmanagementv1.ComplianceControlValidation{
				Certified: core.BoolPtr(true),
				Results: map[string]interface{}{"anyKey": "anyValue"},
			}

			complianceControlModel := &catalogmanagementv1.ComplianceControl{
				SccProfile: complianceControlSccProfileModel,
				Family: complianceControlFamilyModel,
				Goals: []catalogmanagementv1.Goal{*goalModel},
				Validation: complianceControlValidationModel,
			}

			versionModel := &catalogmanagementv1.Version{
				CRN: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				Flavor: flavorModel,
				Sha: core.StringPtr("testString"),
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				OfferingID: &offeringIDLink,
				CatalogID: &catalogIDLink,
				KindID: core.StringPtr("testString"),
				Tags: []string{"testString"},
				RepoURL: core.StringPtr("testString"),
				SourceURL: core.StringPtr("testString"),
				TgzURL: core.StringPtr("testString"),
				Configuration: []catalogmanagementv1.Configuration{*configurationModel},
				Outputs: []catalogmanagementv1.Output{*outputModel},
				IamPermissions: []catalogmanagementv1.IamPermission{*iamPermissionModel},
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
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
				VersionLocator: &versionIDLink,
				LongDescription: core.StringPtr("testString"),
				LongDescriptionI18n: make(map[string]string),
				WhitelistedAccounts: []string{"testString"},
				ImagePullKeyName: core.StringPtr("testString"),
				DeprecatePending: deprecatePendingModel,
				SolutionInfo: solutionInfoModel,
				IsConsumable: core.BoolPtr(true),
				Compliance: []catalogmanagementv1.ComplianceControl{*complianceControlModel},
			}

			deploymentModel := &catalogmanagementv1.Deployment{
				ID: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ShortDescription: core.StringPtr("testString"),
				LongDescription: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				Tags: []string{"testString"},
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			planModel := &catalogmanagementv1.Plan{
				ID: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ShortDescription: core.StringPtr("testString"),
				LongDescription: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				Tags: []string{"testString"},
				AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Deployments: []catalogmanagementv1.Deployment{*deploymentModel},
			}

			kindModel := &catalogmanagementv1.Kind{
				ID: core.StringPtr("testString"),
				FormatKind: core.StringPtr("testString"),
				InstallKind: core.StringPtr("testString"),
				TargetKind: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				Tags: []string{"testString"},
				AdditionalFeatures: []catalogmanagementv1.Feature{*featureModel},
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Versions: []catalogmanagementv1.Version{*versionModel},
				Plans: []catalogmanagementv1.Plan{*planModel},
			}

			providerInfoModel := &catalogmanagementv1.ProviderInfo{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			repoInfoModel := &catalogmanagementv1.RepoInfo{
				Token: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
			}

			imagePullKeyModel := &catalogmanagementv1.ImagePullKey{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			supportWaitTimeModel := &catalogmanagementv1.SupportWaitTime{
				Value: core.Int64Ptr(int64(38)),
				Type: core.StringPtr("testString"),
			}

			supportTimeModel := &catalogmanagementv1.SupportTime{
				Day: core.Int64Ptr(int64(38)),
				StartTime: core.StringPtr("testString"),
				EndTime: core.StringPtr("testString"),
			}

			supportAvailabilityModel := &catalogmanagementv1.SupportAvailability{
				Times: []catalogmanagementv1.SupportTime{*supportTimeModel},
				Timezone: core.StringPtr("testString"),
				AlwaysAvailable: core.BoolPtr(true),
			}

			supportDetailModel := &catalogmanagementv1.SupportDetail{
				Type: core.StringPtr("testString"),
				Contact: core.StringPtr("testString"),
				ResponseWaitTime: supportWaitTimeModel,
				Availability: supportAvailabilityModel,
			}

			supportEscalationModel := &catalogmanagementv1.SupportEscalation{
				EscalationWaitTime: supportWaitTimeModel,
				ResponseWaitTime: supportWaitTimeModel,
				Contact: core.StringPtr("testString"),
			}

			supportModel := &catalogmanagementv1.Support{
				URL: core.StringPtr("testString"),
				Process: core.StringPtr("testString"),
				ProcessI18n: make(map[string]string),
				Locations: []string{"testString"},
				SupportDetails: []catalogmanagementv1.SupportDetail{*supportDetailModel},
				SupportEscalation: supportEscalationModel,
				SupportType: core.StringPtr("testString"),
			}

			learnMoreLinksModel := &catalogmanagementv1.LearnMoreLinks{
				FirstParty: core.StringPtr("testString"),
				ThirdParty: core.StringPtr("testString"),
			}

			constraintModel := &catalogmanagementv1.Constraint{
				Type: core.StringPtr("testString"),
				Rule: core.StringPtr("testString"),
			}

			badgeModel := &catalogmanagementv1.Badge{
				ID: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Description: core.StringPtr("testString"),
				DescriptionI18n: make(map[string]string),
				Icon: core.StringPtr("testString"),
				Authority: core.StringPtr("testString"),
				Tag: core.StringPtr("testString"),
				LearnMoreLinks: learnMoreLinksModel,
				Constraints: []catalogmanagementv1.Constraint{*constraintModel},
			}

			createOfferingOptions := &catalogmanagementv1.CreateOfferingOptions{
				CatalogIdentifier: &catalogIDLink,
				URL: core.StringPtr("testString"),
				CRN: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Name: core.StringPtr("testString"),
				OfferingIconURL: core.StringPtr("testString"),
				OfferingDocsURL: core.StringPtr("testString"),
				OfferingSupportURL: core.StringPtr("testString"),
				Tags: []string{"testString"},
				Keywords: []string{"testString"},
				Rating: ratingModel,
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				ShortDescription: core.StringPtr("testString"),
				ShortDescriptionI18n: make(map[string]string),
				LongDescription: core.StringPtr("testString"),
				LongDescriptionI18n: make(map[string]string),
				Features: []catalogmanagementv1.Feature{*featureModel},
				Kinds: []catalogmanagementv1.Kind{*kindModel},
				PcManaged: core.BoolPtr(true),
				PublishApproved: core.BoolPtr(true),
				ShareWithAll: core.BoolPtr(true),
				ShareWithIBM: core.BoolPtr(true),
				ShareEnabled: core.BoolPtr(true),
				PermitRequestIBMPublicPublish: core.BoolPtr(true),
				IBMPublishApproved: core.BoolPtr(true),
				PublicPublishApproved: core.BoolPtr(true),
				PublicOriginalCRN: core.StringPtr("testString"),
				PublishPublicCRN: core.StringPtr("testString"),
				PortalApprovalRecord: core.StringPtr("testString"),
				PortalUIURL: core.StringPtr("testString"),
				CatalogID: &catalogIDLink,
				CatalogName: core.StringPtr("testString"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				Disclaimer: core.StringPtr("testString"),
				Hidden: core.BoolPtr(true),
				Provider: core.StringPtr("testString"),
				ProviderInfo: providerInfoModel,
				RepoInfo: repoInfoModel,
				ImagePullKeys: []catalogmanagementv1.ImagePullKey{*imagePullKeyModel},
				Support: supportModel,
				Media: []catalogmanagementv1.MediaItem{*mediaItemModel},
				DeprecatePending: deprecatePendingModel,
				ProductKind: core.StringPtr("testString"),
				Badges: []catalogmanagementv1.Badge{*badgeModel},
			}

			offering, response, err := catalogManagementService.CreateOffering(createOfferingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(offering).ToNot(BeNil())

			offeringIDLink = *offering.ID
			fmt.Fprintf(GinkgoWriter, "Saved offeringIDLink value: %v\n", offeringIDLink)
			offeringRevLink = *offering.Rev
			fmt.Fprintf(GinkgoWriter, "Saved offeringRevLink value: %v\n", offeringRevLink)
			versionLocatorLink = *offering.Kinds[0].Versions[0].VersionLocator
			fmt.Fprintf(GinkgoWriter, "Saved versionLocatorLink value: %v\n", versionLocatorLink)
			versionIDLink = *offering.Kinds[0].Versions[0].VersionLocator
			fmt.Fprintf(GinkgoWriter, "Saved versionIDLink value: %v\n", versionIDLink)
			versionRevLink = *offering.Kinds[0].Versions[0].Rev
			fmt.Fprintf(GinkgoWriter, "Saved versionRevLink value: %v\n", versionRevLink)
		})
	})

	Describe(`ReloadOffering - Reload offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReloadOffering(reloadOfferingOptions *ReloadOfferingOptions)`, func() {
			flavorModel := &catalogmanagementv1.Flavor{
				Name: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Index: core.Int64Ptr(int64(38)),
			}

			reloadOfferingOptions := &catalogmanagementv1.ReloadOfferingOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				TargetVersion: core.StringPtr("testString"),
				Tags: []string{"testString"},
				Content: CreateMockByteArray("This is a mock byte array value."),
				TargetKinds: []string{"testString"},
				FormatKind: core.StringPtr("testString"),
				Flavor: flavorModel,
				WorkingDirectory: core.StringPtr("testString"),
				Zipurl: core.StringPtr("testString"),
				RepoType: core.StringPtr("testString"),
			}

			offering, response, err := catalogManagementService.ReloadOffering(reloadOfferingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

			offeringRevLink = *offering.Rev
			fmt.Fprintf(GinkgoWriter, "Saved offeringRevLink value: %v\n", offeringRevLink)
		})
	})

	Describe(`UpdateOffering - Update offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateOffering(updateOfferingOptions *UpdateOfferingOptions)`, func() {
			jsonPatchOperationModel := &catalogmanagementv1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
			}

			updateOfferingOptions := &catalogmanagementv1.UpdateOfferingOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				IfMatch: core.StringPtr("testString"),
				Updates: []catalogmanagementv1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			offering, response, err := catalogManagementService.UpdateOffering(updateOfferingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offering).ToNot(BeNil())

			offeringRevLink = *offering.Rev
			fmt.Fprintf(GinkgoWriter, "Saved offeringRevLink value: %v\n", offeringRevLink)
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
				CurrentEntered: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Pending: core.StringPtr("testString"),
				PendingRequested: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Previous: core.StringPtr("testString"),
			}

			createObjectOptions := &catalogmanagementv1.CreateObjectOptions{
				CatalogIdentifier: &catalogIDLink,
				Name: core.StringPtr("testString"),
				CRN: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
				ParentID: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Label: core.StringPtr("testString"),
				Tags: []string{"testString"},
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				ShortDescription: core.StringPtr("testString"),
				ShortDescriptionI18n: make(map[string]string),
				Kind: core.StringPtr("testString"),
				Publish: publishObjectModel,
				State: stateModel,
				CatalogID: &catalogIDLink,
				CatalogName: core.StringPtr("testString"),
				Data: map[string]interface{}{"anyKey": "anyValue"},
			}

			catalogObject, response, err := catalogManagementService.CreateObject(createObjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogObject).ToNot(BeNil())

			objectIDLink = *catalogObject.ID
			fmt.Fprintf(GinkgoWriter, "Saved objectIDLink value: %v\n", objectIDLink)
			objectRevLink = *catalogObject.Rev
			fmt.Fprintf(GinkgoWriter, "Saved objectRevLink value: %v\n", objectRevLink)
		})
	})

	Describe(`GetObject - Get catalog object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetObject(getObjectOptions *GetObjectOptions)`, func() {
			getObjectOptions := &catalogmanagementv1.GetObjectOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
			}

			catalogObject, response, err := catalogManagementService.GetObject(getObjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogObject).ToNot(BeNil())

			objectRevLink = *catalogObject.Rev
			fmt.Fprintf(GinkgoWriter, "Saved objectRevLink value: %v\n", objectRevLink)
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
				CurrentEntered: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Pending: core.StringPtr("testString"),
				PendingRequested: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Previous: core.StringPtr("testString"),
			}

			replaceObjectOptions := &catalogmanagementv1.ReplaceObjectOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				Name: core.StringPtr("testString"),
				CRN: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
				ParentID: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Label: core.StringPtr("testString"),
				Tags: []string{"testString"},
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				ShortDescription: core.StringPtr("testString"),
				ShortDescriptionI18n: make(map[string]string),
				Kind: core.StringPtr("testString"),
				Publish: publishObjectModel,
				State: stateModel,
				CatalogID: &catalogIDLink,
				CatalogName: core.StringPtr("testString"),
				Data: map[string]interface{}{"anyKey": "anyValue"},
			}

			catalogObject, response, err := catalogManagementService.ReplaceObject(replaceObjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogObject).ToNot(BeNil())

			objectRevLink = *catalogObject.Rev
			fmt.Fprintf(GinkgoWriter, "Saved objectRevLink value: %v\n", objectRevLink)
		})
	})

	Describe(`ListCatalogAccountAudits - Get catalog account audit logs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCatalogAccountAudits(listCatalogAccountAuditsOptions *ListCatalogAccountAuditsOptions) with pagination`, func(){
			listCatalogAccountAuditsOptions := &catalogmanagementv1.ListCatalogAccountAuditsOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			listCatalogAccountAuditsOptions.Start = nil
			listCatalogAccountAuditsOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.AuditLogDigest
			for {
				auditLogs, response, err := catalogManagementService.ListCatalogAccountAudits(listCatalogAccountAuditsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLogs).ToNot(BeNil())
				allResults = append(allResults, auditLogs.Audits...)

				listCatalogAccountAuditsOptions.Start, err = auditLogs.GetNextStart()
				Expect(err).To(BeNil())

				if listCatalogAccountAuditsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListCatalogAccountAudits(listCatalogAccountAuditsOptions *ListCatalogAccountAuditsOptions) using CatalogAccountAuditsPager`, func(){
			listCatalogAccountAuditsOptions := &catalogmanagementv1.ListCatalogAccountAuditsOptions{
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewCatalogAccountAuditsPager(listCatalogAccountAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.AuditLogDigest
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewCatalogAccountAuditsPager(listCatalogAccountAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListCatalogAccountAudits() returned a total of %d item(s) using CatalogAccountAuditsPager.\n", len(allResults))
		})
	})

	Describe(`GetCatalogAccountAudit - Get a catalog account audit log entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalogAccountAudit(getCatalogAccountAuditOptions *GetCatalogAccountAuditOptions)`, func() {
			getCatalogAccountAuditOptions := &catalogmanagementv1.GetCatalogAccountAuditOptions{
				AuditlogIdentifier: core.StringPtr("testString"),
				Lookupnames: core.BoolPtr(true),
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
				Catalog: &catalogIDLink,
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

	Describe(`ListCatalogAudits - Get catalog audit logs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCatalogAudits(listCatalogAuditsOptions *ListCatalogAuditsOptions) with pagination`, func(){
			listCatalogAuditsOptions := &catalogmanagementv1.ListCatalogAuditsOptions{
				CatalogIdentifier: &catalogIDLink,
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			listCatalogAuditsOptions.Start = nil
			listCatalogAuditsOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.AuditLogDigest
			for {
				auditLogs, response, err := catalogManagementService.ListCatalogAudits(listCatalogAuditsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLogs).ToNot(BeNil())
				allResults = append(allResults, auditLogs.Audits...)

				listCatalogAuditsOptions.Start, err = auditLogs.GetNextStart()
				Expect(err).To(BeNil())

				if listCatalogAuditsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListCatalogAudits(listCatalogAuditsOptions *ListCatalogAuditsOptions) using CatalogAuditsPager`, func(){
			listCatalogAuditsOptions := &catalogmanagementv1.ListCatalogAuditsOptions{
				CatalogIdentifier: &catalogIDLink,
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewCatalogAuditsPager(listCatalogAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.AuditLogDigest
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewCatalogAuditsPager(listCatalogAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListCatalogAudits() returned a total of %d item(s) using CatalogAuditsPager.\n", len(allResults))
		})
	})

	Describe(`GetCatalogAudit - Get a catalog audit log entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalogAudit(getCatalogAuditOptions *GetCatalogAuditOptions)`, func() {
			getCatalogAuditOptions := &catalogmanagementv1.GetCatalogAuditOptions{
				CatalogIdentifier: &catalogIDLink,
				AuditlogIdentifier: core.StringPtr("testString"),
				Lookupnames: core.BoolPtr(true),
			}

			auditLog, response, err := catalogManagementService.GetCatalogAudit(getCatalogAuditOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditLog).ToNot(BeNil())
		})
	})

	Describe(`ListEnterpriseAudits - Get enterprise audit logs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListEnterpriseAudits(listEnterpriseAuditsOptions *ListEnterpriseAuditsOptions) with pagination`, func(){
			listEnterpriseAuditsOptions := &catalogmanagementv1.ListEnterpriseAuditsOptions{
				EnterpriseIdentifier: core.StringPtr("testString"),
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			listEnterpriseAuditsOptions.Start = nil
			listEnterpriseAuditsOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.AuditLogDigest
			for {
				auditLogs, response, err := catalogManagementService.ListEnterpriseAudits(listEnterpriseAuditsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLogs).ToNot(BeNil())
				allResults = append(allResults, auditLogs.Audits...)

				listEnterpriseAuditsOptions.Start, err = auditLogs.GetNextStart()
				Expect(err).To(BeNil())

				if listEnterpriseAuditsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListEnterpriseAudits(listEnterpriseAuditsOptions *ListEnterpriseAuditsOptions) using EnterpriseAuditsPager`, func(){
			listEnterpriseAuditsOptions := &catalogmanagementv1.ListEnterpriseAuditsOptions{
				EnterpriseIdentifier: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewEnterpriseAuditsPager(listEnterpriseAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.AuditLogDigest
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewEnterpriseAuditsPager(listEnterpriseAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListEnterpriseAudits() returned a total of %d item(s) using EnterpriseAuditsPager.\n", len(allResults))
		})
	})

	Describe(`GetEnterpriseAudit - Get an enterprise audit log entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEnterpriseAudit(getEnterpriseAuditOptions *GetEnterpriseAuditOptions)`, func() {
			getEnterpriseAuditOptions := &catalogmanagementv1.GetEnterpriseAuditOptions{
				EnterpriseIdentifier: core.StringPtr("testString"),
				AuditlogIdentifier: core.StringPtr("testString"),
				Lookupnames: core.BoolPtr(true),
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
		It(`GetConsumptionOfferings(getConsumptionOfferingsOptions *GetConsumptionOfferingsOptions) with pagination`, func(){
			getConsumptionOfferingsOptions := &catalogmanagementv1.GetConsumptionOfferingsOptions{
				Digest: core.BoolPtr(true),
				Catalog: &catalogIDLink,
				Select: core.StringPtr("all"),
				IncludeHidden: core.BoolPtr(true),
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(38)),
			}

			getConsumptionOfferingsOptions.Offset = nil
			getConsumptionOfferingsOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.Offering
			for {
				offeringSearchResult, response, err := catalogManagementService.GetConsumptionOfferings(getConsumptionOfferingsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offeringSearchResult).ToNot(BeNil())
				allResults = append(allResults, offeringSearchResult.Resources...)

				getConsumptionOfferingsOptions.Offset, err = offeringSearchResult.GetNextOffset()
				Expect(err).To(BeNil())

				if getConsumptionOfferingsOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetConsumptionOfferings(getConsumptionOfferingsOptions *GetConsumptionOfferingsOptions) using GetConsumptionOfferingsPager`, func(){
			getConsumptionOfferingsOptions := &catalogmanagementv1.GetConsumptionOfferingsOptions{
				Digest: core.BoolPtr(true),
				Catalog: &catalogIDLink,
				Select: core.StringPtr("all"),
				IncludeHidden: core.BoolPtr(true),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewGetConsumptionOfferingsPager(getConsumptionOfferingsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.Offering
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewGetConsumptionOfferingsPager(getConsumptionOfferingsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetConsumptionOfferings() returned a total of %d item(s) using GetConsumptionOfferingsPager.\n", len(allResults))
		})
	})

	Describe(`ListOfferings - Get list of offerings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListOfferings(listOfferingsOptions *ListOfferingsOptions) with pagination`, func(){
			listOfferingsOptions := &catalogmanagementv1.ListOfferingsOptions{
				CatalogIdentifier: &catalogIDLink,
				Digest: core.BoolPtr(true),
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(38)),
				Name: core.StringPtr("testString"),
				Sort: core.StringPtr("testString"),
				IncludeHidden: core.BoolPtr(true),
			}

			listOfferingsOptions.Offset = nil
			listOfferingsOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.Offering
			for {
				offeringSearchResult, response, err := catalogManagementService.ListOfferings(listOfferingsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(offeringSearchResult).ToNot(BeNil())
				allResults = append(allResults, offeringSearchResult.Resources...)

				listOfferingsOptions.Offset, err = offeringSearchResult.GetNextOffset()
				Expect(err).To(BeNil())

				if listOfferingsOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListOfferings(listOfferingsOptions *ListOfferingsOptions) using OfferingsPager`, func(){
			listOfferingsOptions := &catalogmanagementv1.ListOfferingsOptions{
				CatalogIdentifier: &catalogIDLink,
				Digest: core.BoolPtr(true),
				Limit: core.Int64Ptr(int64(10)),
				Name: core.StringPtr("testString"),
				Sort: core.StringPtr("testString"),
				IncludeHidden: core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewOfferingsPager(listOfferingsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.Offering
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewOfferingsPager(listOfferingsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListOfferings() returned a total of %d item(s) using OfferingsPager.\n", len(allResults))
		})
	})

	Describe(`ListOfferingAudits - Get offering audit logs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListOfferingAudits(listOfferingAuditsOptions *ListOfferingAuditsOptions) with pagination`, func(){
			listOfferingAuditsOptions := &catalogmanagementv1.ListOfferingAuditsOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			listOfferingAuditsOptions.Start = nil
			listOfferingAuditsOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.AuditLogDigest
			for {
				auditLogs, response, err := catalogManagementService.ListOfferingAudits(listOfferingAuditsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLogs).ToNot(BeNil())
				allResults = append(allResults, auditLogs.Audits...)

				listOfferingAuditsOptions.Start, err = auditLogs.GetNextStart()
				Expect(err).To(BeNil())

				if listOfferingAuditsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListOfferingAudits(listOfferingAuditsOptions *ListOfferingAuditsOptions) using OfferingAuditsPager`, func(){
			listOfferingAuditsOptions := &catalogmanagementv1.ListOfferingAuditsOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewOfferingAuditsPager(listOfferingAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.AuditLogDigest
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewOfferingAuditsPager(listOfferingAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListOfferingAudits() returned a total of %d item(s) using OfferingAuditsPager.\n", len(allResults))
		})
	})

	Describe(`GetOfferingAudit - Get an offering audit log entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOfferingAudit(getOfferingAuditOptions *GetOfferingAuditOptions)`, func() {
			getOfferingAuditOptions := &catalogmanagementv1.GetOfferingAuditOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				AuditlogIdentifier: core.StringPtr("testString"),
				Lookupnames: core.BoolPtr(true),
			}

			auditLog, response, err := catalogManagementService.GetOfferingAudit(getOfferingAuditOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditLog).ToNot(BeNil())
		})
	})

	Describe(`SetOfferingPublish - Set offering publish approval settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetOfferingPublish(setOfferingPublishOptions *SetOfferingPublishOptions)`, func() {
			setOfferingPublishOptions := &catalogmanagementv1.SetOfferingPublishOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				ApprovalType: core.StringPtr("pc_managed"),
				Approved: core.StringPtr("true"),
				PortalRecord: core.StringPtr("testString"),
				PortalURL: core.StringPtr("testString"),
				XApproverToken: core.StringPtr("testString"),
			}

			approvalResult, response, err := catalogManagementService.SetOfferingPublish(setOfferingPublishOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(approvalResult).ToNot(BeNil())
		})
	})

	Describe(`DeprecateOffering - Allows offering to be deprecated`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeprecateOffering(deprecateOfferingOptions *DeprecateOfferingOptions)`, func() {
			deprecateOfferingOptions := &catalogmanagementv1.DeprecateOfferingOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				Setting: core.StringPtr("true"),
				Description: core.StringPtr("testString"),
				DaysUntilDeprecate: core.Int64Ptr(int64(38)),
			}

			response, err := catalogManagementService.DeprecateOffering(deprecateOfferingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`ShareOffering - Allows offering to be shared`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ShareOffering(shareOfferingOptions *ShareOfferingOptions)`, func() {
			shareOfferingOptions := &catalogmanagementv1.ShareOfferingOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				IBM: core.BoolPtr(true),
				Public: core.BoolPtr(true),
				Enabled: core.BoolPtr(true),
			}

			shareSetting, response, err := catalogManagementService.ShareOffering(shareOfferingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareSetting).ToNot(BeNil())
		})
	})

	Describe(`GetOfferingAccess - Check for account ID in offering access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOfferingAccess(getOfferingAccessOptions *GetOfferingAccessOptions)`, func() {
			getOfferingAccessOptions := &catalogmanagementv1.GetOfferingAccessOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				AccessIdentifier: core.StringPtr("testString"),
			}

			access, response, err := catalogManagementService.GetOfferingAccess(getOfferingAccessOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(access).ToNot(BeNil())
		})
	})

	Describe(`GetOfferingAccessList - Get offering access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOfferingAccessList(getOfferingAccessListOptions *GetOfferingAccessListOptions) with pagination`, func(){
			getOfferingAccessListOptions := &catalogmanagementv1.GetOfferingAccessListOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			getOfferingAccessListOptions.Start = nil
			getOfferingAccessListOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.Access
			for {
				accessListResult, response, err := catalogManagementService.GetOfferingAccessList(getOfferingAccessListOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(accessListResult).ToNot(BeNil())
				allResults = append(allResults, accessListResult.Resources...)

				getOfferingAccessListOptions.Start, err = accessListResult.GetNextStart()
				Expect(err).To(BeNil())

				if getOfferingAccessListOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetOfferingAccessList(getOfferingAccessListOptions *GetOfferingAccessListOptions) using GetOfferingAccessListPager`, func(){
			getOfferingAccessListOptions := &catalogmanagementv1.GetOfferingAccessListOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewGetOfferingAccessListPager(getOfferingAccessListOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.Access
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewGetOfferingAccessListPager(getOfferingAccessListOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetOfferingAccessList() returned a total of %d item(s) using GetOfferingAccessListPager.\n", len(allResults))
		})
	})

	Describe(`AddOfferingAccessList - Add accesses to offering access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddOfferingAccessList(addOfferingAccessListOptions *AddOfferingAccessListOptions)`, func() {
			addOfferingAccessListOptions := &catalogmanagementv1.AddOfferingAccessListOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				Accesses: []string{"testString"},
			}

			accessListResult, response, err := catalogManagementService.AddOfferingAccessList(addOfferingAccessListOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(accessListResult).ToNot(BeNil())
		})
	})

	Describe(`GetOfferingUpdates - Get version updates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOfferingUpdates(getOfferingUpdatesOptions *GetOfferingUpdatesOptions)`, func() {
			getOfferingUpdatesOptions := &catalogmanagementv1.GetOfferingUpdatesOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				Kind: core.StringPtr("testString"),
				XAuthRefreshToken: core.StringPtr("testString"),
				Target: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				Namespace: core.StringPtr("testString"),
				Sha: core.StringPtr("testString"),
				Channel: core.StringPtr("testString"),
				Namespaces: []string{"testString"},
				AllNamespaces: core.BoolPtr(true),
			}

			versionUpdateDescriptor, response, err := catalogManagementService.GetOfferingUpdates(getOfferingUpdatesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(versionUpdateDescriptor).ToNot(BeNil())
		})
	})

	Describe(`GetOfferingSource - Get offering source`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOfferingSource(getOfferingSourceOptions *GetOfferingSourceOptions)`, func() {
			getOfferingSourceOptions := &catalogmanagementv1.GetOfferingSourceOptions{
				Version: core.StringPtr("testString"),
				Accept: core.StringPtr("application/yaml"),
				CatalogID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Kind: core.StringPtr("testString"),
				Channel: core.StringPtr("testString"),
				AsIs: core.BoolPtr(true),
			}

			result, response, err := catalogManagementService.GetOfferingSource(getOfferingSourceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetOfferingSourceURL - Get offering source URL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOfferingSourceURL(getOfferingSourceURLOptions *GetOfferingSourceURLOptions)`, func() {
			getOfferingSourceURLOptions := &catalogmanagementv1.GetOfferingSourceURLOptions{
				KeyIdentifier: core.StringPtr("testString"),
				Accept: core.StringPtr("application/yaml"),
				CatalogID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			result, response, err := catalogManagementService.GetOfferingSourceURL(getOfferingSourceURLOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
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
				LicenseID: core.StringPtr("testString"),
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

	Describe(`ArchiveVersion - Archive version immediately`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ArchiveVersion(archiveVersionOptions *ArchiveVersionOptions)`, func() {
			archiveVersionOptions := &catalogmanagementv1.ArchiveVersionOptions{
				VersionLocID: core.StringPtr("testString"),
			}

			response, err := catalogManagementService.ArchiveVersion(archiveVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`SetDeprecateVersion - Sets version to be deprecated in a certain time period`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetDeprecateVersion(setDeprecateVersionOptions *SetDeprecateVersionOptions)`, func() {
			setDeprecateVersionOptions := &catalogmanagementv1.SetDeprecateVersionOptions{
				VersionLocID: core.StringPtr("testString"),
				Setting: core.StringPtr("true"),
				Description: core.StringPtr("testString"),
				DaysUntilDeprecate: core.Int64Ptr(int64(38)),
			}

			response, err := catalogManagementService.SetDeprecateVersion(setDeprecateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`ConsumableVersion - Make version consumable for sharing`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ConsumableVersion(consumableVersionOptions *ConsumableVersionOptions)`, func() {
			consumableVersionOptions := &catalogmanagementv1.ConsumableVersionOptions{
				VersionLocID: core.StringPtr("testString"),
			}

			response, err := catalogManagementService.ConsumableVersion(consumableVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`SuspendVersion - Suspend a version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SuspendVersion(suspendVersionOptions *SuspendVersionOptions)`, func() {
			suspendVersionOptions := &catalogmanagementv1.SuspendVersionOptions{
				VersionLocID: core.StringPtr("testString"),
			}

			response, err := catalogManagementService.SuspendVersion(suspendVersionOptions)
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
			flavorModel := &catalogmanagementv1.Flavor{
				Name: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				LabelI18n: make(map[string]string),
				Index: core.Int64Ptr(int64(38)),
			}

			copyVersionOptions := &catalogmanagementv1.CopyVersionOptions{
				VersionLocID: core.StringPtr("testString"),
				Tags: []string{"testString"},
				Content: CreateMockByteArray("This is a mock byte array value."),
				TargetKinds: []string{"testString"},
				FormatKind: core.StringPtr("testString"),
				Flavor: flavorModel,
				WorkingDirectory: core.StringPtr("testString"),
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

	Describe(`CopyFromPreviousVersion - Copy values from a previous version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CopyFromPreviousVersion(copyFromPreviousVersionOptions *CopyFromPreviousVersionOptions)`, func() {
			copyFromPreviousVersionOptions := &catalogmanagementv1.CopyFromPreviousVersionOptions{
				VersionLocID: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
				VersionLocIDToCopyFrom: core.StringPtr("testString"),
			}

			response, err := catalogManagementService.CopyFromPreviousVersion(copyFromPreviousVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
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

	Describe(`DeprecateVersion - Deprecate version immediately - use /archive instead`, func() {
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
		It(`GetNamespaces(getNamespacesOptions *GetNamespacesOptions) with pagination`, func(){
			getNamespacesOptions := &catalogmanagementv1.GetNamespacesOptions{
				ClusterID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				XAuthRefreshToken: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(38)),
			}

			getNamespacesOptions.Offset = nil
			getNamespacesOptions.Limit = core.Int64Ptr(1)

			var allResults []string
			for {
				namespaceSearchResult, response, err := catalogManagementService.GetNamespaces(getNamespacesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(namespaceSearchResult).ToNot(BeNil())
				allResults = append(allResults, namespaceSearchResult.Resources...)

				getNamespacesOptions.Offset, err = namespaceSearchResult.GetNextOffset()
				Expect(err).To(BeNil())

				if getNamespacesOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetNamespaces(getNamespacesOptions *GetNamespacesOptions) using GetNamespacesPager`, func(){
			getNamespacesOptions := &catalogmanagementv1.GetNamespacesOptions{
				ClusterID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				XAuthRefreshToken: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewGetNamespacesPager(getNamespacesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []string
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewGetNamespacesPager(getNamespacesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetNamespaces() returned a total of %d item(s) using GetNamespacesPager.\n", len(allResults))
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
				Channel: core.StringPtr("testString"),
				InstallPlan: core.StringPtr("testString"),
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
				Channel: core.StringPtr("testString"),
				InstallPlan: core.StringPtr("testString"),
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
			deployRequestBodyOverrideValuesModel := &catalogmanagementv1.DeployRequestBodyOverrideValues{
				VsiInstanceName: core.StringPtr("testString"),
				VPCProfile: core.StringPtr("testString"),
				SubnetID: core.StringPtr("testString"),
				VPCID: core.StringPtr("testString"),
				SubnetZone: core.StringPtr("testString"),
				SSHKeyID: core.StringPtr("testString"),
				VPCRegion: core.StringPtr("testString"),
			}
			deployRequestBodyOverrideValuesModel.SetProperty("foo", core.StringPtr("testString"))

			deployRequestBodyEnvironmentVariablesItemModel := &catalogmanagementv1.DeployRequestBodyEnvironmentVariablesItem{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				Secure: core.BoolPtr(true),
			}

			deployRequestBodySchematicsModel := &catalogmanagementv1.DeployRequestBodySchematics{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Tags: []string{"testString"},
				ResourceGroupID: core.StringPtr("testString"),
				TerraformVersion: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			installVersionOptions := &catalogmanagementv1.InstallVersionOptions{
				VersionLocID: core.StringPtr("testString"),
				XAuthRefreshToken: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				Namespace: core.StringPtr("testString"),
				OverrideValues: deployRequestBodyOverrideValuesModel,
				EnvironmentVariables: []catalogmanagementv1.DeployRequestBodyEnvironmentVariablesItem{*deployRequestBodyEnvironmentVariablesItemModel},
				EntitlementApikey: core.StringPtr("testString"),
				Schematics: deployRequestBodySchematicsModel,
				Script: core.StringPtr("testString"),
				ScriptID: core.StringPtr("testString"),
				VersionLocatorID: core.StringPtr("testString"),
				VcenterID: core.StringPtr("testString"),
				VcenterLocation: core.StringPtr("testString"),
				VcenterUser: core.StringPtr("testString"),
				VcenterPassword: core.StringPtr("testString"),
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
			deployRequestBodyOverrideValuesModel := &catalogmanagementv1.DeployRequestBodyOverrideValues{
				VsiInstanceName: core.StringPtr("testString"),
				VPCProfile: core.StringPtr("testString"),
				SubnetID: core.StringPtr("testString"),
				VPCID: core.StringPtr("testString"),
				SubnetZone: core.StringPtr("testString"),
				SSHKeyID: core.StringPtr("testString"),
				VPCRegion: core.StringPtr("testString"),
			}
			deployRequestBodyOverrideValuesModel.SetProperty("foo", core.StringPtr("testString"))

			deployRequestBodyEnvironmentVariablesItemModel := &catalogmanagementv1.DeployRequestBodyEnvironmentVariablesItem{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				Secure: core.BoolPtr(true),
			}

			deployRequestBodySchematicsModel := &catalogmanagementv1.DeployRequestBodySchematics{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Tags: []string{"testString"},
				ResourceGroupID: core.StringPtr("testString"),
				TerraformVersion: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			preinstallVersionOptions := &catalogmanagementv1.PreinstallVersionOptions{
				VersionLocID: core.StringPtr("testString"),
				XAuthRefreshToken: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				Namespace: core.StringPtr("testString"),
				OverrideValues: deployRequestBodyOverrideValuesModel,
				EnvironmentVariables: []catalogmanagementv1.DeployRequestBodyEnvironmentVariablesItem{*deployRequestBodyEnvironmentVariablesItemModel},
				EntitlementApikey: core.StringPtr("testString"),
				Schematics: deployRequestBodySchematicsModel,
				Script: core.StringPtr("testString"),
				ScriptID: core.StringPtr("testString"),
				VersionLocatorID: core.StringPtr("testString"),
				VcenterID: core.StringPtr("testString"),
				VcenterLocation: core.StringPtr("testString"),
				VcenterUser: core.StringPtr("testString"),
				VcenterPassword: core.StringPtr("testString"),
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
			deployRequestBodyOverrideValuesModel := &catalogmanagementv1.DeployRequestBodyOverrideValues{
				VsiInstanceName: core.StringPtr("testString"),
				VPCProfile: core.StringPtr("testString"),
				SubnetID: core.StringPtr("testString"),
				VPCID: core.StringPtr("testString"),
				SubnetZone: core.StringPtr("testString"),
				SSHKeyID: core.StringPtr("testString"),
				VPCRegion: core.StringPtr("testString"),
			}
			deployRequestBodyOverrideValuesModel.SetProperty("foo", core.StringPtr("testString"))

			deployRequestBodyEnvironmentVariablesItemModel := &catalogmanagementv1.DeployRequestBodyEnvironmentVariablesItem{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				Secure: core.BoolPtr(true),
			}

			deployRequestBodySchematicsModel := &catalogmanagementv1.DeployRequestBodySchematics{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Tags: []string{"testString"},
				ResourceGroupID: core.StringPtr("testString"),
				TerraformVersion: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			validateInstallOptions := &catalogmanagementv1.ValidateInstallOptions{
				VersionLocID: core.StringPtr("testString"),
				XAuthRefreshToken: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				Namespace: core.StringPtr("testString"),
				OverrideValues: deployRequestBodyOverrideValuesModel,
				EnvironmentVariables: []catalogmanagementv1.DeployRequestBodyEnvironmentVariablesItem{*deployRequestBodyEnvironmentVariablesItemModel},
				EntitlementApikey: core.StringPtr("testString"),
				Schematics: deployRequestBodySchematicsModel,
				Script: core.StringPtr("testString"),
				ScriptID: core.StringPtr("testString"),
				VersionLocatorID: core.StringPtr("testString"),
				VcenterID: core.StringPtr("testString"),
				VcenterLocation: core.StringPtr("testString"),
				VcenterUser: core.StringPtr("testString"),
				VcenterPassword: core.StringPtr("testString"),
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

	Describe(`SearchObjects - List objects across catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SearchObjects(searchObjectsOptions *SearchObjectsOptions) with pagination`, func(){
			searchObjectsOptions := &catalogmanagementv1.SearchObjectsOptions{
				Query: core.StringPtr("testString"),
				Kind: core.StringPtr("vpe"),
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(38)),
				Collapse: core.BoolPtr(true),
				Digest: core.BoolPtr(true),
			}

			searchObjectsOptions.Offset = nil
			searchObjectsOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.CatalogObject
			for {
				objectSearchResult, response, err := catalogManagementService.SearchObjects(searchObjectsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(objectSearchResult).ToNot(BeNil())
				allResults = append(allResults, objectSearchResult.Resources...)

				searchObjectsOptions.Offset, err = objectSearchResult.GetNextOffset()
				Expect(err).To(BeNil())

				if searchObjectsOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`SearchObjects(searchObjectsOptions *SearchObjectsOptions) using SearchObjectsPager`, func(){
			searchObjectsOptions := &catalogmanagementv1.SearchObjectsOptions{
				Query: core.StringPtr("testString"),
				Kind: core.StringPtr("vpe"),
				Limit: core.Int64Ptr(int64(10)),
				Collapse: core.BoolPtr(true),
				Digest: core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewSearchObjectsPager(searchObjectsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.CatalogObject
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewSearchObjectsPager(searchObjectsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "SearchObjects() returned a total of %d item(s) using SearchObjectsPager.\n", len(allResults))
		})
	})

	Describe(`ListObjects - List objects within a catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListObjects(listObjectsOptions *ListObjectsOptions) with pagination`, func(){
			listObjectsOptions := &catalogmanagementv1.ListObjectsOptions{
				CatalogIdentifier: &catalogIDLink,
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(38)),
				Name: core.StringPtr("testString"),
				Sort: core.StringPtr("testString"),
			}

			listObjectsOptions.Offset = nil
			listObjectsOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.CatalogObject
			for {
				objectListResult, response, err := catalogManagementService.ListObjects(listObjectsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(objectListResult).ToNot(BeNil())
				allResults = append(allResults, objectListResult.Resources...)

				listObjectsOptions.Offset, err = objectListResult.GetNextOffset()
				Expect(err).To(BeNil())

				if listObjectsOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListObjects(listObjectsOptions *ListObjectsOptions) using ObjectsPager`, func(){
			listObjectsOptions := &catalogmanagementv1.ListObjectsOptions{
				CatalogIdentifier: &catalogIDLink,
				Limit: core.Int64Ptr(int64(10)),
				Name: core.StringPtr("testString"),
				Sort: core.StringPtr("testString"),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewObjectsPager(listObjectsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.CatalogObject
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewObjectsPager(listObjectsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListObjects() returned a total of %d item(s) using ObjectsPager.\n", len(allResults))
		})
	})

	Describe(`ListObjectAudits - Get object audit logs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListObjectAudits(listObjectAuditsOptions *ListObjectAuditsOptions) with pagination`, func(){
			listObjectAuditsOptions := &catalogmanagementv1.ListObjectAuditsOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			listObjectAuditsOptions.Start = nil
			listObjectAuditsOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.AuditLogDigest
			for {
				auditLogs, response, err := catalogManagementService.ListObjectAudits(listObjectAuditsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLogs).ToNot(BeNil())
				allResults = append(allResults, auditLogs.Audits...)

				listObjectAuditsOptions.Start, err = auditLogs.GetNextStart()
				Expect(err).To(BeNil())

				if listObjectAuditsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListObjectAudits(listObjectAuditsOptions *ListObjectAuditsOptions) using ObjectAuditsPager`, func(){
			listObjectAuditsOptions := &catalogmanagementv1.ListObjectAuditsOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewObjectAuditsPager(listObjectAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.AuditLogDigest
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewObjectAuditsPager(listObjectAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListObjectAudits() returned a total of %d item(s) using ObjectAuditsPager.\n", len(allResults))
		})
	})

	Describe(`GetObjectAudit - Get an object audit log entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetObjectAudit(getObjectAuditOptions *GetObjectAuditOptions)`, func() {
			getObjectAuditOptions := &catalogmanagementv1.GetObjectAuditOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				AuditlogIdentifier: core.StringPtr("testString"),
				Lookupnames: core.BoolPtr(true),
			}

			auditLog, response, err := catalogManagementService.GetObjectAudit(getObjectAuditOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditLog).ToNot(BeNil())
		})
	})

	Describe(`ConsumableShareObject - Make object consumable for sharing`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ConsumableShareObject(consumableShareObjectOptions *ConsumableShareObjectOptions)`, func() {
			consumableShareObjectOptions := &catalogmanagementv1.ConsumableShareObjectOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
			}

			response, err := catalogManagementService.ConsumableShareObject(consumableShareObjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`ShareObject - Allows object to be shared`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ShareObject(shareObjectOptions *ShareObjectOptions)`, func() {
			shareObjectOptions := &catalogmanagementv1.ShareObjectOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				IBM: core.BoolPtr(true),
				Public: core.BoolPtr(true),
				Enabled: core.BoolPtr(true),
			}

			shareSetting, response, err := catalogManagementService.ShareObject(shareObjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareSetting).ToNot(BeNil())
		})
	})

	Describe(`GetObjectAccessList - Get object access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetObjectAccessList(getObjectAccessListOptions *GetObjectAccessListOptions) with pagination`, func(){
			getObjectAccessListOptions := &catalogmanagementv1.GetObjectAccessListOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			getObjectAccessListOptions.Start = nil
			getObjectAccessListOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.Access
			for {
				accessListResult, response, err := catalogManagementService.GetObjectAccessList(getObjectAccessListOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(accessListResult).ToNot(BeNil())
				allResults = append(allResults, accessListResult.Resources...)

				getObjectAccessListOptions.Start, err = accessListResult.GetNextStart()
				Expect(err).To(BeNil())

				if getObjectAccessListOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetObjectAccessList(getObjectAccessListOptions *GetObjectAccessListOptions) using GetObjectAccessListPager`, func(){
			getObjectAccessListOptions := &catalogmanagementv1.GetObjectAccessListOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewGetObjectAccessListPager(getObjectAccessListOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.Access
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewGetObjectAccessListPager(getObjectAccessListOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetObjectAccessList() returned a total of %d item(s) using GetObjectAccessListPager.\n", len(allResults))
		})
	})

	Describe(`GetObjectAccess - Check for account ID in object access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetObjectAccess(getObjectAccessOptions *GetObjectAccessOptions)`, func() {
			getObjectAccessOptions := &catalogmanagementv1.GetObjectAccessOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				AccessIdentifier: core.StringPtr("testString"),
			}

			access, response, err := catalogManagementService.GetObjectAccess(getObjectAccessOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(access).ToNot(BeNil())
		})
	})

	Describe(`CreateObjectAccess - Add account ID to object access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateObjectAccess(createObjectAccessOptions *CreateObjectAccessOptions)`, func() {
			createObjectAccessOptions := &catalogmanagementv1.CreateObjectAccessOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				AccessIdentifier: core.StringPtr("testString"),
			}

			response, err := catalogManagementService.CreateObjectAccess(createObjectAccessOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
		})
	})

	Describe(`GetObjectAccessListDeprecated - Get object access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetObjectAccessListDeprecated(getObjectAccessListDeprecatedOptions *GetObjectAccessListDeprecatedOptions) with pagination`, func(){
			getObjectAccessListDeprecatedOptions := &catalogmanagementv1.GetObjectAccessListDeprecatedOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(38)),
			}

			getObjectAccessListDeprecatedOptions.Offset = nil
			getObjectAccessListDeprecatedOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.Access
			for {
				objectAccessListResult, response, err := catalogManagementService.GetObjectAccessListDeprecated(getObjectAccessListDeprecatedOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(objectAccessListResult).ToNot(BeNil())
				allResults = append(allResults, objectAccessListResult.Resources...)

				getObjectAccessListDeprecatedOptions.Offset, err = objectAccessListResult.GetNextOffset()
				Expect(err).To(BeNil())

				if getObjectAccessListDeprecatedOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetObjectAccessListDeprecated(getObjectAccessListDeprecatedOptions *GetObjectAccessListDeprecatedOptions) using GetObjectAccessListDeprecatedPager`, func(){
			getObjectAccessListDeprecatedOptions := &catalogmanagementv1.GetObjectAccessListDeprecatedOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewGetObjectAccessListDeprecatedPager(getObjectAccessListDeprecatedOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.Access
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewGetObjectAccessListDeprecatedPager(getObjectAccessListDeprecatedOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetObjectAccessListDeprecated() returned a total of %d item(s) using GetObjectAccessListDeprecatedPager.\n", len(allResults))
		})
	})

	Describe(`AddObjectAccessList - Add accesses to object access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddObjectAccessList(addObjectAccessListOptions *AddObjectAccessListOptions)`, func() {
			addObjectAccessListOptions := &catalogmanagementv1.AddObjectAccessListOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				Accesses: []string{"testString"},
			}

			accessListBulkResponse, response, err := catalogManagementService.AddObjectAccessList(addObjectAccessListOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(accessListBulkResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateOfferingInstance - Create an offering resource instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateOfferingInstance(createOfferingInstanceOptions *CreateOfferingInstanceOptions)`, func() {
			offeringInstanceLastOperationModel := &catalogmanagementv1.OfferingInstanceLastOperation{
				Operation: core.StringPtr("testString"),
				State: core.StringPtr("testString"),
				Message: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Code: core.StringPtr("testString"),
			}

			createOfferingInstanceOptions := &catalogmanagementv1.CreateOfferingInstanceOptions{
				XAuthRefreshToken: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Rev: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
				CRN: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				OfferingID: core.StringPtr("testString"),
				KindFormat: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				VersionID: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				ClusterRegion: core.StringPtr("testString"),
				ClusterNamespaces: []string{"testString"},
				ClusterAllNamespaces: core.BoolPtr(true),
				SchematicsWorkspaceID: core.StringPtr("testString"),
				InstallPlan: core.StringPtr("testString"),
				Channel: core.StringPtr("testString"),
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				ResourceGroupID: core.StringPtr("testString"),
				Location: core.StringPtr("testString"),
				Disabled: core.BoolPtr(true),
				Account: core.StringPtr("testString"),
				LastOperation: offeringInstanceLastOperationModel,
				KindTarget: core.StringPtr("testString"),
				Sha: core.StringPtr("testString"),
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
			offeringInstanceLastOperationModel := &catalogmanagementv1.OfferingInstanceLastOperation{
				Operation: core.StringPtr("testString"),
				State: core.StringPtr("testString"),
				Message: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Code: core.StringPtr("testString"),
			}

			putOfferingInstanceOptions := &catalogmanagementv1.PutOfferingInstanceOptions{
				InstanceIdentifier: core.StringPtr("testString"),
				XAuthRefreshToken: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Rev: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
				CRN: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				OfferingID: core.StringPtr("testString"),
				KindFormat: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				VersionID: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				ClusterRegion: core.StringPtr("testString"),
				ClusterNamespaces: []string{"testString"},
				ClusterAllNamespaces: core.BoolPtr(true),
				SchematicsWorkspaceID: core.StringPtr("testString"),
				InstallPlan: core.StringPtr("testString"),
				Channel: core.StringPtr("testString"),
				Created: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Updated: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Metadata: map[string]interface{}{"anyKey": "anyValue"},
				ResourceGroupID: core.StringPtr("testString"),
				Location: core.StringPtr("testString"),
				Disabled: core.BoolPtr(true),
				Account: core.StringPtr("testString"),
				LastOperation: offeringInstanceLastOperationModel,
				KindTarget: core.StringPtr("testString"),
				Sha: core.StringPtr("testString"),
			}

			offeringInstance, response, err := catalogManagementService.PutOfferingInstance(putOfferingInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(offeringInstance).ToNot(BeNil())
		})
	})

	Describe(`ListOfferingInstanceAudits - Get offering instance audit logs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListOfferingInstanceAudits(listOfferingInstanceAuditsOptions *ListOfferingInstanceAuditsOptions) with pagination`, func(){
			listOfferingInstanceAuditsOptions := &catalogmanagementv1.ListOfferingInstanceAuditsOptions{
				InstanceIdentifier: core.StringPtr("testString"),
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			listOfferingInstanceAuditsOptions.Start = nil
			listOfferingInstanceAuditsOptions.Limit = core.Int64Ptr(1)

			var allResults []catalogmanagementv1.AuditLogDigest
			for {
				auditLogs, response, err := catalogManagementService.ListOfferingInstanceAudits(listOfferingInstanceAuditsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(auditLogs).ToNot(BeNil())
				allResults = append(allResults, auditLogs.Audits...)

				listOfferingInstanceAuditsOptions.Start, err = auditLogs.GetNextStart()
				Expect(err).To(BeNil())

				if listOfferingInstanceAuditsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListOfferingInstanceAudits(listOfferingInstanceAuditsOptions *ListOfferingInstanceAuditsOptions) using OfferingInstanceAuditsPager`, func(){
			listOfferingInstanceAuditsOptions := &catalogmanagementv1.ListOfferingInstanceAuditsOptions{
				InstanceIdentifier: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Lookupnames: core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := catalogManagementService.NewOfferingInstanceAuditsPager(listOfferingInstanceAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []catalogmanagementv1.AuditLogDigest
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = catalogManagementService.NewOfferingInstanceAuditsPager(listOfferingInstanceAuditsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListOfferingInstanceAudits() returned a total of %d item(s) using OfferingInstanceAuditsPager.\n", len(allResults))
		})
	})

	Describe(`GetOfferingInstanceAudit - Get an offering instance audit log entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOfferingInstanceAudit(getOfferingInstanceAuditOptions *GetOfferingInstanceAuditOptions)`, func() {
			getOfferingInstanceAuditOptions := &catalogmanagementv1.GetOfferingInstanceAuditOptions{
				InstanceIdentifier: core.StringPtr("testString"),
				AuditlogIdentifier: core.StringPtr("testString"),
				Lookupnames: core.BoolPtr(true),
			}

			auditLog, response, err := catalogManagementService.GetOfferingInstanceAudit(getOfferingInstanceAuditOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditLog).ToNot(BeNil())
		})
	})

	Describe(`DeleteOfferingAccessList - Delete accesses from offering access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteOfferingAccessList(deleteOfferingAccessListOptions *DeleteOfferingAccessListOptions)`, func() {
			deleteOfferingAccessListOptions := &catalogmanagementv1.DeleteOfferingAccessListOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
				Accesses: []string{"testString"},
			}

			accessListBulkResponse, response, err := catalogManagementService.DeleteOfferingAccessList(deleteOfferingAccessListOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessListBulkResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteOffering - Delete offering`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteOffering(deleteOfferingOptions *DeleteOfferingOptions)`, func() {
			deleteOfferingOptions := &catalogmanagementv1.DeleteOfferingOptions{
				CatalogIdentifier: &catalogIDLink,
				OfferingID: &offeringIDLink,
			}

			response, err := catalogManagementService.DeleteOffering(deleteOfferingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`DeleteObjectAccess - Remove account ID from object access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteObjectAccess(deleteObjectAccessOptions *DeleteObjectAccessOptions)`, func() {
			deleteObjectAccessOptions := &catalogmanagementv1.DeleteObjectAccessOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				AccessIdentifier: core.StringPtr("testString"),
			}

			response, err := catalogManagementService.DeleteObjectAccess(deleteObjectAccessOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`DeleteObjectAccessList - Delete accesses from object access list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteObjectAccessList(deleteObjectAccessListOptions *DeleteObjectAccessListOptions)`, func() {
			deleteObjectAccessListOptions := &catalogmanagementv1.DeleteObjectAccessListOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
				Accesses: []string{"testString"},
			}

			accessListBulkResponse, response, err := catalogManagementService.DeleteObjectAccessList(deleteObjectAccessListOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessListBulkResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteObject - Delete catalog object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteObject(deleteObjectOptions *DeleteObjectOptions)`, func() {
			deleteObjectOptions := &catalogmanagementv1.DeleteObjectOptions{
				CatalogIdentifier: &catalogIDLink,
				ObjectIdentifier: &objectIDLink,
			}

			response, err := catalogManagementService.DeleteObject(deleteObjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`DeleteCatalog - Delete catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCatalog(deleteCatalogOptions *DeleteCatalogOptions)`, func() {
			deleteCatalogOptions := &catalogmanagementv1.DeleteCatalogOptions{
				CatalogIdentifier: &catalogIDLink,
			}

			response, err := catalogManagementService.DeleteCatalog(deleteCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

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
				ClusterID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				VersionLocatorID: core.StringPtr("testString"),
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
				XAuthRefreshToken: core.StringPtr("testString"),
			}

			response, err := catalogManagementService.DeleteOfferingInstance(deleteOfferingInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
