// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"context"
	"fmt"
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

	azureEnv "github.com/Azure/go-autorest/autorest/azure"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/arm2pulumi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/defaults"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources/customresources"
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
	PulumiPartnerID        = "a90539d8-a7a6-5826-95c4-1fbef22d4b22"
	apiVersionResources    = "2020-10-01"
	createBeforeDeleteFlag = "__createBeforeDelete"
)

type azureNativeProvider struct {
	rpc.UnimplementedResourceProviderServer

	azureClient      azure.AzureClient
	host             *provider.HostClient
	name             string
	version          string
	subscriptionID   string
	environment      azureEnv.Environment
	resourceMap      *resources.PartialAzureAPIMetadata
	config           map[string]string
	schemaBytes      []byte
	metadataBytes    []byte
	fullPkgSpec      *schema.PackageSpec
	fullResourceMap  *resources.AzureAPIMetadata
	converter        *convert.SdkShapeConverter
	customResources  map[string]*customresources.CustomResource
	rgLocationMap    map[string]string
	lookupType       resources.TypeLookupFunc
	skipReadOnUpdate bool
}

func makeProvider(host *provider.HostClient, name, version string, schemaBytes []byte,
	azureAPIResourcesBytes []byte) (rpc.ResourceProviderServer, error) {

	resourceMap, err := LoadMetadataPartial(azureAPIResourcesBytes)
	if err != nil {
		return nil, err
	}

	converter := convert.NewSdkShapeConverterPartial(resourceMap.Types)

	// Return the new provider
	p := &azureNativeProvider{
		// This client will be regnerated with correct environment and authorizer in Configure.
		azureClient:   azure.NewAzureClient(azureEnv.PublicCloud, nil, azure.BuildUserAgent(PulumiPartnerID)),
		host:          host,
		name:          name,
		version:       version,
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

func (k *azureNativeProvider) lookupResourceFromURN(urn resource.URN) (*resources.AzureAPIResource, error) {
	resourceKey := string(urn.Type())
	res, ok, err := k.LookupResource(resourceKey)
	if err != nil {
		return nil, errors.Errorf("Decoding resource spec %s", resourceKey)
	}
	if !ok {
		return nil, errors.Errorf("Resource type %s not found", resourceKey)
	}
	return &res, nil
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
	env, err := azureEnv.EnvironmentFromName(envName)
	if err != nil {
		env, err = azureEnv.EnvironmentFromName(fmt.Sprintf("AZURE%sCLOUD", envName))
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

	userAgent := k.getUserAgent()

	k.azureClient = azure.NewAzureClient(env, resourceManagerAuth, userAgent)

	azCoreTokenCredential := azCoreTokenCredential{p: k}
	k.customResources, err = customresources.BuildCustomResources(&env, k.azureClient, k.LookupResource, k.subscriptionID,
		resourceManagerBearerAuth, resourceManagerAuth, keyVaultBearerAuth, userAgent, azCoreTokenCredential)
	if err != nil {
		return nil, fmt.Errorf("initializing custom resources: %w", err)
	}

	k.skipReadOnUpdate = k.getConfig("skipReadOnUpdate", "ARM_SKIP_READ_ON_UPDATE") == "true"

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
		id, query, err := crud.PrepareAzureRESTIdAndQuery(
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
		body := crud.PrepareAzureRESTBody(id, parameters, nil, args.Mappable(), k.converter)

		var response map[string]interface{}
		if res.GetParameters != nil {
			response, err = k.azureClient.Get(ctx, id, res.APIVersion)
		} else if res.PostParameters != nil {
			if body == nil {
				body = map[string]interface{}{}
			}
			response, err = k.azureClient.Post(ctx, id, body, query)
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

	res, err := k.lookupResourceFromURN(urn)
	if err != nil {
		return nil, err
	}

	k.applyDefaults(ctx, req.Urn, req.RandomSeed, *res, olds, news)
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
	response, err := k.azureClient.Get(ctx, id, apiVersionResources)
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
	res, err := k.lookupResourceFromURN(urn)
	if err != nil {
		return nil, err
	}

	detailedDiff := diff(k.lookupTypeDefault, *res, oldInputs, newResInputs)
	if detailedDiff == nil {
		return &rpc.DiffResponse{
			Changes:             rpc.DiffResponse_DIFF_NONE,
			Replaces:            []string{},
			Stables:             []string{},
			DeleteBeforeReplace: false,
		}, nil
	}

	// Based on the detailed diff above, calculate the list of changes and replacements.
	changes, replaces := calculateChangesAndReplacements(detailedDiff, oldInputs, newResInputs, oldState, *res)

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

	res, err := k.lookupResourceFromURN(urn)
	if err != nil {
		return nil, err
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

	crudClient := crud.NewResourceCrudClient(k.azureClient, k.lookupType, k.converter, k.subscriptionID, res)

	id, queryParams, err := crudClient.PrepareAzureRESTIdAndQuery(inputs)
	if err != nil {
		return nil, err
	}

	ctx, cancel := azureContext(ctx, req.Timeout)
	defer cancel()

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

		// Create the custom resource and retrieve its outputs, which already match the SDK shape.
		outputs, err = customRes.Create(ctx, id, inputs)
		if err != nil {
			return nil, azure.AzureError(err)
		}
	default:
		id, outputs, err = k.defaultCreate(ctx, req, inputs, id, queryParams, crudClient)
		if err != nil {
			return nil, err
		}
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

func (k *azureNativeProvider) defaultCreate(ctx context.Context, req *rpc.CreateRequest, inputs resource.PropertyMap, id string,
	queryParams map[string]any, crudClient crud.ResourceCrudClient) (string, map[string]any, error) {
	bodyParams := crudClient.PrepareAzureRESTBody(id, inputs)

	// First check if the resource already exists - we want to try our best to avoid updating instead of creating here
	// (though it's technically impossible since the only operation supported is an upsert).
	err := crudClient.CanCreate(ctx, id)
	if err != nil {
		return id, nil, err
	}

	crudClient.SetUnsetSubresourcePropertiesToDefaults(bodyParams, bodyParams, true)

	response, created, err := crudClient.CreateOrUpdate(ctx, id, bodyParams, queryParams)
	if err != nil {
		if created {
			return id, nil, crudClient.HandleErrorWithCheckpoint(ctx, err, id, inputs, req.GetProperties())
		}
		return id, nil, azure.AzureError(err)
	}

	// Read the canonical ID from the response.
	if azureId, ok := response["id"].(string); ok {
		id = azureId
	}

	// Read the state of the resource from its Read endpoint. While we already have resource state from
	// the PUT reponse, it may not match precisely the shape of Read reponses that we also use for Refresh
	// operations. That discrepancy may cause noisy diffs that are hard for users to reconcile.
	if !k.skipReadOnUpdate {
		readResponse, err := crudClient.Read(ctx, id)
		if err != nil {
			// Since this read operation was introduced in a minor version of the provider, we choose to ignore
			// any errors here to avoid user-facing regressions. If no warnings are reported, we should be able
			// to promote this situation to an error.
			k.host.Log(ctx, diag.Warning, resource.URN(req.GetUrn()), `Failed to read resource after Create. Please report this issue.
	Verbose logs contain more information, see https://www.pulumi.com/docs/support/troubleshooting/#verbose-logging.`)
			logging.V(9).Infof("failed read resource %q after creation: %s", id, err.Error())
		} else if readResponse != nil {
			response = readResponse
		}
	}

	// Map the raw response to the shape of outputs that the SDKs expect.
	outputs := crudClient.ResponseBodyToSdkOutputs(response)
	return id, outputs, nil
}

// Properties pointing to sub-resources that can be maintained as separate resources might not be
// present in the inputs because the user wants to manage them as standalone resources. However,
// such a property might be required by Azure even if it's not annotated as such in the spec, e.g.,
// Key Vault's accessPolicies. Therefore, we set these properties to their default value here,
// an empty array. For more details, see section "Sub-resources" in CONTRIBUTING.md.
//
// During create, no sub-resources can exist yet so there's no danger of overwriting existing values.
//
// The `input` param is used to determine the unset sub-resource properties. They are then reset in
// the `output` parameter which is modified in-place.
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
func (k *azureNativeProvider) setUnsetSubresourcePropertiesToDefaults(res resources.AzureAPIResource,
	input, output map[string]interface{}, outputIsInApiShape bool) {
	unset := k.findUnsetPropertiesToMaintain(&res, input, outputIsInApiShape)

	for _, p := range unset {
		cur := output
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

type propertyPath struct {
	path         []string
	propertyName string
}

func (k *azureNativeProvider) findUnsetPropertiesToMaintain(res *resources.AzureAPIResource, bodyParams map[string]interface{}, returnApiShapePaths bool) []propertyPath {
	missingProperties := []propertyPath{}
	for _, path := range res.PathsToSubResourcePropertiesToMaintain(returnApiShapePaths, k.lookupType) {
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

// currentResourceStateCheckpoint reads the resource state by ID, converts it to outputs map, and
// produces a checkpoint with these outputs and given inputs.
func (k *azureNativeProvider) currentResourceStateCheckpoint(ctx context.Context, id string, res resources.AzureAPIResource, inputs resource.PropertyMap) (*structpb.Struct, error) {
	getResp, getErr := k.azureClient.Get(ctx, id, res.APIVersion)
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

	res, err := k.lookupResourceFromURN(urn)
	if err != nil {
		return nil, err
	}

	url := id + res.ReadPath
	crudClient := crud.NewResourceCrudClient(k.azureClient, k.lookupType, k.converter, k.subscriptionID, res)

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
		err = k.azureClient.Head(ctx, url, res.APIVersion)
		response = oldState.Mappable()
		outputs = crudClient.ResponseBodyToSdkOutputs(response)
	default:
		response, err = crudClient.Read(ctx, id)
		outputs = crudClient.ResponseBodyToSdkOutputs(response)
	}
	if err != nil {
		if reqErr, ok := err.(*azureEnv.RequestError); ok && reqErr.StatusCode == http.StatusNotFound {
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

	var outputsWithoutIgnores = outputs
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
		plainOldState := mappableOldState(*res, oldState)
		// 2. Project old outputs to their corresponding input shape (exclude read-only properties).
		oldInputProjection := k.converter.SdkOutputsToSdkInputs(res.PutParameters, plainOldState)
		// 3a. Remove sub-resource properties from new outputs which weren't set in the old inputs.
		// If the user didn't specify them inline originally, we don't want to push them into the inputs now.
		outputsWithoutIgnores = k.resetUnsetSubResourceProperties(ctx, urn, outputs, inputs, res)
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
	obj := checkpointObject(inputs, outputsWithoutIgnores)

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

// removeUnsetSubResourceProperties resets sub-resource properties in the outputs if they weren't set in the old inputs.
// If the user didn't specify them inline originally, we don't want to push them into the inputs now.
// For more details, see section "Sub-resources" in CONTRIBUTING.md.
func (k *azureNativeProvider) resetUnsetSubResourceProperties(ctx context.Context, urn resource.URN, sdkResponse map[string]any,
	oldInputs resource.PropertyMap, res *resources.AzureAPIResource) map[string]any {
	// Take a deep copy so we don't modify the original which is also used later for diffing.
	copy := deepcopy.Copy(sdkResponse)
	result, ok := copy.(map[string]interface{})
	if !ok {
		// This should never happen.
		k.host.Log(ctx, diag.Warning, urn, `Failed to remove unset sub-resource properties. Please report this issue.
Verbose logs contain more information, see https://www.pulumi.com/docs/support/troubleshooting/#verbose-logging.`)
		logging.V(9).Infof("failed to remove unset sub-resource properties: failed to cast copy value, expected map[string]interface{}, found %v", copy)
		return sdkResponse
	}

	k.setUnsetSubresourcePropertiesToDefaults(*res, oldInputs.Mappable(), result, false)

	return result
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

	res, err := k.lookupResourceFromURN(urn)
	if err != nil {
		return nil, err
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

	restoreDefaultInputsForRemovedProperties(inputs, *res, oldState)

	crudClient := crud.NewResourceCrudClient(k.azureClient, k.lookupType, k.converter, k.subscriptionID, res)

	id, queryParams, err := crudClient.PrepareAzureRESTIdAndQuery(inputs)
	if err != nil {
		return nil, err
	}

	var outputs map[string]interface{}
	customRes, isCustom := k.customResources[res.Path]
	if isCustom && customRes.Update != nil {
		outputs, err = customRes.Update(ctx, id, inputs, oldState)
		if err != nil {
			return nil, azure.AzureError(err)
		}
	} else {
		// Map the raw response to the shape of outputs that the SDKs expect.
		outputs, err = k.defaultUpdate(ctx, req, inputs, id, queryParams, crudClient)
		if err != nil {
			return nil, err
		}
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

func (k *azureNativeProvider) defaultUpdate(ctx context.Context, req *rpc.UpdateRequest, inputs resource.PropertyMap,
	id string, queryParams map[string]any, crudClient crud.ResourceCrudClient) (map[string]any, error) {

	bodyParams := crudClient.PrepareAzureRESTBody(id, inputs)

	ctx, cancel := azureContext(ctx, req.Timeout)
	defer cancel()

	err := crudClient.MaintainSubResourcePropertiesIfNotSet(ctx, id, bodyParams)
	if err != nil {
		return nil, fmt.Errorf("failed maintaining unset sub-resource properties: %w", err)
	}

	response, updated, err := crudClient.CreateOrUpdate(ctx, id, bodyParams, queryParams)
	if err != nil {
		if updated {
			return nil, crudClient.HandleErrorWithCheckpoint(ctx, err, id, inputs, req.GetNews())
		}
		return nil, azure.AzureError(err)
	}

	// Read the state of the resource from its Read endpoint. While we already have resource state from
	// the PUT reponse, it may not match precisely the shape of Read reponses that we also use for Refresh
	// operations. That discrepancy may cause noisy diffs that are hard for users to reconcile.
	if !k.skipReadOnUpdate {
		readResponse, err := crudClient.Read(ctx, id)
		if err != nil {
			// Since this read operation was introduced in a minor version of the provider, we choose to ignore
			// any errors here to avoid user-facing regressions. If no warnings are reported, we should be able
			// to promote this situation to an error.
			k.host.Log(ctx, diag.Warning, resource.URN(req.GetUrn()), `Failed to read resource after Update. Please report this issue.
	Verbose logs contain more information, see https://www.pulumi.com/docs/support/troubleshooting/#verbose-logging.`)
			logging.V(9).Infof("failed read resource %q after creation: %s", id, err.Error())
		} else if readResponse != nil {
			response = readResponse
		}
	}

	// Map the raw response to the shape of outputs that the SDKs expect.
	return crudClient.ResponseBodyToSdkOutputs(response), nil
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

	res, err := k.lookupResourceFromURN(urn)
	if err != nil {
		return nil, err
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
			return nil, azure.AzureError(err)
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

				// Submit the `PUT` or `PATCH` against the ARM endpoint
				op := k.azureClient.Put
				if res.UpdateMethod == "PATCH" {
					op = k.azureClient.Patch
				}
				_, _, err := op(ctx, id, requestBody, queryParams, res.PutAsyncStyle)
				if err != nil {
					return nil, azure.AzureError(err)
				}
			}
		}
	default:
		err := k.azureClient.Delete(ctx, id, res.APIVersion, res.DeleteAsyncStyle, nil)
		if err != nil {
			return nil, azure.AzureError(err)
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
	return azure.BuildUserAgent(partnerID)
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
	case azureEnv.USGovernmentCloud.Name:
		return environments.USGovernmentL4
	case azureEnv.ChinaCloud.Name:
		return environments.China
	default:
		return environments.Global
	}
}
