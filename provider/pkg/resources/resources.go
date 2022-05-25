// Copyright 2016-2020, Pulumi Corporation.

package resources

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gedex/inflector"
)

// AzureAPIParameter represents a parameter of a Azure REST API endpoint.
type AzureAPIParameter struct {
	Name string `json:"name"`
	// Location defines the parameter's place the HTTP request: "path", "query", or "body".
	Location string `json:"location"`
	// IsRequired is true for mandatory parameters.
	IsRequired bool `json:"required,omitempty"`
	// Value contains metadata for path/query parameters.
	Value *AzureAPIProperty `json:"value"`
	// Body contains metadata for the body parameter.
	Body *AzureAPIType `json:"body,omitempty"`
}

type AutoNameKind string

const (
	AutoNameRandom AutoNameKind = "random"
	AutoNameCopy   AutoNameKind = "copy"
	AutoNameUuid   AutoNameKind = "uuid"
)

// AzureAPIProperty represents validation constraints for a single parameter or body property.
type AzureAPIProperty struct {
	Type                 string            `json:"type,omitempty"`
	Items                *AzureAPIProperty `json:"items,omitempty"`
	AdditionalProperties *AzureAPIProperty `json:"additionalProperties,omitempty"`
	OneOf                []string          `json:"oneOf,omitempty"`
	Ref                  string            `json:"$ref,omitempty"`
	Const                interface{}       `json:"const,omitempty"`
	Minimum              *float64          `json:"minimum,omitempty"`
	Maximum              *float64          `json:"maximum,omitempty"`
	MinLength            *int64            `json:"minLength,omitempty"`
	MaxLength            *int64            `json:"maxLength,omitempty"`
	Pattern              string            `json:"pattern,omitempty"`
	// The name in the SDK if different from the wire-serialized name, empty otherwise.
	SdkName string `json:"sdkName,omitempty"`
	// The names of container properties that were "flattened" during SDK generation, i.e. extra layers that exist
	// in the API payload but do not exist in the SDK. The order is from the top-most to bottom-most.
	Containers []string `json:"containers,omitempty"`
	// Whether a change in the value of the property requires a replacement of the whole resource
	// (i.e., no in-place updates allowed).
	ForceNew bool `json:"forceNew,omitempty"`
	// If the property is a resource name where we should apply auto-naming, this will contain the kind of
	// auto-naming strategy. Possible values are:
	// - "copy" for 1-to-1 copy of the resource's logical name.
	// - "random" for adding a random suffix to resource's logical name.
	AutoName AutoNameKind `json:"autoname,omitempty"`
}

// AzureAPIType represents the shape of an object property.
type AzureAPIType struct {
	Properties         map[string]AzureAPIProperty `json:"properties,omitempty"`
	RequiredProperties []string                    `json:"required,omitempty"`
}

// AzureAPIResource is a resource in Azure REST API.
type AzureAPIResource struct {
	APIVersion string `json:"apiVersion"`
	Path       string `json:"path"`
	// HTTP method to create/update the resource. Defaults to PUT if empty.
	UpdateMethod  string                      `json:"updateMethod,omitempty"`
	PutParameters []AzureAPIParameter         `json:"PUT"`
	Response      map[string]AzureAPIProperty `json:"response"`
	// A singleton resource is created by Azure with its parent.
	// It can't be created or deleted explicitly on its own.
	Singleton bool `json:"singleton,omitempty"`
	// DefaultBody is the default state of a resource for resources that are
	// created automatically by Azure. Note:
	// - omitempty is not set to distinct between nil (no default) and empty maps (empty default)
	// - DefaultBody applies to all singletons but also some non-singleton resources that can be
	//   deleted despite being created automatically.
	DefaultBody map[string]interface{} `json:"defaultBody"`
	// For long-running operations, defines the style of the Asynchronous Operation, see
	// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#asynchronous-operations
	// Possible values: nil (not async), "azure-async-operation", "location", "original-uri", see
	// https://github.com/Azure/autorest/blob/master/docs/extensions/readme.md#x-ms-long-running-operation
	PutAsyncStyle    string `json:"putAsyncStyle,omitempty"`
	DeleteAsyncStyle string `json:"deleteAsyncStyle,omitempty"`
	// HTTP method to read resource state. Defaults to GET if empty.
	ReadMethod string `json:"readMethod,omitempty"`
	// Path to append to the resource ID to produce the URL for to read resource state.
	ReadPath string `json:"readPath,omitempty"`
	// By default, we populate the `location` property of every resource to the location of its resource
	// group or the configured value. AutoLocationDisabled can override this default behavior.
	AutoLocationDisabled bool `json:"autoLocationDisabled,omitempty"`
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
type PartialAzureAPIMetadata struct {
	Types     PartialMap[AzureAPIType]     `json:"types"`
	Resources PartialMap[AzureAPIResource] `json:"resources"`
	Invokes   PartialMap[AzureAPIInvoke]   `json:"invokes"`
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
	"aad":                          "Aad",
	"aadiam":                       "AadIam",
	"dbformariadb":                 "DBforMariaDB",
	"dbformysql":                   "DBforMySQL",
	"dbforpostgresql":              "DBforPostgreSQL",
	"powerbidedicated":             "PowerBIDedicated",
	"servicefabricmanagedclusters": "ServiceFabric", // https://github.com/Azure/azure-rest-api-specs/issues/15867
	"visualstudio":                 "VisualStudio",
}

// For the cases below, we use folder (SDK) name for module names instead of the ARM name.
var folderModulePattern = regexp.MustCompile(`.*/specification/([a-z]+)/resource-manager/.*`)
var folderModuleNames = map[string]string{
	"videoanalyzer": "VideoAnalyzer",
	"webpubsub":     "WebPubSub",
}

// ResourceProvider returns a provider name given Open API spec file and resource's API URI.
func ResourceProvider(filePath, apiUri string) string {
	// Start with extracting the provider from the folder path. If the folder name is explicitly listed,
	// use it as the provider name. This is the new style we use for newer resources after 1.0. Older
	// resources to be migrated as part of https://github.com/pulumi/pulumi-azure-native/issues/690.
	subMatches := folderModulePattern.FindStringSubmatch(filePath)
	if len(subMatches) > 1 {
		moduleAlias := subMatches[1]
		if name, ok := folderModuleNames[moduleAlias]; ok {
			return name
		}
	}
	// We extract the provider name from two sources:
	// - from the folder name of the Open API spec
	// - from the URI of the API endpoint (we take the last provider in the URI)
	fileProvider := resourceProvider(filePath, "")
	apiProvider := resourceProvider(apiUri, "Resources")
	// We proceed with the endpoint if both provider values match. This way, we avoid flukes and
	// declarations outside of the current API provider.
	if strings.ToLower(fileProvider) != strings.ToLower(apiProvider) {
		return ""
	}
	return fileProvider
}

func resourceProvider(path, defaultValue string) string {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return ""
	}

	for i := len(parts) - 2; i >= 0; i-- {
		part := parts[i]
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

	return defaultValue
}

var verbReplacer = strings.NewReplacer("GetProperties", "", "Get", "", "getByName", "", "get", "", "ListByResourceName", "", "List", "", "list", "", "CheckEntityExists", "")
var wellKnownNames = map[string]string{
	"AssessmentsMetadata": "AssessmentMetadata",
	"Caches":              "Cache",
	"Metadata":            "Metadata",
	"Mediaservices":       "MediaService",
	"Redis":               "Redis",
}

// ResourceName constructs a name of a resource based on Get or List operation ID,
// e.g. "Managers_GetActivationKey" -> "ManagerActivationKey".
func ResourceName(operationID string) string {
	normalizedID := strings.ReplaceAll(operationID, "-", "_")
	parts := strings.Split(normalizedID, "_")
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
	subNameParts := splitCamelCase(inflector.Singularize(subName))
	for i := 0; i < len(nameParts); i++ {
		if sliceHasPrefix(subNameParts, nameParts[i:]) {
			namePrefix := strings.Join(nameParts[:i], "")
			return namePrefix + subName
		}
	}

	return name + subName
}

var referenceNameReplacer = strings.NewReplacer("CreateOrUpdateParameters", "", "Create", "", "Request", "")

// DiscriminatedResourceName returns the name of a resource variant based on the variant's schema name.
func DiscriminatedResourceName(referenceName string) string {
	return referenceNameReplacer.Replace(referenceName)
}

// AutoNamer decides on the per-property auto-naming property for a given API path.
type AutoNamer struct {
	path string
}

// NewAutoNamer creates a new AutoNamer for a given API path.
func NewAutoNamer(path string) AutoNamer {
	return AutoNamer{path}
}

// AutoName returns auto-naming strategy ("random", "copy") for a given property name and a resource path.
// The second value is true if auto-naming should be applied.
func (a *AutoNamer) AutoName(name, format string) (AutoNameKind, bool) {
	suffix := fmt.Sprintf("{%s}", name)
	if !strings.HasSuffix(a.path, suffix) {
		return "", false
	}

	if format == "uuid" {
		return AutoNameUuid, true
	}

	switch name {
	// Endpoints and custom domains both produce URIs, so they are globally unique.
	case "endpointName", "customDomainName":
		return AutoNameRandom, true
	// Work around https://github.com/Azure/azure-rest-api-specs/issues/1684.
	case "roleAssignmentName", "roleDefinitionId":
		return AutoNameUuid, true
	}

	parts := strings.Split(strings.ToLower(a.path), "microsoft.")
	resourcePath := parts[len(parts)-1]
	if len(parts) == 1 || strings.Count(resourcePath, "{") == 1 {
		return AutoNameRandom, true
	}

	return AutoNameCopy, true
}

// AutoLocationDisabled returns true auto-location should not be applied for a given resource path.
func AutoLocationDisabled(path string) bool {
	switch strings.ToLower(path) {
	case "/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.resources/deployments/{deploymentname}":
		// Template deployment.
		return true
	case "/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.network/loadbalancers/{loadbalancername}/backendaddresspools/{backendaddresspoolname}":
		// Load balancer backend address pool, see https://github.com/pulumi/pulumi-azure-native/issues/819.
		return true
	case "/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.documentdb/databaseaccounts/{accountname}/cassandrakeyspaces/{keyspacename}",
		"/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.documentdb/databaseaccounts/{accountname}/cassandrakeyspaces/{keyspacename}/tables/{tablename}",
		"/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.documentdb/databaseaccounts/{accountname}/gremlindatabases/{databasename}",
		"/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.documentdb/databaseaccounts/{accountname}/gremlindatabases/{databasename}/graphs/{graphname}",
		"/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.documentdb/databaseaccounts/{accountname}/mongodbdatabases/{databasename}",
		"/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.documentdb/databaseaccounts/{accountname}/mongodbdatabases/{databasename}/collections/{collectionname}",
		"/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.documentdb/databaseaccounts/{accountname}/sqldatabases/{databasename}",
		"/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.documentdb/databaseaccounts/{accountname}/sqldatabases/{databasename}/containers/{containername}",
		"/subscriptions/{subscriptionid}/resourcegroups/{resourcegroupname}/providers/microsoft.documentdb/databaseaccounts/{accountname}/tables/{tablename}":
		// Child resources of Cosmos DB accounts.
		return true
	default:
		return false
	}
}

func sliceHasPrefix(slice, subSlice []string) bool {
	if len(subSlice) > len(slice) {
		return false
	}

	for i := range subSlice {
		if slice[i] != subSlice[i] {
			return false
		}
	}

	return true
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
