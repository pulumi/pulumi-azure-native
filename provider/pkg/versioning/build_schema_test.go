package versioning

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestDrop(t *testing.T) {
	pkgSpec := schema.PackageSpec{
		Resources: map[string]schema.ResourceSpec{
			"azure-native:machinelearningservices/v20200515preview:ACIService":    {},
			"azure-native:machinelearningservices/v20200515:ACIService":           {},
			"azure-native:machinelearningservices/v20200515preview:SomethingElse": {},
		},
		Functions: map[string]schema.FunctionSpec{
			"azure-native:machinelearningservices/v20200515preview:getACIService":    {},
			"azure-native:machinelearningservices/v20200515:getACIService":           {},
			"azure-native:machinelearningservices/v20200515preview:getSomethingElse": {},
		},
	}

	removable := ResourceRemovals{
		"azure-native:machinelearningservices/v20200515preview:ACIService": "azure-native:machinelearningservices/v20210401:ACIService",
		"azure-native:machinelearningservices/v20200515preview:AKSService": "azure-native:machinelearningservices/v20210401:AKSService",
	}

	dropFromSchema(&pkgSpec, removable)

	assert.NotContains(t, pkgSpec.Resources, "azure-native:machinelearningservices/v20200515preview:ACIService")
	assert.Contains(t, pkgSpec.Resources, "azure-native:machinelearningservices/v20200515:ACIService")
	assert.Contains(t, pkgSpec.Resources, "azure-native:machinelearningservices/v20200515preview:SomethingElse")

	assert.NotContains(t, pkgSpec.Functions, "azure-native:machinelearningservices/v20200515preview:getACIService")
	assert.Contains(t, pkgSpec.Functions, "azure-native:machinelearningservices/v20200515:getACIService")
	assert.Contains(t, pkgSpec.Functions, "azure-native:machinelearningservices/v20200515preview:getSomethingElse")
}
