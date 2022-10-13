


package policyinsights

import (
	"context"
	"reflect"

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

func LookupRemediationAtManagementGroupOutput(ctx *pulumi.Context, args LookupRemediationAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupRemediationAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemediationAtManagementGroupResult, error) {
			args := v.(LookupRemediationAtManagementGroupArgs)
			r, err := LookupRemediationAtManagementGroup(ctx, &args, opts...)
			var s LookupRemediationAtManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemediationAtManagementGroupResultOutput)
}

type LookupRemediationAtManagementGroupOutputArgs struct {
	ManagementGroupId         pulumi.StringInput `pulumi:"managementGroupId"`
	ManagementGroupsNamespace pulumi.StringInput `pulumi:"managementGroupsNamespace"`
	RemediationName           pulumi.StringInput `pulumi:"remediationName"`
}

func (LookupRemediationAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtManagementGroupArgs)(nil)).Elem()
}


type LookupRemediationAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupRemediationAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtManagementGroupResult)(nil)).Elem()
}

func (o LookupRemediationAtManagementGroupResultOutput) ToLookupRemediationAtManagementGroupResultOutput() LookupRemediationAtManagementGroupResultOutput {
	return o
}

func (o LookupRemediationAtManagementGroupResultOutput) ToLookupRemediationAtManagementGroupResultOutputWithContext(ctx context.Context) LookupRemediationAtManagementGroupResultOutput {
	return o
}

func (o LookupRemediationAtManagementGroupResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupRemediationAtManagementGroupResultOutput) DeploymentStatus() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) RemediationDeploymentSummaryResponse {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponseOutput)
}

func (o LookupRemediationAtManagementGroupResultOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *RemediationFiltersResponse { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

func (o LookupRemediationAtManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemediationAtManagementGroupResultOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

func (o LookupRemediationAtManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRemediationAtManagementGroupResultOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *string { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtManagementGroupResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtManagementGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRemediationAtManagementGroupResultOutput) ResourceDiscoveryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *string { return v.ResourceDiscoveryMode }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemediationAtManagementGroupResultOutput{})
}
