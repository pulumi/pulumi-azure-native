package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewConnectionType(ctx, "connectionType", &automation.ConnectionTypeArgs{
			AutomationAccountName: pulumi.String("myAutomationAccount22"),
			ConnectionTypeName:    pulumi.String("myCT"),
			FieldDefinitions: automation.FieldDefinitionMap{
				"myBoolField": &automation.FieldDefinitionArgs{
					IsEncrypted: pulumi.Bool(false),
					IsOptional:  pulumi.Bool(false),
					Type:        pulumi.String("bool"),
				},
				"myStringField": &automation.FieldDefinitionArgs{
					IsEncrypted: pulumi.Bool(false),
					IsOptional:  pulumi.Bool(false),
					Type:        pulumi.String("string"),
				},
				"myStringFieldEncrypted": &automation.FieldDefinitionArgs{
					IsEncrypted: pulumi.Bool(true),
					IsOptional:  pulumi.Bool(false),
					Type:        pulumi.String("string"),
				},
			},
			IsGlobal:          pulumi.Bool(false),
			Name:              pulumi.String("myCT"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
