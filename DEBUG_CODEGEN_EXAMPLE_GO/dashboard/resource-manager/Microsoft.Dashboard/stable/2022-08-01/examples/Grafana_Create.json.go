package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dashboard/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := dashboard.NewGrafana(ctx, "grafana", &dashboard.GrafanaArgs{
Identity: &dashboard.ManagedServiceIdentityArgs{
Type: pulumi.String("SystemAssigned"),
},
Location: pulumi.String("West US"),
Properties: dashboard.ManagedGrafanaPropertiesResponse{
ApiKey: pulumi.String("Enabled"),
DeterministicOutboundIP: pulumi.String("Enabled"),
GrafanaIntegrations: interface{}{
AzureMonitorWorkspaceIntegrations: dashboard.AzureMonitorWorkspaceIntegrationArray{
&dashboard.AzureMonitorWorkspaceIntegrationArgs{
AzureMonitorWorkspaceResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"),
},
},
},
PublicNetworkAccess: pulumi.String("Enabled"),
ZoneRedundancy: pulumi.String("Enabled"),
},
ResourceGroupName: pulumi.String("myResourceGroup"),
Sku: &dashboard.ResourceSkuArgs{
Name: pulumi.String("Standard"),
},
Tags: pulumi.StringMap{
"Environment": pulumi.String("Dev"),
},
WorkspaceName: pulumi.String("myWorkspace"),
})
if err != nil {
return err
}
return nil
})
}
