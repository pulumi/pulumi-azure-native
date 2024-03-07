package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewFleet(ctx, "fleet", &containerservice.FleetArgs{
			FleetName: pulumi.String("fleet1"),
			HubProfile: &containerservice.FleetHubProfileArgs{
				DnsPrefix: pulumi.String("dnsprefix1"),
			},
			Location:          pulumi.String("East US"),
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"archv2": pulumi.String(""),
				"tier":   pulumi.String("production"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
