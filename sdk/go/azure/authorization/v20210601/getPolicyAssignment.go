


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyAssignment(ctx *pulumi.Context, args *LookupPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentResult, error) {
	var rv LookupPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20210601:getPolicyAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyAssignmentArgs struct {
	PolicyAssignmentName string `pulumi:"policyAssignmentName"`
	Scope                string `pulumi:"scope"`
}


type LookupPolicyAssignmentResult struct {
	Description           *string                                 `pulumi:"description"`
	DisplayName           *string                                 `pulumi:"displayName"`
	EnforcementMode       *string                                 `pulumi:"enforcementMode"`
	Id                    string                                  `pulumi:"id"`
	Identity              *IdentityResponse                       `pulumi:"identity"`
	Location              *string                                 `pulumi:"location"`
	Metadata              interface{}                             `pulumi:"metadata"`
	Name                  string                                  `pulumi:"name"`
	NonComplianceMessages []NonComplianceMessageResponse          `pulumi:"nonComplianceMessages"`
	NotScopes             []string                                `pulumi:"notScopes"`
	Parameters            map[string]ParameterValuesValueResponse `pulumi:"parameters"`
	PolicyDefinitionId    *string                                 `pulumi:"policyDefinitionId"`
	Scope                 string                                  `pulumi:"scope"`
	SystemData            SystemDataResponse                      `pulumi:"systemData"`
	Type                  string                                  `pulumi:"type"`
}


func (val *LookupPolicyAssignmentResult) Defaults() *LookupPolicyAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnforcementMode) {
		enforcementMode_ := "Default"
		tmp.EnforcementMode = &enforcementMode_
	}
	return &tmp
}

func LookupPolicyAssignmentOutput(ctx *pulumi.Context, args LookupPolicyAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyAssignmentResult, error) {
			args := v.(LookupPolicyAssignmentArgs)
			r, err := LookupPolicyAssignment(ctx, &args, opts...)
			var s LookupPolicyAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyAssignmentResultOutput)
}

type LookupPolicyAssignmentOutputArgs struct {
	PolicyAssignmentName pulumi.StringInput `pulumi:"policyAssignmentName"`
	Scope                pulumi.StringInput `pulumi:"scope"`
}

func (LookupPolicyAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyAssignmentArgs)(nil)).Elem()
}


type LookupPolicyAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyAssignmentResult)(nil)).Elem()
}

func (o LookupPolicyAssignmentResultOutput) ToLookupPolicyAssignmentResultOutput() LookupPolicyAssignmentResultOutput {
	return o
}

func (o LookupPolicyAssignmentResultOutput) ToLookupPolicyAssignmentResultOutputWithContext(ctx context.Context) LookupPolicyAssignmentResultOutput {
	return o
}

func (o LookupPolicyAssignmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) EnforcementMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.EnforcementMode }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyAssignmentResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupPolicyAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicyAssignmentResultOutput) NonComplianceMessages() NonComplianceMessageResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) []NonComplianceMessageResponse { return v.NonComplianceMessages }).(NonComplianceMessageResponseArrayOutput)
}

func (o LookupPolicyAssignmentResultOutput) NotScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) []string { return v.NotScopes }).(pulumi.StringArrayOutput)
}

func (o LookupPolicyAssignmentResultOutput) Parameters() ParameterValuesValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) map[string]ParameterValuesValueResponse { return v.Parameters }).(ParameterValuesValueResponseMapOutput)
}

func (o LookupPolicyAssignmentResultOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupPolicyAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPolicyAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyAssignmentResultOutput{})
}
