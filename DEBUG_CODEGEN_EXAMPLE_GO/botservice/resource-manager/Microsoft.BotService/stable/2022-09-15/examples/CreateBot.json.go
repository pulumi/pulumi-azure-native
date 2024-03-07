package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/botservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := botservice.NewBot(ctx, "bot", &botservice.BotArgs{
			Kind:     pulumi.String("sdk"),
			Location: pulumi.String("West US"),
			Properties: &botservice.BotPropertiesArgs{
				CmekKeyVaultUrl:                   pulumi.String("https://myCmekKey"),
				Description:                       pulumi.String("The description of the bot"),
				DeveloperAppInsightKey:            pulumi.String("appinsightskey"),
				DeveloperAppInsightsApiKey:        pulumi.String("appinsightsapikey"),
				DeveloperAppInsightsApplicationId: pulumi.String("appinsightsappid"),
				DisableLocalAuth:                  pulumi.Bool(true),
				DisplayName:                       pulumi.String("The Name of the bot"),
				Endpoint:                          pulumi.String("http://mybot.coffee"),
				IconUrl:                           pulumi.String("http://myicon"),
				IsCmekEnabled:                     pulumi.Bool(true),
				LuisAppIds: pulumi.StringArray{
					pulumi.String("luisappid1"),
					pulumi.String("luisappid2"),
				},
				LuisKey:                     pulumi.String("luiskey"),
				MsaAppId:                    pulumi.String("exampleappid"),
				MsaAppMSIResourceId:         pulumi.String("/subscriptions/foo/resourcegroups/bar/providers/microsoft.managedidentity/userassignedidentities/sampleId"),
				MsaAppTenantId:              pulumi.String("exampleapptenantid"),
				MsaAppType:                  pulumi.String("UserAssignedMSI"),
				PublicNetworkAccess:         pulumi.String("Enabled"),
				SchemaTransformationVersion: pulumi.String("1.0"),
			},
			ResourceGroupName: pulumi.String("OneResourceGroupName"),
			ResourceName:      pulumi.String("samplebotname"),
			Sku: &botservice.SkuArgs{
				Name: pulumi.String("S1"),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
