package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicenetworking/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicenetworking.NewTrafficControllerInterface(ctx, "trafficControllerInterface", &servicenetworking.TrafficControllerInterfaceArgs{
			Location:          pulumi.String("NorthCentralUS"),
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			TrafficControllerName: pulumi.String("tc1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
