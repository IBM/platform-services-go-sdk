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

package iampolicymanagementv1_test

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"os"
	"strconv"
)

// Below are examples on how to use IAM Policy Management service
//
// The following environment variables are assumed to be defined when running examples below:
//
// IAM_POLICY_MANAGEMENT_URL=https://iam.cloud.ibm.com
// IAM_POLICY_MANAGEMENT_AUTH_TYPE=iam
// IAM_POLICY_MANAGEMENT_AUTH_URL=https://iam.cloud.ibm.com/identity/token
// IAM_POLICY_MANAGEMENT_APIKEY= <YOUR_APIKEY>
// IAM_POLICY_MANAGEMENT_TEST_ACCOUNT_ID= <YOUR_ACCOUNT_ID>
//
// Alternatively, above environment variables can be placed in a "credentials" file

const externalConfigFile = "../iam_policy_management.env"

var (
	// TODO: Align
	iamPolicyManagementService *iampolicymanagementv1.IamPolicyManagementV1
	config       map[string]string
	configLoaded bool = false

	testUserId        = "IBMid-GoSDK" + strconv.Itoa(rand.Intn(100000))
	testServiceName   = "iam-groups"
	testAccountId      string
	testPolicyId       string
	testPolicyETag     string
	testCustomRoleId   string
	testCustomRoleETag string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`IamPolicyManagementV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(iampolicymanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			testAccountId = config["TEST_ACCOUNT_ID"]

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

			iamPolicyManagementServiceOptions := &iampolicymanagementv1.IamPolicyManagementV1Options{}

			iamPolicyManagementService, err = iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(iamPolicyManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(iamPolicyManagementService).ToNot(BeNil())
		})
	})

	Describe(`IamPolicyManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicy request example`, func() {
			// begin-create_policy

			subjectAttribute := &iampolicymanagementv1.SubjectAttribute{
				Name: core.StringPtr("iam_id"),
				Value: &testUserId,
			}
			policySubjects := &iampolicymanagementv1.PolicySubject{
				Attributes: []iampolicymanagementv1.SubjectAttribute{*subjectAttribute},
			}
			policyRoles := &iampolicymanagementv1.PolicyRole{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Viewer"),
			}
			accountIdResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name: core.StringPtr("accountId"),
				Value: core.StringPtr(testAccountId),
				Operator: core.StringPtr("stringEquals"),
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name: core.StringPtr("serviceName"),
				Value: core.StringPtr(testServiceName),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResourceTag := &iampolicymanagementv1.ResourceTag{
				Name: core.StringPtr("project"),
				Value: core.StringPtr("prototype"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResources := &iampolicymanagementv1.PolicyResource{
				Attributes: []iampolicymanagementv1.ResourceAttribute{
					*accountIdResourceAttribute, *serviceNameResourceAttribute},
				Tags: []iampolicymanagementv1.ResourceTag{*policyResourceTag},
			}

			options := iamPolicyManagementService.NewCreatePolicyOptions(
				"access",
				[]iampolicymanagementv1.PolicySubject{*policySubjects},
				[]iampolicymanagementv1.PolicyRole{*policyRoles},
				[]iampolicymanagementv1.PolicyResource{*policyResources},
			)

			policy, response, err := iamPolicyManagementService.CreatePolicy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-create_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())

			testPolicyId = *policy.ID
		})
		It(`GetPolicy request example`, func() {
			// begin-get_policy

			options := iamPolicyManagementService.NewGetPolicyOptions(
				testPolicyId,
			)

			policy, response, err := iamPolicyManagementService.GetPolicy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-get_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())

			testPolicyETag = response.GetHeaders().Get("ETag")
		})
		It(`UpdatePolicy request example`, func() {
			// begin-update_policy

			subjectAttribute := &iampolicymanagementv1.SubjectAttribute{
				Name: core.StringPtr("iam_id"),
				Value: &testUserId,
			}
			policySubjects := &iampolicymanagementv1.PolicySubject{
				Attributes: []iampolicymanagementv1.SubjectAttribute{*subjectAttribute},
			}
			accountIdResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name: core.StringPtr("accountId"),
				Value: core.StringPtr(testAccountId),
				Operator: core.StringPtr("stringEquals"),
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name: core.StringPtr("serviceName"),
				Value: core.StringPtr(testServiceName),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResources := &iampolicymanagementv1.PolicyResource{
				Attributes: []iampolicymanagementv1.ResourceAttribute{
					*accountIdResourceAttribute, *serviceNameResourceAttribute},
			}
			updatedPolicyRoles := &iampolicymanagementv1.PolicyRole{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Editor"),
			}

			options := iamPolicyManagementService.NewUpdatePolicyOptions(
				testPolicyId,
				testPolicyETag,
				"access",
				[]iampolicymanagementv1.PolicySubject{*policySubjects},
				[]iampolicymanagementv1.PolicyRole{*updatedPolicyRoles},
				[]iampolicymanagementv1.PolicyResource{*policyResources},
			)

			policy, response, err := iamPolicyManagementService.UpdatePolicy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-update_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())

		})
		It(`ListPolicies request example`, func() {
			// begin-list_policies

			options := iamPolicyManagementService.NewListPoliciesOptions(
				testAccountId,
			)
			options.SetIamID(testUserId)
			options.SetFormat("include_last_permit")

			policyList, response, err := iamPolicyManagementService.ListPolicies(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyList, "", "  ")
			fmt.Println(string(b))

			// end-list_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyList).ToNot(BeNil())

		})
		It(`DeletePolicy request example`, func() {
			// begin-delete_policy

			options := iamPolicyManagementService.NewDeletePolicyOptions(
				testPolicyId,
			)

			response, err := iamPolicyManagementService.DeletePolicy(options)
			if err != nil {
				panic(err)
			}

			// end-delete_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`CreateRole request example`, func() {
			// begin-create_role

			options := iamPolicyManagementService.NewCreateRoleOptions(
				"IAM Groups read access",
				[]string{"iam-groups.groups.read"},
				"ExampleRoleIAMGroups",
				testAccountId,
				testServiceName,
			)

			customRole, response, err := iamPolicyManagementService.CreateRole(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customRole, "", "  ")
			fmt.Println(string(b))

			// end-create_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(customRole).ToNot(BeNil())

			testCustomRoleId = *customRole.ID
		})
		It(`GetRole request example`, func() {
			// begin-get_role

			options := iamPolicyManagementService.NewGetRoleOptions(
				testCustomRoleId,
			)

			customRole, response, err := iamPolicyManagementService.GetRole(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customRole, "", "  ")
			fmt.Println(string(b))

			// end-get_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customRole).ToNot(BeNil())

			testCustomRoleETag = response.Headers.Get("ETag")

		})
		It(`UpdateRole request example`, func() {
			// begin-update_role

			updatedRoleActions := []string{"iam-groups.groups.read", "iam-groups.groups.list"}

			options := iamPolicyManagementService.NewUpdateRoleOptions(
				testCustomRoleId,
				testCustomRoleETag,
			)
			options.SetActions(updatedRoleActions)

			customRole, response, err := iamPolicyManagementService.UpdateRole(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customRole, "", "  ")
			fmt.Println(string(b))

			// end-update_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customRole).ToNot(BeNil())

		})
		It(`ListRoles request example`, func() {
			// begin-list_roles

			options := iamPolicyManagementService.NewListRolesOptions()
			options.SetAccountID(testAccountId)

			roleList, response, err := iamPolicyManagementService.ListRoles(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(roleList, "", "  ")
			fmt.Println(string(b))

			// end-list_roles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(roleList).ToNot(BeNil())

		})
		It(`DeleteRole request example`, func() {
			// begin-delete_role

			options := iamPolicyManagementService.NewDeleteRoleOptions(
				testCustomRoleId,
			)

			response, err := iamPolicyManagementService.DeleteRole(options)
			if err != nil {
				panic(err)
			}

			// end-delete_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
