


package v20210701preview

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
	case "azure-native:security/v20210701preview:CustomAssessmentAutomation":
		r = &CustomAssessmentAutomation{}
	case "azure-native:security/v20210701preview:CustomEntityStoreAssignment":
		r = &CustomEntityStoreAssignment{}
	case "azure-native:security/v20210701preview:SecurityConnector":
		r = &SecurityConnector{}
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
		"security/v20210701preview",
		&module{version},
	)
}
