package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewConnection(ctx, "connection", &automation.ConnectionArgs{
			AutomationAccountName: pulumi.String("myAutomationAccount28"),
			ConnectionName:        pulumi.String("mysConnection"),
			ConnectionType: &automation.ConnectionTypeAssociationPropertyArgs{
				Name: pulumi.String("Azure"),
			},
			Description: pulumi.String("my description goes here"),
			FieldDefinitionValues: pulumi.StringMap{
				"AutomationCertificateName": pulumi.String("mysCertificateName"),
				"SubscriptionID":            pulumi.String("subid"),
			},
			Name:              pulumi.String("mysConnection"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
