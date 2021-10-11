


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:databoxedge/v20210601:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountArgs struct {
	DeviceName         string `pulumi:"deviceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageAccountName string `pulumi:"storageAccountName"`
}


type LookupStorageAccountResult struct {
	BlobEndpoint               string             `pulumi:"blobEndpoint"`
	ContainerCount             int                `pulumi:"containerCount"`
	DataPolicy                 string             `pulumi:"dataPolicy"`
	Description                *string            `pulumi:"description"`
	Id                         string             `pulumi:"id"`
	Name                       string             `pulumi:"name"`
	StorageAccountCredentialId *string            `pulumi:"storageAccountCredentialId"`
	StorageAccountStatus       *string            `pulumi:"storageAccountStatus"`
	SystemData                 SystemDataResponse `pulumi:"systemData"`
	Type                       string             `pulumi:"type"`
}
