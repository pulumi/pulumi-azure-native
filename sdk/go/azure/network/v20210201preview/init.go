


package v20210201preview

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
	case "azure-native:network/v20210201preview:AdminRule":
		r = &AdminRule{}
	case "azure-native:network/v20210201preview:AdminRuleCollection":
		r = &AdminRuleCollection{}
	case "azure-native:network/v20210201preview:ConnectivityConfiguration":
		r = &ConnectivityConfiguration{}
	case "azure-native:network/v20210201preview:DefaultAdminRule":
		r = &DefaultAdminRule{}
	case "azure-native:network/v20210201preview:DefaultUserRule":
		r = &DefaultUserRule{}
	case "azure-native:network/v20210201preview:NetworkGroup":
		r = &NetworkGroup{}
	case "azure-native:network/v20210201preview:NetworkManager":
		r = &NetworkManager{}
	case "azure-native:network/v20210201preview:NetworkSecurityPerimeter":
		r = &NetworkSecurityPerimeter{}
	case "azure-native:network/v20210201preview:NspAccessRule":
		r = &NspAccessRule{}
	case "azure-native:network/v20210201preview:NspAssociation":
		r = &NspAssociation{}
	case "azure-native:network/v20210201preview:NspAssociationsProxy":
		r = &NspAssociationsProxy{}
	case "azure-native:network/v20210201preview:NspProfile":
		r = &NspProfile{}
	case "azure-native:network/v20210201preview:SecurityAdminConfiguration":
		r = &SecurityAdminConfiguration{}
	case "azure-native:network/v20210201preview:SecurityUserConfiguration":
		r = &SecurityUserConfiguration{}
	case "azure-native:network/v20210201preview:UserRule":
		r = &UserRule{}
	case "azure-native:network/v20210201preview:UserRuleCollection":
		r = &UserRuleCollection{}
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
		"network/v20210201preview",
		&module{version},
	)
}
