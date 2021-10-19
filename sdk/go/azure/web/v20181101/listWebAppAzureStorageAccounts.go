


package v20181101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppAzureStorageAccounts(ctx *pulumi.Context, args *ListWebAppAzureStorageAccountsArgs, opts ...pulumi.InvokeOption) (*ListWebAppAzureStorageAccountsResult, error) {
	var rv ListWebAppAzureStorageAccountsResult
	err := ctx.Invoke("azure-native:web/v20181101:listWebAppAzureStorageAccounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppAzureStorageAccountsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppAzureStorageAccountsResult struct {
	Id         string                                   `pulumi:"id"`
	Kind       *string                                  `pulumi:"kind"`
	Name       string                                   `pulumi:"name"`
	Properties map[string]AzureStorageInfoValueResponse `pulumi:"properties"`
	Type       string                                   `pulumi:"type"`
}
