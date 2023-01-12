


package v20230101preview

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
	case "azure-native:databoxedge/v20230101preview:Addon":
		r = &Addon{}
	case "azure-native:databoxedge/v20230101preview:ArcAddon":
		r = &ArcAddon{}
	case "azure-native:databoxedge/v20230101preview:BandwidthSchedule":
		r = &BandwidthSchedule{}
	case "azure-native:databoxedge/v20230101preview:CloudEdgeManagementRole":
		r = &CloudEdgeManagementRole{}
	case "azure-native:databoxedge/v20230101preview:Container":
		r = &Container{}
	case "azure-native:databoxedge/v20230101preview:Device":
		r = &Device{}
	case "azure-native:databoxedge/v20230101preview:FileEventTrigger":
		r = &FileEventTrigger{}
	case "azure-native:databoxedge/v20230101preview:IoTRole":
		r = &IoTRole{}
	case "azure-native:databoxedge/v20230101preview:KubernetesRole":
		r = &KubernetesRole{}
	case "azure-native:databoxedge/v20230101preview:MECRole":
		r = &MECRole{}
	case "azure-native:databoxedge/v20230101preview:MonitoringConfig":
		r = &MonitoringConfig{}
	case "azure-native:databoxedge/v20230101preview:Order":
		r = &Order{}
	case "azure-native:databoxedge/v20230101preview:PeriodicTimerEventTrigger":
		r = &PeriodicTimerEventTrigger{}
	case "azure-native:databoxedge/v20230101preview:Role":
		r = &Role{}
	case "azure-native:databoxedge/v20230101preview:Share":
		r = &Share{}
	case "azure-native:databoxedge/v20230101preview:StorageAccount":
		r = &StorageAccount{}
	case "azure-native:databoxedge/v20230101preview:StorageAccountCredential":
		r = &StorageAccountCredential{}
	case "azure-native:databoxedge/v20230101preview:Trigger":
		r = &Trigger{}
	case "azure-native:databoxedge/v20230101preview:User":
		r = &User{}
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
		"databoxedge/v20230101preview",
		&module{version},
	)
}
