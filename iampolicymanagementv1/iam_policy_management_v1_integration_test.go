//go:build integration
// +build integration

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

package iampolicymanagementv1_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

var _ = Describe("IAM Policy Management - Integration Tests", func() {
	const externalConfigFile = "../iam_policy_management.env"

	var (
		service *iampolicymanagementv1.IamPolicyManagementV1

		err          error
		config       map[string]string
		configLoaded bool = false

		testAccountID     string
		etagHeader        string = "ETag"
		testPolicyETag    string = ""
		testV2PolicyETag  string = ""
		testPolicyId      string = ""
		testV2PolicyId    string = ""
		testUserId        string = "IBMid-GoSDK" + strconv.Itoa(rand.Intn(100000))
		testViewerRoleCrn string = "crn:v1:bluemix:public:iam::::role:Viewer"
		testEditorRoleCrn string = "crn:v1:bluemix:public:iam::::role:Editor"
		testServiceName   string = "iam-groups"

		testCustomRoleId          string = ""
		testCustomRoleETag        string = ""
		testCustomRoleName        string = "TestGoRole" + strconv.Itoa(rand.Intn(100000))
		testServiceRoleCrn        string = "crn:v1:bluemix:public:iam-identity::::serviceRole:ServiceIdCreator"
		testPolicyTemplateID      string = ""
		testPolicyTemplateETag    string = ""
		testPolicyTemplateVersion string = ""
		testPolicyAssignmentId    string = ""
		examplePolicyTemplateName        = "PolicySampleTemplateTestV1"
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping...")
		}
	}

	It("Successfully load the configuration", func() {
		err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
		if err != nil {
			Skip("Could not set IBM_CREDENTIALS_FILE environment variable: " + err.Error())
		}

		config, err = core.GetServiceProperties(iampolicymanagementv1.DefaultServiceName)
		if err == nil {
			testAccountID = config["TEST_ACCOUNT_ID"]
			if testAccountID != "" {
				configLoaded = true
			}
		}
		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}
	})

	It(`Successfully created IamPolicyManagementV1 service instance`, func() {
		shouldSkipTest()

		service, err = iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(
			&iampolicymanagementv1.IamPolicyManagementV1Options{},
		)

		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())

		core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
		service.EnableRetries(4, 30*time.Second)
	})

	Describe("Create an access policy", func() {

		It("Successfully created an access policy", func() {
			shouldSkipTest()

			// Construct an instance of the ResourceAttribute model
			accountIdResourceAttribute := new(iampolicymanagementv1.ResourceAttribute)
			accountIdResourceAttribute.Name = core.StringPtr("accountId")
			accountIdResourceAttribute.Value = core.StringPtr(testAccountID)
			accountIdResourceAttribute.Operator = core.StringPtr("stringEquals")

			serviceNameResourceAttribute := new(iampolicymanagementv1.ResourceAttribute)
			serviceNameResourceAttribute.Name = core.StringPtr("serviceType")
			serviceNameResourceAttribute.Value = core.StringPtr("service")
			serviceNameResourceAttribute.Operator = core.StringPtr("stringEquals")

			policyResourceTag := new(iampolicymanagementv1.ResourceTag)
			policyResourceTag.Name = core.StringPtr("project")
			policyResourceTag.Value = core.StringPtr("prototype")
			policyResourceTag.Operator = core.StringPtr("stringEquals")

			// Construct an instance of the SubjectAttribute model
			subjectAttribute := new(iampolicymanagementv1.SubjectAttribute)
			subjectAttribute.Name = core.StringPtr("iam_id")
			subjectAttribute.Value = core.StringPtr(testUserId)

			// Construct an instance of the PolicyResource model
			policyResource := new(iampolicymanagementv1.PolicyResource)
			policyResource.Attributes = []iampolicymanagementv1.ResourceAttribute{*accountIdResourceAttribute, *serviceNameResourceAttribute}
			policyResource.Tags = []iampolicymanagementv1.ResourceTag{*policyResourceTag}

			// Construct an instance of the PolicyRole model
			policyRole := new(iampolicymanagementv1.PolicyRole)
			policyRole.RoleID = core.StringPtr(testViewerRoleCrn)

			// Construct an instance of the PolicySubject model
			policySubject := new(iampolicymanagementv1.PolicySubject)
			policySubject.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttribute}

			// Construct an instance of the CreatePolicyOptions model
			options := new(iampolicymanagementv1.CreatePolicyOptions)
			options.Type = core.StringPtr("access")
			options.Subjects = []iampolicymanagementv1.PolicySubject{*policySubject}
			options.Roles = []iampolicymanagementv1.PolicyRole{*policyRole}
			options.Resources = []iampolicymanagementv1.PolicyResource{*policyResource}
			options.AcceptLanguage = core.StringPtr("en")

			policy, detailedResponse, err := service.CreatePolicy(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreatePolicy() result:\n%s\n", common.ToJSON(policy))
			Expect(policy.Type).To(Equal(options.Type))
			Expect(policy.Subjects).To(Equal(options.Subjects))
			Expect(policy.Roles[0].RoleID).To(Equal(options.Roles[0].RoleID))
			Expect(policy.Resources).To(Equal(options.Resources))

			testPolicyId = *policy.ID
		})
	})

	Describe("Get an access policy", func() {

		It("Successfully retrieved an access policy", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))

			options := service.NewGetPolicyOptions(testPolicyId)
			policy, detailedResponse, err := service.GetPolicy(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetPolicy() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(testPolicyId))

			testPolicyETag = detailedResponse.GetHeaders().Get(etagHeader)
		})
	})

	Describe("Update an access policy", func() {

		It("Successfully updated an access policy", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))
			Expect(testPolicyETag).To(Not(BeNil()))

			// Construct an instance of the ResourceAttribute model
			accountIdResourceAttribute := new(iampolicymanagementv1.ResourceAttribute)
			accountIdResourceAttribute.Name = core.StringPtr("accountId")
			accountIdResourceAttribute.Value = core.StringPtr(testAccountID)
			accountIdResourceAttribute.Operator = core.StringPtr("stringEquals")

			serviceNameResourceAttribute := new(iampolicymanagementv1.ResourceAttribute)
			serviceNameResourceAttribute.Name = core.StringPtr("serviceType")
			serviceNameResourceAttribute.Value = core.StringPtr("service")
			serviceNameResourceAttribute.Operator = core.StringPtr("stringEquals")

			// Construct an instance of the SubjectAttribute model
			subjectAttribute := new(iampolicymanagementv1.SubjectAttribute)
			subjectAttribute.Name = core.StringPtr("iam_id")
			subjectAttribute.Value = core.StringPtr(testUserId)

			policyResourceTag := new(iampolicymanagementv1.ResourceTag)
			policyResourceTag.Name = core.StringPtr("project")
			policyResourceTag.Value = core.StringPtr("prototype")
			policyResourceTag.Operator = core.StringPtr("stringEquals")

			// Construct an instance of the PolicyResource model
			policyResource := new(iampolicymanagementv1.PolicyResource)
			policyResource.Attributes = []iampolicymanagementv1.ResourceAttribute{*accountIdResourceAttribute, *serviceNameResourceAttribute}
			policyResource.Tags = []iampolicymanagementv1.ResourceTag{*policyResourceTag}

			// Construct an instance of the PolicyRole model
			policyRole := new(iampolicymanagementv1.PolicyRole)
			policyRole.RoleID = core.StringPtr(testEditorRoleCrn)

			// Construct an instance of the PolicySubject model
			policySubject := new(iampolicymanagementv1.PolicySubject)
			policySubject.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttribute}

			// Construct an instance of the CreatePolicyOptions model
			options := new(iampolicymanagementv1.ReplacePolicyOptions)
			options.PolicyID = core.StringPtr(testPolicyId)
			options.IfMatch = core.StringPtr(testPolicyETag)
			options.Type = core.StringPtr("access")
			options.Subjects = []iampolicymanagementv1.PolicySubject{*policySubject}
			options.Roles = []iampolicymanagementv1.PolicyRole{*policyRole}
			options.Resources = []iampolicymanagementv1.PolicyResource{*policyResource}

			policy, detailedResponse, err := service.ReplacePolicy(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ReplacePolicy() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(testPolicyId))
			Expect(policy.Type).To(Equal(options.Type))
			Expect(policy.Subjects).To(Equal(options.Subjects))
			Expect(policy.Roles[0].RoleID).To(Equal(options.Roles[0].RoleID))
			Expect(policy.Resources).To(Equal(options.Resources))

			testPolicyETag = detailedResponse.GetHeaders().Get(etagHeader)

		})
	})

	Describe("Patch an access policy", func() {

		It("Successfully patched an access policy", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))
			Expect(testPolicyETag).To(Not(BeNil()))

			// Construct an instance of the UpdatePolicyStateOptions model
			options := new(iampolicymanagementv1.UpdatePolicyStateOptions)
			options.PolicyID = &testPolicyId
			options.IfMatch = core.StringPtr(testPolicyETag)
			options.State = core.StringPtr("active")

			policy, detailedResponse, err := service.UpdatePolicyState(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "UpdatePolicyState() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(testPolicyId))
			Expect(policy.State).To(Equal(options.State))

		})
	})

	Describe("List access policies", func() {

		It("Successfully listed the account's access policies", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))

			options := service.NewListPoliciesOptions(testAccountID)
			options.SetIamID(testUserId)
			result, detailedResponse, err := service.ListPolicies(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListPolicies() result:\n%s\n", common.ToJSON(result))

			// confirm the test policy is present
			testPolicyPresent := false
			for _, policy := range result.Policies {
				if *policy.ID == testPolicyId {
					testPolicyPresent = true
				}
			}
			Expect(testPolicyPresent).To(BeTrue())
		})
	})

	Describe("Create a v2 access policy", func() {

		It("Successfully created a v2 access policy", func() {
			shouldSkipTest()

			// Construct an instance of the ResourceAttribute model
			accountIdResourceAttribute := new(iampolicymanagementv1.V2PolicyResourceAttribute)
			accountIdResourceAttribute.Key = core.StringPtr("accountId")
			accountIdResourceAttribute.Value = core.StringPtr(testAccountID)
			accountIdResourceAttribute.Operator = core.StringPtr("stringEquals")

			serviceNameResourceAttribute := new(iampolicymanagementv1.V2PolicyResourceAttribute)
			serviceNameResourceAttribute.Key = core.StringPtr("serviceType")
			serviceNameResourceAttribute.Value = core.StringPtr("service")
			serviceNameResourceAttribute.Operator = core.StringPtr("stringEquals")

			policyResourceTag := new(iampolicymanagementv1.V2PolicyResourceTag)
			policyResourceTag.Key = core.StringPtr("project")
			policyResourceTag.Value = core.StringPtr("prototype")
			policyResourceTag.Operator = core.StringPtr("stringEquals")

			// Construct an instance of the SubjectAttribute model
			subjectAttribute := new(iampolicymanagementv1.V2PolicySubjectAttribute)
			subjectAttribute.Key = core.StringPtr("iam_id")
			subjectAttribute.Operator = core.StringPtr("stringEquals")
			subjectAttribute.Value = core.StringPtr(testUserId)

			// Construct an instance of the V2PolicyResource model
			policyResource := new(iampolicymanagementv1.V2PolicyResource)
			policyResource.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*accountIdResourceAttribute, *serviceNameResourceAttribute}
			policyResource.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*policyResourceTag}

			// Construct an instance of the Roles model
			policyRole := new(iampolicymanagementv1.Roles)
			policyRole.RoleID = core.StringPtr(testViewerRoleCrn)

			// Construct an instance of the PolicySubject model
			policySubject := new(iampolicymanagementv1.V2PolicySubject)
			policySubject.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*subjectAttribute}

			// Contruct and instance of PolicyControl model
			control := new(iampolicymanagementv1.Control)
			grant := new(iampolicymanagementv1.Grant)
			grant.Roles = []iampolicymanagementv1.Roles{*policyRole}
			control.Grant = grant

			// Construct an instance of Policy Rule Attribute
			weeklyConditionAttribute := new(iampolicymanagementv1.RuleAttribute)
			weeklyConditionAttribute.Key = core.StringPtr("{{environment.attributes.day_of_week}}")
			weeklyConditionAttribute.Operator = core.StringPtr("dayOfWeekAnyOf")
			weeklyConditionAttribute.Value = []string{"1+00:00", "2+00:00", "3+00:00", "4+00:00", "5+00:00"}

			startConditionAttribute := new(iampolicymanagementv1.RuleAttribute)
			startConditionAttribute.Key = core.StringPtr("{{environment.attributes.current_time}}")
			startConditionAttribute.Operator = core.StringPtr("timeGreaterThanOrEquals")
			startConditionAttribute.Value = core.StringPtr("09:00:00+00:00")

			endConditionAttribute := new(iampolicymanagementv1.RuleAttribute)
			endConditionAttribute.Key = core.StringPtr("{{environment.attributes.current_time}}")
			endConditionAttribute.Operator = core.StringPtr("timeLessThanOrEquals")
			endConditionAttribute.Value = core.StringPtr("17:00:00+00:00")

			policyRule := new(iampolicymanagementv1.V2PolicyRule)
			policyRule.Operator = core.StringPtr("and")
			policyRule.Conditions = []iampolicymanagementv1.RuleAttribute{*weeklyConditionAttribute, *startConditionAttribute, *endConditionAttribute}

			// Construct an instance of the CreateV2PolicyOptions model
			options := new(iampolicymanagementv1.CreateV2PolicyOptions)
			options.Type = core.StringPtr("access")
			options.Subject = policySubject
			options.Control = control
			options.Resource = policyResource
			options.Pattern = core.StringPtr("time-based-conditions:weekly:custom-hours")
			options.Rule = policyRule
			options.AcceptLanguage = core.StringPtr("en")

			policy, detailedResponse, err := service.CreateV2Policy(options)
			controlResponse := new(iampolicymanagementv1.ControlResponse)
			controlResponse.Grant = grant
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateV2Policy() result:\n%s\n", common.ToJSON(policy))
			Expect(policy.Type).To(Equal(options.Type))
			Expect(policy.Subject.Attributes[0].Value).To(Equal(&testUserId))
			Expect(policy.Control).To(Equal(controlResponse))
			Expect(policy.Resource.Attributes[0].Value).To(Equal(testAccountID))

			testV2PolicyId = *policy.ID
		})
	})

	Describe("Get a v2 access policy", func() {

		It("Successfully retrieved a v2 access policy", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))

			options := service.NewGetV2PolicyOptions(testV2PolicyId)
			policy, detailedResponse, err := service.GetV2Policy(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetV2Policy() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(testV2PolicyId))

			testV2PolicyETag = detailedResponse.GetHeaders().Get(etagHeader)
		})
	})

	Describe("Update a v2 access policy", func() {

		It("Successfully updated a v2 access policy", func() {
			shouldSkipTest()
			Expect(testV2PolicyId).To(Not(BeNil()))
			Expect(testV2PolicyETag).To(Not(BeNil()))

			// Construct an instance of the ResourceAttribute model
			accountIdResourceAttribute := new(iampolicymanagementv1.V2PolicyResourceAttribute)
			accountIdResourceAttribute.Key = core.StringPtr("accountId")
			accountIdResourceAttribute.Value = core.StringPtr(testAccountID)
			accountIdResourceAttribute.Operator = core.StringPtr("stringEquals")

			serviceNameResourceAttribute := new(iampolicymanagementv1.V2PolicyResourceAttribute)
			serviceNameResourceAttribute.Key = core.StringPtr("serviceType")
			serviceNameResourceAttribute.Value = core.StringPtr("service")
			serviceNameResourceAttribute.Operator = core.StringPtr("stringEquals")

			// Construct an instance of the SubjectAttribute model
			subjectAttribute := new(iampolicymanagementv1.V2PolicySubjectAttribute)
			subjectAttribute.Key = core.StringPtr("iam_id")
			subjectAttribute.Operator = core.StringPtr("stringEquals")
			subjectAttribute.Value = core.StringPtr(testUserId)

			// Construct an instance of the V2PolicyBaseResource model
			policyResource := new(iampolicymanagementv1.V2PolicyResource)
			policyResource.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*accountIdResourceAttribute, *serviceNameResourceAttribute}

			// Construct an instance of the Roles model
			policyRole := new(iampolicymanagementv1.Roles)
			policyRole.RoleID = core.StringPtr(testViewerRoleCrn)

			// Construct an instance of the PolicySubject model
			policySubject := new(iampolicymanagementv1.V2PolicySubject)
			policySubject.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*subjectAttribute}

			// Contruct and instance of PolicyControl model
			control := new(iampolicymanagementv1.Control)
			grant := new(iampolicymanagementv1.Grant)
			grant.Roles = []iampolicymanagementv1.Roles{*policyRole}
			control.Grant = grant

			// Construct an instance of Policy Rule Attribute
			weeklyConditionAttribute := new(iampolicymanagementv1.RuleAttribute)
			weeklyConditionAttribute.Key = core.StringPtr("{{environment.attributes.day_of_week}}")
			weeklyConditionAttribute.Operator = core.StringPtr("dayOfWeekAnyOf")
			weeklyConditionAttribute.Value = []string{"1+00:00", "2+00:00", "3+00:00", "4+00:00", "5+00:00"}

			startConditionAttribute := new(iampolicymanagementv1.RuleAttribute)
			startConditionAttribute.Key = core.StringPtr("{{environment.attributes.current_time}}")
			startConditionAttribute.Operator = core.StringPtr("timeGreaterThanOrEquals")
			startConditionAttribute.Value = core.StringPtr("09:00:00+00:00")

			endConditionAttribute := new(iampolicymanagementv1.RuleAttribute)
			endConditionAttribute.Key = core.StringPtr("{{environment.attributes.current_time}}")
			endConditionAttribute.Operator = core.StringPtr("timeLessThanOrEquals")
			endConditionAttribute.Value = core.StringPtr("17:00:00+00:00")

			policyRule := new(iampolicymanagementv1.V2PolicyRule)
			policyRule.Operator = core.StringPtr("and")
			policyRule.Conditions = []iampolicymanagementv1.RuleAttribute{*weeklyConditionAttribute, *startConditionAttribute, *endConditionAttribute}

			// Construct an instance of the ReplaceV2PolicyOptions model
			options := new(iampolicymanagementv1.ReplaceV2PolicyOptions)
			options.ID = core.StringPtr(testV2PolicyId)
			options.IfMatch = core.StringPtr(testV2PolicyETag)
			options.Type = core.StringPtr("access")
			options.Subject = policySubject
			options.Control = control
			options.Resource = policyResource
			options.Pattern = core.StringPtr("time-based-conditions:weekly:custom-hours")
			options.Rule = policyRule

			policy, detailedResponse, err := service.ReplaceV2Policy(options)
			controlResponse := new(iampolicymanagementv1.ControlResponse)
			controlResponse.Grant = grant
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ReplaceV2Policy() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(testV2PolicyId))
			Expect(policy.Type).To(Equal(options.Type))
			Expect(policy.Subject.Attributes[0].Value).To(Equal(&testUserId))
			Expect(policy.Control).To(Equal(controlResponse))
			Expect(policy.Resource.Attributes[0].Value).To(Equal(testAccountID))

			newV2PolicyEtag := detailedResponse.GetHeaders().Get(etagHeader)
			Expect(newV2PolicyEtag).ToNot(Equal(testV2PolicyETag))

		})
	})

	Describe("List v2 access policies", func() {

		It("Successfully listed the account's v2 access policies", func() {
			shouldSkipTest()
			Expect(testV2PolicyId).To(Not(BeNil()))

			options := service.NewListV2PoliciesOptions(testAccountID)
			options.SetIamID(testUserId)
			options.SetSort("-id")
			result, detailedResponse, err := service.ListV2Policies(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListV2Policies() result:\n%s\n", common.ToJSON(result))

			// confirm the test policy is present
			testPolicyPresent := false
			for _, policy := range result.Policies {
				if *policy.ID == testV2PolicyId {
					testPolicyPresent = true
				}
			}
			Expect(testPolicyPresent).To(BeTrue())
		})
	})

	Describe("Create custom role", func() {
		It("Successfully created custom role", func() {
			shouldSkipTest()

			actions := []string{"iam-groups.groups.read"}
			options := service.NewCreateRoleOptions(
				testCustomRoleName,
				actions,
				testCustomRoleName,
				testAccountID,
				testServiceName)
			options.SetDescription("GO SDK test role")
			result, detailedResponse, err := service.CreateRole(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateRole() result:\n%s\n", common.ToJSON(result))

			testCustomRoleId = *result.ID
		})
	})

	Describe("Get a custom role", func() {
		It("Successfully retrieved a custom role", func() {
			shouldSkipTest()
			Expect(testCustomRoleId).To(Not(BeNil()))

			options := service.NewGetRoleOptions(testCustomRoleId)
			result, detailedResponse, err := service.GetRole(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetRole() result:\n%s\n", common.ToJSON(result))
			Expect(*result.ID).To(Equal(testCustomRoleId))

			testCustomRoleETag = detailedResponse.GetHeaders().Get(etagHeader)
		})
	})

	Describe("Update custom roles", func() {
		It("Successfully updated a custom role", func() {
			shouldSkipTest()
			Expect(testCustomRoleId).To(Not(BeNil()))
			Expect(testPolicyETag).To(Not(BeNil()))

			actions := []string{"iam-groups.groups.read"}
			options := service.NewReplaceRoleOptions(
				testCustomRoleId,
				testCustomRoleETag,
				testCustomRoleName,
				actions,
			)
			options.SetDescription("GO SDK test role udpated")
			options.SetDisplayName("GO SDK test role udpated")
			result, detailedResponse, err := service.ReplaceRole(options)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ReplaceRole() result:\n%s\n", common.ToJSON(result))
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testCustomRoleId))

		})
	})

	Describe("List custom roles", func() {
		It("Successfully listed the account's custom roles", func() {
			shouldSkipTest()
			Expect(testCustomRoleId).To(Not(BeNil()))

			options := service.NewListRolesOptions()
			options.SetAccountID(testAccountID)
			result, detailedResponse, err := service.ListRoles(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListRoles() result:\n%s\n", common.ToJSON(result))

			// confirm the test policy is present
			testRolePresent := false
			for _, role := range result.CustomRoles {
				if *role.ID == testCustomRoleId {
					testRolePresent = true
				}
			}
			Expect(testRolePresent).To(BeTrue())
		})
	})

	Describe("List V2 roles", func() {
		It("Successfully listed the roles when account_id and service_group_id present", func() {
			shouldSkipTest()

			options := service.NewListRolesOptions()
			options.SetAccountID(testAccountID)
			options.SetServiceGroupID("IAM")
			result, detailedResponse, err := service.ListRoles(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListRoles() result:\n%s\n", common.ToJSON(result))

			// confirm the system's viewer and service roles are present
			testSystemRolePresent := false
			testServiceRolePresent := false
			for _, role := range result.SystemRoles {
				if *role.CRN == testViewerRoleCrn {
					testSystemRolePresent = true
				}
			}

			for _, role := range result.ServiceRoles {
				if *role.CRN == testServiceRoleCrn {
					testServiceRolePresent = true
				}
			}

			Expect(testSystemRolePresent).To(BeTrue())
			Expect(testServiceRolePresent).To(BeTrue())
		})
	})

	Describe(`CreatePolicyTemplate - Create a policy template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicyTemplate(createPolicyTemplateOptions *CreatePolicyTemplateOptions)`, func() {
			v2PolicyResourceAttributeModel := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("iam-access-management"),
			}

			v2PolicyResourceModel := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel},
			}

			rolesModel := &iampolicymanagementv1.Roles{
				RoleID: core.StringPtr(testViewerRoleCrn),
			}

			grantModel := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*rolesModel},
			}

			controlModel := &iampolicymanagementv1.Control{
				Grant: grantModel,
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Test Policy For Template"),
				Resource:    v2PolicyResourceModel,
				Control:     controlModel,
			}

			createPolicyTemplateOptions := &iampolicymanagementv1.CreatePolicyTemplateOptions{
				Name:           &examplePolicyTemplateName,
				AccountID:      &testAccountID,
				Policy:         templatePolicyModel,
				Description:    core.StringPtr("Test PolicySampleTemplate"),
				Committed:      core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("default"),
			}

			policyTemplate, response, err := service.CreatePolicyTemplate(createPolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())
			Expect(policyTemplate.Name).To(Equal(core.StringPtr(examplePolicyTemplateName)))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))

			testPolicyTemplateID = *policyTemplate.ID
		})
	})

	Describe(`ListPolicyTemplates - Get policy templates by attributes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPolicyTemplates(listPolicyTemplatesOptions *ListPolicyTemplatesOptions)`, func() {
			listPolicyTemplatesOptions := &iampolicymanagementv1.ListPolicyTemplatesOptions{
				AccountID:      &testAccountID,
				AcceptLanguage: core.StringPtr("default"),
			}

			policyTemplateCollection, response, err := service.ListPolicyTemplates(listPolicyTemplatesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplateCollection).ToNot(BeNil())

			Expect(policyTemplateCollection.PolicyTemplates[0].Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplateCollection.PolicyTemplates[0].AccountID).To(Equal(&testAccountID))
		})
	})

	Describe(`GetPolicyTemplate - Retrieve latest policy template version by template ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPolicyTemplate(getPolicyTemplateOptions *GetPolicyTemplateOptions)`, func() {
			getPolicyTemplateOptions := &iampolicymanagementv1.GetPolicyTemplateOptions{
				PolicyTemplateID: &testPolicyTemplateID,
			}

			policyTemplate, response, err := service.GetPolicyTemplate(getPolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())

			Expect(policyTemplate.Name).To(Equal(&examplePolicyTemplateName))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(&testAccountID))
		})
	})

	Describe(`CreatePolicyTemplateVersion - Create a new policy template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicyTemplateVersion(createPolicyTemplateVersionOptions *CreatePolicyTemplateVersionOptions)`, func() {
			v2PolicyResourceAttributeModel := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("watson"),
			}

			v2PolicyResourceModel := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel},
			}

			rolesModel := &iampolicymanagementv1.Roles{
				RoleID: &testEditorRoleCrn,
			}

			grantModel := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*rolesModel},
			}

			controlModel := &iampolicymanagementv1.Control{
				Grant: grantModel,
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Watson Policy Template"),
				Resource:    v2PolicyResourceModel,
				Control:     controlModel,
			}

			createPolicyTemplateVersionOptions := &iampolicymanagementv1.CreatePolicyTemplateVersionOptions{
				PolicyTemplateID: &testPolicyTemplateID,
				Policy:           templatePolicyModel,
				Description:      core.StringPtr("Watson Policy Template version"),
			}

			policyTemplate, response, err := service.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())

			Expect(policyTemplate.Name).To(Equal(&examplePolicyTemplateName))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(&testAccountID))

			testPolicyTemplateVersion = *policyTemplate.Version
			testPolicyTemplateETag = response.GetHeaders().Get(etagHeader)
		})
	})

	Describe(`ListPolicyTemplateVersions - Retrieve policy template versions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPolicyTemplateVersions(listPolicyTemplateVersionsOptions *ListPolicyTemplateVersionsOptions)`, func() {
			listPolicyTemplateVersionsOptions := &iampolicymanagementv1.ListPolicyTemplateVersionsOptions{
				PolicyTemplateID: &testPolicyTemplateID,
			}

			policyTemplateVersionsCollection, response, err := service.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplateVersionsCollection).ToNot(BeNil())
		})
	})

	Describe(`ReplacePolicyTemplate - Update a policy template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplacePolicyTemplate(replacePolicyTemplateOptions *ReplacePolicyTemplateOptions)`, func() {
			v2PolicyResourceAttributeModel := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("watson"),
			}

			v2PolicyResourceModel := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel},
			}

			rolesModel := &iampolicymanagementv1.Roles{
				RoleID: &testViewerRoleCrn,
			}

			grantModel := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*rolesModel},
			}

			controlModel := &iampolicymanagementv1.Control{
				Grant: grantModel,
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Version Update"),
				Resource:    v2PolicyResourceModel,
				Control:     controlModel,
			}

			replacePolicyTemplateOptions := &iampolicymanagementv1.ReplacePolicyTemplateOptions{
				PolicyTemplateID: &testPolicyTemplateID,
				Version:          &testPolicyTemplateVersion,
				IfMatch:          &testPolicyTemplateETag,
				Policy:           templatePolicyModel,
				Description:      core.StringPtr("Template version update"),
			}

			policyTemplate, response, err := service.ReplacePolicyTemplate(replacePolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())

			Expect(policyTemplate.Version).To(Equal(core.StringPtr("2")))
			Expect(policyTemplate.Name).To(Equal(&examplePolicyTemplateName))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(&testAccountID))

			testPolicyTemplateETag = response.GetHeaders().Get(etagHeader)

		})
	})

	Describe(`GetPolicyTemplateVersion - Retrieve a policy template version by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPolicyTemplateVersion(getPolicyTemplateVersionOptions *GetPolicyTemplateVersionOptions)`, func() {
			getPolicyTemplateVersionOptions := &iampolicymanagementv1.GetPolicyTemplateVersionOptions{
				PolicyTemplateID: &testPolicyTemplateID,
				Version:          &testPolicyTemplateVersion,
			}

			policyTemplate, response, err := service.GetPolicyTemplateVersion(getPolicyTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())

			Expect(policyTemplate.Version).To(Equal(core.StringPtr("2")))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(&testAccountID))
		})
	})

	Describe(`CommitPolicyTemplate - Commit a policy template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CommitPolicyTemplate(commitPolicyTemplateOptions *CommitPolicyTemplateOptions)`, func() {
			commitPolicyTemplateOptions := &iampolicymanagementv1.CommitPolicyTemplateOptions{
				PolicyTemplateID: &testPolicyTemplateID,
				Version:          &testPolicyTemplateVersion,
				IfMatch:          &testPolicyTemplateETag,
			}

			response, err := service.CommitPolicyTemplate(commitPolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`ListPolicyAssignments - Get policies template assignments by attributes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPolicyAssignments(listPolicyAssignmentsOptions *ListPolicyAssignmentsOptions)`, func() {
			listPolicyAssignmentsOptions := &iampolicymanagementv1.ListPolicyAssignmentsOptions{
				AccountID:      core.StringPtr(testAccountID),
				AcceptLanguage: core.StringPtr("default"),
			}

			polcyTemplateAssignmentCollection, response, err := service.ListPolicyAssignments(listPolicyAssignmentsOptions)
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
	})

	Describe(`GetPolicyAssignment - Retrieve a policy assignment by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPolicyAssignment(getPolicyAssignmentOptions *GetPolicyAssignmentOptions)`, func() {
			getPolicyAssignmentOptions := &iampolicymanagementv1.GetPolicyAssignmentOptions{
				AssignmentID: core.StringPtr(testPolicyAssignmentId),
			}

			policyAssignmentRecord, response, err := service.GetPolicyAssignment(getPolicyAssignmentOptions)
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
	})

	Describe(`DeletePolicyTemplateVersion - Delete a policy template version by ID and version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePolicyTemplateVersion(deletePolicyTemplateVersionOptions *DeletePolicyTemplateVersionOptions)`, func() {
			deletePolicyTemplateVersionOptions := &iampolicymanagementv1.DeletePolicyTemplateVersionOptions{
				PolicyTemplateID: &testPolicyTemplateID,
				Version:          &testPolicyTemplateVersion,
			}

			response, err := service.DeletePolicyTemplateVersion(deletePolicyTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePolicyTemplate - Delete a policy template by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePolicyTemplate(deletePolicyTemplateOptions *DeletePolicyTemplateOptions)`, func() {
			deletePolicyTemplateOptions := &iampolicymanagementv1.DeletePolicyTemplateOptions{
				PolicyTemplateID: &testPolicyTemplateID,
			}

			response, err := service.DeletePolicyTemplate(deletePolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	// clean up all test groups
	AfterSuite(func() {
		if !configLoaded {
			return
		}

		fmt.Fprintf(GinkgoWriter, "Cleaning up test groups...\n")

		// list all policies in the account
		policyOptions := service.NewListPoliciesOptions(testAccountID)
		policyOptions.SetIamID(testUserId)
		policyResult, policyDetailedResponse, err := service.ListPolicies(policyOptions)
		Expect(err).To(BeNil())
		Expect(policyDetailedResponse.StatusCode).To(Equal(200))

		for _, policy := range policyResult.Policies {

			// delete the test policy (or any test policy older than 5 minutes)
			createdAt, err := time.Parse(time.RFC3339, policy.CreatedAt.String())
			if err != nil {
				fmt.Fprintf(GinkgoWriter, "time.Parse error occurred: %v\n", err)
				fmt.Fprintf(GinkgoWriter, "Cleanup of policy (%v) failed\n", *policy.ID)
				continue
			}
			fiveMinutesAgo := time.Now().Add(-(time.Duration(5) * time.Minute))
			if strings.Contains(*policy.Href, "v2/policies") {
				if *policy.ID == testV2PolicyId || createdAt.Before(fiveMinutesAgo) {
					options := service.NewDeleteV2PolicyOptions(*policy.ID)
					detailedResponse, err := service.DeleteV2Policy(options)
					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(204))
				}
			} else {
				if *policy.ID == testPolicyId || createdAt.Before(fiveMinutesAgo) {
					options := service.NewDeletePolicyOptions(*policy.ID)
					detailedResponse, err := service.DeletePolicy(options)
					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(204))
				}
			}
		}

		// List all custom roles in the account
		roleOptions := service.NewListRolesOptions()
		roleOptions.SetAccountID(testAccountID)
		roleResult, roleDetailedResponse, err := service.ListRoles(roleOptions)
		Expect(err).To(BeNil())
		Expect(roleDetailedResponse.StatusCode).To(Equal(200))

		for _, role := range roleResult.CustomRoles {

			// delete the role (or any test role older than 5 minutes)
			createdAt, err := time.Parse(time.RFC3339, role.CreatedAt.String())
			if err != nil {
				fmt.Fprintf(GinkgoWriter, "time.Parse error occurred: %v\n", err)
				fmt.Fprintf(GinkgoWriter, "Cleanup of role (%v) failed\n", *role.ID)
				continue
			}
			fiveMinutesAgo := time.Now().Add(-(time.Duration(5) * time.Minute))

			if *role.ID == testCustomRoleId || createdAt.Before(fiveMinutesAgo) {
				options := service.NewDeleteRoleOptions(*role.ID)
				detailedResponse, err := service.DeleteRole(options)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			}
		}

		fmt.Fprintf(GinkgoWriter, "Cleanup finished!\n")
	})
})
