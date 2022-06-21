package versioning

import (
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
)

func FindDeprecations(specVersions SpecVersions, curatedVersion openapi.CuratedVersion) openapi.ProviderVersionList {
	deprecationCutoff := "2021-01-01"
	olderVersions := findOlderVersions(specVersions, curatedVersion)
	for name, versions := range olderVersions {
		filteredVersions := []openapi.ApiVersion{}
		for _, version := range versions {
			if version < deprecationCutoff {
				filteredVersions = append(filteredVersions, version)
			}
		}
		olderVersions[name] = filteredVersions
	}
	return olderVersions
}
