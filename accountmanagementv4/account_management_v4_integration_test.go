//go:build integration

/**
 * (C) Copyright IBM Corp. 2026.
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

package accountmanagementv4_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/accountmanagementv4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the accountmanagementv4 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`AccountManagementV4 Integration Tests`, func() {
	const externalConfigFile = "../account_management_v4.env"

	var (
		err                      error
		accountManagementService *accountmanagementv4.AccountManagementV4
		serviceURL               string
		config                   map[string]string
		accountID                string
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
			config, err = core.GetServiceProperties(accountmanagementv4.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			accountID = config["ACCOUNT_ID"]
			if accountID == "" {
				Skip("Unable to load account id configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			fmt.Fprintf(GinkgoWriter, "Account ID: %v\n", accountID)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			accountManagementServiceOptions := &accountmanagementv4.AccountManagementV4Options{}

			accountManagementService, err = accountmanagementv4.NewAccountManagementV4UsingExternalConfig(accountManagementServiceOptions)
			Expect(err).To(BeNil())
			Expect(accountManagementService).ToNot(BeNil())
			Expect(accountManagementService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			accountManagementService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetAccount - Get Account by Account ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccount(getAccountOptions *GetAccountOptions)`, func() {
			getAccountOptions := &accountmanagementv4.GetAccountOptions{
				AccountID: core.StringPtr(accountID),
			}

			accountResponse, response, err := accountManagementService.GetAccount(getAccountOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountResponse).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
