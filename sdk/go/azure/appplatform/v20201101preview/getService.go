


package v20201101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:appplatform/v20201101preview:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServiceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupServiceResult struct {
	Id         string                            `pulumi:"id"`
	Location   *string                           `pulumi:"location"`
	Name       string                            `pulumi:"name"`
	Properties ClusterResourcePropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                      `pulumi:"sku"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       string                            `pulumi:"type"`
}


func (val *LookupServiceResult) Defaults() *LookupServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}
