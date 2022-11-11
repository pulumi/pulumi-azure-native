


package v20191109

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabasePrincipals(ctx *pulumi.Context, args *ListDatabasePrincipalsArgs, opts ...pulumi.InvokeOption) (*ListDatabasePrincipalsResult, error) {
	var rv ListDatabasePrincipalsResult
	err := ctx.Invoke("azure-native:kusto/v20191109:listDatabasePrincipals", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabasePrincipalsArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDatabasePrincipalsResult struct {
	Value []DatabasePrincipalResponse `pulumi:"value"`
}

func ListDatabasePrincipalsOutput(ctx *pulumi.Context, args ListDatabasePrincipalsOutputArgs, opts ...pulumi.InvokeOption) ListDatabasePrincipalsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatabasePrincipalsResult, error) {
			args := v.(ListDatabasePrincipalsArgs)
			r, err := ListDatabasePrincipals(ctx, &args, opts...)
			var s ListDatabasePrincipalsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDatabasePrincipalsResultOutput)
}

type ListDatabasePrincipalsOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDatabasePrincipalsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasePrincipalsArgs)(nil)).Elem()
}


type ListDatabasePrincipalsResultOutput struct{ *pulumi.OutputState }

func (ListDatabasePrincipalsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasePrincipalsResult)(nil)).Elem()
}

func (o ListDatabasePrincipalsResultOutput) ToListDatabasePrincipalsResultOutput() ListDatabasePrincipalsResultOutput {
	return o
}

func (o ListDatabasePrincipalsResultOutput) ToListDatabasePrincipalsResultOutputWithContext(ctx context.Context) ListDatabasePrincipalsResultOutput {
	return o
}

func (o ListDatabasePrincipalsResultOutput) Value() DatabasePrincipalResponseArrayOutput {
	return o.ApplyT(func(v ListDatabasePrincipalsResult) []DatabasePrincipalResponse { return v.Value }).(DatabasePrincipalResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabasePrincipalsResultOutput{})
}
