// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualMachineScaleSetVMExtensionsServer is a fake server for instances of the armcompute.VirtualMachineScaleSetVMExtensionsClient type.
type VirtualMachineScaleSetVMExtensionsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualMachineScaleSetVMExtensionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters armcompute.VirtualMachineScaleSetVMExtension, options *armcompute.VirtualMachineScaleSetVMExtensionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualMachineScaleSetVMExtensionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *armcompute.VirtualMachineScaleSetVMExtensionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachineScaleSetVMExtensionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *armcompute.VirtualMachineScaleSetVMExtensionsClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineScaleSetVMExtensionsClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method VirtualMachineScaleSetVMExtensionsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, options *armcompute.VirtualMachineScaleSetVMExtensionsClientListOptions) (resp azfake.Responder[armcompute.VirtualMachineScaleSetVMExtensionsClientListResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method VirtualMachineScaleSetVMExtensionsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters armcompute.VirtualMachineScaleSetVMExtensionUpdate, options *armcompute.VirtualMachineScaleSetVMExtensionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineScaleSetVMExtensionsServerTransport creates a new instance of VirtualMachineScaleSetVMExtensionsServerTransport with the provided implementation.
// The returned VirtualMachineScaleSetVMExtensionsServerTransport instance is connected to an instance of armcompute.VirtualMachineScaleSetVMExtensionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineScaleSetVMExtensionsServerTransport(srv *VirtualMachineScaleSetVMExtensionsServer) *VirtualMachineScaleSetVMExtensionsServerTransport {
	return &VirtualMachineScaleSetVMExtensionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientDeleteResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientUpdateResponse]](),
	}
}

// VirtualMachineScaleSetVMExtensionsServerTransport connects instances of armcompute.VirtualMachineScaleSetVMExtensionsClient to instances of VirtualMachineScaleSetVMExtensionsServer.
// Don't use this type directly, use NewVirtualMachineScaleSetVMExtensionsServerTransport instead.
type VirtualMachineScaleSetVMExtensionsServerTransport struct {
	srv                 *VirtualMachineScaleSetVMExtensionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientDeleteResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for VirtualMachineScaleSetVMExtensionsServerTransport.
func (v *VirtualMachineScaleSetVMExtensionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if virtualMachineScaleSetVMExtensionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = virtualMachineScaleSetVMExtensionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VirtualMachineScaleSetVMExtensionsClient.BeginCreateOrUpdate":
				res.resp, res.err = v.dispatchBeginCreateOrUpdate(req)
			case "VirtualMachineScaleSetVMExtensionsClient.BeginDelete":
				res.resp, res.err = v.dispatchBeginDelete(req)
			case "VirtualMachineScaleSetVMExtensionsClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "VirtualMachineScaleSetVMExtensionsClient.List":
				res.resp, res.err = v.dispatchList(req)
			case "VirtualMachineScaleSetVMExtensionsClient.BeginUpdate":
				res.resp, res.err = v.dispatchBeginUpdate(req)
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

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<vmExtensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineScaleSetVMExtension](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		instanceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
		if err != nil {
			return nil, err
		}
		vmExtensionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmExtensionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, vmScaleSetNameParam, instanceIDParam, vmExtensionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<vmExtensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		instanceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
		if err != nil {
			return nil, err
		}
		vmExtensionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmExtensionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, vmScaleSetNameParam, instanceIDParam, vmExtensionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<vmExtensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmScaleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
	if err != nil {
		return nil, err
	}
	instanceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
	if err != nil {
		return nil, err
	}
	vmExtensionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmExtensionName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armcompute.VirtualMachineScaleSetVMExtensionsClientGetOptions
	if expandParam != nil {
		options = &armcompute.VirtualMachineScaleSetVMExtensionsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, vmScaleSetNameParam, instanceIDParam, vmExtensionNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineScaleSetVMExtension, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if v.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmScaleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
	if err != nil {
		return nil, err
	}
	instanceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armcompute.VirtualMachineScaleSetVMExtensionsClientListOptions
	if expandParam != nil {
		options = &armcompute.VirtualMachineScaleSetVMExtensionsClientListOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.List(req.Context(), resourceGroupNameParam, vmScaleSetNameParam, instanceIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineScaleSetVMExtensionsListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := v.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<vmExtensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineScaleSetVMExtensionUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		instanceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
		if err != nil {
			return nil, err
		}
		vmExtensionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmExtensionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceGroupNameParam, vmScaleSetNameParam, instanceIDParam, vmExtensionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		v.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		v.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to VirtualMachineScaleSetVMExtensionsServerTransport
var virtualMachineScaleSetVMExtensionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
