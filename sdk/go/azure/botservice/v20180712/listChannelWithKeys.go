


package v20180712

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListChannelWithKeys(ctx *pulumi.Context, args *ListChannelWithKeysArgs, opts ...pulumi.InvokeOption) (*ListChannelWithKeysResult, error) {
	var rv ListChannelWithKeysResult
	err := ctx.Invoke("azure-native:botservice/v20180712:listChannelWithKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListChannelWithKeysArgs struct {
	ChannelName       string `pulumi:"channelName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListChannelWithKeysResult struct {
	Etag       *string           `pulumi:"etag"`
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func ListChannelWithKeysOutput(ctx *pulumi.Context, args ListChannelWithKeysOutputArgs, opts ...pulumi.InvokeOption) ListChannelWithKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListChannelWithKeysResult, error) {
			args := v.(ListChannelWithKeysArgs)
			r, err := ListChannelWithKeys(ctx, &args, opts...)
			var s ListChannelWithKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListChannelWithKeysResultOutput)
}

type ListChannelWithKeysOutputArgs struct {
	ChannelName       pulumi.StringInput `pulumi:"channelName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListChannelWithKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListChannelWithKeysArgs)(nil)).Elem()
}


type ListChannelWithKeysResultOutput struct{ *pulumi.OutputState }

func (ListChannelWithKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListChannelWithKeysResult)(nil)).Elem()
}

func (o ListChannelWithKeysResultOutput) ToListChannelWithKeysResultOutput() ListChannelWithKeysResultOutput {
	return o
}

func (o ListChannelWithKeysResultOutput) ToListChannelWithKeysResultOutputWithContext(ctx context.Context) ListChannelWithKeysResultOutput {
	return o
}

func (o ListChannelWithKeysResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ListChannelWithKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListChannelWithKeysResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListChannelWithKeysResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ListChannelWithKeysResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListChannelWithKeysResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o ListChannelWithKeysResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ListChannelWithKeysResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListChannelWithKeysResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListChannelWithKeysResultOutput{})
}
