


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReadWriteDatabase(ctx *pulumi.Context, args *LookupReadWriteDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupReadWriteDatabaseResult, error) {
	var rv LookupReadWriteDatabaseResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getReadWriteDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReadWriteDatabaseArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	KustoPoolName     string `pulumi:"kustoPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupReadWriteDatabaseResult struct {
	HotCachePeriod    *string                    `pulumi:"hotCachePeriod"`
	Id                string                     `pulumi:"id"`
	IsFollowed        bool                       `pulumi:"isFollowed"`
	Kind              string                     `pulumi:"kind"`
	Location          *string                    `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	SoftDeletePeriod  *string                    `pulumi:"softDeletePeriod"`
	Statistics        DatabaseStatisticsResponse `pulumi:"statistics"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Type              string                     `pulumi:"type"`
}
