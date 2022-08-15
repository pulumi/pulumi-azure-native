


package v20180701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStreamingPolicy(ctx *pulumi.Context, args *LookupStreamingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupStreamingPolicyResult, error) {
	var rv LookupStreamingPolicyResult
	err := ctx.Invoke("azure-native:media/v20180701:getStreamingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingPolicyArgs struct {
	AccountName         string `pulumi:"accountName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	StreamingPolicyName string `pulumi:"streamingPolicyName"`
}


type LookupStreamingPolicyResult struct {
	CommonEncryptionCbcs        *CommonEncryptionCbcsResponse `pulumi:"commonEncryptionCbcs"`
	CommonEncryptionCenc        *CommonEncryptionCencResponse `pulumi:"commonEncryptionCenc"`
	Created                     string                        `pulumi:"created"`
	DefaultContentKeyPolicyName *string                       `pulumi:"defaultContentKeyPolicyName"`
	EnvelopeEncryption          *EnvelopeEncryptionResponse   `pulumi:"envelopeEncryption"`
	Id                          string                        `pulumi:"id"`
	Name                        string                        `pulumi:"name"`
	NoEncryption                *NoEncryptionResponse         `pulumi:"noEncryption"`
	Type                        string                        `pulumi:"type"`
}

func LookupStreamingPolicyOutput(ctx *pulumi.Context, args LookupStreamingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupStreamingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamingPolicyResult, error) {
			args := v.(LookupStreamingPolicyArgs)
			r, err := LookupStreamingPolicy(ctx, &args, opts...)
			var s LookupStreamingPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStreamingPolicyResultOutput)
}

type LookupStreamingPolicyOutputArgs struct {
	AccountName         pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	StreamingPolicyName pulumi.StringInput `pulumi:"streamingPolicyName"`
}

func (LookupStreamingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingPolicyArgs)(nil)).Elem()
}


type LookupStreamingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupStreamingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingPolicyResult)(nil)).Elem()
}

func (o LookupStreamingPolicyResultOutput) ToLookupStreamingPolicyResultOutput() LookupStreamingPolicyResultOutput {
	return o
}

func (o LookupStreamingPolicyResultOutput) ToLookupStreamingPolicyResultOutputWithContext(ctx context.Context) LookupStreamingPolicyResultOutput {
	return o
}

func (o LookupStreamingPolicyResultOutput) CommonEncryptionCbcs() CommonEncryptionCbcsResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingPolicyResult) *CommonEncryptionCbcsResponse { return v.CommonEncryptionCbcs }).(CommonEncryptionCbcsResponsePtrOutput)
}

func (o LookupStreamingPolicyResultOutput) CommonEncryptionCenc() CommonEncryptionCencResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingPolicyResult) *CommonEncryptionCencResponse { return v.CommonEncryptionCenc }).(CommonEncryptionCencResponsePtrOutput)
}

func (o LookupStreamingPolicyResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingPolicyResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupStreamingPolicyResultOutput) DefaultContentKeyPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingPolicyResult) *string { return v.DefaultContentKeyPolicyName }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingPolicyResultOutput) EnvelopeEncryption() EnvelopeEncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingPolicyResult) *EnvelopeEncryptionResponse { return v.EnvelopeEncryption }).(EnvelopeEncryptionResponsePtrOutput)
}

func (o LookupStreamingPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStreamingPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStreamingPolicyResultOutput) NoEncryption() NoEncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingPolicyResult) *NoEncryptionResponse { return v.NoEncryption }).(NoEncryptionResponsePtrOutput)
}

func (o LookupStreamingPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamingPolicyResultOutput{})
}
