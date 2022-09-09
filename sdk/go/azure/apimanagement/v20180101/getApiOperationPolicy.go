


package v20180101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiOperationPolicy(ctx *pulumi.Context, args *LookupApiOperationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApiOperationPolicyResult, error) {
	var rv LookupApiOperationPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getApiOperationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiOperationPolicyArgs struct {
	ApiId             string `pulumi:"apiId"`
	OperationId       string `pulumi:"operationId"`
	PolicyId          string `pulumi:"policyId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiOperationPolicyResult struct {
	ContentFormat *string `pulumi:"contentFormat"`
	Id            string  `pulumi:"id"`
	Name          string  `pulumi:"name"`
	PolicyContent string  `pulumi:"policyContent"`
	Type          string  `pulumi:"type"`
}


func (val *LookupApiOperationPolicyResult) Defaults() *LookupApiOperationPolicyResult {
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

func LookupApiOperationPolicyOutput(ctx *pulumi.Context, args LookupApiOperationPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupApiOperationPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiOperationPolicyResult, error) {
			args := v.(LookupApiOperationPolicyArgs)
			r, err := LookupApiOperationPolicy(ctx, &args, opts...)
			var s LookupApiOperationPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiOperationPolicyResultOutput)
}

type LookupApiOperationPolicyOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	OperationId       pulumi.StringInput `pulumi:"operationId"`
	PolicyId          pulumi.StringInput `pulumi:"policyId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiOperationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiOperationPolicyArgs)(nil)).Elem()
}


type LookupApiOperationPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupApiOperationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiOperationPolicyResult)(nil)).Elem()
}

func (o LookupApiOperationPolicyResultOutput) ToLookupApiOperationPolicyResultOutput() LookupApiOperationPolicyResultOutput {
	return o
}

func (o LookupApiOperationPolicyResultOutput) ToLookupApiOperationPolicyResultOutputWithContext(ctx context.Context) LookupApiOperationPolicyResultOutput {
	return o
}

func (o LookupApiOperationPolicyResultOutput) ContentFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiOperationPolicyResult) *string { return v.ContentFormat }).(pulumi.StringPtrOutput)
}

func (o LookupApiOperationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiOperationPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiOperationPolicyResultOutput) PolicyContent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationPolicyResult) string { return v.PolicyContent }).(pulumi.StringOutput)
}

func (o LookupApiOperationPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiOperationPolicyResultOutput{})
}
