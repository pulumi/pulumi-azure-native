


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAccountKeys(ctx *pulumi.Context, args *ListAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListAccountKeysResult, error) {
	var rv ListAccountKeysResult
	err := ctx.Invoke("azure-native:engagementfabric/v20180901preview:listAccountKeys", args, &rv, opts...)
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
	Value []KeyDescriptionResponse `pulumi:"value"`
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

func (o ListAccountKeysResultOutput) Value() KeyDescriptionResponseArrayOutput {
	return o.ApplyT(func(v ListAccountKeysResult) []KeyDescriptionResponse { return v.Value }).(KeyDescriptionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAccountKeysResultOutput{})
}
