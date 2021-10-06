


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlResourceSqlStoredProcedure(ctx *pulumi.Context, args *LookupSqlResourceSqlStoredProcedureArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlStoredProcedureResult, error) {
	var rv LookupSqlResourceSqlStoredProcedureResult
	err := ctx.Invoke("azure-native:documentdb/v20210301preview:getSqlResourceSqlStoredProcedure", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlStoredProcedureArgs struct {
	AccountName         string `pulumi:"accountName"`
	ContainerName       string `pulumi:"containerName"`
	DatabaseName        string `pulumi:"databaseName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	StoredProcedureName string `pulumi:"storedProcedureName"`
}


type LookupSqlResourceSqlStoredProcedureResult struct {
	Id       string                                           `pulumi:"id"`
	Identity *ManagedServiceIdentityResponse                  `pulumi:"identity"`
	Location *string                                          `pulumi:"location"`
	Name     string                                           `pulumi:"name"`
	Resource *SqlStoredProcedureGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                                `pulumi:"tags"`
	Type     string                                           `pulumi:"type"`
}
