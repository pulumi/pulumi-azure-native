


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
	case "azure-native:netapp/v20220501:Account":
		r = &Account{}
	case "azure-native:netapp/v20220501:Backup":
		r = &Backup{}
	case "azure-native:netapp/v20220501:BackupPolicy":
		r = &BackupPolicy{}
	case "azure-native:netapp/v20220501:Pool":
		r = &Pool{}
	case "azure-native:netapp/v20220501:Snapshot":
		r = &Snapshot{}
	case "azure-native:netapp/v20220501:SnapshotPolicy":
		r = &SnapshotPolicy{}
	case "azure-native:netapp/v20220501:Subvolume":
		r = &Subvolume{}
	case "azure-native:netapp/v20220501:Volume":
		r = &Volume{}
	case "azure-native:netapp/v20220501:VolumeGroup":
		r = &VolumeGroup{}
	case "azure-native:netapp/v20220501:VolumeQuotaRule":
		r = &VolumeQuotaRule{}
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
		"netapp/v20220501",
		&module{version},
	)
}
