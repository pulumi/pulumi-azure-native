package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cognitiveservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognitiveservices.NewAccount(ctx, "account", &cognitiveservices.AccountArgs{
			AccountName: pulumi.String("testCreate1"),
			Identity: &cognitiveservices.IdentityArgs{
				Type: cognitiveservices.ResourceIdentityTypeSystemAssigned,
			},
			Kind:              pulumi.String("CognitiveServices"),
			Location:          pulumi.String("West US"),
			Properties:        nil,
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &cognitiveservices.SkuArgs{
				Name: pulumi.String("S0"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
