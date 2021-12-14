


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebPubSub(ctx *pulumi.Context, args *LookupWebPubSubArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubResult, error) {
	var rv LookupWebPubSubResult
	err := ctx.Invoke("azure-native:webpubsub/v20210901preview:getWebPubSub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebPubSubArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWebPubSubResult struct {
	DisableAadAuth             *bool                               `pulumi:"disableAadAuth"`
	DisableLocalAuth           *bool                               `pulumi:"disableLocalAuth"`
	EventHandler               *EventHandlerSettingsResponse       `pulumi:"eventHandler"`
	ExternalIP                 string                              `pulumi:"externalIP"`
	HostName                   string                              `pulumi:"hostName"`
	Id                         string                              `pulumi:"id"`
	Identity                   *ManagedIdentityResponse            `pulumi:"identity"`
	LiveTraceConfiguration     *LiveTraceConfigurationResponse     `pulumi:"liveTraceConfiguration"`
	Location                   *string                             `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	NetworkACLs                *WebPubSubNetworkACLsResponse       `pulumi:"networkACLs"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	PublicPort                 int                                 `pulumi:"publicPort"`
	ServerPort                 int                                 `pulumi:"serverPort"`
	SharedPrivateLinkResources []SharedPrivateLinkResourceResponse `pulumi:"sharedPrivateLinkResources"`
	Sku                        *ResourceSkuResponse                `pulumi:"sku"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Tls                        *WebPubSubTlsSettingsResponse       `pulumi:"tls"`
	Type                       string                              `pulumi:"type"`
	Version                    string                              `pulumi:"version"`
}


func (val *LookupWebPubSubResult) Defaults() *LookupWebPubSubResult {
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
	tmp.LiveTraceConfiguration = tmp.LiveTraceConfiguration.Defaults()

	tmp.NetworkACLs = tmp.NetworkACLs.Defaults()

	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	tmp.Tls = tmp.Tls.Defaults()

	return &tmp
}
