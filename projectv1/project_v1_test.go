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

package projectv1_test

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
	"github.com/damianovesperini/platform-services-go-sdk/projectv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe(`ProjectV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(projectService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(projectService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
				URL: "https://projectv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(projectService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PROJECT_URL": "https://projectv1/api",
				"PROJECT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				projectService, serviceErr := projectv1.NewProjectV1UsingExternalConfig(&projectv1.ProjectV1Options{
				})
				Expect(projectService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := projectService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != projectService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(projectService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(projectService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				projectService, serviceErr := projectv1.NewProjectV1UsingExternalConfig(&projectv1.ProjectV1Options{
					URL: "https://testService/api",
				})
				Expect(projectService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(projectService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := projectService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != projectService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(projectService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(projectService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				projectService, serviceErr := projectv1.NewProjectV1UsingExternalConfig(&projectv1.ProjectV1Options{
				})
				err := projectService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(projectService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := projectService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != projectService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(projectService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(projectService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PROJECT_URL": "https://projectv1/api",
				"PROJECT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			projectService, serviceErr := projectv1.NewProjectV1UsingExternalConfig(&projectv1.ProjectV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(projectService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PROJECT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			projectService, serviceErr := projectv1.NewProjectV1UsingExternalConfig(&projectv1.ProjectV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(projectService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = projectv1.GetServiceURLForRegion("INVALID_REGION")
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.Configs = []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.CreateProject(createProjectOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
				}))
			})
			It(`Invoke CreateProject successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.Configs = []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.CreateProjectWithContext(ctx, createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.CreateProjectWithContext(ctx, createProjectOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
				}))
			})
			It(`Invoke CreateProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.CreateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.Configs = []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProject with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.Configs = []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProjectOptions model with no property values
				createProjectOptionsModelNew := new(projectv1.CreateProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.CreateProject(createProjectOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.Configs = []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.CreateProject(createProjectOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Complete = core.BoolPtr(false)
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ListProjects(listProjectsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"limit": 1, "total_count": 0, "first": {"href": "Href", "start": "Start"}, "last": {"href": "Href", "start": "Start"}, "previous": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "projects": [{"id": "ID", "name": "Name", "description": "Description", "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}]}`)
				}))
			})
			It(`Invoke ListProjects successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Complete = core.BoolPtr(false)
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ListProjectsWithContext(ctx, listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ListProjectsWithContext(ctx, listProjectsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"limit": 1, "total_count": 0, "first": {"href": "Href", "start": "Start"}, "last": {"href": "Href", "start": "Start"}, "previous": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "projects": [{"id": "ID", "name": "Name", "description": "Description", "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}]}`)
				}))
			})
			It(`Invoke ListProjects successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ListProjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Complete = core.BoolPtr(false)
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProjects with error: Operation request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Complete = core.BoolPtr(false)
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ListProjects(listProjectsOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Complete = core.BoolPtr(false)
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ListProjects(listProjectsOptionsModel)
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
				responseObject := new(projectv1.ProjectCollection)
				nextObject := new(projectv1.PaginationLink)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(projectv1.ProjectCollection)
	
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
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"projects":[{"id":"ID","name":"Name","description":"Description","metadata":{"crn":"crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::","created_at":"2019-01-01T12:00:00.000Z","cumulative_needs_attention_view":[{"event":"Event","event_id":"EventID","config_id":"ConfigID","config_version":13}],"cumulative_needs_attention_view_err":"CumulativeNeedsAttentionViewErr","location":"Location","resource_group":"ResourceGroup","state":"State","event_notifications_crn":"EventNotificationsCrn"}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"projects":[{"id":"ID","name":"Name","description":"Description","metadata":{"crn":"crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::","created_at":"2019-01-01T12:00:00.000Z","cumulative_needs_attention_view":[{"event":"Event","event_id":"EventID","config_id":"ConfigID","config_version":13}],"cumulative_needs_attention_view_err":"CumulativeNeedsAttentionViewErr","location":"Location","resource_group":"ResourceGroup","state":"State","event_notifications_crn":"EventNotificationsCrn"}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ProjectsPager.GetNext successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				listProjectsOptionsModel := &projectv1.ListProjectsOptions{
					Limit: core.Int64Ptr(int64(10)),
					Complete: core.BoolPtr(false),
				}

				pager, err := projectService.NewProjectsPager(listProjectsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []projectv1.ProjectCollectionMemberWithMetadata
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ProjectsPager.GetAll successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				listProjectsOptionsModel := &projectv1.ListProjectsOptions{
					Limit: core.Int64Ptr(int64(10)),
					Complete: core.BoolPtr(false),
				}

				pager, err := projectService.NewProjectsPager(listProjectsOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.ExcludeConfigs = core.BoolPtr(false)
				getProjectOptionsModel.Complete = core.BoolPtr(false)
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetProject(getProjectOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
				}))
			})
			It(`Invoke GetProject successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.ExcludeConfigs = core.BoolPtr(false)
				getProjectOptionsModel.Complete = core.BoolPtr(false)
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetProjectWithContext(ctx, getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetProjectWithContext(ctx, getProjectOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
				}))
			})
			It(`Invoke GetProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.ExcludeConfigs = core.BoolPtr(false)
				getProjectOptionsModel.Complete = core.BoolPtr(false)
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProject with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.ExcludeConfigs = core.BoolPtr(false)
				getProjectOptionsModel.Complete = core.BoolPtr(false)
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectOptions model with no property values
				getProjectOptionsModelNew := new(projectv1.GetProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetProject(getProjectOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.ExcludeConfigs = core.BoolPtr(false)
				getProjectOptionsModel.Complete = core.BoolPtr(false)
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetProject(getProjectOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.JSONPatchOperation = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.UpdateProject(updateProjectOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
				}))
			})
			It(`Invoke UpdateProject successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.JSONPatchOperation = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.UpdateProjectWithContext(ctx, updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.UpdateProjectWithContext(ctx, updateProjectOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "id": "ID", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
				}))
			})
			It(`Invoke UpdateProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.UpdateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.JSONPatchOperation = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProject with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.JSONPatchOperation = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProjectOptions model with no property values
				updateProjectOptionsModelNew := new(projectv1.UpdateProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.UpdateProject(updateProjectOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.JSONPatchOperation = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.UpdateProject(updateProjectOptionsModel)
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

					// TODO: Add check for destroy query parameter
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectService.DeleteProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProjectOptions model
				deleteProjectOptionsModel := new(projectv1.DeleteProjectOptions)
				deleteProjectOptionsModel.ID = core.StringPtr("testString")
				deleteProjectOptionsModel.Destroy = core.BoolPtr(false)
				deleteProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectService.DeleteProject(deleteProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProject with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteProjectOptions model
				deleteProjectOptionsModel := new(projectv1.DeleteProjectOptions)
				deleteProjectOptionsModel.ID = core.StringPtr("testString")
				deleteProjectOptionsModel.Destroy = core.BoolPtr(false)
				deleteProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectService.DeleteProject(deleteProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProjectOptions model with no property values
				deleteProjectOptionsModelNew := new(projectv1.DeleteProjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectService.DeleteProject(deleteProjectOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.NewName = core.StringPtr("env-stage")
				createConfigOptionsModel.NewLocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.NewID = core.StringPtr("testString")
				createConfigOptionsModel.NewLabels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.NewDescription = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.NewInput = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				createConfigOptionsModel.NewSetting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.CreateConfig(createConfigOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.NewName = core.StringPtr("env-stage")
				createConfigOptionsModel.NewLocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.NewID = core.StringPtr("testString")
				createConfigOptionsModel.NewLabels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.NewDescription = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.NewInput = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				createConfigOptionsModel.NewSetting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.CreateConfigWithContext(ctx, createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.CreateConfigWithContext(ctx, createConfigOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.CreateConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.NewName = core.StringPtr("env-stage")
				createConfigOptionsModel.NewLocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.NewID = core.StringPtr("testString")
				createConfigOptionsModel.NewLabels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.NewDescription = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.NewInput = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				createConfigOptionsModel.NewSetting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.NewName = core.StringPtr("env-stage")
				createConfigOptionsModel.NewLocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.NewID = core.StringPtr("testString")
				createConfigOptionsModel.NewLabels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.NewDescription = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.NewInput = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				createConfigOptionsModel.NewSetting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateConfigOptions model with no property values
				createConfigOptionsModelNew := new(projectv1.CreateConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.CreateConfig(createConfigOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.NewName = core.StringPtr("env-stage")
				createConfigOptionsModel.NewLocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.NewID = core.StringPtr("testString")
				createConfigOptionsModel.NewLabels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.NewDescription = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.NewInput = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				createConfigOptionsModel.NewSetting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.CreateConfig(createConfigOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectv1.ListConfigsOptions)
				listConfigsOptionsModel.ID = core.StringPtr("testString")
				listConfigsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigsOptionsModel.Version = core.StringPtr("active")
				listConfigsOptionsModel.Complete = core.BoolPtr(false)
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ListConfigs(listConfigsOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectv1.ListConfigsOptions)
				listConfigsOptionsModel.ID = core.StringPtr("testString")
				listConfigsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigsOptionsModel.Version = core.StringPtr("active")
				listConfigsOptionsModel.Complete = core.BoolPtr(false)
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ListConfigsWithContext(ctx, listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ListConfigsWithContext(ctx, listConfigsOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ListConfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectv1.ListConfigsOptions)
				listConfigsOptionsModel.ID = core.StringPtr("testString")
				listConfigsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigsOptionsModel.Version = core.StringPtr("active")
				listConfigsOptionsModel.Complete = core.BoolPtr(false)
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListConfigs with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectv1.ListConfigsOptions)
				listConfigsOptionsModel.ID = core.StringPtr("testString")
				listConfigsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigsOptionsModel.Version = core.StringPtr("active")
				listConfigsOptionsModel.Complete = core.BoolPtr(false)
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListConfigsOptions model with no property values
				listConfigsOptionsModelNew := new(projectv1.ListConfigsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.ListConfigs(listConfigsOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectv1.ListConfigsOptions)
				listConfigsOptionsModel.ID = core.StringPtr("testString")
				listConfigsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigsOptionsModel.Version = core.StringPtr("active")
				listConfigsOptionsModel.Complete = core.BoolPtr(false)
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ListConfigs(listConfigsOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectv1.GetConfigOptions)
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigOptionsModel.Version = core.StringPtr("active")
				getConfigOptionsModel.Complete = core.BoolPtr(false)
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetConfig(getConfigOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectv1.GetConfigOptions)
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigOptionsModel.Version = core.StringPtr("active")
				getConfigOptionsModel.Complete = core.BoolPtr(false)
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetConfigWithContext(ctx, getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetConfigWithContext(ctx, getConfigOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectv1.GetConfigOptions)
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigOptionsModel.Version = core.StringPtr("active")
				getConfigOptionsModel.Complete = core.BoolPtr(false)
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectv1.GetConfigOptions)
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigOptionsModel.Version = core.StringPtr("active")
				getConfigOptionsModel.Complete = core.BoolPtr(false)
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigOptions model with no property values
				getConfigOptionsModelNew := new(projectv1.GetConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetConfig(getConfigOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectv1.GetConfigOptions)
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigOptionsModel.Version = core.StringPtr("active")
				getConfigOptionsModel.Complete = core.BoolPtr(false)
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetConfig(getConfigOptionsModel)
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
					// TODO: Add check for complete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ConfigID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateConfigOptionsModel.Complete = core.BoolPtr(false)
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.UpdateConfig(updateConfigOptionsModel)
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

					// TODO: Add check for complete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke UpdateConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ConfigID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateConfigOptionsModel.Complete = core.BoolPtr(false)
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.UpdateConfigWithContext(ctx, updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.UpdateConfigWithContext(ctx, updateConfigOptionsModel)
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

					// TODO: Add check for complete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke UpdateConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.UpdateConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ConfigID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateConfigOptionsModel.Complete = core.BoolPtr(false)
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ConfigID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateConfigOptionsModel.Complete = core.BoolPtr(false)
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateConfigOptions model with no property values
				updateConfigOptionsModelNew := new(projectv1.UpdateConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.UpdateConfig(updateConfigOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ConfigID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateConfigOptionsModel.Complete = core.BoolPtr(false)
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.UpdateConfig(updateConfigOptionsModel)
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
					// TODO: Add check for draft_only query parameter
					// TODO: Add check for destroy query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.ConfigID = core.StringPtr("testString")
				deleteConfigOptionsModel.DraftOnly = core.BoolPtr(false)
				deleteConfigOptionsModel.Destroy = core.BoolPtr(false)
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.DeleteConfig(deleteConfigOptionsModel)
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

					// TODO: Add check for draft_only query parameter
					// TODO: Add check for destroy query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name"}`)
				}))
			})
			It(`Invoke DeleteConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.ConfigID = core.StringPtr("testString")
				deleteConfigOptionsModel.DraftOnly = core.BoolPtr(false)
				deleteConfigOptionsModel.Destroy = core.BoolPtr(false)
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.DeleteConfigWithContext(ctx, deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.DeleteConfigWithContext(ctx, deleteConfigOptionsModel)
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

					// TODO: Add check for draft_only query parameter
					// TODO: Add check for destroy query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name"}`)
				}))
			})
			It(`Invoke DeleteConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.DeleteConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.ConfigID = core.StringPtr("testString")
				deleteConfigOptionsModel.DraftOnly = core.BoolPtr(false)
				deleteConfigOptionsModel.Destroy = core.BoolPtr(false)
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.ConfigID = core.StringPtr("testString")
				deleteConfigOptionsModel.DraftOnly = core.BoolPtr(false)
				deleteConfigOptionsModel.Destroy = core.BoolPtr(false)
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteConfigOptions model with no property values
				deleteConfigOptionsModelNew := new(projectv1.DeleteConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.DeleteConfig(deleteConfigOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.ConfigID = core.StringPtr("testString")
				deleteConfigOptionsModel.DraftOnly = core.BoolPtr(false)
				deleteConfigOptionsModel.Destroy = core.BoolPtr(false)
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.DeleteConfig(deleteConfigOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigDiffOptions model
				getConfigDiffOptionsModel := new(projectv1.GetConfigDiffOptions)
				getConfigDiffOptionsModel.ID = core.StringPtr("testString")
				getConfigDiffOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetConfigDiff(getConfigDiffOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetConfigDiff(getConfigDiffOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigDiffOptions model
				getConfigDiffOptionsModel := new(projectv1.GetConfigDiffOptions)
				getConfigDiffOptionsModel.ID = core.StringPtr("testString")
				getConfigDiffOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetConfigDiffWithContext(ctx, getConfigDiffOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetConfigDiff(getConfigDiffOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetConfigDiffWithContext(ctx, getConfigDiffOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetConfigDiff(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigDiffOptions model
				getConfigDiffOptionsModel := new(projectv1.GetConfigDiffOptions)
				getConfigDiffOptionsModel.ID = core.StringPtr("testString")
				getConfigDiffOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetConfigDiff(getConfigDiffOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfigDiff with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigDiffOptions model
				getConfigDiffOptionsModel := new(projectv1.GetConfigDiffOptions)
				getConfigDiffOptionsModel.ID = core.StringPtr("testString")
				getConfigDiffOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetConfigDiff(getConfigDiffOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigDiffOptions model with no property values
				getConfigDiffOptionsModelNew := new(projectv1.GetConfigDiffOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetConfigDiff(getConfigDiffOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigDiffOptions model
				getConfigDiffOptionsModel := new(projectv1.GetConfigDiffOptions)
				getConfigDiffOptionsModel.ID = core.StringPtr("testString")
				getConfigDiffOptionsModel.ConfigID = core.StringPtr("testString")
				getConfigDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetConfigDiff(getConfigDiffOptionsModel)
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
	Describe(`ForceApprove(forceApproveOptions *ForceApproveOptions) - Operation response error`, func() {
		forceApprovePath := "/v1/projects/testString/configs/testString/force_approve"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(forceApprovePath))
					Expect(req.Method).To(Equal("POST"))
					// TODO: Add check for complete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ForceApprove with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ForceApproveOptions model
				forceApproveOptionsModel := new(projectv1.ForceApproveOptions)
				forceApproveOptionsModel.ID = core.StringPtr("testString")
				forceApproveOptionsModel.ConfigID = core.StringPtr("testString")
				forceApproveOptionsModel.Comment = core.StringPtr("Approving the changes")
				forceApproveOptionsModel.Complete = core.BoolPtr(false)
				forceApproveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ForceApprove(forceApproveOptions *ForceApproveOptions)`, func() {
		forceApprovePath := "/v1/projects/testString/configs/testString/force_approve"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(forceApprovePath))
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

					// TODO: Add check for complete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke ForceApprove successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ForceApproveOptions model
				forceApproveOptionsModel := new(projectv1.ForceApproveOptions)
				forceApproveOptionsModel.ID = core.StringPtr("testString")
				forceApproveOptionsModel.ConfigID = core.StringPtr("testString")
				forceApproveOptionsModel.Comment = core.StringPtr("Approving the changes")
				forceApproveOptionsModel.Complete = core.BoolPtr(false)
				forceApproveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ForceApproveWithContext(ctx, forceApproveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ForceApproveWithContext(ctx, forceApproveOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(forceApprovePath))
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

					// TODO: Add check for complete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke ForceApprove successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ForceApprove(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ForceApproveOptions model
				forceApproveOptionsModel := new(projectv1.ForceApproveOptions)
				forceApproveOptionsModel.ID = core.StringPtr("testString")
				forceApproveOptionsModel.ConfigID = core.StringPtr("testString")
				forceApproveOptionsModel.Comment = core.StringPtr("Approving the changes")
				forceApproveOptionsModel.Complete = core.BoolPtr(false)
				forceApproveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ForceApprove with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ForceApproveOptions model
				forceApproveOptionsModel := new(projectv1.ForceApproveOptions)
				forceApproveOptionsModel.ID = core.StringPtr("testString")
				forceApproveOptionsModel.ConfigID = core.StringPtr("testString")
				forceApproveOptionsModel.Comment = core.StringPtr("Approving the changes")
				forceApproveOptionsModel.Complete = core.BoolPtr(false)
				forceApproveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ForceApproveOptions model with no property values
				forceApproveOptionsModelNew := new(projectv1.ForceApproveOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.ForceApprove(forceApproveOptionsModelNew)
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
			It(`Invoke ForceApprove successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ForceApproveOptions model
				forceApproveOptionsModel := new(projectv1.ForceApproveOptions)
				forceApproveOptionsModel.ID = core.StringPtr("testString")
				forceApproveOptionsModel.ConfigID = core.StringPtr("testString")
				forceApproveOptionsModel.Comment = core.StringPtr("Approving the changes")
				forceApproveOptionsModel.Complete = core.BoolPtr(false)
				forceApproveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ForceApprove(forceApproveOptionsModel)
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
	Describe(`Approve(approveOptions *ApproveOptions) - Operation response error`, func() {
		approvePath := "/v1/projects/testString/configs/testString/approve"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(approvePath))
					Expect(req.Method).To(Equal("POST"))
					// TODO: Add check for complete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Approve with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ApproveOptions model
				approveOptionsModel := new(projectv1.ApproveOptions)
				approveOptionsModel.ID = core.StringPtr("testString")
				approveOptionsModel.ConfigID = core.StringPtr("testString")
				approveOptionsModel.Comment = core.StringPtr("Approving the changes")
				approveOptionsModel.Complete = core.BoolPtr(false)
				approveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.Approve(approveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.Approve(approveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Approve(approveOptions *ApproveOptions)`, func() {
		approvePath := "/v1/projects/testString/configs/testString/approve"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(approvePath))
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

					// TODO: Add check for complete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke Approve successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ApproveOptions model
				approveOptionsModel := new(projectv1.ApproveOptions)
				approveOptionsModel.ID = core.StringPtr("testString")
				approveOptionsModel.ConfigID = core.StringPtr("testString")
				approveOptionsModel.Comment = core.StringPtr("Approving the changes")
				approveOptionsModel.Complete = core.BoolPtr(false)
				approveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ApproveWithContext(ctx, approveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.Approve(approveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ApproveWithContext(ctx, approveOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(approvePath))
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

					// TODO: Add check for complete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke Approve successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.Approve(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ApproveOptions model
				approveOptionsModel := new(projectv1.ApproveOptions)
				approveOptionsModel.ID = core.StringPtr("testString")
				approveOptionsModel.ConfigID = core.StringPtr("testString")
				approveOptionsModel.Comment = core.StringPtr("Approving the changes")
				approveOptionsModel.Complete = core.BoolPtr(false)
				approveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.Approve(approveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Approve with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ApproveOptions model
				approveOptionsModel := new(projectv1.ApproveOptions)
				approveOptionsModel.ID = core.StringPtr("testString")
				approveOptionsModel.ConfigID = core.StringPtr("testString")
				approveOptionsModel.Comment = core.StringPtr("Approving the changes")
				approveOptionsModel.Complete = core.BoolPtr(false)
				approveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.Approve(approveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ApproveOptions model with no property values
				approveOptionsModelNew := new(projectv1.ApproveOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.Approve(approveOptionsModelNew)
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
			It(`Invoke Approve successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ApproveOptions model
				approveOptionsModel := new(projectv1.ApproveOptions)
				approveOptionsModel.ID = core.StringPtr("testString")
				approveOptionsModel.ConfigID = core.StringPtr("testString")
				approveOptionsModel.Comment = core.StringPtr("Approving the changes")
				approveOptionsModel.Complete = core.BoolPtr(false)
				approveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.Approve(approveOptionsModel)
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
	Describe(`CheckConfig(checkConfigOptions *CheckConfigOptions) - Operation response error`, func() {
		checkConfigPath := "/v1/projects/testString/configs/testString/check"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkConfigPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// TODO: Add check for complete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CheckConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectv1.CheckConfigOptions)
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.ConfigID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Complete = core.BoolPtr(false)
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CheckConfig(checkConfigOptions *CheckConfigOptions)`, func() {
		checkConfigPath := "/v1/projects/testString/configs/testString/check"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkConfigPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// TODO: Add check for complete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke CheckConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectv1.CheckConfigOptions)
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.ConfigID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Complete = core.BoolPtr(false)
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.CheckConfigWithContext(ctx, checkConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.CheckConfigWithContext(ctx, checkConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(checkConfigPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// TODO: Add check for complete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke CheckConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.CheckConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectv1.CheckConfigOptions)
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.ConfigID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Complete = core.BoolPtr(false)
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CheckConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectv1.CheckConfigOptions)
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.ConfigID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Complete = core.BoolPtr(false)
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CheckConfigOptions model with no property values
				checkConfigOptionsModelNew := new(projectv1.CheckConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.CheckConfig(checkConfigOptionsModelNew)
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
			It(`Invoke CheckConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectv1.CheckConfigOptions)
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.ConfigID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Complete = core.BoolPtr(false)
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.CheckConfig(checkConfigOptionsModel)
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
	Describe(`InstallConfig(installConfigOptions *InstallConfigOptions) - Operation response error`, func() {
		installConfigPath := "/v1/projects/testString/configs/testString/install"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(installConfigPath))
					Expect(req.Method).To(Equal("POST"))
					// TODO: Add check for complete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke InstallConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectv1.InstallConfigOptions)
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.ConfigID = core.StringPtr("testString")
				installConfigOptionsModel.Complete = core.BoolPtr(false)
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`InstallConfig(installConfigOptions *InstallConfigOptions)`, func() {
		installConfigPath := "/v1/projects/testString/configs/testString/install"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(installConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for complete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke InstallConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectv1.InstallConfigOptions)
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.ConfigID = core.StringPtr("testString")
				installConfigOptionsModel.Complete = core.BoolPtr(false)
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.InstallConfigWithContext(ctx, installConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.InstallConfigWithContext(ctx, installConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(installConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for complete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "required": true}], "output": [{"name": "Name", "description": "Description", "value": ["Value"]}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke InstallConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.InstallConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectv1.InstallConfigOptions)
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.ConfigID = core.StringPtr("testString")
				installConfigOptionsModel.Complete = core.BoolPtr(false)
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke InstallConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectv1.InstallConfigOptions)
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.ConfigID = core.StringPtr("testString")
				installConfigOptionsModel.Complete = core.BoolPtr(false)
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the InstallConfigOptions model with no property values
				installConfigOptionsModelNew := new(projectv1.InstallConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.InstallConfig(installConfigOptionsModelNew)
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
			It(`Invoke InstallConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectv1.InstallConfigOptions)
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.ConfigID = core.StringPtr("testString")
				installConfigOptionsModel.Complete = core.BoolPtr(false)
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.InstallConfig(installConfigOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectService.UninstallConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UninstallConfigOptions model
				uninstallConfigOptionsModel := new(projectv1.UninstallConfigOptions)
				uninstallConfigOptionsModel.ID = core.StringPtr("testString")
				uninstallConfigOptionsModel.ConfigID = core.StringPtr("testString")
				uninstallConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectService.UninstallConfig(uninstallConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UninstallConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the UninstallConfigOptions model
				uninstallConfigOptionsModel := new(projectv1.UninstallConfigOptions)
				uninstallConfigOptionsModel.ID = core.StringPtr("testString")
				uninstallConfigOptionsModel.ConfigID = core.StringPtr("testString")
				uninstallConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectService.UninstallConfig(uninstallConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UninstallConfigOptions model with no property values
				uninstallConfigOptionsModelNew := new(projectv1.UninstallConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectService.UninstallConfig(uninstallConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchematicsJob(getSchematicsJobOptions *GetSchematicsJobOptions) - Operation response error`, func() {
		getSchematicsJobPath := "/v1/projects/testString/configs/testString/job/plan"
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetSchematicsJobOptions model
				getSchematicsJobOptionsModel := new(projectv1.GetSchematicsJobOptions)
				getSchematicsJobOptionsModel.ID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.ConfigID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
				getSchematicsJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetSchematicsJob(getSchematicsJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetSchematicsJob(getSchematicsJobOptionsModel)
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
		getSchematicsJobPath := "/v1/projects/testString/configs/testString/job/plan"
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetSchematicsJobOptions model
				getSchematicsJobOptionsModel := new(projectv1.GetSchematicsJobOptions)
				getSchematicsJobOptionsModel.ID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.ConfigID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
				getSchematicsJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetSchematicsJobWithContext(ctx, getSchematicsJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetSchematicsJob(getSchematicsJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetSchematicsJobWithContext(ctx, getSchematicsJobOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetSchematicsJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchematicsJobOptions model
				getSchematicsJobOptionsModel := new(projectv1.GetSchematicsJobOptions)
				getSchematicsJobOptionsModel.ID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.ConfigID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
				getSchematicsJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetSchematicsJob(getSchematicsJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSchematicsJob with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetSchematicsJobOptions model
				getSchematicsJobOptionsModel := new(projectv1.GetSchematicsJobOptions)
				getSchematicsJobOptionsModel.ID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.ConfigID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
				getSchematicsJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetSchematicsJob(getSchematicsJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSchematicsJobOptions model with no property values
				getSchematicsJobOptionsModelNew := new(projectv1.GetSchematicsJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetSchematicsJob(getSchematicsJobOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetSchematicsJobOptions model
				getSchematicsJobOptionsModel := new(projectv1.GetSchematicsJobOptions)
				getSchematicsJobOptionsModel.ID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.ConfigID = core.StringPtr("testString")
				getSchematicsJobOptionsModel.Action = core.StringPtr("plan")
				getSchematicsJobOptionsModel.Since = core.Int64Ptr(int64(38))
				getSchematicsJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetSchematicsJob(getSchematicsJobOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigID = core.StringPtr("testString")
				getCostEstimateOptionsModel.Version = core.StringPtr("active")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetCostEstimate(getCostEstimateOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigID = core.StringPtr("testString")
				getCostEstimateOptionsModel.Version = core.StringPtr("active")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetCostEstimateWithContext(ctx, getCostEstimateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetCostEstimateWithContext(ctx, getCostEstimateOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetCostEstimate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigID = core.StringPtr("testString")
				getCostEstimateOptionsModel.Version = core.StringPtr("active")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCostEstimate with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigID = core.StringPtr("testString")
				getCostEstimateOptionsModel.Version = core.StringPtr("active")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetCostEstimate(getCostEstimateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCostEstimateOptions model with no property values
				getCostEstimateOptionsModelNew := new(projectv1.GetCostEstimateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetCostEstimate(getCostEstimateOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetCostEstimateOptions model
				getCostEstimateOptionsModel := new(projectv1.GetCostEstimateOptions)
				getCostEstimateOptionsModel.ID = core.StringPtr("testString")
				getCostEstimateOptionsModel.ConfigID = core.StringPtr("testString")
				getCostEstimateOptionsModel.Version = core.StringPtr("active")
				getCostEstimateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetCostEstimate(getCostEstimateOptionsModel)
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
	Describe(`PostCrnToken(postCrnTokenOptions *PostCrnTokenOptions) - Operation response error`, func() {
		postCrnTokenPath := "/v1/projects/testString/token"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postCrnTokenPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostCrnToken with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PostCrnTokenOptions model
				postCrnTokenOptionsModel := new(projectv1.PostCrnTokenOptions)
				postCrnTokenOptionsModel.ID = core.StringPtr("testString")
				postCrnTokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.PostCrnToken(postCrnTokenOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.PostCrnToken(postCrnTokenOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostCrnToken(postCrnTokenOptions *PostCrnTokenOptions)`, func() {
		postCrnTokenPath := "/v1/projects/testString/token"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postCrnTokenPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"acces_token": "AccesToken", "expiration": 10}`)
				}))
			})
			It(`Invoke PostCrnToken successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the PostCrnTokenOptions model
				postCrnTokenOptionsModel := new(projectv1.PostCrnTokenOptions)
				postCrnTokenOptionsModel.ID = core.StringPtr("testString")
				postCrnTokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.PostCrnTokenWithContext(ctx, postCrnTokenOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.PostCrnToken(postCrnTokenOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.PostCrnTokenWithContext(ctx, postCrnTokenOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postCrnTokenPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"acces_token": "AccesToken", "expiration": 10}`)
				}))
			})
			It(`Invoke PostCrnToken successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.PostCrnToken(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostCrnTokenOptions model
				postCrnTokenOptionsModel := new(projectv1.PostCrnTokenOptions)
				postCrnTokenOptionsModel.ID = core.StringPtr("testString")
				postCrnTokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.PostCrnToken(postCrnTokenOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostCrnToken with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PostCrnTokenOptions model
				postCrnTokenOptionsModel := new(projectv1.PostCrnTokenOptions)
				postCrnTokenOptionsModel.ID = core.StringPtr("testString")
				postCrnTokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.PostCrnToken(postCrnTokenOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostCrnTokenOptions model with no property values
				postCrnTokenOptionsModelNew := new(projectv1.PostCrnTokenOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.PostCrnToken(postCrnTokenOptionsModelNew)
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
			It(`Invoke PostCrnToken successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PostCrnTokenOptions model
				postCrnTokenOptionsModel := new(projectv1.PostCrnTokenOptions)
				postCrnTokenOptionsModel.ID = core.StringPtr("testString")
				postCrnTokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.PostCrnToken(postCrnTokenOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectv1.NotificationEvent)
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.TriggeredBy = core.StringPtr("user-iam-id")
				notificationEventModel.ActionURL = core.StringPtr("actionable/url")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the PostNotificationOptions model
				postNotificationOptionsModel := new(projectv1.PostNotificationOptions)
				postNotificationOptionsModel.ID = core.StringPtr("testString")
				postNotificationOptionsModel.Notifications = []projectv1.NotificationEvent{*notificationEventModel}
				postNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.PostNotification(postNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.PostNotification(postNotificationOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"notifications": [{"event": "Event", "target": "Target", "source": "Source", "triggered_by": "TriggeredBy", "action_url": "ActionURL", "data": {"anyKey": "anyValue"}, "id": "ID", "status": "Status", "reasons": [{"anyKey": "anyValue"}]}]}`)
				}))
			})
			It(`Invoke PostNotification successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectv1.NotificationEvent)
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.TriggeredBy = core.StringPtr("user-iam-id")
				notificationEventModel.ActionURL = core.StringPtr("actionable/url")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the PostNotificationOptions model
				postNotificationOptionsModel := new(projectv1.PostNotificationOptions)
				postNotificationOptionsModel.ID = core.StringPtr("testString")
				postNotificationOptionsModel.Notifications = []projectv1.NotificationEvent{*notificationEventModel}
				postNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.PostNotificationWithContext(ctx, postNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.PostNotification(postNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.PostNotificationWithContext(ctx, postNotificationOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"notifications": [{"event": "Event", "target": "Target", "source": "Source", "triggered_by": "TriggeredBy", "action_url": "ActionURL", "data": {"anyKey": "anyValue"}, "id": "ID", "status": "Status", "reasons": [{"anyKey": "anyValue"}]}]}`)
				}))
			})
			It(`Invoke PostNotification successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.PostNotification(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectv1.NotificationEvent)
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.TriggeredBy = core.StringPtr("user-iam-id")
				notificationEventModel.ActionURL = core.StringPtr("actionable/url")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the PostNotificationOptions model
				postNotificationOptionsModel := new(projectv1.PostNotificationOptions)
				postNotificationOptionsModel.ID = core.StringPtr("testString")
				postNotificationOptionsModel.Notifications = []projectv1.NotificationEvent{*notificationEventModel}
				postNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.PostNotification(postNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostNotification with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectv1.NotificationEvent)
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.TriggeredBy = core.StringPtr("user-iam-id")
				notificationEventModel.ActionURL = core.StringPtr("actionable/url")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the PostNotificationOptions model
				postNotificationOptionsModel := new(projectv1.PostNotificationOptions)
				postNotificationOptionsModel.ID = core.StringPtr("testString")
				postNotificationOptionsModel.Notifications = []projectv1.NotificationEvent{*notificationEventModel}
				postNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.PostNotification(postNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostNotificationOptions model with no property values
				postNotificationOptionsModelNew := new(projectv1.PostNotificationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.PostNotification(postNotificationOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectv1.NotificationEvent)
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.TriggeredBy = core.StringPtr("user-iam-id")
				notificationEventModel.ActionURL = core.StringPtr("actionable/url")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the PostNotificationOptions model
				postNotificationOptionsModel := new(projectv1.PostNotificationOptions)
				postNotificationOptionsModel.ID = core.StringPtr("testString")
				postNotificationOptionsModel.Notifications = []projectv1.NotificationEvent{*notificationEventModel}
				postNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.PostNotification(postNotificationOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetNotificationsOptions model
				getNotificationsOptionsModel := new(projectv1.GetNotificationsOptions)
				getNotificationsOptionsModel.ID = core.StringPtr("testString")
				getNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetNotifications(getNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetNotifications(getNotificationsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"notifications": [{"event": "Event", "target": "Target", "source": "Source", "triggered_by": "TriggeredBy", "action_url": "ActionURL", "data": {"anyKey": "anyValue"}, "id": "ID"}]}`)
				}))
			})
			It(`Invoke GetNotifications successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetNotificationsOptions model
				getNotificationsOptionsModel := new(projectv1.GetNotificationsOptions)
				getNotificationsOptionsModel.ID = core.StringPtr("testString")
				getNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetNotificationsWithContext(ctx, getNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetNotifications(getNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetNotificationsWithContext(ctx, getNotificationsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"notifications": [{"event": "Event", "target": "Target", "source": "Source", "triggered_by": "TriggeredBy", "action_url": "ActionURL", "data": {"anyKey": "anyValue"}, "id": "ID"}]}`)
				}))
			})
			It(`Invoke GetNotifications successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetNotificationsOptions model
				getNotificationsOptionsModel := new(projectv1.GetNotificationsOptions)
				getNotificationsOptionsModel.ID = core.StringPtr("testString")
				getNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetNotifications(getNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetNotifications with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetNotificationsOptions model
				getNotificationsOptionsModel := new(projectv1.GetNotificationsOptions)
				getNotificationsOptionsModel.ID = core.StringPtr("testString")
				getNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetNotifications(getNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetNotificationsOptions model with no property values
				getNotificationsOptionsModelNew := new(projectv1.GetNotificationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetNotifications(getNotificationsOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetNotificationsOptions model
				getNotificationsOptionsModel := new(projectv1.GetNotificationsOptions)
				getNotificationsOptionsModel.ID = core.StringPtr("testString")
				getNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetNotifications(getNotificationsOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectService.ReceivePulsarCatalogEvents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PulsarEventPrototypeCollection model
				pulsarEventPrototypeCollectionModel := new(projectv1.PulsarEventPrototypeCollection)
				pulsarEventPrototypeCollectionModel.EventType = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Timestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				pulsarEventPrototypeCollectionModel.Publisher = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.AccountID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Version = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.EventProperties = map[string]interface{}{"anyKey": "anyValue"}
				pulsarEventPrototypeCollectionModel.EventID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ReceivePulsarCatalogEventsOptions model
				receivePulsarCatalogEventsOptionsModel := new(projectv1.ReceivePulsarCatalogEventsOptions)
				receivePulsarCatalogEventsOptionsModel.PulsarCatalogEvents = []projectv1.PulsarEventPrototypeCollection{*pulsarEventPrototypeCollectionModel}
				receivePulsarCatalogEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectService.ReceivePulsarCatalogEvents(receivePulsarCatalogEventsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ReceivePulsarCatalogEvents with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PulsarEventPrototypeCollection model
				pulsarEventPrototypeCollectionModel := new(projectv1.PulsarEventPrototypeCollection)
				pulsarEventPrototypeCollectionModel.EventType = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Timestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				pulsarEventPrototypeCollectionModel.Publisher = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.AccountID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Version = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.EventProperties = map[string]interface{}{"anyKey": "anyValue"}
				pulsarEventPrototypeCollectionModel.EventID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ReceivePulsarCatalogEventsOptions model
				receivePulsarCatalogEventsOptionsModel := new(projectv1.ReceivePulsarCatalogEventsOptions)
				receivePulsarCatalogEventsOptionsModel.PulsarCatalogEvents = []projectv1.PulsarEventPrototypeCollection{*pulsarEventPrototypeCollectionModel}
				receivePulsarCatalogEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectService.ReceivePulsarCatalogEvents(receivePulsarCatalogEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ReceivePulsarCatalogEventsOptions model with no property values
				receivePulsarCatalogEventsOptionsModelNew := new(projectv1.ReceivePulsarCatalogEventsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectService.ReceivePulsarCatalogEvents(receivePulsarCatalogEventsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReceivePulsarEventNotificationEvents(receivePulsarEventNotificationEventsOptions *ReceivePulsarEventNotificationEventsOptions)`, func() {
		receivePulsarEventNotificationEventsPath := "/v1/pulsar/event_notification_events"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(receivePulsarEventNotificationEventsPath))
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
			It(`Invoke ReceivePulsarEventNotificationEvents successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectService.ReceivePulsarEventNotificationEvents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PulsarEventPrototypeCollection model
				pulsarEventPrototypeCollectionModel := new(projectv1.PulsarEventPrototypeCollection)
				pulsarEventPrototypeCollectionModel.EventType = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Timestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				pulsarEventPrototypeCollectionModel.Publisher = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.AccountID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Version = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.EventProperties = map[string]interface{}{"anyKey": "anyValue"}
				pulsarEventPrototypeCollectionModel.EventID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ReceivePulsarEventNotificationEventsOptions model
				receivePulsarEventNotificationEventsOptionsModel := new(projectv1.ReceivePulsarEventNotificationEventsOptions)
				receivePulsarEventNotificationEventsOptionsModel.PulsarEventNotificationEvents = []projectv1.PulsarEventPrototypeCollection{*pulsarEventPrototypeCollectionModel}
				receivePulsarEventNotificationEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectService.ReceivePulsarEventNotificationEvents(receivePulsarEventNotificationEventsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ReceivePulsarEventNotificationEvents with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PulsarEventPrototypeCollection model
				pulsarEventPrototypeCollectionModel := new(projectv1.PulsarEventPrototypeCollection)
				pulsarEventPrototypeCollectionModel.EventType = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Timestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				pulsarEventPrototypeCollectionModel.Publisher = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.AccountID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Version = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.EventProperties = map[string]interface{}{"anyKey": "anyValue"}
				pulsarEventPrototypeCollectionModel.EventID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ReceivePulsarEventNotificationEventsOptions model
				receivePulsarEventNotificationEventsOptionsModel := new(projectv1.ReceivePulsarEventNotificationEventsOptions)
				receivePulsarEventNotificationEventsOptionsModel.PulsarEventNotificationEvents = []projectv1.PulsarEventPrototypeCollection{*pulsarEventPrototypeCollectionModel}
				receivePulsarEventNotificationEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectService.ReceivePulsarEventNotificationEvents(receivePulsarEventNotificationEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ReceivePulsarEventNotificationEventsOptions model with no property values
				receivePulsarEventNotificationEventsOptionsModelNew := new(projectv1.ReceivePulsarEventNotificationEventsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectService.ReceivePulsarEventNotificationEvents(receivePulsarEventNotificationEventsOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(projectv1.GetHealthOptions)
				getHealthOptionsModel.Info = core.BoolPtr(false)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetHealth(getHealthOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(projectv1.GetHealthOptions)
				getHealthOptionsModel.Info = core.BoolPtr(false)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetHealthWithContext(ctx, getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetHealthWithContext(ctx, getHealthOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetHealth(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(projectv1.GetHealthOptions)
				getHealthOptionsModel.Info = core.BoolPtr(false)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetHealth with error: Operation request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(projectv1.GetHealthOptions)
				getHealthOptionsModel.Info = core.BoolPtr(false)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetHealth(getHealthOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(projectv1.GetHealthOptions)
				getHealthOptionsModel.Info = core.BoolPtr(false)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetHealth(getHealthOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(projectv1.ReplaceServiceInstanceOptions)
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
				result, response, operationErr := projectService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(projectv1.ReplaceServiceInstanceOptions)
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
				_, _, operationErr := projectService.ReplaceServiceInstanceWithContext(ctx, replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ReplaceServiceInstanceWithContext(ctx, replaceServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ReplaceServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(projectv1.ReplaceServiceInstanceOptions)
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
				result, response, operationErr = projectService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceServiceInstance with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(projectv1.ReplaceServiceInstanceOptions)
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
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceServiceInstanceOptions model with no property values
				replaceServiceInstanceOptionsModelNew := new(projectv1.ReplaceServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.ReplaceServiceInstance(replaceServiceInstanceOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(projectv1.ReplaceServiceInstanceOptions)
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
				result, response, operationErr := projectService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(projectv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				deleteServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(projectv1.DeleteServiceInstanceOptions)
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
				_, _, operationErr := projectService.DeleteServiceInstanceWithContext(ctx, deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.DeleteServiceInstanceWithContext(ctx, deleteServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.DeleteServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(projectv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				deleteServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteServiceInstance with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(projectv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				deleteServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteServiceInstanceOptions model with no property values
				deleteServiceInstanceOptionsModelNew := new(projectv1.DeleteServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.DeleteServiceInstance(deleteServiceInstanceOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(projectv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				deleteServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				deleteServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(projectv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.JSONPatchOperation = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(projectv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.JSONPatchOperation = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.UpdateServiceInstanceWithContext(ctx, updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.UpdateServiceInstanceWithContext(ctx, updateServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.UpdateServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(projectv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.JSONPatchOperation = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateServiceInstance with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(projectv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.JSONPatchOperation = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateServiceInstanceOptions model with no property values
				updateServiceInstanceOptionsModelNew := new(projectv1.UpdateServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.UpdateServiceInstance(updateServiceInstanceOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(projectv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.JSONPatchOperation = []projectv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity = core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(false)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(projectv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetLastOperation(getLastOperationOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(projectv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetLastOperationWithContext(ctx, getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetLastOperationWithContext(ctx, getLastOperationOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetLastOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(projectv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLastOperation with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(projectv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLastOperationOptions model with no property values
				getLastOperationOptionsModelNew := new(projectv1.GetLastOperationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetLastOperation(getLastOperationOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(projectv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getLastOperationOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getLastOperationOptionsModel.Operation = core.StringPtr("ABCD")
				getLastOperationOptionsModel.PlanID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("cb54391b-3316-4943-a5a6-a541678c1924")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetLastOperation(getLastOperationOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(projectv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(projectv1.ReplaceServiceInstanceStateOptions)
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
				_, _, operationErr := projectService.ReplaceServiceInstanceStateWithContext(ctx, replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ReplaceServiceInstanceStateWithContext(ctx, replaceServiceInstanceStateOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ReplaceServiceInstanceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(projectv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceServiceInstanceState with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(projectv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceServiceInstanceStateOptions model with no property values
				replaceServiceInstanceStateOptionsModelNew := new(projectv1.ReplaceServiceInstanceStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(projectv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(true)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.ReasonCode = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceInstanceStateOptionsModel.PlanID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.PreviousValues = []string{"testString"}
				replaceServiceInstanceStateOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(projectv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetServiceInstance(getServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(projectv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetServiceInstanceWithContext(ctx, getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetServiceInstanceWithContext(ctx, getServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(projectv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetServiceInstance with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(projectv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetServiceInstance(getServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetServiceInstanceOptions model with no property values
				getServiceInstanceOptionsModelNew := new(projectv1.GetServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetServiceInstance(getServiceInstanceOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceOptions model
				getServiceInstanceOptionsModel := new(projectv1.GetServiceInstanceOptions)
				getServiceInstanceOptionsModel.InstanceID = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetServiceInstance(getServiceInstanceOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(projectv1.GetCatalogOptions)
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetCatalog(getCatalogOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(projectv1.GetCatalogOptions)
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetCatalogWithContext(ctx, getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetCatalogWithContext(ctx, getCatalogOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(projectv1.GetCatalogOptions)
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCatalog with error: Operation request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(projectv1.GetCatalogOptions)
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetCatalog(getCatalogOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(projectv1.GetCatalogOptions)
				getCatalogOptionsModel.XBrokerApiVersion = core.StringPtr("1.0")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetCatalog(getCatalogOptionsModel)
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
		postEventNotificationsIntegrationPath := "/v1/projects/testString/event_notifications"
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				postEventNotificationsIntegrationOptionsModel := new(projectv1.PostEventNotificationsIntegrationOptions)
				postEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationsIntegrationOptionsModel.InstanceCrn = core.StringPtr("CRN of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.Description = core.StringPtr("A sample project source.")
				postEventNotificationsIntegrationOptionsModel.EventNotificationsSourceName = core.StringPtr("project 1 source name for event notifications")
				postEventNotificationsIntegrationOptionsModel.Enabled = core.BoolPtr(true)
				postEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
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
		postEventNotificationsIntegrationPath := "/v1/projects/testString/event_notifications"
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				postEventNotificationsIntegrationOptionsModel := new(projectv1.PostEventNotificationsIntegrationOptions)
				postEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationsIntegrationOptionsModel.InstanceCrn = core.StringPtr("CRN of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.Description = core.StringPtr("A sample project source.")
				postEventNotificationsIntegrationOptionsModel.EventNotificationsSourceName = core.StringPtr("project 1 source name for event notifications")
				postEventNotificationsIntegrationOptionsModel.Enabled = core.BoolPtr(true)
				postEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.PostEventNotificationsIntegrationWithContext(ctx, postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.PostEventNotificationsIntegrationWithContext(ctx, postEventNotificationsIntegrationOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.PostEventNotificationsIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				postEventNotificationsIntegrationOptionsModel := new(projectv1.PostEventNotificationsIntegrationOptions)
				postEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationsIntegrationOptionsModel.InstanceCrn = core.StringPtr("CRN of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.Description = core.StringPtr("A sample project source.")
				postEventNotificationsIntegrationOptionsModel.EventNotificationsSourceName = core.StringPtr("project 1 source name for event notifications")
				postEventNotificationsIntegrationOptionsModel.Enabled = core.BoolPtr(true)
				postEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostEventNotificationsIntegration with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				postEventNotificationsIntegrationOptionsModel := new(projectv1.PostEventNotificationsIntegrationOptions)
				postEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationsIntegrationOptionsModel.InstanceCrn = core.StringPtr("CRN of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.Description = core.StringPtr("A sample project source.")
				postEventNotificationsIntegrationOptionsModel.EventNotificationsSourceName = core.StringPtr("project 1 source name for event notifications")
				postEventNotificationsIntegrationOptionsModel.Enabled = core.BoolPtr(true)
				postEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostEventNotificationsIntegrationOptions model with no property values
				postEventNotificationsIntegrationOptionsModelNew := new(projectv1.PostEventNotificationsIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				postEventNotificationsIntegrationOptionsModel := new(projectv1.PostEventNotificationsIntegrationOptions)
				postEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				postEventNotificationsIntegrationOptionsModel.InstanceCrn = core.StringPtr("CRN of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.Description = core.StringPtr("A sample project source.")
				postEventNotificationsIntegrationOptionsModel.EventNotificationsSourceName = core.StringPtr("project 1 source name for event notifications")
				postEventNotificationsIntegrationOptionsModel.Enabled = core.BoolPtr(true)
				postEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptionsModel)
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
		getEventNotificationsIntegrationPath := "/v1/projects/testString/event_notifications"
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				getEventNotificationsIntegrationOptionsModel := new(projectv1.GetEventNotificationsIntegrationOptions)
				getEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				getEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
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
		getEventNotificationsIntegrationPath := "/v1/projects/testString/event_notifications"
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				getEventNotificationsIntegrationOptionsModel := new(projectv1.GetEventNotificationsIntegrationOptions)
				getEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				getEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetEventNotificationsIntegrationWithContext(ctx, getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetEventNotificationsIntegrationWithContext(ctx, getEventNotificationsIntegrationOptionsModel)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetEventNotificationsIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				getEventNotificationsIntegrationOptionsModel := new(projectv1.GetEventNotificationsIntegrationOptions)
				getEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				getEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEventNotificationsIntegration with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				getEventNotificationsIntegrationOptionsModel := new(projectv1.GetEventNotificationsIntegrationOptions)
				getEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				getEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEventNotificationsIntegrationOptions model with no property values
				getEventNotificationsIntegrationOptionsModelNew := new(projectv1.GetEventNotificationsIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModelNew)
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetEventNotificationsIntegrationOptions model
				getEventNotificationsIntegrationOptionsModel := new(projectv1.GetEventNotificationsIntegrationOptions)
				getEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				getEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptionsModel)
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
		deleteEventNotificationsIntegrationPath := "/v1/projects/testString/event_notifications"
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
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectService.DeleteEventNotificationsIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteEventNotificationsIntegrationOptions model
				deleteEventNotificationsIntegrationOptionsModel := new(projectv1.DeleteEventNotificationsIntegrationOptions)
				deleteEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				deleteEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectService.DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteEventNotificationsIntegration with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteEventNotificationsIntegrationOptions model
				deleteEventNotificationsIntegrationOptionsModel := new(projectv1.DeleteEventNotificationsIntegrationOptions)
				deleteEventNotificationsIntegrationOptionsModel.ID = core.StringPtr("testString")
				deleteEventNotificationsIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectService.DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteEventNotificationsIntegrationOptions model with no property values
				deleteEventNotificationsIntegrationOptionsModelNew := new(projectv1.DeleteEventNotificationsIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectService.DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostTestEventNotification(postTestEventNotificationOptions *PostTestEventNotificationOptions) - Operation response error`, func() {
		postTestEventNotificationPath := "/v1/projects/testString/event_notifications/test"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postTestEventNotificationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostTestEventNotification with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PostTestEventNotificationOptions model
				postTestEventNotificationOptionsModel := new(projectv1.PostTestEventNotificationOptions)
				postTestEventNotificationOptionsModel.ID = core.StringPtr("testString")
				postTestEventNotificationOptionsModel.Ibmendefaultlong = core.StringPtr("long test notification message")
				postTestEventNotificationOptionsModel.Ibmendefaultshort = core.StringPtr("Test notification")
				postTestEventNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.PostTestEventNotification(postTestEventNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.PostTestEventNotification(postTestEventNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostTestEventNotification(postTestEventNotificationOptions *PostTestEventNotificationOptions)`, func() {
		postTestEventNotificationPath := "/v1/projects/testString/event_notifications/test"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postTestEventNotificationPath))
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
			It(`Invoke PostTestEventNotification successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the PostTestEventNotificationOptions model
				postTestEventNotificationOptionsModel := new(projectv1.PostTestEventNotificationOptions)
				postTestEventNotificationOptionsModel.ID = core.StringPtr("testString")
				postTestEventNotificationOptionsModel.Ibmendefaultlong = core.StringPtr("long test notification message")
				postTestEventNotificationOptionsModel.Ibmendefaultshort = core.StringPtr("Test notification")
				postTestEventNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.PostTestEventNotificationWithContext(ctx, postTestEventNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.PostTestEventNotification(postTestEventNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.PostTestEventNotificationWithContext(ctx, postTestEventNotificationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postTestEventNotificationPath))
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
			It(`Invoke PostTestEventNotification successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.PostTestEventNotification(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostTestEventNotificationOptions model
				postTestEventNotificationOptionsModel := new(projectv1.PostTestEventNotificationOptions)
				postTestEventNotificationOptionsModel.ID = core.StringPtr("testString")
				postTestEventNotificationOptionsModel.Ibmendefaultlong = core.StringPtr("long test notification message")
				postTestEventNotificationOptionsModel.Ibmendefaultshort = core.StringPtr("Test notification")
				postTestEventNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.PostTestEventNotification(postTestEventNotificationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostTestEventNotification with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PostTestEventNotificationOptions model
				postTestEventNotificationOptionsModel := new(projectv1.PostTestEventNotificationOptions)
				postTestEventNotificationOptionsModel.ID = core.StringPtr("testString")
				postTestEventNotificationOptionsModel.Ibmendefaultlong = core.StringPtr("long test notification message")
				postTestEventNotificationOptionsModel.Ibmendefaultshort = core.StringPtr("Test notification")
				postTestEventNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.PostTestEventNotification(postTestEventNotificationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostTestEventNotificationOptions model with no property values
				postTestEventNotificationOptionsModelNew := new(projectv1.PostTestEventNotificationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.PostTestEventNotification(postTestEventNotificationOptionsModelNew)
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
			It(`Invoke PostTestEventNotification successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the PostTestEventNotificationOptions model
				postTestEventNotificationOptionsModel := new(projectv1.PostTestEventNotificationOptions)
				postTestEventNotificationOptionsModel.ID = core.StringPtr("testString")
				postTestEventNotificationOptionsModel.Ibmendefaultlong = core.StringPtr("long test notification message")
				postTestEventNotificationOptionsModel.Ibmendefaultshort = core.StringPtr("Test notification")
				postTestEventNotificationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.PostTestEventNotification(postTestEventNotificationOptionsModel)
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
			projectService, _ := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
				URL:           "http://projectv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewApproveOptions successfully`, func() {
				// Construct an instance of the ApproveOptions model
				id := "testString"
				configID := "testString"
				approveOptionsModel := projectService.NewApproveOptions(id, configID)
				approveOptionsModel.SetID("testString")
				approveOptionsModel.SetConfigID("testString")
				approveOptionsModel.SetComment("Approving the changes")
				approveOptionsModel.SetComplete(false)
				approveOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(approveOptionsModel).ToNot(BeNil())
				Expect(approveOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(approveOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(approveOptionsModel.Comment).To(Equal(core.StringPtr("Approving the changes")))
				Expect(approveOptionsModel.Complete).To(Equal(core.BoolPtr(false)))
				Expect(approveOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCheckConfigOptions successfully`, func() {
				// Construct an instance of the CheckConfigOptions model
				id := "testString"
				configID := "testString"
				checkConfigOptionsModel := projectService.NewCheckConfigOptions(id, configID)
				checkConfigOptionsModel.SetID("testString")
				checkConfigOptionsModel.SetConfigID("testString")
				checkConfigOptionsModel.SetXAuthRefreshToken("testString")
				checkConfigOptionsModel.SetVersion("active")
				checkConfigOptionsModel.SetComplete(false)
				checkConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(checkConfigOptionsModel).ToNot(BeNil())
				Expect(checkConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(checkConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(checkConfigOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(checkConfigOptionsModel.Version).To(Equal(core.StringPtr("active")))
				Expect(checkConfigOptionsModel.Complete).To(Equal(core.BoolPtr(false)))
				Expect(checkConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateConfigOptions successfully`, func() {
				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				Expect(projectConfigInputVariableModel).ToNot(BeNil())
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				Expect(projectConfigInputVariableModel.Name).To(Equal(core.StringPtr("account_id")))

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				Expect(projectConfigSettingCollectionModel).ToNot(BeNil())
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")
				Expect(projectConfigSettingCollectionModel.Name).To(Equal(core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")))
				Expect(projectConfigSettingCollectionModel.Value).To(Equal(core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")))

				// Construct an instance of the CreateConfigOptions model
				id := "testString"
				createConfigOptionsNewName := "env-stage"
				createConfigOptionsNewLocatorID := "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"
				createConfigOptionsModel := projectService.NewCreateConfigOptions(id, createConfigOptionsNewName, createConfigOptionsNewLocatorID)
				createConfigOptionsModel.SetID("testString")
				createConfigOptionsModel.SetNewName("env-stage")
				createConfigOptionsModel.SetNewLocatorID("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.SetNewID("testString")
				createConfigOptionsModel.SetNewLabels([]string{"env:stage", "governance:test", "build:0"})
				createConfigOptionsModel.SetNewDescription("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.SetNewInput([]projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel})
				createConfigOptionsModel.SetNewSetting([]projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel})
				createConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createConfigOptionsModel).ToNot(BeNil())
				Expect(createConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createConfigOptionsModel.NewName).To(Equal(core.StringPtr("env-stage")))
				Expect(createConfigOptionsModel.NewLocatorID).To(Equal(core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")))
				Expect(createConfigOptionsModel.NewID).To(Equal(core.StringPtr("testString")))
				Expect(createConfigOptionsModel.NewLabels).To(Equal([]string{"env:stage", "governance:test", "build:0"}))
				Expect(createConfigOptionsModel.NewDescription).To(Equal(core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")))
				Expect(createConfigOptionsModel.NewInput).To(Equal([]projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}))
				Expect(createConfigOptionsModel.NewSetting).To(Equal([]projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}))
				Expect(createConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectOptions successfully`, func() {
				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				Expect(projectConfigInputVariableModel).ToNot(BeNil())
				projectConfigInputVariableModel.Name = core.StringPtr("testString")
				Expect(projectConfigInputVariableModel.Name).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				Expect(projectConfigSettingCollectionModel).ToNot(BeNil())
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")
				Expect(projectConfigSettingCollectionModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigSettingCollectionModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				Expect(projectConfigPrototypeModel).ToNot(BeNil())
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
				Expect(projectConfigPrototypeModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigPrototypeModel.Name).To(Equal(core.StringPtr("common-variables")))
				Expect(projectConfigPrototypeModel.Labels).To(Equal([]string{"testString"}))
				Expect(projectConfigPrototypeModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigPrototypeModel.LocatorID).To(Equal(core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")))
				Expect(projectConfigPrototypeModel.Input).To(Equal([]projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}))
				Expect(projectConfigPrototypeModel.Setting).To(Equal([]projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}))

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsName := "acme-microservice"
				createProjectOptionsModel := projectService.NewCreateProjectOptions(createProjectOptionsName)
				createProjectOptionsModel.SetName("acme-microservice")
				createProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.SetConfigs([]projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel})
				createProjectOptionsModel.SetResourceGroup("Default")
				createProjectOptionsModel.SetLocation("us-south")
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(createProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure.")))
				Expect(createProjectOptionsModel.Configs).To(Equal([]projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}))
				Expect(createProjectOptionsModel.ResourceGroup).To(Equal(core.StringPtr("Default")))
				Expect(createProjectOptionsModel.Location).To(Equal(core.StringPtr("us-south")))
				Expect(createProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteConfigOptions successfully`, func() {
				// Construct an instance of the DeleteConfigOptions model
				id := "testString"
				configID := "testString"
				deleteConfigOptionsModel := projectService.NewDeleteConfigOptions(id, configID)
				deleteConfigOptionsModel.SetID("testString")
				deleteConfigOptionsModel.SetConfigID("testString")
				deleteConfigOptionsModel.SetDraftOnly(false)
				deleteConfigOptionsModel.SetDestroy(false)
				deleteConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteConfigOptionsModel).ToNot(BeNil())
				Expect(deleteConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigOptionsModel.DraftOnly).To(Equal(core.BoolPtr(false)))
				Expect(deleteConfigOptionsModel.Destroy).To(Equal(core.BoolPtr(false)))
				Expect(deleteConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEventNotificationsIntegrationOptions successfully`, func() {
				// Construct an instance of the DeleteEventNotificationsIntegrationOptions model
				id := "testString"
				deleteEventNotificationsIntegrationOptionsModel := projectService.NewDeleteEventNotificationsIntegrationOptions(id)
				deleteEventNotificationsIntegrationOptionsModel.SetID("testString")
				deleteEventNotificationsIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEventNotificationsIntegrationOptionsModel).ToNot(BeNil())
				Expect(deleteEventNotificationsIntegrationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEventNotificationsIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectOptions successfully`, func() {
				// Construct an instance of the DeleteProjectOptions model
				id := "testString"
				deleteProjectOptionsModel := projectService.NewDeleteProjectOptions(id)
				deleteProjectOptionsModel.SetID("testString")
				deleteProjectOptionsModel.SetDestroy(false)
				deleteProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectOptionsModel).ToNot(BeNil())
				Expect(deleteProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectOptionsModel.Destroy).To(Equal(core.BoolPtr(false)))
				Expect(deleteProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteServiceInstanceOptions successfully`, func() {
				// Construct an instance of the DeleteServiceInstanceOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				planID := "cb54391b-3316-4943-a5a6-a541678c1924"
				serviceID := "cb54391b-3316-4943-a5a6-a541678c1924"
				deleteServiceInstanceOptionsModel := projectService.NewDeleteServiceInstanceOptions(instanceID, planID, serviceID)
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
			It(`Invoke NewForceApproveOptions successfully`, func() {
				// Construct an instance of the ForceApproveOptions model
				id := "testString"
				configID := "testString"
				forceApproveOptionsModel := projectService.NewForceApproveOptions(id, configID)
				forceApproveOptionsModel.SetID("testString")
				forceApproveOptionsModel.SetConfigID("testString")
				forceApproveOptionsModel.SetComment("Approving the changes")
				forceApproveOptionsModel.SetComplete(false)
				forceApproveOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(forceApproveOptionsModel).ToNot(BeNil())
				Expect(forceApproveOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(forceApproveOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(forceApproveOptionsModel.Comment).To(Equal(core.StringPtr("Approving the changes")))
				Expect(forceApproveOptionsModel.Complete).To(Equal(core.BoolPtr(false)))
				Expect(forceApproveOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogOptions successfully`, func() {
				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := projectService.NewGetCatalogOptions()
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
				getConfigDiffOptionsModel := projectService.NewGetConfigDiffOptions(id, configID)
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
				getConfigOptionsModel := projectService.NewGetConfigOptions(id, configID)
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
				getCostEstimateOptionsModel := projectService.NewGetCostEstimateOptions(id, configID)
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
				getEventNotificationsIntegrationOptionsModel := projectService.NewGetEventNotificationsIntegrationOptions(id)
				getEventNotificationsIntegrationOptionsModel.SetID("testString")
				getEventNotificationsIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEventNotificationsIntegrationOptionsModel).ToNot(BeNil())
				Expect(getEventNotificationsIntegrationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getEventNotificationsIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHealthOptions successfully`, func() {
				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := projectService.NewGetHealthOptions()
				getHealthOptionsModel.SetInfo(false)
				getHealthOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHealthOptionsModel).ToNot(BeNil())
				Expect(getHealthOptionsModel.Info).To(Equal(core.BoolPtr(false)))
				Expect(getHealthOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLastOperationOptions successfully`, func() {
				// Construct an instance of the GetLastOperationOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				getLastOperationOptionsModel := projectService.NewGetLastOperationOptions(instanceID)
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
				getNotificationsOptionsModel := projectService.NewGetNotificationsOptions(id)
				getNotificationsOptionsModel.SetID("testString")
				getNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getNotificationsOptionsModel).ToNot(BeNil())
				Expect(getNotificationsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectOptions successfully`, func() {
				// Construct an instance of the GetProjectOptions model
				id := "testString"
				getProjectOptionsModel := projectService.NewGetProjectOptions(id)
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
				getSchematicsJobOptionsModel := projectService.NewGetSchematicsJobOptions(id, configID, action)
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
				getServiceInstanceOptionsModel := projectService.NewGetServiceInstanceOptions(instanceID)
				getServiceInstanceOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				getServiceInstanceOptionsModel.SetXBrokerApiVersion("1.0")
				getServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(getServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(getServiceInstanceOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1.0")))
				Expect(getServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInstallConfigOptions successfully`, func() {
				// Construct an instance of the InstallConfigOptions model
				id := "testString"
				configID := "testString"
				installConfigOptionsModel := projectService.NewInstallConfigOptions(id, configID)
				installConfigOptionsModel.SetID("testString")
				installConfigOptionsModel.SetConfigID("testString")
				installConfigOptionsModel.SetComplete(false)
				installConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(installConfigOptionsModel).ToNot(BeNil())
				Expect(installConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(installConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(installConfigOptionsModel.Complete).To(Equal(core.BoolPtr(false)))
				Expect(installConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewJSONPatchOperation successfully`, func() {
				op := "add"
				path := "testString"
				_model, err := projectService.NewJSONPatchOperation(op, path)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListConfigsOptions successfully`, func() {
				// Construct an instance of the ListConfigsOptions model
				id := "testString"
				listConfigsOptionsModel := projectService.NewListConfigsOptions(id)
				listConfigsOptionsModel.SetID("testString")
				listConfigsOptionsModel.SetVersion("active")
				listConfigsOptionsModel.SetComplete(false)
				listConfigsOptionsModel.SetProjectID("testString")
				listConfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigsOptionsModel).ToNot(BeNil())
				Expect(listConfigsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.Version).To(Equal(core.StringPtr("active")))
				Expect(listConfigsOptionsModel.Complete).To(Equal(core.BoolPtr(false)))
				Expect(listConfigsOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectsOptions successfully`, func() {
				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := projectService.NewListProjectsOptions()
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
				_model, err := projectService.NewNotificationEvent(event, target)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPostCrnTokenOptions successfully`, func() {
				// Construct an instance of the PostCrnTokenOptions model
				id := "testString"
				postCrnTokenOptionsModel := projectService.NewPostCrnTokenOptions(id)
				postCrnTokenOptionsModel.SetID("testString")
				postCrnTokenOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postCrnTokenOptionsModel).ToNot(BeNil())
				Expect(postCrnTokenOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(postCrnTokenOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostEventNotificationsIntegrationOptions successfully`, func() {
				// Construct an instance of the PostEventNotificationsIntegrationOptions model
				id := "testString"
				postEventNotificationsIntegrationOptionsInstanceCrn := "CRN of event notifications instance"
				postEventNotificationsIntegrationOptionsModel := projectService.NewPostEventNotificationsIntegrationOptions(id, postEventNotificationsIntegrationOptionsInstanceCrn)
				postEventNotificationsIntegrationOptionsModel.SetID("testString")
				postEventNotificationsIntegrationOptionsModel.SetInstanceCrn("CRN of event notifications instance")
				postEventNotificationsIntegrationOptionsModel.SetDescription("A sample project source.")
				postEventNotificationsIntegrationOptionsModel.SetEventNotificationsSourceName("project 1 source name for event notifications")
				postEventNotificationsIntegrationOptionsModel.SetEnabled(true)
				postEventNotificationsIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postEventNotificationsIntegrationOptionsModel).ToNot(BeNil())
				Expect(postEventNotificationsIntegrationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(postEventNotificationsIntegrationOptionsModel.InstanceCrn).To(Equal(core.StringPtr("CRN of event notifications instance")))
				Expect(postEventNotificationsIntegrationOptionsModel.Description).To(Equal(core.StringPtr("A sample project source.")))
				Expect(postEventNotificationsIntegrationOptionsModel.EventNotificationsSourceName).To(Equal(core.StringPtr("project 1 source name for event notifications")))
				Expect(postEventNotificationsIntegrationOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(postEventNotificationsIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostNotificationOptions successfully`, func() {
				// Construct an instance of the NotificationEvent model
				notificationEventModel := new(projectv1.NotificationEvent)
				Expect(notificationEventModel).ToNot(BeNil())
				notificationEventModel.Event = core.StringPtr("project.create.failed")
				notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
				notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
				notificationEventModel.TriggeredBy = core.StringPtr("user-iam-id")
				notificationEventModel.ActionURL = core.StringPtr("actionable/url")
				notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				Expect(notificationEventModel.Event).To(Equal(core.StringPtr("project.create.failed")))
				Expect(notificationEventModel.Target).To(Equal(core.StringPtr("234234324-3444-4556-224232432")))
				Expect(notificationEventModel.Source).To(Equal(core.StringPtr("id.of.project.service.instance")))
				Expect(notificationEventModel.TriggeredBy).To(Equal(core.StringPtr("user-iam-id")))
				Expect(notificationEventModel.ActionURL).To(Equal(core.StringPtr("actionable/url")))
				Expect(notificationEventModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the PostNotificationOptions model
				id := "testString"
				postNotificationOptionsModel := projectService.NewPostNotificationOptions(id)
				postNotificationOptionsModel.SetID("testString")
				postNotificationOptionsModel.SetNotifications([]projectv1.NotificationEvent{*notificationEventModel})
				postNotificationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postNotificationOptionsModel).ToNot(BeNil())
				Expect(postNotificationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(postNotificationOptionsModel.Notifications).To(Equal([]projectv1.NotificationEvent{*notificationEventModel}))
				Expect(postNotificationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostTestEventNotificationOptions successfully`, func() {
				// Construct an instance of the PostTestEventNotificationOptions model
				id := "testString"
				postTestEventNotificationOptionsModel := projectService.NewPostTestEventNotificationOptions(id)
				postTestEventNotificationOptionsModel.SetID("testString")
				postTestEventNotificationOptionsModel.SetIbmendefaultlong("long test notification message")
				postTestEventNotificationOptionsModel.SetIbmendefaultshort("Test notification")
				postTestEventNotificationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postTestEventNotificationOptionsModel).ToNot(BeNil())
				Expect(postTestEventNotificationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(postTestEventNotificationOptionsModel.Ibmendefaultlong).To(Equal(core.StringPtr("long test notification message")))
				Expect(postTestEventNotificationOptionsModel.Ibmendefaultshort).To(Equal(core.StringPtr("Test notification")))
				Expect(postTestEventNotificationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewProjectPatch successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectv1.OutputValue)
				outputValueModel.Name = core.StringPtr("testString")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"testString"}

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfig model
				projectConfigModel := new(projectv1.ProjectConfig)
				projectConfigModel.ID = core.StringPtr("testString")
				projectConfigModel.Name = core.StringPtr("testString")
				projectConfigModel.Labels = []string{"testString"}
				projectConfigModel.Description = core.StringPtr("testString")
				projectConfigModel.LocatorID = core.StringPtr("testString")
				projectConfigModel.Type = core.StringPtr("terraform_template")
				projectConfigModel.Input = []projectv1.InputVariable{*inputVariableModel}
				projectConfigModel.Output = []projectv1.OutputValue{*outputValueModel}
				projectConfigModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CumulativeNeedsAttention model
				cumulativeNeedsAttentionModel := new(projectv1.CumulativeNeedsAttention)
				cumulativeNeedsAttentionModel.Event = core.StringPtr("testString")
				cumulativeNeedsAttentionModel.EventID = core.StringPtr("testString")
				cumulativeNeedsAttentionModel.ConfigID = core.StringPtr("testString")
				cumulativeNeedsAttentionModel.ConfigVersion = core.Int64Ptr(int64(38))

				// Construct an instance of the ProjectMetadata model
				projectMetadataModel := new(projectv1.ProjectMetadata)
				projectMetadataModel.Crn = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				projectMetadataModel.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				projectMetadataModel.CumulativeNeedsAttentionView = []projectv1.CumulativeNeedsAttention{*cumulativeNeedsAttentionModel}
				projectMetadataModel.CumulativeNeedsAttentionViewErr = core.StringPtr("testString")
				projectMetadataModel.Location = core.StringPtr("testString")
				projectMetadataModel.ResourceGroup = core.StringPtr("testString")
				projectMetadataModel.State = core.StringPtr("testString")
				projectMetadataModel.EventNotificationsCrn = core.StringPtr("testString")

				// Construct an instance of the Project model
				project := new(projectv1.Project)
				project.Name = core.StringPtr("testString")
				project.Description = core.StringPtr("testString")
				project.ID = core.StringPtr("testString")
				project.Crn = core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				project.Configs = []projectv1.ProjectConfig{*projectConfigModel}
				project.Metadata = projectMetadataModel

				projectPatch := projectService.NewProjectPatch(project)
				Expect(projectPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(projectv1.JSONPatchOperation).Path
				}
				Expect(projectPatch).To(MatchAllElements(_path, Elements{
				"/name": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/name")),
					"From": BeNil(),
					"Value": Equal(project.Name),
					}),
				"/description": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/description")),
					"From": BeNil(),
					"Value": Equal(project.Description),
					}),
				"/id": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/id")),
					"From": BeNil(),
					"Value": Equal(project.ID),
					}),
				"/crn": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/crn")),
					"From": BeNil(),
					"Value": Equal(project.Crn),
					}),
				"/configs": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/configs")),
					"From": BeNil(),
					"Value": Equal(project.Configs),
					}),
				"/metadata": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/metadata")),
					"From": BeNil(),
					"Value": Equal(project.Metadata),
					}),
				}))
			})
			It(`Invoke NewProjectConfigPatch successfully`, func() {
				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.Name = core.StringPtr("testString")
				inputVariableModel.Type = core.StringPtr("array")
				inputVariableModel.Required = core.BoolPtr(true)

				// Construct an instance of the OutputValue model
				outputValueModel := new(projectv1.OutputValue)
				outputValueModel.Name = core.StringPtr("testString")
				outputValueModel.Description = core.StringPtr("testString")
				outputValueModel.Value = []string{"testString"}

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfig model
				projectConfig := new(projectv1.ProjectConfig)
				projectConfig.ID = core.StringPtr("testString")
				projectConfig.Name = core.StringPtr("testString")
				projectConfig.Labels = []string{"testString"}
				projectConfig.Description = core.StringPtr("testString")
				projectConfig.LocatorID = core.StringPtr("testString")
				projectConfig.Type = core.StringPtr("terraform_template")
				projectConfig.Input = []projectv1.InputVariable{*inputVariableModel}
				projectConfig.Output = []projectv1.OutputValue{*outputValueModel}
				projectConfig.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				projectConfigPatch := projectService.NewProjectConfigPatch(projectConfig)
				Expect(projectConfigPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(projectv1.JSONPatchOperation).Path
				}
				Expect(projectConfigPatch).To(MatchAllElements(_path, Elements{
				"/id": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/id")),
					"From": BeNil(),
					"Value": Equal(projectConfig.ID),
					}),
				"/name": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/name")),
					"From": BeNil(),
					"Value": Equal(projectConfig.Name),
					}),
				"/labels": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/labels")),
					"From": BeNil(),
					"Value": Equal(projectConfig.Labels),
					}),
				"/description": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/description")),
					"From": BeNil(),
					"Value": Equal(projectConfig.Description),
					}),
				"/locator_id": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/locator_id")),
					"From": BeNil(),
					"Value": Equal(projectConfig.LocatorID),
					}),
				"/type": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/type")),
					"From": BeNil(),
					"Value": Equal(projectConfig.Type),
					}),
				"/input": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/input")),
					"From": BeNil(),
					"Value": Equal(projectConfig.Input),
					}),
				"/output": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/output")),
					"From": BeNil(),
					"Value": Equal(projectConfig.Output),
					}),
				"/setting": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/setting")),
					"From": BeNil(),
					"Value": Equal(projectConfig.Setting),
					}),
				}))
			})
			It(`Invoke NewProjectConfigInputVariable successfully`, func() {
				name := "testString"
				_model, err := projectService.NewProjectConfigInputVariable(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProjectConfigPrototype successfully`, func() {
				name := "testString"
				locatorID := "testString"
				_model, err := projectService.NewProjectConfigPrototype(name, locatorID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProjectConfigSettingCollection successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := projectService.NewProjectConfigSettingCollection(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPulsarEventPrototypeCollection successfully`, func() {
				eventType := "testString"
				timestamp := CreateMockDateTime("2019-01-01T12:00:00.000Z")
				publisher := "testString"
				accountID := "testString"
				version := "testString"
				_model, err := projectService.NewPulsarEventPrototypeCollection(eventType, timestamp, publisher, accountID, version)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReceivePulsarCatalogEventsOptions successfully`, func() {
				// Construct an instance of the PulsarEventPrototypeCollection model
				pulsarEventPrototypeCollectionModel := new(projectv1.PulsarEventPrototypeCollection)
				Expect(pulsarEventPrototypeCollectionModel).ToNot(BeNil())
				pulsarEventPrototypeCollectionModel.EventType = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Timestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				pulsarEventPrototypeCollectionModel.Publisher = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.AccountID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Version = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.EventProperties = map[string]interface{}{"anyKey": "anyValue"}
				pulsarEventPrototypeCollectionModel.EventID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(pulsarEventPrototypeCollectionModel.EventType).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventPrototypeCollectionModel.Timestamp).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(pulsarEventPrototypeCollectionModel.Publisher).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventPrototypeCollectionModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventPrototypeCollectionModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventPrototypeCollectionModel.EventProperties).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(pulsarEventPrototypeCollectionModel.EventID).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventPrototypeCollectionModel.GetProperties()).ToNot(BeEmpty())
				Expect(pulsarEventPrototypeCollectionModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				pulsarEventPrototypeCollectionModel.SetProperties(nil)
				Expect(pulsarEventPrototypeCollectionModel.GetProperties()).To(BeEmpty())

				pulsarEventPrototypeCollectionModelExpectedMap := make(map[string]interface{})
				pulsarEventPrototypeCollectionModelExpectedMap["foo"] = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.SetProperties(pulsarEventPrototypeCollectionModelExpectedMap)
				pulsarEventPrototypeCollectionModelActualMap := pulsarEventPrototypeCollectionModel.GetProperties()
				Expect(pulsarEventPrototypeCollectionModelActualMap).To(Equal(pulsarEventPrototypeCollectionModelExpectedMap))

				// Construct an instance of the ReceivePulsarCatalogEventsOptions model
				pulsarCatalogEvents := []projectv1.PulsarEventPrototypeCollection{}
				receivePulsarCatalogEventsOptionsModel := projectService.NewReceivePulsarCatalogEventsOptions(pulsarCatalogEvents)
				receivePulsarCatalogEventsOptionsModel.SetPulsarCatalogEvents([]projectv1.PulsarEventPrototypeCollection{*pulsarEventPrototypeCollectionModel})
				receivePulsarCatalogEventsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(receivePulsarCatalogEventsOptionsModel).ToNot(BeNil())
				Expect(receivePulsarCatalogEventsOptionsModel.PulsarCatalogEvents).To(Equal([]projectv1.PulsarEventPrototypeCollection{*pulsarEventPrototypeCollectionModel}))
				Expect(receivePulsarCatalogEventsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReceivePulsarEventNotificationEventsOptions successfully`, func() {
				// Construct an instance of the PulsarEventPrototypeCollection model
				pulsarEventPrototypeCollectionModel := new(projectv1.PulsarEventPrototypeCollection)
				Expect(pulsarEventPrototypeCollectionModel).ToNot(BeNil())
				pulsarEventPrototypeCollectionModel.EventType = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Timestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				pulsarEventPrototypeCollectionModel.Publisher = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.AccountID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.Version = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.EventProperties = map[string]interface{}{"anyKey": "anyValue"}
				pulsarEventPrototypeCollectionModel.EventID = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(pulsarEventPrototypeCollectionModel.EventType).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventPrototypeCollectionModel.Timestamp).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(pulsarEventPrototypeCollectionModel.Publisher).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventPrototypeCollectionModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventPrototypeCollectionModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventPrototypeCollectionModel.EventProperties).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(pulsarEventPrototypeCollectionModel.EventID).To(Equal(core.StringPtr("testString")))
				Expect(pulsarEventPrototypeCollectionModel.GetProperties()).ToNot(BeEmpty())
				Expect(pulsarEventPrototypeCollectionModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				pulsarEventPrototypeCollectionModel.SetProperties(nil)
				Expect(pulsarEventPrototypeCollectionModel.GetProperties()).To(BeEmpty())

				pulsarEventPrototypeCollectionModelExpectedMap := make(map[string]interface{})
				pulsarEventPrototypeCollectionModelExpectedMap["foo"] = core.StringPtr("testString")
				pulsarEventPrototypeCollectionModel.SetProperties(pulsarEventPrototypeCollectionModelExpectedMap)
				pulsarEventPrototypeCollectionModelActualMap := pulsarEventPrototypeCollectionModel.GetProperties()
				Expect(pulsarEventPrototypeCollectionModelActualMap).To(Equal(pulsarEventPrototypeCollectionModelExpectedMap))

				// Construct an instance of the ReceivePulsarEventNotificationEventsOptions model
				pulsarEventNotificationEvents := []projectv1.PulsarEventPrototypeCollection{}
				receivePulsarEventNotificationEventsOptionsModel := projectService.NewReceivePulsarEventNotificationEventsOptions(pulsarEventNotificationEvents)
				receivePulsarEventNotificationEventsOptionsModel.SetPulsarEventNotificationEvents([]projectv1.PulsarEventPrototypeCollection{*pulsarEventPrototypeCollectionModel})
				receivePulsarEventNotificationEventsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(receivePulsarEventNotificationEventsOptionsModel).ToNot(BeNil())
				Expect(receivePulsarEventNotificationEventsOptionsModel.PulsarEventNotificationEvents).To(Equal([]projectv1.PulsarEventPrototypeCollection{*pulsarEventPrototypeCollectionModel}))
				Expect(receivePulsarEventNotificationEventsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceServiceInstanceOptions successfully`, func() {
				// Construct an instance of the ReplaceServiceInstanceOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				replaceServiceInstanceOptionsServiceID := "testString"
				replaceServiceInstanceOptionsPlanID := "testString"
				replaceServiceInstanceOptionsModel := projectService.NewReplaceServiceInstanceOptions(instanceID, replaceServiceInstanceOptionsServiceID, replaceServiceInstanceOptionsPlanID)
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
				replaceServiceInstanceStateOptionsModel := projectService.NewReplaceServiceInstanceStateOptions(instanceID, replaceServiceInstanceStateOptionsEnabled)
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
			It(`Invoke NewResourceUpdateResultPatch successfully`, func() {
				// Construct an instance of the ResourceUpdateResult model
				resourceUpdateResult := new(projectv1.ResourceUpdateResult)
				resourceUpdateResult.SetProperty("foo", core.StringPtr("testString"))

				resourceUpdateResultPatch := projectService.NewResourceUpdateResultPatch(resourceUpdateResult)
				Expect(resourceUpdateResultPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(projectv1.JSONPatchOperation).Path
				}
				Expect(resourceUpdateResultPatch).To(MatchAllElements(_path, Elements{
				"/foo": MatchAllFields(Fields{
					"Op": PointTo(Equal(projectv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/foo")),
					"From": BeNil(),
					"Value": Equal(resourceUpdateResult.GetProperty("foo")),
					}),
				}))
			})
			It(`Invoke NewUninstallConfigOptions successfully`, func() {
				// Construct an instance of the UninstallConfigOptions model
				id := "testString"
				configID := "testString"
				uninstallConfigOptionsModel := projectService.NewUninstallConfigOptions(id, configID)
				uninstallConfigOptionsModel.SetID("testString")
				uninstallConfigOptionsModel.SetConfigID("testString")
				uninstallConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uninstallConfigOptionsModel).ToNot(BeNil())
				Expect(uninstallConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(uninstallConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(uninstallConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateConfigOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateConfigOptions model
				id := "testString"
				configID := "testString"
				projectConfig := []projectv1.JSONPatchOperation{}
				updateConfigOptionsModel := projectService.NewUpdateConfigOptions(id, configID, projectConfig)
				updateConfigOptionsModel.SetID("testString")
				updateConfigOptionsModel.SetConfigID("testString")
				updateConfigOptionsModel.SetProjectConfig([]projectv1.JSONPatchOperation{*jsonPatchOperationModel})
				updateConfigOptionsModel.SetComplete(false)
				updateConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateConfigOptionsModel).ToNot(BeNil())
				Expect(updateConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigOptionsModel.ConfigID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigOptionsModel.ProjectConfig).To(Equal([]projectv1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateConfigOptionsModel.Complete).To(Equal(core.BoolPtr(false)))
				Expect(updateConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateProjectOptions model
				id := "testString"
				jsonPatchOperation := []projectv1.JSONPatchOperation{}
				updateProjectOptionsModel := projectService.NewUpdateProjectOptions(id, jsonPatchOperation)
				updateProjectOptionsModel.SetID("testString")
				updateProjectOptionsModel.SetJSONPatchOperation([]projectv1.JSONPatchOperation{*jsonPatchOperationModel})
				updateProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectOptionsModel).ToNot(BeNil())
				Expect(updateProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectOptionsModel.JSONPatchOperation).To(Equal([]projectv1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateServiceInstanceOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(projectv1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateServiceInstanceOptions model
				instanceID := "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"
				jsonPatchOperation := []projectv1.JSONPatchOperation{}
				updateServiceInstanceOptionsModel := projectService.NewUpdateServiceInstanceOptions(instanceID, jsonPatchOperation)
				updateServiceInstanceOptionsModel.SetInstanceID("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")
				updateServiceInstanceOptionsModel.SetJSONPatchOperation([]projectv1.JSONPatchOperation{*jsonPatchOperationModel})
				updateServiceInstanceOptionsModel.SetXBrokerApiVersion("1.0")
				updateServiceInstanceOptionsModel.SetXBrokerApiOriginatingIdentity("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
				updateServiceInstanceOptionsModel.SetAcceptsIncomplete(false)
				updateServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(updateServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")))
				Expect(updateServiceInstanceOptionsModel.JSONPatchOperation).To(Equal([]projectv1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateServiceInstanceOptionsModel.XBrokerApiVersion).To(Equal(core.StringPtr("1.0")))
				Expect(updateServiceInstanceOptionsModel.XBrokerApiOriginatingIdentity).To(Equal(core.StringPtr("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")))
				Expect(updateServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(false)))
				Expect(updateServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
