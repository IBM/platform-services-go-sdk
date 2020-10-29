// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/openservicebrokerv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const externalConfigFile = "../open_service_broker.env"

var (
	openServiceBrokerService *openservicebrokerv1.OpenServiceBrokerV1
	config                   map[string]string
	configLoaded             bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`OpenServiceBrokerV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(openservicebrokerv1.DefaultServiceName)
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

			openServiceBrokerServiceOptions := &openservicebrokerv1.OpenServiceBrokerV1Options{}

			openServiceBrokerService, err = openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(openServiceBrokerServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(openServiceBrokerService).ToNot(BeNil())
		})
	})

	Describe(`OpenServiceBrokerV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServiceInstanceState request example`, func() {
			// begin-get_service_instance_state

			getServiceInstanceStateOptions := openServiceBrokerService.NewGetServiceInstanceStateOptions(
				"testString",
			)

			resp1874644Root, response, err := openServiceBrokerService.GetServiceInstanceState(getServiceInstanceStateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resp1874644Root, "", "  ")
			fmt.Println(string(b))

			// end-get_service_instance_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resp1874644Root).ToNot(BeNil())

		})
		It(`ReplaceServiceInstanceState request example`, func() {
			// begin-replace_service_instance_state

			replaceServiceInstanceStateOptions := openServiceBrokerService.NewReplaceServiceInstanceStateOptions(
				"testString",
			)

			resp2448145Root, response, err := openServiceBrokerService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resp2448145Root, "", "  ")
			fmt.Println(string(b))

			// end-replace_service_instance_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resp2448145Root).ToNot(BeNil())

		})
		It(`ReplaceServiceInstance request example`, func() {
			// begin-replace_service_instance

			replaceServiceInstanceOptions := openServiceBrokerService.NewReplaceServiceInstanceOptions(
				"testString",
			)

			resp2079872Root, response, err := openServiceBrokerService.ReplaceServiceInstance(replaceServiceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resp2079872Root, "", "  ")
			fmt.Println(string(b))

			// end-replace_service_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resp2079872Root).ToNot(BeNil())

		})
		It(`UpdateServiceInstance request example`, func() {
			// begin-update_service_instance

			updateServiceInstanceOptions := openServiceBrokerService.NewUpdateServiceInstanceOptions(
				"testString",
			)

			resp2079874Root, response, err := openServiceBrokerService.UpdateServiceInstance(updateServiceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resp2079874Root, "", "  ")
			fmt.Println(string(b))

			// end-update_service_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resp2079874Root).ToNot(BeNil())

		})
		It(`ListCatalog request example`, func() {
			// begin-list_catalog

			listCatalogOptions := openServiceBrokerService.NewListCatalogOptions()

			resp1874650Root, response, err := openServiceBrokerService.ListCatalog(listCatalogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resp1874650Root, "", "  ")
			fmt.Println(string(b))

			// end-list_catalog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resp1874650Root).ToNot(BeNil())

		})
		It(`GetLastOperation request example`, func() {
			// begin-get_last_operation

			getLastOperationOptions := openServiceBrokerService.NewGetLastOperationOptions(
				"testString",
			)

			resp2079894Root, response, err := openServiceBrokerService.GetLastOperation(getLastOperationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resp2079894Root, "", "  ")
			fmt.Println(string(b))

			// end-get_last_operation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resp2079894Root).ToNot(BeNil())

		})
		It(`ReplaceServiceBinding request example`, func() {
			// begin-replace_service_binding

			replaceServiceBindingOptions := openServiceBrokerService.NewReplaceServiceBindingOptions(
				"testString",
				"testString",
			)

			resp2079876Root, response, err := openServiceBrokerService.ReplaceServiceBinding(replaceServiceBindingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resp2079876Root, "", "  ")
			fmt.Println(string(b))

			// end-replace_service_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resp2079876Root).ToNot(BeNil())

		})
		It(`DeleteServiceInstance request example`, func() {
			// begin-delete_service_instance

			deleteServiceInstanceOptions := openServiceBrokerService.NewDeleteServiceInstanceOptions(
				"testString",
				"testString",
				"testString",
			)

			resp2079874Root, response, err := openServiceBrokerService.DeleteServiceInstance(deleteServiceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resp2079874Root, "", "  ")
			fmt.Println(string(b))

			// end-delete_service_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resp2079874Root).ToNot(BeNil())

		})
		It(`DeleteServiceBinding request example`, func() {
			// begin-delete_service_binding

			deleteServiceBindingOptions := openServiceBrokerService.NewDeleteServiceBindingOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := openServiceBrokerService.DeleteServiceBinding(deleteServiceBindingOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_service_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})
