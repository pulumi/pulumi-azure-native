


package storsimple

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetManagerDevicePublicEncryptionKey(ctx *pulumi.Context, args *GetManagerDevicePublicEncryptionKeyArgs, opts ...pulumi.InvokeOption) (*GetManagerDevicePublicEncryptionKeyResult, error) {
	var rv GetManagerDevicePublicEncryptionKeyResult
	err := ctx.Invoke("azure-native:storsimple:getManagerDevicePublicEncryptionKey", args, &rv, opts...)
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
