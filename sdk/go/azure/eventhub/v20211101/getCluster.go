


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:eventhub/v20211101:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupClusterResult struct {
	CreatedAt  string              `pulumi:"createdAt"`
	Id         string              `pulumi:"id"`
	Location   *string             `pulumi:"location"`
	MetricId   string              `pulumi:"metricId"`
	Name       string              `pulumi:"name"`
	Sku        *ClusterSkuResponse `pulumi:"sku"`
	Status     string              `pulumi:"status"`
	SystemData SystemDataResponse  `pulumi:"systemData"`
	Tags       map[string]string   `pulumi:"tags"`
	Type       string              `pulumi:"type"`
	UpdatedAt  string              `pulumi:"updatedAt"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}


type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.MetricId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Sku() ClusterSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterSkuResponse { return v.Sku }).(ClusterSkuResponsePtrOutput)
}

func (o LookupClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
