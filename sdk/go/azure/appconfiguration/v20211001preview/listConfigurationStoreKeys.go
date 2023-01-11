


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConfigurationStoreKeys(ctx *pulumi.Context, args *ListConfigurationStoreKeysArgs, opts ...pulumi.InvokeOption) (*ListConfigurationStoreKeysResult, error) {
	var rv ListConfigurationStoreKeysResult
	err := ctx.Invoke("azure-native:appconfiguration/v20211001preview:listConfigurationStoreKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConfigurationStoreKeysArgs struct {
	ConfigStoreName   string  `pulumi:"configStoreName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SkipToken         *string `pulumi:"skipToken"`
}


type ListConfigurationStoreKeysResult struct {
	NextLink *string          `pulumi:"nextLink"`
	Value    []ApiKeyResponse `pulumi:"value"`
}

func ListConfigurationStoreKeysOutput(ctx *pulumi.Context, args ListConfigurationStoreKeysOutputArgs, opts ...pulumi.InvokeOption) ListConfigurationStoreKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConfigurationStoreKeysResult, error) {
			args := v.(ListConfigurationStoreKeysArgs)
			r, err := ListConfigurationStoreKeys(ctx, &args, opts...)
			var s ListConfigurationStoreKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConfigurationStoreKeysResultOutput)
}

type ListConfigurationStoreKeysOutputArgs struct {
	ConfigStoreName   pulumi.StringInput    `pulumi:"configStoreName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	SkipToken         pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListConfigurationStoreKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigurationStoreKeysArgs)(nil)).Elem()
}


type ListConfigurationStoreKeysResultOutput struct{ *pulumi.OutputState }

func (ListConfigurationStoreKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigurationStoreKeysResult)(nil)).Elem()
}

func (o ListConfigurationStoreKeysResultOutput) ToListConfigurationStoreKeysResultOutput() ListConfigurationStoreKeysResultOutput {
	return o
}

func (o ListConfigurationStoreKeysResultOutput) ToListConfigurationStoreKeysResultOutputWithContext(ctx context.Context) ListConfigurationStoreKeysResultOutput {
	return o
}

func (o ListConfigurationStoreKeysResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeysResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListConfigurationStoreKeysResultOutput) Value() ApiKeyResponseArrayOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeysResult) []ApiKeyResponse { return v.Value }).(ApiKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConfigurationStoreKeysResultOutput{})
}
