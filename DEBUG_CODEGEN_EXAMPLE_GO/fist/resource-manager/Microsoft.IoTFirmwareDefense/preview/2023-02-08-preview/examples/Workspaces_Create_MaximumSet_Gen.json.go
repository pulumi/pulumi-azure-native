package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotfirmwaredefense/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotfirmwaredefense.NewWorkspace(ctx, "workspace", &iotfirmwaredefense.WorkspaceArgs{
			Location:          pulumi.String("jjwbseilitjgdrhbvvkwviqj"),
			ResourceGroupName: pulumi.String("rgworkspaces"),
			Tags: pulumi.StringMap{
				"key450": pulumi.String("rzqqumbpfsbibnpirsm"),
			},
			WorkspaceName: pulumi.String("E___-3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
