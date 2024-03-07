package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewRuntimeEnvironment(ctx, "runtimeEnvironment", &automation.RuntimeEnvironmentArgs{
			AutomationAccountName: pulumi.String("myAutomationAccount9"),
			DefaultPackages: pulumi.StringMap{
				"Az": pulumi.String("8.3.0"),
			},
			Language:               pulumi.String("PowerShell"),
			Location:               pulumi.String("East US 2"),
			ResourceGroupName:      pulumi.String("rg"),
			RuntimeEnvironmentName: pulumi.String("myRuntimeEnvironmentName"),
			Version:                pulumi.String("7.1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
