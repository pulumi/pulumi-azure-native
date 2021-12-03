


package keyvault

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	var rv LookupSecretResult
	err := ctx.Invoke("azure-native:keyvault:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SecretName        string `pulumi:"secretName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupSecretResult struct {
	Id         string                   `pulumi:"id"`
	Location   string                   `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Properties SecretPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string        `pulumi:"tags"`
	Type       string                   `pulumi:"type"`
}
