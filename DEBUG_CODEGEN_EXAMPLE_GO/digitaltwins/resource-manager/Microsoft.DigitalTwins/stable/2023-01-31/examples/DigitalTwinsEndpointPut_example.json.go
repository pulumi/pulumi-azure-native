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
				AuthenticationType:        "KeyBased",
				EndpointType:              "ServiceBus",
				PrimaryConnectionString:   "Endpoint=sb://mysb.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=xyzxyzoX4=;EntityPath=abcabc",
				SecondaryConnectionString: "Endpoint=sb://mysb.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=xyzxyzoX4=;EntityPath=abcabc",
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
