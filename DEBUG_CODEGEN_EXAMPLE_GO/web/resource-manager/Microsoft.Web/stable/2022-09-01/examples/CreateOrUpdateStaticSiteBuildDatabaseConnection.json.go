package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewStaticSiteBuildDatabaseConnection(ctx, "staticSiteBuildDatabaseConnection", &web.StaticSiteBuildDatabaseConnectionArgs{
			ConnectionIdentity:     pulumi.String("SystemAssigned"),
			ConnectionString:       pulumi.String("AccountEndpoint=https://exampleDatabaseName.documents.azure.com:443/;Database=mydb;"),
			DatabaseConnectionName: pulumi.String("default"),
			EnvironmentName:        pulumi.String("default"),
			Name:                   pulumi.String("testStaticSite0"),
			Region:                 pulumi.String("West US 2"),
			ResourceGroupName:      pulumi.String("rg"),
			ResourceId:             pulumi.String("/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/databaseRG/providers/Microsoft.DocumentDB/databaseAccounts/exampleDatabaseName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
