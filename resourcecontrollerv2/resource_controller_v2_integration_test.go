// +build integrationSKIP

package resourcecontrollerv2_test

/**
 * Copyright 2020 IBM All Rights Reserved.
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
	"os"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/resourcecontrollerv2"
)

var service *resourcecontrollerv2.ResourceControllerV2
var serviceErr error

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create the authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
	}

	service, serviceErr = resourcecontrollerv2.NewResourceControllerV2(
		&resourcecontrollerv2.ResourceControllerV2Options{
			Authenticator: authenticator,
		})
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}

func TestResourceInstanceOperations(t *testing.T) {
	shouldSkipTest(t)

	// List of resource instances - Test 1
	_, _, returnValueErr := service.ListResourceInstances(nil)
	assert.NotNil(t, returnValueErr)

	listResourceInstancesOptions := service.NewListResourceInstancesOptions()
	listResourceInstancesOptions.SetGuid("")
	listResourceInstancesOptions.SetName("")
	listResourceInstancesOptions.SetResourceGroupID("")
	listResourceInstancesOptions.SetResourceID("")
	listResourceInstancesOptions.SetResourcePlanID("")

	result, response, reqErr := service.ListResourceInstances(listResourceInstancesOptions)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Nil(t, reqErr)

	// Create a new resource instance - Test 2
	name := "test-integration-cos-instance"
	target := "bluemix-global"
	resourceGroup := os.Getenv("RESOURCE_GROUP")
	resourcePlanID := os.Getenv("RESOURCE_PLAN_ID")

	rcInstanceOptions := service.NewCreateResourceInstanceOptions(name, target, resourceGroup, resourcePlanID)
	rcInstanceOptions.SetTags([]string{"integration-test"})
	rcInstanceOptions.SetParameters(map[string]interface{}{})
	rcInstanceOptions.SetHeaders(map[string]string{"Content-Type": "application/json"})

	assert.NotNil(t, rcInstanceOptions)

	resourceInstance, response, reqErr := service.CreateResourceInstance(rcInstanceOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, resourceInstance)
	assert.NotNil(t, response)
	assert.Equal(t, 201, response.GetStatusCode())

	// Get a resource by ID - Test 3
	resourceID := resourceInstance.Guid

	getResourceByIdOptions := service.NewGetResourceInstanceOptions(*resourceID)

	resourceInstance, response, reqErr = service.GetResourceInstance(getResourceByIdOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, resourceInstance)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())

	// Update resource instance by ID - Test 4
	name = "update-integration-cos-instance"
	target = "bluemix-global"
	resourcePlanID = os.Getenv("RESOURCE_PLAN_ID")

	updateResourceInstaceOptions := service.NewUpdateResourceInstanceOptions(*resourceID)
	updateResourceInstaceOptions.SetName(name)
	updateResourceInstaceOptions.SetResourcePlanID(resourcePlanID)
	updateResourceInstaceOptions.SetParameters(map[string]interface{}{})

	resourceInstance, response, reqErr = service.UpdateResourceInstance(updateResourceInstaceOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, resourceInstance)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())

	updateName := resourceInstance.Name
	assert.Equal(t, name, updateName)

	// Delete resource instance - Test 5
	time.Sleep(20 * time.Second) // Allows time for the update of the resource
	deleteResourceInstanceOptions := service.NewDeleteResourceInstanceOptions(*resourceID)

	resourceInstance, response, reqErr = service.DeleteResourceInstance(deleteResourceInstanceOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, response)
	
	// Apparently, the DeleteResourceInstance operation could return either a 202 or a 204 status code.
	// For 202, the API definition defines the response as a ResourceInstance.
	// For 204, there is no response object.
	// assert.Equal(t, 204, response.GetStatusCode())
	if response.GetStatusCode() == 204 {
		assert.Nil(t, resourceInstance)
	} else {
		assert.Equal(t, 202, response.GetStatusCode())
		assert.NotNil(t, resourceInstance)
	}
}

func TestResourceKeysOperations(t *testing.T) {
	shouldSkipTest(t)

	// List of resource instances - Test 1
	_, _, returnValueErr := service.ListResourceKeys(nil)
	assert.NotNil(t, returnValueErr)

	listResourceKeysOptions := service.NewListResourceKeysOptions()
	listResourceKeysOptions.SetGuid("")
	listResourceKeysOptions.SetName("")
	listResourceKeysOptions.SetResourceGroupID("")
	listResourceKeysOptions.SetResourceID("")

	result, response, reqErr := service.ListResourceKeys(listResourceKeysOptions)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Nil(t, reqErr)

	// Create a new resource instance - Preq for Test 2
	name := "test-integration-cos-instance"
	target := "bluemix-global"
	resourceGroup := os.Getenv("RESOURCE_GROUP")
	resourcePlanID := os.Getenv("RESOURCE_PLAN_ID")

	rcInstanceOptions := service.NewCreateResourceInstanceOptions(name, target, resourceGroup, resourcePlanID)
	rcInstanceOptions.SetTags([]string{"integration-test"})

	resourceInstance, response, reqErr := service.CreateResourceInstance(rcInstanceOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, resourceInstance)
	assert.NotNil(t, response)
	assert.Equal(t, 201, response.GetStatusCode())

	resourceID := resourceInstance.Guid
	resourceCrn := resourceInstance.Crn
	
	roleCrn := "Writer"
	name = "integration-test-key"

	rcKeyOptions := service.NewCreateResourceKeyOptions(name, *resourceCrn)
	rcKeyOptions.SetRole(roleCrn)

	resourceKey, response, reqErr := service.CreateResourceKey(rcKeyOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, resourceKey)
	assert.NotNil(t, response)
	assert.Equal(t, 201, response.GetStatusCode())

	// Get resource key by ID - Test 3
	resourceKeyID := resourceKey.Guid

	getRCKeyOptions := service.NewGetResourceKeyOptions(*resourceKeyID)

	resourceKey, response, reqErr = service.GetResourceKey(getRCKeyOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, resourceKey)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, name, *resourceKey.Name)

	// Update resource key by ID - Test 4
	updateName := "update-test-key"
	updateRCKeyOptions := service.NewUpdateResourceKeyOptions(*resourceKeyID, updateName)

	resourceKey, response, reqErr = service.UpdateResourceKey(updateRCKeyOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, resourceKey)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, updateName, *resourceKey.Name)

	// Delete resource key - Test 5
	deleteRCKeyOptions := service.NewDeleteResourceKeyOptions(*resourceKeyID)

	response, reqErr = service.DeleteResourceKey(deleteRCKeyOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())

	// Delete resource instance - Postreq
	time.Sleep(20 * time.Second) // Allows time for the delete operation
	deleteResourceInstanceOptions := service.NewDeleteResourceInstanceOptions(*resourceID)

	resourceInstance, response, reqErr = service.DeleteResourceInstance(deleteResourceInstanceOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
}
