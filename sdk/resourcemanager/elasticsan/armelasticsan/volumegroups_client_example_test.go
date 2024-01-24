//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_ListByElasticSan_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_NewListByElasticSanPager_volumeGroupsListByElasticSanMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeGroupsClient().NewListByElasticSanPager("resourcegroupname", "elasticsanname", nil)
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
		// page.VolumeGroupList = armelasticsan.VolumeGroupList{
		// 	Value: []*armelasticsan.VolumeGroup{
		// 		{
		// 			Name: to.Ptr("cr"),
		// 			Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"),
		// 			SystemData: &armelasticsan.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
		// 				CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
		// 				CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("bcclmbseed"),
		// 				LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 			},
		// 			Identity: &armelasticsan.Identity{
		// 				Type: to.Ptr(armelasticsan.IdentityTypeNone),
		// 				PrincipalID: to.Ptr("ihsiwrwdofymkhquaxcrtfmmrsygw"),
		// 				TenantID: to.Ptr("gtkzkjsy"),
		// 				UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
		// 					"key7482": &armelasticsan.UserAssignedIdentity{
		// 						ClientID: to.Ptr("jaczsquolgxwpznljbmdupn"),
		// 						PrincipalID: to.Ptr("vfdzizicxcfcqecgsmshz"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armelasticsan.VolumeGroupProperties{
		// 				Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
		// 				EncryptionProperties: &armelasticsan.EncryptionProperties{
		// 					EncryptionIdentity: &armelasticsan.EncryptionIdentity{
		// 						EncryptionUserAssignedIdentity: to.Ptr("im"),
		// 					},
		// 					KeyVaultProperties: &armelasticsan.KeyVaultProperties{
		// 						CurrentVersionedKeyExpirationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
		// 						CurrentVersionedKeyIdentifier: to.Ptr("rnpxhtzkquzyoepwbwktbwb"),
		// 						KeyName: to.Ptr("sftaiernmrzypnrkpakrrawxcbsqzc"),
		// 						KeyVaultURI: to.Ptr("https://microsoft.com/axmblwp"),
		// 						KeyVersion: to.Ptr("c"),
		// 						LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
		// 					},
		// 				},
		// 				NetworkACLs: &armelasticsan.NetworkRuleSet{
		// 					VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
		// 						{
		// 							Action: to.Ptr(armelasticsan.ActionAllow),
		// 							VirtualNetworkResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}"),
		// 					}},
		// 				},
		// 				PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
		// 					{
		// 						Name: to.Ptr("{privateEndpointConnectionName}"),
		// 						Type: to.Ptr("Microsoft.ElasticSan/elasticSans/privateEndpointConnections"),
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/privateEndpointConnections/{privateEndpointConnectionName}"),
		// 						SystemData: &armelasticsan.SystemData{
		// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
		// 							CreatedBy: to.Ptr("otfifnrahdshqombvtg"),
		// 							CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 							LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
		// 							LastModifiedBy: to.Ptr("jnaxavnlhrboshtidtib"),
		// 							LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 						},
		// 						Properties: &armelasticsan.PrivateEndpointConnectionProperties{
		// 							GroupIDs: []*string{
		// 								to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}")},
		// 								PrivateEndpoint: &armelasticsan.PrivateEndpoint{
		// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
		// 									Description: to.Ptr("Auto-Approved"),
		// 									ActionsRequired: to.Ptr("None"),
		// 									Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
		// 								},
		// 								ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 							},
		// 					}},
		// 					ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
		// 					ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_ListByElasticSan_MinimumSet_Gen.json
func ExampleVolumeGroupsClient_NewListByElasticSanPager_volumeGroupsListByElasticSanMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeGroupsClient().NewListByElasticSanPager("resourcegroupname", "elasticsanname", nil)
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
		// page.VolumeGroupList = armelasticsan.VolumeGroupList{
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Create_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_BeginCreate_volumeGroupsCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeGroupsClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", armelasticsan.VolumeGroup{
		Identity: &armelasticsan.Identity{
			Type: to.Ptr(armelasticsan.IdentityTypeNone),
			UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
				"key7482": {},
			},
		},
		Properties: &armelasticsan.VolumeGroupProperties{
			Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithCustomerManagedKey),
			EncryptionProperties: &armelasticsan.EncryptionProperties{
				EncryptionIdentity: &armelasticsan.EncryptionIdentity{
					EncryptionUserAssignedIdentity: to.Ptr("im"),
				},
				KeyVaultProperties: &armelasticsan.KeyVaultProperties{
					KeyName:     to.Ptr("sftaiernmrzypnrkpakrrawxcbsqzc"),
					KeyVaultURI: to.Ptr("https://microsoft.com/axmblwp"),
					KeyVersion:  to.Ptr("c"),
				},
			},
			NetworkACLs: &armelasticsan.NetworkRuleSet{
				VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
					{
						Action:                   to.Ptr(armelasticsan.ActionAllow),
						VirtualNetworkResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}"),
					}},
			},
			ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
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
	// res.VolumeGroup = armelasticsan.VolumeGroup{
	// 	Name: to.Ptr("cr"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Identity: &armelasticsan.Identity{
	// 		Type: to.Ptr(armelasticsan.IdentityTypeNone),
	// 		PrincipalID: to.Ptr("ihsiwrwdofymkhquaxcrtfmmrsygw"),
	// 		TenantID: to.Ptr("gtkzkjsy"),
	// 		UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
	// 			"key7482": &armelasticsan.UserAssignedIdentity{
	// 				ClientID: to.Ptr("jaczsquolgxwpznljbmdupn"),
	// 				PrincipalID: to.Ptr("vfdzizicxcfcqecgsmshz"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armelasticsan.VolumeGroupProperties{
	// 		Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 		EncryptionProperties: &armelasticsan.EncryptionProperties{
	// 			EncryptionIdentity: &armelasticsan.EncryptionIdentity{
	// 				EncryptionUserAssignedIdentity: to.Ptr("im"),
	// 			},
	// 			KeyVaultProperties: &armelasticsan.KeyVaultProperties{
	// 				CurrentVersionedKeyExpirationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
	// 				CurrentVersionedKeyIdentifier: to.Ptr("rnpxhtzkquzyoepwbwktbwb"),
	// 				KeyName: to.Ptr("sftaiernmrzypnrkpakrrawxcbsqzc"),
	// 				KeyVaultURI: to.Ptr("https://microsoft.com/axmblwp"),
	// 				KeyVersion: to.Ptr("c"),
	// 				LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
	// 			},
	// 		},
	// 		NetworkACLs: &armelasticsan.NetworkRuleSet{
	// 			VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
	// 				{
	// 					Action: to.Ptr(armelasticsan.ActionAllow),
	// 					VirtualNetworkResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}"),
	// 			}},
	// 		},
	// 		ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Create_MinimumSet_Gen.json
func ExampleVolumeGroupsClient_BeginCreate_volumeGroupsCreateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeGroupsClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", armelasticsan.VolumeGroup{}, nil)
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
	// res.VolumeGroup = armelasticsan.VolumeGroup{
	// 	Name: to.Ptr("cr"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Identity: &armelasticsan.Identity{
	// 		Type: to.Ptr(armelasticsan.IdentityTypeNone),
	// 		PrincipalID: to.Ptr("ihsiwrwdofymkhquaxcrtfmmrsygw"),
	// 		TenantID: to.Ptr("gtkzkjsy"),
	// 		UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
	// 			"key7482": &armelasticsan.UserAssignedIdentity{
	// 				ClientID: to.Ptr("jaczsquolgxwpznljbmdupn"),
	// 				PrincipalID: to.Ptr("vfdzizicxcfcqecgsmshz"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armelasticsan.VolumeGroupProperties{
	// 		Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 		EncryptionProperties: &armelasticsan.EncryptionProperties{
	// 			EncryptionIdentity: &armelasticsan.EncryptionIdentity{
	// 				EncryptionUserAssignedIdentity: to.Ptr("im"),
	// 			},
	// 			KeyVaultProperties: &armelasticsan.KeyVaultProperties{
	// 				CurrentVersionedKeyExpirationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
	// 				CurrentVersionedKeyIdentifier: to.Ptr("rnpxhtzkquzyoepwbwktbwb"),
	// 				KeyName: to.Ptr("sftaiernmrzypnrkpakrrawxcbsqzc"),
	// 				KeyVaultURI: to.Ptr("https://microsoft.com/axmblwp"),
	// 				KeyVersion: to.Ptr("c"),
	// 				LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
	// 			},
	// 		},
	// 		NetworkACLs: &armelasticsan.NetworkRuleSet{
	// 			VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
	// 				{
	// 					Action: to.Ptr(armelasticsan.ActionAllow),
	// 					VirtualNetworkResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}"),
	// 			}},
	// 		},
	// 		PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("gewxykc"),
	// 				Type: to.Ptr("ailymcedgvxbqklmqtlty"),
	// 				ID: to.Ptr("opcjchensdf"),
	// 				SystemData: &armelasticsan.SystemData{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 					CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 					CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 					LastModifiedBy: to.Ptr("bcclmbseed"),
	// 					LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 				},
	// 				Properties: &armelasticsan.PrivateEndpointConnectionProperties{
	// 					GroupIDs: []*string{
	// 						to.Ptr("bolviufgqnyid")},
	// 						PrivateEndpoint: &armelasticsan.PrivateEndpoint{
	// 							ID: to.Ptr("ehxmltubeltzmgcqxocakaansat"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("nahklgxicbqjbbvcdrkljqdhprruys"),
	// 							ActionsRequired: to.Ptr("sairafcqpvucoy"),
	// 							Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
	// 						},
	// 						ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 					},
	// 			}},
	// 			ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Update_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_BeginUpdate_volumeGroupsUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeGroupsClient().BeginUpdate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", armelasticsan.VolumeGroupUpdate{
		Identity: &armelasticsan.Identity{
			Type: to.Ptr(armelasticsan.IdentityTypeNone),
			UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
				"key7482": {},
			},
		},
		Properties: &armelasticsan.VolumeGroupUpdateProperties{
			Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
			EncryptionProperties: &armelasticsan.EncryptionProperties{
				EncryptionIdentity: &armelasticsan.EncryptionIdentity{
					EncryptionUserAssignedIdentity: to.Ptr("im"),
				},
				KeyVaultProperties: &armelasticsan.KeyVaultProperties{
					KeyName:     to.Ptr("sftaiernmrzypnrkpakrrawxcbsqzc"),
					KeyVaultURI: to.Ptr("https://microsoft.com/axmblwp"),
					KeyVersion:  to.Ptr("c"),
				},
			},
			NetworkACLs: &armelasticsan.NetworkRuleSet{
				VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
					{
						Action:                   to.Ptr(armelasticsan.ActionAllow),
						VirtualNetworkResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}"),
					}},
			},
			ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
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
	// res.VolumeGroup = armelasticsan.VolumeGroup{
	// 	Name: to.Ptr("cr"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Identity: &armelasticsan.Identity{
	// 		Type: to.Ptr(armelasticsan.IdentityTypeNone),
	// 		PrincipalID: to.Ptr("ihsiwrwdofymkhquaxcrtfmmrsygw"),
	// 		TenantID: to.Ptr("gtkzkjsy"),
	// 		UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
	// 			"key7482": &armelasticsan.UserAssignedIdentity{
	// 				ClientID: to.Ptr("jaczsquolgxwpznljbmdupn"),
	// 				PrincipalID: to.Ptr("vfdzizicxcfcqecgsmshz"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armelasticsan.VolumeGroupProperties{
	// 		Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 		EncryptionProperties: &armelasticsan.EncryptionProperties{
	// 			EncryptionIdentity: &armelasticsan.EncryptionIdentity{
	// 				EncryptionUserAssignedIdentity: to.Ptr("im"),
	// 			},
	// 			KeyVaultProperties: &armelasticsan.KeyVaultProperties{
	// 				CurrentVersionedKeyExpirationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
	// 				CurrentVersionedKeyIdentifier: to.Ptr("rnpxhtzkquzyoepwbwktbwb"),
	// 				KeyName: to.Ptr("sftaiernmrzypnrkpakrrawxcbsqzc"),
	// 				KeyVaultURI: to.Ptr("https://microsoft.com/axmblwp"),
	// 				KeyVersion: to.Ptr("c"),
	// 				LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
	// 			},
	// 		},
	// 		NetworkACLs: &armelasticsan.NetworkRuleSet{
	// 			VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
	// 				{
	// 					Action: to.Ptr(armelasticsan.ActionAllow),
	// 					VirtualNetworkResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}"),
	// 			}},
	// 		},
	// 		PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("{privateEndpointConnectionName}"),
	// 				Type: to.Ptr("Microsoft.ElasticSan/elasticSans/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/privateEndpointConnections/{privateEndpointConnectionName}"),
	// 				SystemData: &armelasticsan.SystemData{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
	// 					CreatedBy: to.Ptr("otfifnrahdshqombvtg"),
	// 					CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
	// 					LastModifiedBy: to.Ptr("jnaxavnlhrboshtidtib"),
	// 					LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 				},
	// 				Properties: &armelasticsan.PrivateEndpointConnectionProperties{
	// 					GroupIDs: []*string{
	// 						to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}")},
	// 						PrivateEndpoint: &armelasticsan.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Auto-Approved"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
	// 						},
	// 						ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 					},
	// 			}},
	// 			ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Update_MinimumSet_Gen.json
func ExampleVolumeGroupsClient_BeginUpdate_volumeGroupsUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeGroupsClient().BeginUpdate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", armelasticsan.VolumeGroupUpdate{}, nil)
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
	// res.VolumeGroup = armelasticsan.VolumeGroup{
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Delete_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_BeginDelete_volumeGroupsDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeGroupsClient().BeginDelete(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Delete_MinimumSet_Gen.json
func ExampleVolumeGroupsClient_BeginDelete_volumeGroupsDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeGroupsClient().BeginDelete(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Get_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_Get_volumeGroupsGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeGroupsClient().Get(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VolumeGroup = armelasticsan.VolumeGroup{
	// 	Name: to.Ptr("cr"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Identity: &armelasticsan.Identity{
	// 		Type: to.Ptr(armelasticsan.IdentityTypeNone),
	// 		PrincipalID: to.Ptr("ihsiwrwdofymkhquaxcrtfmmrsygw"),
	// 		TenantID: to.Ptr("gtkzkjsy"),
	// 		UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
	// 			"key7482": &armelasticsan.UserAssignedIdentity{
	// 				ClientID: to.Ptr("jaczsquolgxwpznljbmdupn"),
	// 				PrincipalID: to.Ptr("vfdzizicxcfcqecgsmshz"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armelasticsan.VolumeGroupProperties{
	// 		Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 		EncryptionProperties: &armelasticsan.EncryptionProperties{
	// 			EncryptionIdentity: &armelasticsan.EncryptionIdentity{
	// 				EncryptionUserAssignedIdentity: to.Ptr("im"),
	// 			},
	// 			KeyVaultProperties: &armelasticsan.KeyVaultProperties{
	// 				CurrentVersionedKeyExpirationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
	// 				CurrentVersionedKeyIdentifier: to.Ptr("rnpxhtzkquzyoepwbwktbwb"),
	// 				KeyName: to.Ptr("sftaiernmrzypnrkpakrrawxcbsqzc"),
	// 				KeyVaultURI: to.Ptr("https://microsoft.com/axmblwp"),
	// 				KeyVersion: to.Ptr("c"),
	// 				LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
	// 			},
	// 		},
	// 		NetworkACLs: &armelasticsan.NetworkRuleSet{
	// 			VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
	// 				{
	// 					Action: to.Ptr(armelasticsan.ActionAllow),
	// 					VirtualNetworkResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}"),
	// 			}},
	// 		},
	// 		PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("{privateEndpointConnectionName}"),
	// 				Type: to.Ptr("Microsoft.ElasticSan/elasticSans/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/privateEndpointConnections/{privateEndpointConnectionName}"),
	// 				SystemData: &armelasticsan.SystemData{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
	// 					CreatedBy: to.Ptr("otfifnrahdshqombvtg"),
	// 					CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
	// 					LastModifiedBy: to.Ptr("jnaxavnlhrboshtidtib"),
	// 					LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 				},
	// 				Properties: &armelasticsan.PrivateEndpointConnectionProperties{
	// 					GroupIDs: []*string{
	// 						to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}")},
	// 						PrivateEndpoint: &armelasticsan.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Auto-Approved"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
	// 						},
	// 						ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 					},
	// 			}},
	// 			ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Get_MinimumSet_Gen.json
func ExampleVolumeGroupsClient_Get_volumeGroupsGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeGroupsClient().Get(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VolumeGroup = armelasticsan.VolumeGroup{
	// }
}
