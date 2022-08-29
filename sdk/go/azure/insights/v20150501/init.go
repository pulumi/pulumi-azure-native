


package v20150501

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
	case "azure-native:insights/v20150501:AnalyticsItem":
		r = &AnalyticsItem{}
	case "azure-native:insights/v20150501:Component":
		r = &Component{}
	case "azure-native:insights/v20150501:ComponentCurrentBillingFeature":
		r = &ComponentCurrentBillingFeature{}
	case "azure-native:insights/v20150501:ExportConfiguration":
		r = &ExportConfiguration{}
	case "azure-native:insights/v20150501:Favorite":
		r = &Favorite{}
	case "azure-native:insights/v20150501:MyWorkbook":
		r = &MyWorkbook{}
	case "azure-native:insights/v20150501:ProactiveDetectionConfiguration":
		r = &ProactiveDetectionConfiguration{}
	case "azure-native:insights/v20150501:WebTest":
		r = &WebTest{}
	case "azure-native:insights/v20150501:Workbook":
		r = &Workbook{}
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
		"insights/v20150501",
		&module{version},
	)
}
