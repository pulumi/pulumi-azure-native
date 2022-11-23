


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCassandraDataCenter(ctx *pulumi.Context, args *LookupCassandraDataCenterArgs, opts ...pulumi.InvokeOption) (*LookupCassandraDataCenterResult, error) {
	var rv LookupCassandraDataCenterResult
	err := ctx.Invoke("azure-native:documentdb/v20210701preview:getCassandraDataCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraDataCenterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DataCenterName    string `pulumi:"dataCenterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCassandraDataCenterResult struct {
	Id         string                               `pulumi:"id"`
	Name       string                               `pulumi:"name"`
	Properties DataCenterResourceResponseProperties `pulumi:"properties"`
	Type       string                               `pulumi:"type"`
}

func LookupCassandraDataCenterOutput(ctx *pulumi.Context, args LookupCassandraDataCenterOutputArgs, opts ...pulumi.InvokeOption) LookupCassandraDataCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCassandraDataCenterResult, error) {
			args := v.(LookupCassandraDataCenterArgs)
			r, err := LookupCassandraDataCenter(ctx, &args, opts...)
			var s LookupCassandraDataCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCassandraDataCenterResultOutput)
}

type LookupCassandraDataCenterOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	DataCenterName    pulumi.StringInput `pulumi:"dataCenterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCassandraDataCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraDataCenterArgs)(nil)).Elem()
}


type LookupCassandraDataCenterResultOutput struct{ *pulumi.OutputState }

func (LookupCassandraDataCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraDataCenterResult)(nil)).Elem()
}

func (o LookupCassandraDataCenterResultOutput) ToLookupCassandraDataCenterResultOutput() LookupCassandraDataCenterResultOutput {
	return o
}

func (o LookupCassandraDataCenterResultOutput) ToLookupCassandraDataCenterResultOutputWithContext(ctx context.Context) LookupCassandraDataCenterResultOutput {
	return o
}

func (o LookupCassandraDataCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraDataCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCassandraDataCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraDataCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCassandraDataCenterResultOutput) Properties() DataCenterResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupCassandraDataCenterResult) DataCenterResourceResponseProperties { return v.Properties }).(DataCenterResourceResponsePropertiesOutput)
}

func (o LookupCassandraDataCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraDataCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCassandraDataCenterResultOutput{})
}
