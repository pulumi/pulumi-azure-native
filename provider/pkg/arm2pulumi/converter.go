package arm2pulumi

import (
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/provider"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"path/filepath"
	"regexp"
	"strings"
)

var stripScopeRegex = regexp.MustCompile(`.*providers/`)

func newResourceTokenConverter(metadata *provider.AzureAPIMetadata) *resourceTokenConverter {
	stableResourceTypeToTokenMap := map[string][]tokenWrapper{}
	previewResourceTypeToTokenMap := map[string][]tokenWrapper{}
	for k, v := range metadata.Resources {
		armResourceType := parseArmResourceType(v.Path)
		if armResourceType == "" {
			continue
		}
		resourceTypeToTokenMap := stableResourceTypeToTokenMap
		if openapi.IsPreview(v.APIVersion) {
			resourceTypeToTokenMap = previewResourceTypeToTokenMap
		}
		existing := resourceTypeToTokenMap[armResourceType]
		insertAt := 0
		duplicate := false
		for i, e := range existing {
			if e.apiVersion == v.APIVersion {
				duplicate = true
				break
			}
			if e.apiVersion > v.APIVersion {
				insertAt = i
				break
			}
			insertAt = i + 1
		}
		if duplicate {
			continue
		}
		existing = append(existing, tokenWrapper{})
		copy(existing[insertAt+1:], existing[insertAt:])
		existing[insertAt] = tokenWrapper{apiVersion: v.APIVersion, resourceToken: k}
		resourceTypeToTokenMap[armResourceType] = existing
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
