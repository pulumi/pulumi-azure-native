


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
	case "azure-native:azurestackhci/v20210701preview:galleryimageRetrieve":
		r = &GalleryimageRetrieve{}
	case "azure-native:azurestackhci/v20210701preview:networkinterfaceRetrieve":
		r = &NetworkinterfaceRetrieve{}
	case "azure-native:azurestackhci/v20210701preview:virtualharddiskRetrieve":
		r = &VirtualharddiskRetrieve{}
	case "azure-native:azurestackhci/v20210701preview:virtualmachineRetrieve":
		r = &VirtualmachineRetrieve{}
	case "azure-native:azurestackhci/v20210701preview:virtualnetworkRetrieve":
		r = &VirtualnetworkRetrieve{}
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
		"azurestackhci/v20210701preview",
		&module{version},
	)
}
