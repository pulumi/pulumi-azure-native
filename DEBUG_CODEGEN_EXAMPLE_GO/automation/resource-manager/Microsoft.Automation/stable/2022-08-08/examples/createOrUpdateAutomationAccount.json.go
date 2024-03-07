package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewAutomationAccount(ctx, "automationAccount", &automation.AutomationAccountArgs{
			AutomationAccountName: pulumi.String("myAutomationAccount9"),
			Location:              pulumi.String("East US 2"),
			Name:                  pulumi.String("myAutomationAccount9"),
			ResourceGroupName:     pulumi.String("rg"),
			Sku: &automation.SkuArgs{
				Name: pulumi.String("Free"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
