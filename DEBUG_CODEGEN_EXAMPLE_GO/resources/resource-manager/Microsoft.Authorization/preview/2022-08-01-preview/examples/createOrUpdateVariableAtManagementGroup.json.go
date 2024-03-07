package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewVariableAtManagementGroup(ctx, "variableAtManagementGroup", &authorization.VariableAtManagementGroupArgs{
			Columns: []authorization.PolicyVariableColumnArgs{
				{
					ColumnName: pulumi.String("TestColumn"),
				},
			},
			ManagementGroupId: pulumi.String("DevOrg"),
			VariableName:      pulumi.String("DemoTestVariable"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
