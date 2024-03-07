package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewEndpoint(ctx, "endpoint", &network.EndpointArgs{
			EndpointName:   pulumi.String("My%20external%20endpoint"),
			EndpointStatus: pulumi.String("Enabled"),
			EndpointType:   pulumi.String("ExternalEndpoints"),
			GeoMapping: pulumi.StringArray{
				pulumi.String("GEO-AS"),
				pulumi.String("GEO-AF"),
			},
			Name:              pulumi.String("My external endpoint"),
			ProfileName:       pulumi.String("azuresdkfornetautoresttrafficmanager8224"),
			ResourceGroupName: pulumi.String("azuresdkfornetautoresttrafficmanager2191"),
			Target:            pulumi.String("foobar.contoso.com"),
			Type:              pulumi.String("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
