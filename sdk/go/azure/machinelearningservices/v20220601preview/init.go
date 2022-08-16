


package v20220601preview

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
	case "azure-native:machinelearningservices/v20220601preview:BatchDeployment":
		r = &BatchDeployment{}
	case "azure-native:machinelearningservices/v20220601preview:BatchEndpoint":
		r = &BatchEndpoint{}
	case "azure-native:machinelearningservices/v20220601preview:CodeContainer":
		r = &CodeContainer{}
	case "azure-native:machinelearningservices/v20220601preview:CodeVersion":
		r = &CodeVersion{}
	case "azure-native:machinelearningservices/v20220601preview:ComponentContainer":
		r = &ComponentContainer{}
	case "azure-native:machinelearningservices/v20220601preview:ComponentVersion":
		r = &ComponentVersion{}
	case "azure-native:machinelearningservices/v20220601preview:Compute":
		r = &Compute{}
	case "azure-native:machinelearningservices/v20220601preview:DataContainer":
		r = &DataContainer{}
	case "azure-native:machinelearningservices/v20220601preview:DataVersion":
		r = &DataVersion{}
	case "azure-native:machinelearningservices/v20220601preview:Datastore":
		r = &Datastore{}
	case "azure-native:machinelearningservices/v20220601preview:EnvironmentContainer":
		r = &EnvironmentContainer{}
	case "azure-native:machinelearningservices/v20220601preview:EnvironmentVersion":
		r = &EnvironmentVersion{}
	case "azure-native:machinelearningservices/v20220601preview:Job":
		r = &Job{}
	case "azure-native:machinelearningservices/v20220601preview:LabelingJob":
		r = &LabelingJob{}
	case "azure-native:machinelearningservices/v20220601preview:ModelContainer":
		r = &ModelContainer{}
	case "azure-native:machinelearningservices/v20220601preview:ModelVersion":
		r = &ModelVersion{}
	case "azure-native:machinelearningservices/v20220601preview:OnlineDeployment":
		r = &OnlineDeployment{}
	case "azure-native:machinelearningservices/v20220601preview:OnlineEndpoint":
		r = &OnlineEndpoint{}
	case "azure-native:machinelearningservices/v20220601preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:machinelearningservices/v20220601preview:Schedule":
		r = &Schedule{}
	case "azure-native:machinelearningservices/v20220601preview:Workspace":
		r = &Workspace{}
	case "azure-native:machinelearningservices/v20220601preview:WorkspaceConnection":
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
		"machinelearningservices/v20220601preview",
		&module{version},
	)
}
