


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDeviceRegistrationKey(ctx *pulumi.Context, args *ListDeviceRegistrationKeyArgs, opts ...pulumi.InvokeOption) (*ListDeviceRegistrationKeyResult, error) {
	var rv ListDeviceRegistrationKeyResult
	err := ctx.Invoke("azure-native:hybridnetwork/v20200101preview:listDeviceRegistrationKey", args, &rv, opts...)
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
