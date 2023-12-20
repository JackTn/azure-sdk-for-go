//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armcontainerservice.ClientFactory type.
type ServerFactory struct {
	AgentPoolsServer                  AgentPoolsServer
	MachinesServer                    MachinesServer
	MaintenanceConfigurationsServer   MaintenanceConfigurationsServer
	ManagedClusterSnapshotsServer     ManagedClusterSnapshotsServer
	ManagedClustersServer             ManagedClustersServer
	OperationStatusResultServer       OperationStatusResultServer
	OperationsServer                  OperationsServer
	PrivateEndpointConnectionsServer  PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer        PrivateLinkResourcesServer
	ResolvePrivateLinkServiceIDServer ResolvePrivateLinkServiceIDServer
	SnapshotsServer                   SnapshotsServer
	TrustedAccessRoleBindingsServer   TrustedAccessRoleBindingsServer
	TrustedAccessRolesServer          TrustedAccessRolesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armcontainerservice.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armcontainerservice.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                 *ServerFactory
	trMu                                sync.Mutex
	trAgentPoolsServer                  *AgentPoolsServerTransport
	trMachinesServer                    *MachinesServerTransport
	trMaintenanceConfigurationsServer   *MaintenanceConfigurationsServerTransport
	trManagedClusterSnapshotsServer     *ManagedClusterSnapshotsServerTransport
	trManagedClustersServer             *ManagedClustersServerTransport
	trOperationStatusResultServer       *OperationStatusResultServerTransport
	trOperationsServer                  *OperationsServerTransport
	trPrivateEndpointConnectionsServer  *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer        *PrivateLinkResourcesServerTransport
	trResolvePrivateLinkServiceIDServer *ResolvePrivateLinkServiceIDServerTransport
	trSnapshotsServer                   *SnapshotsServerTransport
	trTrustedAccessRoleBindingsServer   *TrustedAccessRoleBindingsServerTransport
	trTrustedAccessRolesServer          *TrustedAccessRolesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "AgentPoolsClient":
		initServer(s, &s.trAgentPoolsServer, func() *AgentPoolsServerTransport { return NewAgentPoolsServerTransport(&s.srv.AgentPoolsServer) })
		resp, err = s.trAgentPoolsServer.Do(req)
	case "MachinesClient":
		initServer(s, &s.trMachinesServer, func() *MachinesServerTransport { return NewMachinesServerTransport(&s.srv.MachinesServer) })
		resp, err = s.trMachinesServer.Do(req)
	case "MaintenanceConfigurationsClient":
		initServer(s, &s.trMaintenanceConfigurationsServer, func() *MaintenanceConfigurationsServerTransport {
			return NewMaintenanceConfigurationsServerTransport(&s.srv.MaintenanceConfigurationsServer)
		})
		resp, err = s.trMaintenanceConfigurationsServer.Do(req)
	case "ManagedClusterSnapshotsClient":
		initServer(s, &s.trManagedClusterSnapshotsServer, func() *ManagedClusterSnapshotsServerTransport {
			return NewManagedClusterSnapshotsServerTransport(&s.srv.ManagedClusterSnapshotsServer)
		})
		resp, err = s.trManagedClusterSnapshotsServer.Do(req)
	case "ManagedClustersClient":
		initServer(s, &s.trManagedClustersServer, func() *ManagedClustersServerTransport {
			return NewManagedClustersServerTransport(&s.srv.ManagedClustersServer)
		})
		resp, err = s.trManagedClustersServer.Do(req)
	case "OperationStatusResultClient":
		initServer(s, &s.trOperationStatusResultServer, func() *OperationStatusResultServerTransport {
			return NewOperationStatusResultServerTransport(&s.srv.OperationStatusResultServer)
		})
		resp, err = s.trOperationStatusResultServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "ResolvePrivateLinkServiceIDClient":
		initServer(s, &s.trResolvePrivateLinkServiceIDServer, func() *ResolvePrivateLinkServiceIDServerTransport {
			return NewResolvePrivateLinkServiceIDServerTransport(&s.srv.ResolvePrivateLinkServiceIDServer)
		})
		resp, err = s.trResolvePrivateLinkServiceIDServer.Do(req)
	case "SnapshotsClient":
		initServer(s, &s.trSnapshotsServer, func() *SnapshotsServerTransport { return NewSnapshotsServerTransport(&s.srv.SnapshotsServer) })
		resp, err = s.trSnapshotsServer.Do(req)
	case "TrustedAccessRoleBindingsClient":
		initServer(s, &s.trTrustedAccessRoleBindingsServer, func() *TrustedAccessRoleBindingsServerTransport {
			return NewTrustedAccessRoleBindingsServerTransport(&s.srv.TrustedAccessRoleBindingsServer)
		})
		resp, err = s.trTrustedAccessRoleBindingsServer.Do(req)
	case "TrustedAccessRolesClient":
		initServer(s, &s.trTrustedAccessRolesServer, func() *TrustedAccessRolesServerTransport {
			return NewTrustedAccessRolesServerTransport(&s.srv.TrustedAccessRolesServer)
		})
		resp, err = s.trTrustedAccessRolesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
