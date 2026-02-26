//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/accountmanagementv4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the account_management service.
//
// The following configuration properties are assumed to be defined:
// ACCOUNT_MANAGEMENT_URL=<service base url>
// ACCOUNT_MANAGEMENT_AUTH_TYPE=iam
// ACCOUNT_MANAGEMENT_APIKEY=<IAM apikey>
// ACCOUNT_MANAGEMENT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`AccountManagementV4 Examples Tests`, func() {

	const externalConfigFile = "../account_management_v4.env"

	var (
		accountManagementService *accountmanagementv4.AccountManagementV4
		config                   map[string]string
		accountID                string
		serviceURL               string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(accountmanagementv4.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
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
			var err error

			// begin-common

			accountManagementServiceOptions := &accountmanagementv4.AccountManagementV4Options{}

			accountManagementService, err = accountmanagementv4.NewAccountManagementV4UsingExternalConfig(accountManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(accountManagementService).ToNot(BeNil())
		})
	})

	Describe(`AccountManagementV4 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccount request example`, func() {
			fmt.Println("\nGetAccount() result:")
			// begin-getAccount

			getAccountOptions := accountManagementService.NewGetAccountOptions(
				accountID,
			)

			accountResponse, response, err := accountManagementService.GetAccount(getAccountOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountResponse, "", "  ")
			fmt.Println(string(b))

			// end-getAccount

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountResponse).ToNot(BeNil())
		})
	})
})
