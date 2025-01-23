package providerlist

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
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

// ResourceType represents a single resource type within a provider namespace of the `az provider list` output.
// This is used to indicate which versions of the API are available for a given resource type.
type ResourceType struct {
	ResourceType      string       `json:"resourceType"`
	DefaultApiVersion *string      `json:"defaultApiVersion"`
	ApiVersions       []ApiVersion `json:"apiVersions"`
	Locations         []string     `json:"locations"`
}

// Provider represents an Azure provider namespace of the `az provider list` output
type Provider struct {
	Namespace     string         `json:"namespace"`
	ResourceTypes []ResourceType `json:"resourceTypes"`
}

// ProviderList represents the top level element of the `az provider list` output
type ProviderList struct {
	Providers []Provider `json:"providers"`
}

// ReadProviderList reads provider_list.json, normalises casing and indexes for fast lookup
func ReadProviderList(providerListPath string) (*ProviderList, error) {
	jsonFile, err := os.Open(providerListPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var providers []Provider
	err = json.Unmarshal(byteValue, &providers)
	if err != nil {
		return nil, err
	}

	return &ProviderList{Providers: providers}, nil
}

type ActiveVersionChecker interface {
	HasProviderVersion(namespace string, version openapi.ApiVersion) bool
}

type providerListIndex struct {
	providerResourceVersionLookup ProviderPathVersions
	providerVersions              map[LoweredProviderName]ApiVersions
}

func (index *providerListIndex) HasProviderVersion(namespace string, version openapi.ApiVersion) bool {
	namespace = strings.ToLower(namespace)
	versions, ok := index.providerVersions[namespace]
	if !ok {
		return false
	}
	return versions.Has(string(version))
}

// Ensure providerListIndex implements ProviderListActiveVersionChecker
var _ ActiveVersionChecker = &providerListIndex{}

func (providers ProviderList) Index() *providerListIndex {
	providerLiveVersions := make(ProviderPathVersions)
	providerVersions := make(map[LoweredProviderName]ApiVersions)

	for _, provider := range providers.Providers {
		namespace := strings.ToLower(provider.Namespace)
		if !strings.HasPrefix(namespace, "microsoft.") {
			continue
		}

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
		providerVersions[namespace] = allVersions
		providerLiveVersions[namespace] = pathVersions
	}

	return &providerListIndex{
		providerResourceVersionLookup: providerLiveVersions,
		providerVersions:              providerVersions,
	}
}
