


package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSignalRKeys(ctx *pulumi.Context, args *ListSignalRKeysArgs, opts ...pulumi.InvokeOption) (*ListSignalRKeysResult, error) {
	var rv ListSignalRKeysResult
	err := ctx.Invoke("azure-native:signalrservice/v20200501:listSignalRKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSignalRKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListSignalRKeysResult struct {
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}
