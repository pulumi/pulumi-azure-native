// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/golang/glog"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/arm2pulumi"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/version"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v2/resource/provider"
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	rpc "github.com/pulumi/pulumi/sdk/v2/proto/go"
)

const (
	// Microsoft's Pulumi Partner ID.
	PulumiPartnerID = "a90539d8-a7a6-5826-95c4-1fbef22d4b22"
)

type azureNextGenProvider struct {
	host            *provider.HostClient
	name            string
	version         string
	subscriptionID  string
	environment     azure.Environment
	client          autorest.Client
	resourceMap     *resources.AzureAPIMetadata
	config          map[string]string
	schemaBytes     []byte
	converter       *resources.SdkShapeConverter
	customResources map[string]*resources.CustomResource
}

func makeProvider(host *provider.HostClient, name, version string, schemaBytes []byte,
	azureAPIResourcesBytes []byte) (rpc.ResourceProviderServer, error) {
	// Creating a REST client, defaulting to Pulumi Partner ID until the Configure method is invoked.
	client := autorest.NewClientWithUserAgent(buildUserAgent(PulumiPartnerID))
	// Set a long timeout of 2 hours for now.
	client.PollingDuration = 120 * time.Minute

	resourceMap, err := LoadMetadata(azureAPIResourcesBytes)
	if err != nil {
		return nil, err
	}

	// Return the new provider
	return &azureNextGenProvider{
		host:        host,
		name:        name,
		version:     version,
		client:      client,
		resourceMap: resourceMap,
		config:      map[string]string{},
		schemaBytes: schemaBytes,
		converter:   &resources.SdkShapeConverter{Types: resourceMap.Types},
	}, nil
}

// LoadMetadata deserializes the provided compressed json byte array into
// an AzureAPIMetadata in memory
func LoadMetadata(azureAPIResourcesBytes []byte) (*resources.AzureAPIMetadata, error) {
	var resourceMap resources.AzureAPIMetadata

	uncompressed, err := gzip.NewReader(bytes.NewReader(azureAPIResourcesBytes))
	if err != nil {
		return nil, errors.Wrap(err, "expand compressed metadata")
	}
	if err = json.NewDecoder(uncompressed).Decode(&resourceMap); err != nil {
		return nil, errors.Wrap(err, "unmarshalling resource map")
	}
	if err = uncompressed.Close(); err != nil {
		return nil, errors.Wrap(err, "closing uncompress stream for metadata")
	}
	return &resourceMap, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *azureNextGenProvider) Configure(ctx context.Context,
	req *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
	for key, val := range req.GetVariables() {
		k.config[strings.TrimPrefix(key, "azure-nextgen:config:")] = val
	}

	k.setLoggingContext(ctx)

	authConfig, err := k.getAuthConfig()
	if err != nil {
		return nil, errors.Wrap(err, "building auth config")
	}

	envName := authConfig.Environment
	env, err := azure.EnvironmentFromName(envName)
	if err != nil {
		env, err = azure.EnvironmentFromName(fmt.Sprintf("AZURE%sCLOUD", envName))
		if err != nil {
			return nil, errors.Wrapf(err, "environment %q was not found", envName)
		}
	}
	k.environment = env

	tokenAuth, bearerAuth, err := k.getAuthorizers(authConfig)
	if err != nil {
		return nil, errors.Wrap(err, "building authorizer")
	}

	k.subscriptionID = authConfig.SubscriptionID
	k.client.Authorizer = tokenAuth
	k.client.UserAgent = k.getUserAgent()

	k.customResources = resources.BuildCustomResources(&env, k.subscriptionID, bearerAuth, tokenAuth, k.client.UserAgent)

	return &rpc.ConfigureResponse{
		SupportsPreview: true,
	}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *azureNextGenProvider) Invoke(ctx context.Context, req *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	label := fmt.Sprintf("%s.Invoke(%s)", k.name, req.Tok)
	glog.V(9).Infof("%s executing", label)

	args, err := plugin.UnmarshalProperties(req.GetArgs(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.args", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	var outputs map[string]interface{}
	switch req.Tok {
	case "azure-nextgen:authorization/latest:getClientConfig":
		auth, err := k.getAuthConfig()
		if err != nil {
			return nil, fmt.Errorf("getting auth config: %w", err)
		}
		objectId, err := auth.GetAuthenticatedObjectID(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting authenticated object ID: %w", err)
		}
		outputs = map[string]interface{}{
			"clientId":       auth.ClientID,
			"objectId":       objectId,
			"subscriptionId": auth.SubscriptionID,
			"tenantId":       auth.TenantID,
		}
	case "azure-nextgen:authorization/latest:getClientToken":
		auth, err := k.getAuthConfig()
		if err != nil {
			return nil, fmt.Errorf("getting auth config: %w", err)
		}
		endpoint := k.environment.ResourceManagerEndpoint
		if endpointArg := args["endpoint"]; endpointArg.HasValue() && endpointArg.IsString() {
			endpoint = endpointArg.StringValue()
		}
		token, err := k.getOAuthToken(ctx, auth, endpoint)
		if err != nil {
			return nil, err
		}
		outputs = map[string]interface{}{"token": token}
	case "azure-nextgen:armtemplate:decode":
		var text string
		if textArg := args["text"]; textArg.HasValue() && textArg.IsString() {
			text = textArg.StringValue()
		} else {
			return nil, errors.New("missing required field 'text' of type 'string'")
		}

		uncompressed, err := gzip.NewReader(bytes.NewReader(k.schemaBytes))
		if err != nil {
			return nil, errors.Wrap(err, "expand compressed schema")
		}

		var pkgSpec schema.PackageSpec
		if err = json.NewDecoder(uncompressed).Decode(&pkgSpec); err != nil {
			return nil, fmt.Errorf("deserializing schema: %w", err)
		}

		res, err := arm2pulumi.DecodeTemplate(&pkgSpec, k.resourceMap, text)
		if err != nil {
			return nil, fmt.Errorf("rendering template: %w", err)
		}
		outputs = map[string]interface{}{"result": res}
	default:
		res, ok := k.resourceMap.Invokes[req.Tok]
		if !ok {
			return nil, errors.Errorf("Invoke type %s not found", req.Tok)
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
				"subscriptionId": k.subscriptionID,
				"api-version":    res.APIVersion,
			},
		)
		if err != nil {
			return nil, err
		}

		var response map[string]interface{}
		if res.GetParameters != nil {
			response, err = k.azureGet(ctx, id, res.APIVersion)
		} else if res.PostParameters != nil {
			response, err = k.azurePost(ctx, id, body, query)
		} else {
			return nil, errors.Errorf("neither GET nor POST is defined for %s", label)
		}

		if err != nil {
			return nil, errors.Wrapf(err, "request failed %v", id)
		}

		// Map the raw response to the shape of outputs that the SDKs expect.
		outputs = k.converter.BodyPropertiesToSDK(res.Response, response)
	}

	// Serialize and return RPC outputs.
	result, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.response", label), KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	return &rpc.InvokeResponse{Return: result}, nil
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (k *azureNextGenProvider) StreamInvoke(_ *rpc.InvokeRequest, _ rpc.ResourceProvider_StreamInvokeServer) error {
	panic("StreamInvoke not implemented")
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (k *azureNextGenProvider) Check(_ context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
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

	// Validate inputs against PUT parameters.
	for _, param := range res.PutParameters {
		if param.Name == "subscriptionId" || param.Name == "api-version" {
			// These parameters are supplied by the provider, not user's code.
			continue
		}

		if param.Location == "body" {
			// Body parameter is a collection of properties, so validate all properties in its type.
			failures = append(failures, k.validateType("", param.Body, inputMap)...)

			continue
		}

		name := param.Name
		if param.Value.SdkName != "" {
			name = param.Value.SdkName
		}

		if value, has := inputMap[name]; has {
			// Check if the value matches the parameter constraints (recursively).
			failures = append(failures, k.validateProperty(name, param.Value, value)...)
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
func (k *azureNextGenProvider) validateType(ctx string, typ *resources.AzureAPIType,
	values map[string]interface{}) []*rpc.CheckFailure {
	var failures []*rpc.CheckFailure

	for _, name := range typ.RequiredProperties {
		prop := typ.Properties[name]
		if prop.SdkName != "" {
			name = prop.SdkName
		}
		if _, has := values[name]; !has {
			propCtx := name
			if ctx != "" {
				propCtx = fmt.Sprintf("%s.%s", ctx, name)
			}
			failures = append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("missing required property '%s'", propCtx),
			})
		}
	}

	for name, prop := range typ.Properties {
		p := prop // https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		if prop.SdkName != "" {
			name = prop.SdkName
		}
		value, ok := values[name]
		if !ok {
			continue
		}

		propCtx := name
		if ctx != "" {
			propCtx = fmt.Sprintf("%s.%s", ctx, name)
		}
		failures = append(failures, k.validateProperty(propCtx, &p, value)...)
	}
	return failures
}

// validateProperty checks the property value against its metadata.
func (k *azureNextGenProvider) validateProperty(ctx string, prop *resources.AzureAPIProperty,
	value interface{}) []*rpc.CheckFailure {
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

		failures = append(failures, k.validateType(ctx, &typ, value)...)
	case []interface{}:
		if prop.Type != "array" {
			return append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("'%s' should be of type '%s' but got an array", ctx, prop.Type),
			})
		}

		// Array: validate each item.
		for index, item := range value {
			itemCtx := fmt.Sprintf("%s[%d]", ctx, index)
			failures = append(failures, k.validateProperty(itemCtx, prop.Items, item)...)
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
				Reason: fmt.Sprintf("'%s' is too short (%d): at least %d characters required",
					ctx, length, *prop.MinLength),
			})
		}
		if prop.MaxLength != nil && length > *prop.MaxLength {
			failures = append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("'%s' is too long (%d): at most %d characters allowed",
					ctx, length, *prop.MaxLength),
			})
		}
		if prop.Pattern != "" {
			pattern, err := regexp.Compile(prop.Pattern)
			// TODO: Support ECMA-262 regexp https://github.com/pulumi/pulumi-azure-nextgen-provider/issues/164
			if err == nil && !pattern.MatchString(value) {
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
	case float64:
		if prop.Type != "number" && prop.Type != "integer" {
			return append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("'%s' should be of type '%s' but got a number", ctx, prop.Type),
			})
		}
		if prop.Type == "integer" {
			if _, frac := math.Modf(math.Abs(value)); frac > 1e-9 && frac < 1.0-1e-9 {
				failures = append(failures, &rpc.CheckFailure{
					Reason: fmt.Sprintf("'%s' must be an integer but got %f", ctx, value),
				})
			}
		}

		// Validate min/max inclusive ranges per https://swagger.io/docs/specification/data-models/data-types/#numbers
		if prop.Minimum != nil && value < *prop.Minimum {
			failures = append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("'%s' is too low (%f < %f)", ctx, value, *prop.Minimum),
			})
		}
		if prop.Maximum != nil && value > *prop.Maximum {
			failures = append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("'%s' is too high (%f > %f)", ctx, value, *prop.Maximum),
			})
		}
	}

	return failures
}

func (k *azureNextGenProvider) GetSchema(_ context.Context, req *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}

	uncompressed, err := gzip.NewReader(bytes.NewReader(k.schemaBytes))
	if err != nil {
		return nil, errors.Wrap(err, "expand compressed bytes for schema")
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, uncompressed)
	if err != nil {
		return nil, errors.Wrap(err, "closing read stream for schema")
	}

	if err = uncompressed.Close(); err != nil {
		return nil, errors.Wrap(err, "closing uncompress stream for schema")
	}

	return &rpc.GetSchemaResponse{Schema: buf.String()}, nil
}

// CheckConfig validates the configuration for this provider.
func (k *azureNextGenProvider) CheckConfig(_ context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (k *azureNextGenProvider) DiffConfig(context.Context, *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	return &rpc.DiffResponse{
		Changes:             0,
		Replaces:            []string{},
		Stables:             []string{},
		DeleteBeforeReplace: false,
	}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (k *azureNextGenProvider) Diff(_ context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Diff(%s)", k.name, urn)

	// Retrieve the old state.
	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	// Extract old inputs from the `__inputs` field of the old state.
	oldInputs := parseCheckpointObject(oldState)
	if oldInputs == nil {
		// Protect against a crash for the transition from pre-__inputs state files.
		// This shouldn't happen in any real user's stack.
		oldInputs = resource.PropertyMap{}
		glog.V(9).Infof("no __inputs found for '%s'", urn)
	}

	// Get new resource inputs. The user is submitting these as an update.
	newResInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.news", label),
		KeepUnknowns: true,
		SkipNulls:    true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs %v %v",
			oldState, newResInputs)
	}

	// Calculate the difference between old and new inputs.
	diff := oldInputs.Diff(newResInputs)

	if diff == nil {
		return &rpc.DiffResponse{
			Changes:             rpc.DiffResponse_DIFF_NONE,
			Replaces:            []string{},
			Stables:             []string{},
			DeleteBeforeReplace: false,
		}, nil
	}

	resourceKey := string(urn.Type())
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("Resource type %s not found", resourceKey)
	}

	// Calculate the detailed diff object containing information about replacements.
	detailedDiff := calculateDetailedDiff(&res, diff)

	// Based on the detailed diff above, calculate the list of changes and replacements.
	var changes, replaces []string
	for k, v := range detailedDiff {
		parts := strings.Split(k, ".")
		changes = append(changes, parts[0])
		v.InputDiff = true

		switch v.Kind {
		case rpc.PropertyDiff_ADD_REPLACE:
			// Special case: previously, the property input had no value but is now set to X.
			// If the output contains this property and it's X, that's not a replacement.
			// Workaround for https://github.com/pulumi/pulumi-azure-nextgen/issues/238
			// TODO: remove this block before GA.
			key := resource.PropertyKey(k)
			_, hasOldInput := oldInputs[key]
			newInputValue, hasNewInput := newResInputs[key]
			outputValue, hasOutput := oldState[key]
			if !hasOldInput && hasNewInput && hasOutput && newInputValue == outputValue {
				v.Kind = rpc.PropertyDiff_ADD
			} else {
				replaces = append(replaces, k)
			}
		case rpc.PropertyDiff_DELETE_REPLACE, rpc.PropertyDiff_UPDATE_REPLACE:
			replaces = append(replaces, k)
		}
	}

	// TODO: Define delete-before-replace only if autonaming is off.
	deleteBeforeReplace := len(replaces) > 0
	changeType := rpc.DiffResponse_DIFF_NONE
	if len(detailedDiff) > 0 {
		changeType = rpc.DiffResponse_DIFF_SOME
	}

	response := rpc.DiffResponse{
		Changes:             changeType,
		Replaces:            replaces,
		Stables:             []string{},
		DeleteBeforeReplace: deleteBeforeReplace,
		Diffs:               changes,
		DetailedDiff:        detailedDiff,
		HasDetailedDiff:     true,
	}

	return &response, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (k *azureNextGenProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
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

	if req.GetPreview() {
		// Calculate the preview of outputs using the current inputs and marking all other response properties
		// as computed.
		previewState, err := plugin.MarshalProperties(
			k.converter.PreviewOutputs(inputs, res.Response),
			plugin.MarshalOptions{Label: fmt.Sprintf("%s.previewState", label), KeepUnknowns: true},
		)
		if err != nil {
			return nil, err
		}

		return &rpc.CreateResponse{
			Properties: previewState,
		}, nil
	}

	// Construct ARM REST API body and query from intputs
	id, bodyParams, queryParams, err := k.prepareAzureRESTInputs(
		res.Path,
		res.PutParameters,
		inputs.Mappable(),
		map[string]interface{}{
			"subscriptionId": k.subscriptionID,
			"api-version":    res.APIVersion,
		},
	)
	if err != nil {
		return nil, err
	}

	var outputs map[string]interface{}
	customRes, isCustom := k.customResources[res.Path]
	switch {
	case isCustom && customRes.Create != nil:
		// First check if the resource already exists - we want to try our best to avoid updating instead of creating here.
		_, exists, err := customRes.Read(ctx, inputs)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, fmt.Errorf("cannot create already existing resource %q", id)
		}

		// Create the custom resource and retrieve its outputs, which already match the SDK shape.
		outputs, err = customRes.Create(ctx, inputs)
		if err != nil {
			return nil, err
		}
	default:
		// First check if the resource already exists - we want to try our best to avoid updating instead of creating here
		// (though it's technically impossible since the only operation supported is an upsert).
		err = k.azureCanCreate(ctx, id, &res)
		if err != nil {
			return nil, err
		}

		// Submit the `PUT` against the ARM endpoint
		response, err := k.azureCreateOrUpdate(ctx, id, bodyParams, queryParams, res.PutAsyncStyle)
		if err != nil {
			return nil, err
		}

		// Read the canonical ID from the response.
		if azureId, ok := response["id"].(string); ok {
			id = azureId
		}

		// Map the raw response to the shape of outputs that the SDKs expect.
		outputs = k.converter.BodyPropertiesToSDK(res.Response, response)
	}

	// Store both outputs and inputs into the state.
	obj := checkpointObject(inputs, outputs)

	// Serialize and return RPC outputs
	checkpoint, err := plugin.MarshalProperties(
		obj,
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	return &rpc.CreateResponse{
		Id:         id,
		Properties: checkpoint,
	}, nil
}

// Read the current live state associated with a resource.
func (k *azureNextGenProvider) Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Read(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)
	id := req.GetId()

	// Retrieve the old state.
	oldState, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	resourceKey := string(urn.Type())
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("Resource type '%s' not found", resourceKey)
	}

	url := id + res.ReadPath

	var outputs map[string]interface{}
	customRes, isCustom := k.customResources[res.Path]
	var response map[string]interface{}
	switch {
	case isCustom && customRes.Read != nil:
		var exists bool
		response, exists, err = customRes.Read(ctx, oldState)
		if err != nil {
			return nil, err
		}
		if !exists {
			return &rpc.ReadResponse{Id: ""}, nil
		}
		outputs = response
	case res.ReadMethod == "HEAD":
		err = k.azureHead(ctx, url, res.APIVersion)
		response = oldState.Mappable()
		outputs = k.converter.BodyPropertiesToSDK(res.Response, response)
	case res.ReadMethod == "POST":
		bodyParams := map[string]interface{}{}
		queryParams := map[string]interface{}{
			"api-version": res.APIVersion,
		}
		response, err = k.azurePost(ctx, url, bodyParams, queryParams)
		outputs = k.converter.BodyPropertiesToSDK(res.Response, response)
	default:
		response, err = k.azureGet(ctx, url, res.APIVersion)
		outputs = k.converter.BodyPropertiesToSDK(res.Response, response)
	}
	if err != nil {
		if reqErr, ok := err.(*azure.RequestError); ok && reqErr.StatusCode == http.StatusNotFound {
			// 404 means that the resource was deleted.
			return &rpc.ReadResponse{Id: ""}, nil
		}
		return nil, err
	}

	// Extract old inputs from the `__inputs` field of the old state.
	inputs := parseCheckpointObject(oldState)
	if inputs == nil {
		// There may be no old state (i.e., importing a new resource).
		// Extract inputs from resource's ID and response body.
		pathItems, err := resources.ParseResourceID(id, res.Path)
		if err != nil {
			return nil, err
		}
		inputMap := k.converter.ResponseToSdkInputs(res.PutParameters, pathItems, response)
		inputs = resource.NewPropertyMapFromMap(inputMap)
	} else {
		// It's hard to infer the changes in the inputs shape based on the outputs without false positives.
		// The current approach is complicated but it's aimed to minimize the noise while refreshing:
		// 0. We have "old" inputs and outputs before refresh and "new" outputs read from Azure.
		// 1. Project old outputs to their corresponding input shape (exclude read-only properties).
		oldInputProjection := k.converter.SDKOutputsToSDKInputs(res.PutParameters, oldState.Mappable())
		// 2. Project new outputs to their corresponding input shape (exclude read-only properties).
		newInputProjection := k.converter.SDKOutputsToSDKInputs(res.PutParameters, outputs)
		// 3. Calculate the difference between two projections. This should give us actual significant changes
		// that happened in Azure between the last resource update and its current state.
		oldInputPropertyMap := resource.NewPropertyMapFromMap(oldInputProjection)
		newInputPropertyMap := resource.NewPropertyMapFromMap(newInputProjection)
		diff := oldInputPropertyMap.Diff(newInputPropertyMap)
		// 4. Apply this difference to the actual inputs (not a projection) that we have in state.
		inputs = applyDiff(inputs, diff)
	}

	// Store both outputs and inputs into the state.
	obj := checkpointObject(inputs, outputs)

	// Serialize and return RPC outputs.
	checkpoint, err := plugin.MarshalProperties(
		obj,
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	inputsRecord, err := plugin.MarshalProperties(
		inputs,
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.inputs", label), KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &rpc.ReadResponse{Id: id, Properties: checkpoint, Inputs: inputsRecord}, nil
}

// Update updates an existing resource with new values.
func (k *azureNextGenProvider) Update(ctx context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
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

	if req.GetPreview() {
		// The preview outputs are inputs + a limited list of outputs that are universally immutable.
		// We know that their values won't change, so it's safe to propagate the values to dependent
		// resources during the preview.
		outputs := k.converter.PreviewOutputs(inputs, res.Response)

		oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
			Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
		})
		if err != nil {
			return nil, err
		}

		stableOutputs := []string{"name", "location"}
		for _, name := range stableOutputs {
			key := resource.PropertyKey(name)
			if value, ok := oldState[key]; ok {
				outputs[key] = value
			}
		}

		previewOutputs, err := plugin.MarshalProperties(
			outputs,
			plugin.MarshalOptions{Label: fmt.Sprintf("%s.previewOutputs", label), KeepUnknowns: true, SkipNulls: true},
		)
		if err != nil {
			return nil, err
		}

		return &rpc.UpdateResponse{
			Properties: previewOutputs,
		}, nil
	}

	id, bodyParams, queryParams, err := k.prepareAzureRESTInputs(
		res.Path,
		res.PutParameters,
		inputs.Mappable(),
		map[string]interface{}{
			"subscriptionId": k.subscriptionID,
			"api-version":    res.APIVersion,
		},
	)
	if err != nil {
		return nil, err
	}

	var outputs map[string]interface{}
	customRes, isCustom := k.customResources[res.Path]
	switch {
	case isCustom && customRes.Update != nil:
		outputs, err = customRes.Update(ctx, inputs)
		if err != nil {
			return nil, err
		}
	default:
		response, err := k.azureCreateOrUpdate(ctx, id, bodyParams, queryParams, res.PutAsyncStyle)
		if err != nil {
			return nil, err
		}

		// Map the raw response to the shape of outputs that the SDKs expect.
		outputs = k.converter.BodyPropertiesToSDK(res.Response, response)
	}

	// Store both outputs and inputs into the state.
	obj := checkpointObject(inputs, outputs)

	// Serialize and return RPC outputs.
	checkpoint, err := plugin.MarshalProperties(
		obj,
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	return &rpc.UpdateResponse{
		Properties: checkpoint,
	}, nil
}

// Delete tears down an existing resource with the given ID. If it fails, the resource is assumed
// to still exist.
func (k *azureNextGenProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Delete(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)
	id := req.GetId()
	resourceKey := string(urn.Type())
	res, ok := k.resourceMap.Resources[resourceKey]
	if !ok {
		return nil, errors.Errorf("Resource type '%s' not found", resourceKey)
	}

	customRes, isCustom := k.customResources[res.Path]

	switch {
	case isCustom && customRes.Delete != nil:
		state, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
			Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: false, SkipNulls: true, KeepSecrets: false,
		})
		if err != nil {
			return nil, err
		}
		inputs := parseCheckpointObject(state)
		if inputs == nil {
			return nil, errors.Wrapf(err, "resource %s inputs are empty", label)
		}
		// Our hand-crafted implementation of DELETE operation.
		err = customRes.Delete(ctx, inputs)
		if err != nil {
			return nil, err
		}
	case res.Singleton:
		// Singleton resources can't be deleted (or created), set them to the default state.
		for _, param := range res.PutParameters {
			if param.Location == "body" {
				requestBody := k.converter.SdkPropertiesToRequestBody(param.Body.Properties, res.DefaultBody)

				queryParams := map[string]interface{}{"api-version": res.APIVersion}
				_, err := k.azureCreateOrUpdate(ctx, id, requestBody, queryParams, res.PutAsyncStyle)
				if err != nil {
					return nil, err
				}
			}
		}
	default:
		err := k.azureDelete(ctx, id, res.APIVersion, res.DeleteAsyncStyle)
		if err != nil {
			return nil, errors.Wrapf(err, "deleting %s", label)
		}
	}
	return &pbempty.Empty{}, nil
}

// Construct creates a new component resource.
func (k *azureNextGenProvider) Construct(_ context.Context, _ *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
	panic("Construct not implemented")
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (k *azureNextGenProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: k.version,
	}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (k *azureNextGenProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	// TODO
	return &pbempty.Empty{}, nil
}

func (k *azureNextGenProvider) azureCreateOrUpdate(
	ctx context.Context,
	id string,
	bodyProps map[string]interface{},
	queryParameters map[string]interface{},
	asyncStyle string) (map[string]interface{}, error) {

	decorators := []autorest.PrepareDecorator{
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(k.environment.ResourceManagerEndpoint),
		autorest.WithPath(id),
		autorest.WithQueryParameters(queryParameters),
	}
	if bodyProps != nil {
		decorators = append(decorators, autorest.WithJSON(bodyProps))
	}
	preparer := autorest.CreatePreparer(decorators...)
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return nil, err
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(
		k.client,
		prepReq,
		withSenderLogging(),
		azure.DoRetryWithRegistration(k.client),
	)
	if err != nil {
		return nil, err
	}

	// Some APIs are explicitly marked `x-ms-long-running-operation` and we are only supposed to do the
	// Future+WaitForCompletion+GetResult steps in that case. However, if we get 202, we don't want to
	// consider this a failure - so try following the awaiting protocol in case the service hasn't marked
	// its API as long-running by an oversight.
	if asyncStyle != "" || resp.StatusCode == http.StatusAccepted {
		// Ignore the style value for now, let go-autorest handle the headers.
		future, err := azure.NewFutureFromResponse(resp)
		if err != nil {
			return nil, err
		}
		err = future.WaitForCompletionRef(ctx, k.client)
		if err != nil {
			return nil, err
		}
		resp, err = future.GetResult(k.client)
		if err != nil {
			return nil, err
		}
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

func (k *azureNextGenProvider) azureDelete(ctx context.Context, id string, apiVersion string, asyncStyle string) error {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(k.environment.ResourceManagerEndpoint),
		autorest.WithPath(id),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return err
	}
	resp, err := autorest.SendWithSender(
		k.client,
		prepReq,
		withSenderLogging(),
		azure.DoRetryWithRegistration(k.client),
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
		err = future.WaitForCompletionRef(ctx, k.client)
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
		resp, err = future.GetResult(k.client)
		if err != nil {
			return err
		}
	}

	var result map[string]interface{}
	err = autorest.Respond(
		resp,
		k.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing(),
	)
	if err != nil {
		return err
	}
	return nil
}

// azureCanCreate asserts that a resource with a given ID and API version can be created
// or returns an error otherwise.
func (k *azureNextGenProvider) azureCanCreate(ctx context.Context, id string, res *resources.AzureAPIResource) error {
	queryParameters := map[string]interface{}{
		"api-version": res.APIVersion,
	}
	op := autorest.AsGet()
	switch res.ReadMethod {
	case "HEAD":
		op = autorest.AsHead()
	case "POST":
		op = autorest.AsPost()
	}
	preparer := autorest.CreatePreparer(
		op,
		autorest.WithBaseURL(k.environment.ResourceManagerEndpoint),
		autorest.WithPath(id + res.ReadPath),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return err
	}
	resp, err := autorest.SendWithSender(
		k.client,
		prepReq,
		withSenderLogging(),
		azure.DoRetryWithRegistration(k.client),
	)
	if err != nil {
		return err
	}

	switch {
	case http.StatusOK == resp.StatusCode && res.Singleton:
		// Singleton resources always exist, so OK is expected.
		return nil
	case http.StatusNotFound == resp.StatusCode && res.Singleton:
		return fmt.Errorf("parent resource does not exist for resource '%s'", id)
	case http.StatusOK == resp.StatusCode && res.DefaultBody != nil:
		// This resource is automatically created with a parent and set to its default state. It can be deleted though.
		// Validate that its current shape is in the default state to avoid unintended adoption and destructive
		// actions.
		// NOTE: We may reconsider and relax this restriction when we get more examples of such resources.
		// The difference between "take any singleton resource as-is" and "require default body for deletable resources"
		// isn't very principled but is based on what subjectively feels best for the current examples.
		var outputs map[string]interface{}
		err = autorest.Respond(
			resp,
			k.client.ByInspecting(),
			autorest.ByUnmarshallingJSON(&outputs),
			autorest.ByClosing())
		if err != nil {
			return err
		}
		if !k.converter.IsDefaultResponse(res.PutParameters, outputs, res.DefaultBody) {
			return fmt.Errorf("cannot create already existing subresource '%s'", id)
		}
		return nil
	case http.StatusNoContent == resp.StatusCode:
		return fmt.Errorf("cannot create already existing resource '%s'", id)
	case http.StatusOK == resp.StatusCode:
		// Usually, 200 means that the resource already exists and we shouldn't try to create it.
		// However, unfortunately, some APIs return 200 with an empty body for non-existing resources.
		// Our strategy here is to try to parse the response body and see if it's a valid non-empty JSON.
		// If it is, we assume the resource exists.
		var outputs map[string]interface{}
		err = autorest.Respond(
			resp,
			k.client.ByInspecting(),
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
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("cannot check existence of resource '%s': status code %d, %s", id, resp.StatusCode, body)
	}
}

func (k *azureNextGenProvider) azureHead(ctx context.Context, id string, apiVersion string) error {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(k.environment.ResourceManagerEndpoint),
		autorest.WithPath(id),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return err
	}
	resp, err := autorest.SendWithSender(
		k.client,
		prepReq,
		withSenderLogging(),
		azure.DoRetryWithRegistration(k.client),
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

func (k *azureNextGenProvider) azureGet(ctx context.Context, id string,
	apiVersion string) (map[string]interface{}, error) {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(k.environment.ResourceManagerEndpoint),
		autorest.WithPath(id),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return nil, err
	}
	resp, err := autorest.SendWithSender(
		k.client,
		prepReq,
		withSenderLogging(),
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
		forceRequestErrorForStatusNotFound,
		autorest.ByUnmarshallingJSON(&outputs),
		autorest.ByClosing())
	if err != nil {
		return nil, err
	}
	return outputs, nil
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

func (k *azureNextGenProvider) azurePost(
	ctx context.Context,
	id string,
	bodyProps map[string]interface{},
	queryParameters map[string]interface{}) (map[string]interface{}, error) {

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(k.environment.ResourceManagerEndpoint),
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
		withSenderLogging(),
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

func (k *azureNextGenProvider) prepareAzureRESTInputs(path string, parameters []resources.AzureAPIParameter, methodInputs,
	clientInputs map[string]interface{}) (string, map[string]interface{}, map[string]interface{}, error) {
	// Schema-driven mapping of inputs into Autorest id/body/query
	params := map[string]map[string]interface{}{
		"body": nil,
		"query": {
			"api-version": clientInputs["api-version"],
		},
		"path": {},
	}

	for _, param := range parameters {
		if param.Location == "body" {
			params["body"] = k.converter.SdkPropertiesToRequestBody(param.Body.Properties, methodInputs)
		} else {
			var val interface{}
			var has bool
			sdkName := param.Name
			if param.Value.SdkName != "" {
				sdkName = param.Value.SdkName
			}
			// Look in both `method` and `client` inputs with `method` first
			val, has = methodInputs[sdkName]
			if !has {
				val, has = clientInputs[sdkName]
			}
			if has {
				params[param.Location][param.Name] = val
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

func (k *azureNextGenProvider) setLoggingContext(ctx context.Context) {
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

func (k *azureNextGenProvider) getConfig(configName, envName string) string {
	if val, ok := k.config[configName]; ok {
		return val
	}

	return os.Getenv(envName)
}

func (k *azureNextGenProvider) getAuthConfig() (*authentication.Config, error) {
	auxTenantsString := k.getConfig("auxiliaryTenantIds", "ARM_AUXILIARY_TENANT_IDS")
	var auxTenants []string
	if auxTenantsString != "" {
		err := json.Unmarshal([]byte(auxTenantsString), &auxTenants)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal '%s' as Auxiliary Tenants", auxTenantsString)
		}
	}
	useMsi := k.getConfig("useMsi", "ARM_USE_MSI") == "true"
	envName := k.getConfig("environment", "ARM_ENVIRONMENT")
	if envName == "" {
		envName = "public"
	}
	builder := &authentication.Builder{
		SubscriptionID: k.getConfig("subscriptionId", "ARM_SUBSCRIPTION_ID"),
		ClientID:       k.getConfig("clientId", "ARM_CLIENT_ID"),
		ClientSecret:   k.getConfig("clientSecret", "ARM_CLIENT_SECRET"),
		TenantID:       k.getConfig("tenantId", "ARM_TENANT_ID"),
		Environment:    envName,
		ClientCertPath: k.getConfig("clientCertificatePath", "ARM_CLIENT_CERTIFICATE_PATH"),
		ClientCertPassword: k.getConfig("clientCertificatePassword",
			"ARM_CLIENT_CERTIFICATE_PASSWORD"),
		MsiEndpoint:        k.getConfig("msiEndpoint", "ARM_MSI_ENDPOINT"),
		AuxiliaryTenantIDs: auxTenants,

		// Feature Toggles
		SupportsClientCertAuth:         true,
		SupportsClientSecretAuth:       true,
		SupportsManagedServiceIdentity: useMsi,
		SupportsAzureCliToken:          true,
		SupportsAuxiliaryTenants:       len(auxTenants) > 0,
	}

	return builder.Build()
}

func (k *azureNextGenProvider) getAuthorizers(authConfig *authentication.Config) (tokenAuth autorest.Authorizer,
	bearerAuth autorest.Authorizer, err error) {
	buildSender := sender.BuildSender("AzureNextGen")
	oauthConfig, err := authConfig.BuildOAuthConfig(k.environment.ActiveDirectoryEndpoint)
	if err != nil {
		return nil, nil, err
	}

	if oauthConfig == nil {
		return nil, nil, fmt.Errorf("unable to configure OAuthConfig for tenant %s", authConfig.TenantID)
	}

	tokenAuth, err = authConfig.GetAuthorizationToken(buildSender, oauthConfig, k.environment.ResourceManagerEndpoint)
	if err != nil {
		return nil, nil, err
	}

	bearerAuth = authConfig.BearerAuthorizerCallback(buildSender, oauthConfig)
	return tokenAuth, bearerAuth, nil
}

func (k *azureNextGenProvider) getOAuthToken(ctx context.Context, auth *authentication.Config, endpoint string) (string, error) {
	buildSender := sender.BuildSender("AzureNextGen")
	oauthConfig, err := auth.BuildOAuthConfig(k.environment.ActiveDirectoryEndpoint)
	authorizer, err := auth.GetAuthorizationToken(buildSender, oauthConfig, endpoint)
	if err != nil {
		return "", fmt.Errorf("getting authorization token: %w", err)
	}
	ba, ok := authorizer.(*autorest.BearerAuthorizer)
	if !ok {
		return "", fmt.Errorf("converting %T to a BearerAuthorizer", authorizer)
	}
	tokenProvider := ba.TokenProvider()
	// the ordering is important here, prefer RefresherWithContext if available
	if refresher, ok := tokenProvider.(adal.RefresherWithContext); ok {
		err = refresher.EnsureFreshWithContext(ctx)
	} else if refresher, ok := tokenProvider.(adal.Refresher); ok {
		err = refresher.EnsureFresh()
	}
	if err != nil {
		return "", fmt.Errorf("error refreshing token from %T", tokenProvider)
	}
	token := tokenProvider.OAuthToken()
	if token == "" {
		return "", fmt.Errorf("empty token from %T", tokenProvider)
	}
	return token, nil
}

// getUserAgent returns a User Agent string for the current provider configuration.
func (k *azureNextGenProvider) getUserAgent() string {
	partnerID := PulumiPartnerID
	customPartnerID := k.getConfig("partnerId", "ARM_PARTNER_ID")
	if customPartnerID != "" {
		partnerID = customPartnerID
	} else {
		disablePartnerID := k.getConfig("disablePulumiPartnerId", "ARM_DISABLE_PULUMI_PARTNER_ID")
		if disablePartnerID == "true" {
			partnerID = ""
		}
	}
	return buildUserAgent(partnerID)
}

// buildUserAgent composes a User Agent string with the provided partner ID.
func buildUserAgent(partnerID string) (userAgent string) {
	userAgent = strings.TrimSpace(fmt.Sprintf("%s pulumi-azure-nextgen/%s",
		autorest.UserAgent(), version.Version))

	// append the CloudShell version to the user agent if it exists
	if azureAgent := os.Getenv("AZURE_HTTP_USER_AGENT"); azureAgent != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, azureAgent)
	}

	// Append partner ID, if it's defined.
	if partnerID != "" {
		userAgent = fmt.Sprintf("%s pid-%s", userAgent, partnerID)
	}

	glog.V(9).Infof("AzureNextGen User Agent: %s", userAgent)
	return
}

// withSenderLogging decorates a sender with logging of HTTP methods, URLs, and statuses.
func withSenderLogging() autorest.SendDecorator {
	return func(s autorest.Sender) autorest.Sender {
		return autorest.SenderFunc(func(r *http.Request) (*http.Response, error) {
			glog.V(9).Infof("Sending %s %s", r.Method, r.URL)
			resp, err := s.Do(r)
			if err != nil {
				glog.V(9).Infof("%s %s received error '%v'", r.Method, r.URL, err)
			} else {
				glog.V(9).Infof("%s %s received %s", r.Method, r.URL, resp.Status)
			}
			return resp, err
		})
	}
}

// checkpointObject puts inputs in the `__inputs` field of the state.
func checkpointObject(inputs resource.PropertyMap, outputs map[string]interface{}) resource.PropertyMap {
	object := resource.NewPropertyMapFromMap(outputs)
	object["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	return object

}

// parseCheckpointObject returns inputs that are saved in the `__inputs` field of the state.
func parseCheckpointObject(obj resource.PropertyMap) resource.PropertyMap {
	if inputs, ok := obj["__inputs"]; ok {
		return inputs.ObjectValue()
	}

	return nil
}
