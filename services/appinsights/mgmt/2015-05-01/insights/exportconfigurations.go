package insights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ExportConfigurationsClient is the composite Swagger for Application Insights Management Client
type ExportConfigurationsClient struct {
	BaseClient
}

// NewExportConfigurationsClient creates an instance of the ExportConfigurationsClient client.
func NewExportConfigurationsClient(subscriptionID string, purgeID string) ExportConfigurationsClient {
	return NewExportConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID, purgeID)
}

// NewExportConfigurationsClientWithBaseURI creates an instance of the ExportConfigurationsClient client.
func NewExportConfigurationsClientWithBaseURI(baseURI string, subscriptionID string, purgeID string) ExportConfigurationsClient {
	return ExportConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID, purgeID)}
}

// Create create a Continuous Export configuration of an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. exportProperties is properties that need to be specified to create a Continuous Export
// configuration of a Application Insights component.
func (client ExportConfigurationsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, exportProperties ApplicationInsightsComponentExportRequest) (result ListApplicationInsightsComponentExportConfiguration, err error) {
	req, err := client.CreatePreparer(ctx, resourceGroupName, resourceName, exportProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ExportConfigurationsClient) CreatePreparer(ctx context.Context, resourceGroupName string, resourceName string, exportProperties ApplicationInsightsComponentExportRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration", pathParameters),
		autorest.WithJSON(exportProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ExportConfigurationsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ExportConfigurationsClient) CreateResponder(resp *http.Response) (result ListApplicationInsightsComponentExportConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a Continuous Export configuration of an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. exportID is the Continuous Export configuration ID. This is unique within a Application
// Insights component.
func (client ExportConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, exportID string) (result ApplicationInsightsComponentExportConfiguration, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, exportID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExportConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, exportID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"exportId":          autorest.Encode("path", exportID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExportConfigurationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExportConfigurationsClient) DeleteResponder(resp *http.Response) (result ApplicationInsightsComponentExportConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get the Continuous Export configuration for this export id.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. exportID is the Continuous Export configuration ID. This is unique within a Application
// Insights component.
func (client ExportConfigurationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, exportID string) (result ApplicationInsightsComponentExportConfiguration, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, exportID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExportConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, exportID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"exportId":          autorest.Encode("path", exportID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExportConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExportConfigurationsClient) GetResponder(resp *http.Response) (result ApplicationInsightsComponentExportConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of Continuous Export configuration of an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource.
func (client ExportConfigurationsClient) List(ctx context.Context, resourceGroupName string, resourceName string) (result ListApplicationInsightsComponentExportConfiguration, err error) {
	req, err := client.ListPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ExportConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExportConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExportConfigurationsClient) ListResponder(resp *http.Response) (result ListApplicationInsightsComponentExportConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update the Continuous Export configuration for this export id.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. exportID is the Continuous Export configuration ID. This is unique within a Application
// Insights component. exportProperties is properties that need to be specified to update the Continuous Export
// configuration.
func (client ExportConfigurationsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, exportID string, exportProperties ApplicationInsightsComponentExportRequest) (result ApplicationInsightsComponentExportConfiguration, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, exportID, exportProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ExportConfigurationsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ExportConfigurationsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, exportID string, exportProperties ApplicationInsightsComponentExportRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"exportId":          autorest.Encode("path", exportID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}", pathParameters),
		autorest.WithJSON(exportProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ExportConfigurationsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ExportConfigurationsClient) UpdateResponder(resp *http.Response) (result ApplicationInsightsComponentExportConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
