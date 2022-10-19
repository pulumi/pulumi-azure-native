


package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttestationAtResourceGroup(ctx *pulumi.Context, args *LookupAttestationAtResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupAttestationAtResourceGroupResult, error) {
	var rv LookupAttestationAtResourceGroupResult
	err := ctx.Invoke("azure-native:policyinsights/v20220901:getAttestationAtResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttestationAtResourceGroupArgs struct {
	AttestationName   string `pulumi:"attestationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAttestationAtResourceGroupResult struct {
	AssessmentDate              *string                       `pulumi:"assessmentDate"`
	Comments                    *string                       `pulumi:"comments"`
	ComplianceState             *string                       `pulumi:"complianceState"`
	Evidence                    []AttestationEvidenceResponse `pulumi:"evidence"`
	ExpiresOn                   *string                       `pulumi:"expiresOn"`
	Id                          string                        `pulumi:"id"`
	LastComplianceStateChangeAt string                        `pulumi:"lastComplianceStateChangeAt"`
	Metadata                    interface{}                   `pulumi:"metadata"`
	Name                        string                        `pulumi:"name"`
	Owner                       *string                       `pulumi:"owner"`
	PolicyAssignmentId          string                        `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                       `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           string                        `pulumi:"provisioningState"`
	SystemData                  SystemDataResponse            `pulumi:"systemData"`
	Type                        string                        `pulumi:"type"`
}

func LookupAttestationAtResourceGroupOutput(ctx *pulumi.Context, args LookupAttestationAtResourceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupAttestationAtResourceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttestationAtResourceGroupResult, error) {
			args := v.(LookupAttestationAtResourceGroupArgs)
			r, err := LookupAttestationAtResourceGroup(ctx, &args, opts...)
			var s LookupAttestationAtResourceGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttestationAtResourceGroupResultOutput)
}

type LookupAttestationAtResourceGroupOutputArgs struct {
	AttestationName   pulumi.StringInput `pulumi:"attestationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAttestationAtResourceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestationAtResourceGroupArgs)(nil)).Elem()
}


type LookupAttestationAtResourceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupAttestationAtResourceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestationAtResourceGroupResult)(nil)).Elem()
}

func (o LookupAttestationAtResourceGroupResultOutput) ToLookupAttestationAtResourceGroupResultOutput() LookupAttestationAtResourceGroupResultOutput {
	return o
}

func (o LookupAttestationAtResourceGroupResultOutput) ToLookupAttestationAtResourceGroupResultOutputWithContext(ctx context.Context) LookupAttestationAtResourceGroupResultOutput {
	return o
}

func (o LookupAttestationAtResourceGroupResultOutput) AssessmentDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) *string { return v.AssessmentDate }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) *string { return v.ComplianceState }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) Evidence() AttestationEvidenceResponseArrayOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) []AttestationEvidenceResponse { return v.Evidence }).(AttestationEvidenceResponseArrayOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) *string { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) LastComplianceStateChangeAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) string { return v.LastComplianceStateChangeAt }).(pulumi.StringOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) PolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) string { return v.PolicyAssignmentId }).(pulumi.StringOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAttestationAtResourceGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtResourceGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttestationAtResourceGroupResultOutput{})
}
