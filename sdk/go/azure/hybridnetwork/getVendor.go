


package hybridnetwork

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVendor(ctx *pulumi.Context, args *LookupVendorArgs, opts ...pulumi.InvokeOption) (*LookupVendorResult, error) {
	var rv LookupVendorResult
	err := ctx.Invoke("azure-native:hybridnetwork:getVendor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVendorArgs struct {
	VendorName string `pulumi:"vendorName"`
}


type LookupVendorResult struct {
	Id                string                `pulumi:"id"`
	Name              string                `pulumi:"name"`
	ProvisioningState string                `pulumi:"provisioningState"`
	Skus              []SubResourceResponse `pulumi:"skus"`
	Type              string                `pulumi:"type"`
}
