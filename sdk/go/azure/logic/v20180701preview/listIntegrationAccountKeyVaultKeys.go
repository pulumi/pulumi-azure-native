


package v20180701preview

import (
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
