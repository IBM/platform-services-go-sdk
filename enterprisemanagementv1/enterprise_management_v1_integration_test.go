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

package enterprisemanagementv1_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
)

const externalConfigFile = "../enterprise-management.env"

var (
	configLoaded bool = false
	service      *enterprisemanagementv1.EnterpriseManagementV1
	testConfig   map[string]string
	amAuth       *core.IamAuthenticator
	email        string = "aminttest+" + strconv.Itoa(rand.Intn(100000)) + "_" + strconv.Itoa(rand.Intn(100000)) + "@mail.test.ibm.com"

	account_id            string
	activationId          string
	owner_iam_id          string
	subscription_id       string
	parent                string
	err                   error
	enterprise_id         string
	enterprise_account_id string
	accountGroupID        string
	crn                   string
	email2                string = "aminttest+" + strconv.Itoa(rand.Intn(1000000)) + "_" + strconv.Itoa(rand.Intn(100000)) + "@mail.test.ibm.com"
	standard_account_id   string
	newAccount            string
	accountGroupID2       string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe("Enterprise Management - Integration Tests", func() {

	It("Successfully load the configuration", func() {
		_, err = os.Stat(externalConfigFile)
		if err == nil {
			err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			if err == nil {
				configLoaded = true
			}
		}
		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}

		options := &enterprisemanagementv1.EnterpriseManagementV1Options{}
		service, err = enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(options)
		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())

		core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
		service.EnableRetries(4, 30*time.Second)

		var svcConfig map[string]string
		svcConfig, err = core.GetServiceProperties(enterprisemanagementv1.DefaultServiceName)
		Expect(err).To(BeNil())
		Expect(svcConfig).ToNot(BeNil())
		Expect(len(svcConfig)).To(Equal(4))

		testConfig, err = core.GetServiceProperties("EMTEST_CONFIG")
		Expect(err).To(BeNil())
		Expect(testConfig).ToNot(BeNil())
		Expect(testConfig["AM_HOST"]).ToNot(BeNil())
		Expect(testConfig["DB_URL"]).ToNot(BeNil())
		Expect(testConfig["DB_USER"]).ToNot(BeNil())
		Expect(testConfig["DB_PASS"]).ToNot(BeNil())
		Expect(testConfig["ACTIVATION_DB_NAME"]).ToNot(BeNil())
		Expect(testConfig["IAM_API_KEY"]).ToNot(BeNil())

		// Construct an IamAuthenticator to use with the Account Mgmt API.
		amAuth = &core.IamAuthenticator{
			URL:    svcConfig["AUTH_URL"],
			ApiKey: testConfig["IAM_API_KEY"],
		}
	})

	It("Successfully create a standard account", func() {
		shouldSkipTest()

		apiUrl := testConfig["AM_HOST"]
		resource := "/coe/v2/accounts"

		u, err := url.ParseRequestURI(apiUrl)
		Expect(err).To(BeNil())

		u.Path = resource
		urlStr := u.String()

		type Subscription struct {
			Type        string `json:"type"`
			State       string `json:"state"`
			Part_number string `json:"part_number"`
		}

		type Ibmid struct {
			Password string `json:"password"`
			Question string `json:"question"`
			Answer   string `json:"answer"`
		}

		type Payload struct {
			Owner_user_id    string         `json:"owner_user_id"`
			Owner_email      string         `json:"owner_email"`
			Owner_first_name string         `json:"owner_first_name"`
			Owner_last_name  string         `json:"owner_last_name"`
			Owner_phone      string         `json:"owner_phone"`
			Owner_company    string         `json:"owner_company"`
			Country_code     string         `json:"country_code"`
			Subscriptions    []Subscription `json:"bluemix_subscriptions"`
			Ibmids           Ibmid          `json:"ibmid"`
		}

		accountPayload := Payload{
			Owner_user_id:    email,
			Owner_email:      email,
			Owner_first_name: "TEST",
			Owner_last_name:  "TEST",
			Owner_phone:      "123456789",
			Owner_company:    "IBM",
			Country_code:     "USA",
			Subscriptions: []Subscription{
				Subscription{
					Type:        "STANDARD",
					State:       "ACTIVE",
					Part_number: "COE-Trial",
				},
			},
			Ibmids: Ibmid{
				Password: "password",
				Question: "question",
				Answer:   "answer",
			},
		}

		// Serialize the request body
		accountPayloadJson, _ := json.Marshal(accountPayload)

		url := urlStr

		// Create a new request using http
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(accountPayloadJson))
		Expect(err).To(BeNil())

		err = amAuth.Authenticate(req)
		Expect(err).To(BeNil())

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		Expect(err).To(BeNil())

		var ok bool
		account_id, ok = data["id"].(string)
		Expect(ok).To(BeTrue())
	})

	It("Successfully get activation code - email", func() {
		shouldSkipTest()

		time.Sleep(20000 * time.Millisecond)

		apiUrl := testConfig["AM_HOST"]
		resource := "/v1/activation-codes/" + email

		u, err := url.ParseRequestURI(apiUrl)
		Expect(err).To(BeNil())

		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		Expect(err).To(BeNil())

		err = amAuth.Authenticate(req)
		Expect(err).To(BeNil())

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).To(BeNil())

		var results map[string]interface{} // TopTracks

		err = json.Unmarshal(body, &results)
		Expect(err).To(BeNil())

		res := results["resources"].([]interface{})

		z := res[0].(map[string]interface{})

		var ok bool
		activationId, ok = z["id"].(string)
		Expect(ok).To(BeTrue())
	})

	It("Successfully activate account", func() {
		shouldSkipTest()

		apiUrl := testConfig["AM_HOST"]
		resource := "/coe/v2/accounts/verify"

		u, err := url.ParseRequestURI(apiUrl)
		Expect(err).To(BeNil())

		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		Expect(err).To(BeNil())

		err = amAuth.Authenticate(req)
		Expect(err).To(BeNil())

		q := req.URL.Query()
		q.Add("token", activationId)
		q.Add("email", email)
		req.URL.RawQuery = q.Encode()

		// Send req using http Client
		client := &http.Client{}

		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).To(BeNil())
		Expect(body).ToNot(BeNil())
	})

	It("Successfully get account", func() {
		shouldSkipTest()

		apiUrl := testConfig["AM_HOST"]
		resource := "/coe/v2/accounts/" + account_id
		u, err := url.ParseRequestURI(apiUrl)
		Expect(err).To(BeNil())

		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		Expect(err).To(BeNil())

		err = amAuth.Authenticate(req)
		Expect(err).To(BeNil())

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		Expect(err).To(BeNil())

		res := data["entity"].(map[string]interface{})
		owner_iam_id = res["owner_iam_id"].(string)

		var ok bool
		subscription_id, ok = res["subscription_id"].(string)
		Expect(ok).To(BeTrue())
	})

	It("Successfully convert account from STANDARD to SUBSCRIPTION", func() {
		shouldSkipTest()

		apiUrl := testConfig["AM_HOST"]
		resource := "/coe/v2/accounts/" + account_id + "/bluemix_subscriptions/" + subscription_id
		u, err := url.ParseRequestURI(apiUrl)
		Expect(err).To(BeNil())

		u.Path = resource
		urlStr := u.String()

		url := urlStr

		type Payment_method struct {
			Start_date string `json:"start_date"`
			End_date   string `json:"end_date"`
		}

		type Payload struct {
			Type                    string         `json:"type"`
			State                   string         `json:"state"`
			Payment_methods         Payment_method `json:"payment_method"`
			Part_number             string         `json:"part_number"`
			Subscription_amount     int            `json:"subscription_amount"`
			Quantity                int            `json:"quantity"`
			Billing_frequency       string         `json:"billing_frequency"`
			Charge_agreement_number string         `json:"charge_agreement_number"`
			Partner_customer_number string         `json:"partner_customer_number"`
			Configuration_id        string         `json:"configuration_id"`
			Order_id_number         string         `json:"order_id_number"`
			Sales_doc_type_code     string         `json:"sales_doc_type_code"`
			Renewal_mode_code       string         `json:"renewal_mode_code"`
			Renewal_date            string         `json:"renewal_date"`
			Terminate_renewal       bool           `json:"terminate_renewal"`
			Line_item_id            int            `json:"line_item_id"`
		}

		accountPayload := Payload{
			Type:  "SUBSCRIPTION",
			State: "ACTIVE",
			Payment_methods: Payment_method{
				Start_date: "2020-03-01T07:00:00.000Z",
				End_date:   "2020-11-30T08:00:00.000Z",
			},
			Subscription_amount:     100,
			Quantity:                10,
			Billing_frequency:       "M",
			Charge_agreement_number: "0099342614",
			Partner_customer_number: "0003615466",
			Configuration_id:        "5900A5D20190517",
			Part_number:             "D019JZX",
			Order_id_number:         "150418156",
			Sales_doc_type_code:     "",
			Renewal_mode_code:       "T",
			Renewal_date:            "",
			Terminate_renewal:       false,
			Line_item_id:            10,
		}

		accountPayloadJson, _ := json.Marshal(accountPayload)

		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(accountPayloadJson))
		Expect(err).To(BeNil())

		err = amAuth.Authenticate(req)
		Expect(err).To(BeNil())

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		buf, err := ioutil.ReadAll(resp.Body)
		Expect(err).To(BeNil())
		Expect(buf).ToNot(BeNil())
	})

	It("Successfully create enterprise", func() {
		shouldSkipTest()

		options := service.NewCreateEnterpriseOptions(account_id, "IBM", owner_iam_id)

		result, detailedResponse, err := service.CreateEnterprise(options)
		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(202))

		enterprise_id = *result.EnterpriseID
		enterprise_account_id = *result.EnterpriseAccountID
		Expect(err).To(BeNil())

	})

	It("Successfully get account", func() {
		shouldSkipTest()

		apiUrl := testConfig["AM_HOST"]
		resource := "/coe/v2/accounts/" + account_id
		u, err := url.ParseRequestURI(apiUrl)
		Expect(err).To(BeNil())

		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		Expect(err).To(BeNil())

		err = amAuth.Authenticate(req)
		Expect(err).To(BeNil())

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		Expect(err).To(BeNil())

		res := data["entity"].(map[string]interface{})

		var ok bool
		parent, ok = res["parent"].(string)
		Expect(ok).To(BeTrue())
	})

	It("Successfully Create Account group", func() {
		shouldSkipTest()

		options := service.NewCreateAccountGroupOptions(parent, "IBM", owner_iam_id)

		result, detailedResponse, err := service.CreateAccountGroup(options)
		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(201))

		accountGroupID = *result.AccountGroupID
	})

	It("Successfully Create Account group", func() {
		shouldSkipTest()

		options := service.NewCreateAccountGroupOptions(parent, "IBM", owner_iam_id)

		result, detailedResponse, err := service.CreateAccountGroup(options)
		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(201))

		accountGroupID2 = *result.AccountGroupID
	})

	It("Successfully List Account groups", func() {
		shouldSkipTest()

		options := service.NewListAccountGroupsOptions()
		options.SetEnterpriseID(enterprise_id)
		options.SetParentAccountGroupID(accountGroupID)
		options.SetParent(parent)
		options.SetLimit(100)

		result, detailedResponse, err := service.ListAccountGroups(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))
	})

	It("Successfully Get Account group", func() {
		shouldSkipTest()

		options := service.NewGetAccountGroupOptions(accountGroupID)

		result, detailedResponse, err := service.GetAccountGroup(options)
		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))

		crn = *result.CRN
	})

	It("Successfully Update Account group", func() {
		shouldSkipTest()

		options := service.NewUpdateAccountGroupOptions(accountGroupID)
		options.SetName("IBM")
		options.SetPrimaryContactIamID(owner_iam_id)

		result, err := service.UpdateAccountGroup(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(result.StatusCode).To(Equal(204))
	})

	It("Successfully create a standard account", func() {
		shouldSkipTest()

		apiUrl := testConfig["AM_HOST"]
		resource := "/coe/v2/accounts"

		u, err := url.ParseRequestURI(apiUrl)
		Expect(err).To(BeNil())

		u.Path = resource
		urlStr := u.String()

		type Subscription struct {
			Type        string `json:"type"`
			State       string `json:"state"`
			Part_number string `json:"part_number"`
		}

		type Ibmid struct {
			Password string `json:"password"`
			Question string `json:"question"`
			Answer   string `json:"answer"`
		}

		type Payload struct {
			Owner_user_id    string         `json:"owner_user_id"`
			Owner_email      string         `json:"owner_email"`
			Owner_first_name string         `json:"owner_first_name"`
			Owner_last_name  string         `json:"owner_last_name"`
			Owner_phone      string         `json:"owner_phone"`
			Owner_company    string         `json:"owner_company"`
			Country_code     string         `json:"country_code"`
			Subscriptions    []Subscription `json:"bluemix_subscriptions"`
			Ibmids           Ibmid          `json:"ibmid"`
		}

		accountPayload := Payload{
			Owner_user_id:    email2,
			Owner_email:      email2,
			Owner_first_name: "TEST",
			Owner_last_name:  "TEST",
			Owner_phone:      "123456789",
			Owner_company:    "IBM",
			Country_code:     "USA",
			Subscriptions: []Subscription{
				Subscription{
					Type:        "STANDARD",
					State:       "ACTIVE",
					Part_number: "COE-Trial",
				},
			},
			Ibmids: Ibmid{
				Password: "password",
				Question: "question",
				Answer:   "answer",
			},
		}

		accountPayloadJson, _ := json.Marshal(accountPayload) //convert the struct to JSON format

		url := urlStr

		// Create a new request using http
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(accountPayloadJson))
		Expect(err).To(BeNil())

		// add authorization header to the req
		amAuth.Authenticate(req)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		Expect(err).To(BeNil())
		accountId := data["id"].(interface{})

		var ok bool
		standard_account_id, ok = accountId.(string)
		Expect(ok).To(BeTrue())
	})

	It("Successfully get activation code - email2", func() {
		shouldSkipTest()

		time.Sleep(20000 * time.Millisecond)

		apiUrl := testConfig["AM_HOST"]
		resource := "/v1/activation-codes/" + email2

		u, err := url.ParseRequestURI(apiUrl)
		Expect(err).To(BeNil())

		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		Expect(err).To(BeNil())

		err = amAuth.Authenticate(req)
		Expect(err).To(BeNil())

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).To(BeNil())

		var results map[string]interface{} // TopTracks

		err = json.Unmarshal(body, &results)
		Expect(err).To(BeNil())

		res := results["resources"].([]interface{})

		z := res[0].(map[string]interface{})

		var ok bool
		activationId, ok = z["id"].(string)
		Expect(ok).To(BeTrue())
	})

	It("Successfully activate account", func() {
		shouldSkipTest()

		apiUrl := testConfig["AM_HOST"]
		resource := "/coe/v2/accounts/verify"

		u, err := url.ParseRequestURI(apiUrl)
		Expect(err).To(BeNil())

		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		Expect(err).To(BeNil())

		q := req.URL.Query()
		q.Add("token", activationId)
		q.Add("email", email2)
		req.URL.RawQuery = q.Encode()

		err = amAuth.Authenticate(req)
		Expect(err).To(BeNil())

		// Send req using http Client
		client := &http.Client{}

		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).To(BeNil())
		Expect(body).ToNot(BeNil())
	})

	It("Successfully get account", func() {
		shouldSkipTest()

		apiUrl := testConfig["AM_HOST"]
		resource := "/coe/v2/accounts/" + standard_account_id
		u, err := url.ParseRequestURI(apiUrl)
		Expect(err).To(BeNil())

		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		Expect(err).To(BeNil())

		err = amAuth.Authenticate(req)
		Expect(err).To(BeNil())

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		Expect(err).To(BeNil())

		res := data["entity"].(map[string]interface{})
		owner_iam_id = res["owner_iam_id"].(string)

		var ok bool
		subscription_id, ok = res["subscription_id"].(string)
		Expect(ok).To(BeTrue())
	})

	It("Successfully Import Account to Enterprise", func() {
		shouldSkipTest()

		time.Sleep(20000 * time.Millisecond)
		options := service.NewImportAccountToEnterpriseOptions(enterprise_id, standard_account_id)
		options.SetParent(parent)

		response, err := service.ImportAccountToEnterprise(options)
		Expect(response).NotTo(BeNil())
		if err == nil {
			Expect(response.StatusCode).To(Equal(202))
		} else {
			Expect(err.Error()).To(ContainSubstring("The account to be imported is in INACTIVE state"))
			Expect(response.StatusCode).To(Equal(400))
		}
	})

	It("Successfully Create Account", func() {
		shouldSkipTest()

		options := service.NewCreateAccountOptions(parent, "IBM", "IBMid-550006JKXX")

		result, response, err := service.CreateAccount(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())

		Expect(response.StatusCode).To(Equal(202))
		newAccount = *result.AccountID
	})

	It("Successfully Get Account", func() {
		shouldSkipTest()

		options := service.NewGetAccountOptions(newAccount)

		result, response, err := service.GetAccount(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(response.StatusCode).To(Equal(200))
	})

	It("Successfully List Accounts", func() {
		shouldSkipTest()

		options := service.NewListAccountsOptions()
		options.SetEnterpriseID(enterprise_id)
		options.SetAccountGroupID(accountGroupID)
		options.SetParent(parent)
		options.SetLimit(100)

		result, response, err := service.ListAccounts(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(response.StatusCode).To(Equal(200))
	})

	It("Successfully List Account groups", func() {
		shouldSkipTest()

		options := service.NewListAccountGroupsOptions()
		options.SetEnterpriseID(enterprise_id)
		options.SetParentAccountGroupID(accountGroupID2)
		options.SetParent(parent)
		options.SetLimit(100)

		result, detailedResponse, err := service.ListAccountGroups(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))
	})

	It("Successfully Move Account within an Enterprise", func() {
		shouldSkipTest()

		options := service.NewUpdateAccountOptions(newAccount, crn)

		result, err := service.UpdateAccount(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(result.StatusCode).To(Equal(202))
	})
})
