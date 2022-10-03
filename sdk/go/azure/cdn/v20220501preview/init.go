


package v20220501preview

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
	case "azure-native:cdn/v20220501preview:AFDCustomDomain":
		r = &AFDCustomDomain{}
	case "azure-native:cdn/v20220501preview:AFDEndpoint":
		r = &AFDEndpoint{}
	case "azure-native:cdn/v20220501preview:AFDOrigin":
		r = &AFDOrigin{}
	case "azure-native:cdn/v20220501preview:AFDOriginGroup":
		r = &AFDOriginGroup{}
	case "azure-native:cdn/v20220501preview:CustomDomain":
		r = &CustomDomain{}
	case "azure-native:cdn/v20220501preview:Endpoint":
		r = &Endpoint{}
	case "azure-native:cdn/v20220501preview:Origin":
		r = &Origin{}
	case "azure-native:cdn/v20220501preview:OriginGroup":
		r = &OriginGroup{}
	case "azure-native:cdn/v20220501preview:Policy":
		r = &Policy{}
	case "azure-native:cdn/v20220501preview:Profile":
		r = &Profile{}
	case "azure-native:cdn/v20220501preview:Route":
		r = &Route{}
	case "azure-native:cdn/v20220501preview:Rule":
		r = &Rule{}
	case "azure-native:cdn/v20220501preview:RuleSet":
		r = &RuleSet{}
	case "azure-native:cdn/v20220501preview:Secret":
		r = &Secret{}
	case "azure-native:cdn/v20220501preview:SecurityPolicy":
		r = &SecurityPolicy{}
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
		"cdn/v20220501preview",
		&module{version},
	)
}
