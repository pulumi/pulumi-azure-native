package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/maintenance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := maintenance.NewMaintenanceConfiguration(ctx, "maintenanceConfiguration", &maintenance.MaintenanceConfigurationArgs{
			Duration:           pulumi.String("05:00"),
			ExpirationDateTime: pulumi.String("9999-12-31 00:00"),
			Location:           pulumi.String("westus2"),
			MaintenanceScope:   pulumi.String("OSImage"),
			Namespace:          pulumi.String("Microsoft.Maintenance"),
			RecurEvery:         pulumi.String("Day"),
			ResourceGroupName:  pulumi.String("examplerg"),
			ResourceName:       pulumi.String("configuration1"),
			StartDateTime:      pulumi.String("2020-04-30 08:00"),
			TimeZone:           pulumi.String("Pacific Standard Time"),
			Visibility:         pulumi.String("Custom"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
