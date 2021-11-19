


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListServiceTestKeys(ctx *pulumi.Context, args *ListServiceTestKeysArgs, opts ...pulumi.InvokeOption) (*ListServiceTestKeysResult, error) {
	var rv ListServiceTestKeysResult
	err := ctx.Invoke("azure-native:appplatform/v20210901preview:listServiceTestKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServiceTestKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListServiceTestKeysResult struct {
	Enabled               *bool   `pulumi:"enabled"`
	PrimaryKey            *string `pulumi:"primaryKey"`
	PrimaryTestEndpoint   *string `pulumi:"primaryTestEndpoint"`
	SecondaryKey          *string `pulumi:"secondaryKey"`
	SecondaryTestEndpoint *string `pulumi:"secondaryTestEndpoint"`
}
