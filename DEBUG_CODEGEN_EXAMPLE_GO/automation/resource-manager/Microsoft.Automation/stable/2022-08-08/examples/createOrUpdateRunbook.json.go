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
			Location:              pulumi.String("East US 2"),
			LogActivityTrace:      pulumi.Int(1),
			LogProgress:           pulumi.Bool(true),
			LogVerbose:            pulumi.Bool(false),
			Name:                  pulumi.String("Get-AzureVMTutorial"),
			PublishContentLink: automation.ContentLinkResponse{
				ContentHash: &automation.ContentHashArgs{
					Algorithm: pulumi.String("SHA256"),
					Value:     pulumi.String("115775B8FF2BE672D8A946BD0B489918C724DDE15A440373CA54461D53010A80"),
				},
				Uri: pulumi.String("https://raw.githubusercontent.com/Azure/azure-quickstart-templates/master/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1"),
			},
			ResourceGroupName: pulumi.String("rg"),
			RunbookName:       pulumi.String("Get-AzureVMTutorial"),
			RunbookType:       pulumi.String("PowerShellWorkflow"),
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
