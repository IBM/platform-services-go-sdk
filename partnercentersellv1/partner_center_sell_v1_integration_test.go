// +build integration

/**
 * (C) Copyright IBM Corp. 2022.
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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the partnercentersellv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PartnerCenterSellV1 Integration Tests`, func() {
	const externalConfigFile = "../partner_center_sell_v1.env"

	var (
		err          error
		partnerCenterSellService *partnercentersellv1.PartnerCenterSellV1
		serviceURL   string
		config       map[string]string
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
			config, err = core.GetServiceProperties(partnercentersellv1.DefaultServiceName)
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
			partnerCenterSellServiceOptions := &partnercentersellv1.PartnerCenterSellV1Options{}

			partnerCenterSellService, err = partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(partnerCenterSellServiceOptions)
			Expect(err).To(BeNil())
			Expect(partnerCenterSellService).ToNot(BeNil())
			Expect(partnerCenterSellService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			partnerCenterSellService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListProducts - List products`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProducts(listProductsOptions *ListProductsOptions)`, func() {
			listProductsOptions := &partnercentersellv1.ListProductsOptions{
			}

			listProductsResponse, response, err := partnerCenterSellService.ListProducts(listProductsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listProductsResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateProduct - Create product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProduct(createProductOptions *CreateProductOptions)`, func() {
			createProductOptions := &partnercentersellv1.CreateProductOptions{
				ProductName: core.StringPtr("testString"),
				TaxAssessment: core.StringPtr("SOFTWARE"),
				ProductType: core.StringPtr("SOFTWARE"),
				MaterialAgreement: core.BoolPtr(true),
			}

			productDetails, response, err := partnerCenterSellService.CreateProduct(createProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productDetails).ToNot(BeNil())
		})
	})

	Describe(`GetProduct - Get product details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProduct(getProductOptions *GetProductOptions)`, func() {
			getProductOptions := &partnercentersellv1.GetProductOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			productDetails, response, err := partnerCenterSellService.GetProduct(getProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productDetails).ToNot(BeNil())
		})
	})

	Describe(`UpdateProduct - Update product details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProduct(updateProductOptions *UpdateProductOptions)`, func() {
			updateProductOptions := &partnercentersellv1.UpdateProductOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				MaterialAgreement: core.BoolPtr(true),
				ProductName: core.StringPtr("testString"),
				TaxAssessment: core.StringPtr("SOFTWARE"),
			}

			productDetails, response, err := partnerCenterSellService.UpdateProduct(updateProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productDetails).ToNot(BeNil())
		})
	})

	Describe(`PublishProduct - Publish an approved product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PublishProduct(publishProductOptions *PublishProductOptions)`, func() {
			publishProductOptions := &partnercentersellv1.PublishProductOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			productDetails, response, err := partnerCenterSellService.PublishProduct(publishProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productDetails).ToNot(BeNil())
		})
	})

	Describe(`SuspendProduct - Suspend a published product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SuspendProduct(suspendProductOptions *SuspendProductOptions)`, func() {
			suspendProductOptions := &partnercentersellv1.SuspendProductOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				Reason: core.StringPtr("testString"),
			}

			productDetails, response, err := partnerCenterSellService.SuspendProduct(suspendProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productDetails).ToNot(BeNil())
		})
	})

	Describe(`DeprecateProduct - Deprecate a published product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeprecateProduct(deprecateProductOptions *DeprecateProductOptions)`, func() {
			deprecateProductOptions := &partnercentersellv1.DeprecateProductOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				Reason: core.StringPtr("testString"),
			}

			productDetails, response, err := partnerCenterSellService.DeprecateProduct(deprecateProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productDetails).ToNot(BeNil())
		})
	})

	Describe(`RestoreProduct - Restore a deprecated product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestoreProduct(restoreProductOptions *RestoreProductOptions)`, func() {
			restoreProductOptions := &partnercentersellv1.RestoreProductOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				Reason: core.StringPtr("testString"),
			}

			productDetails, response, err := partnerCenterSellService.RestoreProduct(restoreProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productDetails).ToNot(BeNil())
		})
	})

	Describe(`ListBadges - List badges`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBadges(listBadgesOptions *ListBadgesOptions)`, func() {
			listBadgesOptions := &partnercentersellv1.ListBadgesOptions{
			}

			cloudBadge, response, err := partnerCenterSellService.ListBadges(listBadgesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudBadge).ToNot(BeNil())
		})
	})

	Describe(`GetBadge - Get badge`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBadge(getBadgeOptions *GetBadgeOptions)`, func() {
			getBadgeOptions := &partnercentersellv1.GetBadgeOptions{
				BadgeID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			cloudBadge, response, err := partnerCenterSellService.GetBadge(getBadgeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudBadge).ToNot(BeNil())
		})
	})

	Describe(`GetCatalog - View a product's catalog data`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {
			getCatalogOptions := &partnercentersellv1.GetCatalogOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			catalogListingDetails, response, err := partnerCenterSellService.GetCatalog(getCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogListingDetails).ToNot(BeNil())
		})
	})

	Describe(`UpdateCatalog - Update a product's catalog data`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateCatalog(updateCatalogOptions *UpdateCatalogOptions)`, func() {
			highlightSectionInputModel := &partnercentersellv1.HighlightSectionInput{
				Description: core.StringPtr("testString"),
				Title: core.StringPtr("testString"),
			}

			mediaSectionInputModel := &partnercentersellv1.MediaSectionInput{
				Caption: core.StringPtr("testString"),
				Thumbnail: core.StringPtr("testString"),
				Type: core.StringPtr("image"),
				URL: core.StringPtr("testString"),
			}

			updateCatalogOptions := &partnercentersellv1.UpdateCatalogOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				CatalogID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				IconURL: core.StringPtr("testString"),
				Keywords: []string{"testString"},
				PricingModel: core.StringPtr("free"),
				Category: core.StringPtr("testString"),
				ProviderType: []string{"ibm_community"},
				Label: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Provider: core.StringPtr("testString"),
				Tags: []string{"testString"},
				DocumentationURL: core.StringPtr("testString"),
				Highlights: []partnercentersellv1.HighlightSectionInput{*highlightSectionInputModel},
				LongDescription: core.StringPtr("testString"),
				Media: []partnercentersellv1.MediaSectionInput{*mediaSectionInputModel},
			}

			catalogListingDetails, response, err := partnerCenterSellService.UpdateCatalog(updateCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogListingDetails).ToNot(BeNil())
		})
	})

	Describe(`RequestCatalogApproval - Request a catalog listing approval`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RequestCatalogApproval(requestCatalogApprovalOptions *RequestCatalogApprovalOptions)`, func() {
			requestCatalogApprovalOptions := &partnercentersellv1.RequestCatalogApprovalOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			resource, response, err := partnerCenterSellService.RequestCatalogApproval(requestCatalogApprovalOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resource).ToNot(BeNil())
		})
	})

	Describe(`ListPlans - List pricing plans that are connected to a product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPlans(listPlansOptions *ListPlansOptions)`, func() {
			listPlansOptions := &partnercentersellv1.ListPlansOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			listPlansResponse, response, err := partnerCenterSellService.ListPlans(listPlansOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listPlansResponse).ToNot(BeNil())
		})
	})

	Describe(`CreatePlan - Create a pricing plan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePlan(createPlanOptions *CreatePlanOptions)`, func() {
			createPlanOptions := &partnercentersellv1.CreatePlanOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				Description: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				Type: core.StringPtr("byol"),
				URL: core.StringPtr("testString"),
			}

			createPlanResponse, response, err := partnerCenterSellService.CreatePlan(createPlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createPlanResponse).ToNot(BeNil())
		})
	})

	Describe(`GetPlan - Get pricing plan by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPlan(getPlanOptions *GetPlanOptions)`, func() {
			getPlanOptions := &partnercentersellv1.GetPlanOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				PricingPlanID: core.StringPtr("testString"),
			}

			license, response, err := partnerCenterSellService.GetPlan(getPlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(license).ToNot(BeNil())
		})
	})

	Describe(`UpdatePlan - Update a pricing plan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePlan(updatePlanOptions *UpdatePlanOptions)`, func() {
			updatePlanOptions := &partnercentersellv1.UpdatePlanOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				PricingPlanID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
				Type: core.StringPtr("byol"),
				URL: core.StringPtr("testString"),
			}

			createPlanResponse, response, err := partnerCenterSellService.UpdatePlan(updatePlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createPlanResponse).ToNot(BeNil())
		})
	})

	Describe(`GetSupport - Get product support details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSupport(getSupportOptions *GetSupportOptions)`, func() {
			getSupportOptions := &partnercentersellv1.GetSupportOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			support, response, err := partnerCenterSellService.GetSupport(getSupportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(support).ToNot(BeNil())
		})
	})

	Describe(`UpdateSupport - Update product support details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSupport(updateSupportOptions *UpdateSupportOptions)`, func() {
			escalationContactsUpdateModel := &partnercentersellv1.EscalationContactsUpdate{
				Email: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			supportDetailsAvailabilityTimesModel := &partnercentersellv1.SupportDetailsAvailabilityTimes{
				Day: core.Int64Ptr(int64(1)),
				EndTime: core.StringPtr("19:30"),
				StartTime: core.StringPtr("10:30"),
			}

			supportDetailsAvailabilityModel := &partnercentersellv1.SupportDetailsAvailability{
				AlwaysAvailable: core.BoolPtr(true),
				Times: []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel},
				Timezone: core.StringPtr("America/Los_Angeles"),
			}

			supportResponseTimesModel := &partnercentersellv1.SupportResponseTimes{
				Type: core.StringPtr("hour"),
				Value: core.Int64Ptr(int64(2)),
			}

			supportDetailsModel := &partnercentersellv1.SupportDetails{
				Availability: supportDetailsAvailabilityModel,
				Contact: core.StringPtr("testString"),
				ResponseWaitTime: supportResponseTimesModel,
				Type: core.StringPtr("email"),
			}

			supportEscalationTimesModel := &partnercentersellv1.SupportEscalationTimes{
				Type: core.StringPtr("hour"),
				Value: core.Int64Ptr(int64(2)),
			}

			supportEscalationModel := &partnercentersellv1.SupportEscalation{
				Contact: core.StringPtr("testString"),
				EscalationWaitTime: supportEscalationTimesModel,
				ResponseWaitTime: supportResponseTimesModel,
			}

			updateSupportOptions := &partnercentersellv1.UpdateSupportOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				EscalationContacts: []partnercentersellv1.EscalationContactsUpdate{*escalationContactsUpdateModel},
				Locations: []string{"US"},
				SupportDetails: []partnercentersellv1.SupportDetails{*supportDetailsModel},
				SupportEscalation: supportEscalationModel,
				SupportType: core.StringPtr("third-party"),
				URL: core.StringPtr("https://my-company.com/support"),
			}

			support, response, err := partnerCenterSellService.UpdateSupport(updateSupportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(support).ToNot(BeNil())
		})
	})

	Describe(`ListSupportChangeRequests - List all change requests related to a given product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSupportChangeRequests(listSupportChangeRequestsOptions *ListSupportChangeRequestsOptions)`, func() {
			listSupportChangeRequestsOptions := &partnercentersellv1.ListSupportChangeRequestsOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			listSupportChangeRequestsResponse, response, err := partnerCenterSellService.ListSupportChangeRequests(listSupportChangeRequestsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listSupportChangeRequestsResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateSupportChangeRequest - Update support data of an already approved or published product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSupportChangeRequest(createSupportChangeRequestOptions *CreateSupportChangeRequestOptions)`, func() {
			supportDetailsAvailabilityTimesModel := &partnercentersellv1.SupportDetailsAvailabilityTimes{
				Day: core.Int64Ptr(int64(1)),
				EndTime: core.StringPtr("19:30"),
				StartTime: core.StringPtr("10:30"),
			}

			supportDetailsAvailabilityModel := &partnercentersellv1.SupportDetailsAvailability{
				AlwaysAvailable: core.BoolPtr(true),
				Times: []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel},
				Timezone: core.StringPtr("America/Los_Angeles"),
			}

			supportResponseTimesModel := &partnercentersellv1.SupportResponseTimes{
				Type: core.StringPtr("hour"),
				Value: core.Int64Ptr(int64(2)),
			}

			supportDetailsModel := &partnercentersellv1.SupportDetails{
				Availability: supportDetailsAvailabilityModel,
				Contact: core.StringPtr("testString"),
				ResponseWaitTime: supportResponseTimesModel,
				Type: core.StringPtr("email"),
			}

			supportEscalationTimesModel := &partnercentersellv1.SupportEscalationTimes{
				Type: core.StringPtr("hour"),
				Value: core.Int64Ptr(int64(2)),
			}

			supportEscalationModel := &partnercentersellv1.SupportEscalation{
				Contact: core.StringPtr("testString"),
				EscalationWaitTime: supportEscalationTimesModel,
				ResponseWaitTime: supportResponseTimesModel,
			}

			supportModel := &partnercentersellv1.Support{
				Locations: []string{"US"},
				Process: core.StringPtr("testString"),
				ProcessI18n: map[string]interface{}{"anyKey": "anyValue"},
				SupportDetails: []partnercentersellv1.SupportDetails{*supportDetailsModel},
				SupportEscalation: supportEscalationModel,
				SupportType: core.StringPtr("third-party"),
				URL: core.StringPtr("https://my-company.com/support"),
			}

			createSupportChangeRequestOptions := &partnercentersellv1.CreateSupportChangeRequestOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				Change: supportModel,
			}

			productDetails, response, err := partnerCenterSellService.CreateSupportChangeRequest(createSupportChangeRequestOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productDetails).ToNot(BeNil())
		})
	})

	Describe(`GetSupportChangeRequest - Get a change request related to a given product by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSupportChangeRequest(getSupportChangeRequestOptions *GetSupportChangeRequestOptions)`, func() {
			getSupportChangeRequestOptions := &partnercentersellv1.GetSupportChangeRequestOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				ChangeRequestID: core.StringPtr("testString"),
			}

			changeRequest, response, err := partnerCenterSellService.GetSupportChangeRequest(getSupportChangeRequestOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(changeRequest).ToNot(BeNil())
		})
	})

	Describe(`UpdateSupportChangeRequest - Update an already created change request`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSupportChangeRequest(updateSupportChangeRequestOptions *UpdateSupportChangeRequestOptions)`, func() {
			supportDetailsAvailabilityTimesModel := &partnercentersellv1.SupportDetailsAvailabilityTimes{
				Day: core.Int64Ptr(int64(1)),
				EndTime: core.StringPtr("19:30"),
				StartTime: core.StringPtr("10:30"),
			}

			supportDetailsAvailabilityModel := &partnercentersellv1.SupportDetailsAvailability{
				AlwaysAvailable: core.BoolPtr(true),
				Times: []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel},
				Timezone: core.StringPtr("America/Los_Angeles"),
			}

			supportResponseTimesModel := &partnercentersellv1.SupportResponseTimes{
				Type: core.StringPtr("hour"),
				Value: core.Int64Ptr(int64(2)),
			}

			supportDetailsModel := &partnercentersellv1.SupportDetails{
				Availability: supportDetailsAvailabilityModel,
				Contact: core.StringPtr("testString"),
				ResponseWaitTime: supportResponseTimesModel,
				Type: core.StringPtr("email"),
			}

			supportEscalationTimesModel := &partnercentersellv1.SupportEscalationTimes{
				Type: core.StringPtr("hour"),
				Value: core.Int64Ptr(int64(2)),
			}

			supportEscalationModel := &partnercentersellv1.SupportEscalation{
				Contact: core.StringPtr("testString"),
				EscalationWaitTime: supportEscalationTimesModel,
				ResponseWaitTime: supportResponseTimesModel,
			}

			supportModel := &partnercentersellv1.Support{
				Locations: []string{"US"},
				Process: core.StringPtr("testString"),
				ProcessI18n: map[string]interface{}{"anyKey": "anyValue"},
				SupportDetails: []partnercentersellv1.SupportDetails{*supportDetailsModel},
				SupportEscalation: supportEscalationModel,
				SupportType: core.StringPtr("third-party"),
				URL: core.StringPtr("https://my-company.com/support"),
			}

			updateSupportChangeRequestOptions := &partnercentersellv1.UpdateSupportChangeRequestOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				ChangeRequestID: core.StringPtr("testString"),
				Change: supportModel,
			}

			productDetails, response, err := partnerCenterSellService.UpdateSupportChangeRequest(updateSupportChangeRequestOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productDetails).ToNot(BeNil())
		})
	})

	Describe(`ListSupportChangeRequestReviews - List review events related to a change request`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptions *ListSupportChangeRequestReviewsOptions)`, func() {
			listSupportChangeRequestReviewsOptions := &partnercentersellv1.ListSupportChangeRequestReviewsOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				ChangeRequestID: core.StringPtr("testString"),
			}

			resource, response, err := partnerCenterSellService.ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resource).ToNot(BeNil())
		})
	})

	Describe(`RequestSupportChangeRequestReview - Update an already created change request`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptions *RequestSupportChangeRequestReviewOptions)`, func() {
			requestSupportChangeRequestReviewOptions := &partnercentersellv1.RequestSupportChangeRequestReviewOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				ChangeRequestID: core.StringPtr("testString"),
			}

			resource, response, err := partnerCenterSellService.RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resource).ToNot(BeNil())
		})
	})

	Describe(`MergeSupportChangeRequest - Merge the approved changeset to the published product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`MergeSupportChangeRequest(mergeSupportChangeRequestOptions *MergeSupportChangeRequestOptions)`, func() {
			mergeSupportChangeRequestOptions := &partnercentersellv1.MergeSupportChangeRequestOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				ChangeRequestID: core.StringPtr("testString"),
			}

			productDetails, response, err := partnerCenterSellService.MergeSupportChangeRequest(mergeSupportChangeRequestOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productDetails).ToNot(BeNil())
		})
	})

	Describe(`RequestSupportApproval - Request approval of support information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RequestSupportApproval(requestSupportApprovalOptions *RequestSupportApprovalOptions)`, func() {
			requestSupportApprovalOptions := &partnercentersellv1.RequestSupportApprovalOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			resource, response, err := partnerCenterSellService.RequestSupportApproval(requestSupportApprovalOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resource).ToNot(BeNil())
		})
	})

	Describe(`RequestProductApproval - Request approval to publish`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RequestProductApproval(requestProductApprovalOptions *RequestProductApprovalOptions)`, func() {
			requestProductApprovalOptions := &partnercentersellv1.RequestProductApprovalOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			resource, response, err := partnerCenterSellService.RequestProductApproval(requestProductApprovalOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resource).ToNot(BeNil())
		})
	})

	Describe(`ListProductApprovals - List approvals`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProductApprovals(listProductApprovalsOptions *ListProductApprovalsOptions)`, func() {
			listProductApprovalsOptions := &partnercentersellv1.ListProductApprovalsOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			listProductApprovalsResponse, response, err := partnerCenterSellService.ListProductApprovals(listProductApprovalsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listProductApprovalsResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteProduct - Delete a draft product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProduct(deleteProductOptions *DeleteProductOptions)`, func() {
			deleteProductOptions := &partnercentersellv1.DeleteProductOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			result, response, err := partnerCenterSellService.DeleteProduct(deleteProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`DeletePlan - Delete a pricing plan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePlan(deletePlanOptions *DeletePlanOptions)`, func() {
			deletePlanOptions := &partnercentersellv1.DeletePlanOptions{
				ProductID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				PricingPlanID: core.StringPtr("testString"),
			}

			createPlanResponse, response, err := partnerCenterSellService.DeletePlan(deletePlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createPlanResponse).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
