package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/peering/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := peering.NewPeering(ctx, "peering", &peering.PeeringArgs{
			Direct: &peering.PeeringPropertiesDirectArgs{
				Connections: peering.DirectConnectionArray{
					&peering.DirectConnectionArgs{
						BandwidthInMbps: pulumi.Int(10000),
						BgpSession: &peering.BgpSessionArgs{
							MaxPrefixesAdvertisedV4:     pulumi.Int(1000),
							MaxPrefixesAdvertisedV6:     pulumi.Int(100),
							MicrosoftSessionIPv4Address: pulumi.String("192.168.0.123"),
							PeerSessionIPv4Address:      pulumi.String("192.168.0.234"),
							SessionPrefixV4:             pulumi.String("192.168.0.0/24"),
						},
						ConnectionIdentifier:   pulumi.String("5F4CB5C7-6B43-4444-9338-9ABC72606C16"),
						PeeringDBFacilityId:    pulumi.Int(99999),
						SessionAddressProvider: pulumi.String("Peer"),
						UseForPeeringService:   pulumi.Bool(true),
					},
				},
				DirectPeeringType: pulumi.String("IxRs"),
				PeerAsn: &peering.SubResourceArgs{
					Id: pulumi.String("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
				},
			},
			Kind:              pulumi.String("Direct"),
			Location:          pulumi.String("eastus"),
			PeeringLocation:   pulumi.String("peeringLocation0"),
			PeeringName:       pulumi.String("peeringName"),
			ResourceGroupName: pulumi.String("rgName"),
			Sku: &peering.PeeringSkuArgs{
				Name: pulumi.String("Premium_Direct_Free"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
