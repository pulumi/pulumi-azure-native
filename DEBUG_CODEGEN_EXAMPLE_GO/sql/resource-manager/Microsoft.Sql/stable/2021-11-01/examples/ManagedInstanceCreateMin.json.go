package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewManagedInstance(ctx, "managedInstance", &sql.ManagedInstanceArgs{
			AdministratorLogin:         pulumi.String("dummylogin"),
			AdministratorLoginPassword: pulumi.String("PLACEHOLDER"),
			LicenseType:                pulumi.String("LicenseIncluded"),
			Location:                   pulumi.String("Japan East"),
			ManagedInstanceName:        pulumi.String("testinstance"),
			ResourceGroupName:          pulumi.String("testrg"),
			Sku: &sql.SkuArgs{
				Name: pulumi.String("GP_Gen4"),
				Tier: pulumi.String("GeneralPurpose"),
			},
			StorageSizeInGB: pulumi.Int(1024),
			SubnetId:        pulumi.String("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
			VCores:          pulumi.Int(8),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
