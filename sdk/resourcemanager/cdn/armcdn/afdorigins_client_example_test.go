//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOrigins_ListByOriginGroup.json
func ExampleAFDOriginsClient_NewListByOriginGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAFDOriginsClient().NewListByOriginGroupPager("RG", "profile1", "origingroup1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AFDOriginListResult = armcdn.AFDOriginListResult{
		// 	Value: []*armcdn.AFDOrigin{
		// 		{
		// 			Name: to.Ptr("origin1"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles/origingroups/origins"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/origingroups/origingroup1/origins/origin1"),
		// 			Properties: &armcdn.AFDOriginProperties{
		// 				EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
		// 				EnforceCertificateNameCheck: to.Ptr(true),
		// 				HostName: to.Ptr("host1.blob.core.windows.net"),
		// 				HTTPPort: to.Ptr[int32](80),
		// 				HTTPSPort: to.Ptr[int32](443),
		// 				OriginGroupName: to.Ptr("origingroup1"),
		// 				OriginHostHeader: to.Ptr("host1.foo.com"),
		// 				DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
		// 				ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOrigins_Get.json
func ExampleAFDOriginsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAFDOriginsClient().Get(ctx, "RG", "profile1", "origingroup1", "origin1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AFDOrigin = armcdn.AFDOrigin{
	// 	Name: to.Ptr("origin1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/origingroups/origins"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/origingroups/origingroup1/origins/origin1"),
	// 	Properties: &armcdn.AFDOriginProperties{
	// 		EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
	// 		EnforceCertificateNameCheck: to.Ptr(true),
	// 		HostName: to.Ptr("host1.blob.core.windows.net"),
	// 		HTTPPort: to.Ptr[int32](80),
	// 		HTTPSPort: to.Ptr[int32](443),
	// 		OriginGroupName: to.Ptr("origingroup1"),
	// 		OriginHostHeader: to.Ptr("host1.foo.com"),
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOrigins_Create.json
func ExampleAFDOriginsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAFDOriginsClient().BeginCreate(ctx, "RG", "profile1", "origingroup1", "origin1", armcdn.AFDOrigin{
		Properties: &armcdn.AFDOriginProperties{
			EnabledState:     to.Ptr(armcdn.EnabledStateEnabled),
			HostName:         to.Ptr("host1.blob.core.windows.net"),
			HTTPPort:         to.Ptr[int32](80),
			HTTPSPort:        to.Ptr[int32](443),
			OriginHostHeader: to.Ptr("host1.foo.com"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AFDOrigin = armcdn.AFDOrigin{
	// 	Name: to.Ptr("origin1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/origingroups/origins"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/origingroups/origingroup1/origins/origin1"),
	// 	Properties: &armcdn.AFDOriginProperties{
	// 		EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
	// 		EnforceCertificateNameCheck: to.Ptr(true),
	// 		HostName: to.Ptr("host1.blob.core.windows.net"),
	// 		HTTPPort: to.Ptr[int32](80),
	// 		HTTPSPort: to.Ptr[int32](443),
	// 		OriginGroupName: to.Ptr("origingroup1"),
	// 		OriginHostHeader: to.Ptr("host1.foo.com"),
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOrigins_Update.json
func ExampleAFDOriginsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAFDOriginsClient().BeginUpdate(ctx, "RG", "profile1", "origingroup1", "origin1", armcdn.AFDOriginUpdateParameters{
		Properties: &armcdn.AFDOriginUpdatePropertiesParameters{
			EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
			HostName:     to.Ptr("host1.blob.core.windows.net"),
			HTTPPort:     to.Ptr[int32](80),
			HTTPSPort:    to.Ptr[int32](443),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AFDOrigin = armcdn.AFDOrigin{
	// 	Name: to.Ptr("origin1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/origingroups/origins"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/origingroups/origingroup1/origins/origin1"),
	// 	Properties: &armcdn.AFDOriginProperties{
	// 		EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
	// 		EnforceCertificateNameCheck: to.Ptr(true),
	// 		HostName: to.Ptr("host1.blob.core.windows.net"),
	// 		HTTPPort: to.Ptr[int32](80),
	// 		HTTPSPort: to.Ptr[int32](443),
	// 		OriginGroupName: to.Ptr("origingroup1"),
	// 		OriginHostHeader: to.Ptr("host1.foo.com"),
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOrigins_Delete.json
func ExampleAFDOriginsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAFDOriginsClient().BeginDelete(ctx, "RG", "profile1", "origingroup1", "origin1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
