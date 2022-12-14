


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
	case "azure-native:app/v20220601preview:Certificate":
		r = &Certificate{}
	case "azure-native:app/v20220601preview:ConnectedEnvironment":
		r = &ConnectedEnvironment{}
	case "azure-native:app/v20220601preview:ConnectedEnvironmentsCertificate":
		r = &ConnectedEnvironmentsCertificate{}
	case "azure-native:app/v20220601preview:ConnectedEnvironmentsDaprComponent":
		r = &ConnectedEnvironmentsDaprComponent{}
	case "azure-native:app/v20220601preview:ConnectedEnvironmentsStorage":
		r = &ConnectedEnvironmentsStorage{}
	case "azure-native:app/v20220601preview:ContainerApp":
		r = &ContainerApp{}
	case "azure-native:app/v20220601preview:ContainerAppsAuthConfig":
		r = &ContainerAppsAuthConfig{}
	case "azure-native:app/v20220601preview:ContainerAppsSourceControl":
		r = &ContainerAppsSourceControl{}
	case "azure-native:app/v20220601preview:DaprComponent":
		r = &DaprComponent{}
	case "azure-native:app/v20220601preview:ManagedEnvironment":
		r = &ManagedEnvironment{}
	case "azure-native:app/v20220601preview:ManagedEnvironmentsStorage":
		r = &ManagedEnvironmentsStorage{}
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
		"app/v20220601preview",
		&module{version},
	)
}
