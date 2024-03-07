package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewJob(ctx, "job", &sql.JobArgs{
			JobAgentName:      pulumi.String("agent1"),
			JobName:           pulumi.String("job1"),
			ResourceGroupName: pulumi.String("group1"),
			ServerName:        pulumi.String("server1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
