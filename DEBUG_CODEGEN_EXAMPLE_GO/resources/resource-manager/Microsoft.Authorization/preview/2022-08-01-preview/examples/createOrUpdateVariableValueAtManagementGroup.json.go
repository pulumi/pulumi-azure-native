package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewVariableValueAtManagementGroup(ctx, "variableValueAtManagementGroup", &authorization.VariableValueAtManagementGroupArgs{
			ManagementGroupId: pulumi.String("DevOrg"),
			Values: authorization.PolicyVariableValueColumnValueArray{
				&authorization.PolicyVariableValueColumnValueArgs{
					ColumnName:  pulumi.String("StringColumn"),
					ColumnValue: pulumi.Any("SampleValue"),
				},
				&authorization.PolicyVariableValueColumnValueArgs{
					ColumnName:  pulumi.String("IntegerColumn"),
					ColumnValue: pulumi.Any(10),
				},
			},
			VariableName:      pulumi.String("DemoTestVariable"),
			VariableValueName: pulumi.String("TestValue"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
