


package v20180401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVpnGateway(ctx *pulumi.Context, args *LookupVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVpnGatewayResult, error) {
	var rv LookupVpnGatewayResult
	err := ctx.Invoke("azure-native:network/v20180401:getVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnGatewayArgs struct {
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupVpnGatewayResult struct {
	BgpSettings       *BgpSettingsResponse    `pulumi:"bgpSettings"`
	Connections       []VpnConnectionResponse `pulumi:"connections"`
	Etag              string                  `pulumi:"etag"`
	Id                *string                 `pulumi:"id"`
	Location          string                  `pulumi:"location"`
	Name              string                  `pulumi:"name"`
	Policies          *PoliciesResponse       `pulumi:"policies"`
	ProvisioningState string                  `pulumi:"provisioningState"`
	Tags              map[string]string       `pulumi:"tags"`
	Type              string                  `pulumi:"type"`
	VirtualHub        *SubResourceResponse    `pulumi:"virtualHub"`
}
