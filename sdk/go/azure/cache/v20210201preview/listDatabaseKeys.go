


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabaseKeys(ctx *pulumi.Context, args *ListDatabaseKeysArgs, opts ...pulumi.InvokeOption) (*ListDatabaseKeysResult, error) {
	var rv ListDatabaseKeysResult
	err := ctx.Invoke("azure-native:cache/v20210201preview:listDatabaseKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseKeysArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDatabaseKeysResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListDatabaseKeysOutput(ctx *pulumi.Context, args ListDatabaseKeysOutputArgs, opts ...pulumi.InvokeOption) ListDatabaseKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatabaseKeysResult, error) {
			args := v.(ListDatabaseKeysArgs)
			r, err := ListDatabaseKeys(ctx, &args, opts...)
			var s ListDatabaseKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDatabaseKeysResultOutput)
}

type ListDatabaseKeysOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDatabaseKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseKeysArgs)(nil)).Elem()
}


type ListDatabaseKeysResultOutput struct{ *pulumi.OutputState }

func (ListDatabaseKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseKeysResult)(nil)).Elem()
}

func (o ListDatabaseKeysResultOutput) ToListDatabaseKeysResultOutput() ListDatabaseKeysResultOutput {
	return o
}

func (o ListDatabaseKeysResultOutput) ToListDatabaseKeysResultOutputWithContext(ctx context.Context) ListDatabaseKeysResultOutput {
	return o
}

func (o ListDatabaseKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o ListDatabaseKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabaseKeysResultOutput{})
}
