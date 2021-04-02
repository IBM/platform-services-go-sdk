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

package globaltaggingv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`GlobalTaggingV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(globalTaggingService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(globalTaggingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
				URL: "https://globaltaggingv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(globalTaggingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_TAGGING_URL":       "https://globaltaggingv1/api",
				"GLOBAL_TAGGING_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1UsingExternalConfig(&globaltaggingv1.GlobalTaggingV1Options{})
				Expect(globalTaggingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := globalTaggingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalTaggingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalTaggingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalTaggingService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1UsingExternalConfig(&globaltaggingv1.GlobalTaggingV1Options{
					URL: "https://testService/api",
				})
				Expect(globalTaggingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := globalTaggingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalTaggingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalTaggingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalTaggingService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1UsingExternalConfig(&globaltaggingv1.GlobalTaggingV1Options{})
				err := globalTaggingService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := globalTaggingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalTaggingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalTaggingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalTaggingService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_TAGGING_URL":       "https://globaltaggingv1/api",
				"GLOBAL_TAGGING_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1UsingExternalConfig(&globaltaggingv1.GlobalTaggingV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(globalTaggingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_TAGGING_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1UsingExternalConfig(&globaltaggingv1.GlobalTaggingV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalTaggingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = globaltaggingv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListTags(listTagsOptions *ListTagsOptions) - Operation response error`, func() {
		listTagsPath := "/v3/tags"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTagsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// TODO: Add check for full_data query parameter

					Expect(req.URL.Query()["attached_to"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["timeout"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["order_by_name"]).To(Equal([]string{"asc"}))

					// TODO: Add check for attached_only query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTags with error: Operation response processing error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the ListTagsOptions model
				listTagsOptionsModel := new(globaltaggingv1.ListTagsOptions)
				listTagsOptionsModel.ImpersonateUser = core.StringPtr("testString")
				listTagsOptionsModel.AccountID = core.StringPtr("testString")
				listTagsOptionsModel.TagType = core.StringPtr("user")
				listTagsOptionsModel.FullData = core.BoolPtr(true)
				listTagsOptionsModel.Providers = []string{"ghost"}
				listTagsOptionsModel.AttachedTo = core.StringPtr("testString")
				listTagsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTagsOptionsModel.Timeout = core.Int64Ptr(int64(0))
				listTagsOptionsModel.OrderByName = core.StringPtr("asc")
				listTagsOptionsModel.AttachedOnly = core.BoolPtr(true)
				listTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalTaggingService.ListTags(listTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalTaggingService.EnableRetries(0, 0)
				result, response, operationErr = globalTaggingService.ListTags(listTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListTags(listTagsOptions *ListTagsOptions)`, func() {
		listTagsPath := "/v3/tags"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTagsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// TODO: Add check for full_data query parameter

					Expect(req.URL.Query()["attached_to"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["timeout"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["order_by_name"]).To(Equal([]string{"asc"}))

					// TODO: Add check for attached_only query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 0, "limit": 1, "items": [{"name": "Name"}]}`)
				}))
			})
			It(`Invoke ListTags successfully with retries`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())
				globalTaggingService.EnableRetries(0, 0)

				// Construct an instance of the ListTagsOptions model
				listTagsOptionsModel := new(globaltaggingv1.ListTagsOptions)
				listTagsOptionsModel.ImpersonateUser = core.StringPtr("testString")
				listTagsOptionsModel.AccountID = core.StringPtr("testString")
				listTagsOptionsModel.TagType = core.StringPtr("user")
				listTagsOptionsModel.FullData = core.BoolPtr(true)
				listTagsOptionsModel.Providers = []string{"ghost"}
				listTagsOptionsModel.AttachedTo = core.StringPtr("testString")
				listTagsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTagsOptionsModel.Timeout = core.Int64Ptr(int64(0))
				listTagsOptionsModel.OrderByName = core.StringPtr("asc")
				listTagsOptionsModel.AttachedOnly = core.BoolPtr(true)
				listTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalTaggingService.ListTagsWithContext(ctx, listTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalTaggingService.DisableRetries()
				result, response, operationErr := globalTaggingService.ListTags(listTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalTaggingService.ListTagsWithContext(ctx, listTagsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTagsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// TODO: Add check for full_data query parameter

					Expect(req.URL.Query()["attached_to"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["timeout"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["order_by_name"]).To(Equal([]string{"asc"}))

					// TODO: Add check for attached_only query parameter

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 0, "limit": 1, "items": [{"name": "Name"}]}`)
				}))
			})
			It(`Invoke ListTags successfully`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalTaggingService.ListTags(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTagsOptions model
				listTagsOptionsModel := new(globaltaggingv1.ListTagsOptions)
				listTagsOptionsModel.ImpersonateUser = core.StringPtr("testString")
				listTagsOptionsModel.AccountID = core.StringPtr("testString")
				listTagsOptionsModel.TagType = core.StringPtr("user")
				listTagsOptionsModel.FullData = core.BoolPtr(true)
				listTagsOptionsModel.Providers = []string{"ghost"}
				listTagsOptionsModel.AttachedTo = core.StringPtr("testString")
				listTagsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTagsOptionsModel.Timeout = core.Int64Ptr(int64(0))
				listTagsOptionsModel.OrderByName = core.StringPtr("asc")
				listTagsOptionsModel.AttachedOnly = core.BoolPtr(true)
				listTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalTaggingService.ListTags(listTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTags with error: Operation request error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the ListTagsOptions model
				listTagsOptionsModel := new(globaltaggingv1.ListTagsOptions)
				listTagsOptionsModel.ImpersonateUser = core.StringPtr("testString")
				listTagsOptionsModel.AccountID = core.StringPtr("testString")
				listTagsOptionsModel.TagType = core.StringPtr("user")
				listTagsOptionsModel.FullData = core.BoolPtr(true)
				listTagsOptionsModel.Providers = []string{"ghost"}
				listTagsOptionsModel.AttachedTo = core.StringPtr("testString")
				listTagsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTagsOptionsModel.Timeout = core.Int64Ptr(int64(0))
				listTagsOptionsModel.OrderByName = core.StringPtr("asc")
				listTagsOptionsModel.AttachedOnly = core.BoolPtr(true)
				listTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalTaggingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalTaggingService.ListTags(listTagsOptionsModel)
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
	Describe(`CreateTag(createTagOptions *CreateTagOptions) - Operation response error`, func() {
		createTagPath := "/v3/tags"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTagPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"access"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTag with error: Operation response processing error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the CreateTagOptions model
				createTagOptionsModel := new(globaltaggingv1.CreateTagOptions)
				createTagOptionsModel.TagNames = []string{"testString"}
				createTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				createTagOptionsModel.AccountID = core.StringPtr("testString")
				createTagOptionsModel.TagType = core.StringPtr("access")
				createTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalTaggingService.CreateTag(createTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalTaggingService.EnableRetries(0, 0)
				result, response, operationErr = globalTaggingService.CreateTag(createTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateTag(createTagOptions *CreateTagOptions)`, func() {
		createTagPath := "/v3/tags"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTagPath))
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

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"access"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"tag_name": "TagName", "is_error": false}]}`)
				}))
			})
			It(`Invoke CreateTag successfully with retries`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())
				globalTaggingService.EnableRetries(0, 0)

				// Construct an instance of the CreateTagOptions model
				createTagOptionsModel := new(globaltaggingv1.CreateTagOptions)
				createTagOptionsModel.TagNames = []string{"testString"}
				createTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				createTagOptionsModel.AccountID = core.StringPtr("testString")
				createTagOptionsModel.TagType = core.StringPtr("access")
				createTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalTaggingService.CreateTagWithContext(ctx, createTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalTaggingService.DisableRetries()
				result, response, operationErr := globalTaggingService.CreateTag(createTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalTaggingService.CreateTagWithContext(ctx, createTagOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTagPath))
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

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"access"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"tag_name": "TagName", "is_error": false}]}`)
				}))
			})
			It(`Invoke CreateTag successfully`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalTaggingService.CreateTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTagOptions model
				createTagOptionsModel := new(globaltaggingv1.CreateTagOptions)
				createTagOptionsModel.TagNames = []string{"testString"}
				createTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				createTagOptionsModel.AccountID = core.StringPtr("testString")
				createTagOptionsModel.TagType = core.StringPtr("access")
				createTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalTaggingService.CreateTag(createTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTag with error: Operation validation and request error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the CreateTagOptions model
				createTagOptionsModel := new(globaltaggingv1.CreateTagOptions)
				createTagOptionsModel.TagNames = []string{"testString"}
				createTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				createTagOptionsModel.AccountID = core.StringPtr("testString")
				createTagOptionsModel.TagType = core.StringPtr("access")
				createTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalTaggingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalTaggingService.CreateTag(createTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTagOptions model with no property values
				createTagOptionsModelNew := new(globaltaggingv1.CreateTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalTaggingService.CreateTag(createTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTagAll(deleteTagAllOptions *DeleteTagAllOptions) - Operation response error`, func() {
		deleteTagAllPath := "/v3/tags"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTagAllPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["providers"]).To(Equal([]string{"ghost"}))

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteTagAll with error: Operation response processing error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the DeleteTagAllOptions model
				deleteTagAllOptionsModel := new(globaltaggingv1.DeleteTagAllOptions)
				deleteTagAllOptionsModel.Providers = core.StringPtr("ghost")
				deleteTagAllOptionsModel.ImpersonateUser = core.StringPtr("testString")
				deleteTagAllOptionsModel.AccountID = core.StringPtr("testString")
				deleteTagAllOptionsModel.TagType = core.StringPtr("user")
				deleteTagAllOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalTaggingService.DeleteTagAll(deleteTagAllOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalTaggingService.EnableRetries(0, 0)
				result, response, operationErr = globalTaggingService.DeleteTagAll(deleteTagAllOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteTagAll(deleteTagAllOptions *DeleteTagAllOptions)`, func() {
		deleteTagAllPath := "/v3/tags"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTagAllPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["providers"]).To(Equal([]string{"ghost"}))

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "errors": true, "items": [{"tag_name": "TagName", "is_error": false}]}`)
				}))
			})
			It(`Invoke DeleteTagAll successfully with retries`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())
				globalTaggingService.EnableRetries(0, 0)

				// Construct an instance of the DeleteTagAllOptions model
				deleteTagAllOptionsModel := new(globaltaggingv1.DeleteTagAllOptions)
				deleteTagAllOptionsModel.Providers = core.StringPtr("ghost")
				deleteTagAllOptionsModel.ImpersonateUser = core.StringPtr("testString")
				deleteTagAllOptionsModel.AccountID = core.StringPtr("testString")
				deleteTagAllOptionsModel.TagType = core.StringPtr("user")
				deleteTagAllOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalTaggingService.DeleteTagAllWithContext(ctx, deleteTagAllOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalTaggingService.DisableRetries()
				result, response, operationErr := globalTaggingService.DeleteTagAll(deleteTagAllOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalTaggingService.DeleteTagAllWithContext(ctx, deleteTagAllOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteTagAllPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["providers"]).To(Equal([]string{"ghost"}))

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "errors": true, "items": [{"tag_name": "TagName", "is_error": false}]}`)
				}))
			})
			It(`Invoke DeleteTagAll successfully`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalTaggingService.DeleteTagAll(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteTagAllOptions model
				deleteTagAllOptionsModel := new(globaltaggingv1.DeleteTagAllOptions)
				deleteTagAllOptionsModel.Providers = core.StringPtr("ghost")
				deleteTagAllOptionsModel.ImpersonateUser = core.StringPtr("testString")
				deleteTagAllOptionsModel.AccountID = core.StringPtr("testString")
				deleteTagAllOptionsModel.TagType = core.StringPtr("user")
				deleteTagAllOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalTaggingService.DeleteTagAll(deleteTagAllOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteTagAll with error: Operation request error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the DeleteTagAllOptions model
				deleteTagAllOptionsModel := new(globaltaggingv1.DeleteTagAllOptions)
				deleteTagAllOptionsModel.Providers = core.StringPtr("ghost")
				deleteTagAllOptionsModel.ImpersonateUser = core.StringPtr("testString")
				deleteTagAllOptionsModel.AccountID = core.StringPtr("testString")
				deleteTagAllOptionsModel.TagType = core.StringPtr("user")
				deleteTagAllOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalTaggingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalTaggingService.DeleteTagAll(deleteTagAllOptionsModel)
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
	Describe(`DeleteTag(deleteTagOptions *DeleteTagOptions) - Operation response error`, func() {
		deleteTagPath := "/v3/tags/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTagPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteTag with error: Operation response processing error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the DeleteTagOptions model
				deleteTagOptionsModel := new(globaltaggingv1.DeleteTagOptions)
				deleteTagOptionsModel.TagName = core.StringPtr("testString")
				deleteTagOptionsModel.Providers = []string{"ghost"}
				deleteTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				deleteTagOptionsModel.AccountID = core.StringPtr("testString")
				deleteTagOptionsModel.TagType = core.StringPtr("user")
				deleteTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalTaggingService.DeleteTag(deleteTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalTaggingService.EnableRetries(0, 0)
				result, response, operationErr = globalTaggingService.DeleteTag(deleteTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteTag(deleteTagOptions *DeleteTagOptions)`, func() {
		deleteTagPath := "/v3/tags/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTagPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"provider": "ghost", "is_error": false}]}`)
				}))
			})
			It(`Invoke DeleteTag successfully with retries`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())
				globalTaggingService.EnableRetries(0, 0)

				// Construct an instance of the DeleteTagOptions model
				deleteTagOptionsModel := new(globaltaggingv1.DeleteTagOptions)
				deleteTagOptionsModel.TagName = core.StringPtr("testString")
				deleteTagOptionsModel.Providers = []string{"ghost"}
				deleteTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				deleteTagOptionsModel.AccountID = core.StringPtr("testString")
				deleteTagOptionsModel.TagType = core.StringPtr("user")
				deleteTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalTaggingService.DeleteTagWithContext(ctx, deleteTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalTaggingService.DisableRetries()
				result, response, operationErr := globalTaggingService.DeleteTag(deleteTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalTaggingService.DeleteTagWithContext(ctx, deleteTagOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteTagPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"provider": "ghost", "is_error": false}]}`)
				}))
			})
			It(`Invoke DeleteTag successfully`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalTaggingService.DeleteTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteTagOptions model
				deleteTagOptionsModel := new(globaltaggingv1.DeleteTagOptions)
				deleteTagOptionsModel.TagName = core.StringPtr("testString")
				deleteTagOptionsModel.Providers = []string{"ghost"}
				deleteTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				deleteTagOptionsModel.AccountID = core.StringPtr("testString")
				deleteTagOptionsModel.TagType = core.StringPtr("user")
				deleteTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalTaggingService.DeleteTag(deleteTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteTag with error: Operation validation and request error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the DeleteTagOptions model
				deleteTagOptionsModel := new(globaltaggingv1.DeleteTagOptions)
				deleteTagOptionsModel.TagName = core.StringPtr("testString")
				deleteTagOptionsModel.Providers = []string{"ghost"}
				deleteTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				deleteTagOptionsModel.AccountID = core.StringPtr("testString")
				deleteTagOptionsModel.TagType = core.StringPtr("user")
				deleteTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalTaggingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalTaggingService.DeleteTag(deleteTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteTagOptions model with no property values
				deleteTagOptionsModelNew := new(globaltaggingv1.DeleteTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalTaggingService.DeleteTag(deleteTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AttachTag(attachTagOptions *AttachTagOptions) - Operation response error`, func() {
		attachTagPath := "/v3/tags/attach"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(attachTagPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AttachTag with error: Operation response processing error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the Resource model
				resourceModel := new(globaltaggingv1.Resource)
				resourceModel.ResourceID = core.StringPtr("testString")
				resourceModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the AttachTagOptions model
				attachTagOptionsModel := new(globaltaggingv1.AttachTagOptions)
				attachTagOptionsModel.Resources = []globaltaggingv1.Resource{*resourceModel}
				attachTagOptionsModel.TagName = core.StringPtr("testString")
				attachTagOptionsModel.TagNames = []string{"testString"}
				attachTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				attachTagOptionsModel.AccountID = core.StringPtr("testString")
				attachTagOptionsModel.TagType = core.StringPtr("user")
				attachTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalTaggingService.AttachTag(attachTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalTaggingService.EnableRetries(0, 0)
				result, response, operationErr = globalTaggingService.AttachTag(attachTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AttachTag(attachTagOptions *AttachTagOptions)`, func() {
		attachTagPath := "/v3/tags/attach"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(attachTagPath))
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

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"resource_id": "ResourceID", "is_error": false}]}`)
				}))
			})
			It(`Invoke AttachTag successfully with retries`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())
				globalTaggingService.EnableRetries(0, 0)

				// Construct an instance of the Resource model
				resourceModel := new(globaltaggingv1.Resource)
				resourceModel.ResourceID = core.StringPtr("testString")
				resourceModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the AttachTagOptions model
				attachTagOptionsModel := new(globaltaggingv1.AttachTagOptions)
				attachTagOptionsModel.Resources = []globaltaggingv1.Resource{*resourceModel}
				attachTagOptionsModel.TagName = core.StringPtr("testString")
				attachTagOptionsModel.TagNames = []string{"testString"}
				attachTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				attachTagOptionsModel.AccountID = core.StringPtr("testString")
				attachTagOptionsModel.TagType = core.StringPtr("user")
				attachTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalTaggingService.AttachTagWithContext(ctx, attachTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalTaggingService.DisableRetries()
				result, response, operationErr := globalTaggingService.AttachTag(attachTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalTaggingService.AttachTagWithContext(ctx, attachTagOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(attachTagPath))
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

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"resource_id": "ResourceID", "is_error": false}]}`)
				}))
			})
			It(`Invoke AttachTag successfully`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalTaggingService.AttachTag(nil)
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
				attachTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				attachTagOptionsModel.AccountID = core.StringPtr("testString")
				attachTagOptionsModel.TagType = core.StringPtr("user")
				attachTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalTaggingService.AttachTag(attachTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AttachTag with error: Operation validation and request error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the Resource model
				resourceModel := new(globaltaggingv1.Resource)
				resourceModel.ResourceID = core.StringPtr("testString")
				resourceModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the AttachTagOptions model
				attachTagOptionsModel := new(globaltaggingv1.AttachTagOptions)
				attachTagOptionsModel.Resources = []globaltaggingv1.Resource{*resourceModel}
				attachTagOptionsModel.TagName = core.StringPtr("testString")
				attachTagOptionsModel.TagNames = []string{"testString"}
				attachTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				attachTagOptionsModel.AccountID = core.StringPtr("testString")
				attachTagOptionsModel.TagType = core.StringPtr("user")
				attachTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalTaggingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalTaggingService.AttachTag(attachTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AttachTagOptions model with no property values
				attachTagOptionsModelNew := new(globaltaggingv1.AttachTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalTaggingService.AttachTag(attachTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DetachTag(detachTagOptions *DetachTagOptions) - Operation response error`, func() {
		detachTagPath := "/v3/tags/detach"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(detachTagPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DetachTag with error: Operation response processing error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the Resource model
				resourceModel := new(globaltaggingv1.Resource)
				resourceModel.ResourceID = core.StringPtr("testString")
				resourceModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the DetachTagOptions model
				detachTagOptionsModel := new(globaltaggingv1.DetachTagOptions)
				detachTagOptionsModel.Resources = []globaltaggingv1.Resource{*resourceModel}
				detachTagOptionsModel.TagName = core.StringPtr("testString")
				detachTagOptionsModel.TagNames = []string{"testString"}
				detachTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				detachTagOptionsModel.AccountID = core.StringPtr("testString")
				detachTagOptionsModel.TagType = core.StringPtr("user")
				detachTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalTaggingService.DetachTag(detachTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalTaggingService.EnableRetries(0, 0)
				result, response, operationErr = globalTaggingService.DetachTag(detachTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DetachTag(detachTagOptions *DetachTagOptions)`, func() {
		detachTagPath := "/v3/tags/detach"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(detachTagPath))
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

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"resource_id": "ResourceID", "is_error": false}]}`)
				}))
			})
			It(`Invoke DetachTag successfully with retries`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())
				globalTaggingService.EnableRetries(0, 0)

				// Construct an instance of the Resource model
				resourceModel := new(globaltaggingv1.Resource)
				resourceModel.ResourceID = core.StringPtr("testString")
				resourceModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the DetachTagOptions model
				detachTagOptionsModel := new(globaltaggingv1.DetachTagOptions)
				detachTagOptionsModel.Resources = []globaltaggingv1.Resource{*resourceModel}
				detachTagOptionsModel.TagName = core.StringPtr("testString")
				detachTagOptionsModel.TagNames = []string{"testString"}
				detachTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				detachTagOptionsModel.AccountID = core.StringPtr("testString")
				detachTagOptionsModel.TagType = core.StringPtr("user")
				detachTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalTaggingService.DetachTagWithContext(ctx, detachTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalTaggingService.DisableRetries()
				result, response, operationErr := globalTaggingService.DetachTag(detachTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalTaggingService.DetachTagWithContext(ctx, detachTagOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(detachTagPath))
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

					Expect(req.URL.Query()["impersonate_user"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["tag_type"]).To(Equal([]string{"user"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"resource_id": "ResourceID", "is_error": false}]}`)
				}))
			})
			It(`Invoke DetachTag successfully`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalTaggingService.DetachTag(nil)
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
				detachTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				detachTagOptionsModel.AccountID = core.StringPtr("testString")
				detachTagOptionsModel.TagType = core.StringPtr("user")
				detachTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalTaggingService.DetachTag(detachTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DetachTag with error: Operation validation and request error`, func() {
				globalTaggingService, serviceErr := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalTaggingService).ToNot(BeNil())

				// Construct an instance of the Resource model
				resourceModel := new(globaltaggingv1.Resource)
				resourceModel.ResourceID = core.StringPtr("testString")
				resourceModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the DetachTagOptions model
				detachTagOptionsModel := new(globaltaggingv1.DetachTagOptions)
				detachTagOptionsModel.Resources = []globaltaggingv1.Resource{*resourceModel}
				detachTagOptionsModel.TagName = core.StringPtr("testString")
				detachTagOptionsModel.TagNames = []string{"testString"}
				detachTagOptionsModel.ImpersonateUser = core.StringPtr("testString")
				detachTagOptionsModel.AccountID = core.StringPtr("testString")
				detachTagOptionsModel.TagType = core.StringPtr("user")
				detachTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalTaggingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalTaggingService.DetachTag(detachTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DetachTagOptions model with no property values
				detachTagOptionsModelNew := new(globaltaggingv1.DetachTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalTaggingService.DetachTag(detachTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			globalTaggingService, _ := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
				URL:           "http://globaltaggingv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAttachTagOptions successfully`, func() {
				// Construct an instance of the Resource model
				resourceModel := new(globaltaggingv1.Resource)
				Expect(resourceModel).ToNot(BeNil())
				resourceModel.ResourceID = core.StringPtr("testString")
				resourceModel.ResourceType = core.StringPtr("testString")
				Expect(resourceModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(resourceModel.ResourceType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AttachTagOptions model
				attachTagOptionsResources := []globaltaggingv1.Resource{}
				attachTagOptionsModel := globalTaggingService.NewAttachTagOptions(attachTagOptionsResources)
				attachTagOptionsModel.SetResources([]globaltaggingv1.Resource{*resourceModel})
				attachTagOptionsModel.SetTagName("testString")
				attachTagOptionsModel.SetTagNames([]string{"testString"})
				attachTagOptionsModel.SetImpersonateUser("testString")
				attachTagOptionsModel.SetAccountID("testString")
				attachTagOptionsModel.SetTagType("user")
				attachTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(attachTagOptionsModel).ToNot(BeNil())
				Expect(attachTagOptionsModel.Resources).To(Equal([]globaltaggingv1.Resource{*resourceModel}))
				Expect(attachTagOptionsModel.TagName).To(Equal(core.StringPtr("testString")))
				Expect(attachTagOptionsModel.TagNames).To(Equal([]string{"testString"}))
				Expect(attachTagOptionsModel.ImpersonateUser).To(Equal(core.StringPtr("testString")))
				Expect(attachTagOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(attachTagOptionsModel.TagType).To(Equal(core.StringPtr("user")))
				Expect(attachTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTagOptions successfully`, func() {
				// Construct an instance of the CreateTagOptions model
				createTagOptionsTagNames := []string{"testString"}
				createTagOptionsModel := globalTaggingService.NewCreateTagOptions(createTagOptionsTagNames)
				createTagOptionsModel.SetTagNames([]string{"testString"})
				createTagOptionsModel.SetImpersonateUser("testString")
				createTagOptionsModel.SetAccountID("testString")
				createTagOptionsModel.SetTagType("access")
				createTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTagOptionsModel).ToNot(BeNil())
				Expect(createTagOptionsModel.TagNames).To(Equal([]string{"testString"}))
				Expect(createTagOptionsModel.ImpersonateUser).To(Equal(core.StringPtr("testString")))
				Expect(createTagOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createTagOptionsModel.TagType).To(Equal(core.StringPtr("access")))
				Expect(createTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTagAllOptions successfully`, func() {
				// Construct an instance of the DeleteTagAllOptions model
				deleteTagAllOptionsModel := globalTaggingService.NewDeleteTagAllOptions()
				deleteTagAllOptionsModel.SetProviders("ghost")
				deleteTagAllOptionsModel.SetImpersonateUser("testString")
				deleteTagAllOptionsModel.SetAccountID("testString")
				deleteTagAllOptionsModel.SetTagType("user")
				deleteTagAllOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTagAllOptionsModel).ToNot(BeNil())
				Expect(deleteTagAllOptionsModel.Providers).To(Equal(core.StringPtr("ghost")))
				Expect(deleteTagAllOptionsModel.ImpersonateUser).To(Equal(core.StringPtr("testString")))
				Expect(deleteTagAllOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTagAllOptionsModel.TagType).To(Equal(core.StringPtr("user")))
				Expect(deleteTagAllOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTagOptions successfully`, func() {
				// Construct an instance of the DeleteTagOptions model
				tagName := "testString"
				deleteTagOptionsModel := globalTaggingService.NewDeleteTagOptions(tagName)
				deleteTagOptionsModel.SetTagName("testString")
				deleteTagOptionsModel.SetProviders([]string{"ghost"})
				deleteTagOptionsModel.SetImpersonateUser("testString")
				deleteTagOptionsModel.SetAccountID("testString")
				deleteTagOptionsModel.SetTagType("user")
				deleteTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTagOptionsModel).ToNot(BeNil())
				Expect(deleteTagOptionsModel.TagName).To(Equal(core.StringPtr("testString")))
				Expect(deleteTagOptionsModel.Providers).To(Equal([]string{"ghost"}))
				Expect(deleteTagOptionsModel.ImpersonateUser).To(Equal(core.StringPtr("testString")))
				Expect(deleteTagOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTagOptionsModel.TagType).To(Equal(core.StringPtr("user")))
				Expect(deleteTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDetachTagOptions successfully`, func() {
				// Construct an instance of the Resource model
				resourceModel := new(globaltaggingv1.Resource)
				Expect(resourceModel).ToNot(BeNil())
				resourceModel.ResourceID = core.StringPtr("testString")
				resourceModel.ResourceType = core.StringPtr("testString")
				Expect(resourceModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(resourceModel.ResourceType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DetachTagOptions model
				detachTagOptionsResources := []globaltaggingv1.Resource{}
				detachTagOptionsModel := globalTaggingService.NewDetachTagOptions(detachTagOptionsResources)
				detachTagOptionsModel.SetResources([]globaltaggingv1.Resource{*resourceModel})
				detachTagOptionsModel.SetTagName("testString")
				detachTagOptionsModel.SetTagNames([]string{"testString"})
				detachTagOptionsModel.SetImpersonateUser("testString")
				detachTagOptionsModel.SetAccountID("testString")
				detachTagOptionsModel.SetTagType("user")
				detachTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(detachTagOptionsModel).ToNot(BeNil())
				Expect(detachTagOptionsModel.Resources).To(Equal([]globaltaggingv1.Resource{*resourceModel}))
				Expect(detachTagOptionsModel.TagName).To(Equal(core.StringPtr("testString")))
				Expect(detachTagOptionsModel.TagNames).To(Equal([]string{"testString"}))
				Expect(detachTagOptionsModel.ImpersonateUser).To(Equal(core.StringPtr("testString")))
				Expect(detachTagOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(detachTagOptionsModel.TagType).To(Equal(core.StringPtr("user")))
				Expect(detachTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTagsOptions successfully`, func() {
				// Construct an instance of the ListTagsOptions model
				listTagsOptionsModel := globalTaggingService.NewListTagsOptions()
				listTagsOptionsModel.SetImpersonateUser("testString")
				listTagsOptionsModel.SetAccountID("testString")
				listTagsOptionsModel.SetTagType("user")
				listTagsOptionsModel.SetFullData(true)
				listTagsOptionsModel.SetProviders([]string{"ghost"})
				listTagsOptionsModel.SetAttachedTo("testString")
				listTagsOptionsModel.SetOffset(int64(0))
				listTagsOptionsModel.SetLimit(int64(1))
				listTagsOptionsModel.SetTimeout(int64(0))
				listTagsOptionsModel.SetOrderByName("asc")
				listTagsOptionsModel.SetAttachedOnly(true)
				listTagsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTagsOptionsModel).ToNot(BeNil())
				Expect(listTagsOptionsModel.ImpersonateUser).To(Equal(core.StringPtr("testString")))
				Expect(listTagsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsOptionsModel.TagType).To(Equal(core.StringPtr("user")))
				Expect(listTagsOptionsModel.FullData).To(Equal(core.BoolPtr(true)))
				Expect(listTagsOptionsModel.Providers).To(Equal([]string{"ghost"}))
				Expect(listTagsOptionsModel.AttachedTo).To(Equal(core.StringPtr("testString")))
				Expect(listTagsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listTagsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listTagsOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listTagsOptionsModel.OrderByName).To(Equal(core.StringPtr("asc")))
				Expect(listTagsOptionsModel.AttachedOnly).To(Equal(core.BoolPtr(true)))
				Expect(listTagsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResource successfully`, func() {
				resourceID := "testString"
				model, err := globalTaggingService.NewResource(resourceID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
