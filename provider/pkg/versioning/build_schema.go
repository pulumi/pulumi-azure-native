package versioning

import (
	"path"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

type ReadSpecsArgs struct {
	// Defaults to ./azure-rest-api-specs
	SpecsDir string
	// e.g. Compute
	NamespaceFilter string
	// e.g. "2019-09-01" or "2019*"
	VersionsFilter string
}

type BuildSchemaArgs struct {
	Specs                   ReadSpecsArgs
	RootDir                 string
	ExcludeExplicitVersions bool
	ExampleLanguages        []string
	Version2                bool
	Version                 string
}

type BuildSchemaResult struct {
	PackageSpec schema.PackageSpec
	Metadata    resources.AzureAPIMetadata
	Version     VersionMetadata
}

func BuildSchema(args BuildSchemaArgs) (*BuildSchemaResult, error) {
	specsDir := args.Specs.SpecsDir
	if specsDir == "" {
		specsDir = path.Join(args.RootDir, "azure-rest-api-specs")
	}
	providers, err := openapi.ReadAzureProviders(specsDir, args.Specs.NamespaceFilter, args.Specs.VersionsFilter)
	if err != nil {
		return nil, err
	}

	versionSources, err := ReadVersionSources(args.RootDir)
	if err != nil {
		return nil, err
	}

	versionMetadata, err := CalculateVersionMetadata(versionSources, providers)
	if err != nil {
		return nil, err
	}

	var removed openapi.ProviderVersionList
	if args.Version2 {
		removed, err = openapi.ReadProviderVersionList(path.Join(args.RootDir, "versions", "v2-removed.json"))
	} else {
		removed, err = openapi.ReadProviderVersionList(path.Join(args.RootDir, "versions", "v1-removed.json"))
	}
	if err != nil {
		return nil, err
	}

	if args.Version2 {
		providers = openapi.ApplyProvidersTransformations(providers, versionMetadata.V2Lock, versionMetadata.V1Lock, nil, removed)
	} else {
		providers = openapi.ApplyProvidersTransformations(providers, versionMetadata.V1Lock, nil, versionMetadata.Deprecated, removed)
	}

	if args.Version2 {
		providers = openapi.RemoveResources(providers, openapi.RemovableResources(versionMetadata.V2ResourcesToRemove))
	}

	if args.ExcludeExplicitVersions {
		providers = openapi.SingleVersion(providers)
	}

	pkgSpec, metadata, examples, err := gen.PulumiSchema(providers)
	if err != nil {
		return nil, err
	}

	if len(args.ExampleLanguages) > 0 {
		// Ensure the spec is stamped with a version - Go gen needs it.
		pkgSpec.Version = args.Version
		err = gen.Examples(pkgSpec, metadata, examples, args.ExampleLanguages)
		if err != nil {
			return nil, err
		}
		// Remove the version again so it doesn't get written to the schema file.
		pkgSpec.Version = ""
	}

	return &BuildSchemaResult{
		PackageSpec: *pkgSpec,
		Metadata:    *metadata,
		Version:     versionMetadata,
	}, nil
}
