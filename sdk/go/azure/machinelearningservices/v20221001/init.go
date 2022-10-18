


package v20221001

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:machinelearningservices/v20221001:BatchDeployment":
		r = &BatchDeployment{}
	case "azure-native:machinelearningservices/v20221001:BatchEndpoint":
		r = &BatchEndpoint{}
	case "azure-native:machinelearningservices/v20221001:CodeContainer":
		r = &CodeContainer{}
	case "azure-native:machinelearningservices/v20221001:CodeVersion":
		r = &CodeVersion{}
	case "azure-native:machinelearningservices/v20221001:ComponentContainer":
		r = &ComponentContainer{}
	case "azure-native:machinelearningservices/v20221001:ComponentVersion":
		r = &ComponentVersion{}
	case "azure-native:machinelearningservices/v20221001:Compute":
		r = &Compute{}
	case "azure-native:machinelearningservices/v20221001:DataContainer":
		r = &DataContainer{}
	case "azure-native:machinelearningservices/v20221001:DataVersion":
		r = &DataVersion{}
	case "azure-native:machinelearningservices/v20221001:Datastore":
		r = &Datastore{}
	case "azure-native:machinelearningservices/v20221001:EnvironmentContainer":
		r = &EnvironmentContainer{}
	case "azure-native:machinelearningservices/v20221001:EnvironmentVersion":
		r = &EnvironmentVersion{}
	case "azure-native:machinelearningservices/v20221001:Job":
		r = &Job{}
	case "azure-native:machinelearningservices/v20221001:ModelContainer":
		r = &ModelContainer{}
	case "azure-native:machinelearningservices/v20221001:ModelVersion":
		r = &ModelVersion{}
	case "azure-native:machinelearningservices/v20221001:OnlineDeployment":
		r = &OnlineDeployment{}
	case "azure-native:machinelearningservices/v20221001:OnlineEndpoint":
		r = &OnlineEndpoint{}
	case "azure-native:machinelearningservices/v20221001:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:machinelearningservices/v20221001:Schedule":
		r = &Schedule{}
	case "azure-native:machinelearningservices/v20221001:Workspace":
		r = &Workspace{}
	case "azure-native:machinelearningservices/v20221001:WorkspaceConnection":
		r = &WorkspaceConnection{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"machinelearningservices/v20221001",
		&module{version},
	)
}
