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

package ibmcloudshellv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/ibmcloudshellv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IBM Cloud Shell service.
//
// The following configuration properties are assumed to be defined:
// IBM_CLOUD_SHELL_URL=<service base url>
// IBM_CLOUD_SHELL_AUTH_TYPE=iam
// IBM_CLOUD_SHELL_APIKEY=<IAM apikey>
// IBM_CLOUD_SHELL_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../ibm_cloud_shell_v1.env"

var (
	ibmCloudShellService *ibmcloudshellv1.IbmCloudShellV1
	config               map[string]string
	configLoaded         bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`IbmCloudShellV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(ibmcloudshellv1.DefaultServiceName)
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

			ibmCloudShellServiceOptions := &ibmcloudshellv1.IbmCloudShellV1Options{}

			ibmCloudShellService, err = ibmcloudshellv1.NewIbmCloudShellV1UsingExternalConfig(ibmCloudShellServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(ibmCloudShellService).ToNot(BeNil())
		})
	})

	Describe(`IbmCloudShellV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSettingsByID request example`, func() {
			fmt.Println("\nGetAccountSettingsByID() result:")
			// begin-get_account_settings_by_id

			getAccountSettingsByIdOptions := ibmCloudShellService.NewGetAccountSettingsByIdOptions(
				AccountID: core.StringPtr("12345678-abcd-1a2b-a1b2-1234567890ab"),
			)

			accountSettings, response, err := ibmCloudShellService.GetAccountSettingsByID(getAccountSettingsByIdOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_account_settings_by_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
		It(`UpdateAccountSettingsByID request example`, func() {
			fmt.Println("\nUpdateAccountSettingsByID() result:")
			// begin-update_account_settings_by_id

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
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-update_account_settings_by_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
	})
})
