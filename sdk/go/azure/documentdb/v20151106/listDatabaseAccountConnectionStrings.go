


package v20151106

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabaseAccountConnectionStrings(ctx *pulumi.Context, args *ListDatabaseAccountConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListDatabaseAccountConnectionStringsResult, error) {
	var rv ListDatabaseAccountConnectionStringsResult
	err := ctx.Invoke("azure-native:documentdb/v20151106:listDatabaseAccountConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseAccountConnectionStringsArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDatabaseAccountConnectionStringsResult struct {
	ConnectionStrings []DatabaseAccountConnectionStringResponse `pulumi:"connectionStrings"`
}

func ListDatabaseAccountConnectionStringsOutput(ctx *pulumi.Context, args ListDatabaseAccountConnectionStringsOutputArgs, opts ...pulumi.InvokeOption) ListDatabaseAccountConnectionStringsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatabaseAccountConnectionStringsResult, error) {
			args := v.(ListDatabaseAccountConnectionStringsArgs)
			r, err := ListDatabaseAccountConnectionStrings(ctx, &args, opts...)
			var s ListDatabaseAccountConnectionStringsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDatabaseAccountConnectionStringsResultOutput)
}

type ListDatabaseAccountConnectionStringsOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDatabaseAccountConnectionStringsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseAccountConnectionStringsArgs)(nil)).Elem()
}


type ListDatabaseAccountConnectionStringsResultOutput struct{ *pulumi.OutputState }

func (ListDatabaseAccountConnectionStringsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseAccountConnectionStringsResult)(nil)).Elem()
}

func (o ListDatabaseAccountConnectionStringsResultOutput) ToListDatabaseAccountConnectionStringsResultOutput() ListDatabaseAccountConnectionStringsResultOutput {
	return o
}

func (o ListDatabaseAccountConnectionStringsResultOutput) ToListDatabaseAccountConnectionStringsResultOutputWithContext(ctx context.Context) ListDatabaseAccountConnectionStringsResultOutput {
	return o
}

func (o ListDatabaseAccountConnectionStringsResultOutput) ConnectionStrings() DatabaseAccountConnectionStringResponseArrayOutput {
	return o.ApplyT(func(v ListDatabaseAccountConnectionStringsResult) []DatabaseAccountConnectionStringResponse {
		return v.ConnectionStrings
	}).(DatabaseAccountConnectionStringResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabaseAccountConnectionStringsResultOutput{})
}
