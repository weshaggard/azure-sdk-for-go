package cdn

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// EndpointsClient is the use these APIs to manage Azure CDN resources through the Azure Resource Manager. You must
// make sure that requests made to these resources are secure. For more information, see
// https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx.
type EndpointsClient struct {
	ManagementClient
}

// NewEndpointsClient creates an instance of the EndpointsClient client.
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return NewEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEndpointsClientWithBaseURI creates an instance of the EndpointsClient client.
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return EndpointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create sends the create request. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// endpointName is name of the endpoint within the CDN profile. endpointProperties is endpoint properties profileName
// is name of the CDN profile within the resource group. resourceGroupName is name of the resource group within the
// Azure subscription.
func (client EndpointsClient) Create(endpointName string, endpointProperties EndpointCreateParameters, profileName string, resourceGroupName string, cancel <-chan struct{}) (<-chan Endpoint, <-chan error) {
	resultChan := make(chan Endpoint, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: endpointProperties,
			Constraints: []validation.Constraint{{Target: "endpointProperties.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "endpointProperties.EndpointPropertiesCreateParameters", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "endpointProperties.EndpointPropertiesCreateParameters.Origins", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "cdn.EndpointsClient", "Create")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result Endpoint
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(endpointName, endpointProperties, profileName, resourceGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client EndpointsClient) CreatePreparer(endpointName string, endpointProperties EndpointCreateParameters, profileName string, resourceGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", pathParameters),
		autorest.WithJSON(endpointProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client EndpointsClient) CreateResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteIfExists sends the delete if exists request. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// endpointName is name of the endpoint within the CDN profile. profileName is name of the CDN profile within the
// resource group. resourceGroupName is name of the resource group within the Azure subscription.
func (client EndpointsClient) DeleteIfExists(endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeleteIfExistsPreparer(endpointName, profileName, resourceGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "DeleteIfExists", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteIfExistsSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "DeleteIfExists", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteIfExistsResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "DeleteIfExists", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeleteIfExistsPreparer prepares the DeleteIfExists request.
func (client EndpointsClient) DeleteIfExistsPreparer(endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteIfExistsSender sends the DeleteIfExists request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) DeleteIfExistsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteIfExistsResponder handles the response to the DeleteIfExists request. The method always
// closes the http.Response Body.
func (client EndpointsClient) DeleteIfExistsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
//
// endpointName is name of the endpoint within the CDN profile. profileName is name of the CDN profile within the
// resource group. resourceGroupName is name of the resource group within the Azure subscription.
func (client EndpointsClient) Get(endpointName string, profileName string, resourceGroupName string) (result Endpoint, err error) {
	req, err := client.GetPreparer(endpointName, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client EndpointsClient) GetPreparer(endpointName string, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EndpointsClient) GetResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByProfile sends the list by profile request.
//
// profileName is name of the CDN profile within the resource group. resourceGroupName is name of the resource group
// within the Azure subscription.
func (client EndpointsClient) ListByProfile(profileName string, resourceGroupName string) (result EndpointListResult, err error) {
	req, err := client.ListByProfilePreparer(profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ListByProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ListByProfile", resp, "Failure sending request")
		return
	}

	result, err = client.ListByProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ListByProfile", resp, "Failure responding to request")
	}

	return
}

// ListByProfilePreparer prepares the ListByProfile request.
func (client EndpointsClient) ListByProfilePreparer(profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByProfileSender sends the ListByProfile request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) ListByProfileSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByProfileResponder handles the response to the ListByProfile request. The method always
// closes the http.Response Body.
func (client EndpointsClient) ListByProfileResponder(resp *http.Response) (result EndpointListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// LoadContent sends the load content request. This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// endpointName is name of the endpoint within the CDN profile. contentFilePaths is the path to the content to be
// loaded. Path should describe a file. profileName is name of the CDN profile within the resource group.
// resourceGroupName is name of the resource group within the Azure subscription.
func (client EndpointsClient) LoadContent(endpointName string, contentFilePaths LoadParameters, profileName string, resourceGroupName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: contentFilePaths,
			Constraints: []validation.Constraint{{Target: "contentFilePaths.ContentPaths", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "cdn.EndpointsClient", "LoadContent")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.LoadContentPreparer(endpointName, contentFilePaths, profileName, resourceGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "LoadContent", nil, "Failure preparing request")
			return
		}

		resp, err := client.LoadContentSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "LoadContent", resp, "Failure sending request")
			return
		}

		result, err = client.LoadContentResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "LoadContent", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// LoadContentPreparer prepares the LoadContent request.
func (client EndpointsClient) LoadContentPreparer(endpointName string, contentFilePaths LoadParameters, profileName string, resourceGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/load", pathParameters),
		autorest.WithJSON(contentFilePaths),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// LoadContentSender sends the LoadContent request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) LoadContentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// LoadContentResponder handles the response to the LoadContent request. The method always
// closes the http.Response Body.
func (client EndpointsClient) LoadContentResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PurgeContent sends the purge content request. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// endpointName is name of the endpoint within the CDN profile. contentFilePaths is the path to the content to be
// purged. Path can describe a file or directory. profileName is name of the CDN profile within the resource group.
// resourceGroupName is name of the resource group within the Azure subscription.
func (client EndpointsClient) PurgeContent(endpointName string, contentFilePaths PurgeParameters, profileName string, resourceGroupName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: contentFilePaths,
			Constraints: []validation.Constraint{{Target: "contentFilePaths.ContentPaths", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "cdn.EndpointsClient", "PurgeContent")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.PurgeContentPreparer(endpointName, contentFilePaths, profileName, resourceGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "PurgeContent", nil, "Failure preparing request")
			return
		}

		resp, err := client.PurgeContentSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "PurgeContent", resp, "Failure sending request")
			return
		}

		result, err = client.PurgeContentResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "PurgeContent", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// PurgeContentPreparer prepares the PurgeContent request.
func (client EndpointsClient) PurgeContentPreparer(endpointName string, contentFilePaths PurgeParameters, profileName string, resourceGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/purge", pathParameters),
		autorest.WithJSON(contentFilePaths),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// PurgeContentSender sends the PurgeContent request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) PurgeContentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// PurgeContentResponder handles the response to the PurgeContent request. The method always
// closes the http.Response Body.
func (client EndpointsClient) PurgeContentResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start sends the start request. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// endpointName is name of the endpoint within the CDN profile. profileName is name of the CDN profile within the
// resource group. resourceGroupName is name of the resource group within the Azure subscription.
func (client EndpointsClient) Start(endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (<-chan Endpoint, <-chan error) {
	resultChan := make(chan Endpoint, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Endpoint
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.StartPreparer(endpointName, profileName, resourceGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Start", nil, "Failure preparing request")
			return
		}

		resp, err := client.StartSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Start", resp, "Failure sending request")
			return
		}

		result, err = client.StartResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Start", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// StartPreparer prepares the Start request.
func (client EndpointsClient) StartPreparer(endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) StartSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client EndpointsClient) StartResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Stop sends the stop request. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// endpointName is name of the endpoint within the CDN profile. profileName is name of the CDN profile within the
// resource group. resourceGroupName is name of the resource group within the Azure subscription.
func (client EndpointsClient) Stop(endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (<-chan Endpoint, <-chan error) {
	resultChan := make(chan Endpoint, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Endpoint
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.StopPreparer(endpointName, profileName, resourceGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Stop", nil, "Failure preparing request")
			return
		}

		resp, err := client.StopSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Stop", resp, "Failure sending request")
			return
		}

		result, err = client.StopResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Stop", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// StopPreparer prepares the Stop request.
func (client EndpointsClient) StopPreparer(endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) StopSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client EndpointsClient) StopResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update sends the update request. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// endpointName is name of the endpoint within the CDN profile. endpointProperties is endpoint properties profileName
// is name of the CDN profile within the resource group. resourceGroupName is name of the resource group within the
// Azure subscription.
func (client EndpointsClient) Update(endpointName string, endpointProperties EndpointUpdateParameters, profileName string, resourceGroupName string, cancel <-chan struct{}) (<-chan Endpoint, <-chan error) {
	resultChan := make(chan Endpoint, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Endpoint
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.UpdatePreparer(endpointName, endpointProperties, profileName, resourceGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Update", nil, "Failure preparing request")
			return
		}

		resp, err := client.UpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Update", resp, "Failure sending request")
			return
		}

		result, err = client.UpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Update", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// UpdatePreparer prepares the Update request.
func (client EndpointsClient) UpdatePreparer(endpointName string, endpointProperties EndpointUpdateParameters, profileName string, resourceGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", pathParameters),
		autorest.WithJSON(endpointProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client EndpointsClient) UpdateResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateCustomDomain sends the validate custom domain request.
//
// endpointName is name of the endpoint within the CDN profile. customDomainProperties is custom domain to validate.
// profileName is name of the CDN profile within the resource group. resourceGroupName is name of the resource group
// within the Azure subscription.
func (client EndpointsClient) ValidateCustomDomain(endpointName string, customDomainProperties ValidateCustomDomainInput, profileName string, resourceGroupName string) (result ValidateCustomDomainOutput, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: customDomainProperties,
			Constraints: []validation.Constraint{{Target: "customDomainProperties.HostName", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "cdn.EndpointsClient", "ValidateCustomDomain")
	}

	req, err := client.ValidateCustomDomainPreparer(endpointName, customDomainProperties, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ValidateCustomDomain", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateCustomDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ValidateCustomDomain", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateCustomDomainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ValidateCustomDomain", resp, "Failure responding to request")
	}

	return
}

// ValidateCustomDomainPreparer prepares the ValidateCustomDomain request.
func (client EndpointsClient) ValidateCustomDomainPreparer(endpointName string, customDomainProperties ValidateCustomDomainInput, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/validateCustomDomain", pathParameters),
		autorest.WithJSON(customDomainProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ValidateCustomDomainSender sends the ValidateCustomDomain request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) ValidateCustomDomainSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ValidateCustomDomainResponder handles the response to the ValidateCustomDomain request. The method always
// closes the http.Response Body.
func (client EndpointsClient) ValidateCustomDomainResponder(resp *http.Response) (result ValidateCustomDomainOutput, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
