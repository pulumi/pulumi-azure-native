package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/blang/semver"

	az "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources/customresources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRestoreDefaultInputs(t *testing.T) {
	inputs := resource.PropertyMap{
		"unrelated": resource.NewStringProperty("foo"),
	}
	olds := resource.PropertyMap{
		"unrelated":      resource.NewStringProperty("foo"),
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{}),
	}

	res := resources.AzureAPIResource{
		DefaultProperties: map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"defaultAction": "Allow",
			},
		},
	}

	err := restoreDefaultInputsForRemovedProperties(inputs, res, olds)
	assert.NoError(t, err)

	// Was not in inputs but was added to reset it back to default.
	assert.Contains(t, inputs, resource.PropertyKey("networkRuleSet"))
}

func TestDoNotRestoreDefaultInputsIfInputPresent(t *testing.T) {
	inputs := resource.PropertyMap{
		"unrelated": resource.NewStringProperty("bar"),
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{
			"defaultAction": resource.NewStringProperty("Deny"),
		}),
	}
	olds := resource.PropertyMap{
		"unrelated":      resource.NewStringProperty("foo"),
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{}),
	}

	res := resources.AzureAPIResource{
		DefaultProperties: map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"defaultAction": "Allow",
			},
		},
	}

	err := restoreDefaultInputsForRemovedProperties(inputs, res, olds)
	assert.NoError(t, err)

	assert.Contains(t, inputs, resource.PropertyKey("networkRuleSet"))
	// Input "deny" was not overwritten with default "allow"
	assert.Equal(t, "Deny", inputs["networkRuleSet"].ObjectValue()["defaultAction"].StringValue())
}

func TestRestoreDefaultInputsIsNoopWithoutDefaultProperties(t *testing.T) {
	inputs := resource.PropertyMap{}

	olds := resource.PropertyMap{
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{}),
	}

	res := resources.AzureAPIResource{} // no defaults

	err := restoreDefaultInputsForRemovedProperties(inputs, res, olds)
	assert.NoError(t, err)
	assert.Empty(t, inputs)

	// same with empty defaults
	res.DefaultProperties = map[string]interface{}{}
	err = restoreDefaultInputsForRemovedProperties(inputs, res, olds)
	assert.NoError(t, err)
	assert.Empty(t, inputs)
}

func TestMappableOldStateIsNoopWithoutDefaults(t *testing.T) {
	res := resources.AzureAPIResource{} // no defaults
	m := map[string]interface{}{"foo": "bar"}
	removeDefaults(res, m)
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, m)
}

func TestMappableOldStatePreservesNonDefaults(t *testing.T) {
	res := resources.AzureAPIResource{
		DefaultProperties: map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"defaultAction": "Allow",
			},
		},
	}
	m := map[string]any{
		"networkRuleSet": map[string]any{
			"defaultAction": "Deny",
		},
	}
	removeDefaults(res, m)
	assert.Equal(t, "Deny", m["networkRuleSet"].(map[string]interface{})["defaultAction"])
}

func TestMappableOldStateRemovesDefaultsThatWereInputs(t *testing.T) {
	res := resources.AzureAPIResource{
		DefaultProperties: map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"defaultAction": "Allow",
			},
		},
	}
	m := map[string]any{
		"__inputs": map[string]any{
			"networkRuleSet": map[string]any{
				"defaultAction": "Allow",
			},
		},
		"networkRuleSet": map[string]any{
			"defaultAction": "Allow",
		},
	}
	removeDefaults(res, m)
	assert.Contains(t, m, "__inputs")
	assert.NotContains(t, m, "networkRuleSet")
}

func TestMappableOldStatePreservesDefaultsThatWereNotInputs(t *testing.T) {
	res := resources.AzureAPIResource{
		DefaultProperties: map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"defaultAction": "Allow",
			},
		},
	}
	m := map[string]any{
		"__inputs": map[string]any{},
		"networkRuleSet": map[string]any{
			"defaultAction": "Allow",
		},
	}
	removeDefaults(res, m)
	assert.Contains(t, m, "__inputs")
	assert.Contains(t, m, "networkRuleSet")
}

func TestResetUnsetSubResourceProperties(t *testing.T) {
	ctx := context.Background()

	res, provider := setUpResourceWithRefAndProviderWithTypeLookup("")

	t.Run("empty", func(t *testing.T) {
		empty := &resources.AzureAPIResource{}
		oldInputs := resource.PropertyMap{}
		sdkResponse := map[string]any{}
		actual := provider.resetUnsetSubResourceProperties(ctx, "urn", sdkResponse, oldInputs, empty)
		expected := map[string]any{}
		assert.Equal(t, expected, actual)
	})

	t.Run("remove", func(t *testing.T) {
		oldInputs := resource.PropertyMap{}
		sdkResponse := map[string]any{
			"properties": map[string]any{
				"accessPolicies": []any{
					"a policy",
				},
			},
		}
		actual := provider.resetUnsetSubResourceProperties(ctx, "urn", sdkResponse, oldInputs, res)
		expected := map[string]any{
			"properties": map[string]any{
				"accessPolicies": []any{},
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("preserve", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			resource.PropertyKey("properties"): resource.NewObjectProperty(resource.PropertyMap{
				resource.PropertyKey("accessPolicies"): resource.NewArrayProperty([]resource.PropertyValue{}),
			}),
		}
		sdkResponse := map[string]any{
			"properties": map[string]any{
				"accessPolicies": []any{},
			},
		}
		actual := provider.resetUnsetSubResourceProperties(ctx, "urn", sdkResponse, oldInputs, res)
		expected := sdkResponse
		assert.Equal(t, expected, actual)
	})
}

// Helper to avoid repeating the same setup code in multiple tests. Returns a resource with a
// "properties" property of type azure-native:keyvault:VaultProperties, which the returned provider
// will return when asked to look up that type.
func setUpResourceWithRefAndProviderWithTypeLookup(version string) (*resources.AzureAPIResource, *azureNativeProvider) {
	res := resources.AzureAPIResource{
		APIVersion: "v20241101",
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"properties": {
							Type:       "object",
							Ref:        "#/types/azure-native:keyvault:VaultProperties",
							Containers: []string{"container"},
						},
					},
				},
			},
		},
	}

	// Setup resource metadata
	provider := azureNativeProvider{
		version: version,
		resourceMap: &resources.PartialAzureAPIMetadata{
			Resources: func() resources.PartialMap[resources.AzureAPIResource] {
				m := resources.NewPartialMap[resources.AzureAPIResource]()
				b, _ := json.Marshal(map[string]any{
					"azure-native:keyvault:Vault": res,
				})
				_ = m.UnmarshalJSON(b)
				return m
			}(),
		},

		// Mock the type lookup to only return the type referenced in the resource above
		lookupType: func(ref string) (*resources.AzureAPIType, bool, error) {
			if ref == "#/types/azure-native:keyvault:VaultProperties" {
				return &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"accessPolicies": {
							Type: "array",
							Items: &resources.AzureAPIProperty{
								Type: "object",
								Ref:  "#/types/azure-native:keyvault:AccessPolicyEntry",
							},
							MaintainSubResourceIfUnset: true,
						},
					},
				}, true, nil
			}
			if ref == "#/types/azure-native:keyvault:AccessPolicyEntry" {
				return &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"permissions": {
							Type: "array",
							Items: &resources.AzureAPIProperty{
								Type: "string",
							},
							Containers: []string{"container2", "container3"},
						}},
				}, true, nil
			}
			return nil, false, nil
		},
	}

	return &res, &provider
}

func TestSetUnsetSubresourcePropertiesToDefaults(t *testing.T) {
	res, provider := setUpResourceWithRefAndProviderWithTypeLookup("")

	t.Run("unchanged", func(t *testing.T) {
		body := map[string]any{
			"container": map[string]any{
				"properties": map[string]any{
					"accessPolicies": []any{},
				},
			},
		}
		provider.setUnsetSubresourcePropertiesToDefaults(*res, body, body, true)
		assert.Equal(t, map[string]any{
			"container": map[string]any{
				"properties": map[string]any{
					"accessPolicies": []any{},
				},
			},
		}, body)
	})

	t.Run("simple missing", func(t *testing.T) {
		body := map[string]any{
			"container": map[string]any{
				"properties": map[string]any{},
			},
		}
		provider.setUnsetSubresourcePropertiesToDefaults(*res, body, body, true)
		assert.Equal(t, map[string]any{
			"container": map[string]any{
				"properties": map[string]any{
					"accessPolicies": []any{},
				},
			},
		}, body)
	})

	t.Run("nested missing", func(t *testing.T) {
		body := map[string]any{}
		provider.setUnsetSubresourcePropertiesToDefaults(*res, body, body, true)
		assert.Equal(t, map[string]any{
			"container": map[string]any{
				"properties": map[string]any{
					"accessPolicies": []any{},
				},
			},
		}, body)
	})

	t.Run("nested missing in SDK shape", func(t *testing.T) {
		body := map[string]any{}
		provider.setUnsetSubresourcePropertiesToDefaults(*res, body, body, false)
		assert.Equal(t, map[string]any{
			"properties": map[string]any{
				"accessPolicies": []any{},
			},
		}, body)
	})
}

func TestInvokeResponseToOutputs(t *testing.T) {
	invoke := resources.AzureAPIInvoke{
		APIVersion: "v20241101",
	}
	conv := convert.NewSdkShapeConverterFull(map[string]resources.AzureAPIType{})

	for _, val := range []any{
		"string",
		42,
		[]string{"a", "b"},
	} {
		t.Run(fmt.Sprintf("single value of type %T", val), func(t *testing.T) {
			p := azureNativeProvider{converter: &conv}
			outputs := p.invokeResponseToOutputs(val, invoke)
			require.Len(t, outputs, 1)
			require.Contains(t, outputs, resources.SingleValueProperty)
			assert.Equal(t, val, outputs[resources.SingleValueProperty])
		})
	}

	t.Run("object", func(t *testing.T) {
		p := azureNativeProvider{converter: &conv}
		outputs := p.invokeResponseToOutputs(map[string]any{"key": "value"}, invoke)
		assert.Empty(t, outputs) // the empty converter doesn't know any properties
	})

	t.Run("special case: GetResource", func(t *testing.T) {
		invoke := resources.AzureAPIInvoke{
			APIVersion:  "v20241101",
			GetResource: true,
		}
		t.Run("azureApiVersion", func(t *testing.T) {
			p := azureNativeProvider{converter: &conv, version: "2.0.0"}
			outputs := p.invokeResponseToOutputs(map[string]any{}, invoke)
			require.NotContains(t, outputs, "azureApiVersion")

			p = azureNativeProvider{converter: &conv, version: "3.0.0"}
			outputs = p.invokeResponseToOutputs(map[string]any{}, invoke)
			require.Contains(t, outputs, "azureApiVersion")
			assert.Equal(t, "v20241101", outputs["azureApiVersion"].(resource.PropertyValue).StringValue())
		})
	})
}

func TestReader(t *testing.T) {
	t.Run("custom Read", func(t *testing.T) {
		var customReads []string
		customRes := &customresources.CustomResource{
			Read: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error) {
				customReads = append(customReads, id)
				return map[string]any{}, true, nil
			},
		}

		azureClient := &az.MockAzureClient{}
		crudClient := crud.NewResourceCrudClient(azureClient, nil, nil, "123", nil)

		r := reader(customRes, crudClient)
		_, err := r(context.Background(), "id1", nil)
		require.NoError(t, err)
		assert.Equal(t, []string{"id1"}, customReads)
		assert.Empty(t, azureClient.GetIds)
	})

	t.Run("no custom Read", func(t *testing.T) {
		resource := &resources.AzureAPIResource{
			Response: map[string]resources.AzureAPIProperty{},
		}

		for _, otherCustomRes := range []*customresources.CustomResource{nil, {} /* custom resource that doesn't implement Read */} {
			azureClient := &az.MockAzureClient{}
			crudClient := crud.NewResourceCrudClient(azureClient, nil, nil, "123", resource)

			r := reader(otherCustomRes, crudClient)
			_, err := r(context.Background(), "id2", nil)
			require.NoError(t, err)
			assert.Contains(t, azureClient.GetIds, "id2")
		}
	})
}

func TestReadAfterWrite(t *testing.T) {
	read := false
	var reader readFunc = func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
		read = true
		return nil, nil
	}

	for _, skipReadOnUpdate := range []bool{true, false} {
		p := azureNativeProvider{
			skipReadOnUpdate: skipReadOnUpdate,
		}
		p.readAfterWrite(context.Background(), "id", "urn", "create", resource.PropertyMap{}, reader)
		assert.Equal(t, !skipReadOnUpdate, read)
	}
}

func TestUsesCorrectAzureClient(t *testing.T) {
	p := azureNativeProvider{}

	t.Run("default", func(t *testing.T) {
		t.Setenv("PULUMI_ENABLE_AZCORE_BACKEND", "")
		client, err := p.newAzureClient(nil, &fake.TokenCredential{}, "pulumi")
		require.NoError(t, err)
		assert.Equal(t, "azureClientImpl", reflect.TypeOf(client).Elem().Name())
	})

	t.Run("Autorest and legacy auth disabled explicitly", func(t *testing.T) {
		t.Setenv("PULUMI_ENABLE_AZCORE_BACKEND", "false")
		client, err := p.newAzureClient(nil, &fake.TokenCredential{}, "pulumi")
		require.NoError(t, err)
		assert.Equal(t, "azureClientImpl", reflect.TypeOf(client).Elem().Name())
	})

	t.Run("Azcore enabled", func(t *testing.T) {
		t.Setenv("PULUMI_ENABLE_AZCORE_BACKEND", "true")
		client, err := p.newAzureClient(nil, &fake.TokenCredential{}, "pulumi")
		require.NoError(t, err)
		assert.Equal(t, "azCoreClient", reflect.TypeOf(client).Elem().Name())
	})
}

func TestAzcoreAzureClientUsesCorrectCloud(t *testing.T) {
	if !util.EnableAzcoreBackend() {
		t.Skip()
	}

	for expectedHost, cloudInstance := range map[string]cloud.Configuration{
		"https://management.azure.com":         cloud.AzurePublic,
		"https://management.chinacloudapi.cn":  cloud.AzureChina,
		"https://management.usgovcloudapi.net": cloud.AzureGovernment,
	} {
		p := azureNativeProvider{
			cloud: cloudInstance,
		}

		client, err := p.newAzureClient(nil, &fake.TokenCredential{}, "pulumi")
		require.NoError(t, err)
		require.NotNil(t, client)

		// Use reflection to get the value of the private 'host' field
		clientValue := reflect.ValueOf(client).Elem()
		hostField := clientValue.FieldByName("host")
		require.True(t, hostField.IsValid(), "host field should be valid (%s)", expectedHost)

		assert.Equal(t, expectedHost, hostField.String())
	}
}

func TestAutorestAzureClientUsesCorrectCloud(t *testing.T) {
	for expectedEnv, environment := range map[string]azure.Environment{
		azure.PublicCloud.Name:       azure.PublicCloud,
		azure.ChinaCloud.Name:        azure.ChinaCloud,
		azure.USGovernmentCloud.Name: azure.USGovernmentCloud,
	} {
		p := azureNativeProvider{
			environment: environment,
		}
		t.Setenv("PULUMI_ENABLE_AZCORE_BACKEND", "false")

		client, err := p.newAzureClient(nil, nil, "pulumi")
		require.NoError(t, err)
		require.NotNil(t, client)

		// Use reflection to get the value of the private 'environment' field
		clientValue := reflect.ValueOf(client).Elem()
		environmentField := clientValue.FieldByName("environment")
		require.True(t, environmentField.IsValid(), "environment field should be valid")
		nameField := environmentField.FieldByName("Name")
		require.True(t, nameField.IsValid(), "environment.name field should be valid")

		assert.Equal(t, expectedEnv, nameField.String())
	}
}

func TestGetTokenEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("explicit", func(t *testing.T) {
		t.Parallel()
		p := azureNativeProvider{}
		endpoint := p.tokenEndpoint(resource.NewStringProperty("https://management.azure.com/"))
		assert.Equal(t, "https://management.azure.com/", endpoint)
	})

	t.Run("implicit public", func(t *testing.T) {
		t.Parallel()
		p := azureNativeProvider{
			environment: azure.PublicCloud,
		}
		endpoint := p.tokenEndpoint(resource.NewNullProperty())
		assert.Equal(t, "https://management.azure.com/", endpoint)
	})

	t.Run("implicit usgov", func(t *testing.T) {
		t.Parallel()
		p := azureNativeProvider{
			environment: azure.USGovernmentCloud,
		}
		endpoint := p.tokenEndpoint(resource.NewNullProperty())
		assert.Equal(t, "https://management.usgovcloudapi.net/", endpoint)
	})

	t.Run("implicit with empty string, public", func(t *testing.T) {
		t.Parallel()
		p := azureNativeProvider{
			environment: azure.PublicCloud,
		}
		endpoint := p.tokenEndpoint(resource.NewStringProperty(""))
		assert.Equal(t, "https://management.azure.com/", endpoint)
	})
}

func TestGetTokenRequestOpts(t *testing.T) {
	t.Parallel()

	opts := tokenRequestOpts("http://endpoint")
	assert.Empty(t, opts.Claims)
	assert.Empty(t, opts.TenantID)
	assert.Equal(t, []string{"http://endpoint/.default"}, opts.Scopes)
}

func TestCustomCreate(t *testing.T) {
	t.Parallel()

	t.Run("resource doesn't exist, uses the custom resource's CanCreate", func(t *testing.T) {
		t.Parallel()

		calledCreate := false
		customRes := &customresources.CustomResource{
			CanCreate: func(ctx context.Context, id string) error {
				return nil
			},
			Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
				calledCreate = true
				return map[string]any{}, nil
			},
		}

		_, err := customCreate(context.Background(), resource.PropertyMap{}, "id", nil, customRes)
		require.NoError(t, err)
		assert.True(t, calledCreate)
	})

	t.Run("resource does exist, uses the custom resource's CanCreate", func(t *testing.T) {
		t.Parallel()

		calledCanCreate := false
		calledCreate := false
		calledRead := false
		customRes := &customresources.CustomResource{
			CanCreate: func(ctx context.Context, id string) error {
				calledCanCreate = true
				return fmt.Errorf("resource already exists")
			},
			Read: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error) {
				calledRead = true
				return map[string]any{}, true, nil
			},
			Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
				calledCreate = true
				return map[string]any{}, nil
			},
		}

		_, err := customCreate(context.Background(), resource.PropertyMap{}, "id", nil, customRes)
		require.Error(t, err)
		assert.True(t, calledCanCreate)
		assert.False(t, calledRead)
		assert.False(t, calledCreate)
	})

	t.Run("resource doesn't exist, uses the custom resource's Read", func(t *testing.T) {
		t.Parallel()

		azureClient := &az.MockAzureClient{}
		crudClient := crud.NewResourceCrudClient(azureClient, nil, nil, "123", nil)

		calledCreate := false
		customRes := &customresources.CustomResource{
			Read: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error) {
				return map[string]any{}, false, nil
			},
			Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
				calledCreate = true
				return map[string]any{}, nil
			},
		}

		_, err := customCreate(context.Background(), resource.PropertyMap{}, "id", crudClient, customRes)
		require.NoError(t, err)
		assert.True(t, calledCreate)
	})

	t.Run("resource does exist, uses the custom resource's Read", func(t *testing.T) {
		t.Parallel()

		azureClient := &az.MockAzureClient{}
		crudClient := crud.NewResourceCrudClient(azureClient, nil, nil, "123", nil)

		calledCreate := false
		customRes := &customresources.CustomResource{
			Read: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error) {
				return map[string]any{}, true, nil
			},
			Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
				calledCreate = true
				return map[string]any{}, nil
			},
		}

		_, err := customCreate(context.Background(), resource.PropertyMap{}, "id", crudClient, customRes)
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot create already existing resource")
		assert.False(t, calledCreate)
	})

	t.Run("resource doesn't exist, uses the regular CrudClient if custom resource has neither Read nor CanCreate", func(t *testing.T) {
		t.Parallel()

		azureClient := &az.MockAzureClient{}
		crudClient := crud.NewResourceCrudClient(azureClient, nil, nil, "123", &resources.AzureAPIResource{})

		calledCreate := false
		customRes := &customresources.CustomResource{
			Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
				calledCreate = true
				return map[string]any{}, nil
			},
		}

		_, err := customCreate(context.Background(), resource.PropertyMap{}, "id", crudClient, customRes)
		require.NoError(t, err)
		assert.True(t, calledCreate)
	})
}

func TestCheckpointObject(t *testing.T) {
	t.Parallel()

	res, _ := setUpResourceWithRefAndProviderWithTypeLookup("")

	t.Run("stores inputs in v2", func(t *testing.T) {
		t.Parallel()

		inputs := resource.PropertyMap{
			"name": resource.NewStringProperty("test"),
		}
		outputs := map[string]any{}

		checkpoint := checkpointObjectVersioned(res, inputs, outputs, semver.MustParse("2.0.0"))
		assert.Contains(t, checkpoint, resource.PropertyKey("__inputs"))
	})

	t.Run("does not store inputs in v3", func(t *testing.T) {
		t.Parallel()

		inputs := resource.PropertyMap{
			"name": resource.NewStringProperty("test"),
		}
		outputs := map[string]any{}

		checkpoint := checkpointObjectVersioned(res, inputs, outputs, semver.MustParse("3.0.0"))
		assert.NotContains(t, checkpoint, resource.PropertyKey("__inputs"))
	})

	t.Run("preserves original state", func(t *testing.T) {
		t.Parallel()

		inputs := resource.PropertyMap{
			"name": resource.NewStringProperty("test"),
		}
		outputs := map[string]any{
			customresources.OriginalStateKey: resource.NewStringProperty("original state"),
		}

		checkpoint := checkpointObjectVersioned(res, inputs, outputs, semver.MustParse("2.0.0"))
		assert.Contains(t, checkpoint, resource.PropertyKey(customresources.OriginalStateKey))
		assert.True(t, checkpoint.ContainsSecrets())

		checkpoint = checkpointObjectVersioned(res, inputs, outputs, semver.MustParse("3.0.0"))
		assert.Contains(t, checkpoint, resource.PropertyKey(customresources.OriginalStateKey))
		assert.True(t, checkpoint.ContainsSecrets())
	})

	t.Run("produces azureApiVersion", func(t *testing.T) {
		inputs := resource.PropertyMap{}
		outputs := map[string]interface{}{}

		checkpoint := checkpointObjectVersioned(res, inputs, outputs, semver.MustParse("2.0.0"))
		assert.Contains(t, checkpoint, resource.PropertyKey("azureApiVersion"))
		assert.Equal(t, "v20241101", checkpoint["azureApiVersion"].StringValue())
	})
}

func TestApplyDefaults(t *testing.T) {
	ctx := context.Background()

	t.Run("azureApiVersion", func(t *testing.T) {
		olds := resource.PropertyMap{}
		news := resource.PropertyMap{}

		res, provider := setUpResourceWithRefAndProviderWithTypeLookup("2.0.0")
		provider.applyDefaults(ctx, "urn", []byte{}, nil, *res, olds, news)
		assert.NotContains(t, news, resource.PropertyKey("azureApiVersion"))

		res, provider = setUpResourceWithRefAndProviderWithTypeLookup("3.0.0")
		provider.applyDefaults(ctx, "urn", []byte{}, nil, *res, olds, news)
		assert.Contains(t, news, resource.PropertyKey("azureApiVersion"))
		assert.Equal(t, "v20241101", news["azureApiVersion"].StringValue())
	})
}

func TestDiff(t *testing.T) {
	ctx := context.Background()

	type args struct {
		olds      resource.PropertyMap
		oldInputs resource.PropertyMap
		news      resource.PropertyMap
	}
	type want struct {
		changes      rpc.DiffResponse_DiffChanges
		detailedDiff map[string]*rpc.PropertyDiff
	}
	tests := []struct {
		name    string
		version string
		args    args
		want    want
		wantErr bool
	}{
		{
			name:    "azureApiVersion_v2_olds",
			version: "2.0.0",
			args: args{
				olds: resource.PropertyMap{
					// v2.90+: output property
					"azureApiVersion": resource.NewStringProperty("v20241101"),
				},
				oldInputs: resource.PropertyMap{},
				news:      resource.PropertyMap{},
			},
			want: want{
				changes: rpc.DiffResponse_DIFF_NONE,
			},
		},
		{
			name:    "azureApiVersion_v2-to-v3_same",
			version: "3.0.0",
			args: args{
				olds: resource.PropertyMap{
					// v2.90+: output property
					"azureApiVersion": resource.NewStringProperty("v20241101"),
				},
				oldInputs: resource.PropertyMap{},
				news: resource.PropertyMap{
					// v3: input property
					"azureApiVersion": resource.NewStringProperty("v20241101"),
				},
			},
			want: want{
				changes: rpc.DiffResponse_DIFF_NONE,
			},
		},
		{
			name:    "azureApiVersion_v2-to-v3_add",
			version: "3.0.0",
			args: args{
				olds:      resource.PropertyMap{},
				oldInputs: resource.PropertyMap{},
				news: resource.PropertyMap{
					// v3: input property
					"azureApiVersion": resource.NewStringProperty("v20241101"),
				},
			},
			want: want{
				changes: rpc.DiffResponse_DIFF_SOME,
				detailedDiff: map[string]*rpc.PropertyDiff{
					"azureApiVersion": {
						Kind:      rpc.PropertyDiff_ADD,
						InputDiff: true,
					},
				},
			},
		},
		{
			name:    "azureApiVersion_v2-to-v3_update",
			version: "3.0.0",
			args: args{
				olds: resource.PropertyMap{
					// v2.90+: output property
					"azureApiVersion": resource.NewStringProperty("v20241101"),
				},
				oldInputs: resource.PropertyMap{},
				news: resource.PropertyMap{
					// v3: input property
					"azureApiVersion": resource.NewStringProperty("v20251101-preview"),
				},
			},
			want: want{
				changes: rpc.DiffResponse_DIFF_SOME,
				detailedDiff: map[string]*rpc.PropertyDiff{
					"azureApiVersion": {
						Kind:      rpc.PropertyDiff_UPDATE,
						InputDiff: false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, provider := setUpResourceWithRefAndProviderWithTypeLookup(tt.version)

			olds, _ := plugin.MarshalProperties(tt.args.olds, plugin.MarshalOptions{})
			oldInputs, _ := plugin.MarshalProperties(tt.args.oldInputs, plugin.MarshalOptions{})
			news, _ := plugin.MarshalProperties(tt.args.news, plugin.MarshalOptions{})

			req := &rpc.DiffRequest{
				Id:        "id",
				Urn:       "urn:pulumi:stack::project::azure-native:keyvault:Vault::id",
				Olds:      olds,
				OldInputs: oldInputs,
				News:      news,
			}
			resp, err := provider.Diff(ctx, req)
			require.NoError(t, err)
			assert.Equal(t, tt.want.changes, resp.Changes, "expected no changes")
			if tt.want.detailedDiff != nil {
				assert.True(t, resp.HasDetailedDiff)
				assert.EqualExportedValues(t, tt.want.detailedDiff, resp.DetailedDiff, "unexpected detailed diff")
			}
		})
	}
}
