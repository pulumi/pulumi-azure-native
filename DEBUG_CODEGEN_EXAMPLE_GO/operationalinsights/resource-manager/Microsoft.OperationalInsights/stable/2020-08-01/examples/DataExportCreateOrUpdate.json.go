package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewDataExport(ctx, "dataExport", &operationalinsights.DataExportArgs{
			DataExportName:    pulumi.String("export1"),
			ResourceGroupName: pulumi.String("RgTest1"),
			ResourceId:        pulumi.String("/subscriptions/192b9f85-a39a-4276-b96d-d5cd351703f9/resourceGroups/OIAutoRest1234/providers/Microsoft.EventHub/namespaces/test"),
			TableNames: pulumi.StringArray{
				pulumi.String("Heartbeat"),
			},
			WorkspaceName: pulumi.String("DeWnTest1234"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
