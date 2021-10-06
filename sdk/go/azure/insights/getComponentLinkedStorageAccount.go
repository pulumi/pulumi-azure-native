


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComponentLinkedStorageAccount(ctx *pulumi.Context, args *LookupComponentLinkedStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupComponentLinkedStorageAccountResult, error) {
	var rv LookupComponentLinkedStorageAccountResult
	err := ctx.Invoke("azure-native:insights:getComponentLinkedStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComponentLinkedStorageAccountArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
	StorageType       string `pulumi:"storageType"`
}


type LookupComponentLinkedStorageAccountResult struct {
	Id                   string  `pulumi:"id"`
	LinkedStorageAccount *string `pulumi:"linkedStorageAccount"`
	Name                 string  `pulumi:"name"`
	Type                 string  `pulumi:"type"`
}
