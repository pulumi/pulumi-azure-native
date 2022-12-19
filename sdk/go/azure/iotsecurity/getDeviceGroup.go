


package iotsecurity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeviceGroup(ctx *pulumi.Context, args *LookupDeviceGroupArgs, opts ...pulumi.InvokeOption) (*LookupDeviceGroupResult, error) {
	var rv LookupDeviceGroupResult
	err := ctx.Invoke("azure-native:iotsecurity:getDeviceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceGroupArgs struct {
	DeviceGroupName     string `pulumi:"deviceGroupName"`
	IotDefenderLocation string `pulumi:"iotDefenderLocation"`
}


type LookupDeviceGroupResult struct {
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupDeviceGroupOutput(ctx *pulumi.Context, args LookupDeviceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceGroupResult, error) {
			args := v.(LookupDeviceGroupArgs)
			r, err := LookupDeviceGroup(ctx, &args, opts...)
			var s LookupDeviceGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeviceGroupResultOutput)
}

type LookupDeviceGroupOutputArgs struct {
	DeviceGroupName     pulumi.StringInput `pulumi:"deviceGroupName"`
	IotDefenderLocation pulumi.StringInput `pulumi:"iotDefenderLocation"`
}

func (LookupDeviceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceGroupArgs)(nil)).Elem()
}


type LookupDeviceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceGroupResult)(nil)).Elem()
}

func (o LookupDeviceGroupResultOutput) ToLookupDeviceGroupResultOutput() LookupDeviceGroupResultOutput {
	return o
}

func (o LookupDeviceGroupResultOutput) ToLookupDeviceGroupResultOutputWithContext(ctx context.Context) LookupDeviceGroupResultOutput {
	return o
}

func (o LookupDeviceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeviceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeviceGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDeviceGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDeviceGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceGroupResultOutput{})
}
