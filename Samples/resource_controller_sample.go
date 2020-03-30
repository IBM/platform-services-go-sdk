package sample

/**
 * Copyright 2019 IBM All Rights Reserved.
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
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/resourcecontrollerv2"
)

// Resource Controller Service Instance
var rcService *resourcecontrollerv2.ResourceControllerV2
var rcServiceErr error

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create the IAM authenticator.
	authenticator := &core.IamAuthenticator{
	    ApiKey: os.Getenv("IAMAPIKEY"),
	}

	if err == nil {
		rcService, rcServiceErr = resourcecontrollerv2.
			NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				Authenticator: authenticator, // Get APIKey from your .env file
			})

		if rcServiceErr == nil {
			customHeaders := http.Header{}
			customHeaders.Add("Content-type", "application/json")
			rcService.Service.SetDefaultHeaders(customHeaders)
		}
	}
}

/************************************** Resource Instance Operations ****************************/

// listAllResourceInstance - Gets the list of all resource instances
func ListAllResourceInstance() {
	listResourceInstancesOptions := rcService.NewListResourceInstancesOptions()

	// Set neccessary parameters
	listResourceInstancesOptions.SetGuid("")

	listResourceInstancesOptions.SetHeaders(map[string]string{})

	listResourceInstancesOptions.SetName("")

	listResourceInstancesOptions.SetResourceGroupID("")

	listResourceInstancesOptions.SetResourceID("")

	listResourceInstancesOptions.SetResourcePlanID("")

	_, detailedResponse, reqErr := rcService.ListResourceInstances(listResourceInstancesOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// createNewResourceInstance - Creates a new resource instance
func CreateNewResourceInstance(name string, target string) {
	// ex name := "test-integration-cos-instance"
	// target refers to the depoyment location
	// ex target := "bluemix-global"

	resourceGroup := os.Getenv("RESOURCE_GROUP")
	resourcePlanID := os.Getenv("RESOURCE_PLAN_ID")

	rcInstanceOptions := rcService.NewCreateResourceInstanceOptions(name, target, resourceGroup, resourcePlanID)
	rcInstanceOptions.SetTags([]string{"integration-test"})
	rcInstanceOptions.SetParameters(map[string]interface{}{})
	rcInstanceOptions.SetHeaders(map[string]string{"Content-Type": "application/json"})

	_, detailedResponse, reqErr := rcService.CreateResourceInstance(rcInstanceOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// getResourceInstance - Gets resource instance based on ID
func GetResourceInstance(resourceID string) {

	getResourceByIDOptions := rcService.NewGetResourceInstanceOptions(resourceID)
	getResourceByIDOptions.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	_, detailedResponse, reqErr := rcService.GetResourceInstance(getResourceByIDOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// updateResourceInstance - Updates resource instance based on ID
func UpdateResourceInstance(name string, target string, resourceID string) {
	// target refers to the depoyment location
	// ex target := "bluemix-global"

	resourcePlanID := os.Getenv("RESOURCE_PLAN_ID")

	updateResourceInstaceOptions := rcService.NewUpdateResourceInstanceOptions(resourceID)
	updateResourceInstaceOptions.SetName(name)
	updateResourceInstaceOptions.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	updateResourceInstaceOptions.SetResourcePlanID(resourcePlanID)
	updateResourceInstaceOptions.SetParameters(map[string]interface{}{})

	_, detailedResponse, reqErr := rcService.UpdateResourceInstance(updateResourceInstaceOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// deleteResourceInstance - delete resource instance based on ID
func DeleteResourceInstance(resourceID string) {

	deleteResourceInstanceOptions := rcService.NewDeleteResourceInstanceOptions(resourceID)

	_, detailedResponse, reqErr := rcService.DeleteResourceInstance(deleteResourceInstanceOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

/************************************** Resource Keys Operations ****************************/

// listAllResourceKeys - Gets all resource keys
func ListAllResourceKeys() {

	listResourceKeysOptions := rcService.NewListResourceKeysOptions()
	listResourceKeysOptions.SetGuid("")

	listResourceKeysOptions.SetHeaders(map[string]string{})

	listResourceKeysOptions.SetName("")

	listResourceKeysOptions.SetResourceGroupID("")

	listResourceKeysOptions.SetResourceID("")

	_, detailedResponse, reqErr := rcService.ListResourceKeys(listResourceKeysOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// createNewResourceKey - Creates a new resource key(credentials) for a resource instance
func CreateNewResourceKey(resourceInstance *resourcecontrollerv2.ResourceInstance, role string, name string) {
	// Resource keys can be created for a resource instances, but the resource instances have to exist

	resourceCrn := *resourceInstance.Crn
	rcKeyOptions := rcService.NewCreateResourceKeyOptions(name, resourceCrn)

	rcKeyOptions.SetRole(role)
	rcKeyOptions.SetHeaders(map[string]string{"Content-Type": "application/json"})

	_, detailedResponse, reqErr := rcService.CreateResourceKey(rcKeyOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// getResourceKey - Get resource key by ID
func GetResourceKey(resourceKeyID string) {

	getRCKeyOptions := rcService.NewGetResourceKeyOptions(resourceKeyID)
	getRCKeyOptions.SetHeaders(map[string]string{"Content-Type": "application/json"})

	_, detailedResponse, reqErr := rcService.GetResourceKey(getRCKeyOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// updateResourceKey - update resource key by ID
func UpdateResourceKey(resourceKeyID string, updateName string) {

	updateRCKeyOptions := rcService.NewUpdateResourceKeyOptions(resourceKeyID, updateName)

	_, detailedResponse, reqErr := rcService.UpdateResourceKey(updateRCKeyOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// deleteResourceKey - delete resource key by ID
func DeleteResourceKey(resourceKeyID string) {

	deleteRCKeyOptions := rcService.NewDeleteResourceKeyOptions(resourceKeyID)

	response, reqErr := rcService.DeleteResourceKey(deleteRCKeyOptions)

	if reqErr == nil {
		fmt.Println(response.String())
	} else {
		fmt.Println(reqErr)
	}
}
