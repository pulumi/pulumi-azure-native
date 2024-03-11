// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"fmt"

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
// The override does two things:
// 1. Defines a new type DashboardPartMetadata with three properties: type, inputs, and settings. Inputs and settings are
// loosely typed collections to allow for arbitrary JSON objects.
// 2. Overrides the 'metadata' property type of the portal:DashboardParts type to be DashboardPartMetadata.
//
// The override effectively removes the discriminated union and replaces it with a single loosely-typed structure.
// We override both input and output (Response) types.
func portalDashboard() *CustomResource {
	return &CustomResource{
		Types: map[string]schema.ComplexTypeSpec{
			resources.BuildToken(PortalMod, "", MetadataType):            dashboardPartMetadataType(),
			resources.BuildToken(PortalMod, "", MetadataType+"Response"): dashboardPartMetadataType(),
		},
		TypeOverrides: map[string]schema.ComplexTypeSpec{
			resources.BuildToken(PortalMod, "", PartsType):            dashboardPartsType(""),
			resources.BuildToken(PortalMod, "", PartsType+"Response"): dashboardPartsType("Response"),
		},
		MetaTypes: map[string]resources.AzureAPIType{
			resources.BuildToken(PortalMod, "", MetadataType):            dashboardPartMetadataApiType(),
			resources.BuildToken(PortalMod, "", MetadataType+"Response"): dashboardPartMetadataApiType(),
		},
		MetaTypeOverrides: map[string]resources.AzureAPIType{
			resources.BuildToken(PortalMod, "", PartsType):            dashboardPartsApiType(""),
			resources.BuildToken(PortalMod, "", PartsType+"Response"): dashboardPartsApiType("Response"),
		},
	}
}

const PortalMod = "portal"
const PartsType = "DashboardParts"
const MetadataType = "DashboardPartMetadata"
const MarkdownMetadataType = "MarkdownPartMetadata"

func dashboardPartsType(suffix string) schema.ComplexTypeSpec {
	return schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "A dashboard part.",
			Type:        "object",
			Properties: map[string]schema.PropertySpec{
				"metadata": {
					TypeSpec: schema.TypeSpec{
						Type: "object",
						Ref:  fmt.Sprintf("#/types/azure-native:portal:DashboardPartMetadata%s", suffix),
					},
					Description: "The dashboard's part metadata.",
				},
				"position": {
					TypeSpec: schema.TypeSpec{
						Type: "object",
						Ref:  fmt.Sprintf("#/types/azure-native:portal:DashboardParts%sPosition", suffix),
					},
					Description: "The dashboard's part position.",
				},
			},
			Required: []string{"position"},
		},
	}
}

func dashboardPartsApiType(suffix string) resources.AzureAPIType {
	return resources.AzureAPIType{
		Properties: map[string]resources.AzureAPIProperty{
			"metadata": {
				Type: "object",
				Ref:  fmt.Sprintf("#/types/azure-native:portal:DashboardPartMetadata%s", suffix),
			},
			"position": {
				Type: "object",
				Ref:  fmt.Sprintf("#/types/azure-native:portal:DashboardParts%sPosition", suffix),
			},
		},
		RequiredProperties: []string{"position"},
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
