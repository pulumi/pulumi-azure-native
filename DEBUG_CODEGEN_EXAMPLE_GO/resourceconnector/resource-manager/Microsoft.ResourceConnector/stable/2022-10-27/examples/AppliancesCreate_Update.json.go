package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resourceconnector/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resourceconnector.NewAppliance(ctx, "appliance", &resourceconnector.ApplianceArgs{
			Distro: pulumi.String("AKSEdge"),
			InfrastructureConfig: &resourceconnector.AppliancePropertiesInfrastructureConfigArgs{
				Provider: pulumi.String("VMWare"),
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("testresourcegroup"),
			ResourceName:      pulumi.String("appliance01"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
