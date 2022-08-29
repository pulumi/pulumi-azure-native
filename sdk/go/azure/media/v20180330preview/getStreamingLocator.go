


package v20180330preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStreamingLocator(ctx *pulumi.Context, args *LookupStreamingLocatorArgs, opts ...pulumi.InvokeOption) (*LookupStreamingLocatorResult, error) {
	var rv LookupStreamingLocatorResult
	err := ctx.Invoke("azure-native:media/v20180330preview:getStreamingLocator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingLocatorArgs struct {
	AccountName          string `pulumi:"accountName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	StreamingLocatorName string `pulumi:"streamingLocatorName"`
}


type LookupStreamingLocatorResult struct {
	AssetName                   string                                          `pulumi:"assetName"`
	ContentKeys                 []StreamingLocatorUserDefinedContentKeyResponse `pulumi:"contentKeys"`
	Created                     string                                          `pulumi:"created"`
	DefaultContentKeyPolicyName *string                                         `pulumi:"defaultContentKeyPolicyName"`
	EndTime                     *string                                         `pulumi:"endTime"`
	Id                          string                                          `pulumi:"id"`
	Name                        string                                          `pulumi:"name"`
	StartTime                   *string                                         `pulumi:"startTime"`
	StreamingLocatorId          *string                                         `pulumi:"streamingLocatorId"`
	StreamingPolicyName         string                                          `pulumi:"streamingPolicyName"`
	Type                        string                                          `pulumi:"type"`
}

func LookupStreamingLocatorOutput(ctx *pulumi.Context, args LookupStreamingLocatorOutputArgs, opts ...pulumi.InvokeOption) LookupStreamingLocatorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamingLocatorResult, error) {
			args := v.(LookupStreamingLocatorArgs)
			r, err := LookupStreamingLocator(ctx, &args, opts...)
			var s LookupStreamingLocatorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStreamingLocatorResultOutput)
}

type LookupStreamingLocatorOutputArgs struct {
	AccountName          pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	StreamingLocatorName pulumi.StringInput `pulumi:"streamingLocatorName"`
}

func (LookupStreamingLocatorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingLocatorArgs)(nil)).Elem()
}


type LookupStreamingLocatorResultOutput struct{ *pulumi.OutputState }

func (LookupStreamingLocatorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingLocatorResult)(nil)).Elem()
}

func (o LookupStreamingLocatorResultOutput) ToLookupStreamingLocatorResultOutput() LookupStreamingLocatorResultOutput {
	return o
}

func (o LookupStreamingLocatorResultOutput) ToLookupStreamingLocatorResultOutputWithContext(ctx context.Context) LookupStreamingLocatorResultOutput {
	return o
}

func (o LookupStreamingLocatorResultOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o LookupStreamingLocatorResultOutput) ContentKeys() StreamingLocatorUserDefinedContentKeyResponseArrayOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) []StreamingLocatorUserDefinedContentKeyResponse {
		return v.ContentKeys
	}).(StreamingLocatorUserDefinedContentKeyResponseArrayOutput)
}

func (o LookupStreamingLocatorResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupStreamingLocatorResultOutput) DefaultContentKeyPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) *string { return v.DefaultContentKeyPolicyName }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingLocatorResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingLocatorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStreamingLocatorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStreamingLocatorResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingLocatorResultOutput) StreamingLocatorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) *string { return v.StreamingLocatorId }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingLocatorResultOutput) StreamingPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.StreamingPolicyName }).(pulumi.StringOutput)
}

func (o LookupStreamingLocatorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamingLocatorResultOutput{})
}
