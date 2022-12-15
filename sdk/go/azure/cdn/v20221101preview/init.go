


package v20221101preview

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
	case "azure-native:cdn/v20221101preview:AFDCustomDomain":
		r = &AFDCustomDomain{}
	case "azure-native:cdn/v20221101preview:AFDEndpoint":
		r = &AFDEndpoint{}
	case "azure-native:cdn/v20221101preview:AFDOrigin":
		r = &AFDOrigin{}
	case "azure-native:cdn/v20221101preview:AFDOriginGroup":
		r = &AFDOriginGroup{}
	case "azure-native:cdn/v20221101preview:CustomDomain":
		r = &CustomDomain{}
	case "azure-native:cdn/v20221101preview:Endpoint":
		r = &Endpoint{}
	case "azure-native:cdn/v20221101preview:Origin":
		r = &Origin{}
	case "azure-native:cdn/v20221101preview:OriginGroup":
		r = &OriginGroup{}
	case "azure-native:cdn/v20221101preview:Policy":
		r = &Policy{}
	case "azure-native:cdn/v20221101preview:Profile":
		r = &Profile{}
	case "azure-native:cdn/v20221101preview:Route":
		r = &Route{}
	case "azure-native:cdn/v20221101preview:Rule":
		r = &Rule{}
	case "azure-native:cdn/v20221101preview:RuleSet":
		r = &RuleSet{}
	case "azure-native:cdn/v20221101preview:Secret":
		r = &Secret{}
	case "azure-native:cdn/v20221101preview:SecurityPolicy":
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
		"cdn/v20221101preview",
		&module{version},
	)
}
