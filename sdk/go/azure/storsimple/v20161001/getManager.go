


package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManager(ctx *pulumi.Context, args *LookupManagerArgs, opts ...pulumi.InvokeOption) (*LookupManagerResult, error) {
	var rv LookupManagerResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagerArgs struct {
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagerResult struct {
	CisIntrinsicSettings *ManagerIntrinsicSettingsResponse `pulumi:"cisIntrinsicSettings"`
	Etag                 *string                           `pulumi:"etag"`
	Id                   string                            `pulumi:"id"`
	Location             string                            `pulumi:"location"`
	Name                 string                            `pulumi:"name"`
	ProvisioningState    string                            `pulumi:"provisioningState"`
	Sku                  *ManagerSkuResponse               `pulumi:"sku"`
	Tags                 map[string]string                 `pulumi:"tags"`
	Type                 string                            `pulumi:"type"`
}
