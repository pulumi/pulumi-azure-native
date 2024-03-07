package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewDataSource(ctx, "dataSource", &operationalinsights.DataSourceArgs{
			DataSourceName: pulumi.String("AzTestDS774"),
			Kind:           pulumi.String("AzureActivityLog"),
			Properties: pulumi.Any{
				LinkedResourceId: "/subscriptions/00000000-0000-0000-0000-00000000000/providers/microsoft.insights/eventtypes/management",
			},
			ResourceGroupName: pulumi.String("OIAutoRest5123"),
			WorkspaceName:     pulumi.String("AzTest9724"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
