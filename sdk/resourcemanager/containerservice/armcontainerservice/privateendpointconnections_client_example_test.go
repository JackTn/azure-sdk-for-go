//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-10-02-preview/examples/PrivateEndpointConnectionsList.json
func ExamplePrivateEndpointConnectionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().List(ctx, "rg1", "clustername1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionListResult = armcontainerservice.PrivateEndpointConnectionListResult{
	// 	Value: []*armcontainerservice.PrivateEndpointConnection{
	// 		{
	// 			Name: to.Ptr("privateendpointconnection1"),
	// 			Type: to.Ptr("Microsoft.Network/privateLinkServices/privateEndpointConnections"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedCluster/clustername1/privateEndpointConnections/privateendpointconnection1"),
	// 			Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
	// 				PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
	// 					ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
	// 				},
	// 				PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
	// 					Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
	// 				},
	// 				ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-10-02-preview/examples/PrivateEndpointConnectionsGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "rg1", "clustername1", "privateendpointconnection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armcontainerservice.PrivateEndpointConnection{
	// 	Name: to.Ptr("privateendpointconnection1"),
	// 	Type: to.Ptr("Microsoft.Network/privateLinkServices/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedCluster/clustername1/privateEndpointConnections/privateendpointconnection1"),
	// 	Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
	// 			Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-10-02-preview/examples/PrivateEndpointConnectionsUpdate.json
func ExamplePrivateEndpointConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Update(ctx, "rg1", "clustername1", "privateendpointconnection1", armcontainerservice.PrivateEndpointConnection{
		Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
				Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armcontainerservice.PrivateEndpointConnection{
	// 	Name: to.Ptr("privateendpointconnection1"),
	// 	Type: to.Ptr("Microsoft.Network/privateLinkServices/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedCluster/clustername1/privateEndpointConnections/privateendpointconnection1"),
	// 	Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
	// 			Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-10-02-preview/examples/PrivateEndpointConnectionsDelete.json
func ExamplePrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginDelete(ctx, "rg1", "clustername1", "privateendpointconnection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
