


package v20211030preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSqlMigrationServiceAuthKeys(ctx *pulumi.Context, args *ListSqlMigrationServiceAuthKeysArgs, opts ...pulumi.InvokeOption) (*ListSqlMigrationServiceAuthKeysResult, error) {
	var rv ListSqlMigrationServiceAuthKeysResult
	err := ctx.Invoke("azure-native:datamigration/v20211030preview:listSqlMigrationServiceAuthKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSqlMigrationServiceAuthKeysArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SqlMigrationServiceName string `pulumi:"sqlMigrationServiceName"`
}


type ListSqlMigrationServiceAuthKeysResult struct {
	AuthKey1 *string `pulumi:"authKey1"`
	AuthKey2 *string `pulumi:"authKey2"`
}

func ListSqlMigrationServiceAuthKeysOutput(ctx *pulumi.Context, args ListSqlMigrationServiceAuthKeysOutputArgs, opts ...pulumi.InvokeOption) ListSqlMigrationServiceAuthKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSqlMigrationServiceAuthKeysResult, error) {
			args := v.(ListSqlMigrationServiceAuthKeysArgs)
			r, err := ListSqlMigrationServiceAuthKeys(ctx, &args, opts...)
			var s ListSqlMigrationServiceAuthKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSqlMigrationServiceAuthKeysResultOutput)
}

type ListSqlMigrationServiceAuthKeysOutputArgs struct {
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlMigrationServiceName pulumi.StringInput `pulumi:"sqlMigrationServiceName"`
}

func (ListSqlMigrationServiceAuthKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSqlMigrationServiceAuthKeysArgs)(nil)).Elem()
}


type ListSqlMigrationServiceAuthKeysResultOutput struct{ *pulumi.OutputState }

func (ListSqlMigrationServiceAuthKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSqlMigrationServiceAuthKeysResult)(nil)).Elem()
}

func (o ListSqlMigrationServiceAuthKeysResultOutput) ToListSqlMigrationServiceAuthKeysResultOutput() ListSqlMigrationServiceAuthKeysResultOutput {
	return o
}

func (o ListSqlMigrationServiceAuthKeysResultOutput) ToListSqlMigrationServiceAuthKeysResultOutputWithContext(ctx context.Context) ListSqlMigrationServiceAuthKeysResultOutput {
	return o
}

func (o ListSqlMigrationServiceAuthKeysResultOutput) AuthKey1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSqlMigrationServiceAuthKeysResult) *string { return v.AuthKey1 }).(pulumi.StringPtrOutput)
}

func (o ListSqlMigrationServiceAuthKeysResultOutput) AuthKey2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSqlMigrationServiceAuthKeysResult) *string { return v.AuthKey2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSqlMigrationServiceAuthKeysResultOutput{})
}
