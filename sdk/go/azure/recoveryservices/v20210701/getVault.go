


package v20210701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVault(ctx *pulumi.Context, args *LookupVaultArgs, opts ...pulumi.InvokeOption) (*LookupVaultResult, error) {
	var rv LookupVaultResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210701:getVault", args, &rv, opts...)
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
	Etag       *string                 `pulumi:"etag"`
	Id         string                  `pulumi:"id"`
	Identity   *IdentityDataResponse   `pulumi:"identity"`
	Location   string                  `pulumi:"location"`
	Name       string                  `pulumi:"name"`
	Properties VaultPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse            `pulumi:"sku"`
	SystemData SystemDataResponse      `pulumi:"systemData"`
	Tags       map[string]string       `pulumi:"tags"`
	Type       string                  `pulumi:"type"`
}
