package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewDotNetComponent(ctx, "dotNetComponent", &app.DotNetComponentArgs{
			ComponentType: pulumi.String("AspireDashboard"),
			Configurations: app.DotNetComponentConfigurationPropertyArray{
				&app.DotNetComponentConfigurationPropertyArgs{
					PropertyName: pulumi.String("dashboard-theme"),
					Value:        pulumi.String("dark"),
				},
			},
			EnvironmentName:   pulumi.String("myenvironment"),
			Name:              pulumi.String("mydotnetcomponent"),
			ResourceGroupName: pulumi.String("examplerg"),
			ServiceBinds: app.DotNetComponentServiceBindArray{
				&app.DotNetComponentServiceBindArgs{
					Name:      pulumi.String("yellowcat"),
					ServiceId: pulumi.String("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/dotNetComponents/yellowcat"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
