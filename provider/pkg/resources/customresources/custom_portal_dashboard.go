// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"fmt"
	"reflect"

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
		tok: "azure-native:portal:Dashboard",
		SchemaF: func(resource *ResourceDefinition) (*ResourceDefinition, error) {
			if resource == nil {
				// This should only happen if we're running with a namespace or version filter for testing, so we'll allow it to pass.
				return nil, nil
			}
			existingDashboardPartsType := resource.Types["azure-native:portal:DashboardParts"]
			if !reflect.DeepEqual(existingDashboardPartsType, expectedDashboardPartsType("")) {
				return nil, fmt.Errorf("unexpected azure-native:portal:DashboardParts type: %#v", existingDashboardPartsType)
			}
			resource.Types["azure-native:portal:DashboardParts"] = dashboardPartsType("")
			existingDashboardPartsTypeResponse := resource.Types["azure-native:portal:DashboardPartsResponse"]
			if !reflect.DeepEqual(existingDashboardPartsTypeResponse, expectedDashboardPartsType("Response")) {
				return nil, fmt.Errorf("unexpected azure-native:portal:DashboardPartsResponse type: %#v", existingDashboardPartsTypeResponse)
			}
			resource.Types["azure-native:portal:DashboardPartsResponse"] = dashboardPartsType("Response")

			existingMarkdownPartMetadataType := resource.Types["azure-native:portal:MarkdownPartMetadata"]
			if !reflect.DeepEqual(existingMarkdownPartMetadataType, expectedMarkdownPartMetadataType("")) {
				return nil, fmt.Errorf("unexpected azure-native:portal:MarkdownPartMetadata type: %#v", existingMarkdownPartMetadataType)
			}
			if _, ok := resource.Types["azure-native:portal:DashboardPartMetadata"]; ok {
				return nil, fmt.Errorf("unexpected existing azure-native:portal:DashboardPartMetadata type")
			}
			resource.Types["azure-native:portal:DashboardPartMetadata"] = dashboardPartMetadataType()
			existingMarkdownPartMetadataResponseType := resource.Types["azure-native:portal:MarkdownPartMetadataResponse"]
			if !reflect.DeepEqual(existingMarkdownPartMetadataResponseType, expectedMarkdownPartMetadataType("Response")) {
				return nil, fmt.Errorf("unexpected azure-native:portal:MarkdownPartMetadataResponse type: %#v", existingMarkdownPartMetadataResponseType)
			}
			if _, ok := resource.Types["azure-native:portal:DashboardPartMetadataResponse"]; ok {
				return nil, fmt.Errorf("unexpected existing azure-native:portal:DashboardPartMetadataResponse type")
			}
			resource.Types["azure-native:portal:DashboardPartMetadataResponse"] = dashboardPartMetadataType()

			resource.MetaTypes["azure-native:portal:DashboardParts"] = dashboardPartsApiType("")
			resource.MetaTypes["azure-native:portal:DashboardPartsResponse"] = dashboardPartsApiType("Response")

			resource.MetaTypes["azure-native:portal:DashboardPartMetadata"] = dashboardPartMetadataApiType()
			resource.MetaTypes["azure-native:portal:DashboardPartMetadataResponse"] = dashboardPartMetadataApiType()
			return resource, nil
		},
	}
}

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
						Ref:  fmt.Sprintf("#/types/azure-native:portal:DashboardPartsPosition%s", suffix),
					},
					Description: "The dashboard's part position.",
				},
			},
			Required: []string{"position"},
		},
	}
}

func expectedDashboardPartsType(suffix string) schema.ComplexTypeSpec {
	return schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "A dashboard part.",
			Properties: map[string]schema.PropertySpec{
				"metadata": {
					TypeSpec: schema.TypeSpec{
						Type: "object",
						Ref:  fmt.Sprintf("#/types/azure-native:portal:MarkdownPartMetadata%s", suffix),
					},
					Description: "The dashboard part's metadata.",
				},
				"position": {
					TypeSpec: schema.TypeSpec{
						Type: "object",
						Ref:  fmt.Sprintf("#/types/azure-native:portal:DashboardPartsPosition%s", suffix),
					},
					Description: "The dashboard's part position.",
				},
			},
			Type:     "object",
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
				Ref:  fmt.Sprintf("#/types/azure-native:portal:DashboardPartsPosition%s", suffix),
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

func expectedMarkdownPartMetadataType(suffix string) schema.ComplexTypeSpec {
	return schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Markdown part metadata.",
			Properties: map[string]schema.PropertySpec{
				"inputs": {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Ref: "pulumi.json#/Any"},
					},
					Description: "Input to dashboard part.",
				},
				"settings": {
					TypeSpec: schema.TypeSpec{
						Type: "object",
						Ref:  fmt.Sprintf("#/types/azure-native:portal:MarkdownPartMetadataSettings%s", suffix),
					},
					Description: "Markdown part settings.",
				},
				"type": {
					TypeSpec: schema.TypeSpec{
						Type: "string",
					},
					Description: "The dashboard part metadata type.\nExpected value is 'Extension/HubsExtension/PartType/MarkdownPart'.",
					Const:       "Extension/HubsExtension/PartType/MarkdownPart",
				},
			},
			Type:     "object",
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
