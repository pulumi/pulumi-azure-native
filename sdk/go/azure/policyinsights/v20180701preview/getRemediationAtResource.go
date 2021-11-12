


package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRemediationAtResource(ctx *pulumi.Context, args *LookupRemediationAtResourceArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtResourceResult, error) {
	var rv LookupRemediationAtResourceResult
	err := ctx.Invoke("azure-native:policyinsights/v20180701preview:getRemediationAtResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtResourceArgs struct {
	RemediationName string `pulumi:"remediationName"`
	ResourceId      string `pulumi:"resourceId"`
}


type LookupRemediationAtResourceResult struct {
	CreatedOn                   string                                `pulumi:"createdOn"`
	DeploymentStatus            *RemediationDeploymentSummaryResponse `pulumi:"deploymentStatus"`
	Filters                     *RemediationFiltersResponse           `pulumi:"filters"`
	Id                          string                                `pulumi:"id"`
	LastUpdatedOn               string                                `pulumi:"lastUpdatedOn"`
	Name                        string                                `pulumi:"name"`
	PolicyAssignmentId          *string                               `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                               `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           string                                `pulumi:"provisioningState"`
	Type                        string                                `pulumi:"type"`
}
