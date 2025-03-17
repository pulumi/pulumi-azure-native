// Copyright 2016-2020, Pulumi Corporation.

package resources

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/gedex/inflector"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
)

// SingleValueProperty is the name of the property that we insert into the schema for non-object type responses of invokes.
const SingleValueProperty = "value"

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

// ModuleName as used within the Pulumi provider e.g. aad
// This is often derived from the "Azure Resource Provider" (namespace) in the Azure spec.
// But sometimes is related to the top-level directory in the specs.
type ModuleName string

func (m ModuleName) Lowered() string {
	return strings.ToLower(string(m))
}

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
	ForceNew                            bool `json:"forceNew,omitempty"`
	ForceNewInferredFromReferencedTypes bool `json:"forceNewInferredFromReferencedTypes,omitempty"`
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
	// Optional. Properties that combined form a unique identifier for elements in this array.
	// Corresponds to the x-ms-identifiers extension in the Azure spec.
	ArrayIdentifiers []string `json:"arrayIdentifiers,omitempty"`
}

// AzureAPIType represents the shape of an object property.
type AzureAPIType struct {
	Properties         map[string]AzureAPIProperty `json:"properties,omitempty"`
	RequiredProperties []string                    `json:"required,omitempty"`
}

// AzureAPIResource is a resource in Azure REST API.
type AzureAPIResource struct {
	// API version in "2020-10-01" format.
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
	// Extra query parameters and their values to include in read (GET, HEAD) requests.
	ReadQueryParams map[string]any `json:"readUrlParams,omitempty"`
	// By default, we populate the `location` property of every resource to the location of its resource
	// group or the configured value. AutoLocationDisabled can override this default behavior.
	AutoLocationDisabled bool `json:"autoLocationDisabled,omitempty"`
	// Containers within the request body that are required even if no properties are set within it.
	RequiredContainers [][]string `json:"requiredContainers,omitempty"`
	// Default values to be used when the property is removed or in importing. Must be top-level properties.
	DefaultProperties map[string]interface{} `json:"defaultProperties,omitempty"`
}

type ResourceLookupFunc func(resourceType string) (AzureAPIResource, bool, error)

type TypeLookupFunc func(ref string) (*AzureAPIType, bool, error)

func TraverseProperties(props map[string]AzureAPIProperty, lookupType TypeLookupFunc, includeContainers bool, f func(propName string, prop AzureAPIProperty, path []string)) {
	// Start the traversal with an empty path and an empty set of seen types for cycle detection.
	traverseProperties(props, lookupType, includeContainers, []string{}, map[string]struct{}{}, f)
}

func traverseProperties(props map[string]AzureAPIProperty,
	lookupType TypeLookupFunc,
	includeContainers bool,
	path []string,
	seen map[string]struct{},
	f func(propName string, prop AzureAPIProperty, path []string)) {
	// Sort the properties to ensure stable output
	propNames := make([]string, 0, len(props))
	for propName := range props {
		propNames = append(propNames, propName)
	}
	sort.StringSlice(propNames).Sort()

	for _, propName := range propNames {
		prop := props[propName]
		pathCopy := append([]string{}, path...)

		if includeContainers {
			pathCopy = append(pathCopy, prop.Containers...)
		}

		ref := prop.Ref
		if ref == "" && prop.Items != nil {
			ref = prop.Items.Ref
		}
		if ref != "" && strings.HasPrefix(ref, "#/types/azure-native") {
			refType, ok, err := lookupType(ref)
			if !ok || err != nil {
				fmt.Printf("Cannot traverse properties of %s: failed to find ref %s: %v\n", propName, ref, err)
				continue
			}
			if _, visited := seen[ref]; !visited {
				seen[ref] = struct{}{}
				nextPath := append(pathCopy, propName)
				traverseProperties(refType.Properties, lookupType, includeContainers, nextPath, seen, f)
			}
		}

		f(propName, prop, pathCopy)
	}
}

func (res *AzureAPIResource) PathsToSubResourcePropertiesToMaintain(includeContainers bool, typeLookup TypeLookupFunc) [][]string {
	result := [][]string{}

	body, hasBody := res.BodyParameter()
	if !hasBody {
		return result
	}

	TraverseProperties(
		body.Body.Properties,
		typeLookup,
		includeContainers,
		func(propName string, prop AzureAPIProperty, path []string) {
			if prop.MaintainSubResourceIfUnset {
				// make a copy of path since the original might be passed to other callbacks
				pathToProperty := append([]string{}, path...)
				pathToProperty = append(pathToProperty, propName)
				result = append(result, pathToProperty)
			}
		})

	return result
}

func (res *AzureAPIResource) LookupProperty(key string) (AzureAPIProperty, bool) {
	if body, ok := res.BodyParameter(); ok {
		if prop, ok := body.Body.Properties[key]; ok {
			return prop, true
		}
	}
	return AzureAPIProperty{}, false
}

// BodyParameter returns the body parameter of a resource's PUT parameters, if any.
func (res *AzureAPIResource) BodyParameter() (*AzureAPIParameter, bool) {
	for _, param := range res.PutParameters {
		if param.Location == "body" || param.Body != nil {
			return &param, true
		}
	}
	return nil, false
}

// AzureAPIExample provides a pointer to examples relevant to a resource from the Azure REST API spec.
type AzureAPIExample struct {
	Description string `json:"description"`
	Location    string `json:"location"`
}

// AzureAPIInvoke is an invocation target (a function) in Azure REST API.
type AzureAPIInvoke struct {
	APIVersion string `json:"apiVersion"`
	Path       string `json:"path"`
	// IsResourceGetter is true if the invoke returns a resource type.
	IsResourceGetter bool                        `json:"isResourceGetter"`
	GetParameters    []AzureAPIParameter         `json:"GET"`
	PostParameters   []AzureAPIParameter         `json:"POST"`
	Response         map[string]AzureAPIProperty `json:"response"`
}

// PartialAzureAPIMetadata is a collection of resources and functions in the Azure REST API surface, supplementing the
// Pulumi schema. It saves memory by using PartialMap internally and unmarshaling only the parts of the spec that are
// needed. Since PartialMap implements MapLike, PartialAzureAPIMetadata can trivially be converted to APIMetadata.
type PartialAzureAPIMetadata struct {
	Types     PartialMap[AzureAPIType]     `json:"types"`
	Resources PartialMap[AzureAPIResource] `json:"resources"`
	Invokes   PartialMap[AzureAPIInvoke]   `json:"invokes"`
}

// APIMetadata is a collection of resources and functions in the Azure REST API surface, supplementing the Pulumi schema.
type APIMetadata struct {
	Types     MapLike[AzureAPIType]     `json:"types"`
	Resources MapLike[AzureAPIResource] `json:"resources"`
	Invokes   MapLike[AzureAPIInvoke]   `json:"invokes"`
}

// UnmarshalJSON implements custom unmarshaling for APIMetadata
func (m *APIMetadata) UnmarshalJSON(data []byte) error {
	p := PartialAzureAPIMetadata{}
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}

	m.Invokes = &p.Invokes
	m.Resources = &p.Resources
	m.Types = &p.Types
	return nil
}

// AzureAPIMetadata is a collection of resources and functions in the Azure REST API surface, supplementing the Pulumi schema.
type AzureAPIMetadata struct {
	Types     map[string]AzureAPIType     `json:"types"`
	Resources map[string]AzureAPIResource `json:"resources"`
	Invokes   map[string]AzureAPIInvoke   `json:"invokes"`
}

// Some resource providers changed capitalization between API versions, but we should map them to the
// same spelling so that folder names and namespaces are consistent. The map below provides such
// canonical names based on which names seems to be used prominently as of 2020.
var wellKnownModuleNames = map[string]string{
	"aad":                          "Aad",
	"aadiam":                       "AadIam",
	"dbformariadb":                 "DBforMariaDB",
	"dbformysql":                   "DBforMySQL",
	"dbforpostgresql":              "DBforPostgreSQL",
	"powerbidedicated":             "PowerBIDedicated",
	"servicefabricmanagedclusters": "ServiceFabric", // https://github.com/Azure/azure-rest-api-specs/issues/15867
	"visualstudio":                 "VisualStudio",
}

type ModuleNaming struct {
	ResolvedName           ModuleName
	SpecFolderName         string
	NamespaceWithoutPrefix string
	RpNamespace            string
}

// GetModuleName returns a module name given Open API spec file and resource's API URI.
// Returns the module name, optional old module name, or error.
func GetModuleName(majorVersion uint64, filePath, apiUri string) (ModuleNaming, error) {
	// We extract the module name from two sources:
	// - from the folder name of the Open API spec
	// - from the URI of the API endpoint (we take the last namespace in the URI)
	specFolderName, specFilePath := getSpecFolderNameAndFilePath(filePath)
	namespaceWithoutPrefixFromSpecFilePath := findNamespaceWithoutPrefixFromPath(filePath, "")
	namespace, err := findSpecNamespace(specFilePath)
	if err != nil {
		return ModuleNaming{}, err
	}
	// Sanity check the new and old methods return consistent results for now.
	if namespaceWithoutPrefix := removeNamespacePrefixAndNormalise(namespace); namespaceWithoutPrefix != namespaceWithoutPrefixFromSpecFilePath {
		return ModuleNaming{}, fmt.Errorf("resolved namespace mismatch: old: %s, new: %s", namespaceWithoutPrefixFromSpecFilePath, namespaceWithoutPrefix)
	}
	// Start with extracting the namespace from the folder path. If the folder name is explicitly listed,
	// use it as the module name. This is the new style we use for newer resources after 1.0. Older
	// resources to be migrated as part of https://github.com/pulumi/pulumi-azure-native/issues/690.
	if override, hasOverride := getNameOverride(majorVersion, specFolderName); hasOverride {
		return ModuleNaming{
			ResolvedName:           override,
			SpecFolderName:         specFolderName,
			NamespaceWithoutPrefix: namespaceWithoutPrefixFromSpecFilePath,
			RpNamespace:            namespace,
		}, nil
	}
	// We proceed with the endpoint if both module values match. This way, we avoid flukes and
	// declarations outside of the current API provider.
	namespaceWithoutPrefixFromResourceUrl := findNamespaceWithoutPrefixFromPath(apiUri, "Resources")
	if !strings.EqualFold(namespaceWithoutPrefixFromSpecFilePath, namespaceWithoutPrefixFromResourceUrl) {
		return ModuleNaming{}, fmt.Errorf("resolved module name mismatch: file: %s, uri: %s", namespaceWithoutPrefixFromSpecFilePath, namespaceWithoutPrefixFromResourceUrl)
	}
	// Ultimately we use the module name from the spec file path.
	return ModuleNaming{
		ResolvedName:           ModuleName(namespaceWithoutPrefixFromSpecFilePath),
		SpecFolderName:         specFolderName,
		NamespaceWithoutPrefix: namespaceWithoutPrefixFromSpecFilePath,
		RpNamespace:            namespace,
	}, nil
}

var v2ModulesNamedByFolder = map[string]ModuleName{
	"videoanalyzer": "VideoAnalyzer",
	"webpubsub":     "WebPubSub",
}

var v3ModulesNamedByFolder = map[string]ModuleName{
	// "Network":
	"dns":            "Dns",
	"dnsresolver":    "DnsResolver",
	"frontdoor":      "FrontDoor",
	"privatedns":     "PrivateDns",
	"trafficmanager": "TrafficManager",
	// Cache:
	"redis":           "Redis",
	"redisenterprise": "RedisEnterprise",
	// Devices:
	"iothub":                     "IoTHub",
	"deviceprovisioningservices": "DeviceProvisioningServices",
	// "DocumentDB":
	"cosmos-db":    "CosmosDB",
	"mongocluster": "MongoCluster",
	// Insights:
	"monitor":             "Monitor",
	"applicationinsights": "ApplicationInsights",
}

// getNameOverride returns a name override for a given folder name, and non-prefixed namespace.
// The second return value is true if an override is found.
func getNameOverride(majorVersion uint64, specFolderName string) (ModuleName, bool) {
	// For the cases below, we use folder (SDK) moduleName for module names instead of the ARM moduleName.
	if moduleName, ok := v2ModulesNamedByFolder[specFolderName]; ok {
		return moduleName, true
	}
	// Disable additional rules for v2 and below.
	// TODO: Remove after v3 release.
	if majorVersion < 3 {
		return "", false
	}
	// New rules for v3 and above which include aliases back to the original namespace.
	if moduleName, ok := v3ModulesNamedByFolder[specFolderName]; ok {
		return moduleName, true
	}
	return "", false
}

// Identify the first segment of the path that contains a namespace.
func findSpecNamespace(path string) (string, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return "", fmt.Errorf("path did not contain at least 3 parts: %s", path)
	}

	for _, part := range parts {
		// It must contain a dot to indicate it's a namespace,
		// but not contain any curly braces which would indicate it's a parameter (used in Microsoft.EventGrid).
		if strings.Contains(part, ".") && !strings.ContainsAny(part, "{}") {
			return part, nil
		}
	}
	return "", fmt.Errorf("no namespace found in path: %s", path)
}

// Remove the namespace prefix, normalise the casing and account for renames.
func removeNamespacePrefixAndNormalise(namespace string) string {
	// Some namespaces can be nested such as Microsoft.Web.Admin, though it's rare.
	_, after, found := strings.Cut(namespace, ".")
	if !found {
		return namespace
	}
	if knownName, ok := wellKnownModuleNames[strings.ToLower(after)]; ok {
		return knownName
	}
	return strings.Title(after)
}

// Locate namespaces based on well known prefixes (Microsoft., microsoft., PaloAltoNetworks.)
// and return the namespace without the prefix.
// If the path is long enough, but no prefix is found, return the default value.
func findNamespaceWithoutPrefixFromPath(path, defaultValue string) string {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return ""
	}

	for i := len(parts) - 2; i >= 0; i-- {
		part := parts[i]
		for _, prefix := range []string{"Microsoft.", "microsoft.", "PaloAltoNetworks."} {
			if strings.HasPrefix(part, prefix) {
				name := strings.Title(strings.TrimPrefix(part, prefix))
				if knownName, ok := wellKnownModuleNames[strings.ToLower(name)]; ok {
					return knownName
				}
				return name
			}
		}
	}

	return defaultValue
}

var folderModulePattern = regexp.MustCompile(`.*/specification/([a-zA-Z0-9-]+)/resource-manager/(.*)`)

func getSpecFolderNameAndFilePath(path string) (string, string) {
	// Note that this returns last match, e.g. the following path matches `C`, `D`:
	// specification/A/resource-manager/B/specification/C/resource-manager/D
	subMatches := folderModulePattern.FindStringSubmatch(path)
	if len(subMatches) > 2 {
		moduleAlias := subMatches[1]
		filePath := subMatches[2]
		return moduleAlias, filePath
	}
	return "", ""
}

var verbReplacer = strings.NewReplacer(
	"GetProperties", "",
	"Get", "",
	"getByName", "",
	"get", "",
	"ListByResourceName", "",
	"List", "",
	"list", "",
	"CheckEntityExists", "",

	// Specifically for DesktopVirtualization getHostPoolRegistrationToken.
	// In v3 we should replace all instances of "retrieve" #3329.
	"RetrieveRegistration", "Registration",
)

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
	return createResourceName(operationID, path, version.GetVersion().Major)
}

func createResourceName(operationID, path string, majorVersion uint64) (string, *NameDisambiguation) {
	if majorVersion >= 3 {
		// Uppercase the first character of operationID
		r := []rune(operationID)
		r[0] = unicode.ToUpper(r[0])
		operationID = string(r)
	}

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

	return handleResourceNameSpecialCases(resourceName, operationID, path, majorVersion)
}

// handleResourceNameSpecialCases returns a modified resource name if the resource is mapped to multiple API paths. For
// instance, the deprecated "single server" resources in `dbformysql` and `dbforpostgresql` are renamed to `SingleServer`.
func handleResourceNameSpecialCases(resourceName, operationID, path string, majorVersion uint64) (string, *NameDisambiguation) {
	newName := func(prefix, newName string) (string, *NameDisambiguation) {
		// Don't generate "RedisRedis"
		if prefix != "" && !strings.HasPrefix(resourceName, prefix) {
			newName = prefix + newName
		}

		var nameDisambiguation *NameDisambiguation
		if newName != resourceName {
			nameDisambiguation = &NameDisambiguation{
				OperationID:       operationID,
				Path:              path,
				GeneratedName:     resourceName,
				DisambiguatedName: newName,
			}
		}

		return newName, nameDisambiguation
	}

	var nameDisambiguation *NameDisambiguation
	lowerPath := strings.ToLower(path)

	// Microsoft.CognitiveServices has global and per-account commitment plans with the same name.
	// The global ones are new, introduced in 2022-12-01, so we rename them.
	// The global plan still has the description "Cognitive Services account commitment plan." - upstream issue?
	if resourceName == "CommitmentPlan" && strings.Contains(path, "/providers/Microsoft.CognitiveServices/commitmentPlans/") {
		resourceName, nameDisambiguation = newName("", "SharedCommitmentPlan")
	}

	// Microsoft.NetApp
	if strings.Contains(lowerPath, "/providers/microsoft.netapp/backupvaults/") {
		resourceName, nameDisambiguation = newName("BackupVault", resourceName)
	}

	// Microsoft.Network
	// Manual override to resolve ambiguity between public and private RecordSet.
	// See https://github.com/pulumi/pulumi-azure-native/issues/583.
	// To be removed with https://github.com/pulumi/pulumi-azure-native/issues/690.
	if resourceName == "RecordSet" && strings.Contains(path, "/providers/Microsoft.Network/privateDnsZones/") {
		resourceName, nameDisambiguation = newName("", "PrivateRecordSet")
	}

	// Microsoft.Network
	// Both are virtual network links, but the other side of the link is a different resource and
	// the links have different properties.
	// https://learn.microsoft.com/en-us/azure/dns/dns-private-resolver-overview#virtual-network-links
	// https://learn.microsoft.com/en-us/azure/dns/private-dns-virtual-network-links
	if resourceName == "VirtualNetworkLink" && strings.Contains(path, "/providers/Microsoft.Network/dnsForwardingRulesets/") {
		resourceName, nameDisambiguation = newName("", "PrivateResolverVirtualNetworkLink")
	}

	// ServiceFabric introduced managed clusters in a new folder 'servicefabricmanagedclusters' but
	// in the same Microsoft.ServiceFabric namespace. We need to disambiguate some resource names.
	if strings.Contains(path, "ServiceFabric/managedclusters") &&
		(resourceName == "Application" ||
			resourceName == "ApplicationType" ||
			resourceName == "ApplicationTypeVersion" ||
			resourceName == "Service") {
		resourceName, nameDisambiguation = newName("ManagedCluster", resourceName)
	}

	if majorVersion < 3 {
		if resourceName == "PrivateEndpointConnection" && strings.Contains(path, "/providers/Microsoft.Cache/redisEnterprise/") {
			resourceName, nameDisambiguation = newName("", "EnterprisePrivateEndpointConnection")
		}
	}

	if majorVersion >= 3 {
		// Microsoft.Cache
		if strings.Contains(lowerPath, "/providers/microsoft.cache/redis/") {
			// resourceName, nameDisambiguation = newName("Redis", resourceName)
		} else if strings.Contains(lowerPath, "/providers/microsoft.cache/redisenterprise/") {
			if resourceName != "RedisEnterprise" {
				resourceName, nameDisambiguation = newName("Enterprise", resourceName)
			}
		}

		// Microsoft.DBforMySQL
		if strings.Contains(lowerPath, "/providers/microsoft.dbformysql/servers/") {
			if resourceName == "Server" {
				resourceName, nameDisambiguation = newName("", "SingleServer")
			} else {
				resourceName, nameDisambiguation = newName("SingleServer", resourceName)
			}
		}

		// Microsoft.DBforPostgreSQL
		if strings.Contains(lowerPath, "/providers/microsoft.dbforpostgresql/servers/") {
			if resourceName == "Server" {
				resourceName, nameDisambiguation = newName("", "SingleServer")
			} else {
				resourceName, nameDisambiguation = newName("SingleServer", resourceName)
			}
		} else if strings.Contains(lowerPath, "/providers/microsoft.dbforpostgresql/servergroupsv2/") {
			resourceName, nameDisambiguation = newName("ServerGroup", resourceName)
		}

		// Microsoft.DocumentDB
		if strings.Contains(lowerPath, "/providers/microsoft.documentdb/mongoclusters/") {
			resourceName, nameDisambiguation = newName("MongoCluster", resourceName)
		}

		// Microsoft.HDInsight
		if strings.Contains(lowerPath, "/providers/microsoft.hdinsight/clusterpools/") {
			resourceName, nameDisambiguation = newName("ClusterPool", resourceName)
		}

		// Microsoft.HybridContainerService
		if strings.Contains(lowerPath, "/providers/microsoft.hybridcontainerservice/provisionedclusterinstances/") {
			if !strings.Contains(resourceName, "ClusterInstance") {
				resourceName, nameDisambiguation = newName("ClusterInstance", resourceName)
			}
		}

		// Microsoft.LabServices
		if strings.Contains(lowerPath, "/providers/microsoft.labservices/labaccounts/") {
			resourceName, nameDisambiguation = newName("LabAccount", resourceName)
		}

		// Microsoft.Migrate
		if strings.Contains(lowerPath, "/providers/microsoft.migrate/assessmentprojects/") {
			resourceName, nameDisambiguation = newName("AssessmentProjects", resourceName)
		}

		// Microsoft.MobileNetwork
		if strings.Contains(lowerPath, "/providers/microsoft.mobilenetwork/simgroups/") {
			resourceName, nameDisambiguation = newName("SimGroup", resourceName)
		}

		// Microsoft.NetApp
		if strings.Contains(lowerPath, "/providers/microsoft.netapp/netappaccounts/") && strings.Contains(lowerPath, "/capacitypools/") {
			resourceName, nameDisambiguation = newName("CapacityPool", resourceName)
		}
	}

	return resourceName, nameDisambiguation
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
