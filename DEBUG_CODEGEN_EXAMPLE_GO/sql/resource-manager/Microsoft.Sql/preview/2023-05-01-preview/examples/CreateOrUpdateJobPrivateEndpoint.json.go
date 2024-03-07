package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewJobPrivateEndpoint(ctx, "jobPrivateEndpoint", &sql.JobPrivateEndpointArgs{
			JobAgentName:                pulumi.String("agent1"),
			PrivateEndpointName:         pulumi.String("endpoint1"),
			ResourceGroupName:           pulumi.String("group1"),
			ServerName:                  pulumi.String("server1"),
			TargetServerAzureResourceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/targetserver1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
