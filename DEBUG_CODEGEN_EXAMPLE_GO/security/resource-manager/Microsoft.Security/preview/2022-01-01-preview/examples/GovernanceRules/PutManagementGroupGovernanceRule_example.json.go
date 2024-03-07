package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewGovernanceRule(ctx, "governanceRule", &security.GovernanceRuleArgs{
			Description: pulumi.String("A rule for a management group"),
			DisplayName: pulumi.String("Management group rule"),
			ExcludedScopes: pulumi.StringArray{
				pulumi.String("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
			},
			GovernanceEmailNotification: &security.GovernanceRuleEmailNotificationArgs{
				DisableManagerEmailNotification: pulumi.Bool(true),
				DisableOwnerEmailNotification:   pulumi.Bool(false),
			},
			IsDisabled:    pulumi.Bool(false),
			IsGracePeriod: pulumi.Bool(true),
			OwnerSource: &security.GovernanceRuleOwnerSourceArgs{
				Type:  pulumi.String("Manually"),
				Value: pulumi.String("user@contoso.com"),
			},
			RemediationTimeframe: pulumi.String("7.00:00:00"),
			RuleId:               pulumi.String("ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
			RulePriority:         pulumi.Int(200),
			RuleType:             pulumi.String("Integrated"),
			Scope:                pulumi.String("providers/Microsoft.Management/managementGroups/contoso"),
			SourceResourceType:   pulumi.String("Assessments"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
