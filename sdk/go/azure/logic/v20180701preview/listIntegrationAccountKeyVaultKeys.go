


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountKeyVaultKeys(ctx *pulumi.Context, args *ListIntegrationAccountKeyVaultKeysArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountKeyVaultKeysResult, error) {
	var rv ListIntegrationAccountKeyVaultKeysResult
	err := ctx.Invoke("azure-native:logic/v20180701preview:listIntegrationAccountKeyVaultKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountKeyVaultKeysArgs struct {
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	KeyVault               KeyVaultReference `pulumi:"keyVault"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	SkipToken              *string           `pulumi:"skipToken"`
}


type ListIntegrationAccountKeyVaultKeysResult struct {
	SkipToken *string               `pulumi:"skipToken"`
	Value     []KeyVaultKeyResponse `pulumi:"value"`
}

func ListIntegrationAccountKeyVaultKeysOutput(ctx *pulumi.Context, args ListIntegrationAccountKeyVaultKeysOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationAccountKeyVaultKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationAccountKeyVaultKeysResult, error) {
			args := v.(ListIntegrationAccountKeyVaultKeysArgs)
			r, err := ListIntegrationAccountKeyVaultKeys(ctx, &args, opts...)
			var s ListIntegrationAccountKeyVaultKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationAccountKeyVaultKeysResultOutput)
}

type ListIntegrationAccountKeyVaultKeysOutputArgs struct {
	IntegrationAccountName pulumi.StringInput     `pulumi:"integrationAccountName"`
	KeyVault               KeyVaultReferenceInput `pulumi:"keyVault"`
	ResourceGroupName      pulumi.StringInput     `pulumi:"resourceGroupName"`
	SkipToken              pulumi.StringPtrInput  `pulumi:"skipToken"`
}

func (ListIntegrationAccountKeyVaultKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountKeyVaultKeysArgs)(nil)).Elem()
}


type ListIntegrationAccountKeyVaultKeysResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationAccountKeyVaultKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountKeyVaultKeysResult)(nil)).Elem()
}

func (o ListIntegrationAccountKeyVaultKeysResultOutput) ToListIntegrationAccountKeyVaultKeysResultOutput() ListIntegrationAccountKeyVaultKeysResultOutput {
	return o
}

func (o ListIntegrationAccountKeyVaultKeysResultOutput) ToListIntegrationAccountKeyVaultKeysResultOutputWithContext(ctx context.Context) ListIntegrationAccountKeyVaultKeysResultOutput {
	return o
}

func (o ListIntegrationAccountKeyVaultKeysResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIntegrationAccountKeyVaultKeysResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListIntegrationAccountKeyVaultKeysResultOutput) Value() KeyVaultKeyResponseArrayOutput {
	return o.ApplyT(func(v ListIntegrationAccountKeyVaultKeysResult) []KeyVaultKeyResponse { return v.Value }).(KeyVaultKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationAccountKeyVaultKeysResultOutput{})
}
