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
	"encoding/base64"
	"encoding/json"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/plugin/version"
	"github.com/IBM/platform-services-go-sdk/projectsv1"
	"github.com/IBM/platform-services-go-sdk/testing_utilities"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/gomega"
)

// Fake senders for ListProjects
type ListProjectsMockSender struct{}

func (f ListProjectsMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.ListProjectsOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.Start).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.Limit).To(Equal(core.Int64Ptr(int64(10))))
	Expect(createdOptions.Complete).To(Equal(core.BoolPtr(false)))
	return testing_utilities.GetMockSuccessResponse()
}

type ListProjectsErrorSender struct{}

func (f ListProjectsErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for GetProject
type GetProjectMockSender struct{}

func (f GetProjectMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.GetProjectOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.ExcludeConfigs).To(Equal(core.BoolPtr(false)))
	Expect(createdOptions.Complete).To(Equal(core.BoolPtr(false)))
	return testing_utilities.GetMockSuccessResponse()
}

type GetProjectErrorSender struct{}

func (f GetProjectErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for DeleteProject
type DeleteProjectMockSender struct{}

func (f DeleteProjectMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.DeleteProjectOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	return testing_utilities.GetMockSuccessResponse()
}

type DeleteProjectErrorSender struct{}

func (f DeleteProjectErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for ListConfigs
type ListConfigsMockSender struct{}

func (f ListConfigsMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.ListConfigsOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.Version).To(Equal(core.StringPtr("active")))
	Expect(createdOptions.Complete).To(Equal(core.BoolPtr(false)))
	return testing_utilities.GetMockSuccessResponse()
}

type ListConfigsErrorSender struct{}

func (f ListConfigsErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for GetConfig
type GetConfigMockSender struct{}

func (f GetConfigMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.GetConfigOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.ConfigID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.Version).To(Equal(core.StringPtr("active")))
	Expect(createdOptions.Complete).To(Equal(core.BoolPtr(false)))
	return testing_utilities.GetMockSuccessResponse()
}

type GetConfigErrorSender struct{}

func (f GetConfigErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for DeleteConfig
type DeleteConfigMockSender struct{}

func (f DeleteConfigMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.DeleteConfigOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.ConfigID).To(Equal(core.StringPtr("testString")))
	return testing_utilities.GetMockSuccessResponse()
}

type DeleteConfigErrorSender struct{}

func (f DeleteConfigErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for GetConfigDiff
type GetConfigDiffMockSender struct{}

func (f GetConfigDiffMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.GetConfigDiffOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.ConfigID).To(Equal(core.StringPtr("testString")))
	return testing_utilities.GetMockSuccessResponse()
}

type GetConfigDiffErrorSender struct{}

func (f GetConfigDiffErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for InstallConfig
type InstallConfigMockSender struct{}

func (f InstallConfigMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.InstallConfigOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.ConfigID).To(Equal(core.StringPtr("testString")))
	return testing_utilities.GetMockSuccessResponse()
}

type InstallConfigErrorSender struct{}

func (f InstallConfigErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for UninstallConfig
type UninstallConfigMockSender struct{}

func (f UninstallConfigMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.UninstallConfigOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.ConfigID).To(Equal(core.StringPtr("testString")))
	return testing_utilities.GetMockSuccessResponse()
}

type UninstallConfigErrorSender struct{}

func (f UninstallConfigErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for GetSchematicsJob
type GetSchematicsJobMockSender struct{}

func (f GetSchematicsJobMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.GetSchematicsJobOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.ConfigID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.Action).To(Equal(core.StringPtr("plan")))
	Expect(createdOptions.Since).To(Equal(core.Int64Ptr(int64(38))))
	return testing_utilities.GetMockSuccessResponse()
}

type GetSchematicsJobErrorSender struct{}

func (f GetSchematicsJobErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for GetCostEstimate
type GetCostEstimateMockSender struct{}

func (f GetCostEstimateMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.GetCostEstimateOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.ConfigID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.Version).To(Equal(core.StringPtr("active")))
	return testing_utilities.GetMockSuccessResponse()
}

type GetCostEstimateErrorSender struct{}

func (f GetCostEstimateErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for PostNotification
type PostNotificationMockSender struct{}

func (f PostNotificationMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	// Construct an instance of the NotificationEvent model
	notificationEventModel := new(projectsv1.NotificationEvent)
	notificationEventModel.Event = core.StringPtr("project.create.failed")
	notificationEventModel.Target = core.StringPtr("234234324-3444-4556-224232432")
	notificationEventModel.Source = core.StringPtr("id.of.project.service.instance")
	notificationEventModel.ActionURL = core.StringPtr("url.for.project.documentation")
	notificationEventModel.Data = map[string]interface{}{"anyKey": "anyValue"}

	createdOptions, ok := optionsModel.(*projectsv1.PostNotificationOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(ResolveModel(createdOptions.Notifications)).To(Equal(ResolveModel([]projectsv1.NotificationEvent{*notificationEventModel})))
	return testing_utilities.GetMockSuccessResponse()
}

type PostNotificationErrorSender struct{}

func (f PostNotificationErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for GetHealth
type GetHealthMockSender struct{}

func (f GetHealthMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.GetHealthOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.Info).To(Equal(core.BoolPtr(false)))
	return testing_utilities.GetMockSuccessResponse()
}

type GetHealthErrorSender struct{}

func (f GetHealthErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for PostEventNotificationsIntegration
type PostEventNotificationsIntegrationMockSender struct{}

func (f PostEventNotificationsIntegrationMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.PostEventNotificationsIntegrationOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.InstanceCrn).To(Equal(core.StringPtr("crn of event notifications instance")))
	Expect(createdOptions.Description).To(Equal(core.StringPtr("A sample project source")))
	Expect(createdOptions.Name).To(Equal(core.StringPtr("Project name")))
	Expect(createdOptions.Enabled).To(Equal(core.BoolPtr(true)))
	Expect(createdOptions.Source).To(Equal(core.StringPtr("CRN of the project instance")))
	return testing_utilities.GetMockSuccessResponse()
}

type PostEventNotificationsIntegrationErrorSender struct{}

func (f PostEventNotificationsIntegrationErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}
// Fake senders for PostEventNotification
type PostEventNotificationMockSender struct{}

func (f PostEventNotificationMockSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	createdOptions, ok := optionsModel.(*projectsv1.PostEventNotificationOptions)
	Expect(ok).To(Equal(true))
	Expect(createdOptions.ID).To(Equal(core.StringPtr("testString")))
	Expect(createdOptions.NewID).To(Equal(core.StringPtr("5f208fef-6b64-413c-aa07-dfed0b46abc1236")))
	Expect(createdOptions.NewSource).To(Equal(core.StringPtr("crn of project")))
	Expect(createdOptions.NewDatacontenttype).To(Equal(core.StringPtr("application/json")))
	Expect(createdOptions.NewIbmendefaultlong).To(Equal(core.StringPtr("long test notification message")))
	Expect(createdOptions.NewIbmendefaultshort).To(Equal(core.StringPtr("Test notification")))
	Expect(createdOptions.NewIbmensourceid).To(Equal(core.StringPtr("crn of project")))
	Expect(createdOptions.NewSpecversion).To(Equal(core.StringPtr("1.0")))
	Expect(createdOptions.NewType).To(Equal(core.StringPtr("com.ibm.cloud.project.project.test_notification")))
	return testing_utilities.GetMockSuccessResponse()
}

type PostEventNotificationErrorSender struct{}

func (f PostEventNotificationErrorSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return testing_utilities.GetMockErrorResponse()
}

//
// Utility functions used by the generated test code
//

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	bin, err := base64.StdEncoding.DecodeString(mockData)

	Expect(err).To(BeNil())

	return &bin
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

// convert struct instance to a generic map with resolved pointers, etc. for comparison
func ResolveModel(model interface{}) interface{} {
	buf, e := json.Marshal(model)
	if e != nil {
		panic(e)
	}

	var data interface{}

	e = json.Unmarshal(buf, &data)
	if e != nil {
		panic(e)
	}

	return data
}

func CheckAnalyticsHeader(serviceInstance *projectsv1.ProjectsV1) {
	header := serviceInstance.Service.DefaultHeaders.Get("X-Original-User-Agent")
	Expect(header).To(Equal("github.com/IBM/platform-services-go-sdk/"+version.GetPluginVersion().String()))
}
