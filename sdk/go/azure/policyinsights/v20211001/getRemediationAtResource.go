


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRemediationAtResource(ctx *pulumi.Context, args *LookupRemediationAtResourceArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtResourceResult, error) {
	var rv LookupRemediationAtResourceResult
	err := ctx.Invoke("azure-native:policyinsights/v20211001:getRemediationAtResource", args, &rv, opts...)
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

func LookupRemediationAtResourceOutput(ctx *pulumi.Context, args LookupRemediationAtResourceOutputArgs, opts ...pulumi.InvokeOption) LookupRemediationAtResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemediationAtResourceResult, error) {
			args := v.(LookupRemediationAtResourceArgs)
			r, err := LookupRemediationAtResource(ctx, &args, opts...)
			var s LookupRemediationAtResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemediationAtResourceResultOutput)
}

type LookupRemediationAtResourceOutputArgs struct {
	RemediationName pulumi.StringInput `pulumi:"remediationName"`
	ResourceId      pulumi.StringInput `pulumi:"resourceId"`
}

func (LookupRemediationAtResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtResourceArgs)(nil)).Elem()
}


type LookupRemediationAtResourceResultOutput struct{ *pulumi.OutputState }

func (LookupRemediationAtResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtResourceResult)(nil)).Elem()
}

func (o LookupRemediationAtResourceResultOutput) ToLookupRemediationAtResourceResultOutput() LookupRemediationAtResourceResultOutput {
	return o
}

func (o LookupRemediationAtResourceResultOutput) ToLookupRemediationAtResourceResultOutputWithContext(ctx context.Context) LookupRemediationAtResourceResultOutput {
	return o
}

func (o LookupRemediationAtResourceResultOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceResultOutput) DeploymentStatus() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) RemediationDeploymentSummaryResponse {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponseOutput)
}

func (o LookupRemediationAtResourceResultOutput) FailureThreshold() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) *RemediationPropertiesResponseFailureThreshold {
		return v.FailureThreshold
	}).(RemediationPropertiesResponseFailureThresholdPtrOutput)
}

func (o LookupRemediationAtResourceResultOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) *RemediationFiltersResponse { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

func (o LookupRemediationAtResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceResultOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) string { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceResultOutput) ParallelDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) *int { return v.ParallelDeployments }).(pulumi.IntPtrOutput)
}

func (o LookupRemediationAtResourceResultOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) *string { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtResourceResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtResourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceResultOutput) ResourceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) *int { return v.ResourceCount }).(pulumi.IntPtrOutput)
}

func (o LookupRemediationAtResourceResultOutput) ResourceDiscoveryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) *string { return v.ResourceDiscoveryMode }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtResourceResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

func (o LookupRemediationAtResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRemediationAtResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemediationAtResourceResultOutput{})
}
