


package v20210515

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMongoDBResourceMongoDBCollection(ctx *pulumi.Context, args *LookupMongoDBResourceMongoDBCollectionArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoDBCollectionResult, error) {
	var rv LookupMongoDBResourceMongoDBCollectionResult
	err := ctx.Invoke("azure-native:documentdb/v20210515:getMongoDBResourceMongoDBCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoDBCollectionArgs struct {
	AccountName       string `pulumi:"accountName"`
	CollectionName    string `pulumi:"collectionName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMongoDBResourceMongoDBCollectionResult struct {
	Id       string                                          `pulumi:"id"`
	Location *string                                         `pulumi:"location"`
	Name     string                                          `pulumi:"name"`
	Options  *MongoDBCollectionGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *MongoDBCollectionGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                               `pulumi:"tags"`
	Type     string                                          `pulumi:"type"`
}
