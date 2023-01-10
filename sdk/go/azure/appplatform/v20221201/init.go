


package v20221201

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
	case "azure-native:appplatform/v20221201:ApiPortal":
		r = &ApiPortal{}
	case "azure-native:appplatform/v20221201:ApiPortalCustomDomain":
		r = &ApiPortalCustomDomain{}
	case "azure-native:appplatform/v20221201:App":
		r = &App{}
	case "azure-native:appplatform/v20221201:Binding":
		r = &Binding{}
	case "azure-native:appplatform/v20221201:BuildServiceAgentPool":
		r = &BuildServiceAgentPool{}
	case "azure-native:appplatform/v20221201:BuildServiceBuilder":
		r = &BuildServiceBuilder{}
	case "azure-native:appplatform/v20221201:BuildpackBinding":
		r = &BuildpackBinding{}
	case "azure-native:appplatform/v20221201:Certificate":
		r = &Certificate{}
	case "azure-native:appplatform/v20221201:ConfigServer":
		r = &ConfigServer{}
	case "azure-native:appplatform/v20221201:ConfigurationService":
		r = &ConfigurationService{}
	case "azure-native:appplatform/v20221201:CustomDomain":
		r = &CustomDomain{}
	case "azure-native:appplatform/v20221201:Deployment":
		r = &Deployment{}
	case "azure-native:appplatform/v20221201:Gateway":
		r = &Gateway{}
	case "azure-native:appplatform/v20221201:GatewayCustomDomain":
		r = &GatewayCustomDomain{}
	case "azure-native:appplatform/v20221201:GatewayRouteConfig":
		r = &GatewayRouteConfig{}
	case "azure-native:appplatform/v20221201:MonitoringSetting":
		r = &MonitoringSetting{}
	case "azure-native:appplatform/v20221201:Service":
		r = &Service{}
	case "azure-native:appplatform/v20221201:ServiceRegistry":
		r = &ServiceRegistry{}
	case "azure-native:appplatform/v20221201:Storage":
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
		"appplatform/v20221201",
		&module{version},
	)
}
