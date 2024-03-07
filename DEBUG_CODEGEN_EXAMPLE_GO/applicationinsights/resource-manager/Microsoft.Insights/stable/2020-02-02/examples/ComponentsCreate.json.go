package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewComponent(ctx, "component", &insights.ComponentArgs{
			ApplicationType:     pulumi.String("web"),
			FlowType:            pulumi.String("Bluefield"),
			Kind:                pulumi.String("web"),
			Location:            pulumi.String("South Central US"),
			RequestSource:       pulumi.String("rest"),
			ResourceGroupName:   pulumi.String("my-resource-group"),
			ResourceName:        pulumi.String("my-component"),
			WorkspaceResourceId: pulumi.String("/subscriptions/subid/resourcegroups/my-resource-group/providers/microsoft.operationalinsights/workspaces/my-workspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
