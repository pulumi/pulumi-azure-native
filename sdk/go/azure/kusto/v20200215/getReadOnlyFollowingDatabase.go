


package v20200215

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReadOnlyFollowingDatabase(ctx *pulumi.Context, args *LookupReadOnlyFollowingDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupReadOnlyFollowingDatabaseResult, error) {
	var rv LookupReadOnlyFollowingDatabaseResult
	err := ctx.Invoke("azure-native:kusto/v20200215:getReadOnlyFollowingDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReadOnlyFollowingDatabaseArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupReadOnlyFollowingDatabaseResult struct {
	AttachedDatabaseConfigurationName string                     `pulumi:"attachedDatabaseConfigurationName"`
	HotCachePeriod                    *string                    `pulumi:"hotCachePeriod"`
	Id                                string                     `pulumi:"id"`
	Kind                              string                     `pulumi:"kind"`
	LeaderClusterResourceId           string                     `pulumi:"leaderClusterResourceId"`
	Location                          *string                    `pulumi:"location"`
	Name                              string                     `pulumi:"name"`
	PrincipalsModificationKind        string                     `pulumi:"principalsModificationKind"`
	ProvisioningState                 string                     `pulumi:"provisioningState"`
	SoftDeletePeriod                  string                     `pulumi:"softDeletePeriod"`
	Statistics                        DatabaseStatisticsResponse `pulumi:"statistics"`
	Type                              string                     `pulumi:"type"`
}
