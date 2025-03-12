// Copyright 2025, Pulumi Corporation.

package provider

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// runtimeParameterize is called when the provider is parameterized at runtime. It doesn't do anything but echo back the
// given name and version.
//
// TODO, tkappler I believe we should store the new metadata in parameters? "When interacting with a provider as part of
// program execution, the parameter is embedded in the SDK, so as to free the program author from having to know whether
// a provider is parameterized or not. In this case, the parameter is a provider-specific bytestring (value). This is
// intended to allow a provider to store arbitrary data that may be more efficient or practical at program execution
// time, after SDK generation has taken place. This value is base-64-encoded when embedded in the SDK."
// https://pulumi-developer-docs.readthedocs.io/latest/docs/architecture/providers/parameterized.html#parameterized-providers
func (p *azureNativeProvider) runtimeParameterize(req *rpc.ParameterizeRequest) (*rpc.ParameterizeResponse, error) {
	switch pt := req.Parameters.(type) {
	case *rpc.ParameterizeRequest_Value:
		// Extract the base64 metadata and populate it in the provider
		metadata, err := deserializeMetadata(pt.Value.GetValue())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to deserialize metadata: %v", err)
		}
		p.parameterizedMetadata = metadata

		return &rpc.ParameterizeResponse{
			Name:    pt.Value.GetName(),
			Version: pt.Value.Version,
		}, nil
	default:
		return nil, fmt.Errorf("cannot handle ParameterizeRequest with parameters of type %T", pt)
	}
}

func (p *azureNativeProvider) Parameterize(ctx context.Context, req *rpc.ParameterizeRequest) (*rpc.ParameterizeResponse, error) {
	logging.V(9).Info("Parameterizing Azure Native provider")

	// "When generating an SDK (e.g. using a pulumi package add command), we need to boot up a provider and parameterize
	// it using only information from the command-line invocation. In this case, the parameter is a string array
	// representing the command-line arguments (args)."
	// https://pulumi-developer-docs.readthedocs.io/latest/docs/architecture/providers/parameterized.html#parameterized-providers
	reqArgs := req.GetArgs()
	if reqArgs == nil {
		return p.runtimeParameterize(req)
	}
	args := reqArgs.Args

	// "aad" "v20221201"
	if len(args) != 2 {
		return nil, status.Errorf(codes.InvalidArgument, "expected 2 arguments (module and API version), got %d", len(args))
	}
	targetModule := args[0]
	targetApiVersion := args[1] // TODO be more accepting of version format and convert
	logging.V(9).Infof("Creating parameterized Azure Native for %s %s", targetModule, targetApiVersion)

	var schema pschema.PackageSpec
	err := json.Unmarshal(p.schemaBytes, &schema)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal schema: %v", err)
	}

	newSchema, err := createParameterizedSchemaForApiVersion(p, schema, targetModule, targetApiVersion)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create parameterized schema: %v", err)
	}
	newPackageName := newSchema.Name

	s, err := json.Marshal(newSchema)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal schema: %v", err)
	}

	s = bytes.ReplaceAll(s, []byte(`"$ref":"#/types/azure-native`), []byte(`"$ref": "#/types/`+newPackageName))

	p.schemaBytes = s

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

// createParameterizedSchemaForApiVersion creates a new schema for the given module and API version by picking the
// required resources, types, and functions from the providers existing schema and metadata. This assumes that the given
// provider has a full schema with all API versions.
//
// There are other ways of obtaining such a schema, for instance, generating it directly from the Azure spec. This way
// was the most pragmatic approach since we already have the full schema in the provider.
//
// All names change because their "module" part needs to match the new schema, and the new package name can't be just
// "azure-native" which already exists.
func createParameterizedSchemaForApiVersion(p *azureNativeProvider, schema pschema.PackageSpec, targetModule, targetApiVersion string) (*pschema.PackageSpec, error) {
	newPackageName := fmt.Sprintf("%s_%s_%s", schema.Name, targetModule, targetApiVersion) // "." is not allowed

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
		return fmt.Sprintf("%s:%s/%s:%s", newPackageName, targetModule, targetApiVersion, name)
	}

	typeNames, err := filterTokens(schema.Types, targetModule, targetApiVersion)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse token: %v", err)
	}
	resourceNames, err := filterTokens(schema.Resources, targetModule, targetApiVersion)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse token: %v", err)
	}
	functionNames, err := filterTokens(schema.Functions, targetModule, targetApiVersion)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse token: %v", err)
	}
	logging.V(9).Infof("Creating parameterized Azure Native for %s %s: found %d types, %d resources, %d functions",
		targetModule, targetApiVersion, len(typeNames), len(resourceNames), len(functionNames))

	metadata := &resources.AzureAPIMetadata{
		Types:     map[string]resources.AzureAPIType{},
		Resources: map[string]resources.AzureAPIResource{},
		Invokes:   map[string]resources.AzureAPIInvoke{},
	}

	for typeTok, typeName := range typeNames {
		newTok := makeToken(typeName)
		newSchema.Types[newTok] = schema.Types[typeTok]

		apiType, ok, err := p.lookupType(typeTok)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to get type %s: %v", typeName, err)
		}
		if !ok {
			logging.Warningf("type %s not found in metadata", typeName)
		} else {
			metadata.Types[newTok] = *apiType
		}
	}

	for resourceTok, resourceName := range resourceNames {
		newTok := makeToken(resourceName)
		newSchema.Resources[newTok] = schema.Resources[resourceTok]

		resource, ok, err := p.resourceMap.Resources.Get(resourceTok)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to get type %s: %v", resourceName, err)
		}
		if !ok {
			logging.Warningf("resource %s not found in metadata", resourceName)
		} else {
			metadata.Resources[newTok] = resource
		}
	}

	for functionTok, functionName := range functionNames {
		newTok := makeToken(functionName)
		newSchema.Functions[newTok] = schema.Functions[functionTok]

		invoke, ok, err := p.resourceMap.Invokes.Get(functionTok)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to get type %s: %v", functionName, err)
		}
		if !ok {
			logging.Warningf("function %s not found in metadata", functionName)
		} else {
			metadata.Invokes[newTok] = invoke
		}
	}

	metadataBytes, err := serializeMetadata(metadata)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to serialize metadata: %v", err)
	}
	newSchema.Parameterization = &pschema.ParameterizationSpec{
		BaseProvider: pschema.BaseProviderSpec{
			Name:    p.name,
			Version: newSchema.Version,
		},
		Parameter: metadataBytes,
	}

	return &newSchema, nil
}

func filterTokens[T any](tokens map[string]T, targetModule, targetApiVersion string) (map[string]string, error) {
	typeNames := map[string]string{}
	for token := range tokens {
		moduleName, version, name, err := resources.ParseToken(token)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to parse token: %v", err)
		}
		if !strings.EqualFold(moduleName, targetModule) || version != targetApiVersion {
			continue
		}
		typeNames[token] = name
	}
	return typeNames, nil
}

func serializeMetadata(metadata *resources.AzureAPIMetadata) ([]byte, error) {
	jsonBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(jsonBytes); err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(buf.Len()))
	base64.StdEncoding.Encode(dst, buf.Bytes())
	return dst, nil
}

func deserializeMetadata(metadata []byte) (*resources.AzureAPIMetadata, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(metadata))
	if err != nil {
		return nil, err
	}

	zr, err := gzip.NewReader(bytes.NewReader(decoded))
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(zr); err != nil {
		return nil, err
	}

	var result resources.AzureAPIMetadata
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
