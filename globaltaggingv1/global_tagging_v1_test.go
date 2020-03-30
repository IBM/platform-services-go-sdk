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
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`GlobalTaggingV1`, func() {
	Describe(`ListTags(listTagsOptions *ListTagsOptions)`, func() {
		bearerToken := "0ui9876453"
		listTagsPath := "/v3/tags"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listTagsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["attached_to"]).To(Equal([]string{"testString"}))


				// TODO: Add check for full_data query parameter

				Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["order_by_name"]).To(Equal([]string{"asc"}))

				Expect(req.URL.Query()["timeout"]).To(Equal([]string{fmt.Sprint(int64(38))}))


				// TODO: Add check for attached_only query parameter

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"total_count": 10, "offset": 6, "limit": 5, "items": [{"name": "Name"}]}`)
			}))
			It(`Invoke ListTags successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListTags(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTagsOptions model
				listTagsOptionsModel := new(globaltaggingv1.ListTagsOptions)
				listTagsOptionsModel.Providers = []string{"ghost"}
				listTagsOptionsModel.AttachedTo = core.StringPtr("testString")
				listTagsOptionsModel.FullData = core.BoolPtr(true)
				listTagsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listTagsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listTagsOptionsModel.OrderByName = core.StringPtr("asc")
				listTagsOptionsModel.Timeout = core.Int64Ptr(int64(38))
				listTagsOptionsModel.AttachedOnly = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListTags(listTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteTagAll(deleteTagAllOptions *DeleteTagAllOptions)`, func() {
		bearerToken := "0ui9876453"
		deleteTagAllPath := "/v3/tags"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteTagAllPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["providers"]).To(Equal([]string{"ghost"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"total_count": 10, "errors": true, "items": [{"tag_name": "TagName", "is_error": false}]}`)
			}))
			It(`Invoke DeleteTagAll successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteTagAll(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteTagAllOptions model
				deleteTagAllOptionsModel := new(globaltaggingv1.DeleteTagAllOptions)
				deleteTagAllOptionsModel.Providers = core.StringPtr("ghost")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteTagAll(deleteTagAllOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteTag(deleteTagOptions *DeleteTagOptions)`, func() {
		bearerToken := "0ui9876453"
		deleteTagPath := "/v3/tags/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteTagPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"results": [{"provider": "ghost", "is_error": false}]}`)
			}))
			It(`Invoke DeleteTag successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteTagOptions model
				deleteTagOptionsModel := new(globaltaggingv1.DeleteTagOptions)
				deleteTagOptionsModel.TagName = core.StringPtr("testString")
				deleteTagOptionsModel.Providers = []string{"ghost"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteTag(deleteTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AttachTag(attachTagOptions *AttachTagOptions)`, func() {
		bearerToken := "0ui9876453"
		attachTagPath := "/v3/tags/attach"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(attachTagPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"results": [{"resource_id": "ResourceID", "is_error": false}]}`)
			}))
			It(`Invoke AttachTag successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.AttachTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Resource model
				resourceModel := new(globaltaggingv1.Resource)
				resourceModel.ResourceID = core.StringPtr("testString")
				resourceModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the AttachTagOptions model
				attachTagOptionsModel := new(globaltaggingv1.AttachTagOptions)
				attachTagOptionsModel.Resources = []globaltaggingv1.Resource{*resourceModel}
				attachTagOptionsModel.TagName = core.StringPtr("testString")
				attachTagOptionsModel.TagNames = []string{"testString"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AttachTag(attachTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DetachTag(detachTagOptions *DetachTagOptions)`, func() {
		bearerToken := "0ui9876453"
		detachTagPath := "/v3/tags/detach"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(detachTagPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"results": [{"resource_id": "ResourceID", "is_error": false}]}`)
			}))
			It(`Invoke DetachTag successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DetachTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Resource model
				resourceModel := new(globaltaggingv1.Resource)
				resourceModel.ResourceID = core.StringPtr("testString")
				resourceModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the DetachTagOptions model
				detachTagOptionsModel := new(globaltaggingv1.DetachTagOptions)
				detachTagOptionsModel.Resources = []globaltaggingv1.Resource{*resourceModel}
				detachTagOptionsModel.TagName = core.StringPtr("testString")
				detachTagOptionsModel.TagNames = []string{"testString"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DetachTag(detachTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a sample service client instance`, func() {
			testService, _ := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
				URL:           "http://globaltaggingv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewResource successfully`, func() {
				resourceID := "testString"
				model, err := testService.NewResource(resourceID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockMap() successfully`, func() {
			mockMap := CreateMockMap()
			Expect(mockMap).ToNot(BeNil())
		})
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, len(mockData))
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Now())
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Now())
	return &d
}
