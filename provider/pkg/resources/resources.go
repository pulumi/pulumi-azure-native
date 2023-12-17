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
	// If the property is represented as a map of string to nothing, the SDK represents it as a list of strings.
	IsStringSet bool `json:"isStringSet,omitempty"`
	// Default is the default value for the parameter, if any
	Default interface{} `json:"default,omitempty"`
	// Some resources are nested beneath other resources, but are also made available as a property on the parent.
	// When updating the parent, if the list of nested resources is not specified as a property on the parent, the child resources will be deleted.
	// When refreshing the parent, the list of nested resources will be populated from the current state - resulting in an unwanted diff.
	// We work around this by setting MaintainSubResourceIfUnset to true for the properties containing a list of nested resources.
	// When updating or refreshing the parent, if the list of nested resources was not originally specified inline, we'll maintain the existing sub-resources transparently.
	MaintainSubResourceIfUnset bool `json:"maintainSubResourceIfUnset,omitempty"`
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
	// Containers within the request body that are required even if no properties are set within it.
	RequiredContainers [][]string `json:"requiredContainers,omitempty"`
	// Default values to be used when the property is removed or in importing. Must be top-level properties.
	DefaultProperties map[string]interface{} `json:"defaultProperties,omitempty"`
	// Some resources are nested beneath other resources, but are also made available as a property on the parent.
	// When updating the parent, if the list of nested resources is not specified as a property on the parent, the child resources will be deleted.
	// When refreshing the parent, the list of nested resources will be populated from the current state - resulting in an unwanted diff.
	// We work around this by adding these properties containing a list of nested resources to SubResourcesToMaintainIfUnset.
	// When updating or refreshing the parent, if the list of nested resources was not originally specified inline, we'll maintain the existing sub-resources transparently.
	// The values are property paths pointing to the property, like ["property"] (top-level) or ["container", "property"].
	SubResourcesToMaintainIfUnset [][]string `json:"subResourcesToMaintainIfUnset,omitempty"`
}

type TypeLookupFunc func(ref string) (*AzureAPIType, bool, error)

func TraverseProperties(props map[string]AzureAPIProperty, lookupType TypeLookupFunc, f func(propName string, prop AzureAPIProperty, path []string)) {
	// Start the traversal with an empty path and an empty set of seen types for cycle detection.
	traverseProperties(props, lookupType, []string{}, map[string]struct{}{}, f)
}

func traverseProperties(props map[string]AzureAPIProperty, lookupType TypeLookupFunc, path []string, seen map[string]struct{}, f func(propName string, prop AzureAPIProperty, path []string)) {
	for propName, prop := range props {
		if prop.Ref != "" {
			refType, ok, err := lookupType(prop.Ref)
			if !ok || err != nil {
				fmt.Printf("Cannot traverse properties of %s: failed to find ref %s: %v\n", propName, prop.Ref, err)
				continue
			}
			if _, visited := seen[prop.Ref]; !visited {
				seen[prop.Ref] = struct{}{}
				traverseProperties(refType.Properties, lookupType, append(path, propName), seen, f)
			}
		}

		f(propName, prop, path)
	}
}

func (res *AzureAPIResource) CollectSubResourceToMaintainIfUnset(typeLookup TypeLookupFunc) {
	body := res.BodyParameter()
	if body == nil {
		return
	}

	result := [][]string{}
	TraverseProperties(
		body.Body.Properties,
		typeLookup,
		func(propName string, prop AzureAPIProperty, path []string) {
			if prop.MaintainSubResourceIfUnset {
				// make a copy of path since the original might be passed to other callbacks
				pathToProperty := append([]string{}, path...)
				result = append(result, append(pathToProperty, propName))
			}
		})

	res.SubResourcesToMaintainIfUnset = result
}

func (res *AzureAPIResource) LookupProperty(key string) (AzureAPIProperty, bool) {
	if body := res.BodyParameter(); body != nil {
		if prop, ok := body.Body.Properties[key]; ok {
			return prop, true
		}
	}
	return AzureAPIProperty{}, false
}

func (res *AzureAPIResource) BodyParameter() *AzureAPIParameter {
	for _, param := range res.PutParameters {
		if param.Location == "body" || param.Body != nil {
			return &param
		}
	}
	return nil
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

type NameDisambiguation struct {
	// The path to the OpenAPI spec file that generates the ambiguous name.
	FileLocation string
	// The operation ID that is ambiguous.
	OperationID string
	// The path of the resource or invoke that is ambiguous.
	Path string
	// The name of the resource or invoke that is ambiguous.
	GeneratedName string
	// The API versions that the name is ambiguous between.
	DisambiguatedName string
}

// ResourceName constructs a name of a resource based on Get or List operation ID,
// e.g. "Managers_GetActivationKey" -> "ManagerActivationKey".
func ResourceName(operationID, path string) (string, *NameDisambiguation) {
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
			name = namePrefix
			break
		}
	}

	resourceName := name + subName

	var nameDisambiuation *NameDisambiguation

	// Special cases

	// Manual override to resolve ambiguity between public and private RecordSet.
	// See https://github.com/pulumi/pulumi-azure-native/issues/583.
	// To be removed with https://github.com/pulumi/pulumi-azure-native/issues/690.
	if resourceName == "RecordSet" && strings.Contains(path, "/providers/Microsoft.Network/privateDnsZones/") {
		newName := "PrivateRecordSet"
		nameDisambiuation = &NameDisambiguation{
			OperationID:       operationID,
			Path:              path,
			GeneratedName:     resourceName,
			DisambiguatedName: newName,
		}
		resourceName = newName
	}

	// Cognitive Services has global and per-account commitment plans with the same name.
	// The global ones are new, introduced in 2022-12-01, so we rename them.
	// TODO,tkappler The global plan still has the description "Cognitive Services account
	// commitment plan." - upstream issue?
	if resourceName == "CommitmentPlan" && strings.Contains(path, "/providers/Microsoft.CognitiveServices/commitmentPlans/") {
		newName := "SharedCommitmentPlan"
		nameDisambiuation = &NameDisambiguation{
			OperationID:       operationID,
			Path:              path,
			GeneratedName:     resourceName,
			DisambiguatedName: newName,
		}
		resourceName = newName
	}

	// Redis and RedisEnterprise are essentially distinct resources sharing the Microsoft.Cache
	// namespace. It works out ok because each API version has only one of them, and of the shared
	// types only PrivateEndpointConnection clashes.
	if resourceName == "PrivateEndpointConnection" && strings.Contains(path, "/providers/Microsoft.Cache/redisEnterprise/") {
		newName := "EnterprisePrivateEndpointConnection"
		nameDisambiuation = &NameDisambiguation{
			OperationID:       operationID,
			Path:              path,
			GeneratedName:     resourceName,
			DisambiguatedName: newName,
		}
		resourceName = newName
	}

	// Both are virtual network links, but the other side of the link is a different resource and
	// the links have different properties.
	// https://learn.microsoft.com/en-us/azure/dns/dns-private-resolver-overview#virtual-network-links
	// https://learn.microsoft.com/en-us/azure/dns/private-dns-virtual-network-links
	if resourceName == "VirtualNetworkLink" && strings.Contains(path, "/providers/Microsoft.Network/dnsForwardingRulesets/") {
		newName := "PrivateResolverVirtualNetworkLink"
		nameDisambiuation = &NameDisambiguation{
			OperationID:       operationID,
			Path:              path,
			GeneratedName:     resourceName,
			DisambiguatedName: newName,
		}
		resourceName = newName
	}

	// ServiceFabric introduced managed clusters in a new folder 'servicefabricmanagedclusters' but
	// in the same Microsoft.ServiceFabric namespace. We need to disambiguate some resource names.
	if strings.Contains(path, "ServiceFabric/managedclusters") &&
		(resourceName == "Application" ||
			resourceName == "ApplicationType" ||
			resourceName == "ApplicationTypeVersion" ||
			resourceName == "Service") {
		newName := "ManagedCluster" + resourceName
		nameDisambiuation = &NameDisambiguation{
			OperationID:       operationID,
			Path:              path,
			GeneratedName:     resourceName,
			DisambiguatedName: newName,
		}
		resourceName = newName
	}

	return resourceName, nameDisambiuation
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
