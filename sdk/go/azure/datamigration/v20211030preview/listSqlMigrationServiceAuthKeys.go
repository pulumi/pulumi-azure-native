


package v20211030preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSqlMigrationServiceAuthKeys(ctx *pulumi.Context, args *ListSqlMigrationServiceAuthKeysArgs, opts ...pulumi.InvokeOption) (*ListSqlMigrationServiceAuthKeysResult, error) {
	var rv ListSqlMigrationServiceAuthKeysResult
	err := ctx.Invoke("azure-native:datamigration/v20211030preview:listSqlMigrationServiceAuthKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSqlMigrationServiceAuthKeysArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SqlMigrationServiceName string `pulumi:"sqlMigrationServiceName"`
}


type ListSqlMigrationServiceAuthKeysResult struct {
	AuthKey1 *string `pulumi:"authKey1"`
	AuthKey2 *string `pulumi:"authKey2"`
}
