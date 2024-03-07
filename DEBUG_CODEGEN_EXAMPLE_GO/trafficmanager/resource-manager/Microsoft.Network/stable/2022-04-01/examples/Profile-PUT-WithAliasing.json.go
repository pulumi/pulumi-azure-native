package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewProfile(ctx, "profile", &network.ProfileArgs{
			AllowedEndpointRecordTypes: pulumi.StringArray{
				pulumi.String("DomainName"),
			},
			DnsConfig: &network.DnsConfigArgs{
				RelativeName: pulumi.String("azuresdkfornetautoresttrafficmanager6192"),
				Ttl:          pulumi.Float64(35),
			},
			Endpoints: network.EndpointTypeArray{
				&network.EndpointTypeArgs{
					EndpointLocation: pulumi.String("North Europe"),
					EndpointStatus:   pulumi.String("Enabled"),
					Name:             pulumi.String("My external endpoint"),
					Target:           pulumi.String("foobar.contoso.com"),
					Type:             pulumi.String("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
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
			ProfileName:          pulumi.String("azuresdkfornetautoresttrafficmanager6192"),
			ProfileStatus:        pulumi.String("Enabled"),
			ResourceGroupName:    pulumi.String("azuresdkfornetautoresttrafficmanager2583"),
			TrafficRoutingMethod: pulumi.String("Performance"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
