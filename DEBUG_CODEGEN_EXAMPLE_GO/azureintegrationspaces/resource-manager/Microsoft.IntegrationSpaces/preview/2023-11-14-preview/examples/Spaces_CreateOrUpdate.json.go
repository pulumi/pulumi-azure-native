package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/integrationspaces/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := integrationspaces.NewSpace(ctx, "space", &integrationspaces.SpaceArgs{
			Description:       pulumi.String("This is the user provided description of the space resource."),
			Location:          pulumi.String("CentralUS"),
			ResourceGroupName: pulumi.String("testrg"),
			SpaceName:         pulumi.String("Space1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("Value1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
