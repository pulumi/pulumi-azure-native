


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttestationAtResource(ctx *pulumi.Context, args *LookupAttestationAtResourceArgs, opts ...pulumi.InvokeOption) (*LookupAttestationAtResourceResult, error) {
	var rv LookupAttestationAtResourceResult
	err := ctx.Invoke("azure-native:policyinsights/v20210101:getAttestationAtResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttestationAtResourceArgs struct {
	AttestationName string `pulumi:"attestationName"`
	ResourceId      string `pulumi:"resourceId"`
}


type LookupAttestationAtResourceResult struct {
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

func LookupAttestationAtResourceOutput(ctx *pulumi.Context, args LookupAttestationAtResourceOutputArgs, opts ...pulumi.InvokeOption) LookupAttestationAtResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttestationAtResourceResult, error) {
			args := v.(LookupAttestationAtResourceArgs)
			r, err := LookupAttestationAtResource(ctx, &args, opts...)
			var s LookupAttestationAtResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttestationAtResourceResultOutput)
}

type LookupAttestationAtResourceOutputArgs struct {
	AttestationName pulumi.StringInput `pulumi:"attestationName"`
	ResourceId      pulumi.StringInput `pulumi:"resourceId"`
}

func (LookupAttestationAtResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestationAtResourceArgs)(nil)).Elem()
}


type LookupAttestationAtResourceResultOutput struct{ *pulumi.OutputState }

func (LookupAttestationAtResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestationAtResourceResult)(nil)).Elem()
}

func (o LookupAttestationAtResourceResultOutput) ToLookupAttestationAtResourceResultOutput() LookupAttestationAtResourceResultOutput {
	return o
}

func (o LookupAttestationAtResourceResultOutput) ToLookupAttestationAtResourceResultOutputWithContext(ctx context.Context) LookupAttestationAtResourceResultOutput {
	return o
}

func (o LookupAttestationAtResourceResultOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceResultOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) *string { return v.ComplianceState }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceResultOutput) Evidence() AttestationEvidenceResponseArrayOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) []AttestationEvidenceResponse { return v.Evidence }).(AttestationEvidenceResponseArrayOutput)
}

func (o LookupAttestationAtResourceResultOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) *string { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAttestationAtResourceResultOutput) LastComplianceStateChangeAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) string { return v.LastComplianceStateChangeAt }).(pulumi.StringOutput)
}

func (o LookupAttestationAtResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAttestationAtResourceResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceResultOutput) PolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) string { return v.PolicyAssignmentId }).(pulumi.StringOutput)
}

func (o LookupAttestationAtResourceResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAttestationAtResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAttestationAtResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttestationAtResourceResultOutput{})
}
