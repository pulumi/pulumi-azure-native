


package v20221111

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReadOnlyFollowingDatabase(ctx *pulumi.Context, args *LookupReadOnlyFollowingDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupReadOnlyFollowingDatabaseResult, error) {
	var rv LookupReadOnlyFollowingDatabaseResult
	err := ctx.Invoke("azure-native:kusto/v20221111:getReadOnlyFollowingDatabase", args, &rv, opts...)
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
	AttachedDatabaseConfigurationName string                              `pulumi:"attachedDatabaseConfigurationName"`
	DatabaseShareOrigin               string                              `pulumi:"databaseShareOrigin"`
	HotCachePeriod                    *string                             `pulumi:"hotCachePeriod"`
	Id                                string                              `pulumi:"id"`
	Kind                              string                              `pulumi:"kind"`
	LeaderClusterResourceId           string                              `pulumi:"leaderClusterResourceId"`
	Location                          *string                             `pulumi:"location"`
	Name                              string                              `pulumi:"name"`
	OriginalDatabaseName              string                              `pulumi:"originalDatabaseName"`
	PrincipalsModificationKind        string                              `pulumi:"principalsModificationKind"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	SoftDeletePeriod                  string                              `pulumi:"softDeletePeriod"`
	Statistics                        DatabaseStatisticsResponse          `pulumi:"statistics"`
	TableLevelSharingProperties       TableLevelSharingPropertiesResponse `pulumi:"tableLevelSharingProperties"`
	Type                              string                              `pulumi:"type"`
}

func LookupReadOnlyFollowingDatabaseOutput(ctx *pulumi.Context, args LookupReadOnlyFollowingDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupReadOnlyFollowingDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReadOnlyFollowingDatabaseResult, error) {
			args := v.(LookupReadOnlyFollowingDatabaseArgs)
			r, err := LookupReadOnlyFollowingDatabase(ctx, &args, opts...)
			var s LookupReadOnlyFollowingDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReadOnlyFollowingDatabaseResultOutput)
}

type LookupReadOnlyFollowingDatabaseOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReadOnlyFollowingDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReadOnlyFollowingDatabaseArgs)(nil)).Elem()
}


type LookupReadOnlyFollowingDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupReadOnlyFollowingDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReadOnlyFollowingDatabaseResult)(nil)).Elem()
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) ToLookupReadOnlyFollowingDatabaseResultOutput() LookupReadOnlyFollowingDatabaseResultOutput {
	return o
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) ToLookupReadOnlyFollowingDatabaseResultOutputWithContext(ctx context.Context) LookupReadOnlyFollowingDatabaseResultOutput {
	return o
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) AttachedDatabaseConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.AttachedDatabaseConfigurationName }).(pulumi.StringOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) DatabaseShareOrigin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.DatabaseShareOrigin }).(pulumi.StringOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) HotCachePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) *string { return v.HotCachePeriod }).(pulumi.StringPtrOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) LeaderClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.LeaderClusterResourceId }).(pulumi.StringOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) OriginalDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.OriginalDatabaseName }).(pulumi.StringOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) PrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.PrincipalsModificationKind }).(pulumi.StringOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) SoftDeletePeriod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.SoftDeletePeriod }).(pulumi.StringOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) DatabaseStatisticsResponse { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) TableLevelSharingProperties() TableLevelSharingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) TableLevelSharingPropertiesResponse {
		return v.TableLevelSharingProperties
	}).(TableLevelSharingPropertiesResponseOutput)
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReadOnlyFollowingDatabaseResultOutput{})
}
