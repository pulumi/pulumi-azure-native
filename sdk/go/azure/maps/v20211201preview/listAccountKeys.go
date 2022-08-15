


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAccountKeys(ctx *pulumi.Context, args *ListAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListAccountKeysResult, error) {
	var rv ListAccountKeysResult
	err := ctx.Invoke("azure-native:maps/v20211201preview:listAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAccountKeysResult struct {
	PrimaryKey              string `pulumi:"primaryKey"`
	PrimaryKeyLastUpdated   string `pulumi:"primaryKeyLastUpdated"`
	SecondaryKey            string `pulumi:"secondaryKey"`
	SecondaryKeyLastUpdated string `pulumi:"secondaryKeyLastUpdated"`
}

func ListAccountKeysOutput(ctx *pulumi.Context, args ListAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAccountKeysResult, error) {
			args := v.(ListAccountKeysArgs)
			r, err := ListAccountKeys(ctx, &args, opts...)
			var s ListAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAccountKeysResultOutput)
}

type ListAccountKeysOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountKeysArgs)(nil)).Elem()
}


type ListAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountKeysResult)(nil)).Elem()
}

func (o ListAccountKeysResultOutput) ToListAccountKeysResultOutput() ListAccountKeysResultOutput {
	return o
}

func (o ListAccountKeysResultOutput) ToListAccountKeysResultOutputWithContext(ctx context.Context) ListAccountKeysResultOutput {
	return o
}

func (o ListAccountKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o ListAccountKeysResultOutput) PrimaryKeyLastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountKeysResult) string { return v.PrimaryKeyLastUpdated }).(pulumi.StringOutput)
}

func (o ListAccountKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func (o ListAccountKeysResultOutput) SecondaryKeyLastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountKeysResult) string { return v.SecondaryKeyLastUpdated }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAccountKeysResultOutput{})
}
