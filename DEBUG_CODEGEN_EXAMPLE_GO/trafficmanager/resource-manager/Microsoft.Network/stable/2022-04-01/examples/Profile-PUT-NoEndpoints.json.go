package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewProfile(ctx, "profile", &network.ProfileArgs{
			DnsConfig: &network.DnsConfigArgs{
				RelativeName: pulumi.String("azsmnet6386"),
				Ttl:          pulumi.Float64(35),
			},
			Location: pulumi.String("global"),
			MonitorConfig: &network.MonitorConfigArgs{
				Path:     pulumi.String("/testpath.aspx"),
				Port:     pulumi.Float64(80),
				Protocol: pulumi.String("HTTP"),
			},
			ProfileName:          pulumi.String("azsmnet6386"),
			ProfileStatus:        pulumi.String("Enabled"),
			ResourceGroupName:    pulumi.String("azuresdkfornetautoresttrafficmanager1421"),
			TrafficRoutingMethod: pulumi.String("Performance"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
