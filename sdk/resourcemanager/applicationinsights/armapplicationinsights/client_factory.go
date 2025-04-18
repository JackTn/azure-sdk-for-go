// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAPIKeysClient creates a new instance of APIKeysClient.
func (c *ClientFactory) NewAPIKeysClient() *APIKeysClient {
	return &APIKeysClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAnalyticsItemsClient creates a new instance of AnalyticsItemsClient.
func (c *ClientFactory) NewAnalyticsItemsClient() *AnalyticsItemsClient {
	return &AnalyticsItemsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAnnotationsClient creates a new instance of AnnotationsClient.
func (c *ClientFactory) NewAnnotationsClient() *AnnotationsClient {
	return &AnnotationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewComponentAvailableFeaturesClient creates a new instance of ComponentAvailableFeaturesClient.
func (c *ClientFactory) NewComponentAvailableFeaturesClient() *ComponentAvailableFeaturesClient {
	return &ComponentAvailableFeaturesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewComponentCurrentBillingFeaturesClient creates a new instance of ComponentCurrentBillingFeaturesClient.
func (c *ClientFactory) NewComponentCurrentBillingFeaturesClient() *ComponentCurrentBillingFeaturesClient {
	return &ComponentCurrentBillingFeaturesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewComponentFeatureCapabilitiesClient creates a new instance of ComponentFeatureCapabilitiesClient.
func (c *ClientFactory) NewComponentFeatureCapabilitiesClient() *ComponentFeatureCapabilitiesClient {
	return &ComponentFeatureCapabilitiesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewComponentLinkedStorageAccountsClient creates a new instance of ComponentLinkedStorageAccountsClient.
func (c *ClientFactory) NewComponentLinkedStorageAccountsClient() *ComponentLinkedStorageAccountsClient {
	return &ComponentLinkedStorageAccountsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewComponentQuotaStatusClient creates a new instance of ComponentQuotaStatusClient.
func (c *ClientFactory) NewComponentQuotaStatusClient() *ComponentQuotaStatusClient {
	return &ComponentQuotaStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewComponentsClient creates a new instance of ComponentsClient.
func (c *ClientFactory) NewComponentsClient() *ComponentsClient {
	return &ComponentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDeletedWorkbooksClient creates a new instance of DeletedWorkbooksClient.
func (c *ClientFactory) NewDeletedWorkbooksClient() *DeletedWorkbooksClient {
	return &DeletedWorkbooksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExportConfigurationsClient creates a new instance of ExportConfigurationsClient.
func (c *ClientFactory) NewExportConfigurationsClient() *ExportConfigurationsClient {
	return &ExportConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFavoritesClient creates a new instance of FavoritesClient.
func (c *ClientFactory) NewFavoritesClient() *FavoritesClient {
	return &FavoritesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLiveTokenClient creates a new instance of LiveTokenClient.
func (c *ClientFactory) NewLiveTokenClient() *LiveTokenClient {
	return &LiveTokenClient{
		internal: c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewProactiveDetectionConfigurationsClient creates a new instance of ProactiveDetectionConfigurationsClient.
func (c *ClientFactory) NewProactiveDetectionConfigurationsClient() *ProactiveDetectionConfigurationsClient {
	return &ProactiveDetectionConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWebTestLocationsClient creates a new instance of WebTestLocationsClient.
func (c *ClientFactory) NewWebTestLocationsClient() *WebTestLocationsClient {
	return &WebTestLocationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWebTestsClient creates a new instance of WebTestsClient.
func (c *ClientFactory) NewWebTestsClient() *WebTestsClient {
	return &WebTestsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkItemConfigurationsClient creates a new instance of WorkItemConfigurationsClient.
func (c *ClientFactory) NewWorkItemConfigurationsClient() *WorkItemConfigurationsClient {
	return &WorkItemConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkbookTemplatesClient creates a new instance of WorkbookTemplatesClient.
func (c *ClientFactory) NewWorkbookTemplatesClient() *WorkbookTemplatesClient {
	return &WorkbookTemplatesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkbooksClient creates a new instance of WorkbooksClient.
func (c *ClientFactory) NewWorkbooksClient() *WorkbooksClient {
	return &WorkbooksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
