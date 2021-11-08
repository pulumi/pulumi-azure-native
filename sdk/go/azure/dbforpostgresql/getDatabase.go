


package dbforpostgresql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:dbforpostgresql:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupDatabaseResult struct {
	Charset   *string `pulumi:"charset"`
	Collation *string `pulumi:"collation"`
	Id        string  `pulumi:"id"`
	Name      string  `pulumi:"name"`
	Type      string  `pulumi:"type"`
}
