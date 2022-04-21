/**
 * (C) Copyright IBM Corp. 2022.
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
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
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

				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Metadata = metadataModel
				createProjectOptionsModel.Spec = specModel
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID"}`)
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

				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Metadata = metadataModel
				createProjectOptionsModel.Spec = specModel
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID"}`)
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

				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Metadata = metadataModel
				createProjectOptionsModel.Spec = specModel
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

				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Metadata = metadataModel
				createProjectOptionsModel.Spec = specModel
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Metadata = metadataModel
				createProjectOptionsModel.Spec = specModel
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
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
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
				listProjectsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(1))
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

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 0, "limit": 1, "total_count": 0, "first": {"href": "Href"}, "last": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "projects": [{"href": "Href", "metadata": {"name": "Name", "description": "Description", "id": "ID", "created_by": "CreatedBy", "created_at": "CreatedAt", "updated_at": "UpdatedAt", "tags": ["Tags"], "repo_url": "RepoURL"}}]}`)
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
				listProjectsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(1))
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

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 0, "limit": 1, "total_count": 0, "first": {"href": "Href"}, "last": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "projects": [{"href": "Href", "metadata": {"name": "Name", "description": "Description", "id": "ID", "created_by": "CreatedBy", "created_at": "CreatedAt", "updated_at": "UpdatedAt", "tags": ["Tags"], "repo_url": "RepoURL"}}]}`)
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
				listProjectsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(1))
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
				listProjectsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(1))
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
				listProjectsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(1))
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
		It(`Invoke GetNextOffset successfully`, func() {
			responseObject := new(projectsv1.ListProjectsResponse)
			nextObject := new(projectsv1.Href)
			nextObject.Href = core.StringPtr("ibm.com?offset=135")
			responseObject.Next = nextObject

			value, err := responseObject.GetNextOffset()
			Expect(err).To(BeNil())
			Expect(value).To(Equal(core.Int64Ptr(int64(135))))
		})
		It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
			responseObject := new(projectsv1.ListProjectsResponse)

			value, err := responseObject.GetNextOffset()
			Expect(err).To(BeNil())
			Expect(value).To(BeNil())
		})
		It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
			responseObject := new(projectsv1.ListProjectsResponse)
			nextObject := new(projectsv1.Href)
			nextObject.Href = core.StringPtr("ibm.com")
			responseObject.Next = nextObject

			value, err := responseObject.GetNextOffset()
			Expect(err).To(BeNil())
			Expect(value).To(BeNil())
		})
		It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
			responseObject := new(projectsv1.ListProjectsResponse)
			nextObject := new(projectsv1.Href)
			nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
			responseObject.Next = nextObject

			value, err := responseObject.GetNextOffset()
			Expect(err).NotTo(BeNil())
			Expect(value).To(BeNil())
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"href": "Href", "metadata": {"name": "Name", "description": "Description", "id": "ID", "created_by": "CreatedBy", "created_at": "CreatedAt", "updated_at": "UpdatedAt", "tags": ["Tags"], "repo_url": "RepoURL"}, "spec": {"configs": [{"name": "Name", "labels": ["Labels"], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "type": "manual", "external_resources_account": "ExternalResourcesAccount"}], "dashboard": {"widgets": [{"name": "Name"}]}}}`)
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"href": "Href", "metadata": {"name": "Name", "description": "Description", "id": "ID", "created_by": "CreatedBy", "created_at": "CreatedAt", "updated_at": "UpdatedAt", "tags": ["Tags"], "repo_url": "RepoURL"}, "spec": {"configs": [{"name": "Name", "labels": ["Labels"], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "type": "manual", "external_resources_account": "ExternalResourcesAccount"}], "dashboard": {"widgets": [{"name": "Name"}]}}}`)
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
	Describe(`UpdateProject(updateProjectOptions *UpdateProjectOptions) - Operation response error`, func() {
		updateProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Prefer"]).ToNot(BeNil())
					Expect(req.Header["Prefer"][0]).To(Equal(fmt.Sprintf("%v", "return=minimal")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProject with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Metadata = metadataModel
				updateProjectOptionsModel.Spec = specModel
				updateProjectOptionsModel.Prefer = core.StringPtr("return=minimal")
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProject(updateProjectOptions *UpdateProjectOptions)`, func() {
		updateProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectPath))
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

					Expect(req.Header["Prefer"]).ToNot(BeNil())
					Expect(req.Header["Prefer"][0]).To(Equal(fmt.Sprintf("%v", "return=minimal")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"href": "Href", "metadata": {"name": "Name", "description": "Description", "id": "ID", "created_by": "CreatedBy", "created_at": "CreatedAt", "updated_at": "UpdatedAt", "tags": ["Tags"], "repo_url": "RepoURL"}, "spec": {"configs": [{"name": "Name", "labels": ["Labels"], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "type": "manual", "external_resources_account": "ExternalResourcesAccount"}], "dashboard": {"widgets": [{"name": "Name"}]}}}`)
				}))
			})
			It(`Invoke UpdateProject successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Metadata = metadataModel
				updateProjectOptionsModel.Spec = specModel
				updateProjectOptionsModel.Prefer = core.StringPtr("return=minimal")
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.UpdateProjectWithContext(ctx, updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.UpdateProjectWithContext(ctx, updateProjectOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectPath))
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

					Expect(req.Header["Prefer"]).ToNot(BeNil())
					Expect(req.Header["Prefer"][0]).To(Equal(fmt.Sprintf("%v", "return=minimal")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"href": "Href", "metadata": {"name": "Name", "description": "Description", "id": "ID", "created_by": "CreatedBy", "created_at": "CreatedAt", "updated_at": "UpdatedAt", "tags": ["Tags"], "repo_url": "RepoURL"}, "spec": {"configs": [{"name": "Name", "labels": ["Labels"], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "type": "manual", "external_resources_account": "ExternalResourcesAccount"}], "dashboard": {"widgets": [{"name": "Name"}]}}}`)
				}))
			})
			It(`Invoke UpdateProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.UpdateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Metadata = metadataModel
				updateProjectOptionsModel.Spec = specModel
				updateProjectOptionsModel.Prefer = core.StringPtr("return=minimal")
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProject with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Metadata = metadataModel
				updateProjectOptionsModel.Spec = specModel
				updateProjectOptionsModel.Prefer = core.StringPtr("return=minimal")
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProjectOptions model with no property values
				updateProjectOptionsModelNew := new(projectsv1.UpdateProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.UpdateProject(updateProjectOptionsModelNew)
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
			It(`Invoke UpdateProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Metadata = metadataModel
				updateProjectOptionsModel.Spec = specModel
				updateProjectOptionsModel.Prefer = core.StringPtr("return=minimal")
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.UpdateProject(updateProjectOptionsModel)
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
	Describe(`InstallProject(installProjectOptions *InstallProjectOptions)`, func() {
		installProjectPath := "/v1/projects/testString/install"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(installProjectPath))
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
			It(`Invoke InstallProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.InstallProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the InstallProjectOptions model
				installProjectOptionsModel := new(projectsv1.InstallProjectOptions)
				installProjectOptionsModel.ID = core.StringPtr("testString")
				installProjectOptionsModel.ConfigNames = []string{"testString"}
				installProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.InstallProject(installProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke InstallProject with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InstallProjectOptions model
				installProjectOptionsModel := new(projectsv1.InstallProjectOptions)
				installProjectOptionsModel.ID = core.StringPtr("testString")
				installProjectOptionsModel.ConfigNames = []string{"testString"}
				installProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.InstallProject(installProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the InstallProjectOptions model with no property values
				installProjectOptionsModelNew := new(projectsv1.InstallProjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.InstallProject(installProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UninstallProject(uninstallProjectOptions *UninstallProjectOptions)`, func() {
		uninstallProjectPath := "/v1/projects/testString/uninstall"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uninstallProjectPath))
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
			It(`Invoke UninstallProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.UninstallProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UninstallProjectOptions model
				uninstallProjectOptionsModel := new(projectsv1.UninstallProjectOptions)
				uninstallProjectOptionsModel.ID = core.StringPtr("testString")
				uninstallProjectOptionsModel.ConfigNames = []string{"testString"}
				uninstallProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.UninstallProject(uninstallProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UninstallProject with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the UninstallProjectOptions model
				uninstallProjectOptionsModel := new(projectsv1.UninstallProjectOptions)
				uninstallProjectOptionsModel.ID = core.StringPtr("testString")
				uninstallProjectOptionsModel.ConfigNames = []string{"testString"}
				uninstallProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.UninstallProject(uninstallProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UninstallProjectOptions model with no property values
				uninstallProjectOptionsModelNew := new(projectsv1.UninstallProjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.UninstallProject(uninstallProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CheckProject(checkProjectOptions *CheckProjectOptions)`, func() {
		checkProjectPath := "/v1/projects/testString/check"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkProjectPath))
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
			It(`Invoke CheckProject successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.CheckProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CheckProjectOptions model
				checkProjectOptionsModel := new(projectsv1.CheckProjectOptions)
				checkProjectOptionsModel.ID = core.StringPtr("testString")
				checkProjectOptionsModel.ConfigNames = []string{"testString"}
				checkProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.CheckProject(checkProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CheckProject with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the CheckProjectOptions model
				checkProjectOptionsModel := new(projectsv1.CheckProjectOptions)
				checkProjectOptionsModel.ID = core.StringPtr("testString")
				checkProjectOptionsModel.ConfigNames = []string{"testString"}
				checkProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.CheckProject(checkProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CheckProjectOptions model with no property values
				checkProjectOptionsModelNew := new(projectsv1.CheckProjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.CheckProject(checkProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectStatus(getProjectStatusOptions *GetProjectStatusOptions) - Operation response error`, func() {
		getProjectStatusPath := "/v1/projects/testString/status"
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
				getProjectStatusOptionsModel.ID = core.StringPtr("testString")
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
		getProjectStatusPath := "/v1/projects/testString/status"
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "name": "Name", "status": "installed", "message": "Message", "pipeline_run": "PipelineRun", "schematics": "Schematics", "computed_statuses": {"mapKey": "anyValue"}, "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}`)
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
				getProjectStatusOptionsModel.ID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "name": "Name", "status": "installed", "message": "Message", "pipeline_run": "PipelineRun", "schematics": "Schematics", "computed_statuses": {"mapKey": "anyValue"}, "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}`)
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
				getProjectStatusOptionsModel.ID = core.StringPtr("testString")
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
				getProjectStatusOptionsModel.ID = core.StringPtr("testString")
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
				getProjectStatusOptionsModel.ID = core.StringPtr("testString")
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
	Describe(`UpdateProjectComputedStatus(updateProjectComputedStatusOptions *UpdateProjectComputedStatusOptions)`, func() {
		updateProjectComputedStatusPath := "/v1/projects/testString/computed_status/cost"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectComputedStatusPath))
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

					res.WriteHeader(204)
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
				response, operationErr := projectsService.UpdateProjectComputedStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateProjectComputedStatusOptions model
				updateProjectComputedStatusOptionsModel := new(projectsv1.UpdateProjectComputedStatusOptions)
				updateProjectComputedStatusOptionsModel.ID = core.StringPtr("testString")
				updateProjectComputedStatusOptionsModel.StatusName = core.StringPtr("cost")
				updateProjectComputedStatusOptionsModel.RequestBody = make(map[string]interface{})
				updateProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.UpdateProjectComputedStatus(updateProjectComputedStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
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
				updateProjectComputedStatusOptionsModel.ID = core.StringPtr("testString")
				updateProjectComputedStatusOptionsModel.StatusName = core.StringPtr("cost")
				updateProjectComputedStatusOptionsModel.RequestBody = make(map[string]interface{})
				updateProjectComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.UpdateProjectComputedStatus(updateProjectComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateProjectComputedStatusOptions model with no property values
				updateProjectComputedStatusOptionsModelNew := new(projectsv1.UpdateProjectComputedStatusOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.UpdateProjectComputedStatus(updateProjectComputedStatusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectConfigStatus(getProjectConfigStatusOptions *GetProjectConfigStatusOptions) - Operation response error`, func() {
		getProjectConfigStatusPath := "/v1/projects/testString/configs/testString/status"
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
				getProjectConfigStatusOptionsModel.ID = core.StringPtr("testString")
				getProjectConfigStatusOptionsModel.Name = core.StringPtr("testString")
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
		getProjectConfigStatusPath := "/v1/projects/testString/configs/testString/status"
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "metadata": {"name": "Name", "description": "Description", "id": "ID", "created_by": "CreatedBy", "created_at": "CreatedAt", "updated_at": "UpdatedAt", "tags": ["Tags"], "repo_url": "RepoURL"}, "status": {"configs": [{"name": "Name", "status": "installed", "message": "Message", "pipeline_run": "PipelineRun", "schematics": "Schematics", "computed_statuses": {"mapKey": "anyValue"}, "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}], "computed_statuses": {"mapKey": "anyValue"}, "services": {"toolchain": {"id": "ID", "status": "installed", "message": "Message", "schematics": "Schematics"}, "schematics": {"id": "ID", "status": "installed", "message": "Message", "schematics": "Schematics"}, "git_repo": {"id": "ID", "status": "installed", "message": "Message", "schematics": "Schematics"}}}}`)
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
				getProjectConfigStatusOptionsModel.ID = core.StringPtr("testString")
				getProjectConfigStatusOptionsModel.Name = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "metadata": {"name": "Name", "description": "Description", "id": "ID", "created_by": "CreatedBy", "created_at": "CreatedAt", "updated_at": "UpdatedAt", "tags": ["Tags"], "repo_url": "RepoURL"}, "status": {"configs": [{"name": "Name", "status": "installed", "message": "Message", "pipeline_run": "PipelineRun", "schematics": "Schematics", "computed_statuses": {"mapKey": "anyValue"}, "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}], "computed_statuses": {"mapKey": "anyValue"}, "services": {"toolchain": {"id": "ID", "status": "installed", "message": "Message", "schematics": "Schematics"}, "schematics": {"id": "ID", "status": "installed", "message": "Message", "schematics": "Schematics"}, "git_repo": {"id": "ID", "status": "installed", "message": "Message", "schematics": "Schematics"}}}}`)
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
				getProjectConfigStatusOptionsModel.ID = core.StringPtr("testString")
				getProjectConfigStatusOptionsModel.Name = core.StringPtr("testString")
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
				getProjectConfigStatusOptionsModel.ID = core.StringPtr("testString")
				getProjectConfigStatusOptionsModel.Name = core.StringPtr("testString")
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
				getProjectConfigStatusOptionsModel.ID = core.StringPtr("testString")
				getProjectConfigStatusOptionsModel.Name = core.StringPtr("testString")
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
	Describe(`UpdateProjectConfigStatus(updateProjectConfigStatusOptions *UpdateProjectConfigStatusOptions)`, func() {
		updateProjectConfigStatusPath := "/v1/projects/testString/configs/testString/status"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectConfigStatusPath))
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

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateProjectConfigStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.UpdateProjectConfigStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("bla bla")
				outputValueModel.Description = core.StringPtr("Bla bla description")
				outputValueModel.Value = core.StringPtr("0")

				// Construct an instance of the UpdateProjectConfigStatusOptions model
				updateProjectConfigStatusOptionsModel := new(projectsv1.UpdateProjectConfigStatusOptions)
				updateProjectConfigStatusOptionsModel.ID = core.StringPtr("testString")
				updateProjectConfigStatusOptionsModel.Name = core.StringPtr("testString")
				updateProjectConfigStatusOptionsModel.NewName = core.StringPtr("Example configuration")
				updateProjectConfigStatusOptionsModel.NewStatus = core.StringPtr("installed")
				updateProjectConfigStatusOptionsModel.NewMessage = core.StringPtr("Ok success")
				updateProjectConfigStatusOptionsModel.NewPipelineRun = core.StringPtr("https://url.to.somewhere")
				updateProjectConfigStatusOptionsModel.NewSchematics = core.StringPtr("e2cpmpm8ex4jpqt0o5-q5k6rkhuca6f42p14iwjg:lohc.xnr8fxhrftrl.io2:wnh36pj-kyf0mwmy:r4g.kho/qo:2")
				updateProjectConfigStatusOptionsModel.NewComputedStatuses = make(map[string]interface{})
				updateProjectConfigStatusOptionsModel.NewOutput = []projectsv1.OutputValue{*outputValueModel}
				updateProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.UpdateProjectConfigStatus(updateProjectConfigStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateProjectConfigStatus with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("bla bla")
				outputValueModel.Description = core.StringPtr("Bla bla description")
				outputValueModel.Value = core.StringPtr("0")

				// Construct an instance of the UpdateProjectConfigStatusOptions model
				updateProjectConfigStatusOptionsModel := new(projectsv1.UpdateProjectConfigStatusOptions)
				updateProjectConfigStatusOptionsModel.ID = core.StringPtr("testString")
				updateProjectConfigStatusOptionsModel.Name = core.StringPtr("testString")
				updateProjectConfigStatusOptionsModel.NewName = core.StringPtr("Example configuration")
				updateProjectConfigStatusOptionsModel.NewStatus = core.StringPtr("installed")
				updateProjectConfigStatusOptionsModel.NewMessage = core.StringPtr("Ok success")
				updateProjectConfigStatusOptionsModel.NewPipelineRun = core.StringPtr("https://url.to.somewhere")
				updateProjectConfigStatusOptionsModel.NewSchematics = core.StringPtr("e2cpmpm8ex4jpqt0o5-q5k6rkhuca6f42p14iwjg:lohc.xnr8fxhrftrl.io2:wnh36pj-kyf0mwmy:r4g.kho/qo:2")
				updateProjectConfigStatusOptionsModel.NewComputedStatuses = make(map[string]interface{})
				updateProjectConfigStatusOptionsModel.NewOutput = []projectsv1.OutputValue{*outputValueModel}
				updateProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.UpdateProjectConfigStatus(updateProjectConfigStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateProjectConfigStatusOptions model with no property values
				updateProjectConfigStatusOptionsModelNew := new(projectsv1.UpdateProjectConfigStatusOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.UpdateProjectConfigStatus(updateProjectConfigStatusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptions *UpdateProjectConfigComputedStatusOptions)`, func() {
		updateProjectConfigComputedStatusPath := "/v1/projects/testString/configs/testString/computed_status/cost"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectConfigComputedStatusPath))
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

					res.WriteHeader(204)
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
				response, operationErr := projectsService.UpdateProjectConfigComputedStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateProjectConfigComputedStatusOptions model
				updateProjectConfigComputedStatusOptionsModel := new(projectsv1.UpdateProjectConfigComputedStatusOptions)
				updateProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("testString")
				updateProjectConfigComputedStatusOptionsModel.Name = core.StringPtr("testString")
				updateProjectConfigComputedStatusOptionsModel.StatusName = core.StringPtr("cost")
				updateProjectConfigComputedStatusOptionsModel.RequestBody = make(map[string]interface{})
				updateProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
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
				updateProjectConfigComputedStatusOptionsModel.ID = core.StringPtr("testString")
				updateProjectConfigComputedStatusOptionsModel.Name = core.StringPtr("testString")
				updateProjectConfigComputedStatusOptionsModel.StatusName = core.StringPtr("cost")
				updateProjectConfigComputedStatusOptionsModel.RequestBody = make(map[string]interface{})
				updateProjectConfigComputedStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateProjectConfigComputedStatusOptions model with no property values
				updateProjectConfigComputedStatusOptionsModelNew := new(projectsv1.UpdateProjectConfigComputedStatusOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptionsModelNew)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			projectsService, _ := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
				URL:           "http://projectsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCheckProjectOptions successfully`, func() {
				// Construct an instance of the CheckProjectOptions model
				id := "testString"
				checkProjectOptionsModel := projectsService.NewCheckProjectOptions(id)
				checkProjectOptionsModel.SetID("testString")
				checkProjectOptionsModel.SetConfigNames([]string{"testString"})
				checkProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(checkProjectOptionsModel).ToNot(BeNil())
				Expect(checkProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(checkProjectOptionsModel.ConfigNames).To(Equal([]string{"testString"}))
				Expect(checkProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectOptions successfully`, func() {
				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				Expect(metadataModel).ToNot(BeNil())
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")
				Expect(metadataModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(metadataModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(metadataModel.ID).To(Equal(core.StringPtr("unique-id-of-the-project")))
				Expect(metadataModel.CreatedBy).To(Equal(core.StringPtr("creator@acme.com")))
				Expect(metadataModel.CreatedAt).To(Equal(core.StringPtr("testString")))
				Expect(metadataModel.UpdatedAt).To(Equal(core.StringPtr("testString")))
				Expect(metadataModel.Tags).To(Equal([]string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}))
				Expect(metadataModel.RepoURL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal(core.StringPtr(`["project:ghost","type:infrastructure"]`)))

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				Expect(configModel).ToNot(BeNil())
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				Expect(configModel.Name).To(Equal(core.StringPtr("common-variables")))
				Expect(configModel.Labels).To(Equal([]string{"testString"}))
				Expect(configModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(configModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(configModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				Expect(widgetModel).ToNot(BeNil())
				widgetModel.Name = core.StringPtr("project-properties")
				Expect(widgetModel.Name).To(Equal(core.StringPtr("project-properties")))

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				Expect(specDashboardModel).ToNot(BeNil())
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}
				Expect(specDashboardModel.Widgets).To(Equal([]projectsv1.Widget{*widgetModel}))

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				Expect(specModel).ToNot(BeNil())
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel
				Expect(specModel.Configs).To(Equal([]projectsv1.ConfigIntf{configModel}))
				Expect(specModel.Dashboard).To(Equal(specDashboardModel))

				// Construct an instance of the CreateProjectOptions model
				var createProjectOptionsMetadata *projectsv1.Metadata = nil
				var createProjectOptionsSpec *projectsv1.Spec = nil
				createProjectOptionsModel := projectsService.NewCreateProjectOptions(createProjectOptionsMetadata, createProjectOptionsSpec)
				createProjectOptionsModel.SetMetadata(metadataModel)
				createProjectOptionsModel.SetSpec(specModel)
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.Metadata).To(Equal(metadataModel))
				Expect(createProjectOptionsModel.Spec).To(Equal(specModel))
				Expect(createProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectOptions successfully`, func() {
				// Construct an instance of the DeleteProjectOptions model
				id := "testString"
				deleteProjectOptionsModel := projectsService.NewDeleteProjectOptions(id)
				deleteProjectOptionsModel.SetID("testString")
				deleteProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectOptionsModel).ToNot(BeNil())
				Expect(deleteProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetProjectConfigStatusOptions successfully`, func() {
				// Construct an instance of the GetProjectConfigStatusOptions model
				id := "testString"
				name := "testString"
				getProjectConfigStatusOptionsModel := projectsService.NewGetProjectConfigStatusOptions(id, name)
				getProjectConfigStatusOptionsModel.SetID("testString")
				getProjectConfigStatusOptionsModel.SetName("testString")
				getProjectConfigStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectConfigStatusOptionsModel).ToNot(BeNil())
				Expect(getProjectConfigStatusOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProjectConfigStatusOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(getProjectConfigStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectOptions successfully`, func() {
				// Construct an instance of the GetProjectOptions model
				id := "testString"
				getProjectOptionsModel := projectsService.NewGetProjectOptions(id)
				getProjectOptionsModel.SetID("testString")
				getProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectOptionsModel).ToNot(BeNil())
				Expect(getProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectStatusOptions successfully`, func() {
				// Construct an instance of the GetProjectStatusOptions model
				id := "testString"
				getProjectStatusOptionsModel := projectsService.NewGetProjectStatusOptions(id)
				getProjectStatusOptionsModel.SetID("testString")
				getProjectStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectStatusOptionsModel).ToNot(BeNil())
				Expect(getProjectStatusOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProjectStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInstallProjectOptions successfully`, func() {
				// Construct an instance of the InstallProjectOptions model
				id := "testString"
				installProjectOptionsModel := projectsService.NewInstallProjectOptions(id)
				installProjectOptionsModel.SetID("testString")
				installProjectOptionsModel.SetConfigNames([]string{"testString"})
				installProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(installProjectOptionsModel).ToNot(BeNil())
				Expect(installProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(installProjectOptionsModel.ConfigNames).To(Equal([]string{"testString"}))
				Expect(installProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectsOptions successfully`, func() {
				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := projectsService.NewListProjectsOptions()
				listProjectsOptionsModel.SetOffset(int64(0))
				listProjectsOptionsModel.SetLimit(int64(1))
				listProjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectsOptionsModel).ToNot(BeNil())
				Expect(listProjectsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listProjectsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listProjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSpecDashboard successfully`, func() {
				widgets := []projectsv1.Widget{}
				_model, err := projectsService.NewSpecDashboard(widgets)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUninstallProjectOptions successfully`, func() {
				// Construct an instance of the UninstallProjectOptions model
				id := "testString"
				uninstallProjectOptionsModel := projectsService.NewUninstallProjectOptions(id)
				uninstallProjectOptionsModel.SetID("testString")
				uninstallProjectOptionsModel.SetConfigNames([]string{"testString"})
				uninstallProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uninstallProjectOptionsModel).ToNot(BeNil())
				Expect(uninstallProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(uninstallProjectOptionsModel.ConfigNames).To(Equal([]string{"testString"}))
				Expect(uninstallProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectComputedStatusOptions successfully`, func() {
				// Construct an instance of the UpdateProjectComputedStatusOptions model
				id := "testString"
				statusName := "cost"
				requestBody := make(map[string]interface{})
				updateProjectComputedStatusOptionsModel := projectsService.NewUpdateProjectComputedStatusOptions(id, statusName, requestBody)
				updateProjectComputedStatusOptionsModel.SetID("testString")
				updateProjectComputedStatusOptionsModel.SetStatusName("cost")
				updateProjectComputedStatusOptionsModel.SetRequestBody(make(map[string]interface{}))
				updateProjectComputedStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectComputedStatusOptionsModel).ToNot(BeNil())
				Expect(updateProjectComputedStatusOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectComputedStatusOptionsModel.StatusName).To(Equal(core.StringPtr("cost")))
				Expect(updateProjectComputedStatusOptionsModel.RequestBody).To(Equal(make(map[string]interface{})))
				Expect(updateProjectComputedStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectConfigComputedStatusOptions successfully`, func() {
				// Construct an instance of the UpdateProjectConfigComputedStatusOptions model
				id := "testString"
				name := "testString"
				statusName := "cost"
				requestBody := make(map[string]interface{})
				updateProjectConfigComputedStatusOptionsModel := projectsService.NewUpdateProjectConfigComputedStatusOptions(id, name, statusName, requestBody)
				updateProjectConfigComputedStatusOptionsModel.SetID("testString")
				updateProjectConfigComputedStatusOptionsModel.SetName("testString")
				updateProjectConfigComputedStatusOptionsModel.SetStatusName("cost")
				updateProjectConfigComputedStatusOptionsModel.SetRequestBody(make(map[string]interface{}))
				updateProjectConfigComputedStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectConfigComputedStatusOptionsModel).ToNot(BeNil())
				Expect(updateProjectConfigComputedStatusOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectConfigComputedStatusOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectConfigComputedStatusOptionsModel.StatusName).To(Equal(core.StringPtr("cost")))
				Expect(updateProjectConfigComputedStatusOptionsModel.RequestBody).To(Equal(make(map[string]interface{})))
				Expect(updateProjectConfigComputedStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectConfigStatusOptions successfully`, func() {
				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("bla bla")
				outputValueModel.Description = core.StringPtr("Bla bla description")
				outputValueModel.Value = core.StringPtr("0")
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("bla bla")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("Bla bla description")))
				Expect(outputValueModel.Value).To(Equal(core.StringPtr("0")))

				// Construct an instance of the UpdateProjectConfigStatusOptions model
				id := "testString"
				name := "testString"
				updateProjectConfigStatusOptionsNewName := "Example configuration"
				updateProjectConfigStatusOptionsNewStatus := "installed"
				updateProjectConfigStatusOptionsNewMessage := "Ok success"
				updateProjectConfigStatusOptionsModel := projectsService.NewUpdateProjectConfigStatusOptions(id, name, updateProjectConfigStatusOptionsNewName, updateProjectConfigStatusOptionsNewStatus, updateProjectConfigStatusOptionsNewMessage)
				updateProjectConfigStatusOptionsModel.SetID("testString")
				updateProjectConfigStatusOptionsModel.SetName("testString")
				updateProjectConfigStatusOptionsModel.SetNewName("Example configuration")
				updateProjectConfigStatusOptionsModel.SetNewStatus("installed")
				updateProjectConfigStatusOptionsModel.SetNewMessage("Ok success")
				updateProjectConfigStatusOptionsModel.SetNewPipelineRun("https://url.to.somewhere")
				updateProjectConfigStatusOptionsModel.SetNewSchematics("e2cpmpm8ex4jpqt0o5-q5k6rkhuca6f42p14iwjg:lohc.xnr8fxhrftrl.io2:wnh36pj-kyf0mwmy:r4g.kho/qo:2")
				updateProjectConfigStatusOptionsModel.SetNewComputedStatuses(make(map[string]interface{}))
				updateProjectConfigStatusOptionsModel.SetNewOutput([]projectsv1.OutputValue{*outputValueModel})
				updateProjectConfigStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectConfigStatusOptionsModel).ToNot(BeNil())
				Expect(updateProjectConfigStatusOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectConfigStatusOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectConfigStatusOptionsModel.NewName).To(Equal(core.StringPtr("Example configuration")))
				Expect(updateProjectConfigStatusOptionsModel.NewStatus).To(Equal(core.StringPtr("installed")))
				Expect(updateProjectConfigStatusOptionsModel.NewMessage).To(Equal(core.StringPtr("Ok success")))
				Expect(updateProjectConfigStatusOptionsModel.NewPipelineRun).To(Equal(core.StringPtr("https://url.to.somewhere")))
				Expect(updateProjectConfigStatusOptionsModel.NewSchematics).To(Equal(core.StringPtr("e2cpmpm8ex4jpqt0o5-q5k6rkhuca6f42p14iwjg:lohc.xnr8fxhrftrl.io2:wnh36pj-kyf0mwmy:r4g.kho/qo:2")))
				Expect(updateProjectConfigStatusOptionsModel.NewComputedStatuses).To(Equal(make(map[string]interface{})))
				Expect(updateProjectConfigStatusOptionsModel.NewOutput).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(updateProjectConfigStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectOptions successfully`, func() {
				// Construct an instance of the Metadata model
				metadataModel := new(projectsv1.Metadata)
				Expect(metadataModel).ToNot(BeNil())
				metadataModel.Name = core.StringPtr("acme-microservice")
				metadataModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				metadataModel.ID = core.StringPtr("unique-id-of-the-project")
				metadataModel.CreatedBy = core.StringPtr("creator@acme.com")
				metadataModel.CreatedAt = core.StringPtr("testString")
				metadataModel.UpdatedAt = core.StringPtr("testString")
				metadataModel.Tags = []string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}
				metadataModel.RepoURL = core.StringPtr("testString")
				Expect(metadataModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(metadataModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(metadataModel.ID).To(Equal(core.StringPtr("unique-id-of-the-project")))
				Expect(metadataModel.CreatedBy).To(Equal(core.StringPtr("creator@acme.com")))
				Expect(metadataModel.CreatedAt).To(Equal(core.StringPtr("testString")))
				Expect(metadataModel.UpdatedAt).To(Equal(core.StringPtr("testString")))
				Expect(metadataModel.Tags).To(Equal([]string{"org:/bu-xyz/application/acme/", "billing:/bu-xyz"}))
				Expect(metadataModel.RepoURL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = core.StringPtr(`["project:ghost","type:infrastructure"]`)
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal(core.StringPtr(`["project:ghost","type:infrastructure"]`)))

				// Construct an instance of the ConfigManualProperty model
				configModel := new(projectsv1.ConfigManualProperty)
				Expect(configModel).ToNot(BeNil())
				configModel.Name = core.StringPtr("common-variables")
				configModel.Labels = []string{"testString"}
				configModel.Output = []projectsv1.OutputValue{*outputValueModel}
				configModel.Type = core.StringPtr("manual")
				configModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				Expect(configModel.Name).To(Equal(core.StringPtr("common-variables")))
				Expect(configModel.Labels).To(Equal([]string{"testString"}))
				Expect(configModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(configModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(configModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				Expect(widgetModel).ToNot(BeNil())
				widgetModel.Name = core.StringPtr("project-properties")
				Expect(widgetModel.Name).To(Equal(core.StringPtr("project-properties")))

				// Construct an instance of the SpecDashboard model
				specDashboardModel := new(projectsv1.SpecDashboard)
				Expect(specDashboardModel).ToNot(BeNil())
				specDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}
				Expect(specDashboardModel.Widgets).To(Equal([]projectsv1.Widget{*widgetModel}))

				// Construct an instance of the Spec model
				specModel := new(projectsv1.Spec)
				Expect(specModel).ToNot(BeNil())
				specModel.Configs = []projectsv1.ConfigIntf{configModel}
				specModel.Dashboard = specDashboardModel
				Expect(specModel.Configs).To(Equal([]projectsv1.ConfigIntf{configModel}))
				Expect(specModel.Dashboard).To(Equal(specDashboardModel))

				// Construct an instance of the UpdateProjectOptions model
				id := "testString"
				var updateProjectOptionsMetadata *projectsv1.Metadata = nil
				var updateProjectOptionsSpec *projectsv1.Spec = nil
				updateProjectOptionsModel := projectsService.NewUpdateProjectOptions(id, updateProjectOptionsMetadata, updateProjectOptionsSpec)
				updateProjectOptionsModel.SetID("testString")
				updateProjectOptionsModel.SetMetadata(metadataModel)
				updateProjectOptionsModel.SetSpec(specModel)
				updateProjectOptionsModel.SetPrefer("return=minimal")
				updateProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectOptionsModel).ToNot(BeNil())
				Expect(updateProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectOptionsModel.Metadata).To(Equal(metadataModel))
				Expect(updateProjectOptionsModel.Spec).To(Equal(specModel))
				Expect(updateProjectOptionsModel.Prefer).To(Equal(core.StringPtr("return=minimal")))
				Expect(updateProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewConfigStatus successfully`, func() {
				name := "testString"
				status := "installed"
				message := "testString"
				_model, err := projectsService.NewConfigStatus(name, status, message)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewInputVariable successfully`, func() {
				name := "testString"
				_model, err := projectsService.NewInputVariable(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMetadata successfully`, func() {
				name := "testString"
				_model, err := projectsService.NewMetadata(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewOutputValue successfully`, func() {
				name := "testString"
				_model, err := projectsService.NewOutputValue(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSchematicsBlueprint successfully`, func() {
				repoURL := "testString"
				definitionFile := "testString"
				_model, err := projectsService.NewSchematicsBlueprint(repoURL, definitionFile)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSpec successfully`, func() {
				configs := []projectsv1.ConfigIntf{}
				_model, err := projectsService.NewSpec(configs)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTerraformTemplate successfully`, func() {
				repoURL := "testString"
				terraformVersion := "testString"
				_model, err := projectsService.NewTerraformTemplate(repoURL, terraformVersion)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewWidget successfully`, func() {
				name := "testString"
				_model, err := projectsService.NewWidget(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigManualProperty successfully`, func() {
				name := "testString"
				typeVar := "manual"
				_model, err := projectsService.NewConfigManualProperty(name, typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigSchematicsBlueprintProperty successfully`, func() {
				name := "testString"
				typeVar := "schematics_blueprint"
				input := []projectsv1.InputVariable{}
				var blueprint *projectsv1.SchematicsBlueprint = nil
				_, err := projectsService.NewConfigSchematicsBlueprintProperty(name, typeVar, input, blueprint)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigTerraformTemplateProperty successfully`, func() {
				name := "testString"
				typeVar := "terraform_template"
				input := []projectsv1.InputVariable{}
				var blueprint *projectsv1.TerraformTemplate = nil
				_, err := projectsService.NewConfigTerraformTemplateProperty(name, typeVar, input, blueprint)
				Expect(err).ToNot(BeNil())
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
