


package v20180701preview

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
	case "azure-native:policyinsights/v20180701preview:RemediationAtManagementGroup":
		r = &RemediationAtManagementGroup{}
	case "azure-native:policyinsights/v20180701preview:RemediationAtResource":
		r = &RemediationAtResource{}
	case "azure-native:policyinsights/v20180701preview:RemediationAtResourceGroup":
		r = &RemediationAtResourceGroup{}
	case "azure-native:policyinsights/v20180701preview:RemediationAtSubscription":
		r = &RemediationAtSubscription{}
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
		"policyinsights/v20180701preview",
		&module{version},
	)
}
