package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewScheduledAlertRule(ctx, "scheduledAlertRule", &securityinsights.ScheduledAlertRuleArgs{
			ResourceGroupName: pulumi.String("myRg"),
			RuleId:            pulumi.String("myFirstFusionRule"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
