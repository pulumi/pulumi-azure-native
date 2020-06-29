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
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/rpcutil/rpcerror"
	"google.golang.org/grpc/codes"
	"net/http"
	"strings"

	"github.com/pulumi/pulumi/pkg/v2/resource/provider"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v2/proto/go"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
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

	// First try using a managed service principal - fall back to using the CLI
	// TODO(jen20): Extract this into a method that is less messy and supports all the options,
	// and integrate it with provider config
	cliAuthorizer, err := auth.NewAuthorizerFromCLI()
	if err == nil {
		client.Authorizer = cliAuthorizer
	} else {
		envAuthorizer, err := auth.NewAuthorizerFromEnvironment()
		if err != nil {
			return nil, err
		}
		client.Authorizer = envAuthorizer
	}

	resourceMap, err := ResourceMap()
	if err != nil {
		return nil, errors.Wrap(err, "building resource map")
	}

	// Return the new provider
	return &azurermProvider{
		host:        host,
		name:        name,
		version:     version,
		client:      client,
		resourceMap: resourceMap,
		config:      map[string]string{},
	}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *azurermProvider) Configure(_ context.Context, req *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
	for key, val := range req.GetVariables() {
		k.config[strings.TrimPrefix(key, "azurerm:config:")] = val
	}
	return &rpc.ConfigureResponse{}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *azurermProvider) Invoke(context.Context, *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	panic("Invoke not implemented")
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (k *azurermProvider) StreamInvoke(req *rpc.InvokeRequest, server rpc.ResourceProvider_StreamInvokeServer) error {
	panic("StreamInvoke not implemented")
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

func (k *azurermProvider) GetSchema(ctx context.Context, req *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	return nil, rpcerror.New(codes.Unimplemented, "GetSchema is unimplemented")
}

// CheckConfig validates the configuration for this provider.
func (k *azurermProvider) CheckConfig(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (k *azurermProvider) DiffConfig(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	return &rpc.DiffResponse{
		Changes:             0,
		Replaces:            []string{},
		Stables:             []string{},
		DeleteBeforeReplace: false,
	}, nil
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

	subscriptionId := k.config["subscriptionId"]
	if subscriptionId == "" {
		return nil, errors.New("Subscription ID is not found. Please configure azurerm:subscriptionId.")
	}

	resourceKey := string(urn.Type())
	res, ok := k.resourceMap[resourceKey]
	if !ok {
		return nil, errors.Errorf("Resource type %s not found", resourceKey)
	}

	// Construct ARM REST API body and query from intputs
	id, bodyParams, queryParams, err := k.prepareAzureRESTInputs(
		res,
		inputs.Mappable(),
		map[string]interface{}{
			"subscriptionId": subscriptionId,
			"api-version":    res.apiVersion,
		},
	)
	if err != nil {
		return nil, err
	}

	// First check if the resource already exists - we want to try our best to avoid updating instead of creating here
	// (though it's technically impossible since the only operation supported is an upsert).
	_, err = k.azureGet(ctx, id, res.apiVersion)
	if err == nil {
		return nil, fmt.Errorf("cannot create already existing resource %v", id)
	}

	if resErr, ok := err.(*azure.RequestError); !ok || resErr.StatusCode != 404 {
		return nil, errors.Wrapf(err, "cannot check existence of resource %v: %s", id, err.Error())
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
	res := k.resourceMap[string(urn.Type())]
	outputs, err := k.azureGet(ctx, id, res.apiVersion)
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

	res := k.resourceMap[string(urn.Type())]

	id, bodyParams, queryParams, err := k.prepareAzureRESTInputs(
		res,
		inputs.Mappable(),
		map[string]interface{}{
			"subscriptionId": k.config["subscriptionId"], // TODO: Get this from schema or from user?
			"api-version":    res.apiVersion,
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
	resource := k.resourceMap[string(urn.Type())]
	err := k.azureDelete(ctx, id, resource.apiVersion)
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

func (k *azurermProvider) azureDelete(ctx context.Context, id string, apiVersion string) error {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
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

func (k *azurermProvider) azureGet(ctx context.Context, id string, apiVersion string) (map[string]interface{}, error) {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
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
	// Schema-driven mapping of inputs into Autorest id/body/query
	locations := map[string]map[string]interface{}{
		"client": clientInputs,
		"method": methodInputs,
	}
	params := map[string]map[string]interface{}{
		"body":  {},
		"query": {},
		"path":  {},
	}

	spec, err := openapi.NewSpec(resource.swagggerSpecLocation)
	if err != nil {
		return "", nil, nil, errors.Wrapf(err, "failed to load spec %s", resource.swagggerSpecLocation)
	}
	path := spec.Swagger.Paths.Paths[resource.path]
	if path.Put == nil {
		return "", nil, nil, errors.New("No 'put' method on resource")
	}

	nameParam := nameParameter(resource.path)

	for _, param := range path.Put.Parameters {
		param, err := spec.ResolveParameter(param)
		if err != nil {
			return "", nil, nil, err
		}

		if param.In == "body" {
			if param.Schema == nil {
				return "", nil, nil, errors.Errorf("no schema for body parameter '%s'", param.Name)
			}

			schema, err := param.ResolveSchema(param.Schema)
			if err != nil {
				return "", nil, nil, errors.Wrapf(err, "body parameter '%s'", param.Name)
			}

			properties, required, err := k.resolveProperties(*schema)
			if err != nil {
				return "", nil, nil, errors.Wrapf(err, "body parameter '%s'", param.Name)
			}

			for _, name := range required {
				if _, has := methodInputs[name]; !has {
					return "", nil, nil, fmt.Errorf("missing required property '%s' in '%s'", name, param.Name)
				}
			}
			for _, name := range properties {
				if value, has := methodInputs[name]; has {
					params["body"][name] = value
				}
			}
		} else {
			var val interface{}
			var has bool
			loc, hasLoc := param.Extensions.GetString("x-ms-parameter-location")
			if hasLoc {
				val, has = locations[loc][param.Name]
			} else {
				// If not specified where to find it, look in both with `method` first
				val, has = methodInputs[param.Name]
				if !has {
					val, has = clientInputs[param.Name]
				}
			}
			if !has && param.Name == nameParam {
				// Use the universal 'name' parameter in place of a resource-specific name like `accountName`.
				// We should find a better way to do so when we start using the Pulumi schema in the provider.
				val, has = methodInputs["name"]
			}
			if has {
				params[param.In][param.Name] = val
			} else {
				if param.Required {
					return "", nil, nil, fmt.Errorf("missing required property '%s'", param.Name)
				}
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

// resolveProperties returns the slice of schema's property names and the slice of schema's required properties.
func (k *azurermProvider) resolveProperties(schema openapi.Schema) ([]string, []string, error) {
	var properties []string
	var required []string

	for k, _ := range schema.Properties {
		properties = append(properties, k)
	}
	for _, k := range schema.Required {
		required = append(required, k)
	}

	for _, s := range schema.AllOf {
		allOfSchema, err := schema.ResolveSchema(&s)
		if err != nil {
			return nil, nil, err
		}

		ps, rs, err := k.resolveProperties(*allOfSchema)
		if err != nil {
			return nil, nil, err
		}

		properties = append(properties, ps...)
		required = append(required, rs...)
	}

	return properties, required, nil
}

// nameParameter parses the given URL path to find the name of the last template parameter.
func nameParameter(path string) string {
	parts := strings.Split(path, "/")
	for i := len(parts)-1; i >= 0; i-- {
		part := parts[i]
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			return part[1 : len(part)-1]
		}
	}
	return ""
}
