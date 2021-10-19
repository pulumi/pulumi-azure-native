


package v20210801preview

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
	case "azure-native:containerregistry/v20210801preview:ConnectedRegistry":
		r = &ConnectedRegistry{}
	case "azure-native:containerregistry/v20210801preview:ExportPipeline":
		r = &ExportPipeline{}
	case "azure-native:containerregistry/v20210801preview:ImportPipeline":
		r = &ImportPipeline{}
	case "azure-native:containerregistry/v20210801preview:PipelineRun":
		r = &PipelineRun{}
	case "azure-native:containerregistry/v20210801preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:containerregistry/v20210801preview:Registry":
		r = &Registry{}
	case "azure-native:containerregistry/v20210801preview:Replication":
		r = &Replication{}
	case "azure-native:containerregistry/v20210801preview:ScopeMap":
		r = &ScopeMap{}
	case "azure-native:containerregistry/v20210801preview:Token":
		r = &Token{}
	case "azure-native:containerregistry/v20210801preview:Webhook":
		r = &Webhook{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"containerregistry/v20210801preview",
		&module{version},
	)
}
