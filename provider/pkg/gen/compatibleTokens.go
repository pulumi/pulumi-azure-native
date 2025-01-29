package gen

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
)

type Token string
type ModuleLowered string
type ResourceLowered string
type NormalizedPath string

// CompatibleTokensLookup is a lookup table for compatible tokens based on either name matches or their URI paths in Azure.
// Comparisons are done using normalisation for casing and path parameter naming.
type CompatibleTokensLookup struct {
	byLoweredModuleResourceName map[ModuleLowered]map[ResourceLowered]map[Token]struct{}
	byNormalizedPath            map[NormalizedPath]map[Token]struct{}
}

// NewCompatibleTokensLookup creates a new CompatibleResourceLookup from the collection of tokens mapping to their Azure path.
// Note: tokenPaths are outputted for the current provider within the reports directory.
func NewCompatibleTokensLookup(tokenPaths map[string]string) (*CompatibleTokensLookup, error) {
	lookup := CompatibleTokensLookup{
		byLoweredModuleResourceName: map[ModuleLowered]map[ResourceLowered]map[Token]struct{}{},
		byNormalizedPath:            map[NormalizedPath]map[Token]struct{}{},
	}
	if tokenPaths == nil {
		return &lookup, nil
	}
	for token, path := range util.MapOrdered(tokenPaths) {
		moduleName, _, resourceName, err := resources.ParseToken(token)
		if err != nil {
			return nil, fmt.Errorf("failed to parse token for CompatibleTokensLookup %s: %w", token, err)
		}
		lookup.add(openapi.ModuleName(moduleName), resourceName, path, token)
	}
	return &lookup, nil
}

func (lookup *CompatibleTokensLookup) add(moduleName openapi.ModuleName, resourceName openapi.ResourceName, path string, token string) {
	normalizedPath := NormalizedPath(paths.NormalizePath(path))
	moduleLowered := ModuleLowered(moduleName.Lowered())
	resourceLowered := ResourceLowered(strings.ToLower(resourceName))
	typedToken := Token(token)

	if _, ok := lookup.byLoweredModuleResourceName[moduleLowered]; !ok {
		lookup.byLoweredModuleResourceName[moduleLowered] = map[ResourceLowered]map[Token]struct{}{}
	}
	if _, ok := lookup.byLoweredModuleResourceName[moduleLowered][resourceLowered]; !ok {
		lookup.byLoweredModuleResourceName[moduleLowered][resourceLowered] = map[Token]struct{}{}
	}

	lookup.byLoweredModuleResourceName[moduleLowered][resourceLowered][typedToken] = struct{}{}
	if _, ok := lookup.byNormalizedPath[normalizedPath]; !ok {
		lookup.byNormalizedPath[normalizedPath] = map[Token]struct{}{}
	}
	lookup.byNormalizedPath[normalizedPath][typedToken] = struct{}{}
}

func (lookup *CompatibleTokensLookup) IsPopulated() bool {
	return lookup != nil && len(lookup.byLoweredModuleResourceName) > 0 && len(lookup.byNormalizedPath) > 0
}

// FindCompatibleTokens returns a list of all compatible resource tokens for a given module, resource, and path based on the names or path matching.
func (lookup *CompatibleTokensLookup) FindCompatibleTokens(moduleName openapi.ModuleName, resourceName openapi.ResourceName, path string) []string {
	if lookup == nil || lookup.byLoweredModuleResourceName == nil || lookup.byNormalizedPath == nil {
		return nil
	}
	matches := collections.NewOrderableSet[string]()
	moduleLowered := ModuleLowered(moduleName.Lowered())
	resourceLowered := ResourceLowered(strings.ToLower(resourceName))
	normalizedPath := NormalizedPath(paths.NormalizePath(path))
	if byLoweredResourceName, ok := lookup.byLoweredModuleResourceName[moduleLowered]; ok {
		if tokens, ok := byLoweredResourceName[resourceLowered]; ok {
			for token := range tokens {
				matches.Add(string(token))
			}
		}
	}
	if resource, ok := lookup.byNormalizedPath[normalizedPath]; ok {
		for token := range resource {
			matches.Add(string(token))
		}
	}
	return matches.SortedValues()
}
