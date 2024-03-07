package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/deviceupdate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := deviceupdate.NewPrivateEndpointConnectionProxy(ctx, "privateEndpointConnectionProxy", &deviceupdate.PrivateEndpointConnectionProxyArgs{
AccountName: pulumi.String("contoso"),
PrivateEndpointConnectionProxyId: pulumi.String("peexample01"),
RemotePrivateEndpoint: deviceupdate.RemotePrivateEndpointResponse{
Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
ImmutableResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
ImmutableSubscriptionId: pulumi.String("00000000-0000-0000-0000-000000000000"),
Location: pulumi.String("westus2"),
ManualPrivateLinkServiceConnections: deviceupdate.PrivateLinkServiceConnectionArray{
&deviceupdate.PrivateLinkServiceConnectionArgs{
GroupIds: pulumi.StringArray{
pulumi.String("DeviceUpdate"),
},
Name: pulumi.String("{privateEndpointConnectionProxyId}"),
RequestMessage: pulumi.String("Please approve my connection, thanks."),
},
},
PrivateLinkServiceProxies: deviceupdate.PrivateLinkServiceProxyArray{
interface{}{
GroupConnectivityInformation: deviceupdate.GroupConnectivityInformationArray{
},
Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}"),
},
},
},
ResourceGroupName: pulumi.String("test-rg"),
})
if err != nil {
return err
}
return nil
})
}
