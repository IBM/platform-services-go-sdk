/**
 * (C) Copyright IBM Corp. 2023.
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

package projectsv1_test

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
	"github.com/IBM/platform-services-go-sdk/projectsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ProjectsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(projectsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(projectsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
				URL: "https://projectsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(projectsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PROJECTS_URL": "https://projectsv1/api",
				"PROJECTS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				projectsService, serviceErr := projectsv1.NewProjectsV1UsingExternalConfig(&projectsv1.ProjectsV1Options{
				})
				Expect(projectsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := projectsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != projectsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(projectsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(projectsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				projectsService, serviceErr := projectsv1.NewProjectsV1UsingExternalConfig(&projectsv1.ProjectsV1Options{
					URL: "https://testService/api",
				})
				Expect(projectsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(projectsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := projectsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != projectsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(projectsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(projectsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				projectsService, serviceErr := projectsv1.NewProjectsV1UsingExternalConfig(&projectsv1.ProjectsV1Options{
				})
				err := projectsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(projectsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := projectsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != projectsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(projectsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(projectsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PROJECTS_URL": "https://projectsv1/api",
				"PROJECTS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			projectsService, serviceErr := projectsv1.NewProjectsV1UsingExternalConfig(&projectsv1.ProjectsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(projectsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PROJECTS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			projectsService, serviceErr := projectsv1.NewProjectsV1UsingExternalConfig(&projectsv1.ProjectsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(projectsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = projectsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateProject(createProjectOptions *CreateProjectOptions) - Operation response error`, func() {
		createProjectPath := "/v1/projects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"Default"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"us-south"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProject with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProject(createProjectOptions *CreateProjectOptions)`, func() {
		createProjectPath := "/v1/projects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
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

					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"Default"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"us-south"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "Crn", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description"}]}`)
				}))
			})
			It(`Invoke CreateProject successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.CreateProjectWithContext(ctx, createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.CreateProjectWithContext(ctx, createProjectOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
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

					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"Default"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"us-south"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "Crn", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description"}]}`)
				}))
			})
			It(`Invoke CreateProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.CreateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProject with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProjectOptions model with no property values
				createProjectOptionsModelNew := new(projectsv1.CreateProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.CreateProject(createProjectOptionsModelNew)
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
			It(`Invoke CreateProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.CreateProject(createProjectOptionsModel)
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
	Describe(`ListProjects(listProjectsOptions *ListProjectsOptions) - Operation response error`, func() {
		listProjectsPath := "/v1/projects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProjects with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectsv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProjects(listProjectsOptions *ListProjectsOptions)`, func() {
		listProjectsPath := "/v1/projects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 1, "total_count": 0, "first": {"href": "Href", "start": "Start"}, "last": {"href": "Href", "start": "Start"}, "previous": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "projects": [{"name": "Name", "id": "ID", "definition": {"name": "Name", "description": "Description", "id": "ID", "crn": "Crn", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description"}]}, "state": "State"}]}`)
				}))
			})
			It(`Invoke ListProjects successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectsv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.ListProjectsWithContext(ctx, listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.ListProjectsWithContext(ctx, listProjectsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 1, "total_count": 0, "first": {"href": "Href", "start": "Start"}, "last": {"href": "Href", "start": "Start"}, "previous": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "projects": [{"name": "Name", "id": "ID", "definition": {"name": "Name", "description": "Description", "id": "ID", "crn": "Crn", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description"}]}, "state": "State"}]}`)
				}))
			})
			It(`Invoke ListProjects successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.ListProjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectsv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProjects with error: Operation request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectsv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke ListProjects successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectsv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(projectsv1.ProjectListResponseSchema)
				nextObject := new(projectsv1.PaginationLink)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(projectsv1.ProjectListResponseSchema)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"projects":[{"name":"Name","id":"ID","definition":{"name":"Name","description":"Description","id":"ID","crn":"Crn","configs":[{"id":"ID","name":"Name","labels":["Labels"],"description":"Description"}]},"state":"State"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"projects":[{"name":"Name","id":"ID","definition":{"name":"Name","description":"Description","id":"ID","crn":"Crn","configs":[{"id":"ID","name":"Name","labels":["Labels"],"description":"Description"}]},"state":"State"}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ProjectsPager.GetNext successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				listProjectsOptionsModel := &projectsv1.ListProjectsOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := projectsService.NewProjectsPager(listProjectsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []projectsv1.ProjectListResponse
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ProjectsPager.GetAll successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				listProjectsOptionsModel := &projectsv1.ListProjectsOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := projectsService.NewProjectsPager(listProjectsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetProject(getProjectOptions *GetProjectOptions) - Operation response error`, func() {
		getProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["branch"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProject with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectsv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.Branch = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProject(getProjectOptions *GetProjectOptions)`, func() {
		getProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["branch"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "Crn", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description"}]}`)
				}))
			})
			It(`Invoke GetProject successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectsv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.Branch = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetProjectWithContext(ctx, getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetProjectWithContext(ctx, getProjectOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["branch"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "Crn", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description"}]}`)
				}))
			})
			It(`Invoke GetProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectsv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.Branch = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProject with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectsv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.Branch = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectOptions model with no property values
				getProjectOptionsModelNew := new(projectsv1.GetProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetProject(getProjectOptionsModelNew)
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
			It(`Invoke GetProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectsv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.Branch = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetProject(getProjectOptionsModel)
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
	Describe(`ReplaceProject(replaceProjectOptions *ReplaceProjectOptions) - Operation response error`, func() {
		replaceProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceProjectPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["branch"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceProject with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ReplaceProjectOptions model
				replaceProjectOptionsModel := new(projectsv1.ReplaceProjectOptions)
				replaceProjectOptionsModel.ID = core.StringPtr("testString")
				replaceProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				replaceProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				replaceProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				replaceProjectOptionsModel.Branch = core.StringPtr("testString")
				replaceProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.ReplaceProject(replaceProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.ReplaceProject(replaceProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceProject(replaceProjectOptions *ReplaceProjectOptions)`, func() {
		replaceProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceProjectPath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.URL.Query()["branch"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"branch": "Branch", "pr_url": "PrURL"}`)
				}))
			})
			It(`Invoke ReplaceProject successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ReplaceProjectOptions model
				replaceProjectOptionsModel := new(projectsv1.ReplaceProjectOptions)
				replaceProjectOptionsModel.ID = core.StringPtr("testString")
				replaceProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				replaceProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				replaceProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				replaceProjectOptionsModel.Branch = core.StringPtr("testString")
				replaceProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.ReplaceProjectWithContext(ctx, replaceProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.ReplaceProject(replaceProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.ReplaceProjectWithContext(ctx, replaceProjectOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceProjectPath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.URL.Query()["branch"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"branch": "Branch", "pr_url": "PrURL"}`)
				}))
			})
			It(`Invoke ReplaceProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.ReplaceProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ReplaceProjectOptions model
				replaceProjectOptionsModel := new(projectsv1.ReplaceProjectOptions)
				replaceProjectOptionsModel.ID = core.StringPtr("testString")
				replaceProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				replaceProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				replaceProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				replaceProjectOptionsModel.Branch = core.StringPtr("testString")
				replaceProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.ReplaceProject(replaceProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceProject with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ReplaceProjectOptions model
				replaceProjectOptionsModel := new(projectsv1.ReplaceProjectOptions)
				replaceProjectOptionsModel.ID = core.StringPtr("testString")
				replaceProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				replaceProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				replaceProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				replaceProjectOptionsModel.Branch = core.StringPtr("testString")
				replaceProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.ReplaceProject(replaceProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceProjectOptions model with no property values
				replaceProjectOptionsModelNew := new(projectsv1.ReplaceProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.ReplaceProject(replaceProjectOptionsModelNew)
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
			It(`Invoke ReplaceProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ReplaceProjectOptions model
				replaceProjectOptionsModel := new(projectsv1.ReplaceProjectOptions)
				replaceProjectOptionsModel.ID = core.StringPtr("testString")
				replaceProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				replaceProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				replaceProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				replaceProjectOptionsModel.Branch = core.StringPtr("testString")
				replaceProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.ReplaceProject(replaceProjectOptionsModel)
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
	Describe(`DeleteProject(deleteProjectOptions *DeleteProjectOptions)`, func() {
		deleteProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProjectPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["branch"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.DeleteProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProjectOptions model
				deleteProjectOptionsModel := new(projectsv1.DeleteProjectOptions)
				deleteProjectOptionsModel.ID = core.StringPtr("testString")
				deleteProjectOptionsModel.Branch = core.StringPtr("testString")
				deleteProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.DeleteProject(deleteProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProject with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeleteProjectOptions model
				deleteProjectOptionsModel := new(projectsv1.DeleteProjectOptions)
				deleteProjectOptionsModel.ID = core.StringPtr("testString")
				deleteProjectOptionsModel.Branch = core.StringPtr("testString")
				deleteProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.DeleteProject(deleteProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProjectOptions model with no property values
				deleteProjectOptionsModelNew := new(projectsv1.DeleteProjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.DeleteProject(deleteProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectFile(getProjectFileOptions *GetProjectFileOptions)`, func() {
		getProjectFilePath := "/v1/projects/testString/files/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectFilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["branch"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/plain")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"PGgxPlRoaXMgaXMgYSB0ZXN0PC9oMT4="`)
				}))
			})
			It(`Invoke GetProjectFile successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectFileOptions model
				getProjectFileOptionsModel := new(projectsv1.GetProjectFileOptions)
				getProjectFileOptionsModel.ID = core.StringPtr("testString")
				getProjectFileOptionsModel.FilePath = core.StringPtr("testString")
				getProjectFileOptionsModel.Branch = core.StringPtr("testString")
				getProjectFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetProjectFileWithContext(ctx, getProjectFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetProjectFile(getProjectFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetProjectFileWithContext(ctx, getProjectFileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProjectFilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["branch"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "text/plain")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"PGgxPlRoaXMgaXMgYSB0ZXN0PC9oMT4="`)
				}))
			})
			It(`Invoke GetProjectFile successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetProjectFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectFileOptions model
				getProjectFileOptionsModel := new(projectsv1.GetProjectFileOptions)
				getProjectFileOptionsModel.ID = core.StringPtr("testString")
				getProjectFileOptionsModel.FilePath = core.StringPtr("testString")
				getProjectFileOptionsModel.Branch = core.StringPtr("testString")
				getProjectFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetProjectFile(getProjectFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProjectFile with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectFileOptions model
				getProjectFileOptionsModel := new(projectsv1.GetProjectFileOptions)
				getProjectFileOptionsModel.ID = core.StringPtr("testString")
				getProjectFileOptionsModel.FilePath = core.StringPtr("testString")
				getProjectFileOptionsModel.Branch = core.StringPtr("testString")
				getProjectFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetProjectFile(getProjectFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectFileOptions model with no property values
				getProjectFileOptionsModelNew := new(projectsv1.GetProjectFileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetProjectFile(getProjectFileOptionsModelNew)
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
			It(`Invoke GetProjectFile successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectFileOptions model
				getProjectFileOptionsModel := new(projectsv1.GetProjectFileOptions)
				getProjectFileOptionsModel.ID = core.StringPtr("testString")
				getProjectFileOptionsModel.FilePath = core.StringPtr("testString")
				getProjectFileOptionsModel.Branch = core.StringPtr("testString")
				getProjectFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetProjectFile(getProjectFileOptionsModel)
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
	Describe(`MergeProject(mergeProjectOptions *MergeProjectOptions) - Operation response error`, func() {
		mergeProjectPath := "/v1/projects/testString/merge"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(mergeProjectPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke MergeProject with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the MergeProjectOptions model
				mergeProjectOptionsModel := new(projectsv1.MergeProjectOptions)
				mergeProjectOptionsModel.ID = core.StringPtr("testString")
				mergeProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				mergeProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				mergeProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.MergeProject(mergeProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.MergeProject(mergeProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`MergeProject(mergeProjectOptions *MergeProjectOptions)`, func() {
		mergeProjectPath := "/v1/projects/testString/merge"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(mergeProjectPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "definition": {"name": "Name", "description": "Description", "id": "ID", "crn": "Crn", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description"}]}}`)
				}))
			})
			It(`Invoke MergeProject successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the MergeProjectOptions model
				mergeProjectOptionsModel := new(projectsv1.MergeProjectOptions)
				mergeProjectOptionsModel.ID = core.StringPtr("testString")
				mergeProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				mergeProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				mergeProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.MergeProjectWithContext(ctx, mergeProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.MergeProject(mergeProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.MergeProjectWithContext(ctx, mergeProjectOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(mergeProjectPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "definition": {"name": "Name", "description": "Description", "id": "ID", "crn": "Crn", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description"}]}}`)
				}))
			})
			It(`Invoke MergeProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.MergeProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the MergeProjectOptions model
				mergeProjectOptionsModel := new(projectsv1.MergeProjectOptions)
				mergeProjectOptionsModel.ID = core.StringPtr("testString")
				mergeProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				mergeProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				mergeProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.MergeProject(mergeProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke MergeProject with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the MergeProjectOptions model
				mergeProjectOptionsModel := new(projectsv1.MergeProjectOptions)
				mergeProjectOptionsModel.ID = core.StringPtr("testString")
				mergeProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				mergeProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				mergeProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.MergeProject(mergeProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the MergeProjectOptions model with no property values
				mergeProjectOptionsModelNew := new(projectsv1.MergeProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.MergeProject(mergeProjectOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke MergeProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the MergeProjectOptions model
				mergeProjectOptionsModel := new(projectsv1.MergeProjectOptions)
				mergeProjectOptionsModel.ID = core.StringPtr("testString")
				mergeProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				mergeProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				mergeProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.MergeProject(mergeProjectOptionsModel)
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
	Describe(`ValidateProject(validateProjectOptions *ValidateProjectOptions)`, func() {
		validateProjectPath := "/v1/projects/validate"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateProjectPath))
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

					res.WriteHeader(204)
				}))
			})
			It(`Invoke ValidateProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.ValidateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ValidateProjectOptions model
				validateProjectOptionsModel := new(projectsv1.ValidateProjectOptions)
				validateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				validateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				validateProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				validateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.ValidateProject(validateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ValidateProject with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ValidateProjectOptions model
				validateProjectOptionsModel := new(projectsv1.ValidateProjectOptions)
				validateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				validateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				validateProjectOptionsModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				validateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.ValidateProject(validateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ValidateProjectOptions model with no property values
				validateProjectOptionsModelNew := new(projectsv1.ValidateProjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.ValidateProject(validateProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ConfigChanges(configChangesOptions *ConfigChangesOptions) - Operation response error`, func() {
		configChangesPath := "/v1/projects/234234324-3444-4556-224232432/config_changes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(configChangesPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ConfigChanges with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the ConfigChangesOptions model
				configChangesOptionsModel := new(projectsv1.ConfigChangesOptions)
				configChangesOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				configChangesOptionsModel.Source = projectInputModel
				configChangesOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				configChangesOptionsModel.Target = projectInputModel
				configChangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.ConfigChanges(configChangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.ConfigChanges(configChangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ConfigChanges(configChangesOptions *ConfigChangesOptions)`, func() {
		configChangesPath := "/v1/projects/234234324-3444-4556-224232432/config_changes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(configChangesPath))
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
					fmt.Fprintf(res, "%s", `{"added": [{"name": "Name"}], "deleted": [{"name": "Name"}], "changed": [{"name": "Name", "new_name": "NewName"}]}`)
				}))
			})
			It(`Invoke ConfigChanges successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the ConfigChangesOptions model
				configChangesOptionsModel := new(projectsv1.ConfigChangesOptions)
				configChangesOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				configChangesOptionsModel.Source = projectInputModel
				configChangesOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				configChangesOptionsModel.Target = projectInputModel
				configChangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.ConfigChangesWithContext(ctx, configChangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.ConfigChanges(configChangesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.ConfigChangesWithContext(ctx, configChangesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(configChangesPath))
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
					fmt.Fprintf(res, "%s", `{"added": [{"name": "Name"}], "deleted": [{"name": "Name"}], "changed": [{"name": "Name", "new_name": "NewName"}]}`)
				}))
			})
			It(`Invoke ConfigChanges successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.ConfigChanges(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the ConfigChangesOptions model
				configChangesOptionsModel := new(projectsv1.ConfigChangesOptions)
				configChangesOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				configChangesOptionsModel.Source = projectInputModel
				configChangesOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				configChangesOptionsModel.Target = projectInputModel
				configChangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.ConfigChanges(configChangesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ConfigChanges with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the ConfigChangesOptions model
				configChangesOptionsModel := new(projectsv1.ConfigChangesOptions)
				configChangesOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				configChangesOptionsModel.Source = projectInputModel
				configChangesOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				configChangesOptionsModel.Target = projectInputModel
				configChangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.ConfigChanges(configChangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ConfigChangesOptions model with no property values
				configChangesOptionsModelNew := new(projectsv1.ConfigChangesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.ConfigChanges(configChangesOptionsModelNew)
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
			It(`Invoke ConfigChanges successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the ConfigChangesOptions model
				configChangesOptionsModel := new(projectsv1.ConfigChangesOptions)
				configChangesOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				configChangesOptionsModel.Source = projectInputModel
				configChangesOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				configChangesOptionsModel.Target = projectInputModel
				configChangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.ConfigChanges(configChangesOptionsModel)
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
	Describe(`PlanConfig(planConfigOptions *PlanConfigOptions) - Operation response error`, func() {
		planConfigPath := "/v1/projects/testString/configs/testString/plan"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(planConfigPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PlanConfig with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the PlanConfigOptions model
				planConfigOptionsModel := new(projectsv1.PlanConfigOptions)
				planConfigOptionsModel.ID = core.StringPtr("testString")
				planConfigOptionsModel.ConfigName = core.StringPtr("testString")
				planConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.PlanConfig(planConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.PlanConfig(planConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PlanConfig(planConfigOptions *PlanConfigOptions)`, func() {
		planConfigPath := "/v1/projects/testString/configs/testString/plan"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(planConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"name": "Name", "job": "Job", "workspace": "Workspace", "cart_order": "CartOrder", "catalog_error": "CatalogError", "catalog_status_code": 17, "schematics_error": "SchematicsError", "schematics_status_code": 20, "schematics_submitted_at": 21}`)
				}))
			})
			It(`Invoke PlanConfig successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the PlanConfigOptions model
				planConfigOptionsModel := new(projectsv1.PlanConfigOptions)
				planConfigOptionsModel.ID = core.StringPtr("testString")
				planConfigOptionsModel.ConfigName = core.StringPtr("testString")
				planConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.PlanConfigWithContext(ctx, planConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.PlanConfig(planConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.PlanConfigWithContext(ctx, planConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(planConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"name": "Name", "job": "Job", "workspace": "Workspace", "cart_order": "CartOrder", "catalog_error": "CatalogError", "catalog_status_code": 17, "schematics_error": "SchematicsError", "schematics_status_code": 20, "schematics_submitted_at": 21}`)
				}))
			})
			It(`Invoke PlanConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.PlanConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PlanConfigOptions model
				planConfigOptionsModel := new(projectsv1.PlanConfigOptions)
				planConfigOptionsModel.ID = core.StringPtr("testString")
				planConfigOptionsModel.ConfigName = core.StringPtr("testString")
				planConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.PlanConfig(planConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PlanConfig with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the PlanConfigOptions model
				planConfigOptionsModel := new(projectsv1.PlanConfigOptions)
				planConfigOptionsModel.ID = core.StringPtr("testString")
				planConfigOptionsModel.ConfigName = core.StringPtr("testString")
				planConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.PlanConfig(planConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PlanConfigOptions model with no property values
				planConfigOptionsModelNew := new(projectsv1.PlanConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.PlanConfig(planConfigOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke PlanConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the PlanConfigOptions model
				planConfigOptionsModel := new(projectsv1.PlanConfigOptions)
				planConfigOptionsModel.ID = core.StringPtr("testString")
				planConfigOptionsModel.ConfigName = core.StringPtr("testString")
				planConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.PlanConfig(planConfigOptionsModel)
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
	Describe(`InstallConfig(installConfigOptions *InstallConfigOptions)`, func() {
		installConfigPath := "/v1/projects/testString/configs/testString/install"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(installConfigPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke InstallConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.InstallConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectsv1.InstallConfigOptions)
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.ConfigName = core.StringPtr("testString")
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke InstallConfig with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectsv1.InstallConfigOptions)
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.ConfigName = core.StringPtr("testString")
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the InstallConfigOptions model with no property values
				installConfigOptionsModelNew := new(projectsv1.InstallConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.InstallConfig(installConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UninstallConfig(uninstallConfigOptions *UninstallConfigOptions)`, func() {
		uninstallConfigPath := "/v1/projects/testString/configs/testString/uninstall"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uninstallConfigPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UninstallConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.UninstallConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UninstallConfigOptions model
				uninstallConfigOptionsModel := new(projectsv1.UninstallConfigOptions)
				uninstallConfigOptionsModel.ID = core.StringPtr("testString")
				uninstallConfigOptionsModel.ConfigName = core.StringPtr("testString")
				uninstallConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.UninstallConfig(uninstallConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UninstallConfig with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UninstallConfigOptions model
				uninstallConfigOptionsModel := new(projectsv1.UninstallConfigOptions)
				uninstallConfigOptionsModel.ID = core.StringPtr("testString")
				uninstallConfigOptionsModel.ConfigName = core.StringPtr("testString")
				uninstallConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.UninstallConfig(uninstallConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UninstallConfigOptions model with no property values
				uninstallConfigOptionsModelNew := new(projectsv1.UninstallConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.UninstallConfig(uninstallConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchematicsJob(getSchematicsJobOptions *GetSchematicsJobOptions) - Operation response error`, func() {
		getSchematicsJobPath := "/v1/projects/testString/configs/testString/plan/job"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchematicsJobPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["since"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["pull_request"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSchematicsJob with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetSchematicsJobOptions model
				getSchematicsJobOptionsModel := new(projectsv1.GetSchematicsJobOptions)
				getSchematicsJobOptionsModel.ID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.ConfigName = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
				getSchematicsJobOptionsModel.PullRequest = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetSchematicsJob(getSchematicsJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetSchematicsJob(getSchematicsJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchematicsJob(getSchematicsJobOptions *GetSchematicsJobOptions)`, func() {
		getSchematicsJobPath := "/v1/projects/testString/configs/testString/plan/job"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchematicsJobPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["since"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["pull_request"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID"}`)
				}))
			})
			It(`Invoke GetSchematicsJob successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetSchematicsJobOptions model
				getSchematicsJobOptionsModel := new(projectsv1.GetSchematicsJobOptions)
				getSchematicsJobOptionsModel.ID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.ConfigName = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
				getSchematicsJobOptionsModel.PullRequest = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetSchematicsJobWithContext(ctx, getSchematicsJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetSchematicsJob(getSchematicsJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetSchematicsJobWithContext(ctx, getSchematicsJobOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSchematicsJobPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["since"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["pull_request"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID"}`)
				}))
			})
			It(`Invoke GetSchematicsJob successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetSchematicsJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchematicsJobOptions model
				getSchematicsJobOptionsModel := new(projectsv1.GetSchematicsJobOptions)
				getSchematicsJobOptionsModel.ID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.ConfigName = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
				getSchematicsJobOptionsModel.PullRequest = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetSchematicsJob(getSchematicsJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSchematicsJob with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetSchematicsJobOptions model
				getSchematicsJobOptionsModel := new(projectsv1.GetSchematicsJobOptions)
				getSchematicsJobOptionsModel.ID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.ConfigName = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
				getSchematicsJobOptionsModel.PullRequest = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetSchematicsJob(getSchematicsJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSchematicsJobOptions model with no property values
				getSchematicsJobOptionsModelNew := new(projectsv1.GetSchematicsJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetSchematicsJob(getSchematicsJobOptionsModelNew)
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
			It(`Invoke GetSchematicsJob successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetSchematicsJobOptions model
				getSchematicsJobOptionsModel := new(projectsv1.GetSchematicsJobOptions)
				getSchematicsJobOptionsModel.ID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.ConfigName = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
				getSchematicsJobOptionsModel.PullRequest = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetSchematicsJob(getSchematicsJobOptionsModel)
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
	Describe(`GetCostEstimate(getCostEstimateOptions *GetCostEstimateOptions)`, func() {
		getCostEstimatePath := "/v1/projects/testString/configs/testString/cost_estimate"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCostEstimatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["pull_request"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCostEstimate successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.GetCostEstimate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectsv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigName = core.StringPtr("testString")
				getCostEstimateOptionsModel.PullRequest = core.StringPtr("testString")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetCostEstimate with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectsv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigName = core.StringPtr("testString")
				getCostEstimateOptionsModel.PullRequest = core.StringPtr("testString")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetCostEstimateOptions model with no property values
				getCostEstimateOptionsModelNew := new(projectsv1.GetCostEstimateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.GetCostEstimate(getCostEstimateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectStatus(getProjectStatusOptions *GetProjectStatusOptions) - Operation response error`, func() {
		getProjectStatusPath := "/v1/projects/234234324-3444-4556-224232432/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectStatusPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProjectStatus with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectStatusOptions model
				getProjectStatusOptionsModel := new(projectsv1.GetProjectStatusOptions)
				getProjectStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				getProjectStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetProjectStatus(getProjectStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetProjectStatus(getProjectStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectStatus(getProjectStatusOptions *GetProjectStatusOptions)`, func() {
		getProjectStatusPath := "/v1/projects/234234324-3444-4556-224232432/status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "project_crn": "ProjectCrn", "project_name": "ProjectName", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "git_repo": {"url": "URL", "branch": "Branch", "project_id": "ProjectID"}, "toolchain": {"id": "ID"}, "schematics": {"cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}, "credentials": {"api_key_ref": "ApiKeyRef"}, "configs": [{"name": "Name", "pull_request": "PullRequest", "catalog_id": "CatalogID", "offering_id": "OfferingID", "offering_kind_id": "OfferingKindID", "version_id": "VersionID", "offering_version": "OfferingVersion", "offering_fulfilment_kind": "terraform", "cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}], "dashboard": {"widgets": [{"name": "Name"}]}, "computed_statuses": {"mapKey": "anyValue"}, "active_prs": [{"branch": "Branch", "url": "URL"}], "history": [{"timestamp": "2019-01-01T12:00:00.000Z", "code": "Code", "type": "git_repo"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetProjectStatus successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectStatusOptions model
				getProjectStatusOptionsModel := new(projectsv1.GetProjectStatusOptions)
				getProjectStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				getProjectStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetProjectStatusWithContext(ctx, getProjectStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetProjectStatus(getProjectStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetProjectStatusWithContext(ctx, getProjectStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProjectStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "project_crn": "ProjectCrn", "project_name": "ProjectName", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "git_repo": {"url": "URL", "branch": "Branch", "project_id": "ProjectID"}, "toolchain": {"id": "ID"}, "schematics": {"cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}, "credentials": {"api_key_ref": "ApiKeyRef"}, "configs": [{"name": "Name", "pull_request": "PullRequest", "catalog_id": "CatalogID", "offering_id": "OfferingID", "offering_kind_id": "OfferingKindID", "version_id": "VersionID", "offering_version": "OfferingVersion", "offering_fulfilment_kind": "terraform", "cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}], "dashboard": {"widgets": [{"name": "Name"}]}, "computed_statuses": {"mapKey": "anyValue"}, "active_prs": [{"branch": "Branch", "url": "URL"}], "history": [{"timestamp": "2019-01-01T12:00:00.000Z", "code": "Code", "type": "git_repo"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetProjectStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetProjectStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectStatusOptions model
				getProjectStatusOptionsModel := new(projectsv1.GetProjectStatusOptions)
				getProjectStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				getProjectStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetProjectStatus(getProjectStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProjectStatus with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectStatusOptions model
				getProjectStatusOptionsModel := new(projectsv1.GetProjectStatusOptions)
				getProjectStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				getProjectStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetProjectStatus(getProjectStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectStatusOptions model with no property values
				getProjectStatusOptionsModelNew := new(projectsv1.GetProjectStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetProjectStatus(getProjectStatusOptionsModelNew)
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
			It(`Invoke GetProjectStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectStatusOptions model
				getProjectStatusOptionsModel := new(projectsv1.GetProjectStatusOptions)
				getProjectStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				getProjectStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetProjectStatus(getProjectStatusOptionsModel)
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
	Describe(`UpdateProjectStatus(updateProjectStatusOptions *UpdateProjectStatusOptions) - Operation response error`, func() {
		updateProjectStatusPath := "/v1/projects/234234324-3444-4556-224232432/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectStatusPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProjectStatus with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the History model
				historyModel := new(projectsv1.History)
				historyModel.Timestamp = CreateMockDateTime("2022-05-13T12:25:00Z")
				historyModel.Code = core.StringPtr("ISB005I")
				historyModel.Type = core.StringPtr("schematics")

				// Construct an instance of the ServiceInfoGit model
				serviceInfoGitModel := new(projectsv1.ServiceInfoGit)
				serviceInfoGitModel.URL = core.StringPtr("testString")
				serviceInfoGitModel.Branch = core.StringPtr("testString")
				serviceInfoGitModel.ProjectID = core.StringPtr("testString")

				// Construct an instance of the ServiceInfoToolchain model
				serviceInfoToolchainModel := new(projectsv1.ServiceInfoToolchain)
				serviceInfoToolchainModel.ID = core.StringPtr("testString")

				// Construct an instance of the ServiceInfoSchematics model
				serviceInfoSchematicsModel := new(projectsv1.ServiceInfoSchematics)
				serviceInfoSchematicsModel.CartOrderID = core.StringPtr("us-south.cartOrder.5fbc28c0-87fd-42f5-b2fa-d98d28d6ea96-local.d807263d")
				serviceInfoSchematicsModel.WorkspaceID = core.StringPtr("crn:v1:staging:public:schematics:us-south:a/d44e54841dfe41ee869659f80bddd075:459a39a6-3bc3-4524-aaa8-fb693d96b79e:workspace:us-south.workspace.projects-service.a482a249")
				serviceInfoSchematicsModel.CartItemName = core.StringPtr("name")

				// Construct an instance of the UpdateProjectStatusOptions model
				updateProjectStatusOptionsModel := new(projectsv1.UpdateProjectStatusOptions)
				updateProjectStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				updateProjectStatusOptionsModel.State = core.StringPtr("UPDATING")
				updateProjectStatusOptionsModel.History = []projectsv1.History{*historyModel}
				updateProjectStatusOptionsModel.GitRepo = serviceInfoGitModel
				updateProjectStatusOptionsModel.Toolchain = serviceInfoToolchainModel
				updateProjectStatusOptionsModel.Schematics = serviceInfoSchematicsModel
				updateProjectStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.UpdateProjectStatus(updateProjectStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.UpdateProjectStatus(updateProjectStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProjectStatus(updateProjectStatusOptions *UpdateProjectStatusOptions)`, func() {
		updateProjectStatusPath := "/v1/projects/234234324-3444-4556-224232432/status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectStatusPath))
					Expect(req.Method).To(Equal("PATCH"))

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
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "project_crn": "ProjectCrn", "project_name": "ProjectName", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "git_repo": {"url": "URL", "branch": "Branch", "project_id": "ProjectID"}, "toolchain": {"id": "ID"}, "schematics": {"cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}, "credentials": {"api_key_ref": "ApiKeyRef"}, "configs": [{"name": "Name", "pull_request": "PullRequest", "catalog_id": "CatalogID", "offering_id": "OfferingID", "offering_kind_id": "OfferingKindID", "version_id": "VersionID", "offering_version": "OfferingVersion", "offering_fulfilment_kind": "terraform", "cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}], "dashboard": {"widgets": [{"name": "Name"}]}, "computed_statuses": {"mapKey": "anyValue"}, "active_prs": [{"branch": "Branch", "url": "URL"}], "history": [{"timestamp": "2019-01-01T12:00:00.000Z", "code": "Code", "type": "git_repo"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateProjectStatus successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the History model
				historyModel := new(projectsv1.History)
				historyModel.Timestamp = CreateMockDateTime("2022-05-13T12:25:00Z")
				historyModel.Code = core.StringPtr("ISB005I")
				historyModel.Type = core.StringPtr("schematics")

				// Construct an instance of the ServiceInfoGit model
				serviceInfoGitModel := new(projectsv1.ServiceInfoGit)
				serviceInfoGitModel.URL = core.StringPtr("testString")
				serviceInfoGitModel.Branch = core.StringPtr("testString")
				serviceInfoGitModel.ProjectID = core.StringPtr("testString")

				// Construct an instance of the ServiceInfoToolchain model
				serviceInfoToolchainModel := new(projectsv1.ServiceInfoToolchain)
				serviceInfoToolchainModel.ID = core.StringPtr("testString")

				// Construct an instance of the ServiceInfoSchematics model
				serviceInfoSchematicsModel := new(projectsv1.ServiceInfoSchematics)
				serviceInfoSchematicsModel.CartOrderID = core.StringPtr("us-south.cartOrder.5fbc28c0-87fd-42f5-b2fa-d98d28d6ea96-local.d807263d")
				serviceInfoSchematicsModel.WorkspaceID = core.StringPtr("crn:v1:staging:public:schematics:us-south:a/d44e54841dfe41ee869659f80bddd075:459a39a6-3bc3-4524-aaa8-fb693d96b79e:workspace:us-south.workspace.projects-service.a482a249")
				serviceInfoSchematicsModel.CartItemName = core.StringPtr("name")

				// Construct an instance of the UpdateProjectStatusOptions model
				updateProjectStatusOptionsModel := new(projectsv1.UpdateProjectStatusOptions)
				updateProjectStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				updateProjectStatusOptionsModel.State = core.StringPtr("UPDATING")
				updateProjectStatusOptionsModel.History = []projectsv1.History{*historyModel}
				updateProjectStatusOptionsModel.GitRepo = serviceInfoGitModel
				updateProjectStatusOptionsModel.Toolchain = serviceInfoToolchainModel
				updateProjectStatusOptionsModel.Schematics = serviceInfoSchematicsModel
				updateProjectStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.UpdateProjectStatusWithContext(ctx, updateProjectStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.UpdateProjectStatus(updateProjectStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.UpdateProjectStatusWithContext(ctx, updateProjectStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectStatusPath))
					Expect(req.Method).To(Equal("PATCH"))

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
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "project_crn": "ProjectCrn", "project_name": "ProjectName", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "git_repo": {"url": "URL", "branch": "Branch", "project_id": "ProjectID"}, "toolchain": {"id": "ID"}, "schematics": {"cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}, "credentials": {"api_key_ref": "ApiKeyRef"}, "configs": [{"name": "Name", "pull_request": "PullRequest", "catalog_id": "CatalogID", "offering_id": "OfferingID", "offering_kind_id": "OfferingKindID", "version_id": "VersionID", "offering_version": "OfferingVersion", "offering_fulfilment_kind": "terraform", "cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}], "dashboard": {"widgets": [{"name": "Name"}]}, "computed_statuses": {"mapKey": "anyValue"}, "active_prs": [{"branch": "Branch", "url": "URL"}], "history": [{"timestamp": "2019-01-01T12:00:00.000Z", "code": "Code", "type": "git_repo"}], "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateProjectStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.UpdateProjectStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the History model
				historyModel := new(projectsv1.History)
				historyModel.Timestamp = CreateMockDateTime("2022-05-13T12:25:00Z")
				historyModel.Code = core.StringPtr("ISB005I")
				historyModel.Type = core.StringPtr("schematics")

				// Construct an instance of the ServiceInfoGit model
				serviceInfoGitModel := new(projectsv1.ServiceInfoGit)
				serviceInfoGitModel.URL = core.StringPtr("testString")
				serviceInfoGitModel.Branch = core.StringPtr("testString")
				serviceInfoGitModel.ProjectID = core.StringPtr("testString")

				// Construct an instance of the ServiceInfoToolchain model
				serviceInfoToolchainModel := new(projectsv1.ServiceInfoToolchain)
				serviceInfoToolchainModel.ID = core.StringPtr("testString")

				// Construct an instance of the ServiceInfoSchematics model
				serviceInfoSchematicsModel := new(projectsv1.ServiceInfoSchematics)
				serviceInfoSchematicsModel.CartOrderID = core.StringPtr("us-south.cartOrder.5fbc28c0-87fd-42f5-b2fa-d98d28d6ea96-local.d807263d")
				serviceInfoSchematicsModel.WorkspaceID = core.StringPtr("crn:v1:staging:public:schematics:us-south:a/d44e54841dfe41ee869659f80bddd075:459a39a6-3bc3-4524-aaa8-fb693d96b79e:workspace:us-south.workspace.projects-service.a482a249")
				serviceInfoSchematicsModel.CartItemName = core.StringPtr("name")

				// Construct an instance of the UpdateProjectStatusOptions model
				updateProjectStatusOptionsModel := new(projectsv1.UpdateProjectStatusOptions)
				updateProjectStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				updateProjectStatusOptionsModel.State = core.StringPtr("UPDATING")
				updateProjectStatusOptionsModel.History = []projectsv1.History{*historyModel}
				updateProjectStatusOptionsModel.GitRepo = serviceInfoGitModel
				updateProjectStatusOptionsModel.Toolchain = serviceInfoToolchainModel
				updateProjectStatusOptionsModel.Schematics = serviceInfoSchematicsModel
				updateProjectStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.UpdateProjectStatus(updateProjectStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProjectStatus with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the History model
				historyModel := new(projectsv1.History)
				historyModel.Timestamp = CreateMockDateTime("2022-05-13T12:25:00Z")
				historyModel.Code = core.StringPtr("ISB005I")
				historyModel.Type = core.StringPtr("schematics")

				// Construct an instance of the ServiceInfoGit model
				serviceInfoGitModel := new(projectsv1.ServiceInfoGit)
				serviceInfoGitModel.URL = core.StringPtr("testString")
				serviceInfoGitModel.Branch = core.StringPtr("testString")
				serviceInfoGitModel.ProjectID = core.StringPtr("testString")

				// Construct an instance of the ServiceInfoToolchain model
				serviceInfoToolchainModel := new(projectsv1.ServiceInfoToolchain)
				serviceInfoToolchainModel.ID = core.StringPtr("testString")

				// Construct an instance of the ServiceInfoSchematics model
				serviceInfoSchematicsModel := new(projectsv1.ServiceInfoSchematics)
				serviceInfoSchematicsModel.CartOrderID = core.StringPtr("us-south.cartOrder.5fbc28c0-87fd-42f5-b2fa-d98d28d6ea96-local.d807263d")
				serviceInfoSchematicsModel.WorkspaceID = core.StringPtr("crn:v1:staging:public:schematics:us-south:a/d44e54841dfe41ee869659f80bddd075:459a39a6-3bc3-4524-aaa8-fb693d96b79e:workspace:us-south.workspace.projects-service.a482a249")
				serviceInfoSchematicsModel.CartItemName = core.StringPtr("name")

				// Construct an instance of the UpdateProjectStatusOptions model
				updateProjectStatusOptionsModel := new(projectsv1.UpdateProjectStatusOptions)
				updateProjectStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				updateProjectStatusOptionsModel.State = core.StringPtr("UPDATING")
				updateProjectStatusOptionsModel.History = []projectsv1.History{*historyModel}
				updateProjectStatusOptionsModel.GitRepo = serviceInfoGitModel
				updateProjectStatusOptionsModel.Toolchain = serviceInfoToolchainModel
				updateProjectStatusOptionsModel.Schematics = serviceInfoSchematicsModel
				updateProjectStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.UpdateProjectStatus(updateProjectStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProjectStatusOptions model with no property values
				updateProjectStatusOptionsModelNew := new(projectsv1.UpdateProjectStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.UpdateProjectStatus(updateProjectStatusOptionsModelNew)
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
			It(`Invoke UpdateProjectStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the History model
				historyModel := new(projectsv1.History)
				historyModel.Timestamp = CreateMockDateTime("2022-05-13T12:25:00Z")
				historyModel.Code = core.StringPtr("ISB005I")
				historyModel.Type = core.StringPtr("schematics")

				// Construct an instance of the ServiceInfoGit model
				serviceInfoGitModel := new(projectsv1.ServiceInfoGit)
				serviceInfoGitModel.URL = core.StringPtr("testString")
				serviceInfoGitModel.Branch = core.StringPtr("testString")
				serviceInfoGitModel.ProjectID = core.StringPtr("testString")

				// Construct an instance of the ServiceInfoToolchain model
				serviceInfoToolchainModel := new(projectsv1.ServiceInfoToolchain)
				serviceInfoToolchainModel.ID = core.StringPtr("testString")

				// Construct an instance of the ServiceInfoSchematics model
				serviceInfoSchematicsModel := new(projectsv1.ServiceInfoSchematics)
				serviceInfoSchematicsModel.CartOrderID = core.StringPtr("us-south.cartOrder.5fbc28c0-87fd-42f5-b2fa-d98d28d6ea96-local.d807263d")
				serviceInfoSchematicsModel.WorkspaceID = core.StringPtr("crn:v1:staging:public:schematics:us-south:a/d44e54841dfe41ee869659f80bddd075:459a39a6-3bc3-4524-aaa8-fb693d96b79e:workspace:us-south.workspace.projects-service.a482a249")
				serviceInfoSchematicsModel.CartItemName = core.StringPtr("name")

				// Construct an instance of the UpdateProjectStatusOptions model
				updateProjectStatusOptionsModel := new(projectsv1.UpdateProjectStatusOptions)
				updateProjectStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				updateProjectStatusOptionsModel.State = core.StringPtr("UPDATING")
				updateProjectStatusOptionsModel.History = []projectsv1.History{*historyModel}
				updateProjectStatusOptionsModel.GitRepo = serviceInfoGitModel
				updateProjectStatusOptionsModel.Toolchain = serviceInfoToolchainModel
				updateProjectStatusOptionsModel.Schematics = serviceInfoSchematicsModel
				updateProjectStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.UpdateProjectStatus(updateProjectStatusOptionsModel)
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
	Describe(`GetProjectComputedStatus(getProjectComputedStatusOptions *GetProjectComputedStatusOptions) - Operation response error`, func() {
		getProjectComputedStatusPath := "/v1/projects/234234324-3444-4556-224232432/status/cost"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectComputedStatusPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProjectComputedStatus with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectComputedStatusOptions model
				getProjectComputedStatusOptionsModel := new(projectsv1.GetProjectComputedStatusOptions)
				getProjectComputedStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				getProjectComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				getProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetProjectComputedStatus(getProjectComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetProjectComputedStatus(getProjectComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectComputedStatus(getProjectComputedStatusOptions *GetProjectComputedStatusOptions)`, func() {
		getProjectComputedStatusPath := "/v1/projects/234234324-3444-4556-224232432/status/cost"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectComputedStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "computed_statuses": {"mapKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetProjectComputedStatus successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectComputedStatusOptions model
				getProjectComputedStatusOptionsModel := new(projectsv1.GetProjectComputedStatusOptions)
				getProjectComputedStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				getProjectComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				getProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetProjectComputedStatusWithContext(ctx, getProjectComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetProjectComputedStatus(getProjectComputedStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetProjectComputedStatusWithContext(ctx, getProjectComputedStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProjectComputedStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "computed_statuses": {"mapKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetProjectComputedStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetProjectComputedStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectComputedStatusOptions model
				getProjectComputedStatusOptionsModel := new(projectsv1.GetProjectComputedStatusOptions)
				getProjectComputedStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				getProjectComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				getProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetProjectComputedStatus(getProjectComputedStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProjectComputedStatus with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectComputedStatusOptions model
				getProjectComputedStatusOptionsModel := new(projectsv1.GetProjectComputedStatusOptions)
				getProjectComputedStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				getProjectComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				getProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetProjectComputedStatus(getProjectComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectComputedStatusOptions model with no property values
				getProjectComputedStatusOptionsModelNew := new(projectsv1.GetProjectComputedStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetProjectComputedStatus(getProjectComputedStatusOptionsModelNew)
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
			It(`Invoke GetProjectComputedStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectComputedStatusOptions model
				getProjectComputedStatusOptionsModel := new(projectsv1.GetProjectComputedStatusOptions)
				getProjectComputedStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				getProjectComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				getProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetProjectComputedStatus(getProjectComputedStatusOptionsModel)
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
	Describe(`UpdateProjectComputedStatus(updateProjectComputedStatusOptions *UpdateProjectComputedStatusOptions) - Operation response error`, func() {
		updateProjectComputedStatusPath := "/v1/projects/234234324-3444-4556-224232432/status/cost"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectComputedStatusPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProjectComputedStatus with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UpdateProjectComputedStatusOptions model
				updateProjectComputedStatusOptionsModel := new(projectsv1.UpdateProjectComputedStatusOptions)
				updateProjectComputedStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				updateProjectComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				updateProjectComputedStatusOptionsModel.ComputedStatuses = make(map[string]interface{})
				updateProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.UpdateProjectComputedStatus(updateProjectComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.UpdateProjectComputedStatus(updateProjectComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProjectComputedStatus(updateProjectComputedStatusOptions *UpdateProjectComputedStatusOptions)`, func() {
		updateProjectComputedStatusPath := "/v1/projects/234234324-3444-4556-224232432/status/cost"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectComputedStatusPath))
					Expect(req.Method).To(Equal("PATCH"))

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
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "computed_statuses": {"mapKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateProjectComputedStatus successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateProjectComputedStatusOptions model
				updateProjectComputedStatusOptionsModel := new(projectsv1.UpdateProjectComputedStatusOptions)
				updateProjectComputedStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				updateProjectComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				updateProjectComputedStatusOptionsModel.ComputedStatuses = make(map[string]interface{})
				updateProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.UpdateProjectComputedStatusWithContext(ctx, updateProjectComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.UpdateProjectComputedStatus(updateProjectComputedStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.UpdateProjectComputedStatusWithContext(ctx, updateProjectComputedStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectComputedStatusPath))
					Expect(req.Method).To(Equal("PATCH"))

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
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "computed_statuses": {"mapKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateProjectComputedStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.UpdateProjectComputedStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateProjectComputedStatusOptions model
				updateProjectComputedStatusOptionsModel := new(projectsv1.UpdateProjectComputedStatusOptions)
				updateProjectComputedStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				updateProjectComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				updateProjectComputedStatusOptionsModel.ComputedStatuses = make(map[string]interface{})
				updateProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.UpdateProjectComputedStatus(updateProjectComputedStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProjectComputedStatus with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UpdateProjectComputedStatusOptions model
				updateProjectComputedStatusOptionsModel := new(projectsv1.UpdateProjectComputedStatusOptions)
				updateProjectComputedStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				updateProjectComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				updateProjectComputedStatusOptionsModel.ComputedStatuses = make(map[string]interface{})
				updateProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.UpdateProjectComputedStatus(updateProjectComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProjectComputedStatusOptions model with no property values
				updateProjectComputedStatusOptionsModelNew := new(projectsv1.UpdateProjectComputedStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.UpdateProjectComputedStatus(updateProjectComputedStatusOptionsModelNew)
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
			It(`Invoke UpdateProjectComputedStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UpdateProjectComputedStatusOptions model
				updateProjectComputedStatusOptionsModel := new(projectsv1.UpdateProjectComputedStatusOptions)
				updateProjectComputedStatusOptionsModel.ID = core.StringPtr("234234324-3444-4556-224232432")
				updateProjectComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				updateProjectComputedStatusOptionsModel.ComputedStatuses = make(map[string]interface{})
				updateProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.UpdateProjectComputedStatus(updateProjectComputedStatusOptionsModel)
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
	Describe(`ListProjectConfigStatuses(listProjectConfigStatusesOptions *ListProjectConfigStatusesOptions) - Operation response error`, func() {
		listProjectConfigStatusesPath := "/v1/projects/testString/configs/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectConfigStatusesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProjectConfigStatuses with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ListProjectConfigStatusesOptions model
				listProjectConfigStatusesOptionsModel := new(projectsv1.ListProjectConfigStatusesOptions)
				listProjectConfigStatusesOptionsModel.ID = core.StringPtr("testString")
				listProjectConfigStatusesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.ListProjectConfigStatuses(listProjectConfigStatusesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.ListProjectConfigStatuses(listProjectConfigStatusesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProjectConfigStatuses(listProjectConfigStatusesOptions *ListProjectConfigStatusesOptions)`, func() {
		listProjectConfigStatusesPath := "/v1/projects/testString/configs/status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectConfigStatusesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configs": [{"name": "Name", "state": "State"}]}`)
				}))
			})
			It(`Invoke ListProjectConfigStatuses successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the ListProjectConfigStatusesOptions model
				listProjectConfigStatusesOptionsModel := new(projectsv1.ListProjectConfigStatusesOptions)
				listProjectConfigStatusesOptionsModel.ID = core.StringPtr("testString")
				listProjectConfigStatusesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.ListProjectConfigStatusesWithContext(ctx, listProjectConfigStatusesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.ListProjectConfigStatuses(listProjectConfigStatusesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.ListProjectConfigStatusesWithContext(ctx, listProjectConfigStatusesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProjectConfigStatusesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configs": [{"name": "Name", "state": "State"}]}`)
				}))
			})
			It(`Invoke ListProjectConfigStatuses successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.ListProjectConfigStatuses(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProjectConfigStatusesOptions model
				listProjectConfigStatusesOptionsModel := new(projectsv1.ListProjectConfigStatusesOptions)
				listProjectConfigStatusesOptionsModel.ID = core.StringPtr("testString")
				listProjectConfigStatusesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.ListProjectConfigStatuses(listProjectConfigStatusesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProjectConfigStatuses with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ListProjectConfigStatusesOptions model
				listProjectConfigStatusesOptionsModel := new(projectsv1.ListProjectConfigStatusesOptions)
				listProjectConfigStatusesOptionsModel.ID = core.StringPtr("testString")
				listProjectConfigStatusesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.ListProjectConfigStatuses(listProjectConfigStatusesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListProjectConfigStatusesOptions model with no property values
				listProjectConfigStatusesOptionsModelNew := new(projectsv1.ListProjectConfigStatusesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.ListProjectConfigStatuses(listProjectConfigStatusesOptionsModelNew)
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
			It(`Invoke ListProjectConfigStatuses successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ListProjectConfigStatusesOptions model
				listProjectConfigStatusesOptionsModel := new(projectsv1.ListProjectConfigStatusesOptions)
				listProjectConfigStatusesOptionsModel.ID = core.StringPtr("testString")
				listProjectConfigStatusesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.ListProjectConfigStatuses(listProjectConfigStatusesOptionsModel)
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
	Describe(`GetProjectConfigStatus(getProjectConfigStatusOptions *GetProjectConfigStatusOptions) - Operation response error`, func() {
		getProjectConfigStatusPath := "/v1/projects/b0a2c11d-926c-4653-a15b-ed17d7b34b22/configs/example/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectConfigStatusPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProjectConfigStatus with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectConfigStatusOptions model
				getProjectConfigStatusOptionsModel := new(projectsv1.GetProjectConfigStatusOptions)
				getProjectConfigStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigStatusOptionsModel.ConfigName = core.StringPtr("example")
				getProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetProjectConfigStatus(getProjectConfigStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetProjectConfigStatus(getProjectConfigStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectConfigStatus(getProjectConfigStatusOptions *GetProjectConfigStatusOptions)`, func() {
		getProjectConfigStatusPath := "/v1/projects/b0a2c11d-926c-4653-a15b-ed17d7b34b22/configs/example/status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectConfigStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "state": "State"}`)
				}))
			})
			It(`Invoke GetProjectConfigStatus successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectConfigStatusOptions model
				getProjectConfigStatusOptionsModel := new(projectsv1.GetProjectConfigStatusOptions)
				getProjectConfigStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigStatusOptionsModel.ConfigName = core.StringPtr("example")
				getProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetProjectConfigStatusWithContext(ctx, getProjectConfigStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetProjectConfigStatus(getProjectConfigStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetProjectConfigStatusWithContext(ctx, getProjectConfigStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProjectConfigStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "state": "State"}`)
				}))
			})
			It(`Invoke GetProjectConfigStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetProjectConfigStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectConfigStatusOptions model
				getProjectConfigStatusOptionsModel := new(projectsv1.GetProjectConfigStatusOptions)
				getProjectConfigStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigStatusOptionsModel.ConfigName = core.StringPtr("example")
				getProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetProjectConfigStatus(getProjectConfigStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProjectConfigStatus with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectConfigStatusOptions model
				getProjectConfigStatusOptionsModel := new(projectsv1.GetProjectConfigStatusOptions)
				getProjectConfigStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigStatusOptionsModel.ConfigName = core.StringPtr("example")
				getProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetProjectConfigStatus(getProjectConfigStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectConfigStatusOptions model with no property values
				getProjectConfigStatusOptionsModelNew := new(projectsv1.GetProjectConfigStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetProjectConfigStatus(getProjectConfigStatusOptionsModelNew)
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
			It(`Invoke GetProjectConfigStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectConfigStatusOptions model
				getProjectConfigStatusOptionsModel := new(projectsv1.GetProjectConfigStatusOptions)
				getProjectConfigStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigStatusOptionsModel.ConfigName = core.StringPtr("example")
				getProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetProjectConfigStatus(getProjectConfigStatusOptionsModel)
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
	Describe(`GetProjectConfigComputedStatus(getProjectConfigComputedStatusOptions *GetProjectConfigComputedStatusOptions) - Operation response error`, func() {
		getProjectConfigComputedStatusPath := "/v1/projects/b0a2c11d-926c-4653-a15b-ed17d7b34b22/config_statuses/example/cost"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectConfigComputedStatusPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProjectConfigComputedStatus with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectConfigComputedStatusOptions model
				getProjectConfigComputedStatusOptionsModel := new(projectsv1.GetProjectConfigComputedStatusOptions)
				getProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigComputedStatusOptionsModel.ConfigName = core.StringPtr("example")
				getProjectConfigComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				getProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetProjectConfigComputedStatus(getProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetProjectConfigComputedStatus(getProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectConfigComputedStatus(getProjectConfigComputedStatusOptions *GetProjectConfigComputedStatusOptions)`, func() {
		getProjectConfigComputedStatusPath := "/v1/projects/b0a2c11d-926c-4653-a15b-ed17d7b34b22/config_statuses/example/cost"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectConfigComputedStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"href": "Href", "name": "Name", "computed_statuses": {"mapKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetProjectConfigComputedStatus successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectConfigComputedStatusOptions model
				getProjectConfigComputedStatusOptionsModel := new(projectsv1.GetProjectConfigComputedStatusOptions)
				getProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigComputedStatusOptionsModel.ConfigName = core.StringPtr("example")
				getProjectConfigComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				getProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetProjectConfigComputedStatusWithContext(ctx, getProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetProjectConfigComputedStatus(getProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetProjectConfigComputedStatusWithContext(ctx, getProjectConfigComputedStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProjectConfigComputedStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"href": "Href", "name": "Name", "computed_statuses": {"mapKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetProjectConfigComputedStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetProjectConfigComputedStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectConfigComputedStatusOptions model
				getProjectConfigComputedStatusOptionsModel := new(projectsv1.GetProjectConfigComputedStatusOptions)
				getProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigComputedStatusOptionsModel.ConfigName = core.StringPtr("example")
				getProjectConfigComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				getProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetProjectConfigComputedStatus(getProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProjectConfigComputedStatus with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectConfigComputedStatusOptions model
				getProjectConfigComputedStatusOptionsModel := new(projectsv1.GetProjectConfigComputedStatusOptions)
				getProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigComputedStatusOptionsModel.ConfigName = core.StringPtr("example")
				getProjectConfigComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				getProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetProjectConfigComputedStatus(getProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectConfigComputedStatusOptions model with no property values
				getProjectConfigComputedStatusOptionsModelNew := new(projectsv1.GetProjectConfigComputedStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetProjectConfigComputedStatus(getProjectConfigComputedStatusOptionsModelNew)
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
			It(`Invoke GetProjectConfigComputedStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetProjectConfigComputedStatusOptions model
				getProjectConfigComputedStatusOptionsModel := new(projectsv1.GetProjectConfigComputedStatusOptions)
				getProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigComputedStatusOptionsModel.ConfigName = core.StringPtr("example")
				getProjectConfigComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				getProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetProjectConfigComputedStatus(getProjectConfigComputedStatusOptionsModel)
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
	Describe(`UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptions *UpdateProjectConfigComputedStatusOptions) - Operation response error`, func() {
		updateProjectConfigComputedStatusPath := "/v1/projects/b0a2c11d-926c-4653-a15b-ed17d7b34b22/config_statuses/example/cost"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectConfigComputedStatusPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProjectConfigComputedStatus with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UpdateProjectConfigComputedStatusOptions model
				updateProjectConfigComputedStatusOptionsModel := new(projectsv1.UpdateProjectConfigComputedStatusOptions)
				updateProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigComputedStatusOptionsModel.ConfigName = core.StringPtr("example")
				updateProjectConfigComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				updateProjectConfigComputedStatusOptionsModel.RequestBody = make(map[string]interface{})
				updateProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptions *UpdateProjectConfigComputedStatusOptions)`, func() {
		updateProjectConfigComputedStatusPath := "/v1/projects/b0a2c11d-926c-4653-a15b-ed17d7b34b22/config_statuses/example/cost"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectConfigComputedStatusPath))
					Expect(req.Method).To(Equal("PATCH"))

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
					fmt.Fprintf(res, "%s", `{"href": "Href", "name": "Name", "computed_statuses": {"mapKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateProjectConfigComputedStatus successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateProjectConfigComputedStatusOptions model
				updateProjectConfigComputedStatusOptionsModel := new(projectsv1.UpdateProjectConfigComputedStatusOptions)
				updateProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigComputedStatusOptionsModel.ConfigName = core.StringPtr("example")
				updateProjectConfigComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				updateProjectConfigComputedStatusOptionsModel.RequestBody = make(map[string]interface{})
				updateProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.UpdateProjectConfigComputedStatusWithContext(ctx, updateProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.UpdateProjectConfigComputedStatusWithContext(ctx, updateProjectConfigComputedStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectConfigComputedStatusPath))
					Expect(req.Method).To(Equal("PATCH"))

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
					fmt.Fprintf(res, "%s", `{"href": "Href", "name": "Name", "computed_statuses": {"mapKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateProjectConfigComputedStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.UpdateProjectConfigComputedStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateProjectConfigComputedStatusOptions model
				updateProjectConfigComputedStatusOptionsModel := new(projectsv1.UpdateProjectConfigComputedStatusOptions)
				updateProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigComputedStatusOptionsModel.ConfigName = core.StringPtr("example")
				updateProjectConfigComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				updateProjectConfigComputedStatusOptionsModel.RequestBody = make(map[string]interface{})
				updateProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProjectConfigComputedStatus with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UpdateProjectConfigComputedStatusOptions model
				updateProjectConfigComputedStatusOptionsModel := new(projectsv1.UpdateProjectConfigComputedStatusOptions)
				updateProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigComputedStatusOptionsModel.ConfigName = core.StringPtr("example")
				updateProjectConfigComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				updateProjectConfigComputedStatusOptionsModel.RequestBody = make(map[string]interface{})
				updateProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProjectConfigComputedStatusOptions model with no property values
				updateProjectConfigComputedStatusOptionsModelNew := new(projectsv1.UpdateProjectConfigComputedStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptionsModelNew)
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
			It(`Invoke UpdateProjectConfigComputedStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UpdateProjectConfigComputedStatusOptions model
				updateProjectConfigComputedStatusOptionsModel := new(projectsv1.UpdateProjectConfigComputedStatusOptions)
				updateProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigComputedStatusOptionsModel.ConfigName = core.StringPtr("example")
				updateProjectConfigComputedStatusOptionsModel.ComputedStatus = core.StringPtr("cost")
				updateProjectConfigComputedStatusOptionsModel.RequestBody = make(map[string]interface{})
				updateProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptionsModel)
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
	Describe(`PostNotification(postNotificationOptions *PostNotificationOptions) - Operation response error`, func() {
		postNotificationPath := "/v1/projects/testString/event"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postNotificationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostNotification with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectsv1.NotificationEvent)
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.ActionURL = core.StringPtr("url.for.project.documentation")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the PostNotificationOptions model
				postNotificationOptionsModel := new(projectsv1.PostNotificationOptions)
				postNotificationOptionsModel.ID = core.StringPtr("testString")
				postNotificationOptionsModel.Notifications = []projectsv1.NotificationEvent{*notificationEventModel}
				postNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.PostNotification(postNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.PostNotification(postNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostNotification(postNotificationOptions *PostNotificationOptions)`, func() {
		postNotificationPath := "/v1/projects/testString/event"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postNotificationPath))
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
					fmt.Fprintf(res, "%s", `{"notifications": [{"event": "Event", "target": "Target", "source": "Source", "action_url": "ActionURL", "data": {"anyKey": "anyValue"}, "_id": "ID", "status": "Status", "reasons": [{"anyKey": "anyValue"}]}]}`)
				}))
			})
			It(`Invoke PostNotification successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectsv1.NotificationEvent)
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.ActionURL = core.StringPtr("url.for.project.documentation")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the PostNotificationOptions model
				postNotificationOptionsModel := new(projectsv1.PostNotificationOptions)
				postNotificationOptionsModel.ID = core.StringPtr("testString")
				postNotificationOptionsModel.Notifications = []projectsv1.NotificationEvent{*notificationEventModel}
				postNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.PostNotificationWithContext(ctx, postNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.PostNotification(postNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.PostNotificationWithContext(ctx, postNotificationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postNotificationPath))
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
					fmt.Fprintf(res, "%s", `{"notifications": [{"event": "Event", "target": "Target", "source": "Source", "action_url": "ActionURL", "data": {"anyKey": "anyValue"}, "_id": "ID", "status": "Status", "reasons": [{"anyKey": "anyValue"}]}]}`)
				}))
			})
			It(`Invoke PostNotification successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.PostNotification(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectsv1.NotificationEvent)
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.ActionURL = core.StringPtr("url.for.project.documentation")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the PostNotificationOptions model
				postNotificationOptionsModel := new(projectsv1.PostNotificationOptions)
				postNotificationOptionsModel.ID = core.StringPtr("testString")
				postNotificationOptionsModel.Notifications = []projectsv1.NotificationEvent{*notificationEventModel}
				postNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.PostNotification(postNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostNotification with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectsv1.NotificationEvent)
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.ActionURL = core.StringPtr("url.for.project.documentation")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the PostNotificationOptions model
				postNotificationOptionsModel := new(projectsv1.PostNotificationOptions)
				postNotificationOptionsModel.ID = core.StringPtr("testString")
				postNotificationOptionsModel.Notifications = []projectsv1.NotificationEvent{*notificationEventModel}
				postNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.PostNotification(postNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostNotificationOptions model with no property values
				postNotificationOptionsModelNew := new(projectsv1.PostNotificationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.PostNotification(postNotificationOptionsModelNew)
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
			It(`Invoke PostNotification successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectsv1.NotificationEvent)
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.ActionURL = core.StringPtr("url.for.project.documentation")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the PostNotificationOptions model
				postNotificationOptionsModel := new(projectsv1.PostNotificationOptions)
				postNotificationOptionsModel.ID = core.StringPtr("testString")
				postNotificationOptionsModel.Notifications = []projectsv1.NotificationEvent{*notificationEventModel}
				postNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.PostNotification(postNotificationOptionsModel)
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
	Describe(`GetNotifications(getNotificationsOptions *GetNotificationsOptions) - Operation response error`, func() {
		getNotificationsPath := "/v1/projects/testString/event"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getNotificationsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetNotifications with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetNotificationsOptions model
				getNotificationsOptionsModel := new(projectsv1.GetNotificationsOptions)
				getNotificationsOptionsModel.ID = core.StringPtr("testString")
				getNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetNotifications(getNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetNotifications(getNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetNotifications(getNotificationsOptions *GetNotificationsOptions)`, func() {
		getNotificationsPath := "/v1/projects/testString/event"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getNotificationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"notifications": [{"event": "Event", "target": "Target", "source": "Source", "action_url": "ActionURL", "data": {"anyKey": "anyValue"}, "_id": "ID"}]}`)
				}))
			})
			It(`Invoke GetNotifications successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetNotificationsOptions model
				getNotificationsOptionsModel := new(projectsv1.GetNotificationsOptions)
				getNotificationsOptionsModel.ID = core.StringPtr("testString")
				getNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetNotificationsWithContext(ctx, getNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetNotifications(getNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetNotificationsWithContext(ctx, getNotificationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getNotificationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"notifications": [{"event": "Event", "target": "Target", "source": "Source", "action_url": "ActionURL", "data": {"anyKey": "anyValue"}, "_id": "ID"}]}`)
				}))
			})
			It(`Invoke GetNotifications successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetNotificationsOptions model
				getNotificationsOptionsModel := new(projectsv1.GetNotificationsOptions)
				getNotificationsOptionsModel.ID = core.StringPtr("testString")
				getNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetNotifications(getNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetNotifications with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetNotificationsOptions model
				getNotificationsOptionsModel := new(projectsv1.GetNotificationsOptions)
				getNotificationsOptionsModel.ID = core.StringPtr("testString")
				getNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetNotifications(getNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetNotificationsOptions model with no property values
				getNotificationsOptionsModelNew := new(projectsv1.GetNotificationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetNotifications(getNotificationsOptionsModelNew)
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
			It(`Invoke GetNotifications successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetNotificationsOptions model
				getNotificationsOptionsModel := new(projectsv1.GetNotificationsOptions)
				getNotificationsOptionsModel.ID = core.StringPtr("testString")
				getNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetNotifications(getNotificationsOptionsModel)
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
	Describe(`DeleteNotification(deleteNotificationOptions *DeleteNotificationOptions)`, func() {
		deleteNotificationPath := "/v1/projects/testString/event"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteNotificationPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["notification_id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteNotification successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.DeleteNotification(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteNotificationOptions model
				deleteNotificationOptionsModel := new(projectsv1.DeleteNotificationOptions)
				deleteNotificationOptionsModel.ID = core.StringPtr("testString")
				deleteNotificationOptionsModel.NotificationID = core.StringPtr("testString")
				deleteNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.DeleteNotification(deleteNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteNotification with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeleteNotificationOptions model
				deleteNotificationOptionsModel := new(projectsv1.DeleteNotificationOptions)
				deleteNotificationOptionsModel.ID = core.StringPtr("testString")
				deleteNotificationOptionsModel.NotificationID = core.StringPtr("testString")
				deleteNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.DeleteNotification(deleteNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteNotificationOptions model with no property values
				deleteNotificationOptionsModelNew := new(projectsv1.DeleteNotificationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.DeleteNotification(deleteNotificationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReceivePulsarCatalogEvents(receivePulsarCatalogEventsOptions *ReceivePulsarCatalogEventsOptions)`, func() {
		receivePulsarCatalogEventsPath := "/v1/pulsar/catalog_events"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(receivePulsarCatalogEventsPath))
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

					res.WriteHeader(202)
				}))
			})
			It(`Invoke ReceivePulsarCatalogEvents successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.ReceivePulsarCatalogEvents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PulsarEventItem model
				pulsarEventItemModel := new(projectsv1.PulsarEventItem)
				pulsarEventItemModel.EventType = core.StringPtr("testString")
				pulsarEventItemModel.Timestamp = core.StringPtr("testString")
				pulsarEventItemModel.Publisher = core.StringPtr("testString")
				pulsarEventItemModel.AccountID = core.StringPtr("testString")
				pulsarEventItemModel.Version = core.StringPtr("testString")
				pulsarEventItemModel.EventProperties = map[string]interface{}{"anyKey": "anyValue"}
				pulsarEventItemModel.EventID = core.StringPtr("testString")
				pulsarEventItemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ReceivePulsarCatalogEventsOptions model
				receivePulsarCatalogEventsOptionsModel := new(projectsv1.ReceivePulsarCatalogEventsOptions)
				receivePulsarCatalogEventsOptionsModel.PulsarEventItem = []projectsv1.PulsarEventItem{*pulsarEventItemModel}
				receivePulsarCatalogEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.ReceivePulsarCatalogEvents(receivePulsarCatalogEventsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ReceivePulsarCatalogEvents with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the PulsarEventItem model
				pulsarEventItemModel := new(projectsv1.PulsarEventItem)
				pulsarEventItemModel.EventType = core.StringPtr("testString")
				pulsarEventItemModel.Timestamp = core.StringPtr("testString")
				pulsarEventItemModel.Publisher = core.StringPtr("testString")
				pulsarEventItemModel.AccountID = core.StringPtr("testString")
				pulsarEventItemModel.Version = core.StringPtr("testString")
				pulsarEventItemModel.EventProperties = map[string]interface{}{"anyKey": "anyValue"}
				pulsarEventItemModel.EventID = core.StringPtr("testString")
				pulsarEventItemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ReceivePulsarCatalogEventsOptions model
				receivePulsarCatalogEventsOptionsModel := new(projectsv1.ReceivePulsarCatalogEventsOptions)
				receivePulsarCatalogEventsOptionsModel.PulsarEventItem = []projectsv1.PulsarEventItem{*pulsarEventItemModel}
				receivePulsarCatalogEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.ReceivePulsarCatalogEvents(receivePulsarCatalogEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ReceivePulsarCatalogEventsOptions model with no property values
				receivePulsarCatalogEventsOptionsModelNew := new(projectsv1.ReceivePulsarCatalogEventsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.ReceivePulsarCatalogEvents(receivePulsarCatalogEventsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReceiveGitlabEvents(receiveGitlabEventsOptions *ReceiveGitlabEventsOptions)`, func() {
		receiveGitlabEventsPath := "/v1/gitlab_webhook/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(receiveGitlabEventsPath))
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

					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReceiveGitlabEvents successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.ReceiveGitlabEvents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GitLabEvent model
				gitLabEventModel := new(projectsv1.GitLabEvent)
				gitLabEventModel.ObjectKind = core.StringPtr("push")
				gitLabEventModel.SetProperty("event_name", core.StringPtr("push"))
				gitLabEventModel.SetProperty("before", core.StringPtr("62928f2f51fcea33aefad3145c6ae81434a4cacf"))
				gitLabEventModel.SetProperty("after", core.StringPtr("1dfe1e32d92c72a73f930a9babc36de0d6019fce"))

				// Construct an instance of the ReceiveGitlabEventsOptions model
				receiveGitlabEventsOptionsModel := new(projectsv1.ReceiveGitlabEventsOptions)
				receiveGitlabEventsOptionsModel.ID = core.StringPtr("testString")
				receiveGitlabEventsOptionsModel.GitLabEvent = gitLabEventModel
				receiveGitlabEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.ReceiveGitlabEvents(receiveGitlabEventsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ReceiveGitlabEvents with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GitLabEvent model
				gitLabEventModel := new(projectsv1.GitLabEvent)
				gitLabEventModel.ObjectKind = core.StringPtr("push")
				gitLabEventModel.SetProperty("event_name", core.StringPtr("push"))
				gitLabEventModel.SetProperty("before", core.StringPtr("62928f2f51fcea33aefad3145c6ae81434a4cacf"))
				gitLabEventModel.SetProperty("after", core.StringPtr("1dfe1e32d92c72a73f930a9babc36de0d6019fce"))

				// Construct an instance of the ReceiveGitlabEventsOptions model
				receiveGitlabEventsOptionsModel := new(projectsv1.ReceiveGitlabEventsOptions)
				receiveGitlabEventsOptionsModel.ID = core.StringPtr("testString")
				receiveGitlabEventsOptionsModel.GitLabEvent = gitLabEventModel
				receiveGitlabEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.ReceiveGitlabEvents(receiveGitlabEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ReceiveGitlabEventsOptions model with no property values
				receiveGitlabEventsOptionsModelNew := new(projectsv1.ReceiveGitlabEventsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.ReceiveGitlabEvents(receiveGitlabEventsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHealth(getHealthOptions *GetHealthOptions) - Operation response error`, func() {
		getHealthPath := "/v1/health"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHealthPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for info query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHealth with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(projectsv1.GetHealthOptions)
				getHealthOptionsModel.Info = core.BoolPtr(false)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHealth(getHealthOptions *GetHealthOptions)`, func() {
		getHealthPath := "/v1/health"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHealthPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for info query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "version": "Version", "dependencies": {"mapKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetHealth successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(projectsv1.GetHealthOptions)
				getHealthOptionsModel.Info = core.BoolPtr(false)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetHealthWithContext(ctx, getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetHealthWithContext(ctx, getHealthOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getHealthPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for info query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "version": "Version", "dependencies": {"mapKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetHealth successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetHealth(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(projectsv1.GetHealthOptions)
				getHealthOptionsModel.Info = core.BoolPtr(false)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetHealth with error: Operation request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(projectsv1.GetHealthOptions)
				getHealthOptionsModel.Info = core.BoolPtr(false)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke GetHealth successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(projectsv1.GetHealthOptions)
				getHealthOptionsModel.Info = core.BoolPtr(false)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetHealth(getHealthOptionsModel)
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
	Describe(`Notify(notifyOptions *NotifyOptions)`, func() {
		notifyPath := "/v1/notify"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(notifyPath))
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

					res.WriteHeader(202)
				}))
			})
			It(`Invoke Notify successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.Notify(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the NotifyOptions model
				notifyOptionsModel := new(projectsv1.NotifyOptions)
				notifyOptionsModel.ID = core.StringPtr("bccbb195-fff4-4d7e-9078-61b06adc02ab")
				notifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.Notify(notifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke Notify with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the NotifyOptions model
				notifyOptionsModel := new(projectsv1.NotifyOptions)
				notifyOptionsModel.ID = core.StringPtr("bccbb195-fff4-4d7e-9078-61b06adc02ab")
				notifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.Notify(notifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the NotifyOptions model with no property values
				notifyOptionsModelNew := new(projectsv1.NotifyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.Notify(notifyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RegisterPullRequest(registerPullRequestOptions *RegisterPullRequestOptions)`, func() {
		registerPullRequestPath := "/v1/projects/testString/pullrequest"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(registerPullRequestPath))
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

					res.WriteHeader(200)
				}))
			})
			It(`Invoke RegisterPullRequest successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.RegisterPullRequest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RegisterPullRequestOptions model
				registerPullRequestOptionsModel := new(projectsv1.RegisterPullRequestOptions)
				registerPullRequestOptionsModel.ID = core.StringPtr("testString")
				registerPullRequestOptionsModel.Branch = core.StringPtr("test")
				registerPullRequestOptionsModel.URL = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				registerPullRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.RegisterPullRequest(registerPullRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RegisterPullRequest with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the RegisterPullRequestOptions model
				registerPullRequestOptionsModel := new(projectsv1.RegisterPullRequestOptions)
				registerPullRequestOptionsModel.ID = core.StringPtr("testString")
				registerPullRequestOptionsModel.Branch = core.StringPtr("test")
				registerPullRequestOptionsModel.URL = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				registerPullRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.RegisterPullRequest(registerPullRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RegisterPullRequestOptions model with no property values
				registerPullRequestOptionsModelNew := new(projectsv1.RegisterPullRequestOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.RegisterPullRequest(registerPullRequestOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeregisterPullRequest(deregisterPullRequestOptions *DeregisterPullRequestOptions)`, func() {
		deregisterPullRequestPath := "/v1/projects/testString/pullrequest"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deregisterPullRequestPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["url"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeregisterPullRequest successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.DeregisterPullRequest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeregisterPullRequestOptions model
				deregisterPullRequestOptionsModel := new(projectsv1.DeregisterPullRequestOptions)
				deregisterPullRequestOptionsModel.ID = core.StringPtr("testString")
				deregisterPullRequestOptionsModel.URL = core.StringPtr("testString")
				deregisterPullRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.DeregisterPullRequest(deregisterPullRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeregisterPullRequest with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeregisterPullRequestOptions model
				deregisterPullRequestOptionsModel := new(projectsv1.DeregisterPullRequestOptions)
				deregisterPullRequestOptionsModel.ID = core.StringPtr("testString")
				deregisterPullRequestOptionsModel.URL = core.StringPtr("testString")
				deregisterPullRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.DeregisterPullRequest(deregisterPullRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeregisterPullRequestOptions model with no property values
				deregisterPullRequestOptionsModelNew := new(projectsv1.DeregisterPullRequestOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.DeregisterPullRequest(deregisterPullRequestOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePullRequestConfigs(updatePullRequestConfigsOptions *UpdatePullRequestConfigsOptions)`, func() {
		updatePullRequestConfigsPath := "/v1/projects/testString/pullrequest/configs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePullRequestConfigsPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					Expect(req.URL.Query()["state"]).To(Equal([]string{"merge"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdatePullRequestConfigs successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.UpdatePullRequestConfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the UpdatePullRequestConfigsOptions model
				updatePullRequestConfigsOptionsModel := new(projectsv1.UpdatePullRequestConfigsOptions)
				updatePullRequestConfigsOptionsModel.ID = core.StringPtr("testString")
				updatePullRequestConfigsOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				updatePullRequestConfigsOptionsModel.ProjectDefinition = projectInputModel
				updatePullRequestConfigsOptionsModel.State = core.StringPtr("merge")
				updatePullRequestConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.UpdatePullRequestConfigs(updatePullRequestConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdatePullRequestConfigs with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the UpdatePullRequestConfigsOptions model
				updatePullRequestConfigsOptionsModel := new(projectsv1.UpdatePullRequestConfigsOptions)
				updatePullRequestConfigsOptionsModel.ID = core.StringPtr("testString")
				updatePullRequestConfigsOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				updatePullRequestConfigsOptionsModel.ProjectDefinition = projectInputModel
				updatePullRequestConfigsOptionsModel.State = core.StringPtr("merge")
				updatePullRequestConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.UpdatePullRequestConfigs(updatePullRequestConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdatePullRequestConfigsOptions model with no property values
				updatePullRequestConfigsOptionsModelNew := new(projectsv1.UpdatePullRequestConfigsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.UpdatePullRequestConfigs(updatePullRequestConfigsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PlanPullRequestConfigs(planPullRequestConfigsOptions *PlanPullRequestConfigsOptions) - Operation response error`, func() {
		planPullRequestConfigsPath := "/v1/projects/testString/pullrequest/configs/plan"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(planPullRequestConfigsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PlanPullRequestConfigs with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the PlanPullRequestConfigsOptions model
				planPullRequestConfigsOptionsModel := new(projectsv1.PlanPullRequestConfigsOptions)
				planPullRequestConfigsOptionsModel.ID = core.StringPtr("testString")
				planPullRequestConfigsOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				planPullRequestConfigsOptionsModel.ProjectDefinition = projectInputModel
				planPullRequestConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.PlanPullRequestConfigs(planPullRequestConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.PlanPullRequestConfigs(planPullRequestConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PlanPullRequestConfigs(planPullRequestConfigsOptions *PlanPullRequestConfigsOptions)`, func() {
		planPullRequestConfigsPath := "/v1/projects/testString/pullrequest/configs/plan"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(planPullRequestConfigsPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"configs": [{"name": "Name", "job": "Job", "workspace": "Workspace", "cart_order": "CartOrder", "catalog_error": "CatalogError", "catalog_status_code": 17, "schematics_error": "SchematicsError", "schematics_status_code": 20, "schematics_submitted_at": 21}]}`)
				}))
			})
			It(`Invoke PlanPullRequestConfigs successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the PlanPullRequestConfigsOptions model
				planPullRequestConfigsOptionsModel := new(projectsv1.PlanPullRequestConfigsOptions)
				planPullRequestConfigsOptionsModel.ID = core.StringPtr("testString")
				planPullRequestConfigsOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				planPullRequestConfigsOptionsModel.ProjectDefinition = projectInputModel
				planPullRequestConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.PlanPullRequestConfigsWithContext(ctx, planPullRequestConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.PlanPullRequestConfigs(planPullRequestConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.PlanPullRequestConfigsWithContext(ctx, planPullRequestConfigsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(planPullRequestConfigsPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"configs": [{"name": "Name", "job": "Job", "workspace": "Workspace", "cart_order": "CartOrder", "catalog_error": "CatalogError", "catalog_status_code": 17, "schematics_error": "SchematicsError", "schematics_status_code": 20, "schematics_submitted_at": 21}]}`)
				}))
			})
			It(`Invoke PlanPullRequestConfigs successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.PlanPullRequestConfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the PlanPullRequestConfigsOptions model
				planPullRequestConfigsOptionsModel := new(projectsv1.PlanPullRequestConfigsOptions)
				planPullRequestConfigsOptionsModel.ID = core.StringPtr("testString")
				planPullRequestConfigsOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				planPullRequestConfigsOptionsModel.ProjectDefinition = projectInputModel
				planPullRequestConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.PlanPullRequestConfigs(planPullRequestConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PlanPullRequestConfigs with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the PlanPullRequestConfigsOptions model
				planPullRequestConfigsOptionsModel := new(projectsv1.PlanPullRequestConfigsOptions)
				planPullRequestConfigsOptionsModel.ID = core.StringPtr("testString")
				planPullRequestConfigsOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				planPullRequestConfigsOptionsModel.ProjectDefinition = projectInputModel
				planPullRequestConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.PlanPullRequestConfigs(planPullRequestConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PlanPullRequestConfigsOptions model with no property values
				planPullRequestConfigsOptionsModelNew := new(projectsv1.PlanPullRequestConfigsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.PlanPullRequestConfigs(planPullRequestConfigsOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke PlanPullRequestConfigs successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}

				// Construct an instance of the PlanPullRequestConfigsOptions model
				planPullRequestConfigsOptionsModel := new(projectsv1.PlanPullRequestConfigsOptions)
				planPullRequestConfigsOptionsModel.ID = core.StringPtr("testString")
				planPullRequestConfigsOptionsModel.PullRequest = core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				planPullRequestConfigsOptionsModel.ProjectDefinition = projectInputModel
				planPullRequestConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.PlanPullRequestConfigs(planPullRequestConfigsOptionsModel)
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
	Describe(`ProvisionServiceInstance(provisionServiceInstanceOptions *ProvisionServiceInstanceOptions) - Operation response error`, func() {
		provisionServiceInstancePath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(provisionServiceInstancePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					// TODO: Add check for accepts_incomplete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ProvisionServiceInstance with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ProvisionServiceInstanceOptions model
				provisionServiceInstanceOptionsModel := new(projectsv1.ProvisionServiceInstanceOptions)
				provisionServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				provisionServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				provisionServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				provisionServiceInstanceOptionsModel.Context = []string{"testString"}
				provisionServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				provisionServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				provisionServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				provisionServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				provisionServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				provisionServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.ProvisionServiceInstance(provisionServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.ProvisionServiceInstance(provisionServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ProvisionServiceInstance(provisionServiceInstanceOptions *ProvisionServiceInstanceOptions)`, func() {
		provisionServiceInstancePath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(provisionServiceInstancePath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					// TODO: Add check for accepts_incomplete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dashboard_url": "DashboardURL", "operation": "Operation"}`)
				}))
			})
			It(`Invoke ProvisionServiceInstance successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the ProvisionServiceInstanceOptions model
				provisionServiceInstanceOptionsModel := new(projectsv1.ProvisionServiceInstanceOptions)
				provisionServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				provisionServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				provisionServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				provisionServiceInstanceOptionsModel.Context = []string{"testString"}
				provisionServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				provisionServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				provisionServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				provisionServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				provisionServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				provisionServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.ProvisionServiceInstanceWithContext(ctx, provisionServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.ProvisionServiceInstance(provisionServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.ProvisionServiceInstanceWithContext(ctx, provisionServiceInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(provisionServiceInstancePath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					// TODO: Add check for accepts_incomplete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dashboard_url": "DashboardURL", "operation": "Operation"}`)
				}))
			})
			It(`Invoke ProvisionServiceInstance successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.ProvisionServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProvisionServiceInstanceOptions model
				provisionServiceInstanceOptionsModel := new(projectsv1.ProvisionServiceInstanceOptions)
				provisionServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				provisionServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				provisionServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				provisionServiceInstanceOptionsModel.Context = []string{"testString"}
				provisionServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				provisionServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				provisionServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				provisionServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				provisionServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				provisionServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.ProvisionServiceInstance(provisionServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ProvisionServiceInstance with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ProvisionServiceInstanceOptions model
				provisionServiceInstanceOptionsModel := new(projectsv1.ProvisionServiceInstanceOptions)
				provisionServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				provisionServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				provisionServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				provisionServiceInstanceOptionsModel.Context = []string{"testString"}
				provisionServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				provisionServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				provisionServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				provisionServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				provisionServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				provisionServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.ProvisionServiceInstance(provisionServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ProvisionServiceInstanceOptions model with no property values
				provisionServiceInstanceOptionsModelNew := new(projectsv1.ProvisionServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.ProvisionServiceInstance(provisionServiceInstanceOptionsModelNew)
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
			It(`Invoke ProvisionServiceInstance successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ProvisionServiceInstanceOptions model
				provisionServiceInstanceOptionsModel := new(projectsv1.ProvisionServiceInstanceOptions)
				provisionServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				provisionServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				provisionServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				provisionServiceInstanceOptionsModel.Context = []string{"testString"}
				provisionServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				provisionServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				provisionServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				provisionServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				provisionServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				provisionServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.ProvisionServiceInstance(provisionServiceInstanceOptionsModel)
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
	Describe(`DeprovisionServiceInstance(deprovisionServiceInstanceOptions *DeprovisionServiceInstanceOptions)`, func() {
		deprovisionServiceInstancePath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deprovisionServiceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"e1031479-4b41-4159-b7cf-f7792c616fdc"}))
					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
					// TODO: Add check for accepts_incomplete query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeprovisionServiceInstance successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.DeprovisionServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeprovisionServiceInstanceOptions model
				deprovisionServiceInstanceOptionsModel := new(projectsv1.DeprovisionServiceInstanceOptions)
				deprovisionServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deprovisionServiceInstanceOptionsModel.PlanID = core.StringPtr("e1031479-4b41-4159-b7cf-f7792c616fdc")
				deprovisionServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deprovisionServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				deprovisionServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deprovisionServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deprovisionServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.DeprovisionServiceInstance(deprovisionServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeprovisionServiceInstance with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeprovisionServiceInstanceOptions model
				deprovisionServiceInstanceOptionsModel := new(projectsv1.DeprovisionServiceInstanceOptions)
				deprovisionServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deprovisionServiceInstanceOptionsModel.PlanID = core.StringPtr("e1031479-4b41-4159-b7cf-f7792c616fdc")
				deprovisionServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deprovisionServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				deprovisionServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deprovisionServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deprovisionServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.DeprovisionServiceInstance(deprovisionServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeprovisionServiceInstanceOptions model with no property values
				deprovisionServiceInstanceOptionsModelNew := new(projectsv1.DeprovisionServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.DeprovisionServiceInstance(deprovisionServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions)`, func() {
		updateServiceInstancePath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateServiceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))

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

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					// TODO: Add check for accepts_incomplete query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateServiceInstance successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.UpdateServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(projectsv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.ServiceID = []string{"testString"}
				updateServiceInstanceOptionsModel.Context = []string{"testString"}
				updateServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateServiceInstance with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(projectsv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.ServiceID = []string{"testString"}
				updateServiceInstanceOptionsModel.Context = []string{"testString"}
				updateServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateServiceInstanceOptions model with no property values
				updateServiceInstanceOptionsModelNew := new(projectsv1.UpdateServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.UpdateServiceInstance(updateServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLastOperation(getLastOperationOptions *GetLastOperationOptions) - Operation response error`, func() {
		getLastOperationPath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::/last_operation"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLastOperationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"ABCD"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"e1031479-4b41-4159-b7cf-f7792c616fdc"}))
					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLastOperation with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(projectsv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("e1031479-4b41-4159-b7cf-f7792c616fdc")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLastOperation(getLastOperationOptions *GetLastOperationOptions)`, func() {
		getLastOperationPath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::/last_operation"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLastOperationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"ABCD"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"e1031479-4b41-4159-b7cf-f7792c616fdc"}))
					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"state": "State", "description": "Description"}`)
				}))
			})
			It(`Invoke GetLastOperation successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(projectsv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("e1031479-4b41-4159-b7cf-f7792c616fdc")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetLastOperationWithContext(ctx, getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetLastOperationWithContext(ctx, getLastOperationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLastOperationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"ABCD"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"e1031479-4b41-4159-b7cf-f7792c616fdc"}))
					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"state": "State", "description": "Description"}`)
				}))
			})
			It(`Invoke GetLastOperation successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetLastOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(projectsv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("e1031479-4b41-4159-b7cf-f7792c616fdc")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLastOperation with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(projectsv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("e1031479-4b41-4159-b7cf-f7792c616fdc")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLastOperationOptions model with no property values
				getLastOperationOptionsModelNew := new(projectsv1.GetLastOperationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetLastOperation(getLastOperationOptionsModelNew)
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
			It(`Invoke GetLastOperation successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(projectsv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("e1031479-4b41-4159-b7cf-f7792c616fdc")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetLastOperation(getLastOperationOptionsModel)
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
	Describe(`UpdateServiceInstanceState(updateServiceInstanceStateOptions *UpdateServiceInstanceStateOptions) - Operation response error`, func() {
		updateServiceInstanceStatePath := "/bluemix_v1/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateServiceInstanceStatePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateServiceInstanceState with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UpdateServiceInstanceStateOptions model
				updateServiceInstanceStateOptionsModel := new(projectsv1.UpdateServiceInstanceStateOptions)
				updateServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				updateServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				updateServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				updateServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				updateServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				updateServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.UpdateServiceInstanceState(updateServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.UpdateServiceInstanceState(updateServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateServiceInstanceState(updateServiceInstanceStateOptions *UpdateServiceInstanceStateOptions)`, func() {
		updateServiceInstanceStatePath := "/bluemix_v1/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateServiceInstanceStatePath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": "Active", "enabled": "Enabled", "last_active": "LastActive"}`)
				}))
			})
			It(`Invoke UpdateServiceInstanceState successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateServiceInstanceStateOptions model
				updateServiceInstanceStateOptionsModel := new(projectsv1.UpdateServiceInstanceStateOptions)
				updateServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				updateServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				updateServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				updateServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				updateServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				updateServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.UpdateServiceInstanceStateWithContext(ctx, updateServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.UpdateServiceInstanceState(updateServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.UpdateServiceInstanceStateWithContext(ctx, updateServiceInstanceStateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateServiceInstanceStatePath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": "Active", "enabled": "Enabled", "last_active": "LastActive"}`)
				}))
			})
			It(`Invoke UpdateServiceInstanceState successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.UpdateServiceInstanceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateServiceInstanceStateOptions model
				updateServiceInstanceStateOptionsModel := new(projectsv1.UpdateServiceInstanceStateOptions)
				updateServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				updateServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				updateServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				updateServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				updateServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				updateServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.UpdateServiceInstanceState(updateServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateServiceInstanceState with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UpdateServiceInstanceStateOptions model
				updateServiceInstanceStateOptionsModel := new(projectsv1.UpdateServiceInstanceStateOptions)
				updateServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				updateServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				updateServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				updateServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				updateServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				updateServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.UpdateServiceInstanceState(updateServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateServiceInstanceStateOptions model with no property values
				updateServiceInstanceStateOptionsModelNew := new(projectsv1.UpdateServiceInstanceStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.UpdateServiceInstanceState(updateServiceInstanceStateOptionsModelNew)
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
			It(`Invoke UpdateServiceInstanceState successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UpdateServiceInstanceStateOptions model
				updateServiceInstanceStateOptionsModel := new(projectsv1.UpdateServiceInstanceStateOptions)
				updateServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				updateServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				updateServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				updateServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				updateServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				updateServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.UpdateServiceInstanceState(updateServiceInstanceStateOptionsModel)
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
	Describe(`GetServiceInstance(getServiceInstanceOptions *GetServiceInstanceOptions) - Operation response error`, func() {
		getServiceInstancePath := "/bluemix_v1/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceInstancePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetServiceInstance with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(projectsv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetServiceInstance(getServiceInstanceOptions *GetServiceInstanceOptions)`, func() {
		getServiceInstancePath := "/bluemix_v1/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": "Active", "enabled": "Enabled", "last_active": "LastActive"}`)
				}))
			})
			It(`Invoke GetServiceInstance successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(projectsv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetServiceInstanceWithContext(ctx, getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetServiceInstanceWithContext(ctx, getServiceInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getServiceInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": "Active", "enabled": "Enabled", "last_active": "LastActive"}`)
				}))
			})
			It(`Invoke GetServiceInstance successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(projectsv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetServiceInstance with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(projectsv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetServiceInstanceOptions model with no property values
				getServiceInstanceOptionsModelNew := new(projectsv1.GetServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetServiceInstance(getServiceInstanceOptionsModelNew)
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
			It(`Invoke GetServiceInstance successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(projectsv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetServiceInstance(getServiceInstanceOptionsModel)
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
	Describe(`GetCatalog(getCatalogOptions *GetCatalogOptions) - Operation response error`, func() {
		getCatalogPath := "/v2/catalog"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalog with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(projectsv1.GetCatalogOptions)
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {
		getCatalogPath := "/v2/catalog"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"services": [{"bindable": true, "description": "Description", "id": "ID", "metadata": {"display_name": "DisplayName", "documentation_url": "DocumentationURL", "image_url": "ImageURL", "instructions_url": "InstructionsURL", "long_description": "LongDescription", "provider_display_name": "ProviderDisplayName", "support_url": "SupportURL", "terms_url": "TermsURL"}, "name": "Name", "plan_updateable": true, "tags": ["Tags"], "plans": [{"description": "Description", "free": true, "id": "ID", "metadata": {"bullets": ["Bullets"], "display_name": "DisplayName"}, "name": "Name"}]}]}`)
				}))
			})
			It(`Invoke GetCatalog successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(projectsv1.GetCatalogOptions)
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetCatalogWithContext(ctx, getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetCatalogWithContext(ctx, getCatalogOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"services": [{"bindable": true, "description": "Description", "id": "ID", "metadata": {"display_name": "DisplayName", "documentation_url": "DocumentationURL", "image_url": "ImageURL", "instructions_url": "InstructionsURL", "long_description": "LongDescription", "provider_display_name": "ProviderDisplayName", "support_url": "SupportURL", "terms_url": "TermsURL"}, "name": "Name", "plan_updateable": true, "tags": ["Tags"], "plans": [{"description": "Description", "free": true, "id": "ID", "metadata": {"bullets": ["Bullets"], "display_name": "DisplayName"}, "name": "Name"}]}]}`)
				}))
			})
			It(`Invoke GetCatalog successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(projectsv1.GetCatalogOptions)
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCatalog with error: Operation request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(projectsv1.GetCatalogOptions)
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke GetCatalog successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(projectsv1.GetCatalogOptions)
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetCatalog(getCatalogOptionsModel)
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
			projectsService, _ := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
				URL:           "http://projectsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewConfigChangesOptions successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")
				Expect(inputVariableModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Type).To(Equal(core.StringPtr("array")))
				Expect(inputVariableModel.Required).To(Equal(core.BoolPtr(true)))
				Expect(inputVariableModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Default).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				Expect(configSettingItemModel).ToNot(BeNil())
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")
				Expect(configSettingItemModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(configSettingItemModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				Expect(projectConfigInputModel).ToNot(BeNil())
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}
				Expect(projectConfigInputModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigInputModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))
				Expect(projectConfigInputModel.LocatorID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputModel.Input).To(Equal([]projectsv1.InputVariable{*inputVariableModel}))
				Expect(projectConfigInputModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigInputModel.Setting).To(Equal([]projectsv1.ConfigSettingItem{*configSettingItemModel}))

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				Expect(projectInputModel).ToNot(BeNil())
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				Expect(projectInputModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(projectInputModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(projectInputModel.Configs).To(Equal([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel}))

				// Construct an instance of the ConfigChangesOptions model
				id := "234234324-3444-4556-224232432"
				var configChangesOptionsSource *projectsv1.ProjectInput = nil
				configChangesOptionsModel := projectsService.NewConfigChangesOptions(id, configChangesOptionsSource)
				configChangesOptionsModel.SetID("234234324-3444-4556-224232432")
				configChangesOptionsModel.SetSource(projectInputModel)
				configChangesOptionsModel.SetPullRequest("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				configChangesOptionsModel.SetTarget(projectInputModel)
				configChangesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(configChangesOptionsModel).ToNot(BeNil())
				Expect(configChangesOptionsModel.ID).To(Equal(core.StringPtr("234234324-3444-4556-224232432")))
				Expect(configChangesOptionsModel.Source).To(Equal(projectInputModel))
				Expect(configChangesOptionsModel.PullRequest).To(Equal(core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")))
				Expect(configChangesOptionsModel.Target).To(Equal(projectInputModel))
				Expect(configChangesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewConfigSettingItem successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := projectsService.NewConfigSettingItem(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateProjectOptions successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")
				Expect(inputVariableModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Type).To(Equal(core.StringPtr("array")))
				Expect(inputVariableModel.Required).To(Equal(core.BoolPtr(true)))
				Expect(inputVariableModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Default).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				Expect(configSettingItemModel).ToNot(BeNil())
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")
				Expect(configSettingItemModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(configSettingItemModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				Expect(projectConfigInputModel).ToNot(BeNil())
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}
				Expect(projectConfigInputModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigInputModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))
				Expect(projectConfigInputModel.LocatorID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputModel.Input).To(Equal([]projectsv1.InputVariable{*inputVariableModel}))
				Expect(projectConfigInputModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigInputModel.Setting).To(Equal([]projectsv1.ConfigSettingItem{*configSettingItemModel}))

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsName := "acme-microservice"
				createProjectOptionsModel := projectsService.NewCreateProjectOptions(createProjectOptionsName)
				createProjectOptionsModel.SetName("acme-microservice")
				createProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.SetConfigs([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel})
				createProjectOptionsModel.SetResourceGroup("Default")
				createProjectOptionsModel.SetLocation("us-south")
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(createProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(createProjectOptionsModel.Configs).To(Equal([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel}))
				Expect(createProjectOptionsModel.ResourceGroup).To(Equal(core.StringPtr("Default")))
				Expect(createProjectOptionsModel.Location).To(Equal(core.StringPtr("us-south")))
				Expect(createProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteNotificationOptions successfully`, func() {
				// Construct an instance of the DeleteNotificationOptions model
				id := "testString"
				notificationID := "testString"
				deleteNotificationOptionsModel := projectsService.NewDeleteNotificationOptions(id, notificationID)
				deleteNotificationOptionsModel.SetID("testString")
				deleteNotificationOptionsModel.SetNotificationID("testString")
				deleteNotificationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteNotificationOptionsModel).ToNot(BeNil())
				Expect(deleteNotificationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteNotificationOptionsModel.NotificationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteNotificationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectOptions successfully`, func() {
				// Construct an instance of the DeleteProjectOptions model
				id := "testString"
				deleteProjectOptionsModel := projectsService.NewDeleteProjectOptions(id)
				deleteProjectOptionsModel.SetID("testString")
				deleteProjectOptionsModel.SetBranch("testString")
				deleteProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectOptionsModel).ToNot(BeNil())
				Expect(deleteProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectOptionsModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeprovisionServiceInstanceOptions successfully`, func() {
				// Construct an instance of the DeprovisionServiceInstanceOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				planID := "e1031479-4b41-4159-b7cf-f7792c616fdc"
				serviceID := "cb54391b-3316-4943-a5a6-a541678c1924"
				deprovisionServiceInstanceOptionsModel := projectsService.NewDeprovisionServiceInstanceOptions(instanceID, planID, serviceID)
				deprovisionServiceInstanceOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deprovisionServiceInstanceOptionsModel.SetPlanID("e1031479-4b41-4159-b7cf-f7792c616fdc")
				deprovisionServiceInstanceOptionsModel.SetServiceID("cb54391b-3316-4943-a5a6-a541678c1924")
				deprovisionServiceInstanceOptionsModel.SetXBrokerApiVersion("1")
				deprovisionServiceInstanceOptionsModel.SetXBrokerApiOriginatingIdentity("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deprovisionServiceInstanceOptionsModel.SetAcceptsIncomplete(false)
				deprovisionServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deprovisionServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(deprovisionServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(deprovisionServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("e1031479-4b41-4159-b7cf-f7792c616fdc")))
				Expect(deprovisionServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")))
				Expect(deprovisionServiceInstanceOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1")))
				Expect(deprovisionServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity).To(Equal(core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
				Expect(deprovisionServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(false)))
				Expect(deprovisionServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeregisterPullRequestOptions successfully`, func() {
				// Construct an instance of the DeregisterPullRequestOptions model
				id := "testString"
				url := "testString"
				deregisterPullRequestOptionsModel := projectsService.NewDeregisterPullRequestOptions(id, url)
				deregisterPullRequestOptionsModel.SetID("testString")
				deregisterPullRequestOptionsModel.SetURL("testString")
				deregisterPullRequestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deregisterPullRequestOptionsModel).ToNot(BeNil())
				Expect(deregisterPullRequestOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deregisterPullRequestOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(deregisterPullRequestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogOptions successfully`, func() {
				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := projectsService.NewGetCatalogOptions()
				getCatalogOptionsModel.SetXBrokerApiVersion("1")
				getCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogOptionsModel).ToNot(BeNil())
				Expect(getCatalogOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1")))
				Expect(getCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCostEstimateOptions successfully`, func() {
				// Construct an instance of the GetCostEstimateOptions model
				id := "testString"
				configName := "testString"
				getCostEstimateOptionsModel := projectsService.NewGetCostEstimateOptions(id, configName)
				getCostEstimateOptionsModel.SetID("testString")
				getCostEstimateOptionsModel.SetConfigName("testString")
				getCostEstimateOptionsModel.SetPullRequest("testString")
				getCostEstimateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCostEstimateOptionsModel).ToNot(BeNil())
				Expect(getCostEstimateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getCostEstimateOptionsModel.ConfigName).To(Equal(core.StringPtr("testString")))
				Expect(getCostEstimateOptionsModel.PullRequest).To(Equal(core.StringPtr("testString")))
				Expect(getCostEstimateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHealthOptions successfully`, func() {
				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := projectsService.NewGetHealthOptions()
				getHealthOptionsModel.SetInfo(false)
				getHealthOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHealthOptionsModel).ToNot(BeNil())
				Expect(getHealthOptionsModel.Info).To(Equal(core.BoolPtr(false)))
				Expect(getHealthOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLastOperationOptions successfully`, func() {
				// Construct an instance of the GetLastOperationOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				getLastOperationOptionsModel := projectsService.NewGetLastOperationOptions(instanceID)
				getLastOperationOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.SetXBrokerApiVersion("1")
				getLastOperationOptionsModel.SetOperation("ABCD")
				getLastOperationOptionsModel.SetPlanID("e1031479-4b41-4159-b7cf-f7792c616fdc")
				getLastOperationOptionsModel.SetServiceID("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLastOperationOptionsModel).ToNot(BeNil())
				Expect(getLastOperationOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(getLastOperationOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1")))
				Expect(getLastOperationOptionsModel.Operation).To(Equal(core.StringPtr("ABCD")))
				Expect(getLastOperationOptionsModel.PlanID).To(Equal(core.StringPtr("e1031479-4b41-4159-b7cf-f7792c616fdc")))
				Expect(getLastOperationOptionsModel.ServiceID).To(Equal(core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")))
				Expect(getLastOperationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetNotificationsOptions successfully`, func() {
				// Construct an instance of the GetNotificationsOptions model
				id := "testString"
				getNotificationsOptionsModel := projectsService.NewGetNotificationsOptions(id)
				getNotificationsOptionsModel.SetID("testString")
				getNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getNotificationsOptionsModel).ToNot(BeNil())
				Expect(getNotificationsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectComputedStatusOptions successfully`, func() {
				// Construct an instance of the GetProjectComputedStatusOptions model
				id := "234234324-3444-4556-224232432"
				computedStatus := "cost"
				getProjectComputedStatusOptionsModel := projectsService.NewGetProjectComputedStatusOptions(id, computedStatus)
				getProjectComputedStatusOptionsModel.SetID("234234324-3444-4556-224232432")
				getProjectComputedStatusOptionsModel.SetComputedStatus("cost")
				getProjectComputedStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectComputedStatusOptionsModel).ToNot(BeNil())
				Expect(getProjectComputedStatusOptionsModel.ID).To(Equal(core.StringPtr("234234324-3444-4556-224232432")))
				Expect(getProjectComputedStatusOptionsModel.ComputedStatus).To(Equal(core.StringPtr("cost")))
				Expect(getProjectComputedStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectConfigComputedStatusOptions successfully`, func() {
				// Construct an instance of the GetProjectConfigComputedStatusOptions model
				id := "b0a2c11d-926c-4653-a15b-ed17d7b34b22"
				configName := "example"
				computedStatus := "cost"
				getProjectConfigComputedStatusOptionsModel := projectsService.NewGetProjectConfigComputedStatusOptions(id, configName, computedStatus)
				getProjectConfigComputedStatusOptionsModel.SetID("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigComputedStatusOptionsModel.SetConfigName("example")
				getProjectConfigComputedStatusOptionsModel.SetComputedStatus("cost")
				getProjectConfigComputedStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectConfigComputedStatusOptionsModel).ToNot(BeNil())
				Expect(getProjectConfigComputedStatusOptionsModel.ID).To(Equal(core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")))
				Expect(getProjectConfigComputedStatusOptionsModel.ConfigName).To(Equal(core.StringPtr("example")))
				Expect(getProjectConfigComputedStatusOptionsModel.ComputedStatus).To(Equal(core.StringPtr("cost")))
				Expect(getProjectConfigComputedStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectConfigStatusOptions successfully`, func() {
				// Construct an instance of the GetProjectConfigStatusOptions model
				id := "b0a2c11d-926c-4653-a15b-ed17d7b34b22"
				configName := "example"
				getProjectConfigStatusOptionsModel := projectsService.NewGetProjectConfigStatusOptions(id, configName)
				getProjectConfigStatusOptionsModel.SetID("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				getProjectConfigStatusOptionsModel.SetConfigName("example")
				getProjectConfigStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectConfigStatusOptionsModel).ToNot(BeNil())
				Expect(getProjectConfigStatusOptionsModel.ID).To(Equal(core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")))
				Expect(getProjectConfigStatusOptionsModel.ConfigName).To(Equal(core.StringPtr("example")))
				Expect(getProjectConfigStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectFileOptions successfully`, func() {
				// Construct an instance of the GetProjectFileOptions model
				id := "testString"
				filePath := "testString"
				getProjectFileOptionsModel := projectsService.NewGetProjectFileOptions(id, filePath)
				getProjectFileOptionsModel.SetID("testString")
				getProjectFileOptionsModel.SetFilePath("testString")
				getProjectFileOptionsModel.SetBranch("testString")
				getProjectFileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectFileOptionsModel).ToNot(BeNil())
				Expect(getProjectFileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProjectFileOptionsModel.FilePath).To(Equal(core.StringPtr("testString")))
				Expect(getProjectFileOptionsModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(getProjectFileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectOptions successfully`, func() {
				// Construct an instance of the GetProjectOptions model
				id := "testString"
				getProjectOptionsModel := projectsService.NewGetProjectOptions(id)
				getProjectOptionsModel.SetID("testString")
				getProjectOptionsModel.SetBranch("testString")
				getProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectOptionsModel).ToNot(BeNil())
				Expect(getProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProjectOptionsModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(getProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectStatusOptions successfully`, func() {
				// Construct an instance of the GetProjectStatusOptions model
				id := "234234324-3444-4556-224232432"
				getProjectStatusOptionsModel := projectsService.NewGetProjectStatusOptions(id)
				getProjectStatusOptionsModel.SetID("234234324-3444-4556-224232432")
				getProjectStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectStatusOptionsModel).ToNot(BeNil())
				Expect(getProjectStatusOptionsModel.ID).To(Equal(core.StringPtr("234234324-3444-4556-224232432")))
				Expect(getProjectStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSchematicsJobOptions successfully`, func() {
				// Construct an instance of the GetSchematicsJobOptions model
				id := "testString"
				configName := "testString"
				action := "plan"
				getSchematicsJobOptionsModel := projectsService.NewGetSchematicsJobOptions(id, configName, action)
				getSchematicsJobOptionsModel.SetID("testString")
				getSchematicsJobOptionsModel.SetConfigName("testString")
				getSchematicsJobOptionsModel.SetAction("plan")
				getSchematicsJobOptionsModel.SetSince(int64(38))
				getSchematicsJobOptionsModel.SetPullRequest("testString")
				getSchematicsJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSchematicsJobOptionsModel).ToNot(BeNil())
				Expect(getSchematicsJobOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getSchematicsJobOptionsModel.ConfigName).To(Equal(core.StringPtr("testString")))
				Expect(getSchematicsJobOptionsModel.Action).To(Equal(core.StringPtr("plan")))
				Expect(getSchematicsJobOptionsModel.Since).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getSchematicsJobOptionsModel.PullRequest).To(Equal(core.StringPtr("testString")))
				Expect(getSchematicsJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServiceInstanceOptions successfully`, func() {
				// Construct an instance of the GetServiceInstanceOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				getServiceInstanceOptionsModel := projectsService.NewGetServiceInstanceOptions(instanceID)
				getServiceInstanceOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.SetXBrokerApiVersion("1")
				getServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(getServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(getServiceInstanceOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1")))
				Expect(getServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGitLabEvent successfully`, func() {
				objectKind := "testString"
				_model, err := projectsService.NewGitLabEvent(objectKind)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewInputVariable successfully`, func() {
				name := "testString"
				typeVar := "array"
				_model, err := projectsService.NewInputVariable(name, typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewInstallConfigOptions successfully`, func() {
				// Construct an instance of the InstallConfigOptions model
				id := "testString"
				configName := "testString"
				installConfigOptionsModel := projectsService.NewInstallConfigOptions(id, configName)
				installConfigOptionsModel.SetID("testString")
				installConfigOptionsModel.SetConfigName("testString")
				installConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(installConfigOptionsModel).ToNot(BeNil())
				Expect(installConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(installConfigOptionsModel.ConfigName).To(Equal(core.StringPtr("testString")))
				Expect(installConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectConfigStatusesOptions successfully`, func() {
				// Construct an instance of the ListProjectConfigStatusesOptions model
				id := "testString"
				listProjectConfigStatusesOptionsModel := projectsService.NewListProjectConfigStatusesOptions(id)
				listProjectConfigStatusesOptionsModel.SetID("testString")
				listProjectConfigStatusesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectConfigStatusesOptionsModel).ToNot(BeNil())
				Expect(listProjectConfigStatusesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listProjectConfigStatusesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectsOptions successfully`, func() {
				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := projectsService.NewListProjectsOptions()
				listProjectsOptionsModel.SetStart("testString")
				listProjectsOptionsModel.SetLimit(int64(10))
				listProjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectsOptionsModel).ToNot(BeNil())
				Expect(listProjectsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listProjectsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listProjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMergeProjectOptions successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")
				Expect(inputVariableModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Type).To(Equal(core.StringPtr("array")))
				Expect(inputVariableModel.Required).To(Equal(core.BoolPtr(true)))
				Expect(inputVariableModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Default).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				Expect(configSettingItemModel).ToNot(BeNil())
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")
				Expect(configSettingItemModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(configSettingItemModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				Expect(projectConfigInputModel).ToNot(BeNil())
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}
				Expect(projectConfigInputModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigInputModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))
				Expect(projectConfigInputModel.LocatorID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputModel.Input).To(Equal([]projectsv1.InputVariable{*inputVariableModel}))
				Expect(projectConfigInputModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigInputModel.Setting).To(Equal([]projectsv1.ConfigSettingItem{*configSettingItemModel}))

				// Construct an instance of the MergeProjectOptions model
				id := "testString"
				mergeProjectOptionsName := "acme-microservice"
				mergeProjectOptionsModel := projectsService.NewMergeProjectOptions(id, mergeProjectOptionsName)
				mergeProjectOptionsModel.SetID("testString")
				mergeProjectOptionsModel.SetName("acme-microservice")
				mergeProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.SetConfigs([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel})
				mergeProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(mergeProjectOptionsModel).ToNot(BeNil())
				Expect(mergeProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(mergeProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(mergeProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(mergeProjectOptionsModel.Configs).To(Equal([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel}))
				Expect(mergeProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewNotificationEvent successfully`, func() {
				event := "testString"
				target := "testString"
				_model, err := projectsService.NewNotificationEvent(event, target)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNotifyOptions successfully`, func() {
				// Construct an instance of the NotifyOptions model
				notifyOptionsID := "bccbb195-fff4-4d7e-9078-61b06adc02ab"
				notifyOptionsModel := projectsService.NewNotifyOptions(notifyOptionsID)
				notifyOptionsModel.SetID("bccbb195-fff4-4d7e-9078-61b06adc02ab")
				notifyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(notifyOptionsModel).ToNot(BeNil())
				Expect(notifyOptionsModel.ID).To(Equal(core.StringPtr("bccbb195-fff4-4d7e-9078-61b06adc02ab")))
				Expect(notifyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewOutputValue successfully`, func() {
				name := "testString"
				_model, err := projectsService.NewOutputValue(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPlanConfigOptions successfully`, func() {
				// Construct an instance of the PlanConfigOptions model
				id := "testString"
				configName := "testString"
				planConfigOptionsModel := projectsService.NewPlanConfigOptions(id, configName)
				planConfigOptionsModel.SetID("testString")
				planConfigOptionsModel.SetConfigName("testString")
				planConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(planConfigOptionsModel).ToNot(BeNil())
				Expect(planConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(planConfigOptionsModel.ConfigName).To(Equal(core.StringPtr("testString")))
				Expect(planConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPlanPullRequestConfigsOptions successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")
				Expect(inputVariableModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Type).To(Equal(core.StringPtr("array")))
				Expect(inputVariableModel.Required).To(Equal(core.BoolPtr(true)))
				Expect(inputVariableModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Default).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				Expect(configSettingItemModel).ToNot(BeNil())
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")
				Expect(configSettingItemModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(configSettingItemModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				Expect(projectConfigInputModel).ToNot(BeNil())
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}
				Expect(projectConfigInputModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigInputModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))
				Expect(projectConfigInputModel.LocatorID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputModel.Input).To(Equal([]projectsv1.InputVariable{*inputVariableModel}))
				Expect(projectConfigInputModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigInputModel.Setting).To(Equal([]projectsv1.ConfigSettingItem{*configSettingItemModel}))

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				Expect(projectInputModel).ToNot(BeNil())
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				Expect(projectInputModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(projectInputModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(projectInputModel.Configs).To(Equal([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel}))

				// Construct an instance of the PlanPullRequestConfigsOptions model
				id := "testString"
				planPullRequestConfigsOptionsPullRequest := "https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1"
				var planPullRequestConfigsOptionsProjectDefinition *projectsv1.ProjectInput = nil
				planPullRequestConfigsOptionsModel := projectsService.NewPlanPullRequestConfigsOptions(id, planPullRequestConfigsOptionsPullRequest, planPullRequestConfigsOptionsProjectDefinition)
				planPullRequestConfigsOptionsModel.SetID("testString")
				planPullRequestConfigsOptionsModel.SetPullRequest("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				planPullRequestConfigsOptionsModel.SetProjectDefinition(projectInputModel)
				planPullRequestConfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(planPullRequestConfigsOptionsModel).ToNot(BeNil())
				Expect(planPullRequestConfigsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(planPullRequestConfigsOptionsModel.PullRequest).To(Equal(core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")))
				Expect(planPullRequestConfigsOptionsModel.ProjectDefinition).To(Equal(projectInputModel))
				Expect(planPullRequestConfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostNotificationOptions successfully`, func() {
				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectsv1.NotificationEvent)
				Expect(notificationEventModel).ToNot(BeNil())
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.ActionURL = core.StringPtr("url.for.project.documentation")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				Expect(notificationEventModel.Event).To(Equal(core.StringPtr("project.create.failed")))
				Expect(notificationEventModel.Target).To(Equal(core.StringPtr("234234324-3444-4556-224232432")))
				Expect(notificationEventModel.Source).To(Equal(core.StringPtr("id.of.project.service.instance")))
				Expect(notificationEventModel.ActionURL).To(Equal(core.StringPtr("url.for.project.documentation")))
				Expect(notificationEventModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the PostNotificationOptions model
				id := "testString"
				postNotificationOptionsModel := projectsService.NewPostNotificationOptions(id)
				postNotificationOptionsModel.SetID("testString")
				postNotificationOptionsModel.SetNotifications([]projectsv1.NotificationEvent{*notificationEventModel})
				postNotificationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postNotificationOptionsModel).ToNot(BeNil())
				Expect(postNotificationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(postNotificationOptionsModel.Notifications).To(Equal([]projectsv1.NotificationEvent{*notificationEventModel}))
				Expect(postNotificationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewProjectInput successfully`, func() {
				name := "testString"
				_model, err := projectsService.NewProjectInput(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProvisionServiceInstanceOptions successfully`, func() {
				// Construct an instance of the ProvisionServiceInstanceOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				provisionServiceInstanceOptionsServiceID := "testString"
				provisionServiceInstanceOptionsPlanID := "testString"
				provisionServiceInstanceOptionsModel := projectsService.NewProvisionServiceInstanceOptions(instanceID, provisionServiceInstanceOptionsServiceID, provisionServiceInstanceOptionsPlanID)
				provisionServiceInstanceOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				provisionServiceInstanceOptionsModel.SetServiceID("testString")
				provisionServiceInstanceOptionsModel.SetPlanID("testString")
				provisionServiceInstanceOptionsModel.SetContext([]string{"testString"})
				provisionServiceInstanceOptionsModel.SetParameters(map[string]interface{}{"anyKey": "anyValue"})
				provisionServiceInstanceOptionsModel.SetPreviousValues([]string{"testString"})
				provisionServiceInstanceOptionsModel.SetXBrokerApiVersion("1")
				provisionServiceInstanceOptionsModel.SetXBrokerApiOriginatingIdentity("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				provisionServiceInstanceOptionsModel.SetAcceptsIncomplete(false)
				provisionServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(provisionServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(provisionServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(provisionServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(provisionServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(provisionServiceInstanceOptionsModel.Context).To(Equal([]string{"testString"}))
				Expect(provisionServiceInstanceOptionsModel.Parameters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(provisionServiceInstanceOptionsModel.PreviousValues).To(Equal([]string{"testString"}))
				Expect(provisionServiceInstanceOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1")))
				Expect(provisionServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity).To(Equal(core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
				Expect(provisionServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(false)))
				Expect(provisionServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPulsarEventItem successfully`, func() {
				eventType := "testString"
				timestamp := "testString"
				publisher := "testString"
				accountID := "testString"
				version := "testString"
				_model, err := projectsService.NewPulsarEventItem(eventType, timestamp, publisher, accountID, version)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReceiveGitlabEventsOptions successfully`, func() {
				// Construct an instance of the GitLabEvent model
				gitLabEventModel := new(projectsv1.GitLabEvent)
				Expect(gitLabEventModel).ToNot(BeNil())
				gitLabEventModel.ObjectKind = core.StringPtr("testString")
				gitLabEventModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(gitLabEventModel.ObjectKind).To(Equal(core.StringPtr("testString")))
				Expect(gitLabEventModel.GetProperties()).ToNot(BeEmpty())
				Expect(gitLabEventModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				gitLabEventModel.SetProperties(nil)
				Expect(gitLabEventModel.GetProperties()).To(BeEmpty())

				gitLabEventModelExpectedMap := make(map[string]interface{})
				gitLabEventModelExpectedMap["foo"] = core.StringPtr("testString")
				gitLabEventModel.SetProperties(gitLabEventModelExpectedMap)
				gitLabEventModelActualMap := gitLabEventModel.GetProperties()
				Expect(gitLabEventModelActualMap).To(Equal(gitLabEventModelExpectedMap))

				// Construct an instance of the ReceiveGitlabEventsOptions model
				id := "testString"
				var gitLabEvent *projectsv1.GitLabEvent = nil
				receiveGitlabEventsOptionsModel := projectsService.NewReceiveGitlabEventsOptions(id, gitLabEvent)
				receiveGitlabEventsOptionsModel.SetID("testString")
				receiveGitlabEventsOptionsModel.SetGitLabEvent(gitLabEventModel)
				receiveGitlabEventsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(receiveGitlabEventsOptionsModel).ToNot(BeNil())
				Expect(receiveGitlabEventsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(receiveGitlabEventsOptionsModel.GitLabEvent).To(Equal(gitLabEventModel))
				Expect(receiveGitlabEventsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReceivePulsarCatalogEventsOptions successfully`, func() {
				// Construct an instance of the PulsarEventItem model
				pulsarEventItemModel := new(projectsv1.PulsarEventItem)
				Expect(pulsarEventItemModel).ToNot(BeNil())
				pulsarEventItemModel.EventType = core.StringPtr("testString")
				pulsarEventItemModel.Timestamp = core.StringPtr("testString")
				pulsarEventItemModel.Publisher = core.StringPtr("testString")
				pulsarEventItemModel.AccountID = core.StringPtr("testString")
				pulsarEventItemModel.Version = core.StringPtr("testString")
				pulsarEventItemModel.EventProperties = map[string]interface{}{"anyKey": "anyValue"}
				pulsarEventItemModel.EventID = core.StringPtr("testString")
				pulsarEventItemModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(pulsarEventItemModel.EventType).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventItemModel.Timestamp).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventItemModel.Publisher).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventItemModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventItemModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventItemModel.EventProperties).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(pulsarEventItemModel.EventID).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventItemModel.GetProperties()).ToNot(BeEmpty())
				Expect(pulsarEventItemModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				pulsarEventItemModel.SetProperties(nil)
				Expect(pulsarEventItemModel.GetProperties()).To(BeEmpty())

				pulsarEventItemModelExpectedMap := make(map[string]interface{})
				pulsarEventItemModelExpectedMap["foo"] = core.StringPtr("testString")
				pulsarEventItemModel.SetProperties(pulsarEventItemModelExpectedMap)
				pulsarEventItemModelActualMap := pulsarEventItemModel.GetProperties()
				Expect(pulsarEventItemModelActualMap).To(Equal(pulsarEventItemModelExpectedMap))

				// Construct an instance of the ReceivePulsarCatalogEventsOptions model
				pulsarEventItem := []projectsv1.PulsarEventItem{}
				receivePulsarCatalogEventsOptionsModel := projectsService.NewReceivePulsarCatalogEventsOptions(pulsarEventItem)
				receivePulsarCatalogEventsOptionsModel.SetPulsarEventItem([]projectsv1.PulsarEventItem{*pulsarEventItemModel})
				receivePulsarCatalogEventsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(receivePulsarCatalogEventsOptionsModel).ToNot(BeNil())
				Expect(receivePulsarCatalogEventsOptionsModel.PulsarEventItem).To(Equal([]projectsv1.PulsarEventItem{*pulsarEventItemModel}))
				Expect(receivePulsarCatalogEventsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRegisterPullRequestOptions successfully`, func() {
				// Construct an instance of the RegisterPullRequestOptions model
				id := "testString"
				registerPullRequestOptionsModel := projectsService.NewRegisterPullRequestOptions(id)
				registerPullRequestOptionsModel.SetID("testString")
				registerPullRequestOptionsModel.SetBranch("test")
				registerPullRequestOptionsModel.SetURL("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				registerPullRequestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(registerPullRequestOptionsModel).ToNot(BeNil())
				Expect(registerPullRequestOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(registerPullRequestOptionsModel.Branch).To(Equal(core.StringPtr("test")))
				Expect(registerPullRequestOptionsModel.URL).To(Equal(core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")))
				Expect(registerPullRequestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceProjectOptions successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")
				Expect(inputVariableModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Type).To(Equal(core.StringPtr("array")))
				Expect(inputVariableModel.Required).To(Equal(core.BoolPtr(true)))
				Expect(inputVariableModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Default).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				Expect(configSettingItemModel).ToNot(BeNil())
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")
				Expect(configSettingItemModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(configSettingItemModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				Expect(projectConfigInputModel).ToNot(BeNil())
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}
				Expect(projectConfigInputModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigInputModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))
				Expect(projectConfigInputModel.LocatorID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputModel.Input).To(Equal([]projectsv1.InputVariable{*inputVariableModel}))
				Expect(projectConfigInputModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigInputModel.Setting).To(Equal([]projectsv1.ConfigSettingItem{*configSettingItemModel}))

				// Construct an instance of the ReplaceProjectOptions model
				id := "testString"
				replaceProjectOptionsName := "acme-microservice"
				replaceProjectOptionsModel := projectsService.NewReplaceProjectOptions(id, replaceProjectOptionsName)
				replaceProjectOptionsModel.SetID("testString")
				replaceProjectOptionsModel.SetName("acme-microservice")
				replaceProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure")
				replaceProjectOptionsModel.SetConfigs([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel})
				replaceProjectOptionsModel.SetBranch("testString")
				replaceProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceProjectOptionsModel).ToNot(BeNil())
				Expect(replaceProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(replaceProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(replaceProjectOptionsModel.Configs).To(Equal([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel}))
				Expect(replaceProjectOptionsModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(replaceProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUninstallConfigOptions successfully`, func() {
				// Construct an instance of the UninstallConfigOptions model
				id := "testString"
				configName := "testString"
				uninstallConfigOptionsModel := projectsService.NewUninstallConfigOptions(id, configName)
				uninstallConfigOptionsModel.SetID("testString")
				uninstallConfigOptionsModel.SetConfigName("testString")
				uninstallConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uninstallConfigOptionsModel).ToNot(BeNil())
				Expect(uninstallConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(uninstallConfigOptionsModel.ConfigName).To(Equal(core.StringPtr("testString")))
				Expect(uninstallConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectComputedStatusOptions successfully`, func() {
				// Construct an instance of the UpdateProjectComputedStatusOptions model
				id := "234234324-3444-4556-224232432"
				computedStatus := "cost"
				updateProjectComputedStatusOptionsComputedStatuses := make(map[string]interface{})
				updateProjectComputedStatusOptionsModel := projectsService.NewUpdateProjectComputedStatusOptions(id, computedStatus, updateProjectComputedStatusOptionsComputedStatuses)
				updateProjectComputedStatusOptionsModel.SetID("234234324-3444-4556-224232432")
				updateProjectComputedStatusOptionsModel.SetComputedStatus("cost")
				updateProjectComputedStatusOptionsModel.SetComputedStatuses(make(map[string]interface{}))
				updateProjectComputedStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectComputedStatusOptionsModel).ToNot(BeNil())
				Expect(updateProjectComputedStatusOptionsModel.ID).To(Equal(core.StringPtr("234234324-3444-4556-224232432")))
				Expect(updateProjectComputedStatusOptionsModel.ComputedStatus).To(Equal(core.StringPtr("cost")))
				Expect(updateProjectComputedStatusOptionsModel.ComputedStatuses).To(Equal(make(map[string]interface{})))
				Expect(updateProjectComputedStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectConfigComputedStatusOptions successfully`, func() {
				// Construct an instance of the UpdateProjectConfigComputedStatusOptions model
				id := "b0a2c11d-926c-4653-a15b-ed17d7b34b22"
				configName := "example"
				computedStatus := "cost"
				requestBody := make(map[string]interface{})
				updateProjectConfigComputedStatusOptionsModel := projectsService.NewUpdateProjectConfigComputedStatusOptions(id, configName, computedStatus, requestBody)
				updateProjectConfigComputedStatusOptionsModel.SetID("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigComputedStatusOptionsModel.SetConfigName("example")
				updateProjectConfigComputedStatusOptionsModel.SetComputedStatus("cost")
				updateProjectConfigComputedStatusOptionsModel.SetRequestBody(make(map[string]interface{}))
				updateProjectConfigComputedStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectConfigComputedStatusOptionsModel).ToNot(BeNil())
				Expect(updateProjectConfigComputedStatusOptionsModel.ID).To(Equal(core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")))
				Expect(updateProjectConfigComputedStatusOptionsModel.ConfigName).To(Equal(core.StringPtr("example")))
				Expect(updateProjectConfigComputedStatusOptionsModel.ComputedStatus).To(Equal(core.StringPtr("cost")))
				Expect(updateProjectConfigComputedStatusOptionsModel.RequestBody).To(Equal(make(map[string]interface{})))
				Expect(updateProjectConfigComputedStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectStatusOptions successfully`, func() {
				// Construct an instance of the History model
				historyModel := new(projectsv1.History)
				Expect(historyModel).ToNot(BeNil())
				historyModel.Timestamp = CreateMockDateTime("2022-05-13T12:25:00Z")
				historyModel.Code = core.StringPtr("ISB005I")
				historyModel.Type = core.StringPtr("schematics")
				Expect(historyModel.Timestamp).To(Equal(CreateMockDateTime("2022-05-13T12:25:00Z")))
				Expect(historyModel.Code).To(Equal(core.StringPtr("ISB005I")))
				Expect(historyModel.Type).To(Equal(core.StringPtr("schematics")))

				// Construct an instance of the ServiceInfoGit model
				serviceInfoGitModel := new(projectsv1.ServiceInfoGit)
				Expect(serviceInfoGitModel).ToNot(BeNil())
				serviceInfoGitModel.URL = core.StringPtr("testString")
				serviceInfoGitModel.Branch = core.StringPtr("testString")
				serviceInfoGitModel.ProjectID = core.StringPtr("testString")
				Expect(serviceInfoGitModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(serviceInfoGitModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(serviceInfoGitModel.ProjectID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ServiceInfoToolchain model
				serviceInfoToolchainModel := new(projectsv1.ServiceInfoToolchain)
				Expect(serviceInfoToolchainModel).ToNot(BeNil())
				serviceInfoToolchainModel.ID = core.StringPtr("testString")
				Expect(serviceInfoToolchainModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ServiceInfoSchematics model
				serviceInfoSchematicsModel := new(projectsv1.ServiceInfoSchematics)
				Expect(serviceInfoSchematicsModel).ToNot(BeNil())
				serviceInfoSchematicsModel.CartOrderID = core.StringPtr("us-south.cartOrder.5fbc28c0-87fd-42f5-b2fa-d98d28d6ea96-local.d807263d")
				serviceInfoSchematicsModel.WorkspaceID = core.StringPtr("crn:v1:staging:public:schematics:us-south:a/d44e54841dfe41ee869659f80bddd075:459a39a6-3bc3-4524-aaa8-fb693d96b79e:workspace:us-south.workspace.projects-service.a482a249")
				serviceInfoSchematicsModel.CartItemName = core.StringPtr("name")
				Expect(serviceInfoSchematicsModel.CartOrderID).To(Equal(core.StringPtr("us-south.cartOrder.5fbc28c0-87fd-42f5-b2fa-d98d28d6ea96-local.d807263d")))
				Expect(serviceInfoSchematicsModel.WorkspaceID).To(Equal(core.StringPtr("crn:v1:staging:public:schematics:us-south:a/d44e54841dfe41ee869659f80bddd075:459a39a6-3bc3-4524-aaa8-fb693d96b79e:workspace:us-south.workspace.projects-service.a482a249")))
				Expect(serviceInfoSchematicsModel.CartItemName).To(Equal(core.StringPtr("name")))

				// Construct an instance of the UpdateProjectStatusOptions model
				id := "234234324-3444-4556-224232432"
				updateProjectStatusOptionsState := "UPDATING"
				updateProjectStatusOptionsHistory := []projectsv1.History{}
				updateProjectStatusOptionsModel := projectsService.NewUpdateProjectStatusOptions(id, updateProjectStatusOptionsState, updateProjectStatusOptionsHistory)
				updateProjectStatusOptionsModel.SetID("234234324-3444-4556-224232432")
				updateProjectStatusOptionsModel.SetState("UPDATING")
				updateProjectStatusOptionsModel.SetHistory([]projectsv1.History{*historyModel})
				updateProjectStatusOptionsModel.SetGitRepo(serviceInfoGitModel)
				updateProjectStatusOptionsModel.SetToolchain(serviceInfoToolchainModel)
				updateProjectStatusOptionsModel.SetSchematics(serviceInfoSchematicsModel)
				updateProjectStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectStatusOptionsModel).ToNot(BeNil())
				Expect(updateProjectStatusOptionsModel.ID).To(Equal(core.StringPtr("234234324-3444-4556-224232432")))
				Expect(updateProjectStatusOptionsModel.State).To(Equal(core.StringPtr("UPDATING")))
				Expect(updateProjectStatusOptionsModel.History).To(Equal([]projectsv1.History{*historyModel}))
				Expect(updateProjectStatusOptionsModel.GitRepo).To(Equal(serviceInfoGitModel))
				Expect(updateProjectStatusOptionsModel.Toolchain).To(Equal(serviceInfoToolchainModel))
				Expect(updateProjectStatusOptionsModel.Schematics).To(Equal(serviceInfoSchematicsModel))
				Expect(updateProjectStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePullRequestConfigsOptions successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")
				Expect(inputVariableModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Type).To(Equal(core.StringPtr("array")))
				Expect(inputVariableModel.Required).To(Equal(core.BoolPtr(true)))
				Expect(inputVariableModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Default).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				Expect(configSettingItemModel).ToNot(BeNil())
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")
				Expect(configSettingItemModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(configSettingItemModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				Expect(projectConfigInputModel).ToNot(BeNil())
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}
				Expect(projectConfigInputModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigInputModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))
				Expect(projectConfigInputModel.LocatorID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputModel.Input).To(Equal([]projectsv1.InputVariable{*inputVariableModel}))
				Expect(projectConfigInputModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigInputModel.Setting).To(Equal([]projectsv1.ConfigSettingItem{*configSettingItemModel}))

				// Construct an instance of the ProjectInput model
				projectInputModel := new(projectsv1.ProjectInput)
				Expect(projectInputModel).ToNot(BeNil())
				projectInputModel.Name = core.StringPtr("acme-microservice")
				projectInputModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				projectInputModel.Configs = []projectsv1.ProjectConfigInputIntf{projectConfigInputModel}
				Expect(projectInputModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(projectInputModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(projectInputModel.Configs).To(Equal([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel}))

				// Construct an instance of the UpdatePullRequestConfigsOptions model
				id := "testString"
				updatePullRequestConfigsOptionsPullRequest := "https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1"
				updatePullRequestConfigsOptionsModel := projectsService.NewUpdatePullRequestConfigsOptions(id, updatePullRequestConfigsOptionsPullRequest)
				updatePullRequestConfigsOptionsModel.SetID("testString")
				updatePullRequestConfigsOptionsModel.SetPullRequest("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")
				updatePullRequestConfigsOptionsModel.SetProjectDefinition(projectInputModel)
				updatePullRequestConfigsOptionsModel.SetState("merge")
				updatePullRequestConfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePullRequestConfigsOptionsModel).ToNot(BeNil())
				Expect(updatePullRequestConfigsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updatePullRequestConfigsOptionsModel.PullRequest).To(Equal(core.StringPtr("https://dev.us-south.git.test.cloud.ibm.com/org/projects-poc/-/merge_requests/1")))
				Expect(updatePullRequestConfigsOptionsModel.ProjectDefinition).To(Equal(projectInputModel))
				Expect(updatePullRequestConfigsOptionsModel.State).To(Equal(core.StringPtr("merge")))
				Expect(updatePullRequestConfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateServiceInstanceOptions successfully`, func() {
				// Construct an instance of the UpdateServiceInstanceOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				updateServiceInstanceOptionsServiceID := []string{"testString"}
				updateServiceInstanceOptionsModel := projectsService.NewUpdateServiceInstanceOptions(instanceID, updateServiceInstanceOptionsServiceID)
				updateServiceInstanceOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.SetServiceID([]string{"testString"})
				updateServiceInstanceOptionsModel.SetContext([]string{"testString"})
				updateServiceInstanceOptionsModel.SetParameters(map[string]interface{}{"anyKey": "anyValue"})
				updateServiceInstanceOptionsModel.SetPlanID("testString")
				updateServiceInstanceOptionsModel.SetPreviousValues([]string{"testString"})
				updateServiceInstanceOptionsModel.SetXBrokerApiVersion("1")
				updateServiceInstanceOptionsModel.SetXBrokerApiOriginatingIdentity("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.SetAcceptsIncomplete(false)
				updateServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(updateServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(updateServiceInstanceOptionsModel.ServiceID).To(Equal([]string{"testString"}))
				Expect(updateServiceInstanceOptionsModel.Context).To(Equal([]string{"testString"}))
				Expect(updateServiceInstanceOptionsModel.Parameters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceInstanceOptionsModel.PreviousValues).To(Equal([]string{"testString"}))
				Expect(updateServiceInstanceOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1")))
				Expect(updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity).To(Equal(core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
				Expect(updateServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(false)))
				Expect(updateServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateServiceInstanceStateOptions successfully`, func() {
				// Construct an instance of the UpdateServiceInstanceStateOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				updateServiceInstanceStateOptionsEnabled := true
				updateServiceInstanceStateOptionsModel := projectsService.NewUpdateServiceInstanceStateOptions(instanceID, updateServiceInstanceStateOptionsEnabled)
				updateServiceInstanceStateOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceStateOptionsModel.SetEnabled(true)
				updateServiceInstanceStateOptionsModel.SetInitiatorID("testString")
				updateServiceInstanceStateOptionsModel.SetReasonCode(map[string]interface{}{"anyKey": "anyValue"})
				updateServiceInstanceStateOptionsModel.SetPlanID("testString")
				updateServiceInstanceStateOptionsModel.SetPreviousValues([]string{"testString"})
				updateServiceInstanceStateOptionsModel.SetXBrokerApiVersion("1")
				updateServiceInstanceStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateServiceInstanceStateOptionsModel).ToNot(BeNil())
				Expect(updateServiceInstanceStateOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(updateServiceInstanceStateOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateServiceInstanceStateOptionsModel.InitiatorID).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceInstanceStateOptionsModel.ReasonCode).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateServiceInstanceStateOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceInstanceStateOptionsModel.PreviousValues).To(Equal([]string{"testString"}))
				Expect(updateServiceInstanceStateOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1")))
				Expect(updateServiceInstanceStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewValidateProjectOptions successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				inputVariableModel.Value = core.StringPtr("testString")
				inputVariableModel.Default = core.StringPtr("testString")
				Expect(inputVariableModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Type).To(Equal(core.StringPtr("array")))
				Expect(inputVariableModel.Required).To(Equal(core.BoolPtr(true)))
				Expect(inputVariableModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Default).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ConfigSettingItem model
				configSettingItemModel := new(projectsv1.ConfigSettingItem)
				Expect(configSettingItemModel).ToNot(BeNil())
				configSettingItemModel.Name = core.StringPtr("testString")
				configSettingItemModel.Value = core.StringPtr("testString")
				Expect(configSettingItemModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(configSettingItemModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigInputProp model
				projectConfigInputModel := new(projectsv1.ProjectConfigInputProp)
				Expect(projectConfigInputModel).ToNot(BeNil())
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItem{*configSettingItemModel}
				Expect(projectConfigInputModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigInputModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))
				Expect(projectConfigInputModel.LocatorID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputModel.Input).To(Equal([]projectsv1.InputVariable{*inputVariableModel}))
				Expect(projectConfigInputModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigInputModel.Setting).To(Equal([]projectsv1.ConfigSettingItem{*configSettingItemModel}))

				// Construct an instance of the ValidateProjectOptions model
				validateProjectOptionsName := "acme-microservice"
				validateProjectOptionsModel := projectsService.NewValidateProjectOptions(validateProjectOptionsName)
				validateProjectOptionsModel.SetName("acme-microservice")
				validateProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure")
				validateProjectOptionsModel.SetConfigs([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel})
				validateProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(validateProjectOptionsModel).ToNot(BeNil())
				Expect(validateProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(validateProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(validateProjectOptionsModel.Configs).To(Equal([]projectsv1.ProjectConfigInputIntf{projectConfigInputModel}))
				Expect(validateProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewProjectConfigInputProp successfully`, func() {
				locatorID := "testString"
				_model, err := projectsService.NewProjectConfigInputProp(locatorID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProjectConfigInputProjectConfigCommon successfully`, func() {
				name := "testString"
				_model, err := projectsService.NewProjectConfigInputProjectConfigCommon(name)
				Expect(_model).ToNot(BeNil())
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
