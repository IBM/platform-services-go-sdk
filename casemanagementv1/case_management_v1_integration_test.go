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

package casemanagementv1_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/casemanagementv1"
	common "github.com/IBM/platform-services-go-sdk/common"
)

const externalConfigFile = "../case_management.env"

var (
	service *casemanagementv1.CaseManagementV1
	err     error

	configLoaded bool = false

	caseNumber   string
	commentValue = "Test comment"

	offeringType, _    = service.NewOfferingType(casemanagementv1.OfferingType_Group_CrnServiceName, "cloud-object-storage")
	offeringPayload, _ = service.NewOffering("Cloud Object Storage", offeringType)

	resourcePayload = []casemanagementv1.ResourcePayload{casemanagementv1.ResourcePayload{
		Crn: core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/19c52e57800c4d8bb9aefc66b3e49755:61848e72-6ba6-415e-84e2-91f3915e194d::"),
	}}

	watchlistPayload = casemanagementv1.Watchlist{
		Watchlist: []casemanagementv1.User{
			casemanagementv1.User{
				Realm:  core.StringPtr("IBMid"),
				UserID: core.StringPtr("ashwini.pc@ibm.com"),
			},
			casemanagementv1.User{
				Realm:  core.StringPtr("IBMid"),
				UserID: core.StringPtr("bqegarci@us.ibm.com"),
			},
		},
	}
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe("Case Management - Integration Tests", func() {
	It("Successfully load the configuration", func() {
		if _, fileErr := os.Stat(externalConfigFile); fileErr == nil {
			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, _ := core.GetServiceProperties(casemanagementv1.DefaultServiceName)
			if len(config) > 0 {
				configLoaded = true
			}
		}

		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}
	})

	It(`Successfully created CaseManagementV1 service instance`, func() {
		shouldSkipTest()

		service, err = casemanagementv1.NewCaseManagementV1UsingExternalConfig(
			&casemanagementv1.CaseManagementV1Options{},
		)

		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())

		core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags)))
		service.EnableRetries(4, 30*time.Second)

		// Set client timeout.
		service.Service.Client.Timeout = 2 * time.Minute

		fmt.Fprintf(GinkgoWriter, "\nService URL: %s\n", service.Service.GetServiceURL())
	})

	Describe("Create a case", func() {
		var options *casemanagementv1.CreateCaseOptions
		BeforeEach(func() {
			options = service.NewCreateCaseOptions("technical", "Test case for Go SDK", "Test case for Go SDK")
			options.SetSeverity(4)
			options.SetOffering(offeringPayload)
		})

		It("Successfully created a technical case", func() {
			shouldSkipTest()

			ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancelFunc()

			result, detailedResponse, err := service.CreateCaseWithContext(ctx, options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.Number).To(Not(BeNil()))
			Expect(*result.ShortDescription).To(Equal(*options.Subject))
			Expect(*result.Description).To(Equal(*options.Description))
			Expect(int64(*result.Severity)).To(Equal(*options.Severity))

			caseNumber = *result.Number

			fmt.Fprintf(GinkgoWriter, "\nCase number: %s\n", caseNumber)
		})

		It("Bad payload used to create a case", func() {
			shouldSkipTest()
			options.SetType("invalid_type")
			options.Severity = nil
			options.Offering = nil

			ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancelFunc()

			_, detailedResponse, err := service.CreateCaseWithContext(ctx, options)
			Expect(err).To(Not(BeNil()))
			Expect(detailedResponse.StatusCode).To(Not(Equal(200)))
		})
	})

	Describe("Get Cases", func() {
		var options *casemanagementv1.GetCasesOptions

		BeforeEach(func() {
			options = service.NewGetCasesOptions()
		})

		It("Successfully got cases with default params", func() {
			shouldSkipTest()

			result, detailedResponse, err := service.GetCases(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.TotalCount).To(Not(BeNil()))
			Expect(*result.First).To(Not(BeNil()))
			Expect(*result.Next).To(Not(BeNil()))
			Expect(*result.Last).To(Not(BeNil()))
			Expect(result.Cases).To(Not(BeNil()))
		})

		It("Successfully got cases with non-default params", func() {
			shouldSkipTest()

			options.SetOffset(10)
			options.SetLimit(20)
			options.SetFields([]string{
				casemanagementv1.GetCasesOptions_Fields_Number,
				casemanagementv1.GetCasesOptions_Fields_Comments,
				casemanagementv1.GetCasesOptions_Fields_CreatedAt,
			})

			result, detailedResponse, err := service.GetCases(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.TotalCount).To(Not(BeNil()))
			Expect(*result.First).To(Not(BeNil()))
			Expect(*result.Next).To(Not(BeNil()))
			Expect(*result.Last).To(Not(BeNil()))
			Expect(result.Cases).To(Not(BeNil()))

			testCase := result.Cases[0]
			Expect(testCase).To(Not(BeNil()))
			Expect(testCase.Number).To(Not(BeNil()))
			Expect(testCase.Comments).To(Not(BeNil()))
			Expect(testCase.CreatedAt).To(Not(BeNil()))

			// extra properties should be excluded in the response
			Expect(testCase.Severity).To(BeNil())
			Expect(testCase.Contact).To(BeNil())
		})

		It("Failed to get cases with bad params", func() {
			shouldSkipTest()

			options.SetFields([]string{"invalid_fields"})

			_, detailedResponse, err := service.GetCases(options)
			Expect(err).To(Not(BeNil()))
			Expect(detailedResponse.StatusCode).To(Not(Equal(200)))
		})
	})

	Describe("Get a specific case", func() {
		var options *casemanagementv1.GetCaseOptions

		BeforeEach(func() {
			options = service.NewGetCaseOptions(caseNumber)
		})

		It("Successfully got a case with default params", func() {
			shouldSkipTest()

			result, detailedResponse, err := service.GetCase(options)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.Number).To(Equal(caseNumber))
		})

		It("Successfully got a case with field filtering", func() {
			shouldSkipTest()

			options.SetFields([]string{"number", "severity"})
			result, detailedResponse, err := service.GetCase(options)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.Number).To(Equal(caseNumber))
			Expect(result.Severity).To(Not(BeNil()))
			Expect(result.Contact).To(BeNil())
		})

		It("Failed to get a case with bad params", func() {
			shouldSkipTest()

			options.SetFields([]string{"invalid_field"})
			_, detailedResponse, err := service.GetCase(options)
			Expect(err).To(Not(BeNil()))
			Expect(detailedResponse.StatusCode).To(Not(Equal(200)))
		})
	})

	Describe("Add comment", func() {
		var options *casemanagementv1.AddCommentOptions

		BeforeEach(func() {
			options = service.NewAddCommentOptions(caseNumber, commentValue)
		})

		It("Successfully added a comment to a case", func() {
			shouldSkipTest()
			result, detailedResponse, err := service.AddComment(options)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.Value).To(Equal(commentValue))
			Expect(result.AddedAt).To(Not(BeNil()))
			Expect(result.AddedBy).To(Not(BeNil()))
		})
	})

	Describe("Add watchlist", func() {
		var options *casemanagementv1.AddWatchlistOptions

		BeforeEach(func() {
			options = service.NewAddWatchlistOptions(caseNumber)
			options.SetWatchlist(watchlistPayload.Watchlist)
		})

		It("Successfully added users to case watchlist", func() {
			shouldSkipTest()
			result, detailedResponse, err := service.AddWatchlist(options)

			Expect(err).To((BeNil()))
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "AddWatchlist() result:\n%s\n", common.ToJSON(result))
			Expect(len(result.Added)).To(Equal(len(watchlistPayload.Watchlist)))
		})
	})

	Describe("Remove watchlist", func() {
		var options *casemanagementv1.RemoveWatchlistOptions
		BeforeEach(func() {
			options = service.NewRemoveWatchlistOptions(caseNumber)
			options.SetWatchlist(watchlistPayload.Watchlist)
		})

		It("Successfully removed users from case watchlist", func() {
			shouldSkipTest()
			_, detailedResponse, err := service.RemoveWatchlist(options)

			Expect(err).To((BeNil()))
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})
	})

	Describe("Update status", func() {

		It("Succefully resolve a case", func() {
			shouldSkipTest()
			resolvePayload, _ := service.NewResolvePayload(casemanagementv1.ResolvePayload_Action_Resolve, 1)
			options := service.NewUpdateCaseStatusOptions(caseNumber, resolvePayload)

			result, detailedResponse, err := service.UpdateCaseStatus(options)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.Status).To(Equal("Resolved"))
		})

		It("Succefully unresolve a case", func() {
			shouldSkipTest()
			unresolvePayload, _ := service.NewUnresolvePayload(casemanagementv1.UnresolvePayload_Action_Unresolve, "Test unresolve")
			options := service.NewUpdateCaseStatusOptions(caseNumber, unresolvePayload)

			result, detailedResponse, err := service.UpdateCaseStatus(options)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.Status).To(Equal("In Progress"))
		})
	})

	Describe("Modify attachments", func() {
		var fileID string

		It("Successfully uploaded file", func() {
			shouldSkipTest()
			fileInput, _ := service.NewFileWithMetadata(ioutil.NopCloser(strings.NewReader("hello world")))
			fileInput.Filename = core.StringPtr("GO SDK test file.png")
			fileInput.ContentType = core.StringPtr("application/octet-stream")

			filePayload := []casemanagementv1.FileWithMetadata{*fileInput}
			options := service.NewUploadFileOptions(caseNumber, filePayload)

			result, detailedResponse, err := service.UploadFile(options)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Not(BeNil()))
			Expect(*result.Filename).To(Equal(*fileInput.Filename))

			// store file id so that we could remove it in the next test
			fileID = *result.ID
		})

		It("Successfully deleted file", func() {
			shouldSkipTest()

			if fileID == "" {
				Skip("Case does not have target file to remove. Skipping ....")
			}

			options := service.NewDeleteFileOptions(caseNumber, fileID)

			_, detailedResponse, err := service.DeleteFile(options)

			Expect(err).To((BeNil()))
			Expect(detailedResponse.StatusCode).To((Equal(200)))
		})
	})

	Describe("Add Resource", func() {
		It("Successfully added a resource", func() {
			shouldSkipTest()
			crn := *resourcePayload[0].Crn
			options := service.NewAddResourceOptions(caseNumber)
			options.SetCrn(crn)

			result, detailedResponse, err := service.AddResource(options)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To((Equal(200)))
			Expect(*result.Crn).To(Equal(crn))
		})
	})
})
