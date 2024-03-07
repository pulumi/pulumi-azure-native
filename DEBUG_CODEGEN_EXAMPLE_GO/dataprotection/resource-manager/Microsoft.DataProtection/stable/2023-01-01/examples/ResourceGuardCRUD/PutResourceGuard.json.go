package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dataprotection/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dataprotection.NewResourceGuard(ctx, "resourceGuard", &dataprotection.ResourceGuardArgs{
			Location:           pulumi.String("WestUS"),
			ResourceGroupName:  pulumi.String("SampleResourceGroup"),
			ResourceGuardsName: pulumi.String("swaggerExample"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("val1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
