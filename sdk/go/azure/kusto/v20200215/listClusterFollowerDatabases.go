


package v20200215

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListClusterFollowerDatabases(ctx *pulumi.Context, args *ListClusterFollowerDatabasesArgs, opts ...pulumi.InvokeOption) (*ListClusterFollowerDatabasesResult, error) {
	var rv ListClusterFollowerDatabasesResult
	err := ctx.Invoke("azure-native:kusto/v20200215:listClusterFollowerDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterFollowerDatabasesArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListClusterFollowerDatabasesResult struct {
	Value []FollowerDatabaseDefinitionResponse `pulumi:"value"`
}

func ListClusterFollowerDatabasesOutput(ctx *pulumi.Context, args ListClusterFollowerDatabasesOutputArgs, opts ...pulumi.InvokeOption) ListClusterFollowerDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListClusterFollowerDatabasesResult, error) {
			args := v.(ListClusterFollowerDatabasesArgs)
			r, err := ListClusterFollowerDatabases(ctx, &args, opts...)
			var s ListClusterFollowerDatabasesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListClusterFollowerDatabasesResultOutput)
}

type ListClusterFollowerDatabasesOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListClusterFollowerDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterFollowerDatabasesArgs)(nil)).Elem()
}


type ListClusterFollowerDatabasesResultOutput struct{ *pulumi.OutputState }

func (ListClusterFollowerDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterFollowerDatabasesResult)(nil)).Elem()
}

func (o ListClusterFollowerDatabasesResultOutput) ToListClusterFollowerDatabasesResultOutput() ListClusterFollowerDatabasesResultOutput {
	return o
}

func (o ListClusterFollowerDatabasesResultOutput) ToListClusterFollowerDatabasesResultOutputWithContext(ctx context.Context) ListClusterFollowerDatabasesResultOutput {
	return o
}

func (o ListClusterFollowerDatabasesResultOutput) Value() FollowerDatabaseDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ListClusterFollowerDatabasesResult) []FollowerDatabaseDefinitionResponse { return v.Value }).(FollowerDatabaseDefinitionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListClusterFollowerDatabasesResultOutput{})
}
