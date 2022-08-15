


package v20150408

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabaseAccountKeys(ctx *pulumi.Context, args *ListDatabaseAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListDatabaseAccountKeysResult, error) {
	var rv ListDatabaseAccountKeysResult
	err := ctx.Invoke("azure-native:documentdb/v20150408:listDatabaseAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDatabaseAccountKeysResult struct {
	PrimaryMasterKey           string `pulumi:"primaryMasterKey"`
	PrimaryReadonlyMasterKey   string `pulumi:"primaryReadonlyMasterKey"`
	SecondaryMasterKey         string `pulumi:"secondaryMasterKey"`
	SecondaryReadonlyMasterKey string `pulumi:"secondaryReadonlyMasterKey"`
}

func ListDatabaseAccountKeysOutput(ctx *pulumi.Context, args ListDatabaseAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListDatabaseAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatabaseAccountKeysResult, error) {
			args := v.(ListDatabaseAccountKeysArgs)
			r, err := ListDatabaseAccountKeys(ctx, &args, opts...)
			var s ListDatabaseAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDatabaseAccountKeysResultOutput)
}

type ListDatabaseAccountKeysOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDatabaseAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseAccountKeysArgs)(nil)).Elem()
}


type ListDatabaseAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListDatabaseAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseAccountKeysResult)(nil)).Elem()
}

func (o ListDatabaseAccountKeysResultOutput) ToListDatabaseAccountKeysResultOutput() ListDatabaseAccountKeysResultOutput {
	return o
}

func (o ListDatabaseAccountKeysResultOutput) ToListDatabaseAccountKeysResultOutputWithContext(ctx context.Context) ListDatabaseAccountKeysResultOutput {
	return o
}

func (o ListDatabaseAccountKeysResultOutput) PrimaryMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseAccountKeysResult) string { return v.PrimaryMasterKey }).(pulumi.StringOutput)
}

func (o ListDatabaseAccountKeysResultOutput) PrimaryReadonlyMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseAccountKeysResult) string { return v.PrimaryReadonlyMasterKey }).(pulumi.StringOutput)
}

func (o ListDatabaseAccountKeysResultOutput) SecondaryMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseAccountKeysResult) string { return v.SecondaryMasterKey }).(pulumi.StringOutput)
}

func (o ListDatabaseAccountKeysResultOutput) SecondaryReadonlyMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseAccountKeysResult) string { return v.SecondaryReadonlyMasterKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabaseAccountKeysResultOutput{})
}
