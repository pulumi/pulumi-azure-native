package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewHypervCollectorsOperation(ctx, "hypervCollectorsOperation", &migrate.HypervCollectorsOperationArgs{
			AgentProperties: migrate.CollectorAgentPropertiesBaseResponse{
				Id:               pulumi.String("12f1d90f-b3fa-4926-8893-e56803a09af0"),
				LastHeartbeatUtc: pulumi.String("2022-07-07T14:25:35.708325Z"),
				SpnDetails: &migrate.CollectorAgentSpnPropertiesBaseArgs{
					ApplicationId: pulumi.String("e3bd6eaa-980b-40ae-a30e-2a5069ba097c"),
					Audience:      pulumi.String("e3bd6eaa-980b-40ae-a30e-2a5069ba097c"),
					Authority:     pulumi.String("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
					ObjectId:      pulumi.String("01b9f9e2-2d82-414c-adaa-09ce259b6b44"),
					TenantId:      pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
				},
				Version: pulumi.String("2.0.1993.19"),
			},
			DiscoverySiteId:     pulumi.String("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/HyperVSites/test-60527site"),
			HypervCollectorName: pulumi.String("test-697cecollector"),
			ProjectName:         pulumi.String("app18700project"),
			ProvisioningState:   pulumi.String("Succeeded"),
			ResourceGroupName:   pulumi.String("ayagrawRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
