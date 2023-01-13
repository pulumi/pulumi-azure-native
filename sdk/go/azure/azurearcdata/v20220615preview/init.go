


package v20220615preview

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
	case "azure-native:azurearcdata/v20220615preview:ActiveDirectoryConnector":
		r = &ActiveDirectoryConnector{}
	case "azure-native:azurearcdata/v20220615preview:DataController":
		r = &DataController{}
	case "azure-native:azurearcdata/v20220615preview:PostgresInstance":
		r = &PostgresInstance{}
	case "azure-native:azurearcdata/v20220615preview:SqlManagedInstance":
		r = &SqlManagedInstance{}
	case "azure-native:azurearcdata/v20220615preview:SqlServerDatabase":
		r = &SqlServerDatabase{}
	case "azure-native:azurearcdata/v20220615preview:SqlServerInstance":
		r = &SqlServerInstance{}
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
		"azurearcdata/v20220615preview",
		&module{version},
	)
}
