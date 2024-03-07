package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewCustomEntityStoreAssignment(ctx, "customEntityStoreAssignment", &security.CustomEntityStoreAssignmentArgs{
			CustomEntityStoreAssignmentName: pulumi.String("33e7cc6e-a139-4723-a0e5-76993aee0771"),
			Principal:                       pulumi.String("aaduser=f3923a3e-ad57-4752-b1a9-fbf3c8e5e082;72f988bf-86f1-41af-91ab-2d7cd011db47"),
			ResourceGroupName:               pulumi.String("TestResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
