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

// ServerFactory is a fake server for instances of the armquota.ClientFactory type.
type ServerFactory struct {
	// Server contains the fakes for client Client
	Server Server

	// GroupQuotaLimitsServer contains the fakes for client GroupQuotaLimitsClient
	GroupQuotaLimitsServer GroupQuotaLimitsServer

	// GroupQuotaLimitsRequestServer contains the fakes for client GroupQuotaLimitsRequestClient
	GroupQuotaLimitsRequestServer GroupQuotaLimitsRequestServer

	// GroupQuotaSubscriptionAllocationServer contains the fakes for client GroupQuotaSubscriptionAllocationClient
	GroupQuotaSubscriptionAllocationServer GroupQuotaSubscriptionAllocationServer

	// GroupQuotaSubscriptionAllocationRequestServer contains the fakes for client GroupQuotaSubscriptionAllocationRequestClient
	GroupQuotaSubscriptionAllocationRequestServer GroupQuotaSubscriptionAllocationRequestServer

	// GroupQuotaSubscriptionRequestsServer contains the fakes for client GroupQuotaSubscriptionRequestsClient
	GroupQuotaSubscriptionRequestsServer GroupQuotaSubscriptionRequestsServer

	// GroupQuotaSubscriptionsServer contains the fakes for client GroupQuotaSubscriptionsClient
	GroupQuotaSubscriptionsServer GroupQuotaSubscriptionsServer

	// GroupQuotasServer contains the fakes for client GroupQuotasClient
	GroupQuotasServer GroupQuotasServer

	// OperationServer contains the fakes for client OperationClient
	OperationServer OperationServer

	// RequestStatusServer contains the fakes for client RequestStatusClient
	RequestStatusServer RequestStatusServer

	// UsagesServer contains the fakes for client UsagesClient
	UsagesServer UsagesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armquota.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armquota.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                             *ServerFactory
	trMu                                            sync.Mutex
	trServer                                        *ServerTransport
	trGroupQuotaLimitsServer                        *GroupQuotaLimitsServerTransport
	trGroupQuotaLimitsRequestServer                 *GroupQuotaLimitsRequestServerTransport
	trGroupQuotaSubscriptionAllocationServer        *GroupQuotaSubscriptionAllocationServerTransport
	trGroupQuotaSubscriptionAllocationRequestServer *GroupQuotaSubscriptionAllocationRequestServerTransport
	trGroupQuotaSubscriptionRequestsServer          *GroupQuotaSubscriptionRequestsServerTransport
	trGroupQuotaSubscriptionsServer                 *GroupQuotaSubscriptionsServerTransport
	trGroupQuotasServer                             *GroupQuotasServerTransport
	trOperationServer                               *OperationServerTransport
	trRequestStatusServer                           *RequestStatusServerTransport
	trUsagesServer                                  *UsagesServerTransport
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
	case "Client":
		initServer(s, &s.trServer, func() *ServerTransport { return NewServerTransport(&s.srv.Server) })
		resp, err = s.trServer.Do(req)
	case "GroupQuotaLimitsClient":
		initServer(s, &s.trGroupQuotaLimitsServer, func() *GroupQuotaLimitsServerTransport {
			return NewGroupQuotaLimitsServerTransport(&s.srv.GroupQuotaLimitsServer)
		})
		resp, err = s.trGroupQuotaLimitsServer.Do(req)
	case "GroupQuotaLimitsRequestClient":
		initServer(s, &s.trGroupQuotaLimitsRequestServer, func() *GroupQuotaLimitsRequestServerTransport {
			return NewGroupQuotaLimitsRequestServerTransport(&s.srv.GroupQuotaLimitsRequestServer)
		})
		resp, err = s.trGroupQuotaLimitsRequestServer.Do(req)
	case "GroupQuotaSubscriptionAllocationClient":
		initServer(s, &s.trGroupQuotaSubscriptionAllocationServer, func() *GroupQuotaSubscriptionAllocationServerTransport {
			return NewGroupQuotaSubscriptionAllocationServerTransport(&s.srv.GroupQuotaSubscriptionAllocationServer)
		})
		resp, err = s.trGroupQuotaSubscriptionAllocationServer.Do(req)
	case "GroupQuotaSubscriptionAllocationRequestClient":
		initServer(s, &s.trGroupQuotaSubscriptionAllocationRequestServer, func() *GroupQuotaSubscriptionAllocationRequestServerTransport {
			return NewGroupQuotaSubscriptionAllocationRequestServerTransport(&s.srv.GroupQuotaSubscriptionAllocationRequestServer)
		})
		resp, err = s.trGroupQuotaSubscriptionAllocationRequestServer.Do(req)
	case "GroupQuotaSubscriptionRequestsClient":
		initServer(s, &s.trGroupQuotaSubscriptionRequestsServer, func() *GroupQuotaSubscriptionRequestsServerTransport {
			return NewGroupQuotaSubscriptionRequestsServerTransport(&s.srv.GroupQuotaSubscriptionRequestsServer)
		})
		resp, err = s.trGroupQuotaSubscriptionRequestsServer.Do(req)
	case "GroupQuotaSubscriptionsClient":
		initServer(s, &s.trGroupQuotaSubscriptionsServer, func() *GroupQuotaSubscriptionsServerTransport {
			return NewGroupQuotaSubscriptionsServerTransport(&s.srv.GroupQuotaSubscriptionsServer)
		})
		resp, err = s.trGroupQuotaSubscriptionsServer.Do(req)
	case "GroupQuotasClient":
		initServer(s, &s.trGroupQuotasServer, func() *GroupQuotasServerTransport { return NewGroupQuotasServerTransport(&s.srv.GroupQuotasServer) })
		resp, err = s.trGroupQuotasServer.Do(req)
	case "OperationClient":
		initServer(s, &s.trOperationServer, func() *OperationServerTransport { return NewOperationServerTransport(&s.srv.OperationServer) })
		resp, err = s.trOperationServer.Do(req)
	case "RequestStatusClient":
		initServer(s, &s.trRequestStatusServer, func() *RequestStatusServerTransport {
			return NewRequestStatusServerTransport(&s.srv.RequestStatusServer)
		})
		resp, err = s.trRequestStatusServer.Do(req)
	case "UsagesClient":
		initServer(s, &s.trUsagesServer, func() *UsagesServerTransport { return NewUsagesServerTransport(&s.srv.UsagesServer) })
		resp, err = s.trUsagesServer.Do(req)
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
