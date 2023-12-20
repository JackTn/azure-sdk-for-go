//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// AgentPoolsServer is a fake server for instances of the armcontainerservice.AgentPoolsClient type.
type AgentPoolsServer struct {
	// BeginAbortLatestOperation is the fake for method AgentPoolsClient.BeginAbortLatestOperation
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginAbortLatestOperation func(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *armcontainerservice.AgentPoolsClientBeginAbortLatestOperationOptions) (resp azfake.PollerResponder[armcontainerservice.AgentPoolsClientAbortLatestOperationResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method AgentPoolsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters armcontainerservice.AgentPool, options *armcontainerservice.AgentPoolsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcontainerservice.AgentPoolsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AgentPoolsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *armcontainerservice.AgentPoolsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerservice.AgentPoolsClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDeleteMachines is the fake for method AgentPoolsClient.BeginDeleteMachines
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginDeleteMachines func(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, machines armcontainerservice.AgentPoolDeleteMachinesParameter, options *armcontainerservice.AgentPoolsClientBeginDeleteMachinesOptions) (resp azfake.PollerResponder[armcontainerservice.AgentPoolsClientDeleteMachinesResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AgentPoolsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *armcontainerservice.AgentPoolsClientGetOptions) (resp azfake.Responder[armcontainerservice.AgentPoolsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAvailableAgentPoolVersions is the fake for method AgentPoolsClient.GetAvailableAgentPoolVersions
	// HTTP status codes to indicate success: http.StatusOK
	GetAvailableAgentPoolVersions func(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.AgentPoolsClientGetAvailableAgentPoolVersionsOptions) (resp azfake.Responder[armcontainerservice.AgentPoolsClientGetAvailableAgentPoolVersionsResponse], errResp azfake.ErrorResponder)

	// GetUpgradeProfile is the fake for method AgentPoolsClient.GetUpgradeProfile
	// HTTP status codes to indicate success: http.StatusOK
	GetUpgradeProfile func(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *armcontainerservice.AgentPoolsClientGetUpgradeProfileOptions) (resp azfake.Responder[armcontainerservice.AgentPoolsClientGetUpgradeProfileResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AgentPoolsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, resourceName string, options *armcontainerservice.AgentPoolsClientListOptions) (resp azfake.PagerResponder[armcontainerservice.AgentPoolsClientListResponse])

	// BeginUpgradeNodeImageVersion is the fake for method AgentPoolsClient.BeginUpgradeNodeImageVersion
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpgradeNodeImageVersion func(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *armcontainerservice.AgentPoolsClientBeginUpgradeNodeImageVersionOptions) (resp azfake.PollerResponder[armcontainerservice.AgentPoolsClientUpgradeNodeImageVersionResponse], errResp azfake.ErrorResponder)
}

// NewAgentPoolsServerTransport creates a new instance of AgentPoolsServerTransport with the provided implementation.
// The returned AgentPoolsServerTransport instance is connected to an instance of armcontainerservice.AgentPoolsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAgentPoolsServerTransport(srv *AgentPoolsServer) *AgentPoolsServerTransport {
	return &AgentPoolsServerTransport{
		srv:                          srv,
		beginAbortLatestOperation:    newTracker[azfake.PollerResponder[armcontainerservice.AgentPoolsClientAbortLatestOperationResponse]](),
		beginCreateOrUpdate:          newTracker[azfake.PollerResponder[armcontainerservice.AgentPoolsClientCreateOrUpdateResponse]](),
		beginDelete:                  newTracker[azfake.PollerResponder[armcontainerservice.AgentPoolsClientDeleteResponse]](),
		beginDeleteMachines:          newTracker[azfake.PollerResponder[armcontainerservice.AgentPoolsClientDeleteMachinesResponse]](),
		newListPager:                 newTracker[azfake.PagerResponder[armcontainerservice.AgentPoolsClientListResponse]](),
		beginUpgradeNodeImageVersion: newTracker[azfake.PollerResponder[armcontainerservice.AgentPoolsClientUpgradeNodeImageVersionResponse]](),
	}
}

// AgentPoolsServerTransport connects instances of armcontainerservice.AgentPoolsClient to instances of AgentPoolsServer.
// Don't use this type directly, use NewAgentPoolsServerTransport instead.
type AgentPoolsServerTransport struct {
	srv                          *AgentPoolsServer
	beginAbortLatestOperation    *tracker[azfake.PollerResponder[armcontainerservice.AgentPoolsClientAbortLatestOperationResponse]]
	beginCreateOrUpdate          *tracker[azfake.PollerResponder[armcontainerservice.AgentPoolsClientCreateOrUpdateResponse]]
	beginDelete                  *tracker[azfake.PollerResponder[armcontainerservice.AgentPoolsClientDeleteResponse]]
	beginDeleteMachines          *tracker[azfake.PollerResponder[armcontainerservice.AgentPoolsClientDeleteMachinesResponse]]
	newListPager                 *tracker[azfake.PagerResponder[armcontainerservice.AgentPoolsClientListResponse]]
	beginUpgradeNodeImageVersion *tracker[azfake.PollerResponder[armcontainerservice.AgentPoolsClientUpgradeNodeImageVersionResponse]]
}

// Do implements the policy.Transporter interface for AgentPoolsServerTransport.
func (a *AgentPoolsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AgentPoolsClient.BeginAbortLatestOperation":
		resp, err = a.dispatchBeginAbortLatestOperation(req)
	case "AgentPoolsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AgentPoolsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AgentPoolsClient.BeginDeleteMachines":
		resp, err = a.dispatchBeginDeleteMachines(req)
	case "AgentPoolsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AgentPoolsClient.GetAvailableAgentPoolVersions":
		resp, err = a.dispatchGetAvailableAgentPoolVersions(req)
	case "AgentPoolsClient.GetUpgradeProfile":
		resp, err = a.dispatchGetUpgradeProfile(req)
	case "AgentPoolsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AgentPoolsClient.BeginUpgradeNodeImageVersion":
		resp, err = a.dispatchBeginUpgradeNodeImageVersion(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchBeginAbortLatestOperation(req *http.Request) (*http.Response, error) {
	if a.srv.BeginAbortLatestOperation == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginAbortLatestOperation not implemented")}
	}
	beginAbortLatestOperation := a.beginAbortLatestOperation.get(req)
	if beginAbortLatestOperation == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedclusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/abort`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginAbortLatestOperation(req.Context(), resourceGroupNameParam, resourceNameParam, agentPoolNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginAbortLatestOperation = &respr
		a.beginAbortLatestOperation.add(req, beginAbortLatestOperation)
	}

	resp, err := server.PollerResponderNext(beginAbortLatestOperation, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginAbortLatestOperation.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginAbortLatestOperation) {
		a.beginAbortLatestOperation.remove(req)
	}

	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerservice.AgentPool](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, agentPoolNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
		if err != nil {
			return nil, err
		}
		ignorePodDisruptionBudgetUnescaped, err := url.QueryUnescape(qp.Get("ignore-pod-disruption-budget"))
		if err != nil {
			return nil, err
		}
		ignorePodDisruptionBudgetParam, err := parseOptional(ignorePodDisruptionBudgetUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armcontainerservice.AgentPoolsClientBeginDeleteOptions
		if ignorePodDisruptionBudgetParam != nil {
			options = &armcontainerservice.AgentPoolsClientBeginDeleteOptions{
				IgnorePodDisruptionBudget: ignorePodDisruptionBudgetParam,
			}
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, resourceNameParam, agentPoolNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchBeginDeleteMachines(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDeleteMachines == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteMachines not implemented")}
	}
	beginDeleteMachines := a.beginDeleteMachines.get(req)
	if beginDeleteMachines == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deleteMachines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerservice.AgentPoolDeleteMachinesParameter](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDeleteMachines(req.Context(), resourceGroupNameParam, resourceNameParam, agentPoolNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeleteMachines = &respr
		a.beginDeleteMachines.add(req, beginDeleteMachines)
	}

	resp, err := server.PollerResponderNext(beginDeleteMachines, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		a.beginDeleteMachines.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeleteMachines) {
		a.beginDeleteMachines.remove(req)
	}

	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, agentPoolNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AgentPool, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchGetAvailableAgentPoolVersions(req *http.Request) (*http.Response, error) {
	if a.srv.GetAvailableAgentPoolVersions == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAvailableAgentPoolVersions not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/availableAgentPoolVersions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetAvailableAgentPoolVersions(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AgentPoolAvailableVersions, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchGetUpgradeProfile(req *http.Request) (*http.Response, error) {
	if a.srv.GetUpgradeProfile == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUpgradeProfile not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/upgradeProfiles/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetUpgradeProfile(req.Context(), resourceGroupNameParam, resourceNameParam, agentPoolNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AgentPoolUpgradeProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, resourceNameParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcontainerservice.AgentPoolsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchBeginUpgradeNodeImageVersion(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpgradeNodeImageVersion == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpgradeNodeImageVersion not implemented")}
	}
	beginUpgradeNodeImageVersion := a.beginUpgradeNodeImageVersion.get(req)
	if beginUpgradeNodeImageVersion == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/upgradeNodeImageVersion`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpgradeNodeImageVersion(req.Context(), resourceGroupNameParam, resourceNameParam, agentPoolNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpgradeNodeImageVersion = &respr
		a.beginUpgradeNodeImageVersion.add(req, beginUpgradeNodeImageVersion)
	}

	resp, err := server.PollerResponderNext(beginUpgradeNodeImageVersion, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpgradeNodeImageVersion.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpgradeNodeImageVersion) {
		a.beginUpgradeNodeImageVersion.remove(req)
	}

	return resp, nil
}
