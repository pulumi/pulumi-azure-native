


package v20170101

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
	case "azure-native:customerinsights/v20170101:Connector":
		r = &Connector{}
	case "azure-native:customerinsights/v20170101:ConnectorMapping":
		r = &ConnectorMapping{}
	case "azure-native:customerinsights/v20170101:Hub":
		r = &Hub{}
	case "azure-native:customerinsights/v20170101:Kpi":
		r = &Kpi{}
	case "azure-native:customerinsights/v20170101:Link":
		r = &Link{}
	case "azure-native:customerinsights/v20170101:Profile":
		r = &Profile{}
	case "azure-native:customerinsights/v20170101:Relationship":
		r = &Relationship{}
	case "azure-native:customerinsights/v20170101:RelationshipLink":
		r = &RelationshipLink{}
	case "azure-native:customerinsights/v20170101:RoleAssignment":
		r = &RoleAssignment{}
	case "azure-native:customerinsights/v20170101:View":
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
		"customerinsights/v20170101",
		&module{version},
	)
}
