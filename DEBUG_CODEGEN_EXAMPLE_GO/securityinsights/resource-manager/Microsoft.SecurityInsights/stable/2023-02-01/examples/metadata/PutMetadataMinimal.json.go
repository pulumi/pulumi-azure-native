package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewMetadata(ctx, "metadata", &securityinsights.MetadataArgs{
			ContentId:         pulumi.String("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
			Kind:              pulumi.String("AnalyticsRule"),
			MetadataName:      pulumi.String("metadataName"),
			ParentId:          pulumi.String("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
			ResourceGroupName: pulumi.String("myRg"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
