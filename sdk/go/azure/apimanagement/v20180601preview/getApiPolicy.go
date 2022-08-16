


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiPolicy(ctx *pulumi.Context, args *LookupApiPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApiPolicyResult, error) {
	var rv LookupApiPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20180601preview:getApiPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiPolicyArgs struct {
	ApiId             string `pulumi:"apiId"`
	PolicyId          string `pulumi:"policyId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiPolicyResult struct {
	ContentFormat *string `pulumi:"contentFormat"`
	Id            string  `pulumi:"id"`
	Name          string  `pulumi:"name"`
	PolicyContent string  `pulumi:"policyContent"`
	Type          string  `pulumi:"type"`
}


func (val *LookupApiPolicyResult) Defaults() *LookupApiPolicyResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ContentFormat) {
		contentFormat_ := "xml"
		tmp.ContentFormat = &contentFormat_
	}
	return &tmp
}

func LookupApiPolicyOutput(ctx *pulumi.Context, args LookupApiPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupApiPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiPolicyResult, error) {
			args := v.(LookupApiPolicyArgs)
			r, err := LookupApiPolicy(ctx, &args, opts...)
			var s LookupApiPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiPolicyResultOutput)
}

type LookupApiPolicyOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	PolicyId          pulumi.StringInput `pulumi:"policyId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPolicyArgs)(nil)).Elem()
}


type LookupApiPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupApiPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPolicyResult)(nil)).Elem()
}

func (o LookupApiPolicyResultOutput) ToLookupApiPolicyResultOutput() LookupApiPolicyResultOutput {
	return o
}

func (o LookupApiPolicyResultOutput) ToLookupApiPolicyResultOutputWithContext(ctx context.Context) LookupApiPolicyResultOutput {
	return o
}

func (o LookupApiPolicyResultOutput) ContentFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiPolicyResult) *string { return v.ContentFormat }).(pulumi.StringPtrOutput)
}

func (o LookupApiPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiPolicyResultOutput) PolicyContent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPolicyResult) string { return v.PolicyContent }).(pulumi.StringOutput)
}

func (o LookupApiPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiPolicyResultOutput{})
}
