


package storsimple

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDeviceFailoverTars(ctx *pulumi.Context, args *ListDeviceFailoverTarsArgs, opts ...pulumi.InvokeOption) (*ListDeviceFailoverTarsResult, error) {
	var rv ListDeviceFailoverTarsResult
	err := ctx.Invoke("azure-native:storsimple:listDeviceFailoverTars", args, &rv, opts...)
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
