


package v20170119

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupIotHubResource(ctx *pulumi.Context, args *LookupIotHubResourceArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResourceResult, error) {
	var rv LookupIotHubResourceResult
	err := ctx.Invoke("azure-native:devices/v20170119:getIotHubResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotHubResourceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupIotHubResourceResult struct {
	Etag           *string                  `pulumi:"etag"`
	Id             string                   `pulumi:"id"`
	Location       string                   `pulumi:"location"`
	Name           string                   `pulumi:"name"`
	Properties     IotHubPropertiesResponse `pulumi:"properties"`
	Resourcegroup  string                   `pulumi:"resourcegroup"`
	Sku            IotHubSkuInfoResponse    `pulumi:"sku"`
	Subscriptionid string                   `pulumi:"subscriptionid"`
	Tags           map[string]string        `pulumi:"tags"`
	Type           string                   `pulumi:"type"`
}

func LookupIotHubResourceOutput(ctx *pulumi.Context, args LookupIotHubResourceOutputArgs, opts ...pulumi.InvokeOption) LookupIotHubResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotHubResourceResult, error) {
			args := v.(LookupIotHubResourceArgs)
			r, err := LookupIotHubResource(ctx, &args, opts...)
			var s LookupIotHubResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotHubResourceResultOutput)
}

type LookupIotHubResourceOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupIotHubResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubResourceArgs)(nil)).Elem()
}


type LookupIotHubResourceResultOutput struct{ *pulumi.OutputState }

func (LookupIotHubResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubResourceResult)(nil)).Elem()
}

func (o LookupIotHubResourceResultOutput) ToLookupIotHubResourceResultOutput() LookupIotHubResourceResultOutput {
	return o
}

func (o LookupIotHubResourceResultOutput) ToLookupIotHubResourceResultOutputWithContext(ctx context.Context) LookupIotHubResourceResultOutput {
	return o
}

func (o LookupIotHubResourceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupIotHubResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIotHubResourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupIotHubResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIotHubResourceResultOutput) Properties() IotHubPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) IotHubPropertiesResponse { return v.Properties }).(IotHubPropertiesResponseOutput)
}

func (o LookupIotHubResourceResultOutput) Resourcegroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) string { return v.Resourcegroup }).(pulumi.StringOutput)
}

func (o LookupIotHubResourceResultOutput) Sku() IotHubSkuInfoResponseOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) IotHubSkuInfoResponse { return v.Sku }).(IotHubSkuInfoResponseOutput)
}

func (o LookupIotHubResourceResultOutput) Subscriptionid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) string { return v.Subscriptionid }).(pulumi.StringOutput)
}

func (o LookupIotHubResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIotHubResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotHubResourceResultOutput{})
}
