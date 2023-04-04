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

package projectv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/damianovesperini/platform-services-go-sdk/projectv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the project service.
//
// The following configuration properties are assumed to be defined:
// PROJECT_URL=<service base url>
// PROJECT_AUTH_TYPE=iam
// PROJECT_APIKEY=<IAM apikey>
// PROJECT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`ProjectV1 Examples Tests`, func() {

	const externalConfigFile = "../project_v1.env"

	var (
		projectService *projectv1.ProjectV1
		config         map[string]string

		// Variables to hold link values
		projectIdLink string
		configIdLink  string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(projectv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			projectServiceOptions := &projectv1.ProjectV1Options{}

			projectService, err = projectv1.NewProjectV1UsingExternalConfig(projectServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(projectService).ToNot(BeNil())
		})
	})

	Describe(`ProjectV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProject request example`, func() {
			fmt.Println("\nCreateProject() result:")
			// begin-create_project

			projectConfigPrototypeModel := &projectv1.ProjectConfigPrototype{
				Name:      core.StringPtr("common-variables"),
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
			}

			createProjectOptions := projectService.NewCreateProjectOptions(
				"acme-microservice",
			)
			createProjectOptions.SetDescription("A microservice to deploy on top of ACME infrastructure.")
			createProjectOptions.SetConfigs([]projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel})

			project, response, err := projectService.CreateProject(createProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(project, "", "  ")
			fmt.Println(string(b))

			// end-create_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(project).ToNot(BeNil())

			projectIdLink = *project.ID
			fmt.Fprintf(GinkgoWriter, "Saved projectIdLink value: %v\n", projectIdLink)
		})
		It(`ListProjects request example`, func() {
			fmt.Println("\nListProjects() result:")
			// begin-list_projects
			listProjectsOptions := &projectv1.ListProjectsOptions{
				Limit:    core.Int64Ptr(int64(10)),
				Complete: core.BoolPtr(false),
			}

			pager, err := projectService.NewProjectsPager(listProjectsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []projectv1.ProjectCollectionMemberWithMetadata
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_projects
		})
		It(`GetProject request example`, func() {
			fmt.Println("\nGetProject() result:")
			// begin-get_project

			getProjectOptions := projectService.NewGetProjectOptions(
				projectIdLink,
			)

			project, response, err := projectService.GetProject(getProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(project, "", "  ")
			fmt.Println(string(b))

			// end-get_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(project).ToNot(BeNil())
		})
		It(`UpdateProject request example`, func() {
			fmt.Println("\nUpdateProject() result:")
			// begin-update_project

			jsonPatchOperationModel := &projectv1.JSONPatchOperation{
				Op:   core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateProjectOptions := projectService.NewUpdateProjectOptions(
				"testString",
				[]projectv1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			project, response, err := projectService.UpdateProject(updateProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(project, "", "  ")
			fmt.Println(string(b))

			// end-update_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(project).ToNot(BeNil())
		})
		It(`CreateConfig request example`, func() {
			fmt.Println("\nCreateConfig() result:")
			// begin-create_config

			projectConfigInputVariableModel := &projectv1.ProjectConfigInputVariable{
				Name: core.StringPtr("account_id"),
			}

			projectConfigSettingCollectionModel := &projectv1.ProjectConfigSettingCollection{
				Name:  core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT"),
				Value: core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"),
			}

			createConfigOptions := projectService.NewCreateConfigOptions(
				"testString",
				"env-stage",
				"1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global",
			)
			createConfigOptions.SetNewLabels([]string{"env:stage", "governance:test", "build:0"})
			createConfigOptions.SetNewDescription("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
			createConfigOptions.SetNewInput([]projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel})
			createConfigOptions.SetNewSetting([]projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel})

			projectConfig, response, err := projectService.CreateConfig(createConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-create_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfig).ToNot(BeNil())

			configIdLink = *projectConfig.ID
			fmt.Fprintf(GinkgoWriter, "Saved configIdLink value: %v\n", configIdLink)
		})
		It(`ListConfigs request example`, func() {
			fmt.Println("\nListConfigs() result:")
			// begin-list_configs

			listConfigsOptions := projectService.NewListConfigsOptions(
				projectIdLink,
			)
			listConfigsOptions.SetProjectID(projectIdLink)

			projectConfigCollection, response, err := projectService.ListConfigs(listConfigsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfigCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_configs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigCollection).ToNot(BeNil())
		})
		It(`GetConfig request example`, func() {
			fmt.Println("\nGetConfig() result:")
			// begin-get_config

			getConfigOptions := projectService.NewGetConfigOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfig, response, err := projectService.GetConfig(getConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-get_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`UpdateConfig request example`, func() {
			fmt.Println("\nUpdateConfig() result:")
			// begin-update_config

			jsonPatchOperationModel := &projectv1.JSONPatchOperation{
				Op:   core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateConfigOptions := projectService.NewUpdateConfigOptions(
				projectIdLink,
				configIdLink,
				[]projectv1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			projectConfig, response, err := projectService.UpdateConfig(updateConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-update_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`GetConfigDiff request example`, func() {
			fmt.Println("\nGetConfigDiff() result:")
			// begin-get_config_diff

			getConfigDiffOptions := projectService.NewGetConfigDiffOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfigDiff, response, err := projectService.GetConfigDiff(getConfigDiffOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfigDiff, "", "  ")
			fmt.Println(string(b))

			// end-get_config_diff

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDiff).ToNot(BeNil())
		})
		It(`ForceApprove request example`, func() {
			fmt.Println("\nForceApprove() result:")
			// begin-force_approve

			forceApproveOptions := projectService.NewForceApproveOptions(
				projectIdLink,
				configIdLink,
			)
			forceApproveOptions.SetComment("Approving the changes")

			projectConfig, response, err := projectService.ForceApprove(forceApproveOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-force_approve

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`Approve request example`, func() {
			fmt.Println("\nApprove() result:")
			// begin-approve

			approveOptions := projectService.NewApproveOptions(
				projectIdLink,
				configIdLink,
			)
			approveOptions.SetComment("Approving the changes")

			projectConfig, response, err := projectService.Approve(approveOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-approve

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`CheckConfig request example`, func() {
			fmt.Println("\nCheckConfig() result:")
			// begin-check_config

			checkConfigOptions := projectService.NewCheckConfigOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfig, response, err := projectService.CheckConfig(checkConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-check_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`InstallConfig request example`, func() {
			fmt.Println("\nInstallConfig() result:")
			// begin-install_config

			installConfigOptions := projectService.NewInstallConfigOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfig, response, err := projectService.InstallConfig(installConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-install_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`UninstallConfig request example`, func() {
			// begin-uninstall_config

			uninstallConfigOptions := projectService.NewUninstallConfigOptions(
				projectIdLink,
				configIdLink,
			)

			response, err := projectService.UninstallConfig(uninstallConfigOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from UninstallConfig(): %d\n", response.StatusCode)
			}

			// end-uninstall_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`GetSchematicsJob request example`, func() {
			fmt.Println("\nGetSchematicsJob() result:")
			// begin-get_schematics_job

			getSchematicsJobOptions := projectService.NewGetSchematicsJobOptions(
				projectIdLink,
				configIdLink,
				"plan",
			)

			actionJob, response, err := projectService.GetSchematicsJob(getSchematicsJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionJob, "", "  ")
			fmt.Println(string(b))

			// end-get_schematics_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionJob).ToNot(BeNil())
		})
		It(`GetCostEstimate request example`, func() {
			fmt.Println("\nGetCostEstimate() result:")
			// begin-get_cost_estimate

			getCostEstimateOptions := projectService.NewGetCostEstimateOptions(
				projectIdLink,
				configIdLink,
			)

			costEstimate, response, err := projectService.GetCostEstimate(getCostEstimateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(costEstimate, "", "  ")
			fmt.Println(string(b))

			// end-get_cost_estimate

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(costEstimate).ToNot(BeNil())
		})
		It(`PostCrnToken request example`, func() {
			fmt.Println("\nPostCrnToken() result:")
			// begin-post_crn_token

			postCrnTokenOptions := projectService.NewPostCrnTokenOptions(
				projectIdLink,
			)

			projectCrnTokenResponse, response, err := projectService.PostCrnToken(postCrnTokenOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectCrnTokenResponse, "", "  ")
			fmt.Println(string(b))

			// end-post_crn_token

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectCrnTokenResponse).ToNot(BeNil())
		})
		It(`PostNotification request example`, func() {
			fmt.Println("\nPostNotification() result:")
			// begin-post_notification

			notificationEventModel := &projectv1.NotificationEvent{
				Event:       core.StringPtr("project.create.failed"),
				Target:      core.StringPtr("234234324-3444-4556-224232432"),
				Source:      core.StringPtr("id.of.project.service.instance"),
				TriggeredBy: core.StringPtr("user-iam-id"),
				ActionURL:   core.StringPtr("actionable/url"),
			}

			postNotificationOptions := projectService.NewPostNotificationOptions(
				projectIdLink,
			)
			postNotificationOptions.SetNotifications([]projectv1.NotificationEvent{*notificationEventModel})

			notificationsPrototypePostResponse, response, err := projectService.PostNotification(postNotificationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationsPrototypePostResponse, "", "  ")
			fmt.Println(string(b))

			// end-post_notification

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsPrototypePostResponse).ToNot(BeNil())
		})
		It(`GetNotifications request example`, func() {
			fmt.Println("\nGetNotifications() result:")
			// begin-get_notifications

			getNotificationsOptions := projectService.NewGetNotificationsOptions(
				projectIdLink,
			)

			notificationsGetResponse, response, err := projectService.GetNotifications(getNotificationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationsGetResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_notifications

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsGetResponse).ToNot(BeNil())
		})
		It(`ReceivePulsarCatalogEvents request example`, func() {
			// begin-receive_pulsar_catalog_events

			pulsarEventPrototypeCollectionModel := &projectv1.PulsarEventPrototypeCollection{
				EventType: core.StringPtr("create"),
				Timestamp: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Publisher: core.StringPtr("provider"),
				AccountID: core.StringPtr("accountId"),
				Version:   core.StringPtr("v1"),
			}

			receivePulsarCatalogEventsOptions := projectService.NewReceivePulsarCatalogEventsOptions(
				[]projectv1.PulsarEventPrototypeCollection{*pulsarEventPrototypeCollectionModel},
			)

			response, err := projectService.ReceivePulsarCatalogEvents(receivePulsarCatalogEventsOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from ReceivePulsarCatalogEvents(): %d\n", response.StatusCode)
			}

			// end-receive_pulsar_catalog_events

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`ReceivePulsarEventNotificationEvents request example`, func() {
			// begin-receive_pulsar_event_notification_events

			pulsarEventPrototypeCollectionModel := &projectv1.PulsarEventPrototypeCollection{
				EventType: core.StringPtr("testString"),
				Timestamp: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Publisher: core.StringPtr("provider"),
				AccountID: core.StringPtr("accountid"),
				Version:   core.StringPtr("v1"),
			}

			receivePulsarEventNotificationEventsOptions := projectService.NewReceivePulsarEventNotificationEventsOptions(
				[]projectv1.PulsarEventPrototypeCollection{*pulsarEventPrototypeCollectionModel},
			)

			response, err := projectService.ReceivePulsarEventNotificationEvents(receivePulsarEventNotificationEventsOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from ReceivePulsarEventNotificationEvents(): %d\n", response.StatusCode)
			}

			// end-receive_pulsar_event_notification_events

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`GetHealth request example`, func() {
			fmt.Println("\nGetHealth() result:")
			// begin-get_health

			getHealthOptions := projectService.NewGetHealthOptions()

			health, response, err := projectService.GetHealth(getHealthOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(health, "", "  ")
			fmt.Println(string(b))

			// end-get_health

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(health).ToNot(BeNil())
		})
		It(`ReplaceServiceInstance request example`, func() {
			fmt.Println("\nReplaceServiceInstance() result:")
			// begin-replace_service_instance

			replaceServiceInstanceOptions := projectService.NewReplaceServiceInstanceOptions(
				"crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::",
				projectIdLink,
				configIdLink,
			)
			replaceServiceInstanceOptions.SetXBrokerApiVersion("1.0")
			replaceServiceInstanceOptions.SetXBrokerApiOriginatingIdentity("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
			replaceServiceInstanceOptions.SetAcceptsIncomplete(false)

			resourceCreateResponse, response, err := projectService.ReplaceServiceInstance(replaceServiceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceCreateResponse, "", "  ")
			fmt.Println(string(b))

			// end-replace_service_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceCreateResponse).ToNot(BeNil())
		})
		It(`UpdateServiceInstance request example`, func() {
			fmt.Println("\nUpdateServiceInstance() result:")
			// begin-update_service_instance

			jsonPatchOperationModel := &projectv1.JSONPatchOperation{
				Op:   core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateServiceInstanceOptions := projectService.NewUpdateServiceInstanceOptions(
				"crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::",
				[]projectv1.JSONPatchOperation{*jsonPatchOperationModel},
			)
			updateServiceInstanceOptions.SetXBrokerApiVersion("1.0")
			updateServiceInstanceOptions.SetXBrokerApiOriginatingIdentity("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
			updateServiceInstanceOptions.SetAcceptsIncomplete(false)

			resourceUpdateResult, response, err := projectService.UpdateServiceInstance(updateServiceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceUpdateResult, "", "  ")
			fmt.Println(string(b))

			// end-update_service_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceUpdateResult).ToNot(BeNil())
		})
		It(`GetLastOperation request example`, func() {
			fmt.Println("\nGetLastOperation() result:")
			// begin-get_last_operation

			getLastOperationOptions := projectService.NewGetLastOperationOptions(
				"crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::",
			)
			getLastOperationOptions.SetXBrokerApiVersion("1.0")
			getLastOperationOptions.SetOperation("ABCD")
			getLastOperationOptions.SetPlanID("cb54391b-3316-4943-a5a6-a541678c1924")
			getLastOperationOptions.SetServiceID("cb54391b-3316-4943-a5a6-a541678c1924")

			resourceLastOperationGetResponse, response, err := projectService.GetLastOperation(getLastOperationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceLastOperationGetResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_last_operation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceLastOperationGetResponse).ToNot(BeNil())
		})
		It(`ReplaceServiceInstanceState request example`, func() {
			fmt.Println("\nReplaceServiceInstanceState() result:")
			// begin-replace_service_instance_state

			replaceServiceInstanceStateOptions := projectService.NewReplaceServiceInstanceStateOptions(
				"crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::",
				true,
			)
			replaceServiceInstanceStateOptions.SetXBrokerApiVersion("1.0")

			resourceStateResponse, response, err := projectService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceStateResponse, "", "  ")
			fmt.Println(string(b))

			// end-replace_service_instance_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceStateResponse).ToNot(BeNil())
		})
		It(`GetServiceInstance request example`, func() {
			fmt.Println("\nGetServiceInstance() result:")
			// begin-get_service_instance

			getServiceInstanceOptions := projectService.NewGetServiceInstanceOptions(
				"crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::",
			)
			getServiceInstanceOptions.SetXBrokerApiVersion("1.0")

			resourceStateResponse, response, err := projectService.GetServiceInstance(getServiceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceStateResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_service_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceStateResponse).ToNot(BeNil())
		})
		It(`GetCatalog request example`, func() {
			fmt.Println("\nGetCatalog() result:")
			// begin-get_catalog

			getCatalogOptions := projectService.NewGetCatalogOptions()
			getCatalogOptions.SetXBrokerApiVersion("1.0")

			catalogResponse, response, err := projectService.GetCatalog(getCatalogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogResponse).ToNot(BeNil())
		})
		It(`PostEventNotificationsIntegration request example`, func() {
			fmt.Println("\nPostEventNotificationsIntegration() result:")
			// begin-post_event_notifications_integration

			postEventNotificationsIntegrationOptions := projectService.NewPostEventNotificationsIntegrationOptions(
				projectIdLink,
				"CRN of event notifications instance",
			)
			postEventNotificationsIntegrationOptions.SetDescription("A sample project source.")
			postEventNotificationsIntegrationOptions.SetEventNotificationsSourceName("project 1 source name for event notifications")
			postEventNotificationsIntegrationOptions.SetEnabled(true)

			notificationsIntegrationPostResponse, response, err := projectService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationsIntegrationPostResponse, "", "  ")
			fmt.Println(string(b))

			// end-post_event_notifications_integration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsIntegrationPostResponse).ToNot(BeNil())
		})
		It(`GetEventNotificationsIntegration request example`, func() {
			fmt.Println("\nGetEventNotificationsIntegration() result:")
			// begin-get_event_notifications_integration

			getEventNotificationsIntegrationOptions := projectService.NewGetEventNotificationsIntegrationOptions(
				projectIdLink,
			)

			notificationsIntegrationGetResponse, response, err := projectService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationsIntegrationGetResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_event_notifications_integration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsIntegrationGetResponse).ToNot(BeNil())
		})
		It(`PostTestEventNotification request example`, func() {
			fmt.Println("\nPostTestEventNotification() result:")
			// begin-post_test_event_notification

			postTestEventNotificationOptions := projectService.NewPostTestEventNotificationOptions(
				projectIdLink,
			)
			postTestEventNotificationOptions.SetIbmendefaultlong("long test notification message")
			postTestEventNotificationOptions.SetIbmendefaultshort("Test notification")

			notificationsIntegrationTestPostResponse, response, err := projectService.PostTestEventNotification(postTestEventNotificationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationsIntegrationTestPostResponse, "", "  ")
			fmt.Println(string(b))

			// end-post_test_event_notification

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsIntegrationTestPostResponse).ToNot(BeNil())
		})
		It(`DeleteProject request example`, func() {
			// begin-delete_project

			deleteProjectOptions := projectService.NewDeleteProjectOptions(
				projectIdLink,
			)

			response, err := projectService.DeleteProject(deleteProjectOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteProject(): %d\n", response.StatusCode)
			}

			// end-delete_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteConfig request example`, func() {
			fmt.Println("\nDeleteConfig() result:")
			// begin-delete_config

			deleteConfigOptions := projectService.NewDeleteConfigOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfigDelete, response, err := projectService.DeleteConfig(deleteConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfigDelete, "", "  ")
			fmt.Println(string(b))

			// end-delete_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDelete).ToNot(BeNil())
		})
		It(`DeleteServiceInstance request example`, func() {
			fmt.Println("\nDeleteServiceInstance() result:")
			// begin-delete_service_instance

			deleteServiceInstanceOptions := projectService.NewDeleteServiceInstanceOptions(
				"crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::",
				"cb54391b-3316-4943-a5a6-a541678c1924",
				"cb54391b-3316-4943-a5a6-a541678c1924",
			)
			deleteServiceInstanceOptions.SetXBrokerApiVersion("1.0")
			deleteServiceInstanceOptions.SetXBrokerApiOriginatingIdentity("ibmcloud eyJpYW1fbWQiOiJJQk2pZC03MEdOUjcxN2lFIn0=")
			deleteServiceInstanceOptions.SetAcceptsIncomplete(false)

			resourceDeleteResponse, response, err := projectService.DeleteServiceInstance(deleteServiceInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceDeleteResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_service_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceDeleteResponse).ToNot(BeNil())
		})
		It(`DeleteEventNotificationsIntegration request example`, func() {
			// begin-delete_event_notifications_integration

			deleteEventNotificationsIntegrationOptions := projectService.NewDeleteEventNotificationsIntegrationOptions(
				projectIdLink,
			)

			response, err := projectService.DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteEventNotificationsIntegration(): %d\n", response.StatusCode)
			}

			// end-delete_event_notifications_integration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
