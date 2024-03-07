package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/delegatednetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := delegatednetwork.NewControllerDetails(ctx, "controllerDetails", &delegatednetwork.ControllerDetailsArgs{
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("TestRG"),
			ResourceName:      pulumi.String("testcontroller"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
