package gen

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
)

type Token string
type ModuleLowered string
type ResourceLowered string
type NormalizedPath string

type CompatibleTokensLookup struct {
	byLoweredModuleResourceName map[ModuleLowered]map[ResourceLowered]map[Token]struct{}
	byNormalizedPath            map[NormalizedPath]map[Token]struct{}
}

// NewCompatibleTokensLookup creates a new CompatibleResourceLookup from the given Azure modules and default versions.
func NewCompatibleTokensLookup(tokenPaths map[string]string) (*CompatibleTokensLookup, error) {
	lookup := CompatibleTokensLookup{
		byLoweredModuleResourceName: map[ModuleLowered]map[ResourceLowered]map[Token]struct{}{},
		byNormalizedPath:            map[NormalizedPath]map[Token]struct{}{},
	}
	if tokenPaths == nil {
		tokenPaths = map[string]string{}
	}
	for token, path := range util.MapOrdered(tokenPaths) {
		tokenParts := strings.Split(token, ":")
		if len(tokenParts) != 3 {
			return nil, fmt.Errorf("malformed token '%s'", token)
		}
		tokenModule, resourceName := tokenParts[1], tokenParts[2]
		moduleName, _, _ := strings.Cut(tokenModule, "/")
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

func (lookup *CompatibleTokensLookup) FindCompatibleTokens(moduleName openapi.ModuleName, resourceName openapi.ResourceName, path string) []string {
	if lookup == nil || lookup.byLoweredModuleResourceName == nil || lookup.byNormalizedPath == nil {
		return nil
	}
	matches := map[string]struct{}{}
	moduleLowered := ModuleLowered(moduleName.Lowered())
	resourceLowered := ResourceLowered(strings.ToLower(resourceName))
	normalizedPath := NormalizedPath(paths.NormalizePath(path))
	if byLoweredResourceName, ok := lookup.byLoweredModuleResourceName[moduleLowered]; ok {
		if tokens, ok := byLoweredResourceName[resourceLowered]; ok {
			for token := range tokens {
				matches[string(token)] = struct{}{}
			}
		}
	}
	if resource, ok := lookup.byNormalizedPath[normalizedPath]; ok {
		for token := range resource {
			matches[string(token)] = struct{}{}
		}
	}
	return util.SortedKeys(matches)
}
