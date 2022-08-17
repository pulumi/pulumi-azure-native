


package botservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnterpriseChannel(ctx *pulumi.Context, args *LookupEnterpriseChannelArgs, opts ...pulumi.InvokeOption) (*LookupEnterpriseChannelResult, error) {
	var rv LookupEnterpriseChannelResult
	err := ctx.Invoke("azure-native:botservice:getEnterpriseChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterpriseChannelArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupEnterpriseChannelResult struct {
	Etag       *string                             `pulumi:"etag"`
	Id         string                              `pulumi:"id"`
	Kind       *string                             `pulumi:"kind"`
	Location   *string                             `pulumi:"location"`
	Name       string                              `pulumi:"name"`
	Properties EnterpriseChannelPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                        `pulumi:"sku"`
	Tags       map[string]string                   `pulumi:"tags"`
	Type       string                              `pulumi:"type"`
}

func LookupEnterpriseChannelOutput(ctx *pulumi.Context, args LookupEnterpriseChannelOutputArgs, opts ...pulumi.InvokeOption) LookupEnterpriseChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnterpriseChannelResult, error) {
			args := v.(LookupEnterpriseChannelArgs)
			r, err := LookupEnterpriseChannel(ctx, &args, opts...)
			var s LookupEnterpriseChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnterpriseChannelResultOutput)
}

type LookupEnterpriseChannelOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupEnterpriseChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterpriseChannelArgs)(nil)).Elem()
}


type LookupEnterpriseChannelResultOutput struct{ *pulumi.OutputState }

func (LookupEnterpriseChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterpriseChannelResult)(nil)).Elem()
}

func (o LookupEnterpriseChannelResultOutput) ToLookupEnterpriseChannelResultOutput() LookupEnterpriseChannelResultOutput {
	return o
}

func (o LookupEnterpriseChannelResultOutput) ToLookupEnterpriseChannelResultOutputWithContext(ctx context.Context) LookupEnterpriseChannelResultOutput {
	return o
}

func (o LookupEnterpriseChannelResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupEnterpriseChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnterpriseChannelResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupEnterpriseChannelResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEnterpriseChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnterpriseChannelResultOutput) Properties() EnterpriseChannelPropertiesResponseOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) EnterpriseChannelPropertiesResponse { return v.Properties }).(EnterpriseChannelPropertiesResponseOutput)
}

func (o LookupEnterpriseChannelResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupEnterpriseChannelResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEnterpriseChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnterpriseChannelResultOutput{})
}
