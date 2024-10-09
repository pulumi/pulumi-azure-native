// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

type azCoreClient struct {
	host      string
	pipeline  runtime.Pipeline
	userAgent string

	// Exposed internally for tests, to set it at the minimum value for fast tests.
	deletePollingIntervalSeconds int64
	updatePollingIntervalSeconds int64
}

func NewAzCoreClient(tokenCredential azcore.TokenCredential, userAgent string, azureCloud cloud.Configuration, opts *arm.ClientOptions,
) (AzureClient, error) {
	// Hook our logging up to the azcore logger.
	log.SetListener(func(event log.Event, msg string) {
		// Retry logging is very verbose and the number of the retry attempt is already contained
		// in the response event.
		if event != log.EventRetryPolicy {
			logging.V(9).Infof("[azcore] %v: %s", event, msg)
		}
	})

	if opts == nil {
		opts = &arm.ClientOptions{}
	}
	// azcore logging will only happen at log level 9.
	opts.Logging.IncludeBody = true

	// We're trying to configure retries to be similar to the previous Autorest client. The backoff
	// algorithms are hard-coded so they can't be identical.
	// The previous backoff sequence was 30s, 60s, 120s.
	// The new backoff sequence is       20s, 60s, 120s.
	// It's notable, however, that the azcore default backoff sequence is only 1s, 3s, 7s. We might
	// consider moving a bit towards a faster sequence, e.g. 10s, 30s, 70s.
	opts.Retry.RetryDelay = 20 * time.Second
	opts.Retry.MaxRetryDelay = 120 * time.Second

	pipeline, err := armruntime.NewPipeline("pulumi-azure-native", version.Version, tokenCredential,
		runtime.PipelineOptions{}, opts)
	if err != nil {
		return nil, err
	}

	return &azCoreClient{
		host:                         azureCloud.Services[cloud.ResourceManager].Endpoint,
		pipeline:                     pipeline,
		userAgent:                    userAgent,
		deletePollingIntervalSeconds: 30, // same as autorest.DefaultPollingDelay
		updatePollingIntervalSeconds: 10,
	}, nil
}

func (c *azCoreClient) setHeaders(req *policy.Request, contentTypeJson bool) {
	req.Raw().Header.Set("Accept", "application/json")
	req.Raw().Header.Set("User-Agent", c.userAgent)
	if contentTypeJson {
		req.Raw().Header.Set("Content-Type", "application/json; charset=utf-8")
	}
}

func (c *azCoreClient) initRequest(ctx context.Context, method, id string, queryParams map[string]any) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, method, runtime.JoinPaths(c.host, id))
	if err != nil {
		return nil, err
	}

	c.setHeaders(req, method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch)

	urlValues := MapToValues(queryParams)
	// URL-unescape each value before encoding the URL (which encodes all values). Presumably, this
	// prevents double-encoding. It's not clear to me that this is necessary but Autorest does it and
	// we want to match behavior here as closely as possible.
	for key, value := range urlValues {
		for i := range value {
			d, err := url.QueryUnescape(value[i])
			if err != nil {
				return nil, err
			}
			value[i] = d
		}
		urlValues[key] = value
	}
	req.Raw().URL.RawQuery = urlValues.Encode()

	return req, nil
}

func (c *azCoreClient) Get(ctx context.Context, id, apiVersion string, queryParams map[string]any) (any, error) {
	queryParameters := map[string]any{
		"api-version": apiVersion,
	}
	for k, v := range queryParams {
		queryParameters[k] = v
	}

	req, err := c.initRequest(ctx, http.MethodGet, id, queryParameters)
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
		return nil, handleAzCoreResponseError(err, resp)
	}
	return responseBody, nil
}

func (c *azCoreClient) Delete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any,
) error {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	for k, v := range queryParams {
		queryParameters[k] = v
	}

	req, err := c.initRequest(ctx, http.MethodDelete, id, queryParameters)
	if err != nil {
		return err
	}

	resp, err := c.pipeline.Do(req)
	if err != nil {
		return err
	}

	// Some APIs are explicitly marked `x-ms-long-running-operation` and we should only do the
	// poll for the deletion result in that case.
	if asyncStyle != "" {
		pt, err := runtime.NewPoller[any](resp, c.pipeline, nil)
		if err != nil {
			return err
		}
		_, err = pt.PollUntilDone(ctx, &runtime.PollUntilDoneOptions{
			Frequency: time.Duration(c.deletePollingIntervalSeconds * int64(time.Second)),
		})
		if err != nil {
			respErr := err.(*azcore.ResponseError)
			if resp.StatusCode == 202 && respErr.StatusCode == 404 && strings.Contains(respErr.Error(), "ResourceNotFound") {
				// Consider this specific error to be a success of deletion.
				// Work around https://github.com/pulumi/pulumi-azure-nextgen/issues/120
				return nil
			}
		}
		return err
	}

	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent, http.StatusNotFound) {
		return runtime.NewResponseError(resp)
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

	req, err := c.initRequest(ctx, method, id, queryParameters)
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
		return nil, false, handleAzCoreResponseError(err, resp)
	}

	// Some APIs are explicitly marked `x-ms-long-running-operation` and we are only supposed to poll
	// for the result in that case. However, if we get 202, we don't want to
	// consider this a failure - so try following the awaiting protocol in case the service hasn't marked
	// its API as long-running by an oversight.
	created := false
	if asyncStyle != "" || resp.StatusCode == http.StatusAccepted {
		created = resp.StatusCode < 400

		apiVersion := queryParameters["api-version"].(string)
		normalizeLocationHeader(c.host, apiVersion, resp.Header)

		pt, err := runtime.NewPoller[map[string]any](resp, c.pipeline, nil)
		if err != nil {
			return nil, created, err
		}
		outputs, err = pt.PollUntilDone(ctx, &runtime.PollUntilDoneOptions{
			Frequency: time.Duration(c.updatePollingIntervalSeconds * int64(time.Second)),
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

	// Relative polling URLs should never happen since the Azure RPC spec requires absolute polling
	// URLs. We have observed it once, though. We will try to make it absolute in that case.
	// https://github.com/Azure/azure-sdk-for-go/issues/23385
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

func (c *azCoreClient) Post(ctx context.Context, id string, bodyProps map[string]any, queryParameters map[string]any) (any, error) {
	req, err := c.initRequest(ctx, http.MethodPost, id, queryParameters)
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
	return responseBody, handleAzCoreResponseError(err, resp)
}

func (c *azCoreClient) Head(ctx context.Context, id string, apiVersion string) error {
	queryParams := map[string]any{"api-version": apiVersion}
	req, err := c.initRequest(ctx, http.MethodHead, id, queryParams)
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
	queryParams := map[string]any{"api-version": apiVersion}
	req, err := c.initRequest(ctx, readMethod, id+path, queryParams)
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
			return handleAzCoreResponseError(err, resp)
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

// handleAzCoreResponseError checks for certain kinds of errors and returns a more informative error message.
// Most of the time, it will return the original error.
func handleAzCoreResponseError(err error, resp *http.Response) error {
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "unmarshalling type ") {
		// The service returned a non-JSON response, which is unexpected. The JSON unmarshaling error is not
		// useful; return the response body and the HTTP status as an error.
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected response from Azure: '%s', HTTP status: %s", body, resp.Status)
	}
	return err
}

// from go-autorest v0.11.29, utility.go
func ensureValueString(value interface{}) string {
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// MapToValues method converts map[string]interface{} to url.Values.
// from go-autorest v0.11.29, utility.go
func MapToValues(m map[string]interface{}) url.Values {
	v := url.Values{}
	for key, value := range m {
		x := reflect.ValueOf(value)
		if x.Kind() == reflect.Array || x.Kind() == reflect.Slice {
			for i := 0; i < x.Len(); i++ {
				v.Add(key, ensureValueString(x.Index(i)))
			}
		} else {
			v.Add(key, ensureValueString(value))
		}
	}
	return v
}

// CreateTestClient creates a test AzureClient wthat doesn't actually execute requests but instead
// runs the given assertions against them.
func CreateTestClient(t *testing.T, assertions func(t *testing.T, req *http.Request)) (AzureClient, error) {
	transp := &requestAssertingTransporter{
		t:          t,
		assertions: assertions,
	}
	opts := arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Retry:     policy.RetryOptions{MaxRetries: -1}, // speeds up the tests
			Telemetry: policy.TelemetryOptions{Disabled: true},
			Transport: transp,
		},
		DisableRPRegistration: true,
	}
	return NewAzCoreClient(&fake.TokenCredential{}, "pulumi", cloud.AzurePublic, &opts)
}

type requestAssertingTransporter struct {
	t          *testing.T
	assertions func(*testing.T, *http.Request)
}

func (r *requestAssertingTransporter) Do(req *http.Request) (*http.Response, error) {
	r.assertions(r.t, req)
	return nil, nil
}
