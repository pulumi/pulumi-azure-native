package openapi

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestListDiagnosticCategoriesShouldBeAdded(t *testing.T) {
	path := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/diagnosticSettingsCategories"

	version := NewVersionResources()

	pathItem := spec.PathItem{
		PathItemProps: spec.PathItemProps{
			Get: &spec.Operation{
				OperationProps: spec.OperationProps{
					ID: "DiagnosticSettingsCategory_List",
				},
			},
		},
	}

	addGetFunctionIfRequired(version, pathItem, path, nil /* swagger */)

	assert.Empty(t, version.POST_Invokes)
	assert.NotEmpty(t, version.GET_Invokes)

	invoke, ok := version.GET_Invokes["listDiagnosticSettingsCategory"]
	assert.True(t, ok)
	assert.Equal(t, path, invoke.Path)
}
