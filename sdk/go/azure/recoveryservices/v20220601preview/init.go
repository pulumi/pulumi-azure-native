


package v20220601preview

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
	case "azure-native:recoveryservices/v20220601preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:recoveryservices/v20220601preview:ProtectedItem":
		r = &ProtectedItem{}
	case "azure-native:recoveryservices/v20220601preview:ProtectionContainer":
		r = &ProtectionContainer{}
	case "azure-native:recoveryservices/v20220601preview:ProtectionIntent":
		r = &ProtectionIntent{}
	case "azure-native:recoveryservices/v20220601preview:ProtectionPolicy":
		r = &ProtectionPolicy{}
	case "azure-native:recoveryservices/v20220601preview:ResourceGuardProxy":
		r = &ResourceGuardProxy{}
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
		"recoveryservices/v20220601preview",
		&module{version},
	)
}
