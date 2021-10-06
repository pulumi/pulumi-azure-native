


package v20150320

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageInsight(ctx *pulumi.Context, args *LookupStorageInsightArgs, opts ...pulumi.InvokeOption) (*LookupStorageInsightResult, error) {
	var rv LookupStorageInsightResult
	err := ctx.Invoke("azure-native:operationalinsights/v20150320:getStorageInsight", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageInsightArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageInsightName string `pulumi:"storageInsightName"`
	WorkspaceName      string `pulumi:"workspaceName"`
}


type LookupStorageInsightResult struct {
	Containers     []string                     `pulumi:"containers"`
	ETag           *string                      `pulumi:"eTag"`
	Id             string                       `pulumi:"id"`
	Name           string                       `pulumi:"name"`
	Status         StorageInsightStatusResponse `pulumi:"status"`
	StorageAccount StorageAccountResponse       `pulumi:"storageAccount"`
	Tables         []string                     `pulumi:"tables"`
	Tags           map[string]string            `pulumi:"tags"`
	Type           string                       `pulumi:"type"`
}
