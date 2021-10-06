


package v20151106

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAccountMongoDBCollection(ctx *pulumi.Context, args *LookupDatabaseAccountMongoDBCollectionArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountMongoDBCollectionResult, error) {
	var rv LookupDatabaseAccountMongoDBCollectionResult
	err := ctx.Invoke("azure-native:documentdb/v20151106:getDatabaseAccountMongoDBCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountMongoDBCollectionArgs struct {
	AccountName       string `pulumi:"accountName"`
	CollectionName    string `pulumi:"collectionName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountMongoDBCollectionResult struct {
	Id       string               `pulumi:"id"`
	Indexes  []MongoIndexResponse `pulumi:"indexes"`
	Location *string              `pulumi:"location"`
	Name     string               `pulumi:"name"`
	ShardKey map[string]string    `pulumi:"shardKey"`
	Tags     map[string]string    `pulumi:"tags"`
	Type     string               `pulumi:"type"`
}
