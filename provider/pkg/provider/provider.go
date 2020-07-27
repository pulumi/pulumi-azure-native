// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"

	"github.com/pulumi/pulumi/pkg/v2/resource/provider"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v2/proto/go"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/golang/glog"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
)

const (
	ActiveDirectoryEndpoint = "https://login.microsoftonline.com/"
	AuthTokenAudience       = "https://management.azure.com/"
	// DefaultBaseURI is the default URI used for the service
	DefaultBaseURI = "https://management.azure.com"
)

type azurermProvider struct {
	host           *provider.HostClient
	name           string
	version        string
	subscriptionId string
	client         autorest.Client
	resourceMap    *AzureApiMetadata
	config         map[string]string
	schemaBytes    []byte
}

func makeProvider(host *provider.HostClient, name, version string, schemaBytes []byte, azureApiResourcesBytes []byte) (rpc.ResourceProviderServer, error) {
	// Creating a REST client
	client := autorest.NewClientWithUserAgent("pulumi")
	// Set a long timeout of 2 hours for now.
	client.PollingDuration = 120 * time.Minute

	var resourceMap *AzureApiMetadata
	err := json.Unmarshal(azureApiResourcesBytes, &resourceMap)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling resource map")
	}

	// Return the new provider
	return &azurermProvider{
		host:        host,
		name:        name,
		version:     version,
		client:      client,
		resourceMap: resourceMap,
		config:      map[string]string{},
		schemaBytes: schemaBytes,
	}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *azurermProvider) Configure(ctx context.Context, req *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
	for key, val := range req.GetVariables() {
		k.config[strings.TrimPrefix(key, "azurerm:config:")] = val
	}

	k.setLoggingContext(ctx)

	authConfig, err := k.getAuthConfig()
	if err != nil {
		return nil, errors.Wrap(err, "building auth config")
	}

	authorizer, err := k.getAuthorizationToken(authConfig)
	if err != nil {
		return nil, errors.Wrap(err, "building authorizer")
	}

	k.subscriptionId = authConfig.SubscriptionID
	k.client.Authorizer = authorizer

	return &rpc.ConfigureResponse{}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *azurermProvider) Invoke(ctx context.Context, req *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	label := fmt.Sprintf("%s.Invoke(%s)", k.name, req.Tok)
	glog.V(9).Infof("%s executing", label)

	args, err := plugin.UnmarshalProperties(req.GetArgs(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.args", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	res, ok := k.resourceMap.Invokes[req.Tok]
	if !ok {
		return nil, errors.Errorf("Resource type %s not found", req.Tok)
	}

	parameters := res.GetParameters
	if parameters == nil {
		parameters = res.PostParameters
	}

	// Construct ARM REST API path from args.
	id, body, query, err := k.prepareAzureRESTInputs(
		res.Path,
		parameters,
		args.Mappable(),
		map[string]interface{}{
			"subscriptionId": k.subscriptionId,
			"api-version":    res.ApiVersion,
		},
	)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	if res.GetParameters != nil {
		response, err = k.azureGet(ctx, id, res.ApiVersion)
	} else if res.PostParameters != nil {
		response, err = k.azurePost(ctx, id, body, query)
	} else {
		return nil, errors.Errorf("neither GET nor POST is defined for %s", label)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "request failed %v", id)
	}

	// Serialize and return RPC outputs.
	result, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(response),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.response", label), KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	return &rpc.InvokeResponse{Return: result}, nil
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
	var failures []*rpc.CheckFailure

	// Deserialize RPC inputs.
	newResInputs := req.GetNews()
	inputs, err := plugin.UnmarshalProperties(newResInputs, plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), KeepUnknowns: true, SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	resourceKey := string(urn.Type())
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("Resource type %s not found", resourceKey)
	}

	inputMap := inputs.Mappable()
	nameParam := nameParameter(res.Path)

	// Validate inputs against PUT parameters.
	for _, param := range res.PutParameters {
		if param.Name == "subscriptionId" || param.Name == "api-version" {
			// These parameters are supplied by the provider, not user's code.
			continue
		}

		if param.Location == "body" {
			// Body parameter is a collection of properties, so validate all properties in its type.
			for _, failure := range k.validateType("", param.Body, inputMap) {
				failures = append(failures, failure)
			}

			continue
		}

		name := param.Name
		if param.Name == nameParam {
			name = "name"
		}

		if value, has := inputMap[name]; has {
			// Check if the value matches the parameter constraints (recursively).
			for _, failure := range k.validateProperty(name, param.Value, value) {
				failures = append(failures, failure)
			}
		} else if param.IsRequired {
			// Report a missing required parameter.
			failures = append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("missing required property '%s'", name),
			})
		}
	}

	return &rpc.CheckResponse{Inputs: newResInputs, Failures: failures}, nil
}

// validateType checks the all properties and required properties of the given type and value map.
func (k *azurermProvider) validateType(ctx string, typ *AzureApiType, values map[string]interface{}) []*rpc.CheckFailure {
	var failures []*rpc.CheckFailure

	for _, name := range typ.RequiredProperties {
		if _, has := values[name]; !has {
			propCtx := name
			if ctx != "" {
				propCtx = fmt.Sprintf("%s.%s", ctx, name)
			}
			failures = append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("missing required property '%s.%s'", propCtx, name),
			})
		}
	}

	for name, prop := range typ.Properties {
		value, ok := values[name]
		if !ok {
			continue
		}

		propCtx := name
		if ctx != "" {
			propCtx = fmt.Sprintf("%s.%s", ctx, name)
		}
		for _, failure := range k.validateProperty(propCtx, &prop, value) {
			failures = append(failures, failure)
		}
	}
	return failures
}

// validateProperty checks the property value against its metadata.
func (k *azurermProvider) validateProperty(ctx string, prop *AzureApiProperty, value interface{}) []*rpc.CheckFailure {
	var failures []*rpc.CheckFailure

	if _, ok := value.(resource.Computed); ok {
		return failures
	}

	switch value := value.(type) {
	case resource.Computed:
		// Skip an unresolved value.
		return failures
	case map[string]interface{}:
		if prop.Ref == "" {
			return failures
		}

		// Typed object: validate all properties by looking up its type definition.
		typeName := strings.TrimPrefix(prop.Ref, "#/types/")
		typ := k.resourceMap.Types[typeName]

		for _, failure := range k.validateType(ctx, &typ, value) {
			failures = append(failures, failure)
		}
	case []interface{}:
		if prop.Type != "array" {
			return append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("'%s' should be of type '%s' but got an array", ctx, prop.Type),
			})
		}

		// Array: validate each item.
		for index, item := range value {
			itemCtx := fmt.Sprintf("%s[%d]", ctx, index)
			for _, failure := range k.validateProperty(itemCtx, prop.Items, item) {
				failures = append(failures, failure)
			}
		}
	case string:
		if prop.Type != "string" {
			return append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("'%s' should be of type '%s' but got a string", ctx, prop.Type),
			})
		}

		// Validate a string according to https://swagger.io/docs/specification/data-models/data-types/#string

		// Validate min/max length and RegEx pattern (apply to empty strings too).
		length := int64(len(value))
		if prop.MinLength != nil && length < *prop.MinLength {
			failures = append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("'%s' is too short (%d): at least %d characters required", ctx, length, *prop.MinLength),
			})
		}
		if prop.MaxLength != nil && length > *prop.MaxLength {
			failures = append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("'%s' is too long (%d): at most %d characters allowed", ctx, length, *prop.MaxLength),
			})
		}
		if prop.Pattern != "" {
			pattern := regexp.MustCompile(prop.Pattern)
			if !pattern.MatchString(value) {
				failures = append(failures, &rpc.CheckFailure{
					Reason: fmt.Sprintf("'%s' does not match expression '%s'", ctx, prop.Pattern),
				})
			}
		}

		// For closed enums, check that the value belongs to the list.
		if len(prop.Enum) > 0 {
			found := false
			possible := ""
			for _, v := range prop.Enum {
				if possible != "" {
					possible += ","
				}
				possible += fmt.Sprintf("'%s'", v)
				if v == value {
					found = true
				}
			}
			if !found {
				failures = append(failures, &rpc.CheckFailure{
					Reason: fmt.Sprintf("'%s' is '%s' but expected one of [%s]", ctx, value, possible),
				})
			}
		}
	}

	return failures
}

func (k *azurermProvider) GetSchema(ctx context.Context, req *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}

	return &rpc.GetSchemaResponse{Schema: string(k.schemaBytes)}, nil
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

	resourceKey := string(urn.Type())
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("Resource type %s not found", resourceKey)
	}

	// Construct ARM REST API body and query from intputs
	id, bodyParams, queryParams, err := k.prepareAzureRESTInputs(
		res.Path,
		res.PutParameters,
		inputs.Mappable(),
		map[string]interface{}{
			"subscriptionId": k.subscriptionId,
			"api-version":    res.ApiVersion,
		},
	)
	if err != nil {
		return nil, err
	}

	// First check if the resource already exists - we want to try our best to avoid updating instead of creating here
	// (though it's technically impossible since the only operation supported is an upsert).
	_, err = k.azureGet(ctx, id, res.ApiVersion)
	if err == nil {
		return nil, fmt.Errorf("cannot create already existing resource %v", id)
	}

	if resErr, ok := err.(*azure.RequestError); !ok || resErr.StatusCode != 404 {
		return nil, errors.Wrapf(err, "cannot check existence of resource %v: %s", id, err.Error())
	}

	// Submit the `PUT` against the ARM endpoint
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

	resourceKey := string(urn.Type())
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("Resource type '%s' not found", resourceKey)
	}

	outputs, err := k.azureGet(ctx, id, res.ApiVersion)
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

	resourceKey := string(urn.Type())
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("Resource type '%s' not found", resourceKey)
	}

	id, bodyParams, queryParams, err := k.prepareAzureRESTInputs(
		res.Path,
		res.PutParameters,
		inputs.Mappable(),
		map[string]interface{}{
			"subscriptionId": k.subscriptionId,
			"api-version":    res.ApiVersion,
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

// Delete tears down an existing resource with the given ID. If it fails, the resource is assumed
// to still exist.
func (k *azurermProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Delete(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)
	id := req.GetId()
	resourceKey := string(urn.Type())
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("Resource type '%s' not found", resourceKey)
	}
	err := k.azureDelete(ctx, id, res.ApiVersion)
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
	// Future+WaitForCompletion+GetResult steps in that case. Though in general it appears to be safe to do these
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

func (k *azurermProvider) azurePost(
	ctx context.Context,
	id string,
	bodyProps map[string]interface{},
	queryParameters map[string]interface{}) (map[string]interface{}, error) {

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
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
	var outputs map[string]interface{}
	err = autorest.Respond(
		resp,
		k.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&outputs),
		autorest.ByClosing())
	if err != nil {
		return nil, err
	}
	return outputs, nil
}

func (k *azurermProvider) prepareAzureRESTInputs(path string, parameters []AzureApiParameter, methodInputs,
	clientInputs map[string]interface{}) (string, map[string]interface{}, map[string]interface{}, error) {
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

	nameParam := nameParameter(path)

	for _, param := range parameters {
		if param.Location == "body" {
			for _, name := range param.Body.RequiredProperties {
				if _, has := methodInputs[name]; !has {
					return "", nil, nil, fmt.Errorf("missing required property '%s' in '%s'", name, param.Name)
				}
			}
			for name, _ := range param.Body.Properties {
				if value, has := methodInputs[name]; has {
					params["body"][name] = value
				}
			}
		} else {
			var val interface{}
			var has bool
			if param.Source != "" {
				val, has = locations[param.Source][param.Name]
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
				params[param.Location][param.Name] = val
			} else {
				if param.IsRequired {
					return "", nil, nil, fmt.Errorf("missing required property '%s'", param.Name)
				}
			}
		}
	}
	id := path
	for key, value := range params["path"] {
		encodedVal := autorest.Encode("path", value.(string))
		id = strings.Replace(id, "{"+key+"}", encodedVal, -1)
	}
	return id, params["body"], params["query"], nil
}

func (k *azurermProvider) setLoggingContext(ctx context.Context) {
	log.SetOutput(&LogRedirector{
		writers: map[string]func(string) error{
			tfTracePrefix: func(msg string) error { return k.host.Log(ctx, diag.Debug, "", msg) },
			tfDebugPrefix: func(msg string) error { return k.host.Log(ctx, diag.Debug, "", msg) },
			tfInfoPrefix:  func(msg string) error { return k.host.Log(ctx, diag.Info, "", msg) },
			tfWarnPrefix:  func(msg string) error { return k.host.Log(ctx, diag.Warning, "", msg) },
			tfErrorPrefix: func(msg string) error { return k.host.Log(ctx, diag.Error, "", msg) },
		},
	})
}

func (k *azurermProvider) getConfig(configName, envName string) string {
	if val, ok := k.config[configName]; ok {
		return val
	}

	return os.Getenv(envName)
}

func (k *azurermProvider) getAuthConfig() (*authentication.Config, error) {
	builder := &authentication.Builder{
		SubscriptionID: k.getConfig("subscriptionId", "ARM_SUBSCRIPTION_ID"),
		ClientID:       k.getConfig("clientId", "ARM_CLIENT_ID"),
		ClientSecret:   k.getConfig("clientSecret", "ARM_CLIENT_SECRET"),
		TenantID:       k.getConfig("tenantId", "ARM_TENANT_ID"),
		Environment:    k.getConfig("environment", "ARM_ENVIRONMENT"),
		MsiEndpoint:    k.getConfig("msiEndpoint", "ARM_MSI_ENDPOINT"),

		// Feature Toggles
		SupportsClientCertAuth:         true,
		SupportsClientSecretAuth:       true,
		SupportsManagedServiceIdentity: false,
		SupportsAzureCliToken:          true,
		SupportsAuxiliaryTenants:       false,
	}

	return builder.Build()
}

func (k *azurermProvider) getAuthorizationToken(authConfig *authentication.Config) (autorest.Authorizer, error) {
	oauthConfig, err := authConfig.BuildOAuthConfig(ActiveDirectoryEndpoint)
	if err != nil {
		return nil, err
	}

	if oauthConfig == nil {
		return nil, fmt.Errorf("unable to configure OAuthConfig for tenant %s", authConfig.TenantID)
	}

	sender := sender.BuildSender("AzureRM")
	return authConfig.GetAuthorizationToken(sender, oauthConfig, AuthTokenAudience)
}

// nameParameter parses the given URL path to find the name of the last template parameter.
func nameParameter(path string) string {
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			return part[1 : len(part)-1]
		}
	}
	return ""
}
