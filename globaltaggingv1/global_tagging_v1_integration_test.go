// +build integration

package globaltaggingv1_test

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

import (
	"github.com/IBM/go-sdk-core/v3/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/globaltaggingv1"
	"time"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Resource = globaltaggingv1.Resource
type Tag = globaltaggingv1.Tag

var _ = Describe("Global Search and Tagging - Search integration test", func() {

	err := godotenv.Load("../.ghostenv")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	crn := os.Getenv("GST_RESOURCE_CRN")
	tagName := fmt.Sprint("go-sdk-", time.Now().Unix())
	tagElem := Tag{
					Name: core.StringPtr(tagName),
				}

	// Create the authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("GST_IINTERNA_APIKEY"),
		URL: os.Getenv("GST_IAM_URL"),
	}

	options := &globaltaggingv1.GlobalTaggingV1Options {
		Authenticator: authenticator,
	}
	service, err := globaltaggingv1.NewGlobalTaggingV1(options)
	It(`Successfully created NewGlobalTaggingV1 service instance`, func() {
		Expect(err).To(BeNil())
	})

	err = service.SetServiceURL(os.Getenv("GST_TAGS_URL"))
	It(`Successfully change default service URL to point to` + os.Getenv("GST_TAGS_URL"), func() {
		Expect(err).To(BeNil())
	})

	Describe("Call GetAllTags", func() {

		// Construct an instance of the ListTags model
		listTagsModel := service.NewListTagsOptions()

			It("Successfully get all tags", func() {
			result, detailedResponse, err := service.ListTags(listTagsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.Items).NotTo(BeEmpty())
		})
	})

	Describe("Attach a tag", func() {

		// Construct an instance of the AttachTag model
		resource, _ := service.NewResource(crn)
		array := []Resource{*resource}
		attachTagOptions := service.NewAttachTagOptions(array)
		attachTagOptions.SetTagNames([]string{tagName})

		It("Successfully attach a tag", func() {
			result, detailedResponse, err := service.AttachTag(attachTagOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			for _, elem := range result.Results {
				Expect(*elem.IsError).To(Equal(false))
			}

			// checking that tag is attached
			listTagsModel := service.NewListTagsOptions()
			listTagsModel.SetAttachedTo(crn)
			resultCheck, detailedResponseCheck, errCheck := service.ListTags(listTagsModel)
			Expect(errCheck).To(BeNil())
			Expect(detailedResponseCheck.StatusCode).To(Equal(200))
			Expect(resultCheck.Items).NotTo(BeEmpty())
			Expect(resultCheck.Items).To(ContainElement(tagElem))
		})
	})

	Describe("Detach a tag", func() {

		// Construct an instance of the DetachTag model
		resource, _ := service.NewResource(crn)
		array := []Resource{*resource}
		detachTagOptions := service.NewDetachTagOptions(array)
		detachTagOptions.SetTagNames([]string{tagName})


		It("Successfully detached a tag", func() {
			result, detailedResponse, err := service.DetachTag(detachTagOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			for _, elem := range result.Results {
				Expect(*elem.IsError).To(Equal(false))
			}

			// checking that tag is detached
			listTagsModel := service.NewListTagsOptions()
			listTagsModel.SetAttachedTo(crn)
			resultCheck, detailedResponseCheck, errCheck := service.ListTags(listTagsModel)
			Expect(errCheck).To(BeNil())
			Expect(detailedResponseCheck.StatusCode).To(Equal(200))
			Expect(resultCheck.Items).NotTo(ContainElement(tagElem))
		})
	})

	Describe("Delete a tag", func() {

		// Construct an instance of the DeleteTag model
		deleteTagOptions := service.NewDeleteTagOptions(tagName)

		It("Successfully delete a tag", func() {
			result, detailedResponse, err := service.DeleteTag(deleteTagOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			for _, elem := range result.Results {
				Expect(*elem.IsError).To(Equal(false))
			}
		})
	})

})