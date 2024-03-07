package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewVmwareCollectorsOperation(ctx, "vmwareCollectorsOperation", &migrate.VmwareCollectorsOperationArgs{
			AgentProperties: &migrate.CollectorAgentPropertiesBaseArgs{
				Id:               pulumi.String("fe243486-3318-41fa-aaba-c48b5df75308"),
				LastHeartbeatUtc: pulumi.String("2022-03-29T12:10:08.9167289Z"),
				SpnDetails: &migrate.CollectorAgentSpnPropertiesBaseArgs{
					ApplicationId: pulumi.String("82b3e452-c0e8-4662-8347-58282925ae84"),
					Audience:      pulumi.String("82b3e452-c0e8-4662-8347-58282925ae84"),
					Authority:     pulumi.String("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
					ObjectId:      pulumi.String("3fc89111-1405-4938-9214-37aa4739401d"),
					TenantId:      pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
				},
				Version: pulumi.String("1.0.8.383"),
			},
			DiscoverySiteId:     pulumi.String("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/VMwareSites/Vmware2744site"),
			ProjectName:         pulumi.String("app18700project"),
			ProvisioningState:   pulumi.String("Succeeded"),
			ResourceGroupName:   pulumi.String("ayagrawRG"),
			VmWareCollectorName: pulumi.String("Vmware2258collector"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
