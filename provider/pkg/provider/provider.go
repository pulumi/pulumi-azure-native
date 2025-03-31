// Copyright 2016-2024, Pulumi Corporation.

package provider

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/blang/semver"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/manicminer/hamilton/environments"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/deepcopy"
	"github.com/segmentio/encoding/json"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/go-autorest/autorest"
	azureEnv "github.com/Azure/go-autorest/autorest/azure"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/defaults"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources/customresources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
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

	azureClient    azure.AzureClient
	host           *provider.HostClient
	name           string
	version        string
	subscriptionID string
	environment    azureEnv.Environment
	cloud          azcloud.Configuration
	config         map[string]string

	// also known as "metadata"
	resourceMap *resources.APIMetadata

	schemaBytes []byte

	converter        *convert.SdkShapeConverter
	customResources  map[string]*customresources.CustomResource
	rgLocationMap    map[string]string
	lookupType       resources.TypeLookupFunc
	skipReadOnUpdate bool

	context  context.Context
	shutdown context.CancelFunc
}

func makeProvider(host *provider.HostClient, name, version string, schemaBytes []byte,
	azureAPIResourcesBytes []byte) (rpc.ResourceProviderServer, error) {

	resourceMap, err := LoadMetadataPartial(azureAPIResourcesBytes)
	if err != nil {
		return nil, err
	}
	return makeProviderInternal(host, name, version, schemaBytes, resourceMap)
}

func makeProviderInternal(host *provider.HostClient, name, version string, schemaBytes []byte,
	resourceMap *resources.APIMetadata) (*azureNativeProvider, error) {

	converter := convert.NewSdkShapeConverterPartial(resourceMap.Types)

	ctx, shutdown := context.WithCancel(context.Background())

	// Return the new provider
	p := &azureNativeProvider{
		// Note: azureClient and subscriptionId will be initialized in Configure.
		host:          host,
		name:          name,
		version:       version,
		resourceMap:   resourceMap,
		config:        map[string]string{},
		schemaBytes:   schemaBytes,
		converter:     &converter,
		rgLocationMap: map[string]string{},
		context:       ctx,
		shutdown:      shutdown,
	}
	p.lookupType = p.lookupTypeDefault
	return p, nil
}

func (k *azureNativeProvider) getVersion() semver.Version {
	if k.version == "" {
		return version.GetVersion()
	}
	return semver.MustParse(k.version)
}

// LoadMetadataPartial partially deserializes the provided json byte array into an AzureAPIMetadata
// in memory
func LoadMetadataPartial(azureAPIResourcesBytes []byte) (*resources.APIMetadata, error) {
	var resourceMap resources.PartialAzureAPIMetadata
	if _, err := json.Parse(azureAPIResourcesBytes, &resourceMap, json.ZeroCopy); err != nil {
		return nil, errors.Wrap(err, "unmarshalling resource map")
	}
	apiMetadata := resources.APIMetadata{
		Types:     &resourceMap.Types,
		Resources: &resourceMap.Resources,
		Invokes:   &resourceMap.Invokes,
	}
	return &apiMetadata, nil
}

// LoadMetadata deserializes the provided json byte array into an AzureAPIMetadata in memory
func LoadMetadata(azureAPIResourcesBytes []byte) (*resources.AzureAPIMetadata, error) {
	var resourceMap resources.AzureAPIMetadata

	if _, err := json.Parse(azureAPIResourcesBytes, &resourceMap, json.ZeroCopy); err != nil {
		return nil, errors.Wrap(err, "unmarshalling resource map")
	}
	return &resourceMap, nil
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

// newCrudClient implements crud.ResourceCrudClientFactory
func (p *azureNativeProvider) newCrudClient(res *resources.AzureAPIResource) crud.ResourceCrudClient {
	return crud.NewResourceCrudClient(p.azureClient, p.lookupType, p.converter, p.subscriptionID, res)
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

	if k.getVersion().Major >= 3 && (!req.GetSendsOldInputs() || !req.GetSendsOldInputsToDelete()) {
		// https://github.com/pulumi/pulumi-azure-native/issues/2686
		return nil, errors.New("Azure Native provider requires Pulumi CLI v3.74.0 or later")
	}

	for key, val := range req.GetVariables() {
		k.config[strings.TrimPrefix(key, "azure-native:config:")] = val
	}

	k.setLoggingContext(ctx)

	authConfig, err := k.getAuthConfig()
	if err != nil {
		return nil, err
	}

	k.environment, err = authConfig.autorestEnvironment()
	if err != nil {
		return nil, err
	}

	k.cloud = authConfig.cloud()

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

	var credential azcore.TokenCredential
	if util.EnableAzcoreBackend() {
		credential, err = k.newTokenCredential()
		if err != nil {
			return nil, fmt.Errorf("creating Pulumi auth credential: %w", err)
		}
	} else {
		logging.V(9).Infof("Using legacy authentication")
		credential = azCoreTokenCredential{p: k}
	}

	k.azureClient, err = k.newAzureClient(resourceManagerAuth, credential, userAgent)
	if err != nil {
		return nil, fmt.Errorf("creating Azure client: %w", err)
	}

	// When the provider is parameterized, resources and types that custom resources are built on will probably not be available.
	if !k.isParameterized() {
		k.customResources, err = customresources.BuildCustomResources(&k.environment, k.azureClient, k.LookupResource, k.newCrudClient, k.subscriptionID,
			resourceManagerBearerAuth, resourceManagerAuth, keyVaultBearerAuth, userAgent, credential)
		if err != nil {
			return nil, fmt.Errorf("initializing custom resources: %w", err)
		}
	}

	k.skipReadOnUpdate = k.getConfig("skipReadOnUpdate", "ARM_SKIP_READ_ON_UPDATE") == "true"

	return &rpc.ConfigureResponse{
		SupportsPreview:                 true,
		SupportsAutonamingConfiguration: true,
	}, nil
}

func (k *azureNativeProvider) isParameterized() bool {
	return strings.HasPrefix(k.name, "azure-native"+parameterizedNameSeparator)
}

func (k *azureNativeProvider) newAzureClient(armAuth autorest.Authorizer, tokenCred azcore.TokenCredential, userAgent string) (azure.AzureClient, error) {
	if util.EnableAzcoreBackend() {
		logging.V(9).Infof("AzureClient: using azcore and azidentity")
		return azure.NewAzCoreClient(tokenCred, userAgent, k.cloud, nil)
	}
	logging.V(9).Infof("AzureClient: using autorest")
	return azure.NewAzureClient(k.environment, armAuth, userAgent), nil
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
		token, err := k.getClientToken(ctx, auth, args["endpoint"])
		if err != nil {
			return nil, err
		}
		outputs = map[string]interface{}{"token": token}
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
		body, err := crud.PrepareAzureRESTBody(id, parameters, nil, args.Mappable(), k.converter)
		if err != nil {
			if body == nil {
				return nil, fmt.Errorf("error preparing body for %s: %v", label, err)
			}
			// Ignore the error if we've got a body as it's just a warning
		}

		var response any
		if res.GetParameters != nil {
			response, err = k.azureClient.Get(ctx, id, res.APIVersion, nil)
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
		outputs = k.invokeResponseToOutputs(response, res)
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

func (k *azureNativeProvider) getClientToken(ctx context.Context, authConfig *authConfig, endpointArg resource.PropertyValue) (string, error) {
	endpoint := k.tokenEndpoint(endpointArg)

	if util.EnableAzcoreBackend() {
		cred, err := k.newTokenCredential()
		if err != nil {
			return "", err
		}
		t, err := cred.GetToken(ctx, tokenRequestOpts(endpoint))
		if err != nil {
			return "", err
		}
		return t.Token, nil
	}

	// legacy autorest/go-azure-helpers auth
	return k.getOAuthToken(ctx, authConfig, endpoint)
}

// Returns the Azure endpoint where tokens can be requested. If the argument is not null or empty,
// it will be used verbatim.
func (k *azureNativeProvider) tokenEndpoint(endpointArg resource.PropertyValue) string {
	if endpointArg.HasValue() && endpointArg.IsString() && endpointArg.StringValue() != "" {
		return endpointArg.StringValue()
	}
	return k.environment.ResourceManagerEndpoint
}

func tokenRequestOpts(endpoint string) policy.TokenRequestOptions {
	return policy.TokenRequestOptions{
		// "".default" is the well-defined scope for all resources accessible to the user or
		// application. Despite the URL, it doesn't apply only to OIDC.
		// https://learn.microsoft.com/en-us/entra/identity-platform/scopes-oidc#the-default-scope
		Scopes: []string{endpoint + "/.default"},
	}
}

func (k *azureNativeProvider) invokeResponseToOutputs(response any, res resources.AzureAPIInvoke) map[string]any {
	var outputs map[string]any
	if responseMap, ok := response.(map[string]any); ok {
		// Map the raw response to the shape of outputs that the SDKs expect.
		outputs = k.converter.ResponseBodyToSdkOutputs(res.Response, responseMap)
	} else {
		outputs = map[string]any{resources.SingleValueProperty: response}
	}

	if res.IsResourceGetter {
		// resource getters have an azureApiVersion output property.
		if k.getVersion().Major >= 3 {
			if res.APIVersion != "" {
				outputs["azureApiVersion"] = resource.NewStringProperty(res.APIVersion)
			}
		}
	}

	return outputs
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

	k.applyDefaults(ctx, req.Urn, req.RandomSeed, req.Autonaming, *res, olds, news)
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
	response, err := k.azureClient.Get(ctx, id, apiVersionResources, nil)
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

func (k *azureNativeProvider) getDefaultName(urn string, randomSeed []byte,
	autonaming *rpc.CheckRequest_AutonamingOptions, strategy resources.AutoNameKind, key resource.PropertyKey,
	olds resource.PropertyMap) (nameValue *resource.PropertyValue, randomlyNamed bool, ok bool) {
	if v, ok := olds[key]; ok {
		if vf, ok := olds[createBeforeDeleteFlag]; ok && vf.IsBool() {
			return &v, vf.BoolValue(), true
		}
		return &v, false, true
	}

	result := func(s string) *resource.PropertyValue {
		p := resource.NewStringProperty(s)
		return &p
	}

	if autonaming != nil {
		switch autonaming.Mode {
		case rpc.CheckRequest_AutonamingOptions_DISABLE:
			// Do not return any value. If the property is required, Check will fail
			// with a missing required property error.
			return nil, false, false
		case rpc.CheckRequest_AutonamingOptions_ENFORCE:
			return result(autonaming.ProposedName), true, true
		case rpc.CheckRequest_AutonamingOptions_PROPOSE:
			switch strategy {
			case resources.AutoNameRandom, resources.AutoNameCopy:
				return result(autonaming.ProposedName), true, true
			case resources.AutoNameUuid:
				// Do nothing for UUIDs in proposed mode, fall back to the normal logic below.
			}
		}
	}

	urnParts := strings.Split(urn, "::")
	name := urnParts[len(urnParts)-1]
	switch strategy {
	case resources.AutoNameRandom:
		// Resource name is URN name + random suffix.
		random, err := resource.NewUniqueName(randomSeed, name, 8, 0, nil)
		contract.AssertNoError(err)
		return result(random), true, true
	case resources.AutoNameUuid:
		// Resource name is a random UUID. We need to do a similar trick as NewUniqueName so that this is
		// deterministic by randomSeed. We simply ask NewUniqueName for a 32 byte random hex name.
		hexID, err := resource.NewUniqueName(randomSeed, "", 32, 0, nil)
		contract.AssertNoError(err)
		return result(uuid.MustParse(hexID).String()), true, true
	case resources.AutoNameCopy:
		// Resource name is just a copy of the URN name.
		return result(name), false, true
	default:
		panic(fmt.Sprintf("unknown auto-naming strategy %q", strategy))
	}
}

// Apply default values (e.g., location) to user's inputs.
func (k *azureNativeProvider) applyDefaults(ctx context.Context, urn string, randomSeed []byte,
	autonaming *rpc.CheckRequest_AutonamingOptions, res resources.AzureAPIResource, olds, news resource.PropertyMap) {
	for _, par := range res.PutParameters {
		sdkName := par.Name
		if par.Value != nil && par.Value.SdkName != "" {
			sdkName = par.Value.SdkName
		}

		// Auto-naming.
		key := resource.PropertyKey(sdkName)
		if !news.HasValue(key) && par.Value != nil && par.Value.AutoName != "" {
			name, randomlyNamed, ok := k.getDefaultName(urn, randomSeed, autonaming, par.Value.AutoName, key, olds)
			if ok {
				news[key] = *name
				if randomlyNamed {
					news[createBeforeDeleteFlag] = resource.NewBoolProperty(true)
				}
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

	// Set the default value of azureApiVersion to the APIVersion of the inputs.
	// Note that some Azure resource types do not have an APIVersion (e.g. storage:Blob).
	if k.getVersion().Major >= 3 {
		if !news.HasValue("azureApiVersion") && res.APIVersion != "" {
			news["azureApiVersion"] = resource.NewStringProperty(res.APIVersion)
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
		typ, ok, err := k.lookupType(prop.Ref)
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

	// v2-to-v3 migration: apiVersion might be available in the old state and we use it to suppress spurious diffs
	// that would otherwise be produced by mere input-input diffing.
	migratedApiVersion := false
	if k.getVersion().Major >= 3 {
		if _, ok := oldInputs["azureApiVersion"]; !ok {
			if v, ok := oldState["azureApiVersion"]; ok {
				oldInputs["azureApiVersion"] = v
				migratedApiVersion = true
			}
		}
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

	detailedDiff := diff(k.lookupType, *res, oldInputs, newResInputs)
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

	if migratedApiVersion {
		if v, ok := detailedDiff["azureApiVersion"]; ok {
			// in this case, the diff is between the old state and the new inputs.
			v.InputDiff = false
		}
	}

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
	// Use the global context to handle provider shutdown.
	ctx = k.context
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

	customRes, isCustom := k.customResources[res.Path]

	crudClient := crud.NewResourceCrudClient(k.azureClient, k.lookupType, k.converter, k.subscriptionID, res)

	var id string
	var queryParams map[string]any
	if isCustom && customRes.GetIdAndQuery != nil {
		id, queryParams, err = customRes.GetIdAndQuery(ctx, inputs, crudClient)
	} else {
		id, queryParams, err = crudClient.PrepareAzureRESTIdAndQuery(inputs)
	}
	if err != nil {
		return nil, err
	}

	ctx, cancel := azureContext(ctx, req.Timeout)
	defer cancel()

	var outputs map[string]interface{}
	switch {
	case isCustom && customRes.Create != nil:
		outputs, err = customCreate(ctx, inputs, id, crudClient, customRes)
		if err != nil {
			return nil, azure.AzureError(err)
		}
	default:
		id, outputs, err = k.defaultCreate(ctx, req, inputs, id, queryParams, crudClient,
			reader(customRes, crudClient, nil /* previousInputs, none for Create */))
		if err != nil {
			return nil, err
		}
	}

	// Store both outputs and inputs into the state.
	obj := checkpointObject(res, inputs, outputs)

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

func customCreate(ctx context.Context, inputs resource.PropertyMap, id string, crudClient crud.ResourceCrudClient,
	customRes *customresources.CustomResource,
) (map[string]any, error) {
	// First check if the resource already exists - we want to try our best to avoid updating instead of creating here.
	var exists bool
	var err error

	if customRes.CanCreate != nil {
		err = customRes.CanCreate(ctx, id)
		exists = err != nil
	} else if customRes.Read != nil {
		_, exists, err = customRes.Read(ctx, id, inputs)
	} else {
		err = crudClient.CanCreate(ctx, id)
		exists = err != nil
	}
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("cannot create already existing resource %q", id)
	}

	// Create the custom resource and retrieve its outputs, which already match the SDK shape.
	outputs, err := customRes.Create(ctx, id, inputs)
	if err != nil {
		return nil, azure.AzureError(err)
	}
	return outputs, nil
}

func (k *azureNativeProvider) defaultCreate(ctx context.Context, req *rpc.CreateRequest, inputs resource.PropertyMap, id string,
	queryParams map[string]any, crudClient crud.ResourceCrudClient, reader readFunc) (string, map[string]any, error) {
	bodyParams, err := crudClient.PrepareAzureRESTBody(id, inputs)
	if err != nil {
		bodyError := fmt.Errorf("error preparing body for %s: %v", id, err)
		if bodyParams == nil {
			return id, nil, bodyError
		}
		// Ignore the error if we've got a body as it's just a warning
	}

	// First check if the resource already exists - we want to try our best to avoid updating instead of creating here
	// (though it's technically impossible since the only operation supported is an upsert).
	err = crudClient.CanCreate(ctx, id)
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

	if readResponse := k.readAfterWrite(ctx, id, req.GetUrn(), "Create", inputs, reader); readResponse != nil {
		response = readResponse
	}

	return id, response, nil
}

// readFunc is a function that reads the state of a resource.
type readFunc func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error)

// reader returns a function that reads the state of a resource either from a custom Read if it is
// implemented, or via the standard ResourceCrudClient otherwise. In the latter case, the response
// will be converted to SDK shape. Custom Reads are expected to do this on their own.
//
// The `previousInputs` argument is only used for resources where the API version is user-provided.
func reader(customRes *customresources.CustomResource, crudClient crud.ResourceCrudClient, previousInputs resource.PropertyMap) readFunc {
	if customRes != nil && customRes.Read != nil {
		return func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
			response, _, err := customRes.Read(ctx, id, inputs)
			return response, err
		}
	}

	return func(ctx context.Context, id string, _ resource.PropertyMap) (map[string]any, error) {
		overrideApiVersion := ""
		if crudClient.ApiVersionIsUserInput() {
			if apiVersion, ok := previousInputs["apiVersion"]; ok {
				overrideApiVersion = apiVersion.StringValue()
			}
		}

		response, err := crudClient.Read(ctx, id, overrideApiVersion)
		if err == nil {
			// Map the raw response to the shape of outputs that the SDKs expect.
			response = crudClient.ResponseBodyToSdkOutputs(response)
		}
		return response, err
	}
}

// Read the state of the resource from its Read endpoint. While we already have resource state from
// the PUT reponse, it may not match precisely the shape of Read reponses that we also use for Refresh
// operations. That discrepancy may cause noisy diffs that are hard for users to reconcile.
func (p *azureNativeProvider) readAfterWrite(ctx context.Context, id, urn, opForLogging string,
	inputs resource.PropertyMap, reader readFunc,
) map[string]any {
	if p.skipReadOnUpdate {
		return nil
	}

	response, err := reader(ctx, id, inputs)

	if err != nil {
		// Since this read operation was introduced in a minor version of the provider, we choose to ignore
		// any errors here to avoid user-facing regressions. If no warnings are reported, we should be able
		// to promote this situation to an error.
		p.host.Log(ctx, diag.Warning, resource.URN(urn), fmt.Sprintf(`Failed to read resource after %s. Please report this issue.
Verbose logs contain more information, see https://www.pulumi.com/docs/support/troubleshooting/#verbose-logging.`, opForLogging))
		logging.V(9).Infof("failed to read resource %q after %s: %s", id, opForLogging, err.Error())
		return nil
	}

	return response
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

// Depending on the major version, because we turned on SendOldInputs in v3, the previous inputs can be in state under
// `__inputs`, or available as `req.GetInputs()`/`req.GetOldInputs()`.
func getPreviousInputs(state resource.PropertyMap, reqInputs *structpb.Struct, label string) (resource.PropertyMap, error) {
	if reqInputs != nil {
		return plugin.UnmarshalProperties(reqInputs, plugin.MarshalOptions{
			Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: false, SkipNulls: true, KeepSecrets: false,
		})
	}
	if __inputs, ok := state["__inputs"]; ok {
		return __inputs.ObjectValue(), nil
	}
	return nil, nil
}

// Read the current live state associated with a resource.
func (k *azureNativeProvider) Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	// Use the global context to handle provider shutdown.
	ctx = k.context
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

	previousInputs, err := getPreviousInputs(oldState, req.GetInputs(), label)
	if err != nil {
		return nil, err
	}

	res, err := k.lookupResourceFromURN(urn)
	if err != nil {
		return nil, err
	}

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
		url := id + res.ReadPath
		err = k.azureClient.Head(ctx, url, getApiVersion(res, previousInputs))
		response = oldState.Mappable()
		outputs = crudClient.ResponseBodyToSdkOutputs(response)
	default:
		response, err = crudClient.Read(ctx, id, getApiVersion(res, previousInputs))
		outputs = crudClient.ResponseBodyToSdkOutputs(response)
	}
	if err != nil {
		if azure.IsNotFound(err) {
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

	plainOldState := oldState.Mappable()
	if oldState.HasValue(customresources.OriginalStateKey) {
		outputsWithoutIgnores[customresources.OriginalStateKey] = plainOldState[customresources.OriginalStateKey]
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
		removeDefaults(*res, plainOldState)
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
	obj := checkpointObject(res, inputs, outputsWithoutIgnores)

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

// removeDefaults removes the resource's default values from the given state, modifying the map in place.
func removeDefaults(res resources.AzureAPIResource, plainOldState map[string]any) {
	previousInputsRaw, ok := plainOldState["__inputs"]
	if !ok {
		return
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
	// Use the global context to handle provider shutdown.
	ctx = k.context
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
		previousInputs, err := getPreviousInputs(oldState, req.GetOldInputs(), label)
		if err != nil {
			return nil, err
		}
		outputs, err = k.defaultUpdate(ctx, req, inputs, id, queryParams, crudClient, reader(customRes, crudClient, previousInputs))
		if err != nil {
			return nil, err
		}
	}

	// Store both outputs and inputs into the state.
	obj := checkpointObject(res, inputs, outputs)

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
	id string, queryParams map[string]any, crudClient crud.ResourceCrudClient, reader readFunc,
) (map[string]any, error) {

	bodyParams, err := crudClient.PrepareAzureRESTBody(id, inputs)
	if err != nil {
		bodyError := fmt.Errorf("error preparing body for %s: %v", id, err)
		if bodyParams == nil {
			return nil, bodyError
		}
		// Ignore the error if we've got a body as it's just a warning
	}

	ctx, cancel := azureContext(ctx, req.Timeout)
	defer cancel()

	err = crudClient.MaintainSubResourcePropertiesIfNotSet(ctx, id, bodyParams)
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

	if readResponse := k.readAfterWrite(ctx, id, req.GetUrn(), "Update", inputs, reader); readResponse != nil {
		response = readResponse
	}

	return response, nil
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

func isSingleton(res *resources.AzureAPIResource) bool {
	return res.Singleton
}

// Delete tears down an existing resource with the given ID. If it fails, the resource is assumed
// to still exist.
func (k *azureNativeProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error) {
	// Use the global context to handle provider shutdown.
	ctx = k.context
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

	state, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: false, SkipNulls: true, KeepSecrets: false,
	})

	switch {
	case isCustom && customRes.Delete != nil:
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
		err = customRes.Delete(ctx, id, inputs, state)
		if err != nil {
			return nil, azure.AzureError(err)
		}
	case isSingleton(res):
		// Singleton resources can't be deleted (or created), set them to the default state.
		for _, param := range res.PutParameters {
			if defaults.SkipDeleteOperation(res.Path, res.APIVersion) {
				continue
			}
			if param.Location == "body" {
				requestBody, err := k.converter.SdkInputsToRequestBody(param.Body.Properties, res.DefaultBody, id)
				if err != nil {
					// Log conversion errors but continue with the deletion.
					k.host.Log(ctx, diag.Warning, urn, fmt.Sprintf("error converting inputs to request body: %v", err))
				}

				queryParams := map[string]interface{}{"api-version": res.APIVersion}

				// Submit the `PUT` or `PATCH` against the ARM endpoint
				op := k.azureClient.Put
				if res.UpdateMethod == "PATCH" {
					op = k.azureClient.Patch
				}
				_, _, err = op(ctx, id, requestBody, queryParams, res.PutAsyncStyle)
				if err != nil {
					return nil, azure.AzureError(err)
				}
			}
		}
	default:
		previousInputs, err := getPreviousInputs(state, req.GetOldInputs(), label)
		if err != nil {
			return nil, err
		}

		apiVersion := getApiVersion(res, previousInputs)
		err = k.azureClient.Delete(ctx, id, apiVersion, res.DeleteAsyncStyle, nil)
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
	k.shutdown() // Cancel the global provider context.
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

// checkpointObject produces the checkpointed state for the given inputs and outputs.
// In v2, we stored the inputs in an `__inputs` field of the state; removed in v3.
func checkpointObject(res *resources.AzureAPIResource, inputs resource.PropertyMap, outputs map[string]interface{}) resource.PropertyMap {
	return checkpointObjectVersioned(res, inputs, outputs, version.GetVersion())
}

func checkpointObjectVersioned(res *resources.AzureAPIResource, inputs resource.PropertyMap, outputs map[string]interface{}, version semver.Version) resource.PropertyMap {
	object := resource.NewPropertyMapFromMap(outputs)
	if version.Major < 3 {
		object["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	}
	// If this is a custom resource that needs the resource's original state, preserve it as a secret.
	if object.HasValue(customresources.OriginalStateKey) {
		object[customresources.OriginalStateKey] = resource.MakeSecret(object[customresources.OriginalStateKey])
	}

	// emit the actual APIVersion as an output property.
	if res.APIVersion != "" {
		object["azureApiVersion"] = resource.NewStringProperty(res.APIVersion)
	}
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

// getApiVersion usually returns the API version set on `resource` (from metadata). For resources where the user
// specifies the API version as input, it returns the value from `inputs`.
func getApiVersion(resource *resources.AzureAPIResource, inputs resource.PropertyMap) string {
	if resource.ApiVersionIsUserInput {
		if apiVersion, ok := inputs["apiVersion"]; ok {
			return apiVersion.StringValue()
		}
	}
	return resource.APIVersion
}
