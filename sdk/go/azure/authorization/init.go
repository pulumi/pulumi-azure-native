


package authorization

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
	case "azure-native:authorization:AccessReviewHistoryDefinitionById":
		r = &AccessReviewHistoryDefinitionById{}
	case "azure-native:authorization:AccessReviewScheduleDefinitionById":
		r = &AccessReviewScheduleDefinitionById{}
	case "azure-native:authorization:ManagementLockAtResourceGroupLevel":
		r = &ManagementLockAtResourceGroupLevel{}
	case "azure-native:authorization:ManagementLockAtResourceLevel":
		r = &ManagementLockAtResourceLevel{}
	case "azure-native:authorization:ManagementLockAtSubscriptionLevel":
		r = &ManagementLockAtSubscriptionLevel{}
	case "azure-native:authorization:ManagementLockByScope":
		r = &ManagementLockByScope{}
	case "azure-native:authorization:PolicyAssignment":
		r = &PolicyAssignment{}
	case "azure-native:authorization:PolicyDefinition":
		r = &PolicyDefinition{}
	case "azure-native:authorization:PolicyDefinitionAtManagementGroup":
		r = &PolicyDefinitionAtManagementGroup{}
	case "azure-native:authorization:PolicyExemption":
		r = &PolicyExemption{}
	case "azure-native:authorization:PolicySetDefinition":
		r = &PolicySetDefinition{}
	case "azure-native:authorization:PolicySetDefinitionAtManagementGroup":
		r = &PolicySetDefinitionAtManagementGroup{}
	case "azure-native:authorization:PrivateLinkAssociation":
		r = &PrivateLinkAssociation{}
	case "azure-native:authorization:ResourceManagementPrivateLink":
		r = &ResourceManagementPrivateLink{}
	case "azure-native:authorization:RoleAssignment":
		r = &RoleAssignment{}
	case "azure-native:authorization:RoleDefinition":
		r = &RoleDefinition{}
	case "azure-native:authorization:RoleManagementPolicyAssignment":
		r = &RoleManagementPolicyAssignment{}
	case "azure-native:authorization:ScopeAccessReviewHistoryDefinitionById":
		r = &ScopeAccessReviewHistoryDefinitionById{}
	case "azure-native:authorization:ScopeAccessReviewScheduleDefinitionById":
		r = &ScopeAccessReviewScheduleDefinitionById{}
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
		"authorization",
		&module{version},
	)
}
