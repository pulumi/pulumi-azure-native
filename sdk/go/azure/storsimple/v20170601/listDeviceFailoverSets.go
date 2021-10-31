


package v20170601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDeviceFailoverSets(ctx *pulumi.Context, args *ListDeviceFailoverSetsArgs, opts ...pulumi.InvokeOption) (*ListDeviceFailoverSetsResult, error) {
	var rv ListDeviceFailoverSetsResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:listDeviceFailoverSets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDeviceFailoverSetsArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDeviceFailoverSetsResult struct {
	Value []FailoverSetResponse `pulumi:"value"`
}
