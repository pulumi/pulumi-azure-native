


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInstancePool(ctx *pulumi.Context, args *LookupInstancePoolArgs, opts ...pulumi.InvokeOption) (*LookupInstancePoolResult, error) {
	var rv LookupInstancePoolResult
	err := ctx.Invoke("azure-native:sql/v20180601preview:getInstancePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstancePoolArgs struct {
	InstancePoolName  string `pulumi:"instancePoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupInstancePoolResult struct {
	Id          string            `pulumi:"id"`
	LicenseType string            `pulumi:"licenseType"`
	Location    string            `pulumi:"location"`
	Name        string            `pulumi:"name"`
	Sku         *SkuResponse      `pulumi:"sku"`
	SubnetId    string            `pulumi:"subnetId"`
	Tags        map[string]string `pulumi:"tags"`
	Type        string            `pulumi:"type"`
	VCores      int               `pulumi:"vCores"`
}
