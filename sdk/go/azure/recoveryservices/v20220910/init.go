


package v20220910

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
	case "azure-native:recoveryservices/v20220910:ReplicationFabric":
		r = &ReplicationFabric{}
	case "azure-native:recoveryservices/v20220910:ReplicationMigrationItem":
		r = &ReplicationMigrationItem{}
	case "azure-native:recoveryservices/v20220910:ReplicationNetworkMapping":
		r = &ReplicationNetworkMapping{}
	case "azure-native:recoveryservices/v20220910:ReplicationPolicy":
		r = &ReplicationPolicy{}
	case "azure-native:recoveryservices/v20220910:ReplicationProtectedItem":
		r = &ReplicationProtectedItem{}
	case "azure-native:recoveryservices/v20220910:ReplicationProtectionContainerMapping":
		r = &ReplicationProtectionContainerMapping{}
	case "azure-native:recoveryservices/v20220910:ReplicationRecoveryPlan":
		r = &ReplicationRecoveryPlan{}
	case "azure-native:recoveryservices/v20220910:ReplicationRecoveryServicesProvider":
		r = &ReplicationRecoveryServicesProvider{}
	case "azure-native:recoveryservices/v20220910:ReplicationStorageClassificationMapping":
		r = &ReplicationStorageClassificationMapping{}
	case "azure-native:recoveryservices/v20220910:ReplicationvCenter":
		r = &ReplicationvCenter{}
	case "azure-native:recoveryservices/v20220910:Vault":
		r = &Vault{}
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
		"recoveryservices/v20220910",
		&module{version},
	)
}
