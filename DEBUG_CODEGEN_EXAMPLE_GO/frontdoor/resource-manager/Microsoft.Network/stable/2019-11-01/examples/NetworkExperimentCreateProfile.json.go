package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNetworkExperimentProfile(ctx, "networkExperimentProfile", &network.NetworkExperimentProfileArgs{
			EnabledState:      pulumi.String("Enabled"),
			Location:          pulumi.String("WestUs"),
			ProfileName:       pulumi.String("MyProfile"),
			ResourceGroupName: pulumi.String("MyResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
