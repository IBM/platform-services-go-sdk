//go:build examples

/**
 * (C) Copyright IBM Corp. 2025.
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

package partnercentersellv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Partner Center Sell service.
//
// The following configuration properties are assumed to be defined:
// PARTNER_CENTER_SELL_URL=<service base url>
// PARTNER_CENTER_SELL_AUTH_TYPE=iam
// PARTNER_CENTER_SELL_APIKEY=<IAM apikey>
// PARTNER_CENTER_SELL_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// PRODUCT_ID_APPROVED=<product id>
// PARTNER_CENTER_SELL_BADGE_ID=<badge id>
// PARTNER_CENTER_SELL_IAM_REGISTRATION_ID=<iam registration id>

// PARTNER_CENTER_SELL_ALT_AUTH_TYPE=iam
// PARTNER_CENTER_SELL_ALT_APIKEY=<IAM apikey>
// PARTNER_CENTER_SELL_ALT_URL=<service base url>

// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`PartnerCenterSellV1 Examples Tests`, func() {

	const externalConfigFile = "../partner_center_sell_v1.env"

	var (
		partnerCenterSellService    *partnercentersellv1.PartnerCenterSellV1
		partnerCenterSellServiceAlt *partnercentersellv1.PartnerCenterSellV1
		config                      map[string]string

		// Variables to hold link values
		accountId                             string
		productIdWithApprovedProgrammaticName string
		badgeId                               string
		brokerIdLink                          string
		catalogDeploymentIdLink               string
		catalogPlanIdLink                     string
		catalogProductIdLink                  string
		productIdLink                         string
		programmaticNameLink                  string
		registrationIdLink                    string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(partnercentersellv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			partnerCenterSellServiceOptions := &partnercentersellv1.PartnerCenterSellV1Options{}

			partnerCenterSellService, err = partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(partnerCenterSellServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(partnerCenterSellService).ToNot(BeNil())
		})
	})

	Describe(`PartnerCenterSellV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRegistration request example`, func() {
			fmt.Println("\nCreateRegistration() result:")
			// begin-create_registration

			primaryContactModel := &partnercentersellv1.PrimaryContact{
				Name:  core.StringPtr("Company Representative"),
				Email: core.StringPtr("companyrep@email.com"),
			}

			createRegistrationOptions := partnerCenterSellService.NewCreateRegistrationOptions(
				"4a5c3c51b97a446fbb1d0e1ef089823b",
				"Beautiful Company",
				primaryContactModel,
			)

			registration, response, err := partnerCenterSellService.CreateRegistration(createRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(registration, "", "  ")
			fmt.Println(string(b))

			// end-create_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(registration).ToNot(BeNil())

			registrationIdLink = *registration.ID
			fmt.Fprintf(GinkgoWriter, "Saved registrationIdLink value: %v\n", registrationIdLink)
		})
		It(`CreateOnboardingProduct request example`, func() {
			fmt.Println("\nCreateOnboardingProduct() result:")
			// begin-create_onboarding_product

			primaryContactModel := &partnercentersellv1.PrimaryContact{
				Name:  core.StringPtr("name"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			createOnboardingProductOptions := partnerCenterSellService.NewCreateOnboardingProductOptions(
				"service",
				primaryContactModel,
			)

			onboardingProduct, response, err := partnerCenterSellService.CreateOnboardingProduct(createOnboardingProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(onboardingProduct, "", "  ")
			fmt.Println(string(b))

			// end-create_onboarding_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(onboardingProduct).ToNot(BeNil())

			productIdLink = *onboardingProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved productIdLink value: %v\n", productIdLink)
		})
		It(`UpdateOnboardingProduct request example`, func() {
			fmt.Println("\nUpdateOnboardingProduct() result:")
			// begin-update_onboarding_product

			onboardingProductPatchModel := &partnercentersellv1.OnboardingProductPatch{
				Unspsc: core.Float64Ptr(float64(12345)),
			}
			onboardingProductPatchModelAsPatch, asPatchErr := onboardingProductPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateOnboardingProductOptions := partnerCenterSellService.NewUpdateOnboardingProductOptions(
				productIdLink,
				onboardingProductPatchModelAsPatch,
			)

			onboardingProduct, response, err := partnerCenterSellService.UpdateOnboardingProduct(updateOnboardingProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(onboardingProduct, "", "  ")
			fmt.Println(string(b))

			// end-update_onboarding_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(onboardingProduct).ToNot(BeNil())

			productIdLink = *onboardingProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved productIdLink value: %v\n", productIdLink)
		})
		It(`CreateCatalogProduct request example`, func() {
			fmt.Println("\nCreateCatalogProduct() result:")
			// begin-create_catalog_product

			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("IBM"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			globalCatalogOverviewUiTranslatedContentModel := &partnercentersellv1.GlobalCatalogOverviewUITranslatedContent{
				DisplayName:     core.StringPtr("My product display name."),
				Description:     core.StringPtr("My product description."),
				LongDescription: core.StringPtr("My product description long description."),
			}

			globalCatalogOverviewUiModel := &partnercentersellv1.GlobalCatalogOverviewUI{
				En: globalCatalogOverviewUiTranslatedContentModel,
			}

			globalCatalogProductMetadataServiceModel := &partnercentersellv1.GlobalCatalogProductMetadataService{
				RcProvisionable: core.BoolPtr(true),
				IamCompatible:   core.BoolPtr(true),
			}

			globalCatalogProductMetadataModel := &partnercentersellv1.GlobalCatalogProductMetadata{
				RcCompatible: core.BoolPtr(true),
				Service:      globalCatalogProductMetadataServiceModel,
			}

			createCatalogProductOptions := partnerCenterSellService.NewCreateCatalogProductOptions(
				productIdLink,
				"1p-service-08-06",
				true,
				false,
				"service",
				[]string{"keyword", "support_ibm"},
				catalogProductProviderModel,
			)
			createCatalogProductOptions.SetOverviewUi(globalCatalogOverviewUiModel)
			createCatalogProductOptions.SetMetadata(globalCatalogProductMetadataModel)

			globalCatalogProduct, response, err := partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogProduct, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogProduct).ToNot(BeNil())

			catalogProductIdLink = *globalCatalogProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogProductIdLink value: %v\n", catalogProductIdLink)
		})
		It(`UpdateCatalogProduct request example`, func() {
			fmt.Println("\nUpdateCatalogProduct() result:")
			// begin-update_catalog_product

			globalCatalogOverviewUiTranslatedContentModel := &partnercentersellv1.GlobalCatalogOverviewUITranslatedContent{
				DisplayName: core.StringPtr("My updated display name."),
			}

			globalCatalogOverviewUiModel := &partnercentersellv1.GlobalCatalogOverviewUI{
				En: globalCatalogOverviewUiTranslatedContentModel,
			}

			globalCatalogProductPatchModel := &partnercentersellv1.GlobalCatalogProductPatch{
				OverviewUi: globalCatalogOverviewUiModel,
			}
			globalCatalogProductPatchModelAsPatch, asPatchErr := globalCatalogProductPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogProductOptions := partnerCenterSellService.NewUpdateCatalogProductOptions(
				productIdLink,
				catalogProductIdLink,
				globalCatalogProductPatchModelAsPatch,
			)

			globalCatalogProduct, response, err := partnerCenterSellService.UpdateCatalogProduct(updateCatalogProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogProduct, "", "  ")
			fmt.Println(string(b))

			// end-update_catalog_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogProduct).ToNot(BeNil())

			catalogProductIdLink = *globalCatalogProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogProductIdLink value: %v\n", catalogProductIdLink)
		})
		It(`CreateCatalogPlan request example`, func() {
			fmt.Println("\nCreateCatalogPlan() result:")
			// begin-create_catalog_plan

			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("IBM"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			globalCatalogOverviewUiTranslatedContentModel := &partnercentersellv1.GlobalCatalogOverviewUITranslatedContent{
				DisplayName:     core.StringPtr("My plan"),
				Description:     core.StringPtr("My plan description."),
				LongDescription: core.StringPtr("My plan long description."),
			}

			globalCatalogOverviewUiModel := &partnercentersellv1.GlobalCatalogOverviewUI{
				En: globalCatalogOverviewUiTranslatedContentModel,
			}

			globalCatalogPlanMetadataServiceModel := &partnercentersellv1.GlobalCatalogPlanMetadataService{
				RcProvisionable: core.BoolPtr(false),
				IamCompatible:   core.BoolPtr(true),
			}

			globalCatalogMetadataPricingModel := &partnercentersellv1.GlobalCatalogMetadataPricing{
				Type:   core.StringPtr("paid"),
				Origin: core.StringPtr("pricing_catalog"),
			}

			globalCatalogPlanMetadataModel := &partnercentersellv1.GlobalCatalogPlanMetadata{
				RcCompatible: core.BoolPtr(true),
				Service:      globalCatalogPlanMetadataServiceModel,
				Pricing:      globalCatalogMetadataPricingModel,
			}

			createCatalogPlanOptions := partnerCenterSellService.NewCreateCatalogPlanOptions(
				productIdLink,
				catalogProductIdLink,
				"free-plan2",
				true,
				false,
				"plan",
				[]string{"ibm_created"},
				catalogProductProviderModel,
			)
			createCatalogPlanOptions.SetOverviewUi(globalCatalogOverviewUiModel)
			createCatalogPlanOptions.SetMetadata(globalCatalogPlanMetadataModel)

			globalCatalogPlan, response, err := partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogPlan, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogPlan).ToNot(BeNil())

			catalogPlanIdLink = *globalCatalogPlan.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogPlanIdLink value: %v\n", catalogPlanIdLink)
		})
		It(`UpdateCatalogPlan request example`, func() {
			fmt.Println("\nUpdateCatalogPlan() result:")
			// begin-update_catalog_plan

			globalCatalogMetadataPricingModel := &partnercentersellv1.GlobalCatalogMetadataPricing{
				Type:   core.StringPtr("free"),
				Origin: core.StringPtr("pricing_catalog"),
			}

			globalCatalogPlanMetadataModel := &partnercentersellv1.GlobalCatalogPlanMetadata{
				RcCompatible: core.BoolPtr(true),
				Pricing:      globalCatalogMetadataPricingModel,
			}

			globalCatalogPlanPatchModel := &partnercentersellv1.GlobalCatalogPlanPatch{
				Metadata: globalCatalogPlanMetadataModel,
			}
			globalCatalogPlanPatchModelAsPatch, asPatchErr := globalCatalogPlanPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogPlanOptions := partnerCenterSellService.NewUpdateCatalogPlanOptions(
				productIdLink,
				catalogProductIdLink,
				catalogPlanIdLink,
				globalCatalogPlanPatchModelAsPatch,
			)

			globalCatalogPlan, response, err := partnerCenterSellService.UpdateCatalogPlan(updateCatalogPlanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogPlan, "", "  ")
			fmt.Println(string(b))

			// end-update_catalog_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogPlan).ToNot(BeNil())

			catalogPlanIdLink = *globalCatalogPlan.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogPlanIdLink value: %v\n", catalogPlanIdLink)
		})
		It(`CreateCatalogDeployment request example`, func() {
			fmt.Println("\nCreateCatalogDeployment() result:")
			// begin-create_catalog_deployment

			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("IBM"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			globalCatalogDeploymentMetadataServiceModel := &partnercentersellv1.GlobalCatalogDeploymentMetadataService{
				RcProvisionable: core.BoolPtr(true),
				IamCompatible:   core.BoolPtr(true),
			}

			globalCatalogMetadataDeploymentBrokerModel := &partnercentersellv1.GlobalCatalogMetadataDeploymentBroker{
				Name: core.StringPtr("brokerunique1234"),
				Guid: core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3c2"),
			}

			globalCatalogMetadataDeploymentModel := &partnercentersellv1.GlobalCatalogMetadataDeployment{
				Broker:      globalCatalogMetadataDeploymentBrokerModel,
				Location:    core.StringPtr("eu-gb"),
				LocationURL: core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb"),
				TargetCrn:   core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb"),
			}

			globalCatalogDeploymentMetadataModel := &partnercentersellv1.GlobalCatalogDeploymentMetadata{
				RcCompatible: core.BoolPtr(true),
				Service:      globalCatalogDeploymentMetadataServiceModel,
				Deployment:   globalCatalogMetadataDeploymentModel,
			}

			createCatalogDeploymentOptions := partnerCenterSellService.NewCreateCatalogDeploymentOptions(
				productIdLink,
				catalogProductIdLink,
				catalogPlanIdLink,
				"deployment-eu-de",
				true,
				false,
				"deployment",
				[]string{"eu-gb"},
				catalogProductProviderModel,
			)
			createCatalogDeploymentOptions.SetMetadata(globalCatalogDeploymentMetadataModel)

			globalCatalogDeployment, response, err := partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogDeployment, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogDeployment).ToNot(BeNil())

			catalogDeploymentIdLink = *globalCatalogDeployment.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogDeploymentIdLink value: %v\n", catalogDeploymentIdLink)
		})
		It(`UpdateCatalogDeployment request example`, func() {
			fmt.Println("\nUpdateCatalogDeployment() result:")
			// begin-update_catalog_deployment

			globalCatalogMetadataDeploymentBrokerModel := &partnercentersellv1.GlobalCatalogMetadataDeploymentBroker{
				Name: core.StringPtr("another-broker"),
				Guid: core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3cf"),
			}

			globalCatalogMetadataDeploymentModel := &partnercentersellv1.GlobalCatalogMetadataDeployment{
				Broker:      globalCatalogMetadataDeploymentBrokerModel,
				Location:    core.StringPtr("eu-gb"),
				LocationURL: core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb"),
				TargetCrn:   core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb"),
			}

			globalCatalogDeploymentMetadataModel := &partnercentersellv1.GlobalCatalogDeploymentMetadata{
				Deployment: globalCatalogMetadataDeploymentModel,
			}

			globalCatalogDeploymentPatchModel := &partnercentersellv1.GlobalCatalogDeploymentPatch{
				Metadata: globalCatalogDeploymentMetadataModel,
			}
			globalCatalogDeploymentPatchModelAsPatch, asPatchErr := globalCatalogDeploymentPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogDeploymentOptions := partnerCenterSellService.NewUpdateCatalogDeploymentOptions(
				productIdLink,
				catalogProductIdLink,
				catalogPlanIdLink,
				catalogDeploymentIdLink,
				globalCatalogDeploymentPatchModelAsPatch,
			)

			globalCatalogDeployment, response, err := partnerCenterSellService.UpdateCatalogDeployment(updateCatalogDeploymentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogDeployment, "", "  ")
			fmt.Println(string(b))

			// end-update_catalog_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogDeployment).ToNot(BeNil())

			catalogDeploymentIdLink = *globalCatalogDeployment.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogDeploymentIdLink value: %v\n", catalogDeploymentIdLink)
		})
		It(`CreateIamRegistration request example`, func() {
			fmt.Println("\nCreateIamRegistration() result:")
			// begin-create_iam_registration

			iamServiceRegistrationDescriptionObjectModel := &partnercentersellv1.IamServiceRegistrationDescriptionObject{
				Default: core.StringPtr("View dashboard"),
				En:      core.StringPtr("View dashboard"),
				De:      core.StringPtr("View dashboard"),
				Es:      core.StringPtr("View dashboard"),
				Fr:      core.StringPtr("View dashboard"),
				It:      core.StringPtr("View dashboard"),
				Ja:      core.StringPtr("View dashboard"),
				Ko:      core.StringPtr("View dashboard"),
				PtBr:    core.StringPtr("View dashboard"),
				ZhTw:    core.StringPtr("View dashboard"),
				ZhCn:    core.StringPtr("View dashboard"),
			}

			iamServiceRegistrationDisplayNameObjectModel := &partnercentersellv1.IamServiceRegistrationDisplayNameObject{
				Default: core.StringPtr("View dashboard"),
				En:      core.StringPtr("View dashboard"),
				De:      core.StringPtr("View dashboard"),
				Es:      core.StringPtr("View dashboard"),
				Fr:      core.StringPtr("View dashboard"),
				It:      core.StringPtr("View dashboard"),
				Ja:      core.StringPtr("View dashboard"),
				Ko:      core.StringPtr("View dashboard"),
				PtBr:    core.StringPtr("View dashboard"),
				ZhTw:    core.StringPtr("View dashboard"),
				ZhCn:    core.StringPtr("View dashboard"),
			}

			iamServiceRegistrationActionModel := &partnercentersellv1.IamServiceRegistrationAction{
				ID:          core.StringPtr("pet-store.dashboard.view"),
				Roles:       []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"},
				Description: iamServiceRegistrationDescriptionObjectModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
			}

			iamServiceRegistrationSupportedAnonymousAccessAttributesModel := &partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes{
				AccountID:            core.StringPtr("testString"),
				ServiceName:          core.StringPtr("testString"),
				AdditionalProperties: map[string]string{"key1": "testString"},
			}

			iamServiceRegistrationSupportedAnonymousAccessModel := &partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{
				Attributes: iamServiceRegistrationSupportedAnonymousAccessAttributesModel,
				Roles:      []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"},
			}

			supportedAttributesOptionsModel := &partnercentersellv1.SupportedAttributesOptions{
				Operators: []string{"stringMatch", "stringEquals"},
			}

			supportedAttributeUiInputValueModel := &partnercentersellv1.SupportedAttributeUiInputValue{
				Value:       core.StringPtr("staticValue"),
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
			}

			supportedAttributeUiInputGstModel := &partnercentersellv1.SupportedAttributeUiInputGst{
				Query:             core.StringPtr("ghost query"),
				ValuePropertyName: core.StringPtr("instance"),
				InputOptionLabel:  core.StringPtr("{name} - {instance_id}"),
			}

			supportedAttributeUiInputDetailsModel := &partnercentersellv1.SupportedAttributeUiInputDetails{
				Type:   core.StringPtr("gst"),
				Values: []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel},
				Gst:    supportedAttributeUiInputGstModel,
			}

			supportedAttributeUiModel := &partnercentersellv1.SupportedAttributeUi{
				InputType:    core.StringPtr("selector"),
				InputDetails: supportedAttributeUiInputDetailsModel,
			}

			iamServiceRegistrationSupportedAttributeModel := &partnercentersellv1.IamServiceRegistrationSupportedAttribute{
				Key:         core.StringPtr("testAttribute"),
				Options:     supportedAttributesOptionsModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
				Description: iamServiceRegistrationDescriptionObjectModel,
				Ui:          supportedAttributeUiModel,
			}

			supportAuthorizationSubjectAttributeModel := &partnercentersellv1.SupportAuthorizationSubjectAttribute{}

			iamServiceRegistrationSupportedAuthorizationSubjectModel := &partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{
				Attributes: supportAuthorizationSubjectAttributeModel,
				Roles:      []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"},
			}

			iamServiceRegistrationSupportedRoleModel := &partnercentersellv1.IamServiceRegistrationSupportedRole{
				ID:          core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader"),
				Description: iamServiceRegistrationDescriptionObjectModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
			}

			environmentAttributeOptionsModel := &partnercentersellv1.EnvironmentAttributeOptions{
				Hidden: core.BoolPtr(true),
			}

			environmentAttributeModel := &partnercentersellv1.EnvironmentAttribute{
				Key:     core.StringPtr("networkType"),
				Values:  []string{"public"},
				Options: environmentAttributeOptionsModel,
			}

			iamServiceRegistrationSupportedNetworkModel := &partnercentersellv1.IamServiceRegistrationSupportedNetwork{
				EnvironmentAttributes: []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel},
			}

			createIamRegistrationOptions := partnerCenterSellService.NewCreateIamRegistrationOptions(
				productIdLink,
				"pet-store",
			)
			createIamRegistrationOptions.SetEnabled(true)
			createIamRegistrationOptions.SetActions([]partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel})
			createIamRegistrationOptions.SetAdditionalPolicyScopes([]string{"pet-store"})
			createIamRegistrationOptions.SetDisplayName(iamServiceRegistrationDisplayNameObjectModel)
			createIamRegistrationOptions.SetParentIds([]string{})
			createIamRegistrationOptions.SetSupportedAnonymousAccesses([]partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel})
			createIamRegistrationOptions.SetSupportedAttributes([]partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel})
			createIamRegistrationOptions.SetSupportedAuthorizationSubjects([]partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel})
			createIamRegistrationOptions.SetSupportedRoles([]partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel})
			createIamRegistrationOptions.SetSupportedNetwork(iamServiceRegistrationSupportedNetworkModel)

			iamServiceRegistration, response, err := partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iamServiceRegistration, "", "  ")
			fmt.Println(string(b))

			// end-create_iam_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(iamServiceRegistration).ToNot(BeNil())

			programmaticNameLink = *iamServiceRegistration.Name
			fmt.Fprintf(GinkgoWriter, "Saved programmaticNameLink value: %v\n", programmaticNameLink)
		})
		It(`UpdateIamRegistration request example`, func() {
			fmt.Println("\nUpdateIamRegistration() result:")
			// begin-update_iam_registration

			iamServiceRegistrationDescriptionObjectModel := &partnercentersellv1.IamServiceRegistrationDescriptionObject{
				Default: core.StringPtr("View dashboard"),
				En:      core.StringPtr("View dashboard"),
				De:      core.StringPtr("View dashboard"),
				Es:      core.StringPtr("View dashboard"),
				Fr:      core.StringPtr("View dashboard"),
				It:      core.StringPtr("View dashboard"),
				Ja:      core.StringPtr("View dashboard"),
				Ko:      core.StringPtr("View dashboard"),
				PtBr:    core.StringPtr("View dashboard"),
				ZhTw:    core.StringPtr("View dashboard"),
				ZhCn:    core.StringPtr("View dashboard"),
			}

			iamServiceRegistrationDisplayNameObjectModel := &partnercentersellv1.IamServiceRegistrationDisplayNameObject{
				Default: core.StringPtr("View dashboard"),
				En:      core.StringPtr("View dashboard"),
				De:      core.StringPtr("View dashboard"),
				Es:      core.StringPtr("View dashboard"),
				Fr:      core.StringPtr("View dashboard"),
				It:      core.StringPtr("View dashboard"),
				Ja:      core.StringPtr("View dashboard"),
				Ko:      core.StringPtr("View dashboard"),
				PtBr:    core.StringPtr("View dashboard"),
				ZhTw:    core.StringPtr("View dashboard"),
				ZhCn:    core.StringPtr("View dashboard"),
			}

			iamServiceRegistrationActionModel := &partnercentersellv1.IamServiceRegistrationAction{
				ID:          core.StringPtr("pet-store.dashboard.view"),
				Roles:       []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"},
				Description: iamServiceRegistrationDescriptionObjectModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
			}

			iamServiceRegistrationSupportedAnonymousAccessAttributesModel := &partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes{
				AccountID:            core.StringPtr("testString"),
				ServiceName:          core.StringPtr("testString"),
				AdditionalProperties: map[string]string{"key1": "testString"},
			}

			iamServiceRegistrationSupportedAnonymousAccessModel := &partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{
				Attributes: iamServiceRegistrationSupportedAnonymousAccessAttributesModel,
				Roles:      []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"},
			}

			supportedAttributesOptionsModel := &partnercentersellv1.SupportedAttributesOptions{
				Operators: []string{"stringMatch", "stringEquals"},
			}

			supportedAttributeUiInputValueModel := &partnercentersellv1.SupportedAttributeUiInputValue{
				Value:       core.StringPtr("staticValue"),
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
			}

			supportedAttributeUiInputGstModel := &partnercentersellv1.SupportedAttributeUiInputGst{
				Query:             core.StringPtr("ghost query"),
				ValuePropertyName: core.StringPtr("instance"),
				InputOptionLabel:  core.StringPtr("{name} - {instance_id}"),
			}

			supportedAttributeUiInputDetailsModel := &partnercentersellv1.SupportedAttributeUiInputDetails{
				Type:   core.StringPtr("gst"),
				Values: []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel},
				Gst:    supportedAttributeUiInputGstModel,
			}

			supportedAttributeUiModel := &partnercentersellv1.SupportedAttributeUi{
				InputType:    core.StringPtr("selector"),
				InputDetails: supportedAttributeUiInputDetailsModel,
			}

			iamServiceRegistrationSupportedAttributeModel := &partnercentersellv1.IamServiceRegistrationSupportedAttribute{
				Key:         core.StringPtr("testAttribute"),
				Options:     supportedAttributesOptionsModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
				Description: iamServiceRegistrationDescriptionObjectModel,
				Ui:          supportedAttributeUiModel,
			}

			supportAuthorizationSubjectAttributeModel := &partnercentersellv1.SupportAuthorizationSubjectAttribute{}

			iamServiceRegistrationSupportedAuthorizationSubjectModel := &partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{
				Attributes: supportAuthorizationSubjectAttributeModel,
				Roles:      []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"},
			}

			iamServiceRegistrationSupportedRoleModel := &partnercentersellv1.IamServiceRegistrationSupportedRole{
				ID:          core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader"),
				Description: iamServiceRegistrationDescriptionObjectModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
			}

			environmentAttributeOptionsModel := &partnercentersellv1.EnvironmentAttributeOptions{
				Hidden: core.BoolPtr(true),
			}

			environmentAttributeModel := &partnercentersellv1.EnvironmentAttribute{
				Key:     core.StringPtr("networkType"),
				Values:  []string{"public"},
				Options: environmentAttributeOptionsModel,
			}

			iamServiceRegistrationSupportedNetworkModel := &partnercentersellv1.IamServiceRegistrationSupportedNetwork{
				EnvironmentAttributes: []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel},
			}

			iamServiceRegistrationPatchModel := &partnercentersellv1.IamServiceRegistrationPatch{
				Enabled:                        core.BoolPtr(true),
				Actions:                        []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel},
				AdditionalPolicyScopes:         []string{"pet-store"},
				DisplayName:                    iamServiceRegistrationDisplayNameObjectModel,
				ParentIds:                      []string{},
				SupportedAnonymousAccesses:     []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel},
				SupportedAttributes:            []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel},
				SupportedAuthorizationSubjects: []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel},
				SupportedRoles:                 []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel},
				SupportedNetwork:               iamServiceRegistrationSupportedNetworkModel,
			}
			iamServiceRegistrationPatchModelAsPatch, asPatchErr := iamServiceRegistrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateIamRegistrationOptions := partnerCenterSellService.NewUpdateIamRegistrationOptions(
				productIdLink,
				programmaticNameLink,
				iamServiceRegistrationPatchModelAsPatch,
			)

			iamServiceRegistration, response, err := partnerCenterSellService.UpdateIamRegistration(updateIamRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iamServiceRegistration, "", "  ")
			fmt.Println(string(b))

			// end-update_iam_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iamServiceRegistration).ToNot(BeNil())

			programmaticNameLink = *iamServiceRegistration.Name
			fmt.Fprintf(GinkgoWriter, "Saved programmaticNameLink value: %v\n", programmaticNameLink)
		})
		It(`CreateResourceBroker request example`, func() {
			fmt.Println("\nCreateResourceBroker() result:")
			// begin-create_resource_broker

			createResourceBrokerOptions := partnerCenterSellService.NewCreateResourceBrokerOptions(
				"bearer",
				"brokername",
				"https://broker-url-for-my-service.com",
				"provision_through",
			)
			createResourceBrokerOptions.SetAuthUsername("apikey")
			createResourceBrokerOptions.SetResourceGroupCrn("crn:v1:bluemix:public:resource-controller::a/4a5c3c51b97a446fbb1d0e1ef089823b::resource-group:4fae20bd538a4a738475350dfdc1596f")
			createResourceBrokerOptions.SetState("active")
			createResourceBrokerOptions.SetAllowContextUpdates(false)
			createResourceBrokerOptions.SetCatalogType("service")
			createResourceBrokerOptions.SetRegion("global")

			broker, response, err := partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(broker, "", "  ")
			fmt.Println(string(b))

			// end-create_resource_broker

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(broker).ToNot(BeNil())

			brokerIdLink = *broker.ID
			fmt.Fprintf(GinkgoWriter, "Saved brokerIdLink value: %v\n", brokerIdLink)
		})
		It(`GetRegistration request example`, func() {
			fmt.Println("\nGetRegistration() result:")
			// begin-get_registration

			getRegistrationOptions := partnerCenterSellService.NewGetRegistrationOptions(
				registrationIdLink,
			)

			registration, response, err := partnerCenterSellService.GetRegistration(getRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(registration, "", "  ")
			fmt.Println(string(b))

			// end-get_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(registration).ToNot(BeNil())
		})
		It(`UpdateRegistration request example`, func() {
			fmt.Println("\nUpdateRegistration() result:")
			// begin-update_registration

			registrationPatchModel := &partnercentersellv1.RegistrationPatch{}
			registrationPatchModelAsPatch, asPatchErr := registrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateRegistrationOptions := partnerCenterSellService.NewUpdateRegistrationOptions(
				registrationIdLink,
				registrationPatchModelAsPatch,
			)

			registration, response, err := partnerCenterSellService.UpdateRegistration(updateRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(registration, "", "  ")
			fmt.Println(string(b))

			// end-update_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(registration).ToNot(BeNil())
		})
		It(`GetOnboardingProduct request example`, func() {
			fmt.Println("\nGetOnboardingProduct() result:")
			// begin-get_onboarding_product

			getOnboardingProductOptions := partnerCenterSellService.NewGetOnboardingProductOptions(
				productIdLink,
			)

			onboardingProduct, response, err := partnerCenterSellService.GetOnboardingProduct(getOnboardingProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(onboardingProduct, "", "  ")
			fmt.Println(string(b))

			// end-get_onboarding_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(onboardingProduct).ToNot(BeNil())

			productIdLink = *onboardingProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved productIdLink value: %v\n", productIdLink)
		})
		It(`GetCatalogProduct request example`, func() {
			fmt.Println("\nGetCatalogProduct() result:")
			// begin-get_catalog_product

			getCatalogProductOptions := partnerCenterSellService.NewGetCatalogProductOptions(
				productIdLink,
				catalogProductIdLink,
			)

			globalCatalogProduct, response, err := partnerCenterSellService.GetCatalogProduct(getCatalogProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogProduct, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogProduct).ToNot(BeNil())

			catalogProductIdLink = *globalCatalogProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogProductIdLink value: %v\n", catalogProductIdLink)
		})
		It(`GetCatalogPlan request example`, func() {
			fmt.Println("\nGetCatalogPlan() result:")
			// begin-get_catalog_plan

			getCatalogPlanOptions := partnerCenterSellService.NewGetCatalogPlanOptions(
				productIdLink,
				catalogProductIdLink,
				catalogPlanIdLink,
			)

			globalCatalogPlan, response, err := partnerCenterSellService.GetCatalogPlan(getCatalogPlanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogPlan, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogPlan).ToNot(BeNil())

			catalogPlanIdLink = *globalCatalogPlan.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogPlanIdLink value: %v\n", catalogPlanIdLink)
		})
		It(`GetCatalogDeployment request example`, func() {
			fmt.Println("\nGetCatalogDeployment() result:")
			// begin-get_catalog_deployment

			getCatalogDeploymentOptions := partnerCenterSellService.NewGetCatalogDeploymentOptions(
				productIdLink,
				catalogProductIdLink,
				catalogPlanIdLink,
				catalogDeploymentIdLink,
			)

			globalCatalogDeployment, response, err := partnerCenterSellService.GetCatalogDeployment(getCatalogDeploymentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogDeployment, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogDeployment).ToNot(BeNil())

			catalogDeploymentIdLink = *globalCatalogDeployment.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogDeploymentIdLink value: %v\n", catalogDeploymentIdLink)
		})
		It(`GetIamRegistration request example`, func() {
			fmt.Println("\nGetIamRegistration() result:")
			// begin-get_iam_registration

			getIamRegistrationOptions := partnerCenterSellService.NewGetIamRegistrationOptions(
				productIdLink,
				programmaticNameLink,
			)

			iamServiceRegistration, response, err := partnerCenterSellService.GetIamRegistration(getIamRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iamServiceRegistration, "", "  ")
			fmt.Println(string(b))

			// end-get_iam_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iamServiceRegistration).ToNot(BeNil())

			programmaticNameLink = *iamServiceRegistration.Name
			fmt.Fprintf(GinkgoWriter, "Saved programmaticNameLink value: %v\n", programmaticNameLink)
		})
		It(`UpdateResourceBroker request example`, func() {
			fmt.Println("\nUpdateResourceBroker() result:")
			// begin-update_resource_broker

			brokerPatchModel := &partnercentersellv1.BrokerPatch{
				BrokerURL: core.StringPtr("https://my-updated-broker-url.com"),
			}
			brokerPatchModelAsPatch, asPatchErr := brokerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateResourceBrokerOptions := partnerCenterSellService.NewUpdateResourceBrokerOptions(
				brokerIdLink,
				brokerPatchModelAsPatch,
			)

			broker, response, err := partnerCenterSellService.UpdateResourceBroker(updateResourceBrokerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(broker, "", "  ")
			fmt.Println(string(b))

			// end-update_resource_broker

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(broker).ToNot(BeNil())
		})
		It(`GetResourceBroker request example`, func() {
			fmt.Println("\nGetResourceBroker() result:")
			// begin-get_resource_broker

			getResourceBrokerOptions := partnerCenterSellService.NewGetResourceBrokerOptions(
				brokerIdLink,
			)

			broker, response, err := partnerCenterSellService.GetResourceBroker(getResourceBrokerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(broker, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_broker

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(broker).ToNot(BeNil())
		})
		It(`ListProductBadges request example`, func() {
			fmt.Println("\nListProductBadges() result:")
			// begin-list_product_badges

			listProductBadgesOptions := partnerCenterSellService.NewListProductBadgesOptions()

			productBadgeCollection, response, err := partnerCenterSellService.ListProductBadges(listProductBadgesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(productBadgeCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_product_badges

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productBadgeCollection).ToNot(BeNil())
		})
		It(`GetProductBadge request example`, func() {
			fmt.Println("\nGetProductBadge() result:")
			// begin-get_product_badge

			getProductBadgeOptions := partnerCenterSellService.NewGetProductBadgeOptions(
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			)

			productBadge, response, err := partnerCenterSellService.GetProductBadge(getProductBadgeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(productBadge, "", "  ")
			fmt.Println(string(b))

			// end-get_product_badge

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productBadge).ToNot(BeNil())
		})
		It(`DeleteRegistration request example`, func() {
			// begin-delete_registration

			deleteRegistrationOptions := partnerCenterSellService.NewDeleteRegistrationOptions(
				registrationIdLink,
			)

			response, err := partnerCenterSellService.DeleteRegistration(deleteRegistrationOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteRegistration(): %d\n", response.StatusCode)
			}

			// end-delete_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteCatalogDeployment request example`, func() {
			// begin-delete_catalog_deployment

			deleteCatalogDeploymentOptions := partnerCenterSellService.NewDeleteCatalogDeploymentOptions(
				productIdLink,
				catalogProductIdLink,
				catalogPlanIdLink,
				catalogDeploymentIdLink,
			)

			response, err := partnerCenterSellService.DeleteCatalogDeployment(deleteCatalogDeploymentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteCatalogDeployment(): %d\n", response.StatusCode)
			}

			// end-delete_catalog_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteCatalogPlan request example`, func() {
			// begin-delete_catalog_plan

			deleteCatalogPlanOptions := partnerCenterSellService.NewDeleteCatalogPlanOptions(
				productIdLink,
				catalogProductIdLink,
				catalogPlanIdLink,
			)

			response, err := partnerCenterSellService.DeleteCatalogPlan(deleteCatalogPlanOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteCatalogPlan(): %d\n", response.StatusCode)
			}

			// end-delete_catalog_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteCatalogProduct request example`, func() {
			// begin-delete_catalog_product

			deleteCatalogProductOptions := partnerCenterSellService.NewDeleteCatalogProductOptions(
				productIdLink,
				catalogProductIdLink,
			)

			response, err := partnerCenterSellService.DeleteCatalogProduct(deleteCatalogProductOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteCatalogProduct(): %d\n", response.StatusCode)
			}

			// end-delete_catalog_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteIamRegistration request example`, func() {
			// begin-delete_iam_registration

			deleteIamRegistrationOptions := partnerCenterSellService.NewDeleteIamRegistrationOptions(
				productIdLink,
				programmaticNameLink,
			)

			response, err := partnerCenterSellService.DeleteIamRegistration(deleteIamRegistrationOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteIamRegistration(): %d\n", response.StatusCode)
			}

			// end-delete_iam_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteOnboardingProduct request example`, func() {
			// begin-delete_onboarding_product

			deleteOnboardingProductOptions := partnerCenterSellService.NewDeleteOnboardingProductOptions(
				productIdLink,
			)

			response, err := partnerCenterSellService.DeleteOnboardingProduct(deleteOnboardingProductOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteOnboardingProduct(): %d\n", response.StatusCode)
			}

			// end-delete_onboarding_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteResourceBroker request example`, func() {
			// begin-delete_resource_broker

			deleteResourceBrokerOptions := partnerCenterSellService.NewDeleteResourceBrokerOptions(
				brokerIdLink,
			)

			response, err := partnerCenterSellService.DeleteResourceBroker(deleteResourceBrokerOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteResourceBroker(): %d\n", response.StatusCode)
			}

			// end-delete_resource_broker

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
