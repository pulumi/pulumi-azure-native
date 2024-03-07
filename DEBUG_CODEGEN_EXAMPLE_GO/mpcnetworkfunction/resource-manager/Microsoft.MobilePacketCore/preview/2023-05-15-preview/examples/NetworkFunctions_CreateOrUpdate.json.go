package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilepacketcore/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilepacketcore.NewNetworkFunction(ctx, "networkFunction", &mobilepacketcore.NetworkFunctionArgs{
			Capacity:                           pulumi.Int(100000),
			DeploymentNotes:                    pulumi.String("string"),
			Location:                           pulumi.String("eastus"),
			NetworkFunctionAdministrativeState: pulumi.String("Commissioned"),
			NetworkFunctionName:                pulumi.String("nf1"),
			NetworkFunctionType:                pulumi.String("SMF"),
			ResourceGroupName:                  pulumi.String("rg1"),
			Sku:                                pulumi.String("NexusProduction"),
			UserDescription:                    pulumi.String("string"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
