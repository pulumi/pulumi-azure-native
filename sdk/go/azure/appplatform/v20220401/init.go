


package v20220401

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
	case "azure-native:appplatform/v20220401:App":
		r = &App{}
	case "azure-native:appplatform/v20220401:Binding":
		r = &Binding{}
	case "azure-native:appplatform/v20220401:BuildServiceAgentPool":
		r = &BuildServiceAgentPool{}
	case "azure-native:appplatform/v20220401:BuildServiceBuilder":
		r = &BuildServiceBuilder{}
	case "azure-native:appplatform/v20220401:BuildpackBinding":
		r = &BuildpackBinding{}
	case "azure-native:appplatform/v20220401:Certificate":
		r = &Certificate{}
	case "azure-native:appplatform/v20220401:ConfigServer":
		r = &ConfigServer{}
	case "azure-native:appplatform/v20220401:ConfigurationService":
		r = &ConfigurationService{}
	case "azure-native:appplatform/v20220401:CustomDomain":
		r = &CustomDomain{}
	case "azure-native:appplatform/v20220401:Deployment":
		r = &Deployment{}
	case "azure-native:appplatform/v20220401:MonitoringSetting":
		r = &MonitoringSetting{}
	case "azure-native:appplatform/v20220401:Service":
		r = &Service{}
	case "azure-native:appplatform/v20220401:ServiceRegistry":
		r = &ServiceRegistry{}
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
		"appplatform/v20220401",
		&module{version},
	)
}
