package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appconfiguration.NewKeyValue(ctx, "keyValue", &appconfiguration.KeyValueArgs{
			ConfigStoreName:   pulumi.String("contoso"),
			KeyValueName:      pulumi.String("myKey$myLabel"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("tagValue1"),
				"tag2": pulumi.String("tagValue2"),
			},
			Value: pulumi.String("myValue"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
