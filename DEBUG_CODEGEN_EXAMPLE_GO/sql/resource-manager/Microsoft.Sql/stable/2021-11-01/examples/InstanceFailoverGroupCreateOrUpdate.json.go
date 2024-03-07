package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewInstanceFailoverGroup(ctx, "instanceFailoverGroup", &sql.InstanceFailoverGroupArgs{
			FailoverGroupName: pulumi.String("failover-group-test-3"),
			LocationName:      pulumi.String("Japan East"),
			ManagedInstancePairs: []sql.ManagedInstancePairInfoArgs{
				{
					PartnerManagedInstanceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-secondary-mngdInstance"),
					PrimaryManagedInstanceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-primary-mngdInstance"),
				},
			},
			PartnerRegions: []sql.PartnerRegionInfoArgs{
				{
					Location: pulumi.String("Japan West"),
				},
			},
			ReadOnlyEndpoint: &sql.InstanceFailoverGroupReadOnlyEndpointArgs{
				FailoverPolicy: pulumi.String("Disabled"),
			},
			ReadWriteEndpoint: &sql.InstanceFailoverGroupReadWriteEndpointArgs{
				FailoverPolicy:                         pulumi.String("Automatic"),
				FailoverWithDataLossGracePeriodMinutes: pulumi.Int(480),
			},
			ResourceGroupName: pulumi.String("Default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
