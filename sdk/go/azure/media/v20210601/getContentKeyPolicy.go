


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContentKeyPolicy(ctx *pulumi.Context, args *LookupContentKeyPolicyArgs, opts ...pulumi.InvokeOption) (*LookupContentKeyPolicyResult, error) {
	var rv LookupContentKeyPolicyResult
	err := ctx.Invoke("azure-native:media/v20210601:getContentKeyPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentKeyPolicyArgs struct {
	AccountName          string `pulumi:"accountName"`
	ContentKeyPolicyName string `pulumi:"contentKeyPolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupContentKeyPolicyResult struct {
	Created      string                           `pulumi:"created"`
	Description  *string                          `pulumi:"description"`
	Id           string                           `pulumi:"id"`
	LastModified string                           `pulumi:"lastModified"`
	Name         string                           `pulumi:"name"`
	Options      []ContentKeyPolicyOptionResponse `pulumi:"options"`
	PolicyId     string                           `pulumi:"policyId"`
	SystemData   SystemDataResponse               `pulumi:"systemData"`
	Type         string                           `pulumi:"type"`
}

func LookupContentKeyPolicyOutput(ctx *pulumi.Context, args LookupContentKeyPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupContentKeyPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContentKeyPolicyResult, error) {
			args := v.(LookupContentKeyPolicyArgs)
			r, err := LookupContentKeyPolicy(ctx, &args, opts...)
			var s LookupContentKeyPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContentKeyPolicyResultOutput)
}

type LookupContentKeyPolicyOutputArgs struct {
	AccountName          pulumi.StringInput `pulumi:"accountName"`
	ContentKeyPolicyName pulumi.StringInput `pulumi:"contentKeyPolicyName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContentKeyPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentKeyPolicyArgs)(nil)).Elem()
}


type LookupContentKeyPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupContentKeyPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentKeyPolicyResult)(nil)).Elem()
}

func (o LookupContentKeyPolicyResultOutput) ToLookupContentKeyPolicyResultOutput() LookupContentKeyPolicyResultOutput {
	return o
}

func (o LookupContentKeyPolicyResultOutput) ToLookupContentKeyPolicyResultOutputWithContext(ctx context.Context) LookupContentKeyPolicyResultOutput {
	return o
}

func (o LookupContentKeyPolicyResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupContentKeyPolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupContentKeyPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContentKeyPolicyResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupContentKeyPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContentKeyPolicyResultOutput) Options() ContentKeyPolicyOptionResponseArrayOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) []ContentKeyPolicyOptionResponse { return v.Options }).(ContentKeyPolicyOptionResponseArrayOutput)
}

func (o LookupContentKeyPolicyResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

func (o LookupContentKeyPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupContentKeyPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContentKeyPolicyResultOutput{})
}
