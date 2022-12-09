


package v20170907privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:kusto/v20170907privatepreview:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseResult struct {
	Etag                   string                     `pulumi:"etag"`
	HotCachePeriodInDays   *int                       `pulumi:"hotCachePeriodInDays"`
	Id                     string                     `pulumi:"id"`
	Location               string                     `pulumi:"location"`
	Name                   string                     `pulumi:"name"`
	ProvisioningState      string                     `pulumi:"provisioningState"`
	SoftDeletePeriodInDays int                        `pulumi:"softDeletePeriodInDays"`
	Statistics             DatabaseStatisticsResponse `pulumi:"statistics"`
	Tags                   map[string]string          `pulumi:"tags"`
	Type                   string                     `pulumi:"type"`
}

func LookupDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseResult, error) {
			args := v.(LookupDatabaseArgs)
			r, err := LookupDatabase(ctx, &args, opts...)
			var s LookupDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseResultOutput)
}

type LookupDatabaseOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseArgs)(nil)).Elem()
}


type LookupDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutput() LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) HotCachePeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *int { return v.HotCachePeriodInDays }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) SoftDeletePeriodInDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseResult) int { return v.SoftDeletePeriodInDays }).(pulumi.IntOutput)
}

func (o LookupDatabaseResultOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v LookupDatabaseResult) DatabaseStatisticsResponse { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

func (o LookupDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}
