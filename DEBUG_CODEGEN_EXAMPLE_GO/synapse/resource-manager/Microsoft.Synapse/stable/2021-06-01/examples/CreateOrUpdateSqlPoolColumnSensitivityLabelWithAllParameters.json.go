package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewSqlPoolSensitivityLabel(ctx, "sqlPoolSensitivityLabel", &synapse.SqlPoolSensitivityLabelArgs{
			ColumnName:             pulumi.String("myColumn"),
			InformationType:        pulumi.String("PhoneNumber"),
			InformationTypeId:      pulumi.String("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
			LabelId:                pulumi.String("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
			LabelName:              pulumi.String("PII"),
			ResourceGroupName:      pulumi.String("myRG"),
			SchemaName:             pulumi.String("dbo"),
			SensitivityLabelSource: pulumi.String("current"),
			SqlPoolName:            pulumi.String("myDatabase"),
			TableName:              pulumi.String("myTable"),
			WorkspaceName:          pulumi.String("myServer"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
