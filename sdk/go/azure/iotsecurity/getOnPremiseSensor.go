


package iotsecurity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOnPremiseSensor(ctx *pulumi.Context, args *LookupOnPremiseSensorArgs, opts ...pulumi.InvokeOption) (*LookupOnPremiseSensorResult, error) {
	var rv LookupOnPremiseSensorResult
	err := ctx.Invoke("azure-native:iotsecurity:getOnPremiseSensor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOnPremiseSensorArgs struct {
	OnPremiseSensorName string `pulumi:"onPremiseSensorName"`
}


type LookupOnPremiseSensorResult struct {
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupOnPremiseSensorOutput(ctx *pulumi.Context, args LookupOnPremiseSensorOutputArgs, opts ...pulumi.InvokeOption) LookupOnPremiseSensorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOnPremiseSensorResult, error) {
			args := v.(LookupOnPremiseSensorArgs)
			r, err := LookupOnPremiseSensor(ctx, &args, opts...)
			var s LookupOnPremiseSensorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOnPremiseSensorResultOutput)
}

type LookupOnPremiseSensorOutputArgs struct {
	OnPremiseSensorName pulumi.StringInput `pulumi:"onPremiseSensorName"`
}

func (LookupOnPremiseSensorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnPremiseSensorArgs)(nil)).Elem()
}


type LookupOnPremiseSensorResultOutput struct{ *pulumi.OutputState }

func (LookupOnPremiseSensorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnPremiseSensorResult)(nil)).Elem()
}

func (o LookupOnPremiseSensorResultOutput) ToLookupOnPremiseSensorResultOutput() LookupOnPremiseSensorResultOutput {
	return o
}

func (o LookupOnPremiseSensorResultOutput) ToLookupOnPremiseSensorResultOutputWithContext(ctx context.Context) LookupOnPremiseSensorResultOutput {
	return o
}

func (o LookupOnPremiseSensorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnPremiseSensorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOnPremiseSensorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnPremiseSensorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOnPremiseSensorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOnPremiseSensorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOnPremiseSensorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnPremiseSensorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOnPremiseSensorResultOutput{})
}
