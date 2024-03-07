package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewAction(ctx, "action", &securityinsights.ActionArgs{
			ActionId:           pulumi.String("912bec42-cb66-4c03-ac63-1761b6898c3e"),
			LogicAppResourceId: pulumi.String("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.Logic/workflows/MyAlerts"),
			ResourceGroupName:  pulumi.String("myRg"),
			RuleId:             pulumi.String("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
			TriggerUri:         pulumi.String("https://prod-31.northcentralus.logic.azure.com:443/workflows/cd3765391efd48549fd7681ded1d48d7/triggers/manual/paths/invoke?api-version=2016-10-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=signature"),
			WorkspaceName:      pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
