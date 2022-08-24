


package v20200301

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
	case "azure-native:streamanalytics/v20200301:Cluster":
		r = &Cluster{}
	case "azure-native:streamanalytics/v20200301:Function":
		r = &Function{}
	case "azure-native:streamanalytics/v20200301:Input":
		r = &Input{}
	case "azure-native:streamanalytics/v20200301:Output":
		r = &Output{}
	case "azure-native:streamanalytics/v20200301:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:streamanalytics/v20200301:StreamingJob":
		r = &StreamingJob{}
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
		"streamanalytics/v20200301",
		&module{version},
	)
}
