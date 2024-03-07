package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := migrate.NewServerCollector(ctx, "serverCollector", &migrate.ServerCollectorArgs{
ETag: pulumi.String("\"00000606-0000-0d00-0000-605999bf0000\""),
ProjectName: pulumi.String("app11141project"),
Properties: migrate.CollectorPropertiesResponse{
AgentProperties: interface{}{
SpnDetails: &migrate.CollectorBodyAgentSpnPropertiesArgs{
ApplicationId: pulumi.String("ad9f701a-cc08-4421-b51f-b5762d58e9ba"),
Audience: pulumi.String("https://72f988bf-86f1-41af-91ab-2d7cd011db47/app23df4authandaccessaadapp"),
Authority: pulumi.String("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
ObjectId: pulumi.String("b4975e42-9248-4a36-b99f-37eca377ea00"),
TenantId: pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
},
},
DiscoverySiteId: pulumi.String("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/pajindTest/providers/Microsoft.OffAzure/ServerSites/app21141site"),
},
ResourceGroupName: pulumi.String("pajindtest"),
ServerCollectorName: pulumi.String("app23df4collector"),
})
if err != nil {
return err
}
return nil
})
}
