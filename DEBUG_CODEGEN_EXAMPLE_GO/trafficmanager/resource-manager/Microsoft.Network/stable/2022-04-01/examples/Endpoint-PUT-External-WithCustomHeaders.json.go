package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewEndpoint(ctx, "endpoint", &network.EndpointArgs{
			CustomHeaders: network.EndpointPropertiesCustomHeadersArray{
				&network.EndpointPropertiesCustomHeadersArgs{
					Name:  pulumi.String("header-1"),
					Value: pulumi.String("value-1"),
				},
				&network.EndpointPropertiesCustomHeadersArgs{
					Name:  pulumi.String("header-2"),
					Value: pulumi.String("value-2"),
				},
			},
			EndpointLocation:  pulumi.String("North Europe"),
			EndpointName:      pulumi.String("azsmnet7187"),
			EndpointStatus:    pulumi.String("Enabled"),
			EndpointType:      pulumi.String("ExternalEndpoints"),
			Name:              pulumi.String("azsmnet7187"),
			ProfileName:       pulumi.String("azsmnet6386"),
			ResourceGroupName: pulumi.String("azuresdkfornetautoresttrafficmanager1421"),
			Target:            pulumi.String("foobar.contoso.com"),
			Type:              pulumi.String("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
