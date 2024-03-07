package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewWorkloadClassifier(ctx, "workloadClassifier", &sql.WorkloadClassifierArgs{
			DatabaseName:           pulumi.String("testdb"),
			MemberName:             pulumi.String("dbo"),
			ResourceGroupName:      pulumi.String("Default-SQL-SouthEastAsia"),
			ServerName:             pulumi.String("testsvr"),
			WorkloadClassifierName: pulumi.String("wlm_workloadclassifier"),
			WorkloadGroupName:      pulumi.String("wlm_workloadgroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
