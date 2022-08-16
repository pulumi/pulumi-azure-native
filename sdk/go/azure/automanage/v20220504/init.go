


package v20220504

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
	case "azure-native:automanage/v20220504:ConfigurationProfile":
		r = &ConfigurationProfile{}
	case "azure-native:automanage/v20220504:ConfigurationProfileAssignment":
		r = &ConfigurationProfileAssignment{}
	case "azure-native:automanage/v20220504:ConfigurationProfileHCIAssignment":
		r = &ConfigurationProfileHCIAssignment{}
	case "azure-native:automanage/v20220504:ConfigurationProfileHCRPAssignment":
		r = &ConfigurationProfileHCRPAssignment{}
	case "azure-native:automanage/v20220504:ConfigurationProfilesVersion":
		r = &ConfigurationProfilesVersion{}
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
		"automanage/v20220504",
		&module{version},
	)
}
