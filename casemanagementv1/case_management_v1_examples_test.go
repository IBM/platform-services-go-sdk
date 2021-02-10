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

package casemanagementv1_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/casemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const externalConfigFile = "../case_management_v1.env"

var (
	caseManagementService *casemanagementv1.CaseManagementV1
	config                map[string]string
	configLoaded          bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`CaseManagementV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(casemanagementv1.DefaultServiceName)
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

			caseManagementServiceOptions := &casemanagementv1.CaseManagementV1Options{}

			caseManagementService, err = casemanagementv1.NewCaseManagementV1UsingExternalConfig(caseManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(caseManagementService).ToNot(BeNil())
		})
	})

	Describe(`CaseManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCases request example`, func() {
			// begin-getCases

			getCasesOptions := caseManagementService.NewGetCasesOptions()

			caseList, response, err := caseManagementService.GetCases(getCasesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(caseList, "", "  ")
			fmt.Println(string(b))

			// end-getCases

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(caseList).ToNot(BeNil())

		})
		It(`CreateCase request example`, func() {
			// begin-createCase

			createCaseOptions := caseManagementService.NewCreateCaseOptions(
				"technical",
				"testString",
				"testString",
			)

			caseVar, response, err := caseManagementService.CreateCase(createCaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(caseVar, "", "  ")
			fmt.Println(string(b))

			// end-createCase

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(caseVar).ToNot(BeNil())

		})
		It(`GetCase request example`, func() {
			// begin-getCase

			getCaseOptions := caseManagementService.NewGetCaseOptions(
				"testString",
			)

			caseVar, response, err := caseManagementService.GetCase(getCaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(caseVar, "", "  ")
			fmt.Println(string(b))

			// end-getCase

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(caseVar).ToNot(BeNil())

		})
		It(`UpdateCaseStatus request example`, func() {
			// begin-updateCaseStatus

			statusPayloadModel := &casemanagementv1.ResolvePayload{
				Action:         core.StringPtr("resolve"),
				Comment:        core.StringPtr("It was actually a mistake"),
				ResolutionCode: core.Int64Ptr(int64(1)),
			}

			updateCaseStatusOptions := caseManagementService.NewUpdateCaseStatusOptions(
				"testString",
				statusPayloadModel,
			)

			caseVar, response, err := caseManagementService.UpdateCaseStatus(updateCaseStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(caseVar, "", "  ")
			fmt.Println(string(b))

			// end-updateCaseStatus

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(caseVar).ToNot(BeNil())

		})
		It(`AddComment request example`, func() {
			// begin-addComment

			addCommentOptions := caseManagementService.NewAddCommentOptions(
				"testString",
				"This is a test comment",
			)

			comment, response, err := caseManagementService.AddComment(addCommentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(comment, "", "  ")
			fmt.Println(string(b))

			// end-addComment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(comment).ToNot(BeNil())

		})
		It(`AddWatchlist request example`, func() {
			// begin-addWatchlist

			addWatchlistOptions := caseManagementService.NewAddWatchlistOptions(
				"testString",
			)

			watchlistAddResponse, response, err := caseManagementService.AddWatchlist(addWatchlistOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(watchlistAddResponse, "", "  ")
			fmt.Println(string(b))

			// end-addWatchlist

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(watchlistAddResponse).ToNot(BeNil())

		})
		It(`AddResource request example`, func() {
			// begin-addResource

			addResourceOptions := caseManagementService.NewAddResourceOptions(
				"testString",
			)

			resource, response, err := caseManagementService.AddResource(addResourceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resource, "", "  ")
			fmt.Println(string(b))

			// end-addResource

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resource).ToNot(BeNil())

		})
		It(`UploadFile request example`, func() {
			// begin-uploadFile

			uploadFileOptions := caseManagementService.NewUploadFileOptions(
				"testString",
				[]casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt")}},
			)

			attachment, response, err := caseManagementService.UploadFile(uploadFileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachment, "", "  ")
			fmt.Println(string(b))

			// end-uploadFile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())

		})
		It(`DownloadFile request example`, func() {
			// begin-downloadFile

			downloadFileOptions := caseManagementService.NewDownloadFileOptions(
				"testString",
				"testString",
			)

			result, response, err := caseManagementService.DownloadFile(downloadFileOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil {
					panic(err)
				}
			}

			// end-downloadFile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`RemoveWatchlist request example`, func() {
			// begin-removeWatchlist

			removeWatchlistOptions := caseManagementService.NewRemoveWatchlistOptions(
				"testString",
			)

			watchlist, response, err := caseManagementService.RemoveWatchlist(removeWatchlistOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(watchlist, "", "  ")
			fmt.Println(string(b))

			// end-removeWatchlist

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(watchlist).ToNot(BeNil())

		})
		It(`DeleteFile request example`, func() {
			// begin-deleteFile

			deleteFileOptions := caseManagementService.NewDeleteFileOptions(
				"testString",
				"testString",
			)

			attachmentList, response, err := caseManagementService.DeleteFile(deleteFileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentList, "", "  ")
			fmt.Println(string(b))

			// end-deleteFile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentList).ToNot(BeNil())

		})
	})
})
