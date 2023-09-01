package versioning

import (
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/providerlist"
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
	OnlyExplicitVersions    bool
	ExampleLanguages        []string
	Version                 string
}

type BuildSchemaReports struct {
	PathChangesResult
	AllResourcesByVersion         ProvidersVersionResources
	AllResourceVersionsByResource ProviderResourceVersions
	Active                        providerlist.ProviderPathVersionsJson
	Pending                       openapi.ProviderVersionList
	CurationViolations            []CurationViolation
}

func (r BuildSchemaReports) WriteTo(outputDir string) error {
	return gen.EmitFiles(outputDir, gen.FileMap{
		"pathChanges.json":                   r.PathChangesResult,
		"allResourcesByVersion.json":         r.AllResourcesByVersion,
		"allResourceVersionsByResource.json": r.AllResourceVersionsByResource,
		"active.json":                        r.Active,
		"pending.json":                       r.Pending,
		"curationViolations.json":            r.CurationViolations,
	})
}

type BuildSchemaResult struct {
	PackageSpec schema.PackageSpec
	Metadata    resources.AzureAPIMetadata
	Version     VersionMetadata
	Reports     BuildSchemaReports
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

	var versionMetadata VersionMetadata
	if args.OnlyExplicitVersions {
		versionMetadata = VersionMetadata{}
	} else {
		versionMetadata, err = LoadVersionMetadata(args.RootDir, providers, int(majorVersion))
		if err != nil {
			return nil, err
		}
	}

	if !args.OnlyExplicitVersions {
		providers = openapi.ApplyProvidersTransformations(providers, versionMetadata.Lock, versionMetadata.PreviousLock, nil, nil)
	}

	pathChanges := findPathChanges(providers, versionMetadata.Lock, versionMetadata.PreviousLock, versionMetadata.Config)

	if args.ExcludeExplicitVersions {
		providers = openapi.SingleVersion(providers)
	}

	generationResult, err := gen.PulumiSchema(providers, versionMetadata)
	if err != nil {
		return nil, err
	}

	pkgSpec := generationResult.Schema
	metadata := generationResult.Metadata
	examples := generationResult.Examples

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
		Reports: BuildSchemaReports{
			PathChangesResult:             pathChanges,
			AllResourcesByVersion:         versionMetadata.AllResourcesByVersion,
			AllResourceVersionsByResource: versionMetadata.AllResourceVersionsByResource,
			Active:                        versionMetadata.Active,
			Pending:                       versionMetadata.Pending,
			CurationViolations:            versionMetadata.CurationViolations,
		},
	}, nil
}

type PathChange struct {
	CurrentPath  string
	PreviousPath string
	ResourceName string
}

type MissingExpectedResourceVersion struct {
	ProviderName string
	ResourceName string
	Version      string
	IsPrevious   bool
}

type PathChangesResult struct {
	Changes                        []PathChange
	MissingPreviousDefaultVersions []MissingExpectedResourceVersion
}

func findPathChanges(providers openapi.AzureProviders,
	defaultVersion openapi.DefaultVersionLock,
	previousVersion openapi.DefaultVersionLock,
	curations Curations) PathChangesResult {

	result := []PathChange{}
	missingPreviousDefaultVersions := []MissingExpectedResourceVersion{}

	sortedProviderNames := []string{}
	for providerName := range defaultVersion {
		sortedProviderNames = append(sortedProviderNames, providerName)
	}
	sort.Strings(sortedProviderNames)

	for _, providerName := range sortedProviderNames {
		resources := defaultVersion[providerName]
		previousResources, ok := previousVersion[providerName]
		if !ok {
			continue
		}
		providerVersions := providers[providerName]

		sortedResourceNames := []string{}
		for resourceName := range resources {
			sortedResourceNames = append(sortedResourceNames, resourceName)
		}
		sort.Strings(sortedResourceNames)

		for _, resourceName := range sortedResourceNames {
			version := resources[resourceName]
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
				missingPreviousDefaultVersions = append(missingPreviousDefaultVersions, MissingExpectedResourceVersion{
					ProviderName: providerName,
					ResourceName: resourceName,
					Version:      prevVersion,
					IsPrevious:   false,
				})
				continue
			}

			prevSpec, ok := prev.Resources[resourceName]
			if !ok {
				prevSpec, ok = prev.Invokes[resourceName]
			}
			if !ok {
				missingPreviousDefaultVersions = append(missingPreviousDefaultVersions, MissingExpectedResourceVersion{
					ProviderName: providerName,
					ResourceName: resourceName,
					Version:      prevVersion,
					IsPrevious:   true,
				})
				continue
			}

			path := paths.NormalizePath(spec.Path)
			prevPath := paths.NormalizePath(prevSpec.Path)

			if path != prevPath {
				excluded, err := curations.IsExcluded(providerName, resourceName, prevVersion)
				if !excluded && err == nil {
					result = append(result, PathChange{
						CurrentPath:  path,
						PreviousPath: prevPath,
						ResourceName: resourceName,
					})
				}
			}
		}
	}
	return PathChangesResult{
		Changes:                        result,
		MissingPreviousDefaultVersions: missingPreviousDefaultVersions,
	}
}
