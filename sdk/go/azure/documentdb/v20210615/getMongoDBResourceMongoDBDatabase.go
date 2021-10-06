


package v20210615

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMongoDBResourceMongoDBDatabase(ctx *pulumi.Context, args *LookupMongoDBResourceMongoDBDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoDBDatabaseResult, error) {
	var rv LookupMongoDBResourceMongoDBDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20210615:getMongoDBResourceMongoDBDatabase", args, &rv, opts...)
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
	Options  *MongoDBDatabaseGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *MongoDBDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                             `pulumi:"tags"`
	Type     string                                        `pulumi:"type"`
}
