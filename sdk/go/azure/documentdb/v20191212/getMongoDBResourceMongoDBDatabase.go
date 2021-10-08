


package v20191212

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMongoDBResourceMongoDBDatabase(ctx *pulumi.Context, args *LookupMongoDBResourceMongoDBDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoDBDatabaseResult, error) {
	var rv LookupMongoDBResourceMongoDBDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20191212:getMongoDBResourceMongoDBDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoDBDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMongoDBResourceMongoDBDatabaseResult struct {
	Id       string                                        `pulumi:"id"`
	Location *string                                       `pulumi:"location"`
	Name     string                                        `pulumi:"name"`
	Resource *MongoDBDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                             `pulumi:"tags"`
	Type     string                                        `pulumi:"type"`
}
