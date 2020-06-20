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
package resourcecontrollerv2_test

import (
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/satori/go.uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"os"
	"time"
)

const externalConfigFile = "../resource_controller.env"

var (
	service      *resourcecontrollerv2.ResourceControllerV2
	err          error
	configLoaded bool = false

	testAccountId         string = "bc2b2fca0af84354a916dc1de6eee42e"
	testResourceGroupGuid string = "13aa3ee48c3b44ddb64c05c79f7ab8ef"
	testOrgGuid           string = "d35d4f0e-5076-4c89-9361-2522894b6548"
	testSpaceGuid         string = "336ba5f3-f185-488e-ac8d-02195eebb2f3"
	testAppGuid           string = "bf692181-1f0e-46be-9faf-eb0857f4d1d5"
	testRegionId1         string = "global"
	testPlanId1           string = "a10e4820-3685-11e9-b210-d663bd873d93"
	testRegionId2         string = "us-south"
	testPlanId2           string = "2580b607-db64-4883-9793-445b694ed57b"

	//result info
	testInstanceCrn         string
	testInstanceGuid        string
	testAliasCrn            string
	testAliasGuid           string
	testBindingCrn          string
	testBindingGuid         string
	testInstanceKeyCrn      string
	testInstanceKeyGuid     string
	testAliasKeyCrn         string
	testAliasKeyGuid        string
	aliasTargetCrn          string
	bindTargetCrn           string
	testReclaimInstanceCrn  string
	testReclaimInstanceGuid string
	testReclamationId1      string
	testReclamationId2      string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe("Resource Controller - Integration Tests", func() {

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

	It(`Successfully created ResourceControllerV2 service instances`, func() {
		shouldSkipTest()

		service, err = resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(
			&resourcecontrollerv2.ResourceControllerV2Options{},
		)

		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())
	})

	Describe("00 - Create, Retrieve, and Update Resource Instance", func() {
		It("Create Resource Instance", func() {
			shouldSkipTest()

			options := service.NewCreateResourceInstanceOptions(
				"RcSdkInstance1",
				testRegionId1,
				testResourceGroupGuid,
				testPlanId1,
			)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test00-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.CreateResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result.ID).NotTo(BeNil())
			Expect(result.Guid).NotTo(BeNil())
			Expect(result.Crn).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.Crn))
			Expect(*result.Name).To(Equal("RcSdkInstance1"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.ResourcePlanID).To(Equal(testPlanId1))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.Locked).Should(BeFalse())
			Expect(result.LastOperation["type"]).To(Equal("create"))
			Expect(result.LastOperation["async"]).Should(BeFalse())
			Expect(result.LastOperation["state"]).To(Equal("succeeded"))

			testInstanceCrn = *result.ID
			testInstanceGuid = *result.Guid
		})

		It("01 - Get A Resource Instance", func() {
			shouldSkipTest()

			options := service.NewGetResourceInstanceOptions(testInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test01-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testInstanceCrn))
			Expect(*result.Guid).To(Equal(testInstanceGuid))
			Expect(*result.Crn).To(Equal(testInstanceCrn))
			Expect(*result.Name).To(Equal("RcSdkInstance1"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.ResourcePlanID).To(Equal(testPlanId1))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.Locked).Should(BeFalse())
			Expect(result.LastOperation["type"]).To(Equal("create"))
			Expect(result.LastOperation["async"]).Should(BeFalse())
			Expect(result.LastOperation["state"]).To(Equal("succeeded"))
		})

		It("02 - Update A Resource Instance", func() {
			shouldSkipTest()

			options := service.NewUpdateResourceInstanceOptions(testInstanceGuid)
			options = options.SetName("RcSdkInstanceUpdate1")
			params := make(map[string]interface{}, 0)
			params["hello"] = "bye"
			options = options.SetParameters(params)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test02-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.UpdateResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testInstanceCrn))
			Expect(*result.Name).To(Equal("RcSdkInstanceUpdate1"))
			Expect(*result.State).To(Equal("active"))
			Expect(result.LastOperation["type"]).To(Equal("update"))
			Expect(result.LastOperation["sub_type"]).To(Equal("config"))
			Expect(result.LastOperation["async"]).Should(BeFalse())
			Expect(result.LastOperation["state"]).To(Equal("succeeded"))
		})

		Describe("List Resource Instances", func() {
			It("03 - List Resource Instances With No Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceInstancesOptions()
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test03-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceInstances(options)

				//should return one or more instances
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).Should(BeNumerically(">=", int64(1)))
				Expect(len(result.Resources)).Should(BeNumerically(">=", 1))
			})

			It("04 - List Resource Instances With Guid Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceInstancesOptions()
				options = options.SetGuid(testInstanceGuid)
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test04-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceInstances(options)

				//should return list with only newly created instance
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
				Expect(*result.Resources[0].ID).To(Equal(testInstanceCrn))
				Expect(*result.Resources[0].Guid).To(Equal(testInstanceGuid))
				Expect(*result.Resources[0].Name).To(Equal("RcSdkInstanceUpdate1"))
				Expect(*result.Resources[0].State).To(Equal("active"))
				Expect(result.Resources[0].LastOperation["type"]).To(Equal("update"))
				Expect(result.Resources[0].LastOperation["sub_type"]).To(Equal("config"))
				Expect(result.Resources[0].LastOperation["async"]).Should(BeFalse())
				Expect(result.Resources[0].LastOperation["state"]).To(Equal("succeeded"))
			})

			It("05 - List Resource Instances With Name Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceInstancesOptions()
				options = options.SetName("RcSdkInstance1")
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test05-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceInstances(options)

				//name was updated so no instance with that name should exist
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).To(Equal(int64(0)))
				Expect(result.Resources).Should(HaveLen(0))
			})
		})
	})

	Describe("Create, Retrieve, and Update Resource Alias", func() {
		It("06 - Create Resource Alias", func() {
			shouldSkipTest()

			target := "crn:v1:bluemix:public:bluemix:us-south:o/" + testOrgGuid + "::cf-space:" + testSpaceGuid
			aliasTargetCrn = "crn:v1:bluemix:public:cf:us-south:o/" + testOrgGuid + "::cf-space:" + testSpaceGuid
			options := service.NewCreateResourceAliasOptions("RcSdkAlias1", testInstanceGuid, target)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test06-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.CreateResourceAlias(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result.ID).NotTo(BeNil())
			Expect(result.Guid).NotTo(BeNil())
			Expect(result.Crn).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.Crn))
			Expect(*result.Name).To(Equal("RcSdkAlias1"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.TargetCrn).To(Equal(aliasTargetCrn))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.ResourceInstanceID).To(Equal(testInstanceCrn))

			testAliasCrn = *result.ID
			testAliasGuid = *result.Guid
		})

		It("07 - Get A Resource Alias", func() {
			shouldSkipTest()

			options := service.NewGetResourceAliasOptions(testAliasGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test07-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceAlias(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testAliasCrn))
			Expect(*result.Guid).To(Equal(testAliasGuid))
			Expect(*result.Crn).To(Equal(testAliasCrn))
			Expect(*result.Name).To(Equal("RcSdkAlias1"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.TargetCrn).To(Equal(aliasTargetCrn))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.ResourceInstanceID).To(Equal(testInstanceCrn))
		})

		It("08 - Update A Resource Alias", func() {
			shouldSkipTest()

			options := service.NewUpdateResourceAliasOptions(testAliasGuid, "RcSdkAliasUpdate1")
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test08-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.UpdateResourceAlias(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testAliasCrn))
			Expect(*result.Name).To(Equal("RcSdkAliasUpdate1"))
			Expect(*result.State).To(Equal("active"))
		})

		Describe("List Resource Aliases", func() {
			It("09 - List Resource Aliases With No Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceAliasesOptions()
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test09-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceAliases(options)

				//should return one or more aliases
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).Should(BeNumerically(">=", int64(1)))
				Expect(len(result.Resources)).Should(BeNumerically(">=", 1))
			})

			It("10 - List Resource Aliases With Guid Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceAliasesOptions()
				options = options.SetGuid(testAliasGuid)
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test10-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceAliases(options)

				//should return list with only newly created alias
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
				Expect(*result.Resources[0].ID).To(Equal(testAliasCrn))
				Expect(*result.Resources[0].Name).To(Equal("RcSdkAliasUpdate1"))
				Expect(*result.Resources[0].ResourceGroupID).To(Equal(testResourceGroupGuid))
				Expect(*result.Resources[0].TargetCrn).To(Equal(aliasTargetCrn))
				Expect(*result.Resources[0].State).To(Equal("active"))
				Expect(*result.Resources[0].ResourceInstanceID).To(Equal(testInstanceCrn))
			})

			It("11 - List Resource Aliases With Name Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceAliasesOptions()
				options = options.SetName("RcSdkAlias1")
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test11-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceAliases(options)

				//name was updated so no alias with that name should exist
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).To(Equal(int64(0)))
				Expect(result.Resources).Should(HaveLen(0))
			})
		})
	})

	Describe("Create, Retrieve, and Update Resource Binding", func() {
		It("12 - Create Resource Binding", func() {
			shouldSkipTest()

			target := "crn:v1:staging:public:bluemix:us-south:s/" + testSpaceGuid + "::cf-application:" + testAppGuid
			bindTargetCrn = "crn:v1:staging:public:cf:us-south:s/" + testSpaceGuid + "::cf-application:" + testAppGuid
			options := service.NewCreateResourceBindingOptions(testAliasGuid, target)
			options = options.SetName("RcSdkBinding1")
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test12-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.CreateResourceBinding(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result.ID).NotTo(BeNil())
			Expect(result.Guid).NotTo(BeNil())
			Expect(result.Crn).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.Crn))
			Expect(*result.Name).To(Equal("RcSdkBinding1"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.SourceCrn).To(Equal(testAliasCrn))
			Expect(*result.TargetCrn).To(Equal(bindTargetCrn))
			Expect(*result.State).To(Equal("active"))

			testBindingCrn = *result.ID
			testBindingGuid = *result.Guid
		})

		It("13 - Get A Resource Binding", func() {
			shouldSkipTest()

			options := service.NewGetResourceBindingOptions(testBindingGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test13-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceBinding(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testBindingCrn))
			Expect(*result.Guid).To(Equal(testBindingGuid))
			Expect(*result.Crn).To(Equal(testBindingCrn))
			Expect(*result.Name).To(Equal("RcSdkBinding1"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.SourceCrn).To(Equal(testAliasCrn))
			Expect(*result.TargetCrn).To(Equal(bindTargetCrn))
			Expect(*result.State).To(Equal("active"))
		})

		It("14 - Update A Resource Binding", func() {
			shouldSkipTest()

			options := service.NewUpdateResourceBindingOptions(testBindingGuid, "RcSdkBindingUpdate1")
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test14-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.UpdateResourceBinding(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testBindingCrn))
			Expect(*result.Name).To(Equal("RcSdkBindingUpdate1"))
			Expect(*result.State).To(Equal("active"))
		})

		Describe("List Resource Bindings", func() {
			It("15 - List Resource Bindings With No Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceBindingsOptions()
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test15-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceBindings(options)

				//should return one or more bindings
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).Should(BeNumerically(">=", int64(1)))
				Expect(len(result.Resources)).Should(BeNumerically(">=", 1))
			})

			It("16 - List Resource Bindings With Guid Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceBindingsOptions()
				options = options.SetGuid(testBindingGuid)
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test16-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceBindings(options)

				//should return list with only newly created binding
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
				Expect(*result.Resources[0].ID).To(Equal(testBindingCrn))
				Expect(*result.Resources[0].Name).To(Equal("RcSdkBindingUpdate1"))
				Expect(*result.Resources[0].ResourceGroupID).To(Equal(testResourceGroupGuid))
				Expect(*result.Resources[0].SourceCrn).To(Equal(testAliasCrn))
				Expect(*result.Resources[0].TargetCrn).To(Equal(bindTargetCrn))
				Expect(*result.Resources[0].State).To(Equal("active"))
			})

			It("17 - List Resource Bindings With Name Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceAliasesOptions()
				options = options.SetName("RcSdkBinding1")
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test17-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceAliases(options)

				//name was updated so no binding with that name should exist
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).To(Equal(int64(0)))
				Expect(result.Resources).Should(HaveLen(0))
			})
		})
	})

	Describe("Create, Retrieve, and Update Resource Key With Instance Source", func() {
		It("18 - Create Resource Key For Instance", func() {
			shouldSkipTest()

			options := service.NewCreateResourceKeyOptions("RcSdkKey1", testInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test18-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.CreateResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result.ID).NotTo(BeNil())
			Expect(result.Guid).NotTo(BeNil())
			Expect(result.Crn).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.Crn))
			Expect(*result.Name).To(Equal("RcSdkKey1"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.SourceCrn).To(Equal(testInstanceCrn))
			Expect(*result.State).To(Equal("active"))

			testInstanceKeyCrn = *result.ID
			testInstanceKeyGuid = *result.Guid
		})

		It("19 - Get A Resource Key", func() {
			shouldSkipTest()

			options := service.NewGetResourceKeyOptions(testInstanceKeyGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test19-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testInstanceKeyCrn))
			Expect(*result.Guid).To(Equal(testInstanceKeyGuid))
			Expect(*result.Crn).To(Equal(testInstanceKeyCrn))
			Expect(*result.Name).To(Equal("RcSdkKey1"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.SourceCrn).To(Equal(testInstanceCrn))
			Expect(*result.State).To(Equal("active"))
		})

		It("20 - Update A Resource Key", func() {
			shouldSkipTest()

			options := service.NewUpdateResourceKeyOptions(testInstanceKeyGuid, "RcSdkKeyUpdate1")
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test20-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.UpdateResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testInstanceKeyCrn))
			Expect(*result.Name).To(Equal("RcSdkKeyUpdate1"))
			Expect(*result.State).To(Equal("active"))
		})

		Describe("List Resource Keys", func() {
			It("21 - List Resource Keys With No Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceKeysOptions()
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test21-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceKeys(options)

				//should return one or more keys
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).Should(BeNumerically(">=", int64(1)))
				Expect(len(result.Resources)).Should(BeNumerically(">=", 1))
			})

			It("22 - List Resource Keys With Guid Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceKeysOptions()
				options = options.SetGuid(testInstanceKeyGuid)
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test22-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceKeys(options)

				//should return list with only newly created key
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
				Expect(*result.Resources[0].ID).To(Equal(testInstanceKeyCrn))
				Expect(*result.Resources[0].Name).To(Equal("RcSdkKeyUpdate1"))
				Expect(*result.Resources[0].ResourceGroupID).To(Equal(testResourceGroupGuid))
				Expect(*result.Resources[0].SourceCrn).To(Equal(testInstanceCrn))
				Expect(*result.Resources[0].State).To(Equal("active"))
			})

			It("23 - List Resource Keys With Name Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceKeysOptions()
				options = options.SetName("RcSdkKey1")
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test23-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceKeys(options)

				//name was updated so no key with that name should exist
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).To(Equal(int64(0)))
				Expect(result.Resources).Should(HaveLen(0))
			})
		})
	})

	Describe("Create, Retrieve, and Update Resource Key With Alias Source", func() {
		It("24 - Create Resource Key For Alias", func() {
			shouldSkipTest()

			options := service.NewCreateResourceKeyOptions("RcSdkKey2", testAliasGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test24-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.CreateResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result.ID).NotTo(BeNil())
			Expect(result.Guid).NotTo(BeNil())
			Expect(result.Crn).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.Crn))
			Expect(*result.Name).To(Equal("RcSdkKey2"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.SourceCrn).To(Equal(testAliasCrn))
			Expect(*result.State).To(Equal("active"))

			testAliasKeyCrn = *result.ID
			testAliasKeyGuid = *result.Guid
		})

		It("25 - Get A Resource Key", func() {
			shouldSkipTest()

			options := service.NewGetResourceKeyOptions(testAliasKeyGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test25-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testAliasKeyCrn))
			Expect(*result.Guid).To(Equal(testAliasKeyGuid))
			Expect(*result.Crn).To(Equal(testAliasKeyCrn))
			Expect(*result.Name).To(Equal("RcSdkKey2"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.SourceCrn).To(Equal(testAliasCrn))
			Expect(*result.State).To(Equal("active"))
		})

		It("26 - Update A Resource Key", func() {
			shouldSkipTest()

			options := service.NewUpdateResourceKeyOptions(testAliasKeyGuid, "RcSdkKeyUpdate2")
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test26-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.UpdateResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testAliasKeyCrn))
			Expect(*result.Name).To(Equal("RcSdkKeyUpdate2"))
			Expect(*result.State).To(Equal("active"))
		})

		Describe("List Resource Keys", func() {
			It("27 - List Resource Keys With No Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceKeysOptions()
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test27-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceKeys(options)

				//should return two or more keys
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).Should(BeNumerically(">=", int64(2)))
				Expect(len(result.Resources)).Should(BeNumerically(">=", 2))
			})

			It("28 - List Resource Keys With Guid Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceKeysOptions()
				options = options.SetGuid(testAliasKeyGuid)
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test28-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceKeys(options)

				//should return list with only newly created key
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
				Expect(*result.Resources[0].ID).To(Equal(testAliasKeyCrn))
				Expect(*result.Resources[0].Name).To(Equal("RcSdkKeyUpdate2"))
				Expect(*result.Resources[0].ResourceGroupID).To(Equal(testResourceGroupGuid))
				Expect(*result.Resources[0].SourceCrn).To(Equal(testAliasCrn))
				Expect(*result.Resources[0].State).To(Equal("active"))
			})

			It("29 - List Resource Keys With Name Filter", func() {
				shouldSkipTest()

				options := service.NewListResourceKeysOptions()
				options = options.SetName("RcSdkKey2")
				headers := map[string]string{
					"Transaction-Id": "rc-sdk-go-test29-" + uuid.NewV4().String(),
				}
				options = options.SetHeaders(headers)
				result, resp, err := service.ListResourceKeys(options)

				//name was updated so no key with that name should exist
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(*result.RowsCount).To(Equal(int64(0)))
				Expect(result.Resources).Should(HaveLen(0))
			})
		})
	})

	Describe("Delete All Resources", func() {
		It("30 - Delete A Resource Alias With Dependencies - Fail", func() {
			shouldSkipTest()

			options := service.NewDeleteResourceAliasOptions(testAliasGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test30-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			resp, err := service.DeleteResourceAlias(options)

			Expect(resp.StatusCode).To(Equal(400))
			Expect(err).NotTo(BeNil())
		})

		It("31 - Delete A Resource Instance With Dependencies - Fail", func() {
			shouldSkipTest()

			options := service.NewDeleteResourceInstanceOptions(testInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test31-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			resp, err := service.DeleteResourceInstance(options)

			Expect(resp.StatusCode).To(Equal(400))
			Expect(err).NotTo(BeNil())
		})

		It("32 - Delete A Resource Binding", func() {
			shouldSkipTest()

			options := service.NewDeleteResourceBindingOptions(testBindingGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test32-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			resp, err := service.DeleteResourceBinding(options)

			Expect(resp.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())
		})

		It("33 - Verify Resource Binding Was Deleted", func() {
			shouldSkipTest()

			options := service.NewGetResourceBindingOptions(testBindingGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test33-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceBinding(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testBindingCrn))
			Expect(*result.State).To(Equal("removed"))
		})

		It("34 - Delete Resource Keys", func() {
			shouldSkipTest()

			options1 := service.NewDeleteResourceKeyOptions(testInstanceKeyGuid)
			headers1 := map[string]string{
				"Transaction-Id": "rc-sdk-go-test34-" + uuid.NewV4().String(),
			}
			options1 = options1.SetHeaders(headers1)
			resp1, err1 := service.DeleteResourceKey(options1)

			Expect(resp1.StatusCode).To(Equal(204))
			Expect(err1).To(BeNil())

			options2 := service.NewDeleteResourceKeyOptions(testAliasKeyGuid)
			headers2 := map[string]string{
				"Transaction-Id": "rc-sdk-go-test34-" + uuid.NewV4().String(),
			}
			options2 = options2.SetHeaders(headers2)
			resp2, err2 := service.DeleteResourceKey(options2)

			Expect(resp2.StatusCode).To(Equal(204))
			Expect(err2).To(BeNil())
		})

		It("35 - Verify Resource Keys Were Deleted", func() {
			shouldSkipTest()

			options1 := service.NewGetResourceKeyOptions(testInstanceKeyGuid)
			headers1 := map[string]string{
				"Transaction-Id": "rc-sdk-go-test35-" + uuid.NewV4().String(),
			}
			options1 = options1.SetHeaders(headers1)
			result1, resp1, err1 := service.GetResourceKey(options1)

			Expect(err1).To(BeNil())
			Expect(resp1.StatusCode).To(Equal(200))
			Expect(*result1.ID).To(Equal(testInstanceKeyCrn))
			Expect(*result1.State).To(Equal("removed"))

			options2 := service.NewGetResourceKeyOptions(testAliasKeyGuid)
			headers2 := map[string]string{
				"Transaction-Id": "rc-sdk-go-test35-" + uuid.NewV4().String(),
			}
			options2 = options2.SetHeaders(headers2)
			result2, resp2, err2 := service.GetResourceKey(options2)

			Expect(err2).To(BeNil())
			Expect(resp2.StatusCode).To(Equal(200))
			Expect(*result2.ID).To(Equal(testAliasKeyCrn))
			Expect(*result2.State).To(Equal("removed"))
		})

		It("36 - Delete A Resource Alias", func() {
			shouldSkipTest()

			options := service.NewDeleteResourceAliasOptions(testAliasGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test36-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			resp, err := service.DeleteResourceAlias(options)

			Expect(resp.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())
		})

		It("37 - Verify Resource Alias Was Deleted", func() {
			shouldSkipTest()

			options := service.NewGetResourceAliasOptions(testAliasGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test37-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceAlias(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testAliasCrn))
			Expect(*result.State).To(Equal("removed"))
		})
	})

	Describe("Locking and Unlocking Resource Instance", func() {
		It("38 - Lock A Resource Instance", func() {
			shouldSkipTest()

			options := service.NewLockResourceInstanceOptions(testInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test38-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.LockResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testInstanceCrn))
			Expect(*result.Locked).To(BeTrue())
			Expect(result.LastOperation["type"]).To(Equal("lock"))
			Expect(result.LastOperation["async"]).Should(BeFalse())
			Expect(result.LastOperation["state"]).To(Equal("succeeded"))
		})

		It("39 - Update A Locked Resource Instance - Fail", func() {
			shouldSkipTest()

			options := service.NewUpdateResourceInstanceOptions(testInstanceGuid)
			options = options.SetName("RcSdkLockedInstanceUpdate1")
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test39-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			_, resp, err := service.UpdateResourceInstance(options)

			Expect(err).NotTo(BeNil())
			Expect(resp.StatusCode).To(Equal(400))
		})

		It("40 - Delete A Locked Resource Instance - Fail", func() {
			shouldSkipTest()

			options := service.NewDeleteResourceInstanceOptions(testInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test40-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			resp, err := service.DeleteResourceInstance(options)

			Expect(err).NotTo(BeNil())
			Expect(resp.StatusCode).To(Equal(400))
		})

		It("41 - Unlock A Resource Instance", func() {
			shouldSkipTest()

			options := service.NewUnlockResourceInstanceOptions(testInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test41-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.UnlockResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testInstanceCrn))
			Expect(*result.Locked).To(BeFalse())
			Expect(result.LastOperation["type"]).To(Equal("unlock"))
			Expect(result.LastOperation["async"]).Should(BeFalse())
			Expect(result.LastOperation["state"]).To(Equal("succeeded"))
		})
	})

	Describe("Delete Resource Instance", func() {
		It("42 - Delete A Resource Instance", func() {
			shouldSkipTest()

			options := service.NewDeleteResourceInstanceOptions(testInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test42-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			resp, err := service.DeleteResourceInstance(options)

			Expect(resp.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())
		})

		It("43 - Verify Resource Instance Was Deleted", func() {
			shouldSkipTest()

			options := service.NewGetResourceInstanceOptions(testInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test43-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testInstanceCrn))
			Expect(*result.State).To(Equal("removed"))
			Expect(result.LastOperation["type"]).To(Equal("delete"))
			Expect(result.LastOperation["async"]).Should(BeFalse())
			Expect(result.LastOperation["state"]).To(Equal("succeeded"))
		})
	})

	Describe("Resource Reclamation", func() {
		It("44 - Create Resource Instance For Reclamation Enabled Plan", func() {
			shouldSkipTest()

			options := service.NewCreateResourceInstanceOptions(
				"RcSdkReclaimInstance1",
				testRegionId2,
				testResourceGroupGuid,
				testPlanId2,
			)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test44-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.CreateResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result.ID).NotTo(BeNil())
			Expect(result.Guid).NotTo(BeNil())
			Expect(result.Crn).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.Crn))
			Expect(*result.Name).To(Equal("RcSdkReclaimInstance1"))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.ResourcePlanID).To(Equal(testPlanId2))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.Locked).Should(BeFalse())
			Expect(result.LastOperation["type"]).To(Equal("create"))
			Expect(result.LastOperation["async"]).Should(BeFalse())
			Expect(result.LastOperation["state"]).To(Equal("succeeded"))

			testReclaimInstanceCrn = *result.ID
			testReclaimInstanceGuid = *result.Guid
		})

		It("45 - Schedule The Resource Instance For Reclamation", func() {
			shouldSkipTest()

			options := service.NewDeleteResourceInstanceOptions(testReclaimInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test45-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			resp, err := service.DeleteResourceInstance(options)

			Expect(resp.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())

			//wait for reclamation object to be created
			time.Sleep(20 * time.Second)
		})

		It("46 - Verify The Resource Instance Is Pending Reclamation", func() {
			shouldSkipTest()

			options := service.NewGetResourceInstanceOptions(testReclaimInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test46-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testReclaimInstanceCrn))
			Expect(*result.State).To(Equal("pending_reclamation"))
			Expect(result.LastOperation["type"]).To(Equal("reclamation"))
			Expect(result.LastOperation["sub_type"]).To(Equal("pending"))
			Expect(result.LastOperation["async"]).Should(BeFalse())
			Expect(result.LastOperation["state"]).To(Equal("succeeded"))
		})

		It("47 - List Reclamations For Account Id", func() {
			shouldSkipTest()

			options := service.NewListReclamationsOptions()
			options = options.SetAccountID(testAccountId)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test47-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.ListReclamations(options)

			Expect(resp.StatusCode).To(Equal(200))
			Expect(len(result.Resources)).Should(BeNumerically(">=", 1))
			Expect(err).To(BeNil())

			foundReclamation := false
			for _, res := range result.Resources {
				if *res.ResourceInstanceID == testReclaimInstanceGuid {
					Expect(*res.ResourceInstanceID).To(Equal(testReclaimInstanceGuid))
					Expect(*res.AccountID).To(Equal(testAccountId))
					Expect(*res.ResourceGroupID).To(Equal(testResourceGroupGuid))
					Expect(*res.State).To(Equal("SCHEDULED"))

					foundReclamation = true
					testReclamationId1 = *res.ID
				}
			}

			Expect(foundReclamation).To(BeTrue())
		})

		It("48 - Restore A Resource Instance", func() {
			shouldSkipTest()

			options := service.NewRunReclamationActionOptions(testReclamationId1, "restore")
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test48-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.RunReclamationAction(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ResourceInstanceID).To(Equal(testReclaimInstanceGuid))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.State).To(Equal("RESTORING"))

			//wait for reclamation object to be created
			time.Sleep(20 * time.Second)
		})

		It("49 - Verify The Resource Instance Is Restored", func() {
			shouldSkipTest()

			options := service.NewGetResourceInstanceOptions(testReclaimInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test49-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testReclaimInstanceCrn))
			Expect(*result.State).To(Equal("active"))
			Expect(result.LastOperation["type"]).To(Equal("reclamation"))
			Expect(result.LastOperation["sub_type"]).To(Equal("restore"))
			Expect(result.LastOperation["async"]).Should(BeFalse())
			Expect(result.LastOperation["state"]).To(Equal("succeeded"))
		})

		It("50 - Schedule The Resource Instance For Reclamation 2", func() {
			shouldSkipTest()

			options := service.NewDeleteResourceInstanceOptions(testReclaimInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test50-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			resp, err := service.DeleteResourceInstance(options)

			Expect(resp.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())

			//wait for reclamation object to be created
			time.Sleep(20 * time.Second)
		})

		It("51 - List Reclamations For Account and Resource Instance Id", func() {
			shouldSkipTest()

			options := service.NewListReclamationsOptions()
			options = options.SetAccountID(testAccountId)
			options = options.SetResourceInstanceID(testReclaimInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test51-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.ListReclamations(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result.Resources).Should(HaveLen(1))
			Expect(*result.Resources[0].ResourceInstanceID).To(Equal(testReclaimInstanceGuid))
			Expect(*result.Resources[0].AccountID).To(Equal(testAccountId))
			Expect(*result.Resources[0].ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.Resources[0].State).To(Equal("SCHEDULED"))

			testReclamationId2 = *result.Resources[0].ID
		})

		It("52 - Reclaim A Resource Instance", func() {
			shouldSkipTest()

			options := service.NewRunReclamationActionOptions(testReclamationId2, "reclaim")
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test52-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.RunReclamationAction(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ResourceInstanceID).To(Equal(testReclaimInstanceGuid))
			Expect(*result.AccountID).To(Equal(testAccountId))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGuid))
			Expect(*result.State).To(Equal("RECLAIMING"))

			//wait for reclamation object to be created
			time.Sleep(20 * time.Second)
		})

		It("53 - Verify The Resource Instance Is Reclaimed", func() {
			shouldSkipTest()

			options := service.NewGetResourceInstanceOptions(testReclaimInstanceGuid)
			headers := map[string]string{
				"Transaction-Id": "rc-sdk-go-test53-" + uuid.NewV4().String(),
			}
			options = options.SetHeaders(headers)
			result, resp, err := service.GetResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testReclaimInstanceCrn))
			Expect(*result.State).To(Equal("removed"))
			Expect(result.LastOperation["type"]).To(Equal("reclamation"))
			Expect(result.LastOperation["sub_type"]).To(Equal("delete"))
			Expect(result.LastOperation["async"]).Should(BeFalse())
			Expect(result.LastOperation["state"]).To(Equal("succeeded"))
		})
	})
})


// clean up resources
var _ = AfterSuite(func() {
	if !configLoaded {
		return
	}

	fmt.Printf("\n\nCleaning up test resources...\n")
	cleanupResources()
	if testReclaimInstanceGuid != "" {
		cleanupReclamationInstance()
	} else {
		fmt.Printf("Reclamation instance was not created. No cleanup needed.")
	}
})

func cleanupResources() {
	if testInstanceKeyGuid != "" {
		options := service.NewDeleteResourceKeyOptions(testInstanceKeyGuid)
		headers := map[string]string{
			"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
		}
		options = options.SetHeaders(headers)
		resp, err := service.DeleteResourceKey(options)
		if resp.StatusCode == 204 {
			fmt.Printf("Successful cleanup of key %s.\n", testInstanceKeyGuid)
		} else if resp.StatusCode == 410 {
			fmt.Printf("Key %s was already deleted by the tests.\n", testInstanceKeyGuid)
		} else {
			fmt.Printf("Failed to cleanup key %s. Error: %s\n", testInstanceKeyGuid, err.Error())
		}
	} else {
		fmt.Printf("Key for instance was not created. No cleanup needed.\n")
	}

	if testAliasKeyGuid != "" {
		options := service.NewDeleteResourceKeyOptions(testAliasKeyGuid)
		headers := map[string]string{
			"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
		}
		options = options.SetHeaders(headers)
		resp, err := service.DeleteResourceKey(options)
		if resp.StatusCode == 204 {
			fmt.Printf("Successful cleanup of key %s.\n", testAliasKeyGuid)
		} else if resp.StatusCode == 410 {
			fmt.Printf("Key %s was already deleted by the tests.\n", testAliasKeyGuid)
		} else {
			fmt.Printf("Failed to cleanup key %s. Error: %s\n", testAliasKeyGuid, err.Error())
		}
	} else {
		fmt.Printf("Key for alias was not created. No cleanup needed.\n")
	}

	if testBindingGuid != "" {
		options := service.NewDeleteResourceBindingOptions(testBindingGuid)
		headers := map[string]string{
			"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
		}
		options = options.SetHeaders(headers)
		resp, err := service.DeleteResourceBinding(options)
		if resp.StatusCode == 204 {
			fmt.Printf("Successful cleanup of binding %s.\n", testBindingGuid)
		} else if resp.StatusCode == 410 {
			fmt.Printf("Binding %s was already deleted by the tests.\n", testBindingGuid)
		} else {
			fmt.Printf("Failed to cleanup binding %s. Error: %s\n", testBindingGuid, err.Error())
		}
	} else {
		fmt.Printf("Binding was not created. No cleanup needed.\n")
	}

	if testAliasGuid != "" {
		options := service.NewDeleteResourceAliasOptions(testAliasGuid)
		headers := map[string]string{
			"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
		}
		options = options.SetHeaders(headers)
		resp, err := service.DeleteResourceAlias(options)
		if resp.StatusCode == 204 {
			fmt.Printf("Successful cleanup of alias %s.\n", testAliasGuid)
		} else if resp.StatusCode == 410 {
			fmt.Printf("Alias %s was already deleted by the tests.\n", testAliasGuid)
		} else {
			fmt.Printf("Failed to cleanup alias %s. Error: %s\n", testAliasGuid, err.Error())
		}
	} else {
		fmt.Printf("Alias was not created. No cleanup needed.\n")
	}

	if testInstanceGuid != "" {
		cleanupInstance()
	} else {
		fmt.Printf("Instance was not created. No cleanup needed.\n")
	}
}

func cleanupInstance() {
	options := service.NewGetResourceInstanceOptions(testInstanceGuid)
	headers := map[string]string{
		"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
	}
	options = options.SetHeaders(headers)
	result, _, err := service.GetResourceInstance(options)
	if err != nil {
		fmt.Printf("Failed to retrieve instance %s for cleanup.\n", testInstanceGuid)
		return
	}

	if *result.State == "active" && *result.Locked {
		options2 := service.NewUnlockResourceInstanceOptions(testInstanceGuid)
		headers2 := map[string]string{
			"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
		}
		options2 = options2.SetHeaders(headers2)
		_, _, err2 := service.UnlockResourceInstance(options2)
		if err2 != nil {
			fmt.Printf("Failed to unlock instance %s for cleanup. Error: %s\n", testInstanceGuid, err2.Error())
			return
		} 
	}

	options3 := service.NewDeleteResourceInstanceOptions(testInstanceGuid)
	headers3 := map[string]string{
		"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
	}
	options3 = options3.SetHeaders(headers3)
	resp3, err3 := service.DeleteResourceInstance(options3)
	if resp3.StatusCode == 204 {
		fmt.Printf("Successful cleanup of instance %s.\n", testInstanceGuid)
	} else if resp3.StatusCode == 410 {
		fmt.Printf("Instance %s was already deleted by the tests.\n", testInstanceGuid)
	} else {
		fmt.Printf("Failed to cleanup instance %s. Error: %s\n", testInstanceGuid, err3.Error())
	}
}

func cleanupReclamationInstance() {
	options1 := service.NewGetResourceInstanceOptions(testReclaimInstanceGuid)
	headers1 := map[string]string{
		"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
	}
	options1 = options1.SetHeaders(headers1)
	result1, _, err1 := service.GetResourceInstance(options1)
	if err1 != nil {
		fmt.Printf("Failed to retrieve instance %s for cleanup.\n", testReclaimInstanceGuid)
		return
	}

	if *result1.State == "removed" {
		fmt.Printf("Instance %s was already reclaimed by the tests.\n", testReclaimInstanceGuid)
	} else if *result1.State == "pending_reclamation" {
		cleanupInstancePendingReclamation()
	} else {
		options2 := service.NewDeleteResourceInstanceOptions(testReclaimInstanceGuid)
		headers2 := map[string]string{
			"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
		}
		options2 = options2.SetHeaders(headers2)
		resp2, err2 := service.DeleteResourceInstance(options2)
		if resp2.StatusCode == 204 {
			fmt.Printf("Successfully scheduled instance %s for reclamation.\n", testReclaimInstanceGuid)
			time.Sleep(20 * time.Second)
			cleanupInstancePendingReclamation()
		} else {
			fmt.Printf("Failed to schedule active instance %s for reclamation. Error: %s\n", testReclaimInstanceGuid, err2.Error())
		}
	}
}

func cleanupInstancePendingReclamation() {
	options1 := service.NewListReclamationsOptions()
	options1 = options1.SetAccountID(testAccountId)
	options1 = options1.SetResourceInstanceID(testReclaimInstanceGuid)
	headers1 := map[string]string{
		"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
	}
	options1 = options1.SetHeaders(headers1)
	result1, _, err1 := service.ListReclamations(options1)
	if err1 != nil {
		fmt.Printf("Failed to retrieve reclamation to process to reclaim instance %s. Error: %s\n", testReclaimInstanceGuid, err1.Error())
		return
	}

	if len(result1.Resources) == 0 {
		fmt.Printf("Failed to retrieve reclamation to process to reclaim instance %s.\n", testReclaimInstanceGuid)
		return
	}

	reclamationId := *result1.Resources[0].ID
	if *result1.Resources[0].State != "RECLAIMING" {
		options2 := service.NewRunReclamationActionOptions(reclamationId, "reclaim")
		headers2 := map[string]string{
			"Transaction-Id": "rc-sdk-cleanup-" + uuid.NewV4().String(),
		}
		options2 = options2.SetHeaders(headers2)
		_, _, err2 := service.RunReclamationAction(options2)
		if err2 != nil {
			fmt.Printf("Failed to process reclamation %s for instance %s. Error: %s\n", reclamationId, testReclaimInstanceGuid, err2.Error())
		} else {
			fmt.Printf("Successfully reclaimed instance %s.\n", testReclaimInstanceGuid)
		}
	} else {
		fmt.Printf("Instance %s was already reclaimed by the tests.\n", testReclaimInstanceGuid)
	}
}
