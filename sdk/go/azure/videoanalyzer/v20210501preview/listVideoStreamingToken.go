


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVideoStreamingToken(ctx *pulumi.Context, args *ListVideoStreamingTokenArgs, opts ...pulumi.InvokeOption) (*ListVideoStreamingTokenResult, error) {
	var rv ListVideoStreamingTokenResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20210501preview:listVideoStreamingToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVideoStreamingTokenArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VideoName         string `pulumi:"videoName"`
}


type ListVideoStreamingTokenResult struct {
	ExpirationDate string `pulumi:"expirationDate"`
	Token          string `pulumi:"token"`
}

func ListVideoStreamingTokenOutput(ctx *pulumi.Context, args ListVideoStreamingTokenOutputArgs, opts ...pulumi.InvokeOption) ListVideoStreamingTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVideoStreamingTokenResult, error) {
			args := v.(ListVideoStreamingTokenArgs)
			r, err := ListVideoStreamingToken(ctx, &args, opts...)
			var s ListVideoStreamingTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVideoStreamingTokenResultOutput)
}

type ListVideoStreamingTokenOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VideoName         pulumi.StringInput `pulumi:"videoName"`
}

func (ListVideoStreamingTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVideoStreamingTokenArgs)(nil)).Elem()
}


type ListVideoStreamingTokenResultOutput struct{ *pulumi.OutputState }

func (ListVideoStreamingTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVideoStreamingTokenResult)(nil)).Elem()
}

func (o ListVideoStreamingTokenResultOutput) ToListVideoStreamingTokenResultOutput() ListVideoStreamingTokenResultOutput {
	return o
}

func (o ListVideoStreamingTokenResultOutput) ToListVideoStreamingTokenResultOutputWithContext(ctx context.Context) ListVideoStreamingTokenResultOutput {
	return o
}

func (o ListVideoStreamingTokenResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v ListVideoStreamingTokenResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o ListVideoStreamingTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v ListVideoStreamingTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVideoStreamingTokenResultOutput{})
}
