


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkFunction(ctx *pulumi.Context, args *LookupNetworkFunctionArgs, opts ...pulumi.InvokeOption) (*LookupNetworkFunctionResult, error) {
	var rv LookupNetworkFunctionResult
	err := ctx.Invoke("azure-native:hybridnetwork/v20200101preview:getNetworkFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkFunctionArgs struct {
	NetworkFunctionName string `pulumi:"networkFunctionName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupNetworkFunctionResult struct {
	Device                            *SubResourceResponse                       `pulumi:"device"`
	Etag                              *string                                    `pulumi:"etag"`
	Id                                string                                     `pulumi:"id"`
	Location                          string                                     `pulumi:"location"`
	ManagedApplication                SubResourceResponse                        `pulumi:"managedApplication"`
	ManagedApplicationParameters      interface{}                                `pulumi:"managedApplicationParameters"`
	Name                              string                                     `pulumi:"name"`
	NetworkFunctionUserConfigurations []NetworkFunctionUserConfigurationResponse `pulumi:"networkFunctionUserConfigurations"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	ServiceKey                        string                                     `pulumi:"serviceKey"`
	SkuName                           *string                                    `pulumi:"skuName"`
	SkuType                           string                                     `pulumi:"skuType"`
	Tags                              map[string]string                          `pulumi:"tags"`
	Type                              string                                     `pulumi:"type"`
	VendorName                        *string                                    `pulumi:"vendorName"`
	VendorProvisioningState           string                                     `pulumi:"vendorProvisioningState"`
}
