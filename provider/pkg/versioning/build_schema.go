// Copyright 2023-2025, Pulumi Corporation.

package versioning

import (
	"path"
	"sort"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
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
	// ModuleName -> resourceName -> set of paths, to record resources that have conflicting paths.
	PathConflicts                 map[openapi.ModuleName]map[openapi.ResourceName]map[string][]openapi.ApiVersion
	AllResourcesByVersion         ModuleVersionResources
	AllResourceVersionsByResource ModuleResourceVersions
	Pending                       openapi.ModuleVersionList
	CurationViolations            []CurationViolation
	NamingDisambiguations         []resources.NameDisambiguation
	SkippedPOSTEndpoints          map[openapi.ModuleName]map[string]string
	ModuleNameErrors              []openapi.ModuleNameError
	ForceNewTypes                 []gen.ForceNewType
	TypeCaseConflicts             gen.CaseConflicts
	FlattenedPropertyConflicts    map[string]map[string]any
	AllEndpoints                  map[openapi.ModuleName]map[openapi.ResourceName]map[string]*openapi.Endpoint
	InactiveDefaultVersions       map[openapi.ModuleName][]openapi.ApiVersion
	OldApiVersions                map[openapi.ModuleName][]openapi.ApiVersion
	TokenPaths                    map[string]string
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
		"moduleNameErrors.json":              r.ModuleNameErrors,
		"skippedPOSTEndpoints.json":          r.SkippedPOSTEndpoints,
		"typeCaseConflicts.json":             r.TypeCaseConflicts,
		"oldApiVersions.json":                r.OldApiVersions,
		"tokenPaths.json":                    r.TokenPaths,
	})
}

type BuildSchemaResult struct {
	PackageSpec schema.PackageSpec
	Metadata    resources.AzureAPIMetadata
	Version     VersionMetadata
	Reports     BuildSchemaReports
	Modules     openapi.AzureModules
}

func BuildSchema(args BuildSchemaArgs) (*BuildSchemaResult, error) {
	specsDir := args.Specs.SpecsDir
	if specsDir == "" {
		specsDir = path.Join(args.RootDir, "azure-rest-api-specs")
	}
	modules, diagnostics, err := openapi.ReadAzureModules(specsDir, args.Specs.NamespaceFilter, args.Specs.VersionsFilter)
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
		versionMetadata, err = LoadVersionMetadata(args.RootDir, modules, int(providerVersion.Major))
		if err != nil {
			return nil, err
		}
	}

	if !args.OnlyExplicitVersions {
		modules = openapi.ApplyTransformations(modules, versionMetadata.DefaultVersions, versionMetadata.PreviousDefaultVersions, nil, nil)
	}

	pathChanges := findPathChanges(modules, versionMetadata.DefaultVersions, versionMetadata.PreviousDefaultVersions, versionMetadata.Config)

	if args.ExcludeExplicitVersions {
		modules = openapi.SingleVersion(modules)
	}

	buildSchemaReports := BuildSchemaReports{
		NamingDisambiguations:         diagnostics.NamingDisambiguations,
		SkippedPOSTEndpoints:          diagnostics.SkippedPOSTEndpoints,
		ModuleNameErrors:              diagnostics.ModuleNameErrors,
		PathChangesResult:             pathChanges,
		AllResourcesByVersion:         versionMetadata.AllResourcesByVersion,
		AllResourceVersionsByResource: versionMetadata.AllResourceVersionsByResource,
		Pending:                       versionMetadata.Pending,
		CurationViolations:            versionMetadata.CurationViolations,
		AllEndpoints:                  diagnostics.Endpoints,
		InactiveDefaultVersions:       versionMetadata.InactiveDefaultVersions,
		OldApiVersions:                versionMetadata.GetOldApiVersionsPerModule(),
	}

	generationResult, err := gen.PulumiSchema(args.RootDir, modules, versionMetadata, providerVersion)

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
	buildSchemaReports.TokenPaths = generationResult.TokenPaths

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
		Modules:     modules,
	}, nil
}

type PathChange struct {
	CurrentPath  string
	PreviousPath string
	ResourceName string
}

type MissingExpectedResourceVersion struct {
	ModuleName   openapi.ModuleName
	ResourceName string
	Version      openapi.ApiVersion
	IsPrevious   bool
}

type PathChangesResult struct {
	Changes                        []PathChange
	MissingPreviousDefaultVersions []MissingExpectedResourceVersion
}

func findPathChanges(modules openapi.AzureModules,
	defaultVersions openapi.DefaultVersions,
	previousDefaultVersions openapi.DefaultVersions,
	curations Curations) PathChangesResult {

	result := []PathChange{}
	missingPreviousDefaultVersions := []MissingExpectedResourceVersion{}

	for moduleName, resources := range util.MapOrdered(defaultVersions) {
		previousResources, ok := previousDefaultVersions[moduleName]
		if !ok {
			continue
		}
		moduleVersions := modules[moduleName]

		sortedResourceNames := []string{}
		for resourceName := range resources {
			sortedResourceNames = append(sortedResourceNames, resourceName)
		}
		sort.Strings(sortedResourceNames)

		for _, resourceName := range sortedResourceNames {
			defaultVersion := resources[resourceName].ApiVersion
			previousDefault, ok := previousResources[resourceName]
			if !ok {
				continue
			}

			cur := moduleVersions[defaultVersion]
			prev := moduleVersions[previousDefault.ApiVersion]

			spec, ok := cur.Resources[resourceName]
			if !ok {
				spec, ok = cur.Invokes[resourceName]
			}
			if !ok {
				missingPreviousDefaultVersions = append(missingPreviousDefaultVersions, MissingExpectedResourceVersion{
					ModuleName:   moduleName,
					ResourceName: resourceName,
					Version:      previousDefault.ApiVersion,
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
					ModuleName:   moduleName,
					ResourceName: resourceName,
					Version:      previousDefault.ApiVersion,
					IsPrevious:   true,
				})
				continue
			}

			path := paths.NormalizePath(spec.Path)
			prevPath := paths.NormalizePath(prevSpec.Path)

			if path != prevPath {
				excluded, err := curations.IsExcluded(moduleName, resourceName, previousDefault.ApiVersion)
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
