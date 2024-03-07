package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/peering/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := peering.NewPeering(ctx, "peering", &peering.PeeringArgs{
Exchange: peering.PeeringPropertiesExchangeResponse{
Connections: peering.ExchangeConnectionArray{
interface{}{
BgpSession: &peering.BgpSessionArgs{
MaxPrefixesAdvertisedV4: pulumi.Int(1000),
MaxPrefixesAdvertisedV6: pulumi.Int(100),
Md5AuthenticationKey: pulumi.String("test-md5-auth-key"),
PeerSessionIPv4Address: pulumi.String("192.168.2.1"),
PeerSessionIPv6Address: pulumi.String("fd00::1"),
},
ConnectionIdentifier: pulumi.String("CE495334-0E94-4E51-8164-8116D6CD284D"),
PeeringDBFacilityId: pulumi.Int(99999),
},
interface{}{
BgpSession: &peering.BgpSessionArgs{
MaxPrefixesAdvertisedV4: pulumi.Int(1000),
MaxPrefixesAdvertisedV6: pulumi.Int(100),
Md5AuthenticationKey: pulumi.String("test-md5-auth-key"),
PeerSessionIPv4Address: pulumi.String("192.168.2.2"),
PeerSessionIPv6Address: pulumi.String("fd00::2"),
},
ConnectionIdentifier: pulumi.String("CDD8E673-CB07-47E6-84DE-3739F778762B"),
PeeringDBFacilityId: pulumi.Int(99999),
},
},
PeerAsn: &peering.SubResourceArgs{
Id: pulumi.String("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
},
},
Kind: pulumi.String("Exchange"),
Location: pulumi.String("eastus"),
PeeringLocation: pulumi.String("peeringLocation0"),
PeeringName: pulumi.String("peeringName"),
ResourceGroupName: pulumi.String("rgName"),
Sku: &peering.PeeringSkuArgs{
Name: pulumi.String("Basic_Exchange_Free"),
},
})
if err != nil {
return err
}
return nil
})
}
