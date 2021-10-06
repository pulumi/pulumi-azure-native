


package v20200701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSignalR(ctx *pulumi.Context, args *LookupSignalRArgs, opts ...pulumi.InvokeOption) (*LookupSignalRResult, error) {
	var rv LookupSignalRResult
	err := ctx.Invoke("azure-native:signalrservice/v20200701preview:getSignalR", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSignalRResult struct {
	Cors                       *SignalRCorsSettingsResponse        `pulumi:"cors"`
	ExternalIP                 string                              `pulumi:"externalIP"`
	Features                   []SignalRFeatureResponse            `pulumi:"features"`
	HostName                   string                              `pulumi:"hostName"`
	Id                         string                              `pulumi:"id"`
	Identity                   *ManagedIdentityResponse            `pulumi:"identity"`
	Kind                       *string                             `pulumi:"kind"`
	Location                   *string                             `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	NetworkACLs                *SignalRNetworkACLsResponse         `pulumi:"networkACLs"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicPort                 int                                 `pulumi:"publicPort"`
	ServerPort                 int                                 `pulumi:"serverPort"`
	Sku                        *ResourceSkuResponse                `pulumi:"sku"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Tls                        *SignalRTlsSettingsResponse         `pulumi:"tls"`
	Type                       string                              `pulumi:"type"`
	Upstream                   *ServerlessUpstreamSettingsResponse `pulumi:"upstream"`
	Version                    string                              `pulumi:"version"`
}
