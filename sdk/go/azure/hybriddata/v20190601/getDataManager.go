


package v20190601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataManager(ctx *pulumi.Context, args *LookupDataManagerArgs, opts ...pulumi.InvokeOption) (*LookupDataManagerResult, error) {
	var rv LookupDataManagerResult
	err := ctx.Invoke("azure-native:hybriddata/v20190601:getDataManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataManagerArgs struct {
	DataManagerName   string `pulumi:"dataManagerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDataManagerResult struct {
	Etag     *string           `pulumi:"etag"`
	Id       string            `pulumi:"id"`
	Location string            `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Sku      *SkuResponse      `pulumi:"sku"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}
