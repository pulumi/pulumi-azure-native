


package v20201101preview

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
	case "azure-native:appplatform/v20201101preview:App":
		r = &App{}
	case "azure-native:appplatform/v20201101preview:Binding":
		r = &Binding{}
	case "azure-native:appplatform/v20201101preview:Certificate":
		r = &Certificate{}
	case "azure-native:appplatform/v20201101preview:ConfigServer":
		r = &ConfigServer{}
	case "azure-native:appplatform/v20201101preview:CustomDomain":
		r = &CustomDomain{}
	case "azure-native:appplatform/v20201101preview:Deployment":
		r = &Deployment{}
	case "azure-native:appplatform/v20201101preview:MonitoringSetting":
		r = &MonitoringSetting{}
	case "azure-native:appplatform/v20201101preview:Service":
		r = &Service{}
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
		"appplatform/v20201101preview",
		&module{version},
	)
}
