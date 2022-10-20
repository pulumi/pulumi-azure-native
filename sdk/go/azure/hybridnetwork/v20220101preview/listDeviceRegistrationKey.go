


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDeviceRegistrationKey(ctx *pulumi.Context, args *ListDeviceRegistrationKeyArgs, opts ...pulumi.InvokeOption) (*ListDeviceRegistrationKeyResult, error) {
	var rv ListDeviceRegistrationKeyResult
	err := ctx.Invoke("azure-native:hybridnetwork/v20220101preview:listDeviceRegistrationKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDeviceRegistrationKeyArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDeviceRegistrationKeyResult struct {
	RegistrationKey string `pulumi:"registrationKey"`
}

func ListDeviceRegistrationKeyOutput(ctx *pulumi.Context, args ListDeviceRegistrationKeyOutputArgs, opts ...pulumi.InvokeOption) ListDeviceRegistrationKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDeviceRegistrationKeyResult, error) {
			args := v.(ListDeviceRegistrationKeyArgs)
			r, err := ListDeviceRegistrationKey(ctx, &args, opts...)
			var s ListDeviceRegistrationKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDeviceRegistrationKeyResultOutput)
}

type ListDeviceRegistrationKeyOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDeviceRegistrationKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeviceRegistrationKeyArgs)(nil)).Elem()
}


type ListDeviceRegistrationKeyResultOutput struct{ *pulumi.OutputState }

func (ListDeviceRegistrationKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeviceRegistrationKeyResult)(nil)).Elem()
}

func (o ListDeviceRegistrationKeyResultOutput) ToListDeviceRegistrationKeyResultOutput() ListDeviceRegistrationKeyResultOutput {
	return o
}

func (o ListDeviceRegistrationKeyResultOutput) ToListDeviceRegistrationKeyResultOutputWithContext(ctx context.Context) ListDeviceRegistrationKeyResultOutput {
	return o
}

func (o ListDeviceRegistrationKeyResultOutput) RegistrationKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDeviceRegistrationKeyResult) string { return v.RegistrationKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDeviceRegistrationKeyResultOutput{})
}
