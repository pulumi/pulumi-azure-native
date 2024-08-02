package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	web "github.com/pulumi/pulumi-azure-native-sdk/web/v2/v20220901"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		rg, err := resources.NewResourceGroup(ctx, "rg", nil)
		if err != nil {
			return err
		}

		sta, err := storage.NewStorageAccount(ctx, "storage", &storage.StorageAccountArgs{
			ResourceGroupName: rg.Name,
			Kind:              pulumi.String("StorageV2"),
			Sku: &storage.SkuArgs{
				Name: pulumi.String("Standard_LRS"),
			},
		})
		if err != nil {
			return err
		}

		plan, err := web.NewAppServicePlan(ctx, "plan", &web.AppServicePlanArgs{
			ResourceGroupName: rg.Name,
			Sku: &web.SkuDescriptionArgs{
				Name: pulumi.String("B1"),
				Tier: pulumi.String("Basic"),
			},
		})
		if err != nil {
			return err
		}

		storageAccountKeys := storage.ListStorageAccountKeysOutput(ctx, storage.ListStorageAccountKeysOutputArgs{
			ResourceGroupName: rg.Name,
			AccountName:       sta.Name,
		})
		primaryStorageKey := storageAccountKeys.ApplyT(func(keys storage.ListStorageAccountKeysResult) string {
			return keys.Keys[0].Value
		}).(pulumi.StringOutput)

		app, err := web.NewWebApp(ctx, "app", &web.WebAppArgs{
			ResourceGroupName:     rg.Name,
			HttpsOnly:             pulumi.Bool(true),
			ServerFarmId:          plan.ID(),
			ClientAffinityEnabled: pulumi.Bool(true),
			SiteConfig: &web.SiteConfigArgs{
				AppSettings: web.NameValuePairArray{
					&web.NameValuePairArgs{
						Name:  pulumi.String("AzureWebJobsStorage"),
						Value: pulumi.Sprintf("DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s", sta.Name, primaryStorageKey),
					},
					&web.NameValuePairArgs{
						Name:  pulumi.String("FUNCTIONS_EXTENSION_VERSION"),
						Value: pulumi.String("~3"),
					},
					&web.NameValuePairArgs{
						Name:  pulumi.String("FUNCTIONS_WORKER_RUNTIME"),
						Value: pulumi.String("dotnet"),
					},
					&web.NameValuePairArgs{
						Name:  pulumi.String("WEBSITE_CONTENTAZUREFILECONNECTIONSTRING"),
						Value: pulumi.Sprintf("DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s", sta.Name, primaryStorageKey),
					},
					&web.NameValuePairArgs{
						Name:  pulumi.String("WEBSITE_CONTENTSHARE"),
						Value: pulumi.String("devpod"),
					},
				},
			},
			// SiteConfig is modified outside of this resource via the WebApp* resources.
		}, pulumi.IgnoreChanges([]string{"siteConfig", "siteConfig.*"}))
		if err != nil {
			return err
		}
		ctx.Export("endpoint", pulumi.Sprintf("https://%s.azurewebsites.net", app.Name))

		_, err = web.NewWebAppScmAllowed(ctx, "scm", &web.WebAppScmAllowedArgs{
			Allow:             pulumi.Bool(false),
			ResourceGroupName: rg.Name,
			Name:              app.Name,
		})
		if err != nil {
			return err
		}
		_, err = web.NewWebAppFtpAllowed(ctx, "ftp", &web.WebAppFtpAllowedArgs{
			Allow:             pulumi.Bool(false),
			ResourceGroupName: rg.Name,
			Name:              app.Name,
		})
		if err != nil {
			return err
		}

		return nil
	})
}
