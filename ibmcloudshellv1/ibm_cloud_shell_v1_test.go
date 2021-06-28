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

package ibmcloudshellv1_test

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
	"github.com/IBM/platform-services-go-sdk/ibmcloudshellv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`IbmCloudShellV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudShellService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudShellService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
				URL: "https://ibmcloudshellv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudShellService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_SHELL_URL":       "https://ibmcloudshellv1/api",
				"IBM_CLOUD_SHELL_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1UsingExternalConfig(&ibmcloudshellv1.IbmCloudShellV1Options{})
				Expect(ibmCloudShellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudShellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudShellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudShellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudShellService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1UsingExternalConfig(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudShellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudShellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudShellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudShellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudShellService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1UsingExternalConfig(&ibmcloudshellv1.IbmCloudShellV1Options{})
				err := ibmCloudShellService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudShellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudShellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudShellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudShellService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_SHELL_URL":       "https://ibmcloudshellv1/api",
				"IBM_CLOUD_SHELL_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1UsingExternalConfig(&ibmcloudshellv1.IbmCloudShellV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudShellService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_SHELL_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1UsingExternalConfig(&ibmcloudshellv1.IbmCloudShellV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudShellService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmcloudshellv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAccountSettingsByID(getAccountSettingsByIdOptions *GetAccountSettingsByIdOptions) - Operation response error`, func() {
		getAccountSettingsByIDPath := "/api/v1/user/accounts/12345678-abcd-1a2b-a1b2-1234567890ab/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsByIDPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountSettingsByID with error: Operation response processing error`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsByIdOptions model
				getAccountSettingsByIdOptionsModel := new(ibmcloudshellv1.GetAccountSettingsByIdOptions)
				getAccountSettingsByIdOptionsModel.AccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				getAccountSettingsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudShellService.GetAccountSettingsByID(getAccountSettingsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudShellService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudShellService.GetAccountSettingsByID(getAccountSettingsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountSettingsByID(getAccountSettingsByIdOptions *GetAccountSettingsByIdOptions)`, func() {
		getAccountSettingsByIDPath := "/api/v1/user/accounts/12345678-abcd-1a2b-a1b2-1234567890ab/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_id": "ID", "_rev": "Rev", "account_id": "AccountID", "created_at": 9, "created_by": "CreatedBy", "default_enable_new_features": true, "default_enable_new_regions": false, "enabled": false, "features": [{"enabled": false, "key": "Key"}], "regions": [{"enabled": false, "key": "Key"}], "type": "Type", "updated_at": 9, "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetAccountSettingsByID successfully with retries`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())
				ibmCloudShellService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountSettingsByIdOptions model
				getAccountSettingsByIdOptionsModel := new(ibmcloudshellv1.GetAccountSettingsByIdOptions)
				getAccountSettingsByIdOptionsModel.AccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				getAccountSettingsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudShellService.GetAccountSettingsByIDWithContext(ctx, getAccountSettingsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudShellService.DisableRetries()
				result, response, operationErr := ibmCloudShellService.GetAccountSettingsByID(getAccountSettingsByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudShellService.GetAccountSettingsByIDWithContext(ctx, getAccountSettingsByIdOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_id": "ID", "_rev": "Rev", "account_id": "AccountID", "created_at": 9, "created_by": "CreatedBy", "default_enable_new_features": true, "default_enable_new_regions": false, "enabled": false, "features": [{"enabled": false, "key": "Key"}], "regions": [{"enabled": false, "key": "Key"}], "type": "Type", "updated_at": 9, "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetAccountSettingsByID successfully`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudShellService.GetAccountSettingsByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSettingsByIdOptions model
				getAccountSettingsByIdOptionsModel := new(ibmcloudshellv1.GetAccountSettingsByIdOptions)
				getAccountSettingsByIdOptionsModel.AccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				getAccountSettingsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudShellService.GetAccountSettingsByID(getAccountSettingsByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountSettingsByID with error: Operation validation and request error`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsByIdOptions model
				getAccountSettingsByIdOptionsModel := new(ibmcloudshellv1.GetAccountSettingsByIdOptions)
				getAccountSettingsByIdOptionsModel.AccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				getAccountSettingsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudShellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudShellService.GetAccountSettingsByID(getAccountSettingsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountSettingsByIdOptions model with no property values
				getAccountSettingsByIdOptionsModelNew := new(ibmcloudshellv1.GetAccountSettingsByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudShellService.GetAccountSettingsByID(getAccountSettingsByIdOptionsModelNew)
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
			It(`Invoke GetAccountSettingsByID successfully`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsByIdOptions model
				getAccountSettingsByIdOptionsModel := new(ibmcloudshellv1.GetAccountSettingsByIdOptions)
				getAccountSettingsByIdOptionsModel.AccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				getAccountSettingsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudShellService.GetAccountSettingsByID(getAccountSettingsByIdOptionsModel)
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
	Describe(`UpdateAccountSettingsByID(updateAccountSettingsByIdOptions *UpdateAccountSettingsByIdOptions) - Operation response error`, func() {
		updateAccountSettingsByIDPath := "/api/v1/user/accounts/12345678-abcd-1a2b-a1b2-1234567890ab/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsByIDPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountSettingsByID with error: Operation response processing error`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				featureModel := []ibmcloudshellv1.Feature{
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("server.file_manager"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("server.web_preview"),
					},
				}

				regionSettingModel := []ibmcloudshellv1.RegionSetting{
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("eu-de"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("jp-tok"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("us-south"),
					},
				}

				// Construct an instance of the UpdateAccountSettingsByIdOptions model
				updateAccountSettingsByIdOptionsModel := new(ibmcloudshellv1.UpdateAccountSettingsByIdOptions)
				updateAccountSettingsByIdOptionsModel.AccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewID = core.StringPtr("ac-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewRev = core.StringPtr("130-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewAccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewCreatedAt = core.Int64Ptr(int64(1600079615))
				updateAccountSettingsByIdOptionsModel.NewCreatedBy = core.StringPtr("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewFeatures = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewRegions = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewEnabled = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewFeatures = featureModel
				updateAccountSettingsByIdOptionsModel.NewRegions = regionSettingModel
				updateAccountSettingsByIdOptionsModel.NewType = core.StringPtr("account_settings")
				updateAccountSettingsByIdOptionsModel.NewUpdatedAt = core.Int64Ptr(int64(1624359948))
				updateAccountSettingsByIdOptionsModel.NewUpdatedBy = core.StringPtr("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudShellService.UpdateAccountSettingsByID(updateAccountSettingsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudShellService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudShellService.UpdateAccountSettingsByID(updateAccountSettingsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountSettingsByID(updateAccountSettingsByIdOptions *UpdateAccountSettingsByIdOptions)`, func() {
		updateAccountSettingsByIDPath := "/api/v1/user/accounts/12345678-abcd-1a2b-a1b2-1234567890ab/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsByIDPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_id": "ID", "_rev": "Rev", "account_id": "AccountID", "created_at": 9, "created_by": "CreatedBy", "default_enable_new_features": true, "default_enable_new_regions": false, "enabled": false, "features": [{"enabled": false, "key": "Key"}], "regions": [{"enabled": false, "key": "Key"}], "type": "Type", "updated_at": 9, "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke UpdateAccountSettingsByID successfully with retries`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())
				ibmCloudShellService.EnableRetries(0, 0)

				featureModel := []ibmcloudshellv1.Feature{
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("server.file_manager"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("server.web_preview"),
					},
				}

				regionSettingModel := []ibmcloudshellv1.RegionSetting{
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("eu-de"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("jp-tok"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("us-south"),
					},
				}

				// Construct an instance of the UpdateAccountSettingsByIdOptions model
				updateAccountSettingsByIdOptionsModel := new(ibmcloudshellv1.UpdateAccountSettingsByIdOptions)
				updateAccountSettingsByIdOptionsModel.AccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewID = core.StringPtr("ac-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewRev = core.StringPtr("130-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewAccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewCreatedAt = core.Int64Ptr(int64(1600079615))
				updateAccountSettingsByIdOptionsModel.NewCreatedBy = core.StringPtr("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewFeatures = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewRegions = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewEnabled = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewFeatures = featureModel
				updateAccountSettingsByIdOptionsModel.NewRegions = regionSettingModel
				updateAccountSettingsByIdOptionsModel.NewType = core.StringPtr("account_settings")
				updateAccountSettingsByIdOptionsModel.NewUpdatedAt = core.Int64Ptr(int64(1624359948))
				updateAccountSettingsByIdOptionsModel.NewUpdatedBy = core.StringPtr("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudShellService.UpdateAccountSettingsByIDWithContext(ctx, updateAccountSettingsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudShellService.DisableRetries()
				result, response, operationErr := ibmCloudShellService.UpdateAccountSettingsByID(updateAccountSettingsByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudShellService.UpdateAccountSettingsByIDWithContext(ctx, updateAccountSettingsByIdOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsByIDPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_id": "ID", "_rev": "Rev", "account_id": "AccountID", "created_at": 9, "created_by": "CreatedBy", "default_enable_new_features": true, "default_enable_new_regions": false, "enabled": false, "features": [{"enabled": false, "key": "Key"}], "regions": [{"enabled": false, "key": "Key"}], "type": "Type", "updated_at": 9, "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke UpdateAccountSettingsByID successfully`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudShellService.UpdateAccountSettingsByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				featureModel := []ibmcloudshellv1.Feature{
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("server.file_manager"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("server.web_preview"),
					},
				}

				regionSettingModel := []ibmcloudshellv1.RegionSetting{
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("eu-de"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("jp-tok"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("us-south"),
					},
				}

				// Construct an instance of the UpdateAccountSettingsByIdOptions model
				updateAccountSettingsByIdOptionsModel := new(ibmcloudshellv1.UpdateAccountSettingsByIdOptions)
				updateAccountSettingsByIdOptionsModel.AccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewID = core.StringPtr("ac-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewRev = core.StringPtr("130-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewAccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewCreatedAt = core.Int64Ptr(int64(1600079615))
				updateAccountSettingsByIdOptionsModel.NewCreatedBy = core.StringPtr("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewFeatures = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewRegions = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewEnabled = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewFeatures = featureModel
				updateAccountSettingsByIdOptionsModel.NewRegions = regionSettingModel
				updateAccountSettingsByIdOptionsModel.NewType = core.StringPtr("account_settings")
				updateAccountSettingsByIdOptionsModel.NewUpdatedAt = core.Int64Ptr(int64(1624359948))
				updateAccountSettingsByIdOptionsModel.NewUpdatedBy = core.StringPtr("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudShellService.UpdateAccountSettingsByID(updateAccountSettingsByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccountSettingsByID with error: Operation validation and request error`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				featureModel := []ibmcloudshellv1.Feature{
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("server.file_manager"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("server.web_preview"),
					},
				}

				regionSettingModel := []ibmcloudshellv1.RegionSetting{
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("eu-de"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("jp-tok"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("us-south"),
					},
				}

				// Construct an instance of the UpdateAccountSettingsByIdOptions model
				updateAccountSettingsByIdOptionsModel := new(ibmcloudshellv1.UpdateAccountSettingsByIdOptions)
				updateAccountSettingsByIdOptionsModel.AccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewID = core.StringPtr("ac-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewRev = core.StringPtr("130-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewAccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewCreatedAt = core.Int64Ptr(int64(1600079615))
				updateAccountSettingsByIdOptionsModel.NewCreatedBy = core.StringPtr("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewFeatures = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewRegions = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewEnabled = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewFeatures = featureModel
				updateAccountSettingsByIdOptionsModel.NewRegions = regionSettingModel
				updateAccountSettingsByIdOptionsModel.NewType = core.StringPtr("account_settings")
				updateAccountSettingsByIdOptionsModel.NewUpdatedAt = core.Int64Ptr(int64(1624359948))
				updateAccountSettingsByIdOptionsModel.NewUpdatedBy = core.StringPtr("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudShellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudShellService.UpdateAccountSettingsByID(updateAccountSettingsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountSettingsByIdOptions model with no property values
				updateAccountSettingsByIdOptionsModelNew := new(ibmcloudshellv1.UpdateAccountSettingsByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudShellService.UpdateAccountSettingsByID(updateAccountSettingsByIdOptionsModelNew)
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
			It(`Invoke UpdateAccountSettingsByID successfully`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				featureModel := []ibmcloudshellv1.Feature{
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("server.file_manager"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("server.web_preview"),
					},
				}

				regionSettingModel := []ibmcloudshellv1.RegionSetting{
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("eu-de"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("jp-tok"),
					},
					{
						Enabled: core.BoolPtr(true),
						Key:     core.StringPtr("us-south"),
					},
				}

				// Construct an instance of the UpdateAccountSettingsByIdOptions model
				updateAccountSettingsByIdOptionsModel := new(ibmcloudshellv1.UpdateAccountSettingsByIdOptions)
				updateAccountSettingsByIdOptionsModel.AccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewID = core.StringPtr("ac-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewRev = core.StringPtr("130-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewAccountID = core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.NewCreatedAt = core.Int64Ptr(int64(1600079615))
				updateAccountSettingsByIdOptionsModel.NewCreatedBy = core.StringPtr("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewFeatures = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewRegions = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewEnabled = core.BoolPtr(true)
				updateAccountSettingsByIdOptionsModel.NewFeatures = featureModel
				updateAccountSettingsByIdOptionsModel.NewRegions = regionSettingModel
				updateAccountSettingsByIdOptionsModel.NewType = core.StringPtr("account_settings")
				updateAccountSettingsByIdOptionsModel.NewUpdatedAt = core.Int64Ptr(int64(1624359948))
				updateAccountSettingsByIdOptionsModel.NewUpdatedBy = core.StringPtr("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudShellService.UpdateAccountSettingsByID(updateAccountSettingsByIdOptionsModel)
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
			ibmCloudShellService, _ := ibmcloudshellv1.NewIbmCloudShellV1(&ibmcloudshellv1.IbmCloudShellV1Options{
				URL:           "http://ibmcloudshellv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetAccountSettingsByIdOptions successfully`, func() {
				// Construct an instance of the GetAccountSettingsByIdOptions model
				accountID := "12345678-abcd-1a2b-a1b2-1234567890ab"
				getAccountSettingsByIdOptionsModel := ibmCloudShellService.NewGetAccountSettingsByIdOptions(accountID)
				getAccountSettingsByIdOptionsModel.SetAccountID("12345678-abcd-1a2b-a1b2-1234567890ab")
				getAccountSettingsByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountSettingsByIdOptionsModel).ToNot(BeNil())
				Expect(getAccountSettingsByIdOptionsModel.AccountID).To(Equal(core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")))
				Expect(getAccountSettingsByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountSettingsByIdOptions successfully`, func() {
				// Construct an instance of the Feature model
				featureModel := new(ibmcloudshellv1.Feature)
				Expect(featureModel).ToNot(BeNil())
				featureModel.Enabled = core.BoolPtr(true)
				featureModel.Key = core.StringPtr("server.file_manager")
				Expect(featureModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(featureModel.Key).To(Equal(core.StringPtr("server.file_manager")))

				// Construct an instance of the RegionSetting model
				regionSettingModel := new(ibmcloudshellv1.RegionSetting)
				Expect(regionSettingModel).ToNot(BeNil())
				regionSettingModel.Enabled = core.BoolPtr(true)
				regionSettingModel.Key = core.StringPtr("eu-de")
				Expect(regionSettingModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(regionSettingModel.Key).To(Equal(core.StringPtr("eu-de")))

				// Construct an instance of the UpdateAccountSettingsByIdOptions model
				accountID := "12345678-abcd-1a2b-a1b2-1234567890ab"
				updateAccountSettingsByIdOptionsModel := ibmCloudShellService.NewUpdateAccountSettingsByIdOptions(accountID)
				updateAccountSettingsByIdOptionsModel.SetAccountID("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.SetNewID("ac-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.SetNewRev("130-12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.SetNewAccountID("12345678-abcd-1a2b-a1b2-1234567890ab")
				updateAccountSettingsByIdOptionsModel.SetNewCreatedAt(int64(1600079615))
				updateAccountSettingsByIdOptionsModel.SetNewCreatedBy("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.SetNewDefaultEnableNewFeatures(true)
				updateAccountSettingsByIdOptionsModel.SetNewDefaultEnableNewRegions(true)
				updateAccountSettingsByIdOptionsModel.SetNewEnabled(true)
				updateAccountSettingsByIdOptionsModel.SetNewFeatures([]ibmcloudshellv1.Feature{*featureModel})
				updateAccountSettingsByIdOptionsModel.SetNewRegions([]ibmcloudshellv1.RegionSetting{*regionSettingModel})
				updateAccountSettingsByIdOptionsModel.SetNewType("account_settings")
				updateAccountSettingsByIdOptionsModel.SetNewUpdatedAt(int64(1624359948))
				updateAccountSettingsByIdOptionsModel.SetNewUpdatedBy("IBMid-1000000000")
				updateAccountSettingsByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountSettingsByIdOptionsModel).ToNot(BeNil())
				Expect(updateAccountSettingsByIdOptionsModel.AccountID).To(Equal(core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")))
				Expect(updateAccountSettingsByIdOptionsModel.NewID).To(Equal(core.StringPtr("ac-12345678-abcd-1a2b-a1b2-1234567890ab")))
				Expect(updateAccountSettingsByIdOptionsModel.NewRev).To(Equal(core.StringPtr("130-12345678-abcd-1a2b-a1b2-1234567890ab")))
				Expect(updateAccountSettingsByIdOptionsModel.NewAccountID).To(Equal(core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab")))
				Expect(updateAccountSettingsByIdOptionsModel.NewCreatedAt).To(Equal(core.Int64Ptr(int64(1600079615))))
				Expect(updateAccountSettingsByIdOptionsModel.NewCreatedBy).To(Equal(core.StringPtr("IBMid-1000000000")))
				Expect(updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewFeatures).To(Equal(core.BoolPtr(true)))
				Expect(updateAccountSettingsByIdOptionsModel.NewDefaultEnableNewRegions).To(Equal(core.BoolPtr(true)))
				Expect(updateAccountSettingsByIdOptionsModel.NewEnabled).To(Equal(core.BoolPtr(true)))
				Expect(updateAccountSettingsByIdOptionsModel.NewFeatures).To(Equal([]ibmcloudshellv1.Feature{*featureModel}))
				Expect(updateAccountSettingsByIdOptionsModel.NewRegions).To(Equal([]ibmcloudshellv1.RegionSetting{*regionSettingModel}))
				Expect(updateAccountSettingsByIdOptionsModel.NewType).To(Equal(core.StringPtr("account_settings")))
				Expect(updateAccountSettingsByIdOptionsModel.NewUpdatedAt).To(Equal(core.Int64Ptr(int64(1624359948))))
				Expect(updateAccountSettingsByIdOptionsModel.NewUpdatedBy).To(Equal(core.StringPtr("IBMid-1000000000")))
				Expect(updateAccountSettingsByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
