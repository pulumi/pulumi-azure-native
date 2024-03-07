package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/purview/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := purview.NewAccount(ctx, "account", &purview.AccountArgs{
			AccountName:                         pulumi.String("account1"),
			Location:                            pulumi.String("West US 2"),
			ManagedResourceGroupName:            pulumi.String("custom-rgname"),
			ManagedResourcesPublicNetworkAccess: pulumi.String("Enabled"),
			ResourceGroupName:                   pulumi.String("SampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
