package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewStorageInsightConfig(ctx, "storageInsightConfig", &operationalinsights.StorageInsightConfigArgs{
			Containers: pulumi.StringArray{
				pulumi.String("wad-iis-logfiles"),
			},
			ResourceGroupName: pulumi.String("OIAutoRest5123"),
			StorageAccount: &operationalinsights.StorageAccountArgs{
				Id:  pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/OIAutoRest6987/providers/microsoft.storage/storageaccounts/AzTestFakeSA9945"),
				Key: pulumi.String("1234"),
			},
			StorageInsightName: pulumi.String("AzTestSI1110"),
			Tables: pulumi.StringArray{
				pulumi.String("WADWindowsEventLogsTable"),
				pulumi.String("LinuxSyslogVer2v0"),
			},
			WorkspaceName: pulumi.String("aztest5048"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
