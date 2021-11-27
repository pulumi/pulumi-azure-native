


package v20210601

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
	case "azure-native:recoveryservices/v20210601:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:recoveryservices/v20210601:ProtectedItem":
		r = &ProtectedItem{}
	case "azure-native:recoveryservices/v20210601:ProtectionContainer":
		r = &ProtectionContainer{}
	case "azure-native:recoveryservices/v20210601:ProtectionIntent":
		r = &ProtectionIntent{}
	case "azure-native:recoveryservices/v20210601:ProtectionPolicy":
		r = &ProtectionPolicy{}
	case "azure-native:recoveryservices/v20210601:ReplicationFabric":
		r = &ReplicationFabric{}
	case "azure-native:recoveryservices/v20210601:ReplicationMigrationItem":
		r = &ReplicationMigrationItem{}
	case "azure-native:recoveryservices/v20210601:ReplicationNetworkMapping":
		r = &ReplicationNetworkMapping{}
	case "azure-native:recoveryservices/v20210601:ReplicationPolicy":
		r = &ReplicationPolicy{}
	case "azure-native:recoveryservices/v20210601:ReplicationProtectedItem":
		r = &ReplicationProtectedItem{}
	case "azure-native:recoveryservices/v20210601:ReplicationProtectionContainerMapping":
		r = &ReplicationProtectionContainerMapping{}
	case "azure-native:recoveryservices/v20210601:ReplicationRecoveryPlan":
		r = &ReplicationRecoveryPlan{}
	case "azure-native:recoveryservices/v20210601:ReplicationRecoveryServicesProvider":
		r = &ReplicationRecoveryServicesProvider{}
	case "azure-native:recoveryservices/v20210601:ReplicationStorageClassificationMapping":
		r = &ReplicationStorageClassificationMapping{}
	case "azure-native:recoveryservices/v20210601:ReplicationvCenter":
		r = &ReplicationvCenter{}
	case "azure-native:recoveryservices/v20210601:Vault":
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
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"recoveryservices/v20210601",
		&module{version},
	)
}
