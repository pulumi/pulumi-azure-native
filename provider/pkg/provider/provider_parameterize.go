// Copyright 2025, Pulumi Corporation.

package provider

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// parameterizeArgs is the data that is embedded in the SDK when the provider is parameterized. It will be sent to the
// provider on subsequent invocations of Parameterize to be deserialized.
type parameterizeArgs struct {
	Module  string `json:"module"`
	Version string `json:"version"`
}

func serializeParameterizeArgs(args *parameterizeArgs) ([]byte, error) {
	serialized, err := json.Marshal(args)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to serialize parameterize args: %v", err)
	}
	return []byte(base64.StdEncoding.EncodeToString(serialized)), nil
}

func deserializeParameterizeArgs(in []byte) (*parameterizeArgs, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to deserialize parameterize args: %v", err)
	}
	var args parameterizeArgs
	if err := json.Unmarshal(decoded, &args); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to deserialize parameterize args: %v", err)
	}
	return &args, nil
}

// parseApiVersion parses an Azure API version from the given string. Since the input comes from users (via `package
// add`), it tries to be lenient and accept several formats. The returned result, if successful, is an "SDL version" of
// the form v20200101.
func parseApiVersion(version string) (string, error) {
	v := strings.TrimSpace(version)
	v = strings.TrimLeft(v, "vV")

	isApiVersion, err := regexp.MatchString(`^20\d{2}-[01]\d-[0-3]\d(-preview|-privatepreview)?$`, v)
	if err != nil {
		return "", status.Errorf(codes.InvalidArgument, "unexpected error parsing version %s: %v", version, err)
	}
	if isApiVersion {
		apiVersion := openapi.ApiVersion(v)
		return string(apiVersion.ToSdkVersion()), nil
	}

	isDigits, err := regexp.MatchString(`^20\d{6}(preview|privatepreview)?$`, v)
	if err != nil {
		return "", status.Errorf(codes.InvalidArgument, "unexpected error parsing version %s: %v", version, err)
	}
	if isDigits {
		return "v" + v, nil
	}

	return "", status.Errorf(codes.InvalidArgument, "invalid API version: %s", version)
}

// "When generating an SDK (e.g. using a pulumi package add command), we need to boot up a provider and parameterize it
// using only information from the command-line invocation. In this case, the parameter is a string array representing
// the command-line arguments (args)."
// https://pulumi-developer-docs.readthedocs.io/latest/docs/architecture/providers/parameterized.html#parameterized-providers
func getParameterizeArgs(req *rpc.ParameterizeRequest) (*parameterizeArgs, error) {
	switch {
	// initial parameterization
	case req.GetArgs() != nil:
		args := req.GetArgs().Args

		// "aad" "v20221201"
		if len(args) != 2 {
			return nil, status.Errorf(codes.InvalidArgument, "expected 2 arguments (module and API version), got %d", len(args))
		}
		targetApiVersion, err := parseApiVersion(args[1])
		if err != nil {
			return nil, err
		}
		return &parameterizeArgs{
			Module:  args[0],
			Version: targetApiVersion,
		}, nil

	// provider has already been parameterized and the arguments were serialized into the SDK
	case req.GetValue() != nil:
		return deserializeParameterizeArgs(req.GetValue().Value)

	default:
		return nil, status.Errorf(codes.InvalidArgument, "cannot handle ParameterizeRequest with parameters of type %T", req.Parameters)
	}
}

func (p *azureNativeProvider) Parameterize(ctx context.Context, req *rpc.ParameterizeRequest) (*rpc.ParameterizeResponse, error) {
	logging.V(9).Info("Parameterizing Azure Native provider")

	args, err := getParameterizeArgs(req)
	if err != nil {
		return nil, err
	}

	logging.V(9).Infof("Creating parameterized Azure Native for %s %s", args.Module, args.Version)

	var schema pschema.PackageSpec
	err = json.Unmarshal(p.schemaBytes, &schema)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal schema: %v", err)
	}

	newSchema, newMetadata, err := createSchema(p, schema, args.Module, args.Version)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create parameterized schema: %v", err)
	}
	newPackageName := newSchema.Name

	serializedArgs, err := serializeParameterizeArgs(args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to serialize args %v: %v", args, err)
	}

	newSchema.Parameterization = &pschema.ParameterizationSpec{
		BaseProvider: pschema.BaseProviderSpec{
			Name:    p.name,
			Version: newSchema.Version,
		},
		Parameter: serializedArgs,
	}

	s, err := json.Marshal(newSchema)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal schema: %v", err)
	}
	s = updateRefs(s, newPackageName, args.Module, args.Version)

	newMetadata, err = updateMetadataRefs(newMetadata, newPackageName, args.Module, args.Version)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update metadata $refs: %v", err)
	}

	p.schemaBytes = s
	p.resourceMap = newMetadata
	p.name = newPackageName

	if _, found := os.LookupEnv("PULUMI_DEBUG_PARAMETERIZE"); found {
		tmpPath := filepath.Join(os.TempDir(), newPackageName+".json")
		err = os.WriteFile(tmpPath, s, 0644)
		if err != nil {
			logging.Infof("failed to write PULUMI_DEBUG_PARAMETERIZE schema to %s: %v", tmpPath, err)
		} else {
			logging.Infof("wrote PULUMI_DEBUG_PARAMETERIZE schema to %s", tmpPath)
		}
	}

	resp := &rpc.ParameterizeResponse{
		Name:    newPackageName,
		Version: newSchema.Version,
	}
	return resp, nil
}

const parameterizedNameSeparator = "_"

func generateNewPackageName(unparameterizedPackageName, targetModule, targetApiVersion string) string {
	return strings.Join([]string{unparameterizedPackageName, targetModule, targetApiVersion}, parameterizedNameSeparator)
}

// updateRefs updates all `$ref` pointers in the serialized schema to use the new package name, e.g., from `"$ref":
// "#/types/azure-native:..."` to `"$ref": "#/types/azure-native_resources_20240101:..."`.
func updateRefs(serialized []byte, newPackageName, module, apiVersion string) []byte {
	oldRefPrefix := fmt.Sprintf(`"$ref":"#/types/azure-native:%s/%s`, module, apiVersion)
	newRefPrefix := fmt.Sprintf(`"$ref": "#/types/%s:%s`, newPackageName, module)
	return bytes.ReplaceAll(serialized, []byte(oldRefPrefix), []byte(newRefPrefix))
}

// updateMetadataRefs updates all `$ref` pointers in the metadata to use the new package name.
// This implementation uses a JSON round-trip to update the `$ref`'s via a global string-replacement. Not elegant, but effective.
func updateMetadataRefs(metadata *resources.APIMetadata, newPackageName, module, apiVersion string) (*resources.APIMetadata, error) {
	m, err := json.Marshal(metadata)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal metadata: %v", err)
	}
	updated := updateRefs(m, newPackageName, module, apiVersion)
	newMetadata := &resources.APIMetadata{}
	err = json.Unmarshal(updated, newMetadata)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal metadata: %v", err)
	}
	return newMetadata, nil
}

func getAvailableApiVersions(schema pschema.PackageSpec, targetModule string) []string {
	versions := map[string]struct{}{}
	for resourceName := range schema.Resources {
		moduleName, version, _, err := resources.ParseToken(resourceName)
		if err != nil {
			continue
		}
		if moduleName == targetModule && version != "" {
			versions[version] = struct{}{}
		}
	}
	return util.SortedKeys(versions)
}

// createParameterizedSchemaForApiVersion creates a new schema for the given module and API version by picking the
// required resources, types, and functions from the providers existing schema and metadata. This assumes that the given
// provider has a full schema with all API versions.
//
// There are other ways of obtaining such a schema, for instance, generating it directly from the Azure spec. This way
// was the most pragmatic approach since we already have the full schema in the provider.
//
// All names change because their "module" part needs to match the new schema, and the new package name can't be just
// "azure-native" which already exists.
//
// To separate concerns, the `Parameterization` of the new schema is NOT populated yet, the caller is responsible for
// doing that.
func createSchema(p *azureNativeProvider, schema pschema.PackageSpec, targetModule, targetApiVersion string) (*pschema.PackageSpec, *resources.APIMetadata, error) {
	newPackageName := generateNewPackageName(schema.Name, targetModule, targetApiVersion)

	newSchema := pschema.PackageSpec{
		Name:        newPackageName,
		Version:     schema.Version,
		Description: fmt.Sprintf("A specialized Pulumi Azure Native package for %s %s", targetModule, targetApiVersion),
		DisplayName: schema.DisplayName,
		License:     schema.License,
		Keywords:    schema.Keywords,
		Homepage:    schema.Homepage,
		Publisher:   schema.Publisher,
		Repository:  schema.Repository,
		Config:      schema.Config,
		Provider:    schema.Provider,
		Language:    schema.Language,
		Types:       map[string]pschema.ComplexTypeSpec{},
		Resources:   map[string]pschema.ResourceSpec{},
		Functions:   map[string]pschema.FunctionSpec{},
	}

	makeToken := func(name string) string {
		return fmt.Sprintf("%s:%s:%s", newPackageName, targetModule, name)
	}

	typeNames, err := filterTokens(schema.Types, targetModule, targetApiVersion)
	if err != nil {
		return nil, nil, status.Errorf(codes.InvalidArgument, "failed to parse type token: %v", err)
	}
	resourceNames, err := filterTokens(schema.Resources, targetModule, targetApiVersion)
	if err != nil {
		return nil, nil, status.Errorf(codes.InvalidArgument, "failed to parse resource token: %v", err)
	}
	functionNames, err := filterTokens(schema.Functions, targetModule, targetApiVersion)
	if err != nil {
		return nil, nil, status.Errorf(codes.InvalidArgument, "failed to parse function token: %v", err)
	}
	logging.V(9).Infof("Creating parameterized Azure Native for %s %s: found %d types, %d resources, %d functions",
		targetModule, targetApiVersion, len(typeNames), len(resourceNames), len(functionNames))

	if len(resourceNames) == 0 {
		availableVersions := getAvailableApiVersions(schema, targetModule)
		if len(availableVersions) == 0 {
			return nil, nil, status.Errorf(codes.InvalidArgument, "module %s not found. Some modules were renamed in v3 of the provider; "+
				"please see https://www.pulumi.com/registry/packages/azure-native/from-v2-to-v3/#new-module-structure-and-naming-aligned-closer-to-azure-sdks.",
				targetModule)
		}
		return nil, nil, status.Errorf(codes.InvalidArgument, "no resources found for module %s at API version %s. Available API versions: %s",
			targetModule, targetApiVersion, strings.Join(availableVersions, ", "))
	}

	metadataTypes := map[string]resources.AzureAPIType{}
	metadataResources := map[string]resources.AzureAPIResource{}
	metadataInvokes := map[string]resources.AzureAPIInvoke{}

	for typeTok, typeName := range typeNames {
		newTok := makeToken(typeName)
		newSchema.Types[newTok] = schema.Types[typeTok]

		apiType, ok, err := p.lookupType(typeTok)
		if err != nil {
			return nil, nil, status.Errorf(codes.InvalidArgument, "failed to get type %s: %v", typeName, err)
		}
		if !ok {
			logging.Warningf("type %s not found in metadata", typeName)
		} else {
			metadataTypes[newTok] = *apiType
		}
	}

	for resourceTok, resourceName := range resourceNames {
		newTok := makeToken(resourceName)
		schemaResource := schema.Resources[resourceTok]
		newSchema.Resources[newTok] = schemaResource

		resource, ok, err := p.resourceMap.Resources.Get(resourceTok)
		if err != nil {
			return nil, nil, status.Errorf(codes.InvalidArgument, "failed to get type %s: %v", resourceName, err)
		}
		if !ok {
			logging.Warningf("resource %s not found in metadata", resourceName)
		} else {
			metadataResources[newTok] = resource
		}
	}

	for functionTok, functionName := range functionNames {
		newTok := makeToken(functionName)
		newSchema.Functions[newTok] = schema.Functions[functionTok]

		invoke, ok, err := p.resourceMap.Invokes.Get(functionTok)
		if err != nil {
			return nil, nil, status.Errorf(codes.InvalidArgument, "failed to get type %s: %v", functionName, err)
		}
		if !ok {
			logging.Warningf("function %s not found in metadata", functionName)
		} else {
			metadataInvokes[newTok] = invoke
		}
	}

	metadata := &resources.APIMetadata{
		Types:     resources.GoMap[resources.AzureAPIType](metadataTypes),
		Resources: resources.GoMap[resources.AzureAPIResource](metadataResources),
		Invokes:   resources.GoMap[resources.AzureAPIInvoke](metadataInvokes),
	}

	return &newSchema, metadata, nil
}

// filterTokens returns a map of tokens that match the target module and API version.
func filterTokens[T any](tokens map[string]T, targetModule, targetApiVersion string) (map[string]string, error) {
	typeNames := map[string]string{}
	for token := range tokens {
		moduleName, version, name, err := resources.ParseToken(token)
		if err != nil {
			return nil, err
		}
		if !strings.EqualFold(moduleName, targetModule) || version != targetApiVersion {
			continue
		}
		typeNames[token] = name
	}
	return typeNames, nil
}
