package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewDotNetComponent(ctx, "dotNetComponent", &app.DotNetComponentArgs{
			ComponentType: pulumi.String("AspireDashboard"),
			Configurations: []app.DotNetComponentConfigurationPropertyArgs{
				{
					PropertyName: pulumi.String("dashboard-theme"),
					Value:        pulumi.String("dark"),
				},
			},
			EnvironmentName:   pulumi.String("myenvironment"),
			Name:              pulumi.String("mydotnetcomponent"),
			ResourceGroupName: pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
