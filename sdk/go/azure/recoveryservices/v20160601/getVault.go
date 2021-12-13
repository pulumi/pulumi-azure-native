


package v20160601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVault(ctx *pulumi.Context, args *LookupVaultArgs, opts ...pulumi.InvokeOption) (*LookupVaultResult, error) {
	var rv LookupVaultResult
	err := ctx.Invoke("azure-native:recoveryservices/v20160601:getVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVaultArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupVaultResult struct {
	ETag       *string                 `pulumi:"eTag"`
	Id         string                  `pulumi:"id"`
	Identity   *IdentityDataResponse   `pulumi:"identity"`
	Location   string                  `pulumi:"location"`
	Name       string                  `pulumi:"name"`
	Properties VaultPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse            `pulumi:"sku"`
	Tags       map[string]string       `pulumi:"tags"`
	Type       string                  `pulumi:"type"`
}
