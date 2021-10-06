


package datalakestore

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTrustedIdProvider(ctx *pulumi.Context, args *LookupTrustedIdProviderArgs, opts ...pulumi.InvokeOption) (*LookupTrustedIdProviderResult, error) {
	var rv LookupTrustedIdProviderResult
	err := ctx.Invoke("azure-native:datalakestore:getTrustedIdProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrustedIdProviderArgs struct {
	AccountName           string `pulumi:"accountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TrustedIdProviderName string `pulumi:"trustedIdProviderName"`
}


type LookupTrustedIdProviderResult struct {
	Id         string `pulumi:"id"`
	IdProvider string `pulumi:"idProvider"`
	Name       string `pulumi:"name"`
	Type       string `pulumi:"type"`
}
