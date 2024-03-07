package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewVariable(ctx, "variable", &authorization.VariableArgs{
			Columns: []authorization.PolicyVariableColumnArgs{
				{
					ColumnName: pulumi.String("TestColumn"),
				},
			},
			VariableName: pulumi.String("DemoTestVariable"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
