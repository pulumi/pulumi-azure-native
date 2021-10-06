


package customerinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azure-native:customerinsights:getRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAssignmentArgs struct {
	AssignmentName    string `pulumi:"assignmentName"`
	HubName           string `pulumi:"hubName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRoleAssignmentResult struct {
	AssignmentName     string                          `pulumi:"assignmentName"`
	ConflationPolicies *ResourceSetDescriptionResponse `pulumi:"conflationPolicies"`
	Connectors         *ResourceSetDescriptionResponse `pulumi:"connectors"`
	Description        map[string]string               `pulumi:"description"`
	DisplayName        map[string]string               `pulumi:"displayName"`
	Id                 string                          `pulumi:"id"`
	Interactions       *ResourceSetDescriptionResponse `pulumi:"interactions"`
	Kpis               *ResourceSetDescriptionResponse `pulumi:"kpis"`
	Links              *ResourceSetDescriptionResponse `pulumi:"links"`
	Name               string                          `pulumi:"name"`
	Principals         []AssignmentPrincipalResponse   `pulumi:"principals"`
	Profiles           *ResourceSetDescriptionResponse `pulumi:"profiles"`
	ProvisioningState  string                          `pulumi:"provisioningState"`
	RelationshipLinks  *ResourceSetDescriptionResponse `pulumi:"relationshipLinks"`
	Relationships      *ResourceSetDescriptionResponse `pulumi:"relationships"`
	Role               string                          `pulumi:"role"`
	RoleAssignments    *ResourceSetDescriptionResponse `pulumi:"roleAssignments"`
	SasPolicies        *ResourceSetDescriptionResponse `pulumi:"sasPolicies"`
	Segments           *ResourceSetDescriptionResponse `pulumi:"segments"`
	TenantId           string                          `pulumi:"tenantId"`
	Type               string                          `pulumi:"type"`
	Views              *ResourceSetDescriptionResponse `pulumi:"views"`
	WidgetTypes        *ResourceSetDescriptionResponse `pulumi:"widgetTypes"`
}
