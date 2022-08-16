


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPartnerNamespaceSharedAccessKeys(ctx *pulumi.Context, args *ListPartnerNamespaceSharedAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListPartnerNamespaceSharedAccessKeysResult, error) {
	var rv ListPartnerNamespaceSharedAccessKeysResult
	err := ctx.Invoke("azure-native:eventgrid/v20200401preview:listPartnerNamespaceSharedAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPartnerNamespaceSharedAccessKeysArgs struct {
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type ListPartnerNamespaceSharedAccessKeysResult struct {
	Key1 *string `pulumi:"key1"`
	Key2 *string `pulumi:"key2"`
}

func ListPartnerNamespaceSharedAccessKeysOutput(ctx *pulumi.Context, args ListPartnerNamespaceSharedAccessKeysOutputArgs, opts ...pulumi.InvokeOption) ListPartnerNamespaceSharedAccessKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPartnerNamespaceSharedAccessKeysResult, error) {
			args := v.(ListPartnerNamespaceSharedAccessKeysArgs)
			r, err := ListPartnerNamespaceSharedAccessKeys(ctx, &args, opts...)
			var s ListPartnerNamespaceSharedAccessKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPartnerNamespaceSharedAccessKeysResultOutput)
}

type ListPartnerNamespaceSharedAccessKeysOutputArgs struct {
	PartnerNamespaceName pulumi.StringInput `pulumi:"partnerNamespaceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListPartnerNamespaceSharedAccessKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPartnerNamespaceSharedAccessKeysArgs)(nil)).Elem()
}


type ListPartnerNamespaceSharedAccessKeysResultOutput struct{ *pulumi.OutputState }

func (ListPartnerNamespaceSharedAccessKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPartnerNamespaceSharedAccessKeysResult)(nil)).Elem()
}

func (o ListPartnerNamespaceSharedAccessKeysResultOutput) ToListPartnerNamespaceSharedAccessKeysResultOutput() ListPartnerNamespaceSharedAccessKeysResultOutput {
	return o
}

func (o ListPartnerNamespaceSharedAccessKeysResultOutput) ToListPartnerNamespaceSharedAccessKeysResultOutputWithContext(ctx context.Context) ListPartnerNamespaceSharedAccessKeysResultOutput {
	return o
}

func (o ListPartnerNamespaceSharedAccessKeysResultOutput) Key1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListPartnerNamespaceSharedAccessKeysResult) *string { return v.Key1 }).(pulumi.StringPtrOutput)
}

func (o ListPartnerNamespaceSharedAccessKeysResultOutput) Key2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListPartnerNamespaceSharedAccessKeysResult) *string { return v.Key2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPartnerNamespaceSharedAccessKeysResultOutput{})
}
