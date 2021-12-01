


package v20160331

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAccountSqlContainer(ctx *pulumi.Context, args *LookupDatabaseAccountSqlContainerArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountSqlContainerResult, error) {
	var rv LookupDatabaseAccountSqlContainerResult
	err := ctx.Invoke("azure-native:documentdb/v20160331:getDatabaseAccountSqlContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountSqlContainerArgs struct {
	AccountName       string `pulumi:"accountName"`
	ContainerName     string `pulumi:"containerName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountSqlContainerResult struct {
	ConflictResolutionPolicy *ConflictResolutionPolicyResponse `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                              `pulumi:"defaultTtl"`
	Etag                     *string                           `pulumi:"etag"`
	Id                       string                            `pulumi:"id"`
	IndexingPolicy           *IndexingPolicyResponse           `pulumi:"indexingPolicy"`
	Location                 *string                           `pulumi:"location"`
	Name                     string                            `pulumi:"name"`
	PartitionKey             *ContainerPartitionKeyResponse    `pulumi:"partitionKey"`
	Rid                      *string                           `pulumi:"rid"`
	Tags                     map[string]string                 `pulumi:"tags"`
	Ts                       interface{}                       `pulumi:"ts"`
	Type                     string                            `pulumi:"type"`
	UniqueKeyPolicy          *UniqueKeyPolicyResponse          `pulumi:"uniqueKeyPolicy"`
}
