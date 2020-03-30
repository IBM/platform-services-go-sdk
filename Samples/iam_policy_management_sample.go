package sample

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/iampolicymanagementv1"
)

var policyService *iampolicymanagementv1.IamPolicyManagementV1
var policyServiceErr error

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
	policyService, policyServiceErr = iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
		Authenticator: authenticator,
	})

	if policyServiceErr == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		policyService.Service.SetDefaultHeaders(customHeaders)
	}

	// uncomment bellow lines to run functions

	//policyID := createPolicy()
	//eTag := getPolicy(policyID)
	// add manager role to the created policy which already has reader role
	//updatePolicy(eTag, policyID)
	//listPolicies()
	//deletePolicy(policyID)
}

func CreatePolicy() string {
	// hardcode required parameters for subject, role, and resources
	typeOfPolicy := "access"
	accIDKey := "accountId"
	accountIDValue := "91631433ee674cd9ab0ef150b8e7030f"
	//	role : reader
	roleID := "crn:v1:bluemix:public:iam::::serviceRole:Reader"
	//	parameters for subject
	//	in this sample, the subject is a Service ID
	subjectAttributeKey := "iam_id"
	subjectAttributeValue := "iam-ServiceId-251553c1-8ae0-4144-8d31-04a26c7cede2"
	//	parameters for resources
	//  in this sample, the resource is a COS instance
	resServiceKey := "serviceName"
	resServiceNameValue := "cloud-object-storage"
	resInstanceKey := "serviceInstance"
	resInstanceValue := "bbe0175e-2377-4fba-8231-fd9da8df8502"

	// create a PolicyRole variable for role
	policyRole := iampolicymanagementv1.PolicyRequestRolesItem{
		RoleID: &roleID,
	}

	// create an SubjectAttributes variable for subject
	mySubject := iampolicymanagementv1.PolicyRequestSubjectsItem{
		Attributes: []iampolicymanagementv1.PolicyRequestSubjectsItemAttributesItem{
			iampolicymanagementv1.PolicyRequestSubjectsItemAttributesItem{
				Name:  &subjectAttributeKey,
				Value: &subjectAttributeValue,
			},
		},
	}

	// create an ResourceAttributes variable for resources
	myResource := iampolicymanagementv1.PolicyRequestResourcesItem{
		Attributes: []iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem{
			iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem{
				Name:  &accIDKey,
				Value: &accountIDValue,
			},
			iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem{
				Name:  &resServiceKey,
				Value: &resServiceNameValue,
			},
			iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem{
				Name:  &resInstanceKey,
				Value: &resInstanceValue,
			},
		},
	}

	// create a CreatePolicyOptions variable pointing to above variables
	// the CreatePolicyOptions variable will be passed into function CreatePolicy
	createPolicyOptions := &iampolicymanagementv1.CreatePolicyOptions{
		Type: &typeOfPolicy,
		Subjects: []iampolicymanagementv1.PolicyRequestSubjectsItem{mySubject},
		Roles: []iampolicymanagementv1.PolicyRequestRolesItem{policyRole},
		Resources: []iampolicymanagementv1.PolicyRequestResourcesItem{myResource},
	}

	// calling function CreatePolicy to make a request for a new policy
	returnValue, detailedResponse, returnValueErr := policyService.CreatePolicy(createPolicyOptions)

	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	fmt.Println(detailedResponse)

	// return policy_id of the created policy
	return *returnValue.ID
}

func GetPolicy(policyID string) string {

	//	create a GetPolicyOptions variable pointing to required parameters
	//	GetPolicy requires only Policy ID as a parameter
	getPolicyOptions := policyService.NewGetPolicyOptions(policyID)

	// calling GetPolicy to make a request to get a specific policy
	_, detailedResponse, returnValueErr := policyService.GetPolicy(getPolicyOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}
	// 	get ETag
	//  eTag is required to update a policy
	//	we get eTag by making request GetPolicy
	eTag := detailedResponse.Headers.Get("Etag")

	fmt.Println(detailedResponse)

	// return eTag
	return eTag
}

func UpdatePolicy(eTag string, policyID string) {

	//	hardcode required parameters that UpdatePolicy needs to make a request
	typeOfPolicy := "access"
	accIDKey := "accountId"
	accountIDValue := "91631433ee674cd9ab0ef150b8e7030f"
	//	parameter of roles
	roleIDWrite := "crn:v1:bluemix:public:iam::::serviceRole:Writer"
	roleIDManager := "crn:v1:bluemix:public:iam::::serviceRole:Manager"
	//	parameters of subject
	subjectAttributeKey := "iam_id"
	subjectAttributeValue := "iam-ServiceId-251553c1-8ae0-4144-8d31-04a26c7cede2"
	//	parameters of resource
	resAttributeKey := "resourceGroupId"
	resAttributeValue := "96714c6a77b04bd39d042a265428aea4"

	// create a UpdatePolicyOptions variable pointing to above parameters
	// the UpdatePolicyOptions variable will be passed into function UpdatePolicy
	updatePolicyOptions := &iampolicymanagementv1.UpdatePolicyOptions{
		IfMatch: &eTag, 
		PolicyID: &policyID,
		Type: &typeOfPolicy,
		Subjects: []iampolicymanagementv1.PolicyRequestSubjectsItem{
			iampolicymanagementv1.PolicyRequestSubjectsItem{
				Attributes: []iampolicymanagementv1.PolicyRequestSubjectsItemAttributesItem{
					iampolicymanagementv1.PolicyRequestSubjectsItemAttributesItem{
						Name:  &subjectAttributeKey,
						Value: &subjectAttributeValue,
					},
				},
			},
		},
		Roles: []iampolicymanagementv1.PolicyRequestRolesItem{
			iampolicymanagementv1.PolicyRequestRolesItem{
				RoleID: &roleIDWrite,
			}, 
			iampolicymanagementv1.PolicyRequestRolesItem{
				RoleID: &roleIDManager,
			},
		},
		Resources: []iampolicymanagementv1.PolicyRequestResourcesItem{
			iampolicymanagementv1.PolicyRequestResourcesItem{
				Attributes: []iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem{
					iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem{
						Name:  &accIDKey,
						Value: &accountIDValue,
					},
					iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem{
						Name:  &resAttributeKey,
						Value: &resAttributeValue,
					},
				},
			},
		},
	}

	//	call function UpdatePolicy to make a request
	_, detailedResponse, returnValueErr := policyService.UpdatePolicy(updatePolicyOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	fmt.Println(detailedResponse)
}

func ListPolicies() {
	// hardcode required parameters that ListPolicies needs to make a request
	accountID := "91631433ee674cd9ab0ef150b8e7030f"
	// create an ListPoliciesOptions object, which points to required parameters
	listPolicyOptions := policyService.NewListPoliciesOptions(accountID)

	// call function ListPolicies
	_, detailedResponse, returnValueErr := policyService.ListPolicies(listPolicyOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	fmt.Println(detailedResponse)
}

func DeletePolicy(policyID string) {
	// create an DeletePolicyOptions variable, which points to required parameters
	deletePolicyOptions := policyService.NewDeletePolicyOptions(policyID)
	// call function DeletePolicy
	returnValue, returnValueErr := policyService.DeletePolicy(deletePolicyOptions)
	if returnValueErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}
	// returnValue has only StatusCode and does not have any body

	statusCode := returnValue.GetStatusCode()
	fmt.Println(statusCode)
}
