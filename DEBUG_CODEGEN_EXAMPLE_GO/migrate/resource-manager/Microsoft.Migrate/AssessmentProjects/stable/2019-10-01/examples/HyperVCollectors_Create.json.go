package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewHyperVCollector(ctx, "hyperVCollector", &migrate.HyperVCollectorArgs{
			ETag:                pulumi.String("\"00000981-0000-0300-0000-5d74cd5f0000\""),
			HyperVCollectorName: pulumi.String("migrateprojectce73collector"),
			ProjectName:         pulumi.String("migrateprojectce73project"),
			Properties: &migrate.CollectorPropertiesArgs{
				AgentProperties: &migrate.CollectorAgentPropertiesArgs{
					SpnDetails: &migrate.CollectorBodyAgentSpnPropertiesArgs{
						ApplicationId: pulumi.String("827f1053-44dc-439f-b832-05416dcce12b"),
						Audience:      pulumi.String("https://72f988bf-86f1-41af-91ab-2d7cd011db47/migrateprojectce73agentauthaadapp"),
						Authority:     pulumi.String("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
						ObjectId:      pulumi.String("be75098e-c0fc-4ac4-98c7-282ebbcf8370"),
						TenantId:      pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
					},
				},
				DiscoverySiteId: pulumi.String("/subscriptions/8c3c936a-c09b-4de3-830b-3f5f244d72e9/resourceGroups/ContosoITHyperV/providers/Microsoft.OffAzure/HyperVSites/migrateprojectce73site"),
			},
			ResourceGroupName: pulumi.String("contosoithyperv"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
