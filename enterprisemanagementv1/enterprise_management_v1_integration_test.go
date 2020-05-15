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
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"os"
	"strconv"
	"time"
	//"github.com/ibm-developer/ibm-cloud-env-golang"
	//"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
	"github.com/joho/godotenv"
	//"strings"
	"bytes"
	iam "github.ibm.com/BSS/golang-iam"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const externalConfigFile = "../enterprise-management.env"

var (
	//err     error
	configLoaded          bool = false
	service1              *enterprisemanagementv1.EnterpriseManagementV1
	email                 string = "aminttest+" + strconv.Itoa(rand.Intn(100000)) + "_" + strconv.Itoa(rand.Intn(100000)) + "@mail.test.ibm.com"
	managementToken       string
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

		err = godotenv.Load(externalConfigFile)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		options1 := &enterprisemanagementv1.EnterpriseManagementV1Options{
			ServiceName: "enterprise_management",
			URL:         os.Getenv("ENTERPRISE_MANAGEMENT_URL"),
		}
		service1, err = enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(options1)
		Expect(err).To(BeNil())
		Expect(service1).ToNot(BeNil())

	})

	It("Successfully created Service token", func() {

		tc := iam.TokenCache{
			URI:    os.Getenv("EMTEST_CONFIG_IAM_HOST"),
			ApiKey: os.Getenv("EMTEST_CONFIG_IAM_API_KEY"),
		}
		err = tc.Start()
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		managementToken = tc.GetToken()

		time.Sleep(3000 * time.Millisecond)

	})

	It("Creates a standard account", func() {

		apiUrl := os.Getenv("EMTEST_CONFIG_AM_HOST")
		resource := "/coe/v2/accounts"

		u, _ := url.ParseRequestURI(apiUrl)
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

		accountPayloadJson, _ := json.Marshal(accountPayload) //convert the struct to JSON format

		url := urlStr

		// Create a new request using http
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(accountPayloadJson))
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		//

		// add authorization header to the req
		req.Header.Add("Authorization", managementToken)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		accountId := data["id"].(interface{})

		account_id = accountId.(string)
	})

	It("Successfully get activation code", func() {

		time.Sleep(20000 * time.Millisecond)

		apiUrl := os.Getenv("EMTEST_CONFIG_AM_HOST")
		resource := "/v1/activation-codes/" + email

		u, _ := url.ParseRequestURI(apiUrl)
		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		req.Header.Add("Authorization", managementToken)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		var results map[string]interface{} // TopTracks

		err = json.Unmarshal(body, &results)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		res := results["resources"].([]interface{})

		z := res[0].(map[string]interface{})
		activationId = z["id"].(string)

	})

	It("Successfully activate account", func() {

		apiUrl := os.Getenv("EMTEST_CONFIG_AM_HOST")
		resource := "/coe/v2/accounts/verify"

		u, _ := url.ParseRequestURI(apiUrl)
		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		q := req.URL.Query()
		q.Add("token", activationId)
		q.Add("email", email)
		req.URL.RawQuery = q.Encode()
		// Send req using http Client
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

	})

	It("Successfully get account", func() {

		apiUrl := os.Getenv("EMTEST_CONFIG_AM_HOST")
		resource := "/coe/v2/accounts/" + account_id
		u, _ := url.ParseRequestURI(apiUrl)
		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		req.Header.Add("Authorization", managementToken)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		res := data["entity"].(map[string]interface{})
		owner_iam_id = res["owner_iam_id"].(string)
		subscription_id = res["subscription_id"].(string)

		//["owner_iam_id"]
		//["subscription_id"]
	})

	It("Successfully convert account from STANDARD to SUBSCRIPTION", func() {

		apiUrl := os.Getenv("EMTEST_CONFIG_AM_HOST")
		resource := "/coe/v2/accounts/" + account_id + "/bluemix_subscriptions/" + subscription_id
		u, _ := url.ParseRequestURI(apiUrl)
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
				End_date:   "2020-11-30T08:00:00.000",
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
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		req.Header.Add("Authorization", managementToken)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		ioutil.ReadAll(resp.Body)

	})

	It("Successfully create enterprise account", func() {
		options := service1.NewCreateEnterpriseOptions(account_id, "IBM", owner_iam_id)

		result, detailedResponse, err := service1.CreateEnterprise(options)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		Expect(detailedResponse.StatusCode).To(Equal(202))

		enterprise_id = *result.EnterpriseID
		enterprise_account_id = *result.EnterpriseAccountID
		Expect(err).To(BeNil())

	})

	It("Successfully get account", func() {

		apiUrl := os.Getenv("EMTEST_CONFIG_AM_HOST")
		resource := "/coe/v2/accounts/" + account_id
		u, _ := url.ParseRequestURI(apiUrl)
		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		req.Header.Add("Authorization", managementToken)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		res := data["entity"].(map[string]interface{})
		parent = res["parent"].(string)

		//["owner_iam_id"]
		//["subscription_id"]
	})

	It("Successfully Create Account group", func() {

		options := service1.NewCreateAccountGroupOptions(parent, "IBM", owner_iam_id)

		result, detailedResponse, err := service1.CreateAccountGroup(options)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		Expect(detailedResponse.StatusCode).To(Equal(201))

		accountGroupID = *result.AccountGroupID
	})

	It("Successfully Create Account group", func() {

		options := service1.NewCreateAccountGroupOptions(parent, "IBM", owner_iam_id)

		result, detailedResponse, err := service1.CreateAccountGroup(options)

		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		Expect(detailedResponse.StatusCode).To(Equal(201))

		accountGroupID2 = *result.AccountGroupID
	})

	It("Successfully List Account groups", func() {
		options := service1.NewListAccountGroupsOptions()
		options.SetEnterpriseID(enterprise_id)

		options.SetParentAccountGroupID(accountGroupID)

		options.SetParent(parent)

		options.SetLimit(100)

		result, detailedResponse, err := service1.ListAccountGroups(options)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		Expect(result).NotTo(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))
	})

	It("Successfully Get Account group", func() {

		options := service1.NewGetAccountGroupByIdOptions(accountGroupID)
		result, detailedResponse, err := service1.GetAccountGroupByID(options)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		Expect(detailedResponse.StatusCode).To(Equal(200))

		crn = *result.Crn
	})

	It("Successfully Update Account group", func() {
		options := service1.NewUpdateAccountGroupOptions(accountGroupID)
		options.SetName("IBM")
		options.SetPrimaryContactIamID(owner_iam_id)
		result, err := service1.UpdateAccountGroup(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(result.StatusCode).To(Equal(204))
	})

	It("Successfully Get Account Permissible Actions", func() {

		actions := []string{"teststring"}
		options := service1.NewGetAccountPermissibleActionsOptions(enterprise_account_id)
		options.SetActions(actions)
		result, err := service1.GetAccountPermissibleActions(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(result.StatusCode).To(Equal(200))
	})

	It("Creates a standard account", func() {

		apiUrl := os.Getenv("EMTEST_CONFIG_AM_HOST")
		resource := "/coe/v2/accounts"

		u, _ := url.ParseRequestURI(apiUrl)
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
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		//

		// add authorization header to the req
		req.Header.Add("Authorization", managementToken)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		accountId := data["id"].(interface{})

		standard_account_id = accountId.(string)
	})

	It("Successfully get activation code", func() {

		time.Sleep(20000 * time.Millisecond)

		apiUrl := os.Getenv("EMTEST_CONFIG_AM_HOST")
		resource := "/v1/activation-codes/" + email2

		u, _ := url.ParseRequestURI(apiUrl)
		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		req.Header.Add("Authorization", managementToken)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		var results map[string]interface{} // TopTracks

		err = json.Unmarshal(body, &results)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		res := results["resources"].([]interface{})

		z := res[0].(map[string]interface{})
		activationId = z["id"].(string)

	})

	It("Successfully activate account", func() {

		apiUrl := os.Getenv("EMTEST_CONFIG_AM_HOST")
		resource := "/coe/v2/accounts/verify"

		u, _ := url.ParseRequestURI(apiUrl)
		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		q := req.URL.Query()
		q.Add("token", activationId)
		q.Add("email", email2)
		req.URL.RawQuery = q.Encode()
		// Send req using http Client
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

	})

	It("Successfully get account", func() {

		apiUrl := os.Getenv("EMTEST_CONFIG_AM_HOST")
		resource := "/coe/v2/accounts/" + standard_account_id
		u, _ := url.ParseRequestURI(apiUrl)
		u.Path = resource
		urlStr := u.String()

		url := urlStr

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		req.Header.Add("Authorization", managementToken)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{} // TopTracks
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		res := data["entity"].(map[string]interface{})
		owner_iam_id = res["owner_iam_id"].(string)
		subscription_id = res["subscription_id"].(string)

		//["owner_iam_id"]
		//["subscription_id"]
	})

	It("Successfully Import Account to Enterprise", func() {
		options := service1.NewImportAccountToEnterpriseOptions(enterprise_id, standard_account_id)
		options.SetParent(parent)
		result, err := service1.ImportAccountToEnterprise(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(result.StatusCode).To(Equal(202))
	})

	It("Successfully Create Account", func() {
		options := service1.NewCreateAccountOptions(parent, "IBM", "IBMid-550006JKXX")
		result, response, err := service1.CreateAccount(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())

		Expect(response.StatusCode).To(Equal(202))
		newAccount = *result.AccountID
	})

	It("Successfully Get Account", func() {
		options := service1.NewGetAccountByIdOptions(newAccount)
		result, response, err := service1.GetAccountByID(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(response.StatusCode).To(Equal(200))
	})

	It("Successfully List Account", func() {
		options := service1.NewListAccountsOptions()
		options.SetEnterpriseID(enterprise_id)
		options.SetAccountGroupID(accountGroupID)
		options.SetParent(parent)
		options.SetLimit(100)
		result, response, err := service1.ListAccounts(options)
		Expect(err).To(BeNil())
		Expect(result).NotTo(BeNil())
		Expect(response.StatusCode).To(Equal(200))
	})

	It("Successfully List Account groups", func() {
		options := service1.NewListAccountGroupsOptions()
		options.SetEnterpriseID(enterprise_id)

		options.SetParentAccountGroupID(accountGroupID2)

		options.SetParent(parent)

		options.SetLimit(100)

		result, detailedResponse, err := service1.ListAccountGroups(options)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		Expect(result).NotTo(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))
	})

	It("Successfully Move Account within an Enterprise", func() {
		options := service1.NewUpdateAccountOptions(newAccount, crn)

		result, err := service1.UpdateAccount(options)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		Expect(result).NotTo(BeNil())
		Expect(result.StatusCode).To(Equal(202))
	})

	It("Successfully Get Account Permissible Actions", func() {
		options := service1.NewGetAccountPermissibleActionsOptions(account_id)
		actions := []string{"teststring"}
		options.SetActions(actions)

		result, err := service1.GetAccountPermissibleActions(options)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		Expect(result).NotTo(BeNil())
		Expect(result.StatusCode).To(Equal(200))
	})

})
