


package v20190201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDomainSharedAccessKeys(ctx *pulumi.Context, args *ListDomainSharedAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListDomainSharedAccessKeysResult, error) {
	var rv ListDomainSharedAccessKeysResult
	err := ctx.Invoke("azure-native:eventgrid/v20190201preview:listDomainSharedAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDomainSharedAccessKeysArgs struct {
	DomainName        string `pulumi:"domainName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDomainSharedAccessKeysResult struct {
	Key1 *string `pulumi:"key1"`
	Key2 *string `pulumi:"key2"`
}
