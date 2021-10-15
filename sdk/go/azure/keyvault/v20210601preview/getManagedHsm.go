


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedHsm(ctx *pulumi.Context, args *LookupManagedHsmArgs, opts ...pulumi.InvokeOption) (*LookupManagedHsmResult, error) {
	var rv LookupManagedHsmResult
	err := ctx.Invoke("azure-native:keyvault/v20210601preview:getManagedHsm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedHsmArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagedHsmResult struct {
	Id         string                       `pulumi:"id"`
	Location   *string                      `pulumi:"location"`
	Name       string                       `pulumi:"name"`
	Properties ManagedHsmPropertiesResponse `pulumi:"properties"`
	Sku        *ManagedHsmSkuResponse       `pulumi:"sku"`
	SystemData SystemDataResponse           `pulumi:"systemData"`
	Tags       map[string]string            `pulumi:"tags"`
	Type       string                       `pulumi:"type"`
}
