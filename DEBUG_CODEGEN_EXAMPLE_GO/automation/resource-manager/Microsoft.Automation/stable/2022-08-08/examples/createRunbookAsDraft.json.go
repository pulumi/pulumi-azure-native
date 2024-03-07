package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewRunbook(ctx, "runbook", &automation.RunbookArgs{
			AutomationAccountName: pulumi.String("ContoseAutomationAccount"),
			Description:           pulumi.String("Description of the Runbook"),
			Draft:                 nil,
			Location:              pulumi.String("East US 2"),
			LogProgress:           pulumi.Bool(false),
			LogVerbose:            pulumi.Bool(false),
			Name:                  pulumi.String("Get-AzureVMTutorial"),
			ResourceGroupName:     pulumi.String("rg"),
			RunbookName:           pulumi.String("Get-AzureVMTutorial"),
			RunbookType:           pulumi.String("PowerShellWorkflow"),
			Tags: pulumi.StringMap{
				"tag01": pulumi.String("value01"),
				"tag02": pulumi.String("value02"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
