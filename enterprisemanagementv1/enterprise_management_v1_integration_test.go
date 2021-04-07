// +build integration

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

package enterprisemanagementv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the enterprisemanagementv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`EnterpriseManagementV1 Integration Tests`, func() {

	const externalConfigFile = "../enterprise_management.env"

	var (
		err                         error
		enterpriseManagementService *enterprisemanagementv1.EnterpriseManagementV1
		serviceURL                  string
		testConfig                  map[string]string

		enterpriseId                        string
		accountId                           string
		accountIamId                        string
		firstExampleAccountGroupName        = "Example Account Group"
		firstUpdatedExampleAccountGroupName = "Updated Example Account Group"
		secondExampleAccountGroupName       = "Second Example Account Group"
		exampleAccountName                  = "Example Account Name"
		exampleAccountId                    *string
		resultPerPage                       int64 = 1
		firstExampleAccountGroupId          *string
		secondExampleAccountGroupId         *string
		updatedEnterpriseName               = "Updated Enterprise Name"
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			testConfig, err = core.GetServiceProperties(enterprisemanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = testConfig["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			enterpriseId = testConfig["ENTERPRISE_ID"]
			Expect(enterpriseId).ToNot(BeEmpty())

			accountId = testConfig["ACCOUNT_ID"]
			Expect(accountId).ToNot(BeEmpty())

			accountIamId = testConfig["ACCOUNT_IAM_ID"]
			Expect(accountIamId).NotTo(BeEmpty())

			fmt.Fprintf(GinkgoWriter, "Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			enterpriseManagementServiceOptions := &enterprisemanagementv1.EnterpriseManagementV1Options{}

			enterpriseManagementService, err = enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(enterpriseManagementServiceOptions)

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			enterpriseManagementService.EnableRetries(4, 30*time.Second)

			Expect(err).To(BeNil())
			Expect(enterpriseManagementService).ToNot(BeNil())
			Expect(enterpriseManagementService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`CreateAccountGroup - Create an account group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAccountGroup(createAccountGroupOptions *CreateAccountGroupOptions)`, func() {
			var parent = "crn:v1:bluemix:public:enterprise::a/" + accountId + "::enterprise:" + enterpriseId
			createFirstAccountGroupOptions := &enterprisemanagementv1.CreateAccountGroupOptions{
				Parent:              &parent,
				Name:                &firstExampleAccountGroupName,
				PrimaryContactIamID: &accountIamId,
			}

			createFirstExampleAccountGroupResponse, firstExampleResponse, firstError := enterpriseManagementService.CreateAccountGroup(createFirstAccountGroupOptions)

			Expect(firstError).To(BeNil())
			Expect(firstExampleResponse.StatusCode).To(Equal(201))
			Expect(createFirstExampleAccountGroupResponse).ToNot(BeNil())

			firstExampleAccountGroupId = createFirstExampleAccountGroupResponse.AccountGroupID

			createSecondAccountGroupOptions := &enterprisemanagementv1.CreateAccountGroupOptions{
				Parent:              &parent,
				Name:                &secondExampleAccountGroupName,
				PrimaryContactIamID: &accountIamId,
			}

			createSecondExampleAccountGroupResponse, secondExampleResponse, secondError := enterpriseManagementService.CreateAccountGroup(createSecondAccountGroupOptions)

			Expect(secondError).To(BeNil())
			Expect(secondExampleResponse.StatusCode).To(Equal(201))
			Expect(createSecondExampleAccountGroupResponse).ToNot(BeNil())

			secondExampleAccountGroupId = createSecondExampleAccountGroupResponse.AccountGroupID
		})
	})

	Describe(`ListAccountGroups - List account groups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAccountGroups(listAccountGroupsOptions *ListAccountGroupsOptions)`, func() {

			listAccountGroupsOptions := &enterprisemanagementv1.ListAccountGroupsOptions{
				EnterpriseID: &enterpriseId,
			}

			listAccountGroupsResponse, response, err := enterpriseManagementService.ListAccountGroups(listAccountGroupsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listAccountGroupsResponse).ToNot(BeNil())

		})
	})

	Describe(`ListAccountGroups - List account groups with paging`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAccountGroups(listAccountGroupsOptions *ListAccountGroupsOptions)`, func() {

			listAccountGroupsOptionsPage1 := &enterprisemanagementv1.ListAccountGroupsOptions{
				EnterpriseID: &enterpriseId,
				Limit:        &resultPerPage,
			}

			listAccountGroupsResponsePage1, responsePage1, errorPage1 := enterpriseManagementService.ListAccountGroups(listAccountGroupsOptionsPage1)

			Expect(errorPage1).To(BeNil())
			Expect(responsePage1.StatusCode).To(Equal(200))
			Expect(listAccountGroupsResponsePage1).ToNot(BeNil())
			Expect(*listAccountGroupsResponsePage1.RowsCount).To(Equal(resultPerPage))

			docID, err := core.GetQueryParam(listAccountGroupsResponsePage1.NextURL, "next_docid")
			Expect(err).To(BeNil())

			listAccountGroupsOptionsPage2 := &enterprisemanagementv1.ListAccountGroupsOptions{
				EnterpriseID: &enterpriseId,
				NextDocid:    docID,
				Limit:        &resultPerPage,
			}
			listAccountGroupsResponsePage2, responsePage2, errorPage2 := enterpriseManagementService.ListAccountGroups(listAccountGroupsOptionsPage2)

			Expect(errorPage2).To(BeNil())
			Expect(responsePage2.StatusCode).To(Equal(200))
			Expect(listAccountGroupsResponsePage2).ToNot(BeNil())
			Expect(*listAccountGroupsResponsePage2.RowsCount).To(Equal(resultPerPage))
		})
	})

	Describe(`GetAccountGroup - Get account group by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountGroup(getAccountGroupOptions *GetAccountGroupOptions)`, func() {

			getAccountGroupOptions := &enterprisemanagementv1.GetAccountGroupOptions{
				AccountGroupID: firstExampleAccountGroupId,
			}

			accountGroup, response, err := enterpriseManagementService.GetAccountGroup(getAccountGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountGroup).ToNot(BeNil())

		})
	})

	Describe(`UpdateAccountGroup - Update an account group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAccountGroup(updateAccountGroupOptions *UpdateAccountGroupOptions)`, func() {

			updateAccountGroupOptions := &enterprisemanagementv1.UpdateAccountGroupOptions{
				AccountGroupID:      firstExampleAccountGroupId,
				Name:                &firstUpdatedExampleAccountGroupName,
				PrimaryContactIamID: &accountIamId,
			}

			response, err := enterpriseManagementService.UpdateAccountGroup(updateAccountGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`CreateAccount - Create a new account in an enterprise`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAccount(createAccountOptions *CreateAccountOptions)`, func() {
			var parent = "crn:v1:bluemix:public:enterprise::a/" + accountId + "::account-group:" + *firstExampleAccountGroupId
			createAccountOptions := &enterprisemanagementv1.CreateAccountOptions{
				Parent:     &parent,
				Name:       &exampleAccountName,
				OwnerIamID: &accountIamId,
			}

			createAccountResponse, response, err := enterpriseManagementService.CreateAccount(createAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createAccountResponse).ToNot(BeNil())

			exampleAccountId = createAccountResponse.AccountID

		})
	})

	Describe(`ListAccounts - List accounts`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAccounts(listAccountsOptions *ListAccountsOptions)`, func() {

			listAccountsOptions := &enterprisemanagementv1.ListAccountsOptions{
				EnterpriseID: &enterpriseId,
			}

			listAccountsResponse, response, err := enterpriseManagementService.ListAccounts(listAccountsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listAccountsResponse).ToNot(BeNil())

		})
	})

	Describe(`ListAccounts - List accounts with pagination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAccounts(listAccountsOptions *ListAccountsOptions)`, func() {

			listAccountsOptionsPage1 := &enterprisemanagementv1.ListAccountsOptions{
				EnterpriseID: &enterpriseId,
				Limit:        &resultPerPage,
			}

			listAccountsResponsePage1, responsePage1, errorPage1 := enterpriseManagementService.ListAccounts(listAccountsOptionsPage1)

			Expect(errorPage1).To(BeNil())
			Expect(responsePage1.StatusCode).To(Equal(200))
			Expect(listAccountsResponsePage1).ToNot(BeNil())
			Expect(*listAccountsResponsePage1.RowsCount).To(Equal(resultPerPage))

			docID, err := core.GetQueryParam(listAccountsResponsePage1.NextURL, "next_docid")
			Expect(err).To(BeNil())

			listAccountsOptionsPage2 := &enterprisemanagementv1.ListAccountsOptions{
				EnterpriseID: &enterpriseId,
				Limit:        &resultPerPage,
				NextDocid:    docID,
			}

			listAccountsResponsePage2, responsePage2, errorPage2 := enterpriseManagementService.ListAccounts(listAccountsOptionsPage2)

			Expect(errorPage2).To(BeNil())
			Expect(responsePage2.StatusCode).To(Equal(200))
			Expect(listAccountsResponsePage2).ToNot(BeNil())
			Expect(*listAccountsResponsePage2.RowsCount).To(Equal(resultPerPage))

		})
	})

	Describe(`GetAccount - Get account by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccount(getAccountOptions *GetAccountOptions)`, func() {

			getAccountOptions := &enterprisemanagementv1.GetAccountOptions{
				AccountID: exampleAccountId,
			}

			account, response, err := enterpriseManagementService.GetAccount(getAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(account).ToNot(BeNil())

		})
	})

	Describe(`UpdateAccount - Move an account within the enterprise`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAccount(updateAccountOptions *UpdateAccountOptions)`, func() {

			var newParent = "crn:v1:bluemix:public:enterprise::a/" + accountId + "::account-group:" + *secondExampleAccountGroupId
			updateAccountOptions := &enterprisemanagementv1.UpdateAccountOptions{
				AccountID: exampleAccountId,
				Parent:    &newParent,
			}

			response, err := enterpriseManagementService.UpdateAccount(updateAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})

	Describe(`ListEnterprises - List enterprises`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListEnterprises(listEnterprisesOptions *ListEnterprisesOptions)`, func() {

			listEnterprisesOptions := &enterprisemanagementv1.ListEnterprisesOptions{
				AccountID: &accountId,
			}

			listEnterprisesResponse, response, err := enterpriseManagementService.ListEnterprises(listEnterprisesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listEnterprisesResponse).ToNot(BeNil())

		})
	})

	Describe(`GetEnterprise - Get enterprise by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions)`, func() {

			getEnterpriseOptions := &enterprisemanagementv1.GetEnterpriseOptions{
				EnterpriseID: &enterpriseId,
			}

			enterprise, response, err := enterpriseManagementService.GetEnterprise(getEnterpriseOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(enterprise).ToNot(BeNil())

		})
	})

	Describe(`UpdateEnterprise - Update an enterprise`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateEnterprise(updateEnterpriseOptions *UpdateEnterpriseOptions)`, func() {

			updateEnterpriseOptions := &enterprisemanagementv1.UpdateEnterpriseOptions{
				EnterpriseID:        &enterpriseId,
				Name:                &updatedEnterpriseName,
				PrimaryContactIamID: &accountIamId,
			}

			response, err := enterpriseManagementService.UpdateEnterprise(updateEnterpriseOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
