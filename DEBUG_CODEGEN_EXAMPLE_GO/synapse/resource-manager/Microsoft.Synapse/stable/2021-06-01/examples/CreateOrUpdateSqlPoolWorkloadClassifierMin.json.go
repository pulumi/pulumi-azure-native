package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewSqlPoolWorkloadClassifier(ctx, "sqlPoolWorkloadClassifier", &synapse.SqlPoolWorkloadClassifierArgs{
			MemberName:             pulumi.String("dbo"),
			ResourceGroupName:      pulumi.String("sqlcrudtest-6852"),
			SqlPoolName:            pulumi.String("sqlcrudtest-9187"),
			WorkloadClassifierName: pulumi.String("wlm_workloadclassifier"),
			WorkloadGroupName:      pulumi.String("wlm_workloadgroup"),
			WorkspaceName:          pulumi.String("sqlcrudtest-2080"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
