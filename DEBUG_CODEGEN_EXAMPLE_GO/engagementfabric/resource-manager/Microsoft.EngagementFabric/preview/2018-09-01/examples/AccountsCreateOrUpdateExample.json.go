package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/engagementfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := engagementfabric.NewAccount(ctx, "account", &engagementfabric.AccountArgs{
			AccountName:       pulumi.String("ExampleAccount"),
			Location:          pulumi.String("WestUS"),
			ResourceGroupName: pulumi.String("ExampleRg"),
			Sku: &engagementfabric.SKUArgs{
				Name: pulumi.String("B1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
