


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRemediationAtResourceGroup(ctx *pulumi.Context, args *LookupRemediationAtResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtResourceGroupResult, error) {
	var rv LookupRemediationAtResourceGroupResult
	err := ctx.Invoke("azure-native:policyinsights/v20211001:getRemediationAtResourceGroup", args, &rv, opts...)
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
	CorrelationId               string                                         `pulumi:"correlationId"`
	CreatedOn                   string                                         `pulumi:"createdOn"`
	DeploymentStatus            RemediationDeploymentSummaryResponse           `pulumi:"deploymentStatus"`
	FailureThreshold            *RemediationPropertiesResponseFailureThreshold `pulumi:"failureThreshold"`
	Filters                     *RemediationFiltersResponse                    `pulumi:"filters"`
	Id                          string                                         `pulumi:"id"`
	LastUpdatedOn               string                                         `pulumi:"lastUpdatedOn"`
	Name                        string                                         `pulumi:"name"`
	ParallelDeployments         *int                                           `pulumi:"parallelDeployments"`
	PolicyAssignmentId          *string                                        `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                                        `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           string                                         `pulumi:"provisioningState"`
	ResourceCount               *int                                           `pulumi:"resourceCount"`
	ResourceDiscoveryMode       *string                                        `pulumi:"resourceDiscoveryMode"`
	StatusMessage               string                                         `pulumi:"statusMessage"`
	SystemData                  SystemDataResponse                             `pulumi:"systemData"`
	Type                        string                                         `pulumi:"type"`
}

func LookupRemediationAtResourceGroupOutput(ctx *pulumi.Context, args LookupRemediationAtResourceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupRemediationAtResourceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemediationAtResourceGroupResult, error) {
			args := v.(LookupRemediationAtResourceGroupArgs)
			r, err := LookupRemediationAtResourceGroup(ctx, &args, opts...)
			var s LookupRemediationAtResourceGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemediationAtResourceGroupResultOutput)
}

type LookupRemediationAtResourceGroupOutputArgs struct {
	RemediationName   pulumi.StringInput `pulumi:"remediationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRemediationAtResourceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtResourceGroupArgs)(nil)).Elem()
}


type LookupRemediationAtResourceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupRemediationAtResourceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtResourceGroupResult)(nil)).Elem()
}

func (o LookupRemediationAtResourceGroupResultOutput) ToLookupRemediationAtResourceGroupResultOutput() LookupRemediationAtResourceGroupResultOutput {
	return o
}

func (o LookupRemediationAtResourceGroupResultOutput) ToLookupRemediationAtResourceGroupResultOutputWithContext(ctx context.Context) LookupRemediationAtResourceGroupResultOutput {
	return o
}

func (o LookupRemediationAtResourceGroupResultOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) DeploymentStatus() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) RemediationDeploymentSummaryResponse {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponseOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) FailureThreshold() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) *RemediationPropertiesResponseFailureThreshold {
		return v.FailureThreshold
	}).(RemediationPropertiesResponseFailureThresholdPtrOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) *RemediationFiltersResponse { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) string { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) ParallelDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) *int { return v.ParallelDeployments }).(pulumi.IntPtrOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) *string { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) ResourceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) *int { return v.ResourceCount }).(pulumi.IntPtrOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) ResourceDiscoveryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) *string { return v.ResourceDiscoveryMode }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRemediationAtResourceGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemediationAtResourceGroupResultOutput{})
}
