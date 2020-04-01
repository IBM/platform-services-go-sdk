// +build integration

package globalsearchv2_test

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
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var service *globalsearchv2.GlobalSearchV2
var configLoaded = false
const externalConfigFile = "../.ghostenv"

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe("Global Search and Tagging - Search integration test", func() {
	It("Successfully load the configuration", func() {
		err := godotenv.Load(externalConfigFile)
		if err == nil {
			configLoaded = true
		} else {
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

		options := &globalsearchv2.GlobalSearchV2Options{
			Authenticator: authenticator,
			URL:           os.Getenv("GST_API_URL"),
		}
		service, err := globalsearchv2.NewGlobalSearchV2(options)
		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())
	})

	Describe("Call Search v3 api with query 'name:gst-sdk*' all fields", func() {

		// Construct an instance of the SearchOptions model
		searchOptionsModel := service.NewSearchOptions()
		searchOptionsModel.SetQuery(os.Getenv("GST_QUERY"))
		searchOptionsModel.SetFields([]string{"*"})

		It("Successfully list all resources", func() {
			shouldSkipTest()
			result, detailedResponse, err := service.Search(searchOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.Items).To(HaveLen(2))
			for _, elem := range result.Items {
				Expect(elem.GetProperty("doc")).NotTo(BeNil())
				Expect(elem.GetProperty("family")).NotTo(BeNil())
				Expect(elem.GetProperty("type")).NotTo(BeNil())
				Expect(*elem.Crn).NotTo(BeNil())
			}
		})
	})

	Describe("Call Search v3 api with query 'name:gst-sdk*' retrieving only the attributes crn and name", func() {

		// Construct an instance of the SearchOptions model
		searchOptionsModel := service.NewSearchOptions()
		searchOptionsModel.SetQuery(os.Getenv("GST_QUERY"))
		searchOptionsModel.SetLimit(1)
		searchOptionsModel.SetFields([]string{"crn", "name"})

		It("Successfully list resource using cursor", func() {
			shouldSkipTest()
			result, detailedResponse, err := service.Search(searchOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.Items).To(HaveLen(1))
			for _, elem := range result.Items {
				Expect(elem.GetProperty("doc")).To(BeNil())
				Expect(elem.GetProperty("family")).To(BeNil())
				Expect(elem.GetProperty("name")).NotTo(BeNil())
				Expect(*elem.Crn).NotTo(BeNil())
			}
			firstCrn := *result.Items[0].Crn

			search_cursor := *result.SearchCursor
			searchOptionsModelCursor := service.NewSearchOptions()
			searchOptionsModelCursor.SetQuery(os.Getenv("GST_QUERY"))
			searchOptionsModelCursor.SetLimit(1)
			searchOptionsModelCursor.SetFields([]string{"crn", "name"})
			searchOptionsModelCursor.SetSearchCursor(search_cursor)

			resultCursor, detailedResponseCursor, errCursor := service.Search(searchOptionsModelCursor)
			Expect(errCursor).To(BeNil())
			Expect(detailedResponseCursor.StatusCode).To(Equal(200))
			Expect(resultCursor.Items).To(HaveLen(1))
			for _, elem := range resultCursor.Items {
				Expect(elem.GetProperty("doc")).To(BeNil())
				Expect(elem.GetProperty("family")).To(BeNil())
				Expect(elem.GetProperty("name")).NotTo(BeNil())
				Expect(*elem.Crn).NotTo(BeNil())
			}
			secondCrn := *resultCursor.Items[0].Crn

			Expect(firstCrn).NotTo(BeIdenticalTo(secondCrn))
		})
	})

	Describe("Call GetSupportedTypes", func() {

		// Construct an instance of the SearchOptions model
		supportedTypessModel := service.NewGetSupportedTypesOptions()

		It("Successfully list all resources", func() {
			shouldSkipTest()
			
			result, detailedResponse, err := service.GetSupportedTypes(supportedTypessModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.SupportedTypes).To(ContainElement("cf-space"))
			Expect(result.SupportedTypes).NotTo(ContainElement("fake-resource!"))
		})
	})
})
