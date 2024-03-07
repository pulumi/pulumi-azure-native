package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/integrationspaces/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := integrationspaces.NewApplicationResource(ctx, "applicationResource", &integrationspaces.ApplicationResourceArgs{
			ApplicationName:   pulumi.String("Application1"),
			ResourceGroupName: pulumi.String("testrg"),
			ResourceId:        pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Web/sites/LogicApp1"),
			ResourceKind:      pulumi.String("LogicApp"),
			ResourceName:      pulumi.String("Resource1"),
			ResourceType:      pulumi.String("Microsoft.Web/sites"),
			SpaceName:         pulumi.String("Space1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
