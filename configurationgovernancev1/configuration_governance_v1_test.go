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

package configurationgovernancev1_test

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
	"github.com/IBM/platform-services-go-sdk/configurationgovernancev1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ConfigurationGovernanceV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(configurationGovernanceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(configurationGovernanceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
				URL: "https://configurationgovernancev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(configurationGovernanceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIGURATION_GOVERNANCE_URL":       "https://configurationgovernancev1/api",
				"CONFIGURATION_GOVERNANCE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(&configurationgovernancev1.ConfigurationGovernanceV1Options{})
				Expect(configurationGovernanceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := configurationGovernanceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configurationGovernanceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configurationGovernanceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configurationGovernanceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL: "https://testService/api",
				})
				Expect(configurationGovernanceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := configurationGovernanceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configurationGovernanceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configurationGovernanceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configurationGovernanceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(&configurationgovernancev1.ConfigurationGovernanceV1Options{})
				err := configurationGovernanceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := configurationGovernanceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configurationGovernanceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configurationGovernanceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configurationGovernanceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIGURATION_GOVERNANCE_URL":       "https://configurationgovernancev1/api",
				"CONFIGURATION_GOVERNANCE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(&configurationgovernancev1.ConfigurationGovernanceV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(configurationGovernanceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIGURATION_GOVERNANCE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(&configurationgovernancev1.ConfigurationGovernanceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(configurationGovernanceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = configurationgovernancev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateRules(createRulesOptions *CreateRulesOptions) - Operation response error`, func() {
		createRulesPath := "/config/v1/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRulesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRules with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleTargetAttribute model
				ruleTargetAttributeModel := new(configurationgovernancev1.RuleTargetAttribute)
				ruleTargetAttributeModel.Name = core.StringPtr("resource_id")
				ruleTargetAttributeModel.Operator = core.StringPtr("string_equals")
				ruleTargetAttributeModel.Value = core.StringPtr("f0f8f7994e754ff38f9d370201966561")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("disallow")

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"testString"}

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsModel := new(configurationgovernancev1.CreateRulesOptions)
				createRulesOptionsModel.Rules = []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}
				createRulesOptionsModel.TransactionID = core.StringPtr("testString")
				createRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.CreateRules(createRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.CreateRules(createRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateRules(createRulesOptions *CreateRulesOptions)`, func() {
		createRulesPath := "/config/v1/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRulesPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"rules": [{"request_id": "3cebc877-58e7-44a5-a292-32114fa73558", "status_code": 201, "rule": {"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "iam-groups", "resource_kind": "zone", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modification_date": "2019-01-01T12:00:00", "modified_by": "ModifiedBy", "number_of_attachments": 3}, "errors": [{"code": "bad_request", "message": "The rule is missing an account ID"}], "trace": "861263b4-cee3-4514-8d8c-05d17308e6eb"}]}`)
				}))
			})
			It(`Invoke CreateRules successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the RuleTargetAttribute model
				ruleTargetAttributeModel := new(configurationgovernancev1.RuleTargetAttribute)
				ruleTargetAttributeModel.Name = core.StringPtr("resource_id")
				ruleTargetAttributeModel.Operator = core.StringPtr("string_equals")
				ruleTargetAttributeModel.Value = core.StringPtr("f0f8f7994e754ff38f9d370201966561")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("disallow")

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"testString"}

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsModel := new(configurationgovernancev1.CreateRulesOptions)
				createRulesOptionsModel.Rules = []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}
				createRulesOptionsModel.TransactionID = core.StringPtr("testString")
				createRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.CreateRulesWithContext(ctx, createRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.CreateRules(createRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.CreateRulesWithContext(ctx, createRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createRulesPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"rules": [{"request_id": "3cebc877-58e7-44a5-a292-32114fa73558", "status_code": 201, "rule": {"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "iam-groups", "resource_kind": "zone", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modification_date": "2019-01-01T12:00:00", "modified_by": "ModifiedBy", "number_of_attachments": 3}, "errors": [{"code": "bad_request", "message": "The rule is missing an account ID"}], "trace": "861263b4-cee3-4514-8d8c-05d17308e6eb"}]}`)
				}))
			})
			It(`Invoke CreateRules successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.CreateRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleTargetAttribute model
				ruleTargetAttributeModel := new(configurationgovernancev1.RuleTargetAttribute)
				ruleTargetAttributeModel.Name = core.StringPtr("resource_id")
				ruleTargetAttributeModel.Operator = core.StringPtr("string_equals")
				ruleTargetAttributeModel.Value = core.StringPtr("f0f8f7994e754ff38f9d370201966561")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("disallow")

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"testString"}

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsModel := new(configurationgovernancev1.CreateRulesOptions)
				createRulesOptionsModel.Rules = []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}
				createRulesOptionsModel.TransactionID = core.StringPtr("testString")
				createRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.CreateRules(createRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRules with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleTargetAttribute model
				ruleTargetAttributeModel := new(configurationgovernancev1.RuleTargetAttribute)
				ruleTargetAttributeModel.Name = core.StringPtr("resource_id")
				ruleTargetAttributeModel.Operator = core.StringPtr("string_equals")
				ruleTargetAttributeModel.Value = core.StringPtr("f0f8f7994e754ff38f9d370201966561")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("disallow")

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"testString"}

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsModel := new(configurationgovernancev1.CreateRulesOptions)
				createRulesOptionsModel.Rules = []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}
				createRulesOptionsModel.TransactionID = core.StringPtr("testString")
				createRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.CreateRules(createRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRulesOptions model with no property values
				createRulesOptionsModelNew := new(configurationgovernancev1.CreateRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.CreateRules(createRulesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRules(listRulesOptions *ListRulesOptions) - Operation response error`, func() {
		listRulesPath := "/config/v1/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"531fc3e28bfc43c5a2cea07786d93f5c"}))

					// TODO: Add check for attached query parameter

					Expect(req.URL.Query()["labels"]).To(Equal([]string{"SOC2,ITCS300"}))

					Expect(req.URL.Query()["scopes"]).To(Equal([]string{"scope_id"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRules with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configurationgovernancev1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Attached = core.BoolPtr(true)
				listRulesOptionsModel.Labels = core.StringPtr("SOC2,ITCS300")
				listRulesOptionsModel.Scopes = core.StringPtr("scope_id")
				listRulesOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRulesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListRules(listRulesOptions *ListRulesOptions)`, func() {
		listRulesPath := "/config/v1/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"531fc3e28bfc43c5a2cea07786d93f5c"}))

					// TODO: Add check for attached query parameter

					Expect(req.URL.Query()["labels"]).To(Equal([]string{"SOC2,ITCS300"}))

					Expect(req.URL.Query()["scopes"]).To(Equal([]string{"scope_id"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 1000, "total_count": 10, "first": {"href": "Href"}, "last": {"href": "Href"}, "rules": [{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "iam-groups", "resource_kind": "zone", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modification_date": "2019-01-01T12:00:00", "modified_by": "ModifiedBy", "number_of_attachments": 3}]}`)
				}))
			})
			It(`Invoke ListRules successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configurationgovernancev1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Attached = core.BoolPtr(true)
				listRulesOptionsModel.Labels = core.StringPtr("SOC2,ITCS300")
				listRulesOptionsModel.Scopes = core.StringPtr("scope_id")
				listRulesOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRulesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.ListRulesWithContext(ctx, listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.ListRulesWithContext(ctx, listRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"531fc3e28bfc43c5a2cea07786d93f5c"}))

					// TODO: Add check for attached query parameter

					Expect(req.URL.Query()["labels"]).To(Equal([]string{"SOC2,ITCS300"}))

					Expect(req.URL.Query()["scopes"]).To(Equal([]string{"scope_id"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 1000, "total_count": 10, "first": {"href": "Href"}, "last": {"href": "Href"}, "rules": [{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "iam-groups", "resource_kind": "zone", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modification_date": "2019-01-01T12:00:00", "modified_by": "ModifiedBy", "number_of_attachments": 3}]}`)
				}))
			})
			It(`Invoke ListRules successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.ListRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configurationgovernancev1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Attached = core.BoolPtr(true)
				listRulesOptionsModel.Labels = core.StringPtr("SOC2,ITCS300")
				listRulesOptionsModel.Scopes = core.StringPtr("scope_id")
				listRulesOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRulesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRules with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configurationgovernancev1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Attached = core.BoolPtr(true)
				listRulesOptionsModel.Labels = core.StringPtr("SOC2,ITCS300")
				listRulesOptionsModel.Scopes = core.StringPtr("scope_id")
				listRulesOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRulesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListRulesOptions model with no property values
				listRulesOptionsModelNew := new(configurationgovernancev1.ListRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.ListRules(listRulesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRule(getRuleOptions *GetRuleOptions) - Operation response error`, func() {
		getRulePath := "/config/v1/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRule with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configurationgovernancev1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
		getRulePath := "/config/v1/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "iam-groups", "resource_kind": "zone", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modification_date": "2019-01-01T12:00:00", "modified_by": "ModifiedBy", "number_of_attachments": 3}`)
				}))
			})
			It(`Invoke GetRule successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configurationgovernancev1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.GetRuleWithContext(ctx, getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.GetRuleWithContext(ctx, getRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "iam-groups", "resource_kind": "zone", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modification_date": "2019-01-01T12:00:00", "modified_by": "ModifiedBy", "number_of_attachments": 3}`)
				}))
			})
			It(`Invoke GetRule successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.GetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configurationgovernancev1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRule with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configurationgovernancev1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRuleOptions model with no property values
				getRuleOptionsModelNew := new(configurationgovernancev1.GetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.GetRule(getRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRule(updateRuleOptions *UpdateRuleOptions) - Operation response error`, func() {
		updateRulePath := "/config/v1/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRulePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRule with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleTargetAttribute model
				ruleTargetAttributeModel := new(configurationgovernancev1.RuleTargetAttribute)
				ruleTargetAttributeModel.Name = core.StringPtr("testString")
				ruleTargetAttributeModel.Operator = core.StringPtr("string_equals")
				ruleTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("audit_log")

				// Construct an instance of the UpdateRuleOptions model
				updateRuleOptionsModel := new(configurationgovernancev1.UpdateRuleOptions)
				updateRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleOptionsModel.Name = core.StringPtr("Disable public access")
				updateRuleOptionsModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.Target = targetResourceModel
				updateRuleOptionsModel.RequiredConfig = ruleRequiredConfigModel
				updateRuleOptionsModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				updateRuleOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.RuleType = core.StringPtr("user_defined")
				updateRuleOptionsModel.Labels = []string{"testString"}
				updateRuleOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateRule(updateRuleOptions *UpdateRuleOptions)`, func() {
		updateRulePath := "/config/v1/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRulePath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "iam-groups", "resource_kind": "zone", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modification_date": "2019-01-01T12:00:00", "modified_by": "ModifiedBy", "number_of_attachments": 3}`)
				}))
			})
			It(`Invoke UpdateRule successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the RuleTargetAttribute model
				ruleTargetAttributeModel := new(configurationgovernancev1.RuleTargetAttribute)
				ruleTargetAttributeModel.Name = core.StringPtr("testString")
				ruleTargetAttributeModel.Operator = core.StringPtr("string_equals")
				ruleTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("audit_log")

				// Construct an instance of the UpdateRuleOptions model
				updateRuleOptionsModel := new(configurationgovernancev1.UpdateRuleOptions)
				updateRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleOptionsModel.Name = core.StringPtr("Disable public access")
				updateRuleOptionsModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.Target = targetResourceModel
				updateRuleOptionsModel.RequiredConfig = ruleRequiredConfigModel
				updateRuleOptionsModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				updateRuleOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.RuleType = core.StringPtr("user_defined")
				updateRuleOptionsModel.Labels = []string{"testString"}
				updateRuleOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.UpdateRuleWithContext(ctx, updateRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.UpdateRuleWithContext(ctx, updateRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateRulePath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "iam-groups", "resource_kind": "zone", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2019-01-01T12:00:00", "created_by": "CreatedBy", "modification_date": "2019-01-01T12:00:00", "modified_by": "ModifiedBy", "number_of_attachments": 3}`)
				}))
			})
			It(`Invoke UpdateRule successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.UpdateRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleTargetAttribute model
				ruleTargetAttributeModel := new(configurationgovernancev1.RuleTargetAttribute)
				ruleTargetAttributeModel.Name = core.StringPtr("testString")
				ruleTargetAttributeModel.Operator = core.StringPtr("string_equals")
				ruleTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("audit_log")

				// Construct an instance of the UpdateRuleOptions model
				updateRuleOptionsModel := new(configurationgovernancev1.UpdateRuleOptions)
				updateRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleOptionsModel.Name = core.StringPtr("Disable public access")
				updateRuleOptionsModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.Target = targetResourceModel
				updateRuleOptionsModel.RequiredConfig = ruleRequiredConfigModel
				updateRuleOptionsModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				updateRuleOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.RuleType = core.StringPtr("user_defined")
				updateRuleOptionsModel.Labels = []string{"testString"}
				updateRuleOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateRule with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleTargetAttribute model
				ruleTargetAttributeModel := new(configurationgovernancev1.RuleTargetAttribute)
				ruleTargetAttributeModel.Name = core.StringPtr("testString")
				ruleTargetAttributeModel.Operator = core.StringPtr("string_equals")
				ruleTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("audit_log")

				// Construct an instance of the UpdateRuleOptions model
				updateRuleOptionsModel := new(configurationgovernancev1.UpdateRuleOptions)
				updateRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleOptionsModel.Name = core.StringPtr("Disable public access")
				updateRuleOptionsModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.Target = targetResourceModel
				updateRuleOptionsModel.RequiredConfig = ruleRequiredConfigModel
				updateRuleOptionsModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				updateRuleOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.RuleType = core.StringPtr("user_defined")
				updateRuleOptionsModel.Labels = []string{"testString"}
				updateRuleOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRuleOptions model with no property values
				updateRuleOptionsModelNew := new(configurationgovernancev1.UpdateRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.UpdateRule(updateRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
		deleteRulePath := "/config/v1/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRule successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := configurationGovernanceService.DeleteRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(configurationgovernancev1.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.TransactionID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = configurationGovernanceService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRule with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(configurationgovernancev1.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.TransactionID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := configurationGovernanceService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRuleOptions model with no property values
				deleteRuleOptionsModelNew := new(configurationgovernancev1.DeleteRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = configurationGovernanceService.DeleteRule(deleteRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAttachments(createAttachmentsOptions *CreateAttachmentsOptions) - Operation response error`, func() {
		createAttachmentsPath := "/config/v1/rules/testString/attachments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAttachmentsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAttachments with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the AttachmentRequest model
				attachmentRequestModel := new(configurationgovernancev1.AttachmentRequest)
				attachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				attachmentRequestModel.IncludedScope = ruleScopeModel
				attachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}

				// Construct an instance of the CreateAttachmentsOptions model
				createAttachmentsOptionsModel := new(configurationgovernancev1.CreateAttachmentsOptions)
				createAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				createAttachmentsOptionsModel.Attachments = []configurationgovernancev1.AttachmentRequest{*attachmentRequestModel}
				createAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				createAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.CreateAttachments(createAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.CreateAttachments(createAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateAttachments(createAttachmentsOptions *CreateAttachmentsOptions)`, func() {
		createAttachmentsPath := "/config/v1/rules/testString/attachments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAttachmentsPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"attachments": [{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}]}`)
				}))
			})
			It(`Invoke CreateAttachments successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the AttachmentRequest model
				attachmentRequestModel := new(configurationgovernancev1.AttachmentRequest)
				attachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				attachmentRequestModel.IncludedScope = ruleScopeModel
				attachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}

				// Construct an instance of the CreateAttachmentsOptions model
				createAttachmentsOptionsModel := new(configurationgovernancev1.CreateAttachmentsOptions)
				createAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				createAttachmentsOptionsModel.Attachments = []configurationgovernancev1.AttachmentRequest{*attachmentRequestModel}
				createAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				createAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.CreateAttachmentsWithContext(ctx, createAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.CreateAttachments(createAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.CreateAttachmentsWithContext(ctx, createAttachmentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAttachmentsPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"attachments": [{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}]}`)
				}))
			})
			It(`Invoke CreateAttachments successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.CreateAttachments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the AttachmentRequest model
				attachmentRequestModel := new(configurationgovernancev1.AttachmentRequest)
				attachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				attachmentRequestModel.IncludedScope = ruleScopeModel
				attachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}

				// Construct an instance of the CreateAttachmentsOptions model
				createAttachmentsOptionsModel := new(configurationgovernancev1.CreateAttachmentsOptions)
				createAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				createAttachmentsOptionsModel.Attachments = []configurationgovernancev1.AttachmentRequest{*attachmentRequestModel}
				createAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				createAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.CreateAttachments(createAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAttachments with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the AttachmentRequest model
				attachmentRequestModel := new(configurationgovernancev1.AttachmentRequest)
				attachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				attachmentRequestModel.IncludedScope = ruleScopeModel
				attachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}

				// Construct an instance of the CreateAttachmentsOptions model
				createAttachmentsOptionsModel := new(configurationgovernancev1.CreateAttachmentsOptions)
				createAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				createAttachmentsOptionsModel.Attachments = []configurationgovernancev1.AttachmentRequest{*attachmentRequestModel}
				createAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				createAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.CreateAttachments(createAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAttachmentsOptions model with no property values
				createAttachmentsOptionsModelNew := new(configurationgovernancev1.CreateAttachmentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.CreateAttachments(createAttachmentsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAttachments(listAttachmentsOptions *ListAttachmentsOptions) - Operation response error`, func() {
		listAttachmentsPath := "/config/v1/rules/testString/attachments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAttachments with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(configurationgovernancev1.ListAttachmentsOptions)
				listAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				listAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listAttachmentsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAttachments(listAttachmentsOptions *ListAttachmentsOptions)`, func() {
		listAttachmentsPath := "/config/v1/rules/testString/attachments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 1000, "total_count": 10, "first": {"href": "Href"}, "last": {"href": "Href"}, "attachments": [{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}]}`)
				}))
			})
			It(`Invoke ListAttachments successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(configurationgovernancev1.ListAttachmentsOptions)
				listAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				listAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listAttachmentsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.ListAttachmentsWithContext(ctx, listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.ListAttachmentsWithContext(ctx, listAttachmentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 1000, "total_count": 10, "first": {"href": "Href"}, "last": {"href": "Href"}, "attachments": [{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}]}`)
				}))
			})
			It(`Invoke ListAttachments successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.ListAttachments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(configurationgovernancev1.ListAttachmentsOptions)
				listAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				listAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listAttachmentsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAttachments with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(configurationgovernancev1.ListAttachmentsOptions)
				listAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				listAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listAttachmentsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAttachmentsOptions model with no property values
				listAttachmentsOptionsModelNew := new(configurationgovernancev1.ListAttachmentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.ListAttachments(listAttachmentsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAttachment(getAttachmentOptions *GetAttachmentOptions) - Operation response error`, func() {
		getAttachmentPath := "/config/v1/rules/testString/attachments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAttachmentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAttachment with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the GetAttachmentOptions model
				getAttachmentOptionsModel := new(configurationgovernancev1.GetAttachmentOptions)
				getAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				getAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				getAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.GetAttachment(getAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.GetAttachment(getAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAttachment(getAttachmentOptions *GetAttachmentOptions)`, func() {
		getAttachmentPath := "/config/v1/rules/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAttachmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}`)
				}))
			})
			It(`Invoke GetAttachment successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the GetAttachmentOptions model
				getAttachmentOptionsModel := new(configurationgovernancev1.GetAttachmentOptions)
				getAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				getAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				getAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.GetAttachmentWithContext(ctx, getAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.GetAttachment(getAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.GetAttachmentWithContext(ctx, getAttachmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAttachmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}`)
				}))
			})
			It(`Invoke GetAttachment successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.GetAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAttachmentOptions model
				getAttachmentOptionsModel := new(configurationgovernancev1.GetAttachmentOptions)
				getAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				getAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				getAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.GetAttachment(getAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAttachment with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the GetAttachmentOptions model
				getAttachmentOptionsModel := new(configurationgovernancev1.GetAttachmentOptions)
				getAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				getAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				getAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.GetAttachment(getAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAttachmentOptions model with no property values
				getAttachmentOptionsModelNew := new(configurationgovernancev1.GetAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.GetAttachment(getAttachmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAttachment(updateAttachmentOptions *UpdateAttachmentOptions) - Operation response error`, func() {
		updateAttachmentPath := "/config/v1/rules/testString/attachments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAttachmentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAttachment with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the UpdateAttachmentOptions model
				updateAttachmentOptionsModel := new(configurationgovernancev1.UpdateAttachmentOptions)
				updateAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				updateAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				updateAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAttachmentOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateAttachmentOptionsModel.IncludedScope = ruleScopeModel
				updateAttachmentOptionsModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				updateAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				updateAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.UpdateAttachment(updateAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.UpdateAttachment(updateAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateAttachment(updateAttachmentOptions *UpdateAttachmentOptions)`, func() {
		updateAttachmentPath := "/config/v1/rules/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAttachmentPath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}`)
				}))
			})
			It(`Invoke UpdateAttachment successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the UpdateAttachmentOptions model
				updateAttachmentOptionsModel := new(configurationgovernancev1.UpdateAttachmentOptions)
				updateAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				updateAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				updateAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAttachmentOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateAttachmentOptionsModel.IncludedScope = ruleScopeModel
				updateAttachmentOptionsModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				updateAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				updateAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.UpdateAttachmentWithContext(ctx, updateAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.UpdateAttachment(updateAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.UpdateAttachmentWithContext(ctx, updateAttachmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAttachmentPath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}`)
				}))
			})
			It(`Invoke UpdateAttachment successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.UpdateAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the UpdateAttachmentOptions model
				updateAttachmentOptionsModel := new(configurationgovernancev1.UpdateAttachmentOptions)
				updateAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				updateAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				updateAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAttachmentOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateAttachmentOptionsModel.IncludedScope = ruleScopeModel
				updateAttachmentOptionsModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				updateAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				updateAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.UpdateAttachment(updateAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAttachment with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the UpdateAttachmentOptions model
				updateAttachmentOptionsModel := new(configurationgovernancev1.UpdateAttachmentOptions)
				updateAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				updateAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				updateAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAttachmentOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateAttachmentOptionsModel.IncludedScope = ruleScopeModel
				updateAttachmentOptionsModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				updateAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				updateAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.UpdateAttachment(updateAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAttachmentOptions model with no property values
				updateAttachmentOptionsModelNew := new(configurationgovernancev1.UpdateAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.UpdateAttachment(updateAttachmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAttachment(deleteAttachmentOptions *DeleteAttachmentOptions)`, func() {
		deleteAttachmentPath := "/config/v1/rules/testString/attachments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAttachmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAttachment successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := configurationGovernanceService.DeleteAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAttachmentOptions model
				deleteAttachmentOptionsModel := new(configurationgovernancev1.DeleteAttachmentOptions)
				deleteAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				deleteAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				deleteAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = configurationGovernanceService.DeleteAttachment(deleteAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAttachment with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the DeleteAttachmentOptions model
				deleteAttachmentOptionsModel := new(configurationgovernancev1.DeleteAttachmentOptions)
				deleteAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				deleteAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				deleteAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := configurationGovernanceService.DeleteAttachment(deleteAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAttachmentOptions model with no property values
				deleteAttachmentOptionsModelNew := new(configurationgovernancev1.DeleteAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = configurationGovernanceService.DeleteAttachment(deleteAttachmentOptionsModelNew)
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
			configurationGovernanceService, _ := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
				URL:           "http://configurationgovernancev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAttachmentRequest successfully`, func() {
				accountID := "testString"
				var includedScope *configurationgovernancev1.RuleScope = nil
				_, err := configurationGovernanceService.NewAttachmentRequest(accountID, includedScope)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateAttachmentsOptions successfully`, func() {
				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				Expect(ruleScopeModel).ToNot(BeNil())
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")
				Expect(ruleScopeModel.Note).To(Equal(core.StringPtr("My enterprise")))
				Expect(ruleScopeModel.ScopeID).To(Equal(core.StringPtr("282cf433ac91493ba860480d92519990")))
				Expect(ruleScopeModel.ScopeType).To(Equal(core.StringPtr("enterprise")))

				// Construct an instance of the AttachmentRequest model
				attachmentRequestModel := new(configurationgovernancev1.AttachmentRequest)
				Expect(attachmentRequestModel).ToNot(BeNil())
				attachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				attachmentRequestModel.IncludedScope = ruleScopeModel
				attachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				Expect(attachmentRequestModel.AccountID).To(Equal(core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")))
				Expect(attachmentRequestModel.IncludedScope).To(Equal(ruleScopeModel))
				Expect(attachmentRequestModel.ExcludedScopes).To(Equal([]configurationgovernancev1.RuleScope{*ruleScopeModel}))

				// Construct an instance of the CreateAttachmentsOptions model
				ruleID := "testString"
				createAttachmentsOptionsAttachments := []configurationgovernancev1.AttachmentRequest{}
				createAttachmentsOptionsModel := configurationGovernanceService.NewCreateAttachmentsOptions(ruleID, createAttachmentsOptionsAttachments)
				createAttachmentsOptionsModel.SetRuleID("testString")
				createAttachmentsOptionsModel.SetAttachments([]configurationgovernancev1.AttachmentRequest{*attachmentRequestModel})
				createAttachmentsOptionsModel.SetTransactionID("testString")
				createAttachmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAttachmentsOptionsModel).ToNot(BeNil())
				Expect(createAttachmentsOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentsOptionsModel.Attachments).To(Equal([]configurationgovernancev1.AttachmentRequest{*attachmentRequestModel}))
				Expect(createAttachmentsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateRuleRequest successfully`, func() {
				var rule *configurationgovernancev1.RuleRequest = nil
				_, err := configurationGovernanceService.NewCreateRuleRequest(rule)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateRulesOptions successfully`, func() {
				// Construct an instance of the RuleTargetAttribute model
				ruleTargetAttributeModel := new(configurationgovernancev1.RuleTargetAttribute)
				Expect(ruleTargetAttributeModel).ToNot(BeNil())
				ruleTargetAttributeModel.Name = core.StringPtr("resource_id")
				ruleTargetAttributeModel.Operator = core.StringPtr("string_equals")
				ruleTargetAttributeModel.Value = core.StringPtr("f0f8f7994e754ff38f9d370201966561")
				Expect(ruleTargetAttributeModel.Name).To(Equal(core.StringPtr("resource_id")))
				Expect(ruleTargetAttributeModel.Operator).To(Equal(core.StringPtr("string_equals")))
				Expect(ruleTargetAttributeModel.Value).To(Equal(core.StringPtr("f0f8f7994e754ff38f9d370201966561")))

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				Expect(targetResourceModel).ToNot(BeNil())
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}
				Expect(targetResourceModel.ServiceName).To(Equal(core.StringPtr("iam-groups")))
				Expect(targetResourceModel.ResourceKind).To(Equal(core.StringPtr("service")))
				Expect(targetResourceModel.AdditionalTargetAttributes).To(Equal([]configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}))

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				Expect(ruleConditionModel).ToNot(BeNil())
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")
				Expect(ruleConditionModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(ruleConditionModel.Property).To(Equal(core.StringPtr("public_access_enabled")))
				Expect(ruleConditionModel.Operator).To(Equal(core.StringPtr("is_false")))
				Expect(ruleConditionModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				Expect(ruleRequiredConfigModel).ToNot(BeNil())
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}
				Expect(ruleRequiredConfigModel.Description).To(Equal(core.StringPtr("Public access check")))
				Expect(ruleRequiredConfigModel.And).To(Equal([]configurationgovernancev1.RuleConditionIntf{ruleConditionModel}))

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				Expect(enforcementActionModel).ToNot(BeNil())
				enforcementActionModel.Action = core.StringPtr("disallow")
				Expect(enforcementActionModel.Action).To(Equal(core.StringPtr("disallow")))

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				Expect(ruleRequestModel).ToNot(BeNil())
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"testString"}
				Expect(ruleRequestModel.AccountID).To(Equal(core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")))
				Expect(ruleRequestModel.Name).To(Equal(core.StringPtr("Disable public access")))
				Expect(ruleRequestModel.Description).To(Equal(core.StringPtr("Ensure that public access to account resources is disabled.")))
				Expect(ruleRequestModel.RuleType).To(Equal(core.StringPtr("user_defined")))
				Expect(ruleRequestModel.Target).To(Equal(targetResourceModel))
				Expect(ruleRequestModel.RequiredConfig).To(Equal(ruleRequiredConfigModel))
				Expect(ruleRequestModel.EnforcementActions).To(Equal([]configurationgovernancev1.EnforcementAction{*enforcementActionModel}))
				Expect(ruleRequestModel.Labels).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				Expect(createRuleRequestModel).ToNot(BeNil())
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel
				Expect(createRuleRequestModel.RequestID).To(Equal(core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")))
				Expect(createRuleRequestModel.Rule).To(Equal(ruleRequestModel))

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsRules := []configurationgovernancev1.CreateRuleRequest{}
				createRulesOptionsModel := configurationGovernanceService.NewCreateRulesOptions(createRulesOptionsRules)
				createRulesOptionsModel.SetRules([]configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel})
				createRulesOptionsModel.SetTransactionID("testString")
				createRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRulesOptionsModel).ToNot(BeNil())
				Expect(createRulesOptionsModel.Rules).To(Equal([]configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}))
				Expect(createRulesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAttachmentOptions successfully`, func() {
				// Construct an instance of the DeleteAttachmentOptions model
				ruleID := "testString"
				attachmentID := "testString"
				deleteAttachmentOptionsModel := configurationGovernanceService.NewDeleteAttachmentOptions(ruleID, attachmentID)
				deleteAttachmentOptionsModel.SetRuleID("testString")
				deleteAttachmentOptionsModel.SetAttachmentID("testString")
				deleteAttachmentOptionsModel.SetTransactionID("testString")
				deleteAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAttachmentOptionsModel).ToNot(BeNil())
				Expect(deleteAttachmentOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAttachmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRuleOptions successfully`, func() {
				// Construct an instance of the DeleteRuleOptions model
				ruleID := "testString"
				deleteRuleOptionsModel := configurationGovernanceService.NewDeleteRuleOptions(ruleID)
				deleteRuleOptionsModel.SetRuleID("testString")
				deleteRuleOptionsModel.SetTransactionID("testString")
				deleteRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRuleOptionsModel).ToNot(BeNil())
				Expect(deleteRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEnforcementAction successfully`, func() {
				action := "audit_log"
				model, err := configurationGovernanceService.NewEnforcementAction(action)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetAttachmentOptions successfully`, func() {
				// Construct an instance of the GetAttachmentOptions model
				ruleID := "testString"
				attachmentID := "testString"
				getAttachmentOptionsModel := configurationGovernanceService.NewGetAttachmentOptions(ruleID, attachmentID)
				getAttachmentOptionsModel.SetRuleID("testString")
				getAttachmentOptionsModel.SetAttachmentID("testString")
				getAttachmentOptionsModel.SetTransactionID("testString")
				getAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAttachmentOptionsModel).ToNot(BeNil())
				Expect(getAttachmentOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRuleOptions successfully`, func() {
				// Construct an instance of the GetRuleOptions model
				ruleID := "testString"
				getRuleOptionsModel := configurationGovernanceService.NewGetRuleOptions(ruleID)
				getRuleOptionsModel.SetRuleID("testString")
				getRuleOptionsModel.SetTransactionID("testString")
				getRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRuleOptionsModel).ToNot(BeNil())
				Expect(getRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAttachmentsOptions successfully`, func() {
				// Construct an instance of the ListAttachmentsOptions model
				ruleID := "testString"
				listAttachmentsOptionsModel := configurationGovernanceService.NewListAttachmentsOptions(ruleID)
				listAttachmentsOptionsModel.SetRuleID("testString")
				listAttachmentsOptionsModel.SetTransactionID("testString")
				listAttachmentsOptionsModel.SetLimit(int64(1000))
				listAttachmentsOptionsModel.SetOffset(int64(38))
				listAttachmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAttachmentsOptionsModel).ToNot(BeNil())
				Expect(listAttachmentsOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(listAttachmentsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAttachmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRulesOptions successfully`, func() {
				// Construct an instance of the ListRulesOptions model
				accountID := "531fc3e28bfc43c5a2cea07786d93f5c"
				listRulesOptionsModel := configurationGovernanceService.NewListRulesOptions(accountID)
				listRulesOptionsModel.SetAccountID("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.SetTransactionID("testString")
				listRulesOptionsModel.SetAttached(true)
				listRulesOptionsModel.SetLabels("SOC2,ITCS300")
				listRulesOptionsModel.SetScopes("scope_id")
				listRulesOptionsModel.SetLimit(int64(1000))
				listRulesOptionsModel.SetOffset(int64(38))
				listRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRulesOptionsModel).ToNot(BeNil())
				Expect(listRulesOptionsModel.AccountID).To(Equal(core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")))
				Expect(listRulesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.Attached).To(Equal(core.BoolPtr(true)))
				Expect(listRulesOptionsModel.Labels).To(Equal(core.StringPtr("SOC2,ITCS300")))
				Expect(listRulesOptionsModel.Scopes).To(Equal(core.StringPtr("scope_id")))
				Expect(listRulesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(listRulesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRuleRequest successfully`, func() {
				name := "testString"
				description := "testString"
				var target *configurationgovernancev1.TargetResource = nil
				var requiredConfig configurationgovernancev1.RuleRequiredConfigIntf = nil
				enforcementActions := []configurationgovernancev1.EnforcementAction{}
				_, err := configurationGovernanceService.NewRuleRequest(name, description, target, requiredConfig, enforcementActions)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewRuleScope successfully`, func() {
				scopeID := "testString"
				scopeType := "enterprise"
				model, err := configurationGovernanceService.NewRuleScope(scopeID, scopeType)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleSingleProperty successfully`, func() {
				property := "public_access_enabled"
				operator := "is_true"
				model, err := configurationGovernanceService.NewRuleSingleProperty(property, operator)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleTargetAttribute successfully`, func() {
				name := "testString"
				operator := "string_equals"
				model, err := configurationGovernanceService.NewRuleTargetAttribute(name, operator)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTargetResource successfully`, func() {
				serviceName := "iam-groups"
				resourceKind := "zone"
				model, err := configurationGovernanceService.NewTargetResource(serviceName, resourceKind)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateAttachmentOptions successfully`, func() {
				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				Expect(ruleScopeModel).ToNot(BeNil())
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")
				Expect(ruleScopeModel.Note).To(Equal(core.StringPtr("My enterprise")))
				Expect(ruleScopeModel.ScopeID).To(Equal(core.StringPtr("282cf433ac91493ba860480d92519990")))
				Expect(ruleScopeModel.ScopeType).To(Equal(core.StringPtr("enterprise")))

				// Construct an instance of the UpdateAttachmentOptions model
				ruleID := "testString"
				attachmentID := "testString"
				ifMatch := "testString"
				updateAttachmentOptionsAccountID := "531fc3e28bfc43c5a2cea07786d93f5c"
				var updateAttachmentOptionsIncludedScope *configurationgovernancev1.RuleScope = nil
				updateAttachmentOptionsModel := configurationGovernanceService.NewUpdateAttachmentOptions(ruleID, attachmentID, ifMatch, updateAttachmentOptionsAccountID, updateAttachmentOptionsIncludedScope)
				updateAttachmentOptionsModel.SetRuleID("testString")
				updateAttachmentOptionsModel.SetAttachmentID("testString")
				updateAttachmentOptionsModel.SetIfMatch("testString")
				updateAttachmentOptionsModel.SetAccountID("531fc3e28bfc43c5a2cea07786d93f5c")
				updateAttachmentOptionsModel.SetIncludedScope(ruleScopeModel)
				updateAttachmentOptionsModel.SetExcludedScopes([]configurationgovernancev1.RuleScope{*ruleScopeModel})
				updateAttachmentOptionsModel.SetTransactionID("testString")
				updateAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAttachmentOptionsModel).ToNot(BeNil())
				Expect(updateAttachmentOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(updateAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateAttachmentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAttachmentOptionsModel.AccountID).To(Equal(core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")))
				Expect(updateAttachmentOptionsModel.IncludedScope).To(Equal(ruleScopeModel))
				Expect(updateAttachmentOptionsModel.ExcludedScopes).To(Equal([]configurationgovernancev1.RuleScope{*ruleScopeModel}))
				Expect(updateAttachmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateRuleOptions successfully`, func() {
				// Construct an instance of the RuleTargetAttribute model
				ruleTargetAttributeModel := new(configurationgovernancev1.RuleTargetAttribute)
				Expect(ruleTargetAttributeModel).ToNot(BeNil())
				ruleTargetAttributeModel.Name = core.StringPtr("testString")
				ruleTargetAttributeModel.Operator = core.StringPtr("string_equals")
				ruleTargetAttributeModel.Value = core.StringPtr("testString")
				Expect(ruleTargetAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(ruleTargetAttributeModel.Operator).To(Equal(core.StringPtr("string_equals")))
				Expect(ruleTargetAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				Expect(targetResourceModel).ToNot(BeNil())
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}
				Expect(targetResourceModel.ServiceName).To(Equal(core.StringPtr("iam-groups")))
				Expect(targetResourceModel.ResourceKind).To(Equal(core.StringPtr("service")))
				Expect(targetResourceModel.AdditionalTargetAttributes).To(Equal([]configurationgovernancev1.RuleTargetAttribute{*ruleTargetAttributeModel}))

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				Expect(ruleRequiredConfigModel).ToNot(BeNil())
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")
				Expect(ruleRequiredConfigModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(ruleRequiredConfigModel.Property).To(Equal(core.StringPtr("public_access_enabled")))
				Expect(ruleRequiredConfigModel.Operator).To(Equal(core.StringPtr("is_false")))
				Expect(ruleRequiredConfigModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				Expect(enforcementActionModel).ToNot(BeNil())
				enforcementActionModel.Action = core.StringPtr("audit_log")
				Expect(enforcementActionModel.Action).To(Equal(core.StringPtr("audit_log")))

				// Construct an instance of the UpdateRuleOptions model
				ruleID := "testString"
				ifMatch := "testString"
				updateRuleOptionsName := "Disable public access"
				updateRuleOptionsDescription := "Ensure that public access to account resources is disabled."
				var updateRuleOptionsTarget *configurationgovernancev1.TargetResource = nil
				var updateRuleOptionsRequiredConfig configurationgovernancev1.RuleRequiredConfigIntf = nil
				updateRuleOptionsEnforcementActions := []configurationgovernancev1.EnforcementAction{}
				updateRuleOptionsModel := configurationGovernanceService.NewUpdateRuleOptions(ruleID, ifMatch, updateRuleOptionsName, updateRuleOptionsDescription, updateRuleOptionsTarget, updateRuleOptionsRequiredConfig, updateRuleOptionsEnforcementActions)
				updateRuleOptionsModel.SetRuleID("testString")
				updateRuleOptionsModel.SetIfMatch("testString")
				updateRuleOptionsModel.SetName("Disable public access")
				updateRuleOptionsModel.SetDescription("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.SetTarget(targetResourceModel)
				updateRuleOptionsModel.SetRequiredConfig(ruleRequiredConfigModel)
				updateRuleOptionsModel.SetEnforcementActions([]configurationgovernancev1.EnforcementAction{*enforcementActionModel})
				updateRuleOptionsModel.SetAccountID("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.SetRuleType("user_defined")
				updateRuleOptionsModel.SetLabels([]string{"testString"})
				updateRuleOptionsModel.SetTransactionID("testString")
				updateRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRuleOptionsModel).ToNot(BeNil())
				Expect(updateRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(updateRuleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateRuleOptionsModel.Name).To(Equal(core.StringPtr("Disable public access")))
				Expect(updateRuleOptionsModel.Description).To(Equal(core.StringPtr("Ensure that public access to account resources is disabled.")))
				Expect(updateRuleOptionsModel.Target).To(Equal(targetResourceModel))
				Expect(updateRuleOptionsModel.RequiredConfig).To(Equal(ruleRequiredConfigModel))
				Expect(updateRuleOptionsModel.EnforcementActions).To(Equal([]configurationgovernancev1.EnforcementAction{*enforcementActionModel}))
				Expect(updateRuleOptionsModel.AccountID).To(Equal(core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")))
				Expect(updateRuleOptionsModel.RuleType).To(Equal(core.StringPtr("user_defined")))
				Expect(updateRuleOptionsModel.Labels).To(Equal([]string{"testString"}))
				Expect(updateRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRuleConditionAndLvl2 successfully`, func() {
				and := []configurationgovernancev1.RuleSingleProperty{}
				model, err := configurationGovernanceService.NewRuleConditionAndLvl2(and)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleConditionOrLvl2 successfully`, func() {
				or := []configurationgovernancev1.RuleSingleProperty{}
				model, err := configurationGovernanceService.NewRuleConditionOrLvl2(or)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleConditionSingleProperty successfully`, func() {
				property := "public_access_enabled"
				operator := "is_true"
				model, err := configurationGovernanceService.NewRuleConditionSingleProperty(property, operator)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleRequiredConfigSingleProperty successfully`, func() {
				property := "public_access_enabled"
				operator := "is_true"
				model, err := configurationGovernanceService.NewRuleRequiredConfigSingleProperty(property, operator)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleRequiredConfigMultiplePropertiesConditionAnd successfully`, func() {
				and := []configurationgovernancev1.RuleConditionIntf{}
				model, err := configurationGovernanceService.NewRuleRequiredConfigMultiplePropertiesConditionAnd(and)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleRequiredConfigMultiplePropertiesConditionOr successfully`, func() {
				or := []configurationgovernancev1.RuleConditionIntf{}
				model, err := configurationGovernanceService.NewRuleRequiredConfigMultiplePropertiesConditionOr(or)
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
