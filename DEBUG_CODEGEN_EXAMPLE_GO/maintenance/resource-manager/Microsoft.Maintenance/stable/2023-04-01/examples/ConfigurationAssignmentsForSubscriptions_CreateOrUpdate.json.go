package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/maintenance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := maintenance.NewConfigurationAssignmentsForSubscription(ctx, "configurationAssignmentsForSubscription", &maintenance.ConfigurationAssignmentsForSubscriptionArgs{
			ConfigurationAssignmentName: pulumi.String("workervmConfiguration"),
			Filter: maintenance.ConfigurationAssignmentFilterPropertiesResponse{
				Locations: pulumi.StringArray{
					pulumi.String("Japan East"),
					pulumi.String("UK South"),
				},
				ResourceGroups: pulumi.StringArray{
					pulumi.String("RG1"),
					pulumi.String("RG2"),
				},
				ResourceTypes: pulumi.StringArray{
					pulumi.String("Microsoft.HybridCompute/machines"),
					pulumi.String("Microsoft.Compute/virtualMachines"),
				},
				TagSettings: &maintenance.TagSettingsPropertiesArgs{
					FilterOperator: maintenance.TagOperatorsAny,
					Tags: pulumi.StringArrayMap{
						"tag1": pulumi.StringArray{
							pulumi.String("tag1Value1"),
							pulumi.String("tag1Value2"),
							pulumi.String("tag1Value3"),
						},
						"tag2": pulumi.StringArray{
							pulumi.String("tag2Value1"),
							pulumi.String("tag2Value2"),
							pulumi.String("tag2Value3"),
						},
					},
				},
			},
			MaintenanceConfigurationId: pulumi.String("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
