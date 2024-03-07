package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/signalrservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := signalrservice.NewSignalR(ctx, "signalR", &signalrservice.SignalRArgs{
Cors: &signalrservice.SignalRCorsSettingsArgs{
AllowedOrigins: pulumi.StringArray{
pulumi.String("https://foo.com"),
pulumi.String("https://bar.com"),
},
},
DisableAadAuth: pulumi.Bool(false),
DisableLocalAuth: pulumi.Bool(false),
Features: []signalrservice.SignalRFeatureArgs{
{
Flag: pulumi.String("ServiceMode"),
Properties: nil,
Value: pulumi.String("Serverless"),
},
{
Flag: pulumi.String("EnableConnectivityLogs"),
Properties: nil,
Value: pulumi.String("True"),
},
{
Flag: pulumi.String("EnableMessagingLogs"),
Properties: nil,
Value: pulumi.String("False"),
},
{
Flag: pulumi.String("EnableLiveTrace"),
Properties: nil,
Value: pulumi.String("False"),
},
},
Identity: &signalrservice.ManagedIdentityArgs{
Type: pulumi.String("SystemAssigned"),
},
Kind: pulumi.String("SignalR"),
LiveTraceConfiguration: signalrservice.LiveTraceConfigurationResponse{
Categories: signalrservice.LiveTraceCategoryArray{
&signalrservice.LiveTraceCategoryArgs{
Enabled: pulumi.String("true"),
Name: pulumi.String("ConnectivityLogs"),
},
},
Enabled: pulumi.String("false"),
},
Location: pulumi.String("eastus"),
NetworkACLs: signalrservice.SignalRNetworkACLsResponse{
DefaultAction: pulumi.String("Deny"),
PrivateEndpoints: signalrservice.PrivateEndpointACLArray{
&signalrservice.PrivateEndpointACLArgs{
Allow: pulumi.StringArray{
pulumi.String("ServerConnection"),
},
Name: pulumi.String("mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
},
},
PublicNetwork: &signalrservice.NetworkACLArgs{
Allow: pulumi.StringArray{
pulumi.String("ClientConnection"),
},
},
},
PublicNetworkAccess: pulumi.String("Enabled"),
ResourceGroupName: pulumi.String("myResourceGroup"),
ResourceName: pulumi.String("mySignalRService"),
Serverless: &signalrservice.ServerlessSettingsArgs{
ConnectionTimeoutInSeconds: pulumi.Int(5),
},
Sku: &signalrservice.ResourceSkuArgs{
Capacity: pulumi.Int(1),
Name: pulumi.String("Premium_P1"),
Tier: pulumi.String("Premium"),
},
Tags: pulumi.StringMap{
"key1": pulumi.String("value1"),
},
Tls: &signalrservice.SignalRTlsSettingsArgs{
ClientCertEnabled: pulumi.Bool(false),
},
Upstream: signalrservice.ServerlessUpstreamSettingsResponse{
Templates: signalrservice.UpstreamTemplateArray{
interface{}{
Auth: interface{}{
ManagedIdentity: &signalrservice.ManagedIdentitySettingsArgs{
Resource: pulumi.String("api://example"),
},
Type: pulumi.String("ManagedIdentity"),
},
CategoryPattern: pulumi.String("*"),
EventPattern: pulumi.String("connect,disconnect"),
HubPattern: pulumi.String("*"),
UrlTemplate: pulumi.String("https://example.com/chat/api/connect"),
},
},
},
})
if err != nil {
return err
}
return nil
})
}
