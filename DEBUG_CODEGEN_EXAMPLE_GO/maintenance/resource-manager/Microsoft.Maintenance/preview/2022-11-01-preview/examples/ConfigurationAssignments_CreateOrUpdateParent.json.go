package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/maintenance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := maintenance.NewConfigurationAssignmentParent(ctx, "configurationAssignmentParent", &maintenance.ConfigurationAssignmentParentArgs{
			ConfigurationAssignmentName: pulumi.String("workervmPolicy"),
			MaintenanceConfigurationId:  pulumi.String("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/policy1"),
			ProviderName:                pulumi.String("Microsoft.Compute"),
			ResourceGroupName:           pulumi.String("examplerg"),
			ResourceName:                pulumi.String("smdvm1"),
			ResourceParentName:          pulumi.String("smdtest1"),
			ResourceParentType:          pulumi.String("virtualMachineScaleSets"),
			ResourceType:                pulumi.String("virtualMachines"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
