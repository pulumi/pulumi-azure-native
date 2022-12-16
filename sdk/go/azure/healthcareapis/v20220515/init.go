


package v20220515

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
	case "azure-native:healthcareapis/v20220515:DicomService":
		r = &DicomService{}
	case "azure-native:healthcareapis/v20220515:FhirService":
		r = &FhirService{}
	case "azure-native:healthcareapis/v20220515:IotConnector":
		r = &IotConnector{}
	case "azure-native:healthcareapis/v20220515:IotConnectorFhirDestination":
		r = &IotConnectorFhirDestination{}
	case "azure-native:healthcareapis/v20220515:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:healthcareapis/v20220515:Service":
		r = &Service{}
	case "azure-native:healthcareapis/v20220515:Workspace":
		r = &Workspace{}
	case "azure-native:healthcareapis/v20220515:WorkspacePrivateEndpointConnection":
		r = &WorkspacePrivateEndpointConnection{}
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
		"healthcareapis/v20220515",
		&module{version},
	)
}
