package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewJobTargetGroup(ctx, "jobTargetGroup", &sql.JobTargetGroupArgs{
			JobAgentName: pulumi.String("agent1"),
			Members: []sql.JobTargetArgs{
				{
					DatabaseName:   pulumi.String("database1"),
					MembershipType: sql.JobTargetGroupMembershipTypeExclude,
					ServerName:     pulumi.String("server1"),
					Type:           pulumi.String("SqlDatabase"),
				},
				{
					MembershipType:    sql.JobTargetGroupMembershipTypeInclude,
					RefreshCredential: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/testCredential"),
					ServerName:        pulumi.String("server1"),
					Type:              pulumi.String("SqlServer"),
				},
				{
					ElasticPoolName:   pulumi.String("pool1"),
					MembershipType:    sql.JobTargetGroupMembershipTypeInclude,
					RefreshCredential: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/testCredential"),
					ServerName:        pulumi.String("server2"),
					Type:              pulumi.String("SqlElasticPool"),
				},
				{
					MembershipType:    sql.JobTargetGroupMembershipTypeInclude,
					RefreshCredential: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/testCredential"),
					ServerName:        pulumi.String("server3"),
					ShardMapName:      pulumi.String("shardMap1"),
					Type:              pulumi.String("SqlShardMap"),
				},
			},
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
