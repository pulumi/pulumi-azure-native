package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewJavaComponent(ctx, "javaComponent", &app.JavaComponentArgs{
			ComponentType: pulumi.String("SpringBootAdmin"),
			Configurations: app.JavaComponentConfigurationPropertyArray{
				&app.JavaComponentConfigurationPropertyArgs{
					PropertyName: pulumi.String("spring.boot.admin.ui.enable-toasts"),
					Value:        pulumi.String("true"),
				},
				&app.JavaComponentConfigurationPropertyArgs{
					PropertyName: pulumi.String("spring.boot.admin.monitor.status-interval"),
					Value:        pulumi.String("10000ms"),
				},
			},
			EnvironmentName:   pulumi.String("myenvironment"),
			Name:              pulumi.String("myjavacomponent"),
			ResourceGroupName: pulumi.String("examplerg"),
			ServiceBinds: app.JavaComponentServiceBindArray{
				&app.JavaComponentServiceBindArgs{
					Name:      pulumi.String("yellowcat"),
					ServiceId: pulumi.String("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/javaComponents/yellowcat"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
