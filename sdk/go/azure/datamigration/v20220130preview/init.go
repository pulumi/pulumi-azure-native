


package v20220130preview

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
	case "azure-native:datamigration/v20220130preview:File":
		r = &File{}
	case "azure-native:datamigration/v20220130preview:Project":
		r = &Project{}
	case "azure-native:datamigration/v20220130preview:Service":
		r = &Service{}
	case "azure-native:datamigration/v20220130preview:ServiceTask":
		r = &ServiceTask{}
	case "azure-native:datamigration/v20220130preview:SqlMigrationService":
		r = &SqlMigrationService{}
	case "azure-native:datamigration/v20220130preview:Task":
		r = &Task{}
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
		"datamigration/v20220130preview",
		&module{version},
	)
}
