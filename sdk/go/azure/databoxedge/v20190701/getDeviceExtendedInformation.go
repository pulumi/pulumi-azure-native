


package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDeviceExtendedInformation(ctx *pulumi.Context, args *GetDeviceExtendedInformationArgs, opts ...pulumi.InvokeOption) (*GetDeviceExtendedInformationResult, error) {
	var rv GetDeviceExtendedInformationResult
	err := ctx.Invoke("azure-native:databoxedge/v20190701:getDeviceExtendedInformation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDeviceExtendedInformationArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetDeviceExtendedInformationResult struct {
	EncryptionKey           *string `pulumi:"encryptionKey"`
	EncryptionKeyThumbprint *string `pulumi:"encryptionKeyThumbprint"`
	Id                      string  `pulumi:"id"`
	Name                    string  `pulumi:"name"`
	ResourceKey             string  `pulumi:"resourceKey"`
	Type                    string  `pulumi:"type"`
}
