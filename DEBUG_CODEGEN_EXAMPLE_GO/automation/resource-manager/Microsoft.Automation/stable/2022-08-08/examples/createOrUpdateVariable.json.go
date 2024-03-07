package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewVariable(ctx, "variable", &automation.VariableArgs{
			AutomationAccountName: pulumi.String("sampleAccount9"),
			Description:           pulumi.String("my description"),
			IsEncrypted:           pulumi.Bool(false),
			Name:                  pulumi.String("sampleVariable"),
			ResourceGroupName:     pulumi.String("rg"),
			Value:                 pulumi.String("\"ComputerName.domain.com\""),
			VariableName:          pulumi.String("sampleVariable"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
