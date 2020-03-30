package sample

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/usagereportsv4"
)

var usageReportService *usagereportsv4.UsageReportsV4
var usageReportServiceErr error

func init() {

	// loading environment variables from file .env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	// Create the IAM authenticator.
	authenticator := &core.IamAuthenticator{
			ApiKey: os.Getenv("IAMAPIKEY"),
	}

	// create a new service of usage report
	// APIKEY of an account is needed to access services on that account
	usageReportService, usageReportServiceErr = usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
		Authenticator: authenticator,
	})

	if usageReportServiceErr == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		usageReportService.Service.SetDefaultHeaders(customHeaders)
	}

	// uncomment below lines to run functions

	//getAccountSummary()
	//getAccountUsage()
	//getResourceGroupUsage()
	//getOrganizationUsage()
	//getAccountInstancesUsage()
	//getResourceGroupInstancesUsage()
	//getOrganizationInstancesUsage()

}

func GetAccountSummary() {

	// hardcode variables that are required parameters for GetAccountSummary request
	accountID := "a5d2a90b46ac45089fa28d0cb74c49af"
	// billingMonth has format "yyyy-mm"
	billingMonth := "2019-05"

	// create a GetAccountSummaryOptions struct pointing to above parameters
	// a GetAccountSummaryOptions object will be passed into function GetAccountSummary
	getAccountSummaryOptions := usageReportService.NewGetAccountSummaryOptions(accountID, billingMonth)

	// making a request by calling GetAccountSummary function
	_, detailedResponse, returnValueErr := usageReportService.GetAccountSummary(getAccountSummaryOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	fmt.Println(detailedResponse)
}

func GetAccountUsage() {

	// hardcode variables that are required parameters for GetAccountUsage request
	accountID := "a5d2a90b46ac45089fa28d0cb74c49af"
	// billingMonth has format "yyyy-mm"
	billingmonth := "2019-05"

	// create a GetAccountUsageOptions variable pointing to above parameters
	// a GetAccountUsageOptions object will be passed into function GetAccountUsage
	getAccountUsageOptions := usageReportService.NewGetAccountUsageOptions(accountID, billingmonth)

	// making a request by calling GetAccountUsage function
	_, detailedResponse, returnValueErr := usageReportService.GetAccountUsage(getAccountUsageOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	fmt.Println(detailedResponse)
}

func GetResourceGroupUsage() {

	// hardcode variables that are required parameters for GetResourceGroupUsage request
	accountID := "91631433ee674cd9ab0ef150b8e7030f"
	resourceGroupID := "96714c6a77b04bd39d042a265428aea4"
	// billingMonth has format "yyyy-mm"
	billingmonth := "2019-05"

	// create a GetResourceGroupUsageOptions variable pointing to above parameters
	// a GetResourceGroupUsageOptions object will be passed into function GetResourceGroupUsage
	getResourceGroupUsageOptions := usageReportService.NewGetResourceGroupUsageOptions(accountID, resourceGroupID, billingmonth)

	// making a request by calling GetResourceGroupUsage function
	_, detailedResponse, returnValueErr := usageReportService.GetResourceGroupUsage(getResourceGroupUsageOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	fmt.Println(detailedResponse)
}

func GetOrganizationUsage() {
	// hardcode variables that are required parameters for GetOrganizationUsage request
	accountID := "a5d2a90b46ac45089fa28d0cb74c49af"
	organizationID := "us-south:c9d7aff6-a1ae-4916-912d-42fbc25e0e1d"
	// billingMonth has format "yyyy-mm"
	billingmonth := "2019-05"

	// create a GetOrganizationUsageOptions variable pointing to above parameters
	// a GetOrganizationUsageOptions object will be passed into function GetOrganizationUsage
	getOrganizationUsageOptions := usageReportService.NewGetOrganizationUsageOptions(accountID, organizationID, billingmonth)

	// making a request by calling GetOrganizationUsage function
	_, detailedResponse, returnValueErr := usageReportService.GetOrganizationUsage(getOrganizationUsageOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	fmt.Println(detailedResponse)
}

func GetAccountInstancesUsage(t *testing.T) {
	// hardcode variables that are required parameters for GetAccountInstancesUsage request
	accountID := "a5d2a90b46ac45089fa28d0cb74c49af"
	// billingMonth has format "yyyy-mm"
	billingmonth := "2019-05"

	// create a GetOrganizationUsageOptions variable pointing to above parameters
	// a GetOrganizationUsageOptions object will be passed into function GetOrganizationUsage
	getAccountInstancesUsageOptions := usageReportService.NewGetAccountInstancesUsageOptions(accountID, billingmonth)

	// making request by calling GetAccountInstancesUsage function
	_, detailedResponse, returnValueErr := usageReportService.GetAccountInstancesUsage(getAccountInstancesUsageOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	fmt.Println(detailedResponse)
}
func GetResourceGroupInstancesUsage(t *testing.T) {
	// hardcode variables that are required parameters for GetResourceGroupInstancesUsage request
	accountID := "91631433ee674cd9ab0ef150b8e7030f"
	resourceGroupID := "96714c6a77b04bd39d042a265428aea4"
	// billingMonth has format "yyyy-mm"
	billingmonth := "2019-05"

	// create a GetResourceGroupInstancesUsageOptions variable pointing to above parameters
	// a GetResourceGroupInstancesUsageOptions object will be passed into function GetResourceGroupInstancesUsage
	getResourceGroupInstancesUsageOptions := usageReportService.NewGetResourceGroupInstancesUsageOptions(accountID, resourceGroupID, billingmonth)
	// making request by calling GetResourceGroupInstancesUsage function
	_, detailedResponse, returnValueErr := usageReportService.GetResourceGroupInstancesUsage(getResourceGroupInstancesUsageOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	fmt.Println(detailedResponse)
}

func GetOrganizationInstancesUsage(t *testing.T) {
	// hardcode variables that are required parameters for GetOrganizationInstancesUsage request
	accountID := "a5d2a90b46ac45089fa28d0cb74c49af"
	organizationID := "us-south:c9d7aff6-a1ae-4916-912d-42fbc25e0e1d"
	// billingMonth has format "yyyy-mm"
	billingmonth := "2019-05"

	// create a GetOrganizationInstancesUsageOptions variable pointing to above parameters
	// a GetOrganizationInstancesUsageOptions object will be passed into function GetOrganizationInstancesUsage
	getOrganizationInstancesUsageOptions := usageReportService.NewGetOrganizationInstancesUsageOptions(accountID, organizationID, billingmonth)
	// making request by calling GetOrganizationInstancesUsage function
	_, detailedResponse, returnValueErr := usageReportService.GetOrganizationInstancesUsage(getOrganizationInstancesUsageOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	fmt.Println(detailedResponse)
}
