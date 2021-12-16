


package v20200401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVault(ctx *pulumi.Context, args *LookupVaultArgs, opts ...pulumi.InvokeOption) (*LookupVaultResult, error) {
	var rv LookupVaultResult
	err := ctx.Invoke("azure-native:keyvault/v20200401preview:getVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVaultArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupVaultResult struct {
	Id         string                  `pulumi:"id"`
	Location   *string                 `pulumi:"location"`
	Name       string                  `pulumi:"name"`
	Properties VaultPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse      `pulumi:"systemData"`
	Tags       map[string]string       `pulumi:"tags"`
	Type       string                  `pulumi:"type"`
}


func (val *LookupVaultResult) Defaults() *LookupVaultResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
