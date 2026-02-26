/**
 * (C) Copyright IBM Corp. 2026.
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

package accountmanagementv4_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/accountmanagementv4"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`AccountManagementV4`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(accountManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(accountManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
				URL: "https://accountmanagementv4/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(accountManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ACCOUNT_MANAGEMENT_URL":       "https://accountmanagementv4/api",
				"ACCOUNT_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4UsingExternalConfig(&accountmanagementv4.AccountManagementV4Options{})
				Expect(accountManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := accountManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != accountManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(accountManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(accountManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4UsingExternalConfig(&accountmanagementv4.AccountManagementV4Options{
					URL: "https://testService/api",
				})
				Expect(accountManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(accountManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := accountManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != accountManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(accountManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(accountManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4UsingExternalConfig(&accountmanagementv4.AccountManagementV4Options{})
				err := accountManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(accountManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(accountManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := accountManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != accountManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(accountManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(accountManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ACCOUNT_MANAGEMENT_URL":       "https://accountmanagementv4/api",
				"ACCOUNT_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4UsingExternalConfig(&accountmanagementv4.AccountManagementV4Options{})

			It(`Instantiate service client with error`, func() {
				Expect(accountManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ACCOUNT_MANAGEMENT_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4UsingExternalConfig(&accountmanagementv4.AccountManagementV4Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(accountManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = accountmanagementv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAccount(getAccountOptions *GetAccountOptions) - Operation response error`, func() {
		getAccountPath := "/v4/accounts/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccount with error: Operation response processing error`, func() {
				accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(accountManagementService).ToNot(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(accountmanagementv4.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := accountManagementService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				accountManagementService.EnableRetries(0, 0)
				result, response, operationErr = accountManagementService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccount(getAccountOptions *GetAccountOptions)`, func() {
		getAccountPath := "/v4/accounts/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "owner": "Owner", "owner_userid": "OwnerUserid", "owner_iamid": "OwnerIamid", "type": "Type", "status": "Status", "linked_softlayer_account": "LinkedSoftlayerAccount", "team_directory_enabled": false, "traits": {"eu_supported": false, "poc": false, "hippa": false}}`)
				}))
			})
			It(`Invoke GetAccount successfully with retries`, func() {
				accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(accountManagementService).ToNot(BeNil())
				accountManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(accountmanagementv4.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := accountManagementService.GetAccountWithContext(ctx, getAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				accountManagementService.DisableRetries()
				result, response, operationErr := accountManagementService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = accountManagementService.GetAccountWithContext(ctx, getAccountOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "owner": "Owner", "owner_userid": "OwnerUserid", "owner_iamid": "OwnerIamid", "type": "Type", "status": "Status", "linked_softlayer_account": "LinkedSoftlayerAccount", "team_directory_enabled": false, "traits": {"eu_supported": false, "poc": false, "hippa": false}}`)
				}))
			})
			It(`Invoke GetAccount successfully`, func() {
				accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(accountManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := accountManagementService.GetAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(accountmanagementv4.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = accountManagementService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccount with error: Operation validation and request error`, func() {
				accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(accountManagementService).ToNot(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(accountmanagementv4.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := accountManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := accountManagementService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountOptions model with no property values
				getAccountOptionsModelNew := new(accountmanagementv4.GetAccountOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = accountManagementService.GetAccount(getAccountOptionsModelNew)
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
			It(`Invoke GetAccount successfully`, func() {
				accountManagementService, serviceErr := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(accountManagementService).ToNot(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(accountmanagementv4.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := accountManagementService.GetAccount(getAccountOptionsModel)
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
			accountManagementService, _ := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
				URL:           "http://accountmanagementv4modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetAccountOptions successfully`, func() {
				// Construct an instance of the GetAccountOptions model
				accountID := "testString"
				getAccountOptionsModel := accountManagementService.NewGetAccountOptions(accountID)
				getAccountOptionsModel.SetAccountID("testString")
				getAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountOptionsModel).ToNot(BeNil())
				Expect(getAccountOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
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
