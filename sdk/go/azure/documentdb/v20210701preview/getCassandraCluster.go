


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCassandraCluster(ctx *pulumi.Context, args *LookupCassandraClusterArgs, opts ...pulumi.InvokeOption) (*LookupCassandraClusterResult, error) {
	var rv LookupCassandraClusterResult
	err := ctx.Invoke("azure-native:documentdb/v20210701preview:getCassandraCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCassandraClusterResult struct {
	Id         string                            `pulumi:"id"`
	Identity   *ManagedServiceIdentityResponse   `pulumi:"identity"`
	Location   *string                           `pulumi:"location"`
	Name       string                            `pulumi:"name"`
	Properties ClusterResourceResponseProperties `pulumi:"properties"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       string                            `pulumi:"type"`
}

func LookupCassandraClusterOutput(ctx *pulumi.Context, args LookupCassandraClusterOutputArgs, opts ...pulumi.InvokeOption) LookupCassandraClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCassandraClusterResult, error) {
			args := v.(LookupCassandraClusterArgs)
			r, err := LookupCassandraCluster(ctx, &args, opts...)
			var s LookupCassandraClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCassandraClusterResultOutput)
}

type LookupCassandraClusterOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCassandraClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraClusterArgs)(nil)).Elem()
}


type LookupCassandraClusterResultOutput struct{ *pulumi.OutputState }

func (LookupCassandraClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraClusterResult)(nil)).Elem()
}

func (o LookupCassandraClusterResultOutput) ToLookupCassandraClusterResultOutput() LookupCassandraClusterResultOutput {
	return o
}

func (o LookupCassandraClusterResultOutput) ToLookupCassandraClusterResultOutputWithContext(ctx context.Context) LookupCassandraClusterResultOutput {
	return o
}

func (o LookupCassandraClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCassandraClusterResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupCassandraClusterResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupCassandraClusterResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCassandraClusterResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCassandraClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCassandraClusterResultOutput) Properties() ClusterResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupCassandraClusterResult) ClusterResourceResponseProperties { return v.Properties }).(ClusterResourceResponsePropertiesOutput)
}

func (o LookupCassandraClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCassandraClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCassandraClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCassandraClusterResultOutput{})
}
