


package appplatform

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
	case "azure-native:appplatform:ApiPortal":
		r = &ApiPortal{}
	case "azure-native:appplatform:ApiPortalCustomDomain":
		r = &ApiPortalCustomDomain{}
	case "azure-native:appplatform:App":
		r = &App{}
	case "azure-native:appplatform:Binding":
		r = &Binding{}
	case "azure-native:appplatform:BuildServiceAgentPool":
		r = &BuildServiceAgentPool{}
	case "azure-native:appplatform:BuildServiceBuilder":
		r = &BuildServiceBuilder{}
	case "azure-native:appplatform:BuildpackBinding":
		r = &BuildpackBinding{}
	case "azure-native:appplatform:Certificate":
		r = &Certificate{}
	case "azure-native:appplatform:ConfigServer":
		r = &ConfigServer{}
	case "azure-native:appplatform:ConfigurationService":
		r = &ConfigurationService{}
	case "azure-native:appplatform:CustomDomain":
		r = &CustomDomain{}
	case "azure-native:appplatform:Deployment":
		r = &Deployment{}
	case "azure-native:appplatform:Gateway":
		r = &Gateway{}
	case "azure-native:appplatform:GatewayCustomDomain":
		r = &GatewayCustomDomain{}
	case "azure-native:appplatform:GatewayRouteConfig":
		r = &GatewayRouteConfig{}
	case "azure-native:appplatform:MonitoringSetting":
		r = &MonitoringSetting{}
	case "azure-native:appplatform:Service":
		r = &Service{}
	case "azure-native:appplatform:ServiceRegistry":
		r = &ServiceRegistry{}
	case "azure-native:appplatform:Storage":
		r = &Storage{}
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
		"appplatform",
		&module{version},
	)
}
