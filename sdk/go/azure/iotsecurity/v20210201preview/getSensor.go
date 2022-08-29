


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSensor(ctx *pulumi.Context, args *LookupSensorArgs, opts ...pulumi.InvokeOption) (*LookupSensorResult, error) {
	var rv LookupSensorResult
	err := ctx.Invoke("azure-native:iotsecurity/v20210201preview:getSensor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSensorArgs struct {
	Scope      string `pulumi:"scope"`
	SensorName string `pulumi:"sensorName"`
}


type LookupSensorResult struct {
	ConnectivityTime   string             `pulumi:"connectivityTime"`
	DynamicLearning    bool               `pulumi:"dynamicLearning"`
	Id                 string             `pulumi:"id"`
	LearningMode       bool               `pulumi:"learningMode"`
	Name               string             `pulumi:"name"`
	SensorStatus       string             `pulumi:"sensorStatus"`
	SensorType         *string            `pulumi:"sensorType"`
	SensorVersion      string             `pulumi:"sensorVersion"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	TiAutomaticUpdates *bool              `pulumi:"tiAutomaticUpdates"`
	TiStatus           string             `pulumi:"tiStatus"`
	TiVersion          string             `pulumi:"tiVersion"`
	Type               string             `pulumi:"type"`
	Zone               *string            `pulumi:"zone"`
}

func LookupSensorOutput(ctx *pulumi.Context, args LookupSensorOutputArgs, opts ...pulumi.InvokeOption) LookupSensorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSensorResult, error) {
			args := v.(LookupSensorArgs)
			r, err := LookupSensor(ctx, &args, opts...)
			var s LookupSensorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSensorResultOutput)
}

type LookupSensorOutputArgs struct {
	Scope      pulumi.StringInput `pulumi:"scope"`
	SensorName pulumi.StringInput `pulumi:"sensorName"`
}

func (LookupSensorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSensorArgs)(nil)).Elem()
}


type LookupSensorResultOutput struct{ *pulumi.OutputState }

func (LookupSensorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSensorResult)(nil)).Elem()
}

func (o LookupSensorResultOutput) ToLookupSensorResultOutput() LookupSensorResultOutput {
	return o
}

func (o LookupSensorResultOutput) ToLookupSensorResultOutputWithContext(ctx context.Context) LookupSensorResultOutput {
	return o
}

func (o LookupSensorResultOutput) ConnectivityTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensorResult) string { return v.ConnectivityTime }).(pulumi.StringOutput)
}

func (o LookupSensorResultOutput) DynamicLearning() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSensorResult) bool { return v.DynamicLearning }).(pulumi.BoolOutput)
}

func (o LookupSensorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSensorResultOutput) LearningMode() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSensorResult) bool { return v.LearningMode }).(pulumi.BoolOutput)
}

func (o LookupSensorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSensorResultOutput) SensorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensorResult) string { return v.SensorStatus }).(pulumi.StringOutput)
}

func (o LookupSensorResultOutput) SensorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSensorResult) *string { return v.SensorType }).(pulumi.StringPtrOutput)
}

func (o LookupSensorResultOutput) SensorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensorResult) string { return v.SensorVersion }).(pulumi.StringOutput)
}

func (o LookupSensorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSensorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSensorResultOutput) TiAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSensorResult) *bool { return v.TiAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o LookupSensorResultOutput) TiStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensorResult) string { return v.TiStatus }).(pulumi.StringOutput)
}

func (o LookupSensorResultOutput) TiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensorResult) string { return v.TiVersion }).(pulumi.StringOutput)
}

func (o LookupSensorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensorResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSensorResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSensorResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSensorResultOutput{})
}
