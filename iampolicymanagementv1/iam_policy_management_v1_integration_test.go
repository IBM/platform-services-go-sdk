// +build integrationSKIP

package iampolicymanagementv1_test

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

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/iampolicymanagementv1"
)

var service *iampolicymanagementv1.IamPolicyManagementV1
var serviceErr error
var policyID string
var eTag string

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	// Create the authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
	}

	service, serviceErr = iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
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

func TestCreatePolicy(t *testing.T) {
	shouldSkipTest(t)

	// hardcode parameters
	typeAccess := "access"
	accIDKey := "accountId"
	accountID := "91631433ee674cd9ab0ef150b8e7030f"
	resAttKey := "resourceGroupId"
	resAttValue := "96714c6a77b04bd39d042a265428aea4"
	roleID := "crn:v1:bluemix:public:iam::::serviceRole:Writer"
	subjectAttKey := "iam_id"
	subjectAttValue := "iam-ServiceId-251553c1-8ae0-4144-8d31-04a26c7cede2"

	// Create an instance of the options model so we can invoke CreatePolicy.
	createPolicyOptionsModel := &iampolicymanagementv1.CreatePolicyOptions{
		Type: &typeAccess,
		Subjects: []iampolicymanagementv1.SubjectAttributes{
			iampolicymanagementv1.SubjectAttributes{
				Attributes: []iampolicymanagementv1.Attribute{
					iampolicymanagementv1.Attribute{
						Name:  &subjectAttKey,
						Value: &subjectAttValue,
					},
				},
			},
		},
		Roles: []iampolicymanagementv1.PolicyRole{
			iampolicymanagementv1.PolicyRole{
				RoleID: &roleID,
			},
		},
		Resources: []iampolicymanagementv1.ResourceAttributes{
			iampolicymanagementv1.ResourceAttributes{
				Attributes: []iampolicymanagementv1.ResourceAttribute{
					iampolicymanagementv1.ResourceAttribute{
						Name:  &accIDKey,
						Value: &accountID,
					},
					iampolicymanagementv1.ResourceAttribute{
						Name:  &resAttKey,
						Value: &resAttValue,
					},
				},
			},
		},
	}

	// Invoke the CreatePolicy operation.
	result, returnValue, returnValueErr := service.CreatePolicy(createPolicyOptionsModel)
	assert.Nil(t, returnValueErr)
	assert.NotNil(t, returnValue)
	assert.NotNil(t, result)
	assert.Equal(t, 201, returnValue.StatusCode)

	// save policy_id of the created policy
	policyID = *result.ID
}

func TestGetPolicy(t *testing.T) {
	shouldSkipTest(t)

	// Invoke the GetPolicy operation.
	result, returnValue, returnValueErr := service.GetPolicy(nil)
	assert.Nil(t, returnValueErr)
	assert.NotNil(t, returnValue)
	assert.NotNil(t, result)
	assert.Equal(t, 200, returnValue.StatusCode)

	// Retrieve the ETag header from the response.
	eTag = returnValue.Headers.Get("Etag")
	assert.NotNil(t, eTag)
}

func TestUpdatePolicy(t *testing.T) {
	shouldSkipTest(t)

	// hardcode parameters
	typeAccess := "access"
	accIDKey := "accountId"
	accountID := "91631433ee674cd9ab0ef150b8e7030f"
	resAttKey := "resourceGroupId"
	resAttValue := "96714c6a77b04bd39d042a265428aea4"
	roleID_writer := "crn:v1:bluemix:public:iam::::serviceRole:Writer"
	roleID_manager := "crn:v1:bluemix:public:iam::::serviceRole:Manager"
	subjectAttKey := "iam_id"
	subjectAttValue := "iam-ServiceId-251553c1-8ae0-4144-8d31-04a26c7cede2"

	// Create an instance of the options model so we can invoke UpdatePolicy.
	updatePolicyOptionsModel := &iampolicymanagementv1.UpdatePolicyOptions{
		Type: &typeAccess,
		Subjects: []iampolicymanagementv1.SubjectAttributes{
			iampolicymanagementv1.SubjectAttributes{
				Attributes: []iampolicymanagementv1.Attribute{
					iampolicymanagementv1.Attribute{
						Name:  &subjectAttKey,
						Value: &subjectAttValue,
					},
				},
			},
		},
		Roles: []iampolicymanagementv1.PolicyRole{
			iampolicymanagementv1.PolicyRole{
				RoleID: &roleID_writer,
			},
			iampolicymanagementv1.PolicyRole{
				RoleID: &roleID_manager,
			},
		},
		Resources: []iampolicymanagementv1.ResourceAttributes{
			iampolicymanagementv1.ResourceAttributes{
				Attributes: []iampolicymanagementv1.ResourceAttribute{
					iampolicymanagementv1.ResourceAttribute{
						Name:  &accIDKey,
						Value: &accountID,
					},
					iampolicymanagementv1.ResourceAttribute{
						Name:  &resAttKey,
						Value: &resAttValue,
					},
				},
			},
		},
	}

	// Invoke the UpdatePolicy operation.
	result, returnValue, returnValueErr := service.UpdatePolicy(updatePolicyOptionsModel)
	assert.NotNil(t, result)
	assert.Nil(t, returnValueErr)
	assert.NotNil(t, returnValue)
	assert.Equal(t, 200, returnValue.StatusCode)
}

func TestListPolicies(t *testing.T) {
	shouldSkipTest(t)

	//	hardcode parameters
	accountID := "91631433ee674cd9ab0ef150b8e7030f"

	// Invoke the ListPolicies operation.
	result, returnValue, returnValueErr := service.ListPolicies(
		&iampolicymanagementv1.ListPoliciesOptions{
			AccountID: &accountID,
		},
	)
	assert.Nil(t, returnValueErr)
	assert.NotNil(t, result)
	assert.NotNil(t, returnValue)
	assert.Equal(t, 200, returnValue.StatusCode)
}

func TestDeletePolicy(t *testing.T) {
	shouldSkipTest(t)
	
	// Invoke the DeletePolicy operation, using the policy id set above by CreatePolicy().
	returnValue, returnValueErr := service.DeletePolicy(service.NewDeletePolicyOptions(policyID))
	assert.Nil(t, returnValueErr)
	assert.NotNil(t, returnValue)
	assert.Equal(t, 204, returnValue.StatusCode)
}
