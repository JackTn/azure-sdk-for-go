//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ReservationRecommendationsClient contains the methods for the ReservationRecommendations group.
// Don't use this type directly, use NewReservationRecommendationsClient() instead.
type ReservationRecommendationsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewReservationRecommendationsClient creates a new instance of ReservationRecommendationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReservationRecommendationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ReservationRecommendationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ReservationRecommendationsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - List of recommendations for purchasing reserved instances.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// resourceScope - The scope associated with reservation recommendations operations. This includes '/subscriptions/{subscriptionId}/'
// for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resource group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}'
// for BillingAccount scope, and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile
// scope
// options - ReservationRecommendationsClientListOptions contains the optional parameters for the ReservationRecommendationsClient.List
// method.
func (client *ReservationRecommendationsClient) NewListPager(resourceScope string, options *ReservationRecommendationsClientListOptions) *runtime.Pager[ReservationRecommendationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationRecommendationsClientListResponse]{
		More: func(page ReservationRecommendationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationRecommendationsClientListResponse) (ReservationRecommendationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceScope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReservationRecommendationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReservationRecommendationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReservationRecommendationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReservationRecommendationsClient) listCreateRequest(ctx context.Context, resourceScope string, options *ReservationRecommendationsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Consumption/reservationRecommendations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReservationRecommendationsClient) listHandleResponse(resp *http.Response) (ReservationRecommendationsClientListResponse, error) {
	result := ReservationRecommendationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationRecommendationsListResult); err != nil {
		return ReservationRecommendationsClientListResponse{}, err
	}
	return result, nil
}
