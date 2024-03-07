package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewMachineGroup(ctx, "machineGroup", &operationalinsights.MachineGroupArgs{
			Count:            pulumi.Int(1),
			DisplayName:      pulumi.String("Foo"),
			Kind:             pulumi.String("machineGroup"),
			MachineGroupName: pulumi.String("ccfbf4bf-dc08-4371-9e9b-00a8d875d45a"),
			Machines: []operationalinsights.MachineReferenceWithHintsArgs{
				{
					Id:   pulumi.String("/subscriptions/63BE4E24-FDF0-4E9C-9342-6A5D5A359722/resourceGroups/rg-sm/providers/Microsoft.OperationalInsights/workspaces/D6F79F14-E563-469B-84B5-9286D2803B2F/machines/m-0fe4b501-7ac9-41d7-a4e1-1591a0789519"),
					Kind: pulumi.String("ref:machinewithhints"),
				},
			},
			ResourceGroupName: pulumi.String("rg-sm"),
			WorkspaceName:     pulumi.String("D6F79F14-E563-469B-84B5-9286D2803B2F"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
