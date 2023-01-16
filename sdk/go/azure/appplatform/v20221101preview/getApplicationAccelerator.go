


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationAccelerator(ctx *pulumi.Context, args *LookupApplicationAcceleratorArgs, opts ...pulumi.InvokeOption) (*LookupApplicationAcceleratorResult, error) {
	var rv LookupApplicationAcceleratorResult
	err := ctx.Invoke("azure-native:appplatform/v20221101preview:getApplicationAccelerator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApplicationAcceleratorArgs struct {
	ApplicationAcceleratorName string `pulumi:"applicationAcceleratorName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	ServiceName                string `pulumi:"serviceName"`
}


type LookupApplicationAcceleratorResult struct {
	Id         string                                   `pulumi:"id"`
	Name       string                                   `pulumi:"name"`
	Properties ApplicationAcceleratorPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                             `pulumi:"sku"`
	SystemData SystemDataResponse                       `pulumi:"systemData"`
	Type       string                                   `pulumi:"type"`
}


func (val *LookupApplicationAcceleratorResult) Defaults() *LookupApplicationAcceleratorResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}

func LookupApplicationAcceleratorOutput(ctx *pulumi.Context, args LookupApplicationAcceleratorOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationAcceleratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationAcceleratorResult, error) {
			args := v.(LookupApplicationAcceleratorArgs)
			r, err := LookupApplicationAccelerator(ctx, &args, opts...)
			var s LookupApplicationAcceleratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationAcceleratorResultOutput)
}

type LookupApplicationAcceleratorOutputArgs struct {
	ApplicationAcceleratorName pulumi.StringInput `pulumi:"applicationAcceleratorName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName                pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApplicationAcceleratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationAcceleratorArgs)(nil)).Elem()
}


type LookupApplicationAcceleratorResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationAcceleratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationAcceleratorResult)(nil)).Elem()
}

func (o LookupApplicationAcceleratorResultOutput) ToLookupApplicationAcceleratorResultOutput() LookupApplicationAcceleratorResultOutput {
	return o
}

func (o LookupApplicationAcceleratorResultOutput) ToLookupApplicationAcceleratorResultOutputWithContext(ctx context.Context) LookupApplicationAcceleratorResultOutput {
	return o
}

func (o LookupApplicationAcceleratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationAcceleratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationAcceleratorResultOutput) Properties() ApplicationAcceleratorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) ApplicationAcceleratorPropertiesResponse {
		return v.Properties
	}).(ApplicationAcceleratorPropertiesResponseOutput)
}

func (o LookupApplicationAcceleratorResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupApplicationAcceleratorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupApplicationAcceleratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationAcceleratorResultOutput{})
}
