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

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// propertyBag keeps the schema and metadata properties for a single API type.
type propertyBag struct {
	description        string
	specs              map[string]pschema.PropertySpec
	properties         map[string]resources.AzureAPIProperty
	requiredSpecs      codegen.StringSet
	requiredProperties codegen.StringSet
	requiredContainers RequiredContainers
}

type RequiredContainers [][]string

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

			// Check that none of the inner properties already exists on the outer type. This
			// causes a conflict when flattening, and will probably need to be handled in v3.
			for propName := range bag.properties {
				if _, has := result.properties[propName]; has {
					m.flattenedPropertyConflicts[fmt.Sprintf("%s.%s", name, propName)] = struct{}{}
				}
			}

			// Adjust every property to mark them as flattened.
			newProperties := map[string]resources.AzureAPIProperty{}
			for n, value := range bag.properties {
				// The order of containers is important, so we prepend the outermost name.
				value.Containers = append([]string{name}, value.Containers...)
				newProperties[n] = value
			}
			bag.properties = newProperties

			newRequiredContainers := make(RequiredContainers, len(bag.requiredContainers))
			for i, containers := range bag.requiredContainers {
				newRequiredContainers[i] = append([]string{name}, containers...)
			}
			for _, requiredName := range resolvedSchema.Required {
				if requiredName == name {
					newRequiredContainers = append(newRequiredContainers, []string{name})
				}
			}
			bag.requiredContainers = newRequiredContainers

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
	// Identify properties which are aso available as standalone resources and mark them to be maintained if not specified inline.
	// Ignore types as we only support top-level resource properties
	// Ignore outputs as this is only affecting the input args of a resource, not the resource outputs.
	// We only consider arrays (with Items) because in order for the sub-resource to be a standalone resource it must be able to have many instances.
	// There is other kinds of sub-resources where there's only a single instance within the parent resource but these are not handled here.
	// They are currently handled by the openapi.default module - where we have to add a special case for them *because* they're not managed by the parent and don't have their own delete method.
	maintainSubResourceIfUnset := false
	if !isType && !isOutput && schema.Items != nil && schema.Items.Schema != nil {
		itemsRef := schema.Items.Schema.Ref.String()
		for _, nestedRef := range m.nestedResourceBodyRefs {
			if itemsRef == nestedRef {
				maintainSubResourceIfUnset = true
			}
		}
	}

	// Special case for KeyVault access policies - they are a resource that can be defined inline in Vaults or
	// stand-alone. The logic above cannot detect that because they are defined as a custom resource. #594
	if m.resourceName == "Vault" && m.prov == "KeyVault" && name == "accessPolicies" {
		maintainSubResourceIfUnset = true
	}

	description, err := getPropertyDescription(schema, context, maintainSubResourceIfUnset)
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
		logging.V(5).Infof("Default value '%v' can't be specified for an object property %q\n", schema.Default, name)
		defaultValue = nil
	}
	// #2187 - there are array properties with default value '[]' in the schema which we don't support
	if defaultValue != nil && typeSpec.Type == "array" {
		logging.V(5).Infof("Default value '%v' can't be specified for an array property %q\n", schema.Default, name)
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
			logging.V(5).Infof("Default value '%v' appears to be an object %q\n", schema.Default, name)
			defaultValue = nil
		}
	}

	// If there's no object value type, then it's just a set of strings which we'll represent as a string
	// array in the SDK, but leave the metadata to indicate we need to convert it.
	isStringSet := typeSpec.Type == "object" && typeSpec.AdditionalProperties == nil && typeSpec.Ref == ""

	forceNewSpec := noForceNew
	if !isOutput {
		forceNewSpec = m.forceNew(resolvedProperty, name, isType)
	}

	schemaProperty := pschema.PropertySpec{
		Description:          description,
		Default:              defaultValue,
		TypeSpec:             *typeSpec,
		WillReplaceOnChanges: forceNewSpec == forceNew,
	}

	if isStringSet {
		schemaProperty.Type = "array"
		schemaProperty.AdditionalProperties = nil
		schemaProperty.Items = &pschema.TypeSpec{
			Type: "string",
		}
	}

	metadataProperty := resources.AzureAPIProperty{
		OneOf:                               m.getOneOfValues(typeSpec),
		Ref:                                 schemaProperty.Ref,
		Items:                               m.itemTypeToProperty(typeSpec.Items),
		AdditionalProperties:                m.itemTypeToProperty(typeSpec.AdditionalProperties),
		ForceNew:                            forceNewSpec == forceNew,
		ForceNewInferredFromReferencedTypes: forceNewSpec == forceNewSetOnReferencedType,
		IsStringSet:                         isStringSet,
		Default:                             defaultValue,
		MaintainSubResourceIfUnset:          maintainSubResourceIfUnset,
	}

	if identifiers, ok := schema.Extensions.GetStringSlice(extensionIdentifiers); ok && typeSpec.Type == "array" {
		metadataProperty.ArrayIdentifiers = identifiers
	}

	// Input types only get extra information attached
	if !isOutput {
		if m.isEnum(&schemaProperty.TypeSpec) {
			metadataProperty = resources.AzureAPIProperty{
				Type:     "string",
				Default:  defaultValue,
				ForceNew: forceNewSpec == forceNew,
			}
		} else {
			// Set additional properties when it's an input
			metadataProperty.Type = typeSpec.Type
			metadataProperty.Minimum = resolvedProperty.Minimum
			metadataProperty.Maximum = resolvedProperty.Maximum
			metadataProperty.MinLength = resolvedProperty.MinLength
			metadataProperty.MaxLength = resolvedProperty.MaxLength
			metadataProperty.Pattern = resolvedProperty.Pattern
		}
	}

	return &schemaProperty, &metadataProperty, nil
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
	bag.requiredContainers = mergeRequiredContainers(bag.requiredContainers, other.requiredContainers)
}

func mergeRequiredContainers(a, b RequiredContainers) RequiredContainers {
	if len(a) == 0 && len(b) == 0 {
		return nil
	}
	result := make(RequiredContainers, 0, len(a)+len(b))
	// Index each container by concatenating its elements with a separator.
	// This is necessary because we can't compare slices directly.
	index := map[string]bool{}
	for _, containers := range a {
		index[strings.Join(containers, ".")] = true
		result = append(result, containers)
	}
	for _, containers := range b {
		if _, ok := index[strings.Join(containers, ".")]; !ok {
			result = append(result, containers)
		}
	}
	return result
}

type forceNewMetadata string

const (
	forceNew                    forceNewMetadata = "ForceNew"
	forceNewSetOnReferencedType forceNewMetadata = "ForceNewSetOnReferencedType"
	noForceNew                  forceNewMetadata = "NoForceNew"
)

// forceNew returns true if a change to a given property requires a replacement in the resource
// that is currently being generated, based on forceNewMap and the "x-ms-mutability" API spec extension.
// The second return value is true if the property forcing replacement is a property in a referenced type
// (i.e., a property of a property).
func (m *moduleGenerator) forceNew(schema *openapi.Schema, propertyName string, isType bool) forceNewMetadata {
	// Mutability extension signals whether a property can be updated in-place. Lack of the extension means
	// updatable by default.
	// Note: a non-updatable property at a subtype level (a property of a property of a resource) does not
	// mandate the replacement of the whole resource.
	// Example: `StorageAccount.encryption.services.blob.keyType` is non-updatable, but a user can remove `blob`
	// and then re-add it with the new `keyType` without replacing the whole storage account (which would be
	// very disruptive).

	hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)

	if hasMutabilityInfo && forcesRecreate {
		if isType {
			m.forceNewTypes = append(m.forceNewTypes, ForceNewType{
				Module:        m.module,
				Provider:      m.prov,
				ResourceName:  m.resourceName,
				ReferenceName: schema.ReferenceContext.ReferenceName,
				Property:      propertyName,
			})
			return forceNewSetOnReferencedType
		}
		return forceNew
	}

	if resourceMap, ok := forceNewMap[m.prov]; ok {
		if properties, ok := resourceMap[m.resourceName]; ok {
			if properties.Has(propertyName) {
				return forceNew
			}
		}
	}

	return noForceNew
}

// propChangeForcesRecreate returns two booleans.
// The first one indicates whether the schema has mutability extension.
// The second one indicates whether the property requires recreation to change.
func propChangeForcesRecreate(schema *openapi.Schema) (bool, bool) {
	hasUpdate, hasCreate := false, false
	if mutability, ok := schema.Extensions.GetStringSlice(extensionMutability); ok {
		for _, v := range mutability {
			switch v {
			case extensionMutabilityCreate:
				hasCreate = true
			case extensionMutabilityUpdate:
				hasUpdate = true
			}
		}
		if hasCreate && !hasUpdate {
			return true, true // has mutability info and is forces recreation
		}
		return true, false // has mutability info but is not updatable
	}
	return false, false // does not have mutability info
}

// itemTypeToProperty converts a type of an element in an array or a dictionary to a corresponding
// API property definition of that element type. It only converts the relevant subset of properties,
// and does so recursively.
func (m *moduleGenerator) itemTypeToProperty(typ *schema.TypeSpec) *resources.AzureAPIProperty {
	if typ == nil || typ.Ref == convert.TypeAny {
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

func (m *moduleGenerator) getOneOfValues(property *pschema.TypeSpec) (values []string) {
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
