


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStreamingLocatorContentKeys(ctx *pulumi.Context, args *ListStreamingLocatorContentKeysArgs, opts ...pulumi.InvokeOption) (*ListStreamingLocatorContentKeysResult, error) {
	var rv ListStreamingLocatorContentKeysResult
	err := ctx.Invoke("azure-native:media/v20180601preview:listStreamingLocatorContentKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStreamingLocatorContentKeysArgs struct {
	AccountName          string `pulumi:"accountName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	StreamingLocatorName string `pulumi:"streamingLocatorName"`
}


type ListStreamingLocatorContentKeysResult struct {
	ContentKeys []StreamingLocatorContentKeyResponse `pulumi:"contentKeys"`
}

func ListStreamingLocatorContentKeysOutput(ctx *pulumi.Context, args ListStreamingLocatorContentKeysOutputArgs, opts ...pulumi.InvokeOption) ListStreamingLocatorContentKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStreamingLocatorContentKeysResult, error) {
			args := v.(ListStreamingLocatorContentKeysArgs)
			r, err := ListStreamingLocatorContentKeys(ctx, &args, opts...)
			var s ListStreamingLocatorContentKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStreamingLocatorContentKeysResultOutput)
}

type ListStreamingLocatorContentKeysOutputArgs struct {
	AccountName          pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	StreamingLocatorName pulumi.StringInput `pulumi:"streamingLocatorName"`
}

func (ListStreamingLocatorContentKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStreamingLocatorContentKeysArgs)(nil)).Elem()
}


type ListStreamingLocatorContentKeysResultOutput struct{ *pulumi.OutputState }

func (ListStreamingLocatorContentKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStreamingLocatorContentKeysResult)(nil)).Elem()
}

func (o ListStreamingLocatorContentKeysResultOutput) ToListStreamingLocatorContentKeysResultOutput() ListStreamingLocatorContentKeysResultOutput {
	return o
}

func (o ListStreamingLocatorContentKeysResultOutput) ToListStreamingLocatorContentKeysResultOutputWithContext(ctx context.Context) ListStreamingLocatorContentKeysResultOutput {
	return o
}

func (o ListStreamingLocatorContentKeysResultOutput) ContentKeys() StreamingLocatorContentKeyResponseArrayOutput {
	return o.ApplyT(func(v ListStreamingLocatorContentKeysResult) []StreamingLocatorContentKeyResponse {
		return v.ContentKeys
	}).(StreamingLocatorContentKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStreamingLocatorContentKeysResultOutput{})
}
