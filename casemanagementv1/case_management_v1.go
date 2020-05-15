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

// Package casemanagementv1 : Operations and models for the CaseManagementV1 service
package casemanagementv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"io"
	"reflect"
	"strings"
)

// CaseManagementV1 : Case management API for creating cases, getting case statuses, adding comments to a case, adding
// and removing users from a case watchlist, downloading and adding attachments, and more.
//
// Version: 1.0.0
type CaseManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://support-center.cloud.ibm.com/case-management/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "case_management"

// CaseManagementV1Options : Service options
type CaseManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewCaseManagementV1UsingExternalConfig : constructs an instance of CaseManagementV1 with passed in options and external configuration.
func NewCaseManagementV1UsingExternalConfig(options *CaseManagementV1Options) (caseManagement *CaseManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	caseManagement, err = NewCaseManagementV1(options)
	if err != nil {
		return
	}

	err = caseManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = caseManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewCaseManagementV1 : constructs an instance of CaseManagementV1 with passed in options.
func NewCaseManagementV1(options *CaseManagementV1Options) (service *CaseManagementV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &CaseManagementV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (caseManagement *CaseManagementV1) SetServiceURL(url string) error {
	return caseManagement.Service.SetServiceURL(url)
}

// GetCases : Get cases in account
// Get cases in the account which is specified by the content of the IAM token.
func (caseManagement *CaseManagementV1) GetCases(getCasesOptions *GetCasesOptions) (result *CaseList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCasesOptions, "getCasesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCasesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "GetCases")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getCasesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getCasesOptions.Offset))
	}
	if getCasesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getCasesOptions.Limit))
	}
	if getCasesOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*getCasesOptions.Search))
	}
	if getCasesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*getCasesOptions.Sort))
	}
	if getCasesOptions.Status != nil {
		builder.AddQuery("status", strings.Join(getCasesOptions.Status, ","))
	}
	if getCasesOptions.Fields != nil {
		builder.AddQuery("fields", strings.Join(getCasesOptions.Fields, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCaseList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateCase : Create a case
// Create a case in the account.
func (caseManagement *CaseManagementV1) CreateCase(createCaseOptions *CreateCaseOptions) (result *Case, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCaseOptions, "createCaseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCaseOptions, "createCaseOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "CreateCase")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createCaseOptions.Type != nil {
		body["type"] = createCaseOptions.Type
	}
	if createCaseOptions.Subject != nil {
		body["subject"] = createCaseOptions.Subject
	}
	if createCaseOptions.Description != nil {
		body["description"] = createCaseOptions.Description
	}
	if createCaseOptions.Severity != nil {
		body["severity"] = createCaseOptions.Severity
	}
	if createCaseOptions.Eu != nil {
		body["eu"] = createCaseOptions.Eu
	}
	if createCaseOptions.Offering != nil {
		body["offering"] = createCaseOptions.Offering
	}
	if createCaseOptions.Resources != nil {
		body["resources"] = createCaseOptions.Resources
	}
	if createCaseOptions.Watchlist != nil {
		body["watchlist"] = createCaseOptions.Watchlist
	}
	if createCaseOptions.InvoiceNumber != nil {
		body["invoice_number"] = createCaseOptions.InvoiceNumber
	}
	if createCaseOptions.SlaCreditRequest != nil {
		body["sla_credit_request"] = createCaseOptions.SlaCreditRequest
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCase)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetCase : Get a case in account
// Get a case in the account that is specified by the case number.
func (caseManagement *CaseManagementV1) GetCase(getCaseOptions *GetCaseOptions) (result *Case, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCaseOptions, "getCaseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCaseOptions, "getCaseOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases"}
	pathParameters := []string{*getCaseOptions.CaseNumber}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "GetCase")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getCaseOptions.Fields != nil {
		builder.AddQuery("fields", strings.Join(getCaseOptions.Fields, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCase)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateCaseStatus : Update case status
// Mark the case as resolved or unresolved, or accept the provided resolution.
func (caseManagement *CaseManagementV1) UpdateCaseStatus(updateCaseStatusOptions *UpdateCaseStatusOptions) (result *Case, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCaseStatusOptions, "updateCaseStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCaseStatusOptions, "updateCaseStatusOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases", "status"}
	pathParameters := []string{*updateCaseStatusOptions.CaseNumber}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCaseStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "UpdateCaseStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(updateCaseStatusOptions.StatusPayload)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCase)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AddComment : Add comment to case
// Add a comment to a case.
func (caseManagement *CaseManagementV1) AddComment(addCommentOptions *AddCommentOptions) (result *Comment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addCommentOptions, "addCommentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addCommentOptions, "addCommentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases", "comments"}
	pathParameters := []string{*addCommentOptions.CaseNumber}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range addCommentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "AddComment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addCommentOptions.Comment != nil {
		body["comment"] = addCommentOptions.Comment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalComment)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AddWatchlist : Add users to watchlist of case
// Add users to the watchlist of case. By adding a user to the watchlist of the case, you are granting them read and
// write permissions, so the user can view the case, receive updates, and make updates to the case. Note that the user
// must be in the account to be added to the watchlist.
func (caseManagement *CaseManagementV1) AddWatchlist(addWatchlistOptions *AddWatchlistOptions) (result *WatchlistAddResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addWatchlistOptions, "addWatchlistOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addWatchlistOptions, "addWatchlistOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases", "watchlist"}
	pathParameters := []string{*addWatchlistOptions.CaseNumber}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range addWatchlistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "AddWatchlist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addWatchlistOptions.Watchlist != nil {
		body["watchlist"] = addWatchlistOptions.Watchlist
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWatchlistAddResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// RemoveWatchlist : Remove users from watchlist of case
// Remove users from the watchlist of a case.
func (caseManagement *CaseManagementV1) RemoveWatchlist(removeWatchlistOptions *RemoveWatchlistOptions) (result *Watchlist, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeWatchlistOptions, "removeWatchlistOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeWatchlistOptions, "removeWatchlistOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases", "watchlist"}
	pathParameters := []string{*removeWatchlistOptions.CaseNumber}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeWatchlistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "RemoveWatchlist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if removeWatchlistOptions.Watchlist != nil {
		body["watchlist"] = removeWatchlistOptions.Watchlist
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWatchlist)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AddResource : Add a resource to case
// Add a resource to case by specifying the Cloud Resource Name (CRN), or id and type if attaching a class iaaS
// resource.
func (caseManagement *CaseManagementV1) AddResource(addResourceOptions *AddResourceOptions) (result *Resource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addResourceOptions, "addResourceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addResourceOptions, "addResourceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases", "resources"}
	pathParameters := []string{*addResourceOptions.CaseNumber}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range addResourceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "AddResource")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addResourceOptions.Crn != nil {
		body["crn"] = addResourceOptions.Crn
	}
	if addResourceOptions.Type != nil {
		body["type"] = addResourceOptions.Type
	}
	if addResourceOptions.ID != nil {
		body["id"] = addResourceOptions.ID
	}
	if addResourceOptions.Note != nil {
		body["note"] = addResourceOptions.Note
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResource)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UploadFile : Add attachment(s) to case
// You can add attachments to a case to provide more information for the support team about the issue that you're
// experiencing.
func (caseManagement *CaseManagementV1) UploadFile(uploadFileOptions *UploadFileOptions) (result *Attachment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(uploadFileOptions, "uploadFileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(uploadFileOptions, "uploadFileOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases", "attachments"}
	pathParameters := []string{*uploadFileOptions.CaseNumber}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range uploadFileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "UploadFile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	for _, item := range uploadFileOptions.File {
		builder.AddFormData("file", core.StringNilMapper(item.Filename), core.StringNilMapper(item.ContentType), item.Data)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachment)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DownloadFile : Download an attachment
// Download an attachment from a case.
func (caseManagement *CaseManagementV1) DownloadFile(downloadFileOptions *DownloadFileOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(downloadFileOptions, "downloadFileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(downloadFileOptions, "downloadFileOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases", "attachments"}
	pathParameters := []string{*downloadFileOptions.CaseNumber, *downloadFileOptions.FileID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range downloadFileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "DownloadFile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/octet-stream")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = caseManagement.Service.Request(request, &result)

	return
}

// DeleteFile : Remove attachment from case
// Remove an attachment from a case.
func (caseManagement *CaseManagementV1) DeleteFile(deleteFileOptions *DeleteFileOptions) (result *AttachmentList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteFileOptions, "deleteFileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteFileOptions, "deleteFileOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"cases", "attachments"}
	pathParameters := []string{*deleteFileOptions.CaseNumber, *deleteFileOptions.FileID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(caseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteFileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "DeleteFile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachmentList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AddCommentOptions : The AddComment options.
type AddCommentOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required"`

	// Comment to add to the case.
	Comment *string `json:"comment" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddCommentOptions : Instantiate AddCommentOptions
func (*CaseManagementV1) NewAddCommentOptions(caseNumber string, comment string) *AddCommentOptions {
	return &AddCommentOptions{
		CaseNumber: core.StringPtr(caseNumber),
		Comment: core.StringPtr(comment),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (options *AddCommentOptions) SetCaseNumber(caseNumber string) *AddCommentOptions {
	options.CaseNumber = core.StringPtr(caseNumber)
	return options
}

// SetComment : Allow user to set Comment
func (options *AddCommentOptions) SetComment(comment string) *AddCommentOptions {
	options.Comment = core.StringPtr(comment)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddCommentOptions) SetHeaders(param map[string]string) *AddCommentOptions {
	options.Headers = param
	return options
}

// AddResourceOptions : The AddResource options.
type AddResourceOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required"`

	// Cloud Resource Name of the resource.
	Crn *string `json:"crn,omitempty"`

	// Only used to attach Classic IaaS devices which have no CRN.
	Type *string `json:"type,omitempty"`

	// Only used to attach Classic IaaS devices which have no CRN. Id of Classic IaaS device. This is deprecated in favor
	// of the crn field.
	ID *float64 `json:"id,omitempty"`

	// A note about this resource.
	Note *string `json:"note,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddResourceOptions : Instantiate AddResourceOptions
func (*CaseManagementV1) NewAddResourceOptions(caseNumber string) *AddResourceOptions {
	return &AddResourceOptions{
		CaseNumber: core.StringPtr(caseNumber),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (options *AddResourceOptions) SetCaseNumber(caseNumber string) *AddResourceOptions {
	options.CaseNumber = core.StringPtr(caseNumber)
	return options
}

// SetCrn : Allow user to set Crn
func (options *AddResourceOptions) SetCrn(crn string) *AddResourceOptions {
	options.Crn = core.StringPtr(crn)
	return options
}

// SetType : Allow user to set Type
func (options *AddResourceOptions) SetType(typeVar string) *AddResourceOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetID : Allow user to set ID
func (options *AddResourceOptions) SetID(id float64) *AddResourceOptions {
	options.ID = core.Float64Ptr(id)
	return options
}

// SetNote : Allow user to set Note
func (options *AddResourceOptions) SetNote(note string) *AddResourceOptions {
	options.Note = core.StringPtr(note)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddResourceOptions) SetHeaders(param map[string]string) *AddResourceOptions {
	options.Headers = param
	return options
}

// AddWatchlistOptions : The AddWatchlist options.
type AddWatchlistOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required"`

	// Array of user ID objects.
	Watchlist []User `json:"watchlist,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddWatchlistOptions : Instantiate AddWatchlistOptions
func (*CaseManagementV1) NewAddWatchlistOptions(caseNumber string) *AddWatchlistOptions {
	return &AddWatchlistOptions{
		CaseNumber: core.StringPtr(caseNumber),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (options *AddWatchlistOptions) SetCaseNumber(caseNumber string) *AddWatchlistOptions {
	options.CaseNumber = core.StringPtr(caseNumber)
	return options
}

// SetWatchlist : Allow user to set Watchlist
func (options *AddWatchlistOptions) SetWatchlist(watchlist []User) *AddWatchlistOptions {
	options.Watchlist = watchlist
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddWatchlistOptions) SetHeaders(param map[string]string) *AddWatchlistOptions {
	options.Headers = param
	return options
}

// Attachment : Details of an attachment.
type Attachment struct {
	// Unique identifier of the attachment in database.
	ID *string `json:"id,omitempty"`

	// Name of the attachment.
	Filename *string `json:"filename,omitempty"`

	// Size of the attachment in bytes.
	SizeInBytes *int64 `json:"size_in_bytes,omitempty"`

	// Date time of uploading.
	CreatedAt *string `json:"created_at,omitempty"`

	// URL of the attachment used to download.
	URL *string `json:"url,omitempty"`
}


// UnmarshalAttachment unmarshals an instance of Attachment from the specified map of raw messages.
func UnmarshalAttachment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Attachment)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "filename", &obj.Filename)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size_in_bytes", &obj.SizeInBytes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentList : List of attachments in the case.
type AttachmentList struct {
	// New attachments array.
	Attachments []Attachment `json:"attachments,omitempty"`
}


// UnmarshalAttachmentList unmarshals an instance of AttachmentList from the specified map of raw messages.
func UnmarshalAttachmentList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentList)
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Case : The support case.
type Case struct {
	// Number/ID of the case.
	Number *string `json:"number,omitempty"`

	// A short description of what the case is about.
	ShortDescription *string `json:"short_description,omitempty"`

	// A full description of what the case is about.
	Description *string `json:"description,omitempty"`

	// Date time of case creation in UTC.
	CreatedAt *string `json:"created_at,omitempty"`

	// User info in a case.
	CreatedBy *User `json:"created_by,omitempty"`

	// Date time of the last update on the case in UTC.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// User info in a case.
	UpdatedBy *User `json:"updated_by,omitempty"`

	// Name of the console to interact with the contact.
	ContactType *string `json:"contact_type,omitempty"`

	// User info in a case.
	Contact *User `json:"contact,omitempty"`

	// Status of the case.
	Status *string `json:"status,omitempty"`

	// The severity of the case.
	Severity *float64 `json:"severity,omitempty"`

	// Support tier of the account.
	SupportTier *string `json:"support_tier,omitempty"`

	// Standard reasons of resolving case.
	Resolution *string `json:"resolution,omitempty"`

	// Notes of case closing.
	CloseNotes *string `json:"close_notes,omitempty"`

	// EU support.
	Eu *CaseEu `json:"eu,omitempty"`

	Watchlist []User `json:"watchlist,omitempty"`

	// List of attachments/files of the case.
	Attachments []Attachment `json:"attachments,omitempty"`

	// Offering details.
	Offering *Offering `json:"offering,omitempty"`

	// List of attached resources.
	Resources []Resource `json:"resources,omitempty"`

	// List of comments/updates sorted in chronological order.
	Comments []Comment `json:"comments,omitempty"`
}

// Constants associated with the Case.ContactType property.
// Name of the console to interact with the contact.
const (
	Case_ContactType_CloudSupportCenter = "Cloud Support Center"
	Case_ContactType_ImsConsole = "IMS Console"
)

// Constants associated with the Case.SupportTier property.
// Support tier of the account.
const (
	Case_SupportTier_Basic = "Basic"
	Case_SupportTier_Free = "Free"
	Case_SupportTier_Premium = "Premium"
	Case_SupportTier_Standard = "Standard"
)


// UnmarshalCase unmarshals an instance of Case from the specified map of raw messages.
func UnmarshalCase(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Case)
	err = core.UnmarshalPrimitive(m, "number", &obj.Number)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "short_description", &obj.ShortDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "created_by", &obj.CreatedBy, UnmarshalUser)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "updated_by", &obj.UpdatedBy, UnmarshalUser)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "contact_type", &obj.ContactType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contact", &obj.Contact, UnmarshalUser)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "severity", &obj.Severity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "support_tier", &obj.SupportTier)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resolution", &obj.Resolution)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "close_notes", &obj.CloseNotes)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "eu", &obj.Eu, UnmarshalCaseEu)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "watchlist", &obj.Watchlist, UnmarshalUser)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "offering", &obj.Offering, UnmarshalOffering)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "comments", &obj.Comments, UnmarshalComment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CaseEu : EU support.
type CaseEu struct {
	// Identifying whether the case has EU Support.
	Support *bool `json:"support,omitempty"`

	// Information about the data center.
	DataCenter *string `json:"data_center,omitempty"`
}


// UnmarshalCaseEu unmarshals an instance of CaseEu from the specified map of raw messages.
func UnmarshalCaseEu(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CaseEu)
	err = core.UnmarshalPrimitive(m, "support", &obj.Support)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_center", &obj.DataCenter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CaseList : Response of a GET /cases request.
type CaseList struct {
	// Total number of cases satisfying the query.
	TotalCount *int64 `json:"total_count,omitempty"`

	// URL to related pages of cases.
	First *PaginationLink `json:"first,omitempty"`

	// URL to related pages of cases.
	Next *PaginationLink `json:"next,omitempty"`

	// URL to related pages of cases.
	Previous *PaginationLink `json:"previous,omitempty"`

	// URL to related pages of cases.
	Last *PaginationLink `json:"last,omitempty"`

	// List of cases.
	Cases []Case `json:"cases,omitempty"`
}


// UnmarshalCaseList unmarshals an instance of CaseList from the specified map of raw messages.
func UnmarshalCaseList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CaseList)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cases", &obj.Cases, UnmarshalCase)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CasePayloadEu : Specify if the case should be treated as EU regulated. Only one of the following properties is required. Call EU
// support utility endpoint to determine which property must be specified for your account.
type CasePayloadEu struct {
	// indicating whether the case is EU supported.
	Supported *bool `json:"supported,omitempty"`

	// If EU supported utility endpoint specifies datacenter then pass the datacenter id to mark a case as EU supported.
	DataCenter *int64 `json:"data_center,omitempty"`
}


// UnmarshalCasePayloadEu unmarshals an instance of CasePayloadEu from the specified map of raw messages.
func UnmarshalCasePayloadEu(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CasePayloadEu)
	err = core.UnmarshalPrimitive(m, "supported", &obj.Supported)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_center", &obj.DataCenter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Comment : A comment in a case.
type Comment struct {
	// The comment.
	Value *string `json:"value,omitempty"`

	// Timestamp of when comment is added.
	AddedAt *string `json:"added_at,omitempty"`

	// User info in a case.
	AddedBy *User `json:"added_by,omitempty"`
}


// UnmarshalComment unmarshals an instance of Comment from the specified map of raw messages.
func UnmarshalComment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Comment)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "added_at", &obj.AddedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "added_by", &obj.AddedBy, UnmarshalUser)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCaseOptions : The CreateCase options.
type CreateCaseOptions struct {
	Type *string `json:"type" validate:"required"`

	// Subject of the case.
	Subject *string `json:"subject" validate:"required"`

	// Detailed description of the issue.
	Description *string `json:"description" validate:"required"`

	// Severity of the case. Smaller values mean higher severity.
	Severity *int64 `json:"severity,omitempty"`

	// Specify if the case should be treated as EU regulated. Only one of the following properties is required. Call EU
	// support utility endpoint to determine which property must be specified for your account.
	Eu *CasePayloadEu `json:"eu,omitempty"`

	// Payload to specify the offering of a case.
	Offering *OfferingPayload `json:"offering,omitempty"`

	// List of resources to attach to case. If attaching Classic IaaS devices use type and id fields if Cloud Resource Name
	// (CRN) is unavialable. Otherwise pass the resource CRN. The resource list must be consistent with the value selected
	// for the resource offering.
	Resources []ResourcePayload `json:"resources,omitempty"`

	// Array of user IDs to add to the watchlist.
	Watchlist []UserIdAndRealm `json:"watchlist,omitempty"`

	// Invoice number of "Billing and Invoice" case type.
	InvoiceNumber *string `json:"invoice_number,omitempty"`

	// Flag to indicate if case is for an Service Level Agreement (SLA) credit request.
	SlaCreditRequest *bool `json:"sla_credit_request,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCaseOptions.Type property.
const (
	CreateCaseOptions_Type_AccountAndAccess = "account_and_access"
	CreateCaseOptions_Type_BillingAndInvoice = "billing_and_invoice"
	CreateCaseOptions_Type_Sales = "sales"
	CreateCaseOptions_Type_Technical = "technical"
)

// NewCreateCaseOptions : Instantiate CreateCaseOptions
func (*CaseManagementV1) NewCreateCaseOptions(typeVar string, subject string, description string) *CreateCaseOptions {
	return &CreateCaseOptions{
		Type: core.StringPtr(typeVar),
		Subject: core.StringPtr(subject),
		Description: core.StringPtr(description),
	}
}

// SetType : Allow user to set Type
func (options *CreateCaseOptions) SetType(typeVar string) *CreateCaseOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetSubject : Allow user to set Subject
func (options *CreateCaseOptions) SetSubject(subject string) *CreateCaseOptions {
	options.Subject = core.StringPtr(subject)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateCaseOptions) SetDescription(description string) *CreateCaseOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetSeverity : Allow user to set Severity
func (options *CreateCaseOptions) SetSeverity(severity int64) *CreateCaseOptions {
	options.Severity = core.Int64Ptr(severity)
	return options
}

// SetEu : Allow user to set Eu
func (options *CreateCaseOptions) SetEu(eu *CasePayloadEu) *CreateCaseOptions {
	options.Eu = eu
	return options
}

// SetOffering : Allow user to set Offering
func (options *CreateCaseOptions) SetOffering(offering *OfferingPayload) *CreateCaseOptions {
	options.Offering = offering
	return options
}

// SetResources : Allow user to set Resources
func (options *CreateCaseOptions) SetResources(resources []ResourcePayload) *CreateCaseOptions {
	options.Resources = resources
	return options
}

// SetWatchlist : Allow user to set Watchlist
func (options *CreateCaseOptions) SetWatchlist(watchlist []UserIdAndRealm) *CreateCaseOptions {
	options.Watchlist = watchlist
	return options
}

// SetInvoiceNumber : Allow user to set InvoiceNumber
func (options *CreateCaseOptions) SetInvoiceNumber(invoiceNumber string) *CreateCaseOptions {
	options.InvoiceNumber = core.StringPtr(invoiceNumber)
	return options
}

// SetSlaCreditRequest : Allow user to set SlaCreditRequest
func (options *CreateCaseOptions) SetSlaCreditRequest(slaCreditRequest bool) *CreateCaseOptions {
	options.SlaCreditRequest = core.BoolPtr(slaCreditRequest)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCaseOptions) SetHeaders(param map[string]string) *CreateCaseOptions {
	options.Headers = param
	return options
}

// DeleteFileOptions : The DeleteFile options.
type DeleteFileOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required"`

	// Unique identifier of a file.
	FileID *string `json:"file_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteFileOptions : Instantiate DeleteFileOptions
func (*CaseManagementV1) NewDeleteFileOptions(caseNumber string, fileID string) *DeleteFileOptions {
	return &DeleteFileOptions{
		CaseNumber: core.StringPtr(caseNumber),
		FileID: core.StringPtr(fileID),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (options *DeleteFileOptions) SetCaseNumber(caseNumber string) *DeleteFileOptions {
	options.CaseNumber = core.StringPtr(caseNumber)
	return options
}

// SetFileID : Allow user to set FileID
func (options *DeleteFileOptions) SetFileID(fileID string) *DeleteFileOptions {
	options.FileID = core.StringPtr(fileID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteFileOptions) SetHeaders(param map[string]string) *DeleteFileOptions {
	options.Headers = param
	return options
}

// DownloadFileOptions : The DownloadFile options.
type DownloadFileOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required"`

	// Unique identifier of a file.
	FileID *string `json:"file_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDownloadFileOptions : Instantiate DownloadFileOptions
func (*CaseManagementV1) NewDownloadFileOptions(caseNumber string, fileID string) *DownloadFileOptions {
	return &DownloadFileOptions{
		CaseNumber: core.StringPtr(caseNumber),
		FileID: core.StringPtr(fileID),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (options *DownloadFileOptions) SetCaseNumber(caseNumber string) *DownloadFileOptions {
	options.CaseNumber = core.StringPtr(caseNumber)
	return options
}

// SetFileID : Allow user to set FileID
func (options *DownloadFileOptions) SetFileID(fileID string) *DownloadFileOptions {
	options.FileID = core.StringPtr(fileID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DownloadFileOptions) SetHeaders(param map[string]string) *DownloadFileOptions {
	options.Headers = param
	return options
}

// FileWithMetadata : A file with its associated metadata.
type FileWithMetadata struct {
	// The data / content for the file.
	Data io.ReadCloser `json:"data" validate:"required"`

	// The filename of the file.
	Filename *string `json:"filename,omitempty"`

	// The content type of the file.
	ContentType *string `json:"content_type,omitempty"`
}


// NewFileWithMetadata : Instantiate FileWithMetadata (Generic Model Constructor)
func (*CaseManagementV1) NewFileWithMetadata(data io.ReadCloser) (model *FileWithMetadata, err error) {
	model = &FileWithMetadata{
		Data: data,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalFileWithMetadata unmarshals an instance of FileWithMetadata from the specified map of raw messages.
func UnmarshalFileWithMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FileWithMetadata)
	err = core.UnmarshalPrimitive(m, "data", &obj.Data)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "filename", &obj.Filename)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "content_type", &obj.ContentType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetCaseOptions : The GetCase options.
type GetCaseOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required"`

	// Seleted fields of interest instead of the entire case information.
	Fields []string `json:"fields,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetCaseOptions.Fields property.
const (
	GetCaseOptions_Fields_AgentCloseOnly = "agent_close_only"
	GetCaseOptions_Fields_Attachments = "attachments"
	GetCaseOptions_Fields_CloseNotes = "close_notes"
	GetCaseOptions_Fields_Comments = "comments"
	GetCaseOptions_Fields_Contact = "contact"
	GetCaseOptions_Fields_ContactType = "contact_type"
	GetCaseOptions_Fields_CreatedAt = "created_at"
	GetCaseOptions_Fields_CreatedBy = "created_by"
	GetCaseOptions_Fields_Description = "description"
	GetCaseOptions_Fields_Eu = "eu"
	GetCaseOptions_Fields_InvoiceNumber = "invoice_number"
	GetCaseOptions_Fields_Number = "number"
	GetCaseOptions_Fields_Offering = "offering"
	GetCaseOptions_Fields_Resolution = "resolution"
	GetCaseOptions_Fields_Resources = "resources"
	GetCaseOptions_Fields_Severity = "severity"
	GetCaseOptions_Fields_ShortDescription = "short_description"
	GetCaseOptions_Fields_Status = "status"
	GetCaseOptions_Fields_SupportTier = "support_tier"
	GetCaseOptions_Fields_UpdatedAt = "updated_at"
	GetCaseOptions_Fields_UpdatedBy = "updated_by"
	GetCaseOptions_Fields_Watchlist = "watchlist"
)

// NewGetCaseOptions : Instantiate GetCaseOptions
func (*CaseManagementV1) NewGetCaseOptions(caseNumber string) *GetCaseOptions {
	return &GetCaseOptions{
		CaseNumber: core.StringPtr(caseNumber),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (options *GetCaseOptions) SetCaseNumber(caseNumber string) *GetCaseOptions {
	options.CaseNumber = core.StringPtr(caseNumber)
	return options
}

// SetFields : Allow user to set Fields
func (options *GetCaseOptions) SetFields(fields []string) *GetCaseOptions {
	options.Fields = fields
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCaseOptions) SetHeaders(param map[string]string) *GetCaseOptions {
	options.Headers = param
	return options
}

// GetCasesOptions : The GetCases options.
type GetCasesOptions struct {
	// Number of cases should be skipped.
	Offset *int64 `json:"offset,omitempty"`

	// Number of cases should be returned.
	Limit *int64 `json:"limit,omitempty"`

	// String that a case might contain.
	Search *string `json:"search,omitempty"`

	// Sort field and direction. If omitted, default to descending of updated date. Prefix "~" signifies sort in
	// descending.
	Sort *string `json:"sort,omitempty"`

	// Case status filter.
	Status []string `json:"status,omitempty"`

	// Seleted fields of interest instead of the entire case information.
	Fields []string `json:"fields,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetCasesOptions.Status property.
const (
	GetCasesOptions_Status_Closed = "closed"
	GetCasesOptions_Status_InProgress = "in_progress"
	GetCasesOptions_Status_New = "new"
	GetCasesOptions_Status_ResolutionProvided = "resolution_provided"
	GetCasesOptions_Status_Resolved = "resolved"
	GetCasesOptions_Status_WaitingOnClient = "waiting_on_client"
)

// Constants associated with the GetCasesOptions.Fields property.
const (
	GetCasesOptions_Fields_AgentCloseOnly = "agent_close_only"
	GetCasesOptions_Fields_Attachments = "attachments"
	GetCasesOptions_Fields_CloseNotes = "close_notes"
	GetCasesOptions_Fields_Comments = "comments"
	GetCasesOptions_Fields_Contact = "contact"
	GetCasesOptions_Fields_ContactType = "contact_type"
	GetCasesOptions_Fields_CreatedAt = "created_at"
	GetCasesOptions_Fields_CreatedBy = "created_by"
	GetCasesOptions_Fields_Description = "description"
	GetCasesOptions_Fields_Eu = "eu"
	GetCasesOptions_Fields_InvoiceNumber = "invoice_number"
	GetCasesOptions_Fields_Number = "number"
	GetCasesOptions_Fields_Offering = "offering"
	GetCasesOptions_Fields_Resolution = "resolution"
	GetCasesOptions_Fields_Resources = "resources"
	GetCasesOptions_Fields_Severity = "severity"
	GetCasesOptions_Fields_ShortDescription = "short_description"
	GetCasesOptions_Fields_Status = "status"
	GetCasesOptions_Fields_SupportTier = "support_tier"
	GetCasesOptions_Fields_UpdatedAt = "updated_at"
	GetCasesOptions_Fields_UpdatedBy = "updated_by"
	GetCasesOptions_Fields_Watchlist = "watchlist"
)

// NewGetCasesOptions : Instantiate GetCasesOptions
func (*CaseManagementV1) NewGetCasesOptions() *GetCasesOptions {
	return &GetCasesOptions{}
}

// SetOffset : Allow user to set Offset
func (options *GetCasesOptions) SetOffset(offset int64) *GetCasesOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetCasesOptions) SetLimit(limit int64) *GetCasesOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetSearch : Allow user to set Search
func (options *GetCasesOptions) SetSearch(search string) *GetCasesOptions {
	options.Search = core.StringPtr(search)
	return options
}

// SetSort : Allow user to set Sort
func (options *GetCasesOptions) SetSort(sort string) *GetCasesOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetStatus : Allow user to set Status
func (options *GetCasesOptions) SetStatus(status []string) *GetCasesOptions {
	options.Status = status
	return options
}

// SetFields : Allow user to set Fields
func (options *GetCasesOptions) SetFields(fields []string) *GetCasesOptions {
	options.Fields = fields
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCasesOptions) SetHeaders(param map[string]string) *GetCasesOptions {
	options.Headers = param
	return options
}

// Offering : Offering details.
type Offering struct {
	// Name of the offering.
	Name *string `json:"name" validate:"required"`

	Type *OfferingType `json:"type" validate:"required"`
}


// UnmarshalOffering unmarshals an instance of Offering from the specified map of raw messages.
func UnmarshalOffering(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Offering)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "type", &obj.Type, UnmarshalOfferingType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OfferingPayload : Payload to specify the offering of a case.
type OfferingPayload struct {
	// Offering name.
	Name *string `json:"name" validate:"required"`

	// Offering type.
	Type *OfferingPayloadType `json:"type" validate:"required"`
}


// NewOfferingPayload : Instantiate OfferingPayload (Generic Model Constructor)
func (*CaseManagementV1) NewOfferingPayload(name string, typeVar *OfferingPayloadType) (model *OfferingPayload, err error) {
	model = &OfferingPayload{
		Name: core.StringPtr(name),
		Type: typeVar,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalOfferingPayload unmarshals an instance of OfferingPayload from the specified map of raw messages.
func UnmarshalOfferingPayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OfferingPayload)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "type", &obj.Type, UnmarshalOfferingPayloadType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OfferingPayloadType : Offering type.
type OfferingPayloadType struct {
	// Offering type group. "crn_service_name" is strongly prefered over "category" as the latter is legacy and will be
	// deprecated in the future.
	Group *string `json:"group" validate:"required"`

	// CRN service name of the offering.
	Key *string `json:"key" validate:"required"`

	// Optional. Platform kind of the offering.
	Kind *string `json:"kind,omitempty"`

	// Offering id in the catalog. This alone is enough to identify the offering.
	ID *string `json:"id,omitempty"`
}

// Constants associated with the OfferingPayloadType.Group property.
// Offering type group. "crn_service_name" is strongly prefered over "category" as the latter is legacy and will be
// deprecated in the future.
const (
	OfferingPayloadType_Group_Category = "category"
	OfferingPayloadType_Group_CrnServiceName = "crn_service_name"
)


// NewOfferingPayloadType : Instantiate OfferingPayloadType (Generic Model Constructor)
func (*CaseManagementV1) NewOfferingPayloadType(group string, key string) (model *OfferingPayloadType, err error) {
	model = &OfferingPayloadType{
		Group: core.StringPtr(group),
		Key: core.StringPtr(key),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalOfferingPayloadType unmarshals an instance of OfferingPayloadType from the specified map of raw messages.
func UnmarshalOfferingPayloadType(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OfferingPayloadType)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OfferingType : OfferingType struct
type OfferingType struct {
	// indicating whether this is an offering or a broad category.
	Group *string `json:"group" validate:"required"`

	// crn service name of the offering or the value of the category.
	Key *string `json:"key" validate:"required"`

	// catalog id of the offering.
	ID *string `json:"id,omitempty"`

	// kind of the offering.
	Kind *string `json:"kind,omitempty"`
}


// UnmarshalOfferingType unmarshals an instance of OfferingType from the specified map of raw messages.
func UnmarshalOfferingType(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OfferingType)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginationLink : URL to related pages of cases.
type PaginationLink struct {
	Href *string `json:"href,omitempty"`
}


// UnmarshalPaginationLink unmarshals an instance of PaginationLink from the specified map of raw messages.
func UnmarshalPaginationLink(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginationLink)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RemoveWatchlistOptions : The RemoveWatchlist options.
type RemoveWatchlistOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required"`

	// Array of user ID objects.
	Watchlist []User `json:"watchlist,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveWatchlistOptions : Instantiate RemoveWatchlistOptions
func (*CaseManagementV1) NewRemoveWatchlistOptions(caseNumber string) *RemoveWatchlistOptions {
	return &RemoveWatchlistOptions{
		CaseNumber: core.StringPtr(caseNumber),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (options *RemoveWatchlistOptions) SetCaseNumber(caseNumber string) *RemoveWatchlistOptions {
	options.CaseNumber = core.StringPtr(caseNumber)
	return options
}

// SetWatchlist : Allow user to set Watchlist
func (options *RemoveWatchlistOptions) SetWatchlist(watchlist []User) *RemoveWatchlistOptions {
	options.Watchlist = watchlist
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveWatchlistOptions) SetHeaders(param map[string]string) *RemoveWatchlistOptions {
	options.Headers = param
	return options
}

// Resource : A resource record of a case.
type Resource struct {
	// ID of the resource.
	Crn *string `json:"crn,omitempty"`

	// Name of the resource.
	Name *string `json:"name,omitempty"`

	// Type of resource.
	Type *string `json:"type,omitempty"`

	// URL of resource.
	URL *string `json:"url,omitempty"`

	// Note about resource.
	Note *string `json:"note,omitempty"`
}


// UnmarshalResource unmarshals an instance of Resource from the specified map of raw messages.
func UnmarshalResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resource)
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "note", &obj.Note)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourcePayload : Payload to add a resource to a case.
type ResourcePayload struct {
	// Cloud Resource Name of the resource.
	Crn *string `json:"crn,omitempty"`

	// Only used to attach Classic IaaS devices which have no CRN.
	Type *string `json:"type,omitempty"`

	// Only used to attach Classic IaaS devices which have no CRN. Id of Classic IaaS device. This is deprecated in favor
	// of the crn field.
	ID *float64 `json:"id,omitempty"`

	// A note about this resource.
	Note *string `json:"note,omitempty"`
}


// UnmarshalResourcePayload unmarshals an instance of ResourcePayload from the specified map of raw messages.
func UnmarshalResourcePayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourcePayload)
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "note", &obj.Note)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusPayload : Payload to update status of the case.
// Models which "extend" this model:
// - StatusPayloadResolvePayload
// - StatusPayloadUnresolvePayload
// - StatusPayloadAcceptPayload
type StatusPayload struct {
	// action to perform on the case.
	Action *string `json:"action" validate:"required"`

	// comment of resolution.
	Comment *string `json:"comment,omitempty"`

	// * 1: Client error
	// * 2: Defect found with Component/Service
	// * 3: Documentation Error
	// * 4: Sollution found in forums
	// * 5: Solution found in public Documentation
	// * 6: Solution no longer required
	// * 7: Solution provided by IBM outside of support case
	// * 8: Solution provided by IBM support engineer.
	ResolutionCode *int64 `json:"resolution_code,omitempty"`
}

// Constants associated with the StatusPayload.Action property.
// action to perform on the case.
const (
	StatusPayload_Action_Accept = "accept"
	StatusPayload_Action_Resolve = "resolve"
	StatusPayload_Action_Unresolve = "unresolve"
)

func (*StatusPayload) isaStatusPayload() bool {
	return true
}

type StatusPayloadIntf interface {
	isaStatusPayload() bool
}

// UnmarshalStatusPayload unmarshals an instance of StatusPayload from the specified map of raw messages.
func UnmarshalStatusPayload(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "action", &discValue)
	if err != nil {
		err = fmt.Errorf("error unmarshalling discriminator property 'action': %s", err.Error())
		return
	}
	if discValue == "" {
		err = fmt.Errorf("required discriminator property 'action' not found in JSON object")
		return
	}
	if discValue == "resolve" {
		err = core.UnmarshalModel(m, "", result, UnmarshalStatusPayloadResolvePayload)
	} else if discValue == "unresolve" {
		err = core.UnmarshalModel(m, "", result, UnmarshalStatusPayloadUnresolvePayload)
	} else if discValue == "accept" {
		err = core.UnmarshalModel(m, "", result, UnmarshalStatusPayloadAcceptPayload)
	} else {
		err = fmt.Errorf("unrecognized value for discriminator property 'action': %s", discValue)
	}
	return
}

// UpdateCaseStatusOptions : The UpdateCaseStatus options.
type UpdateCaseStatusOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required"`

	// Payload to update status of the case.
	StatusPayload StatusPayloadIntf `json:"StatusPayload" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCaseStatusOptions : Instantiate UpdateCaseStatusOptions
func (*CaseManagementV1) NewUpdateCaseStatusOptions(caseNumber string, statusPayload StatusPayloadIntf) *UpdateCaseStatusOptions {
	return &UpdateCaseStatusOptions{
		CaseNumber: core.StringPtr(caseNumber),
		StatusPayload: statusPayload,
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (options *UpdateCaseStatusOptions) SetCaseNumber(caseNumber string) *UpdateCaseStatusOptions {
	options.CaseNumber = core.StringPtr(caseNumber)
	return options
}

// SetStatusPayload : Allow user to set StatusPayload
func (options *UpdateCaseStatusOptions) SetStatusPayload(statusPayload StatusPayloadIntf) *UpdateCaseStatusOptions {
	options.StatusPayload = statusPayload
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCaseStatusOptions) SetHeaders(param map[string]string) *UpdateCaseStatusOptions {
	options.Headers = param
	return options
}

// UploadFileOptions : The UploadFile options.
type UploadFileOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required"`

	// file of supported types, 8MB in size limit.
	File []FileWithMetadata `json:"file" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUploadFileOptions : Instantiate UploadFileOptions
func (*CaseManagementV1) NewUploadFileOptions(caseNumber string, file []FileWithMetadata) *UploadFileOptions {
	return &UploadFileOptions{
		CaseNumber: core.StringPtr(caseNumber),
		File: file,
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (options *UploadFileOptions) SetCaseNumber(caseNumber string) *UploadFileOptions {
	options.CaseNumber = core.StringPtr(caseNumber)
	return options
}

// SetFile : Allow user to set File
func (options *UploadFileOptions) SetFile(file []FileWithMetadata) *UploadFileOptions {
	options.File = file
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UploadFileOptions) SetHeaders(param map[string]string) *UploadFileOptions {
	options.Headers = param
	return options
}

// User : User info in a case.
type User struct {
	// Full name of the user.
	Name *string `json:"name,omitempty"`

	// the ID realm.
	Realm *string `json:"realm" validate:"required"`

	// unique user ID in the realm specified by the type.
	UserID *string `json:"user_id" validate:"required"`
}

// Constants associated with the User.Realm property.
// the ID realm.
const (
	User_Realm_Bss = "BSS"
	User_Realm_Ibmid = "IBMid"
	User_Realm_Sl = "SL"
)


// NewUser : Instantiate User (Generic Model Constructor)
func (*CaseManagementV1) NewUser(realm string, userID string) (model *User, err error) {
	model = &User{
		Realm: core.StringPtr(realm),
		UserID: core.StringPtr(userID),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalUser unmarshals an instance of User from the specified map of raw messages.
func UnmarshalUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(User)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "realm", &obj.Realm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserIdAndRealm : user ID and realm.
type UserIdAndRealm struct {
	// the ID realm.
	Realm *string `json:"realm" validate:"required"`

	// unique user ID in the realm specified by the type.
	UserID *string `json:"user_id" validate:"required"`
}

// Constants associated with the UserIdAndRealm.Realm property.
// the ID realm.
const (
	UserIdAndRealm_Realm_Bss = "BSS"
	UserIdAndRealm_Realm_Ibmid = "IBMid"
	UserIdAndRealm_Realm_Sl = "SL"
)


// NewUserIdAndRealm : Instantiate UserIdAndRealm (Generic Model Constructor)
func (*CaseManagementV1) NewUserIdAndRealm(realm string, userID string) (model *UserIdAndRealm, err error) {
	model = &UserIdAndRealm{
		Realm: core.StringPtr(realm),
		UserID: core.StringPtr(userID),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalUserIdAndRealm unmarshals an instance of UserIdAndRealm from the specified map of raw messages.
func UnmarshalUserIdAndRealm(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserIdAndRealm)
	err = core.UnmarshalPrimitive(m, "realm", &obj.Realm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Watchlist : Payload to add/remove users to/from the case watchlist.
type Watchlist struct {
	// Array of user ID objects.
	Watchlist []User `json:"watchlist,omitempty"`
}


// UnmarshalWatchlist unmarshals an instance of Watchlist from the specified map of raw messages.
func UnmarshalWatchlist(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Watchlist)
	err = core.UnmarshalModel(m, "watchlist", &obj.Watchlist, UnmarshalUser)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WatchlistAddResponse : Response of a request adding to watchlist.
type WatchlistAddResponse struct {
	// List of added user.
	Added []User `json:"added,omitempty"`

	// List of failed to add user.
	Failed []User `json:"failed,omitempty"`
}


// UnmarshalWatchlistAddResponse unmarshals an instance of WatchlistAddResponse from the specified map of raw messages.
func UnmarshalWatchlistAddResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WatchlistAddResponse)
	err = core.UnmarshalModel(m, "added", &obj.Added, UnmarshalUser)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failed", &obj.Failed, UnmarshalUser)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusPayloadAcceptPayload : Payload to accept the proposed resolution of the case.
// This model "extends" StatusPayload
type StatusPayloadAcceptPayload struct {
	// action to perform on the case.
	Action *string `json:"action" validate:"required"`

	// Comment about accepting the proposed resolution.
	Comment *string `json:"comment,omitempty"`
}

// Constants associated with the StatusPayloadAcceptPayload.Action property.
// action to perform on the case.
const (
	StatusPayloadAcceptPayload_Action_Accept = "accept"
	StatusPayloadAcceptPayload_Action_Resolve = "resolve"
	StatusPayloadAcceptPayload_Action_Unresolve = "unresolve"
)


// NewStatusPayloadAcceptPayload : Instantiate StatusPayloadAcceptPayload (Generic Model Constructor)
func (*CaseManagementV1) NewStatusPayloadAcceptPayload(action string) (model *StatusPayloadAcceptPayload, err error) {
	model = &StatusPayloadAcceptPayload{
		Action: core.StringPtr(action),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*StatusPayloadAcceptPayload) isaStatusPayload() bool {
	return true
}

// UnmarshalStatusPayloadAcceptPayload unmarshals an instance of StatusPayloadAcceptPayload from the specified map of raw messages.
func UnmarshalStatusPayloadAcceptPayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StatusPayloadAcceptPayload)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusPayloadResolvePayload : Payload to resolve the case.
// This model "extends" StatusPayload
type StatusPayloadResolvePayload struct {
	// action to perform on the case.
	Action *string `json:"action" validate:"required"`

	// comment of resolution.
	Comment *string `json:"comment,omitempty"`

	// * 1: Client error
	// * 2: Defect found with Component/Service
	// * 3: Documentation Error
	// * 4: Sollution found in forums
	// * 5: Solution found in public Documentation
	// * 6: Solution no longer required
	// * 7: Solution provided by IBM outside of support case
	// * 8: Solution provided by IBM support engineer.
	ResolutionCode *int64 `json:"resolution_code" validate:"required"`
}

// Constants associated with the StatusPayloadResolvePayload.Action property.
// action to perform on the case.
const (
	StatusPayloadResolvePayload_Action_Accept = "accept"
	StatusPayloadResolvePayload_Action_Resolve = "resolve"
	StatusPayloadResolvePayload_Action_Unresolve = "unresolve"
)


// NewStatusPayloadResolvePayload : Instantiate StatusPayloadResolvePayload (Generic Model Constructor)
func (*CaseManagementV1) NewStatusPayloadResolvePayload(action string, resolutionCode int64) (model *StatusPayloadResolvePayload, err error) {
	model = &StatusPayloadResolvePayload{
		Action: core.StringPtr(action),
		ResolutionCode: core.Int64Ptr(resolutionCode),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*StatusPayloadResolvePayload) isaStatusPayload() bool {
	return true
}

// UnmarshalStatusPayloadResolvePayload unmarshals an instance of StatusPayloadResolvePayload from the specified map of raw messages.
func UnmarshalStatusPayloadResolvePayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StatusPayloadResolvePayload)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resolution_code", &obj.ResolutionCode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusPayloadUnresolvePayload : Payload to unresolve the case.
// This model "extends" StatusPayload
type StatusPayloadUnresolvePayload struct {
	// action to perform on the case.
	Action *string `json:"action" validate:"required"`

	// Comment why the case should be unresolved.
	Comment *string `json:"comment" validate:"required"`
}

// Constants associated with the StatusPayloadUnresolvePayload.Action property.
// action to perform on the case.
const (
	StatusPayloadUnresolvePayload_Action_Accept = "accept"
	StatusPayloadUnresolvePayload_Action_Resolve = "resolve"
	StatusPayloadUnresolvePayload_Action_Unresolve = "unresolve"
)


// NewStatusPayloadUnresolvePayload : Instantiate StatusPayloadUnresolvePayload (Generic Model Constructor)
func (*CaseManagementV1) NewStatusPayloadUnresolvePayload(action string, comment string) (model *StatusPayloadUnresolvePayload, err error) {
	model = &StatusPayloadUnresolvePayload{
		Action: core.StringPtr(action),
		Comment: core.StringPtr(comment),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*StatusPayloadUnresolvePayload) isaStatusPayload() bool {
	return true
}

// UnmarshalStatusPayloadUnresolvePayload unmarshals an instance of StatusPayloadUnresolvePayload from the specified map of raw messages.
func UnmarshalStatusPayloadUnresolvePayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StatusPayloadUnresolvePayload)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
