package providerlist

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
)

// LoweredProviderName e.g. analysisservices
type LoweredProviderName = string

// ResourcePath is a lowered resource path e.g. locations or locations/checknameavailability
type ResourcePath = string

// ApiVersions StringSet of versions e.g. 2019-03-01-preview or 2018-05-05
type ApiVersions = codegen.StringSet

// ApiVersion e.g. 2020-01-02
type ApiVersion = string

// ProviderPathVersions is a map of lowered provider names to Api Versions e.g. `analysisservices -> "locations/checknameavailability" -> [2019-03-01-preview, 2018-05-05]`
type ProviderPathVersions = map[LoweredProviderName]map[ResourcePath]ApiVersions

// ReadProviderList reads provider_list.json, normalises casing and indexes for fast lookup
func ReadProviderList(providerListPath string) (ProviderPathVersions, error) {
	jsonFile, err := os.Open(providerListPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var providers []provider
	err = json.Unmarshal(byteValue, &providers)
	if err != nil {
		return nil, err
	}

	return toProviderPathVersions(providers), nil
}

type ProviderPathVersionsJson = map[LoweredProviderName]map[ResourcePath][]ApiVersion

// FormatProviderPathVersionsJson prepares the active path versions for writing to JSON – replacing the string set with a string array
func FormatProviderPathVersionsJson(activePathVersions ProviderPathVersions) ProviderPathVersionsJson {
	formatted := ProviderPathVersionsJson{}
	for providerName, paths := range activePathVersions {
		formattedProvider := map[string][]string{}
		for resourcePath, versions := range paths {
			sortedVersions := versions.SortedValues()
			if sortedVersions == nil {
				sortedVersions = []string{}
			}
			formattedProvider[resourcePath] = sortedVersions
		}
		formatted[providerName] = formattedProvider
	}
	return formatted
}

type provider struct {
	Namespace     string    `json:"namespace"`
	ResourceTypes []provRes `json:"resourceTypes"`
}

type provRes struct {
	ResourceType string   `json:"resourceType"`
	ApiVersions  []string `json:"apiVersions"`
}

func toProviderPathVersions(providers []provider) ProviderPathVersions {
	providerLiveVersions := make(ProviderPathVersions)

	for _, provider := range providers {
		namespace := strings.ToLower(provider.Namespace)
		if !strings.HasPrefix(namespace, "microsoft.") {
			continue
		}
		providerName := strings.TrimPrefix(namespace, "microsoft.")

		pathVersions := map[ResourcePath]ApiVersions{}
		allVersions := codegen.NewStringSet()
		for _, resourceType := range provider.ResourceTypes {
			versions := codegen.NewStringSet()
			for _, version := range resourceType.ApiVersions {
				allVersions.Add(version)
				versions.Add(version)
			}
			rtName := strings.ToLower(resourceType.ResourceType)
			pathVersions[rtName] = versions
		}

		pathVersions[""] = allVersions
		providerLiveVersions[providerName] = pathVersions
	}

	return providerLiveVersions
}
