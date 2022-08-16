


package v20210210

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProtectionPolicy(ctx *pulumi.Context, args *LookupProtectionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupProtectionPolicyResult, error) {
	var rv LookupProtectionPolicyResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210210:getProtectionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectionPolicyArgs struct {
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupProtectionPolicyResult struct {
	ETag       *string           `pulumi:"eTag"`
	Id         string            `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupProtectionPolicyOutput(ctx *pulumi.Context, args LookupProtectionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupProtectionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProtectionPolicyResult, error) {
			args := v.(LookupProtectionPolicyArgs)
			r, err := LookupProtectionPolicy(ctx, &args, opts...)
			var s LookupProtectionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProtectionPolicyResultOutput)
}

type LookupProtectionPolicyOutputArgs struct {
	PolicyName        pulumi.StringInput `pulumi:"policyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VaultName         pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupProtectionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionPolicyArgs)(nil)).Elem()
}


type LookupProtectionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupProtectionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionPolicyResult)(nil)).Elem()
}

func (o LookupProtectionPolicyResultOutput) ToLookupProtectionPolicyResultOutput() LookupProtectionPolicyResultOutput {
	return o
}

func (o LookupProtectionPolicyResultOutput) ToLookupProtectionPolicyResultOutputWithContext(ctx context.Context) LookupProtectionPolicyResultOutput {
	return o
}

func (o LookupProtectionPolicyResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupProtectionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProtectionPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupProtectionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProtectionPolicyResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupProtectionPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProtectionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProtectionPolicyResultOutput{})
}
