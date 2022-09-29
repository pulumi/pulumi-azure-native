


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	var rv LookupPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20210401preview:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyArgs struct {
	Format            *string `pulumi:"format"`
	PolicyId          string  `pulumi:"policyId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type LookupPolicyResult struct {
	Format *string `pulumi:"format"`
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Type   string  `pulumi:"type"`
	Value  string  `pulumi:"value"`
}


func (val *LookupPolicyResult) Defaults() *LookupPolicyResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Format) {
		format_ := "xml"
		tmp.Format = &format_
	}
	return &tmp
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

type LookupPolicyOutputArgs struct {
	Format            pulumi.StringPtrInput `pulumi:"format"`
	PolicyId          pulumi.StringInput    `pulumi:"policyId"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput    `pulumi:"serviceName"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}


type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
