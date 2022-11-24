


package v20220808

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
	case "azure-native:automation/v20220808:AutomationAccount":
		r = &AutomationAccount{}
	case "azure-native:automation/v20220808:Certificate":
		r = &Certificate{}
	case "azure-native:automation/v20220808:Connection":
		r = &Connection{}
	case "azure-native:automation/v20220808:ConnectionType":
		r = &ConnectionType{}
	case "azure-native:automation/v20220808:Credential":
		r = &Credential{}
	case "azure-native:automation/v20220808:DscConfiguration":
		r = &DscConfiguration{}
	case "azure-native:automation/v20220808:DscNodeConfiguration":
		r = &DscNodeConfiguration{}
	case "azure-native:automation/v20220808:HybridRunbookWorker":
		r = &HybridRunbookWorker{}
	case "azure-native:automation/v20220808:HybridRunbookWorkerGroup":
		r = &HybridRunbookWorkerGroup{}
	case "azure-native:automation/v20220808:JobSchedule":
		r = &JobSchedule{}
	case "azure-native:automation/v20220808:Module":
		r = &Module{}
	case "azure-native:automation/v20220808:Python2Package":
		r = &Python2Package{}
	case "azure-native:automation/v20220808:Python3Package":
		r = &Python3Package{}
	case "azure-native:automation/v20220808:Runbook":
		r = &Runbook{}
	case "azure-native:automation/v20220808:Schedule":
		r = &Schedule{}
	case "azure-native:automation/v20220808:SourceControl":
		r = &SourceControl{}
	case "azure-native:automation/v20220808:Variable":
		r = &Variable{}
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
		"automation/v20220808",
		&module{version},
	)
}
