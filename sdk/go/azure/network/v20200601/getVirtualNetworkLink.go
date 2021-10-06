


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkLink(ctx *pulumi.Context, args *LookupVirtualNetworkLinkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkLinkResult, error) {
	var rv LookupVirtualNetworkLinkResult
	err := ctx.Invoke("azure-native:network/v20200601:getVirtualNetworkLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkLinkArgs struct {
	PrivateZoneName        string `pulumi:"privateZoneName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	VirtualNetworkLinkName string `pulumi:"virtualNetworkLinkName"`
}


type LookupVirtualNetworkLinkResult struct {
	Etag                    *string              `pulumi:"etag"`
	Id                      string               `pulumi:"id"`
	Location                *string              `pulumi:"location"`
	Name                    string               `pulumi:"name"`
	ProvisioningState       string               `pulumi:"provisioningState"`
	RegistrationEnabled     *bool                `pulumi:"registrationEnabled"`
	Tags                    map[string]string    `pulumi:"tags"`
	Type                    string               `pulumi:"type"`
	VirtualNetwork          *SubResourceResponse `pulumi:"virtualNetwork"`
	VirtualNetworkLinkState string               `pulumi:"virtualNetworkLinkState"`
}
