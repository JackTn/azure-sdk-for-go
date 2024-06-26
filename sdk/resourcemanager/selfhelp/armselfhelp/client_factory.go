//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armselfhelp

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
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
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

// NewCheckNameAvailabilityClient creates a new instance of CheckNameAvailabilityClient.
func (c *ClientFactory) NewCheckNameAvailabilityClient() *CheckNameAvailabilityClient {
	return &CheckNameAvailabilityClient{
		internal: c.internal,
	}
}

// NewDiagnosticsClient creates a new instance of DiagnosticsClient.
func (c *ClientFactory) NewDiagnosticsClient() *DiagnosticsClient {
	return &DiagnosticsClient{
		internal: c.internal,
	}
}

// NewDiscoverySolutionClient creates a new instance of DiscoverySolutionClient.
func (c *ClientFactory) NewDiscoverySolutionClient() *DiscoverySolutionClient {
	return &DiscoverySolutionClient{
		internal: c.internal,
	}
}

// NewDiscoverySolutionNLPSubscriptionScopeClient creates a new instance of DiscoverySolutionNLPSubscriptionScopeClient.
func (c *ClientFactory) NewDiscoverySolutionNLPSubscriptionScopeClient() *DiscoverySolutionNLPSubscriptionScopeClient {
	return &DiscoverySolutionNLPSubscriptionScopeClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDiscoverySolutionNLPTenantScopeClient creates a new instance of DiscoverySolutionNLPTenantScopeClient.
func (c *ClientFactory) NewDiscoverySolutionNLPTenantScopeClient() *DiscoverySolutionNLPTenantScopeClient {
	return &DiscoverySolutionNLPTenantScopeClient{
		internal: c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewSimplifiedSolutionsClient creates a new instance of SimplifiedSolutionsClient.
func (c *ClientFactory) NewSimplifiedSolutionsClient() *SimplifiedSolutionsClient {
	return &SimplifiedSolutionsClient{
		internal: c.internal,
	}
}

// NewSolutionClient creates a new instance of SolutionClient.
func (c *ClientFactory) NewSolutionClient() *SolutionClient {
	return &SolutionClient{
		internal: c.internal,
	}
}

// NewSolutionSelfHelpClient creates a new instance of SolutionSelfHelpClient.
func (c *ClientFactory) NewSolutionSelfHelpClient() *SolutionSelfHelpClient {
	return &SolutionSelfHelpClient{
		internal: c.internal,
	}
}

// NewTroubleshootersClient creates a new instance of TroubleshootersClient.
func (c *ClientFactory) NewTroubleshootersClient() *TroubleshootersClient {
	return &TroubleshootersClient{
		internal: c.internal,
	}
}
