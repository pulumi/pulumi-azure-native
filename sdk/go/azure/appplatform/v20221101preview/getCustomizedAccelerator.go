


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomizedAccelerator(ctx *pulumi.Context, args *LookupCustomizedAcceleratorArgs, opts ...pulumi.InvokeOption) (*LookupCustomizedAcceleratorResult, error) {
	var rv LookupCustomizedAcceleratorResult
	err := ctx.Invoke("azure-native:appplatform/v20221101preview:getCustomizedAccelerator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCustomizedAcceleratorArgs struct {
	ApplicationAcceleratorName string `pulumi:"applicationAcceleratorName"`
	CustomizedAcceleratorName  string `pulumi:"customizedAcceleratorName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	ServiceName                string `pulumi:"serviceName"`
}


type LookupCustomizedAcceleratorResult struct {
	Id         string                                  `pulumi:"id"`
	Name       string                                  `pulumi:"name"`
	Properties CustomizedAcceleratorPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                            `pulumi:"sku"`
	SystemData SystemDataResponse                      `pulumi:"systemData"`
	Type       string                                  `pulumi:"type"`
}


func (val *LookupCustomizedAcceleratorResult) Defaults() *LookupCustomizedAcceleratorResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}

func LookupCustomizedAcceleratorOutput(ctx *pulumi.Context, args LookupCustomizedAcceleratorOutputArgs, opts ...pulumi.InvokeOption) LookupCustomizedAcceleratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomizedAcceleratorResult, error) {
			args := v.(LookupCustomizedAcceleratorArgs)
			r, err := LookupCustomizedAccelerator(ctx, &args, opts...)
			var s LookupCustomizedAcceleratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomizedAcceleratorResultOutput)
}

type LookupCustomizedAcceleratorOutputArgs struct {
	ApplicationAcceleratorName pulumi.StringInput `pulumi:"applicationAcceleratorName"`
	CustomizedAcceleratorName  pulumi.StringInput `pulumi:"customizedAcceleratorName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName                pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCustomizedAcceleratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomizedAcceleratorArgs)(nil)).Elem()
}


type LookupCustomizedAcceleratorResultOutput struct{ *pulumi.OutputState }

func (LookupCustomizedAcceleratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomizedAcceleratorResult)(nil)).Elem()
}

func (o LookupCustomizedAcceleratorResultOutput) ToLookupCustomizedAcceleratorResultOutput() LookupCustomizedAcceleratorResultOutput {
	return o
}

func (o LookupCustomizedAcceleratorResultOutput) ToLookupCustomizedAcceleratorResultOutputWithContext(ctx context.Context) LookupCustomizedAcceleratorResultOutput {
	return o
}

func (o LookupCustomizedAcceleratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomizedAcceleratorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomizedAcceleratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomizedAcceleratorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomizedAcceleratorResultOutput) Properties() CustomizedAcceleratorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCustomizedAcceleratorResult) CustomizedAcceleratorPropertiesResponse { return v.Properties }).(CustomizedAcceleratorPropertiesResponseOutput)
}

func (o LookupCustomizedAcceleratorResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomizedAcceleratorResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupCustomizedAcceleratorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomizedAcceleratorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCustomizedAcceleratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomizedAcceleratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomizedAcceleratorResultOutput{})
}
