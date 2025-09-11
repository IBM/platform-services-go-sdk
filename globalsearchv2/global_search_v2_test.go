/**
 * (C) Copyright IBM Corp. 2025.
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
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

				clone := globalSearchService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalSearchService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalSearchService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalSearchService.Service.Options.Authenticator))
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

				clone := globalSearchService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalSearchService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalSearchService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalSearchService.Service.Options.Authenticator))
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

				clone := globalSearchService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalSearchService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalSearchService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalSearchService.Service.Options.Authenticator))
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
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = globalsearchv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Search(searchOptions *SearchOptions) - Operation response error`, func() {
		searchPath := "/v3/resources/search"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(searchPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["timeout"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["is_deleted"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["is_reclaimed"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["can_tag"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["is_project_resource"]).To(Equal([]string{"false"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Search with error: Operation response processing error`, func() {
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())

				// Construct an instance of the SearchRequestFirstCall model
				searchRequestModel := new(globalsearchv2.SearchRequestFirstCall)
				searchRequestModel.Query = core.StringPtr("testString")
				searchRequestModel.Fields = []string{"testString"}

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(globalsearchv2.SearchOptions)
				searchOptionsModel.Body = searchRequestModel
				searchOptionsModel.XRequestID = core.StringPtr("testString")
				searchOptionsModel.XCorrelationID = core.StringPtr("testString")
				searchOptionsModel.AccountID = core.StringPtr("testString")
				searchOptionsModel.Limit = core.Int64Ptr(int64(10))
				searchOptionsModel.Timeout = core.Int64Ptr(int64(0))
				searchOptionsModel.Sort = []string{"testString"}
				searchOptionsModel.IsDeleted = core.StringPtr("false")
				searchOptionsModel.IsReclaimed = core.StringPtr("false")
				searchOptionsModel.ImpersonateUser = core.StringPtr("testString")
				searchOptionsModel.CanTag = core.StringPtr("false")
				searchOptionsModel.IsProjectResource = core.StringPtr("false")
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalSearchService.Search(searchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalSearchService.EnableRetries(0, 0)
				result, response, operationErr = globalSearchService.Search(searchOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(searchPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["timeout"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["is_deleted"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["is_reclaimed"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["can_tag"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["is_project_resource"]).To(Equal([]string{"false"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"search_cursor": "SearchCursor", "limit": 5, "items": [{"crn": "CRN"}]}`)
				}))
			})
			It(`Invoke Search successfully with retries`, func() {
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())
				globalSearchService.EnableRetries(0, 0)

				// Construct an instance of the SearchRequestFirstCall model
				searchRequestModel := new(globalsearchv2.SearchRequestFirstCall)
				searchRequestModel.Query = core.StringPtr("testString")
				searchRequestModel.Fields = []string{"testString"}

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(globalsearchv2.SearchOptions)
				searchOptionsModel.Body = searchRequestModel
				searchOptionsModel.XRequestID = core.StringPtr("testString")
				searchOptionsModel.XCorrelationID = core.StringPtr("testString")
				searchOptionsModel.AccountID = core.StringPtr("testString")
				searchOptionsModel.Limit = core.Int64Ptr(int64(10))
				searchOptionsModel.Timeout = core.Int64Ptr(int64(0))
				searchOptionsModel.Sort = []string{"testString"}
				searchOptionsModel.IsDeleted = core.StringPtr("false")
				searchOptionsModel.IsReclaimed = core.StringPtr("false")
				searchOptionsModel.ImpersonateUser = core.StringPtr("testString")
				searchOptionsModel.CanTag = core.StringPtr("false")
				searchOptionsModel.IsProjectResource = core.StringPtr("false")
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalSearchService.SearchWithContext(ctx, searchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalSearchService.DisableRetries()
				result, response, operationErr := globalSearchService.Search(searchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalSearchService.SearchWithContext(ctx, searchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(searchPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["timeout"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["is_deleted"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["is_reclaimed"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["can_tag"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["is_project_resource"]).To(Equal([]string{"false"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"search_cursor": "SearchCursor", "limit": 5, "items": [{"crn": "CRN"}]}`)
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

				// Construct an instance of the SearchRequestFirstCall model
				searchRequestModel := new(globalsearchv2.SearchRequestFirstCall)
				searchRequestModel.Query = core.StringPtr("testString")
				searchRequestModel.Fields = []string{"testString"}

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(globalsearchv2.SearchOptions)
				searchOptionsModel.Body = searchRequestModel
				searchOptionsModel.XRequestID = core.StringPtr("testString")
				searchOptionsModel.XCorrelationID = core.StringPtr("testString")
				searchOptionsModel.AccountID = core.StringPtr("testString")
				searchOptionsModel.Limit = core.Int64Ptr(int64(10))
				searchOptionsModel.Timeout = core.Int64Ptr(int64(0))
				searchOptionsModel.Sort = []string{"testString"}
				searchOptionsModel.IsDeleted = core.StringPtr("false")
				searchOptionsModel.IsReclaimed = core.StringPtr("false")
				searchOptionsModel.ImpersonateUser = core.StringPtr("testString")
				searchOptionsModel.CanTag = core.StringPtr("false")
				searchOptionsModel.IsProjectResource = core.StringPtr("false")
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalSearchService.Search(searchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Search with error: Operation validation and request error`, func() {
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())

				// Construct an instance of the SearchRequestFirstCall model
				searchRequestModel := new(globalsearchv2.SearchRequestFirstCall)
				searchRequestModel.Query = core.StringPtr("testString")
				searchRequestModel.Fields = []string{"testString"}

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(globalsearchv2.SearchOptions)
				searchOptionsModel.Body = searchRequestModel
				searchOptionsModel.XRequestID = core.StringPtr("testString")
				searchOptionsModel.XCorrelationID = core.StringPtr("testString")
				searchOptionsModel.AccountID = core.StringPtr("testString")
				searchOptionsModel.Limit = core.Int64Ptr(int64(10))
				searchOptionsModel.Timeout = core.Int64Ptr(int64(0))
				searchOptionsModel.Sort = []string{"testString"}
				searchOptionsModel.IsDeleted = core.StringPtr("false")
				searchOptionsModel.IsReclaimed = core.StringPtr("false")
				searchOptionsModel.ImpersonateUser = core.StringPtr("testString")
				searchOptionsModel.CanTag = core.StringPtr("false")
				searchOptionsModel.IsProjectResource = core.StringPtr("false")
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalSearchService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalSearchService.Search(searchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SearchOptions model with no property values
				searchOptionsModelNew := new(globalsearchv2.SearchOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalSearchService.Search(searchOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke Search successfully`, func() {
				globalSearchService, serviceErr := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalSearchService).ToNot(BeNil())

				// Construct an instance of the SearchRequestFirstCall model
				searchRequestModel := new(globalsearchv2.SearchRequestFirstCall)
				searchRequestModel.Query = core.StringPtr("testString")
				searchRequestModel.Fields = []string{"testString"}

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(globalsearchv2.SearchOptions)
				searchOptionsModel.Body = searchRequestModel
				searchOptionsModel.XRequestID = core.StringPtr("testString")
				searchOptionsModel.XCorrelationID = core.StringPtr("testString")
				searchOptionsModel.AccountID = core.StringPtr("testString")
				searchOptionsModel.Limit = core.Int64Ptr(int64(10))
				searchOptionsModel.Timeout = core.Int64Ptr(int64(0))
				searchOptionsModel.Sort = []string{"testString"}
				searchOptionsModel.IsDeleted = core.StringPtr("false")
				searchOptionsModel.IsReclaimed = core.StringPtr("false")
				searchOptionsModel.ImpersonateUser = core.StringPtr("testString")
				searchOptionsModel.CanTag = core.StringPtr("false")
				searchOptionsModel.IsProjectResource = core.StringPtr("false")
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := globalSearchService.Search(searchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
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
			It(`Invoke NewSearchOptions successfully`, func() {
				// Construct an instance of the SearchRequestFirstCall model
				searchRequestModel := new(globalsearchv2.SearchRequestFirstCall)
				Expect(searchRequestModel).ToNot(BeNil())
				searchRequestModel.Query = core.StringPtr("testString")
				searchRequestModel.Fields = []string{"testString"}
				Expect(searchRequestModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(searchRequestModel.Fields).To(Equal([]string{"testString"}))

				// Construct an instance of the SearchOptions model
				var body globalsearchv2.SearchRequestIntf = nil
				searchOptionsModel := globalSearchService.NewSearchOptions(body)
				searchOptionsModel.SetBody(searchRequestModel)
				searchOptionsModel.SetXRequestID("testString")
				searchOptionsModel.SetXCorrelationID("testString")
				searchOptionsModel.SetAccountID("testString")
				searchOptionsModel.SetLimit(int64(10))
				searchOptionsModel.SetTimeout(int64(0))
				searchOptionsModel.SetSort([]string{"testString"})
				searchOptionsModel.SetIsDeleted("false")
				searchOptionsModel.SetIsReclaimed("false")
				searchOptionsModel.SetImpersonateUser("testString")
				searchOptionsModel.SetCanTag("false")
				searchOptionsModel.SetIsProjectResource("false")
				searchOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(searchOptionsModel).ToNot(BeNil())
				Expect(searchOptionsModel.Body).To(Equal(searchRequestModel))
				Expect(searchOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(searchOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(searchOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(searchOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(searchOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(0))))
				Expect(searchOptionsModel.Sort).To(Equal([]string{"testString"}))
				Expect(searchOptionsModel.IsDeleted).To(Equal(core.StringPtr("false")))
				Expect(searchOptionsModel.IsReclaimed).To(Equal(core.StringPtr("false")))
				Expect(searchOptionsModel.ImpersonateUser).To(Equal(core.StringPtr("testString")))
				Expect(searchOptionsModel.CanTag).To(Equal(core.StringPtr("false")))
				Expect(searchOptionsModel.IsProjectResource).To(Equal(core.StringPtr("false")))
				Expect(searchOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSearchRequestFirstCall successfully`, func() {
				query := "testString"
				_model, err := globalSearchService.NewSearchRequestFirstCall(query)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSearchRequestNextCall successfully`, func() {
				searchCursor := "testString"
				_model, err := globalSearchService.NewSearchRequestNextCall(searchCursor)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalSearchRequest successfully`, func() {
			// Construct an instance of the model.
			model := new(globalsearchv2.SearchRequest)
			model.Query = core.StringPtr("testString")
			model.Fields = []string{"testString"}
			model.SearchCursor = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalsearchv2.SearchRequest
			err = globalsearchv2.UnmarshalSearchRequest(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSearchRequestFirstCall successfully`, func() {
			// Construct an instance of the model.
			model := new(globalsearchv2.SearchRequestFirstCall)
			model.Query = core.StringPtr("testString")
			model.Fields = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalsearchv2.SearchRequestFirstCall
			err = globalsearchv2.UnmarshalSearchRequestFirstCall(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSearchRequestNextCall successfully`, func() {
			// Construct an instance of the model.
			model := new(globalsearchv2.SearchRequestNextCall)
			model.SearchCursor = core.StringPtr("testString")
			model.Query = core.StringPtr("testString")
			model.Fields = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalsearchv2.SearchRequestNextCall
			err = globalsearchv2.UnmarshalSearchRequestNextCall(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
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
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
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
