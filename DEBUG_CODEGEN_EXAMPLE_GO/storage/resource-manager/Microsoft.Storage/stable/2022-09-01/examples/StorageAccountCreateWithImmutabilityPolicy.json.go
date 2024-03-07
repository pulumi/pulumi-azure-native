package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewStorageAccount(ctx, "storageAccount", &storage.StorageAccountArgs{
			AccountName: pulumi.String("sto4445"),
			ExtendedLocation: &storage.ExtendedLocationArgs{
				Name: pulumi.String("losangeles001"),
				Type: pulumi.String("EdgeZone"),
			},
			ImmutableStorageWithVersioning: storage.ImmutableStorageAccountResponse{
				Enabled: pulumi.Bool(true),
				ImmutabilityPolicy: &storage.AccountImmutabilityPolicyPropertiesArgs{
					AllowProtectedAppendWrites:            pulumi.Bool(true),
					ImmutabilityPeriodSinceCreationInDays: pulumi.Int(15),
					State:                                 pulumi.String("Unlocked"),
				},
			},
			Kind:              pulumi.String("Storage"),
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("res9101"),
			Sku: &storage.SkuArgs{
				Name: pulumi.String("Standard_GRS"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
