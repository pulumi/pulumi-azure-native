


package v20160319

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAccountGremlinGraph(ctx *pulumi.Context, args *LookupDatabaseAccountGremlinGraphArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountGremlinGraphResult, error) {
	var rv LookupDatabaseAccountGremlinGraphResult
	err := ctx.Invoke("azure-native:documentdb/v20160319:getDatabaseAccountGremlinGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDatabaseAccountGremlinGraphArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	GraphName         string `pulumi:"graphName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountGremlinGraphResult struct {
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


func (val *LookupDatabaseAccountGremlinGraphResult) Defaults() *LookupDatabaseAccountGremlinGraphResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}
