


package v20210315

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlResourceSqlDatabase(ctx *pulumi.Context, args *LookupSqlResourceSqlDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlDatabaseResult, error) {
	var rv LookupSqlResourceSqlDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20210315:getSqlResourceSqlDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSqlResourceSqlDatabaseResult struct {
	Id       string                                    `pulumi:"id"`
	Location *string                                   `pulumi:"location"`
	Name     string                                    `pulumi:"name"`
	Options  *SqlDatabaseGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *SqlDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                         `pulumi:"tags"`
	Type     string                                    `pulumi:"type"`
}
