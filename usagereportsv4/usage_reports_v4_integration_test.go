// +build integrationSKIP

package usagereportsv4_test

/**
 * (C) Copyright IBM Corp. 2020.
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

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/usagereportsv4"
)

var service *usagereportsv4.UsageReportsV4
var serviceErr error

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	// Create the authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
	}

	service, serviceErr = usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
		Authenticator: authenticator,
	})

	if serviceErr == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		service.Service.SetDefaultHeaders(customHeaders)
	}
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}

func TestGetAccountSummary(t *testing.T) {
	shouldSkipTest(t)
	//tokenType := "BEARER"
	//accountID := "91631433ee674cd9ab0ef150b8e7030f"
	accountID := "a5d2a90b46ac45089fa28d0cb74c49af"
	billingmonth := "2019-05"

	getAccountSummaryOptions := service.NewGetAccountSummaryOptions(accountID, billingmonth)
	accountSummary, returnValue, returnValueErr := service.GetAccountSummary(getAccountSummaryOptions)

	assert.Nil(t, returnValueErr)
	assert.NotNil(t, accountSummary)
	assert.NotNil(t, returnValue)
	assert.Equal(t, 200, returnValue.StatusCode)
}

func TestGetAccountUsage(t *testing.T) {
	shouldSkipTest(t)
	//accountID := "91631433ee674cd9ab0ef150b8e7030f"
	accountID := "a5d2a90b46ac45089fa28d0cb74c49af"
	billingmonth := "2019-05"

	getAccountUsageOptions := service.NewGetAccountUsageOptions(accountID, billingmonth)
	accountUsage, returnValue, returnValueErr := service.GetAccountUsage(getAccountUsageOptions)

	assert.Nil(t, returnValueErr)
	assert.NotNil(t, accountUsage)
	assert.NotNil(t, returnValue)
	assert.Equal(t, 200, returnValue.StatusCode)
}

func TestGetResourceGroupUsage(t *testing.T) {
	shouldSkipTest(t)
	accountID := "91631433ee674cd9ab0ef150b8e7030f"
	resourceGroupID := "96714c6a77b04bd39d042a265428aea4"
	billingmonth := "2019-05"

	getResourceGroupUsageOptions := service.NewGetResourceGroupUsageOptions(accountID, resourceGroupID, billingmonth)
	resGrpUsage, returnValue, returnValueErr := service.GetResourceGroupUsage(getResourceGroupUsageOptions)
	assert.Nil(t, returnValueErr)
	assert.NotNil(t, resGrpUsage)
	assert.NotNil(t, returnValue)
	assert.Equal(t, 200, returnValue.StatusCode)
}

func TestGetOrganizationUsage(t *testing.T) {
	shouldSkipTest(t)
	//accountID := "91631433ee674cd9ab0ef150b8e7030f"
	accountID := "a5d2a90b46ac45089fa28d0cb74c49af"
	//organizationID := "96714c6a77b04bd39d042a265428aea4"
	organizationID := "us-south:c9d7aff6-a1ae-4916-912d-42fbc25e0e1d"
	//organizationID := "a5d2a90b46ac45089fa28d0cb74c49af"
	billingmonth := "2019-05"

	getOrganizationUsageOptions := service.NewGetOrganizationUsageOptions(accountID, organizationID, billingmonth)
	orgUsage, returnValue, returnValueErr := service.GetOrganizationUsage(getOrganizationUsageOptions)
	assert.Nil(t, returnValueErr)
	assert.NotNil(t, orgUsage)
	assert.NotNil(t, returnValue)
	assert.Equal(t, 200, returnValue.StatusCode)
}

func TestGetAccountInstancesUsage(t *testing.T) {
	shouldSkipTest(t)
	//accountID := "91631433ee674cd9ab0ef150b8e7030f"
	accountID := "a5d2a90b46ac45089fa28d0cb74c49af"
	//organizationID := "96714c6a77b04bd39d042a265428aea4"
	billingmonth := "2019-05"

	getAccountInstancesUsageOptions := service.NewGetAccountInstancesUsageOptions(accountID, billingmonth)
	acctUsage, returnValue, returnValueErr := service.GetAccountInstancesUsage(getAccountInstancesUsageOptions)
	assert.Nil(t, returnValueErr)
	assert.NotNil(t, acctUsage)
	assert.NotNil(t, returnValue)
	assert.Equal(t, 200, returnValue.StatusCode)
}

func TestGetResourceGroupInstancesUsage(t *testing.T) {
	shouldSkipTest(t)
	//accountID := "91631433ee674cd9ab0ef150b8e7030f"
	accountID := "91631433ee674cd9ab0ef150b8e7030f"
	resourceGroupID := "96714c6a77b04bd39d042a265428aea4"
	billingmonth := "2019-05"

	getResourceGroupInstancesUsageOptions := service.NewGetResourceGroupInstancesUsageOptions(accountID, resourceGroupID, billingmonth)
	resGrpUsage, returnValue, returnValueErr := service.GetResourceGroupInstancesUsage(getResourceGroupInstancesUsageOptions)
	assert.Nil(t, returnValueErr)
	assert.NotNil(t, resGrpUsage)
	assert.NotNil(t, returnValue)
	assert.Equal(t, 200, returnValue.StatusCode)
}

func TestGetOrganizationInstancesUsage(t *testing.T) {
	shouldSkipTest(t)
	//accountID := "91631433ee674cd9ab0ef150b8e7030f"
	accountID := "a5d2a90b46ac45089fa28d0cb74c49af"
	//organizationID := "96714c6a77b04bd39d042a265428aea4"
	organizationID := "us-south:c9d7aff6-a1ae-4916-912d-42fbc25e0e1d"
	billingmonth := "2019-05"

	getOrganizationInstancesUsageOptions := service.NewGetOrganizationInstancesUsageOptions(accountID, organizationID, billingmonth)
	orgUsage, returnValue, returnValueErr := service.GetOrganizationInstancesUsage(getOrganizationInstancesUsageOptions)
	assert.Nil(t, returnValueErr)
	assert.NotNil(t, orgUsage)
	assert.NotNil(t, returnValue)
	assert.Equal(t, 200, returnValue.StatusCode)
}
