package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewJobStep(ctx, "jobStep", &sql.JobStepArgs{
			Action: &sql.JobStepActionArgs{
				Value: pulumi.String("select 1"),
			},
			Credential:        pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0"),
			JobAgentName:      pulumi.String("agent1"),
			JobName:           pulumi.String("job1"),
			ResourceGroupName: pulumi.String("group1"),
			ServerName:        pulumi.String("server1"),
			StepName:          pulumi.String("step1"),
			TargetGroup:       pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
