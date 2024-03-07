package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewProtectionContainer(ctx, "protectionContainer", &recoveryservices.ProtectionContainerArgs{
			ContainerName: pulumi.String("StorageContainer;Storage;SwaggerTestRg;swaggertestsa"),
			FabricName:    pulumi.String("Azure"),
			Properties: recoveryservices.AzureStorageContainer{
				AcquireStorageAccountLock: "Acquire",
				BackupManagementType:      "AzureStorage",
				ContainerType:             "StorageContainer",
				FriendlyName:              "swaggertestsa",
				SourceResourceId:          "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsa",
			},
			ResourceGroupName: pulumi.String("SwaggerTestRg"),
			VaultName:         pulumi.String("swaggertestvault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
