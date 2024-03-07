package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewServerAdvisor(ctx, "serverAdvisor", &sql.ServerAdvisorArgs{
			AdvisorName:       pulumi.String("CreateIndex"),
			AutoExecuteStatus: sql.AutoExecuteStatusDisabled,
			ResourceGroupName: pulumi.String("workloadinsight-demos"),
			ServerName:        pulumi.String("misosisvr"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
