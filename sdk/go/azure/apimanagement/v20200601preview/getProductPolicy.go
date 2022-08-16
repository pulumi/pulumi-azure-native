


package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProductPolicy(ctx *pulumi.Context, args *LookupProductPolicyArgs, opts ...pulumi.InvokeOption) (*LookupProductPolicyResult, error) {
	var rv LookupProductPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20200601preview:getProductPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupProductPolicyArgs struct {
	Format            *string `pulumi:"format"`
	PolicyId          string  `pulumi:"policyId"`
	ProductId         string  `pulumi:"productId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type LookupProductPolicyResult struct {
	Format *string `pulumi:"format"`
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Type   string  `pulumi:"type"`
	Value  string  `pulumi:"value"`
}


func (val *LookupProductPolicyResult) Defaults() *LookupProductPolicyResult {
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

func LookupProductPolicyOutput(ctx *pulumi.Context, args LookupProductPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupProductPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProductPolicyResult, error) {
			args := v.(LookupProductPolicyArgs)
			r, err := LookupProductPolicy(ctx, &args, opts...)
			var s LookupProductPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProductPolicyResultOutput)
}

type LookupProductPolicyOutputArgs struct {
	Format            pulumi.StringPtrInput `pulumi:"format"`
	PolicyId          pulumi.StringInput    `pulumi:"policyId"`
	ProductId         pulumi.StringInput    `pulumi:"productId"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput    `pulumi:"serviceName"`
}

func (LookupProductPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductPolicyArgs)(nil)).Elem()
}


type LookupProductPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupProductPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductPolicyResult)(nil)).Elem()
}

func (o LookupProductPolicyResultOutput) ToLookupProductPolicyResultOutput() LookupProductPolicyResultOutput {
	return o
}

func (o LookupProductPolicyResultOutput) ToLookupProductPolicyResultOutputWithContext(ctx context.Context) LookupProductPolicyResultOutput {
	return o
}

func (o LookupProductPolicyResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProductPolicyResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupProductPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProductPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProductPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupProductPolicyResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductPolicyResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProductPolicyResultOutput{})
}
