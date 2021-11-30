


package eventgrid

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPartnerNamespaceSharedAccessKeys(ctx *pulumi.Context, args *ListPartnerNamespaceSharedAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListPartnerNamespaceSharedAccessKeysResult, error) {
	var rv ListPartnerNamespaceSharedAccessKeysResult
	err := ctx.Invoke("azure-native:eventgrid:listPartnerNamespaceSharedAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPartnerNamespaceSharedAccessKeysArgs struct {
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type ListPartnerNamespaceSharedAccessKeysResult struct {
	Key1 *string `pulumi:"key1"`
	Key2 *string `pulumi:"key2"`
}
