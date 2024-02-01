// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/manicminer/hamilton/environments"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/deepcopy"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/rpcutil/rpcerror"
	"github.com/segmentio/encoding/json"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/arm2pulumi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/defaults"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources/customresources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// Microsoft's Pulumi Partner ID.
	PulumiPartnerID = "a90539d8-a7a6-5826-95c4-1fbef22d4b22"
	requestFormat   = `HTTP Request Begin %[1]s %[2]s ===================================================
%[3]s
===================================================== HTTP Request End %[1]s %[2]s
`
	responseFormat = `HTTP Response Begin %[1]s [%[2]s ===================================================
%[3]s
===================================================== HTTP Response End %[1]s %[2]s
`
	apiVersionResources    = "2020-10-01"
	createBeforeDeleteFlag = "__createBeforeDelete"
)

type azureNativeProvider struct {
	rpc.UnimplementedResourceProviderServer

	host            *provider.HostClient
	name            string
	version         string
	subscriptionID  string
	environment     azure.Environment
	client          autorest.Client
	resourceMap     *resources.PartialAzureAPIMetadata
	config          map[string]string
	schemaBytes     []byte
	metadataBytes   []byte
	fullPkgSpec     *schema.PackageSpec
	fullResourceMap *resources.AzureAPIMetadata
	converter       *convert.SdkShapeConverter
	customResources map[string]*customresources.CustomResource
	rgLocationMap   map[string]string
	lookupType      resources.TypeLookupFunc
}

func makeProvider(host *provider.HostClient, name, version string, schemaBytes []byte,
	azureAPIResourcesBytes []byte) (rpc.ResourceProviderServer, error) {
	autorest.Count429AsRetry = false
	// Creating a REST client, defaulting to Pulumi Partner ID until the Configure method is invoked.
	client := autorest.NewClientWithUserAgent(buildUserAgent(PulumiPartnerID))
	// Log requests
	client.RequestInspector = withInspection()
	// Log responses
	client.ResponseInspector = byInspecting()

	resourceMap, err := LoadMetadataPartial(azureAPIResourcesBytes)
	if err != nil {
		return nil, err
	}

	converter := convert.NewSdkShapeConverterPartial(resourceMap.Types)

	// Return the new provider
	p := &azureNativeProvider{
		host:          host,
		name:          name,
		version:       version,
		client:        client,
		resourceMap:   resourceMap,
		config:        map[string]string{},
		schemaBytes:   schemaBytes,
		converter:     &converter,
		rgLocationMap: map[string]string{},
	}
	p.lookupType = p.lookupTypeDefault
	return p, nil
}

// LoadMetadataPartial partially deserializes the provided json byte array into an AzureAPIMetadata
// in memory
func LoadMetadataPartial(azureAPIResourcesBytes []byte) (*resources.PartialAzureAPIMetadata, error) {
	var resourceMap resources.PartialAzureAPIMetadata

	if _, err := json.Parse(azureAPIResourcesBytes, &resourceMap, json.ZeroCopy); err != nil {
		return nil, errors.Wrap(err, "unmarshalling resource map")
	}
	return &resourceMap, nil
}

// LoadMetadata deserializes the provided json byte array into an AzureAPIMetadata in memory
func LoadMetadata(azureAPIResourcesBytes []byte) (*resources.AzureAPIMetadata, error) {
	var resourceMap resources.AzureAPIMetadata

	if _, err := json.Parse(azureAPIResourcesBytes, &resourceMap, json.ZeroCopy); err != nil {
		return nil, errors.Wrap(err, "unmarshalling resource map")
	}
	return &resourceMap, nil
}

func (k *azureNativeProvider) getFullMetadata() (*resources.AzureAPIMetadata, error) {
	if k.fullResourceMap != nil {
		return k.fullResourceMap, nil
	}

	if _, err := json.Parse(k.metadataBytes, k.fullResourceMap, json.ZeroCopy); err != nil {
		return nil, errors.Wrap(err, "unmarshalling metadata")
	}

	return k.fullResourceMap, nil
}

// Looks up a type reference, with or without the #/types/ prefix, in the resource map.
// Typically used via lookupType so it can be overridden for tests.
func (k *azureNativeProvider) lookupTypeDefault(ref string) (*resources.AzureAPIType, bool, error) {
	typeName := strings.TrimPrefix(ref, "#/types/")
	t, ok, err := k.resourceMap.Types.Get(typeName)
	return &t, ok, err
}

func (k *azureNativeProvider) LookupResource(resourceType string) (resources.AzureAPIResource, bool, error) {
	return k.resourceMap.Resources.Get(resourceType)
}

func (p *azureNativeProvider) Attach(context context.Context, req *rpc.PluginAttach) (*emptypb.Empty, error) {
	host, err := provider.NewHostClient(req.GetAddress())
	if err != nil {
		return nil, err
	}
	p.host = host
	return &pbempty.Empty{}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *azureNativeProvider) Configure(ctx context.Context,
	req *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {

	for key, val := range req.GetVariables() {
		k.config[strings.TrimPrefix(key, "azure-native:config:")] = val
	}

	k.setLoggingContext(ctx)

	authConfig, err := k.getAuthConfig()
	if err != nil {
		return nil, err
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

	hamiltonEnv := k.autorestEnvToHamiltonEnv()

	// The ctx Context given by gRPC is request-scoped and will be canceled after this request. We
	// need the authorizers to function across requests.
	authCtx := context.Background()

	tokenAuth, bearerAuth, err := k.makeAuthorizerFactories(authCtx, authConfig)
	if err != nil {
		return nil, fmt.Errorf("auth setup: %w", err)
	}

	resourceManagerAuth, err := tokenAuth(hamiltonEnv.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building authorizer for %s: %w", hamiltonEnv.ResourceManager.Endpoint, err)
	}

	resourceManagerBearerAuth, err := bearerAuth(hamiltonEnv.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building bearer authorizer for %s: %w", hamiltonEnv.ResourceManager.Endpoint, err)
	}

	keyVaultBearerAuth, err := bearerAuth(hamiltonEnv.KeyVault)
	if err != nil {
		return nil, fmt.Errorf("building bearer authorizer for %s: %w", hamiltonEnv.KeyVault.Endpoint, err)
	}

	k.subscriptionID = authConfig.SubscriptionID
	k.client.Authorizer = resourceManagerAuth
	k.client.UserAgent = k.getUserAgent()

	azCoreTokenCredential := azCoreTokenCredential{p: k}
	var azureClient customresources.AzureClient = k
	k.customResources, err = customresources.BuildCustomResources(&env, azureClient, k.subscriptionID,
		resourceManagerBearerAuth, resourceManagerAuth, keyVaultBearerAuth,
		k.client.UserAgent, azCoreTokenCredential)
	if err != nil {
		return nil, fmt.Errorf("initializing custom resources: %w", err)
	}

	return &rpc.ConfigureResponse{
		SupportsPreview: true,
	}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *azureNativeProvider) Invoke(ctx context.Context, req *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	label := fmt.Sprintf("%s.Invoke(%s)", k.name, req.Tok)
	logging.V(9).Infof("%s executing", label)

	args, err := plugin.UnmarshalProperties(req.GetArgs(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.args", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	var outputs map[string]interface{}
	switch req.Tok {
	case "azure-native:authorization:getClientConfig":
		auth, err := k.getAuthConfig()
		if err != nil {
			return nil, fmt.Errorf("getting auth config: %w", err)
		}
		objectId := ""
		if auth.GetAuthenticatedObjectID != nil {
			objectIdPtr, err := auth.GetAuthenticatedObjectID(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting authenticated object ID: %w", err)
			}
			if objectIdPtr == nil {
				return nil, fmt.Errorf("getting authenticated object ID")
			}
			objectId = *objectIdPtr
		}
		outputs = map[string]interface{}{
			"clientId":       auth.ClientID,
			"objectId":       objectId,
			"subscriptionId": auth.SubscriptionID,
			"tenantId":       auth.TenantID,
		}
	case "azure-native:authorization:getClientToken":
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
	case "azure-native:armtemplate:decode":
		var text string
		if textArg := args["text"]; textArg.HasValue() && textArg.IsString() {
			text = textArg.StringValue()
		} else {
			return nil, errors.New("missing required field 'text' of type 'string'")
		}

		if k.fullPkgSpec == nil {
			if _, err := json.Parse(k.schemaBytes, k.fullPkgSpec, json.ZeroCopy); err != nil {
				return nil, fmt.Errorf("deserializing schema: %w", err)
			}
		}

		resourceMap, err := k.getFullMetadata()
		if err != nil {
			return nil, fmt.Errorf("deserializing ")
		}

		res, err := arm2pulumi.DecodeTemplate(k.fullPkgSpec, resourceMap, text)
		if err != nil {
			return nil, fmt.Errorf("rendering template: %w", err)
		}
		outputs = map[string]interface{}{"result": res}
	default:
		res, ok, err := k.resourceMap.Invokes.Get(req.Tok)
		if err != nil {
			return nil, errors.Errorf("Decoding invoke spec %s", req.Tok)
		}
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
			nil,
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
			if body == nil {
				body = map[string]interface{}{}
			}
			response, err = k.azurePost(ctx, id, body, query)
		} else {
			return nil, errors.Errorf("neither GET nor POST is defined for %s", label)
		}

		if err != nil {
			return nil, errors.Wrapf(err, "request failed %v", id)
		}

		// Map the raw response to the shape of outputs that the SDKs expect.
		outputs = k.converter.ResponseBodyToSdkOutputs(res.Response, response)
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
func (k *azureNativeProvider) StreamInvoke(_ *rpc.InvokeRequest, _ rpc.ResourceProvider_StreamInvokeServer) error {
	return status.Error(codes.Unimplemented, "StreamInvoke is not yet implemented")
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (k *azureNativeProvider) Check(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Check(%s)", k.name, urn)
	logging.V(9).Infof("%s executing", label)
	var failures []*rpc.CheckFailure

	// Deserialize RPC inputs.
	oldResInputs := req.GetOlds()
	olds, err := plugin.UnmarshalProperties(oldResInputs, plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}
	newResInputs := req.GetNews()
	news, err := plugin.UnmarshalProperties(newResInputs, plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.news", label), KeepUnknowns: true, SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	resourceKey := string(urn.Type())
	res, ok, err := k.LookupResource(resourceKey)
	if err != nil {
		return nil, errors.Errorf("Decoding resource spec %s", resourceKey)
	}
	if !ok {
		return nil, errors.Errorf("Resource type %s not found", resourceKey)
	}

	k.applyDefaults(ctx, req.Urn, req.RandomSeed, res, olds, news)
	inputMap := news.Mappable()

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

	resInputs, err := plugin.MarshalProperties(news, plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.resInputs", label), KeepUnknowns: true})
	if err != nil {
		return nil, err
	}

	return &rpc.CheckResponse{Inputs: resInputs, Failures: failures}, nil
}

// Get a default location property for the given inputs.
func (k *azureNativeProvider) getDefaultLocation(ctx context.Context, olds, news resource.PropertyMap) *resource.PropertyValue {
	result := func(s string) *resource.PropertyValue {
		p := resource.NewStringProperty(s)
		return &p
	}

	// 1. Check if old inputs define a location.
	if v, ok := olds["location"]; ok {
		return &v
	}

	// 2. Check if the config has a fixed location.
	if v, ok := k.config["location"]; ok {
		return result(v)
	}

	// 3. If there is a link to a resource group, try to use its location.
	rg, ok := news["resourceGroupName"]
	if !ok {
		return nil
	}

	// 3a. Resource group is unknown yet - make location unknown too.
	if rg.ContainsUnknowns() {
		computed := resource.MakeComputed(resource.NewStringProperty(""))
		return &computed
	}

	// 3b. Resource group name is known and its location is already in the cache: use the cached value.
	rgName := rg.StringValue()
	if v, ok := k.rgLocationMap[rgName]; ok {
		return result(v)
	}

	// 3c. Retrieve the resource group's properties from ARM API.
	id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", k.subscriptionID, rgName)
	response, err := k.azureGet(ctx, id, apiVersionResources)
	if err != nil {
		logging.V(9).Infof("failed to lookup the location of resource group %q: %v", rgName, err)
		return nil
	}

	v, ok := response["location"].(string)
	if !ok {
		logging.V(9).Infof("no location for resource group %q", rgName)
		return nil
	}

	// Save the location to the cache and return the value.
	k.rgLocationMap[rgName] = v
	return result(v)
}

func (k *azureNativeProvider) getDefaultName(urn string, randomSeed []byte, strategy resources.AutoNameKind,
	key resource.PropertyKey, olds resource.PropertyMap) (resource.PropertyValue, bool) {
	if v, ok := olds[key]; ok {
		if vf, ok := olds[createBeforeDeleteFlag]; ok && vf.IsBool() {
			return v, vf.BoolValue()
		}
		return v, false
	}

	urnParts := strings.Split(urn, "::")
	name := urnParts[len(urnParts)-1]
	switch strategy {
	case resources.AutoNameRandom:
		// Resource name is URN name + random suffix.
		random, err := resource.NewUniqueName(randomSeed, name, 8, 0, nil)
		contract.AssertNoError(err)
		return resource.NewStringProperty(random), true
	case resources.AutoNameUuid:
		// Resource name is a random UUID. We need to do a similar trick as NewUniqueName so that this is
		// deterministic by randomSeed. We simply ask NewUniqueName for a 32 byte random hex name.
		hexID, err := resource.NewUniqueName(randomSeed, "", 32, 0, nil)
		contract.AssertNoError(err)
		return resource.NewStringProperty(uuid.MustParse(hexID).String()), true
	case resources.AutoNameCopy:
		// Resource name is just a copy of the URN name.
		return resource.NewStringProperty(name), false
	default:
		panic(fmt.Sprintf("unknown auto-naming strategy %q", strategy))
	}
}

// Apply default values (e.g., location) to user's inputs.
func (k *azureNativeProvider) applyDefaults(ctx context.Context, urn string, randomSeed []byte,
	res resources.AzureAPIResource, olds, news resource.PropertyMap) {
	for _, par := range res.PutParameters {
		sdkName := par.Name
		if par.Value != nil && par.Value.SdkName != "" {
			sdkName = par.Value.SdkName
		}

		// Auto-naming.
		key := resource.PropertyKey(sdkName)
		if !news.HasValue(key) && par.Value != nil && par.Value.AutoName != "" {
			name, randomlyNamed := k.getDefaultName(urn, randomSeed, par.Value.AutoName, key, olds)
			news[key] = name
			if randomlyNamed {
				news[createBeforeDeleteFlag] = resource.NewBoolProperty(true)
			}
		}

		// Auto-location.
		if !news.HasValue("location") && par.Body != nil && !res.AutoLocationDisabled {
			if _, ok := par.Body.Properties["location"]; ok {
				v := k.getDefaultLocation(ctx, olds, news)
				if v != nil {
					news["location"] = *v
				}
			}
		}
	}
}

// validateType checks the all properties and required properties of the given type and value map.
func (k *azureNativeProvider) validateType(ctx string, typ *resources.AzureAPIType,
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
			reason := fmt.Sprintf("missing required property '%s'", propCtx)
			if propCtx == "location" {
				reason += ". Either set it explicitly or configure it with 'pulumi config set azure-native:location <value>'."
			}
			failures = append(failures, &rpc.CheckFailure{
				Reason: reason,
			})
		}
	}

	for name, prop := range typ.Properties {
		p := prop // https://go.dev/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
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
func (k *azureNativeProvider) validateProperty(ctx string, prop *resources.AzureAPIProperty,
	value interface{}) []*rpc.CheckFailure {
	var failures []*rpc.CheckFailure

	if _, ok := value.(resource.Computed); ok {
		return failures
	}

	if prop == nil || prop.Ref == convert.TypeAny {
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
		typ, ok, err := k.lookupTypeDefault(prop.Ref)
		if err != nil {
			return append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("error decoding type spec '%s': %v", typeName, err),
			})
		}
		if !ok {
			return append(failures, &rpc.CheckFailure{
				Reason: fmt.Sprintf("schema type '%s' not found", typeName),
			})
		}

		failures = append(failures, k.validateType(ctx, typ, value)...)
	case []interface{}:
		if prop.Type != "array" && !prop.IsStringSet {
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
			// TODO: Support ECMA-262 regexp https://github.com/pulumi/pulumi-azure-native/issues/164
			if err == nil && !pattern.MatchString(value) {
				failures = append(failures, &rpc.CheckFailure{
					Reason: fmt.Sprintf("'%s' does not match expression '%s'", ctx, prop.Pattern),
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

func (k *azureNativeProvider) GetSchema(_ context.Context, req *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}

	return &rpc.GetSchemaResponse{Schema: string(k.schemaBytes)}, nil
}

// CheckConfig validates the configuration for this provider.
func (k *azureNativeProvider) CheckConfig(_ context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (k *azureNativeProvider) DiffConfig(context.Context, *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	return &rpc.DiffResponse{
		Changes:             0,
		Replaces:            []string{},
		Stables:             []string{},
		DeleteBeforeReplace: false,
	}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (k *azureNativeProvider) Diff(_ context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Diff(%s)", k.name, urn)

	// Retrieve the old state.
	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	// Extract old inputs from the OldInputs, if available, or the `__inputs` field of the old state as a fallback.
	oldInputs, err := parseCheckpointObject(req.OldInputs, oldState, label)
	if err != nil {
		return nil, err
	}
	if oldInputs == nil {
		// Protect against a crash for the transition from pre-__inputs state files.
		// This shouldn't happen in any real user's stack.
		oldInputs = resource.PropertyMap{}
		logging.V(9).Infof("no __inputs found for '%s'", urn)
	}

	// Get new resource inputs. The user is submitting these as an update.
	newResInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.news", label),
		KeepUnknowns: true,
		SkipNulls:    true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs %v %v",
			oldState, newResInputs)
	}

	// Get the resource definition for looking up additional metadata.
	resourceKey := string(urn.Type())
	res, ok, err := k.LookupResource(resourceKey)
	if err != nil {
		return nil, errors.Errorf("Decoding resource spec %s", resourceKey)
	}
	if !ok {
		return nil, errors.Errorf("Resource type %s not found", resourceKey)
	}

	detailedDiff := diff(k.lookupTypeDefault, res, oldInputs, newResInputs)
	if detailedDiff == nil {
		return &rpc.DiffResponse{
			Changes:             rpc.DiffResponse_DIFF_NONE,
			Replaces:            []string{},
			Stables:             []string{},
			DeleteBeforeReplace: false,
		}, nil
	}

	// Based on the detailed diff above, calculate the list of changes and replacements.
	changes, replaces := calculateChangesAndReplacements(detailedDiff, oldInputs, newResInputs, oldState, res)

	// TODO: implement create-before-delete for children of randomly auto-named resources.
	deleteBeforeReplace := len(replaces) > 0
	if v, ok := oldInputs[createBeforeDeleteFlag]; ok && v.IsBool() {
		deleteBeforeReplace = !v.BoolValue()
	}
	changeType := rpc.DiffResponse_DIFF_NONE
	if len(changes) > 0 {
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
func (k *azureNativeProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Create(%s)", k.name, urn)
	logging.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), KeepUnknowns: true, SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	resourceKey := string(urn.Type())
	res, ok, err := k.LookupResource(resourceKey)
	if err != nil {
		return nil, errors.Errorf("Decoding resource spec %s", resourceKey)
	}
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

	// Construct ARM REST API body and query from inputs
	id, bodyParams, queryParams, err := k.prepareAzureRESTInputs(
		res.Path,
		res.PutParameters,
		res.RequiredContainers,
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
		_, exists, err := customRes.Read(ctx, id, inputs)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, fmt.Errorf("cannot create already existing resource %q", id)
		}

		ctx, cancel := azureContext(ctx, req.Timeout)
		defer cancel()

		// Create the custom resource and retrieve its outputs, which already match the SDK shape.
		outputs, err = customRes.Create(ctx, inputs)
		if err != nil {
			return nil, azureError(err)
		}
	default:
		// First check if the resource already exists - we want to try our best to avoid updating instead of creating here
		// (though it's technically impossible since the only operation supported is an upsert).
		err = k.azureCanCreate(ctx, id, &res)
		if err != nil {
			return nil, err
		}

		ctx, cancel := azureContext(ctx, req.Timeout)
		defer cancel()

		k.setUnsetSubresourcePropertiesToDefaults(res, bodyParams)

		// Submit the `PUT` against the ARM endpoint
		response, created, err := k.azureCreateOrUpdate(ctx, id, bodyParams, queryParams, res.UpdateMethod, res.PutAsyncStyle)
		if err != nil {
			if created {
				// Resource was created but failed to fully initialize.
				// Try reading its state by ID and return a partial error if succeeded.
				checkpoint, getErr := k.currentResourceStateCheckpoint(ctx, id, res, inputs)
				if getErr != nil {
					return nil, azureError(errors.Wrapf(err, "resource partially created but read failed %s", getErr))
				}
				return nil, partialError(id, err, checkpoint, req.GetProperties())
			}
			return nil, azureError(err)
		}

		// Read the canonical ID from the response.
		if azureId, ok := response["id"].(string); ok {
			id = azureId
		}

		// Map the raw response to the shape of outputs that the SDKs expect.
		outputs = k.converter.ResponseBodyToSdkOutputs(res.Response, response)
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

// Properties pointing to sub-resources that can be maintained as separate resources might not be
// present in the inputs because the user wants to manage them as standalone resources. However,
// auch a property might be required by Azure even if it's not annotated as such in the spec, e.g.,
// Key Vault's accessPolicies. Therefore, we set these properties to their default value here,
// an empty array.
// During create, no sub-resources can exist yet so there's no danger of overwriting existing values.
//
// Implementation note: we should make it possible to write custom resources that call code from
// the default implementation as needed. This would allow us to cleanly implement special logic
// like for Key Vault into custom resources without duplicating much code. In the Key Vault case,
// the custom Read() would look like
//
//	provider.azureCanCreate(ctx, id, &res)
//	setUnsetSubresourcePropertiesToDefaults(res, bodyParams) // custom
//	k.azureCreateOrUpdate
//	...
func (k *azureNativeProvider) setUnsetSubresourcePropertiesToDefaults(res resources.AzureAPIResource, bodyParams map[string]interface{}) {
	unset := k.findUnsetPropertiesToMaintain(&res, bodyParams)
	for _, p := range unset {
		cur := bodyParams
		for _, pathEl := range p.path[:len(p.path)-1] {
			curObj, ok := cur[pathEl]
			if !ok {
				newContainer := map[string]any{}
				cur[pathEl] = newContainer
				cur = newContainer
				continue
			}
			cur, ok = curObj.(map[string]any)
			if !ok {
				break
			}
		}

		cur[p.path[len(p.path)-1]] = []any{}
	}
}

// currentResourceStateCheckpoint reads the resource state by ID, converts it to outputs map, and
// produces a checkpoint with these outputs and given inputs.
func (k *azureNativeProvider) currentResourceStateCheckpoint(ctx context.Context, id string, res resources.AzureAPIResource, inputs resource.PropertyMap) (*structpb.Struct, error) {
	getResp, getErr := k.azureGet(ctx, id, res.APIVersion)
	if getErr != nil {
		return nil, getErr
	}
	outputs := k.converter.ResponseBodyToSdkOutputs(res.Response, getResp)
	obj := checkpointObject(inputs, outputs)
	return plugin.MarshalProperties(
		obj,
		plugin.MarshalOptions{
			Label:        "currentResourceStateCheckpoint.checkpoint",
			KeepSecrets:  true,
			KeepUnknowns: true,
			SkipNulls:    true,
		},
	)
}

// Read the current live state associated with a resource.
func (k *azureNativeProvider) Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Read(%s)", k.name, urn)
	logging.V(9).Infof("%s executing", label)
	id := req.GetId()

	// Retrieve the old state.
	oldState, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	resourceKey := string(urn.Type())
	res, ok, err := k.LookupResource(resourceKey)
	if err != nil {
		return nil, errors.Errorf("Decoding resource spec %s", resourceKey)
	}
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
		response, exists, err = customRes.Read(ctx, id, oldState)
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
		outputs = k.converter.ResponseBodyToSdkOutputs(res.Response, response)
	case res.ReadMethod == "POST":
		bodyParams := map[string]interface{}{}
		queryParams := map[string]interface{}{
			"api-version": res.APIVersion,
		}
		response, err = k.azurePost(ctx, url, bodyParams, queryParams)
		outputs = k.converter.ResponseBodyToSdkOutputs(res.Response, response)
	default:
		response, err = k.azureGet(ctx, url, res.APIVersion)
		outputs = k.converter.ResponseBodyToSdkOutputs(res.Response, response)
	}
	if err != nil {
		if reqErr, ok := err.(*azure.RequestError); ok && reqErr.StatusCode == http.StatusNotFound {
			// 404 means that the resource was deleted.
			return &rpc.ReadResponse{Id: ""}, nil
		}
		return nil, err
	}

	// Extract old inputs from the `__inputs` field of the old state.
	inputs, err := parseCheckpointObject(req.GetInputs(), oldState, label)
	if err != nil {
		return nil, err
	}
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
		// We can't read inputs directly, so we calculate an approximation based on old inputs,
		// old outputs, and new outputs. Inputs are usually a subset of outputs, but the complication
		// is that a lot of outputs are default values assigned on the service side, and we don't
		// want to copy all of them to inputs to avoid noisy diffs.
		// Instead, we compare old outputs to new outputs, and if properties changed, we assume they
		// aren't default values anymore, and we need to copy them to new inputs.

		// The current approach is complicated but it's aimed to minimize the noise while refreshing:
		// 0. We have "old" inputs and outputs before refresh and "new" outputs read from Azure.

		// 1. If we previously reset inputs to their default value, remove them so we don't get them in
		// the projected output. This would cause unnecessary changes on refresh.
		plainOldState := mappableOldState(res, oldState)
		// 2. Project old outputs to their corresponding input shape (exclude read-only properties).
		oldInputProjection := k.converter.SdkOutputsToSdkInputs(res.PutParameters, plainOldState)
		// 3a. Remove sub-resource properties from new outputs which weren't set in the old inputs.
		// If the user didn't specify them inline originally, we don't want to push them into the inputs now.
		outputsWithoutIgnores := k.removeUnsetSubResourceProperties(ctx, urn, outputs, inputs, &res)
		// 3b. Project new outputs to their corresponding input shape (exclude read-only properties).
		newInputProjection := k.converter.SdkOutputsToSdkInputs(res.PutParameters, outputsWithoutIgnores)
		// 4. Calculate the difference between two projections. This should give us actual significant changes
		// that happened in Azure between the last resource update and its current state.
		oldInputPropertyMap := resource.NewPropertyMapFromMap(oldInputProjection)
		newInputPropertyMap := resource.NewPropertyMapFromMap(newInputProjection)
		diff := oldInputPropertyMap.Diff(newInputPropertyMap)
		// 5. Apply this difference to the actual inputs (not a projection) that we have in state.
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

// Converts oldState into a serializable object map, with the resource's default values from metadata removed.
func mappableOldState(res resources.AzureAPIResource, oldState resource.PropertyMap) map[string]interface{} {
	plainOldState := oldState.Mappable()
	previousInputsRaw, ok := plainOldState["__inputs"]
	if !ok {
		return plainOldState
	}
	previousInputs := previousInputsRaw.(map[string]interface{})

	for property, defaultValue := range res.DefaultProperties {
		_, wasInPreviousInputs := previousInputs[property]
		val, ok := plainOldState[property]
		if ok && wasInPreviousInputs && reflect.DeepEqual(val, defaultValue) {
			logging.V(5).Infof("removing property %q with default value from old state", property)
			delete(plainOldState, property)
		}
	}
	return plainOldState
}

// removeUnsetSubResourceProperties removes sub-resource properties from new outputs which weren't set in the old inputs.
// If the user didn't specify them inline originally, we don't want to push them into the inputs now.
func (k *azureNativeProvider) removeUnsetSubResourceProperties(ctx context.Context, urn resource.URN, sdkResponse map[string]interface{}, oldInputs resource.PropertyMap, res *resources.AzureAPIResource) map[string]interface{} {
	propertiesToRemove := k.findUnsetPropertiesToMaintain(res, oldInputs.Mappable())

	if len(propertiesToRemove) == 0 {
		return sdkResponse
	}

	// Take a deep copy so we don't modify the original which is also used later for diffing.
	copy := deepcopy.Copy(sdkResponse)
	result, ok := copy.(map[string]interface{})
	if !ok {
		// This should never happen.
		k.host.Log(ctx, diag.Warning, urn, "Failed to remove unset sub-resource properties. Please report this issue. See verbose logs for more information.")
		logging.V(9).Infof("failed to remove unset sub-resource properties: failed to cast copy value, expected map[string]interface{}, found %v", copy)
		return sdkResponse
	}

	for _, prop := range propertiesToRemove {
		deleteFromMap(result, prop.path)
	}
	return result
}

func deleteFromMap(m map[string]interface{}, path []string) bool {
	container := m
	for i, key := range path {
		if i == len(path)-1 {
			_, found := container[key]
			if found {
				delete(container, key)
			}
			return found
		}

		value, ok := container[key]
		if !ok {
			return false
		}
		container, ok = value.(map[string]interface{})
		if !ok {
			return false
		}
	}
	return false
}

// Update updates an existing resource with new values.
func (k *azureNativeProvider) Update(ctx context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Update(%s)", k.name, urn)
	logging.V(9).Infof("%s executing", label)
	inputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), KeepUnknowns: true, SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	resourceKey := string(urn.Type())
	res, ok, err := k.LookupResource(resourceKey)
	if err != nil {
		return nil, errors.Errorf("Decoding resource spec %s", resourceKey)
	}
	if !ok {
		return nil, errors.Errorf("Resource type '%s' not found", resourceKey)
	}

	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	if req.GetPreview() {
		// The preview outputs are inputs + a limited list of outputs that are universally immutable.
		// We know that their values won't change, so it's safe to propagate the values to dependent
		// resources during the preview.
		outputs := k.converter.PreviewOutputs(inputs, res.Response)

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

	restoreDefaultInputsForRemovedProperties(inputs, res, oldState)

	id, bodyParams, queryParams, err := k.prepareAzureRESTInputs(
		res.Path,
		res.PutParameters,
		res.RequiredContainers,
		inputs.Mappable(),
		map[string]interface{}{
			"subscriptionId": k.subscriptionID,
			"api-version":    res.APIVersion,
		},
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := azureContext(ctx, req.Timeout)
	defer cancel()

	var outputs map[string]interface{}
	customRes, isCustom := k.customResources[res.Path]
	switch {
	case isCustom && customRes.Update != nil:
		outputs, err = customRes.Update(ctx, inputs)
		if err != nil {
			return nil, azureError(err)
		}
	default:
		err := k.maintainSubResourcePropertiesIfNotSet(ctx, &res, id, bodyParams)
		if err != nil {
			return nil, fmt.Errorf("failed maintaining unset sub-resource properties: %w", err)
		}
		response, updated, err := k.azureCreateOrUpdate(ctx, id, bodyParams, queryParams, res.UpdateMethod, res.PutAsyncStyle)
		if err != nil {
			if updated {
				// Resource was partially updated but the operation failed to complete.
				// Try reading its state by ID and return a partial error if succeeded.
				checkpoint, getErr := k.currentResourceStateCheckpoint(ctx, id, res, inputs)
				if getErr != nil {
					return nil, azureError(errors.Wrapf(err, "resource updated but read failed %s", getErr))
				}
				return nil, partialError(id, err, checkpoint, req.GetNews())
			}
			return nil, azureError(err)
		}

		// Map the raw response to the shape of outputs that the SDKs expect.
		outputs = k.converter.ResponseBodyToSdkOutputs(res.Response, response)
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

func restoreDefaultInputsForRemovedProperties(inputs resource.PropertyMap, res resources.AzureAPIResource, oldState resource.PropertyMap) error {
	for property, defaultValue := range res.DefaultProperties {
		key := resource.PropertyKey(property)
		if !inputs.HasValue(key) && oldState.HasValue(key) {
			logging.V(5).Infof("setting removed property %q to default value", property)
			inputs[key] = resource.NewPropertyValue(defaultValue)
		}
	}

	return nil
}

// Delete tears down an existing resource with the given ID. If it fails, the resource is assumed
// to still exist.
func (k *azureNativeProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Delete(%s)", k.name, urn)
	logging.V(9).Infof("%s executing", label)
	id := req.GetId()
	resourceKey := string(urn.Type())
	res, ok, err := k.LookupResource(resourceKey)
	if err != nil {
		return nil, errors.Errorf("Decoding resource spec %s", resourceKey)
	}
	if !ok {
		return nil, errors.Errorf("Resource type '%s' not found", resourceKey)
	}

	customRes, isCustom := k.customResources[res.Path]

	ctx, cancel := azureContext(ctx, req.Timeout)
	defer cancel()

	switch {
	case isCustom && customRes.Delete != nil:
		state, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
			Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: false, SkipNulls: true, KeepSecrets: false,
		})
		if err != nil {
			return nil, err
		}
		inputs, err := parseCheckpointObject(req.OldInputs, state, label)
		if err != nil {
			return nil, err
		}
		if inputs == nil {
			return nil, errors.Wrapf(err, "resource %s inputs are empty", label)
		}
		// Our hand-crafted implementation of DELETE operation.
		err = customRes.Delete(ctx, id, inputs)
		if err != nil {
			return nil, azureError(err)
		}
	case res.Singleton:
		// Singleton resources can't be deleted (or created), set them to the default state.
		for _, param := range res.PutParameters {
			if defaults.SkipDeleteOperation(res.Path) {
				continue
			}
			if param.Location == "body" {
				requestBody := k.converter.SdkInputsToRequestBody(param.Body.Properties, res.DefaultBody, id)

				queryParams := map[string]interface{}{"api-version": res.APIVersion}
				_, _, err := k.azureCreateOrUpdate(ctx, id, requestBody, queryParams, res.UpdateMethod, res.PutAsyncStyle)
				if err != nil {
					return nil, azureError(err)
				}
			}
		}
	default:
		err := k.AzureDelete(ctx, id, res.APIVersion, res.DeleteAsyncStyle, nil)
		if err != nil {
			return nil, azureError(err)
		}
	}
	return &pbempty.Empty{}, nil
}

// Construct creates a new component resource.
func (k *azureNativeProvider) Construct(_ context.Context, _ *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

// Call dynamically executes a method in the provider associated with a component resource.
func (k *azureNativeProvider) Call(_ context.Context, _ *rpc.CallRequest) (*rpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (k *azureNativeProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: k.version,
	}, nil
}

// GetMapping eturns plugin specific mapping data for the given provider name. This implementation
// is a stub to satisfy the new convert.Mapper interface from pulumi/pulumi#11510.
func (k *azureNativeProvider) GetMapping(context.Context, *rpc.GetMappingRequest) (*rpc.GetMappingResponse, error) {
	return &rpc.GetMappingResponse{}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (k *azureNativeProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	// TODO
	return &pbempty.Empty{}, nil
}

type propertyPath struct {
	path         []string
	propertyName string
}

func (k *azureNativeProvider) maintainSubResourcePropertiesIfNotSet(ctx context.Context, res *resources.AzureAPIResource, id string, bodyParams map[string]interface{}) error {
	// Identify the properties we need to read
	missingProperties := k.findUnsetPropertiesToMaintain(res, bodyParams)

	if len(missingProperties) == 0 {
		// Everything's already specified explicitly by the user, no need to do read.
		return nil
	}

	// Read the current resource state.
	state, err := k.azureGet(ctx, id+res.ReadPath, res.APIVersion)
	if err != nil {
		return fmt.Errorf("reading cloud state: %w", err)
	}

	writtenProperties := writePropertiesToBody(missingProperties, bodyParams, state)
	for writtenProperty, writtenValue := range writtenProperties {
		logging.V(9).Infof("Maintaining remote value for property: %s.%s = %v", id, writtenProperty, writtenValue)
	}

	return nil
}

func writePropertiesToBody(missingProperties []propertyPath, bodyParams map[string]interface{}, remoteState map[string]interface{}) map[string]interface{} {
	writtenProperties := map[string]interface{}{}
	for _, prop := range missingProperties {
		currentBodyContainer := bodyParams
		currentStateContainer := remoteState
		for _, containerName := range prop.path {
			innerBodyContainer, bodyOk := currentBodyContainer[containerName]
			innerStateContainer, stateOk := currentStateContainer[containerName]
			// If the container doesn't exist in either body or state, create it and continue iterating.
			// But if it doesn't exist in either, there is no point in continuing.
			if !bodyOk && !stateOk {
				break
			}
			if !bodyOk {
				innerBodyContainer = map[string]interface{}{}
				currentBodyContainer[containerName] = innerBodyContainer
			}
			if !stateOk {
				innerStateContainer = map[string]interface{}{}
				currentStateContainer[containerName] = innerStateContainer
			}
			innerBodyObj, innerBodyIsObject := innerBodyContainer.(map[string]interface{})
			innerStateObj, innerStateIsObject := innerStateContainer.(map[string]interface{})
			if !innerBodyIsObject || !innerStateIsObject { // we've reached a leaf node (primitive type)
				break
			}
			currentBodyContainer = innerBodyObj
			currentStateContainer = innerStateObj
		}

		stateValue, ok := currentStateContainer[prop.propertyName]
		if ok {
			currentBodyContainer[prop.propertyName] = stateValue
			writtenProperties[fmt.Sprintf("%s.%s", strings.Join(prop.path, "."), prop.propertyName)] = stateValue
		}
	}
	return writtenProperties
}

func (k *azureNativeProvider) findUnsetPropertiesToMaintain(res *resources.AzureAPIResource, bodyParams map[string]interface{}) []propertyPath {
	missingProperties := []propertyPath{}
	for _, path := range res.PathsToSubResourcePropertiesToMaintain(true /* includeContainers i.e. API-shape */, k.lookupType) {
		curBody := bodyParams
		for i, pathEl := range path {
			v, ok := curBody[pathEl]
			if !ok {
				missingProperties = append(missingProperties, propertyPath{
					path:         path,
					propertyName: pathEl,
				})
				break
			}

			// At the end of the path we don't need to go deeper via references and map lookups.
			if i == len(path)-1 {
				break
			}

			curBody, ok = v.(map[string]interface{})
			if !ok {
				missingProperties = append(missingProperties, propertyPath{
					path:         path,
					propertyName: pathEl,
				})
				break
			}
		}
	}

	return missingProperties
}

func (k *azureNativeProvider) azureCreateOrUpdate(
	ctx context.Context,
	id string,
	bodyProps map[string]interface{},
	queryParameters map[string]interface{},
	updateMethod string,
	asyncStyle string) (map[string]interface{}, bool, error) {

	var op autorest.PrepareDecorator
	switch updateMethod {
	case "PATCH":
		op = autorest.AsPatch()
	default:
		op = autorest.AsPut()
	}

	decorators := []autorest.PrepareDecorator{
		autorest.AsContentType("application/json; charset=utf-8"),
		op,
		autorest.WithBaseURL(k.environment.ResourceManagerEndpoint),
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
		k.client,
		prepReq,
		azure.DoRetryWithRegistration(k.client),
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
		err = future.WaitForCompletionRef(ctx, k.client)
		if err != nil {
			return nil, created, err
		}
		resp, err = future.GetResult(k.client)
		if err != nil {
			return nil, created, err
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
		return nil, created, err
	}
	return outputs, true, nil
}

func (k *azureNativeProvider) AzureDelete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	for k, v := range queryParams {
		queryParameters[k] = v
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

	err = autorest.Respond(
		resp,
		k.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing(),
	)
	if err != nil {
		return err
	}
	return nil
}

// azureCanCreate asserts that a resource with a given ID and API version can be created
// or returns an error otherwise.
func (k *azureNativeProvider) azureCanCreate(ctx context.Context, id string, res *resources.AzureAPIResource) error {
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
		autorest.WithPath(id+res.ReadPath),
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

	switch {
	case http.StatusOK == resp.StatusCode && res.Singleton:
		// Singleton resources always exist, so OK is expected.
		return nil
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
		if res.ReadMethod == "HEAD" {
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

func (k *azureNativeProvider) azureHead(ctx context.Context, id string, apiVersion string) error {
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

func (k *azureNativeProvider) azureGet(ctx context.Context, id string,
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

func (k *azureNativeProvider) azurePost(
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

func (k *azureNativeProvider) prepareAzureRESTInputs(path string, parameters []resources.AzureAPIParameter, requiredContainers gen.RequiredContainers, methodInputs,
	clientInputs map[string]interface{}) (string, map[string]interface{}, map[string]interface{}, error) {
	// Schema-driven mapping of inputs into Autorest id/body/query
	params := map[string]map[string]interface{}{
		"query": {
			"api-version": clientInputs["api-version"],
		},
		"path": {},
	}

	// Build maps of path and query parameters.
	for _, param := range parameters {
		if param.Location == "body" {
			continue
		}
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

	// Calculate resource ID based on path parameter values.
	id := path
	for key, value := range params["path"] {
		encodedVal := autorest.Encode("path", value.(string))
		id = strings.Replace(id, "{"+key+"}", encodedVal, -1)
	}

	// Build the body JSON.
	var body map[string]interface{}
	for _, param := range parameters {
		if param.Location != "body" {
			continue
		}
		body = k.converter.SdkInputsToRequestBody(param.Body.Properties, methodInputs, id)
		break
	}

	// Ensure all required containers are created.
	for _, containers := range requiredContainers {
		currentContainer := body
		for _, containerName := range containers {
			innerContainer, ok := currentContainer[containerName]
			if !ok {
				innerContainer = map[string]interface{}{}
				currentContainer[containerName] = innerContainer
			}
			currentContainer, ok = innerContainer.(map[string]interface{})
			if !ok {
				break
			}
		}
	}

	return id, body, params["query"], nil
}

func (k *azureNativeProvider) setLoggingContext(ctx context.Context) {
	log.SetOutput(NewTerraformLogRedirector(ctx, k.host))
}

func (k *azureNativeProvider) getConfig(configName, envName string) string {
	if val, ok := k.config[configName]; ok {
		return val
	}

	return os.Getenv(envName)
}

// getUserAgent returns a User Agent string for the current provider configuration.
func (k *azureNativeProvider) getUserAgent() string {
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
	userAgent = strings.TrimSpace(fmt.Sprintf("%s pulumi-azure-native/%s",
		autorest.UserAgent(), version.Version))

	// append the CloudShell version to the user agent if it exists
	if azureAgent := os.Getenv("AZURE_HTTP_USER_AGENT"); azureAgent != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, azureAgent)
	}

	// Append partner ID, if it's defined.
	if partnerID != "" {
		userAgent = fmt.Sprintf("%s pid-%s", userAgent, partnerID)
	}

	logging.V(9).Infof("AzureNative User Agent: %s", userAgent)
	return
}

// azureContext returns a new context with a timeout - either the one explicitly specified,
// or the default one (2 hours).
func azureContext(ctx context.Context, timeoutSeconds float64) (context.Context, context.CancelFunc) {
	d := 120 * time.Minute
	if timeoutSeconds > 0 {
		d = time.Duration(timeoutSeconds) * time.Second
	}
	return context.WithTimeout(ctx, d)
}

// azureError catches common errors and substitutes them with more user-friendly ones.
func azureError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		return errors.New("operation timed out")
	}
	if requestError, ok := err.(azure.RequestError); ok {
		if requestError.DetailedError.Message != "" {
			return fmt.Errorf("%w. %s", err, requestError.DetailedError.Message)
		}
	}
	return err
}

// partialError creates an error for resources that did not complete an operation in progress.
// The last known state of the object is included in the error so that it can be checkpointed.
func partialError(id string, err error, state *structpb.Struct, inputs *structpb.Struct) error {
	detail := rpc.ErrorResourceInitFailed{
		Id:         id,
		Properties: state,
		Reasons:    []string{err.Error()},
		Inputs:     inputs,
	}
	return rpcerror.WithDetails(rpcerror.New(codes.Unknown, err.Error()), &detail)
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

// checkpointObject puts inputs in the `__inputs` field of the state.
func checkpointObject(inputs resource.PropertyMap, outputs map[string]interface{}) resource.PropertyMap {
	object := resource.NewPropertyMapFromMap(outputs)
	object["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	return object
}

// parseCheckpointObject returns the old inputs either from the RPC request or that are saved in the `__inputs` field of the state.
func parseCheckpointObject(reqOldInputs *structpb.Struct, obj resource.PropertyMap, label string) (resource.PropertyMap, error) {
	// Prefer the old inputs from the RPC request if available because these are much easier for a user to edit in the state file.
	// In the next major version of the provider, we'll remove the `__inputs` field from the state and require that the user is on
	// a version of the CLI that includes the old inputs in the RPC request. See https://github.com/pulumi/pulumi-azure-native/issues/2686
	if reqOldInputs != nil {
		oldInputs, err := plugin.UnmarshalProperties(reqOldInputs, plugin.MarshalOptions{
			Label: fmt.Sprintf("%s.oldInputs", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
		})
		if err != nil {
			return nil, err
		}
		return oldInputs, nil
	}
	if inputs, ok := obj["__inputs"]; ok {
		return inputs.ObjectValue(), nil
	}

	return nil, nil
}

func (k *azureNativeProvider) autorestEnvToHamiltonEnv() environments.Environment {
	switch k.environment.Name {
	case azure.USGovernmentCloud.Name:
		return environments.USGovernmentL4
	case azure.ChinaCloud.Name:
		return environments.China
	default:
		return environments.Global
	}
}
