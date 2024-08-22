//go:build examples

/**
 * (C) Copyright IBM Corp. 2024.
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
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Partner Center Sell service.
//
// The following configuration properties are assumed to be defined:
// PARTNER_CENTER_SELL_URL=<service base url>
// PARTNER_CENTER_SELL_AUTH_TYPE=iam
// PARTNER_CENTER_SELL_APIKEY=<IAM apikey>
// PARTNER_CENTER_SELL_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`PartnerCenterSellV1 Examples Tests`, func() {

	const externalConfigFile = "../partner_center_sell_v1.env"

	var (
		partnerCenterSellService              *partnercentersellv1.PartnerCenterSellV1
		config                                map[string]string
		accountId                             string
		registrationId                        string
		productIdWithApprovedProgrammaticName string
		productId                             string
		catalogProductId                      string
		catalogPlanId                         string
		catalogDeploymentId                   string
		brokerId                              string
		iamServiceRegistrationId              string
		badgeId                               string
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

			accountId = config["ACCOUNT_ID"]
			Expect(accountId).ToNot(BeEmpty())

			productIdWithApprovedProgrammaticName = config["PRODUCT_ID_APPROVED"]
			Expect(productIdWithApprovedProgrammaticName).ToNot(BeEmpty())

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
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
			}

			createRegistrationOptions := partnerCenterSellService.NewCreateRegistrationOptions(
				accountId,
				"Sample_company_sdk",
				primaryContactModel,
			)

			registration, response, err := partnerCenterSellService.CreateRegistration(createRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(registration, "", "  ")
			registrationId = *registration.ID

			fmt.Println(string(b))

			// end-create_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(registration).ToNot(BeNil())
		})
		It(`GetRegistration request example`, func() {
			fmt.Println("\nGetRegistration() result:")
			// begin-get_registration

			getRegistrationOptions := partnerCenterSellService.NewGetRegistrationOptions(
				registrationId,
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
				registrationId,
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
		It(`DeleteRegistration request example`, func() {
			// begin-delete_registration

			deleteRegistrationOptions := partnerCenterSellService.NewDeleteRegistrationOptions(
				registrationId,
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
		It("Login to account containing a programmatic name approved product", func() {
			var err error

			// begin-reauth
			const externalConfigFile2 = "../partner_center_sell_approved_user_v1.env"
			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile2)

			partnerCenterSellServiceOptions := &partnercentersellv1.PartnerCenterSellV1Options{}

			partnerCenterSellService, err = partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(partnerCenterSellServiceOptions)

			if err != nil {
				panic(err)
			}

			fmt.Println(*partnerCenterSellService)
			// end-reauth

			Expect(partnerCenterSellService).ToNot(BeNil())
		})
		It(`CreateOnboardingProduct request example`, func() {
			fmt.Println("\nCreateOnboardingProduct() result:")
			// begin-create_onboarding_product

			primaryContactModel := &partnercentersellv1.PrimaryContact{
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
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
			productId = *onboardingProduct.ID
			fmt.Println(string(b))

			// end-create_onboarding_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(onboardingProduct).ToNot(BeNil())
		})
		It(`GetOnboardingProduct request example`, func() {
			fmt.Println("\nGetOnboardingProduct() result:")
			// begin-get_onboarding_product

			getOnboardingProductOptions := partnerCenterSellService.NewGetOnboardingProductOptions(
				productId,
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
		})
		It(`UpdateOnboardingProduct request example`, func() {
			fmt.Println("\nUpdateOnboardingProduct() result:")
			// begin-update_onboarding_product

			onboardingProductPatchModel := &partnercentersellv1.OnboardingProductPatch{}
			onboardingProductPatchModelAsPatch, asPatchErr := onboardingProductPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateOnboardingProductOptions := partnerCenterSellService.NewUpdateOnboardingProductOptions(
				productId,
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
		})

		It(`CreateCatalogProduct request example`, func() {
			var err error

			// begin-common

			partnerCenterSellServiceOptions := &partnercentersellv1.PartnerCenterSellV1Options{}

			partnerCenterSellService, err = partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(partnerCenterSellServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(partnerCenterSellService).ToNot(BeNil())

			fmt.Println("\nCreateCatalogProduct() result:")
			// begin-create_catalog_product

			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
			}
			var randomInteger = strconv.Itoa(rand.IntN(1000))
			catalogProductName := fmt.Sprintf("gc-product-example-%s", randomInteger)

			createCatalogProductOptions := partnerCenterSellService.NewCreateCatalogProductOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductName,
				true,
				false,
				"service",
				[]string{"sample"},
				catalogProductProviderModel,
			)

			globalCatalogProduct, response, err := partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogProduct, "", "  ")
			catalogProductId = *globalCatalogProduct.ID
			fmt.Println(string(b))

			// end-create_catalog_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogProduct).ToNot(BeNil())
		})
		It(`GetCatalogProduct request example`, func() {
			fmt.Println("\nGetCatalogProduct() result:")
			// begin-get_catalog_product

			getCatalogProductOptions := partnerCenterSellService.NewGetCatalogProductOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductId,
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
		})
		It(`UpdateCatalogProduct request example`, func() {
			fmt.Println("\nUpdateCatalogProduct() result:")
			// begin-update_catalog_product

			globalCatalogProductPatchModel := &partnercentersellv1.GlobalCatalogProductPatch{}
			globalCatalogProductPatchModelAsPatch, asPatchErr := globalCatalogProductPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogProductOptions := partnerCenterSellService.NewUpdateCatalogProductOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductId,
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
		})
		It(`CreateCatalogPlan request example`, func() {
			fmt.Println("\nCreateCatalogPlan() result:")
			// begin-create_catalog_plan
			pricing := &partnercentersellv1.GlobalCatalogMetadataPricing{
				Type:   core.StringPtr("Paid"),
				Origin: core.StringPtr("pricing_catalog"),
			}

			ui := &partnercentersellv1.GlobalCatalogMetadataUI{}

			metaDataModel := &partnercentersellv1.GlobalCatalogPlanMetadata{
				RcCompatible: core.BoolPtr(true),
				Ui:           ui,
				Pricing:      pricing,
			}

			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
			}
			createCatalogPlanOptions := partnerCenterSellService.NewCreateCatalogPlanOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductId,
				"test-plan",
				true,
				false,
				"plan",
				[]string{"tag"},
				catalogProductProviderModel,
			)
			createCatalogPlanOptions.SetMetadata(metaDataModel)

			globalCatalogPlan, response, err := partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogPlan, "", "  ")
			catalogPlanId = *globalCatalogPlan.ID
			fmt.Println(string(b))

			// end-create_catalog_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogPlan).ToNot(BeNil())
		})
		It(`GetCatalogPlan request example`, func() {
			fmt.Println("\nGetCatalogPlan() result:")
			// begin-get_catalog_plan

			getCatalogPlanOptions := partnerCenterSellService.NewGetCatalogPlanOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductId,
				catalogPlanId,
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
		})
		It(`UpdateCatalogPlan request example`, func() {
			fmt.Println("\nUpdateCatalogPlan() result:")
			// begin-update_catalog_plan

			globalCatalogPlanPatchModel := &partnercentersellv1.GlobalCatalogPlanPatch{}
			globalCatalogPlanPatchModelAsPatch, asPatchErr := globalCatalogPlanPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogPlanOptions := partnerCenterSellService.NewUpdateCatalogPlanOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductId,
				catalogPlanId,
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
		})
		It(`CreateCatalogDeployment request example`, func() {
			fmt.Println("\nCreateCatalogDeployment() result:")
			// begin-create_catalog_deployment

			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
			}

			service := &partnercentersellv1.GlobalCatalogMetadataService{
				RcProvisionable: core.BoolPtr(true),
				IamCompatible:   core.BoolPtr(true),
			}

			metaDataModelDeployment := &partnercentersellv1.GlobalCatalogDeploymentMetadata{
				RcCompatible: core.BoolPtr(true),
				Service:      service,
			}

			createCatalogDeploymentOptions := partnerCenterSellService.NewCreateCatalogDeploymentOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductId,
				catalogPlanId,
				"sample-deployment",
				true,
				false,
				"deployment",
				[]string{"us"},
				catalogProductProviderModel,
			)

			createCatalogDeploymentOptions.SetMetadata(metaDataModelDeployment)

			globalCatalogDeployment, response, err := partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogDeployment, "", "  ")
			catalogDeploymentId = *globalCatalogDeployment.ID
			fmt.Println(string(b))

			// end-create_catalog_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogDeployment).ToNot(BeNil())
		})
		It(`GetCatalogDeployment request example`, func() {
			fmt.Println("\nGetCatalogDeployment() result:")
			// begin-get_catalog_deployment

			getCatalogDeploymentOptions := partnerCenterSellService.NewGetCatalogDeploymentOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductId,
				catalogPlanId,
				catalogDeploymentId,
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
		})
		It(`UpdateCatalogDeployment request example`, func() {
			fmt.Println("\nUpdateCatalogDeployment() result:")
			// begin-update_catalog_deployment

			globalCatalogDeploymentPatchModel := &partnercentersellv1.GlobalCatalogDeploymentPatch{}
			globalCatalogDeploymentPatchModelAsPatch, asPatchErr := globalCatalogDeploymentPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogDeploymentOptions := partnerCenterSellService.NewUpdateCatalogDeploymentOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductId,
				catalogPlanId,
				catalogDeploymentId,
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
		})
		It(`CreateIamRegistration request example`, func() {
			fmt.Println("\nCreateIamRegistration() result:")
			// begin-create_iam_registration

			createIamRegistrationOptions := partnerCenterSellService.NewCreateIamRegistrationOptions(
				productIdWithApprovedProgrammaticName,
			)

			createIamRegistrationOptions.SetName("iam-sample-for-sdk")

			iamServiceRegistration, response, err := partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iamServiceRegistration, "", "  ")
			iamServiceRegistrationId = *iamServiceRegistration.Name
			fmt.Println(string(b))

			// end-create_iam_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(iamServiceRegistration).ToNot(BeNil())
		})
		It(`UpdateIamRegistration request example`, func() {
			fmt.Println("\nUpdateIamRegistration() result:")
			// begin-update_iam_registration

			iamServiceRegistrationPatchModel := &partnercentersellv1.IamServiceRegistrationPatch{}
			iamServiceRegistrationPatchModelAsPatch, asPatchErr := iamServiceRegistrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateIamRegistrationOptions := partnerCenterSellService.NewUpdateIamRegistrationOptions(
				productIdWithApprovedProgrammaticName,
				iamServiceRegistrationId,
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
		})
		It(`GetIamRegistration request example`, func() {
			fmt.Println("\nGetIamRegistration() result:")
			// begin-get_iam_registration

			getIamRegistrationOptions := partnerCenterSellService.NewGetIamRegistrationOptions(
				productIdWithApprovedProgrammaticName,
				iamServiceRegistrationId,
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
		})
		It(`CreateResourceBroker request example`, func() {
			fmt.Println("\nCreateResourceBroker() result:")
			// begin-create_resource_broker
			var randomInteger = strconv.Itoa(rand.IntN(1000))
			brokerUrl := fmt.Sprintf("https://broker-url-for-my-service.com/%s", randomInteger)
			brokerUserName := fmt.Sprintf("petra_test_user_name_%s", randomInteger)
			brokerName := fmt.Sprintf("petra_test_%s", randomInteger)

			createResourceBrokerOptions := partnerCenterSellService.NewCreateResourceBrokerOptions(
				brokerUserName,
				"petra_test_user_pass",
				"basic",
				brokerName,
				brokerUrl,
				"provision_through",
			)

			broker, response, err := partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(broker, "", "  ")
			brokerId = *broker.ID
			fmt.Println(string(b))

			// end-create_resource_broker

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(broker).ToNot(BeNil())
		})
		It(`UpdateResourceBroker request example`, func() {
			fmt.Println("\nUpdateResourceBroker() result:")
			// begin-update_resource_broker

			brokerPatchModel := &partnercentersellv1.BrokerPatch{}
			brokerPatchModelAsPatch, asPatchErr := brokerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateResourceBrokerOptions := partnerCenterSellService.NewUpdateResourceBrokerOptions(
				brokerId,
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
				brokerId,
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

		It(`DeleteCatalogDeployment request example`, func() {
			// begin-delete_catalog_deployment

			deleteCatalogDeploymentOptions := partnerCenterSellService.NewDeleteCatalogDeploymentOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductId,
				catalogPlanId,
				catalogDeploymentId,
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
				productIdWithApprovedProgrammaticName,
				catalogProductId,
				catalogPlanId,
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

		It(`DeleteOnboardingProduct request example`, func() {
			// begin-delete_onboarding_product

			deleteOnboardingProductOptions := partnerCenterSellService.NewDeleteOnboardingProductOptions(
				productId,
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

		It(`DeleteIamRegistration request example`, func() {
			// begin-delete_iam_registration

			deleteIamRegistrationOptions := partnerCenterSellService.NewDeleteIamRegistrationOptions(
				productIdWithApprovedProgrammaticName,
				iamServiceRegistrationId,
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

		It(`DeleteCatalogProduct request example`, func() {
			// begin-delete_catalog_product

			deleteCatalogProductOptions := partnerCenterSellService.NewDeleteCatalogProductOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductId,
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
		It(`DeleteResourceBroker request example`, func() {
			// begin-delete_resource_broker

			deleteResourceBrokerOptions := partnerCenterSellService.NewDeleteResourceBrokerOptions(
				brokerId,
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

	It("Login to account containing badges", func() {
		var err error

		// begin-reauth
		const externalConfigFile2 = "../partner_center_sell_badged_user_v1.env"
		os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile2)

		partnerCenterSellServiceOptions := &partnercentersellv1.PartnerCenterSellV1Options{}

		partnerCenterSellService, err = partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(partnerCenterSellServiceOptions)

		if err != nil {
			panic(err)
		}

		// end-reauth

		Expect(err).To(BeNil())
		Expect(partnerCenterSellService).ToNot(BeNil())

		core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
		partnerCenterSellService.EnableRetries(4, 30*time.Second)
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
		badgeId = *productBadgeCollection.ProductBadges[0].ID
		fmt.Println(badgeId)

	})
	It(`GetProductBadge request example`, func() {
		fmt.Println("\nGetProductBadge() result:")
		// begin-get_product_badge

		getProductBadgeOptions := partnerCenterSellService.NewGetProductBadgeOptions(
			CreateMockUUID(badgeId),
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
	It(`ListBadges request example`, func() {
		fmt.Println("\nListBadges() result:")
		// begin-list_badges

		listBadgesOptions := partnerCenterSellService.NewListBadgesOptions()

		cloudBadges, response, err := partnerCenterSellService.ListBadges(listBadgesOptions)
		if err != nil {
			panic(err)
		}
		b, _ := json.MarshalIndent(cloudBadges, "", "  ")
		fmt.Println(string(b))
		// end-list_badges

		Expect(err).To(BeNil())
		Expect(response.StatusCode).To(Equal(200))
		Expect(cloudBadges).ToNot(BeNil())
	})
	It(`GetBadge request example`, func() {
		fmt.Println("\nGetBadge() result:")
		// begin-get_badge

		getBadgeOptions := partnerCenterSellService.NewGetBadgeOptions(
			CreateMockUUID(badgeId),
		)

		cloudBadge, response, err := partnerCenterSellService.GetBadge(getBadgeOptions)
		if err != nil {
			panic(err)
		}
		b, _ := json.MarshalIndent(cloudBadge, "", "  ")
		fmt.Println(string(b))

		// end-get_badge

		Expect(err).To(BeNil())
		Expect(response.StatusCode).To(Equal(200))
		Expect(cloudBadge).ToNot(BeNil())
	})
})
