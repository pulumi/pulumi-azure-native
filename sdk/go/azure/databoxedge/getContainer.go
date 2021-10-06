


package databoxedge

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainer(ctx *pulumi.Context, args *LookupContainerArgs, opts ...pulumi.InvokeOption) (*LookupContainerResult, error) {
	var rv LookupContainerResult
	err := ctx.Invoke("azure-native:databoxedge:getContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerArgs struct {
	ContainerName      string `pulumi:"containerName"`
	DeviceName         string `pulumi:"deviceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageAccountName string `pulumi:"storageAccountName"`
}


type LookupContainerResult struct {
	ContainerStatus string                 `pulumi:"containerStatus"`
	CreatedDateTime string                 `pulumi:"createdDateTime"`
	DataFormat      string                 `pulumi:"dataFormat"`
	Id              string                 `pulumi:"id"`
	Name            string                 `pulumi:"name"`
	RefreshDetails  RefreshDetailsResponse `pulumi:"refreshDetails"`
	SystemData      SystemDataResponse     `pulumi:"systemData"`
	Type            string                 `pulumi:"type"`
}
