


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyFragment(ctx *pulumi.Context, args *LookupPolicyFragmentArgs, opts ...pulumi.InvokeOption) (*LookupPolicyFragmentResult, error) {
	var rv LookupPolicyFragmentResult
	err := ctx.Invoke("azure-native:apimanagement/v20211201preview:getPolicyFragment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyFragmentArgs struct {
	Format            *string `pulumi:"format"`
	Id                string  `pulumi:"id"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type LookupPolicyFragmentResult struct {
	Description *string `pulumi:"description"`
	Format      *string `pulumi:"format"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	Type        string  `pulumi:"type"`
	Value       string  `pulumi:"value"`
}


func (val *LookupPolicyFragmentResult) Defaults() *LookupPolicyFragmentResult {
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

func LookupPolicyFragmentOutput(ctx *pulumi.Context, args LookupPolicyFragmentOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyFragmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyFragmentResult, error) {
			args := v.(LookupPolicyFragmentArgs)
			r, err := LookupPolicyFragment(ctx, &args, opts...)
			var s LookupPolicyFragmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyFragmentResultOutput)
}

type LookupPolicyFragmentOutputArgs struct {
	Format            pulumi.StringPtrInput `pulumi:"format"`
	Id                pulumi.StringInput    `pulumi:"id"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput    `pulumi:"serviceName"`
}

func (LookupPolicyFragmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyFragmentArgs)(nil)).Elem()
}


type LookupPolicyFragmentResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyFragmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyFragmentResult)(nil)).Elem()
}

func (o LookupPolicyFragmentResultOutput) ToLookupPolicyFragmentResultOutput() LookupPolicyFragmentResultOutput {
	return o
}

func (o LookupPolicyFragmentResultOutput) ToLookupPolicyFragmentResultOutputWithContext(ctx context.Context) LookupPolicyFragmentResultOutput {
	return o
}

func (o LookupPolicyFragmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyFragmentResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyFragmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyFragmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicyFragmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPolicyFragmentResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyFragmentResultOutput{})
}
