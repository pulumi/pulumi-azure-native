// Copyright 2016-2021, Pulumi Corporation.
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

package gen

import (
	"fmt"
	"strings"

	"github.com/go-openapi/spec"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func (m *moduleGenerator) genTypeSpec(propertyName string, schema *spec.Schema, context *openapi.ReferenceContext, isOutput bool) (*pschema.TypeSpec, error) {
	resolvedSchema, err := context.ResolveSchema(schema)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve schema: %w", err)
	}

	// Type specification specifies either a primitive type (e.g. 'string') or a reference to a separately defined
	// strongly-typed object. Either `primitiveTypeName` or `referencedTypeName` should be filled, but not both.
	var primitiveTypeName string
	if len(resolvedSchema.Type) > 0 {
		primitiveTypeName = resolvedSchema.Type[0]
		if primitiveTypeName == "integer" && resolvedSchema.Format == "int64" {
			// Pulumi schema's integer is limited to 32 bits. Model a 64-bit integer as a number.
			primitiveTypeName = "number"
		}
	}

	// If this is an enum and it's an input type, generate the enum type definition.
	enumExtension, isEnum := resolvedSchema.Extensions[extensionEnum].(map[string]interface{})
	if !isOutput && isEnum && resolvedSchema.Enum != nil &&
		// There are some boolean-, integer-, and number- based definitions but they aren't allowed by the Azure specs.
		// Ignore both to preserve over-the-wire format even if they say it's an enum to be modelled as a string.
		primitiveTypeName != "boolean" && primitiveTypeName != "integer" && primitiveTypeName != "number" {
		return m.genEnumType(schema, context, enumExtension, propertyName)
	}

	switch {
	case len(resolvedSchema.Properties) > 0 || len(resolvedSchema.AllOf) > 0:
		ptr := schema.Ref.GetPointer()
		var tok string
		if ptr != nil && !ptr.IsEmpty() {
			tok = m.typeName(resolvedSchema.ReferenceContext, isOutput)
		} else {
			// Inline properties have no type in the Open API schema, so we use parent type's name + property name.
			tok = m.typeName(context, isOutput) + m.inlineTypeName(context, propertyName)
		}

		// If an object type is referenced, add its definition to the type map.
		discriminatedType, ok, err := m.genDiscriminatedType(resolvedSchema, isOutput)
		if err != nil {
			return nil, fmt.Errorf("failed to generate discriminated type: %w", err)
		}
		if ok {
			return discriminatedType, nil
		}

		tok = m.caseSensitiveTypes.normalizeTokenCase(tok)
		if _, ok := m.visitedTypes[tok]; !ok {
			m.visitedTypes[tok] = true

			props, err := m.genProperties(resolvedSchema, genPropertiesVariant{
				isOutput:   isOutput,
				isType:     true,
				isResponse: false,
			})
			if err != nil {
				return nil, err
			}

			// Don't generate a type definition for a typed object with zero properties.
			if len(props.specs) == 0 {
				delete(m.visitedTypes, tok)
				return nil, nil
			}

			// Describe special syntax for relative self-references of sub-resources.
			if strings.HasSuffix(tok, ":SubResource") {
				if prop, ok := props.specs["id"]; ok {
					prop.Description = `Sub-resource ID. Both absolute resource ID and a relative resource ID are accepted.
An absolute ID starts with /subscriptions/ and contains the entire ID of the parent resource and the ID of the sub-resource in the end.
A relative ID replaces the ID of the parent resource with a token '$self', followed by the sub-resource ID itself.
Example of a relative ID: $self/frontEndConfigurations/my-frontend.`
					props.specs["id"] = prop
				}
			}

			// incompatible type "azure-native:network:SecurityRule" for resource "LoadBalancer" ("azure-native:network:LoadBalancer"): required properties do not match: only required in A: priority
			if tok == "azure-native:network:SecurityRule" {
				props.requiredProperties.Delete("priority")
				props.requiredSpecs.Delete("priority")
			}

			// incompatible type "azure-native:containerinstance:Container" for resource "ContainerGroupProfile"
			// ("azure-native:containerinstance:ContainerGroupProfile"): required properties do not match: only
			// required in B: image,resources
			if tok == "azure-native:containerinstance:Container" {
				props.requiredProperties.Delete("image")
				props.requiredSpecs.Delete("image")
				props.requiredProperties.Delete("resources")
				props.requiredSpecs.Delete("resources")
			}

			spec := pschema.ComplexTypeSpec{
				ObjectTypeSpec: pschema.ObjectTypeSpec{
					Description: resolvedSchema.Description,
					Type:        "object",
					Properties:  props.specs,
					Required:    props.requiredSpecs.SortedValues(),
				},
			}

			if existing, has := m.pkg.Types[tok]; has {
				merged, err := mergeTypes(spec, existing, isOutput)
				if err != nil {
					return nil, fmt.Errorf("incompatible type %q for resource %q (%q): %w", tok, m.resourceName,
						m.resourceToken, err)
				}
				spec = *merged
			}

			m.pkg.Types[tok] = spec

			m.metadata.Types[tok] = resources.AzureAPIType{
				Properties:         props.properties,
				RequiredProperties: props.requiredProperties.SortedValues(),
			}
		}
		return &pschema.TypeSpec{
			Type: "object",
			Ref:  fmt.Sprintf("#/types/%s", tok),
		}, nil

	case len(resolvedSchema.Enum) > 0:
		if primitiveTypeName == "" || primitiveTypeName == "object" {
			// Default Enum properties to strings if the type isn't specified.
			primitiveTypeName = "string"
		}
		return &pschema.TypeSpec{Type: primitiveTypeName}, nil

	case resolvedSchema.Items != nil && resolvedSchema.Items.Schema != nil:
		// Special case for VirtualMachineScaleSet.diskControllerType. The property is erroneously
		// defined as a string _and_ array in the spec. Primitive string is correct. #2639,
		// upstream issue Azure/azure-rest-api-specs/issues/25202.
		// Note: as of 2024-08, going with the primitive type over the array for all such
		// properties causes a large diff in the azurestackhci SDK, so we opted not to do that.
		if primitiveTypeName != "array" && m.resourceName == "VirtualMachineScaleSet" && propertyName == "diskControllerType" {
			return &pschema.TypeSpec{Type: primitiveTypeName}, nil
		}

		// Resolve the element type for array types.
		property := resolvedSchema.Properties[propertyName]
		resolvedProperty, err := resolvedSchema.ResolveSchema(&property)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve schema for array element: %w", err)
		}
		itemsSpec, _, err := m.genProperty(propertyName, resolvedSchema.Items.Schema, context, resolvedProperty, genPropertiesVariant{
			isOutput: isOutput,
			isType:   true,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to generate type spec for array element: %w", err)
		}

		// Don't generate a type definition for a typed array with empty item type.
		if itemsSpec == nil {
			return nil, nil
		}

		return &pschema.TypeSpec{
			Type:  "array",
			Items: &itemsSpec.TypeSpec,
		}, nil

	case resolvedSchema.AdditionalProperties != nil && resolvedSchema.AdditionalProperties.Schema != nil:
		// Define the type of maps (untyped objects).
		additionalProperties, err := m.genTypeSpec(propertyName, resolvedSchema.AdditionalProperties.Schema, resolvedSchema.ReferenceContext, isOutput)
		if err != nil {
			return nil, fmt.Errorf("failed to generate type spec for additional properties: %w", err)
		}

		return &pschema.TypeSpec{
			Type:                 "object",
			AdditionalProperties: additionalProperties,
		}, nil

	case primitiveTypeName == "" || primitiveTypeName == "object":
		// Open API v2:
		// > A schema without a type matches any data type – numbers, strings, objects, and so on.
		// Azure uses a 'naked' object type for the same purpose: to specify 'any' type.
		return &pschema.TypeSpec{
			Ref: convert.TypeAny,
		}, nil

	default:
		// Primitive type ('string', 'integer', etc.)
		return &pschema.TypeSpec{Type: primitiveTypeName}, nil
	}
}

// inlineTypeName returns a type name suffix to be used as a type name for properties defined inline in the spec.
// It defaults to the TitleCased property name but it also keeps a map of potential collisions. If a collision occurs,
// (i.e. an inline property `foo` has inline properties and one of them is also `foo`, that can also be down several
// levels), then the property name is duplicated (to get `FooFoo` in this example).
func (m *moduleGenerator) inlineTypeName(ctx *openapi.ReferenceContext, propertyName string) string {
	result := strings.Title(propertyName)
	if ex, ok := m.inlineTypes[ctx]; ok {
		for {
			if !ex.Has(result) {
				break
			}
			result += strings.Title(propertyName)
		}
	} else {
		m.inlineTypes[ctx] = codegen.NewStringSet()
	}
	m.inlineTypes[ctx].Add(result)
	return result
}

// genEnumType generates the enum type. The propertyName serves as a fallback for the enum's name
// if none is given via the `name` attribute.
func (m *moduleGenerator) genEnumType(schema *spec.Schema, context *openapi.ReferenceContext,
	enumExtension map[string]interface{}, propertyName string,
) (*pschema.TypeSpec, error) {
	resolvedSchema, err := context.ResolveSchema(schema)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve schema for enum: %w", err)
	}

	description, err := getPropertyDescription(schema, context, false /* maintainSubResourceIfUnset */)
	if err != nil {
		return nil, fmt.Errorf("failed to get property description: %w", err)
	}

	var enumName string
	if name, ok := enumExtension["name"].(string); ok {
		enumName = name
	} else if propertyName != "" {
		enumName = propertyName
	} else {
		return nil, fmt.Errorf("name key missing from enum metadata and no property name available")
	}
	enumName = m.typeNameOverride(ToUpperCamel(enumName))

	tok := fmt.Sprintf("%s:%s:%s", m.pkg.Name, m.module, enumName)

	enumSpec := &pschema.ComplexTypeSpec{
		Enum: []pschema.EnumValueSpec{},
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: description,
			Type:        "string", // This provider only has string enums
		},
	}

	// Ignore additional enum values that only vary by case as this isn't supported by Pulumi.
	enumExists := func(enumVal pschema.EnumValueSpec) bool {
		for _, existing := range enumSpec.Enum {
			if strings.EqualFold(existing.Value.(string), enumVal.Value.(string)) {
				return true
			}
		}
		return false
	}

	// We prefer the enumExtension if available as it also provides the name and description.
	if values, ok := enumExtension["values"].([]interface{}); ok {
		for _, val := range values {
			if val, ok := val.(map[string]interface{}); ok {
				enumVal := pschema.EnumValueSpec{
					Value: fmt.Sprintf("%v", val["value"]),
				}
				if name, ok := val["name"].(string); ok {
					enumVal.Name = name
				}
				if description, ok := val["description"].(string); ok {
					enumVal.Description = description
				}
				if !enumExists(enumVal) {
					enumSpec.Enum = append(enumSpec.Enum, enumVal)
				}
			}
		}
	} else {
		// Fall back to the enum values defined in the schema if enumExtensions not available.
		for _, val := range resolvedSchema.Enum {
			enumVal := pschema.EnumValueSpec{Value: fmt.Sprintf("%v", val)}
			// Override the name for the values for this Enum since it contains unfortunately
			// named values like `Input` and `Output` which collide especially for Go Codegen.
			if strings.HasPrefix(string(m.module), "datafactory") && enumName == "ScriptActivityParameterDirection" {
				enumVal.Name = fmt.Sprintf("Value%s", val)
			}
			if !enumExists(enumVal) {
				enumSpec.Enum = append(enumSpec.Enum, enumVal)
			}
		}
	}

	m.pkg.Types[tok] = *enumSpec

	referencedTypeName := fmt.Sprintf("#/types/%s", tok)

	modelAsString, ok := enumExtension["modelAsString"].(bool)
	if !ok || modelAsString {
		return &pschema.TypeSpec{
			OneOf: []pschema.TypeSpec{
				{Type: "string"},
				{Ref: referencedTypeName},
			},
		}, nil
	}

	return &pschema.TypeSpec{
		Ref: referencedTypeName,
	}, nil
}

// genDiscriminatedType generates polymorphic types (base type and subtypes) if the schema specifies a discriminator property.
// If no error occurs, the bool result indicates whether a discriminated (union) type is detected. If true, the TypeSpec
// result points to the specification of the union type.
func (m *moduleGenerator) genDiscriminatedType(resolvedSchema *openapi.Schema, isOutput bool) (*pschema.TypeSpec, bool, error) {
	if resolvedSchema.Discriminator == "" {
		return nil, false, nil
	}

	discriminator := resolvedSchema.Discriminator
	if _, ok := resolvedSchema.Properties[discriminator]; !ok {
		return nil, false, nil
	}

	prop := resolvedSchema.Properties[discriminator]
	if prop.ReadOnly && !isOutput {
		return nil, false, nil
	}

	var oneOf []pschema.TypeSpec
	mapping := map[string]string{}
	subtypes, err := resolvedSchema.FindSubtypes()
	if err != nil {
		return nil, false, err
	}
	for _, subtype := range subtypes {
		typ, err := m.genTypeSpec("", subtype, resolvedSchema.ReferenceContext, isOutput)
		if err != nil {
			return nil, false, err
		}
		oneOf = append(oneOf, *typ)
		subtypeSchema, err := resolvedSchema.ResolveSchema(subtype)
		if err != nil {
			return nil, false, err
		}
		discriminatorValue := getDiscriminatorValue(subtypeSchema)
		mapping[discriminatorValue] = typ.Ref
	}

	switch len(oneOf) {
	case 0:
		// Type specifies a discriminator but doesn't have actual subtypes. Ignore the discriminator in this case.
		return nil, false, nil
	case 1:
		// There is just one subtype specified: use it as a definite type.
		return &oneOf[0], true, nil
	default:
		sdkDiscriminator := discriminator
		if clientName, ok := prop.Extensions.GetString(extensionClientName); ok {
			sdkDiscriminator = clientName
		}
		sdkDiscriminator = ToLowerCamel(sdkDiscriminator)

		// Union type for two or more types.
		return &pschema.TypeSpec{
			OneOf: oneOf,
			Discriminator: &pschema.DiscriminatorSpec{
				PropertyName: sdkDiscriminator,
				Mapping:      mapping,
			},
		}, true, nil
	}
}

// typeNameOverrides is a manually-maintained map of alternative names when a type name represents two or more
// distinct types in the same module. This can happen if there are multiple Open API spec files
// for the same RP and version, and each of those files has its own definition of the type under the same name.
// That happens a lot (there are many files for several RPs) but most of the time the definitions are similar
// enough to treat them as same. The following map lists all exceptions from this rule.
// If a new mismatch is introduced in a newly published spec, our codegen will catch the difference
// (see compatibleTypes) and fail. We will have to extend the list below with a new exception to unblock codegen.
var typeNameOverrides = map[string]string{
	// Upstream commit 1a49c263cc introduced AwsConnector.EmrCluster(Summary) and .AwsEksClusterProperties which
	// both have ClusterStatus, but of different types.
	"AwsConnector.EmrCluster.ClusterStatus":        "EmrClusterStatus",
	"AwsConnector.EmrClusterSummary.ClusterStatus": "EmrClusterStatus",
	// Upstream commit 1a49c263cc introduced several resources that all have an "Endpoint" property. Some are
	// strings, the ones below are objects, all of them slightly different.
	"AwsConnector.DaxCluster.Endpoint":      "DaxClusterEndpoint",
	"AwsConnector.RdsDbCluster.Endpoint":    "RdsDbClusterEndpoint",
	"AwsConnector.RdsDbInstance.Endpoint":   "RdsDbInstanceEndpoint",
	"AwsConnector.RedshiftCluster.Endpoint": "RedshiftClusterEndpoint",
	// Upstream commit 1a49c263cc introduced TODO
	"AwsConnector.DaxCluster.NotificationConfiguration":                  "DaxClusterNotificationConfiguration",
	"AwsConnector.AutoScalingAutoScalingGroup.NotificationConfiguration": "AutoScalingGroupNotificationConfiguration",

	// SKU for Redis Enterprise is different from SKU for Redis. Keep them as separate types.
	"Cache.RedisEnterprise.Sku": "EnterpriseSku",
	// This one is not a disambiguation but a fix for a type name "String" that is not descriptive and leads to
	// generating invalid Java.
	"DatabaseWatcher.Target.String": "TargetCollectionStatus",
	// SingleServer is different from just "Server", see the exception in resources.go ResourceName().
	"DBforPostgreSQL.SingleServer.Sku": "SingleServerSku",
	// Devices RP comes from "deviceprovisioningservices" and "iothub" which are similar but slightly different.
	// In particular, the IP Filter Rule has more properties in the DPS version.
	"Devices.IotDpsResource.IpFilterRule": "TargetIpFilterRule",
	// See the exception for Microsoft.HDInsight in resources.go ResourceName().
	"HDInsight.ClusterPoolCluster.Sku":            "ClusterPoolSku",
	"HDInsight.ClusterPoolCluster.ComputeProfile": "ClusterPoolComputeProfile",
	"HDInsight.ClusterPoolCluster.SshProfile":     "ClusterPoolSshProfile",

	"HybridContainerService.ClusterInstanceAgentPool.Status": "AgentPoolProvisioningStatus",
	// Workbook vs. MyWorkbook types are slightly different. Probably, a bug in the spec, but we have to disambiguate.
	"Insights.MyWorkbook.ManagedIdentity":                  "MyManagedIdentity",
	"Insights.MyWorkbook.UserAssignedIdentities":           "MyUserAssignedIdentities",
	"ManagedNetworkFabric.NetworkFabric.OptionBProperties": "FabricOptionBProperties",
	// Experiment's endpoint is a much narrower type compared to endpoints in other network resources.
	"Network.Experiment.Endpoint": "ExperimentEndpoint",
	// These are all FrontDoor types. FrontDoor shares a bunch of type names with generate Network provider,
	// but defines them in its own way.
	"Network.Policy.ManagedRuleGroupOverride": "FrontDoorManagedRuleGroupOverride",
	"Network.Policy.ManagedRuleOverride":      "FrontDoorManagedRuleOverride",
	"Network.Policy.ManagedRuleSet":           "FrontDoorManagedRuleSet",
	"Network.Policy.MatchCondition":           "FrontDoorMatchCondition",
	"Network.Policy.MatchVariable":            "FrontDoorMatchVariable",
	"Network.Policy.PolicySettings":           "FrontDoorPolicySettings",
	// IpConfigurationResponse conflicts with IPConfigurationResponse used for existing networking resources.
	// This avoids Dotnet SDK failing to build since codegen currently elides the IPConfigurationResponse
	// output types since it assumes the existing one is sufficient.
	"Network.InboundEndpoint.IpConfiguration": "InboundEndpointIPConfiguration",
	// The following two types are read-only, while the same types in another spec are writable.
	"RecoveryServices.Vault.PrivateEndpointConnection":         "VaultPrivateEndpointConnection",
	"RecoveryServices.Vault.PrivateLinkServiceConnectionState": "VaultPrivateLinkServiceConnectionState",
	// Watchlist resources only appear in a preview spec and not in stable specs. Anyway, the shapes of their
	// types are slightly different from later specs, so we have to disambiguate for top-level resources.
	"SecurityInsights.Watchlist.UserInfo":     "WatchlistUserInfo",
	"SecurityInsights.WatchlistItem.UserInfo": "WatchlistUserInfo",

	// These generate enums which conflict with other resource types.
	"BillingBenefits.Discount.DiscountTypeProductFamily":            "DiscountProductFamily",
	"BillingBenefits.Discount.DiscountTypeProduct":                  "DiscountProduct",
	"BillingBenefits.Discount.DiscountTypeCustomPrice":              "DiscountCustomPrice",
	"BillingBenefits.Discount.DiscountTypeCustomPriceMultiCurrency": "DiscountCustomPriceMultiCurrency",
}

var typeNameOverridesV3 = map[string]string{
	"CosmosDB.MongoCluster.CreateMode": "MongoClusterCreateMode",

	// DocumentDB.MongoCluster from /mongocluster uses a different private endpoint connection.
	// The MongoCluster resource has the same name in v2 and v3, but the private endpoint connection was disambiguated in v3.
	"DocumentDB.MongoCluster.PrivateEndpointConnection": "MongoClusterPrivateEndpointConnection",

	"DBforMySQL.SingleServer.ServerVersion":                "SingleServerVersion",
	"DBforMySQL.SingleServer.SkuTier":                      "SingleServerSkuTier",
	"DBforPostgreSQL.ServerGroupCluster.AuthConfig":        "ServerGroupClusterAuthConfig",
	"DBforPostgreSQL.ServerGroupCluster.DataEncryption":    "ServerGroupClusterDataEncryption",
	"DBforPostgreSQL.ServerGroupCluster.MaintenanceWindow": "ServerGroupClusterMaintenanceWindow",
	"DBforPostgreSQL.SingleServer.CreateMode":              "SingleServerCreateMode",
	"DBforPostgreSQL.SingleServer.IdentityType":            "SingleServerIdentityProperties",
	"DBforPostgreSQL.SingleServer.PrincipalType":           "SingleServerPrincipalType",
	"DBforPostgreSQL.SingleServer.ServerVersion":           "SingleServerVersion",
	"DBforPostgreSQL.SingleServer.SkuTier":                 "SingleServerSkuTier",

	"Migrate.AssessmentProjectsAssessmentsOperation.AzureOfferCode":          "AssessmentProjectsAssessmentsOperationAzureOfferCode",
	"Migrate.AssessmentProjectsAssessmentsOperation.AzureVmFamily":           "AssessmentProjectsAssessmentsOperationAzureVmFamily",
	"Migrate.AssessmentProjectsAssessmentsOperation.AvsAssessmentProperties": "AssessmentProjectsAssessmentsOperationAvsAssessmentProperties",
}

func (m *moduleGenerator) typeNameOverride(typeName string) string {
	key := fmt.Sprintf("%s.%s.%s", m.moduleName, m.resourceName, typeName)
	if v, ok := typeNameOverrides[key]; ok {
		return v
	}
	if version.GetVersion().Major >= 3 {
		if v, ok := typeNameOverridesV3[key]; ok {
			return v
		}
	}
	return typeName
}

func (m *moduleGenerator) typeName(ctx *openapi.ReferenceContext, isOutput bool) string {
	suffix := ""
	if isOutput {
		suffix = "Response"
	}
	standardName := ToUpperCamel(MakeLegalIdentifier(ctx.ReferenceName))
	referenceName := m.typeNameOverride(standardName)
	return fmt.Sprintf("azure-native:%s:%s%s", m.module, referenceName, suffix)
}
