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
package openservicebrokerv1_test

import (
	"github.com/IBM/platform-services-go-sdk/openservicebrokerv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"

	"fmt"
	"os"
)

const externalConfigFile = "../open_service_broker.env"

var (
	service      *openservicebrokerv1.OpenServiceBrokerV1
	err          error
	configLoaded bool = false

	testAccountId         string = "bc2b2fca0af84354a916dc1de6eee42e"
	testResourceGroupGuid string = "13aa3ee48c3b44ddb64c05c79f7ab8ef"
	testOrgGuid           string = "d35d4f0e-5076-4c89-9361-2522894b6548"
	testSpaceGuid         string = "336ba5f3-f185-488e-ac8d-02195eebb2f3"
	testAppGuid           string = "bf692181-1f0e-46be-9faf-eb0857f4d1d5"
	testPlanId1           string = "a10e4820-3685-11e9-b210-d663bd873d93"
	testPlanId2           string = "a10e4410-3685-11e9-b210-d663bd873d933"
	testInstanceId        string = "crn:v1:staging:public:bss-monitor:global:a/bc2b2fca0af84354a916dc1de6eee42e:sdkTestInstance::"
	testBindingId         string = "crn:v1:staging:public:bss-monitor:us-south:a/bc2b2fca0af84354a916dc1de6eee42e:sdkTestInstance:resource-binding:sdkTestBinding"
	testInstanceIdEscaped string = "crn%3Av1%3Astaging%3Apublic%3Abss-monitor%3Aglobal%3Aa%2Fbc2b2fca0af84354a916dc1de6eee42e%3AsdkTestInstance%3A%3A"
	testBindingIdEscaped  string = "crn%3Av1%3Astaging%3Apublic%3Abss-monitor%3Aus-south%3Aa%2Fbc2b2fca0af84354a916dc1de6eee42e%3AsdkTestInstance%3Aresource-binding%3AsdkTestBinding"
	testDashboardUrl      string = "http://www.example.com/crn%3Av1%3Astaging%3Apublic%3Abss-monitor%3Aglobal%3Aa%2Fbc2b2fca0af84354a916dc1de6eee42e%3AsdkTestInstance%3A%3A"
	testServiceId         string = "a10e46ae-3685-11e9-b210-d663bd873d93"
	testReasonCode        string = "test_reason"
	testInitiatorId       string = "test_initiator"
	transactionId         string = uuid.NewV4().String()
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe("Open Service Broker - Integration Tests", func() {

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
	})

	It(`Successfully created OpenServiceBrokerV1 service instances`, func() {
		shouldSkipTest()

		service, err = openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(
			&openservicebrokerv1.OpenServiceBrokerV1Options{},
		)

		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())

		fmt.Printf("\nTransaction Id for Test Run: %s\n", transactionId)
	})

	It("00 - Create Service Instance", func() {
		shouldSkipTest()

		platform := "ibmcloud"
		contextOpt := &openservicebrokerv1.Context{
			AccountID: &testAccountId,
			Crn:       &testInstanceId,
			Platform:  &platform,
		}

		paramsOpt := make(map[string]string, 0)
		paramsOpt["hello"] = "bye"

		options := service.NewReplaceServiceInstanceOptions(testInstanceIdEscaped)
		options = options.SetPlanID(testPlanId1)
		options = options.SetServiceID(testServiceId)
		options = options.SetOrganizationGuid(testOrgGuid)
		options = options.SetSpaceGuid(testSpaceGuid)
		options = options.SetContext(contextOpt)
		options = options.SetParameters(paramsOpt)
		options = options.SetAcceptsIncomplete(true)

		headers := map[string]string{
			"Transaction-Id": "osb-sdk-go-test00-" + transactionId,
		}
		options = options.SetHeaders(headers)
		result, resp, err := service.ReplaceServiceInstance(options)

		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(201))
		Expect(result).NotTo(BeNil())
		Expect(*result.DashboardURL).To(Equal(testDashboardUrl))
	})

	It("01 - Update Service Instance", func() {
		shouldSkipTest()

		platform := "cf"
		contextOpt := &openservicebrokerv1.Context{
			AccountID: &testAccountId,
			Crn:       &testInstanceId,
			Platform:  &platform,
		}

		paramsOpt := make(map[string]string, 0)
		paramsOpt["hello"] = "hi"

		previousValues := make(map[string]string, 0)
		previousValues["plan_id"] = testPlanId1

		options := service.NewUpdateServiceInstanceOptions(testInstanceIdEscaped)
		options = options.SetPlanID(testPlanId2)
		options = options.SetServiceID(testServiceId)
		options = options.SetContext(contextOpt)
		options = options.SetParameters(paramsOpt)
		options = options.SetAcceptsIncomplete(true)

		headers := map[string]string{
			"Transaction-Id": "osb-sdk-go-test01-" + transactionId,
		}
		options = options.SetHeaders(headers)
		result, resp, err := service.UpdateServiceInstance(options)

		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))
		Expect(result).NotTo(BeNil())
	})

	It("02 - Disable Service Instance State", func() {
		shouldSkipTest()

		options := service.NewReplaceServiceInstanceStateOptions(testInstanceIdEscaped)
		options = options.SetEnabled(false)
		options = options.SetInitiatorID(testInitiatorId)

		headers := map[string]string{
			"Transaction-Id": "osb-sdk-go-test02-" + transactionId,
		}
		options = options.SetHeaders(headers)
		result, resp, err := service.ReplaceServiceInstanceState(options)

		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))
		Expect(result).NotTo(BeNil())
	})

	It("03 - Enable Service Instance State", func() {
		shouldSkipTest()

		options := service.NewReplaceServiceInstanceStateOptions(testInstanceIdEscaped)
		options = options.SetEnabled(true)
		options = options.SetInitiatorID(testInitiatorId)

		headers := map[string]string{
			"Transaction-Id": "osb-sdk-go-test03-" + transactionId,
		}
		options = options.SetHeaders(headers)
		result, resp, err := service.ReplaceServiceInstanceState(options)

		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))
		Expect(result).NotTo(BeNil())
	})

	It("04 - Bind Service Instance", func() {
		shouldSkipTest()

		paramsOpt := make(map[string]string, 0)
		paramsOpt["hello"] = "bye"

		bindResource := &openservicebrokerv1.BindResource{
			AccountID:    &testAccountId,
			ServiceidCrn: &testAppGuid,
		}

		options := service.NewReplaceServiceBindingOptions(testBindingIdEscaped, testInstanceIdEscaped)
		options = options.SetPlanID(testPlanId2)
		options = options.SetServiceID(testServiceId)
		options = options.SetParameters(paramsOpt)
		options = options.SetBindResource(bindResource)

		headers := map[string]string{
			"Transaction-Id": "osb-sdk-go-test04-" + transactionId,
		}
		options = options.SetHeaders(headers)
		result, resp, err := service.ReplaceServiceBinding(options)

		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(201))
		Expect(result).NotTo(BeNil())
		Expect(result.Credentials).NotTo(BeNil())
		Expect(result.Credentials.(map[string]interface{})["credField"]).To(Equal("credValue"))
	})

	It("05 - Get Service Instance State", func() {
		shouldSkipTest()

		options := service.NewGetServiceInstanceStateOptions(testInstanceIdEscaped)

		headers := map[string]string{
			"Transaction-Id": "osb-sdk-go-test05-" + transactionId,
		}
		options = options.SetHeaders(headers)
		result, resp, err := service.GetServiceInstanceState(options)

		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))
		Expect(result).NotTo(BeNil())
	})

	It("06 - Get Catalog Metadata", func() {
		shouldSkipTest()

		options := service.NewListCatalogOptions()

		headers := map[string]string{
			"Transaction-Id": "osb-sdk-go-test06-" + transactionId,
		}
		options = options.SetHeaders(headers)
		result, resp, err := service.ListCatalog(options)

		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))
		Expect(result).NotTo(BeNil())
		Expect(*result.Services[0].ID).To(Equal(testServiceId))
		Expect(*result.Services[0].Name).To(Equal("bss-monitor"))
		Expect(*result.Services[0].Bindable).Should(BeTrue())
		Expect(*result.Services[0].PlanUpdateable).Should(BeTrue())

		foundPlan1 := false
		foundPlan2 := false
		for _, plan := range result.Services[0].Plans {
			if *plan.ID == testPlanId1 {
				foundPlan1 = true
			} else if *plan.ID == testPlanId2 {
				foundPlan2 = true
			}
		}
		Expect(foundPlan1 && foundPlan2).Should(BeTrue())
	})

	It("07 - Delete Service Binding", func() {
		shouldSkipTest()

		options := service.NewDeleteServiceBindingOptions(testBindingIdEscaped, testInstanceIdEscaped, testPlanId1, testServiceId)

		headers := map[string]string{
			"Transaction-Id": "osb-sdk-go-test07-" + transactionId,
		}
		options = options.SetHeaders(headers)
		resp, err := service.DeleteServiceBinding(options)

		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))
	})

	It("08 - Delete Service Instance", func() {
		shouldSkipTest()

		options := service.NewDeleteServiceInstanceOptions(testServiceId, testPlanId1, testInstanceIdEscaped)

		headers := map[string]string{
			"Transaction-Id": "osb-sdk-go-test08-" + transactionId,
		}
		options = options.SetHeaders(headers)
		result, resp, err := service.DeleteServiceInstance(options)

		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))
		Expect(result).NotTo(BeNil())
	})

})
