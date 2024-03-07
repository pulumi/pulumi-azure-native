package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewIntegrationAccountBatchConfiguration(ctx, "integrationAccountBatchConfiguration", &logic.IntegrationAccountBatchConfigurationArgs{
			BatchConfigurationName: pulumi.String("testBatchConfiguration"),
			IntegrationAccountName: pulumi.String("testIntegrationAccount"),
			Location:               pulumi.String("westus"),
			Properties: &logic.BatchConfigurationPropertiesArgs{
				BatchGroupName: pulumi.String("DEFAULT"),
				ReleaseCriteria: &logic.BatchReleaseCriteriaArgs{
					BatchSize:    pulumi.Int(234567),
					MessageCount: pulumi.Int(10),
					Recurrence: &logic.WorkflowTriggerRecurrenceArgs{
						Frequency: pulumi.String("Minute"),
						Interval:  pulumi.Int(1),
						StartTime: pulumi.String("2017-03-24T11:43:00"),
						TimeZone:  pulumi.String("India Standard Time"),
					},
				},
			},
			ResourceGroupName: pulumi.String("testResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
