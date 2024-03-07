package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewOrigin(ctx, "origin", &cdn.OriginArgs{
			Enabled:                    pulumi.Bool(true),
			EndpointName:               pulumi.String("endpoint1"),
			HostName:                   pulumi.String("www.someDomain.net"),
			HttpPort:                   pulumi.Int(80),
			HttpsPort:                  pulumi.Int(443),
			OriginHostHeader:           pulumi.String("www.someDomain.net"),
			OriginName:                 pulumi.String("www-someDomain-net"),
			Priority:                   pulumi.Int(1),
			PrivateLinkApprovalMessage: pulumi.String("Please approve the connection request for this Private Link"),
			PrivateLinkLocation:        pulumi.String("eastus"),
			PrivateLinkResourceId:      pulumi.String("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Network/privateLinkServices/pls1"),
			ProfileName:                pulumi.String("profile1"),
			ResourceGroupName:          pulumi.String("RG"),
			Weight:                     pulumi.Int(50),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
