// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"context"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type azCoreClient struct {
	host     string
	pipeline runtime.Pipeline
}

func NewAzCoreClient(tenantId, userAgent string) AzureClient {
	credentials, err := newAzureCredentials()
	if err != nil {
		panic(err)
	}

	pipeline, err := armruntime.NewPipeline("pulumi-azure-native", "v2", credentials,
		runtime.PipelineOptions{}, &arm.ClientOptions{})
	if err != nil {
		panic(err)
	}

	return &azCoreClient{
		host:     cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint,
		pipeline: pipeline,
	}
}

func (c *azCoreClient) Get(ctx context.Context, id string, apiVersion string) (any, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(c.host, id))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	resp, err := c.pipeline.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}

	var responseBody map[string]interface{}
	if err := runtime.UnmarshalAsJSON(resp, &responseBody); err != nil {
		return nil, err
	}
	return responseBody, nil
}

func (c *azCoreClient) Delete(ctx context.Context, id, apiVersion, asyncStyle string,
	queryParams map[string]any) error {
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(c.host, id))
	if err != nil {
		return err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	resp, err := c.pipeline.Do(req)
	if err != nil {
		return err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return runtime.NewResponseError(resp)
	}
	pt, err := runtime.NewPoller[interface{}](resp, c.pipeline, nil)
	if err != nil {
		return err
	}
	_, err = pt.PollUntilDone(ctx, &runtime.PollUntilDoneOptions{
		Frequency: 10 * time.Second,
	})
	return err
}

func (c *azCoreClient) Put(ctx context.Context, id string, bodyProps map[string]interface{},
	queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(c.host, id))
	if err != nil {
		return nil, false, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", queryParameters["api-version"].(string))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	err = runtime.MarshalAsJSON(req, bodyProps)
	if err != nil {
		return nil, false, err
	}
	resp, err := c.pipeline.Do(req)
	if err != nil {
		return nil, false, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, false, runtime.NewResponseError(resp)
	}
	pt, err := runtime.NewPoller[map[string]interface{}](resp, c.pipeline, nil)
	if err != nil {
		return nil, false, err
	}
	pollResp, err := pt.PollUntilDone(ctx, &runtime.PollUntilDoneOptions{
		Frequency: 10 * time.Second,
	})
	if err != nil {
		return nil, true, err
	}
	return pollResp, true, nil
}

func (c *azCoreClient) Post(ctx context.Context, id string, bodyProps map[string]interface{},
	queryParameters map[string]interface{}) (any, error) {
	panic("Post not implemented")
}

func (c *azCoreClient) Head(ctx context.Context, id string, apiVersion string) error {
	panic("Head not implemented")
}

func (c *azCoreClient) Patch(ctx context.Context, id string, bodyProps map[string]interface{},
	queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error) {
	panic("Patch not implemented")
}

// CanCreate asserts that a resource with a given ID and API version can be created
// or returns an error otherwise.
func (c *azCoreClient) CanCreate(ctx context.Context, id, path, apiVersion, readMethod string,
	isSingletonResource, hasDefaultBody bool, isDefaultResponse func(map[string]any) bool) error {
	// TODO: not implemented
	return nil
}

func newAzureCredentials() (*azidentity.ChainedTokenCredential, error) {
	cli, err := azidentity.NewAzureCLICredential(&azidentity.AzureCLICredentialOptions{})
	if err != nil {
		return nil, err
	}

	sources := []azcore.TokenCredential{cli}
	chain, err := azidentity.NewChainedTokenCredential(sources, nil)
	if err != nil {
		return nil, err
	}

	return chain, nil
}
