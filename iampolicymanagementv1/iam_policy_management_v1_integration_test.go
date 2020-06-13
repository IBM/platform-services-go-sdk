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
	"math/rand"
	"os"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/go-sdk-core/v4/core"
  "github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

 const externalConfigFile = "../iam_policy_management.env"

 var (
	 service             *iampolicymanagementv1.IamPolicyManagementV1
	 err                 error
	 config              map[string]string
	 configLoaded        bool   = false

	 testAccountID       string
   etagHeader          string = "ETag"
   testPolicyETag      string = ""
   testPolicyId        string = ""
   testUserId          string = "IBMid-GoSDK" + strconv.Itoa(rand.Intn(100000))
   testViewerRoleCrn   string = "crn:v1:bluemix:public:iam::::role:Viewer"
   testEditorRoleCrn   string = "crn:v1:bluemix:public:iam::::role:Editor"
   testServiceName     string = "iam-groups"

   testCustomRoleId    string = ""
   testCustomRoleETag  string = ""
   testCustomRoleName  string = "TestGoRole" + strconv.Itoa(rand.Intn(100000))
 )

 func shouldSkipTest() {
	 if !configLoaded {
		 Skip("External configuration is not available, skipping...")
	 }
 }

 var _ = Describe("IAM Policy Management - Integration Tests", func() {
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
       serviceNameResourceAttribute.Name = core.StringPtr("serviceName")
       serviceNameResourceAttribute.Value = core.StringPtr(testServiceName)
       serviceNameResourceAttribute.Operator = core.StringPtr("stringEquals")

       // Construct an instance of the SubjectAttribute model
       subjectAttribute := new(iampolicymanagementv1.SubjectAttribute)
       subjectAttribute.Name = core.StringPtr("iam_id")
       subjectAttribute.Value = core.StringPtr(testUserId)

       // Construct an instance of the PolicyResource model
       policyResource := new(iampolicymanagementv1.PolicyResource)
       policyResource.Attributes = []iampolicymanagementv1.ResourceAttribute{*accountIdResourceAttribute, *serviceNameResourceAttribute}

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
			 result, detailedResponse, err := service.GetPolicy(options)
			 Expect(err).To(BeNil())
			 Expect(detailedResponse.StatusCode).To(Equal(200))
			 Expect(*result.ID).To(Equal(testPolicyId))

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
       serviceNameResourceAttribute.Name = core.StringPtr("serviceName")
       serviceNameResourceAttribute.Value = core.StringPtr(testServiceName)
       serviceNameResourceAttribute.Operator = core.StringPtr("stringEquals")

       // Construct an instance of the SubjectAttribute model
       subjectAttribute := new(iampolicymanagementv1.SubjectAttribute)
       subjectAttribute.Name = core.StringPtr("iam_id")
       subjectAttribute.Value = core.StringPtr(testUserId)

       // Construct an instance of the PolicyResource model
       policyResource := new(iampolicymanagementv1.PolicyResource)
       policyResource.Attributes = []iampolicymanagementv1.ResourceAttribute{*accountIdResourceAttribute, *serviceNameResourceAttribute}

       // Construct an instance of the PolicyRole model
       policyRole := new(iampolicymanagementv1.PolicyRole)
       policyRole.RoleID = core.StringPtr(testEditorRoleCrn)

       // Construct an instance of the PolicySubject model
       policySubject := new(iampolicymanagementv1.PolicySubject)
       policySubject.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttribute}

       // Construct an instance of the CreatePolicyOptions model
       options := new(iampolicymanagementv1.UpdatePolicyOptions)
       options.PolicyID = core.StringPtr(testPolicyId)
       options.IfMatch = core.StringPtr(testPolicyETag)
       options.Type = core.StringPtr("access")
       options.Subjects = []iampolicymanagementv1.PolicySubject{*policySubject}
       options.Roles = []iampolicymanagementv1.PolicyRole{*policyRole}
       options.Resources = []iampolicymanagementv1.PolicyResource{*policyResource}

			 policy, detailedResponse, err := service.UpdatePolicy(options)
			 Expect(err).To(BeNil())
			 Expect(detailedResponse.StatusCode).To(Equal(200))
			 Expect(*policy.ID).To(Equal(testPolicyId))
       Expect(policy.Type).To(Equal(options.Type))
       Expect(policy.Subjects).To(Equal(options.Subjects))
       Expect(policy.Roles[0].RoleID).To(Equal(options.Roles[0].RoleID))
       Expect(policy.Resources).To(Equal(options.Resources))

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
      Expect(*result.ID).To(Equal(testCustomRoleId))

      testCustomRoleETag = detailedResponse.GetHeaders().Get(etagHeader)
    })
  })

  Describe("Update custom roles", func() {
    It("Successfully updated a custom role", func() {
      shouldSkipTest()
      Expect(testCustomRoleId).To(Not(BeNil()))
      Expect(testPolicyETag).To(Not(BeNil()))

      options := service.NewUpdateRoleOptions(
        testCustomRoleId,
        testCustomRoleETag)
      options.SetDescription("GO SDK test role udpated")
      options.SetDisplayName("GO SDK test role udpated")
      result, detailedResponse, err := service.UpdateRole(options)
      Expect(err).To(BeNil())
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

})


 // clean up all test groups
 var _ = AfterSuite(func() {
	 if !configLoaded {
		 return
	 }

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
			 fmt.Printf("time.Parse error occurred: %v", err)
			 fmt.Printf("Cleanup of policy (%v) failed", *policy.ID)
			 continue
		 }
		 fiveMinutesAgo := time.Now().Add(-(time.Duration(5) * time.Minute))

		 if *policy.ID == testPolicyId || createdAt.Before(fiveMinutesAgo) {
			 options := service.NewDeletePolicyOptions(*policy.ID)
			 detailedResponse, err := service.DeletePolicy(options)
			 Expect(err).To(BeNil())
			 Expect(detailedResponse.StatusCode).To(Equal(204))
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
			 fmt.Printf("time.Parse error occurred: %v", err)
			 fmt.Printf("Cleanup of role (%v) failed", *role.ID)
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

 })
