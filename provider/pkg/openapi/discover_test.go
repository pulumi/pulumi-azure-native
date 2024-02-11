package openapi

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/defaults"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListDiagnosticCategoriesShouldBeAdded(t *testing.T) {
	path := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/diagnosticSettingsCategories"

	version := NewVersionResources()

	swagger := Spec{
		Swagger: spec.Swagger{
			SwaggerProps: spec.SwaggerProps{
				Info: &spec.Info{
					InfoProps: spec.InfoProps{
						Version: "2020-01-01",
					},
				},
				Paths: &spec.Paths{
					Paths: map[string]spec.PathItem{
						path: {
							PathItemProps: spec.PathItemProps{
								Get: &spec.Operation{
									OperationProps: spec.OperationProps{
										ID: "DiagnosticSettingsCategory_List",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	addResourcesAndInvokes(version, "/file/path", path, "insights", &swagger)

	assert.Empty(t, version.Resources)
	assert.NotEmpty(t, version.Invokes)

	invoke, ok := version.Invokes["listDiagnosticSettingsCategory"]
	assert.True(t, ok)
	assert.Equal(t, path, invoke.Path)
}

func TestDefaultState(t *testing.T) {
	makeSwagger := func(path string) Spec {
		return Spec{
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Info: &spec.Info{
						InfoProps: spec.InfoProps{
							Version: "2020-01-01",
						},
					},
					Paths: &spec.Paths{
						Paths: map[string]spec.PathItem{
							path: {
								// Needs GET, DELETE, PUT to be discovered as a resource
								PathItemProps: spec.PathItemProps{
									Get: &spec.Operation{
										OperationProps: spec.OperationProps{
											ID: "DiagnosticSettingsCategory_Get",
										},
									},
									Put: &spec.Operation{
										OperationProps: spec.OperationProps{
											ID: "DiagnosticSettingsCategory_Put",
										},
									},
									Delete: &spec.Operation{
										OperationProps: spec.OperationProps{
											ID: "DiagnosticSettingsCategory_Delete",
										},
									},
								},
							},
						},
					},
				},
			},
		}
	}

	t.Run("No default", func(t *testing.T) {
		path := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/diagnosticSettingsCategories/{categoryName}"
		swagger := makeSwagger(path)

		version := NewVersionResources()

		addResourcesAndInvokes(version, "/file/path", path, "insights", &swagger)

		require.NotEmpty(t, version.Resources)
		res, ok := version.Resources["DiagnosticSettingsCategory"]
		require.True(t, ok)
		assert.Nil(t, res.DefaultBody)
	})

	t.Run("With default", func(t *testing.T) {
		path := "/{resourceId}/providers/Microsoft.Security/advancedThreatProtectionSettings/{settingName}"
		def := defaults.GetDefaultResourceState(path)
		require.NotNil(t, def)

		swagger := makeSwagger(path)
		version := NewVersionResources()

		addResourcesAndInvokes(version, "/file/path", path, "insights", &swagger)

		require.NotEmpty(t, version.Resources)
		res, ok := version.Resources["DiagnosticSettingsCategory"]
		require.True(t, ok)
		assert.Equal(t, def.State, res.DefaultBody)
	})
}
