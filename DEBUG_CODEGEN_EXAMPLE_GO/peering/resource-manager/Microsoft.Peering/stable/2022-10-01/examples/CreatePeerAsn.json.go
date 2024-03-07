package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/peering/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := peering.NewPeerAsn(ctx, "peerAsn", &peering.PeerAsnArgs{
			PeerAsn:     pulumi.Int(65000),
			PeerAsnName: pulumi.String("peerAsnName"),
			PeerContactDetail: []peering.ContactDetailArgs{
				{
					Email: pulumi.String("noc@contoso.com"),
					Phone: pulumi.String("+1 (234) 567-8999"),
					Role:  pulumi.String("Noc"),
				},
				{
					Email: pulumi.String("abc@contoso.com"),
					Phone: pulumi.String("+1 (234) 567-8900"),
					Role:  pulumi.String("Policy"),
				},
				{
					Email: pulumi.String("xyz@contoso.com"),
					Phone: pulumi.String("+1 (234) 567-8900"),
					Role:  pulumi.String("Technical"),
				},
			},
			PeerName: pulumi.String("Contoso"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
