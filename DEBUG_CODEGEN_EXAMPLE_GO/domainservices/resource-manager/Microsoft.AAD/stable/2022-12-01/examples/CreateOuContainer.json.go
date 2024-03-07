package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/aad/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aad.NewOuContainer(ctx, "ouContainer", &aad.OuContainerArgs{
			AccountName:       pulumi.String("AccountName1"),
			DomainServiceName: pulumi.String("OuContainer.com"),
			OuContainerName:   pulumi.String("OuContainer1"),
			Password:          pulumi.String("<password>"),
			ResourceGroupName: pulumi.String("OuContainerResourceGroup"),
			Spn:               pulumi.String("Spn1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
