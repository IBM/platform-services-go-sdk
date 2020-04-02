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
	"fmt"
	"os"
	"time"
	
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
)

type Resource = globaltaggingv1.Resource
type Tag = globaltaggingv1.Tag

var service *globaltaggingv1.GlobalTaggingV1
var configLoaded = false

var crn string

var tagName string = fmt.Sprint("go-sdk-", time.Now().Unix())

var tagElem = Tag{
	Name: core.StringPtr(tagName),
}

const externalConfigFile = "../ghost.env"

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe("Global Search and Tagging - Tagging integration test", func() {
	It("Successfully load the configuration", func() {
		err := godotenv.Load(externalConfigFile)
		if err == nil {
			crn = os.Getenv("GST_RESOURCE_CRN")
			if crn != "" {
				configLoaded = true
			}
		}
		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}
	})

	It("Successfully construct service", func() {
		shouldSkipTest()

		// Create the authenticator.
		authenticator := &core.IamAuthenticator{
			ApiKey: os.Getenv("GST_IINTERNA_APIKEY"),
			URL:    os.Getenv("GST_IAM_URL"),
		}

		options := &globaltaggingv1.GlobalTaggingV1Options{
			Authenticator: authenticator,
			URL:           os.Getenv("GST_TAGS_URL"),
		}
		var err error
		service, err = globaltaggingv1.NewGlobalTaggingV1(options)
		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())
	})

	Describe("Call GetAllTags", func() {

		It("Successfully get all tags", func() {
			shouldSkipTest()

			// Construct an instance of the ListTags model
			listTagsModel := service.NewListTagsOptions()

			var err error
			result, detailedResponse, err := service.ListTags(listTagsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.Items).NotTo(BeEmpty())
		})
	})

	Describe("Attach a tag", func() {

		It("Successfully attach a tag", func() {
			shouldSkipTest()

			// Construct an instance of the Resource model
			resource, _ := service.NewResource(crn)
			array := []Resource{*resource}
			attachTagOptions := service.NewAttachTagOptions(array)
			attachTagOptions.SetTagNames([]string{tagName})

			var err error
			result, detailedResponse, err := service.AttachTag(attachTagOptions)

			//fmt.Println("attach tag response", detailedResponse)

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

		It("Successfully detached a tag", func() {
			shouldSkipTest()

			// Construct an instance of the DetachTag model
			resource, _ := service.NewResource(crn)
			array := []Resource{*resource}
			detachTagOptions := service.NewDetachTagOptions(array)
			detachTagOptions.SetTagNames([]string{tagName})

			var err error
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

		It("Successfully delete a tag", func() {
			shouldSkipTest()

			// Construct an instance of the DeleteTag model
			deleteTagOptions := service.NewDeleteTagOptions(tagName)

			var err error
			result, detailedResponse, err := service.DeleteTag(deleteTagOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			for _, elem := range result.Results {
				Expect(*elem.IsError).To(Equal(false))
			}
		})
	})
})
