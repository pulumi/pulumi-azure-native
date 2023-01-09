


package v20221201preview

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
	case "azure-native:databoxedge/v20221201preview:Addon":
		r = &Addon{}
	case "azure-native:databoxedge/v20221201preview:ArcAddon":
		r = &ArcAddon{}
	case "azure-native:databoxedge/v20221201preview:BandwidthSchedule":
		r = &BandwidthSchedule{}
	case "azure-native:databoxedge/v20221201preview:CloudEdgeManagementRole":
		r = &CloudEdgeManagementRole{}
	case "azure-native:databoxedge/v20221201preview:Container":
		r = &Container{}
	case "azure-native:databoxedge/v20221201preview:Device":
		r = &Device{}
	case "azure-native:databoxedge/v20221201preview:FileEventTrigger":
		r = &FileEventTrigger{}
	case "azure-native:databoxedge/v20221201preview:IoTRole":
		r = &IoTRole{}
	case "azure-native:databoxedge/v20221201preview:KubernetesRole":
		r = &KubernetesRole{}
	case "azure-native:databoxedge/v20221201preview:MECRole":
		r = &MECRole{}
	case "azure-native:databoxedge/v20221201preview:MonitoringConfig":
		r = &MonitoringConfig{}
	case "azure-native:databoxedge/v20221201preview:Order":
		r = &Order{}
	case "azure-native:databoxedge/v20221201preview:PeriodicTimerEventTrigger":
		r = &PeriodicTimerEventTrigger{}
	case "azure-native:databoxedge/v20221201preview:Role":
		r = &Role{}
	case "azure-native:databoxedge/v20221201preview:Share":
		r = &Share{}
	case "azure-native:databoxedge/v20221201preview:StorageAccount":
		r = &StorageAccount{}
	case "azure-native:databoxedge/v20221201preview:StorageAccountCredential":
		r = &StorageAccountCredential{}
	case "azure-native:databoxedge/v20221201preview:Trigger":
		r = &Trigger{}
	case "azure-native:databoxedge/v20221201preview:User":
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
		"databoxedge/v20221201preview",
		&module{version},
	)
}
