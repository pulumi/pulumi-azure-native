


package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDeviceFailoverTars(ctx *pulumi.Context, args *ListDeviceFailoverTarsArgs, opts ...pulumi.InvokeOption) (*ListDeviceFailoverTarsResult, error) {
	var rv ListDeviceFailoverTarsResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:listDeviceFailoverTars", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDeviceFailoverTarsArgs struct {
	ManagerName       string   `pulumi:"managerName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	SourceDeviceName  string   `pulumi:"sourceDeviceName"`
	VolumeContainers  []string `pulumi:"volumeContainers"`
}


type ListDeviceFailoverTarsResult struct {
	Value []FailoverTargetResponse `pulumi:"value"`
}

func ListDeviceFailoverTarsOutput(ctx *pulumi.Context, args ListDeviceFailoverTarsOutputArgs, opts ...pulumi.InvokeOption) ListDeviceFailoverTarsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDeviceFailoverTarsResult, error) {
			args := v.(ListDeviceFailoverTarsArgs)
			r, err := ListDeviceFailoverTars(ctx, &args, opts...)
			var s ListDeviceFailoverTarsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDeviceFailoverTarsResultOutput)
}

type ListDeviceFailoverTarsOutputArgs struct {
	ManagerName       pulumi.StringInput      `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput      `pulumi:"resourceGroupName"`
	SourceDeviceName  pulumi.StringInput      `pulumi:"sourceDeviceName"`
	VolumeContainers  pulumi.StringArrayInput `pulumi:"volumeContainers"`
}

func (ListDeviceFailoverTarsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeviceFailoverTarsArgs)(nil)).Elem()
}


type ListDeviceFailoverTarsResultOutput struct{ *pulumi.OutputState }

func (ListDeviceFailoverTarsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeviceFailoverTarsResult)(nil)).Elem()
}

func (o ListDeviceFailoverTarsResultOutput) ToListDeviceFailoverTarsResultOutput() ListDeviceFailoverTarsResultOutput {
	return o
}

func (o ListDeviceFailoverTarsResultOutput) ToListDeviceFailoverTarsResultOutputWithContext(ctx context.Context) ListDeviceFailoverTarsResultOutput {
	return o
}

func (o ListDeviceFailoverTarsResultOutput) Value() FailoverTargetResponseArrayOutput {
	return o.ApplyT(func(v ListDeviceFailoverTarsResult) []FailoverTargetResponse { return v.Value }).(FailoverTargetResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDeviceFailoverTarsResultOutput{})
}
