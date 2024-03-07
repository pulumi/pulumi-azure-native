package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/fluidrelay/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fluidrelay.NewFluidRelayServer(ctx, "fluidRelayServer", &fluidrelay.FluidRelayServerArgs{
			FluidRelayServerName: pulumi.String("myFluidRelayServer"),
			Identity: &fluidrelay.IdentityArgs{
				Type: fluidrelay.ResourceIdentityTypeSystemAssigned,
			},
			Location:      pulumi.String("west-us"),
			ResourceGroup: pulumi.String("myResourceGroup"),
			Storagesku:    pulumi.String("basic"),
			Tags: pulumi.StringMap{
				"Category": pulumi.String("sales"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
