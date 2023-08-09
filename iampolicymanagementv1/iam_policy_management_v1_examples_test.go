//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2023.
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
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IAM Policy Management service.
//
// The following configuration properties are assumed to be defined:
//
// IAM_POLICY_MANAGEMENT_URL=<service url>
// IAM_POLICY_MANAGEMENT_AUTH_TYPE=iam
// IAM_POLICY_MANAGEMENT_AUTH_URL=<IAM token service URL - omit this if using the production environment>
// IAM_POLICY_MANAGEMENT_APIKEY=<YOUR_APIKEY>
// IAM_POLICY_MANAGEMENT_TEST_ACCOUNT_ID=<YOUR_ACCOUNT_ID>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of config file>
//
// Location of our config file.

var _ = Describe(`IamPolicyManagementV1 Examples Tests`, func() {
	const externalConfigFile = "../iam_policy_management.env"

	var (
		// TODO: Align
		iamPolicyManagementService *iampolicymanagementv1.IamPolicyManagementV1
		config                     map[string]string
		configLoaded               bool = false

		exampleUserID                = "IBMid-user1"
		exampleServiceName           = "iam-groups"
		exampleAccountID             string
		examplePolicyID              string
		examplePolicyETag            string
		exampleCustomRoleID          string
		exampleCustomRoleETag        string
		examplePolicyTemplateName    = "PolicySampleTemplateTest"
		examplePolicyTemplateID      string
		examplePolicyTemplateETag    string
		examplePolicyTemplateVersion string
		testPolicyAssignmentId       string
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping tests...")
		}
	}

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

			exampleAccountID = config["TEST_ACCOUNT_ID"]

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
			fmt.Println("\nCreatePolicy() result:")
			// begin-create_policy

			subjectAttribute := &iampolicymanagementv1.SubjectAttribute{
				Name:  core.StringPtr("iam_id"),
				Value: &exampleUserID,
			}
			policySubjects := &iampolicymanagementv1.PolicySubject{
				Attributes: []iampolicymanagementv1.SubjectAttribute{*subjectAttribute},
			}
			policyRoles := &iampolicymanagementv1.PolicyRole{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Viewer"),
			}
			accountIDResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name:     core.StringPtr("accountId"),
				Value:    core.StringPtr(exampleAccountID),
				Operator: core.StringPtr("stringEquals"),
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name:     core.StringPtr("serviceType"),
				Value:    core.StringPtr("service"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResourceTag := &iampolicymanagementv1.ResourceTag{
				Name:     core.StringPtr("project"),
				Value:    core.StringPtr("prototype"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResources := &iampolicymanagementv1.PolicyResource{
				Attributes: []iampolicymanagementv1.ResourceAttribute{
					*accountIDResourceAttribute, *serviceNameResourceAttribute},
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

			examplePolicyID = *policy.ID
		})
		It(`GetPolicy request example`, func() {
			fmt.Println("\nGetPolicy() result:")
			// begin-get_policy

			options := iamPolicyManagementService.NewGetPolicyOptions(
				examplePolicyID,
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

			examplePolicyETag = response.GetHeaders().Get("ETag")
		})
		It(`ReplacePolicy request example`, func() {
			fmt.Println("\nReplacePolicy() result:")
			// begin-replace_policy

			subjectAttribute := &iampolicymanagementv1.SubjectAttribute{
				Name:  core.StringPtr("iam_id"),
				Value: &exampleUserID,
			}
			policySubjects := &iampolicymanagementv1.PolicySubject{
				Attributes: []iampolicymanagementv1.SubjectAttribute{*subjectAttribute},
			}
			accountIDResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name:     core.StringPtr("accountId"),
				Value:    core.StringPtr(exampleAccountID),
				Operator: core.StringPtr("stringEquals"),
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name:     core.StringPtr("serviceType"),
				Value:    core.StringPtr("service"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResourceTag := &iampolicymanagementv1.ResourceTag{
				Name:     core.StringPtr("project"),
				Value:    core.StringPtr("prototype"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResources := &iampolicymanagementv1.PolicyResource{
				Attributes: []iampolicymanagementv1.ResourceAttribute{
					*accountIDResourceAttribute, *serviceNameResourceAttribute},
				Tags: []iampolicymanagementv1.ResourceTag{*policyResourceTag},
			}
			updatedPolicyRoles := &iampolicymanagementv1.PolicyRole{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Editor"),
			}

			options := iamPolicyManagementService.NewReplacePolicyOptions(
				examplePolicyID,
				examplePolicyETag,
				"access",
				[]iampolicymanagementv1.PolicySubject{*policySubjects},
				[]iampolicymanagementv1.PolicyRole{*updatedPolicyRoles},
				[]iampolicymanagementv1.PolicyResource{*policyResources},
			)

			policy, response, err := iamPolicyManagementService.ReplacePolicy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-replace_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())

			examplePolicyETag = response.GetHeaders().Get("ETag")
		})
		It(`UpdatePolicy request example`, func() {
			fmt.Println("\nUpdatePolicyState() result:")
			// begin-update_policy_state

			options := iamPolicyManagementService.NewUpdatePolicyStateOptions(
				examplePolicyID,
				examplePolicyETag,
			)

			options.SetState("active")

			policy, response, err := iamPolicyManagementService.UpdatePolicyState(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-update_policy_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())

		})
		It(`ListPolicies request example`, func() {
			fmt.Println("\nListPolicies() result:")
			// begin-list_policies

			options := iamPolicyManagementService.NewListPoliciesOptions(
				exampleAccountID,
			)
			options.SetIamID(exampleUserID)
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
				examplePolicyID,
			)

			response, err := iamPolicyManagementService.DeletePolicy(options)
			if err != nil {
				panic(err)
			}

			// end-delete_policy
			fmt.Printf("\nDeletePolicy() response status code: %d\n", response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`CreateV2Policy request example`, func() {
			fmt.Println("\nCreateV2Policy() result:")
			// begin-create_v2_policy

			subjectAttribute := &iampolicymanagementv1.V2PolicySubjectAttribute{
				Key:      core.StringPtr("iam_id"),
				Operator: core.StringPtr("stringEquals"),
				Value:    &exampleUserID,
			}
			policySubject := &iampolicymanagementv1.V2PolicySubject{
				Attributes: []iampolicymanagementv1.V2PolicySubjectAttribute{*subjectAttribute},
			}
			policyRole := &iampolicymanagementv1.Roles{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Viewer"),
			}
			v2PolicyGrant := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*policyRole},
			}
			v2PolicyControl := &iampolicymanagementv1.Control{
				Grant: v2PolicyGrant,
			}
			accountIDResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("accountId"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr(exampleAccountID),
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceType"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("service"),
			}
			policyResourceTag := &iampolicymanagementv1.V2PolicyResourceTag{
				Key:      core.StringPtr("project"),
				Value:    core.StringPtr("prototype"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResource := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{
					*accountIDResourceAttribute, *serviceNameResourceAttribute},
				Tags: []iampolicymanagementv1.V2PolicyResourceTag{*policyResourceTag},
			}
			weeklyConditionAttribute := &iampolicymanagementv1.RuleAttribute{
				Key:      core.StringPtr("{{environment.attributes.day_of_week}}"),
				Operator: core.StringPtr("dayOfWeekAnyOf"),
				Value:    []string{"1+00:00", "2+00:00", "3+00:00", "4+00:00", "5+00:00"},
			}
			startConditionAttribute := &iampolicymanagementv1.RuleAttribute{
				Key:      core.StringPtr("{{environment.attributes.current_time}}"),
				Operator: core.StringPtr("timeGreaterThanOrEquals"),
				Value:    core.StringPtr("09:00:00+00:00"),
			}
			endConditionAttribute := &iampolicymanagementv1.RuleAttribute{
				Key:      core.StringPtr("{{environment.attributes.current_time}}"),
				Operator: core.StringPtr("timeLessThanOrEquals"),
				Value:    core.StringPtr("17:00:00+00:00"),
			}
			policyRule := &iampolicymanagementv1.V2PolicyRule{
				Operator: core.StringPtr("and"),
				Conditions: []iampolicymanagementv1.RuleAttribute{
					*weeklyConditionAttribute, *startConditionAttribute, *endConditionAttribute},
			}

			options := iamPolicyManagementService.NewCreateV2PolicyOptions(
				v2PolicyControl,
				"access",
			)
			options.SetSubject(policySubject)
			options.SetResource(policyResource)
			options.SetRule(policyRule)
			options.SetPattern(*core.StringPtr("time-based-conditions:weekly:custom-hours"))

			policy, response, err := iamPolicyManagementService.CreateV2Policy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-create_v2_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())

			examplePolicyID = *policy.ID
		})
		It(`GetV2Policy request example`, func() {
			fmt.Println("\nGetV2Policy() result:")
			// begin-get_v2_policy

			options := iamPolicyManagementService.NewGetV2PolicyOptions(
				examplePolicyID,
			)

			policy, response, err := iamPolicyManagementService.GetV2Policy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-get_v2_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())

			examplePolicyETag = response.GetHeaders().Get("ETag")
		})
		It(`ReplaceV2Policy request example`, func() {
			fmt.Println("\nReplaceV2Policy() result:")
			// begin-replace_v2_policy

			subjectAttribute := &iampolicymanagementv1.V2PolicySubjectAttribute{
				Key:      core.StringPtr("iam_id"),
				Operator: core.StringPtr("stringEquals"),
				Value:    &exampleUserID,
			}
			policySubject := &iampolicymanagementv1.V2PolicySubject{
				Attributes: []iampolicymanagementv1.V2PolicySubjectAttribute{*subjectAttribute},
			}
			updatedPolicyRole := &iampolicymanagementv1.Roles{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Editor"),
			}
			v2PolicyGrant := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*updatedPolicyRole},
			}
			v2PolicyControl := &iampolicymanagementv1.Control{
				Grant: v2PolicyGrant,
			}
			accountIDResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("accountId"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr(exampleAccountID),
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceType"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("service"),
			}
			policyResource := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{
					*accountIDResourceAttribute, *serviceNameResourceAttribute},
			}

			options := iamPolicyManagementService.NewReplaceV2PolicyOptions(
				examplePolicyID,
				examplePolicyETag,
				v2PolicyControl,
				"access",
			)
			weeklyConditionAttribute := &iampolicymanagementv1.RuleAttribute{
				Key:      core.StringPtr("{{environment.attributes.day_of_week}}"),
				Operator: core.StringPtr("dayOfWeekAnyOf"),
				Value:    []string{"1+00:00", "2+00:00", "3+00:00", "4+00:00"},
			}
			startConditionAttribute := &iampolicymanagementv1.RuleAttribute{
				Key:      core.StringPtr("{{environment.attributes.current_time}}"),
				Operator: core.StringPtr("timeGreaterThanOrEquals"),
				Value:    core.StringPtr("09:00:00+00:00"),
			}
			endConditionAttribute := &iampolicymanagementv1.RuleAttribute{
				Key:      core.StringPtr("{{environment.attributes.current_time}}"),
				Operator: core.StringPtr("timeLessThanOrEquals"),
				Value:    core.StringPtr("17:00:00+00:00"),
			}
			policyRule := &iampolicymanagementv1.V2PolicyRule{
				Operator: core.StringPtr("and"),
				Conditions: []iampolicymanagementv1.RuleAttribute{
					*weeklyConditionAttribute, *startConditionAttribute, *endConditionAttribute},
			}
			options.SetRule(policyRule)
			options.SetPattern(*core.StringPtr("time-based-conditions:weekly:custom-hours"))
			options.SetSubject(policySubject)
			options.SetResource(policyResource)

			policy, response, err := iamPolicyManagementService.ReplaceV2Policy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-replace_v2_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())

			examplePolicyETag = response.GetHeaders().Get("ETag")
		})
		It(`ListV2Policies request example`, func() {
			fmt.Println("\nListV2Policies() result:")
			// begin-list_v2_policies

			options := iamPolicyManagementService.NewListV2PoliciesOptions(
				exampleAccountID,
			)
			options.SetIamID(exampleUserID)
			options.SetFormat("include_last_permit")
			options.SetSort("-id")

			policyList, response, err := iamPolicyManagementService.ListV2Policies(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyList, "", "  ")
			fmt.Println(string(b))

			// end-list_v2_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyList).ToNot(BeNil())

		})
		It(`DeleteV2Policy request example`, func() {
			// begin-delete_v2_policy

			options := iamPolicyManagementService.NewDeleteV2PolicyOptions(
				examplePolicyID,
			)

			response, err := iamPolicyManagementService.DeleteV2Policy(options)
			if err != nil {
				panic(err)
			}

			// end-delete_delete_v2_policypolicy
			fmt.Printf("\nDeleteV2Policy() response status code: %d\n", response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`CreateRole request example`, func() {
			fmt.Println("\nCreateRole() result:")
			// begin-create_role

			options := iamPolicyManagementService.NewCreateRoleOptions(
				"IAM Groups read access",
				[]string{"iam-groups.groups.read"},
				"ExampleRoleIAMGroups",
				exampleAccountID,
				exampleServiceName,
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

			exampleCustomRoleID = *customRole.ID
		})
		It(`GetRole request example`, func() {
			fmt.Println("\nGetRole() result:")
			// begin-get_role

			options := iamPolicyManagementService.NewGetRoleOptions(
				exampleCustomRoleID,
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

			exampleCustomRoleETag = response.Headers.Get("ETag")

		})
		It(`ReplaceRole request example`, func() {
			fmt.Println("\nReplaceRole() result:")
			// begin-replace_role

			updatedRoleActions := []string{"iam-groups.groups.read", "iam-groups.groups.list"}

			options := iamPolicyManagementService.NewReplaceRoleOptions(
				exampleCustomRoleID,
				exampleCustomRoleETag,
				"ExampleRoleIAMGroups",
				updatedRoleActions,
			)

			customRole, response, err := iamPolicyManagementService.ReplaceRole(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customRole, "", "  ")
			fmt.Println(string(b))

			// end-replace_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customRole).ToNot(BeNil())

		})
		It(`ListRoles request example`, func() {
			fmt.Println("\nListRoles() result:")
			// begin-list_roles

			options := iamPolicyManagementService.NewListRolesOptions()
			options.SetAccountID(exampleAccountID)

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
				exampleCustomRoleID,
			)

			response, err := iamPolicyManagementService.DeleteRole(options)
			if err != nil {
				panic(err)
			}

			// end-delete_role
			fmt.Printf("\nDeleteRole() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListPolicyTemplates request example`, func() {
			fmt.Println("\nListPolicyTemplates() result:")
			// begin-list_policy_templates

			listPolicyTemplatesOptions := iamPolicyManagementService.NewListPolicyTemplatesOptions(
				exampleAccountID,
			)

			policyTemplateCollection, response, err := iamPolicyManagementService.ListPolicyTemplates(listPolicyTemplatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplateCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_policy_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplateCollection).ToNot(BeNil())
		})
		It(`CreatePolicyTemplate request example`, func() {
			fmt.Println("\nCreatePolicyTemplate() result:")
			// begin-create_policy_template

			policyRole := &iampolicymanagementv1.Roles{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Viewer"),
			}
			v2PolicyGrant := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*policyRole},
			}
			v2PolicyControl := &iampolicymanagementv1.Control{
				Grant: v2PolicyGrant,
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("iam-access-management"),
			}
			policyResource := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{
					*serviceNameResourceAttribute},
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Test Template"),
				Resource:    policyResource,
				Control:     v2PolicyControl,
			}

			createPolicyTemplateOptions := iamPolicyManagementService.NewCreatePolicyTemplateOptions(
				examplePolicyTemplateName,
				exampleAccountID,
				templatePolicyModel,
			)

			policyTemplate, response, err := iamPolicyManagementService.CreatePolicyTemplate(createPolicyTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplate, "", "  ")
			fmt.Println(string(b))

			// end-create_policy_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())

			examplePolicyTemplateID = *policyTemplate.ID

		})
		It(`GetPolicyTemplate request example`, func() {
			fmt.Println("\nGetPolicyTemplate() result:")
			// begin-get_policy_template

			getPolicyTemplateOptions := iamPolicyManagementService.NewGetPolicyTemplateOptions(
				examplePolicyTemplateID,
			)

			policyTemplate, response, err := iamPolicyManagementService.GetPolicyTemplate(getPolicyTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplate, "", "  ")
			fmt.Println(string(b))

			// end-get_policy_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate.AccountID).ToNot(BeNil())
			Expect(policyTemplate.Version).ToNot(BeNil())
			Expect(policyTemplate.Name).ToNot(BeNil())
			Expect(policyTemplate.Policy).ToNot(BeNil())

			examplePolicyTemplateETag = response.GetHeaders().Get("ETag")

		})

		It(`CreatePolicyTemplateVersion request example`, func() {
			fmt.Println("\nCreatePolicyTemplateVersion() result:")
			// begin-create_policy_template_version

			v2PolicyGrant := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{
					{core.StringPtr("crn:v1:bluemix:public:iam::::role:Viewer")},
					{core.StringPtr("crn:v1:bluemix:public:iam::::role:Administrator")},
				},
			}

			v2PolicyControl := &iampolicymanagementv1.Control{
				Grant: v2PolicyGrant,
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("watson"),
			}
			policyResource := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{
					*serviceNameResourceAttribute},
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Test Template v2"),
				Resource:    policyResource,
				Control:     v2PolicyControl,
			}

			createPolicyTemplateVersionOptions := iamPolicyManagementService.NewCreatePolicyTemplateVersionOptions(
				examplePolicyTemplateID,
				templatePolicyModel,
			)

			policyTemplate, response, err := iamPolicyManagementService.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplate, "", "  ")
			fmt.Println(string(b))

			// end-create_policy_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())

			examplePolicyTemplateVersion = *policyTemplate.Version
			examplePolicyTemplateETag = response.GetHeaders().Get("ETag")

		})

		It(`ListPolicyTemplateVersions request example`, func() {
			fmt.Println("\nListPolicyTemplateVersions() result:")
			// begin-list_policy_template_versions

			listPolicyTemplateVersionsOptions := iamPolicyManagementService.NewListPolicyTemplateVersionsOptions(
				examplePolicyTemplateID,
			)

			policyTemplateVersionsCollection, response, err := iamPolicyManagementService.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplateVersionsCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_policy_template_versions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplateVersionsCollection).ToNot(BeNil())
		})

		It(`ReplacePolicyTemplate request example`, func() {
			fmt.Println("\nReplacePolicyTemplate() result:")
			// begin-replace_policy_template
			v2PolicyGrant := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{
					{core.StringPtr("crn:v1:bluemix:public:iam::::role:Viewer")},
					{core.StringPtr("crn:v1:bluemix:public:iam::::role:Administrator")},
				},
			}

			v2PolicyControl := &iampolicymanagementv1.Control{
				Grant: v2PolicyGrant,
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("watson"),
			}
			policyResource := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{
					*serviceNameResourceAttribute},
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Test Template v2"),
				Resource:    policyResource,
				Control:     v2PolicyControl,
			}

			replacePolicyTemplateOptions := iamPolicyManagementService.NewReplacePolicyTemplateOptions(
				examplePolicyTemplateID,
				examplePolicyTemplateVersion,
				examplePolicyTemplateETag,
				templatePolicyModel,
			)

			policyTemplate, response, err := iamPolicyManagementService.ReplacePolicyTemplate(replacePolicyTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplate, "", "  ")
			fmt.Println(string(b))

			// end-replace_policy_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())

			examplePolicyTemplateVersion = *policyTemplate.Version
			examplePolicyTemplateETag = response.GetHeaders().Get("ETag")
		})
		It(`GetPolicyTemplateVersion request example`, func() {
			fmt.Println("\nGetPolicyTemplateVersion() result:")
			// begin-get_policy_template_version

			getPolicyTemplateVersionOptions := iamPolicyManagementService.NewGetPolicyTemplateVersionOptions(
				examplePolicyTemplateID,
				examplePolicyTemplateVersion,
			)

			policyTemplate, response, err := iamPolicyManagementService.GetPolicyTemplateVersion(getPolicyTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplate, "", "  ")
			fmt.Println(string(b))

			// end-get_policy_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())
		})

		It(`CommitPolicyTemplate request example`, func() {
			fmt.Println("\nCommitPolicyTemplate() result:")
			// begin-commit_policy_template

			commitPolicyTemplateOptions := iamPolicyManagementService.NewCommitPolicyTemplateOptions(
				examplePolicyTemplateID,
				examplePolicyTemplateVersion,
				examplePolicyTemplateETag,
			)

			response, err := iamPolicyManagementService.CommitPolicyTemplate(commitPolicyTemplateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from CommitPolicyTemplate(): %d\n", response.StatusCode)
			}

			// end-commit_policy_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`ListPolicyAssignments request example`, func() {
			fmt.Println("\nListPolicyAssignments() result:")
			// begin-list_Policy Assignments

			listPolicyAssignmentsOptions := iamPolicyManagementService.NewListPolicyAssignmentsOptions(
				exampleAccountID,
			)

			polcyTemplateAssignmentCollection, response, err := iamPolicyManagementService.ListPolicyAssignments(listPolicyAssignmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(polcyTemplateAssignmentCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_Policy Assignments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(polcyTemplateAssignmentCollection).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].TemplateID).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].TargetType).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].TemplateVersion).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].Target).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].AssignmentID).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].Options).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].Status).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].AccountID).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].Resources).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].CreatedAt).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].CreatedByID).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].LastModifiedAt).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].LastModifiedByID).ToNot(BeNil())
			Expect(polcyTemplateAssignmentCollection.Assignments[0].Href).ToNot(BeNil())

			testPolicyAssignmentId = *polcyTemplateAssignmentCollection.Assignments[0].ID
		})

		It(`GetPolicyAssignment request example`, func() {
			fmt.Println("\nGetPolicyAssignment() result:")
			// begin-get_policy_assignment

			getPolicyAssignmentOptions := iamPolicyManagementService.NewGetPolicyAssignmentOptions(
				testPolicyAssignmentId,
			)

			policyAssignmentRecord, response, err := iamPolicyManagementService.GetPolicyAssignment(getPolicyAssignmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyAssignmentRecord, "", "  ")
			fmt.Println(string(b))

			// end-get_policy_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyAssignmentRecord).ToNot(BeNil())
			Expect(policyAssignmentRecord.TemplateID).ToNot(BeNil())
			Expect(policyAssignmentRecord.TargetType).ToNot(BeNil())
			Expect(policyAssignmentRecord.TemplateVersion).ToNot(BeNil())
			Expect(policyAssignmentRecord.Target).ToNot(BeNil())
			Expect(policyAssignmentRecord.AssignmentID).ToNot(BeNil())
			Expect(policyAssignmentRecord.Options).ToNot(BeNil())
			Expect(policyAssignmentRecord.Status).ToNot(BeNil())
			Expect(policyAssignmentRecord.AccountID).ToNot(BeNil())
			Expect(policyAssignmentRecord.Resources).ToNot(BeNil())
			Expect(policyAssignmentRecord.CreatedAt).ToNot(BeNil())
			Expect(policyAssignmentRecord.CreatedByID).ToNot(BeNil())
			Expect(policyAssignmentRecord.LastModifiedAt).ToNot(BeNil())
			Expect(policyAssignmentRecord.LastModifiedByID).ToNot(BeNil())
			Expect(policyAssignmentRecord.Href).ToNot(BeNil())
		})

		It(`DeletePolicyTemplateVersion request example`, func() {
			// begin-delete_policy_template_version

			deletePolicyTemplateVersionOptions := iamPolicyManagementService.NewDeletePolicyTemplateVersionOptions(
				examplePolicyTemplateID,
				examplePolicyTemplateVersion,
			)

			response, err := iamPolicyManagementService.DeletePolicyTemplateVersion(deletePolicyTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeletePolicyTemplateVersion(): %d\n", response.StatusCode)
			}

			// end-delete_policy_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`DeletePolicyTemplate request example`, func() {
			// begin-delete_policy_template

			deletePolicyTemplateOptions := iamPolicyManagementService.NewDeletePolicyTemplateOptions(
				examplePolicyTemplateID,
			)

			response, err := iamPolicyManagementService.DeletePolicyTemplate(deletePolicyTemplateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeletePolicyTemplate(): %d\n", response.StatusCode)
			}

			// end-delete_policy_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

	})
})
