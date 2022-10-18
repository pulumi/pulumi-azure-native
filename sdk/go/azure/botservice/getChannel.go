


package botservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupChannel(ctx *pulumi.Context, args *LookupChannelArgs, opts ...pulumi.InvokeOption) (*LookupChannelResult, error) {
	var rv LookupChannelResult
	err := ctx.Invoke("azure-native:botservice:getChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelArgs struct {
	ChannelName       string `pulumi:"channelName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupChannelResult struct {
	Etag       *string           `pulumi:"etag"`
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
	Zones      []string          `pulumi:"zones"`
}

func LookupChannelOutput(ctx *pulumi.Context, args LookupChannelOutputArgs, opts ...pulumi.InvokeOption) LookupChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelResult, error) {
			args := v.(LookupChannelArgs)
			r, err := LookupChannel(ctx, &args, opts...)
			var s LookupChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChannelResultOutput)
}

type LookupChannelOutputArgs struct {
	ChannelName       pulumi.StringInput `pulumi:"channelName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelArgs)(nil)).Elem()
}


type LookupChannelResultOutput struct{ *pulumi.OutputState }

func (LookupChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelResult)(nil)).Elem()
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutput() LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutputWithContext(ctx context.Context) LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupChannelResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupChannelResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupChannelResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupChannelResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupChannelResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupChannelResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupChannelResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupChannelResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupChannelResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelResultOutput{})
}
