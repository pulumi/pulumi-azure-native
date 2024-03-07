package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewHybridRunbookWorkerGroup(ctx, "hybridRunbookWorkerGroup", &automation.HybridRunbookWorkerGroupArgs{
			AutomationAccountName: pulumi.String("testaccount"),
			Credential: &automation.RunAsCredentialAssociationPropertyArgs{
				Name: pulumi.String("myRunAsCredentialName"),
			},
			HybridRunbookWorkerGroupName: pulumi.String("TestHybridGroup"),
			ResourceGroupName:            pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
