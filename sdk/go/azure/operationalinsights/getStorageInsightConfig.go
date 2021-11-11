


package operationalinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageInsightConfig(ctx *pulumi.Context, args *LookupStorageInsightConfigArgs, opts ...pulumi.InvokeOption) (*LookupStorageInsightConfigResult, error) {
	var rv LookupStorageInsightConfigResult
	err := ctx.Invoke("azure-native:operationalinsights:getStorageInsightConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageInsightConfigArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageInsightName string `pulumi:"storageInsightName"`
	WorkspaceName      string `pulumi:"workspaceName"`
}


type LookupStorageInsightConfigResult struct {
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
