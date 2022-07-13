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
					Expect(req.Header["X-Iam-Api"]).ToNot(BeNil())
					Expect(req.Header["X-Iam-Api"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"Default"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"us-south"}))
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				createProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
				createProjectOptionsModel.XIamApi = core.StringPtr("testString")
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

					Expect(req.Header["X-Iam-Api"]).ToNot(BeNil())
					Expect(req.Header["X-Iam-Api"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"Default"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"us-south"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "crn": "Crn", "location": "us-south", "resource_group": "ResourceGroup"}`)
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				createProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
				createProjectOptionsModel.XIamApi = core.StringPtr("testString")
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

					Expect(req.Header["X-Iam-Api"]).ToNot(BeNil())
					Expect(req.Header["X-Iam-Api"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"Default"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"us-south"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "crn": "Crn", "location": "us-south", "resource_group": "ResourceGroup"}`)
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				createProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
				createProjectOptionsModel.XIamApi = core.StringPtr("testString")
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				createProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
				createProjectOptionsModel.XIamApi = core.StringPtr("testString")
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				createProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
				createProjectOptionsModel.XIamApi = core.StringPtr("testString")
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
				listProjectsOptionsModel.Start = core.StringPtr("testString")
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

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 1, "total_count": 0, "first": {"href": "Href", "start": "Start"}, "last": {"href": "Href", "start": "Start"}, "previous": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "projects": [{"name": "Name", "id": "ID", "crn": "Crn", "location": "us-south", "resource_group": "ResourceGroup", "state": "CREATING"}]}`)
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

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 1, "total_count": 0, "first": {"href": "Href", "start": "Start"}, "last": {"href": "Href", "start": "Start"}, "previous": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "projects": [{"name": "Name", "id": "ID", "crn": "Crn", "location": "us-south", "resource_group": "ResourceGroup", "state": "CREATING"}]}`)
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
				listProjectsOptionsModel.Start = core.StringPtr("testString")
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
				listProjectsOptionsModel.Start = core.StringPtr("testString")
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
		It(`Invoke GetNextStart successfully`, func() {
			responseObject := new(projectsv1.ListProjectsResponse)
			nextObject := new(projectsv1.PaginationLink)
			nextObject.Start = core.StringPtr("abc-123")
			responseObject.Next = nextObject

			value, err := responseObject.GetNextStart()
			Expect(err).To(BeNil())
			Expect(value).To(Equal(core.StringPtr("abc-123")))
		})
		It(`Invoke GetNextStart without a "Next" property in the response`, func() {
			responseObject := new(projectsv1.ListProjectsResponse)

			value, err := responseObject.GetNextStart()
			Expect(err).To(BeNil())
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "configs": [{"name": "Name", "labels": ["Labels"], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "type": "manual", "external_resources_account": "ExternalResourcesAccount"}], "dashboard": {"widgets": [{"name": "Name"}]}}`)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "configs": [{"name": "Name", "labels": ["Labels"], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "type": "manual", "external_resources_account": "ExternalResourcesAccount"}], "dashboard": {"widgets": [{"name": "Name"}]}}`)
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				updateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				updateProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				updateProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"branch": "Branch", "url": "URL"}`)
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				updateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				updateProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				updateProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"branch": "Branch", "url": "URL"}`)
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				updateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				updateProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				updateProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				updateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				updateProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				updateProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				updateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				updateProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				updateProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the MergeProjectOptions model
				mergeProjectOptionsModel := new(projectsv1.MergeProjectOptions)
				mergeProjectOptionsModel.ID = core.StringPtr("testString")
				mergeProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				mergeProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				mergeProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "crn": "Crn", "location": "us-south", "resource_group": "ResourceGroup"}`)
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the MergeProjectOptions model
				mergeProjectOptionsModel := new(projectsv1.MergeProjectOptions)
				mergeProjectOptionsModel.ID = core.StringPtr("testString")
				mergeProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				mergeProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				mergeProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "crn": "Crn", "location": "us-south", "resource_group": "ResourceGroup"}`)
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the MergeProjectOptions model
				mergeProjectOptionsModel := new(projectsv1.MergeProjectOptions)
				mergeProjectOptionsModel.ID = core.StringPtr("testString")
				mergeProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				mergeProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				mergeProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the MergeProjectOptions model
				mergeProjectOptionsModel := new(projectsv1.MergeProjectOptions)
				mergeProjectOptionsModel.ID = core.StringPtr("testString")
				mergeProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				mergeProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				mergeProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the MergeProjectOptions model
				mergeProjectOptionsModel := new(projectsv1.MergeProjectOptions)
				mergeProjectOptionsModel.ID = core.StringPtr("testString")
				mergeProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				mergeProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				mergeProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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

					res.WriteHeader(200)
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the ValidateProjectOptions model
				validateProjectOptionsModel := new(projectsv1.ValidateProjectOptions)
				validateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				validateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				validateProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				validateProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				widgetModel.Name = core.StringPtr("project-properties")

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}

				// Construct an instance of the ValidateProjectOptions model
				validateProjectOptionsModel := new(projectsv1.ValidateProjectOptions)
				validateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				validateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				validateProjectOptionsModel.Configs = []projectsv1.ProjectConfigIntf{projectConfigModel}
				validateProjectOptionsModel.Dashboard = projectPrototypeDashboardModel
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
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "project_crn": "ProjectCrn", "project_name": "ProjectName", "location": "us-south", "resource_group": "ResourceGroup", "state": "CREATING", "git_repo": {"url": "URL", "branch": "Branch", "project_id": "ProjectID"}, "toolchain": {"crn": "Crn", "guid": "Guid"}, "schematics": {"cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}, "credentials": {"api_key_ref": "ApiKeyRef"}, "configs": [{"name": "Name", "labels": ["Labels"], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "type": "manual", "external_resources_account": "ExternalResourcesAccount"}], "dashboard": {"widgets": [{"name": "Name"}]}, "computed_statuses": {"mapKey": "anyValue"}, "active_prs": [{"branch": "Branch", "url": "URL"}], "history": [{"timestamp": "2019-01-01T12:00:00.000Z", "code": "Code", "type": "git_repo"}]}`)
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
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "project_crn": "ProjectCrn", "project_name": "ProjectName", "location": "us-south", "resource_group": "ResourceGroup", "state": "CREATING", "git_repo": {"url": "URL", "branch": "Branch", "project_id": "ProjectID"}, "toolchain": {"crn": "Crn", "guid": "Guid"}, "schematics": {"cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}, "credentials": {"api_key_ref": "ApiKeyRef"}, "configs": [{"name": "Name", "labels": ["Labels"], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "type": "manual", "external_resources_account": "ExternalResourcesAccount"}], "dashboard": {"widgets": [{"name": "Name"}]}, "computed_statuses": {"mapKey": "anyValue"}, "active_prs": [{"branch": "Branch", "url": "URL"}], "history": [{"timestamp": "2019-01-01T12:00:00.000Z", "code": "Code", "type": "git_repo"}]}`)
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
				serviceInfoToolchainModel.Crn = core.StringPtr("testString")
				serviceInfoToolchainModel.Guid = core.StringPtr("testString")

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
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "project_crn": "ProjectCrn", "project_name": "ProjectName", "location": "us-south", "resource_group": "ResourceGroup", "state": "CREATING", "git_repo": {"url": "URL", "branch": "Branch", "project_id": "ProjectID"}, "toolchain": {"crn": "Crn", "guid": "Guid"}, "schematics": {"cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}, "credentials": {"api_key_ref": "ApiKeyRef"}, "configs": [{"name": "Name", "labels": ["Labels"], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "type": "manual", "external_resources_account": "ExternalResourcesAccount"}], "dashboard": {"widgets": [{"name": "Name"}]}, "computed_statuses": {"mapKey": "anyValue"}, "active_prs": [{"branch": "Branch", "url": "URL"}], "history": [{"timestamp": "2019-01-01T12:00:00.000Z", "code": "Code", "type": "git_repo"}]}`)
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
				serviceInfoToolchainModel.Crn = core.StringPtr("testString")
				serviceInfoToolchainModel.Guid = core.StringPtr("testString")

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
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "project_crn": "ProjectCrn", "project_name": "ProjectName", "location": "us-south", "resource_group": "ResourceGroup", "state": "CREATING", "git_repo": {"url": "URL", "branch": "Branch", "project_id": "ProjectID"}, "toolchain": {"crn": "Crn", "guid": "Guid"}, "schematics": {"cart_order_id": "CartOrderID", "workspace_id": "WorkspaceID", "cart_item_name": "CartItemName"}, "credentials": {"api_key_ref": "ApiKeyRef"}, "configs": [{"name": "Name", "labels": ["Labels"], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "type": "manual", "external_resources_account": "ExternalResourcesAccount"}], "dashboard": {"widgets": [{"name": "Name"}]}, "computed_statuses": {"mapKey": "anyValue"}, "active_prs": [{"branch": "Branch", "url": "URL"}], "history": [{"timestamp": "2019-01-01T12:00:00.000Z", "code": "Code", "type": "git_repo"}]}`)
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
				serviceInfoToolchainModel.Crn = core.StringPtr("testString")
				serviceInfoToolchainModel.Guid = core.StringPtr("testString")

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
				serviceInfoToolchainModel.Crn = core.StringPtr("testString")
				serviceInfoToolchainModel.Guid = core.StringPtr("testString")

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
				serviceInfoToolchainModel.Crn = core.StringPtr("testString")
				serviceInfoToolchainModel.Guid = core.StringPtr("testString")

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
		listProjectConfigStatusesPath := "/v1/projects/testString/config_statuses"
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
		listProjectConfigStatusesPath := "/v1/projects/testString/config_statuses"
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
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "config_statuses": [{"href": "Href", "name": "Name", "status": "INSTALLING", "messages": ["Messages"], "pipeline_run": "PipelineRun", "schematics_resource_id": "SchematicsResourceID", "computed_statuses": {"mapKey": "anyValue"}, "output": [{"name": "Name", "description": "Description", "value": ["Value"]}]}]}`)
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
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "href": "Href", "config_statuses": [{"href": "Href", "name": "Name", "status": "INSTALLING", "messages": ["Messages"], "pipeline_run": "PipelineRun", "schematics_resource_id": "SchematicsResourceID", "computed_statuses": {"mapKey": "anyValue"}, "output": [{"name": "Name", "description": "Description", "value": ["Value"]}]}]}`)
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
		getProjectConfigStatusPath := "/v1/projects/b0a2c11d-926c-4653-a15b-ed17d7b34b22/config_statuses/example"
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
		getProjectConfigStatusPath := "/v1/projects/b0a2c11d-926c-4653-a15b-ed17d7b34b22/config_statuses/example"
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "name": "Name", "status": "INSTALLING", "messages": ["Messages"], "pipeline_run": "PipelineRun", "schematics_resource_id": "SchematicsResourceID", "computed_statuses": {"mapKey": "anyValue"}, "output": [{"name": "Name", "description": "Description", "value": ["Value"]}]}`)
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "name": "Name", "status": "INSTALLING", "messages": ["Messages"], "pipeline_run": "PipelineRun", "schematics_resource_id": "SchematicsResourceID", "computed_statuses": {"mapKey": "anyValue"}, "output": [{"name": "Name", "description": "Description", "value": ["Value"]}]}`)
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
	Describe(`UpdateProjectConfigStatus(updateProjectConfigStatusOptions *UpdateProjectConfigStatusOptions) - Operation response error`, func() {
		updateProjectConfigStatusPath := "/v1/projects/b0a2c11d-926c-4653-a15b-ed17d7b34b22/config_statuses/example"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectConfigStatusPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProjectConfigStatus with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("vpc_id")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"testString"}

				// Construct an instance of the UpdateProjectConfigStatusOptions model
				updateProjectConfigStatusOptionsModel := new(projectsv1.UpdateProjectConfigStatusOptions)
				updateProjectConfigStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigStatusOptionsModel.ConfigName = core.StringPtr("example")
				updateProjectConfigStatusOptionsModel.Status = core.StringPtr("UPDATING_FAILED")
				updateProjectConfigStatusOptionsModel.Messages = []string{"Config installation failed"}
				updateProjectConfigStatusOptionsModel.PipelineRun = core.StringPtr("https://url.to.somewhere.failed.install")
				updateProjectConfigStatusOptionsModel.SchematicsResourceID = core.StringPtr("eu-de.workspace.schematicstestkshama.240ff36b")
				updateProjectConfigStatusOptionsModel.Output = []projectsv1.OutputValue{*outputValueModel}
				updateProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.UpdateProjectConfigStatus(updateProjectConfigStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.UpdateProjectConfigStatus(updateProjectConfigStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProjectConfigStatus(updateProjectConfigStatusOptions *UpdateProjectConfigStatusOptions)`, func() {
		updateProjectConfigStatusPath := "/v1/projects/b0a2c11d-926c-4653-a15b-ed17d7b34b22/config_statuses/example"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectConfigStatusPath))
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "name": "Name", "status": "INSTALLING", "messages": ["Messages"], "pipeline_run": "PipelineRun", "schematics_resource_id": "SchematicsResourceID", "computed_statuses": {"mapKey": "anyValue"}, "output": [{"name": "Name", "description": "Description", "value": ["Value"]}]}`)
				}))
			})
			It(`Invoke UpdateProjectConfigStatus successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("vpc_id")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"testString"}

				// Construct an instance of the UpdateProjectConfigStatusOptions model
				updateProjectConfigStatusOptionsModel := new(projectsv1.UpdateProjectConfigStatusOptions)
				updateProjectConfigStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigStatusOptionsModel.ConfigName = core.StringPtr("example")
				updateProjectConfigStatusOptionsModel.Status = core.StringPtr("UPDATING_FAILED")
				updateProjectConfigStatusOptionsModel.Messages = []string{"Config installation failed"}
				updateProjectConfigStatusOptionsModel.PipelineRun = core.StringPtr("https://url.to.somewhere.failed.install")
				updateProjectConfigStatusOptionsModel.SchematicsResourceID = core.StringPtr("eu-de.workspace.schematicstestkshama.240ff36b")
				updateProjectConfigStatusOptionsModel.Output = []projectsv1.OutputValue{*outputValueModel}
				updateProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.UpdateProjectConfigStatusWithContext(ctx, updateProjectConfigStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.UpdateProjectConfigStatus(updateProjectConfigStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.UpdateProjectConfigStatusWithContext(ctx, updateProjectConfigStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectConfigStatusPath))
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "name": "Name", "status": "INSTALLING", "messages": ["Messages"], "pipeline_run": "PipelineRun", "schematics_resource_id": "SchematicsResourceID", "computed_statuses": {"mapKey": "anyValue"}, "output": [{"name": "Name", "description": "Description", "value": ["Value"]}]}`)
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
				result, response, operationErr := projectsService.UpdateProjectConfigStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("vpc_id")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"testString"}

				// Construct an instance of the UpdateProjectConfigStatusOptions model
				updateProjectConfigStatusOptionsModel := new(projectsv1.UpdateProjectConfigStatusOptions)
				updateProjectConfigStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigStatusOptionsModel.ConfigName = core.StringPtr("example")
				updateProjectConfigStatusOptionsModel.Status = core.StringPtr("UPDATING_FAILED")
				updateProjectConfigStatusOptionsModel.Messages = []string{"Config installation failed"}
				updateProjectConfigStatusOptionsModel.PipelineRun = core.StringPtr("https://url.to.somewhere.failed.install")
				updateProjectConfigStatusOptionsModel.SchematicsResourceID = core.StringPtr("eu-de.workspace.schematicstestkshama.240ff36b")
				updateProjectConfigStatusOptionsModel.Output = []projectsv1.OutputValue{*outputValueModel}
				updateProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.UpdateProjectConfigStatus(updateProjectConfigStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				outputValueModel.Name = core.StringPtr("vpc_id")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"testString"}

				// Construct an instance of the UpdateProjectConfigStatusOptions model
				updateProjectConfigStatusOptionsModel := new(projectsv1.UpdateProjectConfigStatusOptions)
				updateProjectConfigStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigStatusOptionsModel.ConfigName = core.StringPtr("example")
				updateProjectConfigStatusOptionsModel.Status = core.StringPtr("UPDATING_FAILED")
				updateProjectConfigStatusOptionsModel.Messages = []string{"Config installation failed"}
				updateProjectConfigStatusOptionsModel.PipelineRun = core.StringPtr("https://url.to.somewhere.failed.install")
				updateProjectConfigStatusOptionsModel.SchematicsResourceID = core.StringPtr("eu-de.workspace.schematicstestkshama.240ff36b")
				updateProjectConfigStatusOptionsModel.Output = []projectsv1.OutputValue{*outputValueModel}
				updateProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.UpdateProjectConfigStatus(updateProjectConfigStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProjectConfigStatusOptions model with no property values
				updateProjectConfigStatusOptionsModelNew := new(projectsv1.UpdateProjectConfigStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.UpdateProjectConfigStatus(updateProjectConfigStatusOptionsModelNew)
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
			It(`Invoke UpdateProjectConfigStatus successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				outputValueModel.Name = core.StringPtr("vpc_id")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"testString"}

				// Construct an instance of the UpdateProjectConfigStatusOptions model
				updateProjectConfigStatusOptionsModel := new(projectsv1.UpdateProjectConfigStatusOptions)
				updateProjectConfigStatusOptionsModel.ID = core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigStatusOptionsModel.ConfigName = core.StringPtr("example")
				updateProjectConfigStatusOptionsModel.Status = core.StringPtr("UPDATING_FAILED")
				updateProjectConfigStatusOptionsModel.Messages = []string{"Config installation failed"}
				updateProjectConfigStatusOptionsModel.PipelineRun = core.StringPtr("https://url.to.somewhere.failed.install")
				updateProjectConfigStatusOptionsModel.SchematicsResourceID = core.StringPtr("eu-de.workspace.schematicstestkshama.240ff36b")
				updateProjectConfigStatusOptionsModel.Output = []projectsv1.OutputValue{*outputValueModel}
				updateProjectConfigStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.UpdateProjectConfigStatus(updateProjectConfigStatusOptionsModel)
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
				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				Expect(projectConfigModel).ToNot(BeNil())
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				Expect(projectConfigModel.Name).To(Equal(core.StringPtr("common-variables")))
				Expect(projectConfigModel.Labels).To(Equal([]string{"testString"}))
				Expect(projectConfigModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				Expect(widgetModel).ToNot(BeNil())
				widgetModel.Name = core.StringPtr("project-properties")
				Expect(widgetModel.Name).To(Equal(core.StringPtr("project-properties")))

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				Expect(projectPrototypeDashboardModel).ToNot(BeNil())
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}
				Expect(projectPrototypeDashboardModel.Widgets).To(Equal([]projectsv1.Widget{*widgetModel}))

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsName := "acme-microservice"
				createProjectOptionsModel := projectsService.NewCreateProjectOptions(createProjectOptionsName)
				createProjectOptionsModel.SetName("acme-microservice")
				createProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.SetConfigs([]projectsv1.ProjectConfigIntf{projectConfigModel})
				createProjectOptionsModel.SetDashboard(projectPrototypeDashboardModel)
				createProjectOptionsModel.SetXIamApi("testString")
				createProjectOptionsModel.SetResourceGroup("Default")
				createProjectOptionsModel.SetLocation("us-south")
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(createProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(createProjectOptionsModel.Configs).To(Equal([]projectsv1.ProjectConfigIntf{projectConfigModel}))
				Expect(createProjectOptionsModel.Dashboard).To(Equal(projectPrototypeDashboardModel))
				Expect(createProjectOptionsModel.XIamApi).To(Equal(core.StringPtr("testString")))
				Expect(createProjectOptionsModel.ResourceGroup).To(Equal(core.StringPtr("Default")))
				Expect(createProjectOptionsModel.Location).To(Equal(core.StringPtr("us-south")))
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
			It(`Invoke NewGetHealthOptions successfully`, func() {
				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := projectsService.NewGetHealthOptions()
				getHealthOptionsModel.SetInfo(false)
				getHealthOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHealthOptionsModel).ToNot(BeNil())
				Expect(getHealthOptionsModel.Info).To(Equal(core.BoolPtr(false)))
				Expect(getHealthOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				id := "234234324-3444-4556-224232432"
				getProjectStatusOptionsModel := projectsService.NewGetProjectStatusOptions(id)
				getProjectStatusOptionsModel.SetID("234234324-3444-4556-224232432")
				getProjectStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectStatusOptionsModel).ToNot(BeNil())
				Expect(getProjectStatusOptionsModel.ID).To(Equal(core.StringPtr("234234324-3444-4556-224232432")))
				Expect(getProjectStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInputVariable successfully`, func() {
				name := "testString"
				_model, err := projectsService.NewInputVariable(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
				listProjectsOptionsModel.SetLimit(int64(1))
				listProjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectsOptionsModel).ToNot(BeNil())
				Expect(listProjectsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listProjectsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listProjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMergeProjectOptions successfully`, func() {
				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				Expect(projectConfigModel).ToNot(BeNil())
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				Expect(projectConfigModel.Name).To(Equal(core.StringPtr("common-variables")))
				Expect(projectConfigModel.Labels).To(Equal([]string{"testString"}))
				Expect(projectConfigModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				Expect(widgetModel).ToNot(BeNil())
				widgetModel.Name = core.StringPtr("project-properties")
				Expect(widgetModel.Name).To(Equal(core.StringPtr("project-properties")))

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				Expect(projectPrototypeDashboardModel).ToNot(BeNil())
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}
				Expect(projectPrototypeDashboardModel.Widgets).To(Equal([]projectsv1.Widget{*widgetModel}))

				// Construct an instance of the MergeProjectOptions model
				id := "testString"
				mergeProjectOptionsName := "acme-microservice"
				mergeProjectOptionsModel := projectsService.NewMergeProjectOptions(id, mergeProjectOptionsName)
				mergeProjectOptionsModel.SetID("testString")
				mergeProjectOptionsModel.SetName("acme-microservice")
				mergeProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure")
				mergeProjectOptionsModel.SetConfigs([]projectsv1.ProjectConfigIntf{projectConfigModel})
				mergeProjectOptionsModel.SetDashboard(projectPrototypeDashboardModel)
				mergeProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(mergeProjectOptionsModel).ToNot(BeNil())
				Expect(mergeProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(mergeProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(mergeProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(mergeProjectOptionsModel.Configs).To(Equal([]projectsv1.ProjectConfigIntf{projectConfigModel}))
				Expect(mergeProjectOptionsModel.Dashboard).To(Equal(projectPrototypeDashboardModel))
				Expect(mergeProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewProjectPrototypeDashboard successfully`, func() {
				widgets := []projectsv1.Widget{}
				_model, err := projectsService.NewProjectPrototypeDashboard(widgets)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
			It(`Invoke NewSchematicsBlueprint successfully`, func() {
				repoURL := "testString"
				definitionFile := "testString"
				_model, err := projectsService.NewSchematicsBlueprint(repoURL, definitionFile)
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
			It(`Invoke NewUpdateProjectConfigStatusOptions successfully`, func() {
				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("vpc_id")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"testString"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("vpc_id")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateProjectConfigStatusOptions model
				id := "b0a2c11d-926c-4653-a15b-ed17d7b34b22"
				configName := "example"
				updateProjectConfigStatusOptionsStatus := "UPDATING_FAILED"
				updateProjectConfigStatusOptionsMessages := []string{"Config installation failed"}
				updateProjectConfigStatusOptionsModel := projectsService.NewUpdateProjectConfigStatusOptions(id, configName, updateProjectConfigStatusOptionsStatus, updateProjectConfigStatusOptionsMessages)
				updateProjectConfigStatusOptionsModel.SetID("b0a2c11d-926c-4653-a15b-ed17d7b34b22")
				updateProjectConfigStatusOptionsModel.SetConfigName("example")
				updateProjectConfigStatusOptionsModel.SetStatus("UPDATING_FAILED")
				updateProjectConfigStatusOptionsModel.SetMessages([]string{"Config installation failed"})
				updateProjectConfigStatusOptionsModel.SetPipelineRun("https://url.to.somewhere.failed.install")
				updateProjectConfigStatusOptionsModel.SetSchematicsResourceID("eu-de.workspace.schematicstestkshama.240ff36b")
				updateProjectConfigStatusOptionsModel.SetOutput([]projectsv1.OutputValue{*outputValueModel})
				updateProjectConfigStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectConfigStatusOptionsModel).ToNot(BeNil())
				Expect(updateProjectConfigStatusOptionsModel.ID).To(Equal(core.StringPtr("b0a2c11d-926c-4653-a15b-ed17d7b34b22")))
				Expect(updateProjectConfigStatusOptionsModel.ConfigName).To(Equal(core.StringPtr("example")))
				Expect(updateProjectConfigStatusOptionsModel.Status).To(Equal(core.StringPtr("UPDATING_FAILED")))
				Expect(updateProjectConfigStatusOptionsModel.Messages).To(Equal([]string{"Config installation failed"}))
				Expect(updateProjectConfigStatusOptionsModel.PipelineRun).To(Equal(core.StringPtr("https://url.to.somewhere.failed.install")))
				Expect(updateProjectConfigStatusOptionsModel.SchematicsResourceID).To(Equal(core.StringPtr("eu-de.workspace.schematicstestkshama.240ff36b")))
				Expect(updateProjectConfigStatusOptionsModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(updateProjectConfigStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectOptions successfully`, func() {
				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				Expect(projectConfigModel).ToNot(BeNil())
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				Expect(projectConfigModel.Name).To(Equal(core.StringPtr("common-variables")))
				Expect(projectConfigModel.Labels).To(Equal([]string{"testString"}))
				Expect(projectConfigModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				Expect(widgetModel).ToNot(BeNil())
				widgetModel.Name = core.StringPtr("project-properties")
				Expect(widgetModel.Name).To(Equal(core.StringPtr("project-properties")))

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				Expect(projectPrototypeDashboardModel).ToNot(BeNil())
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}
				Expect(projectPrototypeDashboardModel.Widgets).To(Equal([]projectsv1.Widget{*widgetModel}))

				// Construct an instance of the UpdateProjectOptions model
				id := "testString"
				updateProjectOptionsName := "acme-microservice"
				updateProjectOptionsModel := projectsService.NewUpdateProjectOptions(id, updateProjectOptionsName)
				updateProjectOptionsModel.SetID("testString")
				updateProjectOptionsModel.SetName("acme-microservice")
				updateProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure")
				updateProjectOptionsModel.SetConfigs([]projectsv1.ProjectConfigIntf{projectConfigModel})
				updateProjectOptionsModel.SetDashboard(projectPrototypeDashboardModel)
				updateProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectOptionsModel).ToNot(BeNil())
				Expect(updateProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(updateProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(updateProjectOptionsModel.Configs).To(Equal([]projectsv1.ProjectConfigIntf{projectConfigModel}))
				Expect(updateProjectOptionsModel.Dashboard).To(Equal(projectPrototypeDashboardModel))
				Expect(updateProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				serviceInfoToolchainModel.Crn = core.StringPtr("testString")
				serviceInfoToolchainModel.Guid = core.StringPtr("testString")
				Expect(serviceInfoToolchainModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(serviceInfoToolchainModel.Guid).To(Equal(core.StringPtr("testString")))

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
			It(`Invoke NewValidateProjectOptions successfully`, func() {
				// Construct an instance of the OutputValue model
				outputValueModel := new(projectsv1.OutputValue)
				Expect(outputValueModel).ToNot(BeNil())
				outputValueModel.Name = core.StringPtr("tags")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"project:ghost", "type:infrastructure"}
				Expect(outputValueModel.Name).To(Equal(core.StringPtr("tags")))
				Expect(outputValueModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(outputValueModel.Value).To(Equal([]string{"project:ghost", "type:infrastructure"}))

				// Construct an instance of the ProjectConfigManualProperty model
				projectConfigModel := new(projectsv1.ProjectConfigManualProperty)
				Expect(projectConfigModel).ToNot(BeNil())
				projectConfigModel.Name = core.StringPtr("common-variables")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Output = []projectsv1.OutputValue{*outputValueModel}
				projectConfigModel.Type = core.StringPtr("manual")
				projectConfigModel.ExternalResourcesAccount = core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")
				Expect(projectConfigModel.Name).To(Equal(core.StringPtr("common-variables")))
				Expect(projectConfigModel.Labels).To(Equal([]string{"testString"}))
				Expect(projectConfigModel.Output).To(Equal([]projectsv1.OutputValue{*outputValueModel}))
				Expect(projectConfigModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigModel.ExternalResourcesAccount).To(Equal(core.StringPtr("e5ed08b9203bad3e4b6f57f0d1675a88")))

				// Construct an instance of the Widget model
				widgetModel := new(projectsv1.Widget)
				Expect(widgetModel).ToNot(BeNil())
				widgetModel.Name = core.StringPtr("project-properties")
				Expect(widgetModel.Name).To(Equal(core.StringPtr("project-properties")))

				// Construct an instance of the ProjectPrototypeDashboard model
				projectPrototypeDashboardModel := new(projectsv1.ProjectPrototypeDashboard)
				Expect(projectPrototypeDashboardModel).ToNot(BeNil())
				projectPrototypeDashboardModel.Widgets = []projectsv1.Widget{*widgetModel}
				Expect(projectPrototypeDashboardModel.Widgets).To(Equal([]projectsv1.Widget{*widgetModel}))

				// Construct an instance of the ValidateProjectOptions model
				validateProjectOptionsName := "acme-microservice"
				validateProjectOptionsModel := projectsService.NewValidateProjectOptions(validateProjectOptionsName)
				validateProjectOptionsModel.SetName("acme-microservice")
				validateProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure")
				validateProjectOptionsModel.SetConfigs([]projectsv1.ProjectConfigIntf{projectConfigModel})
				validateProjectOptionsModel.SetDashboard(projectPrototypeDashboardModel)
				validateProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(validateProjectOptionsModel).ToNot(BeNil())
				Expect(validateProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(validateProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(validateProjectOptionsModel.Configs).To(Equal([]projectsv1.ProjectConfigIntf{projectConfigModel}))
				Expect(validateProjectOptionsModel.Dashboard).To(Equal(projectPrototypeDashboardModel))
				Expect(validateProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewWidget successfully`, func() {
				name := "testString"
				_model, err := projectsService.NewWidget(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProjectConfigManualProperty successfully`, func() {
				name := "testString"
				typeVar := "manual"
				_model, err := projectsService.NewProjectConfigManualProperty(name, typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProjectConfigSchematicsBlueprintProperty successfully`, func() {
				name := "testString"
				typeVar := "schematics_blueprint"
				input := []projectsv1.InputVariable{}
				_model, err := projectsService.NewProjectConfigSchematicsBlueprintProperty(name, typeVar, input)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProjectConfigTerraformTemplateProperty successfully`, func() {
				name := "testString"
				typeVar := "terraform_template"
				input := []projectsv1.InputVariable{}
				_model, err := projectsService.NewProjectConfigTerraformTemplateProperty(name, typeVar, input)
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
