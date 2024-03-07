package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewMaintenanceConfiguration(ctx, "maintenanceConfiguration", &containerservice.MaintenanceConfigurationArgs{
			ConfigName: pulumi.String("default"),
			NotAllowedTime: containerservice.TimeSpanArray{
				&containerservice.TimeSpanArgs{
					End:   pulumi.String("2020-11-30T12:00:00Z"),
					Start: pulumi.String("2020-11-26T03:00:00Z"),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("clustername1"),
			TimeInWeek: containerservice.TimeInWeekArray{
				&containerservice.TimeInWeekArgs{
					Day: pulumi.String("Monday"),
					HourSlots: pulumi.IntArray{
						pulumi.Int(1),
						pulumi.Int(2),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
