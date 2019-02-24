// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/resource/plugin"
	"github.com/pulumi/pulumi/pkg/resource/provider"
	rpc "github.com/pulumi/pulumi/sdk/proto/go"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/go-openapi/spec"
	"github.com/golang/glog"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
)

const (
	// DefaultBaseURI is the default URI used for the service
	DefaultBaseURI = "https://management.azure.com"
)

type azurermProvider struct {
	host        *provider.HostClient
	name        string
	version     string
	client      autorest.Client
	resourceMap map[string]Resource
	config      map[string]string
}

func makeProvider(host *provider.HostClient, name, version string) (rpc.ResourceProviderServer, error) {
	// Creating a REST client
	client := autorest.NewClientWithUserAgent("pulumi")
	authorizer, err := auth.NewAuthorizerFromCLI()
	if err != nil {
		return nil, err
	}
	client.Authorizer = authorizer

	// Return the new provider
	return &azurermProvider{
		host:        host,
		name:        name,
		version:     version,
		client:      client,
		resourceMap: ResourceMap(),
		config:      map[string]string{},
	}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *azurermProvider) Configure(_ context.Context, req *rpc.ConfigureRequest) (*pbempty.Empty, error) {
	for key, val := range req.GetVariables() {
		k.config[strings.TrimPrefix(key, "azurerm:config:")] = val
	}
	return &pbempty.Empty{}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *azurermProvider) Invoke(context.Context, *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	panic("Invoke not implemented")
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (k *azurermProvider) Check(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Check(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)

	newResInputs := req.GetNews()
	return &rpc.CheckResponse{Inputs: newResInputs, Failures: nil}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (k *azurermProvider) Diff(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	// TODO: Actually figure out whether any properties require replacement
	// * Anything that goes in `path` requies replacement
	return &rpc.DiffResponse{
		Changes:             0,
		Replaces:            []string{},
		Stables:             []string{},
		DeleteBeforeReplace: false,
	}, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (k *azurermProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Create(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), KeepUnknowns: true, SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	// Construct ARM REST API body and query from intputs
	id, bodyParams, queryParams, err := k.prepareAzureRESTInputs(
		k.resourceMap[string(urn.Type())],
		inputs.Mappable(),
		map[string]interface{}{
			"subscriptionId": k.config["subscriptionId"],
			"api-version":    "2018-10-01", // TODO: Get this from schema or from user?
		},
	)
	if err != nil {
		return nil, err
	}

	// First check if the resource already exists - we want to try our best to avoid updating instead of creating here
	// (though it's technically impossible since the only operation supported is an upsert).
	_, err = k.azureGet(ctx, id)
	if resErr, ok := err.(azure.RequestError); ok && resErr.StatusCode == 404 {
		return nil, fmt.Errorf("cannot create already existing resource %v", id)
	}

	//  Submit the `PUT` against the ARM endpoint
	outputs, err := k.azureCreateOrUpdate(ctx, id, bodyParams, queryParams)
	if err != nil {
		return nil, err
	}

	// Serialize and return RPC outputs
	outputProperties, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.autonamedInputs", label), KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	return &rpc.CreateResponse{
		Id:         id,
		Properties: outputProperties,
	}, nil
}

// Read the current live state associated with a resource.
func (k *azurermProvider) Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Read(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)
	id := req.GetId()
	outputs, err := k.azureGet(ctx, id)
	if err != nil {
		return nil, err
	}
	outputProperties, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.autonamedInputs", label), KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	return &rpc.ReadResponse{Id: id, Properties: outputProperties}, nil
}

// Update updates an existing resource with new values.
func (k *azurermProvider) Update(ctx context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Update(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)
	inputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), KeepUnknowns: true, SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	id, bodyParams, queryParams, err := k.prepareAzureRESTInputs(
		k.resourceMap[string(urn.Type())],
		inputs.Mappable(),
		map[string]interface{}{
			"subscriptionId": k.config["subscriptionId"], // TODO: Get this from schema or from user?
			"api-version":    "2018-10-01",               // TODO: Get this from schema or from user?
		},
	)
	if err != nil {
		return nil, err
	}

	outputs, err := k.azureCreateOrUpdate(ctx, id, bodyParams, queryParams)
	if err != nil {
		return nil, err
	}

	outputProperties, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.autonamedInputs", label), KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	return &rpc.UpdateResponse{
		Properties: outputProperties,
	}, nil
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (k *azurermProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Delete(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)
	id := req.GetId()
	err := k.azureDelete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pbempty.Empty{}, nil
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (k *azurermProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: k.version,
	}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (k *azurermProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	// TODO
	return &pbempty.Empty{}, nil
}

func (k *azurermProvider) azureCreateOrUpdate(
	ctx context.Context,
	id string,
	bodyProps map[string]interface{},
	queryParameters map[string]interface{}) (map[string]interface{}, error) {

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(DefaultBaseURI),
		autorest.WithPath(id),
		autorest.WithJSON(bodyProps),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return nil, err
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(
		k.client,
		prepReq,
		azure.DoRetryWithRegistration(k.client),
	)
	if err != nil {
		return nil, err
	}
	// TODO: Some APIs are explicitly marked `x-ms-long-running-operation` and possibly we should only do the
	// Future+WaitForCompletion+GetResult steps in that case.  Though in general it appears to be safe to do these
	// always.
	future, err := azure.NewFutureFromResponse(resp)
	if err != nil {
		return nil, err
	}
	err = future.WaitForCompletionRef(ctx, k.client)
	if err != nil {
		return nil, err
	}
	res, err := future.GetResult(k.client)
	if err != nil {
		return nil, err
	}
	var outputs map[string]interface{}
	err = autorest.Respond(
		res,
		k.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&outputs),
		autorest.ByClosing())
	if err != nil {
		return nil, err
	}
	return outputs, nil
}

func (k *azurermProvider) azureDelete(ctx context.Context, id string) error {
	const APIVersion = "2018-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(DefaultBaseURI),
		autorest.WithPath(id),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return err
	}
	resp, err := autorest.SendWithSender(
		k.client,
		prepReq,
		azure.DoRetryWithRegistration(k.client),
	)
	if err != nil {
		return err
	}
	future, err := azure.NewFutureFromResponse(resp)
	if err != nil {
		return err
	}
	err = future.WaitForCompletionRef(ctx, k.client)
	if err != nil {
		return err
	}
	res, err := future.GetResult(k.client)
	if err != nil {
		return err
	}
	var result map[string]interface{}
	err = autorest.Respond(
		res,
		k.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (k *azurermProvider) azureGet(ctx context.Context, id string) (map[string]interface{}, error) {
	const APIVersion = "2018-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(DefaultBaseURI),
		autorest.WithPath(id),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return nil, err
	}
	resp, err := autorest.SendWithSender(
		k.client,
		prepReq,
		azure.DoRetryWithRegistration(k.client),
	)
	if err != nil {
		return nil, err
	}
	var outputs map[string]interface{}
	err = autorest.Respond(
		resp,
		k.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&outputs),
		autorest.ByClosing())
	if err != nil {
		return nil, err
	}
	return outputs, nil
}

func (k *azurermProvider) prepareAzureRESTInputs(resource Resource, methodInputs, clientInputs map[string]interface{}) (string, map[string]interface{}, map[string]interface{}, error) {
	// Schema-drivenn mapping of inputs into Autorest id/body/query
	locations := map[string]map[string]interface{}{
		"client": clientInputs,
		"method": methodInputs,
	}
	params := map[string]map[string]interface{}{
		"body":  map[string]interface{}{},
		"query": map[string]interface{}{},
		"path":  map[string]interface{}{},
	}
	resSpec, err := resource.LoadSpec()
	if err != nil {
		return "", nil, nil, errors.Wrapf(err, "failed to load spec")
	}
	path := resSpec.Paths.Paths[resource.path]
	if path.Put == nil {
		return "", nil, nil, fmt.Errorf("No 'put' method on resource")
	}
	for _, param := range path.Put.Parameters {
		// TODO: Can we resolve "$refs" globally instead of doing this here?
		if ptr := param.Ref.Ref.GetPointer(); ptr != nil && !ptr.IsEmpty() {
			base, err := url.Parse(resource.swagggerSpecLocation)
			if err != nil {
				return "", nil, nil, err
			}
			relative, err := url.Parse(param.Ref.RemoteURI())
			if err != nil {
				return "", nil, nil, err
			}
			finalURL := base.ResolveReference(relative)
			newSpec, err := loadSwaggerSpec(finalURL.String())
			if err != nil {
				return "", nil, nil, err
			}

			p2, _, err := ptr.Get(newSpec)
			if err != nil {
				return "", nil, nil, err
			}
			param = p2.(spec.Parameter)
		}
		var val interface{}
		var has bool
		loc, hasLoc := param.Extensions.GetString("x-ms-parameter-location")
		if hasLoc {
			val, has = locations[loc][param.Name]
		} else {
			// If not specified where to find it, look in both with `method` first
			val, has = locations["method"][param.Name]
			if !has {
				val, has = locations["client"][param.Name]
			}
		}
		if has {
			if param.In == "body" {
				// TODO: It seems Autorest inlines "body" parameters, seemingly assuming there will only be one?  Or
				// that if there are multple they are all merged?
				for k, v := range val.(map[string]interface{}) {
					params["body"][k] = v
				}
			} else {
				params[param.In][param.Name] = val
			}
		} else {
			if param.Required {
				return "", nil, nil, fmt.Errorf("missing required property '%s'", param.Name)
			}
		}
	}
	id := resource.path
	for key, value := range params["path"] {
		encodedVal := autorest.Encode("path", value.(string))
		id = strings.Replace(id, "{"+key+"}", encodedVal, -1)
	}
	return id, params["body"], params["query"], nil

}
