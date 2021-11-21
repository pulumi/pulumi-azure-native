


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebPubSub(ctx *pulumi.Context, args *LookupWebPubSubArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubResult, error) {
	var rv LookupWebPubSubResult
	err := ctx.Invoke("azure-native:webpubsub/v20211001:getWebPubSub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWebPubSubResult struct {
	DisableAadAuth             *bool                               `pulumi:"disableAadAuth"`
	DisableLocalAuth           *bool                               `pulumi:"disableLocalAuth"`
	ExternalIP                 string                              `pulumi:"externalIP"`
	HostName                   string                              `pulumi:"hostName"`
	HostNamePrefix             string                              `pulumi:"hostNamePrefix"`
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
	ResourceLogConfiguration   *ResourceLogConfigurationResponse   `pulumi:"resourceLogConfiguration"`
	ServerPort                 int                                 `pulumi:"serverPort"`
	SharedPrivateLinkResources []SharedPrivateLinkResourceResponse `pulumi:"sharedPrivateLinkResources"`
	Sku                        *ResourceSkuResponse                `pulumi:"sku"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Tls                        *WebPubSubTlsSettingsResponse       `pulumi:"tls"`
	Type                       string                              `pulumi:"type"`
	Version                    string                              `pulumi:"version"`
}
