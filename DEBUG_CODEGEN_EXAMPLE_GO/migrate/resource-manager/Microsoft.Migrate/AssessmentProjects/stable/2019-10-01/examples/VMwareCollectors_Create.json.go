package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := migrate.NewVMwareCollector(ctx, "vMwareCollector", &migrate.VMwareCollectorArgs{
ETag: pulumi.String("\"01003d32-0000-0d00-0000-5d74d2e50000\""),
ProjectName: pulumi.String("abgoyalWEselfhostb72bproject"),
Properties: migrate.CollectorPropertiesResponse{
AgentProperties: interface{}{
SpnDetails: &migrate.CollectorBodyAgentSpnPropertiesArgs{
ApplicationId: pulumi.String("fc717575-8173-4b21-92a5-658b655e613e"),
Audience: pulumi.String("https://72f988bf-86f1-41af-91ab-2d7cd011db47/PortalvCenterbc2fagentauthaadapp"),
Authority: pulumi.String("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
ObjectId: pulumi.String("29d94f38-db94-4980-aec0-0cfd55ab1cd0"),
TenantId: pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
},
},
DiscoverySiteId: pulumi.String("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.OffAzure/VMwareSites/PortalvCenterbc2fsite"),
},
ResourceGroupName: pulumi.String("abgoyal-westEurope"),
VmWareCollectorName: pulumi.String("PortalvCenterbc2fcollector"),
})
if err != nil {
return err
}
return nil
})
}
