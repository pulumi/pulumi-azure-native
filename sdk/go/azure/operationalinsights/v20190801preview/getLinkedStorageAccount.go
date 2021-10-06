


package v20190801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedStorageAccount(ctx *pulumi.Context, args *LookupLinkedStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupLinkedStorageAccountResult, error) {
	var rv LookupLinkedStorageAccountResult
	err := ctx.Invoke("azure-native:operationalinsights/v20190801preview:getLinkedStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedStorageAccountArgs struct {
	DataSourceType    string `pulumi:"dataSourceType"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupLinkedStorageAccountResult struct {
	DataSourceType    string   `pulumi:"dataSourceType"`
	Id                string   `pulumi:"id"`
	Name              string   `pulumi:"name"`
	StorageAccountIds []string `pulumi:"storageAccountIds"`
	Type              string   `pulumi:"type"`
}
