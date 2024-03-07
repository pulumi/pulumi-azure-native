package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resources.NewAzurePowerShellScript(ctx, "azurePowerShellScript", &resources.AzurePowerShellScriptArgs{
			Arguments:           pulumi.String("-Location 'westus' -Name \"*rg2\""),
			AzPowerShellVersion: pulumi.String("1.7.0"),
			CleanupPreference:   pulumi.String("Always"),
			Kind:                pulumi.String("AzurePowerShell"),
			Location:            pulumi.String("westus"),
			ResourceGroupName:   pulumi.String("script-rg"),
			RetentionInterval:   pulumi.String("PT7D"),
			ScriptContent:       pulumi.String("Param([string]$Location,[string]$Name) $deploymentScriptOutputs['test'] = 'value' Get-AzResourceGroup -Location $Location -Name $Name"),
			ScriptName:          pulumi.String("MyDeploymentScript"),
			SupportingScriptUris: pulumi.StringArray{
				pulumi.String("https://uri1.to.supporting.script"),
				pulumi.String("https://uri2.to.supporting.script"),
			},
			Timeout: pulumi.String("PT1H"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
