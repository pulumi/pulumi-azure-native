


package v20220101preview

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
	case "azure-native:appplatform/v20220101preview:ApiPortal":
		r = &ApiPortal{}
	case "azure-native:appplatform/v20220101preview:ApiPortalCustomDomain":
		r = &ApiPortalCustomDomain{}
	case "azure-native:appplatform/v20220101preview:App":
		r = &App{}
	case "azure-native:appplatform/v20220101preview:Binding":
		r = &Binding{}
	case "azure-native:appplatform/v20220101preview:BuildServiceBuilder":
		r = &BuildServiceBuilder{}
	case "azure-native:appplatform/v20220101preview:BuildpackBinding":
		r = &BuildpackBinding{}
	case "azure-native:appplatform/v20220101preview:Certificate":
		r = &Certificate{}
	case "azure-native:appplatform/v20220101preview:ConfigurationService":
		r = &ConfigurationService{}
	case "azure-native:appplatform/v20220101preview:CustomDomain":
		r = &CustomDomain{}
	case "azure-native:appplatform/v20220101preview:Deployment":
		r = &Deployment{}
	case "azure-native:appplatform/v20220101preview:Gateway":
		r = &Gateway{}
	case "azure-native:appplatform/v20220101preview:GatewayCustomDomain":
		r = &GatewayCustomDomain{}
	case "azure-native:appplatform/v20220101preview:GatewayRouteConfig":
		r = &GatewayRouteConfig{}
	case "azure-native:appplatform/v20220101preview:Service":
		r = &Service{}
	case "azure-native:appplatform/v20220101preview:ServiceRegistry":
		r = &ServiceRegistry{}
	case "azure-native:appplatform/v20220101preview:Storage":
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
		"appplatform/v20220101preview",
		&module{version},
	)
}
