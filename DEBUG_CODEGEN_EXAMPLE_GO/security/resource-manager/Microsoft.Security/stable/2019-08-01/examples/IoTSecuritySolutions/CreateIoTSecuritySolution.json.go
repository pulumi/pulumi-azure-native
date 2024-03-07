package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewIotSecuritySolution(ctx, "iotSecuritySolution", &security.IotSecuritySolutionArgs{
			DisabledDataSources: pulumi.StringArray{},
			DisplayName:         pulumi.String("Solution Default"),
			Export:              pulumi.StringArray{},
			IotHubs: pulumi.StringArray{
				pulumi.String("/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub"),
			},
			Location: pulumi.String("East Us"),
			RecommendationsConfiguration: []security.RecommendationConfigurationPropertiesArgs{
				{
					RecommendationType: pulumi.String("IoT_OpenPorts"),
					Status:             pulumi.String("Disabled"),
				},
				{
					RecommendationType: pulumi.String("IoT_SharedCredentials"),
					Status:             pulumi.String("Disabled"),
				},
			},
			ResourceGroupName:       pulumi.String("MyGroup"),
			SolutionName:            pulumi.String("default"),
			Status:                  pulumi.String("Enabled"),
			Tags:                    nil,
			UnmaskedIpLoggingStatus: pulumi.String("Enabled"),
			UserDefinedResources: &security.UserDefinedResourcesPropertiesArgs{
				Query: pulumi.String("where type != \"microsoft.devices/iothubs\" | where name contains \"iot\""),
				QuerySubscriptions: pulumi.StringArray{
					pulumi.String("075423e9-7d33-4166-8bdf-3920b04e3735"),
				},
			},
			Workspace: pulumi.String("/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
