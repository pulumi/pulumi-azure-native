// Copyright 2016-2024, Pulumi Corporation.

package provider

import (
	"context"
	"testing"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
)

func TestGetDefaultName(t *testing.T) {
	t.Parallel()

	p := &azureNativeProvider{}
	const testUrn = "urn:pulumi:stack::project::azure-native:storage:StorageAccount::test"
	testSeed := []byte("seed")

	t.Run("returns old name if exists", func(t *testing.T) {
		oldName := resource.NewStringProperty("existing-name")
		olds := resource.PropertyMap{
			"name": oldName,
		}

		result, randomlyNamed, ok := p.getDefaultName(
			testUrn,
			testSeed,
			nil,
			resources.AutoNameRandom,
			"name",
			olds,
		)

		assert.True(t, ok)
		assert.False(t, randomlyNamed)
		assert.Equal(t, &oldName, result)
	})

	t.Run("returns old name and createBeforeDelete if flag exists", func(t *testing.T) {
		oldName := resource.NewStringProperty("existing-name")
		olds := resource.PropertyMap{
			"name":                 oldName,
			createBeforeDeleteFlag: resource.NewBoolProperty(true),
		}

		result, randomlyNamed, ok := p.getDefaultName(
			testUrn,
			testSeed,
			nil,
			resources.AutoNameRandom,
			"name",
			olds,
		)

		assert.True(t, ok)
		assert.True(t, randomlyNamed)
		assert.Equal(t, &oldName, result)
	})

	t.Run("fails if auto-naming is disabled", func(t *testing.T) {
		result, _, ok := p.getDefaultName(
			testUrn,
			testSeed,
			&pulumirpc.CheckRequest_AutonamingOptions{
				Mode: pulumirpc.CheckRequest_AutonamingOptions_DISABLE,
			},
			resources.AutoNameRandom,
			"name",
			resource.PropertyMap{},
		)

		assert.False(t, ok)
		assert.Nil(t, result)
	})

	t.Run("returns proposed name for enforce mode", func(t *testing.T) {
		result, randomlyNamed, ok := p.getDefaultName(
			testUrn,
			testSeed,
			&pulumirpc.CheckRequest_AutonamingOptions{
				Mode:         pulumirpc.CheckRequest_AutonamingOptions_ENFORCE,
				ProposedName: "proposed-uuid-name",
			},
			resources.AutoNameUuid, // UUID would otherwise take priority over proposed name
			"name",
			resource.PropertyMap{},
		)

		assert.True(t, ok)
		assert.True(t, randomlyNamed)
		assert.Equal(t, "proposed-uuid-name", result.StringValue())
	})

	t.Run("returns URN name for copy strategy without auto-naming options", func(t *testing.T) {
		result, randomlyNamed, ok := p.getDefaultName(
			testUrn,
			testSeed,
			nil,
			resources.AutoNameCopy,
			"name",
			resource.PropertyMap{},
		)

		assert.True(t, ok)
		assert.False(t, randomlyNamed)
		assert.Equal(t, "test", result.StringValue())
	})

	t.Run("returns random name for random strategy without auto-naming options", func(t *testing.T) {
		result, randomlyNamed, ok := p.getDefaultName(
			testUrn,
			testSeed,
			nil,
			resources.AutoNameRandom,
			"name",
			resource.PropertyMap{},
		)

		assert.True(t, ok)
		assert.True(t, randomlyNamed)
		assert.Equal(t, "test614742c2", result.StringValue())
	})

	t.Run("returns proposed name for random strategy with propose mode", func(t *testing.T) {
		result, randomlyNamed, ok := p.getDefaultName(
			testUrn,
			testSeed,
			&pulumirpc.CheckRequest_AutonamingOptions{
				Mode:         pulumirpc.CheckRequest_AutonamingOptions_PROPOSE,
				ProposedName: "proposed-random-name",
			},
			resources.AutoNameRandom,
			"name",
			resource.PropertyMap{},
		)

		assert.True(t, ok)
		assert.True(t, randomlyNamed)
		assert.Equal(t, "proposed-random-name", result.StringValue())
	})

	t.Run("returns proposed name for copy strategy with propose mode", func(t *testing.T) {
		result, randomlyNamed, ok := p.getDefaultName(
			testUrn,
			testSeed,
			&pulumirpc.CheckRequest_AutonamingOptions{
				Mode:         pulumirpc.CheckRequest_AutonamingOptions_PROPOSE,
				ProposedName: "proposed-copy-name",
			},
			resources.AutoNameCopy,
			"name",
			resource.PropertyMap{},
		)

		assert.True(t, ok)
		assert.True(t, randomlyNamed)
		assert.Equal(t, "proposed-copy-name", result.StringValue())
	})

	t.Run("returns UUID for UUID strategy with propose mode", func(t *testing.T) {
		result, randomlyNamed, ok := p.getDefaultName(
			testUrn,
			testSeed,
			&pulumirpc.CheckRequest_AutonamingOptions{
				Mode:         pulumirpc.CheckRequest_AutonamingOptions_PROPOSE,
				ProposedName: "proposed-uuid-name",
			},
			resources.AutoNameUuid,
			"name",
			resource.PropertyMap{},
		)

		assert.True(t, ok)
		assert.True(t, randomlyNamed)
		assert.Equal(t, "1ec836c1-4f8e-3f38-4557-7ed3bc41c8ad", result.StringValue())
	})
}

func TestCheck_AutoNaming(t *testing.T) {
	t.Parallel()
	p := &azureNativeProvider{}
	const testUrn = "urn:pulumi:stack::project::azure-native:storage:StorageAccount::test"

	t.Run("applies auto-naming when enabled", func(t *testing.T) {
		// Setup resource metadata
		p.resourceMap = &resources.PartialAzureAPIMetadata{
			Resources: func() resources.PartialMap[resources.AzureAPIResource] {
				m := resources.NewPartialMap[resources.AzureAPIResource]()
				m.UnmarshalJSON([]byte(`{
						"azure-native:storage:StorageAccount": {
							"apiVersion": "2021-01-01",
							"PUT": [{
								"name": "name",
								"value": {
									"autoName": "random"
								}
							}]
						}
					}`))
				return m
			}(),
		}

		// Call Check with auto-naming options
		resp, err := p.Check(context.Background(), &pulumirpc.CheckRequest{
			Urn:        testUrn,
			News:       &structpb.Struct{}, // Empty inputs
			Olds:       &structpb.Struct{}, // Empty old state
			RandomSeed: []byte("seed"),
			Autonaming: &pulumirpc.CheckRequest_AutonamingOptions{
				Mode:         pulumirpc.CheckRequest_AutonamingOptions_ENFORCE,
				ProposedName: "proposed-name",
			},
		})

		// Verify the response
		assert.NoError(t, err)
		inputs, err := plugin.UnmarshalProperties(resp.Inputs, plugin.MarshalOptions{})
		assert.NoError(t, err)
		assert.Equal(t, "proposed-name", inputs["name"].StringValue())
		assert.True(t, inputs[createBeforeDeleteFlag].BoolValue())
	})

	t.Run("returns error when auto-naming is disabled", func(t *testing.T) {
		// Setup resource metadata
		p.resourceMap = &resources.PartialAzureAPIMetadata{
			Resources: func() resources.PartialMap[resources.AzureAPIResource] {
				m := resources.NewPartialMap[resources.AzureAPIResource]()
				m.UnmarshalJSON([]byte(`{
						"azure-native:storage:StorageAccount": {
							"apiVersion": "2021-01-01",
							"PUT": [{
								"name": "name",
								"required": true,
								"value": {
									"autoName": "random"
								}
							}]
						}
					}`))
				return m
			}(),
		}

		// Call Check with auto-naming disabled
		resp, err := p.Check(context.Background(), &pulumirpc.CheckRequest{
			Urn:        testUrn,
			News:       &structpb.Struct{}, // Empty inputs
			Olds:       &structpb.Struct{}, // Empty old state
			RandomSeed: []byte("seed"),
			Autonaming: &pulumirpc.CheckRequest_AutonamingOptions{
				Mode: pulumirpc.CheckRequest_AutonamingOptions_DISABLE,
			},
		})

		// Verify the error
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Failures, 1)
		assert.Contains(t, resp.Failures[0].Reason, "missing required property 'name'")
	})
}
