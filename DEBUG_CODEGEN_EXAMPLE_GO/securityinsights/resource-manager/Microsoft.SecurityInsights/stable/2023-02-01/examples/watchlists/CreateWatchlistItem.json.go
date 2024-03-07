package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := securityinsights.NewWatchlistItem(ctx, "watchlistItem", &securityinsights.WatchlistItemArgs{
ItemsKeyValue: pulumi.Any{
Business tier: "10.0.2.0/24",
Data tier: "10.0.2.0/24",
Gateway subnet: "10.0.255.224/27",
Private DMZ in: "10.0.0.0/27",
Public DMZ out: "10.0.0.96/27",
Web Tier: "10.0.1.0/24",
},
ResourceGroupName: pulumi.String("myRg"),
WatchlistAlias: pulumi.String("highValueAsset"),
WatchlistItemId: pulumi.String("82ba292c-dc97-4dfc-969d-d4dd9e666842"),
WorkspaceName: pulumi.String("myWorkspace"),
})
if err != nil {
return err
}
return nil
})
}
