package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/digitaltwins/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := digitaltwins.NewDigitalTwinsEndpoint(ctx, "digitalTwinsEndpoint", &digitaltwins.DigitalTwinsEndpointArgs{
			EndpointName: pulumi.String("myServiceBus"),
			Properties: digitaltwins.ServiceBus{
				AuthenticationType: "IdentityBased",
				EndpointType:       "ServiceBus",
				EndpointUri:        "sb://mysb.servicebus.windows.net/",
				EntityPath:         "mysbtopic",
			},
			ResourceGroupName: pulumi.String("resRg"),
			ResourceName:      pulumi.String("myDigitalTwinsService"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
