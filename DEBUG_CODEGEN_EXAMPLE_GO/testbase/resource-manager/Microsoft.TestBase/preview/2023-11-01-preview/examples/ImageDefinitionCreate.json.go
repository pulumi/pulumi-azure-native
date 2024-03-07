package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/testbase/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := testbase.NewImageDefinition(ctx, "imageDefinition", &testbase.ImageDefinitionArgs{
			Architecture:        pulumi.String("x64"),
			ImageDefinitionName: pulumi.String("contoso-image-def"),
			OsState:             pulumi.String("Generalized"),
			ResourceGroupName:   pulumi.String("contoso-rg1"),
			SecurityType:        pulumi.String("Standard"),
			TestBaseAccountName: pulumi.String("contoso-testBaseAccount1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
