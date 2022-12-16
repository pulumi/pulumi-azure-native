


package v20221014preview

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
	case "azure-native:desktopvirtualization/v20221014preview:Application":
		r = &Application{}
	case "azure-native:desktopvirtualization/v20221014preview:ApplicationGroup":
		r = &ApplicationGroup{}
	case "azure-native:desktopvirtualization/v20221014preview:HostPool":
		r = &HostPool{}
	case "azure-native:desktopvirtualization/v20221014preview:MSIXPackage":
		r = &MSIXPackage{}
	case "azure-native:desktopvirtualization/v20221014preview:PrivateEndpointConnectionByHostPool":
		r = &PrivateEndpointConnectionByHostPool{}
	case "azure-native:desktopvirtualization/v20221014preview:PrivateEndpointConnectionByWorkspace":
		r = &PrivateEndpointConnectionByWorkspace{}
	case "azure-native:desktopvirtualization/v20221014preview:ScalingPlan":
		r = &ScalingPlan{}
	case "azure-native:desktopvirtualization/v20221014preview:ScalingPlanPooledSchedule":
		r = &ScalingPlanPooledSchedule{}
	case "azure-native:desktopvirtualization/v20221014preview:Workspace":
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
		"desktopvirtualization/v20221014preview",
		&module{version},
	)
}
