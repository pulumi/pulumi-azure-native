


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReadWriteDatabase(ctx *pulumi.Context, args *LookupReadWriteDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupReadWriteDatabaseResult, error) {
	var rv LookupReadWriteDatabaseResult
	err := ctx.Invoke("azure-native:synapse/v20210401preview:getReadWriteDatabase", args, &rv, opts...)
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

func LookupReadWriteDatabaseOutput(ctx *pulumi.Context, args LookupReadWriteDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupReadWriteDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReadWriteDatabaseResult, error) {
			args := v.(LookupReadWriteDatabaseArgs)
			r, err := LookupReadWriteDatabase(ctx, &args, opts...)
			var s LookupReadWriteDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReadWriteDatabaseResultOutput)
}

type LookupReadWriteDatabaseOutputArgs struct {
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	KustoPoolName     pulumi.StringInput `pulumi:"kustoPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupReadWriteDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReadWriteDatabaseArgs)(nil)).Elem()
}


type LookupReadWriteDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupReadWriteDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReadWriteDatabaseResult)(nil)).Elem()
}

func (o LookupReadWriteDatabaseResultOutput) ToLookupReadWriteDatabaseResultOutput() LookupReadWriteDatabaseResultOutput {
	return o
}

func (o LookupReadWriteDatabaseResultOutput) ToLookupReadWriteDatabaseResultOutputWithContext(ctx context.Context) LookupReadWriteDatabaseResultOutput {
	return o
}

func (o LookupReadWriteDatabaseResultOutput) HotCachePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) *string { return v.HotCachePeriod }).(pulumi.StringPtrOutput)
}

func (o LookupReadWriteDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReadWriteDatabaseResultOutput) IsFollowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) bool { return v.IsFollowed }).(pulumi.BoolOutput)
}

func (o LookupReadWriteDatabaseResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupReadWriteDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReadWriteDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReadWriteDatabaseResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupReadWriteDatabaseResultOutput) SoftDeletePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) *string { return v.SoftDeletePeriod }).(pulumi.StringPtrOutput)
}

func (o LookupReadWriteDatabaseResultOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) DatabaseStatisticsResponse { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

func (o LookupReadWriteDatabaseResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupReadWriteDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReadWriteDatabaseResultOutput{})
}
