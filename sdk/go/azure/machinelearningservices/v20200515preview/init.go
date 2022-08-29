


package v20200515preview

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
	case "azure-native:machinelearningservices/v20200515preview:ACIService":
		r = &ACIService{}
	case "azure-native:machinelearningservices/v20200515preview:AKSService":
		r = &AKSService{}
	case "azure-native:machinelearningservices/v20200515preview:EndpointVariant":
		r = &EndpointVariant{}
	case "azure-native:machinelearningservices/v20200515preview:LinkedWorkspace":
		r = &LinkedWorkspace{}
	case "azure-native:machinelearningservices/v20200515preview:MachineLearningCompute":
		r = &MachineLearningCompute{}
	case "azure-native:machinelearningservices/v20200515preview:MachineLearningService":
		r = &MachineLearningService{}
	case "azure-native:machinelearningservices/v20200515preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:machinelearningservices/v20200515preview:Workspace":
		r = &Workspace{}
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
		"machinelearningservices/v20200515preview",
		&module{version},
	)
}
