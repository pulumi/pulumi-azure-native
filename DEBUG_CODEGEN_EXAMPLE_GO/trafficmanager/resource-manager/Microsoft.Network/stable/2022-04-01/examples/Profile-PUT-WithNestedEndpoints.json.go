package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewProfile(ctx, "profile", &network.ProfileArgs{
			DnsConfig: &network.DnsConfigArgs{
				RelativeName: pulumi.String("parentprofile"),
				Ttl:          pulumi.Float64(35),
			},
			Endpoints: []network.EndpointTypeArgs{
				{
					EndpointStatus:        pulumi.String("Enabled"),
					MinChildEndpoints:     pulumi.Float64(2),
					MinChildEndpointsIPv4: pulumi.Float64(1),
					MinChildEndpointsIPv6: pulumi.Float64(2),
					Name:                  pulumi.String("MyFirstNestedEndpoint"),
					Priority:              pulumi.Float64(1),
					Target:                pulumi.String("firstnestedprofile.tmpreview.watmtest.azure-test.net"),
					Type:                  pulumi.String("Microsoft.Network/trafficManagerProfiles/nestedEndpoints"),
					Weight:                pulumi.Float64(1),
				},
				{
					EndpointStatus:        pulumi.String("Enabled"),
					MinChildEndpoints:     pulumi.Float64(2),
					MinChildEndpointsIPv4: pulumi.Float64(2),
					MinChildEndpointsIPv6: pulumi.Float64(1),
					Name:                  pulumi.String("MySecondNestedEndpoint"),
					Priority:              pulumi.Float64(2),
					Target:                pulumi.String("secondnestedprofile.tmpreview.watmtest.azure-test.net"),
					Type:                  pulumi.String("Microsoft.Network/trafficManagerProfiles/nestedEndpoints"),
					Weight:                pulumi.Float64(1),
				},
			},
			Location: pulumi.String("global"),
			MonitorConfig: &network.MonitorConfigArgs{
				IntervalInSeconds:         pulumi.Float64(10),
				Path:                      pulumi.String("/testpath.aspx"),
				Port:                      pulumi.Float64(80),
				Protocol:                  pulumi.String("HTTP"),
				TimeoutInSeconds:          pulumi.Float64(5),
				ToleratedNumberOfFailures: pulumi.Float64(2),
			},
			ProfileName:          pulumi.String("parentprofile"),
			ProfileStatus:        pulumi.String("Enabled"),
			ResourceGroupName:    pulumi.String("myresourcegroup"),
			TrafficRoutingMethod: pulumi.String("Priority"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
