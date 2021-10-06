


package v20210415

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGremlinResourceGremlinDatabase(ctx *pulumi.Context, args *LookupGremlinResourceGremlinDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupGremlinResourceGremlinDatabaseResult, error) {
	var rv LookupGremlinResourceGremlinDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20210415:getGremlinResourceGremlinDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGremlinResourceGremlinDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupGremlinResourceGremlinDatabaseResult struct {
	Id       string                                        `pulumi:"id"`
	Location *string                                       `pulumi:"location"`
	Name     string                                        `pulumi:"name"`
	Options  *GremlinDatabaseGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *GremlinDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                             `pulumi:"tags"`
	Type     string                                        `pulumi:"type"`
}
