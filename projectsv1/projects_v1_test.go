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
					res.WriteHeader(201)
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

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInput model
				projectConfigInputModel := new(projectsv1.ProjectConfigInput)
				projectConfigInputModel.ID = core.StringPtr("testString")
				projectConfigInputModel.Name = core.StringPtr("common-variables")
				projectConfigInputModel.Labels = []string{"testString"}
				projectConfigInputModel.Description = core.StringPtr("testString")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigInput{*projectConfigInputModel}
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": {"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}, "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
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

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInput model
				projectConfigInputModel := new(projectsv1.ProjectConfigInput)
				projectConfigInputModel.ID = core.StringPtr("testString")
				projectConfigInputModel.Name = core.StringPtr("common-variables")
				projectConfigInputModel.Labels = []string{"testString"}
				projectConfigInputModel.Description = core.StringPtr("testString")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigInput{*projectConfigInputModel}
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": {"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}, "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
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

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInput model
				projectConfigInputModel := new(projectsv1.ProjectConfigInput)
				projectConfigInputModel.ID = core.StringPtr("testString")
				projectConfigInputModel.Name = core.StringPtr("common-variables")
				projectConfigInputModel.Labels = []string{"testString"}
				projectConfigInputModel.Description = core.StringPtr("testString")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigInput{*projectConfigInputModel}
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

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInput model
				projectConfigInputModel := new(projectsv1.ProjectConfigInput)
				projectConfigInputModel.ID = core.StringPtr("testString")
				projectConfigInputModel.Name = core.StringPtr("common-variables")
				projectConfigInputModel.Labels = []string{"testString"}
				projectConfigInputModel.Description = core.StringPtr("testString")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigInput{*projectConfigInputModel}
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
					res.WriteHeader(201)
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

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInput model
				projectConfigInputModel := new(projectsv1.ProjectConfigInput)
				projectConfigInputModel.ID = core.StringPtr("testString")
				projectConfigInputModel.Name = core.StringPtr("common-variables")
				projectConfigInputModel.Labels = []string{"testString"}
				projectConfigInputModel.Description = core.StringPtr("testString")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectsv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.Configs = []projectsv1.ProjectConfigInput{*projectConfigInputModel}
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
					// TODO: Add check for complete query parameter
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
				listProjectsOptionsModel.Complete = core.BoolPtr(false)
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
					// TODO: Add check for complete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 1, "total_count": 0, "first": {"href": "Href", "start": "Start"}, "last": {"href": "Href", "start": "Start"}, "previous": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "projects": [{"id": "ID", "name": "Name", "description": "Description", "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": {"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}, "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State"}}]}`)
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
				listProjectsOptionsModel.Complete = core.BoolPtr(false)
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
					// TODO: Add check for complete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 1, "total_count": 0, "first": {"href": "Href", "start": "Start"}, "last": {"href": "Href", "start": "Start"}, "previous": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "projects": [{"id": "ID", "name": "Name", "description": "Description", "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": {"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}, "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State"}}]}`)
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
				listProjectsOptionsModel.Complete = core.BoolPtr(false)
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
				listProjectsOptionsModel.Complete = core.BoolPtr(false)
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
				listProjectsOptionsModel.Complete = core.BoolPtr(false)
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
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"projects":[{"id":"ID","name":"Name","description":"Description","metadata":{"crn":"crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::","created_at":"2019-01-01T12:00:00.000Z","cumulative_needs_attention_view":{"event":"Event","event_id":"EventID","config_id":"ConfigID","config_version":13},"cumulative_needs_attention_view_err":"CumulativeNeedsAttentionViewErr","location":"Location","resource_group":"ResourceGroup","state":"State"}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"projects":[{"id":"ID","name":"Name","description":"Description","metadata":{"crn":"crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::","created_at":"2019-01-01T12:00:00.000Z","cumulative_needs_attention_view":{"event":"Event","event_id":"EventID","config_id":"ConfigID","config_version":13},"cumulative_needs_attention_view_err":"CumulativeNeedsAttentionViewErr","location":"Location","resource_group":"ResourceGroup","state":"State"}}],"total_count":2,"limit":1}`)
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
					Complete: core.BoolPtr(false),
				}

				pager, err := projectsService.NewProjectsPager(listProjectsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []projectsv1.ProjectListItem
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
					Complete: core.BoolPtr(false),
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
					// TODO: Add check for exclude_configs query parameter
					// TODO: Add check for complete query parameter
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
				getProjectOptionsModel.ExcludeConfigs = core.BoolPtr(false)
				getProjectOptionsModel.Complete = core.BoolPtr(false)
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

					// TODO: Add check for exclude_configs query parameter
					// TODO: Add check for complete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": {"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}, "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
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
				getProjectOptionsModel.ExcludeConfigs = core.BoolPtr(false)
				getProjectOptionsModel.Complete = core.BoolPtr(false)
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

					// TODO: Add check for exclude_configs query parameter
					// TODO: Add check for complete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": {"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}, "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
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
				getProjectOptionsModel.ExcludeConfigs = core.BoolPtr(false)
				getProjectOptionsModel.Complete = core.BoolPtr(false)
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
				getProjectOptionsModel.ExcludeConfigs = core.BoolPtr(false)
				getProjectOptionsModel.Complete = core.BoolPtr(false)
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
				getProjectOptionsModel.ExcludeConfigs = core.BoolPtr(false)
				getProjectOptionsModel.Complete = core.BoolPtr(false)
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
					Expect(req.Method).To(Equal("PATCH"))
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

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				updateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description"}`)
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

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				updateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description"}`)
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

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				updateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
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

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				updateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
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

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectsv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				updateProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure")
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
	Describe(`CreateConfig(createConfigOptions *CreateConfigOptions) - Operation response error`, func() {
		createConfigPath := "/v1/projects/testString/configs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateConfig with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				configSettingItemsModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectsv1.CreateConfigOptions)
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.NewName = core.StringPtr("env-stage")
				createConfigOptionsModel.NewLocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.NewID = core.StringPtr("testString")
				createConfigOptionsModel.NewLabels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.NewDescription = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n")
				createConfigOptionsModel.NewType = core.StringPtr("terraform_template")
				createConfigOptionsModel.NewInput = []projectsv1.InputVariable{*inputVariableModel}
				createConfigOptionsModel.NewSetting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateConfig(createConfigOptions *CreateConfigOptions)`, func() {
		createConfigPath := "/v1/projects/testString/configs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke CreateConfig successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				configSettingItemsModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectsv1.CreateConfigOptions)
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.NewName = core.StringPtr("env-stage")
				createConfigOptionsModel.NewLocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.NewID = core.StringPtr("testString")
				createConfigOptionsModel.NewLabels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.NewDescription = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n")
				createConfigOptionsModel.NewType = core.StringPtr("terraform_template")
				createConfigOptionsModel.NewInput = []projectsv1.InputVariable{*inputVariableModel}
				createConfigOptionsModel.NewSetting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.CreateConfigWithContext(ctx, createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.CreateConfigWithContext(ctx, createConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createConfigPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke CreateConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.CreateConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				configSettingItemsModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectsv1.CreateConfigOptions)
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.NewName = core.StringPtr("env-stage")
				createConfigOptionsModel.NewLocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.NewID = core.StringPtr("testString")
				createConfigOptionsModel.NewLabels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.NewDescription = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n")
				createConfigOptionsModel.NewType = core.StringPtr("terraform_template")
				createConfigOptionsModel.NewInput = []projectsv1.InputVariable{*inputVariableModel}
				createConfigOptionsModel.NewSetting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateConfig with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				configSettingItemsModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectsv1.CreateConfigOptions)
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.NewName = core.StringPtr("env-stage")
				createConfigOptionsModel.NewLocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.NewID = core.StringPtr("testString")
				createConfigOptionsModel.NewLabels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.NewDescription = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n")
				createConfigOptionsModel.NewType = core.StringPtr("terraform_template")
				createConfigOptionsModel.NewInput = []projectsv1.InputVariable{*inputVariableModel}
				createConfigOptionsModel.NewSetting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateConfigOptions model with no property values
				createConfigOptionsModelNew := new(projectsv1.CreateConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.CreateConfig(createConfigOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				configSettingItemsModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectsv1.CreateConfigOptions)
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.NewName = core.StringPtr("env-stage")
				createConfigOptionsModel.NewLocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.NewID = core.StringPtr("testString")
				createConfigOptionsModel.NewLabels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.NewDescription = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n")
				createConfigOptionsModel.NewType = core.StringPtr("terraform_template")
				createConfigOptionsModel.NewInput = []projectsv1.InputVariable{*inputVariableModel}
				createConfigOptionsModel.NewSetting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.CreateConfig(createConfigOptionsModel)
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
	Describe(`ListConfigs(listConfigsOptions *ListConfigsOptions) - Operation response error`, func() {
		listConfigsPath := "/v1/projects/testString/configs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// TODO: Add check for complete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListConfigs with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectsv1.ListConfigsOptions)
				listConfigsOptionsModel.ID = core.StringPtr("testString")
				listConfigsOptionsModel.Version = core.StringPtr("active")
				listConfigsOptionsModel.Complete = core.BoolPtr(false)
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigs(listConfigsOptions *ListConfigsOptions)`, func() {
		listConfigsPath := "/v1/projects/testString/configs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// TODO: Add check for complete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}]}`)
				}))
			})
			It(`Invoke ListConfigs successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectsv1.ListConfigsOptions)
				listConfigsOptionsModel.ID = core.StringPtr("testString")
				listConfigsOptionsModel.Version = core.StringPtr("active")
				listConfigsOptionsModel.Complete = core.BoolPtr(false)
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.ListConfigsWithContext(ctx, listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.ListConfigsWithContext(ctx, listConfigsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listConfigsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// TODO: Add check for complete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}]}`)
				}))
			})
			It(`Invoke ListConfigs successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.ListConfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectsv1.ListConfigsOptions)
				listConfigsOptionsModel.ID = core.StringPtr("testString")
				listConfigsOptionsModel.Version = core.StringPtr("active")
				listConfigsOptionsModel.Complete = core.BoolPtr(false)
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListConfigs with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectsv1.ListConfigsOptions)
				listConfigsOptionsModel.ID = core.StringPtr("testString")
				listConfigsOptionsModel.Version = core.StringPtr("active")
				listConfigsOptionsModel.Complete = core.BoolPtr(false)
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListConfigsOptions model with no property values
				listConfigsOptionsModelNew := new(projectsv1.ListConfigsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.ListConfigs(listConfigsOptionsModelNew)
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
			It(`Invoke ListConfigs successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectsv1.ListConfigsOptions)
				listConfigsOptionsModel.ID = core.StringPtr("testString")
				listConfigsOptionsModel.Version = core.StringPtr("active")
				listConfigsOptionsModel.Complete = core.BoolPtr(false)
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.ListConfigs(listConfigsOptionsModel)
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
	Describe(`GetConfig(getConfigOptions *GetConfigOptions) - Operation response error`, func() {
		getConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// TODO: Add check for complete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConfig with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectsv1.GetConfigOptions)
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigOptionsModel.Version = core.StringPtr("active")
				getConfigOptionsModel.Complete = core.BoolPtr(false)
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfig(getConfigOptions *GetConfigOptions)`, func() {
		getConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// TODO: Add check for complete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke GetConfig successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectsv1.GetConfigOptions)
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigOptionsModel.Version = core.StringPtr("active")
				getConfigOptionsModel.Complete = core.BoolPtr(false)
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetConfigWithContext(ctx, getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetConfigWithContext(ctx, getConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getConfigPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// TODO: Add check for complete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke GetConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectsv1.GetConfigOptions)
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigOptionsModel.Version = core.StringPtr("active")
				getConfigOptionsModel.Complete = core.BoolPtr(false)
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfig with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectsv1.GetConfigOptions)
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigOptionsModel.Version = core.StringPtr("active")
				getConfigOptionsModel.Complete = core.BoolPtr(false)
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigOptions model with no property values
				getConfigOptionsModelNew := new(projectsv1.GetConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetConfig(getConfigOptionsModelNew)
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
			It(`Invoke GetConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectsv1.GetConfigOptions)
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigOptionsModel.Version = core.StringPtr("active")
				getConfigOptionsModel.Complete = core.BoolPtr(false)
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetConfig(getConfigOptionsModel)
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
	Describe(`UpdateConfig(updateConfigOptions *UpdateConfigOptions) - Operation response error`, func() {
		updateConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateConfig with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectConfigInputProjectConfigInputSchematicsTemplate model
				updateProjectConfigInputModel := new(projectsv1.UpdateProjectConfigInputProjectConfigInputSchematicsTemplate)
				updateProjectConfigInputModel.Name = core.StringPtr("testString")
				updateProjectConfigInputModel.Labels = []string{"testString"}
				updateProjectConfigInputModel.Description = core.StringPtr("testString")
				updateProjectConfigInputModel.LocatorID = core.StringPtr("testString")
				updateProjectConfigInputModel.Type = core.StringPtr("terraform_template")
				updateProjectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				updateProjectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectsv1.UpdateConfigOptions)
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ConfigID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = updateProjectConfigInputModel
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateConfig(updateConfigOptions *UpdateConfigOptions)`, func() {
		updateConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke UpdateConfig successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectConfigInputProjectConfigInputSchematicsTemplate model
				updateProjectConfigInputModel := new(projectsv1.UpdateProjectConfigInputProjectConfigInputSchematicsTemplate)
				updateProjectConfigInputModel.Name = core.StringPtr("testString")
				updateProjectConfigInputModel.Labels = []string{"testString"}
				updateProjectConfigInputModel.Description = core.StringPtr("testString")
				updateProjectConfigInputModel.LocatorID = core.StringPtr("testString")
				updateProjectConfigInputModel.Type = core.StringPtr("terraform_template")
				updateProjectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				updateProjectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectsv1.UpdateConfigOptions)
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ConfigID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = updateProjectConfigInputModel
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.UpdateConfigWithContext(ctx, updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.UpdateConfigWithContext(ctx, updateConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke UpdateConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.UpdateConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectConfigInputProjectConfigInputSchematicsTemplate model
				updateProjectConfigInputModel := new(projectsv1.UpdateProjectConfigInputProjectConfigInputSchematicsTemplate)
				updateProjectConfigInputModel.Name = core.StringPtr("testString")
				updateProjectConfigInputModel.Labels = []string{"testString"}
				updateProjectConfigInputModel.Description = core.StringPtr("testString")
				updateProjectConfigInputModel.LocatorID = core.StringPtr("testString")
				updateProjectConfigInputModel.Type = core.StringPtr("terraform_template")
				updateProjectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				updateProjectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectsv1.UpdateConfigOptions)
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ConfigID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = updateProjectConfigInputModel
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateConfig with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectConfigInputProjectConfigInputSchematicsTemplate model
				updateProjectConfigInputModel := new(projectsv1.UpdateProjectConfigInputProjectConfigInputSchematicsTemplate)
				updateProjectConfigInputModel.Name = core.StringPtr("testString")
				updateProjectConfigInputModel.Labels = []string{"testString"}
				updateProjectConfigInputModel.Description = core.StringPtr("testString")
				updateProjectConfigInputModel.LocatorID = core.StringPtr("testString")
				updateProjectConfigInputModel.Type = core.StringPtr("terraform_template")
				updateProjectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				updateProjectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectsv1.UpdateConfigOptions)
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ConfigID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = updateProjectConfigInputModel
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateConfigOptions model with no property values
				updateConfigOptionsModelNew := new(projectsv1.UpdateConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.UpdateConfig(updateConfigOptionsModelNew)
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
			It(`Invoke UpdateConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectConfigInputProjectConfigInputSchematicsTemplate model
				updateProjectConfigInputModel := new(projectsv1.UpdateProjectConfigInputProjectConfigInputSchematicsTemplate)
				updateProjectConfigInputModel.Name = core.StringPtr("testString")
				updateProjectConfigInputModel.Labels = []string{"testString"}
				updateProjectConfigInputModel.Description = core.StringPtr("testString")
				updateProjectConfigInputModel.LocatorID = core.StringPtr("testString")
				updateProjectConfigInputModel.Type = core.StringPtr("terraform_template")
				updateProjectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				updateProjectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectsv1.UpdateConfigOptions)
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ConfigID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = updateProjectConfigInputModel
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.UpdateConfig(updateConfigOptionsModel)
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
	Describe(`DeleteConfig(deleteConfigOptions *DeleteConfigOptions) - Operation response error`, func() {
		deleteConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteConfig with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectsv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.ConfigID = core.StringPtr("testString")
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteConfig(deleteConfigOptions *DeleteConfigOptions)`, func() {
		deleteConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name"}`)
				}))
			})
			It(`Invoke DeleteConfig successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectsv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.ConfigID = core.StringPtr("testString")
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.DeleteConfigWithContext(ctx, deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.DeleteConfigWithContext(ctx, deleteConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name"}`)
				}))
			})
			It(`Invoke DeleteConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.DeleteConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectsv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.ConfigID = core.StringPtr("testString")
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteConfig with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectsv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.ConfigID = core.StringPtr("testString")
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteConfigOptions model with no property values
				deleteConfigOptionsModelNew := new(projectsv1.DeleteConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.DeleteConfig(deleteConfigOptionsModelNew)
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
			It(`Invoke DeleteConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectsv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.ConfigID = core.StringPtr("testString")
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.DeleteConfig(deleteConfigOptionsModel)
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
	Describe(`GetConfigDiff(getConfigDiffOptions *GetConfigDiffOptions) - Operation response error`, func() {
		getConfigDiffPath := "/v1/projects/testString/configs/testString/diff"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigDiffPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConfigDiff with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetConfigDiffOptions model
				getConfigDiffOptionsModel := new(projectsv1.GetConfigDiffOptions)
				getConfigDiffOptionsModel.ID = core.StringPtr("testString")
				getConfigDiffOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetConfigDiff(getConfigDiffOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetConfigDiff(getConfigDiffOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfigDiff(getConfigDiffOptions *GetConfigDiffOptions)`, func() {
		getConfigDiffPath := "/v1/projects/testString/configs/testString/diff"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigDiffPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"added": {"input": [{"name": "Name", "type": "array"}]}, "changed": {"input": [{"name": "Name", "type": "array"}]}, "removed": {"input": [{"name": "Name", "type": "array"}]}}`)
				}))
			})
			It(`Invoke GetConfigDiff successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigDiffOptions model
				getConfigDiffOptionsModel := new(projectsv1.GetConfigDiffOptions)
				getConfigDiffOptionsModel.ID = core.StringPtr("testString")
				getConfigDiffOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetConfigDiffWithContext(ctx, getConfigDiffOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetConfigDiff(getConfigDiffOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetConfigDiffWithContext(ctx, getConfigDiffOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getConfigDiffPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"added": {"input": [{"name": "Name", "type": "array"}]}, "changed": {"input": [{"name": "Name", "type": "array"}]}, "removed": {"input": [{"name": "Name", "type": "array"}]}}`)
				}))
			})
			It(`Invoke GetConfigDiff successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetConfigDiff(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigDiffOptions model
				getConfigDiffOptionsModel := new(projectsv1.GetConfigDiffOptions)
				getConfigDiffOptionsModel.ID = core.StringPtr("testString")
				getConfigDiffOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetConfigDiff(getConfigDiffOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfigDiff with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetConfigDiffOptions model
				getConfigDiffOptionsModel := new(projectsv1.GetConfigDiffOptions)
				getConfigDiffOptionsModel.ID = core.StringPtr("testString")
				getConfigDiffOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetConfigDiff(getConfigDiffOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigDiffOptions model with no property values
				getConfigDiffOptionsModelNew := new(projectsv1.GetConfigDiffOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetConfigDiff(getConfigDiffOptionsModelNew)
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
			It(`Invoke GetConfigDiff successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetConfigDiffOptions model
				getConfigDiffOptionsModel := new(projectsv1.GetConfigDiffOptions)
				getConfigDiffOptionsModel.ID = core.StringPtr("testString")
				getConfigDiffOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetConfigDiff(getConfigDiffOptionsModel)
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
	Describe(`CreateDraftAction(createDraftActionOptions *CreateDraftActionOptions) - Operation response error`, func() {
		createDraftActionPath := "/v1/projects/testString/configs/testString/draft/merge"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDraftActionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDraftAction with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the CreateDraftActionOptions model
				createDraftActionOptionsModel := new(projectsv1.CreateDraftActionOptions)
				createDraftActionOptionsModel.ID = core.StringPtr("testString")
				createDraftActionOptionsModel.ConfigID = core.StringPtr("testString")
				createDraftActionOptionsModel.Action = core.StringPtr("merge")
				createDraftActionOptionsModel.Comment = core.StringPtr("Approving the changes")
				createDraftActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.CreateDraftAction(createDraftActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.CreateDraftAction(createDraftActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDraftAction(createDraftActionOptions *CreateDraftActionOptions)`, func() {
		createDraftActionPath := "/v1/projects/testString/configs/testString/draft/merge"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDraftActionPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke CreateDraftAction successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the CreateDraftActionOptions model
				createDraftActionOptionsModel := new(projectsv1.CreateDraftActionOptions)
				createDraftActionOptionsModel.ID = core.StringPtr("testString")
				createDraftActionOptionsModel.ConfigID = core.StringPtr("testString")
				createDraftActionOptionsModel.Action = core.StringPtr("merge")
				createDraftActionOptionsModel.Comment = core.StringPtr("Approving the changes")
				createDraftActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.CreateDraftActionWithContext(ctx, createDraftActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.CreateDraftAction(createDraftActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.CreateDraftActionWithContext(ctx, createDraftActionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDraftActionPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke CreateDraftAction successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.CreateDraftAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDraftActionOptions model
				createDraftActionOptionsModel := new(projectsv1.CreateDraftActionOptions)
				createDraftActionOptionsModel.ID = core.StringPtr("testString")
				createDraftActionOptionsModel.ConfigID = core.StringPtr("testString")
				createDraftActionOptionsModel.Action = core.StringPtr("merge")
				createDraftActionOptionsModel.Comment = core.StringPtr("Approving the changes")
				createDraftActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.CreateDraftAction(createDraftActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDraftAction with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the CreateDraftActionOptions model
				createDraftActionOptionsModel := new(projectsv1.CreateDraftActionOptions)
				createDraftActionOptionsModel.ID = core.StringPtr("testString")
				createDraftActionOptionsModel.ConfigID = core.StringPtr("testString")
				createDraftActionOptionsModel.Action = core.StringPtr("merge")
				createDraftActionOptionsModel.Comment = core.StringPtr("Approving the changes")
				createDraftActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.CreateDraftAction(createDraftActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDraftActionOptions model with no property values
				createDraftActionOptionsModelNew := new(projectsv1.CreateDraftActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.CreateDraftAction(createDraftActionOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateDraftAction successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the CreateDraftActionOptions model
				createDraftActionOptionsModel := new(projectsv1.CreateDraftActionOptions)
				createDraftActionOptionsModel.ID = core.StringPtr("testString")
				createDraftActionOptionsModel.ConfigID = core.StringPtr("testString")
				createDraftActionOptionsModel.Action = core.StringPtr("merge")
				createDraftActionOptionsModel.Comment = core.StringPtr("Approving the changes")
				createDraftActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.CreateDraftAction(createDraftActionOptionsModel)
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
	Describe(`CheckConfig(checkConfigOptions *CheckConfigOptions)`, func() {
		checkConfigPath := "/v1/projects/testString/configs/testString/check"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkConfigPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke CheckConfig successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.CheckConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectsv1.CheckConfigOptions)
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.ConfigID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CheckConfig with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectsv1.CheckConfigOptions)
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.ConfigID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CheckConfigOptions model with no property values
				checkConfigOptionsModelNew := new(projectsv1.CheckConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.CheckConfig(checkConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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
				installConfigOptionsModel.ConfigID = core.StringPtr("testString")
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
				installConfigOptionsModel.ConfigID = core.StringPtr("testString")
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
				uninstallConfigOptionsModel.ConfigID = core.StringPtr("testString")
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
				uninstallConfigOptionsModel.ConfigID = core.StringPtr("testString")
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
				getSchematicsJobOptionsModel.ConfigID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
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
				getSchematicsJobOptionsModel.ConfigID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
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
				getSchematicsJobOptionsModel.ConfigID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
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
				getSchematicsJobOptionsModel.ConfigID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
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
				getSchematicsJobOptionsModel.ConfigID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
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
	Describe(`GetCostEstimate(getCostEstimateOptions *GetCostEstimateOptions) - Operation response error`, func() {
		getCostEstimatePath := "/v1/projects/testString/configs/testString/cost_estimate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCostEstimatePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCostEstimate with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectsv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigID = core.StringPtr("testString")
				getCostEstimateOptionsModel.Version = core.StringPtr("active")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCostEstimate(getCostEstimateOptions *GetCostEstimateOptions)`, func() {
		getCostEstimatePath := "/v1/projects/testString/configs/testString/cost_estimate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCostEstimatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{}`)
				}))
			})
			It(`Invoke GetCostEstimate successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectsv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigID = core.StringPtr("testString")
				getCostEstimateOptionsModel.Version = core.StringPtr("active")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetCostEstimateWithContext(ctx, getCostEstimateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetCostEstimateWithContext(ctx, getCostEstimateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCostEstimatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{}`)
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
				result, response, operationErr := projectsService.GetCostEstimate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectsv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigID = core.StringPtr("testString")
				getCostEstimateOptionsModel.Version = core.StringPtr("active")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				getCostEstimateOptionsModel.ConfigID = core.StringPtr("testString")
				getCostEstimateOptionsModel.Version = core.StringPtr("active")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCostEstimateOptions model with no property values
				getCostEstimateOptionsModelNew := new(projectsv1.GetCostEstimateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetCostEstimate(getCostEstimateOptionsModelNew)
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
			It(`Invoke GetCostEstimate successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectsv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigID = core.StringPtr("testString")
				getCostEstimateOptionsModel.Version = core.StringPtr("active")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetCostEstimate(getCostEstimateOptionsModel)
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
				receivePulsarCatalogEventsOptionsModel.PulsarCatalogEvents = []projectsv1.PulsarEventItem{*pulsarEventItemModel}
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
				receivePulsarCatalogEventsOptionsModel.PulsarCatalogEvents = []projectsv1.PulsarEventItem{*pulsarEventItemModel}
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "version": "Version", "dependencies": {"anyKey": "anyValue"}}`)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "version": "Version", "dependencies": {"anyKey": "anyValue"}}`)
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
	Describe(`ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions) - Operation response error`, func() {
		replaceServiceInstancePath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceInstancePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					// TODO: Add check for accepts_incomplete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceServiceInstance with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(projectsv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = []string{"testString"}
				replaceServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions)`, func() {
		replaceServiceInstancePath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceInstancePath))
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
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
			It(`Invoke ReplaceServiceInstance successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(projectsv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = []string{"testString"}
				replaceServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.ReplaceServiceInstanceWithContext(ctx, replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.ReplaceServiceInstanceWithContext(ctx, replaceServiceInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceInstancePath))
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					// TODO: Add check for accepts_incomplete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dashboard_url": "DashboardURL", "operation": "Operation"}`)
				}))
			})
			It(`Invoke ReplaceServiceInstance successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.ReplaceServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(projectsv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = []string{"testString"}
				replaceServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceServiceInstance with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(projectsv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = []string{"testString"}
				replaceServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceServiceInstanceOptions model with no property values
				replaceServiceInstanceOptionsModelNew := new(projectsv1.ReplaceServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.ReplaceServiceInstance(replaceServiceInstanceOptionsModelNew)
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
			It(`Invoke ReplaceServiceInstance successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(projectsv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = []string{"testString"}
				replaceServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
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
	Describe(`DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions) - Operation response error`, func() {
		deleteServiceInstancePath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteServiceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
					// TODO: Add check for accepts_incomplete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteServiceInstance with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(projectsv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				deleteServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions)`, func() {
		deleteServiceInstancePath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteServiceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
					// TODO: Add check for accepts_incomplete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{}`)
				}))
			})
			It(`Invoke DeleteServiceInstance successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(projectsv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				deleteServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.DeleteServiceInstanceWithContext(ctx, deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.DeleteServiceInstanceWithContext(ctx, deleteServiceInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteServiceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
					// TODO: Add check for accepts_incomplete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{}`)
				}))
			})
			It(`Invoke DeleteServiceInstance successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.DeleteServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(projectsv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				deleteServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteServiceInstance with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(projectsv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				deleteServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteServiceInstanceOptions model with no property values
				deleteServiceInstanceOptionsModelNew := new(projectsv1.DeleteServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.DeleteServiceInstance(deleteServiceInstanceOptionsModelNew)
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
			It(`Invoke DeleteServiceInstance successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(projectsv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				deleteServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
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
	Describe(`UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions) - Operation response error`, func() {
		updateServiceInstancePath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateServiceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					// TODO: Add check for accepts_incomplete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateServiceInstance with error: Operation response processing error`, func() {
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
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions)`, func() {
		updateServiceInstancePath := "/v2/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with timeout`, func() {
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					// TODO: Add check for accepts_incomplete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{}`)
				}))
			})
			It(`Invoke UpdateServiceInstance successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(projectsv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.ServiceID = []string{"testString"}
				updateServiceInstanceOptionsModel.Context = []string{"testString"}
				updateServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.UpdateServiceInstanceWithContext(ctx, updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.UpdateServiceInstanceWithContext(ctx, updateServiceInstanceOptionsModel)
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["X-Broker-Api-Originating-Identity"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Originating-Identity"][0]).To(Equal(fmt.Sprintf("%v", "ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
					// TODO: Add check for accepts_incomplete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{}`)
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
				result, response, operationErr := projectsService.UpdateServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(projectsv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.ServiceID = []string{"testString"}
				updateServiceInstanceOptionsModel.Context = []string{"testString"}
				updateServiceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateServiceInstanceOptions model with no property values
				updateServiceInstanceOptionsModelNew := new(projectsv1.UpdateServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.UpdateServiceInstance(updateServiceInstanceOptionsModelNew)
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
			It(`Invoke UpdateServiceInstance successfully`, func() {
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
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"ABCD"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
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
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"ABCD"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
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
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"ABCD"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"cb54391b-3316-4943-a5a6-a541678c1924"}))
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
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
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
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
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
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
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
	Describe(`ReplaceServiceInstanceState(replaceServiceInstanceStateOptions *ReplaceServiceInstanceStateOptions) - Operation response error`, func() {
		replaceServiceInstanceStatePath := "/bluemix_v1/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceInstanceStatePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Broker-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceServiceInstanceState with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(projectsv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceServiceInstanceState(replaceServiceInstanceStateOptions *ReplaceServiceInstanceStateOptions)`, func() {
		replaceServiceInstanceStatePath := "/bluemix_v1/service_instances/crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceInstanceStatePath))
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": "Active", "enabled": "Enabled", "last_active": "LastActive"}`)
				}))
			})
			It(`Invoke ReplaceServiceInstanceState successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(projectsv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.ReplaceServiceInstanceStateWithContext(ctx, replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.ReplaceServiceInstanceStateWithContext(ctx, replaceServiceInstanceStateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceInstanceStatePath))
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": "Active", "enabled": "Enabled", "last_active": "LastActive"}`)
				}))
			})
			It(`Invoke ReplaceServiceInstanceState successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.ReplaceServiceInstanceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(projectsv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceServiceInstanceState with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(projectsv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceServiceInstanceStateOptions model with no property values
				replaceServiceInstanceStateOptionsModelNew := new(projectsv1.ReplaceServiceInstanceStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModelNew)
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
			It(`Invoke ReplaceServiceInstanceState successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(projectsv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
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
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
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
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
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
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
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
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
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
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
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
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
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
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
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
					Expect(req.Header["X-Broker-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
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
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
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
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
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
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
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
	Describe(`PostEventNotificationsIntegration(postEventNotificationsIntegrationOptions *PostEventNotificationsIntegrationOptions) - Operation response error`, func() {
		postEventNotificationsIntegrationPath := "/v1/projects/testString/integrations/event_notifications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postEventNotificationsIntegrationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostEventNotificationsIntegration with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				postEventNotificationsIntegrationOptionsModel := new(projectsv1.PostEventNotificationsIntegrationOptions)
				postEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationsIntegrationOptionsModel.InstanceCrn = core.StringPtr("crn of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.Description = core.StringPtr("A sample project source")
				postEventNotificationsIntegrationOptionsModel.Name = core.StringPtr("Project name")
				postEventNotificationsIntegrationOptionsModel.Enabled = core.BoolPtr(true)
				postEventNotificationsIntegrationOptionsModel.Source = core.StringPtr("CRN of the project instance")
				postEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostEventNotificationsIntegration(postEventNotificationsIntegrationOptions *PostEventNotificationsIntegrationOptions)`, func() {
		postEventNotificationsIntegrationPath := "/v1/projects/testString/integrations/event_notifications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postEventNotificationsIntegrationPath))
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
					fmt.Fprintf(res, "%s", `{"description": "Description", "name": "Name", "enabled": false, "id": "ID", "type": "Type", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke PostEventNotificationsIntegration successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				postEventNotificationsIntegrationOptionsModel := new(projectsv1.PostEventNotificationsIntegrationOptions)
				postEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationsIntegrationOptionsModel.InstanceCrn = core.StringPtr("crn of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.Description = core.StringPtr("A sample project source")
				postEventNotificationsIntegrationOptionsModel.Name = core.StringPtr("Project name")
				postEventNotificationsIntegrationOptionsModel.Enabled = core.BoolPtr(true)
				postEventNotificationsIntegrationOptionsModel.Source = core.StringPtr("CRN of the project instance")
				postEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.PostEventNotificationsIntegrationWithContext(ctx, postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.PostEventNotificationsIntegrationWithContext(ctx, postEventNotificationsIntegrationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postEventNotificationsIntegrationPath))
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
					fmt.Fprintf(res, "%s", `{"description": "Description", "name": "Name", "enabled": false, "id": "ID", "type": "Type", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke PostEventNotificationsIntegration successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.PostEventNotificationsIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				postEventNotificationsIntegrationOptionsModel := new(projectsv1.PostEventNotificationsIntegrationOptions)
				postEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationsIntegrationOptionsModel.InstanceCrn = core.StringPtr("crn of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.Description = core.StringPtr("A sample project source")
				postEventNotificationsIntegrationOptionsModel.Name = core.StringPtr("Project name")
				postEventNotificationsIntegrationOptionsModel.Enabled = core.BoolPtr(true)
				postEventNotificationsIntegrationOptionsModel.Source = core.StringPtr("CRN of the project instance")
				postEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostEventNotificationsIntegration with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				postEventNotificationsIntegrationOptionsModel := new(projectsv1.PostEventNotificationsIntegrationOptions)
				postEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationsIntegrationOptionsModel.InstanceCrn = core.StringPtr("crn of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.Description = core.StringPtr("A sample project source")
				postEventNotificationsIntegrationOptionsModel.Name = core.StringPtr("Project name")
				postEventNotificationsIntegrationOptionsModel.Enabled = core.BoolPtr(true)
				postEventNotificationsIntegrationOptionsModel.Source = core.StringPtr("CRN of the project instance")
				postEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostEventNotificationsIntegrationOptions model with no property values
				postEventNotificationsIntegrationOptionsModelNew := new(projectsv1.PostEventNotificationsIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModelNew)
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
			It(`Invoke PostEventNotificationsIntegration successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				postEventNotificationsIntegrationOptionsModel := new(projectsv1.PostEventNotificationsIntegrationOptions)
				postEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationsIntegrationOptionsModel.InstanceCrn = core.StringPtr("crn of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.Description = core.StringPtr("A sample project source")
				postEventNotificationsIntegrationOptionsModel.Name = core.StringPtr("Project name")
				postEventNotificationsIntegrationOptionsModel.Enabled = core.BoolPtr(true)
				postEventNotificationsIntegrationOptionsModel.Source = core.StringPtr("CRN of the project instance")
				postEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
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
	Describe(`GetEventNotificationsIntegration(getEventNotificationsIntegrationOptions *GetEventNotificationsIntegrationOptions) - Operation response error`, func() {
		getEventNotificationsIntegrationPath := "/v1/projects/testString/integrations/event_notifications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEventNotificationsIntegrationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEventNotificationsIntegration with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				getEventNotificationsIntegrationOptionsModel := new(projectsv1.GetEventNotificationsIntegrationOptions)
				getEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				getEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEventNotificationsIntegration(getEventNotificationsIntegrationOptions *GetEventNotificationsIntegrationOptions)`, func() {
		getEventNotificationsIntegrationPath := "/v1/projects/testString/integrations/event_notifications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEventNotificationsIntegrationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Description", "name": "Name", "enabled": false, "id": "ID", "type": "Type", "updated_at": "2019-01-01T12:00:00.000Z", "topic_count": 10, "topic_names": ["TopicNames"]}`)
				}))
			})
			It(`Invoke GetEventNotificationsIntegration successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				getEventNotificationsIntegrationOptionsModel := new(projectsv1.GetEventNotificationsIntegrationOptions)
				getEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				getEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.GetEventNotificationsIntegrationWithContext(ctx, getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.GetEventNotificationsIntegrationWithContext(ctx, getEventNotificationsIntegrationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getEventNotificationsIntegrationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Description", "name": "Name", "enabled": false, "id": "ID", "type": "Type", "updated_at": "2019-01-01T12:00:00.000Z", "topic_count": 10, "topic_names": ["TopicNames"]}`)
				}))
			})
			It(`Invoke GetEventNotificationsIntegration successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.GetEventNotificationsIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				getEventNotificationsIntegrationOptionsModel := new(projectsv1.GetEventNotificationsIntegrationOptions)
				getEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				getEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEventNotificationsIntegration with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				getEventNotificationsIntegrationOptionsModel := new(projectsv1.GetEventNotificationsIntegrationOptions)
				getEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				getEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEventNotificationsIntegrationOptions model with no property values
				getEventNotificationsIntegrationOptionsModelNew := new(projectsv1.GetEventNotificationsIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModelNew)
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
			It(`Invoke GetEventNotificationsIntegration successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				getEventNotificationsIntegrationOptionsModel := new(projectsv1.GetEventNotificationsIntegrationOptions)
				getEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				getEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
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
	Describe(`DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptions *DeleteEventNotificationsIntegrationOptions)`, func() {
		deleteEventNotificationsIntegrationPath := "/v1/projects/testString/integrations/event_notifications"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEventNotificationsIntegrationPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteEventNotificationsIntegration successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectsService.DeleteEventNotificationsIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteEventNotificationsIntegrationOptions model
				deleteEventNotificationsIntegrationOptionsModel := new(projectsv1.DeleteEventNotificationsIntegrationOptions)
				deleteEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				deleteEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectsService.DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteEventNotificationsIntegration with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the DeleteEventNotificationsIntegrationOptions model
				deleteEventNotificationsIntegrationOptionsModel := new(projectsv1.DeleteEventNotificationsIntegrationOptions)
				deleteEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				deleteEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectsService.DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteEventNotificationsIntegrationOptions model with no property values
				deleteEventNotificationsIntegrationOptionsModelNew := new(projectsv1.DeleteEventNotificationsIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectsService.DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostEventNotification(postEventNotificationOptions *PostEventNotificationOptions) - Operation response error`, func() {
		postEventNotificationPath := "/v1/projects/testString/integrations/event_notifications/notifications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postEventNotificationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostEventNotification with error: Operation response processing error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the PostEventNotificationOptions model
				postEventNotificationOptionsModel := new(projectsv1.PostEventNotificationOptions)
				postEventNotificationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationOptionsModel.NewID = core.StringPtr("5f208fef-6b64-413c-aa07-dfed0b46abc1236")
				postEventNotificationOptionsModel.NewSource = core.StringPtr("crn of project")
				postEventNotificationOptionsModel.NewDatacontenttype = core.StringPtr("application/json")
				postEventNotificationOptionsModel.NewIbmendefaultlong = core.StringPtr("long test notification message")
				postEventNotificationOptionsModel.NewIbmendefaultshort = core.StringPtr("Test notification")
				postEventNotificationOptionsModel.NewIbmensourceid = core.StringPtr("crn of project")
				postEventNotificationOptionsModel.NewSpecversion = core.StringPtr("1.0")
				postEventNotificationOptionsModel.NewType = core.StringPtr("com.ibm.cloud.project.project.test_notification")
				postEventNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectsService.PostEventNotification(postEventNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectsService.EnableRetries(0, 0)
				result, response, operationErr = projectsService.PostEventNotification(postEventNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostEventNotification(postEventNotificationOptions *PostEventNotificationOptions)`, func() {
		postEventNotificationPath := "/v1/projects/testString/integrations/event_notifications/notifications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postEventNotificationPath))
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
					fmt.Fprintf(res, "%s", `{"datacontenttype": "Datacontenttype", "ibmendefaultlong": "Ibmendefaultlong", "ibmendefaultshort": "Ibmendefaultshort", "ibmensourceid": "Ibmensourceid", "id": "ID", "source": "Source", "specversion": "Specversion", "type": "Type"}`)
				}))
			})
			It(`Invoke PostEventNotification successfully with retries`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())
				projectsService.EnableRetries(0, 0)

				// Construct an instance of the PostEventNotificationOptions model
				postEventNotificationOptionsModel := new(projectsv1.PostEventNotificationOptions)
				postEventNotificationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationOptionsModel.NewID = core.StringPtr("5f208fef-6b64-413c-aa07-dfed0b46abc1236")
				postEventNotificationOptionsModel.NewSource = core.StringPtr("crn of project")
				postEventNotificationOptionsModel.NewDatacontenttype = core.StringPtr("application/json")
				postEventNotificationOptionsModel.NewIbmendefaultlong = core.StringPtr("long test notification message")
				postEventNotificationOptionsModel.NewIbmendefaultshort = core.StringPtr("Test notification")
				postEventNotificationOptionsModel.NewIbmensourceid = core.StringPtr("crn of project")
				postEventNotificationOptionsModel.NewSpecversion = core.StringPtr("1.0")
				postEventNotificationOptionsModel.NewType = core.StringPtr("com.ibm.cloud.project.project.test_notification")
				postEventNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectsService.PostEventNotificationWithContext(ctx, postEventNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectsService.DisableRetries()
				result, response, operationErr := projectsService.PostEventNotification(postEventNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectsService.PostEventNotificationWithContext(ctx, postEventNotificationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postEventNotificationPath))
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
					fmt.Fprintf(res, "%s", `{"datacontenttype": "Datacontenttype", "ibmendefaultlong": "Ibmendefaultlong", "ibmendefaultshort": "Ibmendefaultshort", "ibmensourceid": "Ibmensourceid", "id": "ID", "source": "Source", "specversion": "Specversion", "type": "Type"}`)
				}))
			})
			It(`Invoke PostEventNotification successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectsService.PostEventNotification(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostEventNotificationOptions model
				postEventNotificationOptionsModel := new(projectsv1.PostEventNotificationOptions)
				postEventNotificationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationOptionsModel.NewID = core.StringPtr("5f208fef-6b64-413c-aa07-dfed0b46abc1236")
				postEventNotificationOptionsModel.NewSource = core.StringPtr("crn of project")
				postEventNotificationOptionsModel.NewDatacontenttype = core.StringPtr("application/json")
				postEventNotificationOptionsModel.NewIbmendefaultlong = core.StringPtr("long test notification message")
				postEventNotificationOptionsModel.NewIbmendefaultshort = core.StringPtr("Test notification")
				postEventNotificationOptionsModel.NewIbmensourceid = core.StringPtr("crn of project")
				postEventNotificationOptionsModel.NewSpecversion = core.StringPtr("1.0")
				postEventNotificationOptionsModel.NewType = core.StringPtr("com.ibm.cloud.project.project.test_notification")
				postEventNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectsService.PostEventNotification(postEventNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostEventNotification with error: Operation validation and request error`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the PostEventNotificationOptions model
				postEventNotificationOptionsModel := new(projectsv1.PostEventNotificationOptions)
				postEventNotificationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationOptionsModel.NewID = core.StringPtr("5f208fef-6b64-413c-aa07-dfed0b46abc1236")
				postEventNotificationOptionsModel.NewSource = core.StringPtr("crn of project")
				postEventNotificationOptionsModel.NewDatacontenttype = core.StringPtr("application/json")
				postEventNotificationOptionsModel.NewIbmendefaultlong = core.StringPtr("long test notification message")
				postEventNotificationOptionsModel.NewIbmendefaultshort = core.StringPtr("Test notification")
				postEventNotificationOptionsModel.NewIbmensourceid = core.StringPtr("crn of project")
				postEventNotificationOptionsModel.NewSpecversion = core.StringPtr("1.0")
				postEventNotificationOptionsModel.NewType = core.StringPtr("com.ibm.cloud.project.project.test_notification")
				postEventNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectsService.PostEventNotification(postEventNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostEventNotificationOptions model with no property values
				postEventNotificationOptionsModelNew := new(projectsv1.PostEventNotificationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectsService.PostEventNotification(postEventNotificationOptionsModelNew)
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
			It(`Invoke PostEventNotification successfully`, func() {
				projectsService, serviceErr := projectsv1.NewProjectsV1(&projectsv1.ProjectsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectsService).ToNot(BeNil())

				// Construct an instance of the PostEventNotificationOptions model
				postEventNotificationOptionsModel := new(projectsv1.PostEventNotificationOptions)
				postEventNotificationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationOptionsModel.NewID = core.StringPtr("5f208fef-6b64-413c-aa07-dfed0b46abc1236")
				postEventNotificationOptionsModel.NewSource = core.StringPtr("crn of project")
				postEventNotificationOptionsModel.NewDatacontenttype = core.StringPtr("application/json")
				postEventNotificationOptionsModel.NewIbmendefaultlong = core.StringPtr("long test notification message")
				postEventNotificationOptionsModel.NewIbmendefaultshort = core.StringPtr("Test notification")
				postEventNotificationOptionsModel.NewIbmensourceid = core.StringPtr("crn of project")
				postEventNotificationOptionsModel.NewSpecversion = core.StringPtr("1.0")
				postEventNotificationOptionsModel.NewType = core.StringPtr("com.ibm.cloud.project.project.test_notification")
				postEventNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectsService.PostEventNotification(postEventNotificationOptionsModel)
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
			It(`Invoke NewCheckConfigOptions successfully`, func() {
				// Construct an instance of the CheckConfigOptions model
				id := "testString"
				configID := "testString"
				checkConfigOptionsModel := projectsService.NewCheckConfigOptions(id, configID)
				checkConfigOptionsModel.SetID("testString")
				checkConfigOptionsModel.SetConfigID("testString")
				checkConfigOptionsModel.SetXAuthRefreshToken("testString")
				checkConfigOptionsModel.SetVersion("active")
				checkConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(checkConfigOptionsModel).ToNot(BeNil())
				Expect(checkConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(checkConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(checkConfigOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(checkConfigOptionsModel.Version).To(Equal(core.StringPtr("active")))
				Expect(checkConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewConfigSettingItems successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := projectsService.NewConfigSettingItems(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateConfigOptions successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.Name = core.StringPtr("account_id")
				inputVariableModel.Type = core.StringPtr("string")
				inputVariableModel.Required = core.BoolPtr(true)
				Expect(inputVariableModel.Name).To(Equal(core.StringPtr("account_id")))
				Expect(inputVariableModel.Type).To(Equal(core.StringPtr("string")))
				Expect(inputVariableModel.Required).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				Expect(configSettingItemsModel).ToNot(BeNil())
				configSettingItemsModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				configSettingItemsModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")
				Expect(configSettingItemsModel.Name).To(Equal(core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")))
				Expect(configSettingItemsModel.Value).To(Equal(core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")))

				// Construct an instance of the CreateConfigOptions model
				id := "testString"
				createConfigOptionsNewName := "env-stage"
				createConfigOptionsNewLocatorID := "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"
				createConfigOptionsModel := projectsService.NewCreateConfigOptions(id, createConfigOptionsNewName, createConfigOptionsNewLocatorID)
				createConfigOptionsModel.SetID("testString")
				createConfigOptionsModel.SetNewName("env-stage")
				createConfigOptionsModel.SetNewLocatorID("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.SetNewID("testString")
				createConfigOptionsModel.SetNewLabels([]string{"env:stage", "governance:test", "build:0"})
				createConfigOptionsModel.SetNewDescription("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n")
				createConfigOptionsModel.SetNewType("terraform_template")
				createConfigOptionsModel.SetNewInput([]projectsv1.InputVariable{*inputVariableModel})
				createConfigOptionsModel.SetNewSetting([]projectsv1.ConfigSettingItems{*configSettingItemsModel})
				createConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createConfigOptionsModel).ToNot(BeNil())
				Expect(createConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createConfigOptionsModel.NewName).To(Equal(core.StringPtr("env-stage")))
				Expect(createConfigOptionsModel.NewLocatorID).To(Equal(core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")))
				Expect(createConfigOptionsModel.NewID).To(Equal(core.StringPtr("testString")))
				Expect(createConfigOptionsModel.NewLabels).To(Equal([]string{"env:stage", "governance:test", "build:0"}))
				Expect(createConfigOptionsModel.NewDescription).To(Equal(core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n")))
				Expect(createConfigOptionsModel.NewType).To(Equal(core.StringPtr("terraform_template")))
				Expect(createConfigOptionsModel.NewInput).To(Equal([]projectsv1.InputVariable{*inputVariableModel}))
				Expect(createConfigOptionsModel.NewSetting).To(Equal([]projectsv1.ConfigSettingItems{*configSettingItemsModel}))
				Expect(createConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDraftActionOptions successfully`, func() {
				// Construct an instance of the CreateDraftActionOptions model
				id := "testString"
				configID := "testString"
				action := "merge"
				createDraftActionOptionsModel := projectsService.NewCreateDraftActionOptions(id, configID, action)
				createDraftActionOptionsModel.SetID("testString")
				createDraftActionOptionsModel.SetConfigID("testString")
				createDraftActionOptionsModel.SetAction("merge")
				createDraftActionOptionsModel.SetComment("Approving the changes")
				createDraftActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDraftActionOptionsModel).ToNot(BeNil())
				Expect(createDraftActionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createDraftActionOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(createDraftActionOptionsModel.Action).To(Equal(core.StringPtr("merge")))
				Expect(createDraftActionOptionsModel.Comment).To(Equal(core.StringPtr("Approving the changes")))
				Expect(createDraftActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectOptions successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectsv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)
				Expect(inputVariableModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(inputVariableModel.Type).To(Equal(core.StringPtr("array")))
				Expect(inputVariableModel.Required).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigSettingItems model
				configSettingItemsModel := new(projectsv1.ConfigSettingItems)
				Expect(configSettingItemsModel).ToNot(BeNil())
				configSettingItemsModel.Name = core.StringPtr("testString")
				configSettingItemsModel.Value = core.StringPtr("testString")
				Expect(configSettingItemsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(configSettingItemsModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigInput model
				projectConfigInputModel := new(projectsv1.ProjectConfigInput)
				Expect(projectConfigInputModel).ToNot(BeNil())
				projectConfigInputModel.ID = core.StringPtr("testString")
				projectConfigInputModel.Name = core.StringPtr("common-variables")
				projectConfigInputModel.Labels = []string{"testString"}
				projectConfigInputModel.Description = core.StringPtr("testString")
				projectConfigInputModel.LocatorID = core.StringPtr("testString")
				projectConfigInputModel.Type = core.StringPtr("manual")
				projectConfigInputModel.Input = []projectsv1.InputVariable{*inputVariableModel}
				projectConfigInputModel.Setting = []projectsv1.ConfigSettingItems{*configSettingItemsModel}
				Expect(projectConfigInputModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputModel.Name).To(Equal(core.StringPtr("common-variables")))
				Expect(projectConfigInputModel.Labels).To(Equal([]string{"testString"}))
				Expect(projectConfigInputModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputModel.LocatorID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigInputModel.Input).To(Equal([]projectsv1.InputVariable{*inputVariableModel}))
				Expect(projectConfigInputModel.Setting).To(Equal([]projectsv1.ConfigSettingItems{*configSettingItemsModel}))

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsName := "acme-microservice"
				createProjectOptionsModel := projectsService.NewCreateProjectOptions(createProjectOptionsName)
				createProjectOptionsModel.SetName("acme-microservice")
				createProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure")
				createProjectOptionsModel.SetConfigs([]projectsv1.ProjectConfigInput{*projectConfigInputModel})
				createProjectOptionsModel.SetResourceGroup("Default")
				createProjectOptionsModel.SetLocation("us-south")
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(createProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(createProjectOptionsModel.Configs).To(Equal([]projectsv1.ProjectConfigInput{*projectConfigInputModel}))
				Expect(createProjectOptionsModel.ResourceGroup).To(Equal(core.StringPtr("Default")))
				Expect(createProjectOptionsModel.Location).To(Equal(core.StringPtr("us-south")))
				Expect(createProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteConfigOptions successfully`, func() {
				// Construct an instance of the DeleteConfigOptions model
				id := "testString"
				configID := "testString"
				deleteConfigOptionsModel := projectsService.NewDeleteConfigOptions(id, configID)
				deleteConfigOptionsModel.SetID("testString")
				deleteConfigOptionsModel.SetConfigID("testString")
				deleteConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteConfigOptionsModel).ToNot(BeNil())
				Expect(deleteConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEventNotificationsIntegrationOptions successfully`, func() {
				// Construct an instance of the DeleteEventNotificationsIntegrationOptions model
				id := "testString"
				deleteEventNotificationsIntegrationOptionsModel := projectsService.NewDeleteEventNotificationsIntegrationOptions(id)
				deleteEventNotificationsIntegrationOptionsModel.SetID("testString")
				deleteEventNotificationsIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEventNotificationsIntegrationOptionsModel).ToNot(BeNil())
				Expect(deleteEventNotificationsIntegrationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEventNotificationsIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteNotificationOptions successfully`, func() {
				// Construct an instance of the DeleteNotificationOptions model
				id := "testString"
				deleteNotificationOptionsModel := projectsService.NewDeleteNotificationOptions(id)
				deleteNotificationOptionsModel.SetID("testString")
				deleteNotificationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteNotificationOptionsModel).ToNot(BeNil())
				Expect(deleteNotificationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteNotificationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewDeleteServiceInstanceOptions successfully`, func() {
				// Construct an instance of the DeleteServiceInstanceOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				planID := "cb54391b-3316-4943-a5a6-a541678c1924"
				serviceID := "cb54391b-3316-4943-a5a6-a541678c1924"
				deleteServiceInstanceOptionsModel := projectsService.NewDeleteServiceInstanceOptions(instanceID, planID, serviceID)
				deleteServiceInstanceOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deleteServiceInstanceOptionsModel.SetPlanID("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.SetServiceID("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.SetXBrokerApiVersion("1.0")
				deleteServiceInstanceOptionsModel.SetXBrokerApiOriginatingIdentity("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deleteServiceInstanceOptionsModel.SetAcceptsIncomplete(false)
				deleteServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(deleteServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(deleteServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")))
				Expect(deleteServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")))
				Expect(deleteServiceInstanceOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1.0")))
				Expect(deleteServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity).To(Equal(core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
				Expect(deleteServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(false)))
				Expect(deleteServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogOptions successfully`, func() {
				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := projectsService.NewGetCatalogOptions()
				getCatalogOptionsModel.SetXBrokerApiVersion("1.0")
				getCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogOptionsModel).ToNot(BeNil())
				Expect(getCatalogOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1.0")))
				Expect(getCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigDiffOptions successfully`, func() {
				// Construct an instance of the GetConfigDiffOptions model
				id := "testString"
				configID := "testString"
				getConfigDiffOptionsModel := projectsService.NewGetConfigDiffOptions(id, configID)
				getConfigDiffOptionsModel.SetID("testString")
				getConfigDiffOptionsModel.SetConfigID("testString")
				getConfigDiffOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigDiffOptionsModel).ToNot(BeNil())
				Expect(getConfigDiffOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigDiffOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigDiffOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigOptions successfully`, func() {
				// Construct an instance of the GetConfigOptions model
				id := "testString"
				configID := "testString"
				getConfigOptionsModel := projectsService.NewGetConfigOptions(id, configID)
				getConfigOptionsModel.SetID("testString")
				getConfigOptionsModel.SetConfigID("testString")
				getConfigOptionsModel.SetVersion("active")
				getConfigOptionsModel.SetComplete(false)
				getConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigOptionsModel).ToNot(BeNil())
				Expect(getConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigOptionsModel.Version).To(Equal(core.StringPtr("active")))
				Expect(getConfigOptionsModel.Complete).To(Equal(core.BoolPtr(false)))
				Expect(getConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCostEstimateOptions successfully`, func() {
				// Construct an instance of the GetCostEstimateOptions model
				id := "testString"
				configID := "testString"
				getCostEstimateOptionsModel := projectsService.NewGetCostEstimateOptions(id, configID)
				getCostEstimateOptionsModel.SetID("testString")
				getCostEstimateOptionsModel.SetConfigID("testString")
				getCostEstimateOptionsModel.SetVersion("active")
				getCostEstimateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCostEstimateOptionsModel).ToNot(BeNil())
				Expect(getCostEstimateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getCostEstimateOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(getCostEstimateOptionsModel.Version).To(Equal(core.StringPtr("active")))
				Expect(getCostEstimateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEventNotificationsIntegrationOptions successfully`, func() {
				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				id := "testString"
				getEventNotificationsIntegrationOptionsModel := projectsService.NewGetEventNotificationsIntegrationOptions(id)
				getEventNotificationsIntegrationOptionsModel.SetID("testString")
				getEventNotificationsIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEventNotificationsIntegrationOptionsModel).ToNot(BeNil())
				Expect(getEventNotificationsIntegrationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getEventNotificationsIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				getLastOperationOptionsModel.SetXBrokerApiVersion("1.0")
				getLastOperationOptionsModel.SetOperation("ABCD")
				getLastOperationOptionsModel.SetPlanID("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.SetServiceID("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLastOperationOptionsModel).ToNot(BeNil())
				Expect(getLastOperationOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(getLastOperationOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1.0")))
				Expect(getLastOperationOptionsModel.Operation).To(Equal(core.StringPtr("ABCD")))
				Expect(getLastOperationOptionsModel.PlanID).To(Equal(core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")))
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
			It(`Invoke NewGetProjectOptions successfully`, func() {
				// Construct an instance of the GetProjectOptions model
				id := "testString"
				getProjectOptionsModel := projectsService.NewGetProjectOptions(id)
				getProjectOptionsModel.SetID("testString")
				getProjectOptionsModel.SetExcludeConfigs(false)
				getProjectOptionsModel.SetComplete(false)
				getProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectOptionsModel).ToNot(BeNil())
				Expect(getProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProjectOptionsModel.ExcludeConfigs).To(Equal(core.BoolPtr(false)))
				Expect(getProjectOptionsModel.Complete).To(Equal(core.BoolPtr(false)))
				Expect(getProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSchematicsJobOptions successfully`, func() {
				// Construct an instance of the GetSchematicsJobOptions model
				id := "testString"
				configID := "testString"
				action := "plan"
				getSchematicsJobOptionsModel := projectsService.NewGetSchematicsJobOptions(id, configID, action)
				getSchematicsJobOptionsModel.SetID("testString")
				getSchematicsJobOptionsModel.SetConfigID("testString")
				getSchematicsJobOptionsModel.SetAction("plan")
				getSchematicsJobOptionsModel.SetSince(int64(38))
				getSchematicsJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSchematicsJobOptionsModel).ToNot(BeNil())
				Expect(getSchematicsJobOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getSchematicsJobOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(getSchematicsJobOptionsModel.Action).To(Equal(core.StringPtr("plan")))
				Expect(getSchematicsJobOptionsModel.Since).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getSchematicsJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServiceInstanceOptions successfully`, func() {
				// Construct an instance of the GetServiceInstanceOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				getServiceInstanceOptionsModel := projectsService.NewGetServiceInstanceOptions(instanceID)
				getServiceInstanceOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.SetXBrokerApiVersion("1.0")
				getServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(getServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(getServiceInstanceOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1.0")))
				Expect(getServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				configID := "testString"
				installConfigOptionsModel := projectsService.NewInstallConfigOptions(id, configID)
				installConfigOptionsModel.SetID("testString")
				installConfigOptionsModel.SetConfigID("testString")
				installConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(installConfigOptionsModel).ToNot(BeNil())
				Expect(installConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(installConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(installConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigsOptions successfully`, func() {
				// Construct an instance of the ListConfigsOptions model
				id := "testString"
				listConfigsOptionsModel := projectsService.NewListConfigsOptions(id)
				listConfigsOptionsModel.SetID("testString")
				listConfigsOptionsModel.SetVersion("active")
				listConfigsOptionsModel.SetComplete(false)
				listConfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigsOptionsModel).ToNot(BeNil())
				Expect(listConfigsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.Version).To(Equal(core.StringPtr("active")))
				Expect(listConfigsOptionsModel.Complete).To(Equal(core.BoolPtr(false)))
				Expect(listConfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectsOptions successfully`, func() {
				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := projectsService.NewListProjectsOptions()
				listProjectsOptionsModel.SetStart("testString")
				listProjectsOptionsModel.SetLimit(int64(10))
				listProjectsOptionsModel.SetComplete(false)
				listProjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectsOptionsModel).ToNot(BeNil())
				Expect(listProjectsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listProjectsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listProjectsOptionsModel.Complete).To(Equal(core.BoolPtr(false)))
				Expect(listProjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewNotificationEvent successfully`, func() {
				event := "testString"
				target := "testString"
				_model, err := projectsService.NewNotificationEvent(event, target)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPostEventNotificationOptions successfully`, func() {
				// Construct an instance of the PostEventNotificationOptions model
				id := "testString"
				postEventNotificationOptionsNewID := "5f208fef-6b64-413c-aa07-dfed0b46abc1236"
				postEventNotificationOptionsNewSource := "crn of project"
				postEventNotificationOptionsModel := projectsService.NewPostEventNotificationOptions(id, postEventNotificationOptionsNewID, postEventNotificationOptionsNewSource)
				postEventNotificationOptionsModel.SetID("testString")
				postEventNotificationOptionsModel.SetNewID("5f208fef-6b64-413c-aa07-dfed0b46abc1236")
				postEventNotificationOptionsModel.SetNewSource("crn of project")
				postEventNotificationOptionsModel.SetNewDatacontenttype("application/json")
				postEventNotificationOptionsModel.SetNewIbmendefaultlong("long test notification message")
				postEventNotificationOptionsModel.SetNewIbmendefaultshort("Test notification")
				postEventNotificationOptionsModel.SetNewIbmensourceid("crn of project")
				postEventNotificationOptionsModel.SetNewSpecversion("1.0")
				postEventNotificationOptionsModel.SetNewType("com.ibm.cloud.project.project.test_notification")
				postEventNotificationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postEventNotificationOptionsModel).ToNot(BeNil())
				Expect(postEventNotificationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(postEventNotificationOptionsModel.NewID).To(Equal(core.StringPtr("5f208fef-6b64-413c-aa07-dfed0b46abc1236")))
				Expect(postEventNotificationOptionsModel.NewSource).To(Equal(core.StringPtr("crn of project")))
				Expect(postEventNotificationOptionsModel.NewDatacontenttype).To(Equal(core.StringPtr("application/json")))
				Expect(postEventNotificationOptionsModel.NewIbmendefaultlong).To(Equal(core.StringPtr("long test notification message")))
				Expect(postEventNotificationOptionsModel.NewIbmendefaultshort).To(Equal(core.StringPtr("Test notification")))
				Expect(postEventNotificationOptionsModel.NewIbmensourceid).To(Equal(core.StringPtr("crn of project")))
				Expect(postEventNotificationOptionsModel.NewSpecversion).To(Equal(core.StringPtr("1.0")))
				Expect(postEventNotificationOptionsModel.NewType).To(Equal(core.StringPtr("com.ibm.cloud.project.project.test_notification")))
				Expect(postEventNotificationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostEventNotificationRequest successfully`, func() {
				id := "testString"
				source := "testString"
				_model, err := projectsService.NewPostEventNotificationRequest(id, source)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPostEventNotificationsIntegrationOptions successfully`, func() {
				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				id := "testString"
				postEventNotificationsIntegrationOptionsInstanceCrn := "crn of event notifications instance"
				postEventNotificationsIntegrationOptionsModel := projectsService.NewPostEventNotificationsIntegrationOptions(id, postEventNotificationsIntegrationOptionsInstanceCrn)
				postEventNotificationsIntegrationOptionsModel.SetID("testString")
				postEventNotificationsIntegrationOptionsModel.SetInstanceCrn("crn of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.SetDescription("A sample project source")
				postEventNotificationsIntegrationOptionsModel.SetName("Project name")
				postEventNotificationsIntegrationOptionsModel.SetEnabled(true)
				postEventNotificationsIntegrationOptionsModel.SetSource("CRN of the project instance")
				postEventNotificationsIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postEventNotificationsIntegrationOptionsModel).ToNot(BeNil())
				Expect(postEventNotificationsIntegrationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(postEventNotificationsIntegrationOptionsModel.InstanceCrn).To(Equal(core.StringPtr("crn of event notifications instance")))
				Expect(postEventNotificationsIntegrationOptionsModel.Description).To(Equal(core.StringPtr("A sample project source")))
				Expect(postEventNotificationsIntegrationOptionsModel.Name).To(Equal(core.StringPtr("Project name")))
				Expect(postEventNotificationsIntegrationOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(postEventNotificationsIntegrationOptionsModel.Source).To(Equal(core.StringPtr("CRN of the project instance")))
				Expect(postEventNotificationsIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewProjectConfigInput successfully`, func() {
				name := "testString"
				locatorID := "testString"
				_model, err := projectsService.NewProjectConfigInput(name, locatorID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
				pulsarCatalogEvents := []projectsv1.PulsarEventItem{}
				receivePulsarCatalogEventsOptionsModel := projectsService.NewReceivePulsarCatalogEventsOptions(pulsarCatalogEvents)
				receivePulsarCatalogEventsOptionsModel.SetPulsarCatalogEvents([]projectsv1.PulsarEventItem{*pulsarEventItemModel})
				receivePulsarCatalogEventsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(receivePulsarCatalogEventsOptionsModel).ToNot(BeNil())
				Expect(receivePulsarCatalogEventsOptionsModel.PulsarCatalogEvents).To(Equal([]projectsv1.PulsarEventItem{*pulsarEventItemModel}))
				Expect(receivePulsarCatalogEventsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceServiceInstanceOptions successfully`, func() {
				// Construct an instance of the ReplaceServiceInstanceOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				replaceServiceInstanceOptionsServiceID := "testString"
				replaceServiceInstanceOptionsPlanID := "testString"
				replaceServiceInstanceOptionsModel := projectsService.NewReplaceServiceInstanceOptions(instanceID, replaceServiceInstanceOptionsServiceID, replaceServiceInstanceOptionsPlanID)
				replaceServiceInstanceOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceOptionsModel.SetServiceID("testString")
				replaceServiceInstanceOptionsModel.SetPlanID("testString")
				replaceServiceInstanceOptionsModel.SetContext([]string{"testString"})
				replaceServiceInstanceOptionsModel.SetParameters(map[string]interface{}{"anyKey": "anyValue"})
				replaceServiceInstanceOptionsModel.SetPreviousValues([]string{"testString"})
				replaceServiceInstanceOptionsModel.SetXBrokerApiVersion("1.0")
				replaceServiceInstanceOptionsModel.SetXBrokerApiOriginatingIdentity("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				replaceServiceInstanceOptionsModel.SetAcceptsIncomplete(false)
				replaceServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(replaceServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(replaceServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceInstanceOptionsModel.Context).To(Equal([]string{"testString"}))
				Expect(replaceServiceInstanceOptionsModel.Parameters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(replaceServiceInstanceOptionsModel.PreviousValues).To(Equal([]string{"testString"}))
				Expect(replaceServiceInstanceOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1.0")))
				Expect(replaceServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity).To(Equal(core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
				Expect(replaceServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(false)))
				Expect(replaceServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceServiceInstanceStateOptions successfully`, func() {
				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				replaceServiceInstanceStateOptionsEnabled := true
				replaceServiceInstanceStateOptionsModel := projectsService.NewReplaceServiceInstanceStateOptions(instanceID, replaceServiceInstanceStateOptionsEnabled)
				replaceServiceInstanceStateOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceStateOptionsModel.SetEnabled(true)
				replaceServiceInstanceStateOptionsModel.SetInitiatorID("testString")
				replaceServiceInstanceStateOptionsModel.SetReasonCode(map[string]interface{}{"anyKey": "anyValue"})
				replaceServiceInstanceStateOptionsModel.SetPlanID("testString")
				replaceServiceInstanceStateOptionsModel.SetPreviousValues([]string{"testString"})
				replaceServiceInstanceStateOptionsModel.SetXBrokerApiVersion("1.0")
				replaceServiceInstanceStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceServiceInstanceStateOptionsModel).ToNot(BeNil())
				Expect(replaceServiceInstanceStateOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(replaceServiceInstanceStateOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(replaceServiceInstanceStateOptionsModel.InitiatorID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceInstanceStateOptionsModel.ReasonCode).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(replaceServiceInstanceStateOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceInstanceStateOptionsModel.PreviousValues).To(Equal([]string{"testString"}))
				Expect(replaceServiceInstanceStateOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1.0")))
				Expect(replaceServiceInstanceStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUninstallConfigOptions successfully`, func() {
				// Construct an instance of the UninstallConfigOptions model
				id := "testString"
				configID := "testString"
				uninstallConfigOptionsModel := projectsService.NewUninstallConfigOptions(id, configID)
				uninstallConfigOptionsModel.SetID("testString")
				uninstallConfigOptionsModel.SetConfigID("testString")
				uninstallConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uninstallConfigOptionsModel).ToNot(BeNil())
				Expect(uninstallConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(uninstallConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(uninstallConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateConfigOptions successfully`, func() {
				// Construct an instance of the UpdateProjectConfigInputProjectConfigManualProperty model
				updateProjectConfigInputModel := new(projectsv1.UpdateProjectConfigInputProjectConfigManualProperty)
				Expect(updateProjectConfigInputModel).ToNot(BeNil())
				updateProjectConfigInputModel.Name = core.StringPtr("testString")
				updateProjectConfigInputModel.Labels = []string{"testString"}
				updateProjectConfigInputModel.Description = core.StringPtr("testString")
				updateProjectConfigInputModel.Type = core.StringPtr("manual")
				updateProjectConfigInputModel.ExternalResourcesAccount = core.StringPtr("testString")
				Expect(updateProjectConfigInputModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectConfigInputModel.Labels).To(Equal([]string{"testString"}))
				Expect(updateProjectConfigInputModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectConfigInputModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(updateProjectConfigInputModel.ExternalResourcesAccount).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateConfigOptions model
				id := "testString"
				configID := "testString"
				var projectConfig projectsv1.UpdateProjectConfigInputIntf = nil
				updateConfigOptionsModel := projectsService.NewUpdateConfigOptions(id, configID, projectConfig)
				updateConfigOptionsModel.SetID("testString")
				updateConfigOptionsModel.SetConfigID("testString")
				updateConfigOptionsModel.SetProjectConfig(updateProjectConfigInputModel)
				updateConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateConfigOptionsModel).ToNot(BeNil())
				Expect(updateConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigOptionsModel.ProjectConfig).To(Equal(updateProjectConfigInputModel))
				Expect(updateConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectOptions successfully`, func() {
				// Construct an instance of the UpdateProjectOptions model
				id := "testString"
				updateProjectOptionsModel := projectsService.NewUpdateProjectOptions(id)
				updateProjectOptionsModel.SetID("testString")
				updateProjectOptionsModel.SetName("acme-microservice")
				updateProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure")
				updateProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectOptionsModel).ToNot(BeNil())
				Expect(updateProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(updateProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure")))
				Expect(updateProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				updateServiceInstanceOptionsModel.SetXBrokerApiVersion("1.0")
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
				Expect(updateServiceInstanceOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1.0")))
				Expect(updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity).To(Equal(core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
				Expect(updateServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(false)))
				Expect(updateServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectConfigInputProjectConfigManualProperty successfully`, func() {
				typeVar := "manual"
				_model, err := projectsService.NewUpdateProjectConfigInputProjectConfigManualProperty(typeVar)
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
