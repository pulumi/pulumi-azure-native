package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/webpubsub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := webpubsub.NewWebPubSubHub(ctx, "webPubSubHub", &webpubsub.WebPubSubHubArgs{
HubName: pulumi.String("exampleHub"),
Properties: webpubsub.WebPubSubHubPropertiesResponse{
AnonymousConnectPolicy: pulumi.String("allow"),
EventHandlers: webpubsub.EventHandlerArray{
interface{}{
Auth: interface{}{
ManagedIdentity: &webpubsub.ManagedIdentitySettingsArgs{
Resource: pulumi.String("abc"),
},
Type: pulumi.String("ManagedIdentity"),
},
SystemEvents: pulumi.StringArray{
pulumi.String("connect"),
pulumi.String("connected"),
},
UrlTemplate: pulumi.String("http://host.com"),
UserEventPattern: pulumi.String("*"),
},
},
EventListeners: []webpubsub.EventListenerArgs{
{
Endpoint: {
EventHubName: pulumi.String("eventHubName1"),
FullyQualifiedNamespace: pulumi.String("example.servicebus.windows.net"),
Type: pulumi.String("EventHub"),
},
Filter: {
SystemEvents: pulumi.StringArray{
pulumi.String("connected"),
pulumi.String("disconnected"),
},
Type: pulumi.String("EventName"),
UserEventPattern: pulumi.String("*"),
},
},
},
},
ResourceGroupName: pulumi.String("myResourceGroup"),
ResourceName: pulumi.String("myWebPubSubService"),
})
if err != nil {
return err
}
return nil
})
}
