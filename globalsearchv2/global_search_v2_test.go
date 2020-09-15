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

package globalsearchv2_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`GlobalSearchV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(globalSearchService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(globalSearchService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
				URL: "https://globalsearchv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(globalSearchService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_SEARCH_URL": "https://globalsearchv2/api",
				"GLOBAL_SEARCH_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(&globalsearchv2.GlobalSearchV2Options{
				})
				Expect(globalSearchService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(&globalsearchv2.GlobalSearchV2Options{
					URL: "https://testService/api",
				})
				Expect(globalSearchService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(&globalsearchv2.GlobalSearchV2Options{
				})
				err := globalSearchService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_SEARCH_URL": "https://globalsearchv2/api",
				"GLOBAL_SEARCH_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(&globalsearchv2.GlobalSearchV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalSearchService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_SEARCH_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(&globalsearchv2.GlobalSearchV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalSearchService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Search(searchOptions *SearchOptions) - Operation response error`, func() {
		searchPath := "/v3/resources/search"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(searchPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["timeout"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Search with error: Operation response processing error`, func() {
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(globalsearchv2.SearchOptions)
				searchOptionsModel.Query = core.StringPtr("testString")
				searchOptionsModel.Fields = []string{"testString"}
				searchOptionsModel.SearchCursor = core.StringPtr("testString")
				searchOptionsModel.TransactionID = core.StringPtr("testString")
				searchOptionsModel.AccountID = core.StringPtr("testString")
				searchOptionsModel.Limit = core.Int64Ptr(int64(1))
				searchOptionsModel.Timeout = core.Int64Ptr(int64(0))
				searchOptionsModel.Sort = []string{"testString"}
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalSearchService.Search(searchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Search(searchOptions *SearchOptions)`, func() {
		searchPath := "/v3/resources/search"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(searchPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["timeout"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"search_cursor": "SearchCursor", "limit": 5, "items": [{"crn": "Crn"}]}`)
				}))
			})
			It(`Invoke Search successfully`, func() {
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalSearchService.Search(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(globalsearchv2.SearchOptions)
				searchOptionsModel.Query = core.StringPtr("testString")
				searchOptionsModel.Fields = []string{"testString"}
				searchOptionsModel.SearchCursor = core.StringPtr("testString")
				searchOptionsModel.TransactionID = core.StringPtr("testString")
				searchOptionsModel.AccountID = core.StringPtr("testString")
				searchOptionsModel.Limit = core.Int64Ptr(int64(1))
				searchOptionsModel.Timeout = core.Int64Ptr(int64(0))
				searchOptionsModel.Sort = []string{"testString"}
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalSearchService.Search(searchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke Search with error: Operation request error`, func() {
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(globalsearchv2.SearchOptions)
				searchOptionsModel.Query = core.StringPtr("testString")
				searchOptionsModel.Fields = []string{"testString"}
				searchOptionsModel.SearchCursor = core.StringPtr("testString")
				searchOptionsModel.TransactionID = core.StringPtr("testString")
				searchOptionsModel.AccountID = core.StringPtr("testString")
				searchOptionsModel.Limit = core.Int64Ptr(int64(1))
				searchOptionsModel.Timeout = core.Int64Ptr(int64(0))
				searchOptionsModel.Sort = []string{"testString"}
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalSearchService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalSearchService.Search(searchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(globalSearchService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(globalSearchService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
				URL: "https://globalsearchv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(globalSearchService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_SEARCH_URL": "https://globalsearchv2/api",
				"GLOBAL_SEARCH_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(&globalsearchv2.GlobalSearchV2Options{
				})
				Expect(globalSearchService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(&globalsearchv2.GlobalSearchV2Options{
					URL: "https://testService/api",
				})
				Expect(globalSearchService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(&globalsearchv2.GlobalSearchV2Options{
				})
				err := globalSearchService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_SEARCH_URL": "https://globalsearchv2/api",
				"GLOBAL_SEARCH_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(&globalsearchv2.GlobalSearchV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalSearchService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_SEARCH_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(&globalsearchv2.GlobalSearchV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalSearchService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetSupportedTypes(getSupportedTypesOptions *GetSupportedTypesOptions) - Operation response error`, func() {
		getSupportedTypesPath := "/v2/resources/supported_types"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportedTypesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSupportedTypes with error: Operation response processing error`, func() {
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())

				// Construct an instance of the GetSupportedTypesOptions model
				getSupportedTypesOptionsModel := new(globalsearchv2.GetSupportedTypesOptions)
				getSupportedTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalSearchService.GetSupportedTypes(getSupportedTypesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSupportedTypes(getSupportedTypesOptions *GetSupportedTypesOptions)`, func() {
		getSupportedTypesPath := "/v2/resources/supported_types"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportedTypesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"supported_types": ["SupportedTypes"]}`)
				}))
			})
			It(`Invoke GetSupportedTypes successfully`, func() {
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalSearchService.GetSupportedTypes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSupportedTypesOptions model
				getSupportedTypesOptionsModel := new(globalsearchv2.GetSupportedTypesOptions)
				getSupportedTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalSearchService.GetSupportedTypes(getSupportedTypesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSupportedTypes with error: Operation request error`, func() {
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())

				// Construct an instance of the GetSupportedTypesOptions model
				getSupportedTypesOptionsModel := new(globalsearchv2.GetSupportedTypesOptions)
				getSupportedTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalSearchService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalSearchService.GetSupportedTypes(getSupportedTypesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			globalSearchService, _ := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
				URL:           "http://globalsearchv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetSupportedTypesOptions successfully`, func() {
				// Construct an instance of the GetSupportedTypesOptions model
				getSupportedTypesOptionsModel := globalSearchService.NewGetSupportedTypesOptions()
				getSupportedTypesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSupportedTypesOptionsModel).ToNot(BeNil())
				Expect(getSupportedTypesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSearchOptions successfully`, func() {
				// Construct an instance of the SearchOptions model
				searchOptionsModel := globalSearchService.NewSearchOptions()
				searchOptionsModel.SetQuery("testString")
				searchOptionsModel.SetFields([]string{"testString"})
				searchOptionsModel.SetSearchCursor("testString")
				searchOptionsModel.SetTransactionID("testString")
				searchOptionsModel.SetAccountID("testString")
				searchOptionsModel.SetLimit(int64(1))
				searchOptionsModel.SetTimeout(int64(0))
				searchOptionsModel.SetSort([]string{"testString"})
				searchOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(searchOptionsModel).ToNot(BeNil())
				Expect(searchOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(searchOptionsModel.Fields).To(Equal([]string{"testString"}))
				Expect(searchOptionsModel.SearchCursor).To(Equal(core.StringPtr("testString")))
				Expect(searchOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(searchOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(searchOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(searchOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(0))))
				Expect(searchOptionsModel.Sort).To(Equal([]string{"testString"}))
				Expect(searchOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
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

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
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
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
