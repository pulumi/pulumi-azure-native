


package policyinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttestationAtSubscription(ctx *pulumi.Context, args *LookupAttestationAtSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupAttestationAtSubscriptionResult, error) {
	var rv LookupAttestationAtSubscriptionResult
	err := ctx.Invoke("azure-native:policyinsights:getAttestationAtSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttestationAtSubscriptionArgs struct {
	AttestationName string `pulumi:"attestationName"`
}


type LookupAttestationAtSubscriptionResult struct {
	Comments                    *string                       `pulumi:"comments"`
	ComplianceState             *string                       `pulumi:"complianceState"`
	Evidence                    []AttestationEvidenceResponse `pulumi:"evidence"`
	ExpiresOn                   *string                       `pulumi:"expiresOn"`
	Id                          string                        `pulumi:"id"`
	LastComplianceStateChangeAt string                        `pulumi:"lastComplianceStateChangeAt"`
	Name                        string                        `pulumi:"name"`
	Owner                       *string                       `pulumi:"owner"`
	PolicyAssignmentId          string                        `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                       `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           string                        `pulumi:"provisioningState"`
	SystemData                  SystemDataResponse            `pulumi:"systemData"`
	Type                        string                        `pulumi:"type"`
}

func LookupAttestationAtSubscriptionOutput(ctx *pulumi.Context, args LookupAttestationAtSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupAttestationAtSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttestationAtSubscriptionResult, error) {
			args := v.(LookupAttestationAtSubscriptionArgs)
			r, err := LookupAttestationAtSubscription(ctx, &args, opts...)
			var s LookupAttestationAtSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttestationAtSubscriptionResultOutput)
}

type LookupAttestationAtSubscriptionOutputArgs struct {
	AttestationName pulumi.StringInput `pulumi:"attestationName"`
}

func (LookupAttestationAtSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestationAtSubscriptionArgs)(nil)).Elem()
}


type LookupAttestationAtSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupAttestationAtSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestationAtSubscriptionResult)(nil)).Elem()
}

func (o LookupAttestationAtSubscriptionResultOutput) ToLookupAttestationAtSubscriptionResultOutput() LookupAttestationAtSubscriptionResultOutput {
	return o
}

func (o LookupAttestationAtSubscriptionResultOutput) ToLookupAttestationAtSubscriptionResultOutputWithContext(ctx context.Context) LookupAttestationAtSubscriptionResultOutput {
	return o
}

func (o LookupAttestationAtSubscriptionResultOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.ComplianceState }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) Evidence() AttestationEvidenceResponseArrayOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) []AttestationEvidenceResponse { return v.Evidence }).(AttestationEvidenceResponseArrayOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) LastComplianceStateChangeAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.LastComplianceStateChangeAt }).(pulumi.StringOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) PolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.PolicyAssignmentId }).(pulumi.StringOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAttestationAtSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttestationAtSubscriptionResultOutput{})
}
