package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewAlertsSuppressionRule(ctx, "alertsSuppressionRule", &security.AlertsSuppressionRuleArgs{
			AlertType:                 pulumi.String("IpAnomaly"),
			AlertsSuppressionRuleName: pulumi.String("dismissIpAnomalyAlerts"),
			Comment:                   pulumi.String("Test VM"),
			ExpirationDateUtc:         pulumi.String("2019-12-01T19:50:47.083633Z"),
			Reason:                    pulumi.String("FalsePositive"),
			State:                     pulumi.String("Enabled"),
			SuppressionAlertsScope: &security.SuppressionAlertsScopeArgs{
				AllOf: security.ScopeElementArray{
					&security.ScopeElementArgs{
						Field: pulumi.String("entities.ip.address"),
					},
					&security.ScopeElementArgs{
						Field: pulumi.String("entities.process.commandline"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
