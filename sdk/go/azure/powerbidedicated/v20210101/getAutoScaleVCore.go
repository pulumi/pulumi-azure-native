


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAutoScaleVCore(ctx *pulumi.Context, args *LookupAutoScaleVCoreArgs, opts ...pulumi.InvokeOption) (*LookupAutoScaleVCoreResult, error) {
	var rv LookupAutoScaleVCoreResult
	err := ctx.Invoke("azure-native:powerbidedicated/v20210101:getAutoScaleVCore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutoScaleVCoreArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VcoreName         string `pulumi:"vcoreName"`
}


type LookupAutoScaleVCoreResult struct {
	CapacityLimit     *int                      `pulumi:"capacityLimit"`
	CapacityObjectId  *string                   `pulumi:"capacityObjectId"`
	Id                string                    `pulumi:"id"`
	Location          string                    `pulumi:"location"`
	Name              string                    `pulumi:"name"`
	ProvisioningState string                    `pulumi:"provisioningState"`
	Sku               AutoScaleVCoreSkuResponse `pulumi:"sku"`
	SystemData        *SystemDataResponse       `pulumi:"systemData"`
	Tags              map[string]string         `pulumi:"tags"`
	Type              string                    `pulumi:"type"`
}

func LookupAutoScaleVCoreOutput(ctx *pulumi.Context, args LookupAutoScaleVCoreOutputArgs, opts ...pulumi.InvokeOption) LookupAutoScaleVCoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutoScaleVCoreResult, error) {
			args := v.(LookupAutoScaleVCoreArgs)
			r, err := LookupAutoScaleVCore(ctx, &args, opts...)
			var s LookupAutoScaleVCoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutoScaleVCoreResultOutput)
}

type LookupAutoScaleVCoreOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VcoreName         pulumi.StringInput `pulumi:"vcoreName"`
}

func (LookupAutoScaleVCoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoScaleVCoreArgs)(nil)).Elem()
}


type LookupAutoScaleVCoreResultOutput struct{ *pulumi.OutputState }

func (LookupAutoScaleVCoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoScaleVCoreResult)(nil)).Elem()
}

func (o LookupAutoScaleVCoreResultOutput) ToLookupAutoScaleVCoreResultOutput() LookupAutoScaleVCoreResultOutput {
	return o
}

func (o LookupAutoScaleVCoreResultOutput) ToLookupAutoScaleVCoreResultOutputWithContext(ctx context.Context) LookupAutoScaleVCoreResultOutput {
	return o
}

func (o LookupAutoScaleVCoreResultOutput) CapacityLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) *int { return v.CapacityLimit }).(pulumi.IntPtrOutput)
}

func (o LookupAutoScaleVCoreResultOutput) CapacityObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) *string { return v.CapacityObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScaleVCoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAutoScaleVCoreResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAutoScaleVCoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAutoScaleVCoreResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAutoScaleVCoreResultOutput) Sku() AutoScaleVCoreSkuResponseOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) AutoScaleVCoreSkuResponse { return v.Sku }).(AutoScaleVCoreSkuResponseOutput)
}

func (o LookupAutoScaleVCoreResultOutput) SystemData() SystemDataResponsePtrOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) *SystemDataResponse { return v.SystemData }).(SystemDataResponsePtrOutput)
}

func (o LookupAutoScaleVCoreResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAutoScaleVCoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutoScaleVCoreResultOutput{})
}
