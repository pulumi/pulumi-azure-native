package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managedidentity/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managedidentity.NewUserAssignedIdentity(ctx, "userAssignedIdentity", &managedidentity.UserAssignedIdentityArgs{
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("rgName"),
			ResourceName:      pulumi.String("resourceName"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
				"key2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
