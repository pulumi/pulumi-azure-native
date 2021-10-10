


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVendorSkuPreview(ctx *pulumi.Context, args *LookupVendorSkuPreviewArgs, opts ...pulumi.InvokeOption) (*LookupVendorSkuPreviewResult, error) {
	var rv LookupVendorSkuPreviewResult
	err := ctx.Invoke("azure-native:hybridnetwork/v20200101preview:getVendorSkuPreview", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVendorSkuPreviewArgs struct {
	PreviewSubscription string `pulumi:"previewSubscription"`
	SkuName             string `pulumi:"skuName"`
	VendorName          string `pulumi:"vendorName"`
}


type LookupVendorSkuPreviewResult struct {
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}
