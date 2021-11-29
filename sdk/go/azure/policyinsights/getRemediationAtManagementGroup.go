


package policyinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRemediationAtManagementGroup(ctx *pulumi.Context, args *LookupRemediationAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtManagementGroupResult, error) {
	var rv LookupRemediationAtManagementGroupResult
	err := ctx.Invoke("azure-native:policyinsights:getRemediationAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtManagementGroupArgs struct {
	ManagementGroupId         string `pulumi:"managementGroupId"`
	ManagementGroupsNamespace string `pulumi:"managementGroupsNamespace"`
	RemediationName           string `pulumi:"remediationName"`
}


type LookupRemediationAtManagementGroupResult struct {
	CreatedOn                   string                               `pulumi:"createdOn"`
	DeploymentStatus            RemediationDeploymentSummaryResponse `pulumi:"deploymentStatus"`
	Filters                     *RemediationFiltersResponse          `pulumi:"filters"`
	Id                          string                               `pulumi:"id"`
	LastUpdatedOn               string                               `pulumi:"lastUpdatedOn"`
	Name                        string                               `pulumi:"name"`
	PolicyAssignmentId          *string                              `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                              `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           string                               `pulumi:"provisioningState"`
	ResourceDiscoveryMode       *string                              `pulumi:"resourceDiscoveryMode"`
	Type                        string                               `pulumi:"type"`
}
