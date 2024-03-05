// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// portalDashboard creates an override for the portal:Dashboard resource types.
//
// This change is needed because the Azure API open API specs define DashboardPartMetadata as a discriminated type but
// with only a concrete type defined. The API supports a multitude of part types, but the open API spec only describes
// one of them. The result of this is that we generate a discriminated union which is so restricted that it can't be used
// in practice.
//
// The override does three things:
// 1. Defines a new type DashboardPartMetadata with three properties: type, inputs, and settings. Inputs and settings are
// loosely typed collections to allow for arbitrary JSON objects.
// 2. Overrides the 'metadata' property type of the portal:DashboardParts type to be DashboardPartMetadata.
// 3. Deletes the redundant type MarkdownPartMetadata that is the only concrete type defined in the specs, currently.
//
// The override effectively removes the discriminated union and replaces it with a single loosely-typed structure.
// We override both input and output (Response) types.
func portalDashboard() *CustomResource {
	return &CustomResource{
		Transformation: portalDashboardTransformation,
	}
}

const PortalMod = "portal"
const PartsType = "DashboardParts"
const MetadataType = "DashboardPartMetadata"
const MarkdownMetadataType = "MarkdownPartMetadata"

func portalDashboardTransformation(types map[string]schema.ComplexTypeSpec, metaTypes map[string]resources.AzureAPIType) {
	// The list of existing type suffixes that need to be deleted for each redundant MarkdownPartMetadata type.
	markdownSubTypes := []string{
		"",
		"Content",
		"Settings",
		"SettingsSettings",
	}

	for tok, typ := range types {
		// Skip all types except for the portal:DashboardParts type.
		mod, apiVersion, typeName, err := resources.ParseToken(tok)
		if err != nil || mod != PortalMod || (typeName != PartsType && typeName != PartsType+"Response") {
			continue
		}

		newTypeTok := resources.BuildToken(PortalMod, apiVersion, MetadataType)
		redundantTypeTok := resources.BuildToken(PortalMod, apiVersion, MarkdownMetadataType)
		if strings.HasSuffix(typeName, "Response") {
			newTypeTok += "Response"
			redundantTypeTok += "Response"
		}

		// Define the new type DashboardPartMetadata.
		types[newTypeTok] = dashboardPartMetadataType()
		metaTypes[newTypeTok] = dashboardPartMetadataApiType()

		// Override the 'metadata' property reference to point to the newly defined DashboardPartMetadata.
		newTypeRef := fmt.Sprintf("#/types/%s", newTypeTok)
		typ.Properties["metadata"] = schema.PropertySpec{
			Description: "The dashboard part's metadata.",
			TypeSpec: schema.TypeSpec{
				Type: "object",
				Ref:  newTypeRef,
			},
		}
		types[tok] = typ
		meta, ok := metaTypes[tok]
		if !ok {
			panic(fmt.Sprintf("metadata type %q not found", tok))
		}
		meta.Properties["metadata"] = resources.AzureAPIProperty{
			Type: "object",
			Ref:  newTypeRef,
		}
		metaTypes[tok] = meta

		// Delete the redundant type MarkdownPartMetadata (the only specific type defined in Open API).
		for _, subType := range markdownSubTypes {
			markdownTypeTok := redundantTypeTok + subType
			delete(types, markdownTypeTok)
			delete(metaTypes, markdownTypeTok)
		}
	}
}

func dashboardPartMetadataType() schema.ComplexTypeSpec {
	return schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "A dashboard part metadata.",
			Type:        "object",
			Properties: map[string]schema.PropertySpec{
				"type": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "The type of dashboard part.",
				},
				"inputs": {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Ref: "pulumi.json#/Any"},
					},
					Description: "Inputs to dashboard part.",
				},
				"settings": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Ref: "pulumi.json#/Any"},
					},
					Description: "Settings of dashboard part.",
				},
			},
			Required: []string{"type"},
		},
	}
}

func dashboardPartMetadataApiType() resources.AzureAPIType {
	return resources.AzureAPIType{
		Properties: map[string]resources.AzureAPIProperty{
			"type": {
				Type: "string",
			},
			"inputs": {
				Type:  "array",
				Items: &resources.AzureAPIProperty{Ref: "pulumi.json#/Any"},
			},
			"settings": {
				Type:                 "object",
				AdditionalProperties: &resources.AzureAPIProperty{Ref: "pulumi.json#/Any"},
			},
		},
		RequiredProperties: []string{"type"},
	}
}
