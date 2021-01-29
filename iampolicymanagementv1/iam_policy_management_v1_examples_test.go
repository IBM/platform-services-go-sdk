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
	"os"
)

const externalConfigFile = "../iam_policy_management_v1.env"

var (
	iamPolicyManagementService *iampolicymanagementv1.IamPolicyManagementV1
	config       map[string]string
	configLoaded bool = false
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
		It(`ListPolicies request example`, func() {
			// begin-list_policies

			listPoliciesOptions := iamPolicyManagementService.NewListPoliciesOptions(
				"testString",
			)

			policyList, response, err := iamPolicyManagementService.ListPolicies(listPoliciesOptions)
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
		It(`CreatePolicy request example`, func() {
			// begin-create_policy

			policySubjectModel := &iampolicymanagementv1.PolicySubject{
			}

			policyRoleModel := &iampolicymanagementv1.PolicyRole{
				RoleID: core.StringPtr("testString"),
			}

			policyResourceModel := &iampolicymanagementv1.PolicyResource{
			}

			createPolicyOptions := iamPolicyManagementService.NewCreatePolicyOptions(
				"testString",
				[]iampolicymanagementv1.PolicySubject{*policySubjectModel},
				[]iampolicymanagementv1.PolicyRole{*policyRoleModel},
				[]iampolicymanagementv1.PolicyResource{*policyResourceModel},
			)

			policy, response, err := iamPolicyManagementService.CreatePolicy(createPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-create_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())

		})
		It(`UpdatePolicy request example`, func() {
			// begin-update_policy

			policySubjectModel := &iampolicymanagementv1.PolicySubject{
			}

			policyRoleModel := &iampolicymanagementv1.PolicyRole{
				RoleID: core.StringPtr("testString"),
			}

			policyResourceModel := &iampolicymanagementv1.PolicyResource{
			}

			updatePolicyOptions := iamPolicyManagementService.NewUpdatePolicyOptions(
				"testString",
				"testString",
				"testString",
				[]iampolicymanagementv1.PolicySubject{*policySubjectModel},
				[]iampolicymanagementv1.PolicyRole{*policyRoleModel},
				[]iampolicymanagementv1.PolicyResource{*policyResourceModel},
			)

			policy, response, err := iamPolicyManagementService.UpdatePolicy(updatePolicyOptions)
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
		It(`GetPolicy request example`, func() {
			// begin-get_policy

			getPolicyOptions := iamPolicyManagementService.NewGetPolicyOptions(
				"testString",
			)

			policy, response, err := iamPolicyManagementService.GetPolicy(getPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-get_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())

		})
		It(`ListRoles request example`, func() {
			// begin-list_roles

			listRolesOptions := iamPolicyManagementService.NewListRolesOptions()

			roleList, response, err := iamPolicyManagementService.ListRoles(listRolesOptions)
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
		It(`CreateRole request example`, func() {
			// begin-create_role

			createRoleOptions := iamPolicyManagementService.NewCreateRoleOptions(
				"testString",
				[]string{"testString"},
				"testString",
				"testString",
				"testString",
			)

			customRole, response, err := iamPolicyManagementService.CreateRole(createRoleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customRole, "", "  ")
			fmt.Println(string(b))

			// end-create_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(customRole).ToNot(BeNil())

		})
		It(`UpdateRole request example`, func() {
			// begin-update_role

			updateRoleOptions := iamPolicyManagementService.NewUpdateRoleOptions(
				"testString",
				"testString",
			)

			customRole, response, err := iamPolicyManagementService.UpdateRole(updateRoleOptions)
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
		It(`GetRole request example`, func() {
			// begin-get_role

			getRoleOptions := iamPolicyManagementService.NewGetRoleOptions(
				"testString",
			)

			customRole, response, err := iamPolicyManagementService.GetRole(getRoleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customRole, "", "  ")
			fmt.Println(string(b))

			// end-get_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customRole).ToNot(BeNil())

		})
		It(`DeleteRole request example`, func() {
			// begin-delete_role

			deleteRoleOptions := iamPolicyManagementService.NewDeleteRoleOptions(
				"testString",
			)

			response, err := iamPolicyManagementService.DeleteRole(deleteRoleOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeletePolicy request example`, func() {
			// begin-delete_policy

			deletePolicyOptions := iamPolicyManagementService.NewDeletePolicyOptions(
				"testString",
			)

			response, err := iamPolicyManagementService.DeletePolicy(deletePolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
