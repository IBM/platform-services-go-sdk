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

package posturemanagementv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/posturemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the posturemanagementv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PostureManagementV1 Integration Tests`, func() {

	const externalConfigFile = "../posture_management.env"

	var (
		err                      error
		postureManagementService *posturemanagementv1.PostureManagementV1
		serviceURL               string
		config                   map[string]string

		accountID   string
		profileName string
		scopesName  string

		profileID *int64
		scopeID   *int64
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
			config, err = core.GetServiceProperties(posturemanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			profileName = config["PROFILE_NAME"]
			Expect(profileName).ToNot(BeEmpty())

			scopesName = config["SCOPES_NAME"]
			Expect(scopesName).ToNot(BeEmpty())

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			postureManagementServiceOptions := &posturemanagementv1.PostureManagementV1Options{}

			postureManagementService, err = posturemanagementv1.NewPostureManagementV1UsingExternalConfig(postureManagementServiceOptions)

			Expect(err).To(BeNil())
			Expect(postureManagementService).ToNot(BeNil())
			Expect(postureManagementService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`ListProfile - List profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfile(listProfileOptions *ListProfileOptions)`, func() {

			listProfileOptions := &posturemanagementv1.ListProfileOptions{
				AccountID: core.StringPtr(accountID),
				Name:      core.StringPtr(profileName),
			}

			profilesList, response, err := postureManagementService.ListProfile(listProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profilesList).ToNot(BeNil())

			profileID = profilesList.Profiles[0].ProfileID
		})
	})

	Describe(`ListScopes - List scopes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListScopes(listScopesOptions *ListScopesOptions)`, func() {

			listScopesOptions := &posturemanagementv1.ListScopesOptions{
				AccountID: core.StringPtr(accountID),
				Name:      core.StringPtr(scopesName),
			}

			scopesList, response, err := postureManagementService.ListScopes(listScopesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopesList).ToNot(BeNil())

			scopeID = scopesList.Scopes[0].ScopeID
		})
	})

	Describe(`CreateValidationScan - Initiate a validation scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateValidationScan(createValidationScanOptions *CreateValidationScanOptions)`, func() {
			Expect(scopeID).ToNot(BeNil())
			Expect(profileID).ToNot(BeNil())

			createValidationScanOptions := &posturemanagementv1.CreateValidationScanOptions{
				AccountID: core.StringPtr(accountID),
				ScopeID:   scopeID,
				ProfileID: profileID,
			}

			result, response, err := postureManagementService.CreateValidationScan(createValidationScanOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
	})
})

//
// Utility functions are declared in the unit test file
//
