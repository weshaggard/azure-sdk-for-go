// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: dab57ee609fffdc578f48519d5cdc980efd8cc00

package operationalinsights

import original "github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2015-11-01-preview/operationalinsights"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type DataSourcesClient = original.DataSourcesClient
type LinkedServicesClient = original.LinkedServicesClient
type DataSourceKind = original.DataSourceKind

const (
	AzureActivityLog		DataSourceKind	= original.AzureActivityLog
	ChangeTrackingCustomRegistry	DataSourceKind	= original.ChangeTrackingCustomRegistry
	ChangeTrackingDefaultPath	DataSourceKind	= original.ChangeTrackingDefaultPath
	ChangeTrackingDefaultRegistry	DataSourceKind	= original.ChangeTrackingDefaultRegistry
	ChangeTrackingPath		DataSourceKind	= original.ChangeTrackingPath
	CustomLog			DataSourceKind	= original.CustomLog
	CustomLogCollection		DataSourceKind	= original.CustomLogCollection
	GenericDataSource		DataSourceKind	= original.GenericDataSource
	IISLogs				DataSourceKind	= original.IISLogs
	LinuxPerformanceCollection	DataSourceKind	= original.LinuxPerformanceCollection
	LinuxPerformanceObject		DataSourceKind	= original.LinuxPerformanceObject
	LinuxSyslog			DataSourceKind	= original.LinuxSyslog
	LinuxSyslogCollection		DataSourceKind	= original.LinuxSyslogCollection
	WindowsEvent			DataSourceKind	= original.WindowsEvent
	WindowsPerformanceCounter	DataSourceKind	= original.WindowsPerformanceCounter
)

type EntityStatus = original.EntityStatus

const (
	Canceled		EntityStatus	= original.Canceled
	Creating		EntityStatus	= original.Creating
	Deleting		EntityStatus	= original.Deleting
	Failed			EntityStatus	= original.Failed
	ProvisioningAccount	EntityStatus	= original.ProvisioningAccount
	Succeeded		EntityStatus	= original.Succeeded
)

type SearchSortEnum = original.SearchSortEnum

const (
	Asc	SearchSortEnum	= original.Asc
	Desc	SearchSortEnum	= original.Desc
)

type SkuNameEnum = original.SkuNameEnum

const (
	Free		SkuNameEnum	= original.Free
	PerNode		SkuNameEnum	= original.PerNode
	Premium		SkuNameEnum	= original.Premium
	Standalone	SkuNameEnum	= original.Standalone
	Standard	SkuNameEnum	= original.Standard
	Unlimited	SkuNameEnum	= original.Unlimited
)

type StorageInsightState = original.StorageInsightState

const (
	ERROR	StorageInsightState	= original.ERROR
	OK	StorageInsightState	= original.OK
)

type CoreSummary = original.CoreSummary
type DataSource = original.DataSource
type DataSourceFilter = original.DataSourceFilter
type DataSourceListResult = original.DataSourceListResult
type IntelligencePack = original.IntelligencePack
type LinkedService = original.LinkedService
type LinkedServiceListResult = original.LinkedServiceListResult
type LinkedServiceProperties = original.LinkedServiceProperties
type LinkTarget = original.LinkTarget
type ListIntelligencePack = original.ListIntelligencePack
type ListLinkTarget = original.ListLinkTarget
type ManagementGroup = original.ManagementGroup
type ManagementGroupProperties = original.ManagementGroupProperties
type MetricName = original.MetricName
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SavedSearch = original.SavedSearch
type SavedSearchesListResult = original.SavedSearchesListResult
type SavedSearchProperties = original.SavedSearchProperties
type SearchError = original.SearchError
type SearchGetSchemaResponse = original.SearchGetSchemaResponse
type SearchHighlight = original.SearchHighlight
type SearchMetadata = original.SearchMetadata
type SearchMetadataSchema = original.SearchMetadataSchema
type SearchParameters = original.SearchParameters
type SearchResultsResponse = original.SearchResultsResponse
type SearchSchemaValue = original.SearchSchemaValue
type SearchSort = original.SearchSort
type SharedKeys = original.SharedKeys
type Sku = original.Sku
type StorageAccount = original.StorageAccount
type StorageInsight = original.StorageInsight
type StorageInsightListResult = original.StorageInsightListResult
type StorageInsightProperties = original.StorageInsightProperties
type StorageInsightStatus = original.StorageInsightStatus
type Tag = original.Tag
type UsageMetric = original.UsageMetric
type Workspace = original.Workspace
type WorkspaceListManagementGroupsResult = original.WorkspaceListManagementGroupsResult
type WorkspaceListResult = original.WorkspaceListResult
type WorkspaceListUsagesResult = original.WorkspaceListUsagesResult
type WorkspaceProperties = original.WorkspaceProperties
type SavedSearchesClient = original.SavedSearchesClient
type StorageInsightsClient = original.StorageInsightsClient
type WorkspacesClient = original.WorkspacesClient

func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewDataSourcesClient(subscriptionID string) DataSourcesClient {
	return original.NewDataSourcesClient(subscriptionID)
}
func NewDataSourcesClientWithBaseURI(baseURI string, subscriptionID string) DataSourcesClient {
	return original.NewDataSourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLinkedServicesClient(subscriptionID string) LinkedServicesClient {
	return original.NewLinkedServicesClient(subscriptionID)
}
func NewLinkedServicesClientWithBaseURI(baseURI string, subscriptionID string) LinkedServicesClient {
	return original.NewLinkedServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSavedSearchesClient(subscriptionID string) SavedSearchesClient {
	return original.NewSavedSearchesClient(subscriptionID)
}
func NewSavedSearchesClientWithBaseURI(baseURI string, subscriptionID string) SavedSearchesClient {
	return original.NewSavedSearchesClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageInsightsClient(subscriptionID string) StorageInsightsClient {
	return original.NewStorageInsightsClient(subscriptionID)
}
func NewStorageInsightsClientWithBaseURI(baseURI string, subscriptionID string) StorageInsightsClient {
	return original.NewStorageInsightsClientWithBaseURI(baseURI, subscriptionID)
}
