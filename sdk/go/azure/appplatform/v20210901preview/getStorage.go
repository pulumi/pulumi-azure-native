


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorage(ctx *pulumi.Context, args *LookupStorageArgs, opts ...pulumi.InvokeOption) (*LookupStorageResult, error) {
	var rv LookupStorageResult
	err := ctx.Invoke("azure-native:appplatform/v20210901preview:getStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	StorageName       string `pulumi:"storageName"`
}


type LookupStorageResult struct {
	Id         string                 `pulumi:"id"`
	Name       string                 `pulumi:"name"`
	Properties StorageAccountResponse `pulumi:"properties"`
	SystemData SystemDataResponse     `pulumi:"systemData"`
	Type       string                 `pulumi:"type"`
}
