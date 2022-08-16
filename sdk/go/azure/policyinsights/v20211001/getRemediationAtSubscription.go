


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRemediationAtSubscription(ctx *pulumi.Context, args *LookupRemediationAtSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtSubscriptionResult, error) {
	var rv LookupRemediationAtSubscriptionResult
	err := ctx.Invoke("azure-native:policyinsights/v20211001:getRemediationAtSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtSubscriptionArgs struct {
	RemediationName string `pulumi:"remediationName"`
}


type LookupRemediationAtSubscriptionResult struct {
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

func LookupRemediationAtSubscriptionOutput(ctx *pulumi.Context, args LookupRemediationAtSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupRemediationAtSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemediationAtSubscriptionResult, error) {
			args := v.(LookupRemediationAtSubscriptionArgs)
			r, err := LookupRemediationAtSubscription(ctx, &args, opts...)
			var s LookupRemediationAtSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemediationAtSubscriptionResultOutput)
}

type LookupRemediationAtSubscriptionOutputArgs struct {
	RemediationName pulumi.StringInput `pulumi:"remediationName"`
}

func (LookupRemediationAtSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtSubscriptionArgs)(nil)).Elem()
}


type LookupRemediationAtSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupRemediationAtSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtSubscriptionResult)(nil)).Elem()
}

func (o LookupRemediationAtSubscriptionResultOutput) ToLookupRemediationAtSubscriptionResultOutput() LookupRemediationAtSubscriptionResultOutput {
	return o
}

func (o LookupRemediationAtSubscriptionResultOutput) ToLookupRemediationAtSubscriptionResultOutputWithContext(ctx context.Context) LookupRemediationAtSubscriptionResultOutput {
	return o
}

func (o LookupRemediationAtSubscriptionResultOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) DeploymentStatus() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) RemediationDeploymentSummaryResponse {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponseOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) FailureThreshold() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *RemediationPropertiesResponseFailureThreshold {
		return v.FailureThreshold
	}).(RemediationPropertiesResponseFailureThresholdPtrOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *RemediationFiltersResponse { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) ParallelDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *int { return v.ParallelDeployments }).(pulumi.IntPtrOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *string { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) ResourceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *int { return v.ResourceCount }).(pulumi.IntPtrOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) ResourceDiscoveryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *string { return v.ResourceDiscoveryMode }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRemediationAtSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemediationAtSubscriptionResultOutput{})
}
