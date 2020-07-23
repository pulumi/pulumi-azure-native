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
	// IsRequired is true for mandatory parameters.
	IsRequired bool `json:"required,omitempty"`
	// Value contains metadata for path/query parameters.
	Value *AzureApiProperty `json:"value"`
	// Body contains metadata for the body parameter.
	Body *AzureApiType `json:"body,omitempty"`
}

// AzureApiProperty represents validation constraints for a single parameter or body property.
type AzureApiProperty struct {
	Type      string            `json:"type,omitempty"`
	Items     *AzureApiProperty `json:"items,omitempty"`
	Enum      []string          `json:"enum,omitempty"`
	Ref       string            `json:"$ref,omitempty"`
	MinLength *int64            `json:"minLength,omitempty"`
	MaxLength *int64            `json:"maxLength,omitempty"`
	Pattern   string            `json:"pattern,omitempty"`
}

// AzureApiType represents the shape of an object property.
type AzureApiType struct {
	Properties         map[string]AzureApiProperty `json:"properties,omitempty"`
	RequiredProperties []string                    `json:"required,omitempty"`
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
	Types     map[string]AzureApiType     `json:"types"`
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

// ResourceProvider returns a provider name given resource's PUT path.
func ResourceProvider(path string) string {
	if path == "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}" {
		return "Core"
	}

	parts := strings.Split(path, "/")

	// The path is /subscriptions/id/resoursegroups/id/providers/Microsoft.SomeProvider/foos/id/bars/id/...
	if len(parts) < 8 {
		return ""
	}

	armProvider := parts[6]
	// TODO: handle non-Microsoft providers.
	if !strings.HasPrefix(armProvider, "Microsoft.") {
		return ""
	}

	return strings.TrimPrefix(armProvider, "Microsoft.")
}

var verbReplacer *strings.Replacer
var wellKnownNames map[string]string

func init() {
	verbReplacer = strings.NewReplacer("GetProperties", "", "Get", "", "getByName", "", "get", "", "List", "")
	wellKnownNames = map[string]string {
		"Redis": "Redis",
		"Caches": "Cache",
		"AssessmentsMetadata": "AssessmentMetadata",
		"Mediaservices": "MediaService",
	}
}

// ResourceName constructs a name of a resource based on Get or List operation ID,
// e.g. "Managers_GetActivationKey" -> "ManagerActivationKey".
func ResourceName(operationId string) string {
	parts := strings.Split(operationId, "_")
	var name, verb string
	if len(parts) == 1 {
		verb = parts[0]
	} else {
		if v, ok := wellKnownNames[parts[0]]; ok {
			name = v
		} else {
			name = inflector.Singularize(parts[0])
		}
		verb = parts[1]
	}

	subName := verbReplacer.Replace(verb)

	if strings.HasPrefix(subName, name) {
		return subName
	}

	return name + subName
}
