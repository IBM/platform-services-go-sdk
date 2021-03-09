// +build integration

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

package securitycompliancev1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/securitycompliancev1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the securitycompliancev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`SecurityComplianceV1 Integration Tests`, func() {

	const externalConfigFile = "../security_compliance.env"

	var (
		err                       error
		securityComplianceService *securitycompliancev1.SecurityComplianceV1
		serviceURL                string
		config                    map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(securitycompliancev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			securityComplianceServiceOptions := &securitycompliancev1.SecurityComplianceV1Options{}

			securityComplianceService, err = securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(securityComplianceServiceOptions)

			Expect(err).To(BeNil())
			Expect(securityComplianceService).ToNot(BeNil())
			Expect(securityComplianceService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`CreateValidationScan - Initiate a validation scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateValidationScan(createValidationScanOptions *CreateValidationScanOptions)`, func() {

			createValidationScanOptions := &securitycompliancev1.CreateValidationScanOptions{
				AccountID:      core.StringPtr("testString"),
				ScopeID:        core.Int64Ptr(int64(1)),
				ProfileID:      core.Int64Ptr(int64(6)),
				GroupProfileID: core.Int64Ptr(int64(13)),
			}

			result, response, err := securityComplianceService.CreateValidationScan(createValidationScanOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
	})

	Describe(`ListProfile - List profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfile(listProfileOptions *ListProfileOptions)`, func() {

			listProfileOptions := &securitycompliancev1.ListProfileOptions{
				AccountID: core.StringPtr("testString"),
				Name:      core.StringPtr("testString"),
			}

			profilesList, response, err := securityComplianceService.ListProfile(listProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profilesList).ToNot(BeNil())

		})
	})

	Describe(`ListScopes - List scopes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListScopes(listScopesOptions *ListScopesOptions)`, func() {

			listScopesOptions := &securitycompliancev1.ListScopesOptions{
				AccountID: core.StringPtr("testString"),
				Name:      core.StringPtr("testString"),
			}

			scopesList, response, err := securityComplianceService.ListScopes(listScopesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopesList).ToNot(BeNil())

		})
	})
})

//
// Utility functions are declared in the unit test file
//
