


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedDatabase(ctx *pulumi.Context, args *LookupManagedDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupManagedDatabaseResult, error) {
	var rv LookupManagedDatabaseResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getManagedDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedDatabaseArgs struct {
	DatabaseName        string `pulumi:"databaseName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedDatabaseResult struct {
	CatalogCollation         *string           `pulumi:"catalogCollation"`
	Collation                *string           `pulumi:"collation"`
	CreationDate             string            `pulumi:"creationDate"`
	DefaultSecondaryLocation string            `pulumi:"defaultSecondaryLocation"`
	EarliestRestorePoint     string            `pulumi:"earliestRestorePoint"`
	FailoverGroupId          string            `pulumi:"failoverGroupId"`
	Id                       string            `pulumi:"id"`
	Location                 string            `pulumi:"location"`
	Name                     string            `pulumi:"name"`
	Status                   string            `pulumi:"status"`
	Tags                     map[string]string `pulumi:"tags"`
	Type                     string            `pulumi:"type"`
}
