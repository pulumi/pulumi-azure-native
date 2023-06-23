package versioning

import (
	"fmt"
	"log"
	"path"
	"strconv"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
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

	majorVersion, err := strconv.ParseInt(strings.Split(args.Version, ".")[0], 10, 64)
	if err != nil {
		return nil, err
	}

	versionMetadata, err := LoadVersionMetadata(args.RootDir, providers, int(majorVersion))
	if err != nil {
		return nil, err
	}

	providers = openapi.ApplyProvidersTransformations(providers, versionMetadata.Lock, versionMetadata.PreviousLock, nil, nil)

	pathChanges := findPathChanges(providers, versionMetadata.Lock, versionMetadata.PreviousLock, versionMetadata.Config)
	printPathChanges(pathChanges)

	if args.ExcludeExplicitVersions {
		providers = openapi.SingleVersion(providers)
	}

	pkgSpec, metadata, examples, err := gen.PulumiSchema(providers, versionMetadata)
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

type pathChange struct {
	currentPath  string
	previousPath string
	resourceName string
}

func findPathChanges(providers openapi.AzureProviders,
	defaultVersion openapi.DefaultVersionLock,
	previousVersion openapi.DefaultVersionLock,
	curations Curations) []pathChange {

	result := []pathChange{}

	for providerName, resources := range defaultVersion {
		previousResources, ok := previousVersion[providerName]
		if !ok {
			continue
		}
		providerVersions := providers[providerName]

		for resourceName, version := range resources {
			previousVersion, ok := previousResources[resourceName]
			if !ok {
				continue
			}

			cur := providerVersions[openapi.ApiToSdkVersion(version)]
			prevVersion := openapi.ApiToSdkVersion(previousVersion)
			prev := providerVersions[prevVersion]

			spec, ok := cur.Resources[resourceName]
			if !ok {
				spec, ok = cur.Invokes[resourceName]
			}
			if !ok {
				log.Printf("Warning: could not find current default resource %s/%s in OpenAPI spec.\n", resourceName, version)
				continue
			}

			prevSpec, ok := prev.Resources[resourceName]
			if !ok {
				prevSpec, ok = prev.Invokes[resourceName]
			}
			if !ok {
				log.Printf("Warning: could not find previous default resource %s/%s in OpenAPI spec.\n", resourceName, version)
				continue
			}

			path := paths.NormalizePath(spec.Path)
			prevPath := paths.NormalizePath(prevSpec.Path)

			if path != prevPath {
				excluded, err := curations.IsExcluded(providerName, resourceName, prevVersion)
				if !excluded && err == nil {
					result = append(result, pathChange{
						currentPath:  path,
						previousPath: prevPath,
						resourceName: resourceName,
					})
				}
			}
		}
	}
	return result
}

func printPathChanges(changes []pathChange) {
	const fmtStr = "[V1->V2 path change] %s: %s...\n    ...%s\n    ...%s\n"

	for _, change := range changes {
		cur := change.currentPath
		prev := change.previousPath

		// Find the first index where the paths differ so we can print the common prefix only once.
		idx := 0
		for idx < len(cur) && idx < len(prev) && cur[idx] == prev[idx] {
			idx++
		}
		fmt.Printf(fmtStr, change.resourceName, prev[:idx], cur[idx:], prev[idx:])
	}
}
