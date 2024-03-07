package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicenetworking/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicenetworking.NewFrontendsInterface(ctx, "frontendsInterface", &servicenetworking.FrontendsInterfaceArgs{
			FrontendName:          pulumi.String("fe1"),
			Location:              pulumi.String("NorthCentralUS"),
			ResourceGroupName:     pulumi.String("rg1"),
			TrafficControllerName: pulumi.String("tc1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
