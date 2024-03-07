package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewJob(ctx, "job", &sql.JobArgs{
			Description:       pulumi.String("my favourite job"),
			JobAgentName:      pulumi.String("agent1"),
			JobName:           pulumi.String("job1"),
			ResourceGroupName: pulumi.String("group1"),
			Schedule: &sql.JobScheduleArgs{
				Enabled:   pulumi.Bool(true),
				EndTime:   pulumi.String("2015-09-24T23:59:59Z"),
				Interval:  pulumi.String("PT5M"),
				StartTime: pulumi.String("2015-09-24T18:30:01Z"),
				Type:      sql.JobScheduleTypeRecurring,
			},
			ServerName: pulumi.String("server1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
