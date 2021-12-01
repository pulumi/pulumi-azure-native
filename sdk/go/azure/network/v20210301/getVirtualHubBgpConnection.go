


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualHubBgpConnection(ctx *pulumi.Context, args *LookupVirtualHubBgpConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubBgpConnectionResult, error) {
	var rv LookupVirtualHubBgpConnectionResult
	err := ctx.Invoke("azure-native:network/v20210301:getVirtualHubBgpConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubBgpConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupVirtualHubBgpConnectionResult struct {
	ConnectionState             string               `pulumi:"connectionState"`
	Etag                        string               `pulumi:"etag"`
	HubVirtualNetworkConnection *SubResourceResponse `pulumi:"hubVirtualNetworkConnection"`
	Id                          *string              `pulumi:"id"`
	Name                        *string              `pulumi:"name"`
	PeerAsn                     *float64             `pulumi:"peerAsn"`
	PeerIp                      *string              `pulumi:"peerIp"`
	ProvisioningState           string               `pulumi:"provisioningState"`
	Type                        string               `pulumi:"type"`
}
