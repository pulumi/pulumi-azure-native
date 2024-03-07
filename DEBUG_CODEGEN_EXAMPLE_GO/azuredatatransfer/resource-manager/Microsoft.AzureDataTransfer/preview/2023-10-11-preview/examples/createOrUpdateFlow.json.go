package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azuredatatransfer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuredatatransfer.NewFlow(ctx, "flow", &azuredatatransfer.FlowArgs{
			ConnectionName: pulumi.String("testConnection"),
			FlowName:       pulumi.String("testFlow"),
			Location:       pulumi.String("East US"),
			Properties: azuredatatransfer.FlowPropertiesResponse{
				Connection: &azuredatatransfer.SelectedResourceArgs{
					Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection"),
				},
				FlowType:             pulumi.String("Blob"),
				StorageAccountName:   pulumi.String("testsa"),
				StorageContainerName: pulumi.String("testcontainer"),
			},
			ResourceGroupName: pulumi.String("testRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
