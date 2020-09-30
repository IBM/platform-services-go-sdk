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

package globaltaggingv1_test

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the globaltaggingv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`GlobalTaggingV1 Integration Tests`, func() {

	const externalConfigFile = "../global_tagging.env"

	var (
		err                  error
		globalTaggingService *globaltaggingv1.GlobalTaggingV1
		serviceURL           string
		config               map[string]string

		crn     string
		tagName string = fmt.Sprint("go-sdk-", time.Now().Unix())
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
			config, err = core.GetServiceProperties(globaltaggingv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			crn = config["RESOURCE_CRN"]
			if crn == "" {
				Skip("Unable to load RESOURCE_CRN configuration property, skipping tests")
			}

			fmt.Printf("\nService URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			globalTaggingServiceOptions := &globaltaggingv1.GlobalTaggingV1Options{}

			globalTaggingService, err = globaltaggingv1.NewGlobalTaggingV1UsingExternalConfig(globalTaggingServiceOptions)

			Expect(err).To(BeNil())
			Expect(globalTaggingService).ToNot(BeNil())
			Expect(globalTaggingService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`ListTags - Get all tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTags(listTagsOptions *ListTagsOptions)`, func() {

			listTagsOptions := &globaltaggingv1.ListTagsOptions{
				Offset: core.Int64Ptr(0),
				Limit:  core.Int64Ptr(1000),
			}

			tagList, response, err := globalTaggingService.ListTags(listTagsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tagList).ToNot(BeNil())
			// fmt.Printf("\nListTags(all) response:\n%s", toJson(tagList))

			Expect(tagList.Items).ToNot(BeEmpty())
		})
	})

	Describe(`AttachTag - Attach tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AttachTag(attachTagOptions *AttachTagOptions)`, func() {

			resourceModel := &globaltaggingv1.Resource{
				ResourceID: &crn,
			}

			attachTagOptions := &globaltaggingv1.AttachTagOptions{
				Resources: []globaltaggingv1.Resource{*resourceModel},
				TagNames:  []string{tagName},
			}

			tagResults, response, err := globalTaggingService.AttachTag(attachTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tagResults).ToNot(BeNil())
			// fmt.Printf("\nAttachTag() response:\n%s", toJson(tagResults))

			Expect(tagResults.Results).ToNot(BeEmpty())
			for _, elem := range tagResults.Results {
				Expect(*elem.IsError).To(Equal(false))
			}

			// Make sure the tag was in fact attached.
			tagNames := getTagNamesForResource(globalTaggingService, crn)
			// fmt.Print("\nResource now has these tags: ", tagNames)
			Expect(tagNames).ToNot(BeEmpty())
			Expect(tagNames).To(ContainElement(tagName))
		})
	})

	Describe(`DetachTag - Detach tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DetachTag(detachTagOptions *DetachTagOptions)`, func() {

			resourceModel := &globaltaggingv1.Resource{
				ResourceID: &crn,
			}

			detachTagOptions := &globaltaggingv1.DetachTagOptions{
				Resources: []globaltaggingv1.Resource{*resourceModel},
				TagNames:  []string{tagName},
			}

			tagResults, response, err := globalTaggingService.DetachTag(detachTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tagResults).ToNot(BeNil())
			// fmt.Printf("\nDetachTag() response:\n%s", toJson(tagResults))

			Expect(tagResults.Results).ToNot(BeEmpty())
			for _, elem := range tagResults.Results {
				Expect(*elem.IsError).To(Equal(false))
			}

			// Make sure the tag was in fact detached.
			tagNames := getTagNamesForResource(globalTaggingService, crn)
			// fmt.Print("\nResource now has these tags: ", tagNames)
			Expect(tagNames).ToNot(ContainElement(tagName))
		})
	})

	Describe(`DeleteTag - Delete an unused tag`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTag(deleteTagOptions *DeleteTagOptions)`, func() {

			deleteTagOptions := &globaltaggingv1.DeleteTagOptions{
				TagName: &tagName,
			}

			deleteTagResults, response, err := globalTaggingService.DeleteTag(deleteTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteTagResults).ToNot(BeNil())
			for _, elem := range deleteTagResults.Results {
				Expect(*elem.IsError).To(Equal(false))
			}
		})
	})

	Describe(`DeleteTagAll - Delete all unused tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTagAll(deleteTagAllOptions *DeleteTagAllOptions)`, func() {

			deleteTagAllOptions := &globaltaggingv1.DeleteTagAllOptions{}

			deleteTagsResult, response, err := globalTaggingService.DeleteTagAll(deleteTagAllOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteTagsResult).ToNot(BeNil())
			// fmt.Printf("\nDetachTagAll() response:\n%s", toJson(deleteTagsResult))
			// fmt.Printf("\nDeleteTagAll deleted %d unused tags.\n", *deleteTagsResult.TotalCount)
		})
	})
})

func getTagNamesForResource(service *globaltaggingv1.GlobalTaggingV1, resourceID string) []string {
	listTagsOptions := &globaltaggingv1.ListTagsOptions{
		AttachedTo: &resourceID,
	}
	tagList, response, err := service.ListTags(listTagsOptions)
	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))

	tagNames := []string{}
	for _, tag := range tagList.Items {
		tagNames = append(tagNames, *tag.Name)
	}

	return tagNames
}

func toJson(obj interface{}) string {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)

}
