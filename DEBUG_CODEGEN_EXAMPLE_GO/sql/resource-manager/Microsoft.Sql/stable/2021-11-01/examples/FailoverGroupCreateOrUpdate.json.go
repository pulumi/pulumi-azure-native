package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewFailoverGroup(ctx, "failoverGroup", &sql.FailoverGroupArgs{
			Databases: pulumi.StringArray{
				pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1"),
				pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-2"),
			},
			FailoverGroupName: pulumi.String("failover-group-test-3"),
			PartnerServers: []sql.PartnerInfoArgs{
				{
					Id: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
				},
			},
			ReadOnlyEndpoint: &sql.FailoverGroupReadOnlyEndpointArgs{
				FailoverPolicy: pulumi.String("Disabled"),
			},
			ReadWriteEndpoint: &sql.FailoverGroupReadWriteEndpointArgs{
				FailoverPolicy:                         pulumi.String("Automatic"),
				FailoverWithDataLossGracePeriodMinutes: pulumi.Int(480),
			},
			ResourceGroupName: pulumi.String("Default"),
			ServerName:        pulumi.String("failover-group-primary-server"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
