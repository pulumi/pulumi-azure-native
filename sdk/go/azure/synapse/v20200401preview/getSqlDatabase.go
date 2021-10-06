


package v20200401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlDatabase(ctx *pulumi.Context, args *LookupSqlDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupSqlDatabaseResult, error) {
	var rv LookupSqlDatabaseResult
	err := ctx.Invoke("azure-native:synapse/v20200401preview:getSqlDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDatabaseArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SqlDatabaseName   string `pulumi:"sqlDatabaseName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSqlDatabaseResult struct {
	Collation         *string                           `pulumi:"collation"`
	DataRetention     *SqlDatabaseDataRetentionResponse `pulumi:"dataRetention"`
	DatabaseGuid      string                            `pulumi:"databaseGuid"`
	Id                string                            `pulumi:"id"`
	Location          string                            `pulumi:"location"`
	Name              string                            `pulumi:"name"`
	Status            string                            `pulumi:"status"`
	StorageRedundancy *string                           `pulumi:"storageRedundancy"`
	SystemData        SystemDataResponse                `pulumi:"systemData"`
	Tags              map[string]string                 `pulumi:"tags"`
	Type              string                            `pulumi:"type"`
}
