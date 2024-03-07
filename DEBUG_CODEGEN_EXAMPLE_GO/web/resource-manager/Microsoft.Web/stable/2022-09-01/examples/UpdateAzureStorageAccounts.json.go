package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewWebAppAzureStorageAccounts(ctx, "webAppAzureStorageAccounts", &web.WebAppAzureStorageAccountsArgs{
			Name: pulumi.String("sitef6141"),
			Properties: web.AzureStorageInfoValueMap{
				"account1": &web.AzureStorageInfoValueArgs{
					AccessKey:   pulumi.String("26515^%@#*"),
					AccountName: pulumi.String("testsa"),
					MountPath:   pulumi.String("/mounts/a/files"),
					ShareName:   pulumi.String("web"),
					Type:        web.AzureStorageTypeAzureFiles,
				},
			},
			ResourceGroupName: pulumi.String("testrg123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
