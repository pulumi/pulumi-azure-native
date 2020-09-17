// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"strings"

	"github.com/gedex/inflector"
)

// AzureAPIParameter represents a parameter of a Azure REST API endpoint.
type AzureAPIParameter struct {
	Name string `json:"name"`
	// Location defines the parameter's place the HTTP request: "path", "query", or "body".
	Location string `json:"location"`
	// Source defines the value source: "method" (resource arguments) or "client" (provider configuration).
	Source string `json:"source,omitempty"`
	// IsRequired is true for mandatory parameters.
	IsRequired bool `json:"required,omitempty"`
	// Value contains metadata for path/query parameters.
	Value *AzureAPIProperty `json:"value"`
	// Body contains metadata for the body parameter.
	Body *AzureAPIType `json:"body,omitempty"`
}

// AzureAPIProperty represents validation constraints for a single parameter or body property.
type AzureAPIProperty struct {
	Type      string            `json:"type,omitempty"`
	Items     *AzureAPIProperty `json:"items,omitempty"`
	Enum      []string          `json:"enum,omitempty"`
	OneOf     []string          `json:"oneOf,omitempty"`
	Ref       string            `json:"$ref,omitempty"`
	Const     interface{}       `json:"const,omitempty"`
	Minimum   *float64          `json:"minimum,omitempty"`
	Maximum   *float64          `json:"maximum,omitempty"`
	MinLength *int64            `json:"minLength,omitempty"`
	MaxLength *int64            `json:"maxLength,omitempty"`
	Pattern   string            `json:"pattern,omitempty"`
	// The name in the SDK if different from the wire-serialized name, empty otherwise.
	SdkName string `json:"sdkName,omitempty"`
	// The name of a container property that was "flattened" during SDK generation, i.e. an extra layer that exists
	// in the API payload but does not exist in the SDK.
	Container string `json:"container,omitempty"`
	// Whether a change in the value of the property requires a replacement of the whole resource
	// (i.e., no in-place updates allowed).
	ForceNew bool `json:"forceNew,omitempty"`
}

// AzureAPIType represents the shape of an object property.
type AzureAPIType struct {
	Properties         map[string]AzureAPIProperty `json:"properties,omitempty"`
	RequiredProperties []string                    `json:"required,omitempty"`
}

// AzureAPIResource is a resource in Azure REST API.
type AzureAPIResource struct {
	APIVersion    string                      `json:"apiVersion"`
	Path          string                      `json:"path"`
	GetParameters []AzureAPIParameter         `json:"GET"`
	PutParameters []AzureAPIParameter         `json:"PUT"`
	Response      map[string]AzureAPIProperty `json:"response"`
}

// AzureAPIExample provides a pointer to examples relevant to a resource from the Azure REST API spec.
type AzureAPIExample struct {
	Description string `json:"description"`
	Location    string `json:"location"`
}

// AzureAPIInvoke is an invocation target (a function) in Azure REST API.
type AzureAPIInvoke struct {
	APIVersion     string                      `json:"apiVersion"`
	Path           string                      `json:"path"`
	GetParameters  []AzureAPIParameter         `json:"GET"`
	PostParameters []AzureAPIParameter         `json:"POST"`
	Response       map[string]AzureAPIProperty `json:"response"`
}

// AzureAPIMetadata is a collection of all resources and functions in the Azure REST API surface.
type AzureAPIMetadata struct {
	Types     map[string]AzureAPIType     `json:"types"`
	Resources map[string]AzureAPIResource `json:"resources"`
	Invokes   map[string]AzureAPIInvoke   `json:"invokes"`
}

// Some resource providers changed capitalization between API versions, but we should map them to the
// same spelling so that folder names and namespaces are consistent. The map below provides such
// canonical names based on which names seems to be used prominently as of 2020.
var wellKnownProviderNames = map[string]string{
	"aad":             "Aad",
	"aadiam":          "AadIam",
	"dbformariadb":    "DBforMariaDB",
	"dbformysql":      "DBforMySQL",
	"dbforpostgresql": "DBforPostgreSQL",
	"visualstudio":    "VisualStudio",
}

// ResourceProvider returns a provider name given resource's PUT path.
func ResourceProvider(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return ""
	}

	for _, part := range parts {
		if strings.HasPrefix(part, "Microsoft.") {
			name := strings.TrimPrefix(part, "Microsoft.")
			if knownName, ok := wellKnownProviderNames[strings.ToLower(name)]; ok {
				return knownName
			}
			return name
		}
		if strings.HasPrefix(part, "microsoft.") {
			name := strings.Title(strings.TrimPrefix(part, "microsoft."))
			if knownName, ok := wellKnownProviderNames[strings.ToLower(name)]; ok {
				return knownName
			}
			return name
		}
	}

	// This could cause some undesired resources in the Resources namespace, but it looks okay for now.
	return "Resources"
}

var verbReplacer = strings.NewReplacer("GetProperties", "", "Get", "", "getByName", "", "get", "", "List", "")
var wellKnownNames = map[string]string{
	"Redis":               "Redis",
	"Caches":              "Cache",
	"AssessmentsMetadata": "AssessmentMetadata",
	"Mediaservices":       "MediaService",
}

// ResourceName constructs a name of a resource based on Get or List operation ID,
// e.g. "Managers_GetActivationKey" -> "ManagerActivationKey".
func ResourceName(operationID string) string {
	parts := strings.Split(operationID, "_")
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

	// Sometimes, name or its part is included in the operation name. We want to eliminate obvious duplication
	// in the resulting resource name.
	nameParts := splitCamelCase(name)
	for i := 0; i < len(nameParts); i++ {
		namePrefix := strings.Join(nameParts[:i], "")
		nameSuffix := strings.Join(nameParts[i:], "")
		if strings.HasPrefix(subName, nameSuffix) {
			return namePrefix + subName
		}
	}

	return name + subName
}

func splitCamelCase(s string) (entries []string) {
	current := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			if current != "" {
				entries = append(entries, current)
			}
			current = string(v)
		} else {
			current += string(v)
		}
	}
	entries = append(entries, current)
	return
}
