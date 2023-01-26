


package customerinsights

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
	case "azure-native:customerinsights:Connector":
		r = &Connector{}
	case "azure-native:customerinsights:ConnectorMapping":
		r = &ConnectorMapping{}
	case "azure-native:customerinsights:Hub":
		r = &Hub{}
	case "azure-native:customerinsights:Kpi":
		r = &Kpi{}
	case "azure-native:customerinsights:Link":
		r = &Link{}
	case "azure-native:customerinsights:Prediction":
		r = &Prediction{}
	case "azure-native:customerinsights:Profile":
		r = &Profile{}
	case "azure-native:customerinsights:Relationship":
		r = &Relationship{}
	case "azure-native:customerinsights:RelationshipLink":
		r = &RelationshipLink{}
	case "azure-native:customerinsights:RoleAssignment":
		r = &RoleAssignment{}
	case "azure-native:customerinsights:View":
		r = &View{}
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
		"customerinsights",
		&module{version},
	)
}
