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

package projectsv1_test

import (
	"github.com/IBM/platform-services-go-sdk/plugin/commands/projectsv1"
	"github.com/IBM/platform-services-go-sdk/testing_utilities"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var positiveFakeUtils = testing_utilities.NewPositiveTestUtilities()
var negativeFakeUtils = testing_utilities.NewNegativeTestUtilities()

// Test suite
var _ = Describe("ProjectsV1", func() {
	// ensure the service instance is newly created during each test
	BeforeEach(func() {
	  projectsv1.ServiceInstance = nil
	})

	Describe("ListProjects", func() {
		// put together mock arguments
		Start := `--start=testString`
		Limit := `--limit=10`
		Complete := `--complete=false`

		args := []string{
			Start,
			Limit,
			Complete,
		}

		It("Puts together ListProjects options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewListProjectsCommandRunner(positiveFakeUtils, ListProjectsMockSender{}, nil)
			command := projectsv1.GetListProjectsCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
		It("Uses dedicated sender for ListProjects when --all-pages is set", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewListProjectsCommandRunner(positiveFakeUtils, nil, ListProjectsMockSender{})
			command := projectsv1.GetListProjectsCommand(runner)

			newArgs := append(args, "--all-pages")
			command.SetArgs(newArgs)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for ListProjects", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewListProjectsCommandRunner(negativeFakeUtils, ListProjectsErrorSender{}, nil)
			command := projectsv1.GetListProjectsCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("GetProject", func() {
		// put together mock arguments
		ID := `--id=testString`
		ExcludeConfigs := `--exclude-configs=false`
		Complete := `--complete=false`

		args := []string{
			ID,
			ExcludeConfigs,
			Complete,
		}

		It("Puts together GetProject options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetProjectCommandRunner(positiveFakeUtils, GetProjectMockSender{})
			command := projectsv1.GetGetProjectCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for GetProject", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetProjectCommandRunner(negativeFakeUtils, GetProjectErrorSender{})
			command := projectsv1.GetGetProjectCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("DeleteProject", func() {
		// put together mock arguments
		ID := `--id=testString`

		args := []string{
			ID,
		}

		It("Puts together DeleteProject options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewDeleteProjectCommandRunner(positiveFakeUtils, DeleteProjectMockSender{})
			command := projectsv1.GetDeleteProjectCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for DeleteProject", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewDeleteProjectCommandRunner(negativeFakeUtils, DeleteProjectErrorSender{})
			command := projectsv1.GetDeleteProjectCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("ListConfigs", func() {
		// put together mock arguments
		ID := `--id=testString`
		Version := `--version=active`
		Complete := `--complete=false`

		args := []string{
			ID,
			Version,
			Complete,
		}

		It("Puts together ListConfigs options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewListConfigsCommandRunner(positiveFakeUtils, ListConfigsMockSender{})
			command := projectsv1.GetListConfigsCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for ListConfigs", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewListConfigsCommandRunner(negativeFakeUtils, ListConfigsErrorSender{})
			command := projectsv1.GetListConfigsCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("GetConfig", func() {
		// put together mock arguments
		ID := `--id=testString`
		ConfigID := `--config-id=testString`
		Version := `--version=active`
		Complete := `--complete=false`

		args := []string{
			ID,
			ConfigID,
			Version,
			Complete,
		}

		It("Puts together GetConfig options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetConfigCommandRunner(positiveFakeUtils, GetConfigMockSender{})
			command := projectsv1.GetGetConfigCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for GetConfig", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetConfigCommandRunner(negativeFakeUtils, GetConfigErrorSender{})
			command := projectsv1.GetGetConfigCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("DeleteConfig", func() {
		// put together mock arguments
		ID := `--id=testString`
		ConfigID := `--config-id=testString`

		args := []string{
			ID,
			ConfigID,
		}

		It("Puts together DeleteConfig options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewDeleteConfigCommandRunner(positiveFakeUtils, DeleteConfigMockSender{})
			command := projectsv1.GetDeleteConfigCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for DeleteConfig", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewDeleteConfigCommandRunner(negativeFakeUtils, DeleteConfigErrorSender{})
			command := projectsv1.GetDeleteConfigCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("GetConfigDiff", func() {
		// put together mock arguments
		ID := `--id=testString`
		ConfigID := `--config-id=testString`

		args := []string{
			ID,
			ConfigID,
		}

		It("Puts together GetConfigDiff options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetConfigDiffCommandRunner(positiveFakeUtils, GetConfigDiffMockSender{})
			command := projectsv1.GetGetConfigDiffCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for GetConfigDiff", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetConfigDiffCommandRunner(negativeFakeUtils, GetConfigDiffErrorSender{})
			command := projectsv1.GetGetConfigDiffCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("InstallConfig", func() {
		// put together mock arguments
		ID := `--id=testString`
		ConfigID := `--config-id=testString`

		args := []string{
			ID,
			ConfigID,
		}

		It("Puts together InstallConfig options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewInstallConfigCommandRunner(positiveFakeUtils, InstallConfigMockSender{})
			command := projectsv1.GetInstallConfigCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for InstallConfig", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewInstallConfigCommandRunner(negativeFakeUtils, InstallConfigErrorSender{})
			command := projectsv1.GetInstallConfigCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("UninstallConfig", func() {
		// put together mock arguments
		ID := `--id=testString`
		ConfigID := `--config-id=testString`

		args := []string{
			ID,
			ConfigID,
		}

		It("Puts together UninstallConfig options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewUninstallConfigCommandRunner(positiveFakeUtils, UninstallConfigMockSender{})
			command := projectsv1.GetUninstallConfigCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for UninstallConfig", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewUninstallConfigCommandRunner(negativeFakeUtils, UninstallConfigErrorSender{})
			command := projectsv1.GetUninstallConfigCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("GetSchematicsJob", func() {
		// put together mock arguments
		ID := `--id=testString`
		ConfigID := `--config-id=testString`
		Action := `--action=plan`
		Since := `--since=38`

		args := []string{
			ID,
			ConfigID,
			Action,
			Since,
		}

		It("Puts together GetSchematicsJob options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetSchematicsJobCommandRunner(positiveFakeUtils, GetSchematicsJobMockSender{})
			command := projectsv1.GetGetSchematicsJobCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for GetSchematicsJob", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetSchematicsJobCommandRunner(negativeFakeUtils, GetSchematicsJobErrorSender{})
			command := projectsv1.GetGetSchematicsJobCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("GetCostEstimate", func() {
		// put together mock arguments
		ID := `--id=testString`
		ConfigID := `--config-id=testString`
		Version := `--version=active`

		args := []string{
			ID,
			ConfigID,
			Version,
		}

		It("Puts together GetCostEstimate options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetCostEstimateCommandRunner(positiveFakeUtils, GetCostEstimateMockSender{})
			command := projectsv1.GetGetCostEstimateCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for GetCostEstimate", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetCostEstimateCommandRunner(negativeFakeUtils, GetCostEstimateErrorSender{})
			command := projectsv1.GetGetCostEstimateCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("PostNotification", func() {
		// put together mock arguments
		ID := `--id=testString`
		Notifications := `--notifications=[{"event": "project.create.failed", "target": "234234324-3444-4556-224232432", "source": "id.of.project.service.instance", "action_url": "url.for.project.documentation", "data": {"anyKey": "anyValue"}}]`

		args := []string{
			ID,
			Notifications,
		}

		It("Puts together PostNotification options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewPostNotificationCommandRunner(positiveFakeUtils, PostNotificationMockSender{})
			command := projectsv1.GetPostNotificationCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for PostNotification", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewPostNotificationCommandRunner(negativeFakeUtils, PostNotificationErrorSender{})
			command := projectsv1.GetPostNotificationCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("GetHealth", func() {
		// put together mock arguments
		Info := `--info=false`

		args := []string{
			Info,
		}

		It("Puts together GetHealth options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetHealthCommandRunner(positiveFakeUtils, GetHealthMockSender{})
			command := projectsv1.GetGetHealthCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for GetHealth", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewGetHealthCommandRunner(negativeFakeUtils, GetHealthErrorSender{})
			command := projectsv1.GetGetHealthCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("PostEventNotificationsIntegration", func() {
		// put together mock arguments
		ID := `--id=testString`
		InstanceCrn := `--instance-crn=crn of event notifications instance`
		Description := `--description=A sample project source`
		Name := `--name=Project name`
		Enabled := `--enabled=true`
		Source := `--source=CRN of the project instance`

		args := []string{
			ID,
			InstanceCrn,
			Description,
			Name,
			Enabled,
			Source,
		}

		It("Puts together PostEventNotificationsIntegration options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewPostEventNotificationsIntegrationCommandRunner(positiveFakeUtils, PostEventNotificationsIntegrationMockSender{})
			command := projectsv1.GetPostEventNotificationsIntegrationCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for PostEventNotificationsIntegration", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewPostEventNotificationsIntegrationCommandRunner(negativeFakeUtils, PostEventNotificationsIntegrationErrorSender{})
			command := projectsv1.GetPostEventNotificationsIntegrationCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	Describe("PostEventNotification", func() {
		// put together mock arguments
		ID := `--id=testString`
		NewID := `--new-id=5f208fef-6b64-413c-aa07-dfed0b46abc1236`
		NewSource := `--new-source=crn of project`
		NewDatacontenttype := `--new-datacontenttype=application/json`
		NewIbmendefaultlong := `--new-ibmendefaultlong=long test notification message`
		NewIbmendefaultshort := `--new-ibmendefaultshort=Test notification`
		NewIbmensourceid := `--new-ibmensourceid=crn of project`
		NewSpecversion := `--new-specversion=1.0`
		NewType := `--new-type=com.ibm.cloud.project.project.test_notification`

		args := []string{
			ID,
			NewID,
			NewSource,
			NewDatacontenttype,
			NewIbmendefaultlong,
			NewIbmendefaultshort,
			NewIbmensourceid,
			NewSpecversion,
			NewType,
		}

		It("Puts together PostEventNotification options model", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewPostEventNotificationCommandRunner(positiveFakeUtils, PostEventNotificationMockSender{})
			command := projectsv1.GetPostEventNotificationCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})

		It("Tests error handling for PostEventNotification", func() {
			projectsv1.Service = &testing_utilities.TestServiceCommandHelper{}
			runner := projectsv1.NewPostEventNotificationCommandRunner(negativeFakeUtils, PostEventNotificationErrorSender{})
			command := projectsv1.GetPostEventNotificationCommand(runner)

			command.SetArgs(args)

			_, err := command.ExecuteC()
			Expect(err).To(BeNil())
		})
	})

	// Test the Service Command Getter
	It("Gets the parent command for ProjectsV1", func() {
		projectsv1.InitializeService(positiveFakeUtils)
		command := projectsv1.GetProjectsV1Command(positiveFakeUtils)
		Expect(command.Use).To(Equal("project [command] [options]"))
	})

	// Test the analytics header
	It("Sets a custom analytics header on the service instance", func() {
		projectsv1.InitializeService(positiveFakeUtils)
		command := projectsv1.GetProjectsV1Command(positiveFakeUtils)
		// need to explicitly set empty args or the ginkgo default args get picked up and cause errors
		command.SetArgs([]string{})
		command.SetHelpFunc(func(c *cobra.Command, a []string) {
			// calling the parent command without a subcommand will print the help menu
			// stub out the method for the test
		})

		err := command.Execute()
		Expect(err).To(BeNil())

		flagSet := command.Flags()
		projectsv1.Service.InitializeServiceInstance(flagSet)

		CheckAnalyticsHeader(projectsv1.ServiceInstance)
	})

	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockMap() successfully`, func() {
			mockMap := CreateMockMap()
			Expect(mockMap).ToNot(BeNil())
		})

		It(`Invoke CreateMockByteArray() successfully`, func() {
			bytes := CreateMockByteArray("VGhpcyBpcyBhIG1vY2sgYnl0ZSBhcnJheSB2YWx1ZS4=")
			Expect(bytes).ToNot(BeNil())
		})

		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})

		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})

		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})

		It(`Invoke ResolveModel() successfully`, func() {
			model := ResolveModel(map[string]string{
				"foo": "bar",
			})
			Expect(model).ToNot(BeNil())
		})
	})
})
