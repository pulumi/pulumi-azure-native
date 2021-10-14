


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKey(ctx *pulumi.Context, args *LookupKeyArgs, opts ...pulumi.InvokeOption) (*LookupKeyResult, error) {
	var rv LookupKeyResult
	err := ctx.Invoke("azure-native:synapse/v20210501:getKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKeyArgs struct {
	KeyName           string `pulumi:"keyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupKeyResult struct {
	Id          string  `pulumi:"id"`
	IsActiveCMK *bool   `pulumi:"isActiveCMK"`
	KeyVaultUrl *string `pulumi:"keyVaultUrl"`
	Name        string  `pulumi:"name"`
	Type        string  `pulumi:"type"`
}
