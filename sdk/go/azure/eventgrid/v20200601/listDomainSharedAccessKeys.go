


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDomainSharedAccessKeys(ctx *pulumi.Context, args *ListDomainSharedAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListDomainSharedAccessKeysResult, error) {
	var rv ListDomainSharedAccessKeysResult
	err := ctx.Invoke("azure-native:eventgrid/v20200601:listDomainSharedAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDomainSharedAccessKeysArgs struct {
	DomainName        string `pulumi:"domainName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDomainSharedAccessKeysResult struct {
	Key1 *string `pulumi:"key1"`
	Key2 *string `pulumi:"key2"`
}

func ListDomainSharedAccessKeysOutput(ctx *pulumi.Context, args ListDomainSharedAccessKeysOutputArgs, opts ...pulumi.InvokeOption) ListDomainSharedAccessKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDomainSharedAccessKeysResult, error) {
			args := v.(ListDomainSharedAccessKeysArgs)
			r, err := ListDomainSharedAccessKeys(ctx, &args, opts...)
			var s ListDomainSharedAccessKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDomainSharedAccessKeysResultOutput)
}

type ListDomainSharedAccessKeysOutputArgs struct {
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDomainSharedAccessKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainSharedAccessKeysArgs)(nil)).Elem()
}


type ListDomainSharedAccessKeysResultOutput struct{ *pulumi.OutputState }

func (ListDomainSharedAccessKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainSharedAccessKeysResult)(nil)).Elem()
}

func (o ListDomainSharedAccessKeysResultOutput) ToListDomainSharedAccessKeysResultOutput() ListDomainSharedAccessKeysResultOutput {
	return o
}

func (o ListDomainSharedAccessKeysResultOutput) ToListDomainSharedAccessKeysResultOutputWithContext(ctx context.Context) ListDomainSharedAccessKeysResultOutput {
	return o
}

func (o ListDomainSharedAccessKeysResultOutput) Key1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListDomainSharedAccessKeysResult) *string { return v.Key1 }).(pulumi.StringPtrOutput)
}

func (o ListDomainSharedAccessKeysResultOutput) Key2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListDomainSharedAccessKeysResult) *string { return v.Key2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDomainSharedAccessKeysResultOutput{})
}
