


package v20160331

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAccountGremlinDatabase(ctx *pulumi.Context, args *LookupDatabaseAccountGremlinDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountGremlinDatabaseResult, error) {
	var rv LookupDatabaseAccountGremlinDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20160331:getDatabaseAccountGremlinDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountGremlinDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountGremlinDatabaseResult struct {
	Etag     *string           `pulumi:"etag"`
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Rid      *string           `pulumi:"rid"`
	Tags     map[string]string `pulumi:"tags"`
	Ts       interface{}       `pulumi:"ts"`
	Type     string            `pulumi:"type"`
}
