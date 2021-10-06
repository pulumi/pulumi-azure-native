


package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlPool(ctx *pulumi.Context, args *LookupSqlPoolArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolResult, error) {
	var rv LookupSqlPoolResult
	err := ctx.Invoke("azure-native:synapse/v20190601preview:getSqlPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SqlPoolName       string `pulumi:"sqlPoolName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSqlPoolResult struct {
	Collation             *string           `pulumi:"collation"`
	CreateMode            *string           `pulumi:"createMode"`
	CreationDate          *string           `pulumi:"creationDate"`
	Id                    string            `pulumi:"id"`
	Location              string            `pulumi:"location"`
	MaxSizeBytes          *float64          `pulumi:"maxSizeBytes"`
	Name                  string            `pulumi:"name"`
	ProvisioningState     *string           `pulumi:"provisioningState"`
	RecoverableDatabaseId *string           `pulumi:"recoverableDatabaseId"`
	RestorePointInTime    *string           `pulumi:"restorePointInTime"`
	Sku                   *SkuResponse      `pulumi:"sku"`
	SourceDatabaseId      *string           `pulumi:"sourceDatabaseId"`
	Status                *string           `pulumi:"status"`
	Tags                  map[string]string `pulumi:"tags"`
	Type                  string            `pulumi:"type"`
}
