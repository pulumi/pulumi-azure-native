


package videoanalyzer

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
	case "azure-native:videoanalyzer:AccessPolicy":
		r = &AccessPolicy{}
	case "azure-native:videoanalyzer:EdgeModule":
		r = &EdgeModule{}
	case "azure-native:videoanalyzer:LivePipeline":
		r = &LivePipeline{}
	case "azure-native:videoanalyzer:PipelineJob":
		r = &PipelineJob{}
	case "azure-native:videoanalyzer:PipelineTopology":
		r = &PipelineTopology{}
	case "azure-native:videoanalyzer:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:videoanalyzer:Video":
		r = &Video{}
	case "azure-native:videoanalyzer:VideoAnalyzer":
		r = &VideoAnalyzer{}
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
		"videoanalyzer",
		&module{version},
	)
}
