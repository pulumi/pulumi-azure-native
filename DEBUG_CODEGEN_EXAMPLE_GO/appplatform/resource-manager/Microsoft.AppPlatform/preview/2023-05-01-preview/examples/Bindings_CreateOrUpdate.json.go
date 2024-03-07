package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewBinding(ctx, "binding", &appplatform.BindingArgs{
			AppName:     pulumi.String("myapp"),
			BindingName: pulumi.String("mybinding"),
			Properties: &appplatform.BindingResourcePropertiesArgs{
				BindingParameters: pulumi.StringMap{
					"apiType":      pulumi.String("SQL"),
					"databaseName": pulumi.String("db1"),
				},
				Key:        pulumi.String("xxxx"),
				ResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DocumentDB/databaseAccounts/my-cosmosdb-1"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
