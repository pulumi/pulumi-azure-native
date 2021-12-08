


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkSecurityPerimeter(ctx *pulumi.Context, args *LookupNetworkSecurityPerimeterArgs, opts ...pulumi.InvokeOption) (*LookupNetworkSecurityPerimeterResult, error) {
	var rv LookupNetworkSecurityPerimeterResult
	err := ctx.Invoke("azure-native:network:getNetworkSecurityPerimeter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkSecurityPerimeterArgs struct {
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupNetworkSecurityPerimeterResult struct {
	Description       *string           `pulumi:"description"`
	DisplayName       *string           `pulumi:"displayName"`
	Etag              string            `pulumi:"etag"`
	Id                string            `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}
