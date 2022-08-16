


package v20180712

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBot(ctx *pulumi.Context, args *LookupBotArgs, opts ...pulumi.InvokeOption) (*LookupBotResult, error) {
	var rv LookupBotResult
	err := ctx.Invoke("azure-native:botservice/v20180712:getBot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupBotResult struct {
	Etag       *string               `pulumi:"etag"`
	Id         string                `pulumi:"id"`
	Kind       *string               `pulumi:"kind"`
	Location   *string               `pulumi:"location"`
	Name       string                `pulumi:"name"`
	Properties BotPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse          `pulumi:"sku"`
	Tags       map[string]string     `pulumi:"tags"`
	Type       string                `pulumi:"type"`
}

func LookupBotOutput(ctx *pulumi.Context, args LookupBotOutputArgs, opts ...pulumi.InvokeOption) LookupBotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBotResult, error) {
			args := v.(LookupBotArgs)
			r, err := LookupBot(ctx, &args, opts...)
			var s LookupBotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBotResultOutput)
}

type LookupBotOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupBotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotArgs)(nil)).Elem()
}


type LookupBotResultOutput struct{ *pulumi.OutputState }

func (LookupBotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotResult)(nil)).Elem()
}

func (o LookupBotResultOutput) ToLookupBotResultOutput() LookupBotResultOutput {
	return o
}

func (o LookupBotResultOutput) ToLookupBotResultOutputWithContext(ctx context.Context) LookupBotResultOutput {
	return o
}

func (o LookupBotResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupBotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupBotResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupBotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBotResultOutput) Properties() BotPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBotResult) BotPropertiesResponse { return v.Properties }).(BotPropertiesResponseOutput)
}

func (o LookupBotResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupBotResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupBotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBotResultOutput{})
}
