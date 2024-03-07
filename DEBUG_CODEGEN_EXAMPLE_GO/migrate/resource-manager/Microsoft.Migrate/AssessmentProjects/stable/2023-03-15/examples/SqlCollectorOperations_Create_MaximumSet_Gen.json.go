package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewSqlCollectorOperation(ctx, "sqlCollectorOperation", &migrate.SqlCollectorOperationArgs{
			AgentProperties: &migrate.CollectorAgentPropertiesBaseArgs{
				Id: pulumi.String("630da710-4d44-41f7-a189-72fe3db5502b-agent"),
				SpnDetails: &migrate.CollectorAgentSpnPropertiesBaseArgs{
					ApplicationId: pulumi.String("db9c4c3d-477c-4d5a-817b-318276713565"),
					Audience:      pulumi.String("db9c4c3d-477c-4d5a-817b-318276713565"),
					Authority:     pulumi.String("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
					ObjectId:      pulumi.String("e50236ad-ad07-47d4-af71-ed7b52d200d5"),
					TenantId:      pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
				},
			},
			CollectorName:     pulumi.String("fci-test0c1esqlsitecollector"),
			DiscoverySiteId:   pulumi.String("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.OffAzure/MasterSites/fci-ankit-test6065mastersite/SqlSites/fci-ankit-test6065sqlsites"),
			ProjectName:       pulumi.String("fci-test6904project"),
			ResourceGroupName: pulumi.String("rgmigrate"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
