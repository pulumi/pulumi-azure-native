package versioning

import (
	"path"
	"sort"

	"github.com/blang/semver"
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
	OnlyExplicitVersions    bool
	ExampleLanguages        []string
	Version                 string
}

type BuildSchemaReports struct {
	PathChangesResult
	// providerName -> resourceName -> set of paths, to record resources that have conflicting paths.
	PathConflicts                 map[openapi.ProviderName]map[openapi.ResourceName]map[string][]openapi.ApiVersion
	AllResourcesByVersion         ProvidersVersionResources
	AllResourceVersionsByResource ProviderResourceVersions
	Pending                       openapi.ProviderVersionList
	CurationViolations            []CurationViolation
	NamingDisambiguations         []resources.NameDisambiguation
	SkippedPOSTEndpoints          map[string]map[string]string
	ProviderNameErrors            []openapi.ProviderNameError
	ForceNewTypes                 []gen.ForceNewType
	TypeCaseConflicts             gen.CaseConflicts
	FlattenedPropertyConflicts    map[openapi.ProviderName]map[string]struct{}
	AllEndpoints                  map[openapi.ProviderName]map[openapi.ResourceName]map[string]*openapi.Endpoint
	InactiveDefaultVersions       map[openapi.ProviderName][]openapi.ApiVersion
}

func (r BuildSchemaReports) WriteTo(outputDir string) ([]string, error) {
	return gen.EmitFiles(outputDir, gen.FileMap{
		"allEndpoints.json":                  r.AllEndpoints,
		"allResourcesByVersion.json":         r.AllResourcesByVersion,
		"allResourceVersionsByResource.json": r.AllResourceVersionsByResource,
		"curationViolations.json":            r.CurationViolations,
		"flattenedPropertyConflicts.json":    r.FlattenedPropertyConflicts,
		"forceNewTypes.json":                 r.ForceNewTypes,
		"inactiveDefaultVersions.json":       r.InactiveDefaultVersions,
		"namingDisambiguations.json":         r.NamingDisambiguations,
		"pathChanges.json":                   r.PathChangesResult,
		"pathConflicts.json":                 r.PathConflicts,
		"pending.json":                       r.Pending,
		"providerNameErrors.json":            r.ProviderNameErrors,
		"skippedPOSTEndpoints.json":          r.SkippedPOSTEndpoints,
		"typeCaseConflicts.json":             r.TypeCaseConflicts,
	})
}

type BuildSchemaResult struct {
	PackageSpec schema.PackageSpec
	Metadata    resources.AzureAPIMetadata
	Version     VersionMetadata
	Reports     BuildSchemaReports
	Providers   openapi.AzureProviders
}

func BuildSchema(args BuildSchemaArgs) (*BuildSchemaResult, error) {
	specsDir := args.Specs.SpecsDir
	if specsDir == "" {
		specsDir = path.Join(args.RootDir, "azure-rest-api-specs")
	}
	providers, diagnostics, err := openapi.ReadAzureProviders(specsDir, args.Specs.NamespaceFilter, args.Specs.VersionsFilter)
	if err != nil {
		return nil, err
	}

	providerVersion, err := semver.Parse(args.Version)
	if err != nil {
		return nil, err
	}

	var versionMetadata VersionMetadata
	if args.OnlyExplicitVersions {
		versionMetadata = VersionMetadata{}
	} else {
		versionMetadata, err = LoadVersionMetadata(args.RootDir, providers, int(providerVersion.Major))
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

	buildSchemaReports := BuildSchemaReports{
		NamingDisambiguations:         diagnostics.NamingDisambiguations,
		SkippedPOSTEndpoints:          diagnostics.SkippedPOSTEndpoints,
		ProviderNameErrors:            diagnostics.ProviderNameErrors,
		PathChangesResult:             pathChanges,
		AllResourcesByVersion:         versionMetadata.AllResourcesByVersion,
		AllResourceVersionsByResource: versionMetadata.AllResourceVersionsByResource,
		Pending:                       versionMetadata.Pending,
		CurationViolations:            versionMetadata.CurationViolations,
		AllEndpoints:                  diagnostics.Endpoints,
		InactiveDefaultVersions:       versionMetadata.InactiveDefaultVersions,
	}

	generationResult, err := gen.PulumiSchema(args.RootDir, providers, versionMetadata, providerVersion)

	if err != nil {
		return &BuildSchemaResult{
			PackageSpec: schema.PackageSpec{},
			Metadata:    resources.AzureAPIMetadata{},
			Version:     versionMetadata,
			Reports:     buildSchemaReports,
		}, err
	}

	buildSchemaReports.ForceNewTypes = generationResult.ForceNewTypes
	buildSchemaReports.TypeCaseConflicts = generationResult.TypeCaseConflicts
	buildSchemaReports.FlattenedPropertyConflicts = generationResult.FlattenedPropertyConflicts
	buildSchemaReports.PathConflicts = generationResult.PathConflicts

	pkgSpec := generationResult.Schema
	metadata := generationResult.Metadata
	examples := generationResult.Examples

	if len(args.ExampleLanguages) > 0 {
		// Ensure the spec is stamped with a version - Go gen needs it.
		pkgSpec.Version = args.Version
		err = gen.Examples(args.RootDir, pkgSpec, metadata, examples, args.ExampleLanguages)
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
		Reports:     buildSchemaReports,
		Providers:   providers,
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
	Version      openapi.SdkVersion
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
				excluded, err := curations.IsExcluded(providerName, resourceName, previousVersion)
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
