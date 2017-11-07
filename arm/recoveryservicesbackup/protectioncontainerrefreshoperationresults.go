package recoveryservicesbackup

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
	"net/http"
)

// ProtectionContainerRefreshOperationResultsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectionContainerRefreshOperationResultsClient struct {
	ManagementClient
}

// NewProtectionContainerRefreshOperationResultsClient creates an instance of the
// ProtectionContainerRefreshOperationResultsClient client.
func NewProtectionContainerRefreshOperationResultsClient(subscriptionID string) ProtectionContainerRefreshOperationResultsClient {
	return NewProtectionContainerRefreshOperationResultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectionContainerRefreshOperationResultsClientWithBaseURI creates an instance of the
// ProtectionContainerRefreshOperationResultsClient client.
func NewProtectionContainerRefreshOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) ProtectionContainerRefreshOperationResultsClient {
	return ProtectionContainerRefreshOperationResultsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get provides the result of the refresh operation triggered by the BeginRefresh operation.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. fabricName is fabric name associated with the container. operationID is
// operation ID associated with the operation whose result needs to be fetched.
func (client ProtectionContainerRefreshOperationResultsClient) Get(vaultName string, resourceGroupName string, fabricName string, operationID string) (result autorest.Response, err error) {
	req, err := client.GetPreparer(vaultName, resourceGroupName, fabricName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionContainerRefreshOperationResultsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionContainerRefreshOperationResultsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionContainerRefreshOperationResultsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProtectionContainerRefreshOperationResultsClient) GetPreparer(vaultName string, resourceGroupName string, fabricName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/operationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectionContainerRefreshOperationResultsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectionContainerRefreshOperationResultsClient) GetResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
