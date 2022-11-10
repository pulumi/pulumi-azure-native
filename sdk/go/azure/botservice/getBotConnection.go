


package botservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBotConnection(ctx *pulumi.Context, args *LookupBotConnectionArgs, opts ...pulumi.InvokeOption) (*LookupBotConnectionResult, error) {
	var rv LookupBotConnectionResult
	err := ctx.Invoke("azure-native:botservice:getBotConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBotConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupBotConnectionResult struct {
	Etag       *string                             `pulumi:"etag"`
	Id         string                              `pulumi:"id"`
	Kind       *string                             `pulumi:"kind"`
	Location   *string                             `pulumi:"location"`
	Name       string                              `pulumi:"name"`
	Properties ConnectionSettingPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                        `pulumi:"sku"`
	Tags       map[string]string                   `pulumi:"tags"`
	Type       string                              `pulumi:"type"`
	Zones      []string                            `pulumi:"zones"`
}

func LookupBotConnectionOutput(ctx *pulumi.Context, args LookupBotConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupBotConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBotConnectionResult, error) {
			args := v.(LookupBotConnectionArgs)
			r, err := LookupBotConnection(ctx, &args, opts...)
			var s LookupBotConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBotConnectionResultOutput)
}

type LookupBotConnectionOutputArgs struct {
	ConnectionName    pulumi.StringInput `pulumi:"connectionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupBotConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotConnectionArgs)(nil)).Elem()
}


type LookupBotConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupBotConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotConnectionResult)(nil)).Elem()
}

func (o LookupBotConnectionResultOutput) ToLookupBotConnectionResultOutput() LookupBotConnectionResultOutput {
	return o
}

func (o LookupBotConnectionResultOutput) ToLookupBotConnectionResultOutputWithContext(ctx context.Context) LookupBotConnectionResultOutput {
	return o
}

func (o LookupBotConnectionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupBotConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBotConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupBotConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupBotConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBotConnectionResultOutput) Properties() ConnectionSettingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) ConnectionSettingPropertiesResponse { return v.Properties }).(ConnectionSettingPropertiesResponseOutput)
}

func (o LookupBotConnectionResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupBotConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBotConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBotConnectionResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBotConnectionResultOutput{})
}
