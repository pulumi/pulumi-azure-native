


package v20210801

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
	case "azure-native:netapp/v20210801:Account":
		r = &Account{}
	case "azure-native:netapp/v20210801:Backup":
		r = &Backup{}
	case "azure-native:netapp/v20210801:BackupPolicy":
		r = &BackupPolicy{}
	case "azure-native:netapp/v20210801:Pool":
		r = &Pool{}
	case "azure-native:netapp/v20210801:Snapshot":
		r = &Snapshot{}
	case "azure-native:netapp/v20210801:SnapshotPolicy":
		r = &SnapshotPolicy{}
	case "azure-native:netapp/v20210801:Volume":
		r = &Volume{}
	case "azure-native:netapp/v20210801:VolumeGroup":
		r = &VolumeGroup{}
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
		"netapp/v20210801",
		&module{version},
	)
}
