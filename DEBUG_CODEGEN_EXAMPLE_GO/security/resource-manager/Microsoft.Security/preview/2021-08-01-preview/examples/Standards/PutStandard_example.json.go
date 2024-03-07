package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewStandard(ctx, "standard", &security.StandardArgs{
			Category: pulumi.String("SecurityCenter"),
			Components: security.StandardComponentPropertiesArray{
				&security.StandardComponentPropertiesArgs{
					Key: pulumi.String("1195afff-c881-495e-9bc5-1486211ae03f"),
				},
				&security.StandardComponentPropertiesArgs{
					Key: pulumi.String("dbd0cb49-b563-45e7-9724-889e799fa648"),
				},
			},
			Description:       pulumi.String("description of Azure Test Security Standard 1"),
			DisplayName:       pulumi.String("Azure Test Security Standard 1"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			StandardId:        pulumi.String("8bb8be0a-6010-4789-812f-e4d661c4ed0e"),
			SupportedClouds: security.StandardSupportedCloudsArray{
				security.StandardSupportedCloudsGCP,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
