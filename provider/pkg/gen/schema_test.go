// Copyright 2016-2020, Pulumi Corporation.

package gen

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestTypeAliasFormatting(t *testing.T) {
	generator := packageGenerator{
		pkg:        &pschema.PackageSpec{Name: "azure-native"},
		apiVersion: "v20220222",
	}

	actual := generator.makeTypeAlias("Compute", "VirtualMachine", "v18851225")
	assert.NotNil(t, actual)
	assert.Equal(t, "azure-native:compute/v18851225:VirtualMachine", *actual.Type)

	actual = generator.makeTypeAlias("Compute", "VirtualMachine", "")
	assert.NotNil(t, actual)
	assert.Equal(t, "azure-native:compute:VirtualMachine", *actual.Type)
}

func TestAliases(t *testing.T) {
	generator := packageGenerator{
		pkg:        &pschema.PackageSpec{Name: "azure-native"},
		apiVersion: "v20220222",
	}

	resource := &resourceVariant{
		ResourceSpec: &openapi.ResourceSpec{
			CompatibleVersions: []string{"v20210111"},
		},
		typeName: "PrivateLinkForAzureAd",
	}

	aliases := generator.generateAliases("Insights", resource, "privateLinkForAzureAd")
	actual := codegen.NewStringSet()
	for _, alias := range aliases {
		actual.Add(*alias.Type)
	}
	expected := codegen.NewStringSet(
		"azure-native:insights/v20210111:privateLinkForAzureAd",
		"azure-native:insights/v20210111:PrivateLinkForAzureAd",
		"azure-native:insights/v20220222:privateLinkForAzureAd",
	)
	assert.Equal(t, expected, actual)
}
