// Copyright 2021, Pulumi Corporation.  All rights reserved.

package arm2pulumi

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

var stripScopeRegex = regexp.MustCompile(`.*providers/`)

var isLatestRegex = regexp.MustCompile(`azure-native:[^/]*:.*`)

func newResourceTokenConverter(metadata *resources.AzureAPIMetadata) *resourceTokenConverter {
	stableResourceTypeToTokenWrapperSet := map[string]map[string]tokenWrapper{}
	previewResourceTypeToTokenWrapperSet := map[string]map[string]tokenWrapper{}
	for k, v := range metadata.Resources {
		if isLatestRegex.MatchString(k) {
			continue // Skip latest since metadata will also contain the versioned variant anyway.
			// This helps with keeping the output deterministic.
		}
		armResourceType := parseArmResourceType(v.Path)
		if armResourceType == "" {
			continue
		}
		resourceTypeToTokenMap := stableResourceTypeToTokenWrapperSet
		if openapi.IsPreview(v.APIVersion) {
			resourceTypeToTokenMap = previewResourceTypeToTokenWrapperSet
		}
		existing := resourceTypeToTokenMap[armResourceType]
		if existing == nil {
			existing = map[string]tokenWrapper{}
		}
		if _, ok := existing[v.APIVersion]; ok {
			continue
		}
		existing[v.APIVersion] = tokenWrapper{apiVersion: v.APIVersion, resourceToken: k}
		resourceTypeToTokenMap[armResourceType] = existing
	}

	stableResourceTypeToTokenMap := map[string][]tokenWrapper{}
	for k, m := range stableResourceTypeToTokenWrapperSet {
		var tokenWrappers []tokenWrapper
		sorted := codegen.SortedKeys(m)
		for _, s := range sorted {
			tokenWrappers = append(tokenWrappers, m[s])
		}
		stableResourceTypeToTokenMap[k] = tokenWrappers
	}

	previewResourceTypeToTokenMap := map[string][]tokenWrapper{}
	for k, m := range previewResourceTypeToTokenWrapperSet {
		var tokenWrappers []tokenWrapper
		sorted := codegen.SortedKeys(m)
		for _, s := range sorted {
			tokenWrappers = append(tokenWrappers, m[s])
		}
		previewResourceTypeToTokenMap[k] = tokenWrappers
	}

	return &resourceTokenConverter{
		stableResourceTypeToTokenMap:  stableResourceTypeToTokenMap,
		previewResourceTypeToTokenMap: previewResourceTypeToTokenMap,
	}
}

func parseArmResourceType(path string) string {
	if !stripScopeRegex.MatchString(path) {
		return ""
	}
	stripped := stripScopeRegex.ReplaceAllString(path, "")
	token := strings.ToLower(filepath.Clean(stripped))
	parts := strings.Split(token, "/")
	var tokenParts []string
	for _, part := range parts {
		contract.Assertf(len(part) > 0, "invalid path '%s'", path)
		if part[0] == '{' && part[len(part)-1] == '}' {
			continue
		}
		tokenParts = append(tokenParts, part)
	}
	return strings.Join(tokenParts, "/")
}

type tokenWrapper struct {
	apiVersion    string
	resourceToken string
}

type resourceTokenConverter struct {
	stableResourceTypeToTokenMap  map[string][]tokenWrapper
	previewResourceTypeToTokenMap map[string][]tokenWrapper
}

func (r *resourceTokenConverter) convert(armResourceType string, apiVersion string) string {
	resourceTypeToTokenMap := r.stableResourceTypeToTokenMap
	if openapi.IsPreview(apiVersion) {
		resourceTypeToTokenMap = r.previewResourceTypeToTokenMap
	}
	found := resourceTypeToTokenMap[strings.ToLower(armResourceType)]
	for _, f := range found {
		if f.apiVersion >= apiVersion {
			return f.resourceToken
		}
	}
	return ""
}
