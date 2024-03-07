package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewJobTargetGroup(ctx, "jobTargetGroup", &sql.JobTargetGroupArgs{
			JobAgentName:      pulumi.String("agent1"),
			Members:           sql.JobTargetArray{},
			ResourceGroupName: pulumi.String("group1"),
			ServerName:        pulumi.String("server1"),
			TargetGroupName:   pulumi.String("targetGroup1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
