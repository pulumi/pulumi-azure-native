package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewCredential(ctx, "credential", &automation.CredentialArgs{
			AutomationAccountName: pulumi.String("myAutomationAccount18"),
			CredentialName:        pulumi.String("myCredential"),
			Description:           pulumi.String("my description goes here"),
			Name:                  pulumi.String("myCredential"),
			Password:              pulumi.String("<password>"),
			ResourceGroupName:     pulumi.String("rg"),
			UserName:              pulumi.String("mylingaiah"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
