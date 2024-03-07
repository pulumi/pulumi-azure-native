package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automanage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automanage.NewConfigurationProfileHCRPAssignment(ctx, "configurationProfileHCRPAssignment", &automanage.ConfigurationProfileHCRPAssignmentArgs{
			ConfigurationProfileAssignmentName: pulumi.String("default"),
			MachineName:                        pulumi.String("myMachineName"),
			Properties: &automanage.ConfigurationProfileAssignmentPropertiesArgs{
				ConfigurationProfile: pulumi.String("/providers/Microsoft.Automanage/bestPractices/AzureBestPracticesProduction"),
			},
			ResourceGroupName: pulumi.String("myResourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
