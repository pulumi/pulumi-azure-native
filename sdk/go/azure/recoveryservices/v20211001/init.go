


package v20211001

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
	case "azure-native:recoveryservices/v20211001:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:recoveryservices/v20211001:ProtectedItem":
		r = &ProtectedItem{}
	case "azure-native:recoveryservices/v20211001:ProtectionContainer":
		r = &ProtectionContainer{}
	case "azure-native:recoveryservices/v20211001:ProtectionIntent":
		r = &ProtectionIntent{}
	case "azure-native:recoveryservices/v20211001:ProtectionPolicy":
		r = &ProtectionPolicy{}
	case "azure-native:recoveryservices/v20211001:ResourceGuardProxy":
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
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"recoveryservices/v20211001",
		&module{version},
	)
}
