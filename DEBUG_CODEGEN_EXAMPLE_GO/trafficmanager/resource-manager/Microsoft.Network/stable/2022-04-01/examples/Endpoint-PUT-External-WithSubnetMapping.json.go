package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewEndpoint(ctx, "endpoint", &network.EndpointArgs{
			EndpointName:      pulumi.String("My%20external%20endpoint"),
			EndpointStatus:    pulumi.String("Enabled"),
			EndpointType:      pulumi.String("ExternalEndpoints"),
			Name:              pulumi.String("My external endpoint"),
			ProfileName:       pulumi.String("azuresdkfornetautoresttrafficmanager8224"),
			ResourceGroupName: pulumi.String("azuresdkfornetautoresttrafficmanager2191"),
			Subnets: []network.EndpointPropertiesSubnetsArgs{
				{
					First: pulumi.String("1.2.3.0"),
					Scope: pulumi.Int(24),
				},
				{
					First: pulumi.String("25.26.27.28"),
					Last:  pulumi.String("29.30.31.32"),
				},
			},
			Target: pulumi.String("foobar.contoso.com"),
			Type:   pulumi.String("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
