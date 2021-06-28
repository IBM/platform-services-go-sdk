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

package ibmcloudshellv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/ibmcloudshellv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the ibmcloudshellv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`IbmCloudShellV1 Integration Tests`, func() {

	const externalConfigFile = "../ibm_cloud_shell_v1.env"

	var (
		err                  error
		ibmCloudShellService *ibmcloudshellv1.IbmCloudShellV1
		serviceURL           string
		config               map[string]string
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
			config, err = core.GetServiceProperties(ibmcloudshellv1.DefaultServiceName)
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

			ibmCloudShellServiceOptions := &ibmcloudshellv1.IbmCloudShellV1Options{}

			ibmCloudShellService, err = ibmcloudshellv1.NewIbmCloudShellV1UsingExternalConfig(ibmCloudShellServiceOptions)

			Expect(err).To(BeNil())
			Expect(ibmCloudShellService).ToNot(BeNil())
			Expect(ibmCloudShellService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`GetAccountSettingsByID - Get account settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSettingsByID(getAccountSettingsByIdOptions *GetAccountSettingsByIdOptions)`, func() {

			getAccountSettingsByIdOptions := &ibmcloudshellv1.GetAccountSettingsByIdOptions{
				AccountID: core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab"),
			}

			accountSettings, response, err := ibmCloudShellService.GetAccountSettingsByID(getAccountSettingsByIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
	})

	Describe(`UpdateAccountSettingsByID - Update account settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAccountSettingsByID(updateAccountSettingsByIdOptions *UpdateAccountSettingsByIdOptions)`, func() {

			featureModel := []ibmcloudshellv1.Feature{
				{
					Enabled: core.BoolPtr(true),
					Key:     core.StringPtr("server.file_manager"),
				},
				{
					Enabled: core.BoolPtr(true),
					Key:     core.StringPtr("server.web_preview"),
				},
			}

			regionSettingModel := []ibmcloudshellv1.RegionSetting{
				{
					Enabled: core.BoolPtr(true),
					Key:     core.StringPtr("eu-de"),
				},
				{
					Enabled: core.BoolPtr(true),
					Key:     core.StringPtr("jp-tok"),
				},
				{
					Enabled: core.BoolPtr(true),
					Key:     core.StringPtr("us-south"),
				},
			}

			accountID := "12345678-abcd-1a2b-a1b2-1234567890ab"
			updateAccountSettingsByIdOptions := &ibmcloudshellv1.UpdateAccountSettingsByIdOptions{
				AccountID:                   core.StringPtr(accountID),
				NewID:                       core.StringPtr("ac" + accountID),
				NewRev:                      core.StringPtr("130" + accountID),
				NewAccountID:                core.StringPtr(accountID),
				NewCreatedAt:                core.Int64Ptr(int64(1600079615)),
				NewCreatedBy:                core.StringPtr("IBMid-1000000000"),
				NewDefaultEnableNewFeatures: core.BoolPtr(true),
				NewDefaultEnableNewRegions:  core.BoolPtr(true),
				NewEnabled:                  core.BoolPtr(true),
				NewFeatures:                 featureModel,
				NewRegions:                  regionSettingModel,
				NewType:                     core.StringPtr("account_settings"),
				NewUpdatedAt:                core.Int64Ptr(int64(1624359948)),
				NewUpdatedBy:                core.StringPtr("IBMid-1000000000"),
			}

			accountSettings, response, err := ibmCloudShellService.UpdateAccountSettingsByID(updateAccountSettingsByIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
	})
})

//
// Utility functions are declared in the unit test file
//
