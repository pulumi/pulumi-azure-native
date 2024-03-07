package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewManagedNetworkSettingsRule(ctx, "managedNetworkSettingsRule", &machinelearningservices.ManagedNetworkSettingsRuleArgs{
			Properties: machinelearningservices.FqdnOutboundRule{
				Category:    "UserDefined",
				Destination: "some_string",
				Status:      "Active",
				Type:        "FQDN",
			},
			ResourceGroupName: pulumi.String("test-rg"),
			RuleName:          pulumi.String("some_string"),
			WorkspaceName:     pulumi.String("aml-workspace-name"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
