


package v20191101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStorageAccountSasTokens(ctx *pulumi.Context, args *ListStorageAccountSasTokensArgs, opts ...pulumi.InvokeOption) (*ListStorageAccountSasTokensResult, error) {
	var rv ListStorageAccountSasTokensResult
	err := ctx.Invoke("azure-native:datalakeanalytics/v20191101preview:listStorageAccountSasTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStorageAccountSasTokensArgs struct {
	AccountName        string `pulumi:"accountName"`
	ContainerName      string `pulumi:"containerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageAccountName string `pulumi:"storageAccountName"`
}


type ListStorageAccountSasTokensResult struct {
	NextLink string                        `pulumi:"nextLink"`
	Value    []SasTokenInformationResponse `pulumi:"value"`
}
