package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewAFDOrigin(ctx, "afdOrigin", &cdn.AFDOriginArgs{
			EnabledState:      pulumi.String("Enabled"),
			HostName:          pulumi.String("host1.blob.core.windows.net"),
			HttpPort:          pulumi.Int(80),
			HttpsPort:         pulumi.Int(443),
			OriginGroupName:   pulumi.String("origingroup1"),
			OriginHostHeader:  pulumi.String("host1.foo.com"),
			OriginName:        pulumi.String("origin1"),
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
