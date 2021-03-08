// +build examples

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
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/securitycompliancev1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

//
// This file provides an example of how to use the Security Compliance service.
//
// The following configuration properties are assumed to be defined:
// SECURITY_COMPLIANCE_URL=<service base url>
// SECURITY_COMPLIANCE_AUTH_TYPE=iam
// SECURITY_COMPLIANCE_APIKEY=<IAM apikey>
// SECURITY_COMPLIANCE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../security_compliance_v1.env"

var (
	securityComplianceService *securitycompliancev1.SecurityComplianceV1
	config       map[string]string
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`SecurityComplianceV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(securitycompliancev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			securityComplianceServiceOptions := &securitycompliancev1.SecurityComplianceV1Options{}

			securityComplianceService, err = securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(securityComplianceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(securityComplianceService).ToNot(BeNil())
		})
	})

	Describe(`SecurityComplianceV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateValidationScan request example`, func() {
			// begin-create_validation_scan

			createValidationScanOptions := securityComplianceService.NewCreateValidationScanOptions(
				"testString",
			)

			result, response, err := securityComplianceService.CreateValidationScan(createValidationScanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-create_validation_scan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`ListProfile request example`, func() {
			// begin-list_profile

			listProfileOptions := securityComplianceService.NewListProfileOptions(
				"testString",
			)

			profilesList, response, err := securityComplianceService.ListProfile(listProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profilesList, "", "  ")
			fmt.Println(string(b))

			// end-list_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profilesList).ToNot(BeNil())

		})
		It(`ListScopes request example`, func() {
			// begin-list_scopes

			listScopesOptions := securityComplianceService.NewListScopesOptions(
				"testString",
			)

			scopesList, response, err := securityComplianceService.ListScopes(listScopesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scopesList, "", "  ")
			fmt.Println(string(b))

			// end-list_scopes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopesList).ToNot(BeNil())

		})
	})
})
