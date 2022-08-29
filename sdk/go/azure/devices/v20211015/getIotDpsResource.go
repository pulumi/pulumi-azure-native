


package v20211015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotDpsResource(ctx *pulumi.Context, args *LookupIotDpsResourceArgs, opts ...pulumi.InvokeOption) (*LookupIotDpsResourceResult, error) {
	var rv LookupIotDpsResourceResult
	err := ctx.Invoke("azure-native:devices/v20211015:getIotDpsResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotDpsResourceArgs struct {
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupIotDpsResourceResult struct {
	Etag       *string                             `pulumi:"etag"`
	Id         string                              `pulumi:"id"`
	Location   string                              `pulumi:"location"`
	Name       string                              `pulumi:"name"`
	Properties IotDpsPropertiesDescriptionResponse `pulumi:"properties"`
	Sku        IotDpsSkuInfoResponse               `pulumi:"sku"`
	SystemData SystemDataResponse                  `pulumi:"systemData"`
	Tags       map[string]string                   `pulumi:"tags"`
	Type       string                              `pulumi:"type"`
}

func LookupIotDpsResourceOutput(ctx *pulumi.Context, args LookupIotDpsResourceOutputArgs, opts ...pulumi.InvokeOption) LookupIotDpsResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotDpsResourceResult, error) {
			args := v.(LookupIotDpsResourceArgs)
			r, err := LookupIotDpsResource(ctx, &args, opts...)
			var s LookupIotDpsResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotDpsResourceResultOutput)
}

type LookupIotDpsResourceOutputArgs struct {
	ProvisioningServiceName pulumi.StringInput `pulumi:"provisioningServiceName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIotDpsResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotDpsResourceArgs)(nil)).Elem()
}


type LookupIotDpsResourceResultOutput struct{ *pulumi.OutputState }

func (LookupIotDpsResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotDpsResourceResult)(nil)).Elem()
}

func (o LookupIotDpsResourceResultOutput) ToLookupIotDpsResourceResultOutput() LookupIotDpsResourceResultOutput {
	return o
}

func (o LookupIotDpsResourceResultOutput) ToLookupIotDpsResourceResultOutputWithContext(ctx context.Context) LookupIotDpsResourceResultOutput {
	return o
}

func (o LookupIotDpsResourceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupIotDpsResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIotDpsResourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupIotDpsResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIotDpsResourceResultOutput) Properties() IotDpsPropertiesDescriptionResponseOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) IotDpsPropertiesDescriptionResponse { return v.Properties }).(IotDpsPropertiesDescriptionResponseOutput)
}

func (o LookupIotDpsResourceResultOutput) Sku() IotDpsSkuInfoResponseOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) IotDpsSkuInfoResponse { return v.Sku }).(IotDpsSkuInfoResponseOutput)
}

func (o LookupIotDpsResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupIotDpsResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIotDpsResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotDpsResourceResultOutput{})
}
