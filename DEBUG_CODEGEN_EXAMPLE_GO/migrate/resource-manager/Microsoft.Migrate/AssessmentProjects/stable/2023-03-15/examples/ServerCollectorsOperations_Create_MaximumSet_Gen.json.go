package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewServerCollectorsOperation(ctx, "serverCollectorsOperation", &migrate.ServerCollectorsOperationArgs{
			AgentProperties: migrate.CollectorAgentPropertiesBaseResponse{
				Id: pulumi.String("498e4965-bbb1-47c2-8613-345baff9c509"),
				SpnDetails: &migrate.CollectorAgentSpnPropertiesBaseArgs{
					ApplicationId: pulumi.String("65153d2f-9afb-44e8-b3ca-1369150b7354"),
					Audience:      pulumi.String("65153d2f-9afb-44e8-b3ca-1369150b7354"),
					Authority:     pulumi.String("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
					ObjectId:      pulumi.String("ddde6f96-87c8-420b-9d4d-f16a5090519e"),
					TenantId:      pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
				},
			},
			DiscoverySiteId:     pulumi.String("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/ServerSites/walter7155site"),
			ProjectName:         pulumi.String("app18700project"),
			ProvisioningState:   pulumi.String("Succeeded"),
			ResourceGroupName:   pulumi.String("ayagrawRG"),
			ServerCollectorName: pulumi.String("walter389fcollector"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
