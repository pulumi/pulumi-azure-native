


package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRemediationAtResourceGroup(ctx *pulumi.Context, args *LookupRemediationAtResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtResourceGroupResult, error) {
	var rv LookupRemediationAtResourceGroupResult
	err := ctx.Invoke("azure-native:policyinsights/v20190701:getRemediationAtResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtResourceGroupArgs struct {
	RemediationName   string `pulumi:"remediationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRemediationAtResourceGroupResult struct {
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
