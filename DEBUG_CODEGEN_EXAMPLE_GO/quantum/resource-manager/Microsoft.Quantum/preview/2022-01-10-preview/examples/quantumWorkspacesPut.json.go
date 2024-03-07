package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/quantum/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quantum.NewWorkspace(ctx, "workspace", &quantum.WorkspaceArgs{
			Location: pulumi.String("West US"),
			Providers: quantum.ProviderArray{
				&quantum.ProviderArgs{
					ProviderId:  pulumi.String("Honeywell"),
					ProviderSku: pulumi.String("Basic"),
				},
				&quantum.ProviderArgs{
					ProviderId:  pulumi.String("IonQ"),
					ProviderSku: pulumi.String("Basic"),
				},
				&quantum.ProviderArgs{
					ProviderId:  pulumi.String("OneQBit"),
					ProviderSku: pulumi.String("Basic"),
				},
			},
			ResourceGroupName: pulumi.String("quantumResourcegroup"),
			StorageAccount:    pulumi.String("/subscriptions/1C4B2828-7D49-494F-933D-061373BE28C2/resourceGroups/quantumResourcegroup/providers/Microsoft.Storage/storageAccounts/testStorageAccount"),
			WorkspaceName:     pulumi.String("quantumworkspace1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
