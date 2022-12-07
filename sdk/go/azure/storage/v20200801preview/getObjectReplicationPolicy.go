


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupObjectReplicationPolicy(ctx *pulumi.Context, args *LookupObjectReplicationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupObjectReplicationPolicyResult, error) {
	var rv LookupObjectReplicationPolicyResult
	err := ctx.Invoke("azure-native:storage/v20200801preview:getObjectReplicationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupObjectReplicationPolicyArgs struct {
	AccountName               string `pulumi:"accountName"`
	ObjectReplicationPolicyId string `pulumi:"objectReplicationPolicyId"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupObjectReplicationPolicyResult struct {
	DestinationAccount string                                `pulumi:"destinationAccount"`
	EnabledTime        string                                `pulumi:"enabledTime"`
	Id                 string                                `pulumi:"id"`
	Name               string                                `pulumi:"name"`
	PolicyId           string                                `pulumi:"policyId"`
	Rules              []ObjectReplicationPolicyRuleResponse `pulumi:"rules"`
	SourceAccount      string                                `pulumi:"sourceAccount"`
	Type               string                                `pulumi:"type"`
}

func LookupObjectReplicationPolicyOutput(ctx *pulumi.Context, args LookupObjectReplicationPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupObjectReplicationPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupObjectReplicationPolicyResult, error) {
			args := v.(LookupObjectReplicationPolicyArgs)
			r, err := LookupObjectReplicationPolicy(ctx, &args, opts...)
			var s LookupObjectReplicationPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupObjectReplicationPolicyResultOutput)
}

type LookupObjectReplicationPolicyOutputArgs struct {
	AccountName               pulumi.StringInput `pulumi:"accountName"`
	ObjectReplicationPolicyId pulumi.StringInput `pulumi:"objectReplicationPolicyId"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupObjectReplicationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectReplicationPolicyArgs)(nil)).Elem()
}


type LookupObjectReplicationPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupObjectReplicationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectReplicationPolicyResult)(nil)).Elem()
}

func (o LookupObjectReplicationPolicyResultOutput) ToLookupObjectReplicationPolicyResultOutput() LookupObjectReplicationPolicyResultOutput {
	return o
}

func (o LookupObjectReplicationPolicyResultOutput) ToLookupObjectReplicationPolicyResultOutputWithContext(ctx context.Context) LookupObjectReplicationPolicyResultOutput {
	return o
}

func (o LookupObjectReplicationPolicyResultOutput) DestinationAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.DestinationAccount }).(pulumi.StringOutput)
}

func (o LookupObjectReplicationPolicyResultOutput) EnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.EnabledTime }).(pulumi.StringOutput)
}

func (o LookupObjectReplicationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupObjectReplicationPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupObjectReplicationPolicyResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

func (o LookupObjectReplicationPolicyResultOutput) Rules() ObjectReplicationPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) []ObjectReplicationPolicyRuleResponse { return v.Rules }).(ObjectReplicationPolicyRuleResponseArrayOutput)
}

func (o LookupObjectReplicationPolicyResultOutput) SourceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.SourceAccount }).(pulumi.StringOutput)
}

func (o LookupObjectReplicationPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupObjectReplicationPolicyResultOutput{})
}
