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
	"sort"
	"strconv"
	"strings"

	"github.com/go-openapi/spec"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// propertyBag keeps the schema and metadata properties for a single API type.
type propertyBag struct {
	description        string
	specs              map[string]pschema.PropertySpec
	properties         map[string]resources.AzureAPIProperty
	requiredSpecs      codegen.StringSet
	requiredProperties codegen.StringSet
}

func (m *moduleGenerator) genProperties(resolvedSchema *openapi.Schema, isOutput, isType bool) (*propertyBag, error) {
	result := newPropertyBag()

	// Sort properties to make codegen deterministic.
	var names []string
	for name := range resolvedSchema.Properties {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		property := resolvedSchema.Properties[name]
		resolvedProperty, err := resolvedSchema.ResolveSchema(&property)
		if err != nil {
			return nil, err
		}

		if name == "etag" && !isType && !isOutput {
			// ETags may be defined as inputs to PUT endpoints, but we should never model them as
			// resource inputs because they are a protocol implementation detail rather than
			// a meaningful desired-state property.
			continue
		}
		if name == "clusterID" && property.Description == "The deprecated identity" {
			// TODO: Get rid of this in https://github.com/pulumi/pulumi-azure-native/issues/331
			continue
		}
		// TODO: Workaround for https://github.com/pulumi/pulumi/issues/12629 - remove once merged
		if (name == "conditionSets" && property.Items.Schema.Ref.String() == "#/definitions/GovernanceRuleConditionSets") ||
			(name == "conditionSets" && property.Items.Schema.Ref.String() == "#/definitions/ApplicationConditionSets") {
			continue
		}

		sdkName := name
		if clientName, ok := resolvedProperty.Extensions.GetString(extensionClientName); ok {
			sdkName = firstToLower(clientName)
		}
		// Change the name to lowerCamelCase.
		sdkName = ToLowerCamel(sdkName)

		// Flattened properties aren't modelled in the SDK explicitly: their sub-properties are merged directly to the parent.
		// If the type is marked as a dictionary, ignore the extension and proceed with modeling this property explicitly.
		// We can't flatten dictionaries in a type-safe manner.
		isDict := resolvedProperty.AdditionalProperties != nil
		//TODO: Remove when https://github.com/Azure/azure-rest-api-specs/pull/14550 is rolled back
		workaroundDelegatedNetworkBreakingChange := property.Ref.String() == "#/definitions/OrchestratorResourceProperties" ||
			property.Ref.String() == "#/definitions/DelegatedSubnetProperties" ||
			property.Ref.String() == "#/definitions/DelegatedControllerProperties"
		if flatten, ok := property.Extensions.GetBool(extensionClientFlatten); (ok && flatten && !isDict) || workaroundDelegatedNetworkBreakingChange {
			bag, err := m.genProperties(resolvedProperty, isOutput, isType)
			if err != nil {
				return nil, err
			}

			// Adjust every property to mark them as flattened.
			newProperties := map[string]resources.AzureAPIProperty{}
			for n, value := range bag.properties {
				// The order of containers is important, so we prepend the outermost name.
				value.Containers = append([]string{name}, value.Containers...)
				newProperties[n] = value
			}
			bag.properties = newProperties

			result.merge(bag)
			continue
		}

		// Skip read-only properties for input types and write-only properties for output types.
		if resolvedProperty.ReadOnly && !isOutput {
			continue
		}
		if isOutput && isWriteOnly(resolvedProperty) {
			continue
		}

		propertySpec, apiProperty, err := m.genProperty(name, &property, resolvedSchema.ReferenceContext, resolvedProperty, isOutput, isType)
		if err != nil {
			return nil, err
		}

		// Skip properties that yield degenerate types (e.g., when an input type has only read-only properties).
		if propertySpec == nil {
			continue
		}

		if isOutput && resolvedProperty.ReadOnly {
			result.requiredSpecs.Add(sdkName)
		}

		if sdkName != name {
			apiProperty.SdkName = sdkName
		}

		result.properties[name] = *apiProperty
		result.specs[sdkName] = *propertySpec
	}

	for _, s := range resolvedSchema.AllOf {
		allOfSchema, err := resolvedSchema.ResolveSchema(&s)
		if err != nil {
			return nil, err
		}

		allOfProperties, err := m.genProperties(allOfSchema, isOutput, isType)
		if err != nil {
			return nil, err
		}

		// For a derived type, set the discriminator property to the const value, if any.
		discriminator, discriminatorDesc, isDU, err := m.getDiscriminator(allOfSchema)
		if err != nil {
			return nil, err
		}
		if isDU {
			prop := allOfProperties.properties[discriminator]
			sdkDiscriminator := discriminator
			if prop.SdkName != "" {
				sdkDiscriminator = prop.SdkName
			}

			propSpec := allOfProperties.specs[sdkDiscriminator]
			discriminatorValue := getDiscriminatorValue(resolvedSchema)
			propSpec.Const = discriminatorValue
			propSpec.Type = "string"
			propSpec.Ref = ""
			propSpec.OneOf = nil

			// Add the discriminator value to the property description to help users fill it.
			propSpec.Description = fmt.Sprintf("%s\nExpected value is '%s'.", discriminatorDesc, discriminatorValue)

			allOfProperties.specs[sdkDiscriminator] = propSpec
			prop.Const = discriminatorValue
			allOfProperties.properties[discriminator] = prop
		}

		result.merge(allOfProperties)
	}

	for _, name := range resolvedSchema.Required {
		if prop, ok := result.properties[name]; ok {
			if prop.SdkName != "" {
				result.requiredSpecs.Add(prop.SdkName)
			} else {
				result.requiredSpecs.Add(name)
			}
			result.requiredProperties.Add(name)
		}
	}
	return result, nil
}

func (m *moduleGenerator) genProperty(name string, schema *spec.Schema, context *openapi.ReferenceContext, resolvedProperty *openapi.Schema, isOutput, isType bool) (*pschema.PropertySpec, *resources.AzureAPIProperty, error) {
	description, err := getPropertyDescription(schema, context)
	if err != nil {
		return nil, nil, err
	}

	typeSpec, err := m.genTypeSpec(name, schema, context, isOutput)
	if err != nil {
		return nil, nil, err
	}

	// Nil type means empty: e.g., when an input type has only read-only properties. Bubble the nil up.
	if typeSpec == nil {
		return nil, nil, nil
	}

	// TODO: Remove this switch if https://github.com/Azure/azure-rest-api-specs/issues/13167 is fixed
	defaultValue := schema.Default
	if defaultValue != nil && typeSpec.Type == "object" {
		fmt.Printf("Default value '%v' can't be specified for an object property %q\n", schema.Default, name)
		defaultValue = nil
	}
	// #2187 - there are array properties with default value '[]' in the schema which we don't support
	if defaultValue != nil && typeSpec.Type == "array" {
		fmt.Printf("Default value '%v' can't be specified for an array property %q\n", schema.Default, name)
		defaultValue = nil
	}

	// Convert default values which are represented using strings instead of their actual types
	switch schema.Default.(type) {
	case string:
		defaultString := schema.Default.(string)
		switch typeSpec.Type {
		case "boolean":
			defaultValue, err = strconv.ParseBool(defaultString)
		case "number":
			defaultValue, err = strconv.ParseFloat(defaultString, 32)
		case "integer":
			defaultValue, err = strconv.ParseInt(defaultString, 10, 32)
		}
		if err != nil {
			defaultValue = nil
		}
		// TODO: Find a better way of detecting when we won't support a default value:
		// E.g. #/types/azure-native:machinelearningservices%2Fv20220601preview:LabelingJobResponse/properties/mlAssistConfiguration/default:
		// type Union<azure-native:machinelearningservices/v20220601preview:MLAssistConfigurationDisabledResponse, azure-native:machinelearningservices/v20220601preview:MLAssistConfigurationEnabledResponse>
		// cannot have a constant value; only booleans, integers, numbers and strings may have constant values;
		// HACK: Check if the default value looks like JSON and remove it.
		if strings.HasPrefix(defaultString, "{") {
			fmt.Printf("Default value '%v' appears to be an object %q\n", schema.Default, name)
			defaultValue = nil
		}
	}

	propertySpec := pschema.PropertySpec{
		Description: description,
		Default:     defaultValue,
		TypeSpec:    *typeSpec,
	}

	apiProperty := m.genApiProperty(isOutput, &propertySpec, resolvedProperty, name, isType)

	return &propertySpec, &apiProperty, nil
}

func (m *moduleGenerator) genApiProperty(isOutput bool, propertySpec *pschema.PropertySpec, resolvedProperty *openapi.Schema, name string, isType bool) resources.AzureAPIProperty {
	apiProperty := resources.AzureAPIProperty{
		OneOf:                m.getOneOfValues(propertySpec),
		Ref:                  propertySpec.Ref,
		Items:                m.itemTypeToProperty(propertySpec.Items),
		AdditionalProperties: m.itemTypeToProperty(propertySpec.AdditionalProperties),
	}
	if isOutput {
		return apiProperty
	}

	// Input types only:
	if m.isEnum(&propertySpec.TypeSpec) {
		apiProperty = resources.AzureAPIProperty{Type: "string"}
	} else {
		// Set additional properties when it's an input
		apiProperty.Type = propertySpec.Type
		apiProperty.Minimum = resolvedProperty.Minimum
		apiProperty.Maximum = resolvedProperty.Maximum
		apiProperty.MinLength = resolvedProperty.MinLength
		apiProperty.MaxLength = resolvedProperty.MaxLength
		apiProperty.Pattern = resolvedProperty.Pattern
	}

	// Apply manual metadata about Force New properties.
	if m.forceNew(resolvedProperty, name, isType) {
		apiProperty.ForceNew = true
		propertySpec.WillReplaceOnChanges = true
	}
	return apiProperty
}

func newPropertyBag() *propertyBag {
	return &propertyBag{
		specs:              map[string]pschema.PropertySpec{},
		properties:         map[string]resources.AzureAPIProperty{},
		requiredSpecs:      codegen.NewStringSet(),
		requiredProperties: codegen.NewStringSet(),
	}
}

func (bag *propertyBag) merge(other *propertyBag) {
	for key, value := range other.specs {
		bag.specs[key] = value
	}
	for key, value := range other.properties {
		bag.properties[key] = value
	}
	for key := range other.requiredSpecs {
		bag.requiredSpecs.Add(key)
	}
	for key := range other.requiredProperties {
		bag.requiredProperties.Add(key)
	}
}

// forceNew return true if a property with a given name requires a replacement in the resource
// that is currently being generated, based on forceNewMap.
func (m *moduleGenerator) forceNew(schema *openapi.Schema, propertyName string, isType bool) bool {
	// Mutability extension signals whether a property can be updated in-place. Lack of the extension means
	// updatable by default.
	// Note: a non-updatable property at a subtype level (a property of a property of a resource) does not
	// mandate the replacement of the whole resource. Anyway, it's used very seldom (2 places at the time of writing).
	// Example: `StorageAccount.encryption.services.blob.keyType` is non-updatable, but a user can remove `blob`
	// and then re-add it with the new `keyType` without replacing the whole storage account (which would be
	// very disruptive).
	if mutability, ok := schema.Extensions.GetStringSlice(extensionMutability); ok && !isType {
		for _, v := range mutability {
			if v == extensionMutabilityUpdate {
				return false
			}
		}
		return true
	}

	if resourceMap, ok := forceNewMap[m.prov]; ok {
		if properties, ok := resourceMap[m.resourceName]; ok {
			if properties.Has(propertyName) {
				return true
			}
		}
	}

	return false
}

// itemTypeToProperty converts a type of an element in an array or a dictionary to a corresponding
// API property definition of that element type. It only converts the relevant subset of properties,
// and does so recursively.
func (m *moduleGenerator) itemTypeToProperty(typ *schema.TypeSpec) *resources.AzureAPIProperty {
	if typ == nil || typ.Ref == resources.TypeAny {
		return nil
	}

	if m.isEnum(typ) {
		return &resources.AzureAPIProperty{Type: "string"}
	}

	var oneOf []string
	for _, subType := range typ.OneOf {
		if subType.Ref != "" {
			oneOf = append(oneOf, subType.Ref)
		} else {
			oneOf = append(oneOf, subType.Type)
		}
	}

	return &resources.AzureAPIProperty{
		Type:                 typ.Type,
		Ref:                  typ.Ref,
		OneOf:                oneOf,
		Items:                m.itemTypeToProperty(typ.Items),
		AdditionalProperties: m.itemTypeToProperty(typ.AdditionalProperties),
	}
}

func (m *moduleGenerator) isEnum(typ *schema.TypeSpec) bool {
	for _, subType := range typ.OneOf {
		if m.isEnum(&subType) {
			return true
		}
	}

	if typ.Ref == "" {
		return false
	}

	typeName := strings.TrimPrefix(typ.Ref, "#/types/")
	refTyp := m.pkg.Types[typeName]
	return len(refTyp.Enum) > 0
}

func (m *moduleGenerator) getOneOfValues(property *pschema.PropertySpec) (values []string) {
	for _, value := range property.OneOf {
		values = append(values, value.Ref)
	}
	return
}

// isWriteOnly return true for properties which are annotated with mutability extension that contain no 'read' value.
func isWriteOnly(schema *openapi.Schema) bool {
	mutability, has := schema.Extensions.GetStringSlice(extensionMutability)
	if !has {
		return false
	}
	for _, v := range mutability {
		if v == extensionMutabilityRead {
			return false
		}
	}
	return true
}

// getDiscriminator returns a property name and description for a discriminator if it's defined on the schema.
// The boolean return flag is true when a discriminator is found.
func (m *moduleGenerator) getDiscriminator(resolvedSchema *openapi.Schema) (string, string, bool, error) {
	if resolvedSchema.Discriminator != "" {
		property := resolvedSchema.Properties[resolvedSchema.Discriminator]
		resolvedProperty, err := resolvedSchema.ResolveSchema(&property)
		if err != nil {
			return "", "", false, err
		}
		return resolvedSchema.Discriminator, resolvedProperty.Description, true, nil
	}
	for _, s := range resolvedSchema.AllOf {
		parentSchema, err := resolvedSchema.ResolveSchema(&s)
		if err != nil {
			return "", "", false, err
		}
		parentDiscriminator, parentDescription, has, err := m.getDiscriminator(parentSchema)
		if err != nil {
			return "", "", false, err
		}
		if has {
			return parentDiscriminator, parentDescription, true, nil
		}
	}
	return "", "", false, nil
}
