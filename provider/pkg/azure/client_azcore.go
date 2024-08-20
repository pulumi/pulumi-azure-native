// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

type azCoreClient struct {
	host      string
	pipeline  runtime.Pipeline
	userAgent string
}

func NewAzCoreClient(tokenCredential azcore.TokenCredential, userAgent string) AzureClient {
	// TODO,tkappler version
	pipeline, err := armruntime.NewPipeline("pulumi-azure-native", "v2", tokenCredential,
		runtime.PipelineOptions{}, &arm.ClientOptions{})
	if err != nil {
		panic(err)
	}

	return &azCoreClient{
		host:      cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint, // TODO,tkappler other clouds
		pipeline:  pipeline,
		userAgent: userAgent,
	}
}

func (c *azCoreClient) setHeaders(req *policy.Request) {
	req.Raw().Header.Set("Accept", "application/json")
	req.Raw().Header.Set("User-Agent", c.userAgent)
}

func (c *azCoreClient) Get(ctx context.Context, id string, apiVersion string) (any, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(c.host, id))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	c.setHeaders(req)
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

// TODO,tkappler asyncStyle
func (c *azCoreClient) Delete(ctx context.Context, id, apiVersion, asyncStyle string,
	queryParams map[string]any) error {
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(c.host, id))
	if err != nil {
		return err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	c.setHeaders(req)
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

// TODO,tkappler asyncStyle
func (c *azCoreClient) Put(ctx context.Context, id string, bodyProps map[string]interface{},
	queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(c.host, id))
	if err != nil {
		return nil, false, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", queryParameters["api-version"].(string))
	req.Raw().URL.RawQuery = reqQP.Encode()
	c.setHeaders(req)
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(c.host, id))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", queryParameters["api-version"].(string))
	req.Raw().URL.RawQuery = reqQP.Encode()
	c.setHeaders(req)

	err = runtime.MarshalAsJSON(req, bodyProps)
	if err != nil {
		return nil, err
	}

	resp, err := c.pipeline.Do(req)
	if err != nil {
		return nil, err
	}

	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}

	var responseBody map[string]interface{}
	err = runtime.UnmarshalAsJSON(resp, &responseBody)
	return responseBody, err
}

func (c *azCoreClient) Head(ctx context.Context, id string, apiVersion string) error {
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(c.host, id))
	if err != nil {
		return err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	c.setHeaders(req)

	resp, err := c.pipeline.Do(req)
	if err != nil {
		return err
	}

	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return runtime.NewResponseError(resp)
	}
	return nil
}

// TODO,tkappler asyncStyle
// TODO,tkappler almost identical to Put
func (c *azCoreClient) Patch(ctx context.Context, id string, bodyProps map[string]interface{},
	queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(c.host, id))
	if err != nil {
		return nil, false, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", queryParameters["api-version"].(string))
	req.Raw().URL.RawQuery = reqQP.Encode()
	c.setHeaders(req)
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

// CanCreate asserts that a resource with a given ID and API version can be created
// or returns an error otherwise.
func (c *azCoreClient) CanCreate(ctx context.Context, id, path, apiVersion, readMethod string,
	isSingletonResource, hasDefaultBody bool, isDefaultResponse func(map[string]any) bool,
) error {
	req, err := runtime.NewRequest(ctx, readMethod, runtime.JoinPaths(c.host, id, path))
	if err != nil {
		return err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	c.setHeaders(req)
	resp, err := c.pipeline.Do(req)
	if err != nil {
		return err
	}

	switch {
	case http.StatusOK == resp.StatusCode && isSingletonResource:
		// Singleton resources always exist, so OK is expected.
		return nil
	case http.StatusOK == resp.StatusCode && hasDefaultBody:
		// This resource is automatically created with a parent and set to its default state. It can be deleted though.
		// Validate that its current shape is in the default state to avoid unintended adoption and destructive
		// actions.
		// NOTE: We may reconsider and relax this restriction when we get more examples of such resources.
		// The difference between "take any singleton resource as-is" and "require default body for deletable resources"
		// isn't very principled but is based on what subjectively feels best for the current examples.
		var outputs map[string]interface{}
		if err := runtime.UnmarshalAsJSON(resp, &outputs); err != nil {
			return err
		}
		if !isDefaultResponse(outputs) {
			return fmt.Errorf("cannot create already existing subresource '%s'", id)
		}
		return nil
	case http.StatusNoContent == resp.StatusCode:
		if readMethod == "HEAD" {
			return fmt.Errorf("cannot create already existing resource '%s'", id)
		}
		// A few "linking" resources, like private endpoint connections, return 204 as "does not exist" status code.
		// Treat them as such unless it's a HEAD method treated above.
		return nil
	case http.StatusOK == resp.StatusCode:
		// Usually, 200 means that the resource already exists and we shouldn't try to create it.
		// However, unfortunately, some APIs return 200 with an empty body for non-existing resources.
		// Our strategy here is to try to parse the response body and see if it's a valid non-empty JSON.
		// If it is, we assume the resource exists.
		var outputs map[string]interface{}
		err := runtime.UnmarshalAsJSON(resp, &outputs)
		if err == nil && len(outputs) > 0 {
			return fmt.Errorf("cannot create already existing resource '%s'", id)
		}
		return nil
	case http.StatusNotFound == resp.StatusCode:
		return nil
	default:
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("cannot check existence of resource '%s': status code %d, %s", id, resp.StatusCode, body)
	}
}
