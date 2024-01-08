package openapi

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
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
