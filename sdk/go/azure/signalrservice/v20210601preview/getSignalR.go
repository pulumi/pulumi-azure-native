


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSignalR(ctx *pulumi.Context, args *LookupSignalRArgs, opts ...pulumi.InvokeOption) (*LookupSignalRResult, error) {
	var rv LookupSignalRResult
	err := ctx.Invoke("azure-native:signalrservice/v20210601preview:getSignalR", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSignalRArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSignalRResult struct {
	Cors                       *SignalRCorsSettingsResponse        `pulumi:"cors"`
	DisableAadAuth             *bool                               `pulumi:"disableAadAuth"`
	DisableLocalAuth           *bool                               `pulumi:"disableLocalAuth"`
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
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	PublicPort                 int                                 `pulumi:"publicPort"`
	ServerPort                 int                                 `pulumi:"serverPort"`
	SharedPrivateLinkResources []SharedPrivateLinkResourceResponse `pulumi:"sharedPrivateLinkResources"`
	Sku                        *ResourceSkuResponse                `pulumi:"sku"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Tls                        *SignalRTlsSettingsResponse         `pulumi:"tls"`
	Type                       string                              `pulumi:"type"`
	Upstream                   *ServerlessUpstreamSettingsResponse `pulumi:"upstream"`
	Version                    string                              `pulumi:"version"`
}


func (val *LookupSignalRResult) Defaults() *LookupSignalRResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DisableAadAuth) {
		disableAadAuth_ := false
		tmp.DisableAadAuth = &disableAadAuth_
	}
	if isZero(tmp.DisableLocalAuth) {
		disableLocalAuth_ := false
		tmp.DisableLocalAuth = &disableLocalAuth_
	}
	tmp.NetworkACLs = tmp.NetworkACLs.Defaults()

	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	tmp.Tls = tmp.Tls.Defaults()

	return &tmp
}
