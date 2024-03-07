package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewJobStep(ctx, "jobStep", &sql.JobStepArgs{
			Action: &sql.JobStepActionArgs{
				Source: pulumi.String("Inline"),
				Type:   pulumi.String("TSql"),
				Value:  pulumi.String("select 2"),
			},
			Credential: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred1"),
			ExecutionOptions: &sql.JobStepExecutionOptionsArgs{
				InitialRetryIntervalSeconds:    pulumi.Int(11),
				MaximumRetryIntervalSeconds:    pulumi.Int(222),
				RetryAttempts:                  pulumi.Int(42),
				RetryIntervalBackoffMultiplier: pulumi.Float64(3),
				TimeoutSeconds:                 pulumi.Int(1234),
			},
			JobAgentName: pulumi.String("agent1"),
			JobName:      pulumi.String("job1"),
			Output: &sql.JobStepOutputTypeArgs{
				Credential:        pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0"),
				DatabaseName:      pulumi.String("database3"),
				ResourceGroupName: pulumi.String("group3"),
				SchemaName:        pulumi.String("myschema1234"),
				ServerName:        pulumi.String("server3"),
				SubscriptionId:    pulumi.String("3501b905-a848-4b5d-96e8-b253f62d735a"),
				TableName:         pulumi.String("mytable5678"),
				Type:              pulumi.String("SqlDatabase"),
			},
			ResourceGroupName: pulumi.String("group1"),
			ServerName:        pulumi.String("server1"),
			StepId:            pulumi.Int(1),
			StepName:          pulumi.String("step1"),
			TargetGroup:       pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
