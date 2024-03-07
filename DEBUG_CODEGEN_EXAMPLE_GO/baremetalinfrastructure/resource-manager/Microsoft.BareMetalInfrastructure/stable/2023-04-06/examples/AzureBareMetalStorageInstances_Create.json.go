package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/baremetalinfrastructure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := baremetalinfrastructure.NewAzureBareMetalStorageInstance(ctx, "azureBareMetalStorageInstance", &baremetalinfrastructure.AzureBareMetalStorageInstanceArgs{
			AzureBareMetalStorageInstanceName:             pulumi.String("myAzureBareMetalStorageInstance"),
			AzureBareMetalStorageInstanceUniqueIdentifier: pulumi.String("23415635-4d7e-41dc-9598-8194f22c24e9"),
			Location:          pulumi.String("westus2"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			StorageProperties: &baremetalinfrastructure.StoragePropertiesArgs{
				Generation:        pulumi.String("Gen4"),
				HardwareType:      pulumi.String("NetApp"),
				OfferingType:      pulumi.String("EPIC"),
				ProvisioningState: pulumi.String("Succeeded"),
				StorageBillingProperties: &baremetalinfrastructure.StorageBillingPropertiesArgs{
					AzureBareMetalStorageInstanceSize: pulumi.String(""),
					BillingMode:                       pulumi.String("PAYG"),
				},
				StorageType:  pulumi.String("FC"),
				WorkloadType: pulumi.String("ODB"),
			},
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
