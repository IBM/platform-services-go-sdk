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
 
	 "os"
 )
 
 const externalConfigFile = "../openservicebroker_config.env"
 
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
	 testServiceId         string = "a10e46ae-3685-11e9-b210-d663bd873d93" 
	 testEnable 			  bool = true
	 testReasonCode		  string = "test_reason"
	 testInitiatorId      string = "test_initiator"
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
	 })
 
	 It("Create Service Instance", func() {
		 shouldSkipTest()
 
		 options := service.NewReplaceServiceInstanceOptions(
			 testInstanceId
		 )
		 options = options.SetPlanId(testPlanId1)
		 options = options.SetServiceID(testServiceId)
 
		 headers := map[string]string{
			 "Transaction-Id": "osb-sdk-go-test00-" + uuid.NewV4().String(),
		 }
		 options = options.SetHeaders(headers)
		 result, resp, err := service.ReplaceServiceInstance(options)
 
		 Expect(err).To(BeNil())
		 Expect(resp.StatusCode).To(Equal(201))
		 Expect(result.DashboardURL).NotTo(BeNil())
	 })
 
	 It("Update Service Instance", func() {
		 shouldSkipTest()
 
		 options := service.NewUpdateServiceInstanceOptions(
			 testInstanceId
		 )
		 options = options.SetPlanID(testPlanId1)
		 options = options.SetServiceID(testServiceId)
 
		 headers := map[string]string{
			 "Transaction-Id": "osb-sdk-go-test01-" + uuid.NewV4().String(),
		 }
		 options = options.SetHeaders(headers)
		 result, resp, err := service.UpdateServiceInstance(options)
 
		 Expect(err).To(BeNil())
		 Expect(resp.StatusCode).To(Equal(200))
		 Expect(result).NotTo(BeZero()) //check this, need to check elements of the response or just string?
	 })
 
	 It("Update Service Instance State", func() {
		 shouldSkipTest()
 
		 options := service.NewReplaceStateOptions(
			 testInstanceId
		 )
		 options = options.SetEnabled(testEnable)
		 options = options.SetInitiatorID(testInitiatorId)
 
		 headers := map[string]string{
			 "Transaction-Id": "osb-sdk-go-test02-" + uuid.NewV4().String(),
		 }
		 options = options.SetHeaders(headers)
		 result, resp, err := service.ReplaceState(options)
 
		 Expect(err).To(BeNil())
		 Expect(resp.StatusCode).To(Equal(200))
	 })
 
	 It("Bind Service Instance", func() {
		 shouldSkipTest()
 
		 options := service.NewReplaceServiceBindingOptions(
			 testBindingId, testInstanceId
		 )
 
		 headers := map[string]string{
			 "Transaction-Id": "osb-sdk-go-test03-" + uuid.NewV4().String(),
		 }
		 options = options.SetHeaders(headers)
		 result, resp, err := service.ReplaceServiceBinding(options)
 
		 Expect(err).To(BeNil())
		 Expect(resp.StatusCode).To(Equal(201)) //check, response comes back with 201 but api says 200
		 Expect(result).NotTo(BeZero()) //check, same as with update service, result type is just a string? should they be actual values from a struct?
	 })
 
	 It("Get Service Instance State", func() {
		 shouldSkipTest()
 
		 options := service.NewGetServiceInstanceStateOptions(
			 testInstanceId
		 )
 
		 headers := map[string]string{
			 "Transaction-Id": "osb-sdk-go-test04-" + uuid.NewV4().String(),
		 }
		 options = options.SetHeaders(headers)
		 result, resp, err := service.GetServiceInstanceState(options)
 
		 Expect(err).To(BeNil())
		 Expect(resp.StatusCode).To(Equal(200))
		 Expect(result).NotTo(BeNil()) //check this, actual response doesnt match response in api?
	 })
 
	 It("Get Catalog Metadata", func() {
		 shouldSkipTest()
 
		 options := service.NewGetServiceInstanceStateOptions()
 
		 headers := map[string]string{
			 "Transaction-Id": "osb-sdk-go-test05-" + uuid.NewV4().String(),
		 }
		 options = options.SetHeaders(headers)
		 result, resp, err := service.ListCatalog(options)
 
		 Expect(err).To(BeNil())
		 Expect(resp.StatusCode).To(Equal(200))
		 Expect(result).NotTo(BeNil()) 
	 })
 
	 It("Delete Service Binding", func() {
		 shouldSkipTest()
 
		 options := service.NewDeleteServiceBindingOptions(
			 testBindingId, testInstanceId, testPlanId1, testServiceId
		 )
 
		 headers := map[string]string{
			 "Transaction-Id": "osb-sdk-go-test06-" + uuid.NewV4().String(),
		 }
		 options = options.SetHeaders(headers)
		 result, resp, err := service.DeleteServiceBinding(options)
 
		 Expect(err).To(BeNil())
		 Expect(resp.StatusCode).To(Equal(200)) 
	 })
 
	 It("Delete Service Instance", func() {
		 shouldSkipTest()
 
		 options := service.NewDeleteServiceInstanceOptions(
			 testServiceId, testPlanId1, testInstanceId
		 )
 
		 headers := map[string]string{
			 "Transaction-Id": "osb-sdk-go-test07-" + uuid.NewV4().String(),
		 }
		 options = options.SetHeaders(headers)
		 result, resp, err := service.DeleteServiceInstance(options)
 
		 Expect(err).To(BeNil())
		 Expect(resp.StatusCode).To(Equal(200)) 
	 })
 
 })
 