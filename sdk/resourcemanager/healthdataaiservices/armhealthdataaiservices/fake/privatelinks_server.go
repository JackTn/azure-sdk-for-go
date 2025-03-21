// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinksServer is a fake server for instances of the armhealthdataaiservices.PrivateLinksClient type.
type PrivateLinksServer struct {
	// NewListByDeidServicePager is the fake for method PrivateLinksClient.NewListByDeidServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDeidServicePager func(resourceGroupName string, deidServiceName string, options *armhealthdataaiservices.PrivateLinksClientListByDeidServiceOptions) (resp azfake.PagerResponder[armhealthdataaiservices.PrivateLinksClientListByDeidServiceResponse])
}

// NewPrivateLinksServerTransport creates a new instance of PrivateLinksServerTransport with the provided implementation.
// The returned PrivateLinksServerTransport instance is connected to an instance of armhealthdataaiservices.PrivateLinksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinksServerTransport(srv *PrivateLinksServer) *PrivateLinksServerTransport {
	return &PrivateLinksServerTransport{
		srv:                       srv,
		newListByDeidServicePager: newTracker[azfake.PagerResponder[armhealthdataaiservices.PrivateLinksClientListByDeidServiceResponse]](),
	}
}

// PrivateLinksServerTransport connects instances of armhealthdataaiservices.PrivateLinksClient to instances of PrivateLinksServer.
// Don't use this type directly, use NewPrivateLinksServerTransport instead.
type PrivateLinksServerTransport struct {
	srv                       *PrivateLinksServer
	newListByDeidServicePager *tracker[azfake.PagerResponder[armhealthdataaiservices.PrivateLinksClientListByDeidServiceResponse]]
}

// Do implements the policy.Transporter interface for PrivateLinksServerTransport.
func (p *PrivateLinksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateLinksServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateLinksServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateLinksServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateLinksClient.NewListByDeidServicePager":
				res.resp, res.err = p.dispatchNewListByDeidServicePager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (p *PrivateLinksServerTransport) dispatchNewListByDeidServicePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByDeidServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDeidServicePager not implemented")}
	}
	newListByDeidServicePager := p.newListByDeidServicePager.get(req)
	if newListByDeidServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthDataAIServices/deidServices/(?P<deidServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		deidServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deidServiceName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByDeidServicePager(resourceGroupNameParam, deidServiceNameParam, nil)
		newListByDeidServicePager = &resp
		p.newListByDeidServicePager.add(req, newListByDeidServicePager)
		server.PagerResponderInjectNextLinks(newListByDeidServicePager, req, func(page *armhealthdataaiservices.PrivateLinksClientListByDeidServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDeidServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByDeidServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDeidServicePager) {
		p.newListByDeidServicePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateLinksServerTransport
var privateLinksServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
