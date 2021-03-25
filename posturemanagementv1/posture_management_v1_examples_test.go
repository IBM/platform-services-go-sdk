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

package posturemanagementv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/posturemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Posture Management service.
//
// The following configuration properties are assumed to be defined:
// POSTURE_MANAGEMENT_URL=<service base url>
// POSTURE_MANAGEMENT_AUTH_TYPE=iam
// POSTURE_MANAGEMENT_APIKEY=<IAM apikey>
// POSTURE_MANAGEMENT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// POSTURE_MANAGEMENT_ACCOUNT_ID=<IBM CLOUD ACCOUNT ID>
// POSTURE_MANAGEMENT_SCOPES_NAME=<The name of the scope>
// POSTURE_MANAGEMENT_PROFILE_NAME=<The name of profile - CIS IBM Foundations Benchmark 1.0.0>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../posture_management.env"

var (
	postureManagementService *posturemanagementv1.PostureManagementV1
	config                   map[string]string
	configLoaded             bool = false

	accountID   string
	profileName string
	scopesName  string

	profileID string
	scopeID   string

	groupProfileID = "0"
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`PostureManagementV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(posturemanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			profileName = config["PROFILE_NAME"]
			Expect(profileName).ToNot(BeEmpty())

			scopesName = config["SCOPES_NAME"]
			Expect(scopesName).ToNot(BeEmpty())
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			postureManagementServiceOptions := &posturemanagementv1.PostureManagementV1Options{}

			postureManagementService, err = posturemanagementv1.NewPostureManagementV1UsingExternalConfig(postureManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(postureManagementService).ToNot(BeNil())
		})
	})

	Describe(`PostureManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfiles request example`, func() {
			// begin-list_profiles

			listProfilesOptions := postureManagementService.NewListProfilesOptions(
				accountID,
			)
			listProfilesOptions.SetName(profileName)

			profilesList, response, err := postureManagementService.ListProfiles(listProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profilesList, "", "  ")
			fmt.Printf("\nListProfiles() result:\n%s\n", string(b))

			// end-list_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profilesList).ToNot(BeNil())

			profileID = *profilesList.Profiles[0].ProfileID
		})
		It(`ListScopes request example`, func() {
			// begin-list_scopes

			listScopesOptions := postureManagementService.NewListScopesOptions(
				accountID,
			)
			listScopesOptions.SetName(scopesName)

			scopesList, response, err := postureManagementService.ListScopes(listScopesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scopesList, "", "  ")
			fmt.Printf("\nListScopes() result:\n%s\n", string(b))

			// end-list_scopes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopesList).ToNot(BeNil())

			scopeID = *scopesList.Scopes[0].ScopeID
		})
		It(`CreateValidation request example`, func() {
			// begin-create_validation

			createValidationOptions := postureManagementService.NewCreateValidationOptions(
				accountID,
			)
			createValidationOptions.SetProfileID(profileID)
			createValidationOptions.SetScopeID(scopeID)
			createValidationOptions.SetGroupProfileID(groupProfileID)

			result, response, err := postureManagementService.CreateValidation(createValidationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Printf("\nCreateValidation() result:\n%s\n", string(b))

			// end-create_validation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(result).ToNot(BeNil())

		})
	})
})
