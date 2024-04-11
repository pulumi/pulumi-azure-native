package azure

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const (
	requestFormat = `HTTP Request Begin %[1]s %[2]s ===================================================
%[3]s
===================================================== HTTP Request End %[1]s %[2]s
`
	responseFormat = `HTTP Response Begin %[1]s [%[2]s ===================================================
%[3]s
===================================================== HTTP Response End %[1]s %[2]s
`
)

type AzureDeleter interface {
	Delete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error
}

type AzureClient interface {
	AzureDeleter
	CanCreate(ctx context.Context, id, path, apiVersion, readMethod string, isSingletonResource, hasDefaultBody bool, isDefaultResponse func(map[string]any) bool) error
	Get(ctx context.Context, id string, apiVersion string) (any, error)
	Head(ctx context.Context, id string, apiVersion string) error
	Patch(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error)
	Post(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}) (any, error)
	Put(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error)
}

type azureClientImpl struct {
	environment azure.Environment
	client      autorest.Client
}

func NewAzureClient(environment azure.Environment, authorizer autorest.Authorizer, userAgent string) AzureClient {
	autorest.Count429AsRetry = false

	client := autorest.NewClientWithUserAgent(userAgent)
	client.Authorizer = authorizer
	// Log requests
	client.RequestInspector = withInspection()
	// Log responses
	client.ResponseInspector = byInspecting()

	return &azureClientImpl{environment: environment, client: client}
}

func (a *azureClientImpl) Delete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	for k, v := range queryParams {
		queryParameters[k] = v
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(a.environment.ResourceManagerEndpoint),
		autorest.WithPath(id),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return err
	}
	resp, err := autorest.SendWithSender(
		a.client,
		prepReq,
		azure.DoRetryWithRegistration(a.client),
	)
	if err != nil {
		return err
	}

	// Some APIs are explicitly marked `x-ms-long-running-operation` and we should only do the
	// Future+WaitForCompletion+GetResult steps in that case.
	if asyncStyle != "" {
		future, err := azure.NewFutureFromResponse(resp)
		if err != nil {
			return err
		}
		err = future.WaitForCompletionRef(ctx, a.client)
		if err != nil {
			if detailed, ok := err.(autorest.DetailedError); ok {
				if resp.StatusCode == 202 && detailed.StatusCode == 404 && detailed.Original != nil &&
					strings.Contains(detailed.Original.Error(), "ResourceNotFound") {
					// Consider this specific error to be a success of deletion.
					// Work around https://github.com/pulumi/pulumi-azure-nextgen/issues/120
					// Upstream fix is tracked in https://github.com/Azure/go-autorest/issues/596
					return nil
				}
			}
			return err
		}
		resp, err = future.GetResult(a.client)
		if err != nil {
			return err
		}
	}

	err = autorest.Respond(
		resp,
		a.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *azureClientImpl) Post(
	ctx context.Context,
	id string,
	bodyProps map[string]interface{},
	queryParameters map[string]interface{}) (any, error) {

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(a.environment.ResourceManagerEndpoint),
		autorest.WithPath(id),
		autorest.WithJSON(bodyProps),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return nil, err
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(
		a.client,
		prepReq,
		azure.DoRetryWithRegistration(a.client),
	)
	if err != nil {
		return nil, err
	}
	var outputs any
	err = autorest.Respond(
		resp,
		a.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&outputs),
		autorest.ByClosing())
	if err != nil {
		return nil, handleResponseError(err, resp)
	}
	return outputs, nil
}

func (a *azureClientImpl) Get(ctx context.Context, id string, apiVersion string) (any, error) {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(a.environment.ResourceManagerEndpoint),
		autorest.WithPath(id),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return nil, err
	}
	resp, err := autorest.SendWithSender(
		a.client,
		prepReq,
		azure.DoRetryWithRegistration(a.client),
	)
	if err != nil {
		return nil, err
	}
	var outputs any
	err = autorest.Respond(
		resp,
		a.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		forceRequestErrorForStatusNotFound,
		autorest.ByUnmarshallingJSON(&outputs),
		autorest.ByClosing())
	if err != nil {
		return nil, err
	}
	return outputs, nil
}

func (a *azureClientImpl) Head(ctx context.Context, id string, apiVersion string) error {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(a.environment.ResourceManagerEndpoint),
		autorest.WithPath(id),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return err
	}
	resp, err := autorest.SendWithSender(
		a.client,
		prepReq,
		azure.DoRetryWithRegistration(a.client),
	)
	if err != nil {
		return err
	}
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		forceRequestErrorForStatusNotFound,
		autorest.ByClosing())
	return err
}

// If a Status Code is 404, always return a request error 'StatusNotFound'. This doesn't work out-of-the-box
// in case the service returns an invalid response that failed to parse into an 'autorest.ServiceError'.
// The parsing error (e.g., 'json.UnmarshalTypeError') would mask the 404 response and the provider wouldn't
// be able to make the right decision for a missing resource.
func forceRequestErrorForStatusNotFound(r autorest.Responder) autorest.Responder {
	return autorest.ResponderFunc(func(resp *http.Response) error {
		err := r.Respond(resp)
		if err == nil || !autorest.ResponseHasStatusCode(resp, http.StatusNotFound) {
			return err
		}
		if _, ok := err.(*azure.RequestError); ok {
			return err
		}
		return &azure.RequestError{
			DetailedError: autorest.DetailedError{
				Original:   err,
				StatusCode: http.StatusNotFound,
			},
		}
	})
}

func (a *azureClientImpl) Patch(
	ctx context.Context,
	id string,
	bodyProps map[string]interface{},
	queryParameters map[string]interface{},
	asyncStyle string,
) (map[string]interface{}, bool, error) {
	return a.update(ctx, id, bodyProps, queryParameters, autorest.AsPatch(), asyncStyle)
}

func (a *azureClientImpl) Put(
	ctx context.Context,
	id string,
	bodyProps map[string]interface{},
	queryParameters map[string]interface{},
	asyncStyle string,
) (map[string]interface{}, bool, error) {
	return a.update(ctx, id, bodyProps, queryParameters, autorest.AsPut(), asyncStyle)
}

func (a *azureClientImpl) update(
	ctx context.Context,
	id string,
	bodyProps map[string]interface{},
	queryParameters map[string]interface{},
	op autorest.PrepareDecorator,
	asyncStyle string,
) (map[string]interface{}, bool, error) {
	decorators := []autorest.PrepareDecorator{
		autorest.AsContentType("application/json; charset=utf-8"),
		op,
		autorest.WithBaseURL(a.environment.ResourceManagerEndpoint),
		autorest.WithPath(id),
		autorest.WithQueryParameters(queryParameters),
	}
	if bodyProps != nil {
		decorators = append(decorators, autorest.WithJSON(bodyProps))
	}
	prepReq, err := autorest.Prepare((&http.Request{}).WithContext(ctx), decorators...)
	if err != nil {
		return nil, false, err
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(
		a.client,
		prepReq,
		azure.DoRetryWithRegistration(a.client),
	)
	if err != nil {
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

		// Ignore the style value for now, let go-autorest handle the headers.
		future, err := azure.NewFutureFromResponse(resp)
		if err != nil {
			return nil, created, err
		}
		err = future.WaitForCompletionRef(ctx, a.client)
		if err != nil {
			return nil, created, err
		}
		resp, err = future.GetResult(a.client)
		if err != nil {
			return nil, created, err
		}
	}

	var outputs map[string]interface{}
	err = autorest.Respond(
		resp,
		a.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&outputs),
		autorest.ByClosing())
	if err != nil {
		return nil, created, handleResponseError(err, resp)
	}
	return outputs, true, nil
}

// CanCreate asserts that a resource with a given ID and API version can be created
// or returns an error otherwise.
func (a *azureClientImpl) CanCreate(ctx context.Context, id, path, apiVersion, readMethod string, isSingletonResource, hasDefaultBody bool, isDefaultResponse func(map[string]any) bool) error {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	op := autorest.AsGet()
	switch readMethod {
	case "HEAD":
		op = autorest.AsHead()
	case "POST":
		op = autorest.AsPost()
	}
	preparer := autorest.CreatePreparer(
		op,
		autorest.WithBaseURL(a.environment.ResourceManagerEndpoint),
		autorest.WithPath(id+path),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return err
	}
	resp, err := autorest.SendWithSender(
		a.client,
		prepReq,
		azure.DoRetryWithRegistration(a.client),
	)
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
		err = autorest.Respond(
			resp,
			a.client.ByInspecting(),
			autorest.ByUnmarshallingJSON(&outputs),
			autorest.ByClosing())
		if err != nil {
			return handleResponseError(err, resp)
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
		err = autorest.Respond(
			resp,
			a.client.ByInspecting(),
			autorest.ByUnmarshallingJSON(&outputs),
			autorest.ByClosing())
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

// withInspection is a copy of autorest's LoggingInspector.WithInspector. It uses our glog wrapper
// instead of a go logger which gets complicated in the presence of log redirection via the host.
func withInspection() autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			var body, b bytes.Buffer

			if r.Body != nil {
				defer r.Body.Close()

				r.Body = io.NopCloser(io.TeeReader(r.Body, &body))
				if err := r.Write(&b); err != nil {
					return nil, fmt.Errorf("failed to write response: %v", err)
				}

				logging.V(9).Infof(requestFormat, r.Method, r.URL, b.String())

				r.Body = io.NopCloser(&body)
			} else {
				logging.V(9).Infof(requestFormat, r.Method, r.URL, "Empty body")
			}
			return p.Prepare(r)
		})
	}
}

// byInspecting is acopy of autorest's LoggingInspector.ByInspecting(). It uses our glog wrapper
// instead of a go logger which gets complicated in the presence of log redirection via the host.
func byInspecting() autorest.RespondDecorator {
	return func(r autorest.Responder) autorest.Responder {
		return autorest.ResponderFunc(func(resp *http.Response) error {
			var body, b bytes.Buffer
			defer resp.Body.Close()
			resp.Body = io.NopCloser(io.TeeReader(resp.Body, &body))
			if err := resp.Write(&b); err != nil {
				return fmt.Errorf("failed to write response: %v", err)
			}

			logging.V(9).Infof(responseFormat, resp.Request.Method, resp.Request.URL, b.String())

			resp.Body = io.NopCloser(&body)
			return r.Respond(resp)
		})
	}
}

// handleResponseError checks for certain kinds of errors and returns a more informative error message.
// Most of the time, it will return the original error.
func handleResponseError(err error, resp *http.Response) error {
	if strings.Contains(err.Error(), "autorest/azure: error response cannot be parsed") {
		// The service returned a non-JSON response, which is unexpected. The JSON unmarshaling error is not
		// useful; return the response body and the HTTP status as an error.
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected response from Azure: '%s', HTTP status: %s", body, resp.Status)
	}
	return err
}
