// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestPortalDashboardTransformation(t *testing.T) {
	types := map[string]schema.ComplexTypeSpec{
		"azure-native:portal:DashboardParts": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"metadata": {
						TypeSpec: schema.TypeSpec{
							Type: "object",
							Ref:  "#/types/azure-native:portal:MarkdownPartMetadata",
						},
					},
					"position": {
						TypeSpec: schema.TypeSpec{
							Type: "object",
							Ref:  "#/types/azure-native:portal:DashboardPartsPosition",
						},
					},
				},
			},
		},
		"azure-native:portal/v20240101:DashboardPartsResponse": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Properties: map[string]schema.PropertySpec{
					"metadata": {
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/azure-native:portal:MarkdownPartMetadataResponse",
						},
					},
				},
			},
		},
		"azure-native:portal:MarkdownPartMetadata": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
			},
		},
		"azure-native:portal/v20240101:MarkdownPartMetadataResponseContent": {},
		"azure-native:portal:MarkdownPartMetadataSettings":                  {},
		"azure-native:portal:MarkdownPartMetadataSettingsSettings":          {},
	}

	metaTypes := map[string]resources.AzureAPIType{
		"azure-native:portal:DashboardParts": {
			Properties: map[string]resources.AzureAPIProperty{
				"metadata": {
					Type: "object",
					Ref:  "#/types/azure-native:portal:MarkdownPartMetadata",
				},
				"position": {
					Type: "object",
					Ref:  "#/types/azure-native:portal:DashboardPartsPosition",
				},
			},
		},
		"azure-native:portal/v20240101:DashboardPartsResponse": {
			Properties: map[string]resources.AzureAPIProperty{
				"metadata": {
					Ref: "#/types/azure-native:portal:MarkdownPartMetadataResponse",
				},
			},
		},
		"azure-native:portal:MarkdownPartMetadata": {},
	}

	portalDashboardTransformation(types, metaTypes)

	// Check the default version.
	assert.Equal(t, "#/types/azure-native:portal:DashboardPartMetadata", types["azure-native:portal:DashboardParts"].Properties["metadata"].Ref)
	assert.Equal(t, dashboardPartMetadataType(), types["azure-native:portal:DashboardPartMetadata"])
	assert.Equal(t, "#/types/azure-native:portal:DashboardPartMetadata", metaTypes["azure-native:portal:DashboardParts"].Properties["metadata"].Ref)
	assert.Equal(t, dashboardPartMetadataApiType(), metaTypes["azure-native:portal:DashboardPartMetadata"])

	// Check the explicit version.
	assert.Equal(t, "#/types/azure-native:portal/v20240101:DashboardPartMetadataResponse", types["azure-native:portal/v20240101:DashboardPartsResponse"].Properties["metadata"].Ref)
	assert.Equal(t, dashboardPartMetadataType(), types["azure-native:portal/v20240101:DashboardPartMetadataResponse"])
	assert.Equal(t, "#/types/azure-native:portal/v20240101:DashboardPartMetadataResponse", metaTypes["azure-native:portal/v20240101:DashboardPartsResponse"].Properties["metadata"].Ref)
	assert.Equal(t, dashboardPartMetadataApiType(), metaTypes["azure-native:portal/v20240101:DashboardPartMetadataResponse"])

	// Check that all irrelevant types got removed.
	assert.Equal(t, 4, len(types))
	assert.Equal(t, 4, len(metaTypes))
}
