


package v20220501

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
	case "azure-native:storage/v20220501:BlobContainer":
		r = &BlobContainer{}
	case "azure-native:storage/v20220501:BlobContainerImmutabilityPolicy":
		r = &BlobContainerImmutabilityPolicy{}
	case "azure-native:storage/v20220501:BlobInventoryPolicy":
		r = &BlobInventoryPolicy{}
	case "azure-native:storage/v20220501:BlobServiceProperties":
		r = &BlobServiceProperties{}
	case "azure-native:storage/v20220501:EncryptionScope":
		r = &EncryptionScope{}
	case "azure-native:storage/v20220501:FileServiceProperties":
		r = &FileServiceProperties{}
	case "azure-native:storage/v20220501:FileShare":
		r = &FileShare{}
	case "azure-native:storage/v20220501:LocalUser":
		r = &LocalUser{}
	case "azure-native:storage/v20220501:ManagementPolicy":
		r = &ManagementPolicy{}
	case "azure-native:storage/v20220501:ObjectReplicationPolicy":
		r = &ObjectReplicationPolicy{}
	case "azure-native:storage/v20220501:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:storage/v20220501:Queue":
		r = &Queue{}
	case "azure-native:storage/v20220501:QueueServiceProperties":
		r = &QueueServiceProperties{}
	case "azure-native:storage/v20220501:StorageAccount":
		r = &StorageAccount{}
	case "azure-native:storage/v20220501:Table":
		r = &Table{}
	case "azure-native:storage/v20220501:TableServiceProperties":
		r = &TableServiceProperties{}
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
		"storage/v20220501",
		&module{version},
	)
}
