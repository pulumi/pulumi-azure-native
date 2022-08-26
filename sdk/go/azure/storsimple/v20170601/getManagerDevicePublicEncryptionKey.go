


package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetManagerDevicePublicEncryptionKey(ctx *pulumi.Context, args *GetManagerDevicePublicEncryptionKeyArgs, opts ...pulumi.InvokeOption) (*GetManagerDevicePublicEncryptionKeyResult, error) {
	var rv GetManagerDevicePublicEncryptionKeyResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:getManagerDevicePublicEncryptionKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetManagerDevicePublicEncryptionKeyArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetManagerDevicePublicEncryptionKeyResult struct {
	Key string `pulumi:"key"`
}

func GetManagerDevicePublicEncryptionKeyOutput(ctx *pulumi.Context, args GetManagerDevicePublicEncryptionKeyOutputArgs, opts ...pulumi.InvokeOption) GetManagerDevicePublicEncryptionKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetManagerDevicePublicEncryptionKeyResult, error) {
			args := v.(GetManagerDevicePublicEncryptionKeyArgs)
			r, err := GetManagerDevicePublicEncryptionKey(ctx, &args, opts...)
			var s GetManagerDevicePublicEncryptionKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetManagerDevicePublicEncryptionKeyResultOutput)
}

type GetManagerDevicePublicEncryptionKeyOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetManagerDevicePublicEncryptionKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetManagerDevicePublicEncryptionKeyArgs)(nil)).Elem()
}


type GetManagerDevicePublicEncryptionKeyResultOutput struct{ *pulumi.OutputState }

func (GetManagerDevicePublicEncryptionKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetManagerDevicePublicEncryptionKeyResult)(nil)).Elem()
}

func (o GetManagerDevicePublicEncryptionKeyResultOutput) ToGetManagerDevicePublicEncryptionKeyResultOutput() GetManagerDevicePublicEncryptionKeyResultOutput {
	return o
}

func (o GetManagerDevicePublicEncryptionKeyResultOutput) ToGetManagerDevicePublicEncryptionKeyResultOutputWithContext(ctx context.Context) GetManagerDevicePublicEncryptionKeyResultOutput {
	return o
}

func (o GetManagerDevicePublicEncryptionKeyResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagerDevicePublicEncryptionKeyResult) string { return v.Key }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetManagerDevicePublicEncryptionKeyResultOutput{})
}
