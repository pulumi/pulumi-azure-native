


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSwiftVirtualNetworkConnectionSlot(ctx *pulumi.Context, args *LookupWebAppSwiftVirtualNetworkConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSwiftVirtualNetworkConnectionSlotResult, error) {
	var rv LookupWebAppSwiftVirtualNetworkConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20220301:getWebAppSwiftVirtualNetworkConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSwiftVirtualNetworkConnectionSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppSwiftVirtualNetworkConnectionSlotResult struct {
	Id               string  `pulumi:"id"`
	Kind             *string `pulumi:"kind"`
	Name             string  `pulumi:"name"`
	SubnetResourceId *string `pulumi:"subnetResourceId"`
	SwiftSupported   *bool   `pulumi:"swiftSupported"`
	Type             string  `pulumi:"type"`
}

func LookupWebAppSwiftVirtualNetworkConnectionSlotOutput(ctx *pulumi.Context, args LookupWebAppSwiftVirtualNetworkConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSwiftVirtualNetworkConnectionSlotResult, error) {
			args := v.(LookupWebAppSwiftVirtualNetworkConnectionSlotArgs)
			r, err := LookupWebAppSwiftVirtualNetworkConnectionSlot(ctx, &args, opts...)
			var s LookupWebAppSwiftVirtualNetworkConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput)
}

type LookupWebAppSwiftVirtualNetworkConnectionSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppSwiftVirtualNetworkConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSwiftVirtualNetworkConnectionSlotArgs)(nil)).Elem()
}


type LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSwiftVirtualNetworkConnectionSlotResult)(nil)).Elem()
}

func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) ToLookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput() LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) ToLookupWebAppSwiftVirtualNetworkConnectionSlotResultOutputWithContext(ctx context.Context) LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) *string { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) SwiftSupported() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) *bool { return v.SwiftSupported }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput{})
}
