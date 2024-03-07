package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewAutomation(ctx, "automation", &security.AutomationArgs{
			Actions: pulumi.Array{
				security.AutomationActionLogicApp{
					ActionType:         "LogicApp",
					LogicAppResourceId: "/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1",
					Uri:                "https://exampleTriggerUri1.com",
				},
			},
			AutomationName:    pulumi.String("exampleAutomation"),
			Description:       pulumi.String("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment"),
			IsEnabled:         pulumi.Bool(true),
			Location:          pulumi.String("Central US"),
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
			Scopes: security.AutomationScopeArray{
				&security.AutomationScopeArgs{
					Description: pulumi.String("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
					ScopePath:   pulumi.String("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
				},
			},
			Sources: security.AutomationSourceArray{
				&security.AutomationSourceArgs{
					EventSource: pulumi.String("Assessments"),
				},
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
