// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"github.com/gedex/inflector"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// AzureApiParameter represents a parameter of a Azure REST API endpoint.
type AzureApiParameter struct {
	Name string `json:"name"`
	// Location defines the parameter's place the HTTP request: "path", "query", or "body".
	Location string `json:"location"`
	// Source defines the value source: "method" (resource arguments) or "client" (provider configuration).
	Source string `json:"source,omitempty"`
	// Requires is true for mandatory parameters.
	IsRequired bool `json:"required,omitempty"`
	// Properties contains the names of properties for a body-placed parameter.
	Properties []string `json:"properties,omitempty"`
	// RequiredProperties contains the names of required properties for a body-placed parameter.
	RequiredProperties []string `json:"requiredProperties,omitempty"`
}

// AzureApiResource is a resource in Azure REST API.
type AzureApiResource struct {
	ApiVersion    string              `json:"apiVersion"`
	Path          string              `json:"path"`
	GetParameters []AzureApiParameter `json:"GET"`
	PutParameters []AzureApiParameter `json:"PUT"`
}

// AzureApiInvoke is an invocation target (a function) in Azure REST API.
type AzureApiInvoke struct {
	ApiVersion     string              `json:"apiVersion"`
	Path           string              `json:"path"`
	GetParameters  []AzureApiParameter `json:"GET"`
	PostParameters []AzureApiParameter `json:"POST"`
}

// AzureApiMetadata is a collection of all resources and functions in the Azure REST API surface.
type AzureApiMetadata struct {
	Resources map[string]AzureApiResource `json:"resources"`
	Invokes   map[string]AzureApiInvoke   `json:"invokes"`
}

// blessedVersions contains the preferred versions per resource provider. If a resource provider is not specified,
// the latest stable version is used.
var blessedVersions map[string]string

func init() {
	blessedVersions = map[string]string{
		"Microsoft.Cdn": "2020-03-31", // the later version throws runtime errors
		"Microsoft.Web": "2019-08-01", // an earlier version contains extra files that spoil generation
	}
}

// SwaggerLocations returns a slice of URLs of all known Azure Resource Manager swagger files.
func SwaggerLocations() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := filepath.Glob(dir + "/azure-rest-api-specs/specification/**/resource-manager/Microsoft.*/stable/*/*.json")
	if err != nil {
		return nil, err
	}

	// Sorting alphabetically means the schemas with the latest API version are the last.
	sort.Strings(files)

	// Deduplicate files of the same provider having the same name.
	latest := map[string]string{}
	for _, file := range files {
		parts := strings.Split(file, "/")
		length := len(parts)
		provider := parts[length-4]
		apiVersion := parts[length-2]
		filename := parts[length-1]

		if v, ok := blessedVersions[provider]; ok && v != apiVersion {
			continue
		}

		key := provider + "/" + filename
		latest[key] = file
	}

	var locations []string
	for _, file := range latest {
		locations = append(locations, file)
	}
	sort.Strings(locations)

	return locations, nil
}

// ResourceQualifiedName returns a tuple of (module, resource name) for a given PUT path.
func ResourceQualifiedName(path string) (string, string) {
	if path == "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}" {
		return "core", "ResourceGroup"
	}

	parts := strings.Split(path, "/")

	// The path is /subscriptions/id/resoursegroups/id/providers/Microsoft.SomeProvider/foos/id/bars/id/...
	if len(parts) < 8 || len(parts)%2 != 1 {
		return "", ""
	}

	armProvider := parts[6]
	// TODO: handle non-Microsoft providers.
	if !strings.HasPrefix(armProvider, "Microsoft.") {
		return "", ""
	}

	// Build a name like SomeProvider:FooBar.
	provider := armProvider[10:]
	resource := ""
	for i := 7; i < len(parts); i += 2 {
		if strings.Contains(parts[i], "{") {
			// We don't support this shape of URLs yet. Example: dnszones/{zoneName}/{recordType}/{relativeRecordSetName}.
			return "", ""
		}

		// TODO: generalize this case to a map of well-known aliases.
		switch strings.ToLower(parts[i]) {
		case "redis":
			resource += "Redis"
		case "sites":
			resource += "AppService"
		case "serverfarms":
			resource += "AppServicePlan"
		default:
			// TODO: we may get better singular names from some metadata.
			resource += strings.ReplaceAll(strings.Title(inflector.Singularize(parts[i])), "-", "")
		}
	}
	return provider, resource
}
