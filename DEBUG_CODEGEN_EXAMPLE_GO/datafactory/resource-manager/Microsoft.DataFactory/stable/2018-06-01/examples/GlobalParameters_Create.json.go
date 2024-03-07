package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datafactory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datafactory.NewGlobalParameter(ctx, "globalParameter", &datafactory.GlobalParameterArgs{
			FactoryName:         pulumi.String("exampleFactoryName"),
			GlobalParameterName: pulumi.String("default"),
			Properties: datafactory.GlobalParameterSpecificationMap{
				"waitTime": &datafactory.GlobalParameterSpecificationArgs{
					Type:  pulumi.String("Int"),
					Value: pulumi.Any(5),
				},
			},
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
