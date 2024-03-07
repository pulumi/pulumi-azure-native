package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/digitaltwins/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := digitaltwins.NewDigitalTwin(ctx, "digitalTwin", &digitaltwins.DigitalTwinArgs{
			Location:          pulumi.String("WestUS2"),
			ResourceGroupName: pulumi.String("resRg"),
			ResourceName:      pulumi.String("myDigitalTwinsService"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
