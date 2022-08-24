


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
	case "azure-native:storage/v20210601:BlobContainer":
		r = &BlobContainer{}
	case "azure-native:storage/v20210601:BlobContainerImmutabilityPolicy":
		r = &BlobContainerImmutabilityPolicy{}
	case "azure-native:storage/v20210601:BlobInventoryPolicy":
		r = &BlobInventoryPolicy{}
	case "azure-native:storage/v20210601:BlobServiceProperties":
		r = &BlobServiceProperties{}
	case "azure-native:storage/v20210601:EncryptionScope":
		r = &EncryptionScope{}
	case "azure-native:storage/v20210601:FileServiceProperties":
		r = &FileServiceProperties{}
	case "azure-native:storage/v20210601:FileShare":
		r = &FileShare{}
	case "azure-native:storage/v20210601:ManagementPolicy":
		r = &ManagementPolicy{}
	case "azure-native:storage/v20210601:ObjectReplicationPolicy":
		r = &ObjectReplicationPolicy{}
	case "azure-native:storage/v20210601:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:storage/v20210601:Queue":
		r = &Queue{}
	case "azure-native:storage/v20210601:QueueServiceProperties":
		r = &QueueServiceProperties{}
	case "azure-native:storage/v20210601:StorageAccount":
		r = &StorageAccount{}
	case "azure-native:storage/v20210601:Table":
		r = &Table{}
	case "azure-native:storage/v20210601:TableServiceProperties":
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
		"storage/v20210601",
		&module{version},
	)
}
