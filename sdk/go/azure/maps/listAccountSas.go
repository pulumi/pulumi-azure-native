


package maps

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAccountSas(ctx *pulumi.Context, args *ListAccountSasArgs, opts ...pulumi.InvokeOption) (*ListAccountSasResult, error) {
	var rv ListAccountSasResult
	err := ctx.Invoke("azure-native:maps:listAccountSas", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountSasArgs struct {
	AccountName       string   `pulumi:"accountName"`
	Expiry            string   `pulumi:"expiry"`
	MaxRatePerSecond  int      `pulumi:"maxRatePerSecond"`
	PrincipalId       string   `pulumi:"principalId"`
	Regions           []string `pulumi:"regions"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	SigningKey        string   `pulumi:"signingKey"`
	Start             string   `pulumi:"start"`
}


type ListAccountSasResult struct {
	AccountSasToken string `pulumi:"accountSasToken"`
}
