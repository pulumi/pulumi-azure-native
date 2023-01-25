


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStreamingLocatorPaths(ctx *pulumi.Context, args *ListStreamingLocatorPathsArgs, opts ...pulumi.InvokeOption) (*ListStreamingLocatorPathsResult, error) {
	var rv ListStreamingLocatorPathsResult
	err := ctx.Invoke("azure-native:media/v20180601preview:listStreamingLocatorPaths", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStreamingLocatorPathsArgs struct {
	AccountName          string `pulumi:"accountName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	StreamingLocatorName string `pulumi:"streamingLocatorName"`
}


type ListStreamingLocatorPathsResult struct {
	DownloadPaths  []string                `pulumi:"downloadPaths"`
	StreamingPaths []StreamingPathResponse `pulumi:"streamingPaths"`
}

func ListStreamingLocatorPathsOutput(ctx *pulumi.Context, args ListStreamingLocatorPathsOutputArgs, opts ...pulumi.InvokeOption) ListStreamingLocatorPathsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStreamingLocatorPathsResult, error) {
			args := v.(ListStreamingLocatorPathsArgs)
			r, err := ListStreamingLocatorPaths(ctx, &args, opts...)
			var s ListStreamingLocatorPathsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStreamingLocatorPathsResultOutput)
}

type ListStreamingLocatorPathsOutputArgs struct {
	AccountName          pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	StreamingLocatorName pulumi.StringInput `pulumi:"streamingLocatorName"`
}

func (ListStreamingLocatorPathsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStreamingLocatorPathsArgs)(nil)).Elem()
}


type ListStreamingLocatorPathsResultOutput struct{ *pulumi.OutputState }

func (ListStreamingLocatorPathsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStreamingLocatorPathsResult)(nil)).Elem()
}

func (o ListStreamingLocatorPathsResultOutput) ToListStreamingLocatorPathsResultOutput() ListStreamingLocatorPathsResultOutput {
	return o
}

func (o ListStreamingLocatorPathsResultOutput) ToListStreamingLocatorPathsResultOutputWithContext(ctx context.Context) ListStreamingLocatorPathsResultOutput {
	return o
}

func (o ListStreamingLocatorPathsResultOutput) DownloadPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListStreamingLocatorPathsResult) []string { return v.DownloadPaths }).(pulumi.StringArrayOutput)
}

func (o ListStreamingLocatorPathsResultOutput) StreamingPaths() StreamingPathResponseArrayOutput {
	return o.ApplyT(func(v ListStreamingLocatorPathsResult) []StreamingPathResponse { return v.StreamingPaths }).(StreamingPathResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStreamingLocatorPathsResultOutput{})
}
