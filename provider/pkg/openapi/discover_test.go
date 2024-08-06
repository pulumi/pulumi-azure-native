package openapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/defaults"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

func createTestSwaggerForInvoke(path, operationId string, httpMethod string) Spec {
	op := &spec.Operation{
		OperationProps: spec.OperationProps{
			ID: operationId,
		},
	}

	props := spec.PathItemProps{}
	switch httpMethod {
	case http.MethodGet:
		props.Get = op
	case http.MethodPost:
		props.Post = op
	default:
		panic("unsupported HTTP method " + httpMethod)
	}

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
							PathItemProps: props,
						},
					},
				},
			},
		},
	}
}

func TestAddInvokes(t *testing.T) {
	for tcName, tc := range map[string]struct {
		path        string
		operationId string
		httpMethod  string
		expected    string
	}{
		"listDiagnosticSettingsCategory": {
			path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/diagnosticSettingsCategories",
			operationId: "DiagnosticSettingsCategory_List",
			httpMethod:  http.MethodGet,
			expected:    "listDiagnosticSettingsCategory",
		},
		"retrieveRegistrationToken": {
			path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/retrieveRegistrationToken",
			operationId: "HostPools_RetrieveRegistrationToken",
			httpMethod:  http.MethodPost,
			expected:    "getHostPoolRegistrationToken",
		},
	} {
		version := NewVersionResources()
		spec := createTestSwaggerForInvoke(tc.path, tc.operationId, tc.httpMethod)

		addResourcesAndInvokes(version, "/file/path", tc.path, "foo", &spec)

		assert.Empty(t, version.Resources, tcName)
		assert.NotEmpty(t, version.Invokes, tcName)

		invoke, ok := version.Invokes[tc.expected]
		require.True(t, ok, fmt.Sprintf("%s: found invokes: %v", tcName, maps.Keys(version.Invokes)))
		assert.Equal(t, tc.path, invoke.Path, tcName)
	}
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

	parseSwagger := func(t *testing.T, path string) *ResourceSpec {
		swagger := makeSwagger(path)
		version := NewVersionResources()

		addResourcesAndInvokes(version, "/file/path", path, "insights", &swagger)

		require.NotEmpty(t, version.Resources)
		res, ok := version.Resources["DiagnosticSettingsCategory"]
		require.True(t, ok)
		return res
	}

	t.Run("No default", func(t *testing.T) {
		path := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/diagnosticSettingsCategories/{categoryName}"

		def := defaults.GetDefaultResourceState(path, "2020-01-01")
		require.Nil(t, def)

		res := parseSwagger(t, path)
		assert.Nil(t, res.DefaultBody)
	})

	t.Run("With default", func(t *testing.T) {
		path := "/{resourceId}/providers/Microsoft.Security/advancedThreatProtectionSettings/{settingName}"

		def := defaults.GetDefaultResourceState(path, "2020-01-01")
		require.NotNil(t, def)

		res := parseSwagger(t, path)
		assert.Equal(t, def.State, res.DefaultBody)
	})
}
