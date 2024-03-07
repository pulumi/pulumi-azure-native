package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewWebAppCollectorOperation(ctx, "webAppCollectorOperation", &migrate.WebAppCollectorOperationArgs{
			AgentProperties: &migrate.CollectorAgentPropertiesBaseArgs{
				Id:               pulumi.String("fed93df5-b787-4e3f-a764-e3d2b9101a59-agent"),
				LastHeartbeatUtc: pulumi.String("2023-11-03T05:43:02.078Z"),
				SpnDetails: &migrate.CollectorAgentSpnPropertiesBaseArgs{
					ApplicationId: pulumi.String("2f70d5e8-7adc-4c64-910a-7031079efc6e"),
					Audience:      pulumi.String("2f70d5e8-7adc-4c64-910a-7031079efc6e"),
					Authority:     pulumi.String("https://login.microsoftonline.com/72f988bf-86f1-41af-91ab-2d7cd011db47"),
					ObjectId:      pulumi.String("2f70d5e8-7adc-4c64-910a-7031079efc6e"),
					TenantId:      pulumi.String("2f70d5e8-7adc-4c64-910a-7031079efc6e"),
				},
			},
			CollectorName:     pulumi.String("collector1"),
			DiscoverySiteId:   pulumi.String("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sumukk-ccy-bcs/providers/Microsoft.OffAzure/MasterSites/sumukk-ccy-bcs9880mastersite/WebAppSites/sumukk-ccy-bcs9880webappsites"),
			ProjectName:       pulumi.String("sumukk-ccy-bcs4557project"),
			ResourceGroupName: pulumi.String("rgopenapi"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
