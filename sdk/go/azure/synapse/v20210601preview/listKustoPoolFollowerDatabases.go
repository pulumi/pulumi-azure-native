


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListKustoPoolFollowerDatabases(ctx *pulumi.Context, args *ListKustoPoolFollowerDatabasesArgs, opts ...pulumi.InvokeOption) (*ListKustoPoolFollowerDatabasesResult, error) {
	var rv ListKustoPoolFollowerDatabasesResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:listKustoPoolFollowerDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListKustoPoolFollowerDatabasesArgs struct {
	KustoPoolName     string `pulumi:"kustoPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListKustoPoolFollowerDatabasesResult struct {
	Value []FollowerDatabaseDefinitionResponse `pulumi:"value"`
}

func ListKustoPoolFollowerDatabasesOutput(ctx *pulumi.Context, args ListKustoPoolFollowerDatabasesOutputArgs, opts ...pulumi.InvokeOption) ListKustoPoolFollowerDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListKustoPoolFollowerDatabasesResult, error) {
			args := v.(ListKustoPoolFollowerDatabasesArgs)
			r, err := ListKustoPoolFollowerDatabases(ctx, &args, opts...)
			var s ListKustoPoolFollowerDatabasesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListKustoPoolFollowerDatabasesResultOutput)
}

type ListKustoPoolFollowerDatabasesOutputArgs struct {
	KustoPoolName     pulumi.StringInput `pulumi:"kustoPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListKustoPoolFollowerDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKustoPoolFollowerDatabasesArgs)(nil)).Elem()
}


type ListKustoPoolFollowerDatabasesResultOutput struct{ *pulumi.OutputState }

func (ListKustoPoolFollowerDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKustoPoolFollowerDatabasesResult)(nil)).Elem()
}

func (o ListKustoPoolFollowerDatabasesResultOutput) ToListKustoPoolFollowerDatabasesResultOutput() ListKustoPoolFollowerDatabasesResultOutput {
	return o
}

func (o ListKustoPoolFollowerDatabasesResultOutput) ToListKustoPoolFollowerDatabasesResultOutputWithContext(ctx context.Context) ListKustoPoolFollowerDatabasesResultOutput {
	return o
}

func (o ListKustoPoolFollowerDatabasesResultOutput) Value() FollowerDatabaseDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ListKustoPoolFollowerDatabasesResult) []FollowerDatabaseDefinitionResponse { return v.Value }).(FollowerDatabaseDefinitionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListKustoPoolFollowerDatabasesResultOutput{})
}
