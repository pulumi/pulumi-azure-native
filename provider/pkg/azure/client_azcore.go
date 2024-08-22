// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

type azCoreClient struct {
	host      string
	pipeline  runtime.Pipeline
	userAgent string
}

func NewAzCoreClient(tokenCredential azcore.TokenCredential, userAgent string, azureCloud cloud.Configuration) AzureClient {
	pipeline, err := armruntime.NewPipeline("pulumi-azure-native", version.Version, tokenCredential,
		runtime.PipelineOptions{}, &arm.ClientOptions{})
	if err != nil {
		panic(err)
	}

	return &azCoreClient{
		host:      azureCloud.Services[cloud.ResourceManager].Endpoint,
		pipeline:  pipeline,
		userAgent: userAgent,
	}
}

func (c *azCoreClient) setHeaders(req *policy.Request) {
	req.Raw().Header.Set("Accept", "application/json")
	req.Raw().Header.Set("User-Agent", c.userAgent)
}

func (c *azCoreClient) initRequest(ctx context.Context, method, id, apiVersion string) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, method, runtime.JoinPaths(c.host, id))
	if err != nil {
		return nil, err
	}

	c.setHeaders(req)

	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()

	return req, nil
}

func (c *azCoreClient) Get(ctx context.Context, id string, apiVersion string) (any, error) {
	req, err := c.initRequest(ctx, http.MethodGet, id, apiVersion)
	if err != nil {
		return nil, err
	}

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
	req, err := c.initRequest(ctx, http.MethodDelete, id, apiVersion)
	if err != nil {
		return err
	}

	resp, err := c.pipeline.Do(req)
	if err != nil {
		return err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent, http.StatusNotFound) {
		return runtime.NewResponseError(resp)
	}

	// Some APIs are explicitly marked `x-ms-long-running-operation` and we should only do the
	// Future+WaitForCompletion+GetResult steps in that case.
	if asyncStyle != "" {
		pt, err := runtime.NewPoller[any](resp, c.pipeline, nil)
		if err != nil {
			return err
		}
		_, err = pt.PollUntilDone(ctx, &runtime.PollUntilDoneOptions{
			Frequency: 10 * time.Second,
		})
		if err != nil {
			respErr := err.(*azcore.ResponseError)
			if resp.StatusCode == 202 && respErr.StatusCode == 404 && strings.Contains(respErr.Error(), "ResourceNotFound") {
				// Consider this specific error to be a success of deletion.
				// Work around https://github.com/pulumi/pulumi-azure-nextgen/issues/120
				// Upstream fix is tracked in https://github.com/Azure/go-autorest/issues/596
				return nil
			}
		}
	}
	return err
}

func (c *azCoreClient) Put(ctx context.Context, id string, bodyProps map[string]interface{},
	queryParameters map[string]interface{}, asyncStyle string,
) (map[string]interface{}, bool, error) {
	return c.putOrPatch(ctx, http.MethodPut, id, bodyProps, queryParameters, asyncStyle)
}

func (c *azCoreClient) Patch(ctx context.Context, id string, bodyProps map[string]interface{},
	queryParameters map[string]interface{}, asyncStyle string,
) (map[string]interface{}, bool, error) {
	return c.putOrPatch(ctx, http.MethodPatch, id, bodyProps, queryParameters, asyncStyle)
}

func (c *azCoreClient) putOrPatch(ctx context.Context, method string, id string, bodyProps map[string]any,
	queryParameters map[string]any, asyncStyle string,
) (map[string]any, bool, error) {
	if method != http.MethodPut && method != http.MethodPatch {
		return nil, false, fmt.Errorf("method must be PUT or PATCH, got %s. Please report this issue", method)
	}

	apiVersion := queryParameters["api-version"].(string)

	req, err := c.initRequest(ctx, method, id, apiVersion)
	if err != nil {
		return nil, false, err
	}

	if bodyProps != nil {
		err = runtime.MarshalAsJSON(req, bodyProps)
		if err != nil {
			return nil, false, err
		}
	}

	resp, err := c.pipeline.Do(req)
	if err != nil {
		return nil, false, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, false, runtime.NewResponseError(resp)
	}

	var outputs map[string]any
	if err := runtime.UnmarshalAsJSON(resp, &outputs); err != nil {
		return nil, false, err
	}

	// Some APIs are explicitly marked `x-ms-long-running-operation` and we are only supposed to do the
	// Future+WaitForCompletion+GetResult steps in that case. However, if we get 202, we don't want to
	// consider this a failure - so try following the awaiting protocol in case the service hasn't marked
	// its API as long-running by an oversight.
	created := false
	if asyncStyle != "" || resp.StatusCode == http.StatusAccepted {
		// We have now created a resource. It is very important to ensure that from this point on,
		// any other error below returns the ID using the `pulumirpc.ErrorResourceInitFailed` error
		// details annotation. Otherwise, the resource is leaked. We ensure that we wrap any await
		// errors as a partial error to the RPC.
		created = resp.StatusCode < 400

		normalizeLocationHeader(c.host, apiVersion, resp.Header)

		pt, err := runtime.NewPoller[map[string]interface{}](resp, c.pipeline, nil)
		if err != nil {
			return nil, created, err
		}
		outputs, err = pt.PollUntilDone(ctx, &runtime.PollUntilDoneOptions{
			Frequency: 10 * time.Second,
		})
		if err != nil {
			return nil, created, err
		}
	}

	return outputs, true, nil
}

// The azcore location header Poller validates that the location header is a valid absolute URL.
// In some cases, like user-assigned identities, the location header in the PUT response is a
// relative URL "/subscriptions/...", which we try to make absolute.
// Also, add the api version query param if missing.
func normalizeLocationHeader(host, apiVersion string, headers http.Header) {
	locUrlStr := headers.Get("Location")
	if locUrlStr == "" {
		return
	}

	locUrl, err := url.Parse(locUrlStr)
	if err != nil {
		logging.V(3).Infof("Location header '%s' is not a valid URL: %v", locUrlStr, err)
		return
	}

	if !locUrl.IsAbs() {
		locUrl.Host = host
		absUrlStr := runtime.JoinPaths(host, locUrlStr)
		absUrl, err := url.Parse(absUrlStr)
		if err != nil {
			logging.V(3).Infof("Location header '%s' is not an absolute URL, failed to make absolute: %s", locUrlStr, absUrlStr)
		} else {
			locUrl = absUrl
			logging.V(9).Infof("Location header '%s' is not an absolute URL, added host '%s'", locUrlStr, host)
		}
	}

	if locUrl.Query().Get("api-version") == "" {
		q := locUrl.Query()
		q.Add("api-version", apiVersion)
		locUrl.RawQuery = q.Encode()
	}

	headers.Set("Location", locUrl.String())
}

func (c *azCoreClient) Post(ctx context.Context, id string, bodyProps map[string]interface{},
	queryParameters map[string]interface{}) (any, error) {
	req, err := c.initRequest(ctx, http.MethodPost, id, queryParameters["api-version"].(string))
	if err != nil {
		return nil, err
	}

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
	req, err := c.initRequest(ctx, http.MethodHead, id, apiVersion)
	if err != nil {
		return err
	}

	resp, err := c.pipeline.Do(req)
	if err != nil {
		return err
	}

	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return runtime.NewResponseError(resp)
	}
	return nil
}

// CanCreate asserts that a resource with a given ID and API version can be created
// or returns an error otherwise.
func (c *azCoreClient) CanCreate(ctx context.Context, id, path, apiVersion, readMethod string,
	isSingletonResource, hasDefaultBody bool, isDefaultResponse func(map[string]any) bool,
) error {
	req, err := c.initRequest(ctx, readMethod, path, apiVersion)
	if err != nil {
		return err
	}

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
